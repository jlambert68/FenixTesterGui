package testCaseUI

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
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

func makeDragNDropTestGUI(textIn *canvas.Text, recIn *canvas.Rectangle, rec2In *canvas.Rectangle, containerIn *fyne.Container) (myCanvasObject fyne.CanvasObject) {

	textRef = textIn
	rectangleRef = recIn
	rectangle2Ref = rec2In
	containerRef = containerIn

	dragFromOneLabel := newDraggableLabel("No 1")
	dragFromTwoLabel := newDraggableLabel("No 2.000000")
	dragFromThreeLabel := newDraggableLabel("No 3..00000000000000000")
	dragFromFourLabel := newDraggableLabel("No 4.0000000000000000000000000000000")
	dragToDrop1Label := newNoneDroppableLabel("No 5..0000000000000000000000000000000000000")
	dragToDrop2Label := newDroppableLabel("No 6.000000000000000000000000000000000000000000000000")
	dragToDrop3Label := newNoneDroppableLabel("No 7.00000000000000000000000000000000000000000000000000000000000000")
	dragToDrop4Label := newDroppableLabel("No 8.00000000000000000000000000000000000000000000000000000000000000000")
	dragToDrop5Label := newNoneDroppableLabel("No 9.0000000000000000000000000000000000000000000000000000000000000000000")
	dragToDrop6Label := newDroppableLabel("No 10.00000000000000000000000000000000000000000000000000000000000000000000000")

	registeredDroppableTargetLabels = append(registeredDroppableTargetLabels, dragToDrop2Label)
	registeredDroppableTargetLabels = append(registeredDroppableTargetLabels, dragToDrop4Label)
	registeredDroppableTargetLabels = append(registeredDroppableTargetLabels, dragToDrop6Label)

	registeredNoneDroppableTargetLabels = append(registeredNoneDroppableTargetLabels, dragToDrop1Label)
	registeredNoneDroppableTargetLabels = append(registeredNoneDroppableTargetLabels, dragToDrop3Label)
	registeredNoneDroppableTargetLabels = append(registeredNoneDroppableTargetLabels, dragToDrop5Label)

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

	//*********************************

	dragToDropAccordianLabel1 := newDroppableLabel("Accordian 1")
	dragToDropAccordianLabel2 := newDroppableLabel("Accordian 2")
	dragToDropAccordianLabel3 := newDroppableLabel("Accordian 3")
	Accordian1 := container.NewMax(dragToDropAccordianLabel1.backgroundRectangle, dragToDropAccordianLabel1)
	Accordian2 := container.NewMax(dragToDropAccordianLabel2.backgroundRectangle, dragToDropAccordianLabel2)
	Accordian3 := container.NewMax(dragToDropAccordianLabel3.backgroundRectangle, dragToDropAccordianLabel3)

	registeredDroppableTargetLabels = append(registeredDroppableTargetLabels, dragToDropAccordianLabel1)
	registeredDroppableTargetLabels = append(registeredDroppableTargetLabels, dragToDropAccordianLabel2)
	registeredDroppableTargetLabels = append(registeredDroppableTargetLabels, dragToDropAccordianLabel3)

	textualRepresentationGrid2 := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Extended"),
		Accordian3)

	// Create GUI Canvas object to be used
	testCaseTextualModelArea2 := container.NewVBox(textualRepresentationGrid2)

	// Create a Canvas Accordion type for grouping the Textual Representations
	testCaseTextualModelAreaAccordionItem2 := widget.NewAccordionItem("Textual Representation of the TestCase 2", testCaseTextualModelArea2)

	testCaseTextualModelAreaAccordion2 := widget.NewAccordion(testCaseTextualModelAreaAccordionItem2)
	myrect := canvas.NewRectangle(color.White)
	myrect.SetMinSize(fyne.NewSize(20, 20))
	testCaseTextualModelAreaAccordion2Area := container.NewHBox(myrect, testCaseTextualModelAreaAccordion2)

	textualRepresentationGrid := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Simple"),
		Accordian1,
		widget.NewLabel("Complex"),
		Accordian2,
		testCaseTextualModelAreaAccordion2Area,
	)

	// Create GUI Canvas object to be used
	testCaseTextualModelArea := container.NewVBox(textualRepresentationGrid)

	// Create a Canvas Accordion type for grouping the Textual Representations
	testCaseTextualModelAreaAccordionItem := widget.NewAccordionItem("Textual Representation of the TestCase", testCaseTextualModelArea)

	testCaseTextualModelAreaAccordion := widget.NewAccordion(testCaseTextualModelAreaAccordionItem)

	canvasTextualRepresentationAccordionObject := container.NewVBox(testCaseTextualModelAreaAccordion)

	//*********************************

	myCanvasObject = container.NewVBox(myText, fromContainer, layout.NewSpacer(), toContainer, dropContainer, canvasTextualRepresentationAccordionObject)

	myCanvasObject.Refresh()
	/*
		containerMinSize := myCanvasObject.MinSize()
		thinHeight := dragToDrop2Label.Size().Height / 2

		for _, thinContainer := range registeredThinDroppableContainers {
			thinContainer.Resize(fyne.NewSize(containerMinSize.Width, thinHeight))
		}

		dragToDrop2Label.backgroundRectangle.Resize(fyne.NewSize(containerMinSize.Width, thinHeight))

		dragToDrop2Label.Refresh()


	*/
	return myCanvasObject
}

