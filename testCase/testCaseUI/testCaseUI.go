package testCaseUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/soundEngine"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
)

// GenerateBaseCanvasObjectForTestCaseUI
// Create the Base-UI-canvas-object for the TestCases object. This base doesn't contain any specific TestCase-parts, and they will be added in other function
func (testCasesUiCanvasObject *TestCasesUiModelStruct) GenerateBaseCanvasObjectForTestCaseUI() (baseCanvasObjectForTestCaseUI fyne.CanvasObject) {

	// Create toolbar for TestCase area
	testCasesUiCanvasObject.TestCaseToolUIBar = widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentRedoIcon(), func() {
			fmt.Println("Reload GUI TestCase from testCaseModel")
		}),

		// New TestCase
		widget.NewToolbarAction(theme.DocumentIcon(), func() {
			commandEngineChannelMessage := sharedCode.ChannelCommandStruct{
				ChannelCommand:  sharedCode.ChannelCommandNewTestCase,
				FirstParameter:  "",
				SecondParameter: "",
				ActiveTestCase:  "",
				ElementType:     sharedCode.BuildingBlock(sharedCode.Undefined),
			}

			// Send command message over channel to Command and Rule Engine
			*testCasesUiCanvasObject.CommandChannelReference <- commandEngineChannelMessage
		}),

		// Open TestCase
		widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
			commandEngineChannelMessage := sharedCode.ChannelCommandStruct{
				ChannelCommand:  sharedCode.ChannelCommandOpenTestCase,
				FirstParameter:  "",
				SecondParameter: "",
				ActiveTestCase:  "",
				ElementType:     sharedCode.BuildingBlock(sharedCode.Undefined),
			}

			// Send command message over channel to Command and Rule Engine
			*testCasesUiCanvasObject.CommandChannelReference <- commandEngineChannelMessage
		}),

		// Save TestCase
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
			commandEngineChannelMessage := sharedCode.ChannelCommandStruct{
				ChannelCommand:  sharedCode.ChannelCommandSaveTestCase,
				FirstParameter:  "",
				SecondParameter: "",
				ActiveTestCase:  "",
				ElementType:     sharedCode.BuildingBlock(sharedCode.Undefined),
			}

			// Send command message over channel to Command and Rule Engine
			*testCasesUiCanvasObject.CommandChannelReference <- commandEngineChannelMessage
		}),

		// Remove Node in TestCase
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {

			var foundTestCase bool
			var testCaseName string

			// If only the "Home-tab" is left then play "angry sound"
			if testCasesUiCanvasObject.TestCasesModelReference.CurrentActiveTestCaseUuid == "" {

				// Trigger Invalid Notification sound
				soundEngine.PlaySoundChannel <- soundEngine.InvalidNotificationSound

				return
			}

			// Loop Map with TestCase-tabs to find relation between TabItem and UUID
			for _, tempTestCaseUITabRefToTestCaseUuidMapStructObject := range testCasesUiCanvasObject.TestCaseUITabRefToTestCaseUuidMap {

				// Is this the TestCaseUuid we are looking for
				if tempTestCaseUITabRefToTestCaseUuidMapStructObject.TestCaseUuid == testCasesUiCanvasObject.TestCasesModelReference.CurrentActiveTestCaseUuid {
					testCaseName = tempTestCaseUITabRefToTestCaseUuidMapStructObject.TestCaseUiTabRef.Text
					foundTestCase = true
					break
				}
			}

			// When TestCase was not found then
			if foundTestCase == false {
				sharedCode.Logger.WithFields(logrus.Fields{
					"ID": "bbb933c2-3a57-40b9-be99-4725ce9727a0",
					"testCasesUiCanvasObject.TestCasesModelReference.CurrentActiveTestCaseUuid": testCasesUiCanvasObject.TestCasesModelReference.CurrentActiveTestCaseUuid,
				}).Fatal("TestCase doesn't exist in Tab-reference-map. This should not happen")

			}

			soundEngine.PlaySoundChannel <- soundEngine.UserNeedToRespondSound

			// Show a confirmation dialog
			dialog.ShowConfirm("Confirm to Close TestCase", "Do you want to close the TestCase without saving it?\n"+testCaseName+"\n"+testCasesUiCanvasObject.TestCasesModelReference.CurrentActiveTestCaseUuid,
				func(confirm bool) {
					if confirm {

						commandEngineChannelMessage := sharedCode.ChannelCommandStruct{
							ChannelCommand:  sharedCode.ChannelCommandCloseOpenTestCaseWithOutSaving,
							FirstParameter:  testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCasesUiCanvasObject.TestCasesModelReference.CurrentActiveTestCaseUuid].CurrentSelectedTestCaseElement.CurrentSelectedTestCaseElementUuid,
							SecondParameter: "",
							ActiveTestCase:  testCasesUiCanvasObject.TestCasesModelReference.CurrentActiveTestCaseUuid,
							ElementType:     sharedCode.BuildingBlock(sharedCode.TestInstruction),
						}

						// Send command message over channel to Command and Rule Engine
						*testCasesUiCanvasObject.CommandChannelReference <- commandEngineChannelMessage
					} else {
						// Do nothing
					}
				}, *sharedCode.FenixMasterWindowPtr)

		}),

		// Execute 'current' TestCase
		widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
			commandEngineChannelMessage := sharedCode.ChannelCommandStruct{
				ChannelCommand:  sharedCode.ChannelCommandExecuteTestCase,
				FirstParameter:  testCasesUiCanvasObject.TestCasesModelReference.CurrentActiveTestCaseUuid,
				SecondParameter: "",
				ActiveTestCase:  "",
				ElementType:     sharedCode.BuildingBlock(sharedCode.Undefined),
			}

			// Send command message over channel to Command and Rule Engine
			*testCasesUiCanvasObject.CommandChannelReference <- commandEngineChannelMessage
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

