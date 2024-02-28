package testCaseUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testCase/testCaseModel"
	"FenixTesterGui/testUIDragNDropStatemachine"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"image/color"
	"strconv"
)

// Generates the graphical structure for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) makeTestCaseGraphicalUIObject(
	testCaseUuid string) (
	testCaseCanvasObject fyne.CanvasObject) {

	treeViewModelForTestCase, err := testCasesUiCanvasObject.TestCasesModelReference.GetTreeViewModelForTestCase(
		testCaseUuid)

	if err != nil {
		errText := widget.NewLabel(err.Error())
		testCaseCanvasObject = container.NewVBox(errText)

		return testCaseCanvasObject

	}

	// Clear state machine for Drag N Drop
	testCasesUiCanvasObject.DragNDropStateMachine = testUIDragNDropStatemachine.StateMachineDragAndDropStruct{}

	// Container holding all TestInstructions, TestInstructionContainers and Bonds
	var testcaseTreeContainer *fyne.Container
	testcaseTreeContainer = container.NewVBox()

	// Start processing model for TestCase
	testCaseCanvasObject = testCasesUiCanvasObject.recursiveMakeTestCaseGraphicalUIObject(
		"",
		&treeViewModelForTestCase,
		nil, 1,
		testCaseUuid,
		testcaseTreeContainer)

	return testcaseTreeContainer

}

