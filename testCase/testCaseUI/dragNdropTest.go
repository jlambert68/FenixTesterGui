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
)

// Statea for handling Drag from Source object
const (
	stateFromSearching = iota
	stateFromFinds
	stateFromGrabs
	stateFromDragging
	stateFromReleasingWithOutTarget
	stateFromEnteringTarget
	stateFromReleasingOnTarget
	stateReleasedOnTarget
)

// State for handling Drop-target object
const (
	stateWaitingForSenderToEnteringTarget = iota
	stateSenderIsDraggingObject
	stateSenderEnteringTarget
	stateSenderReleasingOnTarget
	stateSenderReleasedOnTarget
)

func makeDragNDropTestGUI(textIn *canvas.Text, recIn *canvas.Rectangle, rec2In *canvas.Rectangle, containerIn *fyne.Container) (myCanvasObject fyne.CanvasObject) {

	textRef = textIn
	rectangleRef = recIn
	rectangle2Ref = rec2In
	containerRef = containerIn

	dragFromOneRectangle := newDraggablelabel(color.Gray{0xEE}, "No 1")
	dragFromTwoRectangle := newDraggablelabel(color.Gray{0xBB}, "No 2.000000")
	dragToOneRectangle := newDraggablelabel(color.Gray{0x88}, "No 3..00000000000000000")
	dragToTwoRectangle := newDraggablelabel(color.Gray{0x44}, "No 4.0000000000000000000000000000000")
	dragToDrop1Rectangle := newDropablelabel(color.Gray{0x33}, "No 5..0000000000000000000000000000000000000")
	dragToDrop2Rectangle := newDropablelabel(color.Gray{0x44}, "No 6.000000000000000000000000000000000000000000000000")

	fromContainer := container.NewHBox(dragFromOneRectangle, dragFromTwoRectangle)
	toContainer := container.NewHBox(dragToOneRectangle, dragToTwoRectangle)
	dropContainer := container.NewHBox(dragToDrop1Rectangle, dragToDrop2Rectangle)

	myText := widget.NewLabel("Test Area for Drag n Drop")

	myCanvasObject = container.NewVBox(myText, fromContainer, layout.NewSpacer(), toContainer, dropContainer)

	return myCanvasObject
}

var textRef *canvas.Text
var rectangleRef *canvas.Rectangle
var rectangle2Ref *canvas.Rectangle
var containerRef *fyne.Container

func newDraggablelabel(myColor color.Gray, myNewTitle string) *draggableLabel {
	myLabel := &draggableLabel{}
	myLabel.ExtendBaseWidget(myLabel)
	//myLabel.FillColor = myColor
	//myLabel.StrokeColor = color.Black
	//myLabel.StrokeWidth = 0
	//myLabel.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize), float32(testCaseNodeRectangleSize)))

	myLabel.myTitle = myNewTitle
	myLabel.Text = myNewTitle

	return myLabel
}

func newDropablelabel(myColor color.Gray, myNewTitle string) *droppableLabel {
	myLabel := &droppableLabel{}
	myLabel.ExtendBaseWidget(myLabel)
	//myLabel.FillColor = myColor
	//myLabel.StrokeColor = color.Black
	//myLabel.StrokeWidth = 0
	//myLabel.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize), float32(testCaseNodeRectangleSize)))

	myLabel.myTitle = myNewTitle
	myLabel.Text = myNewTitle

	return myLabel
}

type draggableLabel struct {
	widget.Label
	myTitle string
	hovered bool
}

type droppableLabel struct {
	widget.Label
	myTitle string
	hovered bool
}

// Structure for 'Drag-part of 'Drag-N-Drop' state machine
type StateMachineStruct struct {
	currentState int
	/*
		dragBeforeMouseIn      bool
		mouseInBeforeDragStart bool
		mouseInObjectUuid      string
		dragStart              bool
		dragStartObjectUuid    string
		dragEnd                bool
		MouseIn                bool
		MouseOut               bool
		LightRectangle         bool
		preLightRectangleState bool
		dropIsAllowed          bool
		dragFrom               string
		DropIn                 string
	*/
}

var stateMachineDragFrom StateMachineStruct
var stateMachineTarget StateMachineStruct

// ***** The Object from the Drag starts *****

