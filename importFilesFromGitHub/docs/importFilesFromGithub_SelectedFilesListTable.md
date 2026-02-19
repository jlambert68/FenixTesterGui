# importFilesFromGithub_SelectedFilesListTable.go

## File Overview
- Path: `importFilesFromGitHub/importFilesFromGithub_SelectedFilesListTable.go`
- Package: `importFilesFromGitHub`
- Functions/Methods: `14`
- Imports: `5`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `MouseIn`
- `MouseMoved`
- `MouseOut`
- `Tapped`
- `TappedSecondary`
- `UpdateSelectedFilesTable`

## Imports
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/driver/desktop`
- `fyne.io/fyne/v2/theme`
- `fyne.io/fyne/v2/widget`
- `time`

## Declared Types
- `clickableLabel`
- `customLabel`

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### MouseIn (method on `*customLabel`)
- Signature: `func (*customLabel) MouseIn(*desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`

### MouseIn (method on `*clickableLabel`)
- Signature: `func (*clickableLabel) MouseIn(*desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `if`
- Selector calls: `l.Refresh`

### MouseMoved (method on `*customLabel`)
- Signature: `func (*customLabel) MouseMoved(*desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`

### MouseMoved (method on `*clickableLabel`)
- Signature: `func (*clickableLabel) MouseMoved(*desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`

### MouseOut (method on `*customLabel`)
- Signature: `func (*customLabel) MouseOut()`
- Exported: `true`
- Control-flow features: `none detected`

### MouseOut (method on `*clickableLabel`)
- Signature: `func (*clickableLabel) MouseOut()`
- Exported: `true`
- Control-flow features: `if`
- Selector calls: `l.Refresh`

### Tapped (method on `*customLabel`)
- Signature: `func (*customLabel) Tapped(e *fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `if`
- Selector calls: `l.onDoubleTap`, `now.Sub`, `time.Now`

### Tapped (method on `*clickableLabel`)
- Signature: `func (*clickableLabel) Tapped(e *fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `if`
- Selector calls: `l.onDoubleTap`, `time.Now`, `time.Since`

### TappedSecondary (method on `*customLabel`)
- Signature: `func (*customLabel) TappedSecondary(*fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `none detected`

### TappedSecondary (method on `*clickableLabel`)
- Signature: `func (*clickableLabel) TappedSecondary(*fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `none detected`

### UpdateSelectedFilesTable (method on `*ImportFilesFromGitHubStruct`)
- Signature: `func (*ImportFilesFromGitHubStruct) UpdateSelectedFilesTable()`
- Exported: `true`
- Control-flow features: `if, for/range, switch`
- Internal calls: `float32`
- Selector calls: `clickable.SetText`, `fyne.MeasureText`, `importFilesFromGitHubObject.UpdateSelectedFilesTable`, `importFilesFromGitHubObject.newClickableLabel`, `nonClickable.SetText`, `theme.Padding`, `theme.TextSize`

### generateSelectedFilesListTable (method on `*ImportFilesFromGitHubStruct`)
- Signature: `func (*ImportFilesFromGitHubStruct) generateSelectedFilesListTable(parentWindow fyne.Window)`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Create the UI-list that holds the selected files
- Selector calls: `widget.NewLabel`, `widget.NewTable`

### newClickableLabel (method on `*ImportFilesFromGitHubStruct`)
- Signature: `func (*ImportFilesFromGitHubStruct) newClickableLabel(text string, onDoubleTap func(), tempIsClickable bool) *clickableLabel`
- Exported: `false`
- Control-flow features: `none detected`
- Selector calls: `l.ExtendBaseWidget`, `time.Now`

### newCustomLabel (method on `*ImportFilesFromGitHubStruct`)
- Signature: `func (*ImportFilesFromGitHubStruct) newCustomLabel(text string, onDoubleTap func()) *customLabel`
- Exported: `false`
- Control-flow features: `none detected`
- Selector calls: `l.ExtendBaseWidget`, `time.Now`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
