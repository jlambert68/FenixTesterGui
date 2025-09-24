package listTestSuiteExecutionsUI

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
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"image/color"
	"image/png"
	"log"
	"sort"
	"strconv"
	"sync"
)

// RemoveTestCaseExecutionFromList
// Remove a TestCaseExecution from the List
func RemoveTestCaseExecutionFromList(testCaseExecutionUuidToBeRemoved string,
	testCaseExecutionsModelRef *testCaseExecutionsModel.TestCaseExecutionsModelStruct) {

	// Delete TestCaseExecution from 'testCaseExecutionsThatCanBeViewedByUserMap'
	testCaseExecutionsModelRef.DeleteFromTestCaseExecutionsMap(
		testCaseExecutionsModel.TestCaseExecutionUuidType(testCaseExecutionUuidToBeRemoved))

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


		// Update table-list and update Table
		loadTestCaseExecutionListTableTable(testCaseExecutionsModelRef)
		calculateAndSetCorrectColumnWidths()
		updateTestCaseExecutionsListTable(testCaseExecutionsModelRef)
	*/
}

// Create the UI-list that holds the list of TestCasesMapPtr that the user can edit
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

var updateTestCaseExecutionsListTableMutex = &sync.RWMutex{}

// Update the Table
func updateTestCaseExecutionsListTable(testCaseExecutionsModel *testCaseExecutionsModel.TestCaseExecutionsModelStruct) {

	// Lock function
	updateTestCaseExecutionsListTableMutex.Lock()

	// UnLock function
	defer updateTestCaseExecutionsListTableMutex.Unlock()

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
		clickable.currentTestCaseExecutionUuid = testCaseExecutionsListTableTable[id.Row][testCaseExecutionUuidColumnNumber]
		clickable.currentTestCaseExecutionVersion = 1 //TODO HArdcoded  for now testCaseExecutionsListTableTable[id.Row][testCaseExecutionVersionColumnNumber]
		clickable.currentTestCaseUuid = testCaseExecutionsListTableTable[id.Row][testCaseUuidColumnNumber]
		clickable.currentTestCaseName = testCaseExecutionsListTableTable[id.Row][testCaseNameColumnNumber]

		clickable.onDoubleTap = func() {

			// Open TestCaseExecution
			//openTestCaseExecution(clickable.currentTestCaseExecutionUuid, clickable.testSuiteExecutionsModel)

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

			// Extract if this row is the one that is shown in TestCase preview window
			var tempRowIsSelected bool
			switch selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType {

			case AllExecutionsForOneTestCase:
				tempRowIsSelected = clickable.currentTestCaseExecutionUuid == selectedTestCaseExecutionObjected.
					allExecutionsFoOneTestCaseListObject.testCaseExecutionUuidThatIsShownInPreview

			case OneExecutionPerTestCase:
				tempRowIsSelected = clickable.currentTestCaseExecutionUuid == selectedTestCaseExecutionObjected.
					oneExecutionPerTestCaseListObject.testCaseExecutionUuidThatIsShownInPreview

			case NotDefined:

			}

			// If this row is the one that is shown in TestCase preview window
			if tempRowIsSelected {

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

		// If this Header is 'latestTestCaseExecutionTimeStampColumnNumber' then save reference to it
		//if id.Col == int(latestTestCaseExecutionTimeStampColumnNumber) {
		//	sortableHeaderReference = tempSortableHeaderLabel
		//}

		// Save a reference to the Header in the Header-map
		testCaseExecutionsListTableHeadersMapRef[id.Col] = tempSortableHeaderLabel

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

var loadTestCaseExecutionListTableTableMutex = &sync.RWMutex{}

func loadTestCaseExecutionListTableTable(
	testCaseExecutionsModelObject *testCaseExecutionsModel.TestCaseExecutionsModelStruct,
	retrieveAllExecutionsForSpecificTestCaseUuid bool,
	specificTestCaseUuid string) {

	// Lock function
	loadTestCaseExecutionListTableTableMutex.Lock()

	// Unlock function
	defer loadTestCaseExecutionListTableTableMutex.Unlock()

	testCaseExecutionsListTableTable = nil

	// Get all TestCaseExecutions form 'testCaseExecutionsThatCanBeViewedByUserMap' or from 'allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap'
	var testCaseExecutionsListMessage *[]*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage
	if selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType == OneExecutionPerTestCase {

		// Retrieve latest TestExecutions
		testCaseExecutionsListMessage = testCaseExecutionsModelObject.ReadAllFromTestCaseExecutionsMap()
	} else {

		// Retrieve all TestCaseExecutions for one TestCase

		testCaseExecutionsListMessage, _ = testCaseExecutionsModelObject.
			GetAllTestCaseExecutionsForOneTestCaseUuid(testCaseExecutionsModel.TestCaseUuidType(specificTestCaseUuid))
	}

	if testCaseExecutionsListMessage == nil {
		testCaseExecutionsListMessage = &[]*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage{}
	}

	// Loop all TestCaseExecutions and add to '[][]string'-object for the Table
	for _, tempTestCaseExecution := range *testCaseExecutionsListMessage {

		// Create temporary Row-object for the table
		var tempRowSlice []string

		// Populate the temporary Row-object

		// Column 0:
		// DomainName
		var domainNameForTable string
		domainNameForTable = fmt.Sprintf("%s [%s]",
			tempTestCaseExecution.GetDomainName(),
			tempTestCaseExecution.GetDomainUUID()[0:8])

		tempRowSlice = append(tempRowSlice, domainNameForTable)

		// Column 1:
		// SuiteName
		var suiteNameForTable string
		suiteNameForTable = fmt.Sprintf("%s [%s]",
			tempTestCaseExecution.GetTestSuiteName(),
			tempTestCaseExecution.GetTestSuiteUuid()[0:8])

		tempRowSlice = append(tempRowSlice, suiteNameForTable)

		// Column 2:
		// TestCaseName
		var testCaseNameTable string
		testCaseNameTable = fmt.Sprintf("%s [%s]",
			tempTestCaseExecution.GetTestCaseName(),
			tempTestCaseExecution.GetTestCaseUuid()[0:8])

		tempRowSlice = append(tempRowSlice, testCaseNameTable)

		// Column 3:
		// TestCaseVersion
		tempRowSlice = append(tempRowSlice, strconv.Itoa(int(tempTestCaseExecution.GetTestCaseVersion())))

		// Column 4:
		// TestCaseExecutionUuid
		tempRowSlice = append(tempRowSlice, tempTestCaseExecution.GetTestCaseExecutionUuid())

		// Column 5:
		// LatestTestCaseExecutionStatus
		var tempTestCaseExecutionStatus string

		if tempTestCaseExecution.GetTestCaseExecutionStatus() > 0 {

			tempTestCaseExecutionStatus = detailedExecutionsModel.ExecutionStatusColorMap[int32(tempTestCaseExecution.GetTestCaseExecutionStatus())].ExecutionStatusName
		} else {
			tempTestCaseExecutionStatus = "<no execution>"
		}

		tempRowSlice = append(tempRowSlice, tempTestCaseExecutionStatus)

		// Column 6:
		// TestCaseExecutionStatusStartTimeStamp
		var tempTestCaseExecutionStatusStartTimeStamp string

		if tempTestCaseExecution.GetExecutionStartTimeStamp() != nil {
			tempTestCaseExecutionStatusStartTimeStamp = sharedCode.ConvertGrpcTimeStampToStringForDB(tempTestCaseExecution.
				GetExecutionStartTimeStamp())
		} else {
			tempTestCaseExecutionStatusStartTimeStamp = "<no execution>"
		}
		tempRowSlice = append(tempRowSlice, tempTestCaseExecutionStatusStartTimeStamp)

		// Column 7:
		// TestCaseExecutionStatusUpdateTimeStamp
		var tempTestCaseExecutionFinishTimeStamp string

		if tempTestCaseExecution.GetExecutionStatusUpdateTimeStamp() != nil {
			tempTestCaseExecutionFinishTimeStamp = sharedCode.ConvertGrpcTimeStampToStringForDB(
				tempTestCaseExecution.GetExecutionStatusUpdateTimeStamp())
		} else {
			tempTestCaseExecutionFinishTimeStamp = "<no successful execution yet>"
		}
		tempRowSlice = append(tempRowSlice, tempTestCaseExecutionFinishTimeStamp)

		// Column 8:
		// TestCaseExecutionStatusStopTimeStamp
		var tempTestCaseExecutionStatusStopTimeStamp string

		if tempTestCaseExecution.GetExecutionStartTimeStamp() != nil {
			tempTestCaseExecutionStatusStopTimeStamp = sharedCode.ConvertGrpcTimeStampToStringForDB(tempTestCaseExecution.
				GetExecutionStartTimeStamp())
		} else {
			tempTestCaseExecutionStatusStopTimeStamp = "<TestCase is not not finished>"
		}
		tempRowSlice = append(tempRowSlice, tempTestCaseExecutionStatusStopTimeStamp)

		// Column 9:
		// TestCaseUuid
		tempRowSlice = append(tempRowSlice, tempTestCaseExecution.GetTestCaseUuid())

		// Column 10:
		// DomainUuid
		tempRowSlice = append(tempRowSlice, tempTestCaseExecution.GetDomainUUID())

		// Column 11:
		// TestSuiteUuid
		tempRowSlice = append(tempRowSlice, tempTestCaseExecution.GetTestSuiteUuid())

		// Column 12:
		// TestSuiteExecutionUuid
		tempRowSlice = append(tempRowSlice, tempTestCaseExecution.GetTestSuiteExecutionUuid())

		// Add Row to slice of rows for the table
		testCaseExecutionsListTableTable = append(testCaseExecutionsListTableTable, tempRowSlice)

		// Verify number columns match constant 'numberColumnsInTestCaseExecutionsListUI'
		if len(tempRowSlice) != numberColumnsInTestCaseExecutionsListUI {
			log.Fatalln(fmt.Sprintf("Number of elements in 'tempRowSlice' missmatch contant 'numberColumnsInTestCaseExecutionsListUI'. %d vs %d. ID: %s",
				tempRowSlice,
				numberColumnsInTestCaseExecutionsListUI,
				"0e3edfa7-52b8-4fcc-8243-51494463c641"))
		}

	}

	// Do an initial sort 'testCaseExecutionsListTableTable' descending on 'LastSavedTimeStamp'
	if testCaseExecutionsListMessage != nil && len(*testCaseExecutionsListMessage) > 0 {

		switch selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType {

		case AllExecutionsForOneTestCase:
			selectedTestCaseExecutionObjected.allExecutionsFoOneTestCaseListObject.
				currentSortColumn = initialColumnToSortOn

		case OneExecutionPerTestCase:
			selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.
				currentSortColumn = initialColumnToSortOn

		case NotDefined:

		}

		sort2DStringSlice(testCaseExecutionsListTableTable, initialColumnToSortOn, initialSortDirectionForInitialColumnToSortOn)

	}

}

// Sort the matrix for GUI table, update the Gui and Set correct Sort-icon for sorted Header
func sortGuiTableOnColumn(columnNumber uint8, sortDirection SortingDirectionType) {

	// Update the Gui table with the newly sorted data
	// Update the GUI
	loadTestCaseExecutionListTableTable(
		&testCaseExecutionsModel.TestCaseExecutionsModel,
		true,
		selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.
			testCaseUuidForTestCaseExecutionThatIsShownInPreview)

	// Sort matrix
	sort2DStringSlice(testCaseExecutionsListTableTable, int(columnNumber), sortDirection)

	calculateAndSetCorrectColumnWidths()
	updateTestCaseExecutionsListTable(&testCaseExecutionsModel.TestCaseExecutionsModel)

	// Set current sorted column
	switch selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType {
	case OneExecutionPerTestCase:
		selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.
			previousSortColumn = selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.currentSortColumn
		selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.currentSortColumn = int(columnNumber)
		selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.
			previousHeader = selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.currentHeader
		selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.
			currentHeader = testCaseExecutionsListTableHeadersMapRef[int(columnNumber)]

		selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.
			currentSortColumnsSortDirection = sortDirection

	case AllExecutionsForOneTestCase:

		for _, testCaseExecutionsListTableHeaderRef := range testCaseExecutionsListTableHeadersMapRef {
			testCaseExecutionsListTableHeaderRef.sortImage.unspecifiedImageContainer.Hide()
			testCaseExecutionsListTableHeaderRef.sortImage.ascendingImageContainer.Hide()
			testCaseExecutionsListTableHeaderRef.sortImage.descendingImageContainer.Hide()

			testCaseExecutionsListTableHeaderRef.sortImage.latestSelectedSortOrder = SortingDirectionUnSpecified

		}

		selectedTestCaseExecutionObjected.allExecutionsFoOneTestCaseListObject.previousSortColumn = -1
		selectedTestCaseExecutionObjected.allExecutionsFoOneTestCaseListObject.currentSortColumn = int(columnNumber)

		selectedTestCaseExecutionObjected.allExecutionsFoOneTestCaseListObject.previousHeader = nil
		selectedTestCaseExecutionObjected.allExecutionsFoOneTestCaseListObject.
			currentHeader = testCaseExecutionsListTableHeadersMapRef[int(columnNumber)]

		selectedTestCaseExecutionObjected.allExecutionsFoOneTestCaseListObject.
			currentSortColumnsSortDirection = sortDirection

	default:
		// Handle unexpected cases
		sharedCode.Logger.WithFields(logrus.Fields{
			"Id": "8fd5ae9a-9cfe-42d9-91c1-6b7c084cbc90",
			"selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType": selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType,
		}).Fatalln("Unhandled 'selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType'")

	}

	// Set correct sorting for the header
	testCaseExecutionsListTableHeadersMapRef[int(columnNumber)].sortImage.latestSelectedSortOrder = sortDirection

	// Hide all images first
	testCaseExecutionsListTableHeadersMapRef[int(columnNumber)].sortImage.unspecifiedImageContainer.Hide()
	testCaseExecutionsListTableHeadersMapRef[int(columnNumber)].sortImage.ascendingImageContainer.Hide()
	testCaseExecutionsListTableHeadersMapRef[int(columnNumber)].sortImage.descendingImageContainer.Hide()

	// Show the appropriate image
	if testCaseExecutionsListTableHeadersMapRef[int(columnNumber)].sortImage.isSortable {

		switch testCaseExecutionsListTableHeadersMapRef[int(columnNumber)].sortImage.latestSelectedSortOrder {

		case SortingDirectionUnSpecified:
			testCaseExecutionsListTableHeadersMapRef[int(columnNumber)].sortImage.unspecifiedImageContainer.Show()

		case SortingDirectionAscending:
			testCaseExecutionsListTableHeadersMapRef[int(columnNumber)].sortImage.ascendingImageContainer.Show()

		case SortingDirectionDescending:
			testCaseExecutionsListTableHeadersMapRef[int(columnNumber)].sortImage.descendingImageContainer.Show()

		default:
			// Handle unexpected cases
			sharedCode.Logger.WithFields(logrus.Fields{
				"Id":            "0eb8808b-6314-4160-b923-1e2ec5e4b4f2",
				"sortDirection": sortDirection,
			}).Fatalln("Unhandled 'sortDirection'")
		}

	} else {
		// Handle unexpected cases
		sharedCode.Logger.WithFields(logrus.Fields{
			"Id":           "7e0d2caf-c57b-45a1-8031-d71755338e94",
			"columnNumber": columnNumber,
		}).Error("Column is not sortable")

	}

	// Refresh table
	testCaseExecutionsListTable.Refresh()

}

// Sort the matrix, ascending, for GUI table, update the Gui for 'latestTestCaseExecutionTimeStampColumnNumber'
func sortGuiTableAscendingOnTestCaseExecutionTimeStamp() {

	sortGuiTableOnColumn(latestTestCaseExecutionTimeStampColumnNumber, SortingDirectionAscending)

}

// Sort the matrix, ascending, for GUI table, update the Gui for 'latestTestCaseExecutionTimeStampColumnNumber'
func SortOrReverseSortGuiTable(sortInThisColumn uint8) {

	// Which TestCaseExecutions table is shown
	switch selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType {
	case OneExecutionPerTestCase:
		if selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.currentSortColumn == int(sortInThisColumn) {

			// Switch sort order
			switch selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.currentSortColumnsSortDirection {
			case SortingDirectionAscending:
				// Set new column to be sorted ascending
				sortGuiTableOnColumn(sortInThisColumn, SortingDirectionDescending)

			case SortingDirectionDescending:
				// Set new column to be sorted ascending
				sortGuiTableOnColumn(sortInThisColumn, SortingDirectionAscending)

			case SortingDirectionUnSpecified:
				// Handle unexpected cases
				sharedCode.Logger.WithFields(logrus.Fields{
					"Id": "3fa1a246-2a0d-4bcb-ba51-9d0d2c56c0a6",
					"selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.currentSortColumnsSortDirection": selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.currentSortColumnsSortDirection,
				}).Fatalln("Unhandled 'selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.currentSortColumnsSortDirection'")

			}

		} else {
			// Set new column to be sorted ascending
			sortGuiTableOnColumn(sortInThisColumn, SortingDirectionAscending)
		}

	case AllExecutionsForOneTestCase:

		if selectedTestCaseExecutionObjected.allExecutionsFoOneTestCaseListObject.currentSortColumn == int(sortInThisColumn) {

			// Switch sort order
			switch selectedTestCaseExecutionObjected.allExecutionsFoOneTestCaseListObject.currentSortColumnsSortDirection {
			case SortingDirectionAscending:
				// Set new column to be sorted ascending
				sortGuiTableOnColumn(sortInThisColumn, SortingDirectionDescending)

			case SortingDirectionDescending:
				// Set new column to be sorted ascending
				sortGuiTableOnColumn(sortInThisColumn, SortingDirectionAscending)

			case SortingDirectionUnSpecified:
				// Handle unexpected cases
				sharedCode.Logger.WithFields(logrus.Fields{
					"Id": "a2faa1b5-3b5d-40bd-8749-93054145756e",
					"selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.currentSortColumnsSortDirection": selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.currentSortColumnsSortDirection,
				}).Fatalln("Unhandled 'selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.currentSortColumnsSortDirection'")

			}

		} else {
			// Set new column to be sorted ascending
			sortGuiTableOnColumn(sortInThisColumn, SortingDirectionAscending)
		}

	default:
		// Handle unexpected cases
		sharedCode.Logger.WithFields(logrus.Fields{
			"Id": "8fd5ae9a-9cfe-42d9-91c1-6b7c084cbc90",
			"selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType": selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType,
		}).Fatalln("Unhandled 'selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType'")
	}

}

// Sort the matrix, for GUI table, update the Gui. Use current table sorting and column, if exist
func SortGuiTableOnCurrentColumnAndSorting() {

	// Which TestCaseExecutions table is shown
	switch selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType {
	case OneExecutionPerTestCase:

		// Sort table on current column and sort order, if exist
		switch selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.currentSortColumnsSortDirection {
		case SortingDirectionAscending, SortingDirectionDescending:
			// Set on current column and use current sort direction
			sortGuiTableOnColumn(
				uint8(selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.currentSortColumn),
				selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.currentSortColumnsSortDirection)

		case SortingDirectionUnSpecified:
			// Unspecified sort order, then sort on TestCaseExecutionTimeStamp
			sortGuiTableAscendingOnTestCaseExecutionTimeStamp()
		}

	case AllExecutionsForOneTestCase:

		// Sort table on current column and sort order, if exist
		switch selectedTestCaseExecutionObjected.allExecutionsFoOneTestCaseListObject.currentSortColumnsSortDirection {
		case SortingDirectionAscending, SortingDirectionDescending:
			// Set on current column and use current sort direction
			sortGuiTableOnColumn(
				uint8(selectedTestCaseExecutionObjected.allExecutionsFoOneTestCaseListObject.currentSortColumn),
				selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.currentSortColumnsSortDirection)

		case SortingDirectionUnSpecified:
			// Unspecified sort order, then sort on TestCaseExecutionTimeStamp
			sortGuiTableAscendingOnTestCaseExecutionTimeStamp()
		}

	default:
		// Handle unexpected cases
		sharedCode.Logger.WithFields(logrus.Fields{
			"Id": "9d0e08e7-c547-42e6-ae7c-ee80601b24b5",
			"selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType": selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType,
		}).Fatalln("Unhandled 'selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType'")
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
