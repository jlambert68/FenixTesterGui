package testCaseUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/importFilesFromGitHub"
	"FenixTesterGui/testCase/testCaseModel"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
	"time"
)

type clickableLabel struct {
	widget.Label
	onDoubleTap func()
	lastTapTime time.Time
	isClickable bool
}

func newClickableLabel(text string, onDoubleTap func(), tempIsClickable bool) *clickableLabel {
	l := &clickableLabel{
		widget.Label{Text: text},
		onDoubleTap,
		time.Now(),
		tempIsClickable}
	l.ExtendBaseWidget(l)
	return l
}

func (l *clickableLabel) Tapped(e *fyne.PointEvent) {
	if l.isClickable == false {
		return
	}

	if time.Since(l.lastTapTime) < 500*time.Millisecond {
		if l.onDoubleTap != nil {
			l.onDoubleTap()
		}
	}
	l.lastTapTime = time.Now()
}

func (l *clickableLabel) TappedSecondary(*fyne.PointEvent) {
	// Implement if you need right-click (secondary tap) actions.
}

func (l *clickableLabel) MouseIn(*desktop.MouseEvent)    {}
func (l *clickableLabel) MouseMoved(*desktop.MouseEvent) {}
func (l *clickableLabel) MouseOut()                      {}

// Create the UI-list that holds the selected files
func generateTemplateFilesTable() *widget.Table {

	// Correctly initialize the templateFilesTable as a new table
	var templateFilesTable *widget.Table
	templateFilesTable = widget.NewTable(
		func() (int, int) { return 0, 2 }, // Start with zero rows, 2 columns
		func() fyne.CanvasObject {
			return widget.NewLabel("") // Create cells as labels
		},
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			// This should be filled when updating the table
		},
	)

	return templateFilesTable

}

/*

// Struct for parsing JSON response for files from Github
type GitHubFile struct {
	Name                string `json:"name"`
	Type                string `json:"type"` // "file" or "dir"
	URL                 string `json:"url"`  // URL to fetch contents if it's a directory
	Content             []byte `json:"content"`
	SHA                 string `json:"sha"`
	FileContentAsString string
}
type GitHubFileDetail struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	URL         string `json:"url"`
	DownloadURL string `json:"download_url"`
	Type        string `json:"type"`
	Content     string `json:"content"`
	Encoding    string `json:"encoding"`
	SHA         string `json:"sha"`
	// Include other fields as needed
}

*/

func updateTemplateFilesTable(
	templateFilesTable *widget.Table,
	testCaseUuid string,
	filesAreClickable bool,
	testCasesUiCanvasObject *TestCasesUiModelStruct) {

	var existInMap bool
	var currentTestCase testCaseModel.TestCaseModelStruct

	currentTestCase, existInMap = testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]
	if existInMap == false {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":           "76ce3c8f-7791-44c9-8f18-be1ed0d9544d",
			"testCaseUuid": testCaseUuid,
		}).Fatal("TestCase doesn't exist in TestCaseMap. This should not happen")
	}

	var templateFilesFromGithub []importFilesFromGitHub.GitHubFile
	templateFilesFromGithub = currentTestCase.ImportedTemplateFilesFromGitHub

	templateFilesTable.Length = func() (int, int) {
		return len(templateFilesFromGithub), 2
	}
	templateFilesTable.CreateCell = func() fyne.CanvasObject {
		return newClickableLabel("", func() {}, false)

	}
	templateFilesTable.UpdateCell = func(id widget.TableCellID, cell fyne.CanvasObject) {

		currentTestCase, existInMap = testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]
		if existInMap == false {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":           "76ce3c8f-7791-44c9-8f18-be1ed0d9544d",
				"testCaseUuid": testCaseUuid,
			}).Fatal("TestCase doesn't exist in TestCaseMap. This should not happen")
		}

		templateFilesFromGithub = currentTestCase.ImportedTemplateFilesFromGitHub

		// A cell/row can be clickable to remove file when it is in first column and incoming 'filesAreClickable' is true
		// 1)
		var shouldBeClickable bool
		shouldBeClickable = id.Col == 0 && filesAreClickable == true

		switch shouldBeClickable {
		case true:
			// For the "Name" column, use the clickable label
			clickable := cell.(*clickableLabel)
			clickable.SetText(templateFilesFromGithub[id.Row].Name)
			clickable.isClickable = true

			clickable.onDoubleTap = func() {

				fmt.Println("Not Implemeted,yet!")
				/*
					// Remove the file from templateFilesFromGithub and refresh the list
					for fileIndex, file := range templateFilesFromGithub {
						if file.URL == templateFilesFromGithub[id.Row].URL {
							templateFilesFromGithub = append(templateFilesFromGithub[:fileIndex], templateFilesFromGithub[fileIndex+1:]...)
							*templateFilesFromGithubPtr = templateFilesFromGithub
							templateFilesTable.Unselect(id)
							templateFilesTable.Refresh()
							updateTemplateFilesTable(templateFilesTable, templateFilesFromGithubPtr, filesAreClickable)
							break
						}
					}

				*/

			}

		case false:
			// For the "URL" column, use a regular label

			nonClickable := cell.(*clickableLabel)
			nonClickable.SetText(templateFilesFromGithub[id.Row].URL)
		}
	}

	maxNameWidth := float32(150) // Start with a minimum width
	maxUrlWidth := float32(250)  // Start with a minimum width
	for _, file := range templateFilesFromGithub {
		textNameWidth := fyne.MeasureText(file.Name, theme.TextSize(), fyne.TextStyle{}).Width
		textUrlWidth := fyne.MeasureText(file.URL, theme.TextSize(), fyne.TextStyle{}).Width
		if textNameWidth > maxNameWidth {
			maxNameWidth = textNameWidth
		}
		if textUrlWidth > maxUrlWidth {
			maxUrlWidth = textUrlWidth
		}
	}

	templateFilesTable.SetColumnWidth(0, maxNameWidth+theme.Padding()*4) // Add padding
	templateFilesTable.SetColumnWidth(1, maxUrlWidth+theme.Padding()*4)  // Path column width can be static or calculated similarly

	templateFilesTable.Refresh()

}
