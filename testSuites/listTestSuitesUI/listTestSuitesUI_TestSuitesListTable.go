package listTestSuitesUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/executions/detailedExecutionsModel"
	"FenixTesterGui/soundEngine"
	"FenixTesterGui/testCase/testCaseModel"
	"FenixTesterGui/testSuites/listTestSuitesModel"
	"bytes"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jlambert68/Fast_BitFilter_MetaData/boolbits/boolbits"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"image/color"
	"image/png"
	"log"
	"sort"
	"strconv"
)

// RemoveTestSuiteFromList
// Remove a TestSuite from the List
func (listTestSuiteUIObject *ListTestSuiteUIStruct) RemoveTestSuiteFromList(testSuiteUuidToBeRemoved string, testCasesModel *testCaseModel.TestCasesModelsStruct) {

	// Delete TestSuite from 'TestSuitesThatCanBeEditedByUserMap'
	delete(listTestSuitesModel.TestSuitesThatCanBeEditedByUserMap, testSuiteUuidToBeRemoved)

	// Update table-list and update Table
	listTestSuiteUIObject.loadTestSuiteListTableTable(nil)
	listTestSuiteUIObject.calculateAndSetCorrectColumnWidths()
	listTestSuiteUIObject.updateTestSuitesListTable(testCasesModel)

}

