package listTestCaseExecutionsUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/executions/detailedExecutionsModel"
	"FenixTesterGui/testCaseExecutions/testCaseExecutionsModel"
	"bytes"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
	"image/color"
	"image/png"
	"sort"
	"strconv"
)

// RemoveTestCaseExecutionFromList
// Remove a TestCaseExecution from the List
func RemoveTestCaseExecutionFromList(testCaseExecutionUuidToBeRemoved string,
	testCaseExecutionsModel *testCaseExecutionsModel.TestCaseExecutionsModelStruct) {

	// Delete TestCase from 'TestCasesThatCanBeEditedByUserMap'
	delete(testCaseExecutionsModel.TestCaseExecutionsThatCanBeViewedByUserMap,
		testCaseExecutionUuidToBeRemoved)

	// Delete TestCase from 'TestCasesThatCanBeEditedByUserSlice'
	/*
		for index, tempTestCasesThatCanBeEditedByUser := range testCasesModel.TestCasesThatCanBeEditedByUserSlice {

			// Is this the TestCase to be removed from slice
			if tempTestCasesThatCanBeEditedByUser.TestCaseUuid == testCaseExecutionUuidToBeRemoved {

				// Remove TestCase at index
				testCasesModel.TestCasesThatCanBeEditedByUserSlice = append(
					testCasesModel.TestCasesThatCanBeEditedByUserSlice[:index],
					testCasesModel.TestCasesThatCanBeEditedByUserSlice[index+1:]...)

				break
			}

		}
	*/

	// Update table-list and update Table
	loadTestCaseExecutionListTableTable(testCaseExecutionsModel)
	calculateAndSetCorrectColumnWidths()
	updateTestCaseExecutionsListTable(testCaseExecutionsModel)

}

// Create the UI-list that holds the list of TestCases that the user can edit
func generateTestCaseExecutionsListTable(testCaseExecutionsModel *testCaseExecutionsModel.TestCaseExecutionsModelStruct) {

	// Correctly initialize the selectedFilesTable as a new table
	testCaseExecutionsListTable = widget.NewTable(
		func() (int, int) { return 0, numberColumnsInTestCaseExecutionsListUI }, // Start with zero rows, 8 columns
		func() fyne.CanvasObject {
			return widget.NewLabel("") // Create cells as labels
		},
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			// This should be filled when updating the table
		},
	)

	var err error

	// Load the Image, not active, for sort direction if not already done
	if sortImageUnspecifiedAsImage == nil {
		sortImageUnspecifiedAsImage, err = png.Decode(bytes.NewReader(sortUnspecifiedImageAsByteArray))
		if err != nil {
			sharedCode.Logger.WithFields(logrus.Fields{
				"Id":  "7b887b6e-9c90-4b4c-bc53-a5a750a463cb",
				"err": err,
			}).Fatalln("Failed to decode 'sortUnspecifiedImageAsByteArray'")
		}
	}

	// Load the Image, ascending, for sort direction if not already done
	if sortImageAscendingAsImage == nil {
		sortImageAscendingAsImage, err = png.Decode(bytes.NewReader(sortImageAscendingAsByteArray))
		if err != nil {
			sharedCode.Logger.WithFields(logrus.Fields{
				"Id":  "1e2e622c-afe1-45af-ac01-e2f0ef716b20",
				"err": err,
			}).Fatalln("Failed to decode 'sortImageAscendingAsByteArray'")
		}
	}

	// Load the Image, descending, for sort direction if not already done
	if sortImageDescendingAsImage == nil {
		sortImageDescendingAsImage, err = png.Decode(bytes.NewReader(sortImageDescendingAsByteArray))
		if err != nil {
			sharedCode.Logger.WithFields(logrus.Fields{
				"Id":  "a3826074-308d-4504-8695-85c95aba5eb3",
				"err": err,
			}).Fatalln("Failed to decode 'sortImageDescendingAsByteArray'")
		}
	}

	// Define the Header
	testCaseExecutionsListTable.ShowHeaderRow = true
	testCaseExecutionsListTable.CreateHeader = func() fyne.CanvasObject {
		//return widget.NewLabel("") // Create cells as labels

		var tempNewSortableHeaderLabel *sortableHeaderLabelStruct
		tempNewSortableHeaderLabel = newSortableHeaderLabel("", true, 0)

		// Create the Sort Icons container
		//var newSortIconsContainer *fyne.Container
		//newSortIconsContainer = container.NewStack(tempNewSortableHeaderLabel.sortImage.unspecifiedImageContainer, tempNewSortableHeaderLabel.sortImage)

		//var newSortableHeaderLabelContainer *fyne.Container
		//newSortableHeaderLabelContainer = container.NewHBox(
		//	widget.NewLabel(tempNewSortableHeaderLabel.Text), newSortIconsContainer) //canvas.NewImageFromImage(sortImageUnspecifiedAsImage)) //tempNewSortableHeaderLabel.sortImage)

		return tempNewSortableHeaderLabel

	}

	updateTestCaseExecutionsListTable(testCaseExecutionsModel)
	calculateAndSetCorrectColumnWidths()

}

