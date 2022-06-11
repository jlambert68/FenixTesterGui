package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

/*

* NewTestCase()
[Creates and empty TestCase having a B0-area]

* Remove(ElementToBeRemoved)
[ElementToBeRemoved is a TI or a TIC]

* SwapFromNew(ElementTobeSwappedOut, NewElementTobeSwappedIn)
[ElementTobeSwappedOut is a Bond-area(B0, B10, B11 or B12) and NewElementTobeSwappedIn is a TI or TIC]

* Copy(Element) [Element is a TI or a TIC]

* SwapFromCopyBuffer(ElementTobeSwappedOut, CopyBufferElementTobeSwappedIn)
[ElementTobeSwappedOut is a Bond-area(B0, B10, B11 or B12) and CopyBufferElementTobeSwappedIn is a TI or TIC]

* Cut(Element) [Element is a TI or a TIC]

* SwapFromCutBuffer(ElementTobeSwappedOut, CutBufferElementTobeSwappedIn)
[ElementTobeSwappedOut is a Bond-area(B0, B10, B11 or B12) and CutBufferElementTobeSwappedIn is a TI or TIC]

* UndoLastCommandOnStack()
[Removes the last command from the stack of commands.
Example:
CommandStack = { SwapFromNew_1 ; SwapFromNew_2 ; SwapFromNew_3}
UndoLastCommandOnStack() --> CommandStack = { SwapFromNew_1 ; SwapFromNew_2}
UndoLastCommandOnStack() --> CommandStack = { SwapFromNew_1}]

*UndoUndoLastCommandOnStack()
[Removes the last command from the stack of commands.
Example:
CommandStack = { SwapFromNew_1 ; SwapFromNew_2 ; SwapFromNew_3}
UndoLastCommandOnStack() --> CommandStack = { SwapFromNew_1 ; SwapFromNew_2}
*/

const (
	CommandNewTestcase                = "NewTestCase()"
	CommandRemoveElementFromTestcase  = "Remove(ElementToBeRemoved)"
	CommandSwapFromNewComponent       = "SwapFromNew(ElementTobeSwappedOut, NewElementTobeSwappedIn)"
	CommandCopy                       = "Copy(Element)"
	CommandSwapFromCopyBuffer         = "SwapFromCopyBuffer(ElementTobeSwappedOut, CopyBufferElementTobeSwappedIn)"
	CommandCut                        = "Cut(Element)"
	CommandSwapFromCutBuffer          = "SwapFromCutBuffer(ElementTobeSwappedOut, CutBufferElementTobeSwappedIn)"
	CommandUndoLastCommandOnStack     = "UndoLastCommandOnStack()"
	CommandUndoUndoLastCommandOnStack = "UndoUndoLastCommandOnStack()"
)

func (testCaseUI *testCaseUIStruct) createTestCaseCommandsUI() (testCaseCommandsUIObject fyne.CanvasObject) {

	newTestCaseButton := widget.NewButton(CommandNewTestcase, func() {
		testCaseUI.newTestCase()
	})
	removeButton := widget.NewButton(CommandRemoveElementFromTestcase, func() {
		testCaseUI.remove("x")
	})
	swapFromNewButton := widget.NewButton(CommandSwapFromNewComponent, func() {
		testCaseUI.swapFromNew("x", "xx")
	})
	copyButton := widget.NewButton(CommandCopy, func() {
		testCaseUI.copy("x")
	})
	swapFromCopyBufferButton := widget.NewButton(CommandSwapFromCopyBuffer, func() {
		testCaseUI.swapFromCopyBuffer("x", "xx")
	})
	cutButton := widget.NewButton(CommandCut, func() {
		testCaseUI.cut("")
	})
	swapFromCutBufferButton := widget.NewButton(CommandSwapFromCutBuffer, func() {
		testCaseUI.swapFromCutBuffer("x", "xx")
	})
	undoLastCommandOnStackButton := widget.NewButton(CommandUndoLastCommandOnStack, func() {
		testCaseUI.undoLastCommandOnStack()
	})
	undoUndoLastCommandOnStackButton := widget.NewButton(CommandUndoUndoLastCommandOnStack, func() {
		testCaseUI.undoUndoLastCommandOnStack()
	})

	testCaseCommandsUIObject_temp := container.New(
		layout.NewVBoxLayout(),
		newTestCaseButton,
		removeButton,
		swapFromNewButton,
		copyButton,
		swapFromCopyBufferButton,
		cutButton,
		swapFromCutBufferButton,
		undoLastCommandOnStackButton,
		undoUndoLastCommandOnStackButton)

	testCaseCommandsUIObject = container.NewScroll(testCaseCommandsUIObject_temp)

	return testCaseCommandsUIObject
}

// NewTestCase()
func (testCaseUI *testCaseUIStruct) newTestCase() {

	fmt.Printf("NewTestCase()\n")

}

// Remove(ElementToBeRemoved)
func (testCaseUI *testCaseUIStruct) remove(elementToBeRemoved string) {

	fmt.Printf("Remove(ElementToBeRemoved='%s')\n", elementToBeRemoved)

}

// SwapFromNew(ElementTobeSwappedOut, NewElementTobeSwappedIn)
func (testCaseUI *testCaseUIStruct) swapFromNew(elementTobeSwappedOut string, newElementTobeSwappedIn string) {

	fmt.Printf("SwapFromNew(ElementTobeSwappedOut='%s', NewElementTobeSwappedIn='%s')\n", elementTobeSwappedOut, newElementTobeSwappedIn)

}

// Copy(Element)
func (testCaseUI *testCaseUIStruct) copy(element string) {

	fmt.Printf("Copy(Element='%s')\n", element)

}

// SwapFromCopyBuffer(ElementTobeSwappedOut, CopyBufferElementTobeSwappedIn)
func (testCaseUI *testCaseUIStruct) swapFromCopyBuffer(elementTobeSwappedOut string, copyBufferElementTobeSwappedIn string) {

	fmt.Printf("SwapFromCopyBuffer(ElementTobeSwappedOut='%s', CopyBufferElementTobeSwappedIn='%s')\n", elementTobeSwappedOut, copyBufferElementTobeSwappedIn)

}

// Cut(Element)
func (testCaseUI *testCaseUIStruct) cut(element string) {

	fmt.Printf("Cut(Element='%s')\n", element)

}

// SwapFromCutBuffer(ElementTobeSwappedOut, CutBufferElementTobeSwappedIn)
func (testCaseUI *testCaseUIStruct) swapFromCutBuffer(elementTobeSwappedOut string, cutBufferElementTobeSwappedIn string) {

	fmt.Printf("SwapFromCutBuffer(ElementTobeSwappedOut='%s', CutBufferElementTobeSwappedIn='%s')\n", elementTobeSwappedOut, cutBufferElementTobeSwappedIn)

}

// UndoLastCommandOnStack()
func (testCaseUI *testCaseUIStruct) undoLastCommandOnStack() {

	fmt.Printf("UndoLastCommandOnStack()\n")

}

// UndoUndoLastCommandOnStack()
func (testCaseUI *testCaseUIStruct) undoUndoLastCommandOnStack() {

	fmt.Printf("UndoUndoLastCommandOnStack()\n")

}
