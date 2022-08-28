package testCaseUI

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Generate the Textual Representation Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateNewTextualRepresentationAreaForTestCase(testCaseUuid string) (newTestCaseTextualStructure testCaseTextualStructureStruct, canvasTextualRepresentationAccordionObject fyne.CanvasObject, err error) {

	// Get current TestCase-UI-model
	_, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]

	if existsInMap == true {
		errorId := "3308ff4c-4f70-447a-94c6-18e55e3bc1fc"
		err = errors.New(fmt.Sprintf("testcase-UI-model with sourceUuid '%s' allready exist in 'TestCasesUiModelMap' [ErrorID: %s]", testCaseUuid, errorId))

		return testCaseTextualStructureStruct{}, nil, err
	}

	// Create a new Textual Structure to be used in TestCase-UI-model
	newTestCaseTextualStructure = testCaseTextualStructureStruct{}

	// Set initial values for TestCase Textual Structure - Simple
	newTestCaseTextualStructure.currentTestCaseTextualStructureSimple = binding.NewString()
	newTestCaseTextualStructure.currentTestCaseTextualStructureSimple.Set("'currentTestCaseTextualStructureSimple'")

	// Set initial values for TestCase Textual Structure - Complex
	newTestCaseTextualStructure.currentTestCaseTextualStructureComplex = binding.NewString()
	newTestCaseTextualStructure.currentTestCaseTextualStructureComplex.Set("'currentTestCaseTextualStructureComplex'")

	// Set initial values for TestCase Textual Structure - Simple
	newTestCaseTextualStructure.currentTestCaseTextualStructureExtended = binding.NewString()
	newTestCaseTextualStructure.currentTestCaseTextualStructureExtended.Set("'currentTestCaseTextualStructureExtended'")

	// Create the Labels to be used for showing the TestCase Textual Structures
	testCaseTextualStructureSimpleWidget := widget.NewLabelWithData(newTestCaseTextualStructure.currentTestCaseTextualStructureSimple)
	testCaseTextualStructureComplexWidget := widget.NewLabelWithData(newTestCaseTextualStructure.currentTestCaseTextualStructureComplex)
	testCaseTextualStructureExtendedWidget := widget.NewLabelWithData(newTestCaseTextualStructure.currentTestCaseTextualStructureExtended)

	textualRepresentationGrid := container.New(layout.NewFormLayout(),
		widget.NewLabel("Simple"),
		testCaseTextualStructureSimpleWidget,
		widget.NewLabel("Complex"),
		testCaseTextualStructureComplexWidget,
		widget.NewLabel("Extended"),
		testCaseTextualStructureExtendedWidget)

	// Create GUI Canvas object to be used
	testCaseTextualModelArea := container.NewVBox(textualRepresentationGrid)

	// Create a Canvas Accordion type for grouping the Textual Representations
	testCaseTextualModelAreaAccordionItem := widget.NewAccordionItem("Textual Representation of the TestCase", testCaseTextualModelArea)
	testCaseTextualModelAreaAccordion := widget.NewAccordion(testCaseTextualModelAreaAccordionItem)

	canvasTextualRepresentationAccordionObject = container.NewVBox(testCaseTextualModelAreaAccordion)

	// Save the Accordion object in Textual Structure
	newTestCaseTextualStructure.currentTestCaseGraphicalAccordionObject = testCaseTextualModelAreaAccordion

	return newTestCaseTextualStructure, canvasTextualRepresentationAccordionObject, err
}

// UpdateTextualStructuresForTestCase
// Updates hte Textual Structures (Simple, Complex and Extended) for a specific TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) UpdateTextualStructuresForTestCase(
	testCaseUuid string,
	testCaseTextualStructureSimple string,
	testCaseTextualStructureComplex string,
	testCaseTextualStructureExtended string) (err error) {

	// Get current TestCase
	currentTestCaseUiModel, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]

	if existsInMap == false {
		errorId := "92b67dc9-73af-4669-97be-57ac9b1ea2ea"
		err = errors.New(fmt.Sprintf("testcase with sourceUuid '%s' doesn't exist in 'TestCasesUiModelMap' [ErrorID: %s]", testCaseUuid, errorId))

		return err
	}

	// Update 'Binded' variables in UI-model for TestCase
	currentTestCaseUiModel.currentTestCaseTextualStructure.currentTestCaseTextualStructureSimple.Set(testCaseTextualStructureSimple)
	currentTestCaseUiModel.currentTestCaseTextualStructure.currentTestCaseTextualStructureComplex.Set(testCaseTextualStructureComplex)
	currentTestCaseUiModel.currentTestCaseTextualStructure.currentTestCaseTextualStructureExtended.Set(testCaseTextualStructureExtended)

	currentTestCaseUiModel.TestCaseTextualModelArea.Refresh()

	return err
}
