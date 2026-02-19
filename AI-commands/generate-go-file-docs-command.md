# Command: Generate high-quality docs for every Go file

Use the command below to generate documentation artifacts for each `.go` file in the repository using source analysis (Go AST).

It will:
- Find all `.go` files in the repo
- Create a `docs` subfolder next to each `.go` file
- Generate three documentation artifacts per `.go` file:
  - `docs/<file>_activity.puml`
  - `docs/<file>_sequence.puml`
  - `docs/<file>.md`
- Fill docs with extracted package/function/import/comment/control-flow details

```bash
set -euo pipefail

TMP_GO="$(mktemp /tmp/go-docgen-XXXXXX.go)"
trap 'rm -f "$TMP_GO"' EXIT

cat > "$TMP_GO" <<'EOGO'
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type FuncInfo struct {
	Name          string
	Receiver      string
	Signature     string
	Doc           string
	Exported      bool
	CallsInternal []string
	CallsExternal []string
	HasIf         bool
	HasFor        bool
	HasSwitch     bool
	HasSelect     bool
	HasGo         bool
	HasDefer      bool
	ReturnsError  bool
}

type FileInfo struct {
	Path      string
	Base      string
	Dir       string
	Package   string
	FileDoc   string
	Imports   []string
	Types     []string
	Constants []string
	Variables []string
	Funcs     []FuncInfo
}

func main() {
	files, err := readPathsFromStdin()
	if err != nil {
		fmt.Fprintf(os.Stderr, "read paths: %v\n", err)
		os.Exit(1)
	}

	failed := 0
	for _, p := range files {
		if err := processFile(p); err != nil {
			fmt.Fprintf(os.Stderr, "process %s: %v\n", p, err)
			failed++
		}
	}

	if failed > 0 {
		fmt.Fprintf(os.Stderr, "documentation generation failed for %d file(s)\n", failed)
		os.Exit(1)
	}
}

func readPathsFromStdin() ([]string, error) {
	var out []string
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		p := strings.TrimSpace(s.Text())
		if p == "" {
			continue
		}
		out = append(out, p)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	sort.Strings(out)
	return out, nil
}

func processFile(path string) error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	info := extractFileInfo(fset, f, path)
	docsDir := filepath.Join(info.Dir, "docs")
	if err := os.MkdirAll(docsDir, 0o755); err != nil {
		return err
	}

	activityPath := filepath.Join(docsDir, info.Base+"_activity.puml")
	sequencePath := filepath.Join(docsDir, info.Base+"_sequence.puml")
	markdownPath := filepath.Join(docsDir, info.Base+".md")

	if err := os.WriteFile(activityPath, []byte(buildActivityDiagram(info)), 0o644); err != nil {
		return err
	}
	if err := os.WriteFile(sequencePath, []byte(buildSequenceDiagram(info)), 0o644); err != nil {
		return err
	}
	if err := os.WriteFile(markdownPath, []byte(buildMarkdown(info)), 0o644); err != nil {
		return err
	}

	return nil
}

func extractFileInfo(fset *token.FileSet, f *ast.File, path string) FileInfo {
	base := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	dir := filepath.Dir(path)

	importsSet := map[string]struct{}{}
	typesSet := map[string]struct{}{}
	constSet := map[string]struct{}{}
	varSet := map[string]struct{}{}
	funcs := make([]FuncInfo, 0, 16)

	for _, im := range f.Imports {
		if unq, err := strconv.Unquote(im.Path.Value); err == nil {
			importsSet[unq] = struct{}{}
		}
	}

	for _, d := range f.Decls {
		switch decl := d.(type) {
		case *ast.GenDecl:
			switch decl.Tok {
			case token.TYPE:
				for _, spec := range decl.Specs {
					ts := spec.(*ast.TypeSpec)
					typesSet[ts.Name.Name] = struct{}{}
				}
			case token.CONST:
				for _, spec := range decl.Specs {
					vs := spec.(*ast.ValueSpec)
					for _, n := range vs.Names {
						constSet[n.Name] = struct{}{}
					}
				}
			case token.VAR:
				for _, spec := range decl.Specs {
					vs := spec.(*ast.ValueSpec)
					for _, n := range vs.Names {
						varSet[n.Name] = struct{}{}
					}
				}
			}
		case *ast.FuncDecl:
			funcs = append(funcs, extractFuncInfo(fset, decl))
		}
	}

	imports := setToSortedSlice(importsSet)
	types := setToSortedSlice(typesSet)
	constants := setToSortedSlice(constSet)
	variables := setToSortedSlice(varSet)
	sort.Slice(funcs, func(i, j int) bool { return funcs[i].Name < funcs[j].Name })

	fileDoc := ""
	if f.Doc != nil {
		fileDoc = cleanDoc(f.Doc.Text())
	}

	return FileInfo{
		Path:      filepath.ToSlash(path),
		Base:      base,
		Dir:       dir,
		Package:   f.Name.Name,
		FileDoc:   fileDoc,
		Imports:   imports,
		Types:     types,
		Constants: constants,
		Variables: variables,
		Funcs:     funcs,
	}
}

func extractFuncInfo(fset *token.FileSet, d *ast.FuncDecl) FuncInfo {
	fi := FuncInfo{
		Name:      d.Name.Name,
		Doc:       "",
		Exported:  d.Name.IsExported(),
		Signature: buildFuncSignature(fset, d),
	}

	if d.Doc != nil {
		fi.Doc = cleanDoc(d.Doc.Text())
	}

	if d.Recv != nil && len(d.Recv.List) > 0 {
		fi.Receiver = exprString(fset, d.Recv.List[0].Type)
	}

	internalSet := map[string]struct{}{}
	externalSet := map[string]struct{}{}

	if d.Type.Results != nil {
		for _, r := range d.Type.Results.List {
			if isErrorType(r.Type) {
				fi.ReturnsError = true
				break
			}
		}
	}

	if d.Body != nil {
		ast.Inspect(d.Body, func(n ast.Node) bool {
			switch x := n.(type) {
			case *ast.IfStmt:
				fi.HasIf = true
			case *ast.ForStmt, *ast.RangeStmt:
				fi.HasFor = true
			case *ast.SwitchStmt, *ast.TypeSwitchStmt:
				fi.HasSwitch = true
			case *ast.SelectStmt:
				fi.HasSelect = true
			case *ast.GoStmt:
				fi.HasGo = true
			case *ast.DeferStmt:
				fi.HasDefer = true
			case *ast.CallExpr:
				switch fun := x.Fun.(type) {
				case *ast.Ident:
					if !isBuiltin(fun.Name) {
						internalSet[fun.Name] = struct{}{}
					}
				case *ast.SelectorExpr:
					if pkg, sel, ok := simpleSelector(fun); ok {
						externalSet[pkg+"."+sel] = struct{}{}
					}
				}
			}
			return true
		})
	}

	fi.CallsInternal = setToSortedSlice(internalSet)
	fi.CallsExternal = setToSortedSlice(externalSet)
	return fi
}

func simpleSelector(s *ast.SelectorExpr) (pkg, sel string, ok bool) {
	id, ok := s.X.(*ast.Ident)
	if !ok {
		return "", "", false
	}
	if id.Name == "" || s.Sel == nil || s.Sel.Name == "" {
		return "", "", false
	}
	return id.Name, s.Sel.Name, true
}

func buildFuncSignature(fset *token.FileSet, d *ast.FuncDecl) string {
	params := ""
	results := ""

	if d.Type.Params != nil {
		params = fieldListString(fset, d.Type.Params)
	}
	if d.Type.Results != nil {
		results = fieldListString(fset, d.Type.Results)
	}

	recv := ""
	if d.Recv != nil && len(d.Recv.List) > 0 {
		recv = "(" + exprString(fset, d.Recv.List[0].Type) + ") "
	}

	sig := "func " + recv + d.Name.Name + "(" + params + ")"
	if results != "" {
		if strings.Contains(results, ",") {
			sig += " (" + results + ")"
		} else {
			sig += " " + results
		}
	}
	return sig
}

func fieldListString(fset *token.FileSet, fl *ast.FieldList) string {
	if fl == nil || len(fl.List) == 0 {
		return ""
	}
	parts := make([]string, 0, len(fl.List))
	for _, f := range fl.List {
		t := exprString(fset, f.Type)
		if len(f.Names) == 0 {
			parts = append(parts, t)
			continue
		}
		names := make([]string, 0, len(f.Names))
		for _, n := range f.Names {
			names = append(names, n.Name)
		}
		parts = append(parts, strings.Join(names, ", ")+" "+t)
	}
	return strings.Join(parts, ", ")
}

func exprString(fset *token.FileSet, e ast.Expr) string {
	if e == nil {
		return ""
	}
	var b bytes.Buffer
	_ = printer.Fprint(&b, fset, e)
	return b.String()
}

func isErrorType(e ast.Expr) bool {
	id, ok := e.(*ast.Ident)
	return ok && id.Name == "error"
}

func isBuiltin(name string) bool {
	_, ok := map[string]struct{}{
		"append": {}, "cap": {}, "close": {}, "complex": {}, "copy": {}, "delete": {},
		"imag": {}, "len": {}, "make": {}, "new": {}, "panic": {}, "print": {},
		"println": {}, "real": {}, "recover": {},
	}[name]
	return ok
}

func cleanDoc(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	lines := strings.Split(s, "\n")
	out := make([]string, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "//") {
			line = strings.TrimSpace(strings.TrimPrefix(line, "//"))
		}
		out = append(out, line)
		if len(out) == 2 {
			break
		}
	}
	if len(out) == 0 {
		return ""
	}
	text := strings.Join(out, " ")
	if len(text) > 240 {
		text = text[:240] + "..."
	}
	return text
}

func setToSortedSlice(m map[string]struct{}) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func pickEntryPoints(funcs []FuncInfo) []string {
	out := make([]string, 0, 8)
	seen := map[string]struct{}{}
	for _, f := range funcs {
		if f.Name == "main" || f.Name == "init" || f.Exported {
			if _, exists := seen[f.Name]; !exists {
				seen[f.Name] = struct{}{}
				out = append(out, f.Name)
			}
		}
	}
	sort.Strings(out)
	return out
}

func buildActivityDiagram(fi FileInfo) string {
	var b strings.Builder
	b.WriteString("@startuml\n")
	b.WriteString("title Activity Diagram - " + fi.Base + ".go\n")
	b.WriteString("start\n")
	b.WriteString(":Load " + fi.Path + ";\n")

	funcs := fi.Funcs
	if len(funcs) > 12 {
		funcs = funcs[:12]
	}

	if len(funcs) == 0 {
		b.WriteString(":No functions declared in this file;\n")
	} else {
		for i, f := range funcs {
			step := fmt.Sprintf("%d. %s", i+1, f.Name)
			if f.Receiver != "" {
				step += " [method on " + f.Receiver + "]"
			}
			b.WriteString(":" + step + ";\n")

			if f.HasIf || f.HasSwitch || f.ReturnsError {
				b.WriteString("if (branch or error path?) then (yes)\n")
				b.WriteString(":Handle alternative/error path;\n")
				b.WriteString("else (no)\n")
				b.WriteString(":Continue primary flow;\n")
				b.WriteString("endif\n")
			}
			if f.HasFor {
				b.WriteString(":Iterate loop/range logic;\n")
			}
			if f.HasDefer {
				b.WriteString(":Schedule deferred cleanup;\n")
			}
			if f.HasGo {
				b.WriteString(":Start goroutine/asynchronous work;\n")
			}
		}
	}

	b.WriteString("stop\n")
	b.WriteString("@enduml\n")
	return b.String()
}

func buildSequenceDiagram(fi FileInfo) string {
	var b strings.Builder
	b.WriteString("@startuml\n")
	b.WriteString("title Sequence Diagram - " + fi.Base + ".go\n")
	b.WriteString("actor Caller\n")
	b.WriteString("participant \"" + fi.Base + ".go\" as File\n")

	extSet := map[string]struct{}{}
	for _, f := range fi.Funcs {
		for _, c := range f.CallsExternal {
			left := strings.SplitN(c, ".", 2)[0]
			if left != "" {
				extSet[left] = struct{}{}
			}
		}
	}
	exts := setToSortedSlice(extSet)
	if len(exts) > 8 {
		exts = exts[:8]
	}
	for _, e := range exts {
		alias := strings.Map(func(r rune) rune {
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
				return r
			}
			return '_'
		}, e)
		b.WriteString("participant \"" + e + "\" as " + alias + "\n")
	}

	selected := make([]FuncInfo, 0, 10)
	for _, f := range fi.Funcs {
		if f.Exported {
			selected = append(selected, f)
		}
	}
	if len(selected) == 0 {
		selected = fi.Funcs
	}
	if len(selected) > 10 {
		selected = selected[:10]
	}

	for _, f := range selected {
		b.WriteString("Caller -> File: " + f.Name + "(...)\n")
		b.WriteString("activate File\n")
		external := f.CallsExternal
		if len(external) > 3 {
			external = external[:3]
		}
		for _, call := range external {
			parts := strings.SplitN(call, ".", 2)
			if len(parts) != 2 {
				continue
			}
			alias := strings.Map(func(r rune) rune {
				if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
					return r
				}
				return '_'
			}, parts[0])
			b.WriteString("File -> " + alias + ": " + parts[1] + "(...)\n")
		}
		b.WriteString("File --> Caller: return\n")
		b.WriteString("deactivate File\n")
	}

	if len(selected) == 0 {
		b.WriteString("Caller -> File: (no callable functions detected)\n")
		b.WriteString("File --> Caller: return\n")
	}

	b.WriteString("@enduml\n")
	return b.String()
}

func buildMarkdown(fi FileInfo) string {
	var b strings.Builder
	b.WriteString("# " + fi.Base + ".go\n\n")
	b.WriteString("## File Overview\n")
	b.WriteString("- Path: `" + fi.Path + "`\n")
	b.WriteString("- Package: `" + fi.Package + "`\n")
	b.WriteString(fmt.Sprintf("- Functions/Methods: `%d`\n", len(fi.Funcs)))
	b.WriteString(fmt.Sprintf("- Imports: `%d`\n\n", len(fi.Imports)))

	b.WriteString("## File Purpose\n")
	if fi.FileDoc != "" {
		b.WriteString(fi.FileDoc + "\n\n")
	} else {
		b.WriteString("No concise file-level comment detected. Purpose inferred from declarations below.\n\n")
	}

	entryPoints := pickEntryPoints(fi.Funcs)
	b.WriteString("## Potential Entry Points\n")
	if len(entryPoints) == 0 {
		b.WriteString("- None detected\n\n")
	} else {
		for _, ep := range entryPoints {
			b.WriteString("- `" + ep + "`\n")
		}
		b.WriteString("\n")
	}

	b.WriteString("## Imports\n")
	if len(fi.Imports) == 0 {
		b.WriteString("- None\n\n")
	} else {
		for _, im := range fi.Imports {
			b.WriteString("- `" + im + "`\n")
		}
		b.WriteString("\n")
	}

	b.WriteString("## Declared Types\n")
	if len(fi.Types) == 0 {
		b.WriteString("- None\n\n")
	} else {
		for _, t := range fi.Types {
			b.WriteString("- `" + t + "`\n")
		}
		b.WriteString("\n")
	}

	b.WriteString("## Declared Constants\n")
	if len(fi.Constants) == 0 {
		b.WriteString("- None\n\n")
	} else {
		for _, c := range fi.Constants {
			b.WriteString("- `" + c + "`\n")
		}
		b.WriteString("\n")
	}

	b.WriteString("## Declared Variables\n")
	if len(fi.Variables) == 0 {
		b.WriteString("- None\n\n")
	} else {
		for _, v := range fi.Variables {
			b.WriteString("- `" + v + "`\n")
		}
		b.WriteString("\n")
	}

	b.WriteString("## Functions and Methods\n")
	if len(fi.Funcs) == 0 {
		b.WriteString("No functions or methods declared in this file.\n\n")
	} else {
		for _, f := range fi.Funcs {
			name := f.Name
			if f.Receiver != "" {
				name += " (method on `" + f.Receiver + "`)"
			}
			b.WriteString("### " + name + "\n")
			b.WriteString("- Signature: `" + f.Signature + "`\n")
			b.WriteString("- Exported: `" + fmt.Sprintf("%t", f.Exported) + "`\n")

			flags := make([]string, 0, 8)
			if f.HasIf {
				flags = append(flags, "if")
			}
			if f.HasFor {
				flags = append(flags, "for/range")
			}
			if f.HasSwitch {
				flags = append(flags, "switch")
			}
			if f.HasSelect {
				flags = append(flags, "select")
			}
			if f.HasGo {
				flags = append(flags, "go")
			}
			if f.HasDefer {
				flags = append(flags, "defer")
			}
			if f.ReturnsError {
				flags = append(flags, "returns error")
			}
			if len(flags) == 0 {
				b.WriteString("- Control-flow features: `none detected`\n")
			} else {
				b.WriteString("- Control-flow features: `" + strings.Join(flags, ", ") + "`\n")
			}

			if f.Doc != "" {
				b.WriteString("- Doc: " + f.Doc + "\n")
			}

			if len(f.CallsInternal) > 0 {
				calls := f.CallsInternal
				if len(calls) > 8 {
					calls = calls[:8]
				}
				b.WriteString("- Internal calls: `" + strings.Join(calls, "`, `") + "`\n")
			}
			if len(f.CallsExternal) > 0 {
				calls := f.CallsExternal
				if len(calls) > 8 {
					calls = calls[:8]
				}
				b.WriteString("- Selector calls: `" + strings.Join(calls, "`, `") + "`\n")
			}
			b.WriteString("\n")
		}
	}

	b.WriteString("## Behavioral Summary\n")
	b.WriteString("This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.\n")
	return b.String()
}
EOGO

if command -v rg >/dev/null 2>&1; then
  rg --files -g '**/*.go' | go run "$TMP_GO"
else
  find . -type f -name '*.go' -print | sed 's#^\./##' | go run "$TMP_GO"
fi

```

## Important
- This command updates documentation files in-place.
- Run it from repository root.
- Re-run whenever Go files change significantly.
