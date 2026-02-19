# importFilesFromGithub_FilteredFilesList.go

## File Overview
- Path: `importFilesFromGitHub/importFilesFromGithub_FilteredFilesList.go`
- Package: `importFilesFromGitHub`
- Functions/Methods: `8`
- Imports: `8`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `MouseIn`
- `MouseMoved`
- `MouseOut`
- `Tapped`
- `TappedSecondary`

## Imports
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/dialog`
- `fyne.io/fyne/v2/driver/desktop`
- `fyne.io/fyne/v2/widget`
- `log`
- `regexp`
- `strings`
- `time`

## Declared Types
- `customFilteredLabel`

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### generateFilteredList (method on `*ImportFilesFromGitHubStruct`)
- Signature: `func (*ImportFilesFromGitHubStruct) generateFilteredList(parentWindow fyne.Window)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Selector calls: `widget.NewList`, `importFilesFromGitHubObject.newCustomFilteredLabel`, `importFilesFromGitHubObject.getFileListFromGitHub`, `importFilesFromGitHubObject.filterFileListFromGitHub`, `strings.Split`, `importFilesFromGitHubObject.UpdateSelectedFilesTable`, `dialog.ShowInformation`, `label.Refresh`

### filterFileListFromGitHub (method on `*ImportFilesFromGitHubStruct`)
- Signature: `func (*ImportFilesFromGitHubStruct) filterFileListFromGitHub()`
- Exported: `false`
- Control-flow features: `if, for/range`
- Selector calls: `strings.ReplaceAll`, `regexp.Compile`, `log.Fatalln`, `combinedRegex.MatchString`

### newCustomFilteredLabel (method on `*ImportFilesFromGitHubStruct`)
- Signature: `func (*ImportFilesFromGitHubStruct) newCustomFilteredLabel(text string, onDoubleTap func()) *customFilteredLabel`
- Exported: `false`
- Control-flow features: `none detected`
- Selector calls: `l.ExtendBaseWidget`

### Tapped (method on `*customFilteredLabel`)
- Signature: `func (*customFilteredLabel) Tapped(e *fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `if`
- Selector calls: `time.Now`, `now.Sub`, `l.onDoubleTap`

### TappedSecondary (method on `*customFilteredLabel`)
- Signature: `func (*customFilteredLabel) TappedSecondary(*fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `none detected`

### MouseIn (method on `*customFilteredLabel`)
- Signature: `func (*customFilteredLabel) MouseIn(*desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `l.Refresh`

### MouseMoved (method on `*customFilteredLabel`)
- Signature: `func (*customFilteredLabel) MouseMoved(*desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`

### MouseOut (method on `*customFilteredLabel`)
- Signature: `func (*customFilteredLabel) MouseOut()`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `l.Refresh`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
