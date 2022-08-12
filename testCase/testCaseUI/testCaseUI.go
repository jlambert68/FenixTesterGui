package testCaseUI

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// GenerateBaseCanvasObjectForTestCaseUI
// Create the Base-UI-canvas-object for the TestCases object. This base doesn't contain any specific TestCase-parts, and they will be added in other function
func (testCasesUiCanvasObject *TestCasesObjectForUiStruct) GenerateBaseCanvasObjectForTestCaseUI() (baseCanvasObjectForTestCaseUI fyne.CanvasObject) {

	// Create toolbar for TestCase area
	testCasesUiCanvasObject.TestCaseToolUIBar = widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentRedoIcon(), func() {
			fmt.Println("Reload GUI TestCase from testCaseModel")
		}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
			fmt.Println("Copy Node")
		}),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {
			fmt.Println("Cut Node")
		}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {
			fmt.Println("Past Node")
		}),
	)

	// Create The Tab-object, where each TestCase will have its own Tab
	testCasesUiCanvasObject.TestCasesTabs = container.NewAppTabs(
		container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)

	//tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	// Set the Tabs to be positioned in upper part of the object
	testCasesUiCanvasObject.TestCasesTabs.SetTabLocation(container.TabLocationLeading)

	// Load the TestCase testCaseModel UI area
	currentTestCaseModelAreaUI := uiServer.loadCurrentTestCaseModelAreaUI()

	// Load the TestCase attributes UI area
	currentTestCaseAttributesAreaUI := uiServer.loadCurrentTestCaseAttributesAreaUI()

	// Create the UI area for both TestCase testCaseModel UI and TestCase attributes UI
	testCaseAdaptiveSplitLayoutContainer := newAdaptiveSplit(currentTestCaseModelAreaUI, currentTestCaseAttributesAreaUI)

	// Create the complete TestCase UI area
	testCaseBorderedLayout := layout.NewBorderLayout(testCaseToolUIBar, nil, nil, nil)
	baseCanvasObjectForTestCaseUI = container.New(testCaseBorderedLayout, testCaseToolUIBar, testCaseAdaptiveSplitLayoutContainer)

	return baseCanvasObjectForTestCaseUI
}

// AddNewTestCaseToUIModel
// Add new TestCase to UI-model
func (testCasesUiCanvasObject *TestCasesObjectForUiStruct) AddNewTestCaseToUIModel(testCaseToBeAddedUuid string) (err error) {

	var tabName string

	// Get TestCase Name
	testCaseNameForTab, err := testCasesUiCanvasObject.TestCasesModelReference.GetTestCaseNameUuid(testCaseToBeAddedUuid)

	if err != nil {
		return err
	}

	// Generate short version of UUID to put in TestCase Name
	shortUUid := testCasesUiCanvasObject.TestCasesModelReference.GenerateShortUuidFromFullUuid(testCaseToBeAddedUuid)

	// Build the Tab-name
	if len(testCaseNameForTab) == 0 {
		// New TestCase
		tabName = "<New TestCase>" + " [" + shortUUid + "]"

	} else {
		// Existing TestCase
		tabName = testCaseNameForTab + " [" + shortUUid + "]"

	}

	// Generate the Textual Representation Area for the TestCase
	testCaseTextualModelArea := testCasesUiCanvasObject.generateTextualRepresentationAreaForTestCase()

	// Generate the Graphical Representation Area for the TestCase
	testCaseGraphicalModelArea := testCasesUiCanvasObject.generateGraphicalRepresentationAreaForTestCase()

	// Generate the BaseInformation Area for the TestCase
	testCaseBaseInformationArea := testCasesUiCanvasObject.generateBaseInformationAreaForTestCase()

	// Generate the MetaData Area for the TestCase
	testCaseMetaDataArea := testCasesUiCanvasObject.generateMetaDataAreaForTestCase()

	// Generate the TestCaseAttributes Area for the TestCase
	testCaseAttributesArea := testCasesUiCanvasObject.generateTestCaseAttributesAreaForTestCase()

	// Create the new Tab that should own all UI Objects for the TestCase
	newTestCaseTab := container.NewAppTabs(tabName, testCaseCanvasObject)

	return err
}

