package testUIDragNDropStatemachine

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
)

type DraggableLabel struct {
	widget.Label
	SourceUuid        string
	IsDraggable       bool
	BuildingBlockType int
}

// ***** The Object from the Drag starts *****

// Dragged
// When the user press down the mouse button this event is triggered
func (t *DraggableLabel) Dragged(ev *fyne.DragEvent) {

	switch stateMachineDragAndDrop.sourceStateMachine.currentState {

	case sourceStateSearching:
		return

	case sourceStateFinds:
		// switch state to 'sourceStateGrabs'
		switchStateForSource(sourceStateGrabs)

		return

	case sourceStateGrabs:

		// switch state to 'sourceStateDragging'
		switchStateForSource(sourceStateDragging)
		stateMachineDragAndDrop.SourceUuid = t.SourceUuid
		stateMachineDragAndDrop.SourceType = t.BuildingBlockType

		expandDropAreas()

		return

	case sourceStateDragging:
		switchStateForTarget(targetStateSourceIsDraggingObject)
		// Just continue

	case sourceStateReleasingWithOutTarget:
		return

	case sourceStateEnteringTarget:

	case sourceStateReleasingOnTarget:
		return

	case sourceStateReleasedOnTarget:
		return

	default:
		log.Fatalln("Unhandled state for StateMachine(From): ", stateMachineDragAndDrop.sourceStateMachine.currentState)

	}

	// Change Text of 'Drag N Drop'-object
	textRef.Text = t.Text

	// Change size of 'Drag N Drop'-object text backgrounds
	rectangleRef.SetMinSize(textRef.Size().Add(fyne.NewSize(40, 40)))
	rectangle2Ref.SetMinSize(textRef.Size())

	// Move 'Drag N Drop'-object container, so it is to the right of the mouse-pointer
	diffPos := fyne.Position{
		X: -200,
		Y: 50,
	}
	//newPos := ev.AbsolutePosition.Add(diffPos)
	newPos := ev.AbsolutePosition.Add(diffPos).Add(fyne.NewSize(rectangleRef.Size().Width/2, rectangleRef.Size().Height/2))
	containerRef.Move(newPos)

	// Refresh 'Drag N Drop'-object and show them
	containerRef.Refresh()
	if containerRef.Hidden == true {
		containerRef.Show()
	}

}

// DragEnd
// When the user release the mouse button this event is triggered
func (t *DraggableLabel) DragEnd() {

	switch stateMachineDragAndDrop.sourceStateMachine.currentState {

	case sourceStateSearching:
		return

	case sourceStateFinds:
		return

	case sourceStateGrabs:
		return

	case sourceStateDragging:
		switchStateForSource(sourceStateReleasingWithOutTarget)
		switchStateForTarget(targetStateWaitingForSourceToEnteringTarget)
		stateMachineDragAndDrop.SourceUuid = ""
		stateMachineDragAndDrop.SourceType = 0

		shrinkDropAreas()

	case sourceStateReleasingWithOutTarget:
		stateMachineDragAndDrop.SourceUuid = ""
		// Just continue

	case sourceStateEnteringTarget:
		switchStateForSource(sourceStateReleasingOnTarget)
		switchStateForTarget(targetStateSourceReleasingOnTarget)

		shrinkDropAreas()

		for _, droppableTargetLabel := range stateMachineDragAndDrop.registeredDroppableTargetLabels {
			droppableTargetLabel.BackgroundRectangle.StrokeWidth = 0
			droppableTargetLabel.BackgroundRectangle.StrokeColor = color.RGBA{
				R: 0x00,
				G: 0x00,
				B: 0x00,
				A: 0x00,
			}
			droppableTargetLabel.BackgroundRectangle.FillColor = color.RGBA{
				R: 0x00,
				G: 0x00,
				B: 0x00,
				A: 0x00,
			}
		}

		executeDropAction()

	case sourceStateReleasingOnTarget:
		switchStateForSource(sourceStateReleasedOnTarget)

	case sourceStateReleasedOnTarget:
		switchStateForTarget(targetStateSourceReleasedOnTarget)

	default:
		log.Fatalln("Unhandled state for StateMachine(From): ", stateMachineDragAndDrop.sourceStateMachine.currentState)

	}

	// Hide the 'Drag N Drop'-objects
	containerRef.Hide()
	containerRef.Refresh()

	// switch state to 'sourceStateSearching'
	switchStateForSource(sourceStateSearching)

}

// MouseIn is called when a desktop pointer enters the widget
func (b *DraggableLabel) MouseIn(*desktop.MouseEvent) {

	switch stateMachineDragAndDrop.sourceStateMachine.currentState {

	case sourceStateSearching:
		// Mouse finds draggable object
		if b.IsDraggable == true {
			switchStateForSource(sourceStateFinds)
		}

	case sourceStateFinds:
		return

	case sourceStateGrabs:
		return

	case sourceStateDragging:
		return

	case sourceStateReleasingWithOutTarget:
		return

	case sourceStateEnteringTarget:
		return

	case sourceStateReleasingOnTarget:
		return

	case sourceStateReleasedOnTarget:
		return

	default:
		log.Fatalln("Unhandled state for StateMachine(From): ", stateMachineDragAndDrop.sourceStateMachine.currentState)

	}

}

// MouseMoved is called when a desktop pointer hovers over the widget
func (b *DraggableLabel) MouseMoved(a *desktop.MouseEvent) {

}

// MouseOut is called when a desktop pointer exits the widget
func (b *DraggableLabel) MouseOut() {

	switch stateMachineDragAndDrop.sourceStateMachine.currentState {

	case sourceStateSearching:
		return

	case sourceStateFinds:
		// Mouse leaves Draggable Object before grabbing it
		switchStateForSource(sourceStateSearching)

	case sourceStateGrabs:
		return

	case sourceStateDragging:
		return

	case sourceStateReleasingWithOutTarget:
		return

	case sourceStateEnteringTarget:
		return

	case sourceStateReleasingOnTarget:
		return

	case sourceStateReleasedOnTarget:
		return

	default:
		log.Fatalln("Unhandled state for StateMachine(From): ", stateMachineDragAndDrop.sourceStateMachine.currentState)

	}

}
