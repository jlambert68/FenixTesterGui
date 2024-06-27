package importFilesFromGitHub

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

// Channel used to inform when this window closes and there are values to use
// True = files where picked
// False = No file where picked
var sharedResponseChannel *chan SharedResponseChannelStruct

type SharedResponseChannelStruct struct {
	SharedResponse   bool
	SelectedFilesPtr *[]GitHubFile
}

// A pointer to Fenix Main Window
var fenixMainWindow fyne.Window

// The window for the File Importer
var githubFileImporterWindow fyne.Window

type ImportFilesFromGitHubStruct struct {

	// THe root ApiUrl
	rootApiUrl string

	// The current ApiUrl tp fetch files and folders from
	currentApiUrl string

	fileRegExFilterMap map[string]string

	githubFiles, githubFilesFiltered, selectedFiles []GitHubFile

	// Create a string data binding
	currentPathShowedinGUI binding.String

	selectedFilesTable    *widget.Table
	filteredFileList      *widget.List
	fileFilterPopupButton *widget.Button

	// Create a label with data binding used for showing current GitHub path
	pathLabel *widget.Label

	// The Button that moves upwards in the folder structure in GitHub
	moveUpInFolderStructureButton *widget.Button

	// The button that imports the selected files from GitHub
	importSelectedFilesFromGithubButton *widget.Button

	// The button that cancel everything and closes the window
	cancelButton *widget.Button

	// The Select holding the different URLs to chose from
	githubRepositorySelect *widget.Select

	// The GitHubKey for the selected repository in Select 'githubRepositorySelect'
	currentGitHubApiKey string
}
