# dragNdropTest.go

## File Overview
- Path: `testCase/testCaseUI/dragNdropTest.go`
- Package: `testCaseUI`
- Functions/Methods: `17`
- Imports: `10`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `DragEnd`
- `Dragged`
- `MouseIn`
- `MouseMoved`
- `MouseOut`

## Imports
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/driver/desktop`
- `fyne.io/fyne/v2/layout`
- `fyne.io/fyne/v2/widget`
- `image/color`
- `log`
- `time`

## Declared Types
- `StateMachineSouceAndDestinationStruct`
- `draggableLabel`
- `droppableLabel`
- `noneDroppableLabel`
- `stateMachineDragAndDropStruct`

## Declared Constants
- `sourceStateDragging`
- `sourceStateEnteringTarget`
- `sourceStateFinds`
- `sourceStateGrabs`
- `sourceStateReleasedOnTarget`
- `sourceStateReleasingOnTarget`
- `sourceStateReleasingWithOutTarget`
- `sourceStateSearching`
- `targetStateSourceEnteredTargetWithObject`
- `targetStateSourceIsDraggingObject`
- `targetStateSourceReleasedOnTarget`
- `targetStateSourceReleasingOnTarget`
- `targetStateWaitingForSourceToEnteringTarget`

## Declared Variables
- `DropFour`
- `containerRef`
- `dropContainer`
- `labelStandardHeight`
- `rectangle2Ref`
- `rectangleRef`
- `registeredDroppableTargetLabels`
- `registeredNoneDroppableTargetLabels`
- `registeredThinDroppableContainers`
- `stateMachineDragAndDrop`
- `stateMachineDragFrom`
- `stateMachineTarget`
- `textRef`

## Functions and Methods
### makeDragNDropTestGUI
- Signature: `func makeDragNDropTestGUI(textIn *canvas.Text, recIn *canvas.Rectangle, rec2In *canvas.Rectangle, containerIn *fyne.Container) myCanvasObject fyne.CanvasObject`
- Exported: `false`
- Control-flow features: `none detected`
- Internal calls: `newDraggableLabel`, `newNoneDroppableLabel`, `newDroppableLabel`
- Selector calls: `container.NewMax`, `dragToDrop2Label.Size`, `container.NewHBox`, `container.NewVBox`, `widget.NewLabel`, `container.New`, `layout.NewVBoxLayout`, `widget.NewAccordionItem`

### newDraggableLabel
- Signature: `func newDraggableLabel(uuid string) *draggableLabel`
- Exported: `false`
- Control-flow features: `none detected`
- Selector calls: `draggableLabel.ExtendBaseWidget`

### newDroppableLabel
- Signature: `func newDroppableLabel(uuid string) *droppableLabel`
- Exported: `false`
- Control-flow features: `none detected`
- Selector calls: `droppableLabel.ExtendBaseWidget`, `canvas.NewRectangle`, `droppableLabel.Refresh`, `droppableLabel.Size`

### newNoneDroppableLabel
- Signature: `func newNoneDroppableLabel(uuid string) *noneDroppableLabel`
- Exported: `false`
- Control-flow features: `none detected`
- Selector calls: `nonDroppableLabel.ExtendBaseWidget`

### Dragged (method on `*draggableLabel`)
- Signature: `func (*draggableLabel) Dragged(ev *fyne.DragEvent)`
- Exported: `true`
- Control-flow features: `switch`
- Doc: Dragged When the user press down the mouse button this event is triggered
- Internal calls: `switchStateForSource`, `expandDropAreas`, `switchStateForTarget`
- Selector calls: `log.Fatalln`, `rectangleRef.SetMinSize`, `textRef.Size`, `fyne.NewSize`, `rectangle2Ref.SetMinSize`, `rectangleRef.Size`, `containerRef.Move`, `containerRef.Show`

### DragEnd (method on `*draggableLabel`)
- Signature: `func (*draggableLabel) DragEnd()`
- Exported: `true`
- Control-flow features: `for/range, switch`
- Doc: DragEnd When the user release the mouse button this event is triggered
- Internal calls: `switchStateForSource`, `switchStateForTarget`, `shrinkDropAreas`, `executeDropAction`
- Selector calls: `log.Fatalln`, `containerRef.Hide`, `containerRef.Refresh`

### MouseIn (method on `*draggableLabel`)
- Signature: `func (*draggableLabel) MouseIn(*desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `switch`
- Doc: MouseIn is called when a desktop pointer enters the widget
- Internal calls: `switchStateForSource`
- Selector calls: `log.Fatalln`

### MouseMoved (method on `*draggableLabel`)
- Signature: `func (*draggableLabel) MouseMoved(a *desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: MouseMoved is called when a desktop pointer hovers over the widget

### MouseOut (method on `*draggableLabel`)
- Signature: `func (*draggableLabel) MouseOut()`
- Exported: `true`
- Control-flow features: `switch`
- Doc: MouseOut is called when a desktop pointer exits the widget
- Internal calls: `switchStateForSource`
- Selector calls: `log.Fatalln`

### MouseIn (method on `*droppableLabel`)
- Signature: `func (*droppableLabel) MouseIn(*desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `switch`
- Doc: MouseIn is called when a desktop pointer enters the widget
- Internal calls: `switchStateForSource`, `switchStateForTarget`
- Selector calls: `log.Fatalln`

### MouseMoved (method on `*droppableLabel`)
- Signature: `func (*droppableLabel) MouseMoved(a *desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: MouseMoved is called when a desktop pointer hovers over the widget

### MouseOut (method on `*droppableLabel`)
- Signature: `func (*droppableLabel) MouseOut()`
- Exported: `true`
- Control-flow features: `switch`
- Doc: MouseOut is called when a desktop pointer exits the widget
- Internal calls: `switchStateForSource`, `switchStateForTarget`
- Selector calls: `log.Fatalln`

### switchStateForSource
- Signature: `func switchStateForSource(newState int)`
- Exported: `false`
- Control-flow features: `none detected`

### switchStateForTarget
- Signature: `func switchStateForTarget(newState int)`
- Exported: `false`
- Control-flow features: `none detected`

### expandDropAreas
- Signature: `func expandDropAreas()`
- Exported: `false`
- Control-flow features: `for/range, go`
- Selector calls: `canvas.NewColorRGBAAnimation`, `canvas.Refresh`, `canvas.NewSizeAnimation`, `fyne.NewSize`, `rectangleColorAnimation.Start`, `rectangleSizeAnimation.Start`, `time.Sleep`, `targetLabel.Show`

### shrinkDropAreas
- Signature: `func shrinkDropAreas()`
- Exported: `false`
- Control-flow features: `for/range, go`
- Selector calls: `targetLabel.Hide`, `canvas.NewColorRGBAAnimation`, `canvas.Refresh`, `canvas.NewSizeAnimation`, `fyne.NewSize`, `rectangleColorAnimation.Start`, `rectangleSizeAnimation.Start`, `time.Sleep`

### executeDropAction
- Signature: `func executeDropAction()`
- Exported: `false`
- Control-flow features: `none detected`
- Selector calls: `fmt.Println`, `fmt.Sprintf`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
