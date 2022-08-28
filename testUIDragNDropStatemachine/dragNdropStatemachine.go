package testUIDragNDropStatemachine

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
	"time"
)

// Statea for handling Drag from Source object
const (
	sourceStateSearching              = iota // 0
	sourceStateFinds                         // 1
	sourceStateGrabs                         // 2
	sourceStateDragging                      // 3
	sourceStateReleasingWithOutTarget        // 4
	sourceStateEnteringTarget                // 5
	sourceStateReleasingOnTarget             // 6
	sourceStateReleasedOnTarget              // 7
)

// State for handling Drop-target object
const (
	targetStateWaitingForSourceToEnteringTarget = iota // 0
	targetStateSourceIsDraggingObject                  // 1
	targetStateSourceEnteredTargetWithObject           // 2
	targetStateSourceReleasingOnTarget                 // 3
	targetStateSourceReleasedOnTarget                  // 4
)

/*
func makeDragNDropTestGUI(textIn *canvas.Text, recIn *canvas.Rectangle, rec2In *canvas.Rectangle, containerIn *fyne.Container) (myCanvasObject fyne.CanvasObject) {

	textRef = textIn
	rectangleRef = recIn
	rectangle2Ref = rec2In
	containerRef = containerIn

	dragFromOneLabel := NewDraggableLabel("No 1")
	dragFromTwoLabel := NewDraggableLabel("No 2.000000")
	dragFromThreeLabel := NewDraggableLabel("No 3..00000000000000000")
	dragFromFourLabel := NewDraggableLabel("No 4.0000000000000000000000000000000")
	dragToDrop1Label := newNoneDroppableLabel("No 5..0000000000000000000000000000000000000")
	dragToDrop2Label := NewDroppableLabel("No 6.000000000000000000000000000000000000000000000000")
	dragToDrop3Label := newNoneDroppableLabel("No 7.00000000000000000000000000000000000000000000000000000000000000")
	dragToDrop4Label := NewDroppableLabel("No 8.00000000000000000000000000000000000000000000000000000000000000000")
	dragToDrop5Label := newNoneDroppableLabel("No 9.0000000000000000000000000000000000000000000000000000000000000000000")
	dragToDrop6Label := NewDroppableLabel("No 10.00000000000000000000000000000000000000000000000000000000000000000000000")

	registeredDroppableTargetLabels = append(registeredDroppableTargetLabels, dragToDrop2Label)
	registeredDroppableTargetLabels = append(registeredDroppableTargetLabels, dragToDrop4Label)
	registeredDroppableTargetLabels = append(registeredDroppableTargetLabels, dragToDrop6Label)

	DropOne := container.NewMax(dragToDrop1Label)
	DropTwo := container.NewMax(dragToDrop2Label.backgroundRectangle, dragToDrop2Label)
	DropThree := container.NewMax(dragToDrop3Label)
	DropFour = container.NewMax(dragToDrop4Label.backgroundRectangle, dragToDrop4Label)
	DropFive := container.NewMax(dragToDrop5Label)
	DropSix := container.NewMax(dragToDrop6Label.backgroundRectangle, dragToDrop6Label)

	labelStandardHeight = dragToDrop2Label.Size().Height

	//DropTwoThin := container.NewGridWrap(fyne.NewSize(100, dragToDrop2Label.Size().Height/2), DropTwo)

	//registeredThinDroppableContainers = append(registeredThinDroppableContainers, DropTwoThin)

	fromContainer := container.NewHBox(dragFromOneLabel, dragFromTwoLabel)
	toContainer := container.NewHBox(dragFromThreeLabel, dragFromFourLabel)
	dropContainer = container.NewVBox(DropOne, DropTwo, DropThree, DropFour, DropFive, DropSix)

	myText := widget.NewLabel("Test Area for Drag n Drop")

	myCanvasObject = container.NewVBox(myText, fromContainer, layout.NewSpacer(), toContainer, dropContainer)

	myCanvasObject.Refresh()

	return myCanvasObject
}

*/

//var DropFour *fyne.Container
//var dropContainer *fyne.Container

// Local variables for the Drag n Drop object
var textRef *canvas.Text
var rectangleRef *canvas.Rectangle
var rectangle2Ref *canvas.Rectangle
var containerRef *fyne.Container
var labelStandardHeight float32

//****************************************************
type DraggableLabel struct {
	widget.Label
	SourceUuid  string
	IsDraggable bool
}

type DroppableLabel struct {
	widget.Label
	targetUuid          string
	backgroundRectangle *canvas.Rectangle
}

type noneDroppableLabel struct {
	widget.Label
	uuid string
}

