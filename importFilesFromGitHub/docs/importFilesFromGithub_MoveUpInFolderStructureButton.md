# importFilesFromGithub_MoveUpInFolderStructureButton.go

## File Overview
- Path: `importFilesFromGitHub/importFilesFromGithub_MoveUpInFolderStructureButton.go`
- Package: `importFilesFromGitHub`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `2`
- Imports: `4`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `fmt`
- `fyne.io/fyne/v2/theme`
- `fyne.io/fyne/v2/widget`
- `strings`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### generateMoveUpInFolderStructureButton (method on `*ImportFilesFromGitHubStruct`)
- Signature: `func (*ImportFilesFromGitHubStruct) generateMoveUpInFolderStructureButton()`
- Exported: `false`
- Control-flow features: `if`
- Doc: Generate the Button that moves upwards in the folder structure in GitHub
- External calls: `importFilesFromGitHubObject.filterFileListFromGitHub`, `importFilesFromGitHubObject.getFileListFromGitHub`, `importFilesFromGitHubObject.moveUpInPath`, `strings.Split`, `theme.NavigateBackIcon`, `widget.NewButtonWithIcon`

### moveUpInPath (method on `*ImportFilesFromGitHubStruct`)
- Signature: `func (*ImportFilesFromGitHubStruct) moveUpInPath(currentPath string) (string, error)`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Move one step in the folder structure
- External calls: `fmt.Errorf`, `strings.Join`, `strings.Split`, `strings.TrimRight`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
