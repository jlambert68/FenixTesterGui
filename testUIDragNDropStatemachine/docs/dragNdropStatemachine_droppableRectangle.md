# dragNdropStatemachine_droppableRectangle.go

## File Overview
- Path: `testUIDragNDropStatemachine/dragNdropStatemachine_droppableRectangle.go`
- Package: `testUIDragNDropStatemachine`
- Functions/Methods: `11`
- Imports: `6`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `CreateRenderer`
- `Destroy`
- `Layout`
- `MinSize`
- `MouseIn`
- `MouseMoved`
- `MouseOut`
- `NewDroppableRectangle`
- `Objects`
- `Refresh`
- `Tapped`

## Imports
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/driver/desktop`
- `fyne.io/fyne/v2/widget`
- `image/color`
- `log`

## Declared Types
- `DroppableRectangle`
- `droppableRectRenderer`

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### CreateRenderer (method on `*DroppableRectangle`)
- Signature: `func (*DroppableRectangle) CreateRenderer() fyne.WidgetRenderer`
- Exported: `true`
- Control-flow features: `none detected`

### Destroy (method on `*droppableRectRenderer`)
- Signature: `func (*droppableRectRenderer) Destroy()`
- Exported: `true`
- Control-flow features: `none detected`

### Layout (method on `*droppableRectRenderer`)
- Signature: `func (*droppableRectRenderer) Layout(size fyne.Size)`
- Exported: `true`
- Control-flow features: `none detected`

### MinSize (method on `*droppableRectRenderer`)
- Signature: `func (*droppableRectRenderer) MinSize() fyne.Size`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `fyne.NewSize`

### MouseIn (method on `*DroppableRectangle`)
- Signature: `func (*DroppableRectangle) MouseIn(e *desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `switch`
- Internal calls: `switchStateForSource`, `switchStateForTarget`
- Selector calls: `log.Fatalln`

### MouseMoved (method on `*DroppableRectangle`)
- Signature: `func (*DroppableRectangle) MouseMoved(*desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`

### MouseOut (method on `*DroppableRectangle`)
- Signature: `func (*DroppableRectangle) MouseOut()`
- Exported: `true`
- Control-flow features: `switch`
- Internal calls: `switchStateForSource`, `switchStateForTarget`
- Selector calls: `log.Fatalln`

### NewDroppableRectangle (method on `*StateMachineDragAndDropStruct`)
- Signature: `func (*StateMachineDragAndDropStruct) NewDroppableRectangle(nodeLevel float32, testCaseNodeRectangleSize int, uuid string, testCaseUuid string, tempParentTestInstructionContainerUuid string) *DroppableRectangle`
- Exported: `true`
- Control-flow features: `if`
- Selector calls: `canvas.NewRectangle`, `droppableRectangle.ExtendBaseWidget`, `fyne.NewSize`, `rect.SetMinSize`

### Objects (method on `*droppableRectRenderer`)
- Signature: `func (*droppableRectRenderer) Objects() []fyne.CanvasObject`
- Exported: `true`
- Control-flow features: `none detected`

### Refresh (method on `*droppableRectRenderer`)
- Signature: `func (*droppableRectRenderer) Refresh()`
- Exported: `true`
- Control-flow features: `none detected`

### Tapped (method on `*DroppableRectangle`)
- Signature: `func (*DroppableRectangle) Tapped(*fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `none detected`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
