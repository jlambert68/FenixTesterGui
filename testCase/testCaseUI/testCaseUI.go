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
func (testCasesUiCanvasObject *TestCasesUiModelStruct) GenerateBaseCanvasObjectForTestCaseUI() (baseCanvasObjectForTestCaseUI fyne.CanvasObject) {

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
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")),
	)

	// Set the Tabs to be positioned in upper part of the object
	testCasesUiCanvasObject.TestCasesTabs.SetTabLocation(container.TabLocationTop)

	// Create the complete TestCase UI area
	testCaseBorderedLayout := layout.NewBorderLayout(testCasesUiCanvasObject.TestCaseToolUIBar, nil, nil, nil)
	baseCanvasObjectForTestCaseUI = container.New(testCaseBorderedLayout, testCasesUiCanvasObject.TestCaseToolUIBar, testCasesUiCanvasObject.TestCasesTabs)

	return baseCanvasObjectForTestCaseUI
}

// Generate a new TestCase UI-model
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateNewTestCaseTabObject(testCaseToBeAddedUuid string) (err error) {

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

	// Initiate TestCase UI-model
	testCaseGraphicalAreas := testCaseGraphicalAreasStruct{}

	// Generate the Textual Representation Area for the TestCase
	testCaseTextualModelArea, err := testCasesUiCanvasObject.generateNewTextualRepresentationAreaForTestCase(testCaseToBeAddedUuid)

	if err != nil {
		return err
	}
	// Add newly created Textual Representation Area to object for all graphical parts of one TestCase
	testCaseGraphicalAreas.TestCaseTextualModelArea = testCaseTextualModelArea

	// Generate the Graphical Representation Area for the TestCase
	testCaseGraphicalModelArea, err := testCasesUiCanvasObject.generateGraphicalRepresentationAreaForTestCase(testCaseToBeAddedUuid)

	if err != nil {
		return err
	}

	// Add newly created Graphical Representation Area to object for all graphical parts of one TestCase
	testCaseGraphicalAreas.TestCaseGraphicalModelArea = testCaseGraphicalModelArea

	// Generate the BaseInformation Area for the TestCase
	testCaseBaseInformationArea, err := testCasesUiCanvasObject.generateBaseInformationAreaForTestCase(testCaseToBeAddedUuid)

	if err != nil {
		return err
	}

	// Add newly created BaseInformation Area to object for all graphical parts of one TestCase
	testCaseGraphicalAreas.TestCaseBaseInformationArea = testCaseBaseInformationArea

	// Generate the MetaData Area for the TestCase
	testCaseMetaDataArea, err := testCasesUiCanvasObject.generateMetaDataAreaForTestCase(testCaseToBeAddedUuid)

	if err != nil {
		return err
	}

	// Add newly created MetaData Area to object for all graphical parts of one TestCase
	testCaseGraphicalAreas.TestCaseMetaDataArea = testCaseMetaDataArea

	// Generate the TestCaseAttributes Area for the TestCase
	testCaseAttributesArea, err := testCasesUiCanvasObject.generateTestCaseAttributesAreaForTestCase(testCaseToBeAddedUuid)

	if err != nil {
		return err
	}
	// Add newly created TestCaseAttributes Area to object for all graphical parts of one TestCase
	testCaseGraphicalAreas.TestCaseAttributesArea = testCaseAttributesArea

	// Save TestCase-UI-model in UI-modelMap
	// Check if TestCase allready exists in TestCase-UI-model
	_, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseToBeAddedUuid]

	if existsInMap == true {
		errorId := "db34dee8-1b23-425c-868a-2747959ec682"
		err = errors.New(fmt.Sprintf("testcase-UI-model with uuid '%s' allready exist in 'TestCasesUiModelMap' [ErrorID: %s]", testCaseToBeAddedUuid, errorId))

		return err
	}

	// Create canvas-object for Textual and Graphical Representation
	textualAndGraphicalRepresentations := container.NewVBox(testCaseGraphicalAreas.TestCaseTextualModelArea, testCaseGraphicalAreas.TestCaseGraphicalModelArea)

	// Create canvas-object for BaseInformation, MetaData and TestCaseAttributes
	baseInformationMetaDataTestCaseAttributes := container.NewVBox(testCaseGraphicalAreas.TestCaseBaseInformationArea, testCaseGraphicalAreas.TestCaseMetaDataArea, testCaseGraphicalAreas.TestCaseAttributesArea)

	// Create the UI area for all parts of one TestCase
	testCaseAdaptiveSplitContainer := newAdaptiveSplit(textualAndGraphicalRepresentations, baseInformationMetaDataTestCaseAttributes)

	// Create a new Tab-object
	newTestCaseTabObject := container.NewTabItem(tabName, testCaseAdaptiveSplitContainer)

	// Add tab to existing Tab-object
	testCasesUiCanvasObject.TestCasesTabs.Append(newTestCaseTabObject)

	// Set focus on newly created Tab
	testCasesUiCanvasObject.TestCasesTabs.Select(newTestCaseTabObject)

	return err
}

