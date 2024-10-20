package listTestCasesUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"sync"
)

// The UI-table for the List with TestCase
var testCaseListTable *widget.Table

// The data source used to produce the UI-table for the List with TestCase
var testCaseListTableTable [][]string

// Keeps the number of TestCase that is shown in the list, after local filter is applied
var numberOfTestCasesAfterLocalFilters binding.String

// Keeps the number of TestCase that have been retrieved from the Database
var numberOfTestCasesInTheDatabaseSearch binding.String

var testCaseListTableHeader = []string{
	"DomainName", "TestCaseName", "TestCaseUuid", "TestCaseVersion", "Latest TestCaseExecution Status",
	"Latest TestCaseExecution TimeStamp", "Latest OK Finished TestCaseExecution TimeStamp", "Last Saved TimeStamp", "DomainUuid"}

// The number of visible columns in UI-table for TestCases
const numberColumnsInTestCasesListUI int = 9

// Keeps track of the in which column the TestCaseUUID exist in the data source for the UI-table
const testCaseUuidColumnNumber uint8 = 2

// Keeps track of the in which column the "Latest TestCaseExecution Status" exist in the data source for the UI-table
const latestTestCaseExecutionStatus uint8 = 4

// Keeps track of the in which column the ""Latest TestCaseExecution TimeStamp"" exist in the data source for the UI-table
const latestTestCaseExecutionTimeStamp uint8 = 5

// Keeps track of the in which column the ""Latest TestCaseExecution TimeStamp"" exist in the data source for the UI-table
const latestOkFinishedTestCaseExecutionTimeStamp uint8 = 6

// The row that the mouse is currently hovering above. Used for highlight that row in the UI-Table
var currentRowThatMouseIsHoveringAbove int16 = -1

// Use a mutex to synchronize access to 'currentRowThatMouseIsHoveringAbove'
var currentRowThatMouseIsHoveringAboveMutex sync.Mutex

// The size of the rectangles used for TestInstructionContainers-processing type and the color of the TestInstruction
const testCaseNodeRectangleSize = 40

// The TestCase Preview-container
var testCasePreviewContainerScroll *container.Scroll
var testCasePreviewContainer *fyne.Container

// The Split container have both the TestCase-list and the Preview-container in it
var TestCaseListAndTestCasePreviewSplitContainer *container.Split

// The TestCase that is shown in Preview
var testCaseThatIsShownInPreview string

// SortingDirectionType
// Define type for Sorting Direction
type SortingDirectionType uint8

const (
	SortingDirectionUnSpecified SortingDirectionType = iota
	SortingDirectionAscending
	SortingDirectionDescending
)

// Define initial SortDirection for when the Table is first shown to use
const (
	initialSortDirectionForInitialColumnToSortOn SortingDirectionType = SortingDirectionDescending
	initialColumnToSortOn                        int                  = 7
)

// The current column that the TestCase-list is sorted on
var currentSortColumn int

// The current Columns SortDirect
var currentSortColumnsSortDirection SortingDirectionType

// The previous column that the TestCase-list is sorted on
var previousSortColumn int
