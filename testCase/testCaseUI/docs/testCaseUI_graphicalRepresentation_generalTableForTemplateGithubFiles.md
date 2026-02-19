# testCaseUI_graphicalRepresentation_generalTableForTemplateGithubFiles.go

## File Overview
- Path: `testCase/testCaseUI/testCaseUI_graphicalRepresentation_generalTableForTemplateGithubFiles.go`
- Package: `testCaseUI`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `9`
- Imports: `10`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `MinSize`
- `MouseIn`
- `MouseMoved`
- `MouseOut`
- `Tapped`
- `TappedSecondary`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/importFilesFromGitHub`
- `FenixTesterGui/testCase/testCaseModel`
- `FenixTesterGui/testCase/testCaseUI/templateViewer`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/driver/desktop`
- `fyne.io/fyne/v2/theme`
- `fyne.io/fyne/v2/widget`
- `github.com/sirupsen/logrus`
- `time`

## Declared Types
- `CustomTemplateTable`
- `clickableLabel`

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### MinSize (method on `*CustomTemplateTable`)
- Signature: `func (*CustomTemplateTable) MinSize() fyne.Size`
- Exported: `true`
- Control-flow features: `if`
- Doc: MinSize customizes the minimum size of the CustomTemplateTable.
- Internal calls: `float32`
- External calls: `fyne.MeasureText`, `fyne.NewSize`, `t.Length`, `t.Size`, `theme.TextSize`

### MouseIn (method on `*clickableLabel`)
- Signature: `func (*clickableLabel) MouseIn(*desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`

### MouseMoved (method on `*clickableLabel`)
- Signature: `func (*clickableLabel) MouseMoved(*desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`

### MouseOut (method on `*clickableLabel`)
- Signature: `func (*clickableLabel) MouseOut()`
- Exported: `true`
- Control-flow features: `none detected`

### Tapped (method on `*clickableLabel`)
- Signature: `func (*clickableLabel) Tapped(e *fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `if`
- External calls: `l.onDoubleTap`, `time.Now`, `time.Since`

### TappedSecondary (method on `*clickableLabel`)
- Signature: `func (*clickableLabel) TappedSecondary(*fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `none detected`

### generateTemplateFilesTable
- Signature: `func generateTemplateFilesTable(testCaseUuid string, filesAreClickable bool, testCasesUiCanvasObject *TestCasesUiModelStruct) *CustomTemplateTable`
- Exported: `false`
- Control-flow features: `if, switch`
- Doc: Create the UI-list that holds the selected files
- Internal calls: `newClickableLabel`
- External calls: `clickable.SetText`, `nonClickable.SetText`, `templateFilesTable.ExtendBaseWidget`, `templateViewer.InitiateTemplateViewer`

### newClickableLabel
- Signature: `func newClickableLabel(text string, onDoubleTap func(), tempIsClickable bool) *clickableLabel`
- Exported: `false`
- Control-flow features: `none detected`
- External calls: `l.ExtendBaseWidget`, `time.Now`

### updateColumnAndRowSizes (method on `*CustomTemplateTable`)
- Signature: `func (*CustomTemplateTable) updateColumnAndRowSizes(testCaseUuid string, testCasesUiCanvasObject *TestCasesUiModelStruct, checkIfTemplatesAreChangedButton *widget.Button, viewTemplateButton *widget.Button)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Update size of columns and Rows for the Table
- Internal calls: `float32`
- External calls: `checkIfTemplatesAreChangedButton.Disable`, `checkIfTemplatesAreChangedButton.Enable`, `fyne.MeasureText`, `t.Refresh`, `t.SetColumnWidth`, `t.SetRowHeight`, `theme.Padding`, `theme.TextSize`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
