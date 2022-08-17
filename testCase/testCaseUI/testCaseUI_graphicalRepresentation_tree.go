package testCaseUI

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
)

// Generate the Graphical Representation Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateGraphicalRepresentationAreaForTestCase(testCaseUuid string) (testCaseGraphicalModelArea fyne.CanvasObject, testCaseGraphicalUITree *widget.Tree, testCaseGraphicalModelAreaAccordion *widget.Accordion, err error) {

	// Get current TestCase-UI-model
	_, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]

	if existsInMap == true {
		errorId := "a058d6d3-76bd-4667-802f-5e417f76ad26"
		err = errors.New(fmt.Sprintf("testcase-UI-model with uuid '%s' allready exist in 'TestCasesUiModelMap' [ErrorID: %s]", testCaseUuid, errorId))

		return nil, nil, nil, err
	}

	//testCaseGraphicalModelArea = widget.NewLabel("'testCaseGraphicalModelArea'")

	testCaseGraphicalUITree = testCasesUiCanvasObject.makeTestCaseGraphicalUITree(testCaseUuid)

	//tree.OpenAllBranches()

	//tree.Refresh()

	tempContainer := container.New(layout.NewGridWrapLayout(fyne.NewSize(1, 400)))

	treeExpandedContainer := container.New(layout.NewMaxLayout(), testCaseGraphicalUITree, tempContainer, layout.NewSpacer())

	// Create a Canvas Accordion type for grouping the Graphical Representation of the TestCase
	testCaseGraphicalModelAreaAccordionItem := widget.NewAccordionItem("Graphical Representation of the TestCase", treeExpandedContainer)
	testCaseGraphicalModelAreaAccordion = widget.NewAccordion(testCaseGraphicalModelAreaAccordionItem)

	testCaseGraphicalModelArea = container.NewVBox(testCaseGraphicalModelAreaAccordion)

	return testCaseGraphicalModelArea, testCaseGraphicalUITree, testCaseGraphicalModelAreaAccordion, err
}