var DropFour *fyne.Container
var dropContainer *fyne.Container

var registeredDroppableTargetLabels []*droppableLabel
var registeredNoneDroppableTargetLabels []*noneDroppableLabel
var registeredThinDroppableContainers []*fyne.Container

var textRef *canvas.Text
var rectangleRef *canvas.Rectangle
var rectangle2Ref *canvas.Rectangle
var containerRef *fyne.Container
var labelStandardHeight float32

//****************************************************
type draggableLabel struct {
	widget.Label
	sourceUuid string
}

type droppableLabel struct {
	widget.Label
	targetUuid          string
	backgroundRectangle *canvas.Rectangle
}

type noneDroppableLabel struct {
	widget.Label
	uuid string
}

//****************************************************

func newDraggableLabel(uuid string) *draggableLabel {
	draggableLabel := &draggableLabel{}
	draggableLabel.ExtendBaseWidget(draggableLabel)

	draggableLabel.sourceUuid = uuid
	draggableLabel.Text = uuid

	return draggableLabel
}

func newDroppableLabel(uuid string) *droppableLabel {
	droppableLabel := &droppableLabel{}
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

type stateMachineDragAndDropStruct struct {
	sourceStateMachine StateMachineSouceAndDestinationStruct
	targetStateMachine StateMachineSouceAndDestinationStruct
	sourceUuid         string
	target             droppableLabel
}

// Structure for 'Drag-part of 'Drag-N-Drop' state machine
type StateMachineSouceAndDestinationStruct struct {
	currentState int
	//target       droppableLabel
	//currentUuid  string
}

var stateMachineDragAndDrop stateMachineDragAndDropStruct
var stateMachineDragFrom StateMachineSouceAndDestinationStruct
var stateMachineTarget StateMachineSouceAndDestinationStruct

// ***** The Object from the Drag starts *****

// Dragged
// When the user press down the mouse button this event is triggered
func (t *draggableLabel) Dragged(ev *fyne.DragEvent) {

	switch stateMachineDragFrom.currentState {

	case sourceStateSearching:
		return

	case sourceStateFinds:
		// switch state to 'sourceStateGrabs'
		switchStateForSource(sourceStateGrabs)

		return

	case sourceStateGrabs:
		// switch state to 'sourceStateDragging'
		switchStateForSource(sourceStateDragging)
		stateMachineDragAndDrop.sourceUuid = t.sourceUuid

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
		log.Fatalln("Unhandled state for StateMachine(From): ", stateMachineDragFrom.currentState)

	}

	// Change Text of 'Drag N Drop'-object
	textRef.Text = t.sourceUuid

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
func (t *draggableLabel) DragEnd() {

	switch stateMachineDragFrom.currentState {

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

		for _, droppableTargetLabel := range registeredDroppableTargetLabels {
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
		log.Fatalln("Unhandled state for StateMachine(From): ", stateMachineDragFrom.currentState)

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
func (b *draggableLabel) MouseIn(*desktop.MouseEvent) {

	switch stateMachineDragFrom.currentState {

	case sourceStateSearching:
		// Mouse finds draggable object
		switchStateForSource(sourceStateFinds)

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
		log.Fatalln("Unhandled state for StateMachine(From): ", stateMachineDragFrom.currentState)

	}

}

// MouseMoved is called when a desktop pointer hovers over the widget
func (b *draggableLabel) MouseMoved(a *desktop.MouseEvent) {

}

// MouseOut is called when a desktop pointer exits the widget
func (b *draggableLabel) MouseOut() {

	switch stateMachineDragFrom.currentState {

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
		log.Fatalln("Unhandled state for StateMachine(From): ", stateMachineDragFrom.currentState)

	}

}

// ***** The Object from the Drop Ends *****

// MouseIn is called when a desktop pointer enters the widget
func (b *droppableLabel) MouseIn(*desktop.MouseEvent) {

	switch stateMachineTarget.currentState {

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
		log.Fatalln("Unhandled state for StateMachine(From): ", stateMachineDragFrom.currentState)

	}

}

// MouseMoved is called when a desktop pointer hovers over the widget
func (b *droppableLabel) MouseMoved(a *desktop.MouseEvent) {

}

// MouseOut is called when a desktop pointer exits the widget
func (b *droppableLabel) MouseOut() {

	switch stateMachineTarget.currentState {

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
		log.Fatalln("Unhandled state for StateMachine(From): ", stateMachineDragFrom.currentState)

	}

}

func switchStateForSource(newState int) {
	stateMachineDragFrom.currentState = newState
}

func switchStateForTarget(newState int) {
	stateMachineTarget.currentState = newState
}

func expandDropAreas() {
	for _, targetLabel := range registeredDroppableTargetLabels {

		targetLabel.backgroundRectangle.StrokeWidth = 2
		/*targetLabel.backgroundRectangle.StrokeColor = color.RGBA{
			R: 0xFF,
			G: 0x00,
			B: 0x00,
			A: 0xAA,
		}

		*/
		targetLabel.backgroundRectangle.Show()
		go func(targetReferenceLabel *droppableLabel) {
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
					canvas.Refresh(DropFour)
					canvas.Refresh(dropContainer)
				})

			rectangleColorAnimation.Start()
			rectangleSizeAnimation.Start()
		}(targetLabel)

		//targetLabel.backgroundRectangle.Refresh()
	}

	go func() {
		time.Sleep(400 * time.Millisecond)
		for _, targetLabel := range registeredDroppableTargetLabels {
			targetLabel.Show() // *** NEW ***
		}
	}()
}

func shrinkDropAreas() {
	for _, targetLabel := range registeredDroppableTargetLabels {

		targetLabel.Hide()

		targetLabel.backgroundRectangle.StrokeWidth = 2
		go func(targetReferenceLabel *droppableLabel) {
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
		for _, targetLabel := range registeredDroppableTargetLabels {
			targetLabel.backgroundRectangle.Hide()
			targetLabel.backgroundRectangle.Refresh()

		}
	}()
}

func executeDropAction() {
	fmt.Println(fmt.Sprintf("'%s' was droppen in '%s'", stateMachineDragAndDrop.sourceUuid, stateMachineDragAndDrop.target.targetUuid))
}
