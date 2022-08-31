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

// Generate the Graphical Representation Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateGraphicalRepresentationAreaForTestCase(testCaseUuid string) (testCaseGraphicalModelArea fyne.CanvasObject, graphicalTestCaseUIObject fyne.CanvasObject, testCaseGraphicalModelAreaAccordion2 *widget.Accordion, err error) {

	/*
		// Get current TestCase-UI-model
		_, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]

		if existsInMap == true {
			errorId := "a058d6d3-76bd-4667-802f-5e417f76ad26"
			err = errors.New(fmt.Sprintf("testcase-UI-model with sourceUuid '%s' allready exist in 'TestCasesUiModelMap' [ErrorID: %s]", testCaseUuid, errorId))

			return nil, nil, nil, err
		}

	*/

	//testCaseGraphicalUITree = testCasesUiCanvasObject.makeTestCaseGraphicalUITree(testCaseUuid)

	//tempContainer := container.New(layout.NewGridWrapLayout(fyne.NewSize(1, 400)))

	//treeExpandedContainer := container.New(layout.NewMaxLayout(), testCaseGraphicalUITree, tempContainer, layout.NewSpacer())
	//mytempText := widget.NewLabel("Temp text")

	// Create a Canvas Accordion type for grouping the Graphical Representation of the TestCase
	testCaseGraphicalModelAreaAccordionItem := widget.NewAccordionItem("Graphical Representation of the TestCase", widget.NewLabel("temp")) //treeExpandedContainer)
	testCaseGraphicalModelAreaAccordion := widget.NewAccordion(testCaseGraphicalModelAreaAccordionItem)
	testCaseGraphicalModelAreaAccordion.RemoveIndex(0)

	graphicalTestCaseUIObject = testCasesUiCanvasObject.makeTestCaseGraphicalUIObject(testCaseUuid, testCaseGraphicalModelAreaAccordion)

	testCaseGraphicalModelAreaAccordionItem = widget.NewAccordionItem("Graphical Representation of the TestCase", graphicalTestCaseUIObject) //treeExpandedContainer)
	testCaseGraphicalModelAreaAccordion.Append(testCaseGraphicalModelAreaAccordionItem)

	testCaseGraphicalModelArea = container.NewVBox(testCaseGraphicalModelAreaAccordion)

	return testCaseGraphicalModelArea, graphicalTestCaseUIObject, testCaseGraphicalModelAreaAccordion, err
}

