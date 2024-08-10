package templateViewer

import (
	"FenixTesterGui/importFilesFromGitHub"
	"fyne.io/fyne/v2"
)

// A pointer to Fenix Main Window
var fenixMainWindow fyne.Window

// The window for the File Viewer
var templateViewerWindow fyne.Window

var importedFiles []importFilesFromGitHub.GitHubFile