// Generate the Textual Representation Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateNewTextualRepresentationAreaForTestCase(testCaseUuid string) (testCaseTextualModelArea fyne.CanvasObject, err error) {

	// Get current TestCase-UI-model
	_, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]

	if existsInMap == true {
		errorId := "3308ff4c-4f70-447a-94c6-18e55e3bc1fc"
		err = errors.New(fmt.Sprintf("testcase-UI-model with uuid '%s' allready exist in 'TestCasesUiModelMap' [ErrorID: %s]", testCaseUuid, errorId))

		return nil, err
	}

	// Create a new Textual Structure to be used in TestCase-UI-model
	newTestCaseTextualStructure := currentTestCaseTextualStructureStruct{}

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

	// Create GUI Canvas object to be used
	testCaseTextualModelArea = container.NewVBox(testCaseTextualStructureSimpleWidget, testCaseTextualStructureComplexWidget, testCaseTextualStructureExtendedWidget)

	return testCaseTextualModelArea, err
}

// Generate the Graphical Representation Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateGraphicalRepresentationAreaForTestCase(testCaseUuid string) (testCaseGraphicalModelArea fyne.CanvasObject, err error) {

	// Get current TestCase-UI-model
	_, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]

	if existsInMap == true {
		errorId := "a058d6d3-76bd-4667-802f-5e417f76ad26"
		err = errors.New(fmt.Sprintf("testcase-UI-model with uuid '%s' allready exist in 'TestCasesUiModelMap' [ErrorID: %s]", testCaseUuid, errorId))

		return nil, err
	}

	testCaseGraphicalModelArea = widget.NewLabel("'testCaseGraphicalModelArea'")

	return testCaseGraphicalModelArea, err
}

// Generate the BaseInformation Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateBaseInformationAreaForTestCase(testCaseUuid string) (testCaseBaseInformationArea fyne.CanvasObject, err error) {

	// Get current TestCase-UI-model
	_, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]

	if existsInMap == true {
		errorId := "4b062436-590a-4f2a-9004-181f3f575a4b"
		err = errors.New(fmt.Sprintf("testcase-UI-model with uuid '%s' allready exist in 'TestCasesUiModelMap' [ErrorID: %s]", testCaseUuid, errorId))

		return nil, err
	}

	testCaseBaseInformationArea = widget.NewLabel("'testCaseBaseInformationArea'")

	return testCaseBaseInformationArea, err
}

// Generate the MetaData Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateMetaDataAreaForTestCase(testCaseUuid string) (testCaseMetaDataArea fyne.CanvasObject, err error) {

	// Get current TestCase-UI-model
	_, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]

	if existsInMap == true {
		errorId := "bcb9d984-3106-42b6-9c23-288ec6d26224"
		err = errors.New(fmt.Sprintf("testcase-UI-model with uuid '%s' allready exist in 'TestCasesUiModelMap' [ErrorID: %s]", testCaseUuid, errorId))

		return nil, err
	}

	testCaseMetaDataArea = widget.NewLabel("'testCaseMetaDataArea'")

	return testCaseMetaDataArea, err
}

// Generate the TestCaseAttributes Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateTestCaseAttributesAreaForTestCase(testCaseUuid string) (testCaseAttributesArea fyne.CanvasObject, err error) {

	// Get current TestCase-UI-model
	_, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]

	if existsInMap == true {
		errorId := "c4110d4f-3dca-48bd-a8e4-57cb040fe079"
		err = errors.New(fmt.Sprintf("testcase-UI-model with uuid '%s' allready exist in 'TestCasesUiModelMap' [ErrorID: %s]", testCaseUuid, errorId))

		return nil, err
	}

	testCaseAttributesArea = widget.NewLabel("'testCaseAttributesArea'")

	return testCaseAttributesArea, err
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
		err = errors.New(fmt.Sprintf("testcase with uuid '%s' doesn't exist in 'TestCasesUiModelMap' [ErrorID: %s]", testCaseUuid, errorId))

		return err
	}

	// Update 'Binded' variables in UI-model for TestCase
	currentTestCaseUiModel.currentTestCaseTextualStructure.currentTestCaseTextualStructureSimple.Set(testCaseTextualStructureSimple)
	currentTestCaseUiModel.currentTestCaseTextualStructure.currentTestCaseTextualStructureComplex.Set(testCaseTextualStructureComplex)
	currentTestCaseUiModel.currentTestCaseTextualStructure.currentTestCaseTextualStructureExtended.Set(testCaseTextualStructureExtended)

	return err
}