func (testCasesUiCanvasObject *TestCasesUiModelStruct) GenerateNewTestCaseTabObject(
	testCaseToBeAddedUuid string) (
	err error) {

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
	testCaseGraphicalModelArea,
		graphicalTestCaseUIObject,
		testCaseGraphicalModelAreaAccordion, err := testCasesUiCanvasObject.
		generateGraphicalRepresentationAreaForTestCase(testCaseToBeAddedUuid)

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
	testCaseAttributesArea, testInstructionAttributesAccordion, err := testCasesUiCanvasObject.
		generateTestCaseAttributesAreaForTestCase(testCaseToBeAddedUuid, "") // "" used for first time creation

	if err != nil {
		return err
	}
	// Add newly created TestCaseAttributes Area to object for all graphical parts of one TestCase
	testCaseGraphicalAreas.TestCaseAttributesArea = testCaseAttributesArea
	testCaseGraphicalAreas.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject = testInstructionAttributesAccordion

	// Save TestCase-UI-model in UI-modelMap
	// Check if TestCase already exists in TestCase-UI-model
	_, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseToBeAddedUuid]

	if existsInMap == true {
		errorId := "db34dee8-1b23-425c-868a-2747959ec682"
		err = errors.New(fmt.Sprintf("testcase-UI-model with sourceUuid '%s' allready exist in 'TestCasesUiModelMap' [ErrorID: %s]", testCaseToBeAddedUuid, errorId))

		return err
	}

	// Create canvas-object for Textual and Graphical Representation
	textualAndGraphicalRepresentations := container.NewBorder(
		testCaseGraphicalAreas.TestCaseTextualModelArea,
		nil,
		nil,
		nil,
		testCaseGraphicalAreas.TestCaseGraphicalModelArea)

	// Create canvas-object for BaseInformation, MetaData and TestCaseAttributes
	baseInformationMetaDataTestCaseAttributes := container.NewVBox(
		testCaseGraphicalAreas.TestCaseBaseInformationArea,
		widget.NewSeparator(),
		testCaseGraphicalAreas.TestCaseMetaDataArea,
		widget.NewSeparator(),
	)

	// Create the BorderContainer for BaseInformation, MetaData and TestCaseAttributes...
	baseInformationMetaDataTestCaseAttributesBorderContainer := container.NewBorder(
		baseInformationMetaDataTestCaseAttributes, nil, nil, nil,
		testCaseGraphicalAreas.TestCaseAttributesArea)

	// Create the UI area for all parts of one TestCase
	testCaseAdaptiveSplitContainer := newAdaptiveSplit(
		textualAndGraphicalRepresentations, container.NewWithoutLayout(),
		baseInformationMetaDataTestCaseAttributesBorderContainer, container.NewWithoutLayout())

	// Create a new Tab-object
	newTestCaseTabObject := container.NewTabItem(tabName, testCaseAdaptiveSplitContainer)

	// Add tab to existing Tab-object
	testCasesUiCanvasObject.TestCasesTabs.Append(newTestCaseTabObject)

	testCasesUiCanvasObject.TestCasesTabs.OnSelected = func(tabItem *container.TabItem) {
		fmt.Println("OnSelected")
		fmt.Println(tabItem)

		if TempTestCasesUiCanvasObject != nil {
			var existInMap bool
			var tabItemRefString string

			tabItemRefString = fmt.Sprintf("%p", tabItem)
			_, existInMap = TempTestCasesUiCanvasObject.TestCaseUITabRefToTestCaseUuidMap[tabItemRefString]

			if existInMap == true {
				var testCaseUuid string
				testCaseUuid = TempTestCasesUiCanvasObject.TestCaseUITabRefToTestCaseUuidMap[tabItemRefString].TestCaseUuid

				// Send command 'ChannelCommandChangeActiveTestCase' on command-channle
				commandEngineChannelMessage := sharedCode.ChannelCommandStruct{
					ChannelCommand:  sharedCode.ChannelCommandChangeActiveTestCase,
					FirstParameter:  "",
					SecondParameter: "",
					ActiveTestCase:  testCaseUuid,
					ElementType:     sharedCode.BuildingBlock(sharedCode.Undefined),
				}

				// Send command message over channel to Command and Rule Engine
				sharedCode.CommandChannel <- commandEngineChannelMessage

			}

		}

		//testCaseExecutionsTabPage.Refresh()

	}

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
	//testCaseGraphicalAreas.currentTestCaseGraphicalStructure.currentTestCaseGraphicalTreeComponent = testCaseGraphicalUITree
	testCaseGraphicalAreas.currentTestCaseGraphicalStructure.currentTestCaseGraphicalObject = &graphicalTestCaseUIObject

	// Open 'Accordions' for Textual and Graphical TestCase Representation for TestCase
	testCaseGraphicalAreas.currentTestCaseTextualStructure.currentTestCaseGraphicalAccordionObject.OpenAll()
	testCaseGraphicalAreas.currentTestCaseGraphicalStructure.currentTestCaseGraphicalAccordionObject.OpenAll()

	// Save link between newTestCaseTabObject and TestCaseUuid in Map
	if testCasesUiCanvasObject.TestCaseUITabRefToTestCaseUuidMap == nil {
		// Initiate if nil
		testCasesUiCanvasObject.TestCaseUITabRefToTestCaseUuidMap = make(map[string]TestCaseUITabRefToTestCaseUuidMapstruct)
		TempTestCasesUiCanvasObject = &TestCasesUiModelStruct{}
		TempTestCasesUiCanvasObject.TestCaseUITabRefToTestCaseUuidMap = make(map[string]TestCaseUITabRefToTestCaseUuidMapstruct)
	}

	var newTestCaseTabObjectRefString string
	newTestCaseTabObjectRefString = fmt.Sprintf("%p", newTestCaseTabObject)

	testCasesUiCanvasObject.TestCaseUITabRefToTestCaseUuidMap[newTestCaseTabObjectRefString] = TestCaseUITabRefToTestCaseUuidMapstruct{
		TestCaseUuid:     testCaseToBeAddedUuid,
		TestCaseUiTabRef: newTestCaseTabObject,
	}

	TempTestCasesUiCanvasObject.TestCaseUITabRefToTestCaseUuidMap = testCasesUiCanvasObject.TestCaseUITabRefToTestCaseUuidMap

	// Save TestCase UI-components-Map
	testCasesUiCanvasObject.TestCasesUiModelMap[testCaseToBeAddedUuid] = &testCaseGraphicalAreas

	return err
}
