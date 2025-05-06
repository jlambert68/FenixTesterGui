package listTestCaseExecutionsUI

import (
	"FenixTesterGui/testCaseExecutions/testCaseExecutionsModel"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"sync"
)

// The UI-table for the List with TestCaseExecutions
var testCaseExecutionsListTable *widget.Table

var testCaseExecutionsListTableHeadersMapRef map[int]*sortableHeaderLabelStruct

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

// const testCaseExecutionVersionColumnNumber uint8 = xxx

// Keeps track of the in which column the TestCaseUUID exist in the data source for the UI-table
const testCaseUuidColumnNumber uint8 = 9

// Keeps track of the in which column the TestCaseName exist in the data source for the UI-table
const testCaseNameColumnNumber uint8 = 2

// Keeps track of the in which column the "Latest TestCaseExecution Status" exist in the data source for the UI-table
const latestTestCaseExecutionStatus uint8 = 5

// Keeps track of the in which column the ""Latest TestCaseExecution TimeStamp"" exist in the data source for the UI-table
const latestTestCaseExecutionTimeStamp uint8 = 6

// Keeps track of the in which column the ""Latest TestCaseExecution TimeStamp"" exist in the data source for the UI-table
const latestOkFinishedTestCaseExecutionTimeStamp uint8 = 6

// Keeps track of the in which column the ""Latest TestCaseExecution TimeStamp"" exist in the data source for the UI-table
const latestTestCaseExecutionTimeStampColumnNumber uint8 = 6

// Reference to Header for column for "Latest TestCaseExecution TimeStamp"
//var sortableHeaderReference *sortableHeaderLabelStruct

// The row that the mouse is currently hovering above. Used for highlight that row in the UI-Table
var currentRowThatMouseIsHoveringAbove int16 = -1

// Use a mutex to synchronize access to 'currentRowThatMouseIsHoveringAbove'
var currentRowThatMouseIsHoveringAboveMutex sync.Mutex

// Indicate of mouse left table for the Preview-area
var mouseHasLeftTable bool = true

// Indicate of mouse left TestCaseExecutionPreview for the Tab-explorer area
var mouseHasLeftTestCaseExecutionPreviewTree bool = true

// The size of the rectangles used for TestInstructionContainers-processing type and the color of the TestInstruction
const testCaseNodeRectangleSize = 40

// The size of the rectangles used for indicate status of a TestInstructionExecution
const testCaseExecutionStatusRectangleHeight = 30
const testCaseExecutionStatusRectangleWidth = 15

// The size of the rectangles used for log-status in execution log
const logStatusRectangleHeight = 25
const logStatusRectangleWidth = 25

// The TestCase Preview-container
var testCaseExecutionPreviewContainerScroll *container.Scroll
var testCaseExecutionPreviewContainer *fyne.Container

var testInstructionsExecutionLogContainerScroll *container.Scroll
var testInstructionsExecutionLogContainer *fyne.Container
var testInstructionsExecutionAttributesContainerScroll *container.Scroll
var testInstructionsExecutionAttributesContainer *fyne.Container
var preViewTabs *container.AppTabs
var attributeExplorerTab *container.TabItem
var logsExplorerTab *container.TabItem

//var tempTestCasePreviewTestInstructionExecutionLogSplitContainer

// Mutex for Attributes-map below
var testCaseExecutionAttributesForPreviewMapMutex = &sync.Mutex{}

// Struct holding all attribute-containers for one TestInstructionsExecution and a list with Objects (TI or TIC)
// that is placed below this TestInstruction in indentation level
type testCaseExecutionAttributesForPreviewStruct struct {
	LabelType                                   labelTypeType
	LabelText                                   string
	attributesContainerShouldBeVisible          bool
	testInstructionExecutionAttributesContainer *fyne.Container
	childObjectsWithAttributes                  []testCaseExecutionsModel.TCEoTICoTIEAttributesContainerMapKeyType
}

// The map holding all TestInstructions and their Attributes-containers
var testCaseExecutionAttributesForPreviewMapPtr *map[testCaseExecutionsModel.TCEoTICoTIEAttributesContainerMapKeyType]*testCaseExecutionAttributesForPreviewStruct

