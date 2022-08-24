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
)

func makeDragNDropTestGUI(textIn *canvas.Text, recIn *canvas.Rectangle, rec2In *canvas.Rectangle, containerIn *fyne.Container) (myCanvasObject fyne.CanvasObject) {

	textRef = textIn
	rectangleRef = recIn
	rectangle2Ref = rec2In
	containerRef = containerIn

	dragFromOneRectangle := newDraggableRectangle(color.Gray{0xEE}, "No 1")
	dragFromTwoRectangle := newDraggableRectangle(color.Gray{0xBB}, "No 2.000000")
	dragToOneRectangle := newDraggableRectangle(color.Gray{0x88}, "No 3..00000000000000000")
	dragToTwoRectangle := newDraggableRectangle(color.Gray{0x44}, "No 4.0000000000000000000000000000000")

	fromContainer := container.NewHBox(dragFromOneRectangle, dragFromTwoRectangle)
	toContainer := container.NewHBox(dragToOneRectangle, dragToTwoRectangle)

	myText := widget.NewLabel("Test Area for Drag n Drop")

	myCanvasObject = container.NewVBox(myText, fromContainer, layout.NewSpacer(), toContainer)

	return myCanvasObject
}

var textRef *canvas.Text
var rectangleRef *canvas.Rectangle
var rectangle2Ref *canvas.Rectangle
var containerRef *fyne.Container

func newDraggableRectangle(myColor color.Gray, myNewTitle string) *draggableLabel {
	myRectangle := &draggableLabel{}
	myRectangle.ExtendBaseWidget(myRectangle)
	//myRectangle.FillColor = myColor
	//myRectangle.StrokeColor = color.Black
	//myRectangle.StrokeWidth = 0
	//myRectangle.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize), float32(testCaseNodeRectangleSize)))

	myRectangle.myTitle = myNewTitle
	myRectangle.Text = myNewTitle

	return myRectangle
}

type draggableLabel struct {
	widget.Label
	//widget.Button
	myTitle string
	hovered bool
}

type draggableStateStruct struct {
	dragStart              bool
	dragEnd                bool
	MouseIn                bool
	MouseOut               bool
	LightRectangle         bool
	preLightRectangleState bool
	dropIsAllowed          bool
	dragFrom               string
	DropIn                 string
}

var stateMachine draggableStateStruct

func (t *draggableLabel) Dragged(ev *fyne.DragEvent) {
	//log.Println("I have been 'Dragged': ", t.Position(), " And I am ", t.myTitle)

	t.switchDragStart(t.myTitle)
	//fmt.Println(ev.Position, ev.AbsolutePosition)

	textRef.Text = t.myTitle
	textRef.Refresh()

	rectangleRef.SetMinSize(textRef.Size().Add(fyne.NewSize(50, 50)))
	rectangleRef.Refresh()

	rectangle2Ref.SetMinSize(textRef.Size())
	rectangle2Ref.Refresh()

	diffPos := fyne.Position{
		X: 10,
		Y: -20,
	}

	newPos := ev.AbsolutePosition.Add(diffPos).Add(fyne.NewSize(rectangleRef.Size().Width/2, rectangleRef.Size().Height/2))

	containerRef.Move(newPos)

	containerRef.Refresh()
	textRef.Show()
	rectangle2Ref.Show()
	rectangleRef.Show()

}

func (t *draggableLabel) DragEnd() {
	//log.Println("I have been 'DragEnd' ", t.myTitle)
	t.switchDragEnd(t.myTitle)

	textRef.Hide()
	rectangleRef.Hide()
	rectangle2Ref.Hide()
	containerRef.Refresh()

}

// MouseIn is called when a desktop pointer enters the widget
func (b *draggableLabel) MouseIn(*desktop.MouseEvent) {
	//log.Println("I have been 'MouseIn' ", b.myTitle)
	b.hovered = true
	b.switchMouseIn(b.myTitle)
	//b.Refresh()
}

// MouseMoved is called when a desktop pointer hovers over the widget
func (b *draggableLabel) MouseMoved(a *desktop.MouseEvent) {
	//log.Println("I have been 'MouseMoved' ", b.myTitle)
	//b.switchMouseIn(b.myTitle)
	//b.Refresh()
	//fmt.Println(a.Position, a.AbsolutePosition)
}

// MouseOut is called when a desktop pointer exits the widget
func (b *draggableLabel) MouseOut() {
	//log.Println("I have been 'MouseOut' ", b.myTitle)
	b.hovered = false
	b.switchMouseOut(b.myTitle)
	//b.Refresh()

}

func (b *draggableLabel) switchDragStart(labelNo string) {
	stateMachine.dragStart = true
	stateMachine.dragEnd = false

	stateMachine.dragFrom = labelNo

	b.stateMachineSwitcher(true)
}

func (b *draggableLabel) switchDragEnd(labelNo string) {
	stateMachine.dragStart = false
	stateMachine.dragEnd = true

	b.stateMachineSwitcher(true)
}

func (b *draggableLabel) switchMouseIn(labelNo string) {
	stateMachine.MouseIn = true
	stateMachine.MouseOut = false

	stateMachine.DropIn = labelNo

	b.stateMachineSwitcher(true)
}

func (b *draggableLabel) switchMouseOut(labelNo string) {
	stateMachine.MouseIn = false
	stateMachine.MouseOut = true

	b.stateMachineSwitcher(true)
}

func (b *draggableLabel) stateMachineSwitcher(doAction bool) {

	//fmt.Println(stateMachine.dragStart, stateMachine.dragEnd, stateMachine.MouseIn, stateMachine.MouseOut, stateMachine.LightRectangle, stateMachine.preLightRectangleState, stateMachine.LightRectangle, stateMachine.dropIsAllowed)

	stateMachine.preLightRectangleState = stateMachine.LightRectangle

	if stateMachine.dragStart == true &&
		stateMachine.dragEnd == false &&
		stateMachine.MouseIn == true &&
		stateMachine.MouseOut == false &&
		stateMachine.DropIn != stateMachine.dragFrom {

		stateMachine.LightRectangle = true
		if doAction && stateMachine.preLightRectangleState == false {
			fmt.Println("Light up rectangle no ", stateMachine.DropIn)
		}

	} else {
		stateMachine.LightRectangle = false
		if doAction && stateMachine.preLightRectangleState == true {
			fmt.Println("Turn of Light for rectangle no ", stateMachine.DropIn)
		}
	}

	if stateMachine.dragStart == false &&
		stateMachine.dragEnd == true &&
		stateMachine.MouseIn == true &&
		stateMachine.MouseOut == false &&
		stateMachine.preLightRectangleState == true &&
		stateMachine.DropIn != stateMachine.dragFrom {

		stateMachine.dropIsAllowed = true

		if doAction {
			fmt.Println("Drag rectangle no ", stateMachine.dragFrom, " and drop in rectangle no ", stateMachine.DropIn)
		}

	} else {
		stateMachine.dropIsAllowed = false
	}

	//fmt.Println(stateMachine.dragStart, stateMachine.dragEnd, stateMachine.MouseIn, stateMachine.MouseOut, stateMachine.LightRectangle, stateMachine.preLightRectangleState, stateMachine.LightRectangle, stateMachine.dropIsAllowed)
}
