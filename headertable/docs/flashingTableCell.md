# flashingTableCell.go

## File Overview
- Path: `headertable/flashingTableCell.go`
- Package: `headertable`
- Functions/Methods: `12`
- Imports: `9`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `CreateRenderer`
- `Destroy`
- `DoubleTapped`
- `FlashAddedRow`
- `FlashRowToBeRemoved`
- `Layout`
- `MinSize`
- `NewFlashingTableCell`
- `Objects`
- `Refresh`
- `Tapped`
- `TappedSecondary`

## Imports
- `FenixTesterGui/executions/detailedExecutionsModel`
- `FenixTesterGui/resources`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/widget`
- `image/color`
- `time`

## Declared Types
- `FlashCellWhenAddToTableFunctionType`
- `FlashCellWhenRemoveFromTableFunctionType`
- `FlashingTableCellStruct`
- `TestCaseExecutionMapKeyType`
- `flashingTableCellRenderer`

## Declared Constants
- None

## Declared Variables
- `_`
- `backgroundRectangleBaseColor`
- `headerBackgroundRectangleBaseColor`

## Functions and Methods
### CreateRenderer (method on `*FlashingTableCellStruct`)
- Signature: `func (*FlashingTableCellStruct) CreateRenderer() fyne.WidgetRenderer`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `container.NewMax`

### Destroy (method on `*flashingTableCellRenderer`)
- Signature: `func (*flashingTableCellRenderer) Destroy()`
- Exported: `true`
- Control-flow features: `none detected`

### DoubleTapped (method on `*FlashingTableCellStruct`)
- Signature: `func (*FlashingTableCellStruct) DoubleTapped(_ *fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `if`
- Internal calls: `string`
- Selector calls: `detailedExecutionsModel.RemoveTestCaseExecutionFromSummaryTable`, `detailedExecutionsModel.RetrieveSingleTestCaseExecution`, `fmt.Println`

### FlashAddedRow
- Signature: `func FlashAddedRow(flashingTableCell *FlashingTableCellStruct)`
- Exported: `true`
- Control-flow features: `go`
- Selector calls: `canvas.NewColorRGBAAnimation`, `canvas.Refresh`, `rectangleColorAnimation.Start`

### FlashRowToBeRemoved
- Signature: `func FlashRowToBeRemoved(flashingTableCell *FlashingTableCellStruct)`
- Exported: `true`
- Control-flow features: `go`
- Selector calls: `canvas.NewColorRGBAAnimation`, `canvas.Refresh`, `rectangleColorAnimation.Start`

### Layout (method on `*flashingTableCellRenderer`)
- Signature: `func (*flashingTableCellRenderer) Layout(size fyne.Size)`
- Exported: `true`
- Control-flow features: `none detected`

### MinSize (method on `*flashingTableCellRenderer`)
- Signature: `func (*flashingTableCellRenderer) MinSize() fyne.Size`
- Exported: `true`
- Control-flow features: `none detected`

### NewFlashingTableCell
- Signature: `func NewFlashingTableCell(text string) *FlashingTableCellStruct`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `canvas.NewImageFromResource`, `canvas.NewRectangle`, `newFlashingTableCell.ExtendBaseWidget`, `widget.NewLabel`

### Objects (method on `*flashingTableCellRenderer`)
- Signature: `func (*flashingTableCellRenderer) Objects() []fyne.CanvasObject`
- Exported: `true`
- Control-flow features: `none detected`

### Refresh (method on `*flashingTableCellRenderer`)
- Signature: `func (*flashingTableCellRenderer) Refresh()`
- Exported: `true`
- Control-flow features: `none detected`

### Tapped (method on `*FlashingTableCellStruct`)
- Signature: `func (*FlashingTableCellStruct) Tapped(_ *fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `fmt.Println`

### TappedSecondary (method on `*FlashingTableCellStruct`)
- Signature: `func (*FlashingTableCellStruct) TappedSecondary(_ *fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `fmt.Println`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
