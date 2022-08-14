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
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateGraphicalRepresentationAreaForTestCase(testCaseUuid string) (testCaseGraphicalModelArea fyne.CanvasObject, err error) {

	// Get current TestCase-UI-model
	_, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]

	if existsInMap == true {
		errorId := "a058d6d3-76bd-4667-802f-5e417f76ad26"
		err = errors.New(fmt.Sprintf("testcase-UI-model with uuid '%s' allready exist in 'TestCasesUiModelMap' [ErrorID: %s]", testCaseUuid, errorId))

		return nil, err
	}

	//testCaseGraphicalModelArea = widget.NewLabel("'testCaseGraphicalModelArea'")

	testCaseGraphicalModelArea = testCasesUiCanvasObject.makeTestCaseGraphicalUITree(testCaseUuid)

	// Create a Canvas Accordion type for grouping the Graphical Representation of the TestCase
	testCaseGraphicalModelAreaAccordionItem := widget.NewAccordionItem("Graphical Representation of the TestCase", testCaseGraphicalModelArea)
	testCaseGraphicalModelAreaAccordion := widget.NewAccordion(testCaseGraphicalModelAreaAccordionItem)

	return testCaseGraphicalModelAreaAccordion, err
}

var list map[string][]string
var tree *widget.Tree

func (testCasesUiCanvasObject *TestCasesUiModelStruct) makeTestCaseGraphicalUITree(testCaseUuid string) fyne.CanvasObject {
	list = map[string][]string{
		"":  {"A"},
		"A": {"B", "D"},
		"B": {"C"},
		"C": {"abc"},
		"D": {"E"},
		"E": {"F", "G"},
	}
	/*
		list, err := testCasesUiCanvasObject.TestCasesModelReference.GetTreeViewModelForTestCase()
		if err != nil {
			return nil
		}

		tree = widget.NewTreeWithStrings(list)
		tree.OnSelected = func(id string) {
			fmt.Printf("Tree node selected: %s", id)

		}
		tree.OnUnselected = func(id string) {
			fmt.Printf("Tree node unselected: %s", id)
		}
	*/

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

	tree.OpenAllBranches()

	tree.Refresh()

	tempContainer := container.New(layout.NewGridWrapLayout(fyne.NewSize(1, 400)))

	treeExpandedContainer := container.New(layout.NewMaxLayout(), tree, tempContainer, layout.NewSpacer())

	return treeExpandedContainer

}
