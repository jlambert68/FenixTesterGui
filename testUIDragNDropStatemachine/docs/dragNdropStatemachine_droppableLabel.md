# dragNdropStatemachine_droppableLabel.go

## File Overview
- Path: `testUIDragNDropStatemachine/dragNdropStatemachine_droppableLabel.go`
- Package: `testUIDragNDropStatemachine`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `4`
- Imports: `6`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `MouseIn`
- `MouseMoved`
- `MouseOut`
- `NewDroppableLabel`

## Imports
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/driver/desktop`
- `fyne.io/fyne/v2/widget`
- `image/color`
- `log`

## Declared Types
- `DroppableLabel`

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### MouseIn (method on `*DroppableLabel`)
- Signature: `func (*DroppableLabel) MouseIn(*desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `switch`
- Doc: MouseIn is called when a desktop pointer enters the widget
- Internal calls: `switchStateForSource`, `switchStateForTarget`
- External calls: `log.Fatalln`

### MouseMoved (method on `*DroppableLabel`)
- Signature: `func (*DroppableLabel) MouseMoved(a *desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: MouseMoved is called when a desktop pointer hovers over the widget

### MouseOut (method on `*DroppableLabel`)
- Signature: `func (*DroppableLabel) MouseOut()`
- Exported: `true`
- Control-flow features: `switch`
- Doc: MouseOut is called when a desktop pointer exits the widget
- Internal calls: `switchStateForSource`, `switchStateForTarget`
- External calls: `log.Fatalln`

### NewDroppableLabel (method on `*StateMachineDragAndDropStruct`)
- Signature: `func (*StateMachineDragAndDropStruct) NewDroppableLabel(labelText string, nodeLevel float32, testCaseNodeRectangleSize int, uuid string, testCaseUuid string) *DroppableLabel`
- Exported: `true`
- Control-flow features: `if`
- External calls: `canvas.NewRectangle`, `droppableLabel.ExtendBaseWidget`, `droppableLabel.MinSize`, `droppableLabel.Refresh`, `fyne.NewSize`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
