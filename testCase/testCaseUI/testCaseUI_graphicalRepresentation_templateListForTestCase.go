package testCaseUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/importFilesFromGitHub"
	"FenixTesterGui/testCase/testCaseModel"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
)

// Generate the Template-table Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateTemplateListForTestCaseArea(
	testCaseUuid string) (
	fyne.CanvasObject,
	error) {

	var templatesFilesInTestCaseTable *widget.Table

	var tableAndButtonContainer *fyne.Container
	var tableAccordionItem *widget.AccordionItem
	var accordion *widget.Accordion

	// Generate the Table used for showing Templates in TestCase
	templatesFilesInTestCaseTable = generateTemplateFilesTable()

	// Update the Table used for showing Templates in TestCase
	updateTemplateFilesTable(
		templatesFilesInTestCaseTable,
		testCaseUuid,
		false,
		testCasesUiCanvasObject)

	// Create Button to be able to Add or Remove Template from TestCase

	var responseChannel chan importFilesFromGitHub.SharedResponseChannelStruct
	responseChannel = make(chan importFilesFromGitHub.SharedResponseChannelStruct)

	// Import the Template-files
	githubFilesImporterButton := widget.NewButton("Import files from GitHub", func() {

		var existInMap bool
		var currectTestCase testCaseModel.TestCaseModelStruct

		currectTestCase, existInMap = testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]
		if existInMap == false {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":           "a54bce68-fa84-4b29-aa62-5d47b8bdc7fb",
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
			currectTestCase.ImportedTemplateFilesFromGitHub)

		// Wait for response from Files Selector Window to close
		var channelResponseForSelectedFiles importFilesFromGitHub.SharedResponseChannelStruct

		channelResponseForSelectedFiles = <-responseChannel

		var localCopyForSelectedFiles []importFilesFromGitHub.GitHubFile
		localCopyForSelectedFiles = *channelResponseForSelectedFiles.SelectedFilesPtr

		// Update Template files for TestCase
		currectTestCase.ImportedTemplateFilesFromGitHub = localCopyForSelectedFiles

		// Store back TestCase
		testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid] = currectTestCase

		updateTemplateFilesTable(templatesFilesInTestCaseTable,
			testCaseUuid,
			false,
			testCasesUiCanvasObject)

	})

	// Create Button to be able to check which Templates that are updated
	checkIfTemplatesAreChangedButton := widget.NewButton("Check if Templates are changed", func() {
		// Add button functionality here
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "Check if Templates are changed",
			Content: "Button 'checkIfTemplatesAreChangedButton' clicked",
		})

	})

	// Create an Accordion item for the buttons
	buttonContainer := container.NewHBox(githubFilesImporterButton, checkIfTemplatesAreChangedButton)

	tableAndButtonContainer = container.NewBorder(nil, buttonContainer, nil, nil, templatesFilesInTestCaseTable)
	tableAndButtonContainer.Refresh()

	// Create an Accordion item for the list
	tableAccordionItem = widget.NewAccordionItem("Templates", tableAndButtonContainer)

	accordion = widget.NewAccordion(tableAccordionItem) // widget.NewAccordion(tableAccordionItem)

	// Create the VBox-container that will be returned
	templateListArea := container.NewVBox(accordion)

	return templateListArea, nil

}
