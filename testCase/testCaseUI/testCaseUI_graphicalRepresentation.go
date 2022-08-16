package testCaseUI

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
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
			treeViewModelForTestCase, _ := testCasesUiCanvasObject.TestCasesModelReference.GetTreeViewModelForTestCase(testCaseUuid)
			return treeViewModelForTestCase[uid]
		},
		IsBranch: func(uid string) bool {
			treeViewModelForTestCase, _ := testCasesUiCanvasObject.TestCasesModelReference.GetTreeViewModelForTestCase(testCaseUuid)
			children, ok := treeViewModelForTestCase[uid]

			return ok && len(children) > 0
		},

		CreateNode: func(branch bool) fyne.CanvasObject {
			fmt.Println("CreateNode: ")
			//return newTappableLabel() //widget.NewLabel("Collection Widgets: ")
			return widget.NewLabel("xxxx")
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
			//obj.(*tappableLabel).SetText(uid) //obj.(*widget.Label).SetText(uid) // + time.Now().String())
			obj.(*widget.Label).SetText(uid)
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
