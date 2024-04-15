package testUIDragNDropStatemachine

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
)

type DroppableRectangle struct {
	widget.BaseWidget
	rect                      *canvas.Rectangle
	TargetUuid                string
	IsDroppable               bool
	labelStandardHeight       float32
	nodeLevel                 float32
	testCaseNodeRectangleSize int
	CurrentTestCaseUuid       string
}

func (h *DroppableRectangle) CreateRenderer() fyne.WidgetRenderer {
	return &droppableRectRenderer{rect: h.rect}
}

type droppableRectRenderer struct {
	rect *canvas.Rectangle
}

func (r *droppableRectRenderer) Destroy() {
	// Any necessary teardown can be added here
}

func (r *droppableRectRenderer) Layout(size fyne.Size) {
	r.rect.Resize(size)
}

func (r *droppableRectRenderer) MinSize() fyne.Size {
	return fyne.NewSize(150, 150) // Ensure the minimum size covers the rectangle
}

func (r *droppableRectRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.rect}
}

func (r *droppableRectRenderer) Refresh() {
	r.rect.Refresh()
}

func (h *DroppableRectangle) Tapped(*fyne.PointEvent) {
	// Add tap handling if necessary
}

func (h *DroppableRectangle) MouseIn(e *desktop.MouseEvent) {

	switch stateMachineDragAndDrop.targetStateMachine.currentState {

	case targetStateWaitingForSourceToEnteringTarget:
		return

	case targetStateSourceIsDraggingObject:

		// Verify if this Draggable component can be dropped on this Element

		//if b.IsDroppable == true {
		switchStateForSource(sourceStateEnteringTarget)
		switchStateForTarget(targetStateSourceEnteredTargetWithObject)
		h.rect.FillColor = color.RGBA{
			R: 0x99,
			G: 0x99,
			B: 0x99,
			A: 0x99,
		}

		h.rect.Show()
		h.rect.Refresh()

		stateMachineDragAndDrop.targetDroppableRectangle = *h
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

func (h *DroppableRectangle) MouseOut() {

	switch stateMachineDragAndDrop.targetStateMachine.currentState {

	case targetStateWaitingForSourceToEnteringTarget:
		return

	case targetStateSourceIsDraggingObject:
		return

	case targetStateSourceEnteredTargetWithObject:
		// switch state to 'targetStateSourceIsDraggingObject'
		switchStateForSource(sourceStateDragging)
		switchStateForTarget(targetStateSourceIsDraggingObject)
		h.rect.FillColor = color.RGBA{
			R: 0x33,
			G: 0x33,
			B: 0x33,
			A: 0x22,
		}
		//b.BackgroundRectangle.Hide()
		h.rect.Refresh()

	case targetStateSourceReleasingOnTarget:
		return

	case targetStateSourceReleasedOnTarget:
		return

	default:
		log.Fatalln("Unhandled state for StateMachine(From): ", stateMachineDragAndDrop.targetStateMachine.currentState)

	}

}

func (h *DroppableRectangle) MouseMoved(*desktop.MouseEvent) {
	// Implement if needed
}
