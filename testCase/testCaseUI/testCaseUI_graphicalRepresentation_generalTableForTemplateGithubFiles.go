package testCaseUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/importFilesFromGitHub"
	"FenixTesterGui/testCase/testCaseModel"
	"FenixTesterGui/testCase/testCaseUI/templateViewer"
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
func generateTemplateFilesTable(
	testCaseUuid string,
	filesAreClickable bool,
	testCasesUiCanvasObject *TestCasesUiModelStruct) *CustomTemplateTable {

	var existInMap bool
	var currentTestCase testCaseModel.TestCaseModelStruct

	currentTestCase, existInMap = testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]
	if existInMap == false {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":           "6fb0f1ff-9e16-4576-ae7d-10915065e15f",
			"testCaseUuid": testCaseUuid,
		}).Fatal("TestCase doesn't exist in TestCaseMap. This should not happen")
	}

	var templateFilesFromGithub []importFilesFromGitHub.GitHubFile
	templateFilesFromGithub = currentTestCase.ImportedTemplateFilesFromGitHub

	// Correctly initialize the templateFilesTable as a new table
	templateFilesTable := &CustomTemplateTable{
		widget.Table{
			Length: func() (int, int) {

				currentTestCase, existInMap = testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]
				if existInMap == false {
					sharedCode.Logger.WithFields(logrus.Fields{
						"ID":           "ebf42b07-76d6-4186-aebd-4cf18beb9f1d",
						"testCaseUuid": testCaseUuid,
					}).Fatal("TestCase doesn't exist in TestCaseMap. This should not happen")
				}

				templateFilesFromGithub = currentTestCase.ImportedTemplateFilesFromGitHub

				if templateFilesFromGithub == nil {
					fmt.Println("templateFilesFromGithub == nil")
					return 0, 2
				}
				fmt.Println("len(templateFilesFromGithub)")
				return len(templateFilesFromGithub), 2
			}, // Start with zero rows, 2 columns
			//func() fyne.CanvasObject {
			//	return widget.NewLabel("") // Create cells as labels
			CreateCell: func() fyne.CanvasObject {

				fmt.Println("CreateCell_01")
				return newClickableLabel("",
					func() {},
					false)
			},
			UpdateCell: func(id widget.TableCellID, cell fyne.CanvasObject) {
				fmt.Println("UpdateCell_01", id)

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
				filesAreClickable = true
				var shouldBeClickable bool
				shouldBeClickable = id.Col == 0 && filesAreClickable == true

				switch shouldBeClickable {
				case true:
					// For the "Name" column, use the clickable label
					clickable := cell.(*clickableLabel)
					clickable.SetText(templateFilesFromGithub[id.Row].Name)
					clickable.isClickable = true

					clickable.onDoubleTap = func() {

						templateViewer.InitiateTemplateViewer(
							*sharedCode.FenixMasterWindowPtr,
							*sharedCode.FenixAppPtr,
							&currentTestCase.ImportedTemplateFilesFromGitHub,
							currentTestCase.TestData,
							testCaseUuid,
							templateFilesFromGithub[id.Row].Name,
							testDataPointGroupsSelectSelectedInMainTestCaseArea,
							testDataPointForAGroupSelectSelectedInMainTestCaseArea,
							testDataRowForTestDataPointsSelectSelectedInMainTestCaseArea)
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
			},
		},
	}

	templateFilesTable.ExtendBaseWidget(templateFilesTable)

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

// Update the Table used for showing Templates in TestCase
func updateTemplateFilesTable(
	templateFilesTable *CustomTemplateTable,
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
		fmt.Println("Length", len(templateFilesFromGithub), 2)
		return len(templateFilesFromGithub), 2
	}
	/*templateFilesTable.CreateCell = func() fyne.CanvasObject {
		fmt.Println("CreateCell")
		return newClickableLabel("", func() {}, false)

	}*/
	templateFilesTable.UpdateCell = func(id widget.TableCellID, cell fyne.CanvasObject) {
		fmt.Println("UpdateCell", id)
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

				fmt.Println("Not Implemented,yet!")
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

	/*var oneRowHeight float32
	var numberOfTemplates int
	if len(templateFilesFromGithub) > 0 {
		oneRowHeight = fyne.MeasureText(
			templateFilesFromGithub[0].Name, theme.TextSize(), fyne.TextStyle{}).Height
		numberOfTemplates = len(templateFilesFromGithub)
	}

	*/
}

// Update size of columns and Rows for the Table
func (t *CustomTemplateTable) updateColumnAndRowSizes(
	testCaseUuid string,
	testCasesUiCanvasObject *TestCasesUiModelStruct,
	checkIfTemplatesAreChangedButton *widget.Button,
	viewTemplateButton *widget.Button) {

	var existInMap bool
	var currentTestCase testCaseModel.TestCaseModelStruct

	currentTestCase, existInMap = testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]
	if existInMap == false {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":           "34d04c69-a6e4-44f7-bbf2-c891268ac3b8",
			"testCaseUuid": testCaseUuid,
		}).Fatal("TestCase doesn't exist in TestCaseMap. This should not happen")
	}

	var templateFilesFromGithub []importFilesFromGitHub.GitHubFile
	templateFilesFromGithub = currentTestCase.ImportedTemplateFilesFromGitHub

	if templateFilesFromGithub != nil || len(templateFilesFromGithub) > 0 {
		var nameWidth, maxNameWidth, urlWidth, maxUrlWidth float32
		maxNameWidth = float32(150) // Start with a minimum width
		maxUrlWidth = float32(250)  // Start with a minimum width
		rowHeight := fyne.MeasureText("Text", theme.TextSize(), fyne.TextStyle{}).Height
		for fileCounter, file := range templateFilesFromGithub {
			nameWidth = fyne.MeasureText(file.Name, theme.TextSize(), fyne.TextStyle{}).Width
			urlWidth = fyne.MeasureText(file.URL, theme.TextSize(), fyne.TextStyle{}).Width
			if nameWidth > maxNameWidth {
				maxNameWidth = nameWidth
			}
			if urlWidth > maxUrlWidth {
				maxUrlWidth = urlWidth
			}

			// Set RowHeight
			t.SetRowHeight(fileCounter, rowHeight)
		}

		t.SetColumnWidth(0, maxNameWidth+theme.Padding()*4) // Add padding
		t.SetColumnWidth(1, maxUrlWidth+theme.Padding()*4)  // Path column width can be static or calculated similarly

	}
	//templateFilesTable.Resize(fyne.NewSize(
	//	templateFilesTable.Size().Width, oneRowHeight*float32(numberOfTemplates)))

	t.Refresh()

	// Enable and Disable 'checkIfTemplatesAreChangedButton' and 'viewTemplateButton' depending on number of Templates
	if templateFilesFromGithub == nil || len(templateFilesFromGithub) == 0 {
		checkIfTemplatesAreChangedButton.Disable()
		viewTemplateButton.Disable()
	} else {
		checkIfTemplatesAreChangedButton.Enable()
		viewTemplateButton.Enable()
	}

}

