package testCaseUI

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
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

	// Initiate map with TestCaseUI-models-Map
	testCasesUiCanvasObject.TestCasesUiModelMap = make(map[string]*testCaseGraphicalAreasStruct)

	return baseCanvasObjectForTestCaseUI
}

// GenerateNewTestCaseTabObject
// Generate a new TestCase UI-model
func (testCasesUiCanvasObject *TestCasesUiModelStruct) GenerateNewTestCaseTabObject(testCaseToBeAddedUuid string) (err error) {

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

	// Generate the Textual Binding Objects for Textual Representation and Textual Representation Area for the TestCase
	newTestCaseTextualStructure, canvasTextualRepresentationAccordionObject, err := testCasesUiCanvasObject.generateNewTextualRepresentationAreaForTestCase(testCaseToBeAddedUuid)
	if err != nil {
		return err
	}
	// Add newly created Textual Representation Area to object for all graphical parts of one TestCase
	testCaseGraphicalAreas.TestCaseTextualModelArea = canvasTextualRepresentationAccordionObject

	// Generate the Graphical Representation Area for the TestCase
	testCaseGraphicalModelArea, testCaseGraphicalUITree, testCaseGraphicalModelAreaAccordion, err := testCasesUiCanvasObject.generateGraphicalRepresentationAreaForTestCase(testCaseToBeAddedUuid)

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
	textualAndGraphicalRepresentations := container.NewVBox(
		testCaseGraphicalAreas.TestCaseTextualModelArea,
		widget.NewSeparator(),
		testCaseGraphicalAreas.TestCaseGraphicalModelArea,
		layout.NewSpacer(),
		makeDragNDropTestGUI(testCasesUiCanvasObject.DragNDropText, testCasesUiCanvasObject.DragNDropRectangle, testCasesUiCanvasObject.DragNDropContainer),
		layout.NewSpacer())

	// Create canvas-object for BaseInformation, MetaData and TestCaseAttributes
	baseInformationMetaDataTestCaseAttributes := container.NewVBox(
		testCaseGraphicalAreas.TestCaseBaseInformationArea,
		widget.NewSeparator(),
		testCaseGraphicalAreas.TestCaseMetaDataArea,
		widget.NewSeparator(),
		testCaseGraphicalAreas.TestCaseAttributesArea)
	//layout.NewSpacer())

	// Create the UI area for all parts of one TestCase
	testCaseAdaptiveSplitContainer := newAdaptiveSplit(textualAndGraphicalRepresentations, baseInformationMetaDataTestCaseAttributes)

	// Create a new Tab-object
	newTestCaseTabObject := container.NewTabItem(tabName, testCaseAdaptiveSplitContainer)

	// Add tab to existing Tab-object
	testCasesUiCanvasObject.TestCasesTabs.Append(newTestCaseTabObject)

	// Set focus on newly created Tab
	testCasesUiCanvasObject.TestCasesTabs.Select(newTestCaseTabObject)

	/*
		// Initiate Textual Representations for TestCase
		testCaseGraphicalAreas.currentTestCaseTextualStructure.currentTestCaseTextualStructureSimple = binding.NewString()
		testCaseGraphicalAreas.currentTestCaseTextualStructure.currentTestCaseTextualStructureComplex = binding.NewString()
		testCaseGraphicalAreas.currentTestCaseTextualStructure.currentTestCaseTextualStructureExtended = binding.NewString()


	*/

	// 	Save Textual Binding Objects and Accordion Objectfor Textual Representation
	testCaseGraphicalAreas.currentTestCaseTextualStructure = newTestCaseTextualStructure

	// save Graphical object into TestCase, to be reachable
	testCaseGraphicalAreas.currentTestCaseGraphicalStructure.currentTestCaseGraphicalAccordionObject = testCaseGraphicalModelAreaAccordion
	testCaseGraphicalAreas.currentTestCaseGraphicalStructure.currentTestCaseGraphicalTreeComponent = testCaseGraphicalUITree

	// Open 'Accordions' for Textual and Graphical TestCase Representation for TestCase
	testCaseGraphicalAreas.currentTestCaseTextualStructure.currentTestCaseGraphicalAccordionObject.OpenAll()
	testCaseGraphicalAreas.currentTestCaseGraphicalStructure.currentTestCaseGraphicalAccordionObject.OpenAll()

	// Save TestCase UI-components-Map
	testCasesUiCanvasObject.TestCasesUiModelMap[testCaseToBeAddedUuid] = &testCaseGraphicalAreas

	return err
}