// InitiateStateStateMachine
// InitiateState State machine
func (stateMachine *StateMachineDragAndDropStruct) InitiateStateStateMachine(dragNDropText *canvas.Text, dragNDropRectangleRef *canvas.Rectangle, dragNDropRectangle2Ref *canvas.Rectangle, dragNDropContainerRef *fyne.Container) {
	textRef = dragNDropText
	rectangleRef = dragNDropRectangleRef
	rectangle2Ref = dragNDropRectangle2Ref
	containerRef = dragNDropContainerRef

}

//****************************************************

func (stateMachine *StateMachineDragAndDropStruct) NewDraggableLabel(uuid string) *DraggableLabel {
	draggableLabel := &DraggableLabel{}
	draggableLabel.ExtendBaseWidget(draggableLabel)

	draggableLabel.SourceUuid = uuid
	draggableLabel.Text = uuid

	return draggableLabel
}

func (stateMachine *StateMachineDragAndDropStruct) NewDroppableLabel(uuid string) *DroppableLabel {
	droppableLabel := &DroppableLabel{}
	droppableLabel.ExtendBaseWidget(droppableLabel)

	droppableLabel.targetUuid = uuid

	droppableLabel.backgroundRectangle = canvas.NewRectangle(color.RGBA{
		R: 0x00,
		G: 0x00,
		B: 0x00,
		A: 0x00,
	})

	droppableLabel.Refresh()
	droppableLabel.backgroundRectangle.SetMinSize(droppableLabel.Size())
	droppableLabel.backgroundRectangle.Hide()

	return droppableLabel
}

func newNoneDroppableLabel(uuid string) *noneDroppableLabel {
	nonDroppableLabel := &noneDroppableLabel{}
	nonDroppableLabel.ExtendBaseWidget(nonDroppableLabel)

	nonDroppableLabel.uuid = uuid
	nonDroppableLabel.Text = uuid

	return nonDroppableLabel
}

//****************************************************

type StateMachineDragAndDropStruct struct {
	sourceStateMachine              stateMachineSourceAndDestinationStruct
	targetStateMachine              stateMachineSourceAndDestinationStruct
	registeredDroppableTargetLabels []*DroppableLabel
	sourceUuid                      string
	target                          DroppableLabel
}

// Structure for 'Drag-part of 'Drag-N-Drop' state machine
type stateMachineSourceAndDestinationStruct struct {
	currentState int
}

