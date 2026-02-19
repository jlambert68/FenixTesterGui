# dragNdropTest.go

## File Overview
- Path: `testCase/testCaseUI/dragNdropTest.go`
- Package: `testCaseUI`
- Generated: `2026-02-19T14:23:17+01:00`
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
### DragEnd (method on `*draggableLabel`)
- Signature: `func (*draggableLabel) DragEnd()`
- Exported: `true`
- Control-flow features: `for/range, switch`
- Doc: DragEnd When the user release the mouse button this event is triggered
- Internal calls: `executeDropAction`, `shrinkDropAreas`, `switchStateForSource`, `switchStateForTarget`
- External calls: `containerRef.Hide`, `containerRef.Refresh`, `log.Fatalln`

### Dragged (method on `*draggableLabel`)
- Signature: `func (*draggableLabel) Dragged(ev *fyne.DragEvent)`
- Exported: `true`
- Control-flow features: `switch`
- Doc: Dragged When the user press down the mouse button this event is triggered
- Internal calls: `expandDropAreas`, `switchStateForSource`, `switchStateForTarget`
- External calls: `containerRef.Move`, `containerRef.Show`, `fyne.NewSize`, `log.Fatalln`, `rectangle2Ref.SetMinSize`, `rectangleRef.SetMinSize`, `rectangleRef.Size`, `textRef.Size`

### MouseIn (method on `*draggableLabel`)
- Signature: `func (*draggableLabel) MouseIn(*desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `switch`
- Doc: MouseIn is called when a desktop pointer enters the widget
- Internal calls: `switchStateForSource`
- External calls: `log.Fatalln`

### MouseIn (method on `*droppableLabel`)
- Signature: `func (*droppableLabel) MouseIn(*desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `switch`
- Doc: MouseIn is called when a desktop pointer enters the widget
- Internal calls: `switchStateForSource`, `switchStateForTarget`
- External calls: `log.Fatalln`

### MouseMoved (method on `*droppableLabel`)
- Signature: `func (*droppableLabel) MouseMoved(a *desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: MouseMoved is called when a desktop pointer hovers over the widget

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
- External calls: `log.Fatalln`

### MouseOut (method on `*droppableLabel`)
- Signature: `func (*droppableLabel) MouseOut()`
- Exported: `true`
- Control-flow features: `switch`
- Doc: MouseOut is called when a desktop pointer exits the widget
- Internal calls: `switchStateForSource`, `switchStateForTarget`
- External calls: `log.Fatalln`

### executeDropAction
- Signature: `func executeDropAction()`
- Exported: `false`
- Control-flow features: `none detected`
- External calls: `fmt.Println`, `fmt.Sprintf`

### expandDropAreas
- Signature: `func expandDropAreas()`
- Exported: `false`
- Control-flow features: `for/range, go`
- External calls: `canvas.NewColorRGBAAnimation`, `canvas.NewSizeAnimation`, `canvas.Refresh`, `fyne.NewSize`, `rectangleColorAnimation.Start`, `rectangleSizeAnimation.Start`, `targetLabel.Show`, `time.Sleep`

### makeDragNDropTestGUI
- Signature: `func makeDragNDropTestGUI(textIn *canvas.Text, recIn *canvas.Rectangle, rec2In *canvas.Rectangle, containerIn *fyne.Container) myCanvasObject fyne.CanvasObject`
- Exported: `false`
- Control-flow features: `none detected`
- Internal calls: `newDraggableLabel`, `newDroppableLabel`, `newNoneDroppableLabel`
- External calls: `canvas.NewRectangle`, `container.New`, `container.NewHBox`, `container.NewMax`, `container.NewVBox`, `dragToDrop2Label.Size`, `fyne.NewSize`, `layout.NewSpacer`

### newDraggableLabel
- Signature: `func newDraggableLabel(uuid string) *draggableLabel`
- Exported: `false`
- Control-flow features: `none detected`
- External calls: `draggableLabel.ExtendBaseWidget`

### newDroppableLabel
- Signature: `func newDroppableLabel(uuid string) *droppableLabel`
- Exported: `false`
- Control-flow features: `none detected`
- External calls: `canvas.NewRectangle`, `droppableLabel.ExtendBaseWidget`, `droppableLabel.Refresh`, `droppableLabel.Size`

### newNoneDroppableLabel
- Signature: `func newNoneDroppableLabel(uuid string) *noneDroppableLabel`
- Exported: `false`
- Control-flow features: `none detected`
- External calls: `nonDroppableLabel.ExtendBaseWidget`

### shrinkDropAreas
- Signature: `func shrinkDropAreas()`
- Exported: `false`
- Control-flow features: `for/range, go`
- External calls: `canvas.NewColorRGBAAnimation`, `canvas.NewSizeAnimation`, `canvas.Refresh`, `fyne.NewSize`, `rectangleColorAnimation.Start`, `rectangleSizeAnimation.Start`, `targetLabel.Hide`, `time.Sleep`

### switchStateForSource
- Signature: `func switchStateForSource(newState int)`
- Exported: `false`
- Control-flow features: `none detected`

### switchStateForTarget
- Signature: `func switchStateForTarget(newState int)`
- Exported: `false`
- Control-flow features: `none detected`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
