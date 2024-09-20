package testCaseUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Generate the Graphical Representation Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateGraphicalRepresentationAreaForTestCase(
	testCaseUuid string) (
	testCaseGraphicalModelArea fyne.CanvasObject,
	graphicalTestCaseUIObject fyne.CanvasObject,
	testCaseGraphicalModelAreaAccordion2 *widget.Accordion,
	err error) {

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

	graphicalTestCaseUIObject = testCasesUiCanvasObject.makeTestCaseGraphicalUIObject(testCaseUuid)

	testCaseGraphicalModelAreaAccordionItem = widget.NewAccordionItem("Graphical Representation of the TestCase", graphicalTestCaseUIObject) //treeExpandedContainer)
	testCaseGraphicalModelAreaAccordion.Append(testCaseGraphicalModelAreaAccordionItem)

	//canvasGraphicalRepresentationAccordionObject := container.NewScroll(testCaseGraphicalModelAreaAccordion)

	//testCaseGraphicalModelArea = container.NewBorder(nil, nil, nil, nil, canvasGraphicalRepresentationAccordionObject)
	//testCaseGraphicalModelArea =  container.NewVBox(canvasGraphicalRepresentationAccordionObject)
	//testCaseGraphicalModelAreaWithScroll := container.NewHScroll(testCaseGraphicalModelArea)

	return testCaseGraphicalModelAreaAccordion, graphicalTestCaseUIObject, testCaseGraphicalModelAreaAccordion, err
}