var stateMachineDragAndDrop StateMachineDragAndDropStruct

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
		stateMachineDragAndDrop.sourceUuid = t.SourceUuid

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
	textRef.Text = t.SourceUuid

	// Change size of 'Drag N Drop'-object text backgrounds
	rectangleRef.SetMinSize(textRef.Size().Add(fyne.NewSize(40, 40)))
	rectangle2Ref.SetMinSize(textRef.Size())

	// Move 'Drag N Drop'-object container, so it is to the right of the mouse-pointer
	diffPos := fyne.Position{
		X: 10,
		Y: -20,
	}
	newPos := ev.AbsolutePosition.Add(diffPos).Add(fyne.NewSize(rectangleRef.Size().Width/2, rectangleRef.Size().Height/2))
	containerRef.Move(newPos)

	// Refresh 'Drag N Drop'-object and show them
	containerRef.Refresh()
	textRef.Show()
	rectangle2Ref.Show()
	rectangleRef.Show()

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
		stateMachineDragAndDrop.sourceUuid = ""

		shrinkDropAreas()

	case sourceStateReleasingWithOutTarget:
		stateMachineDragAndDrop.sourceUuid = ""
		// Just continue

	case sourceStateEnteringTarget:
		switchStateForSource(sourceStateReleasingOnTarget)
		switchStateForTarget(targetStateSourceReleasingOnTarget)

		shrinkDropAreas()

		for _, droppableTargetLabel := range stateMachineDragAndDrop.registeredDroppableTargetLabels {
			droppableTargetLabel.backgroundRectangle.StrokeWidth = 0
			droppableTargetLabel.backgroundRectangle.StrokeColor = color.RGBA{
				R: 0x00,
				G: 0x00,
				B: 0x00,
				A: 0x00,
			}
			droppableTargetLabel.backgroundRectangle.FillColor = color.RGBA{
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
	textRef.Hide()
	rectangleRef.Hide()
	rectangle2Ref.Hide()
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

// ***** The Object from the Drop Ends *****

// MouseIn is called when a desktop pointer enters the widget
func (b *DroppableLabel) MouseIn(*desktop.MouseEvent) {

	switch stateMachineDragAndDrop.targetStateMachine.currentState {

	case targetStateWaitingForSourceToEnteringTarget:
		return

	case targetStateSourceIsDraggingObject:
		switchStateForSource(sourceStateEnteringTarget)
		switchStateForTarget(targetStateSourceEnteredTargetWithObject)
		b.backgroundRectangle.FillColor = color.RGBA{
			R: 0x33,
			G: 0x33,
			B: 0x33,
			A: 0x22,
		}

		b.backgroundRectangle.Show()
		b.backgroundRectangle.Refresh()

		stateMachineDragAndDrop.target = *b

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
		b.backgroundRectangle.FillColor = color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0x00,
		}
		//b.backgroundRectangle.Hide()
		b.backgroundRectangle.Refresh()

	case targetStateSourceReleasingOnTarget:
		return

	case targetStateSourceReleasedOnTarget:
		return

	default:
		log.Fatalln("Unhandled state for StateMachine(From): ", stateMachineDragAndDrop.targetStateMachine.currentState)

	}

}

func switchStateForSource(newState int) {
	stateMachineDragAndDrop.sourceStateMachine.currentState = newState
}

func switchStateForTarget(newState int) {
	stateMachineDragAndDrop.targetStateMachine.currentState = newState
}

func expandDropAreas() {
	for _, targetLabel := range stateMachineDragAndDrop.registeredDroppableTargetLabels {

		targetLabel.backgroundRectangle.StrokeWidth = 2

		targetLabel.backgroundRectangle.Show()
		go func(targetReferenceLabel *DroppableLabel) {
			rectangleColorAnimation := canvas.NewColorRGBAAnimation(color.RGBA{
				R: 0x00,
				G: 0x00,
				B: 0x00,
				A: 0x00,
			}, color.RGBA{
				R: 0xFF,
				G: 0x00,
				B: 0x00,
				A: 0xAA,
			}, time.Millisecond*300, func(c color.Color) {
				targetReferenceLabel.backgroundRectangle.StrokeColor = c
				canvas.Refresh(targetReferenceLabel.backgroundRectangle)
			})

			rectangleSizeAnimation := canvas.NewSizeAnimation(
				fyne.NewSize(targetReferenceLabel.backgroundRectangle.Size().Width, 0),
				fyne.NewSize(targetReferenceLabel.backgroundRectangle.Size().Width, labelStandardHeight),
				time.Millisecond*300,
				func(animationSize fyne.Size) {
					targetReferenceLabel.backgroundRectangle.SetMinSize(animationSize)
					canvas.Refresh(targetReferenceLabel.backgroundRectangle)
					//canvas.Refresh(DropFour)
					//canvas.Refresh(dropContainer)
				})

			rectangleColorAnimation.Start()
			rectangleSizeAnimation.Start()
		}(targetLabel)

	}

	go func() {
		time.Sleep(400 * time.Millisecond)
		for _, targetLabel := range stateMachineDragAndDrop.registeredDroppableTargetLabels {
			targetLabel.Show()
		}
	}()
}

func shrinkDropAreas() {
	for _, targetLabel := range stateMachineDragAndDrop.registeredDroppableTargetLabels {

		targetLabel.Hide()

		targetLabel.backgroundRectangle.StrokeWidth = 2
		go func(targetReferenceLabel *DroppableLabel) {
			rectangleColorAnimation := canvas.NewColorRGBAAnimation(color.RGBA{
				R: 0x00,
				G: 0x00,
				B: 0x00,
				A: 0x00,
			}, color.RGBA{
				R: 0xFF,
				G: 0x00,
				B: 0x00,
				A: 0xAA,
			}, time.Millisecond*300, func(c color.Color) {
				targetReferenceLabel.backgroundRectangle.StrokeColor = c
				canvas.Refresh(targetReferenceLabel.backgroundRectangle)
			})

			rectangleSizeAnimation := canvas.NewSizeAnimation(
				fyne.NewSize(targetReferenceLabel.backgroundRectangle.Size().Width, labelStandardHeight),
				fyne.NewSize(targetReferenceLabel.backgroundRectangle.Size().Width, 0),
				time.Millisecond*300,
				func(animationSize fyne.Size) {
					targetReferenceLabel.backgroundRectangle.SetMinSize(animationSize)
					canvas.Refresh(targetReferenceLabel.backgroundRectangle)
					//canvas.Refresh(DropFour)
					//canvas.Refresh(dropContainer)
				})

			rectangleColorAnimation.Start()
			rectangleSizeAnimation.Start()
		}(targetLabel)

	}

	go func() {
		time.Sleep(400 * time.Millisecond)
		for _, targetLabel := range stateMachineDragAndDrop.registeredDroppableTargetLabels {
			targetLabel.backgroundRectangle.Hide()
			targetLabel.backgroundRectangle.Refresh()

		}
	}()
}

func executeDropAction() {
	fmt.Println(fmt.Sprintf("'%s' was droppen in '%s'", stateMachineDragAndDrop.sourceUuid, stateMachineDragAndDrop.target.targetUuid))
}