// Dragged
// When the user press down the mouse button this event is triggered
func (t *draggableLabel) Dragged(ev *fyne.DragEvent) {

	switch stateMachineDragFrom.currentState {

	case stateFromSearching:
		fmt.Println("Dragged: 'stateFromSearching'")
		return

	case stateFromFinds:
		// switch state to 'stateFromGrabs'
		switchStateForFrom(stateFromGrabs)

		return

	case stateFromGrabs:
		// switch state to 'stateFromDragging'
		switchStateForFrom(stateFromDragging)

		return

	case stateFromDragging:
		switchStateForTarget(stateSenderIsDraggingObject)
		// Just continue

	case stateFromReleasingWithOutTarget:
		fmt.Println("Dragged: 'stateFromReleasingWithOutTarget'")
		return

	case stateFromEnteringTarget:
		fmt.Println("Dragged: 'stateFromEnteringTarget'")
		return

	case stateFromReleasingOnTarget:
		fmt.Println("Dragged: 'stateFromReleasingOnTarget'")
		return

	case stateReleasedOnTarget:
		fmt.Println("Dragged: 'stateReleasedOnTarget'")
		return

	default:
		log.Fatalln("Unhandled state for StateMachine(From): ", stateMachineDragFrom.currentState)

	}

	// Change Text of 'Drag N Drop'-object
	textRef.Text = t.myTitle

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

	case stateFromSearching:
		fmt.Println("Dragged: 'stateFromSearching'")
		return

	case stateFromFinds:
		fmt.Println("Dragged: 'stateFromFinds'")
		return

	case stateFromGrabs:
		fmt.Println("Dragged: 'stateFromGrabs'")
		return

	case stateFromDragging:
		// switch state to 'stateFromReleasingWithOutTarget'
		switchStateForFrom(stateFromReleasingWithOutTarget)
		switchStateForTarget(stateWaitingForSenderToEnteringTarget)

	case stateFromReleasingWithOutTarget:
		// Just continue

	case stateFromEnteringTarget:
		fmt.Println("Dragged: 'stateFromEnteringTarget'")
		return

	case stateFromReleasingOnTarget:
		fmt.Println("Dragged: 'stateFromReleasingOnTarget'")
		return

	case stateReleasedOnTarget:
		fmt.Println("Dragged: 'stateReleasedOnTarget'")
		return

	default:
		log.Fatalln("Unhandled state for StateMachine(From): ", stateMachineDragFrom.currentState)

	}

	// Hide the 'Drag N Drop'-objects
	textRef.Hide()
	rectangleRef.Hide()
	rectangle2Ref.Hide()
	containerRef.Refresh()

	// switch state to 'stateFromSearching'
	switchStateForFrom(stateFromSearching)

}

// MouseIn is called when a desktop pointer enters the widget
func (b *draggableLabel) MouseIn(*desktop.MouseEvent) {

	switch stateMachineDragFrom.currentState {

	case stateFromSearching:
		// switch state to 'stateFromFinds'
		switchStateForFrom(stateFromFinds)
		return

	case stateFromFinds:
		fmt.Println("MouseIn: 'stateFromFinds'")
		return

	case stateFromGrabs:
		fmt.Println("MouseIn: 'stateFromGrabs'")
		return

	case stateFromDragging:
		fmt.Println("MouseIn: 'stateFromDragging'")
		return

	case stateFromReleasingWithOutTarget:
		fmt.Println("MouseIn: 'stateFromReleasingWithOutTarget'")
		return

	case stateFromEnteringTarget:
		fmt.Println("MouseIn: 'stateFromEnteringTarget'")
		return

	case stateFromReleasingOnTarget:
		fmt.Println("MouseIn: 'stateFromReleasingOnTarget'")
		return

	case stateReleasedOnTarget:
		fmt.Println("MouseIn: 'stateReleasedOnTarget'")
		return

	default:
		log.Fatalln("Unhandled state for StateMachine(From): ", stateMachineDragFrom.currentState)

	}

}

// MouseMoved is called when a desktop pointer hovers over the widget
func (b *draggableLabel) MouseMoved(a *desktop.MouseEvent) {
	//log.Println("I have been 'MouseMoved' ", b.myTitle)

}