// Generates the graphical structure for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) recursiveMakeTestCaseGraphicalUIObject(
	uuid string,
	testCaseModelForUITree *map[string][]testCaseModel.TestCaseModelAdaptedForUiTreeDataStruct,
	firstAccordion *clickableAccordion,
	nodeTreeLevel float32,
	testCaseUuid string,
	testcaseTreeContainer *fyne.Container) (
	testCaseCanvasObject fyne.CanvasObject) {

	var newTestInstructionAccordion *clickableAccordion

	if firstAccordion != nil {
		newTestInstructionAccordion = firstAccordion
	}

	testCaseModelForUITreeMap := *testCaseModelForUITree

	testCaseNodeChildren := testCaseModelForUITreeMap[uuid]

	//nodeChildrenContainer := container.NewVBox()

	for _, child := range testCaseNodeChildren {

		switch child.NodeTypeEnum {

		// Some kind of TestInstruction
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE:

			// Extract the node name
			var shortedUuid string
			shortedUuid = sharedCode.GenerateShortUuidFromFullUuid(child.Uuid)
			nodeName := fmt.Sprintf("%s [%s]",
				child.NodeName,
				shortedUuid,
			)

			// Create the color for TestInstruction Type
			rectangleColor, err := testCasesUiCanvasObject.convertRGBAHexStringIntoRGBAColor(child.NodeColor)
			if err != nil {
				nodeName = err.Error()
			}

			// Create color to use
			newTransparentColor := color.RGBA{
				R: 0x00,
				G: 0x00,
				B: 0x00,
				A: 0x00,
			}

			// Create indentation colorRectangle to move node to right
			newIndentationRectangle := canvas.NewRectangle(newTransparentColor)
			newIndentationRectangle.StrokeColor = color.Black
			newIndentationRectangle.StrokeWidth = 0
			newIndentationRectangle.SetMinSize(fyne.NewSize(testCaseNodeRectangleSize*nodeTreeLevel, float32(0)))
			newIndentationRectangleContainer := container.NewMax(newIndentationRectangle)

			// Create indentation within TestInstructionContainer

			// Create colorRectangle to show TestInstruction-color
			newTestInstructionColorRectangle := testCasesUiCanvasObject.NewClickableRectangle(
				rectangleColor, testCaseUuid, child.Uuid)

			testInstructionNodeColorContainer := container.NewMax(
				newTestInstructionColorRectangle.colorRectangle,
				newTestInstructionColorRectangle.selectedRectangle,
				newTestInstructionColorRectangle)

			// Create the Accordion-object to hold information about the TestInstruction
			dummyText := widget.NewLabel("this is just a dummy text and might show other TestInstruction-attributes later on")
			newTestInstructionAccordionItem := widget.NewAccordionItem(nodeName, dummyText)
			newTestInstructionAccordion = testCasesUiCanvasObject.newClickableAccordion(
				newTestInstructionAccordionItem, true, testCaseUuid, child.OriginalUuid) //widget.NewAccordion(newTestInstructionAccordionItem)

			// Create the container object to be put on GUI
			nodeContainer := container.NewHBox(
				newIndentationRectangleContainer, testInstructionNodeColorContainer, newTestInstructionAccordion)

			// Add the child
			testcaseTreeContainer.Add(nodeContainer)

		// Some kind of TestInstructionContainer
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE:

			// Create the Accordion-object to hold information about the TestInstructionContainer
			dummyText := widget.NewLabel("this is just a dummy text and might show other TestInstructionContainer-attributes later on")
			newTestInstructionAccordionItem := widget.NewAccordionItem(child.NodeName+" - "+child.Uuid, dummyText)
			newTestInstructionContainerAccordion := testCasesUiCanvasObject.newClickableAccordion(
				newTestInstructionAccordionItem, true, testCaseUuid, child.OriginalUuid) //widget.NewAccordion(newTestInstructionAccordionItem)

			// Create color to use
			newTransparentColor := color.RGBA{
				R: 0x00,
				G: 0x00,
				B: 0x00,
				A: 0x00,
			}

			// Create indentation colorRectangle to move node to right
			newIndentationRectangle := canvas.NewRectangle(newTransparentColor)
			newIndentationRectangle.StrokeColor = color.Black
			newIndentationRectangle.StrokeWidth = 0
			newIndentationRectangle.SetMinSize(fyne.NewSize(testCaseNodeRectangleSize*nodeTreeLevel, float32(0)))
			newIndentationRectangleContainer := container.NewMax(newIndentationRectangle)

			// Create the Horizontal node container object to be put on GUI
			nodeHContainer := container.NewHBox(
				newIndentationRectangleContainer, newTestInstructionContainerAccordion, layout.NewSpacer())

			// Create trailer colorRectangle for TestInstructionContainer
			newITrailerRectangle := canvas.NewRectangle(newTransparentColor)
			newITrailerRectangle.StrokeColor = color.Black
			newITrailerRectangle.StrokeWidth = 0
			newITrailerRectangle.SetMinSize(fyne.NewSize(1, 4))
			newITrailerRectangleContainer := container.NewMax(newITrailerRectangle)

			// Create the node container object to be put on GUI
			nodeContainer := container.NewVBox(nodeHContainer, newITrailerRectangleContainer)

			// Add the child
			testcaseTreeContainer.Add(nodeContainer)

			// Process children for TestInstructionContainer
			_ = testCasesUiCanvasObject.recursiveMakeTestCaseGraphicalUIObject(
				child.Uuid,
				testCaseModelForUITree,
				newTestInstructionContainerAccordion,
				nodeTreeLevel+2.0,
				testCaseUuid,
				testcaseTreeContainer)

		// Some kind of droppable Bond
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10ox_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10oxo_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10xo_BOND:

			newDroppableBondLabel := testCasesUiCanvasObject.DragNDropStateMachine.NewDroppableLabel(
				child.NodeName+" - "+child.Uuid, nodeTreeLevel, testCaseNodeRectangleSize, child.Uuid, testCaseUuid)
			newDroppableBondLabelContainer := container.NewMax(
				newDroppableBondLabel.BackgroundRectangle, newDroppableBondLabel)
			newDroppableBondLabel.Hide()

			// Create color to use
			newTransaparentColor := color.RGBA{
				R: 0x00,
				G: 0x00,
				B: 0x00,
				A: 0x00,
			}

			// Create indentation colorRectangle to move node to right
			newIndentationRectangle := canvas.NewRectangle(newTransaparentColor)
			newIndentationRectangle.StrokeColor = color.Black
			newIndentationRectangle.StrokeWidth = 0
			newIndentationRectangle.SetMinSize(fyne.NewSize(testCaseNodeRectangleSize*nodeTreeLevel, float32(0)))
			newIndentationRectangleContainer := container.NewMax(newIndentationRectangle)

			// Create indentation within TestInstructionContainer
			newTestInstructionColorRectangle := canvas.NewRectangle(newTransaparentColor)
			newTestInstructionColorRectangle.StrokeColor = color.Black
			newTestInstructionColorRectangle.StrokeWidth = 0
			newTestInstructionColorRectangle.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize), float32(0)))
			testInstructionNodeTransparentContainer := container.NewMax(newTestInstructionColorRectangle)

			// Create the none container object to be put on GUI
			nodeContainer := container.NewHBox(
				newIndentationRectangleContainer,
				testInstructionNodeTransparentContainer,
				newDroppableBondLabelContainer,
				layout.NewSpacer())

			// Add the child
			testcaseTreeContainer.Add(nodeContainer)

		// Some kind of non-droppable Bond
		default:

		}

	}

	testCaseCanvasObject = testcaseTreeContainer

	return testCaseCanvasObject

}