func (testCasesUiCanvasObject *TestCasesUiModelStruct) makeTestCaseGraphicalUITree(testCaseUuid string) (tree *widget.Tree) {

	// Check if TestCase already exists, shouldn't do that
	_, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]
	if existsInMap == true {
		errorId := "69447c68-b650-49bd-ab34-2d26964cea05"
		err := errors.New(fmt.Sprintf("testcase with sourceUuid '%s' allready exist in map with all testcases [ErrorID: %s]", testCaseUuid, errorId))

		list := map[string][]string{
			"": {err.Error()},
		}

		tree := widget.NewTreeWithStrings(list)
		return tree
	}

	// Create Tree
	tree = &widget.Tree{
		ChildUIDs: func(uid string) []string {

			// Create slice with children UUIDs
			var childrenUuidSlice []string

			// Get the array
			childrenUuidSlice = testCasesUiCanvasObject.TestCasesModelReference.GetArrayOfTestCaseTreeNodeChildrenData(uid, testCaseUuid)

			return childrenUuidSlice
		},
		IsBranch: func(uid string) bool {
			treeViewModelForTestCase, _ := testCasesUiCanvasObject.TestCasesModelReference.GetTreeViewModelForTestCase(testCaseUuid)
			children, ok := treeViewModelForTestCase[uid]

			return ok && len(children) > 0
		},

		CreateNode: func(branch bool) fyne.CanvasObject {

			//nodeLabel := widget.NewLabel("This is just some text")

			testInstructionNodeColorRectangle := canvas.NewRectangle(color.RGBA{0xff, 0x00, 0x00, 0xff})
			testInstructionNodeColorRectangle.StrokeColor = color.Black
			testInstructionNodeColorRectangle.StrokeWidth = 0

			testInstructionNodeColorRectangle.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize*0.5), float32(testCaseNodeRectangleSize*0.5)))
			testInstructionNodeColorContainer := container.NewMax(testInstructionNodeColorRectangle)

			/*
				nodeTextBackgroundColorectangle := canvas.NewRectangle(color.Gray{0x44}) // RGBA{0x00, 0xFF, 0x00, 0xff})
				nodeTextBackgroundColorectangle.StrokeColor = color.Black
				nodeTextBackgroundColorectangle.StrokeWidth = 00
				labelContainer := container.NewMax(nodeTextBackgroundColorectangle, nodeLabel)
			*/

			newDroppableLabel := testCasesUiCanvasObject.DragNDropStateMachine.NewDroppableLabel("This is just some text", nil, 0, 0, nil)
			newDroppableContainer := container.NewMax(newDroppableLabel.BackgroundRectangle, newDroppableLabel)

			content := container.NewHBox(testInstructionNodeColorContainer, newDroppableContainer) //labelContainer)

			return content
		},

		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {

			var (
				extractedNodeName string
				err               error
			)

			// Extract Node-data
			treeNodeChildData := testCasesUiCanvasObject.TestCasesModelReference.GetTestCaseTreeNodeChildData(uid, testCaseUuid)

			// Secure that treeNodeChildData has correct content
			if uid != treeNodeChildData.Uuid {
				errorId := "2a398319-d1a5-4a8b-9270-deb29746ac6c"
				err = errors.New(fmt.Sprintf("Node-uid '%s' is not the same as UUID '%s' found in NodeData in testcase with sourceUuid '%s' [ErrorID: %s]", uid, treeNodeChildData.Uuid, testCaseUuid, errorId))

				extractedNodeName = err.Error()
			} else {

				// Set Node Name
				extractedNodeName = treeNodeChildData.NodeName + " - " + treeNodeChildData.Uuid
			}
			// Break up into correct Red-Green-Blue
			hexValueAsString := treeNodeChildData.NodeColor //"#FFEEFF"

			// Convert Color
			extractedNodeColor, err := testCasesUiCanvasObject.convertRGBAHexStringIntoRGBAColor(hexValueAsString)
			if err != nil {
				extractedNodeName = err.Error()
			}

			/*
				// Extract if node is droppable
				var nodeIsDroppable bool
				draggedUuid := testCasesUiCanvasObject.DragNDropStateMachine.SourceUuid
				if draggedUuid == "" {
					nodeIsDroppable = false
				} else {
					// Extract Dragged nodes type
					var elementType fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum

					switch testCasesUiCanvasObject.DragNDropStateMachine.SourceType {
					case 1: //gui.TestInstruction:
						elementType = fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION

					case 2: //gui.TestInstructionContainer:
						elementType = fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

					default:
						errorId := "c6e74d79-9268-4c54-8338-2ed23539a5a2"
						err = errors.New(fmt.Sprintf("unknown Source Type [ErrorID: %s]", errorId))

						extractedNodeName = err.Error()
					}

					nodeIsDroppable, err = testCasesUiCanvasObject.CommandAndRuleEngineReference.VerifyIfElementCanBeSwapped(testCaseUuid, treeNodeChildData.Uuid, elementType)
					if err != nil {
						extractedNodeName = err.Error()
					}


				}
			*/

			// Extract TestInstruction Type Color and change
			//hexValueForTestInstructionNodeColorAsString := treeNodeChildData.TestInstructionTypeColor
			if treeNodeChildData.NodeTypeEnum == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
				treeNodeChildData.NodeTypeEnum == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE {
				newRectangleBackgroundWithColorForTestInstructionType := canvas.NewRectangle(color.RGBA{
					R: 255,
					G: 255,
					B: 0,
					A: 0x88,
				})
				newRectangleBackgroundWithColorForTestInstructionType.StrokeColor = color.Black
				newRectangleBackgroundWithColorForTestInstructionType.StrokeWidth = 0

				newRectangleBackgroundWithColorForTestInstructionType.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize), float32(testCaseNodeRectangleSize)))
				obj.(*fyne.Container).Objects[0].(*fyne.Container).Objects[0] = newRectangleBackgroundWithColorForTestInstructionType
			} else {
				// Not a TestInstruction so set it to be Invisible
				newRectangleBackgroundWithColorForTestInstructionType := canvas.NewRectangle(color.RGBA{
					R: 0,
					G: 0,
					B: 0,
					A: 0x00,
				})
				newRectangleBackgroundWithColorForTestInstructionType.StrokeColor = color.Black
				newRectangleBackgroundWithColorForTestInstructionType.StrokeWidth = 0

				newRectangleBackgroundWithColorForTestInstructionType.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize), float32(testCaseNodeRectangleSize)))
				obj.(*fyne.Container).Objects[0].(*fyne.Container).Objects[0] = newRectangleBackgroundWithColorForTestInstructionType

			}
			//obj.(*tappableLabel).SetText(uid) //obj.(*widget.Label).SetText(uid) // + time.Now().String())
			//obj.(*widget.Label).SetText(uid)
			// Set the UUID as Node-Lable
			//obj.(*fyne.Container).Objects[1].Objects[1].(*widget.Label).SetText(uid)
			//obj.(*fyne.Container).Objects[1].(*widget.Accordion).Items[0].Title = uid

			//obj.(*fyne.Container).Objects[1].(*fyne.Container).Objects[1].(*widget.Label).SetText(extractedNodeName)
			obj.(*fyne.Container).Objects[1].(*fyne.Container).Objects[1].(*testUIDragNDropStatemachine.DroppableLabel).SetText(extractedNodeName)
			obj.(*fyne.Container).Objects[1].(*fyne.Container).Objects[1].(*testUIDragNDropStatemachine.DroppableLabel).TargetUuid = extractedNodeName
			//obj.(*fyne.Container).Objects[1].(*testUIDragNDropStatemachine.DroppableLabel).IsDroppable = nodeIsDroppable

			//*testUIDragNDropStatemachine.DroppableLabel

			// Update Node color by replacing rectangle
			newRectangleBackgroundWithColor := canvas.NewRectangle(extractedNodeColor)
			newRectangleBackgroundWithColor.StrokeColor = color.Black
			newRectangleBackgroundWithColor.StrokeWidth = 0
			//obj.(*fyne.Container).Objects[1].(*fyne.Container).Objects[0] = newRectangleBackgroundWithColor

			// Set colored rectangle size to (labelHeight, labelHeight)
			//labelHeight := obj.(*fyne.Container).Objects[1].(*widget.Label).MinSize().Height
			//obj.(*fyne.Container).Objects[0].(*canvas.Rectangle).Resize(fyne.NewSize(labelHeight, labelHeight))
			obj.Refresh()

			//fmt.Println(tree.Size())
		},

		OnSelected: func(uid string) {
			fmt.Println(uid)
			//fmt.Println(uid, uiServer.availableBuildingBlocksModel.getAvailableBuildingBlocksModel()[uid])
			//uiServer.availableBuildingBlocksModel.clickedNodeName = uid

			//if t, ok := list[uid]; ok {
			//	fmt.Println(tree.Root)
			//	fmt.Println(t)

			//}
		},
	}

	return tree

}

