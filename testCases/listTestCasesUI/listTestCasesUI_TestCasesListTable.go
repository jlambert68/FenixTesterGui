package listTestCasesUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testCase/testCaseModel"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"image/color"
	"strconv"
	"time"
)

type clickableLabel struct {
	widget.Label
	onDoubleTap         func()
	lastTapTime         time.Time
	isClickable         bool
	currentRow          int16
	currentTestCaseUuid string
	background          *canvas.Rectangle
	testCasesModel      *testCaseModel.TestCasesModelsStruct
}

func newClickableLabel(text string, onDoubleTap func(), tempIsClickable bool,
	testCasesModel *testCaseModel.TestCasesModelsStruct) *clickableLabel {

	l := &clickableLabel{
		Label:       widget.Label{Text: text},
		onDoubleTap: onDoubleTap,
		lastTapTime: time.Now(),
		isClickable: tempIsClickable,
		currentRow:  -1}

	l.background = canvas.NewRectangle(color.Transparent)
	l.testCasesModel = testCasesModel
	l.currentTestCaseUuid = ""

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

// TappedSecondary
// Implement if you need right-click (secondary tap) actions.
func (l *clickableLabel) TappedSecondary(*fyne.PointEvent) {
	if l.isClickable == false {
		return
	}

	fenixMasterWindow := *sharedCode.FenixMasterWindowPtr
	clipboard := fenixMasterWindow.Clipboard()
	clipboard.SetContent(l.Text)

	// Optional: Notify the user
	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   "Clipboard",
		Content: fmt.Sprintf("'%s' copied to clipboard!", l.Text),
	})

}

func (l *clickableLabel) MouseIn(*desktop.MouseEvent) {

	if l.isClickable == false {
		return
	}

	// Hinder concurrent setting of variable
	currentRowThatMouseIsHoveringAboveMutex.Lock()

	// Set current TestCaseUuid to be highlighted
	currentRowThatMouseIsHoveringAbove = l.currentRow

	// Release variable
	currentRowThatMouseIsHoveringAboveMutex.Unlock()

	l.TextStyle = fyne.TextStyle{Bold: true}
	l.Refresh()
	testCaseListTable.Refresh()

}
func (l *clickableLabel) MouseMoved(*desktop.MouseEvent) {}
func (l *clickableLabel) MouseOut() {

	if l.isClickable == false {
		return
	}

	// Hinder concurrent setting of variable
	currentRowThatMouseIsHoveringAboveMutex.Lock()

	// Set current TestCaseUuid to be highlighted
	currentRowThatMouseIsHoveringAbove = -1

	// Release variable
	currentRowThatMouseIsHoveringAboveMutex.Unlock()

	l.TextStyle = fyne.TextStyle{Bold: false}
	l.Refresh()
	testCaseListTable.Refresh()

}

// Create the UI-list that holds the list of TestCases that the user can edit
func generateTestCasesListTable(testCasesModel *testCaseModel.TestCasesModelsStruct) {

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

	testCaseListTable.ShowHeaderRow = true
	testCaseListTable.CreateHeader = func() fyne.CanvasObject {
		return widget.NewLabel("") // Create cells as labels
	}

	updateTestCasesListTable(testCasesModel)
	calculateAndSetCorrectColumnWidths()

}

// Update the Table
func updateTestCasesListTable(testCasesModel *testCaseModel.TestCasesModelsStruct) {

	testCaseListTable.Length = func() (int, int) {
		return len(testCaseListTableTable), 8
	}
	testCaseListTable.CreateCell = func() fyne.CanvasObject {

		tempNewClickableLabel := newClickableLabel("", func() {}, false, testCasesModel)
		tempContainer := container.NewStack(canvas.NewRectangle(color.Transparent), tempNewClickableLabel)

		return tempContainer

	}
	testCaseListTable.UpdateCell = func(id widget.TableCellID, cell fyne.CanvasObject) {

		clickableContainer := cell.(*fyne.Container)
		clickable := clickableContainer.Objects[1].(*clickableLabel)
		rectangle := clickableContainer.Objects[0].(*canvas.Rectangle)
		clickable.SetText(testCaseListTableTable[id.Row][id.Col])
		clickable.isClickable = true
		clickable.currentRow = int16(id.Row)
		clickable.currentTestCaseUuid = testCaseListTableTable[id.Row][testCaseUuidColumnNumber]

		clickable.onDoubleTap = func() {

			// Open TestCase
			openTestCase(clickable.currentTestCaseUuid, clickable.testCasesModel)

		}

		// Check if this row should be highlighted or not
		if int16(id.Row) == currentRowThatMouseIsHoveringAbove {
			clickable.TextStyle = fyne.TextStyle{Bold: false}
			rectangle.FillColor = color.RGBA{
				R: 0x4A,
				G: 0x4B,
				B: 0x4D,
				A: 0xFF,
			}

		} else {
			clickable.TextStyle = fyne.TextStyle{Bold: false}
			rectangle.FillColor = color.Transparent
		}
		clickableContainer.Refresh()

	}

	testCaseListTable.UpdateHeader = func(id widget.TableCellID, cell fyne.CanvasObject) {
		clickable := cell.(*widget.Label)
		clickable.SetText(testCaseListTableHeader[id.Col])
		clickable.TextStyle = fyne.TextStyle{Bold: true}
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

	var columnWidth float32

	// Set initial value for max width size
	for index, headerValue := range testCaseListTableHeader {

		// Calculate the column width base on this value
		columnWidth = fyne.MeasureText(headerValue, theme.TextSize(), fyne.TextStyle{Bold: true}).Width
		columnsMaxSizeSlice[index] = columnWidth
	}

	// Loop all rows
	for _, tempRow := range testCaseListTableTable {

		// Loop columns for a row to get column width
		for columnIndex, tempColumnValue := range tempRow {

			// Calculate the column width base on this value
			columnWidth = fyne.MeasureText(tempColumnValue, theme.TextSize(), fyne.TextStyle{Bold: true}).Width

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

func loadTestCaseListTableTable(testCasesModel *testCaseModel.TestCasesModelsStruct) {

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
