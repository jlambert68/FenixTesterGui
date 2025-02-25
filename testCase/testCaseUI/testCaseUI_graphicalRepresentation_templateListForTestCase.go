package testCaseUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/importFilesFromGitHub"
	"FenixTesterGui/soundEngine"
	"FenixTesterGui/testCase/testCaseModel"
	"FenixTesterGui/testCase/testCaseUI/templateViewer"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/jlambert68/FenixScriptEngine/luaEngine"
	"github.com/sirupsen/logrus"
)

// Generate the Template-table Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateTemplateListForTestCaseArea(
	testCaseUuid string) (
	fyne.CanvasObject,
	error) {

	var templatesFilesInTestCaseTable *CustomTemplateTable

	var tableAndButtonContainer *fyne.Container
	var tableAccordionItem *widget.AccordionItem
	var accordion *widget.Accordion
	var templateListArea fyne.CanvasObject

	var githubFilesImporterButton *widget.Button
	var checkIfTemplatesAreChangedButton *widget.Button
	var viewTemplateButton *widget.Button

	var existInMap bool
	var currentTestCase testCaseModel.TestCaseModelStruct

	// Initiate Lua-script-Engine. TODO For now only Fenix-Placeholders are supported
	luaEngine.InitiateLuaScriptEngine([]luaEngine.LuaScriptsStruct{})

	currentTestCase, existInMap = testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]
	if existInMap == false {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":           "a54bce68-fa84-4b29-aa62-5d47b8bdc7fb",
			"testCaseUuid": testCaseUuid,
		}).Fatal("TestCase doesn't exist in TestCaseMap. This should not happen")
	}

	// Generate the Table used for showing Templates in TestCase
	templatesFilesInTestCaseTable = generateTemplateFilesTable(
		testCaseUuid,
		false,
		testCasesUiCanvasObject)

	// Update the Table used for showing Templates in TestCase
	//updateTemplateFilesTable(
	//	templatesFilesInTestCaseTable,
	//	testCaseUuid,
	//	false,
	//	testCasesUiCanvasObject)

	// Create Button to be able to Add or Remove Template from TestCase

	var responseChannel chan importFilesFromGitHub.SharedResponseChannelStruct
	responseChannel = make(chan importFilesFromGitHub.SharedResponseChannelStruct)

	// Import the Template-files
	githubFilesImporterButton = widget.NewButton("Import files from GitHub", func() {

		currentTestCase, existInMap = testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]
		if existInMap == false {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":           "59fab568-2da4-43f9-8300-6858eae73431",
				"testCaseUuid": testCaseUuid,
			}).Fatal("TestCase doesn't exist in TestCaseMap. This should not happen")
		}

		var tempFenixMasterWindow fyne.Window
		tempFenixMasterWindow = *sharedCode.FenixMasterWindowPtr
		tempFenixMasterWindow.Hide()

		var localImportFilesFromGitHubObject *importFilesFromGitHub.ImportFilesFromGitHubStruct
		localImportFilesFromGitHubObject = &importFilesFromGitHub.ImportFilesFromGitHubStruct{}

		localImportFilesFromGitHubObject.InitiateImportFilesFromGitHubWindow(
			*sharedCode.TemplateRepositoryApiUrlsPtr,
			*sharedCode.FenixMasterWindowPtr,
			*sharedCode.FenixAppPtr,
			&responseChannel,
			currentTestCase.ImportedTemplateFilesFromGitHub)

		// Wait for response from Files Selector Window to close
		var channelResponseForSelectedFiles importFilesFromGitHub.SharedResponseChannelStruct

		channelResponseForSelectedFiles = <-responseChannel

		var localCopyForSelectedFiles []importFilesFromGitHub.GitHubFile
		localCopyForSelectedFiles = *channelResponseForSelectedFiles.SelectedFilesPtr

		// Update Template files for TestCase
		currentTestCase.ImportedTemplateFilesFromGitHub = localCopyForSelectedFiles

		// Store back TestCase
		testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid] = currentTestCase

		//updateTemplateFilesTable(templatesFilesInTestCaseTable,
		//	testCaseUuid,
		//	false,
		//	testCasesUiCanvasObject)

		//		accordion.Resize(fyne.NewSize(accordion.Size().Width, templatesFilesInTestCaseTable.Size().Height))
		// tableAndButtonContainer.Refresh()
		// accordion.Refresh()

		// Update size of Columns and rows
		templatesFilesInTestCaseTable.updateColumnAndRowSizes(
			testCaseUuid,
			testCasesUiCanvasObject,
			checkIfTemplatesAreChangedButton,
			viewTemplateButton)

		templateListArea.Refresh()

		// Update Attributes area of TestCase
		if testCasesUiCanvasObject.CurrentSelectedTestCaseUIElement != nil {
			testCasesUiCanvasObject.CurrentSelectedTestCaseUIElement.ForceClick()
		}

		// Update ComboBox, if exists, for available Templates to chose from

	})

	// Create Button to be able to check which Templates that are updated
	checkIfTemplatesAreChangedButton = widget.NewButton("Check if Templates are changed", func() {
		// Add button functionality here

		currentTestCase, existInMap = testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]
		if existInMap == false {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":           "59fab568-2da4-43f9-8300-6858eae73431",
				"testCaseUuid": testCaseUuid,
			}).Fatal("TestCase doesn't exist in TestCaseMap. This should not happen")
		}

		// Trigger System Notification sound
		soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "Check if Templates are changed",
			Content: "Button 'checkIfTemplatesAreChangedButton' clicked",
		})

	})

	// Create Button to be able to view Template and the effect of TestData- and PlaceHolder-engine
	viewTemplateButton = widget.NewButton("View Templates", func() {

		currentTestCase, existInMap = testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]
		if existInMap == false {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":           "994ac3c8-2a89-4786-8c70-96bb86fbe70d",
				"testCaseUuid": testCaseUuid,
			}).Fatal("TestCase doesn't exist in TestCaseMap. This should not happen")
		}

		templateViewer.InitiateTemplateViewer(
			*sharedCode.FenixMasterWindowPtr,
			*sharedCode.FenixAppPtr,
			&currentTestCase.ImportedTemplateFilesFromGitHub,
			currentTestCase.TestData,
			testCaseUuid,
			"",
			testDataPointGroupsSelectSelectedInMainTestCaseArea,
			testDataPointForAGroupSelectSelectedInMainTestCaseArea,
			testDataRowForTestDataPointsSelectSelectedInMainTestCaseArea)

	})

	templatesFilesInTestCaseTable.updateColumnAndRowSizes(
		testCaseUuid,
		testCasesUiCanvasObject,
		checkIfTemplatesAreChangedButton,
		viewTemplateButton)

	// Create an Accordion item for the buttons
	buttonContainer := container.NewHBox(githubFilesImporterButton, checkIfTemplatesAreChangedButton, viewTemplateButton)

	tableAndButtonContainer = container.NewBorder(nil, buttonContainer, nil, nil, templatesFilesInTestCaseTable)
	tableAndButtonContainer.Refresh()

	// Create an Accordion item for the list
	tableAccordionItem = widget.NewAccordionItem("Templates", tableAndButtonContainer)

	accordion = widget.NewAccordion(tableAccordionItem) // widget.NewAccordion(tableAccordionItem)

	// Create the VBox-container that will be returned
	templateListArea = container.NewVBox(accordion, widget.NewLabel(""), widget.NewSeparator())

	return templateListArea, nil

}
