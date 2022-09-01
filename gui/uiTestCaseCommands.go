package gui

import (
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strings"
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

var (
	availableTestCasesSelectWidget                widget.Select
	availableBuildingBlocksSelectWidget           widget.Select
	availableBuildingBlocksInTestCaseSelectWidget widget.Select
)

func (uiServer *UIServerStruct) createTestCaseCommandParametersUI() (testCaseCommandsParametersUIObject fyne.CanvasObject) {

	// List alla TestCases
	availableTestCasesSelectWidget = widget.Select{
		DisableableWidget: widget.DisableableWidget{},
		Alignment:         0,
		Selected:          "",
		Options:           uiServer.testCasesModel.ListAvailableTestCases(),
		PlaceHolder:       "",
		OnChanged: func(s string) {
			fmt.Println("I selected %s to live forever..", s)
			availableTestCaseElements, err := uiServer.testCasesModel.ListAllAvailableBuildingBlocksInTestCase(availableTestCasesSelectWidget.Selected)

			if err != nil {
				fmt.Println(err)
			} else {
				availableBuildingBlocksInTestCaseSelectWidget.Options = availableTestCaseElements
			}
			//label1.Text = s
			//label1.Refresh()
		},
	}

	// List all Available BuildingBlocks
	availableBuildingBlocksSelectWidget = widget.Select{
		DisableableWidget: widget.DisableableWidget{},
		Alignment:         0,
		Selected:          "",
		Options:           uiServer.availableBuildingBlocksModel.listAllAvailableBuidlingBlocks(),
		PlaceHolder:       "",
		OnChanged: func(s string) {
			fmt.Printf("I selected %s to live forever..", s)
			//label1.Text = s
			//label1.Refresh()
		},
	}
	/*
		Select(
			uiServer.availableBuildingBlocksModel.listAllAvailableBuidlingBlocks(),
			func(s string) {
				fmt.Printf("I selected %s to live forever..", s)
				//label1.Text = s
				//label1.Refresh()
			})

	*/

	// List all Elements for current TestCase
	availableBuildingBlocksInTestCaseSelectWidget = widget.Select{
		DisableableWidget: widget.DisableableWidget{},
		Alignment:         0,
		Selected:          "",
		Options:           nil,
		PlaceHolder:       "",
		OnChanged: func(s string) {
			fmt.Printf("I selected %s to live forever..", s)
			//label1.Text = s
			//label1.Refresh()
		},
	}

	testCaseCommandsUIObject_temp := container.New(
		layout.NewVBoxLayout(),
		&availableTestCasesSelectWidget,
		&availableBuildingBlocksSelectWidget,
		&availableBuildingBlocksInTestCaseSelectWidget)

	testCaseCommandsParametersUIObject = container.NewScroll(testCaseCommandsUIObject_temp)

	return testCaseCommandsParametersUIObject

}
func (uiServer *UIServerStruct) createTestCaseCommandsUI() (testCaseCommandsUIObject fyne.CanvasObject) {

	// List alla TestCases
	availableTestCasesLabelWidget := widget.NewLabel("Available TestCases")
	availableTestCasesSelectWidget = widget.Select{
		DisableableWidget: widget.DisableableWidget{},
		Alignment:         0,
		Selected:          "",
		Options:           uiServer.testCasesModel.ListAvailableTestCases(),
		PlaceHolder:       "",
		OnChanged: func(s string) {
			fmt.Println("I selected %s to live forever..", s)
			availableTestCaseElements, err := uiServer.testCasesModel.ListAllAvailableBuildingBlocksInTestCase(availableTestCasesSelectWidget.Selected)

			if err != nil {
				fmt.Println(err)
			} else {

				// Update UI with TestCase Textual Representation
				textualTestCaseSimple, textualTestCaseComplex, textualTestCaseExtended, err := uiServer.commandAndRuleEngine.Testcases.CreateTextualTestCase(availableTestCasesSelectWidget.Selected)
				if err != nil {
					fmt.Println(err)
				}

				// Update Textual Representations, in UI-model, for TestCase
				err = uiServer.testCasesUiModel.UpdateTextualStructuresForTestCase(
					availableTestCasesSelectWidget.Selected,
					textualTestCaseSimple,
					textualTestCaseComplex,
					textualTestCaseExtended)

				if err != nil {
					fmt.Println(err)

					return
				}

				availableBuildingBlocksInTestCaseSelectWidget.Options = availableTestCaseElements

				availableBuildingBlocksInTestCaseSelectWidget.Refresh()

			}
			//label1.Text = s
			//label1.Refresh()
		},
	}

	// List all Available BuildingBlocks
	availableBuildingBlocksLabelWidget := widget.NewLabel("Available BuildingBlocks")
	availableBuildingBlocksSelectWidget = widget.Select{
		DisableableWidget: widget.DisableableWidget{},
		Alignment:         0,
		Selected:          "",
		Options:           uiServer.availableBuildingBlocksModel.listAllAvailableBuidlingBlocks(),
		PlaceHolder:       "",
		OnChanged: func(s string) {
			fmt.Printf("I selected %s to live forever..", s)
			//label1.Text = s
			//label1.Refresh()
		},
	}
	/*
		Select(
			uiServer.availableBuildingBlocksModel.listAllAvailableBuidlingBlocks(),
			func(s string) {
				fmt.Printf("I selected %s to live forever..", s)
				//label1.Text = s
				//label1.Refresh()
			})

	*/

	// List all Elements for current TestCase
	availableBuildingBlocksInTestCaseLabelWidget := widget.NewLabel("All Elements for current TestCase")
	availableBuildingBlocksInTestCaseSelectWidget = widget.Select{
		DisableableWidget: widget.DisableableWidget{},
		Alignment:         0,
		Selected:          "",
		Options:           nil,
		PlaceHolder:       "",
		OnChanged: func(s string) {
			fmt.Printf("I selected %s to live forever..", s)
			//label1.Text = s
			//label1.Refresh()
		},
	}

	newTestCaseButton := widget.NewButton(CommandNewTestcase, func() {
		uiServer.newTestCase()
	})
	removeButton := widget.NewButton(CommandRemoveElementFromTestcase, func() {
		uiServer.remove(availableTestCasesSelectWidget.Selected, availableBuildingBlocksInTestCaseSelectWidget.Selected)
	})
	swapFromNewButton := widget.NewButton(CommandSwapFromNewComponent, func() {
		uiServer.swapFromNew(availableTestCasesSelectWidget.Selected, availableBuildingBlocksInTestCaseSelectWidget.Selected, availableBuildingBlocksSelectWidget.Selected)
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

	commandHeaderLabelWidget := widget.NewLabel("Commands")
	commandHeaderLabelCanvasObject := container.New(layout.NewHBoxLayout(), commandHeaderLabelWidget)

	selectWidgetsGrid := container.New(layout.NewFormLayout(),
		availableTestCasesLabelWidget,
		&availableTestCasesSelectWidget,
		availableBuildingBlocksLabelWidget,
		&availableBuildingBlocksSelectWidget,
		availableBuildingBlocksInTestCaseLabelWidget,
		&availableBuildingBlocksInTestCaseSelectWidget)

	separator1Widget := widget.NewSeparator()
	separator2Widget := widget.NewSeparator()

	testCaseCommandsUIObject_temp := container.New(
		layout.NewVBoxLayout(),
		commandHeaderLabelCanvasObject,
		separator1Widget,
		selectWidgetsGrid,
		separator2Widget,
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
	testCaseUuid, err := uiServer.commandAndRuleEngine.NewTestCaseModel()

	if err != nil {
		fmt.Println(err)

		return

	}

	// Update List with available TestCases and Select the New one
	availableTestCasesSelectWidget.Options = uiServer.testCasesModel.ListAvailableTestCases()
	availableTestCasesSelectWidget.Selected = testCaseUuid
	availableTestCasesSelectWidget.Refresh()

	// Clear DropDown for 'Available Building Blocks'
	availableBuildingBlocksSelectWidget.Selected = ""
	availableBuildingBlocksSelectWidget.Refresh()

	// Load avaialble Building Blocks for newly created TestCase
	availableBuildingBlocksInTestCaseSelectWidget.Selected = ""
	availableTestCaseElements, err := uiServer.testCasesModel.ListAllAvailableBuildingBlocksInTestCase(availableTestCasesSelectWidget.Selected)

	if err != nil {
		fmt.Println(err)
	} else {
		availableBuildingBlocksInTestCaseSelectWidget.Options = availableTestCaseElements
	}
	availableBuildingBlocksInTestCaseSelectWidget.Refresh()

	// Update UI with TestCase Textual Representation
	textualTestCaseSimple, textualTestCaseComplex, textualTestCaseExtended, err := uiServer.commandAndRuleEngine.Testcases.CreateTextualTestCase(availableTestCasesSelectWidget.Selected)
	if err != nil {
		fmt.Println(err)

		return
	}

	// Generate UI for New TestCase
	err = uiServer.testCasesUiModel.GenerateNewTestCaseTabObject(testCaseUuid)
	if err != nil {
		fmt.Println(err)

		return
	}

	// Update Textual Representations, in UI-model, for TestCase
	err = uiServer.testCasesUiModel.UpdateTextualStructuresForTestCase(
		testCaseUuid,
		textualTestCaseSimple,
		textualTestCaseComplex,
		textualTestCaseExtended)

	if err != nil {
		fmt.Println(err)

		return
	}

	// Update Graphical TestCase Representation
	err = uiServer.testCasesUiModel.UpdateGraphicalRepresentationForTestCase(testCaseUuid)
	if err != nil {
		fmt.Println(err)

		return
	}

}

// Remove(ElementToBeRemoved)
func (uiServer *UIServerStruct) remove(testCaseUuid string, elementUiNameoBeRemoved string) {

	fmt.Printf("Remove(ElementToBeRemoved='%s' in TestCase='%s')\n", elementUiNameoBeRemoved, testCaseUuid)
	bindedCommandListData.Prepend(CommandRemoveElementFromTestcase)

	// Convert UI-name for element into elements UUID
	elementUuid, err := uiServer.testCasesModel.GetUuidFromUiName(testCaseUuid, elementUiNameoBeRemoved)
	if err != nil {
		fmt.Println(err)

		return
	}

	// Delete Element from TestCase
	err = uiServer.commandAndRuleEngine.DeleteElementFromTestCaseModel(testCaseUuid, elementUuid)
	if err != nil {
		fmt.Println(err)

		return
	}

	// List Available TestCase BuildingBlocks and add the DropDown
	availableTestCaseElements, err := uiServer.testCasesModel.ListAllAvailableBuildingBlocksInTestCase(availableTestCasesSelectWidget.Selected)

	if err != nil {
		fmt.Println(err)
	} else {
		availableBuildingBlocksInTestCaseSelectWidget.Options = availableTestCaseElements
	}

	// Clear DropDown for 'Available Building Blocks' and 'TestCase Building Blocks'
	availableBuildingBlocksSelectWidget.Selected = ""
	availableBuildingBlocksInTestCaseSelectWidget.Selected = ""

	// Update UI with TestCase Textual Representation
	textualTestCaseSimple, textualTestCaseComplex, textualTestCaseExtended, err := uiServer.commandAndRuleEngine.Testcases.CreateTextualTestCase(availableTestCasesSelectWidget.Selected)
	if err != nil {
		fmt.Println(err)
	}

	// Update Textual Representations, in UI-model, for TestCase
	err = uiServer.testCasesUiModel.UpdateTextualStructuresForTestCase(
		testCaseUuid,
		textualTestCaseSimple,
		textualTestCaseComplex,
		textualTestCaseExtended)

	if err != nil {
		fmt.Println(err)

		return
	}

	// Update Graphical TestCase Representation
	err = uiServer.testCasesUiModel.UpdateGraphicalRepresentationForTestCase(testCaseUuid)
	if err != nil {
		fmt.Println(err)

		return
	}

}

// SwapFromNew(ElementTobeSwappedOut, NewElementTobeSwappedIn)
func (uiServer *UIServerStruct) swapFromNew(testCaseUuid string, elementUiNameTobeSwappedOut string, newElementUiNameTobeSwappedIn string) {

	fmt.Printf("SwapFromNew(elementUiNameTobeSwappedOut='%s', newElementUiNameTobeSwappedIn='%s') in TestCase '%s'\n", elementUiNameTobeSwappedOut, newElementUiNameTobeSwappedIn, testCaseUuid)
	bindedCommandListData.Prepend(CommandSwapFromNewComponent)

	// Convert UI-name for element into elements UUID
	elementUuidTobeSwappedOut, err := uiServer.testCasesModel.GetUuidFromUiName(testCaseUuid, elementUiNameTobeSwappedOut)
	if err != nil {
		fmt.Println(err)

		return
	}

	elementUuidTobeSwappedIn, buildingBlockType, err := uiServer.getUuidFromTreeName(newElementUiNameTobeSwappedIn)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(elementUuidTobeSwappedOut, elementUuidTobeSwappedIn)

	// Get the ImmatureElement To Swap In
	var immatureElementToSwapInTestCaseFormat testCaseModel.ImmatureElementStruct

	switch buildingBlockType {

	case TestInstruction:
		tempMap := uiServer.availableBuildingBlocksModel.allImmatureTestInstructionsBuildingBlocks
		immatureElementToSwapInOriginal := tempMap[elementUuidTobeSwappedIn].ImmatureSubTestCaseModel
		immatureElementToSwapInTestCaseFormat = uiServer.availableBuildingBlocksModel.convertGrpcElementModelIntoTestCaseElementModel(immatureElementToSwapInOriginal)

	case TestInstructionContainer:
		tempMap := uiServer.availableBuildingBlocksModel.allImmatureTestInstructionContainerBuildingBlocks
		immatureElementToSwapInOriginal := tempMap[elementUuidTobeSwappedIn].ImmatureSubTestCaseModel
		immatureElementToSwapInTestCaseFormat = uiServer.availableBuildingBlocksModel.convertGrpcElementModelIntoTestCaseElementModel(immatureElementToSwapInOriginal)

	default:

		errorId := "f7f2c257-c571-4050-a8c8-99aaca1dd24f"
		err = errors.New(fmt.Sprintf("unknown Building BLock Type: '%s' [ErrorID: %s]", buildingBlockType, errorId))

		fmt.Println(err)

		// Exit function
		return

	}

	// Execute Swap of Elements
	err = uiServer.commandAndRuleEngine.SwapElementsInTestCaseModel(testCaseUuid, elementUuidTobeSwappedOut, &immatureElementToSwapInTestCaseFormat)
	if err != nil {
		fmt.Println(err)

		return
	}

	// List Available TestCase BuildingBlocks and add the DropDown
	availableTestCaseElements, err := uiServer.testCasesModel.ListAllAvailableBuildingBlocksInTestCase(availableTestCasesSelectWidget.Selected)

	if err != nil {
		fmt.Println(err)
	} else {
		availableBuildingBlocksInTestCaseSelectWidget.Options = availableTestCaseElements
	}

	// Clear DropDown for 'Available Building Blocks' and 'TestCase Building Blocks'
	availableBuildingBlocksSelectWidget.Selected = ""
	availableBuildingBlocksSelectWidget.Refresh()
	availableBuildingBlocksInTestCaseSelectWidget.Selected = ""
	availableBuildingBlocksInTestCaseSelectWidget.Refresh()

	// Update UI with TestCase Textual Representation
	textualTestCaseSimple, textualTestCaseComplex, textualTestCaseExtended, err := uiServer.commandAndRuleEngine.Testcases.CreateTextualTestCase(availableTestCasesSelectWidget.Selected)
	if err != nil {
		fmt.Println(err)
	}

	// Update Textual Representations, in UI-model, for TestCase
	err = uiServer.testCasesUiModel.UpdateTextualStructuresForTestCase(
		testCaseUuid,
		textualTestCaseSimple,
		textualTestCaseComplex,
		textualTestCaseExtended)

	if err != nil {
		fmt.Println(err)

		return
	}

	// Update Graphical TestCase Representation
	err = uiServer.testCasesUiModel.UpdateGraphicalRepresentationForTestCase(testCaseUuid)
	if err != nil {
		fmt.Println(err)

		return
	}

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
	uiServer.testCasesUiModel.UpdateGraphicalRepresentationForTestCase(availableTestCasesSelectWidget.Selected)
	fmt.Println("hello")

}

// GetUuidFromUiName
// Finds the UUID for from a UI-name like ' B0_BOND [3c8a3bc] [BOND] to live forever..'
func (uiServer *UIServerStruct) getUuidFromTreeName(uiTreeName string) (buildingBlockUuid string, buildingBlockType BuildingBlock, err error) {

	// Get first square brackets, for part of UUID
	firstSquareBracketStart := strings.Index(uiTreeName, "[")
	firstSquareBracketEnd := strings.Index(uiTreeName, "]")

	if firstSquareBracketStart == -1 || firstSquareBracketEnd == -1 {
		errorId := "65b8f415-281d-46c9-b133-df1168683a03"
		err = errors.New(fmt.Sprintf("problem with finding first par of '[' or ']' in avaialble building block-name '%s' in testcase '%s' [ErrorID: %s]", uiTreeName, errorId))

		return "", -1, err
	}

	// Get second square brackets, for type
	secondSquareBracketStart := strings.Index(uiTreeName[firstSquareBracketEnd+1:], "[")
	secondSquareBracketEnd := strings.Index(uiTreeName[firstSquareBracketEnd+1:], "]")

	if secondSquareBracketStart == -1 || secondSquareBracketEnd == -1 {
		errorId := "814c80b7-3759-46ad-b9e6-ddaa84f8a642"
		err = errors.New(fmt.Sprintf("problem with finding second par of '[' or ']' in avaialble building block-name '%s' in testcase '%s' [ErrorID: %s]", uiTreeName, errorId))

		return "", -1, err
	}

	// Extract UUID-part
	uuidPart := uiTreeName[firstSquareBracketStart+1 : firstSquareBracketEnd]

	// Extract Type
	elementTypeFromName := uiTreeName[firstSquareBracketEnd+1:][secondSquareBracketStart+1 : secondSquareBracketEnd]

	// Loop all available building blocks and find match
	for _, buildingBlock := range uiServer.availableBuildingBlocksModel.allBuildingBlocksTreeNameToUuid {

		switch elementTypeFromName {

		// TestInstructions
		case "TI":
			if buildingBlock.buildingBlockType == TestInstruction &&
				buildingBlock.uuid[:len(uuidPart)] == uuidPart {

				return buildingBlock.uuid, buildingBlock.buildingBlockType, nil
			}

			// TestInstructionContainers
		case "TIC":
			if buildingBlock.buildingBlockType == TestInstructionContainer &&
				buildingBlock.uuid[:len(uuidPart)] == uuidPart {

				return buildingBlock.uuid, buildingBlock.buildingBlockType, nil
			}

			// Bonds
		default:
			errorId := "70335847-35cd-4551-bea8-59257075723d"
			err = errors.New(fmt.Sprintf("couldn't find avavialbel BuildingBlockType with UI-name '%s' in testcase '%s' [ErrorID: %s]", uiTreeName, errorId))

			return "", -1, err
		}

	}
	return "", -1, err

}