// MouseOut is called when a desktop pointer exits the widget
func (b *draggableLabel) MouseOut() {

	switch stateMachineDragFrom.currentState {

	case stateFromSearching:
		fmt.Println("MouseOut: 'stateFromSearching'")
		return

	case stateFromFinds:
		// switch state to 'stateFromFinds'
		switchStateForFrom(stateFromSearching)
		return

	case stateFromGrabs:
		fmt.Println("MouseOut: 'stateFromGrabs'")
		return

	case stateFromDragging:
		fmt.Println("MouseOut: 'stateFromDragging'")
		return

	case stateFromReleasingWithOutTarget:
		fmt.Println("MouseOut: 'stateFromReleasingWithOutTarget'")
		return

	case stateFromEnteringTarget:
		fmt.Println("MouseOut: 'stateFromEnteringTarget'")
		return

	case stateFromReleasingOnTarget:
		fmt.Println("MouseOut: 'stateFromReleasingOnTarget'")
		return

	case stateReleasedOnTarget:
		fmt.Println("MouseOut: 'stateReleasedOnTarget'")
		return

	default:
		log.Fatalln("Unhandled state for StateMachine(From): ", stateMachineDragFrom.currentState)

	}

}

// ***** The Object from the Drop Ends *****

// MouseIn is called when a desktop pointer enters the widget
func (b *droppableLabel) MouseIn(*desktop.MouseEvent) {

	switch stateMachineTarget.currentState {

	case stateWaitingForSenderToEnteringTarget:
		fmt.Println("MouseOut: 'stateWaitingForSenderToEnteringTarget'")
		return

	case stateSenderIsDraggingObject:
		// switch state to 'stateSenderIsDraggingObject'
		switchStateForFrom(stateSenderEnteringTarget)
		return

	case stateSenderEnteringTarget:
		switchStateForFrom(stateFromEnteringTarget)

	case stateSenderReleasingOnTarget:
		fmt.Println("MouseOut: 'stateSenderReleasingOnTarget'")
		return

	case stateSenderReleasedOnTarget:
		fmt.Println("MouseOut: 'stateFromReleasingWithOutTarget'")
		return

	default:
		log.Fatalln("Unhandled state for StateMachine(From): ", stateMachineDragFrom.currentState)

	}

}

// MouseMoved is called when a desktop pointer hovers over the widget
func (b *droppableLabel) MouseMoved(a *desktop.MouseEvent) {
	//log.Println("I have been 'MouseMoved' ", b.myTitle)

}

// MouseOut is called when a desktop pointer exits the widget
func (b *droppableLabel) MouseOut() {

}

