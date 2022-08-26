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

	dragFromOneLabel := newDraggablelabel(color.Gray{0xEE}, "No 1")
	dragFromTwoLabel := newDraggablelabel(color.Gray{0xBB}, "No 2.000000")
	dragFromThreeLabel := newDraggablelabel(color.Gray{0x88}, "No 3..00000000000000000")
	dragFromFourLabel := newDraggablelabel(color.Gray{0x44}, "No 4.0000000000000000000000000000000")
	dragToDrop1label := newDropablelabel(color.Gray{0x33}, "No 5..0000000000000000000000000000000000000")
	dragToDrop2Label := newDropablelabel(color.Gray{0x44}, "No 6.000000000000000000000000000000000000000000000000")

	DropOne := container.NewMax(dragToDrop1label.backgroundRectangle, dragToDrop1label)
	DropTwo := container.NewMax(dragToDrop2Label.backgroundRectangle, dragToDrop2Label)

	fromContainer := container.NewHBox(dragFromOneLabel, dragFromTwoLabel)
	toContainer := container.NewHBox(dragFromThreeLabel, dragFromFourLabel)
	dropContainer := container.NewHBox(DropOne, DropTwo)

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

	myLabel.backgroundRectangle = canvas.NewRectangle(color.RGBA{
		R: 0x00,
		G: 0xFF,
		B: 0x00,
		A: 0x44,
	})
	myLabel.Refresh()
	myLabel.backgroundRectangle.SetMinSize(myLabel.Size())
	myLabel.backgroundRectangle.Hide()

	return myLabel
}

type draggableLabel struct {
	widget.Label
	myTitle string
	hovered bool
}

type droppableLabel struct {
	widget.Label
	myTitle             string
	hovered             bool
	backgroundRectangle *canvas.Rectangle
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

	case sourceStateSearching:
		fmt.Println("Dragged: 'sourceStateSearching'")
		return

	case sourceStateFinds:
		// switch state to 'sourceStateGrabs'
		switchStateForSource(sourceStateGrabs)

		return

	case sourceStateGrabs:
		// switch state to 'sourceStateDragging'
		switchStateForSource(sourceStateDragging)

		return

	case sourceStateDragging:
		switchStateForTarget(targetStateSourceIsDraggingObject)
		// Just continue

	case sourceStateReleasingWithOutTarget:
		fmt.Println("Dragged: 'sourceStateReleasingWithOutTarget'")
		return

	case sourceStateEnteringTarget:
		fmt.Println("Dragged: 'sourceStateEnteringTarget'")
		return

	case sourceStateReleasingOnTarget:
		fmt.Println("Dragged: 'sourceStateReleasingOnTarget'")
		return

	case sourceStateReleasedOnTarget:
		fmt.Println("Dragged: 'sourceStateReleasedOnTarget'")
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

	case sourceStateSearching:
		fmt.Println("Dragged: 'sourceStateSearching'")
		return

	case sourceStateFinds:
		fmt.Println("Dragged: 'sourceStateFinds'")
		return

	case sourceStateGrabs:
		fmt.Println("Dragged: 'sourceStateGrabs'")
		return

	case sourceStateDragging:
		// switch state to 'sourceStateReleasingWithOutTarget'
		switchStateForSource(sourceStateReleasingWithOutTarget)
		switchStateForTarget(targetStateWaitingForSourceToEnteringTarget)

	case sourceStateReleasingWithOutTarget:
		// Just continue

	case sourceStateEnteringTarget:
		switchStateForSource(sourceStateReleasingOnTarget)
		return

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

	fmt.Println("MouseIn-Source: ", stateMachineDragFrom.currentState, stateMachineTarget.currentState)

}

// MouseMoved is called when a desktop pointer hovers over the widget
func (b *draggableLabel) MouseMoved(a *desktop.MouseEvent) {
	//log.Println("I have been 'MouseMoved' ", b.myTitle)

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

	fmt.Println("MouseOut-Source: ", stateMachineDragFrom.currentState, stateMachineTarget.currentState)

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
		b.backgroundRectangle.Show()
		fmt.Println("Show")

	case targetStateSourceEnteredTargetWithObject:
		return

	case targetStateSourceReleasingOnTarget:
		return

	case targetStateSourceReleasedOnTarget:
		return

	default:
		log.Fatalln("Unhandled state for StateMachine(From): ", stateMachineDragFrom.currentState)

	}

	fmt.Println("MouseIn-Target: ", stateMachineDragFrom.currentState, stateMachineTarget.currentState)

}

// MouseMoved is called when a desktop pointer hovers over the widget
func (b *droppableLabel) MouseMoved(a *desktop.MouseEvent) {
	//log.Println("I have been 'MouseMoved' ", b.myTitle)

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
		b.backgroundRectangle.Hide()
		fmt.Println("Hide")

	case targetStateSourceReleasingOnTarget:
		return

	case targetStateSourceReleasedOnTarget:
		return

	default:
		log.Fatalln("Unhandled state for StateMachine(From): ", stateMachineDragFrom.currentState)

	}

	fmt.Println("MouseOut-Target: ", stateMachineDragFrom.currentState, stateMachineTarget.currentState)

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
func switchStateForSource(newState int) {
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
