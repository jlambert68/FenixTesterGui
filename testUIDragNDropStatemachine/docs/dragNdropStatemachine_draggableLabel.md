# dragNdropStatemachine_draggableLabel.go

## File Overview
- Path: `testUIDragNDropStatemachine/dragNdropStatemachine_draggableLabel.go`
- Package: `testUIDragNDropStatemachine`
- Functions/Methods: `6`
- Imports: `5`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `DragEnd`
- `Dragged`
- `MouseIn`
- `MouseMoved`
- `MouseOut`
- `NewDraggableLabel`

## Imports
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/driver/desktop`
- `fyne.io/fyne/v2/widget`
- `image/color`
- `log`

## Declared Types
- `DraggableLabel`

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### DragEnd (method on `*DraggableLabel`)
- Signature: `func (*DraggableLabel) DragEnd()`
- Exported: `true`
- Control-flow features: `if, for/range, switch`
- Doc: DragEnd When the user release the mouse button this event is triggered
- Internal calls: `executeDropAction`, `shrinkDropAreas`, `switchStateForSource`, `switchStateForTarget`
- Selector calls: `containerRef.Hide`, `containerRef.Refresh`, `log.Fatalln`

### Dragged (method on `*DraggableLabel`)
- Signature: `func (*DraggableLabel) Dragged(ev *fyne.DragEvent)`
- Exported: `true`
- Control-flow features: `if, switch`
- Doc: Dragged When the user press down the mouse button this event is triggered
- Internal calls: `expandDropAreas`, `switchStateForSource`, `switchStateForTarget`
- Selector calls: `containerRef.Move`, `containerRef.Refresh`, `containerRef.Show`, `fyne.NewSize`, `log.Fatalln`, `rectangle2Ref.SetMinSize`, `rectangleRef.SetMinSize`, `rectangleRef.Size`

### MouseIn (method on `*DraggableLabel`)
- Signature: `func (*DraggableLabel) MouseIn(*desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `if, switch`
- Doc: MouseIn is called when a desktop pointer enters the widget
- Internal calls: `switchStateForSource`
- Selector calls: `log.Fatalln`

### MouseMoved (method on `*DraggableLabel`)
- Signature: `func (*DraggableLabel) MouseMoved(a *desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: MouseMoved is called when a desktop pointer hovers over the widget

### MouseOut (method on `*DraggableLabel`)
- Signature: `func (*DraggableLabel) MouseOut()`
- Exported: `true`
- Control-flow features: `switch`
- Doc: MouseOut is called when a desktop pointer exits the widget
- Internal calls: `switchStateForSource`
- Selector calls: `log.Fatalln`

### NewDraggableLabel (method on `*StateMachineDragAndDropStruct`)
- Signature: `func (*StateMachineDragAndDropStruct) NewDraggableLabel(uuid string) *DraggableLabel`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `draggableLabel.ExtendBaseWidget`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