/*
func dragNDropStatMachineSwitchStateWhenMouseInOnDragFrom(labelNo string) {
	fmt.Println("dragNDropStatMachineSwitchStateWhenMouseInOnDragFrom ", labelNo)
	// Change Mouse-state

	if stateMachineDragFrom.dragStart == true {
		stateMachineDragFrom.dragBeforeMouseIn = true
		return
	}

	stateMachineDragFrom.MouseIn = true
	stateMachineDragFrom.MouseOut = false
	stateMachineDragFrom.mouseInObjectUuid = labelNo

	// Trigger State Machine
	stateMachineExecuteEngineDragAndDrop(true)
}

func dragNDropStatMachineSwitchStateWhenMouseOutOnDragFrom(labelNo string) {
	fmt.Println("dragNDropStatMachineSwitchStateWhenMouseOutOnDragFrom ", labelNo)

	// Change Mouse-state
	if stateMachineDragFrom.dragStart == true {
		return
	}
	stateMachineDragFrom.MouseIn = false
	stateMachineDragFrom.MouseOut = true
	stateMachineDragFrom.mouseInObjectUuid = ""

	// Trigger State Machine
	stateMachineExecuteEngineDragAndDrop(true)
}

func dragNDropStatMachineSwitchStateWhenDragStart(labelNo string) {
	// Change Drag-state

	if stateMachineDragFrom.dragBeforeMouseIn == true {
		return
	}
	stateMachineDragFrom.dragStart = true
	stateMachineDragFrom.dragEnd = false

	// Trigger State Machine
	stateMachineExecuteEngineDragAndDrop(true)
}

func dragNDropStatMachineSwitchStateWhenDragEnd(labelNo string) {
	// Change Drag-state

	stateMachineDragFrom.dragBeforeMouseIn = false

	if stateMachineDragFrom.MouseIn == true {
		return
	}
	stateMachineDragFrom.dragStart = false
	stateMachineDragFrom.dragEnd = true
	stateMachineDragFrom.mouseInObjectUuid = ""

	// Trigger State Machine
	stateMachineExecuteEngineDragAndDrop(true)
}

func dragNDropStatMachineSwitchStateWhenMouseIn(labelNo string) {
	fmt.Println("dragNDropStatMachineSwitchStateWhenMouseIn ", labelNo)

	// Change Mouse-state
	stateMachineDragFrom.MouseIn = true
	stateMachineDragFrom.MouseOut = false

	stateMachineDragFrom.DropIn = labelNo

	// Trigger State Machine
	stateMachineExecuteEngineDragAndDrop(true)
}

func dragNDropStatMachineSwitchStateWhenMouseOut(labelNo string) {
	fmt.Println("dragNDropStatMachineSwitchStateWhenMouseOut ", labelNo)
	// Change Mouse-state
	stateMachineDragFrom.MouseIn = false
	stateMachineDragFrom.MouseOut = true

	// Trigger State Machine
	stateMachineExecuteEngineDragAndDrop(true)
}


*/
/*
func stateMachineExecuteEngineDragAndDrop(doAction bool) {

	//fmt.Println(stateMachineDragFrom.dragStart, stateMachineDragFrom.dragEnd, stateMachineDragFrom.MouseIn, stateMachineDragFrom.MouseOut, stateMachineDragFrom.LightRectangle, stateMachineDragFrom.preLightRectangleState, stateMachineDragFrom.LightRectangle, stateMachineDragFrom.dropIsAllowed)

	stateMachineDragFrom.preLightRectangleState = stateMachineDragFrom.LightRectangle

	// Logic for light up Droppable Areas when hover over them with "something"
	if stateMachineDragFrom.dragStart == true &&
		stateMachineDragFrom.dragEnd == false &&
		stateMachineDragFrom.MouseIn == true &&
		stateMachineDragFrom.MouseOut == false &&
		stateMachineDragFrom.DropIn != stateMachineDragFrom.dragFrom {

		stateMachineDragFrom.LightRectangle = true
		if doAction && stateMachineDragFrom.preLightRectangleState == false {
			fmt.Println("Light up rectangle no ", stateMachineDragFrom.DropIn)
		}

	} else {
		stateMachineDragFrom.LightRectangle = false
		if doAction && stateMachineDragFrom.preLightRectangleState == true {
			fmt.Println("Turn of Light for rectangle no ", stateMachineDragFrom.DropIn)
		}
	}

	// Logic for dropping something that has been dragged
	if stateMachineDragFrom.dragStart == false &&
		stateMachineDragFrom.dragEnd == true &&
		stateMachineDragFrom.MouseIn == true &&
		stateMachineDragFrom.MouseOut == false &&
		stateMachineDragFrom.preLightRectangleState == true &&
		stateMachineDragFrom.DropIn != stateMachineDragFrom.dragFrom {

		stateMachineDragFrom.dropIsAllowed = true

		if doAction {
			fmt.Println("Drag rectangle no ", stateMachineDragFrom.dragFrom, " and drop in rectangle no ", stateMachineDragFrom.DropIn)
		}

	} else {
		stateMachineDragFrom.dropIsAllowed = false
	}

	//fmt.Println(stateMachineDragFrom.dragStart, stateMachineDragFrom.dragEnd, stateMachineDragFrom.MouseIn, stateMachineDragFrom.MouseOut, stateMachineDragFrom.LightRectangle, stateMachineDragFrom.preLightRectangleState, stateMachineDragFrom.LightRectangle, stateMachineDragFrom.dropIsAllowed)
}


*/
func switchStateForFrom(newState int) {
	stateMachineDragFrom.currentState = newState
}

func switchStateForTarget(newState int) {
	stateMachineTarget.currentState = newState
}

/*
func executeDragObject() {
	if stateMachineDragFrom.mouseInBeforeDragStart == false {
		return
	}

	dragNDropStatMachineSwitchStateWhenDragStart(t.myTitle)
	//fmt.Println(ev.Position, ev.AbsolutePosition)

	// Change Text of 'Drag N Drop'-object
	textRef.Text = t.myTitle

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

*/