func (testCasesUiCanvasObject *TestCasesUiModelStruct) makeTestCaseGraphicalUITree(testCaseUuid string) (tree *widget.Tree) {

	// Check if TestCase allready exists, shouldn't do that
	_, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]
	if existsInMap == true {
		errorId := "69447c68-b650-49bd-ab34-2d26964cea05"
		err := errors.New(fmt.Sprintf("testcase with uuid '%s' allready exist in map with all testcases [ErrorID: %s]", testCaseUuid, errorId))

		list := map[string][]string{
			"": {err.Error()},
		}

		tree := widget.NewTreeWithStrings(list)
		return tree
	}

	// Create Tree
	tree = &widget.Tree{
		ChildUIDs: func(uid string) []string {
			// treeViewModelMapForTestCase, _ := testCasesUiCanvasObject.TestCasesModelReference.GetTreeViewModelForTestCase(testCaseUuid)

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
			fmt.Println("CreateNode: ")
			//return newTappableLabel() //widget.NewLabel("Collection Widgets: ")

			//return widget.NewLabel("xxxx")

			nodeLabel := widget.NewLabel("This is just some text")

			// Create a Canvas Accordion type for grouping the TestCase Node and any node info to be displayed
			//testCaseNodeAreaAccordionItem := widget.NewAccordionItem("xxxx", nodeLabel)
			//testCaseTNodeAreaAccordion := widget.NewAccordion(testCaseNodeAreaAccordionItem)

			leftRectangle := canvas.NewRectangle(color.RGBA{0xff, 0x00, 0x00, 0xff})
			leftRectangle.StrokeColor = color.Black
			leftRectangle.StrokeWidth = 0

			leftRectangle.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize), float32(testCaseNodeRectangleSize)))

			greeRectangle := canvas.NewRectangle(color.Gray{0x44}) // RGBA{0x00, 0xFF, 0x00, 0xff})
			greeRectangle.StrokeColor = color.Black
			greeRectangle.StrokeWidth = 0
			labelContainer := container.NewMax(greeRectangle, nodeLabel)

			content := container.New(layout.NewBorderLayout(nil, nil, leftRectangle, nil),
				leftRectangle, labelContainer)

			return content
		},

		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			fmt.Println("UpdateNode: ", uid)
			/*
				_, ok := list[uid]
				if !ok {
					fyne.LogError("Missing tutorial panel: "+uid, nil)
					return
				}
			*/

			var (
				extractedNodeName string
				//valueRed          int64
				valueGreen          int64
				valueBlue           int64
				extractedColorRed   uint8
				extractedColorGreen uint8
				extractedColorBlue  uint8
				err                 error
			)

			// Extract Node-data
			treeNodeChildData := testCasesUiCanvasObject.TestCasesModelReference.GetTestCaseTreeNodeChildData(uid, testCaseUuid)

			// Secure that treeNodeChildData has correct content
			if uid != treeNodeChildData.Uuid {
				errorId := "2a398319-d1a5-4a8b-9270-deb29746ac6c"
				err = errors.New(fmt.Sprintf("Node-uid '%s' is not the same as UUID '%s' found in NodeData in testcase with uuid '%s' [ErrorID: %s]", uid, treeNodeChildData.Uuid, testCaseUuid, errorId))

				extractedNodeName = err.Error()
			} else {
				// Set Node Name
				extractedNodeName = treeNodeChildData.Uuid

				// Break up into correct Red-Green-Blue
				hexValueAsString := treeNodeChildData.NodeColor //"#FFEEFF"
				hexRed := hexValueAsString[1:3]
				hexGreen := hexValueAsString[3:5]
				hexBlue := hexValueAsString[5:7]

				valueRed, hexConvertionErr := strconv.ParseInt(hexRed, 16, 64)
				if hexConvertionErr != nil {
					// Hex conversion failed
					errorId := "162667e9-d35e-45be-b1b4-d07877a3cd2c"
					err = errors.New(fmt.Sprintf("hexConvertion for Color  failed with error message: '%s' in testcase with uuid '%s' [ErrorID: %s]", hexConvertionErr, testCaseUuid, errorId))

					extractedNodeName = err.Error()

				}

				// Only do if last Hex-conversion succeeded
				if hexConvertionErr == nil {
					valueGreen, hexConvertionErr = strconv.ParseInt(hexGreen, 16, 64)
					if hexConvertionErr != nil {
						// Hex conversion failed
						errorId := "b2b9fae0-30e3-49df-99d7-5662b78311a3"
						err = errors.New(fmt.Sprintf("hexConvertion for Color  failed with error message: '%s' in testcase with uuid '%s' [ErrorID: %s]", hexConvertionErr, testCaseUuid, errorId))

						extractedNodeName = err.Error()

					}
				}

				// Only do if last Hex-conversion succeeded
				if hexConvertionErr == nil {
					valueBlue, hexConvertionErr = strconv.ParseInt(hexBlue, 16, 64)
					if hexConvertionErr != nil {
						// Hex conversion failed
						errorId := "b2b9fae0-30e3-49df-99d7-5662b78311a3"
						err = errors.New(fmt.Sprintf("hexConvertion for Color  failed with error message: '%s' in testcase with uuid '%s' [ErrorID: %s]", hexConvertionErr, testCaseUuid, errorId))

						extractedNodeName = err.Error()

					}
				}

				// Only convert to color values to use if all Hex-conversion succeeded and no Other Error occured
				if hexConvertionErr == nil && err == nil {
					extractedColorRed = uint8(valueRed)
					extractedColorGreen = uint8(valueGreen)
					extractedColorBlue = uint8(valueBlue)

				} else {
					// Set color to be used when error
					extractedColorRed = uint8(255)
					extractedColorGreen = uint8(0)
					extractedColorBlue = uint8(255)
				}

			}
			//obj.(*tappableLabel).SetText(uid) //obj.(*widget.Label).SetText(uid) // + time.Now().String())
			//obj.(*widget.Label).SetText(uid)
			// Set the UUID as Node-Lable
			//obj.(*fyne.Container).Objects[1].Objects[1].(*widget.Label).SetText(uid)
			//obj.(*fyne.Container).Objects[1].(*widget.Accordion).Items[0].Title = uid
			obj.(*fyne.Container).Objects[1].(*fyne.Container).Objects[1].(*widget.Label).SetText(extractedNodeName)

			// Update Node color by replacing rectangle
			greeRectangle := canvas.NewRectangle(color.RGBA{
				R: extractedColorRed,
				G: extractedColorGreen,
				B: extractedColorBlue,
				A: 0xff,
			}) //color.Gray{0x66}) // RGBA{0x00, 0xFF, 0x00, 0xff})
			greeRectangle.StrokeColor = color.Black
			greeRectangle.StrokeWidth = 0
			obj.(*fyne.Container).Objects[1].(*fyne.Container).Objects[0] = greeRectangle

			// Set colored rectangle size to (labelHeight, labelHeight)
			//labelHeight := obj.(*fyne.Container).Objects[1].(*widget.Label).MinSize().Height
			//obj.(*fyne.Container).Objects[0].(*canvas.Rectangle).Resize(fyne.NewSize(labelHeight, labelHeight))
			//obj.Refresh()

			fmt.Println(tree.Size())
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
