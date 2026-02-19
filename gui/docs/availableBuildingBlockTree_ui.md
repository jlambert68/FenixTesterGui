# availableBuildingBlockTree_ui.go

## File Overview
- Path: `gui/availableBuildingBlockTree_ui.go`
- Package: `gui`
- Functions/Methods: `6`
- Imports: `5`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `DragEnd`
- `Dragged`
- `Tapped`
- `TappedSecondary`

## Imports
- `FenixTesterGui/testUIDragNDropStatemachine`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/widget`
- `log`

## Declared Types
- `tappableLabel`

## Declared Constants
- None

## Declared Variables
- `list`
- `tree`

## Functions and Methods
### makeTreeUI (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) makeTreeUI()`
- Exported: `false`
- Control-flow features: `if`
- Internal calls: `int`

### newTappableLabel
- Signature: `func newTappableLabel() *tappableLabel`
- Exported: `false`
- Control-flow features: `none detected`
- Selector calls: `label.ExtendBaseWidget`

### Tapped (method on `*tappableLabel`)
- Signature: `func (*tappableLabel) Tapped(_ *fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `log.Println`, `fmt.Println`, `t.Position`

### TappedSecondary (method on `*tappableLabel`)
- Signature: `func (*tappableLabel) TappedSecondary(_ *fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `log.Println`

### Dragged (method on `*tappableLabel`)
- Signature: `func (*tappableLabel) Dragged(ev *fyne.DragEvent)`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `log.Println`, `t.Position`, `widget.NewLabel`, `fmt.Println`

### DragEnd (method on `*tappableLabel`)
- Signature: `func (*tappableLabel) DragEnd()`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `log.Println`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
