package listTestCaseExecutionsUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"sync"
)

// The UI-table for the List with TestCaseExecutions
var testCaseExecutionsListTable *widget.Table

// The data source used to produce the UI-table for the List with TestCaseExecutions
var testCaseExecutionsListTableTable [][]string

// Keeps the number of TestCaseExecutions that is shown in the list, after local filter is applied
var numberOfTestCaseExecutionsAfterLocalFilters binding.String

// Keeps the number of TestCaseExecutions that have been retrieved from the Database
var numberOfTestCaseExecutionsInTheDatabaseSearch binding.String

var testCaseExecutionsListTableHeader = []string{
	"DomainName", "TestSuiteName", "TestCaseName", "TestCaseVersion", "TestCaseExecutionUuid", "Latest TestCaseExecution Status",
	"TestCaseExecution Start TimeStamp", "TestCaseExecution Status Update TimeStamp", "TestCaseExecution Finished TimeStamp",
	"TestCaseUuid", "DomainUuid", "TestSuiteUuid", "TestSuiteExecutionUuid"}

const loadAllTestCaseExecutionsForOneTestCaseButtonTextPart1 = "Load all executions for TestCase: "

var loadAllTestCaseExecutionsForOneTestCaseButtonText string

var loadAllTestCaseExecutionsForOneTestCaseButtonReference *widget.Button

// The number of visible columns in UI-table for TestCaseExecutions
const numberColumnsInTestCaseExecutionsListUI int = 13

// Keeps track of the in which column the TestCaseExecutionUUID exist in the data source for the UI-table
const testCaseExecutionUuidColumnNumber uint8 = 4

// Keeps track of the in which column the TestCaseUUID exist in the data source for the UI-table
const testCaseUuidColumnNumber uint8 = 9

// Keeps track of the in which column the TestCaseName exist in the data source for the UI-table
const testCaseNameColumnNumber uint8 = 2

// Keeps track of the in which column the "Latest TestCaseExecution Status" exist in the data source for the UI-table
const latestTestCaseExecutionStatus uint8 = 5

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

// The size of the rectangles used for indicate status of a TestInstructionExecution
const testCaseExecutionStatusRectangleHight = 30
const testCaseExecutionStatusRectangleWidth = 25

// The TestCase Preview-container
var testCaseExecutionPreviewContainerScroll *container.Scroll
var testCaseExecutionPreviewContainer *fyne.Container

// The Split container have both the TestCaseExecutions-list and the Preview-container in it
var TestCaseExecutionListAndTestCaseExecutionPreviewSplitContainer *container.Split

// The TestCaseExecutions that is shown in Preview
var testCaseExecutionThatIsShownInPreview string

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
var currentSortColumn int = -1

// The previous column that the TestCaseExecution-list was sorted on
var previousSortColumn int

var currentHeader *sortableHeaderLabelStruct

var previousHeader *sortableHeaderLabelStruct

// The current Columns SortDirect
var currentSortColumnsSortDirection SortingDirectionType
