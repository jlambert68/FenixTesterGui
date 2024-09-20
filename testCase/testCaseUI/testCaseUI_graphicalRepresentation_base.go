package testCaseUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
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

	myColor := color.RGBA{
		R: 0,
		G: 0,
		B: 0,
		A: 0,
	}

	var myRectangle *canvas.Rectangle
	myRectangle = canvas.NewRectangle(myColor)
	myRectangle.SetMinSize(fyne.NewSize(1, 200))

	graphicalTestCaseUIObject = testCasesUiCanvasObject.makeTestCaseGraphicalUIObject(testCaseUuid)
	//graphicalTestCaseUIObjectContainer := container.NewBorder(nil, nil, nil, nil, graphicalTestCaseUIObject)

	testCaseGraphicalModelAreaAccordionItem = widget.NewAccordionItem("Graphical Representation of the TestCase", graphicalTestCaseUIObject) //treeExpandedContainer)
	testCaseGraphicalModelAreaAccordion.Append(testCaseGraphicalModelAreaAccordionItem)
	testCaseGraphicalModelAreaAccordion.OpenAll()

	// Create a container for the Accordion
	accordionContainer := container.NewVBox(testCaseGraphicalModelAreaAccordion)

	// Wrap the container in a scrollContainer
	accordionScrollContainer := container.NewScroll(accordionContainer)

	// Set a minimum size for the vertical scroll to prevent collapsing
	//verticalScroll.SetMinSize(fyne.NewSize(300, 400)) // Adjust the minimum size based on your needs

	// Finally, wrap the vertical scroll in a horizontal scroll (if needed)
	//scrollContainer := container.NewHScroll(verticalScroll)
	//scrollContainer.Refresh()

	//canvasGraphicalRepresentationAccordionObject := container.NewScroll(testCaseGraphicalModelAreaAccordion)

	//testCaseGraphicalModelArea = container.NewBorder(nil, nil, nil, nil, testCaseGraphicalModelAreaAccordionNewHScroll)
	//testCaseGraphicalModelArea =  container.NewVBox(canvasGraphicalRepresentationAccordionObject)
	//testCaseGraphicalModelAreaWithScroll := container.NewVScroll(testCaseGraphicalModelArea)
	//testCaseGraphicalModelAreaWithScroll.SetMinSize(fyne.NewSize(testCaseGraphicalModelAreaWithScroll.Size().Width, 300))

	return accordionScrollContainer, graphicalTestCaseUIObject, testCaseGraphicalModelAreaAccordion, err
}
