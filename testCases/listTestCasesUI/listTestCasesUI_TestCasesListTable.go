package listTestCasesUI

import (
	"FenixTesterGui/testCase/testCaseModel"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"strconv"
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

func (l *clickableLabel) MouseIn(*desktop.MouseEvent) {
	if l.isClickable == false {
		return
	}

	l.TextStyle = fyne.TextStyle{Bold: true}
	l.Refresh()

}
func (l *clickableLabel) MouseMoved(*desktop.MouseEvent) {}
func (l *clickableLabel) MouseOut() {
	if l.isClickable == false {
		return
	}

	l.TextStyle = fyne.TextStyle{Bold: false}
	l.Refresh()

}

// Create the UI-list that holds the list of TestCases that the user can edit
func generateTestCasesListTable() {

	// Correctly initialize the selectedFilesTable as a new table
	testCaseListTable = widget.NewTable(
		func() (int, int) { return 0, 8 }, // Start with zero rows, 8 columns
		func() fyne.CanvasObject {
			return widget.NewLabel("") // Create cells as labels
		},
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			// This should be filled when updating the table
		},
	)

}

// Update the Table
func updateTestCasesListTable() {

	testCaseListTable.Length = func() (int, int) {
		return len(testCaseListTableTable), 8
	}
	testCaseListTable.CreateCell = func() fyne.CanvasObject {
		return newClickableLabel("", func() {}, false)

	}
	testCaseListTable.UpdateCell = func(id widget.TableCellID, cell fyne.CanvasObject) {

		clickable := cell.(*clickableLabel)
		clickable.SetText(testCaseListTableTable[id.Row][id.Col])
		clickable.isClickable = true

		clickable.onDoubleTap = func() {

			// TODO Open TestCase

		}

	}

	testCaseListTable.Refresh()
}

// TestCaseUuid
// TestCaseVersion
// LatestTestCaseExecutionStatus
// LatestTestCaseExecutionStatusInsertTimeStamp
// LatestFinishedOkTestCaseExecutionStatusInsertTimeStamp
// DomainUuid

func calculateAndSetCorrectColumnWidths() {

	// Initiate slice for keeping track of max column width size
	var columnsMaxSizeSlice []float32
	columnsMaxSizeSlice = make([]float32, 8)

	// Set initial value for max width size
	for index, _ := range columnsMaxSizeSlice {
		columnsMaxSizeSlice[index] = 0
	}

	var columnWidth float32

	// Loop all rows
	for _, tempRow := range testCaseListTableTable {

		// Loop columns for a row to get column width
		for columnIndex, tempColumnValue := range tempRow {

			// Calculate the column width base on this value
			columnWidth = fyne.MeasureText(tempColumnValue, theme.TextSize(), fyne.TextStyle{}).Width

			// If this value is bigger than current then set this to max column size
			if columnWidth > columnsMaxSizeSlice[columnIndex] {
				columnsMaxSizeSlice[columnIndex] = columnWidth
			}

		}

	}

	// Loop columns in table and set column width including some Padding
	for columnIndex, columnWidth := range columnsMaxSizeSlice {
		testCaseListTable.SetColumnWidth(columnIndex, columnWidth+theme.Padding()*4)
	}

	// Refresh the table
	testCaseListTable.Refresh()

}

func loadTestCaseListTableTable(testCasesModel testCaseModel.TestCasesModelsStruct) {

	testCaseListTableTable = nil

	// Loop all TestCases and add to '[][]string'-object for the Table
	for _, tempTestCase := range testCasesModel.TestCasesThatCanBeEditedByUserSlice {

		// Create temporary Row-object for the table
		var tempRowslice []string

		// Populate the temporary Row-object

		// Column 0:
		// DomainName
		var domainNameForTable string
		domainNameForTable = fmt.Sprintf("%s [%s]",
			tempTestCase.GetDomainName(),
			tempTestCase.GetTestCaseUuid()[0:8])

		tempRowslice = append(tempRowslice, domainNameForTable)

		// Column 1:
		// TestCaseName
		tempRowslice = append(tempRowslice, tempTestCase.GetTestCaseName())

		// Column 2:
		// TestCaseUuid
		tempRowslice = append(tempRowslice, tempTestCase.GetTestCaseUuid())

		// Column 3:
		// TestCaseVersion
		tempRowslice = append(tempRowslice, strconv.Itoa(int(tempTestCase.GetTestCaseVersion())))

		// Column 4:
		// LatestTestCaseExecutionStatus
		tempRowslice = append(tempRowslice, fenixGuiTestCaseBuilderServerGrpcApi.TestCaseExecutionStatusEnum_name[int32(
			tempTestCase.GetLatestTestCaseExecutionStatus())])

		// Column 5:
		// LatestTestCaseExecutionStatusInsertTimeStamp
		tempRowslice = append(tempRowslice, tempTestCase.
			GetLatestTestCaseExecutionStatusInsertTimeStamp().String())

		// Column 6:
		// LatestFinishedOkTestCaseExecutionStatusInsertTimeStamp
		tempRowslice = append(tempRowslice, tempTestCase.
			GetLatestFinishedOkTestCaseExecutionStatusInsertTimeStamp().String())

		// Column 7:
		// DomainUuid
		tempRowslice = append(tempRowslice, tempTestCase.GetDomainUuid())

		// Add Row to slice of rows for the table
		testCaseListTableTable = append(testCaseListTableTable, tempRowslice)

	}

}

type customLabel struct {
	widget.Label
	onDoubleTap func()
	lastTap     time.Time
}

func newCustomLabel(text string, onDoubleTap func()) *customLabel {
	l := &customLabel{Label: widget.Label{Text: text}, onDoubleTap: onDoubleTap, lastTap: time.Now()}
	l.ExtendBaseWidget(l)
	return l
}

func (l *customLabel) Tapped(e *fyne.PointEvent) {
	now := time.Now()
	if now.Sub(l.lastTap) < 500*time.Millisecond { // 500 ms as double-click interval
		if l.onDoubleTap != nil {
			l.onDoubleTap()
		}
	}
	l.lastTap = now
}

func (l *customLabel) TappedSecondary(*fyne.PointEvent) {
	// Implement if you need right-click (secondary tap) actions.
}

func (l *customLabel) MouseIn(*desktop.MouseEvent)    {}
func (l *customLabel) MouseMoved(*desktop.MouseEvent) {}
func (l *customLabel) MouseOut()                      {}

/*
type coloredLabelItem struct {
	text  string
	color color.Color
}

func (importFilesFromGitHubObject *ImportFilesFromGitHubStruct) newColoredLabelItem(text string, color color.Color) *coloredLabelItem {
	return &coloredLabelItem{text: text, color: color}
}

func (importFilesFromGitHubObject *ImportFilesFromGitHubStruct) (item *coloredLabelItem) CreateRenderer() fyne.WidgetRenderer {
	label := widget.NewLabel(item.text)
	label.color = item.color
	label.Refresh()

	return widget.NewSimpleRenderer(label)
}

*/