/*
type testCaseTreeNodeStruct struct {
	widget.Label
}

func newTestCaseTreeNode() *testCaseTreeNodeStruct {
	treeNode := &testCaseTreeNodeStruct{}
	treeNode.ExtendBaseWidget(treeNode)
	//label.ExtendBaseWidget(label)
	//icon.SetResource(res)

	left := canvas.NewText("left", color.RGBA{
		R: 0xFF,
		G: 0,
		B: 0,
		A: 0,
	})
	middle := widget.NewLabel("xxxx")
	content := container.New(layout.NewBorderLayout(nil, nil, left, nil),
		left, middle)

	return treeNode
}


*/

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

// Generates the graphical structure for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) makeTestCaseGraphicalUIObject(testCaseUuid string, testCaseGraphicalModelAreaAccordion *widget.Accordion) (testCaseCanvasObject fyne.CanvasObject) {

	treeViewModelForTestCase, err := testCasesUiCanvasObject.TestCasesModelReference.GetTreeViewModelForTestCase(testCaseUuid)

	if err != nil {
		errText := widget.NewLabel(err.Error())
		testCaseCanvasObject = container.NewVBox(errText)

		return testCaseCanvasObject

	}

	// Clear state machine for Drag N Drop
	testCasesUiCanvasObject.DragNDropStateMachine = testUIDragNDropStatemachine.StateMachineDragAndDropStruct{}

	// Start processing model for TestCase
	testCaseCanvasObject, _ = testCasesUiCanvasObject.recursiveMakeTestCaseGraphicalUIObject("", &treeViewModelForTestCase, testCaseGraphicalModelAreaAccordion, 1, testCaseGraphicalModelAreaAccordion)

	return testCaseCanvasObject

}

