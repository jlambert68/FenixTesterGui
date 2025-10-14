package listTestSuiteExecutionsUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/executions/detailedExecutionsModel"
	"FenixTesterGui/testSuiteExecutions/testSuiteExecutionsModel"
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

// RemoveTestSuiteExecutionFromList
// Remove a TestSuiteExecution from the List
func RemoveTestSuiteExecutionFromList(testSuiteExecutionUuidToBeRemoved string,
	testSuiteExecutionsModelRef *testSuiteExecutionsModel.TestSuiteExecutionsModelStruct) {

	// Delete TestSuiteExecution from 'testSuiteExecutionsThatCanBeViewedByUserMap'
	testSuiteExecutionsModelRef.DeleteFromTestSuiteExecutionsMap(
		testSuiteExecutionsModel.TestSuiteExecutionUuidType(testSuiteExecutionUuidToBeRemoved))

	// Delete TestSuite from 'TestSuitesThatCanBeEditedByUserSlice'
	/*
			for index, tempTestSuitesThatCanBeEditedByUser := range testSuitesModel.TestSuitesThatCanBeEditedByUserSlice {

				// Is this the TestSuite to be removed from slice
				if tempTestSuitesThatCanBeEditedByUser.TestSuiteUuid == testSuiteExecutionUuidToBeRemoved {

					// Remove TestSuite at index
					testSuitesModel.TestSuitesThatCanBeEditedByUserSlice = append(
						testSuitesModel.TestSuitesThatCanBeEditedByUserSlice[:index],
						testSuitesModel.TestSuitesThatCanBeEditedByUserSlice[index+1:]...)

					break
				}

			}


		// Update table-list and update Table
		loadTestSuiteExecutionListTableTable(testSuiteExecutionsModelRef)
		calculateAndSetCorrectColumnWidths()
		updateTestSuiteExecutionsListTable(testSuiteExecutionsModelRef)
	*/
}

// Create the UI-list that holds the list of TestSuitesMapPtr that the user can edit
func generateTestSuiteExecutionsListTable(testSuiteExecutionsModel *testSuiteExecutionsModel.TestSuiteExecutionsModelStruct) {

	// Correctly initialize the selectedFilesTable as a new table
	testSuiteExecutionsListTable = widget.NewTable(
		func() (int, int) { return 0, numberColumnsInTestSuiteExecutionsListUI }, // Start with zero rows, 8 columns
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
				"Id":  "a65d2e66-614d-4570-bed2-976d11eca6e9",
				"err": err,
			}).Fatalln("Failed to decode 'sortUnspecifiedImageAsByteArray'")
		}
	}

	// Load the Image, ascending, for sort direction if not already done
	if sortImageAscendingAsImage == nil {
		sortImageAscendingAsImage, err = png.Decode(bytes.NewReader(sortImageAscendingAsByteArray))
		if err != nil {
			sharedCode.Logger.WithFields(logrus.Fields{
				"Id":  "5850cbe4-1aad-4079-a0c1-6c82961f870f",
				"err": err,
			}).Fatalln("Failed to decode 'sortImageAscendingAsByteArray'")
		}
	}

	// Load the Image, descending, for sort direction if not already done
	if sortImageDescendingAsImage == nil {
		sortImageDescendingAsImage, err = png.Decode(bytes.NewReader(sortImageDescendingAsByteArray))
		if err != nil {
			sharedCode.Logger.WithFields(logrus.Fields{
				"Id":  "5ef92c4a-f98e-4b47-bddf-bdc9fd076c6b",
				"err": err,
			}).Fatalln("Failed to decode 'sortImageDescendingAsByteArray'")
		}
	}

	// Define the Header
	testSuiteExecutionsListTable.ShowHeaderRow = true
	testSuiteExecutionsListTable.CreateHeader = func() fyne.CanvasObject {
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

	updateTestSuiteExecutionsListTable(testSuiteExecutionsModel)
	calculateAndSetCorrectColumnWidths()

}

