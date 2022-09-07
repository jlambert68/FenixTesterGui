package testCaseUI

import (
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
func (testCasesUiCanvasObject *TestCasesUiModelStruct) makeTestCaseGraphicalUIObject(testCaseUuid string) (testCaseCanvasObject fyne.CanvasObject) {

	treeViewModelForTestCase, err := testCasesUiCanvasObject.TestCasesModelReference.GetTreeViewModelForTestCase(testCaseUuid)

	if err != nil {
		errText := widget.NewLabel(err.Error())
		testCaseCanvasObject = container.NewVBox(errText)

		return testCaseCanvasObject

	}

	// Clear state machine for Drag N Drop
	testCasesUiCanvasObject.DragNDropStateMachine = testUIDragNDropStatemachine.StateMachineDragAndDropStruct{}

	// Start processing model for TestCase
	testCaseCanvasObject, _ = testCasesUiCanvasObject.recursiveMakeTestCaseGraphicalUIObject("", &treeViewModelForTestCase, nil, 1, testCaseUuid)

	return testCaseCanvasObject

}

// Generates the graphical structure for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) recursiveMakeTestCaseGraphicalUIObject(uuid string, testCaseModelForUITree *map[string][]testCaseModel.TestCaseModelAdaptedForUiTreeDataStruct, firstAccordion *clickableAccordion, nodeTreeLevel float32, testCaseUuid string) (testCaseCanvasObject fyne.CanvasObject, newTestInstructionAccordion2 *clickableAccordion) {

	var childObject fyne.CanvasObject
	//var newTestInstructionAccordion *widget.Accordion
	var newTestInstructionAccordion *clickableAccordion

	if firstAccordion != nil {
		newTestInstructionAccordion = firstAccordion
	}

	testCaseModelForUITreeMap := *testCaseModelForUITree

	testCaseNodeChildren := testCaseModelForUITreeMap[uuid]

	nodeChildrenContainer := container.NewVBox()

	for _, child := range testCaseNodeChildren {

		switch child.NodeTypeEnum {

		// Some kind of TestInstruction
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE:

			// Extract the node name
			nodeName := child.NodeName

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

			// Create indentation rectangle to move node to right
			newIndentationRectangle := canvas.NewRectangle(newTransparentColor)
			newIndentationRectangle.StrokeColor = color.Black
			newIndentationRectangle.StrokeWidth = 0
			newIndentationRectangle.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize*nodeTreeLevel), float32(0)))
			newIndentationRectangleContainer := container.NewMax(newIndentationRectangle)

			// Create indentation within TestInstructionContainer

			// Create rectangle to show TestInstruction-color
			newTestInstructionColorRectangle := testCasesUiCanvasObject.NewClickableRectangle(rectangleColor, testCaseUuid, child.OriginalUuid)
			/*
				newTestInstructionColorRectangle := canvas.NewRectangle(rectangleColor)
				newTestInstructionColorRectangle.StrokeColor = color.Black
				newTestInstructionColorRectangle.StrokeWidth = 0
				newTestInstructionColorRectangle.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize), float32(testCaseNodeRectangleSize)))
			*/
			testInstructionNodeColorContainer := container.NewMax(newTestInstructionColorRectangle.rectangle, newTestInstructionColorRectangle)

			// Create the Accordion-object to hold information about the TestInstruction
			dummyText := widget.NewLabel("this is just a dummy text and might show other TestInstruction-attributes later on")
			newTestInstructionAccordionItem := widget.NewAccordionItem(nodeName, dummyText)
			newTestInstructionAccordion = testCasesUiCanvasObject.newClickableAccordion(newTestInstructionAccordionItem, true, testCaseUuid, child.OriginalUuid) //widget.NewAccordion(newTestInstructionAccordionItem)

			/*
				// Create color to use
				newTICommandtColor := color.RGBA{
					R: 0x33,
					G: 0x33,
					B: 0x33,
					A: 0x33,
				}

				// Create rectangle to show TestInstruction-command
				newTestInstructionCommandColorRectangle := canvas.NewRectangle(newTICommandtColor)
				newTestInstructionCommandColorRectangle.StrokeColor = color.RGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0x77}
				newTestInstructionCommandColorRectangle.StrokeWidth = 2
				newTestInstructionCommandColorRectangle.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize), float32(testCaseNodeRectangleSize)))
				testInstructionCommandColorContainer := container.NewMax(newTestInstructionCommandColorRectangle)


			*/
			// Create the container object to be put on GUI
			nodeContainer := container.NewHBox(newIndentationRectangleContainer, testInstructionNodeColorContainer, newTestInstructionAccordion)

			// Add the child
			nodeChildrenContainer.Add(nodeContainer)

		// Some kind of TestInstructionContainer
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE:

			// Create the Accordion-object to hold information about the TestInstructionContainer
			dummyText := widget.NewLabel("this is just a dummy text")
			newTestInstructionAccordionItem := widget.NewAccordionItem("DummyText", dummyText)
			newTestInstructionContainerAccordion := testCasesUiCanvasObject.newClickableAccordion(newTestInstructionAccordionItem, true, testCaseUuid, child.OriginalUuid) //widget.NewAccordion(newTestInstructionAccordionItem)
			newTestInstructionContainerAccordion.RemoveIndex(0)

			childObject, newTestInstructionAccordion = testCasesUiCanvasObject.recursiveMakeTestCaseGraphicalUIObject(child.Uuid, testCaseModelForUITree, newTestInstructionContainerAccordion, nodeTreeLevel+0.2, testCaseUuid)

			// Create the Accordion-object to hold information about the TestInstructionContainer
			newTestInstructionContainerAccordionItem := widget.NewAccordionItem(child.NodeName+" - "+child.Uuid, childObject)
			newTestInstructionContainerAccordion.Append(newTestInstructionContainerAccordionItem)
			newTestInstructionContainerAccordion.Open(0)

			// Create color to use
			newTransparentColor := color.RGBA{
				R: 0x00,
				G: 0x00,
				B: 0x00,
				A: 0x00,
			}

			// Create indentation rectangle to move node to right
			newIndentationRectangle := canvas.NewRectangle(newTransparentColor)
			newIndentationRectangle.StrokeColor = color.Black
			newIndentationRectangle.StrokeWidth = 0
			newIndentationRectangle.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize*nodeTreeLevel), float32(0)))
			newIndentationRectangleContainer := container.NewMax(newIndentationRectangle)

			// Create indentation within TestInstructionContainer
			/*newTestInstructionColorRectangle := canvas.NewRectangle(newTransparentColor)
			newTestInstructionColorRectangle.StrokeColor = color.Black
			newTestInstructionColorRectangle.StrokeWidth = 0
			newTestInstructionColorRectangle.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize * 0.3), float32(0)))
			testInstructionNodeTransparentContainer := container.NewMax(newTestInstructionColorRectangle)
			*/

			// Create the Horizontal node container object to be put on GUI
			nodeHContainer := container.NewHBox(newIndentationRectangleContainer, newTestInstructionContainerAccordion, layout.NewSpacer())

			// Create trailer rectangle for TestInstructionContainer
			newITrailerRectangle := canvas.NewRectangle(newTransparentColor)
			newITrailerRectangle.StrokeColor = color.Black
			newITrailerRectangle.StrokeWidth = 0
			newITrailerRectangle.SetMinSize(fyne.NewSize(1, 4))
			newITrailerRectangleContainer := container.NewMax(newITrailerRectangle)

			// Create the node container object to be put on GUI
			nodeContainer := container.NewVBox(nodeHContainer, newITrailerRectangleContainer)

			/*
				// Create background for TestInstructionContainer
				newTestInstructionContainerBackgroundFillColor := color.RGBA{
					R: 0x22,
					G: 0x22,
					B: 0x22,
					A: 0x22,
				}
				newTestInstructionContainerBackgroundBorderColor := color.RGBA{
					R: 0xFF,
					G: 0xFF,
					B: 0x22,
					A: 0xFF,
				}
				newTestInstructionContainerBackgroundRectangle := canvas.NewRectangle(newTestInstructionContainerBackgroundFillColor)
				newTestInstructionContainerBackgroundRectangle.StrokeWidth = 2
				newTestInstructionContainerBackgroundRectangle.StrokeColor = newTestInstructionContainerBackgroundBorderColor

				newTestInstructionContainerContainer := container.NewMax(newTestInstructionContainerBackgroundRectangle, nodeContainer)
			*/
			// Add the child
			nodeChildrenContainer.Add(nodeContainer)

		// Some kind of droppable Bond
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10ox_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10oxo_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10xo_BOND:

			newDroppableBondLabel := testCasesUiCanvasObject.DragNDropStateMachine.NewDroppableLabel(child.NodeName+" - "+child.Uuid, nodeTreeLevel, testCaseNodeRectangleSize, child.Uuid, testCaseUuid)
			newDroppableBondLabelContainer := container.NewMax(newDroppableBondLabel.BackgroundRectangle, newDroppableBondLabel)
			newDroppableBondLabel.Hide()

			// Create color to use
			newTransaparentColor := color.RGBA{
				R: 0x00,
				G: 0x00,
				B: 0x00,
				A: 0x00,
			}

			// Create indentation rectangle to move node to right
			newIndentationRectangle := canvas.NewRectangle(newTransaparentColor)
			newIndentationRectangle.StrokeColor = color.Black
			newIndentationRectangle.StrokeWidth = 0
			newIndentationRectangle.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize*nodeTreeLevel), float32(0)))
			newIndentationRectangleContainer := container.NewMax(newIndentationRectangle)

			// Create indentation within TestInstructionContainer
			newTestInstructionColorRectangle := canvas.NewRectangle(newTransaparentColor)
			newTestInstructionColorRectangle.StrokeColor = color.Black
			newTestInstructionColorRectangle.StrokeWidth = 0
			newTestInstructionColorRectangle.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize), float32(0)))
			testInstructionNodeTransparentContainer := container.NewMax(newTestInstructionColorRectangle)

			// Create the none container object to be put on GUI
			nodeContainer := container.NewHBox(newIndentationRectangleContainer, testInstructionNodeTransparentContainer, newDroppableBondLabelContainer, layout.NewSpacer())

			// Add the child
			nodeChildrenContainer.Add(nodeContainer)

		// Some kind of non-droppable Bond
		default:

		}

	}

	testCaseCanvasObject = nodeChildrenContainer

	return testCaseCanvasObject, newTestInstructionAccordion

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
		errorId := "93789d03-f728-40da-a6bd-78f8a96628a5"
		err = errors.New(fmt.Sprintf("color string with hexvalues, '%s', has not the correct lenght, '#AABBCCDDEE' in testcase with sourceUuid '%s' [ErrorID: %s]", rgbaHexString, errorId))

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
