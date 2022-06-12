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

func (uiServer *UIServerStruct) createTestCaseCommandsUI() (testCaseCommandsUIObject fyne.CanvasObject) {

	newTestCaseButton := widget.NewButton(CommandNewTestcase, func() {
		uiServer.newTestCase()
	})
	removeButton := widget.NewButton(CommandRemoveElementFromTestcase, func() {
		uiServer.remove("x")
	})
	swapFromNewButton := widget.NewButton(CommandSwapFromNewComponent, func() {
		uiServer.swapFromNew("x", "xx")
	})
	copyButton := widget.NewButton(CommandCopy, func() {
		uiServer.copy("x")
	})
	swapFromCopyBufferButton := widget.NewButton(CommandSwapFromCopyBuffer, func() {
		uiServer.swapFromCopyBuffer("x", "xx")
	})
	cutButton := widget.NewButton(CommandCut, func() {
		uiServer.cut("")
	})
	swapFromCutBufferButton := widget.NewButton(CommandSwapFromCutBuffer, func() {
		uiServer.swapFromCutBuffer("x", "xx")
	})
	undoLastCommandOnStackButton := widget.NewButton(CommandUndoLastCommandOnStack, func() {
		uiServer.undoLastCommandOnStack()
	})
	undoUndoLastCommandOnStackButton := widget.NewButton(CommandUndoUndoLastCommandOnStack, func() {
		uiServer.undoUndoLastCommandOnStack()
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
func (uiServer *UIServerStruct) newTestCase() {

	fmt.Printf("NewTestCase()\n")
	bindedCommandListData.Prepend(CommandNewTestcase)

}

// Remove(ElementToBeRemoved)
func (uiServer *UIServerStruct) remove(elementToBeRemoved string) {

	fmt.Printf("Remove(ElementToBeRemoved='%s')\n", elementToBeRemoved)
	bindedCommandListData.Prepend(CommandRemoveElementFromTestcase)

}

// SwapFromNew(ElementTobeSwappedOut, NewElementTobeSwappedIn)
func (uiServer *UIServerStruct) swapFromNew(elementTobeSwappedOut string, newElementTobeSwappedIn string) {

	fmt.Printf("SwapFromNew(ElementTobeSwappedOut='%s', NewElementTobeSwappedIn='%s')\n", elementTobeSwappedOut, newElementTobeSwappedIn)
	bindedCommandListData.Prepend(CommandSwapFromNewComponent)

}

// Copy(Element)
func (uiServer *UIServerStruct) copy(element string) {

	fmt.Printf("Copy(Element='%s')\n", element)
	bindedCommandListData.Prepend(CommandCopy)

}

// SwapFromCopyBuffer(ElementTobeSwappedOut, CopyBufferElementTobeSwappedIn)
func (uiServer *UIServerStruct) swapFromCopyBuffer(elementTobeSwappedOut string, copyBufferElementTobeSwappedIn string) {

	fmt.Printf("SwapFromCopyBuffer(ElementTobeSwappedOut='%s', CopyBufferElementTobeSwappedIn='%s')\n", elementTobeSwappedOut, copyBufferElementTobeSwappedIn)
	bindedCommandListData.Prepend(CommandSwapFromCopyBuffer)

}

// Cut(Element)
func (uiServer *UIServerStruct) cut(element string) {

	fmt.Printf("Cut(Element='%s')\n", element)
	bindedCommandListData.Prepend(CommandCut)

}

// SwapFromCutBuffer(ElementTobeSwappedOut, CutBufferElementTobeSwappedIn)
func (uiServer *UIServerStruct) swapFromCutBuffer(elementTobeSwappedOut string, cutBufferElementTobeSwappedIn string) {

	fmt.Printf("SwapFromCutBuffer(ElementTobeSwappedOut='%s', CutBufferElementTobeSwappedIn='%s')\n", elementTobeSwappedOut, cutBufferElementTobeSwappedIn)
	bindedCommandListData.Prepend(CommandSwapFromCutBuffer)

}

// UndoLastCommandOnStack()
func (uiServer *UIServerStruct) undoLastCommandOnStack() {

	fmt.Printf("UndoLastCommandOnStack()\n")
	bindedCommandListData.Prepend(CommandUndoLastCommandOnStack)

}

// UndoUndoLastCommandOnStack()
func (uiServer *UIServerStruct) undoUndoLastCommandOnStack() {

	fmt.Printf("UndoUndoLastCommandOnStack()\n")
	bindedCommandListData.Prepend(CommandUndoUndoLastCommandOnStack)

}