var updateTestSuiteExecutionsListTableMutex = &sync.RWMutex{}

// Update the Table
func updateTestSuiteExecutionsListTable(testSuiteExecutionsModel *testSuiteExecutionsModel.TestSuiteExecutionsModelStruct) {

	// Lock function
	updateTestSuiteExecutionsListTableMutex.Lock()

	// UnLock function
	defer updateTestSuiteExecutionsListTableMutex.Unlock()

	testSuiteExecutionsListTable.Length = func() (int, int) {
		return len(testSuiteExecutionsListTableTable), numberColumnsInTestSuiteExecutionsListUI
	}
	testSuiteExecutionsListTable.CreateCell = func() fyne.CanvasObject {

		tempNewClickableLabel := newClickableTableLabel("", func() {}, false, testSuiteExecutionsModel)
		tempContainer := container.NewStack(canvas.NewRectangle(color.Transparent), tempNewClickableLabel, tempNewClickableLabel.textInsteadOfLabel)

		return tempContainer

	}
	testSuiteExecutionsListTable.UpdateCell = func(id widget.TableCellID, cell fyne.CanvasObject) {

		clickableContainer := cell.(*fyne.Container)
		clickable := clickableContainer.Objects[1].(*clickableTableLabel)
		rectangle := clickableContainer.Objects[0].(*canvas.Rectangle)
		clickable.SetText(testSuiteExecutionsListTableTable[id.Row][id.Col])
		clickable.isClickable = true
		clickable.currentRow = int16(id.Row)
		clickable.currentTestSuiteExecutionUuid = testSuiteExecutionsListTableTable[id.Row][testSuiteExecutionUuidColumnNumber]
		clickable.currentTestSuiteExecutionVersion = 1 //TODO HArdcoded  for now testSuiteExecutionsListTableTable[id.Row][testSuiteExecutionVersionColumnNumber]
		clickable.currentTestSuiteUuid = testSuiteExecutionsListTableTable[id.Row][testSuiteUuidColumnNumber]
		clickable.currentTestSuiteName = testSuiteExecutionsListTableTable[id.Row][testSuiteNameColumnNumber]

		clickable.onDoubleTap = func() {

			// Open TestSuiteExecution
			//openTestSuiteExecution(clickable.currentTestSuiteExecutionUuid, clickable.testSuiteExecutionsModel)

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

			case latestTestSuiteExecutionStatus:

				clickable.textInsteadOfLabel.Text = clickable.Text
				clickable.Text = ""
				clickable.textInsteadOfLabel.Show()
				clickable.Hide()

			default:
				clickable.textInsteadOfLabel.Hide()
				clickable.Show()

			}

		} else {

			// Extract if this row is the one that is shown in TestSuite preview window
			var tempRowIsSelected bool
			switch selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType {

			case AllExecutionsForOneTestSuite:
				tempRowIsSelected = clickable.currentTestSuiteExecutionUuid == selectedTestSuiteExecutionObjected.
					allExecutionsFoOneTestSuiteListObject.testSuiteExecutionUuidThatIsShownInPreview

			case OneExecutionPerTestSuite:
				tempRowIsSelected = clickable.currentTestSuiteExecutionUuid == selectedTestSuiteExecutionObjected.
					oneExecutionPerTestSuiteListObject.testSuiteExecutionUuidThatIsShownInPreview

			case NotDefined:

			}

			// If this row is the one that is shown in TestSuite preview window
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

			case latestTestSuiteExecutionStatus:
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
	testSuiteExecutionsListTable.UpdateHeader = func(id widget.TableCellID, cell fyne.CanvasObject) {
		tempSortableHeaderLabel := cell.(*sortableHeaderLabelStruct)
		tempSortableHeaderLabel.label.SetText(testSuiteExecutionsListTableHeader[id.Col])
		tempSortableHeaderLabel.label.TextStyle = fyne.TextStyle{Bold: true}

		// Set Column number
		tempSortableHeaderLabel.columnNumber = id.Col
		tempSortableHeaderLabel.sortImage.headerColumnNumber = id.Col

		// If this Header is 'latestTestSuiteExecutionTimeStampColumnNumber' then save reference to it
		//if id.Col == int(latestTestSuiteExecutionTimeStampColumnNumber) {
		//	sortableHeaderReference = tempSortableHeaderLabel
		//}

		// Save a reference to the Header in the Header-map
		testSuiteExecutionsListTableHeadersMapRef[id.Col] = tempSortableHeaderLabel

		//tempSortableHeaderLabel.latestSelectedSortOrder = SortingDirectionAscending
		//tempSortableHeaderLabel.updateColumnNumberFunction()

		// Refresh the widget to update the UI
		tempSortableHeaderLabel.Refresh()
	}

	testSuiteExecutionsListTable.Refresh()
}

// TestSuiteUuid
// TestSuiteVersion
// LatestTestSuiteExecutionStatus
// LatestTestSuiteExecutionStatusInsertTimeStamp
// LatestFinishedOkTestSuiteExecutionStatusInsertTimeStamp
// DomainUuid

func calculateAndSetCorrectColumnWidths() {

	// Initiate slice for keeping track of max column width size
	var columnsMaxSizeSlice []float32
	columnsMaxSizeSlice = make([]float32, numberColumnsInTestSuiteExecutionsListUI)

	var columnWidth float32

	// Set initial value for max width size
	for index, headerValue := range testSuiteExecutionsListTableHeader {

		// Calculate the column width base on this value. Add  'float32(30)' to give room for sort direction icon
		columnWidth = fyne.MeasureText(headerValue, theme.TextSize(), fyne.TextStyle{Bold: true}).Width + float32(30) //
		columnsMaxSizeSlice[index] = columnWidth
	}

	// Loop all rows
	for _, tempRow := range testSuiteExecutionsListTableTable {

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
		testSuiteExecutionsListTable.SetColumnWidth(columnIndex, columnWidth+theme.Padding()*4)
	}

	// Refresh the table
	testSuiteExecutionsListTable.Refresh()

}

var loadTestSuiteExecutionListTableTableMutex = &sync.RWMutex{}

func loadTestSuiteExecutionListTableTable(
	testSuiteExecutionsModelObject *testSuiteExecutionsModel.TestSuiteExecutionsModelStruct,
	retrieveAllExecutionsForSpecificTestSuiteUuid bool,
	specificTestSuiteUuid string) {

	// Lock function
	loadTestSuiteExecutionListTableTableMutex.Lock()

	// Unlock function
	defer loadTestSuiteExecutionListTableTableMutex.Unlock()

	testSuiteExecutionsListTableTable = nil

	// Get all TestSuiteExecutions form 'testSuiteExecutionsThatCanBeViewedByUserMap' or from 'allTestSuiteExecutionsForAllTestSuitesThatCanBeViewedByUserMap'
	var testSuiteExecutionsListMessage *[]*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage
	if selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType == OneExecutionPerTestSuite {

		// Retrieve latest TestExecutions
		testSuiteExecutionsListMessage = testSuiteExecutionsModelObject.ReadAllFromTestSuiteExecutionsMap()
	} else {

		// Retrieve all TestSuiteExecutions for one TestSuite

		testSuiteExecutionsListMessage, _ = testSuiteExecutionsModelObject.
			GetAllTestSuiteExecutionsForOneTestSuiteUuid(testSuiteExecutionsModel.TestSuiteUuidType(specificTestSuiteUuid))
	}

	if testSuiteExecutionsListMessage == nil {
		testSuiteExecutionsListMessage = &[]*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage{}
	}

	// Loop all TestSuiteExecutions and add to '[][]string'-object for the Table
	for _, tempTestSuiteExecution := range *testSuiteExecutionsListMessage {

		// Create temporary Row-object for the table
		var tempRowSlice []string

		// Populate the temporary Row-object

		// Column 0:
		// DomainName
		var domainNameForTable string
		domainNameForTable = fmt.Sprintf("%s [%s]",
			tempTestSuiteExecution.GetDomainName(),
			tempTestSuiteExecution.GetDomainUUID()[0:8])

		tempRowSlice = append(tempRowSlice, domainNameForTable)

		// Column 1:
		// SuiteName
		var suiteNameForTable string
		suiteNameForTable = fmt.Sprintf("%s [%s]",
			tempTestSuiteExecution.GetTestSuiteName(),
			tempTestSuiteExecution.GetTestSuiteUuid()[0:8])

		tempRowSlice = append(tempRowSlice, suiteNameForTable)

		// Column 2:
		// TestSuiteName
		var testSuiteNameTable string
		testSuiteNameTable = fmt.Sprintf("%s [%s]",
			tempTestSuiteExecution.GetTestSuiteName(),
			tempTestSuiteExecution.GetTestSuiteUuid()[0:8])

		tempRowSlice = append(tempRowSlice, testSuiteNameTable)

		// Column 3:
		// TestSuiteVersion
		tempRowSlice = append(tempRowSlice, strconv.Itoa(int(tempTestSuiteExecution.GetTestSuiteVersion())))

		// Column 4:
		// TestSuiteExecutionUuid
		tempRowSlice = append(tempRowSlice, tempTestSuiteExecution.GetTestSuiteExecutionUuid())

		// Column 5:
		// LatestTestSuiteExecutionStatus
		var tempTestSuiteExecutionStatus string

		if tempTestSuiteExecution.GetTestSuiteExecutionStatus() > 0 {

			tempTestSuiteExecutionStatus = detailedExecutionsModel.ExecutionStatusColorMap[int32(tempTestSuiteExecution.GetTestSuiteExecutionStatus())].ExecutionStatusName
		} else {
			tempTestSuiteExecutionStatus = "<no execution>"
		}

		tempRowSlice = append(tempRowSlice, tempTestSuiteExecutionStatus)

		// Column 6:
		// TestSuiteExecutionStatusStartTimeStamp
		var tempTestSuiteExecutionStatusStartTimeStamp string

		if tempTestSuiteExecution.GetExecutionStartTimeStamp() != nil {
			tempTestSuiteExecutionStatusStartTimeStamp = sharedCode.ConvertGrpcTimeStampToStringForDB(tempTestSuiteExecution.
				GetExecutionStartTimeStamp())
		} else {
			tempTestSuiteExecutionStatusStartTimeStamp = "<no execution>"
		}
		tempRowSlice = append(tempRowSlice, tempTestSuiteExecutionStatusStartTimeStamp)

		// Column 7:
		// TestSuiteExecutionStatusUpdateTimeStamp
		var tempTestSuiteExecutionFinishTimeStamp string

		if tempTestSuiteExecution.GetExecutionStatusUpdateTimeStamp() != nil {
			tempTestSuiteExecutionFinishTimeStamp = sharedCode.ConvertGrpcTimeStampToStringForDB(
				tempTestSuiteExecution.GetExecutionStatusUpdateTimeStamp())
		} else {
			tempTestSuiteExecutionFinishTimeStamp = "<no successful execution yet>"
		}
		tempRowSlice = append(tempRowSlice, tempTestSuiteExecutionFinishTimeStamp)

		// Column 8:
		// TestSuiteExecutionStatusStopTimeStamp
		var tempTestSuiteExecutionStatusStopTimeStamp string

		if tempTestSuiteExecution.GetExecutionStartTimeStamp() != nil {
			tempTestSuiteExecutionStatusStopTimeStamp = sharedCode.ConvertGrpcTimeStampToStringForDB(tempTestSuiteExecution.
				GetExecutionStartTimeStamp())
		} else {
			tempTestSuiteExecutionStatusStopTimeStamp = "<TestSuite is not not finished>"
		}
		tempRowSlice = append(tempRowSlice, tempTestSuiteExecutionStatusStopTimeStamp)

		// Column 9:
		// TestSuiteUuid
		tempRowSlice = append(tempRowSlice, tempTestSuiteExecution.GetTestSuiteUuid())

		// Column 10:
		// DomainUuid
		tempRowSlice = append(tempRowSlice, tempTestSuiteExecution.GetDomainUUID())

		// Column 11:
		// TestSuiteUuid
		tempRowSlice = append(tempRowSlice, tempTestSuiteExecution.GetTestSuiteUuid())

		// Column 12:
		// TestSuiteExecutionUuid
		tempRowSlice = append(tempRowSlice, tempTestSuiteExecution.GetTestSuiteExecutionUuid())

		// Add Row to slice of rows for the table
		testSuiteExecutionsListTableTable = append(testSuiteExecutionsListTableTable, tempRowSlice)

		// Verify number columns match constant 'numberColumnsInTestSuiteExecutionsListUI'
		if len(tempRowSlice) != numberColumnsInTestSuiteExecutionsListUI {
			log.Fatalln(fmt.Sprintf("Number of elements in 'tempRowSlice' missmatch contant 'numberColumnsInTestSuiteExecutionsListUI'. %d vs %d. ID: %s",
				tempRowSlice,
				numberColumnsInTestSuiteExecutionsListUI,
				"0e3edfa7-52b8-4fcc-8243-51494463c641"))
		}

	}

	// Do an initial sort 'testSuiteExecutionsListTableTable' descending on 'LastSavedTimeStamp'
	if testSuiteExecutionsListMessage != nil && len(*testSuiteExecutionsListMessage) > 0 {

		switch selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType {

		case AllExecutionsForOneTestSuite:
			selectedTestSuiteExecutionObjected.allExecutionsFoOneTestSuiteListObject.
				currentSortColumn = initialColumnToSortOn

		case OneExecutionPerTestSuite:
			selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.
				currentSortColumn = initialColumnToSortOn

		case NotDefined:

		}

		sort2DStringSlice(testSuiteExecutionsListTableTable, initialColumnToSortOn, initialSortDirectionForInitialColumnToSortOn)

	}

}

// Sort the matrix for GUI table, update the Gui and Set correct Sort-icon for sorted Header
func sortGuiTableOnColumn(columnNumber uint8, sortDirection SortingDirectionType) {

	// Update the Gui table with the newly sorted data
	// Update the GUI
	loadTestSuiteExecutionListTableTable(
		&testSuiteExecutionsModel.TestSuiteExecutionsModel,
		true,
		selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.
			testSuiteUuidForTestSuiteExecutionThatIsShownInPreview)

	// Sort matrix
	sort2DStringSlice(testSuiteExecutionsListTableTable, int(columnNumber), sortDirection)

	calculateAndSetCorrectColumnWidths()
	updateTestSuiteExecutionsListTable(&testSuiteExecutionsModel.TestSuiteExecutionsModel)

	// Set current sorted column
	switch selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType {
	case OneExecutionPerTestSuite:
		selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.
			previousSortColumn = selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.currentSortColumn
		selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.currentSortColumn = int(columnNumber)
		selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.
			previousHeader = selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.currentHeader
		selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.
			currentHeader = testSuiteExecutionsListTableHeadersMapRef[int(columnNumber)]

		selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.
			currentSortColumnsSortDirection = sortDirection

	case AllExecutionsForOneTestSuite:

		for _, testSuiteExecutionsListTableHeaderRef := range testSuiteExecutionsListTableHeadersMapRef {
			testSuiteExecutionsListTableHeaderRef.sortImage.unspecifiedImageContainer.Hide()
			testSuiteExecutionsListTableHeaderRef.sortImage.ascendingImageContainer.Hide()
			testSuiteExecutionsListTableHeaderRef.sortImage.descendingImageContainer.Hide()

			testSuiteExecutionsListTableHeaderRef.sortImage.latestSelectedSortOrder = SortingDirectionUnSpecified

		}

		selectedTestSuiteExecutionObjected.allExecutionsFoOneTestSuiteListObject.previousSortColumn = -1
		selectedTestSuiteExecutionObjected.allExecutionsFoOneTestSuiteListObject.currentSortColumn = int(columnNumber)

		selectedTestSuiteExecutionObjected.allExecutionsFoOneTestSuiteListObject.previousHeader = nil
		selectedTestSuiteExecutionObjected.allExecutionsFoOneTestSuiteListObject.
			currentHeader = testSuiteExecutionsListTableHeadersMapRef[int(columnNumber)]

		selectedTestSuiteExecutionObjected.allExecutionsFoOneTestSuiteListObject.
			currentSortColumnsSortDirection = sortDirection

	default:
		// Handle unexpected cases
		sharedCode.Logger.WithFields(logrus.Fields{
			"Id": "a78d2b95-f8c2-41fb-89d2-873815c8745f",
			"selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType": selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType,
		}).Fatalln("Unhandled 'selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType'")

	}

	// Set correct sorting for the header
	testSuiteExecutionsListTableHeadersMapRef[int(columnNumber)].sortImage.latestSelectedSortOrder = sortDirection

	// Hide all images first
	testSuiteExecutionsListTableHeadersMapRef[int(columnNumber)].sortImage.unspecifiedImageContainer.Hide()
	testSuiteExecutionsListTableHeadersMapRef[int(columnNumber)].sortImage.ascendingImageContainer.Hide()
	testSuiteExecutionsListTableHeadersMapRef[int(columnNumber)].sortImage.descendingImageContainer.Hide()

	// Show the appropriate image
	if testSuiteExecutionsListTableHeadersMapRef[int(columnNumber)].sortImage.isSortable {

		switch testSuiteExecutionsListTableHeadersMapRef[int(columnNumber)].sortImage.latestSelectedSortOrder {

		case SortingDirectionUnSpecified:
			testSuiteExecutionsListTableHeadersMapRef[int(columnNumber)].sortImage.unspecifiedImageContainer.Show()

		case SortingDirectionAscending:
			testSuiteExecutionsListTableHeadersMapRef[int(columnNumber)].sortImage.ascendingImageContainer.Show()

		case SortingDirectionDescending:
			testSuiteExecutionsListTableHeadersMapRef[int(columnNumber)].sortImage.descendingImageContainer.Show()

		default:
			// Handle unexpected cases
			sharedCode.Logger.WithFields(logrus.Fields{
				"Id":            "72a579d0-4886-4865-a251-efb7d9eefa93",
				"sortDirection": sortDirection,
			}).Fatalln("Unhandled 'sortDirection'")
		}

	} else {
		// Handle unexpected cases
		sharedCode.Logger.WithFields(logrus.Fields{
			"Id":           "6b26dfa6-b188-4ca7-a489-b91fe6ff68ee",
			"columnNumber": columnNumber,
		}).Error("Column is not sortable")

	}

	// Refresh table
	testSuiteExecutionsListTable.Refresh()

}

// Sort the matrix, ascending, for GUI table, update the Gui for 'latestTestSuiteExecutionTimeStampColumnNumber'
func sortGuiTableAscendingOnTestSuiteExecutionTimeStamp() {

	sortGuiTableOnColumn(latestTestSuiteExecutionTimeStampColumnNumber, SortingDirectionAscending)

}

// Sort the matrix, ascending, for GUI table, update the Gui for 'latestTestSuiteExecutionTimeStampColumnNumber'
func SortOrReverseSortGuiTable(sortInThisColumn uint8) {

	// Which TestSuiteExecutions table is shown
	switch selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType {
	case OneExecutionPerTestSuite:
		if selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.currentSortColumn == int(sortInThisColumn) {

			// Switch sort order
			switch selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.currentSortColumnsSortDirection {
			case SortingDirectionAscending:
				// Set new column to be sorted ascending
				sortGuiTableOnColumn(sortInThisColumn, SortingDirectionDescending)

			case SortingDirectionDescending:
				// Set new column to be sorted ascending
				sortGuiTableOnColumn(sortInThisColumn, SortingDirectionAscending)

			case SortingDirectionUnSpecified:
				// Handle unexpected cases
				sharedCode.Logger.WithFields(logrus.Fields{
					"Id": "ec1069fd-f8eb-4bae-a3e3-91b536a558e3",
					"selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.currentSortColumnsSortDirection": selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.currentSortColumnsSortDirection,
				}).Fatalln("Unhandled 'selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.currentSortColumnsSortDirection'")

			}

		} else {
			// Set new column to be sorted ascending
			sortGuiTableOnColumn(sortInThisColumn, SortingDirectionAscending)
		}

	case AllExecutionsForOneTestSuite:

		if selectedTestSuiteExecutionObjected.allExecutionsFoOneTestSuiteListObject.currentSortColumn == int(sortInThisColumn) {

			// Switch sort order
			switch selectedTestSuiteExecutionObjected.allExecutionsFoOneTestSuiteListObject.currentSortColumnsSortDirection {
			case SortingDirectionAscending:
				// Set new column to be sorted ascending
				sortGuiTableOnColumn(sortInThisColumn, SortingDirectionDescending)

			case SortingDirectionDescending:
				// Set new column to be sorted ascending
				sortGuiTableOnColumn(sortInThisColumn, SortingDirectionAscending)

			case SortingDirectionUnSpecified:
				// Handle unexpected cases
				sharedCode.Logger.WithFields(logrus.Fields{
					"Id": "169f8fef-942b-45e3-b9a8-1af88d74c1ed",
					"selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.currentSortColumnsSortDirection": selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.currentSortColumnsSortDirection,
				}).Fatalln("Unhandled 'selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.currentSortColumnsSortDirection'")

			}

		} else {
			// Set new column to be sorted ascending
			sortGuiTableOnColumn(sortInThisColumn, SortingDirectionAscending)
		}

	default:
		// Handle unexpected cases
		sharedCode.Logger.WithFields(logrus.Fields{
			"Id": "4d165016-4494-47a6-9525-681bee77ee83",
			"selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType": selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType,
		}).Fatalln("Unhandled 'selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType'")
	}

}

// Sort the matrix, for GUI table, update the Gui. Use current table sorting and column, if exist
func SortGuiTableOnCurrentColumnAndSorting() {

	// Which TestSuiteExecutions table is shown
	switch selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType {
	case OneExecutionPerTestSuite:

		// Sort table on current column and sort order, if exist
		switch selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.currentSortColumnsSortDirection {
		case SortingDirectionAscending, SortingDirectionDescending:
			// Set on current column and use current sort direction
			sortGuiTableOnColumn(
				uint8(selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.currentSortColumn),
				selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.currentSortColumnsSortDirection)

		case SortingDirectionUnSpecified:
			// Unspecified sort order, then sort on TestSuiteExecutionTimeStamp
			sortGuiTableAscendingOnTestSuiteExecutionTimeStamp()
		}

	case AllExecutionsForOneTestSuite:

		// Sort table on current column and sort order, if exist
		switch selectedTestSuiteExecutionObjected.allExecutionsFoOneTestSuiteListObject.currentSortColumnsSortDirection {
		case SortingDirectionAscending, SortingDirectionDescending:
			// Set on current column and use current sort direction
			sortGuiTableOnColumn(
				uint8(selectedTestSuiteExecutionObjected.allExecutionsFoOneTestSuiteListObject.currentSortColumn),
				selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.currentSortColumnsSortDirection)

		case SortingDirectionUnSpecified:
			// Unspecified sort order, then sort on TestSuiteExecutionTimeStamp
			sortGuiTableAscendingOnTestSuiteExecutionTimeStamp()
		}

	default:
		// Handle unexpected cases
		sharedCode.Logger.WithFields(logrus.Fields{
			"Id": "4a450af6-6a7c-49b0-b50e-1923c406978c",
			"selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType": selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType,
		}).Fatalln("Unhandled 'selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType'")
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
