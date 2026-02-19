# dragNdropStatemachine.go

## File Overview
- Path: `testUIDragNDropStatemachine/dragNdropStatemachine.go`
- Package: `testUIDragNDropStatemachine`
- Functions/Methods: `8`
- Imports: `8`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `InitiateStateStateMachine`
- `TestInstructionContainerBondBelongingRectangleToMap`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/testCase/testCaseModel`
- `errors`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/widget`
- `image/color`

## Declared Types
- `StateMachineDragAndDropStruct`
- `noneDroppableLabel`
- `stateMachineSourceAndDestinationStruct`
- `targetDroppedTypeType`

## Declared Constants
- `droppableLabelType`
- `droppableRectangleType`
- `sourceStateDragging`
- `sourceStateEnteringTarget`
- `sourceStateFinds`
- `sourceStateGrabs`
- `sourceStateReleasedOnTarget`
- `sourceStateReleasingOnTarget`
- `sourceStateReleasingWithOutTarget`
- `sourceStateSearching`
- `targetDropLabelRectangleHeight`
- `targetDropLabelRectangleWidth`
- `targetDropRectangleHeight`
- `targetDropRectangleWidth`
- `targetStateSourceEnteredTargetWithObject`
- `targetStateSourceIsDraggingObject`
- `targetStateSourceReleasedOnTarget`
- `targetStateSourceReleasingOnTarget`
- `targetStateWaitingForSourceToEnteringTarget`
- `unspecifiedType`

## Declared Variables
- `commandChannelReference`
- `containerRef`
- `labelStandardHeight`
- `rectangle2Ref`
- `rectangleRef`
- `stateMachineDragAndDrop`
- `textRef`

## Functions and Methods
### InitiateStateStateMachine (method on `*StateMachineDragAndDropStruct`)
- Signature: `func (*StateMachineDragAndDropStruct) InitiateStateStateMachine(dragNDropText *canvas.Text, dragNDropRectangleRef *canvas.Rectangle, dragNDropRectangle2Ref *canvas.Rectangle, dragNDropContainerRef *fyne.Container, commandChannelRef *sharedCode.CommandChannelType, testCasesRef *testCaseModel.TestCasesModelsStruct)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: InitiateStateStateMachine InitiateState State machine
- Selector calls: `tempLabel.MinSize`, `tempLabel.Refresh`, `widget.NewLabel`

### TestInstructionContainerBondBelongingRectangleToMap
- Signature: `func TestInstructionContainerBondBelongingRectangleToMap(testInstructionContainerUuid string, rectangle *canvas.Rectangle)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: Add testInstructionContainerBondBelongingRectangle to map

### executeDropAction
- Signature: `func executeDropAction()`
- Exported: `false`
- Control-flow features: `switch`
- Selector calls: `errors.New`, `fmt.Println`, `fmt.Sprintf`, `sharedCode.BuildingBlock`

### expandDropAreas
- Signature: `func expandDropAreas()`
- Exported: `false`
- Control-flow features: `if, for/range`
- Selector calls: `fyne.NewSize`, `targetLabel.Resize`

### newNoneDroppableLabel
- Signature: `func newNoneDroppableLabel(uuid string) *noneDroppableLabel`
- Exported: `false`
- Control-flow features: `none detected`
- Selector calls: `nonDroppableLabel.ExtendBaseWidget`

### shrinkDropAreas
- Signature: `func shrinkDropAreas()`
- Exported: `false`
- Control-flow features: `if, for/range`

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