// Generates the graphical structure for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) recursiveMakeTestCaseGraphicalUIObject(uuid string, testCaseModelForUITree *map[string][]testCaseModel.TestCaseModelAdaptedForUiTreeDataStruct, firstAccordion *widget.Accordion, nodeTreeLevel float32, topLevelAccordian *widget.Accordion) (testCaseCanvasObject fyne.CanvasObject, newTestInstructionAccordion2 *widget.Accordion) {

	var childObject fyne.CanvasObject
	var newTestInstructionAccordion *widget.Accordion
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

			// Create rectangle to show TestInstruction-color
			newTestInstructionColorRectangle := canvas.NewRectangle(rectangleColor)
			newTestInstructionColorRectangle.StrokeColor = color.Black
			newTestInstructionColorRectangle.StrokeWidth = 0
			newTestInstructionColorRectangle.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize), float32(testCaseNodeRectangleSize)))
			testInstructionNodeColorContainer := container.NewMax(newTestInstructionColorRectangle)

			// Create the Accordion-object to hold information about the TestInstruction
			dummyText := widget.NewLabel("this is just a dummy text and might show other TestInstruction-attributes later on")
			newTestInstructionAccordionItem := widget.NewAccordionItem(nodeName, dummyText)
			newTestInstructionAccordion = widget.NewAccordion(newTestInstructionAccordionItem)

			// Create the container object to be put on GUI
			nodeContainer := container.NewHBox(newIndentationRectangleContainer, testInstructionNodeColorContainer, newTestInstructionAccordion, layout.NewSpacer())

			// Add the child
			nodeChildrenContainer.Add(nodeContainer)

		// Some kind of TestInstructionContainer
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE:

			// Create the Accordion-object to hold information about the TestInstructionContainer
			dummyText := widget.NewLabel("this is just a dummy text")
			newTestInstructionAccordionItem := widget.NewAccordionItem("DummyText", dummyText)
			newTestInstructionContainerAccordion := widget.NewAccordion(newTestInstructionAccordionItem)
			newTestInstructionContainerAccordion.RemoveIndex(0)

			childObject, newTestInstructionAccordion = testCasesUiCanvasObject.recursiveMakeTestCaseGraphicalUIObject(child.Uuid, testCaseModelForUITree, newTestInstructionContainerAccordion, nodeTreeLevel+0.2, firstAccordion)

			// Create the Accordion-object to hold information about the TestInstructionContainer
			newTestInstructionContainerAccordionItem := widget.NewAccordionItem(child.NodeName+" - "+child.Uuid, childObject)
			newTestInstructionContainerAccordion.Append(newTestInstructionContainerAccordionItem)
			newTestInstructionContainerAccordion.Open(0)

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
			/*newTestInstructionColorRectangle := canvas.NewRectangle(newTransaparentColor)
			newTestInstructionColorRectangle.StrokeColor = color.Black
			newTestInstructionColorRectangle.StrokeWidth = 0
			newTestInstructionColorRectangle.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize * 0.3), float32(0)))
			testInstructionNodeTransparentContainer := container.NewMax(newTestInstructionColorRectangle)
			*/

			// Create the cone container object to be put on GUI
			nodeContainer := container.NewHBox(newIndentationRectangleContainer, newTestInstructionContainerAccordion, layout.NewSpacer())

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

			newDroppableBondLabel := testCasesUiCanvasObject.DragNDropStateMachine.NewDroppableLabel(child.NodeName+" - "+child.Uuid, newTestInstructionAccordion, nodeTreeLevel, testCaseNodeRectangleSize, firstAccordion)
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

			// Create the cone container object to be put on GUI
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
