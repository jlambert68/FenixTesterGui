package importFilesFromGitHub

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"log"
	"regexp"
	"strings"
	"time"
)

func (importFilesFromGitHubObject *ImportFilesFromGitHubStruct) generateFilteredList(parentWindow fyne.Window) {

	importFilesFromGitHubObject.filteredFileList = widget.NewList(
		func() int {
			return len(importFilesFromGitHubObject.githubFilesFiltered)
		},
		func() fyne.CanvasObject {
			// Create a customFilteredLabel for each item.
			label := importFilesFromGitHubObject.newCustomFilteredLabel("Template", func() {
				// Define double-click action here.
			})

			//var labelContainer *fyne.Container
			//labelContainer = container.NewStack(label, label.backgroundRectangle)

			return label
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			// Update the label text and double-click action for each item.
			label := obj.(*customFilteredLabel)
			label.Text = importFilesFromGitHubObject.githubFilesFiltered[id].Name

			if importFilesFromGitHubObject.githubFilesFiltered[id].Type == "file" {
				label.TextStyle = fyne.TextStyle{Italic: true}
			}

			label.onDoubleTap = func() {

				selectedFile := importFilesFromGitHubObject.githubFilesFiltered[id]
				if selectedFile.Type == "dir" {
					// The item is a directory; fetch its contents
					importFilesFromGitHubObject.getFileListFromGitHub(selectedFile.URL)
					importFilesFromGitHubObject.filterFileListFromGitHub()
					importFilesFromGitHubObject.filteredFileList.Refresh() // Refresh the list to update it with the new contents
					importFilesFromGitHubObject.currentPathShowedinGUI.Set(strings.Split(selectedFile.URL, "?")[0])

					importFilesFromGitHubObject.currentApiUrl = selectedFile.URL
				} else if selectedFile.Type == "file" {
					// Add file to selectedFiles and refresh the list only when if it doesn't exist
					var shouldAddFile bool
					shouldAddFile = true
					for _, existingSelectedFile := range importFilesFromGitHubObject.selectedFiles {
						if existingSelectedFile.URL == selectedFile.URL {
							shouldAddFile = false
							break
						}
					}

					if shouldAddFile == true {
						importFilesFromGitHubObject.selectedFiles = append(importFilesFromGitHubObject.selectedFiles, selectedFile)
						importFilesFromGitHubObject.UpdateSelectedFilesTable()
						importFilesFromGitHubObject.selectedFilesTable.Refresh()

					}

				} else {
					// Show a dialog when other.
					dialog.ShowInformation("Info", "Double-clicked on: "+importFilesFromGitHubObject.githubFiles[id].Name+" with Type "+importFilesFromGitHubObject.githubFiles[id].Type, parentWindow)
				}
			}
			label.Refresh()
		},
	)
}

func (importFilesFromGitHubObject *ImportFilesFromGitHubStruct) filterFileListFromGitHub() {

	var fullRegExFilter string
	var tempGithubFilesFiltered []GitHubFile

	var tempRegex string

	for fileFilter, _ := range importFilesFromGitHubObject.fileRegExFilterMap {
		if fileFilter == "*.*" {
			tempRegex = ".*"
		} else {

			tempRegex = strings.ReplaceAll(fileFilter, "*", "\\")
		}
		tempRegex = tempRegex + "$"

		if len(fullRegExFilter) == 0 {
			fullRegExFilter = fullRegExFilter + tempRegex
		} else {
			fullRegExFilter = fullRegExFilter + "|" + tempRegex
		}
	}

	if len(tempRegex) == 0 {
		tempRegex = `.*`
	}

	combinedRegex, err := regexp.Compile(fullRegExFilter)
	if err != nil {
		log.Fatalln("Error compiling regex:", err)
		return
	}

	for _, githubFile := range importFilesFromGitHubObject.githubFiles {
		if combinedRegex.MatchString(githubFile.Name) == true || githubFile.Type == "dir" {
			tempGithubFilesFiltered = append(tempGithubFilesFiltered, githubFile)
		}
	}

	importFilesFromGitHubObject.githubFilesFiltered = tempGithubFilesFiltered
}

type customFilteredLabel struct {
	widget.Label
	onDoubleTap func()
	lastTap     time.Time
}

func (importFilesFromGitHubObject *ImportFilesFromGitHubStruct) newCustomFilteredLabel(
	text string,
	onDoubleTap func(),
) *customFilteredLabel {

	l := &customFilteredLabel{
		Label:       widget.Label{Text: text},
		onDoubleTap: onDoubleTap}

	l.ExtendBaseWidget(l)
	return l
}

func (l *customFilteredLabel) Tapped(e *fyne.PointEvent) {
	now := time.Now()
	if now.Sub(l.lastTap) < 500*time.Millisecond { // 500 ms as double-click interval
		if l.onDoubleTap != nil {
			l.onDoubleTap()
		}
	}
	l.lastTap = now
}

func (l *customFilteredLabel) TappedSecondary(*fyne.PointEvent) {
	// Implement if you need right-click (secondary tap) actions.
}

func (l *customFilteredLabel) MouseIn(*desktop.MouseEvent) {
	l.TextStyle = fyne.TextStyle{Bold: true}
	l.Refresh()
}
func (l *customFilteredLabel) MouseMoved(*desktop.MouseEvent) {}
func (l *customFilteredLabel) MouseOut() {
	l.TextStyle = fyne.TextStyle{Bold: false}
	l.Refresh()
}