// TestCaseExecutionListAndTestCaseExecutionPreviewSplitContainer
// The Split container have both the TestCaseExecutions-list and the Preview-container in it
var TestCaseExecutionListAndTestCaseExecutionPreviewSplitContainer *container.Split

// Struct for holding all data for the Execution-list when data belongs  "One Execution per TestCase"
type oneExecutionPerTestCaseListObjectStruct struct {

	// The last TestCaseExecution that was picked when we are in "One Execution per TestCase"
	lastSelectedTestCaseExecutionForOneExecutionPerTestCase string

	// The TestCaseUuid for TestCaseExecutions that is shown in Preview
	testCaseUuidForTestCaseExecutionThatIsShownInPreview string

	// The TestCaseExecutions that is shown in Preview
	testCaseExecutionUuidThatIsShownInPreview string

	// The TestCaseExecutionVersion that is shown in Preview
	testCaseExecutionVersionThatIsShownInPreview uint32

	// Is a row selected or not
	isAnyRowSelected bool

	// The current column that the TestCaseExecutions-list is sorted on
	currentSortColumn int

	// The previous column that the TestCaseExecution-list was sorted on
	previousSortColumn int

	currentHeader *sortableHeaderLabelStruct

	previousHeader *sortableHeaderLabelStruct

	// The current Columns SortDirect
	currentSortColumnsSortDirection SortingDirectionType
}

// Struct for holding all data for the Execution-list when data belongs "All Executions for one TestCase"
type allExecutionsFoOneTestCaseListObjectStruct struct {

	// The last TestCaseExecution that was picked when we are in "All Executions for one TestCase"
	lastSelectedTestCaseExecutionForAllExecutionsForOneTestCase string

	// The TestCaseUuid for TestCaseExecutions that is shown in Preview
	testCaseUuidForTestCaseExecutionThatIsShownInPreview string

	// The TestCaseExecutionUuid that is shown in Preview
	testCaseExecutionUuidThatIsShownInPreview string

	// The TestCaseExecutionVersion that is shown in Preview
	testCaseExecutionVersionThatIsShownInPreview uint32

	// Is a row selected or not
	isAnyRowSelected bool

	// The current column that the TestCaseExecutions-list is sorted on
	currentSortColumn int

	// The previous column that the TestCaseExecution-list was sorted on
	previousSortColumn int

	currentHeader *sortableHeaderLabelStruct

	previousHeader *sortableHeaderLabelStruct

	// The current Columns SortDirect
	currentSortColumnsSortDirection SortingDirectionType
}

// Object for keeping if a row is selected and which TestCase/Execution that should be shown
var selectedTestCaseExecutionObjected selectedTestCaseExecutionStruct

// Struct for keeping if a row is selected and which TestCase/Execution that should be shown
type selectedTestCaseExecutionStruct struct {

	// Hold all data for the Execution-list when data belongs 'One Execution per TestCase'
	oneExecutionPerTestCaseListObject oneExecutionPerTestCaseListObjectStruct

	// Hold all data for the Execution-list when data belongs "All Executions for one TestCase"
	allExecutionsFoOneTestCaseListObject allExecutionsFoOneTestCaseListObjectStruct

	// ExecutionsInGuiIsOfType
	// The variable that keeps track on if TestCasesExecutions in the GUI-list comes from
	// "One Execution per TestCase" or "All Executions for one TestCase"
	ExecutionsInGuiIsOfType CurrenExecutionListType
}

// CurrenExecutionListType
// The type that defines if TestCasesExecutions in the GUI-list comes from
// "One Execution per TestCase" or "All Executions for one TestCase"
type CurrenExecutionListType uint8

// The constants that defines if TestCasesExecutions in the GUI-list comes from
// "One Execution per TestCase" or "All Executions for one TestCase"
const (
	NotDefined CurrenExecutionListType = iota
	OneExecutionPerTestCase
	AllExecutionsForOneTestCase
)

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
