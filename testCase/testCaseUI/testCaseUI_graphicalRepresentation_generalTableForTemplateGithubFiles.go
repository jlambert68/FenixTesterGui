package testCaseUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/importFilesFromGitHub"
	"FenixTesterGui/testCase/testCaseModel"
	"FenixTesterGui/testCase/testCaseUI/templateViewer"
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
	var currentTestCasePtr *testCaseModel.TestCaseModelStruct

	currentTestCasePtr, existInMap = testCasesUiCanvasObject.TestCasesModelReference.TestCasesMapPtr[testCaseUuid]
	if existInMap == false {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":           "6fb0f1ff-9e16-4576-ae7d-10915065e15f",
			"testCaseUuid": testCaseUuid,
		}).Fatal("TestCase doesn't exist in TestCaseMap. This should not happen")
	}

	var templateFilesFromGithub []importFilesFromGitHub.GitHubFile
	templateFilesFromGithub = currentTestCasePtr.ImportedTemplateFilesFromGitHub

	// Correctly initialize the templateFilesTable as a new table
	templateFilesTable := &CustomTemplateTable{
		widget.Table{
			Length: func() (int, int) {

				currentTestCasePtr, existInMap = testCasesUiCanvasObject.TestCasesModelReference.TestCasesMapPtr[testCaseUuid]
				if existInMap == false {

					sharedCode.Logger.WithFields(logrus.Fields{
						"ID":           "ebf42b07-76d6-4186-aebd-4cf18beb9f1d",
						"testCaseUuid": testCaseUuid,
					}).Warning("TestCase doesn't exist in TestCaseMap. This should not happen")
				}

				templateFilesFromGithub = currentTestCasePtr.ImportedTemplateFilesFromGitHub

				if templateFilesFromGithub == nil {
					return 0, 2
				}

				return len(templateFilesFromGithub), 2
			}, // Start with zero rows, 2 columns
			//func() fyne.CanvasObject {
			//	return widget.NewLabel("") // Create cells as labels
			CreateCell: func() fyne.CanvasObject {

				return newClickableLabel("",
					func() {},
					false)
			},
			UpdateCell: func(id widget.TableCellID, cell fyne.CanvasObject) {

				currentTestCasePtr, existInMap = testCasesUiCanvasObject.TestCasesModelReference.TestCasesMapPtr[testCaseUuid]
				if existInMap == false {
					sharedCode.Logger.WithFields(logrus.Fields{
						"ID":           "76ce3c8f-7791-44c9-8f18-be1ed0d9544d",
						"testCaseUuid": testCaseUuid,
					}).Fatal("TestCase doesn't exist in TestCaseMap. This should not happen")
				}

				templateFilesFromGithub = currentTestCasePtr.ImportedTemplateFilesFromGitHub

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
							&currentTestCasePtr.ImportedTemplateFilesFromGitHub,
							currentTestCasePtr.TestData,
							testCaseUuid,
							templateFilesFromGithub[id.Row].Name,
							testDataPointGroupsSelectSelectedInMainTestCaseArea,
							testDataPointForAGroupSelectSelectedInMainTestCaseArea,
							testDataRowForTestDataPointsSelectSelectedInMainTestCaseArea)

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

// Update size of columns and Rows for the Table
func (t *CustomTemplateTable) updateColumnAndRowSizes(
	testCaseUuid string,
	testCasesUiCanvasObject *TestCasesUiModelStruct,
	checkIfTemplatesAreChangedButton *widget.Button,
	viewTemplateButton *widget.Button) {

	var existInMap bool
	var currentTestCasePtr *testCaseModel.TestCaseModelStruct

	currentTestCasePtr, existInMap = testCasesUiCanvasObject.TestCasesModelReference.TestCasesMapPtr[testCaseUuid]
	if existInMap == false {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":           "34d04c69-a6e4-44f7-bbf2-c891268ac3b8",
			"testCaseUuid": testCaseUuid,
		}).Fatal("TestCase doesn't exist in TestCaseMap. This should not happen")
	}

	var templateFilesFromGithub []importFilesFromGitHub.GitHubFile
	templateFilesFromGithub = currentTestCasePtr.ImportedTemplateFilesFromGitHub

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
	return fyne.NewSize(t.Size().Width, totalHeight)
}