// CustomTemplateTable is a struct that extends the standard Fyne Table widget.
type CustomTemplateTable struct {
	widget.Table
}

// NewCustomTemplateTable creates a new instance of CustomTemplateTable with custom behavior.
func NewCustomTemplateTable(rowCount, colCount int) *CustomTemplateTable {
	// Create the base table with the necessary parameters.
	baseTable := widget.NewTable(
		func() (int, int) {
			fmt.Println("baseTable-Length")
			return 0, 2
		}, // Start with zero rows, 2 columns
		func() fyne.CanvasObject {
			fmt.Println("baseTable-NewCell")
			return newClickableLabel("", func() {}, false) // Create cells as labels
		},
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			fmt.Println("baseTable-Change")
			// This should be filled when updating the table
			fmt.Println(id)

		},
	)

	// Return the custom table
	return &CustomTemplateTable{*baseTable}
}

/*
// ExtendBaseWidget is needed to satisfy the fyne.Widget interface.
func (t *CustomTemplateTable) ExtendBaseWidget(w fyne.Widget) {
	t.Table.ExtendBaseWidget(w)
}



func (t *CustomTemplateTable) CreateCell() fyne.CanvasObject {

	fmt.Println("CreateCell 2")

	return newClickableLabel("", func() {}, false)

}

*/

// MinSize customizes the minimum size of the CustomTemplateTable.
func (t *CustomTemplateTable) MinSize() fyne.Size {
	// You can override the MinSize to change how the table is sized.
	// Here, we calculate the size based on the number of rows.
	rowHeight := fyne.MeasureText("Text", theme.TextSize(), fyne.TextStyle{}).Height
	tableŔows, _ := t.Length()
	if tableŔows < 10 {
		tableŔows = 10
	}
	totalHeight := float32(tableŔows) * rowHeight
	fmt.Println(tableŔows, t.Size().Width, totalHeight)
	return fyne.NewSize(t.Size().Width, totalHeight)
}