// Generate the Textual Representation Area for the TestCase
func (testCasesUiCanvasObject *TestCasesObjectForUiStruct) generateTextualRepresentationAreaForTestCase(testCaseUuid string) (testCaseTextualModelArea fyne.CanvasObject, err error) {

	// Get current TestCase
	currentTestCaseUiModel, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]

	if existsInMap == false {
		errorId := "3308ff4c-4f70-447a-94c6-18e55e3bc1fc"
		err = errors.New(fmt.Sprintf("testcase with uuid '%s' doesn't exist in 'TestCasesUiModelMap' [ErrorID: %s]", testCaseUuid, errorId))

		return nil, err
	}

	// Set initial values for TestCase Textual Structure - Simple
	currentTestCaseUiModel.currentTestCaseTextualStructure.currentTestCaseTextualStructureSimple = binding.NewString()
	currentTestCaseUiModel.currentTestCaseTextualStructure.currentTestCaseTextualStructureSimple.Set("'currentTestCaseTextualStructureSimple'")

	// Set initial values for TestCase Textual Structure - Complex
	currentTestCaseUiModel.currentTestCaseTextualStructure.currentTestCaseTextualStructureComplex = binding.NewString()
	currentTestCaseUiModel.currentTestCaseTextualStructure.currentTestCaseTextualStructureComplex.Set("'currentTestCaseTextualStructureComplex'")

	// Set initial values for TestCase Textual Structure - Simple
	currentTestCaseUiModel.currentTestCaseTextualStructure.currentTestCaseTextualStructureExtended = binding.NewString()
	currentTestCaseUiModel.currentTestCaseTextualStructure.currentTestCaseTextualStructureExtended.Set("'currentTestCaseTextualStructureExtended'")

	// Create the Labels to be used for showing the TestCase Textual Structures
	testCaseTextualStructureSimpleWidget := widget.NewLabelWithData(currentTestCaseUiModel.currentTestCaseTextualStructure.currentTestCaseTextualStructureSimple)
	testCaseTextualStructureComplexWidget := widget.NewLabelWithData(currentTestCaseUiModel.currentTestCaseTextualStructure.currentTestCaseTextualStructureComplex)
	testCaseTextualStructureExtendedWidget := widget.NewLabelWithData(currentTestCaseUiModel.currentTestCaseTextualStructure.currentTestCaseTextualStructureExtended)

	// Create GUI Canvas object to be used
	testCaseTextualModelArea = container.NewVBox(testCaseTextualStructureSimpleWidget, testCaseTextualStructureComplexWidget, testCaseTextualStructureExtendedWidget)

	return testCaseTextualModelArea, err
}

// Generate the Graphical Representation Area for the TestCase
func (testCasesUiCanvasObject *TestCasesObjectForUiStruct) generateGraphicalRepresentationAreaForTestCase() (testCaseGraphicalModelArea fyne.CanvasObject) {

	testCaseGraphicalModelArea = widget.NewLabel("'testCaseGraphicalModelArea'")

	return testCaseGraphicalModelArea
}

// Generate the BaseInformation Area for the TestCase
func (testCasesUiCanvasObject *TestCasesObjectForUiStruct) generateBaseInformationAreaForTestCase() (testCaseBaseInformationArea fyne.CanvasObject) {

	testCaseBaseInformationArea = widget.NewLabel("'testCaseBaseInformationArea'")

	return testCaseBaseInformationArea
}

// Generate the MetaData Area for the TestCase
func (testCasesUiCanvasObject *TestCasesObjectForUiStruct) generateMetaDataAreaForTestCase() (testCaseMetaDataArea fyne.CanvasObject) {

	testCaseMetaDataArea = widget.NewLabel("'testCaseMetaDataArea'")

	return testCaseMetaDataArea
}

// Generate the TestCaseAttributes Area for the TestCase
func (testCasesUiCanvasObject *TestCasesObjectForUiStruct) generateTestCaseAttributesAreaForTestCase() (testCaseAttributesArea fyne.CanvasObject) {

	testCaseAttributesArea = widget.NewLabel("'testCaseAttributesArea'")

	return testCaseAttributesArea
}

// UpdateTextualStructuresForTestCase
// Updates hte Textual Structures (Simple, Complex and Extended) for a specific TestCase
func (testCasesUiCanvasObject *TestCasesObjectForUiStruct) UpdateTextualStructuresForTestCase(
	testCaseUuid string,
	testCaseTextualStructureSimple string,
	testCaseTextualStructureComplex string,
	testCaseTextualStructureExtended string) (err error) {

	// Get current TestCase
	currentTestCaseUiModel, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]

	if existsInMap == false {
		errorId := "92b67dc9-73af-4669-97be-57ac9b1ea2ea"
		err = errors.New(fmt.Sprintf("testcase with uuid '%s' doesn't exist in 'TestCasesUiModelMap' [ErrorID: %s]", testCaseUuid, errorId))

		return err
	}

	// Update 'Binded' variables in UI-model for TestCase
	currentTestCaseUiModel.currentTestCaseTextualStructure.currentTestCaseTextualStructureSimple.Set(testCaseTextualStructureSimple)
	currentTestCaseUiModel.currentTestCaseTextualStructure.currentTestCaseTextualStructureComplex.Set(testCaseTextualStructureComplex)
	currentTestCaseUiModel.currentTestCaseTextualStructure.currentTestCaseTextualStructureExtended.Set(testCaseTextualStructureExtended)

	return err
}
