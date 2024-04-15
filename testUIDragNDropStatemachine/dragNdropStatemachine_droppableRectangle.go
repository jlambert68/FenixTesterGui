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
	Rectangle                               *canvas.Rectangle
	TargetUuid                              string
	IsDroppable                             bool
	labelStandardHeight                     float32
	nodeLevel                               float32
	testCaseNodeRectangleSize               int
	CurrentTestCaseUuid                     string
	parentTestInstructionContainerUuid      string
	parentTestInstructionContainerRectangle *canvas.Rectangle
}

func (stateMachine *StateMachineDragAndDropStruct) NewDroppableRectangle(
	nodeLevel float32,
	testCaseNodeRectangleSize int,
	uuid string,
	testCaseUuid string,
	tempParentTestInstructionContainerUuid string) *DroppableRectangle {

	color := color.RGBA{
		R: 0x33,
		G: 0x33,
		B: 0x33,
		A: 0x00,
	}

	// Extract parent TestInstructionContainer-Bond-rectangle
	var parentTestInstructionContainerRectangle *canvas.Rectangle
	parentTestInstructionContainerRectangle = stateMachineDragAndDrop.testInstructionContainerBondBelongingRectangleMap[tempParentTestInstructionContainerUuid]

	rect := canvas.NewRectangle(color)
	rect.SetMinSize(fyne.NewSize(targetDropRectangleWidth, targetDropRectangleHeight))
	droppableRectangle := &DroppableRectangle{
		Rectangle:                               rect,
		parentTestInstructionContainerUuid:      tempParentTestInstructionContainerUuid,
		parentTestInstructionContainerRectangle: parentTestInstructionContainerRectangle,
	}

	droppableRectangle.ExtendBaseWidget(droppableRectangle)

	droppableRectangle.TargetUuid = uuid
	droppableRectangle.nodeLevel = nodeLevel
	droppableRectangle.CurrentTestCaseUuid = testCaseUuid

	stateMachineDragAndDrop.registeredDroppableTargetRectangle = append(stateMachineDragAndDrop.registeredDroppableTargetRectangle, droppableRectangle)

	return droppableRectangle
}

func (h *DroppableRectangle) CreateRenderer() fyne.WidgetRenderer {
	return &droppableRectRenderer{rectangle: h.Rectangle}
}

type droppableRectRenderer struct {
	rectangle *canvas.Rectangle
}

func (r *droppableRectRenderer) Destroy() {
	// Any necessary teardown can be added here
}

func (r *droppableRectRenderer) Layout(size fyne.Size) {
	r.rectangle.Resize(size)
}

func (r *droppableRectRenderer) MinSize() fyne.Size {
	return fyne.NewSize(targetDropRectangleWidth, targetDropRectangleHeight) // Ensure the minimum size covers the rectangle
}

func (r *droppableRectRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.rectangle}
}

func (r *droppableRectRenderer) Refresh() {
	r.rectangle.Refresh()
}

func (h *DroppableRectangle) Tapped(*fyne.PointEvent) {
	// Add tap handling if necessary
}

func (h *DroppableRectangle) MouseIn(e *desktop.MouseEvent) {

	// Set targetDroppedType to use
	stateMachineDragAndDrop.targetDroppedType = droppableRectangleType

	switch stateMachineDragAndDrop.targetStateMachine.currentState {

	case targetStateWaitingForSourceToEnteringTarget:
		return

	case targetStateSourceIsDraggingObject:

		// Verify if this Draggable component can be dropped on this Element

		//if b.IsDroppable == true {
		switchStateForSource(sourceStateEnteringTarget)
		switchStateForTarget(targetStateSourceEnteredTargetWithObject)
		h.Rectangle.FillColor = color.RGBA{
			R: 0x99,
			G: 0x99,
			B: 0x99,
			A: 0xFF,
		}

		h.Rectangle.Show()
		h.Rectangle.Refresh()

		stateMachineDragAndDrop.targetDroppableRectangle = *h

		// Show parent TestInstructionContainer
		h.parentTestInstructionContainerRectangle.FillColor = color.White

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
		h.Rectangle.FillColor = color.RGBA{
			R: 0x03,
			G: 0x03,
			B: 0x03,
			A: 0xFF,
		}
		//b.BackgroundRectangle.Hide()
		h.Rectangle.Refresh()

		// Hide parent TestInstructionContainer
		h.parentTestInstructionContainerRectangle.FillColor = color.Transparent

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
