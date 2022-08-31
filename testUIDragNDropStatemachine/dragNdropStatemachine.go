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
	DropTwo := container.NewMax(dragToDrop2Label.BackgroundRectangle, dragToDrop2Label)
	DropThree := container.NewMax(dragToDrop3Label)
	DropFour = container.NewMax(dragToDrop4Label.BackgroundRectangle, dragToDrop4Label)
	DropFive := container.NewMax(dragToDrop5Label)
	DropSix := container.NewMax(dragToDrop6Label.BackgroundRectangle, dragToDrop6Label)

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
	SourceUuid        string
	IsDraggable       bool
	BuildingBlockType int
}

type DroppableLabel struct {
	widget.Label
	topTestCaseAccordion      *widget.Accordion
	parrentAccordion          *widget.Accordion
	TargetUuid                string
	BackgroundRectangle       *canvas.Rectangle
	IsDroppable               bool
	labelStandardHeight       float32
	nodeLevel                 float32
	testCaseNodeRectangleSize int
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

func (stateMachine *StateMachineDragAndDropStruct) NewDroppableLabel(uuid string, accordionReference *widget.Accordion, nodeLevel float32, testCaseNodeRectangleSize int, topTestCaseAccordion *widget.Accordion) *DroppableLabel {
	droppableLabel := &DroppableLabel{}
	droppableLabel.ExtendBaseWidget(droppableLabel)

	droppableLabel.topTestCaseAccordion = topTestCaseAccordion
	droppableLabel.parrentAccordion = accordionReference
	droppableLabel.TargetUuid = uuid
	droppableLabel.Text = uuid
	droppableLabel.nodeLevel = nodeLevel
	droppableLabel.testCaseNodeRectangleSize = testCaseNodeRectangleSize

	droppableLabel.BackgroundRectangle = canvas.NewRectangle(color.RGBA{
		R: 0x00,
		G: 0x00,
		B: 0x00,
		A: 0x00,
	})

	droppableLabel.Refresh()
	droppableLabel.BackgroundRectangle.SetMinSize(droppableLabel.Size())
	droppableLabel.BackgroundRectangle.Hide()

	stateMachineDragAndDrop.registeredDroppableTargetLabels = append(stateMachineDragAndDrop.registeredDroppableTargetLabels, droppableLabel)

	droppableLabel.labelStandardHeight = droppableLabel.MinSize().Height

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
	SourceUuid                      string
	SourceType                      int
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

		// Verify if this Draggable component can be dropped on this Element

		//if b.IsDroppable == true {
		switchStateForSource(sourceStateEnteringTarget)
		switchStateForTarget(targetStateSourceEnteredTargetWithObject)
		b.BackgroundRectangle.FillColor = color.RGBA{
			R: 0x33,
			G: 0x33,
			B: 0x33,
			A: 0x22,
		}

		b.BackgroundRectangle.Show()
		b.BackgroundRectangle.Refresh()

		stateMachineDragAndDrop.target = *b
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
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0x00,
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

func switchStateForSource(newState int) {
	stateMachineDragAndDrop.sourceStateMachine.currentState = newState
}

func switchStateForTarget(newState int) {
	stateMachineDragAndDrop.targetStateMachine.currentState = newState
}

func expandDropAreas() {
	for _, targetLabel := range stateMachineDragAndDrop.registeredDroppableTargetLabels {

		targetLabel.BackgroundRectangle.StrokeWidth = 2

		targetLabel.BackgroundRectangle.Show()
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
				targetReferenceLabel.BackgroundRectangle.StrokeColor = c
				canvas.Refresh(targetReferenceLabel.BackgroundRectangle)
			})

			rectangleWidth := float32(500)

			rectangleSizeAnimation := canvas.NewSizeAnimation(
				fyne.NewSize(rectangleWidth, 0),
				fyne.NewSize(rectangleWidth, targetReferenceLabel.labelStandardHeight),
				time.Millisecond*300,
				func(animationSize fyne.Size) {
					targetReferenceLabel.BackgroundRectangle.SetMinSize(animationSize)
					canvas.Refresh(targetReferenceLabel.BackgroundRectangle)
					targetReferenceLabel.parrentAccordion.Refresh()
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

		targetLabel.BackgroundRectangle.StrokeWidth = 2
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
				targetReferenceLabel.BackgroundRectangle.StrokeColor = c
				canvas.Refresh(targetReferenceLabel.BackgroundRectangle)
			})

			rectangleWidth := float32(500)

			rectangleSizeAnimation := canvas.NewSizeAnimation(
				fyne.NewSize(rectangleWidth, targetReferenceLabel.labelStandardHeight),
				fyne.NewSize(rectangleWidth, 0),
				time.Millisecond*300,
				func(animationSize fyne.Size) {
					targetReferenceLabel.BackgroundRectangle.SetMinSize(animationSize)
					canvas.Refresh(targetReferenceLabel.BackgroundRectangle)
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
			targetLabel.BackgroundRectangle.Hide()
			targetLabel.BackgroundRectangle.Refresh()

		}
	}()
}

func executeDropAction() {
	fmt.Println(fmt.Sprintf("'%s' was droppen in '%s'", stateMachineDragAndDrop.SourceUuid, stateMachineDragAndDrop.target.TargetUuid))
}

//func RegisterDroppableLable()