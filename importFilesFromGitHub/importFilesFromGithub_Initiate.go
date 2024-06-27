package importFilesFromGitHub

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"strings"
)

func (importFilesFromGitHubObject ImportFilesFromGitHubStruct) InitiateImportFilesFromGitHubWindow(
	templateRepositoryApiUrls []*fenixGuiTestCaseBuilderServerGrpcApi.RepositoryApiUrlResponseMessage,
	mainWindow fyne.Window,
	myApp fyne.App,
	responseChannel *chan SharedResponseChannelStruct,
	tempSelectedFiles []GitHubFile) {

	importFilesFromGitHubObject.selectedFiles = tempSelectedFiles

	// Cleare list from Previous Import
	importFilesFromGitHubObject.githubFilesFiltered = nil

	// Disable the main window
	mainWindow.Hide()

	// Store Channel reference in local varaible
	sharedResponseChannel = responseChannel

	// Store reference to Fenix Main Window
	fenixMainWindow = mainWindow

	// Create the window for GitHub files
	githubFileImporterWindow = myApp.NewWindow("Template file importer")

	importFilesFromGitHubObject.rootApiUrl = templateRepositoryApiUrls[0].GetRepositoryApiFullUrl()
	importFilesFromGitHubObject.currentApiUrl = importFilesFromGitHubObject.rootApiUrl

	importFilesFromGitHubObject.currentGitHubApiKey = templateRepositoryApiUrls[0].GetGitHubApiKey()

	importFilesFromGitHubObject.currentPathShowedinGUI = binding.NewString()
	importFilesFromGitHubObject.currentPathShowedinGUI.Set(strings.Split(importFilesFromGitHubObject.currentApiUrl, "?")[0]) // Setting initial value

	importFilesFromGitHubObject.fileRegExFilterMap = make(map[string]string)

	// Retrieve and Filter files from GitHub
	importFilesFromGitHubObject.getFileListFromGitHub(importFilesFromGitHubObject.currentApiUrl)
	importFilesFromGitHubObject.filterFileListFromGitHub()

	// Create the UI-list that holds the selected files
	importFilesFromGitHubObject.generateSelectedFilesListTable(githubFileImporterWindow)

	// Update the table
	importFilesFromGitHubObject.UpdateSelectedFilesTable()

	// Create the UI-list that holds the filtered files and folders from GitHub
	importFilesFromGitHubObject.generateFilteredList(githubFileImporterWindow)

	// Generate the File filter PopUp
	importFilesFromGitHubObject.generateFileFilterPopup(githubFileImporterWindow)

	// Create a label with data binding used for showing current GitHub path
	importFilesFromGitHubObject.pathLabel = widget.NewLabelWithData(importFilesFromGitHubObject.currentPathShowedinGUI)

	// Generate the Button that moves upwards in the folder structure in GitHub
	importFilesFromGitHubObject.generateMoveUpInFolderStructureButton()

	// Generate the button that imports the selected files from GitHub
	importFilesFromGitHubObject.generateImportSelectedFilesFromGithubButton(githubFileImporterWindow)

	// Generate the button that cancel everything and closes the window
	importFilesFromGitHubObject.generateCancelButton(githubFileImporterWindow)

	// Generate the list with URLs to use in Select
	var urlInSelect []string
	for _, templateRepositoryApiUrl := range templateRepositoryApiUrls {
		urlInSelect = append(urlInSelect, templateRepositoryApiUrl.GetRepositoryApiFullUrl())
	}

	// Generate the DropDown that holds the list of Repositories
	importFilesFromGitHubObject.generateGitHubRepositorySelect(urlInSelect, templateRepositoryApiUrls)

	// Set initial size of the window
	githubFileImporterWindow.Resize(fyne.NewSize(1200, 1000))

	var pathSelectorContainer *fyne.Container
	pathSelectorContainer = container.NewBorder(nil, nil, nil, nil, importFilesFromGitHubObject.githubRepositorySelect)

	// Generate the row that holds the up/back button and the path itself
	var pathRowContainer *fyne.Container
	pathRowContainer = container.NewHBox(importFilesFromGitHubObject.moveUpInFolderStructureButton, importFilesFromGitHubObject.pathLabel)

	// Create the top element which has the Filter button and the path.label and the back/upp button
	myTopLayout := container.NewVBox(importFilesFromGitHubObject.fileFilterPopupButton, pathSelectorContainer, pathRowContainer)

	// Create the container that 'selectedFilesTable' will be placed in
	var selectedFilesTableContainer *fyne.Container
	selectedFilesTableContainer = container.NewStack(importFilesFromGitHubObject.selectedFilesTable)

	// Generate the container which has the filtered folders and files to the left and the selected files to the right
	splitContainer := container.NewHSplit(importFilesFromGitHubObject.filteredFileList, selectedFilesTableContainer)
	splitContainer.Offset = 0.5 // Adjust if you want different initial proportions

	// Generate the row that has the import button and the cancel button
	var importCancelRow *fyne.Container
	importCancelRow = container.NewHBox(layout.NewSpacer(), importFilesFromGitHubObject.importSelectedFilesFromGithubButton, importFilesFromGitHubObject.cancelButton)

	// Create the full content that should be showed in the window
	content := container.NewBorder(myTopLayout, importCancelRow, nil, nil, splitContainer)

	// Set content
	githubFileImporterWindow.SetContent(content)

	// Set the callback function for window close event to show the Main window again
	githubFileImporterWindow.SetOnClosed(func() {
		*sharedResponseChannel <- SharedResponseChannelStruct{
			SharedResponse:   false,
			SelectedFilesPtr: &importFilesFromGitHubObject.selectedFiles,
		}
		fenixMainWindow.Show()
	})

	// Show the githubFileImporterWindow
	githubFileImporterWindow.Show()

}