// Update the Table
func updateTestCaseExecutionsListTable(testCaseExecutionsModel *testCaseExecutionsModel.TestCaseExecutionsModelStruct) {

	testCaseExecutionsListTable.Length = func() (int, int) {
		return len(testCaseExecutionsListTableTable), numberColumnsInTestCaseExecutionsListUI
	}
	testCaseExecutionsListTable.CreateCell = func() fyne.CanvasObject {

		tempNewClickableLabel := newClickableTableLabel("", func() {}, false, testCaseExecutionsModel)
		tempContainer := container.NewStack(canvas.NewRectangle(color.Transparent), tempNewClickableLabel, tempNewClickableLabel.textInsteadOfLabel)

		return tempContainer

	}
	testCaseExecutionsListTable.UpdateCell = func(id widget.TableCellID, cell fyne.CanvasObject) {

		clickableContainer := cell.(*fyne.Container)
		clickable := clickableContainer.Objects[1].(*clickableTableLabel)
		rectangle := clickableContainer.Objects[0].(*canvas.Rectangle)
		clickable.SetText(testCaseExecutionsListTableTable[id.Row][id.Col])
		clickable.isClickable = true
		clickable.currentRow = int16(id.Row)
		clickable.currentTestCaseExcutionUuid = testCaseExecutionsListTableTable[id.Row][testCaseExecutionUuidColumnNumber]

		clickable.onDoubleTap = func() {

			// Open TestCase
			openTestCase(clickable.currentTestCaseExcutionUuid, clickable.testCaseExecutionsModel)

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

			// Special handling for certain Columns for Status color and Timestamps
			switch uint8(id.Col) {

			case latestTestCaseExecutionStatus:

				clickable.textInsteadOfLabel.Text = clickable.Text
				clickable.Text = ""
				clickable.textInsteadOfLabel.Show()
				clickable.Hide()

			default:
				clickable.textInsteadOfLabel.Hide()
				clickable.Show()

			}

		} else {

			// If this row is the one that is shown in TestCase preview window
			if clickable.currentTestCaseExcutionUuid == testCaseExecutionThatIsShownInPreview {

				clickable.TextStyle = fyne.TextStyle{Bold: false}
				rectangle.FillColor = color.RGBA{
					R: 0x08,
					G: 0x5C,
					B: 0x04,
					A: 0xFF,
				}
				rectangle.StrokeColor = color.Transparent
				rectangle.StrokeWidth = 3

			} else {

				clickable.TextStyle = fyne.TextStyle{Bold: false}
				rectangle.FillColor = color.Transparent
				rectangle.StrokeColor = color.Transparent
				rectangle.StrokeWidth = 3

			}
			// Special handling for certain Columns for Status color and Timestamps
			switch uint8(id.Col) {

			case latestTestCaseExecutionStatus:
				var statusId uint8
				var statusBackgroundColor color.RGBA
				var statusStrokeColor color.RGBA
				var useStroke bool

				statusId = detailedExecutionsModel.ExecutionStatusColorNameToNumberMap[clickable.Text].ExecutionStatusNumber
				statusBackgroundColor = detailedExecutionsModel.ExecutionStatusColorMap[int32(statusId)].BackgroundColor
				rectangle.FillColor = statusBackgroundColor

				useStroke = detailedExecutionsModel.ExecutionStatusColorMap[int32(statusId)].UseStroke
				if useStroke == true {
					statusStrokeColor = detailedExecutionsModel.ExecutionStatusColorMap[int32(statusId)].StrokeColor
					rectangle.StrokeColor = statusStrokeColor
				}

				// When no background color
				if statusBackgroundColor.R+statusBackgroundColor.G+statusBackgroundColor.B+statusBackgroundColor.A == 0 {

					clickable.Alignment = fyne.TextAlignCenter
					clickable.textInsteadOfLabel.Hide()
					clickable.Show()

				} else {

					clickable.textInsteadOfLabel.Text = clickable.Text
					clickable.textInsteadOfLabel.TextStyle = clickable.TextStyle
					clickable.Text = ""
					clickable.textInsteadOfLabel.Show()
					clickable.Hide()

				}

			default:

			}

		}
		clickableContainer.Refresh()

	}

	// Update the Header
	testCaseExecutionsListTable.UpdateHeader = func(id widget.TableCellID, cell fyne.CanvasObject) {
		tempSortableHeaderLabel := cell.(*sortableHeaderLabelStruct)
		tempSortableHeaderLabel.label.SetText(testCaseExecutionsListTableHeader[id.Col])
		tempSortableHeaderLabel.label.TextStyle = fyne.TextStyle{Bold: true}

		// Set Column number
		tempSortableHeaderLabel.columnNumber = id.Col
		tempSortableHeaderLabel.sortImage.headerColumnNumber = id.Col

		//tempSortableHeaderLabel.latestSelectedSortOrder = SortingDirectionAscending
		//tempSortableHeaderLabel.updateColumnNumberFunction()

		// Refresh the widget to update the UI
		tempSortableHeaderLabel.Refresh()
	}

	testCaseExecutionsListTable.Refresh()
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
	columnsMaxSizeSlice = make([]float32, numberColumnsInTestCaseExecutionsListUI)

	var columnWidth float32

	// Set initial value for max width size
	for index, headerValue := range testCaseExecutionsListTableHeader {

		// Calculate the column width base on this value. Add  'float32(30)' to give room for sort direction icon
		columnWidth = fyne.MeasureText(headerValue, theme.TextSize(), fyne.TextStyle{Bold: true}).Width + float32(30) //
		columnsMaxSizeSlice[index] = columnWidth
	}

	// Loop all rows
	for _, tempRow := range testCaseExecutionsListTableTable {

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
		testCaseExecutionsListTable.SetColumnWidth(columnIndex, columnWidth+theme.Padding()*4)
	}

	// Refresh the table
	testCaseExecutionsListTable.Refresh()

}

