package testCaseUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/importFilesFromGitHub"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var templatesFilesInTestCaseTable *widget.Table
var templatesFilesInTestCase []GitHubFile

// Generate the Template-table Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateTemplateListForTestCaseArea(
	testCaseUuid string) (
	fyne.CanvasObject,
	error) {

	var tableAndButtonContainer *fyne.Container
	var tableAccordionItem *widget.AccordionItem
	var accordion *widget.Accordion

	// Generate the Table used for showing Templates in TestCase
	templatesFilesInTestCaseTable = generateTemplateFilesTable()

	// Update the Table used for showing Templates in TestCase
	updateTemplateFilesTable(templatesFilesInTestCaseTable, &templatesFilesInTestCase, false)

	// Create Button to be able to Add or Remove Template from TestCase

	var responseChannel chan bool
	responseChannel = make(chan bool)
	var selectedFiles *[]importFilesFromGitHub.GitHubFile
	selectedFiles = &[]importFilesFromGitHub.GitHubFile{}
	githubFilesImporterButton := widget.NewButton("Import files from GitHub", func() {
		var tempFenixMasterWindow fyne.Window
		tempFenixMasterWindow = *sharedCode.FenixMasterWindowPtr
		tempFenixMasterWindow.Hide()

		var tempSelectedFiles []importFilesFromGitHub.GitHubFile
		tempSelectedFiles = *selectedFiles
		selectedFiles = importFilesFromGitHub.InitiateImportFilesFromGitHubWindow(
			*sharedCode.TemplateRepositoryApiUrlsPtr, *sharedCode.FenixMasterWindowPtr, *sharedCode.FenixAppPtr, &responseChannel, tempSelectedFiles)
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