// Create the UI-list that holds the list of TestSuitesMapPtr that the user can edit
func (listTestSuiteUIObject *ListTestSuiteUIStruct) generateTestSuitesListTable(testCasesModel *testCaseModel.TestCasesModelsStruct) {

	// Correctly initialize the selectedFilesTable as a new table
	listTestSuiteUIObject.testSuiteListTable = widget.NewTable(
		func() (int, int) {

			if listTestSuiteUIObject.howShouldItBeUsed == UsedForTestSuitesList {
				return 0, numberColumnsInTestSuitesListUIForTestSuitesList
			} else {
				return 0, numberColumnsInTestSuitesListUIForTestSuiteBuilder
			}

		}, // Start with zero rows, 8 columns
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
	listTestSuiteUIObject.testSuiteListTable.ShowHeaderRow = true
	listTestSuiteUIObject.testSuiteListTable.CreateHeader = func() fyne.CanvasObject {
		//return widget.NewLabel("") // Create cells as labels

		var tempNewSortableHeaderLabel *sortableHeaderLabelStruct
		tempNewSortableHeaderLabel = newSortableHeaderLabel("", true, 0, listTestSuiteUIObject)

		// Create the Sort Icons container
		//var newSortIconsContainer *fyne.Container
		//newSortIconsContainer = container.NewStack(tempNewSortableHeaderLabel.sortImage.unspecifiedImageContainer, tempNewSortableHeaderLabel.sortImage)

		//var newSortableHeaderLabelContainer *fyne.Container
		//newSortableHeaderLabelContainer = container.NewHBox(
		//	widget.NewLabel(tempNewSortableHeaderLabel.Text), newSortIconsContainer) //canvas.NewImageFromImage(sortImageUnspecifiedAsImage)) //tempNewSortableHeaderLabel.sortImage)

		return tempNewSortableHeaderLabel

	}

	listTestSuiteUIObject.updateTestSuitesListTable(testCasesModel)
	listTestSuiteUIObject.calculateAndSetCorrectColumnWidths()

}

// Update the Table
func (listTestSuiteUIObject *ListTestSuiteUIStruct) updateTestSuitesListTable(testCasesModel *testCaseModel.TestCasesModelsStruct) {

	listTestSuiteUIObject.testSuiteListTable.Length = func() (int, int) {

		if listTestSuiteUIObject.howShouldItBeUsed == UsedForTestSuitesList {
			return len(listTestSuiteUIObject.testSuitesListTableTable), numberColumnsInTestSuitesListUIForTestSuitesList
		} else {
			return len(listTestSuiteUIObject.testSuitesListTableTable), numberColumnsInTestSuitesListUIForTestSuiteBuilder
		}
	}
	listTestSuiteUIObject.testSuiteListTable.CreateCell = func() fyne.CanvasObject {

		tempNewClickableLabel := newClickableTableLabel("", func() {}, false, testCasesModel, listTestSuiteUIObject)
		tempContainer := container.NewStack(canvas.NewRectangle(color.Transparent), tempNewClickableLabel, tempNewClickableLabel.textInsteadOfLabel)

		return tempContainer

	}
	listTestSuiteUIObject.testSuiteListTable.UpdateCell = func(id widget.TableCellID, cell fyne.CanvasObject) {

		clickableContainer := cell.(*fyne.Container)
		clickable := clickableContainer.Objects[1].(*clickableTableLabel)
		rectangle := clickableContainer.Objects[0].(*canvas.Rectangle)
		clickable.SetText(listTestSuiteUIObject.testSuitesListTableTable[id.Row][id.Col])
		clickable.isClickable = true
		clickable.currentRow = int16(id.Row)
		clickable.currentTestSuiteUuid = listTestSuiteUIObject.testSuitesListTableTable[id.Row][testSuiteUuidColumnNumber+uint8(listTestSuiteUIObject.howShouldItBeUsed)]

		clickable.onDoubleTap = func() {

			if listTestSuiteUIObject.howShouldItBeUsed == UsedForTestSuitesList {
				// Open TestSuite when we are in TestSuiteListing
				listTestSuiteUIObject.openTestSuite(clickable.currentTestSuiteUuid, clickable.testCasesModel)

			} else {
				// Select or UnSelect TestSuite in TestSuiteList
				var selectedTestSuites map[string]*fenixGuiTestCaseBuilderServerGrpcApi.TestCaseInTestSuiteMessage
				//var selectedTestCase *fenixGuiTestCaseBuilderServerGrpcApi.TestCaseInTestSuiteMessage
				var existInMap bool
				selectedTestSuites = *clickable.listTestSuiteUIPtr.selectedTestCasesPtr

				// Check if TestCase is Selected
				_, existInMap = selectedTestSuites[clickable.currentTestSuiteUuid]
				if existInMap == true {
					// UnSelect TestCase
					delete(selectedTestSuites, clickable.currentTestSuiteUuid)

					// Trigger System Notification sound
					soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

					fyne.CurrentApp().SendNotification(&fyne.Notification{
						Title:   "TestCase UnSelected",
						Content: fmt.Sprintf("'%s' was unselected", clickable.currentTestSuiteUuid),
					})

				} else {
					// Select TestSuite
					selectedTestSuites[clickable.currentTestSuiteUuid] = &fenixGuiTestCaseBuilderServerGrpcApi.TestCaseInTestSuiteMessage{
						DomainUuid:   "",
						DomainName:   "",
						TestCaseUuid: clickable.currentTestSuiteUuid,
						TestCaseName: "",
					}

					// Trigger System Notification sound
					soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

					fyne.CurrentApp().SendNotification(&fyne.Notification{
						Title:   "TestSuite Selected",
						Content: fmt.Sprintf("'%s' was selected", clickable.currentTestSuiteUuid),
					})
				}

			}
		}

		// Check if this row should be highlighted or not
		if int16(id.Row) == listTestSuiteUIObject.currentRowThatMouseIsHoveringAbove {
			clickable.TextStyle = fyne.TextStyle{Bold: false}
			rectangle.FillColor = color.RGBA{
				R: 0x4A,
				G: 0x4B,
				B: 0x4D,
				A: 0xFF,
			}

			// Special handling for certain Columns for Status color and Timestamps
			switch uint8(id.Col) {

			case 0:
				// Only check this if TestSuiteBuilder is in play and Column '0'
				if listTestSuiteUIObject.howShouldItBeUsed == UsedForTestSuiteBuilder && id.Col == 0 {
					rectangle.FillColor = color.RGBA{
						R: 0x00,
						G: 0xFF,
						B: 0x00,
						A: 0x00,
					}

					// Check if TestSuite is Selected or UnSelected in TestSuite
					var selectedTestSuites map[string]*fenixGuiTestCaseBuilderServerGrpcApi.TestCaseInTestSuiteMessage
					//var selectedTestCase *fenixGuiTestCaseBuilderServerGrpcApi.TestCaseInTestSuiteMessage
					var existInMap bool
					selectedTestSuites = *clickable.listTestSuiteUIPtr.selectedTestCasesPtr

					// Check if TestCase is Selected
					_, existInMap = selectedTestSuites[clickable.currentTestSuiteUuid]
					if existInMap == true {
						// Is Selected
						clickable.Text = "SELECTED"
						listTestSuiteUIObject.testSuitesListTableTable[id.Row][0] = "SELECTED"

					} else {
						// Is UnSelected
						clickable.Text = "-"
						listTestSuiteUIObject.testSuitesListTableTable[id.Row][0] = "-"
					}

					//listTestSuiteUIObject.testSuiteListTable.Refresh()

					break
				}

			case latestTestSuiteExecutionStatusColumnNumber:
				if listTestSuiteUIObject.howShouldItBeUsed == UsedForTestSuitesList {
					clickable.textInsteadOfLabel.Text = clickable.Text
					clickable.Text = ""
					clickable.textInsteadOfLabel.Show()
					clickable.Hide()

					break
				}

			default:
				clickable.textInsteadOfLabel.Hide()
				clickable.Show()

			}

		} else {

			// If this row is the one that is shown in TestSuite preview window
			if clickable.currentTestSuiteUuid == listTestSuiteUIObject.testSuiteThatIsShownInPreview {

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

			case 0:
				// Only check this if TestSuiteBuilder is in play and Column '0'
				if listTestSuiteUIObject.howShouldItBeUsed == UsedForTestSuiteBuilder && id.Col == 0 {
					rectangle.FillColor = color.RGBA{
						R: 0x00,
						G: 0xFF,
						B: 0x00,
						A: 0x00,
					}

					// Check if TestSuite is Selected or UnSelected in TestSuite
					var selectedTestCases map[string]*fenixGuiTestCaseBuilderServerGrpcApi.TestCaseInTestSuiteMessage
					//var selectedTestCase *fenixGuiTestCaseBuilderServerGrpcApi.TestCaseInTestSuiteMessage
					var existInMap bool
					selectedTestCases = *clickable.listTestSuiteUIPtr.selectedTestCasesPtr

					// Check if TestSuite is Selected
					_, existInMap = selectedTestCases[clickable.currentTestSuiteUuid]
					if existInMap == true {
						// Is Selected
						clickable.Text = "SELECTED"
						listTestSuiteUIObject.testSuitesListTableTable[id.Row][0] = "SELECTED"

					} else {
						// Is UnSelected
						clickable.Text = "-"
						listTestSuiteUIObject.testSuitesListTableTable[id.Row][0] = "-"
					}

					//listTestSuiteUIObject.testSuiteListTable.Refresh()

					break
				}

			case latestTestSuiteExecutionStatusColumnNumber + uint8(listTestSuiteUIObject.howShouldItBeUsed):
				if listTestSuiteUIObject.howShouldItBeUsed == UsedForTestSuitesList {
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

					break
				}

			default:

			}

		}

		clickableContainer.Refresh()

	}

	// Update the Header
	listTestSuiteUIObject.testSuiteListTable.UpdateHeader = func(id widget.TableCellID, cell fyne.CanvasObject) {
		tempSortableHeaderLabel := cell.(*sortableHeaderLabelStruct)

		tempSortableHeaderLabel.label.SetText(listTestSuiteUIObject.testSuiteListTableHeader[id.Col])
		tempSortableHeaderLabel.label.TextStyle = fyne.TextStyle{Bold: true}

		// Set Column number
		tempSortableHeaderLabel.columnNumber = id.Col
		tempSortableHeaderLabel.sortImage.headerColumnNumber = id.Col

		// Save reference to sortableHeaderReference depending on TestSuitesList or TestSuiteBuilder

		if listTestSuiteUIObject.howShouldItBeUsed == UsedForTestSuitesList {
			// If this Header is 'latestTestSuiteExecutionTimeStampColumnNumber' then save reference to it
			if id.Col == int(latestTestSuiteExecutionTimeStampColumnNumber) {
				listTestSuiteUIObject.sortableHeaderReference = tempSortableHeaderLabel
			}
		} else {
			// If this Header is 'latestTestSuiteExecutionTimeStampColumnNumber' then save reference to it
			if id.Col == int(initialColumnToSortOnForTestSuiteBuilder) {
				listTestSuiteUIObject.sortableHeaderReference = tempSortableHeaderLabel
			}
		}

		//tempSortableHeaderLabel.latestSelectedSortOrder = SortingDirectionAscending
		//tempSortableHeaderLabel.updateColumnNumberFunction()

		// Refresh the widget to update the UI
		tempSortableHeaderLabel.Refresh()
	}

	listTestSuiteUIObject.testSuiteListTable.Refresh()
}

// TestSuiteUuid
// TestSuiteVersion
// LatestTestSuiteExecutionStatus
// LatestTestSuiteExecutionStatusInsertTimeStamp
// LatestFinishedOkTestSuiteExecutionStatusInsertTimeStamp
// DomainUuid

func (listTestSuiteUIObject *ListTestSuiteUIStruct) calculateAndSetCorrectColumnWidths() {

	// Initiate slice for keeping track of max column width size
	var columnsMaxSizeSlice []float32
	if listTestSuiteUIObject.howShouldItBeUsed == UsedForTestSuitesList {
		columnsMaxSizeSlice = make([]float32, numberColumnsInTestSuitesListUIForTestSuitesList)
	} else {
		columnsMaxSizeSlice = make([]float32, numberColumnsInTestSuitesListUIForTestSuiteBuilder)
	}

	var columnWidth float32

	// Set initial value for max width size
	for index, headerValue := range listTestSuiteUIObject.testSuiteListTableHeader {

		// Calculate the column width base on this value. Add  'float32(30)' to give room for sort direction icon
		columnWidth = fyne.MeasureText(headerValue, theme.TextSize(), fyne.TextStyle{Bold: true}).Width + float32(30) //
		columnsMaxSizeSlice[index] = columnWidth
	}

	// Loop all rows
	for _, tempRow := range listTestSuiteUIObject.testSuitesListTableTable {

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
		listTestSuiteUIObject.testSuiteListTable.SetColumnWidth(columnIndex, columnWidth+theme.Padding()*4)
	}

	// Refresh the table
	listTestSuiteUIObject.testSuiteListTable.Refresh()

}

func (listTestSuiteUIObject *ListTestSuiteUIStruct) loadTestSuiteListTableTable(
	testSuiteMetaDataFilterEntry *boolbits.Entry) {

	listTestSuiteUIObject.testSuitesListTableTable = nil
	var testSuiteUuid string
	var existInMap bool
	var err error

	// Loop all TestSuitesMapPtr and add to '[][]string'-object for the Table
	for _, tempTestSuite := range listTestSuitesModel.TestSuitesThatCanBeEditedByUserMap {

		// Check and apply MetaData-filter
		if testSuiteMetaDataFilterEntry != nil {

			// Generate Initial Entry for the boolean arithmetic
			var resultEntry *boolbits.Entry
			resultEntry, err = boolbits.NewAllZerosEntry(64)
			if err != nil {
				errorID := "cc06778a-54e0-4281-9a67-088a12bf9107"
				errorMessage := fmt.Sprintf("could not create initial Entry for boolean arithmetic [ErrorID=%s, err='%s']",
					errorID,
					err.Error())

				log.Fatalln(errorMessage)
			}

			// Copy TestSuite-filter
			resultEntry, err = resultEntry.Or(testSuiteMetaDataFilterEntry)

			if err != nil {
				errorID := "4e7bd8b6-b0be-4f79-8d23-679215d074ed"
				errorMessage := fmt.Sprintf("could not do boolean arithmetic, OR [ErrorID=%s, err='%s']",
					errorID,
					err.Error())

				log.Fatalln(errorMessage)
			}

			// Get TestSuiteUuid
			testSuiteUuid = tempTestSuite.NonEditableInformation.GetTestSuiteUuid()

			// Get precomputed MetaDataFilter for TestSuite
			var tempMetaDataFilterforTestSuite *boolbits.Entry
			tempMetaDataFilterforTestSuite, existInMap = listTestSuitesModel.SimpleTestSuiteMetaDataFilterEntryMap[testSuiteUuid]

			// When TestSuite doesn't have any filter set then don't show that TestSuite
			if existInMap == false {
				continue
			}

			// Apply filter
			resultEntry, err = resultEntry.And(tempMetaDataFilterforTestSuite)

			if err != nil {
				errorID := "7547e4ed-a83f-4072-a72f-df4df7d1d777"
				errorMessage := fmt.Sprintf("could not do boolean arithmetic, AND [ErrorID=%s, err='%s']",
					errorID,
					err.Error())

				log.Fatalln(errorMessage)
			}

			// Check if TestSuite-filter matches the user set MetaDataFilter. If not the drop TestSuite for the list
			if resultEntry.Equals(testSuiteMetaDataFilterEntry) == false {
				continue
			}

		}

		// Create temporary Row-object for the table
		var tempRowslice []string

		// Populate the temporary Row-object
		// Column x
		// TestSuiteSelectedForTestSuite
		if listTestSuiteUIObject.howShouldItBeUsed == UsedForTestSuiteBuilder {
			tempRowslice = append(tempRowslice, "-666")
		}

		// Column 0:
		// DomainName
		var domainNameForTable string
		domainNameForTable = fmt.Sprintf("%s [%s]",
			tempTestSuite.NonEditableInformation.GetDomainName(),
			tempTestSuite.NonEditableInformation.GetDomainUuid()[0:8])

		tempRowslice = append(tempRowslice, domainNameForTable)

		// Column 1:
		// TestSuiteName
		tempRowslice = append(tempRowslice, tempTestSuite.NonEditableInformation.GetDomainName())

		// Column 2:
		// TestSuiteUuid
		tempRowslice = append(tempRowslice, tempTestSuite.NonEditableInformation.GetTestSuiteUuid())

		// Column 3:
		// TestSuiteVersion
		if listTestSuiteUIObject.howShouldItBeUsed == UsedForTestSuitesList {
			tempRowslice = append(tempRowslice, strconv.Itoa(int(tempTestSuite.NonEditableInformation.GetTestSuiteVersion())))
		}

		// Column 4:
		// LatestTestSuiteExecutionStatus
		if listTestSuiteUIObject.howShouldItBeUsed == UsedForTestSuitesList {

			var tempLatestTestSuiteExecutionStatus string

			if tempTestSuite.GetLatestTestSuiteExecutionStatus() > 0 {

				tempLatestTestSuiteExecutionStatus = detailedExecutionsModel.ExecutionStatusColorMap[int32(tempTestSuite.GetLatestTestSuiteExecutionStatus())].ExecutionStatusName
			} else {
				tempLatestTestSuiteExecutionStatus = "<no execution>"
			}

			tempRowslice = append(tempRowslice, tempLatestTestSuiteExecutionStatus)

		}

		// Column 5:
		// LatestTestSuiteExecutionStatusInsertTimeStamp
		if listTestSuiteUIObject.howShouldItBeUsed == UsedForTestSuitesList {

			var tempLatestTestSuiteExecutionStatusInsertTimeStamp string

			if tempTestSuite.GetLatestTestSuiteExecutionStatusInsertTimeStamp() != nil {
				tempLatestTestSuiteExecutionStatusInsertTimeStamp = sharedCode.ConvertGrpcTimeStampToStringForDB(tempTestSuite.
					GetLatestTestSuiteExecutionStatusInsertTimeStamp())
			} else {
				tempLatestTestSuiteExecutionStatusInsertTimeStamp = "<no execution>"
			}
			tempRowslice = append(tempRowslice, tempLatestTestSuiteExecutionStatusInsertTimeStamp)

		}

		// Column 6:
		// LatestFinishedOkTestSuiteExecutionStatusInsertTimeStamp
		if listTestSuiteUIObject.howShouldItBeUsed == UsedForTestSuitesList {

			var tempLatestFinishedOkTestSuiteExecutionStatusInsertTimeStamp string

			if tempTestSuite.GetLatestFinishedOkTestSuiteExecutionStatusInsertTimeStamp() != nil {
				tempLatestFinishedOkTestSuiteExecutionStatusInsertTimeStamp = sharedCode.ConvertGrpcTimeStampToStringForDB(
					tempTestSuite.GetLatestFinishedOkTestSuiteExecutionStatusInsertTimeStamp())
			} else {
				tempLatestFinishedOkTestSuiteExecutionStatusInsertTimeStamp = "<no successful execution yet>"
			}
			tempRowslice = append(tempRowslice, tempLatestFinishedOkTestSuiteExecutionStatusInsertTimeStamp)
		}

		// Column 7:
		// LastSavedTimeStamp
		if listTestSuiteUIObject.howShouldItBeUsed == UsedForTestSuitesList {

			var tempLastSavedTimeStamp string

			if tempTestSuite.GetLastSavedTimeStamp() != nil {
				tempLastSavedTimeStamp = sharedCode.ConvertGrpcTimeStampToStringForDB(tempTestSuite.
					GetLastSavedTimeStamp())
			} else {
				tempLastSavedTimeStamp = "<This should not happen, due to it must have been saved!>"
			}
			tempRowslice = append(tempRowslice, tempLastSavedTimeStamp)

		}

		// Column 8:
		// DomainUuid

		if listTestSuiteUIObject.howShouldItBeUsed == UsedForTestSuitesList {

			tempRowslice = append(tempRowslice, tempTestSuite.NonEditableInformation.GetDomainUuid())
		}

		// Add Row to slice of rows for the table
		listTestSuiteUIObject.testSuitesListTableTable = append(listTestSuiteUIObject.testSuitesListTableTable, tempRowslice)

	}

	// Do an initial sort 'testSuitesListTableTable' descending on 'LastSavedTimeStamp'
	if listTestSuitesModel.TestSuitesThatCanBeEditedByUserMap != nil &&
		len(listTestSuitesModel.TestSuitesThatCanBeEditedByUserMap) > 0 {

		if listTestSuiteUIObject.howShouldItBeUsed == UsedForTestSuitesList {
			listTestSuiteUIObject.currentSortColumn = initialColumnToSortOnForTestSuitesList
			listTestSuiteUIObject.sort2DStringSlice(
				listTestSuiteUIObject.testSuitesListTableTable,
				initialColumnToSortOnForTestSuitesList,
				initialSortDirectionForInitialColumnToSortOn)
		} else {
			listTestSuiteUIObject.currentSortColumn = initialColumnToSortOnForTestSuiteBuilder
			listTestSuiteUIObject.sort2DStringSlice(
				listTestSuiteUIObject.testSuitesListTableTable,
				initialColumnToSortOnForTestSuiteBuilder,
				initialSortDirectionForInitialColumnToSortOn)

		}

	}

}

// Sort2DStringSlice sorts a 2D string slice by a specified column index.
// It assumes that the column index is valid for all rows in the slice.
func (listTestSuiteUIObject *ListTestSuiteUIStruct) sort2DStringSlice(data [][]string, columnToSortOn int, sortingDirection SortingDirectionType) {
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
			if listTestSuiteUIObject.howShouldItBeUsed == UsedForTestSuitesList {
				// UsedForTestSuitesList
				return data[i][columnToSortOn] < data[j][columnToSortOn]

			} else {
				// UsedForTestSuiteBuilder
				switch true {
				case data[i][0] == "SELECTED" && data[j][0] == "SELECTED":

					return data[i][columnToSortOn] < data[j][columnToSortOn]

					break

				case data[i][0] == "SELECTED" && data[j][0] != "SELECTED":

					return true

					break

				case data[i][0] != "SELECTED" && data[j][0] == "SELECTED":

					return false

					break

				default:
					// Standard sorting - should not be needed
					return data[i][columnToSortOn] < data[j][columnToSortOn]
				}

			}

		case SortingDirectionDescending:

			if err1 == nil && err2 == nil {
				return num1 > num2
			}

			// Default to string comparison if not numbers.
			// Default to string comparison if not numbers.
			if listTestSuiteUIObject.howShouldItBeUsed == UsedForTestSuitesList {
				// UsedForTestSuitesList
				return data[i][columnToSortOn] < data[j][columnToSortOn]

			} else {
				// UsedForTestSuiteBuilder
				switch true {
				case data[i][0] == "SELECTED" && data[j][0] == "SELECTED":

					return data[i][columnToSortOn] > data[j][columnToSortOn]

					break

				case data[i][0] == "SELECTED" && data[j][0] != "SELECTED":

					return true

					break

				case data[i][0] != "SELECTED" && data[j][0] == "SELECTED":

					return false

					break

				default:
					// Standard sorting - should not be needed
					return data[i][columnToSortOn] > data[j][columnToSortOn]
				}

			}

		}

		// Not important due that switch statement will handle all return values
		return true

	})
}
