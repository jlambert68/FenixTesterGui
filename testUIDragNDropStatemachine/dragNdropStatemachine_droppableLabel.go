package testUIDragNDropStatemachine

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
)

type DroppableLabel struct {
	widget.Label
	//topTestCaseAccordion      *widget.Accordion
	//parrentAccordion           *widget.Accordion
	TargetUuid                string
	BackgroundRectangle       *canvas.Rectangle
	IsDroppable               bool
	labelStandardHeight       float32
	nodeLevel                 float32
	testCaseNodeRectangleSize int
	CurrentTestCaseUuid       string
}

// ***** The Object from the Drop Ends *****

// MouseIn is called when a desktop pointer enters the widget
func (b *DroppableLabel) MouseIn(*desktop.MouseEvent) {

	switch stateMachineDragAndDrop.targetStateMachine.currentState {

	case targetStateWaitingForSourceToEnteringTarget:
		return

	case targetStateSourceIsDraggingObject:

		// Verify if this Draggable component can be dropped on this Element

		//if b.IsDroppable == true {
		switchStateForSource(sourceStateEnteringTarget)
		switchStateForTarget(targetStateSourceEnteredTargetWithObject)
		b.BackgroundRectangle.FillColor = color.RGBA{
			R: 0x99,
			G: 0x99,
			B: 0x99,
			A: 0x99,
		}

		b.BackgroundRectangle.Show()
		b.BackgroundRectangle.Refresh()

		stateMachineDragAndDrop.targetDroppableLabel = *b
		//}

	case targetStateSourceEnteredTargetWithObject:
		return

	case targetStateSourceReleasingOnTarget:
		return

	case targetStateSourceReleasedOnTarget:
		return

	default:
		log.Fatalln("Unhandled state for StateMachine(From): ", stateMachineDragAndDrop.targetStateMachine.currentState)

	}

}

/*
func (b *DroppableLabel) TappedSecondary(_ *fyne.PointEvent) {
	log.Println("I have been Secondary tapped")
}
*/

// MouseMoved is called when a desktop pointer hovers over the widget
func (b *DroppableLabel) MouseMoved(a *desktop.MouseEvent) {

}

// MouseOut is called when a desktop pointer exits the widget
func (b *DroppableLabel) MouseOut() {

	switch stateMachineDragAndDrop.targetStateMachine.currentState {

	case targetStateWaitingForSourceToEnteringTarget:
		return

	case targetStateSourceIsDraggingObject:
		return

	case targetStateSourceEnteredTargetWithObject:
		// switch state to 'targetStateSourceIsDraggingObject'
		switchStateForSource(sourceStateDragging)
		switchStateForTarget(targetStateSourceIsDraggingObject)
		b.BackgroundRectangle.FillColor = color.RGBA{
			R: 0x33,
			G: 0x33,
			B: 0x33,
			A: 0x22,
		}
		//b.BackgroundRectangle.Hide()
		b.BackgroundRectangle.Refresh()

	case targetStateSourceReleasingOnTarget:
		return

	case targetStateSourceReleasedOnTarget:
		return

	default:
		log.Fatalln("Unhandled state for StateMachine(From): ", stateMachineDragAndDrop.targetStateMachine.currentState)

	}

}