func loadTestCaseExecutionListTableTable(testCaseExecutionsModel *testCaseExecutionsModel.TestCaseExecutionsModelStruct) {

	testCaseExecutionsListTableTable = nil

	// Loop all TestCases and add to '[][]string'-object for the Table
	for _, tempTestCaseExecution := range testCaseExecutionsModel.TestCaseExecutionsThatCanBeViewedByUserMap {

		// Create temporary Row-object for the table
		var tempRowslice []string

		// Populate the temporary Row-object

		// Column 0:
		// DomainName
		var domainNameForTable string
		domainNameForTable = fmt.Sprintf("%s [%s]",
			tempTestCaseExecution.GetDomainName(),
			tempTestCaseExecution.GetDomainUUID()[0:8])

		tempRowslice = append(tempRowslice, domainNameForTable)

		// Column 1:
		// TestCaseName
		tempRowslice = append(tempRowslice, tempTestCaseExecution.GetTestCaseName())

		// Column 2:
		// TestCaseUuid
		tempRowslice = append(tempRowslice, tempTestCaseExecution.GetTestCaseUuid())

		// Column 3:
		// TestCaseVersion
		tempRowslice = append(tempRowslice, strconv.Itoa(int(tempTestCaseExecution.GetTestCaseVersion())))

		// Column 4:
		// LatestTestCaseExecutionStatus
		var tempLatestTestCaseExecutionStatus string

		if tempTestCaseExecution.GetLatestTestCaseExecutionStatus() > 0 {

			tempLatestTestCaseExecutionStatus = detailedExecutionsModel.ExecutionStatusColorMap[int32(tempTestCaseExecution.GetLatestTestCaseExecutionStatus())].ExecutionStatusName
		} else {
			tempLatestTestCaseExecutionStatus = "<no execution>"
		}

		tempRowslice = append(tempRowslice, tempLatestTestCaseExecutionStatus)

		// Column 5:
		// LatestTestCaseExecutionStatusInsertTimeStamp
		var tempLatestTestCaseExecutionStatusInsertTimeStamp string

		if tempTestCaseExecution.GetLatestTestCaseExecutionStatusInsertTimeStamp() != nil {
			tempLatestTestCaseExecutionStatusInsertTimeStamp = sharedCode.ConvertGrpcTimeStampToStringForDB(tempTestCaseExecution.
				GetLatestTestCaseExecutionStatusInsertTimeStamp())
		} else {
			tempLatestTestCaseExecutionStatusInsertTimeStamp = "<no execution>"
		}
		tempRowslice = append(tempRowslice, tempLatestTestCaseExecutionStatusInsertTimeStamp)

		// Column 6:
		// LatestFinishedOkTestCaseExecutionStatusInsertTimeStamp
		var tempLatestFinishedOkTestCaseExecutionStatusInsertTimeStamp string

		if tempTestCaseExecution.GetLatestFinishedOkTestCaseExecutionStatusInsertTimeStamp() != nil {
			tempLatestFinishedOkTestCaseExecutionStatusInsertTimeStamp = sharedCode.ConvertGrpcTimeStampToStringForDB(
				tempTestCaseExecution.GetLatestFinishedOkTestCaseExecutionStatusInsertTimeStamp())
		} else {
			tempLatestFinishedOkTestCaseExecutionStatusInsertTimeStamp = "<no successful execution yet>"
		}
		tempRowslice = append(tempRowslice, tempLatestFinishedOkTestCaseExecutionStatusInsertTimeStamp)

		// Column 8:
		// LastSavedTimeStamp
		var tempLastSavedTimeStamp string

		if tempTestCaseExecution.GetLastSavedTimeStamp() != nil {
			tempLastSavedTimeStamp = sharedCode.ConvertGrpcTimeStampToStringForDB(tempTestCaseExecution.
				GetLastSavedTimeStamp())
		} else {
			tempLastSavedTimeStamp = "<This should not happen, due to it must have been saved!>"
		}
		tempRowslice = append(tempRowslice, tempLastSavedTimeStamp)

		// Column 8:
		// DomainUuid
		tempRowslice = append(tempRowslice, tempTestCaseExecution.GetDomainUuid())

		// Add Row to slice of rows for the table
		testCaseExecutionsListTableTable = append(testCaseExecutionsListTableTable, tempRowslice)

	}

	// Do an initial sort 'testCaseExecutionsListTableTable' descending on 'LastSavedTimeStamp'
	if testCaseExecutionsModel.TestCaseExecutionsThatCanBeViewedByUserMap != nil &&
		len(testCaseExecutionsModel.TestCaseExecutionsThatCanBeViewedByUserMap) > 0 {

		currentSortColumn = initialColumnToSortOn
		sort2DStringSlice(testCaseExecutionsListTableTable, initialColumnToSortOn, initialSortDirectionForInitialColumnToSortOn)

	}

}