// Converts a colors in a hex-string into 'color.RGBA'-format. "#FF03AFFF"
func (testCasesUiCanvasObject *TestCasesUiModelStruct) convertRGBAHexStringIntoRGBAColor(rgbaHexString string) (rgbaValue color.RGBA, err error) {

	var (
		extractedColorRed          uint8
		extractedColorGreen        uint8
		extractedColorBlue         uint8
		extractedAlphaColorChannel uint8
	)

	errorColor := color.RGBA{
		R: 0xFF,
		G: 0x00,
		B: 0xFF,
		A: 0xFF}

	// Checka that ther String is of correct length, '#FFEEBB33'
	if len(rgbaHexString) != 9 {
		errorId := "6fe4273f-3a5f-44de-a86f-93b3b5bd5def"
		err = errors.New(fmt.Sprintf("color string with hexvalue, '%s', has not the correct lenght, '#AABBCCDD' in testcase with sourceUuid '%s' [ErrorID: %s]", rgbaHexString, "<unknown>", errorId))

		return errorColor, err
	}

	hexRed := rgbaHexString[1:3]
	hexGreen := rgbaHexString[3:5]
	hexBlue := rgbaHexString[5:7]
	hexAlpha := rgbaHexString[7:9]

	// Hex-conversion for Red
	valueRed, hexConvertionErr := strconv.ParseInt(hexRed, 16, 64)
	if hexConvertionErr != nil {
		// Hex conversion failed
		errorId := "162667e9-d35e-45be-b1b4-d07877a3cd2c"
		err = errors.New(fmt.Sprintf("hexConvertion for Color failed with error message: '%s' [ErrorID: %s]", hexConvertionErr, errorId))

		return errorColor, err

	}

	// Hex-conversion for Green
	valueGreen, hexConvertionErr := strconv.ParseInt(hexGreen, 16, 64)
	if hexConvertionErr != nil {
		// Hex conversion failed
		errorId := "b2b9fae0-30e3-49df-99d7-5662b78311a3"
		err = errors.New(fmt.Sprintf("hexConvertion for Color failed with error message: '%s' [ErrorID: %s]", hexConvertionErr, errorId))

		return errorColor, err

	}

	// Hex-conversion for Blue
	valueBlue, hexConvertionErr := strconv.ParseInt(hexBlue, 16, 64)
	if hexConvertionErr != nil {
		// Hex conversion failed
		errorId := "b2b9fae0-30e3-49df-99d7-5662b78311a3"
		err = errors.New(fmt.Sprintf("hexConvertion for Color failed with error message: '%s' [ErrorID: %s]", hexConvertionErr, errorId))

		return errorColor, err

	}

	// Hex-conversion for Alpha
	valueAlpha, hexConvertionErr := strconv.ParseInt(hexAlpha, 16, 64)
	if hexConvertionErr != nil {
		// Hex conversion failed
		errorId := "f5569252-41f5-48db-8a0a-b217b1460f7d"
		err = errors.New(fmt.Sprintf("hexConvertion for Color failed with error message: '%s' [ErrorID: %s]", hexConvertionErr, errorId))

		return errorColor, err

	}

	// Convert to color values
	extractedColorRed = uint8(valueRed)
	extractedColorGreen = uint8(valueGreen)
	extractedColorBlue = uint8(valueBlue)
	extractedAlphaColorChannel = uint8(valueAlpha)

	rgbaValue = color.RGBA{
		R: extractedColorRed,
		G: extractedColorGreen,
		B: extractedColorBlue,
		A: extractedAlphaColorChannel,
	}

	return rgbaValue, err

}