// Sort2DStringSlice sorts a 2D string slice by a specified column index.
// It assumes that the column index is valid for all rows in the slice.
func sort2DStringSlice(data [][]string, columnToSortOn int, sortingDirection SortingDirectionType) {
	sort.Slice(data, func(i, j int) bool {
		// Adjust the sorting logic as needed.
		// In this case, sorting lexicographically based on the given columnToSortOn
		// You can modify this to handle numeric sorting, etc.

		// Example: sorting as numbers, if needed.
		// First, try to convert strings to integers. If it fails, fall back to string comparison.
		num1, err1 := strconv.Atoi(data[i][columnToSortOn])
		num2, err2 := strconv.Atoi(data[j][columnToSortOn])

		// Handle sorting direction
		switch sortingDirection {

		case SortingDirectionAscending:

			if err1 == nil && err2 == nil {
				return num1 < num2
			}

			// Default to string comparison if not numbers.
			return data[i][columnToSortOn] < data[j][columnToSortOn]

		case SortingDirectionDescending:

			if err1 == nil && err2 == nil {
				return num1 > num2
			}

			// Default to string comparison if not numbers.
			return data[i][columnToSortOn] > data[j][columnToSortOn]
		}

		// Not important due that switch statement will handle all return values
		return true

	})
}
