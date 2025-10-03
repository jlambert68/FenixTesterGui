package listTestSuiteExecutionsUI

import (
	"FenixTesterGui/testCaseExecutions/testCaseExecutionsModel"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"sync"
)

// The UI-table for the List with TestSuiteExecutions
var testSuiteExecutionsListTable *widget.Table

var testSuiteExecutionsListTableHeadersMapRef map[int]*sortableHeaderLabelStruct

// The data source used to produce the UI-table for the List with TestSuiteExecutions
var testSuiteExecutionsListTableTable [][]string

// Keeps the number of TestSuiteExecutions that is shown in the list, after local filter is applied
var numberOfTestSuiteExecutionsAfterLocalFilters binding.String

// Keeps the number of TestSuiteExecutions that have been retrieved from the Database
var numberOfTestSuiteExecutionsInTheDatabaseSearch binding.String

var testSuiteExecutionsListTableHeader = []string{
	"DomainName", "TestSuiteName", "TestSuiteName", "TestSuiteVersion", "TestSuiteExecutionUuid", "Latest TestSuiteExecution Status",
	"TestSuiteExecution Start TimeStamp", "TestSuiteExecution Status Update TimeStamp", "TestSuiteExecution Finished TimeStamp",
	"TestSuiteUuid", "DomainUuid", "TestSuiteUuid", "TestSuiteExecutionUuid"}

const loadAllTestSuiteExecutionsForOneTestSuiteButtonTextPart1 = "Load all executions for TestSuite: "

var loadAllTestSuiteExecutionsForOneTestSuiteButtonText string

var loadAllTestSuiteExecutionsForOneTestSuiteButtonReference *widget.Button

// The number of visible columns in UI-table for TestSuiteExecutions
const numberColumnsInTestSuiteExecutionsListUI int = 13

// Keeps track of the in which column the TestSuiteExecutionUUID exist in the data source for the UI-table
const testSuiteExecutionUuidColumnNumber uint8 = 4

// const testSuiteExecutionVersionColumnNumber uint8 = xxx

// Keeps track of the in which column the TestSuiteUUID exist in the data source for the UI-table
const testSuiteUuidColumnNumber uint8 = 9

// Keeps track of the in which column the TestSuiteName exist in the data source for the UI-table
const testSuiteNameColumnNumber uint8 = 2

// Keeps track of the in which column the "Latest TestSuiteExecution Status" exist in the data source for the UI-table
const latestTestSuiteExecutionStatus uint8 = 5

// Keeps track of the in which column the "Latest TestSuiteExecution TimeStamp" exist in the data source for the UI-table
const latestTestSuiteExecutionTimeStamp uint8 = 6

// Keeps track of the in which column the ""Latest TestSuiteExecution TimeStamp"" exist in the data source for the UI-table
const latestOkFinishedTestSuiteExecutionTimeStamp uint8 = 6

// Keeps track of the in which column the ""Latest TestSuiteExecution TimeStamp"" exist in the data source for the UI-table
const latestTestSuiteExecutionTimeStampColumnNumber uint8 = 6

// Reference to Header for column for "Latest TestSuiteExecution TimeStamp"
//var sortableHeaderReference *sortableHeaderLabelStruct

// The row that the mouse is currently hovering above. Used for highlight that row in the UI-Table
var currentRowThatMouseIsHoveringAbove int16 = -1

// Use a mutex to synchronize access to 'currentRowThatMouseIsHoveringAbove'
var currentRowThatMouseIsHoveringAboveMutex sync.Mutex

// Indicate of mouse left table for the Preview-area
var mouseHasLeftTable bool = true

// Indicate of mouse left TestSuiteExecutionPreview for the Tab-explorer area
var mouseHasLeftTestSuiteExecutionPreviewTree bool = true

// The size of the rectangles used for TestInstructionContainers-processing type and the color of the TestInstruction
const testSuiteNodeRectangleSize = 40

// The size of the rectangles used for indicate status of a TestInstructionExecution
const testCaseExecutionStatusRectangleHeight = 30
const testCaseExecutionStatusRectangleWidth = 15

// The size of the rectangles used for log-status in execution log
const logStatusRectangleHeight = 25
const logStatusRectangleWidth = 25

// Struct holding all variables needed for the PreView
type TestSuiteInstructionPreViewStruct struct {
	// Split container holding TestSuite-PreView, to the left, and Explorer-Tabs to the Right
	testSuitePreviewTestInstructionExecutionLogSplitContainer *container.Split

	// The TestSuite Preview-container
	testSuiteExecutionPreviewContainerScroll *container.Scroll
	testSuiteExecutionPreviewContainer       *fyne.Container

	testInstructionsExecutionDetailsContainerScroll    *container.Scroll
	testInstructionsExecutionLogContainer              *fyne.Container
	testInstructionsExecutionAttributesContainerScroll *container.Scroll
	testInstructionsExecutionAttributesContainer       *fyne.Container
	testInstructionsExecutionDetailsContainer          *fyne.Container
	preViewTabs                                        *container.AppTabs
	attributeExplorerTab                               *container.TabItem
	logsExplorerTab                                    *container.TabItem
	testInstructionDetailsExplorerTab                  *container.TabItem
}

// Variables holding all objects used in PreView to the right of the TestSuiteExecutions-list
var TestSuiteInstructionPreViewObject *TestSuiteInstructionPreViewStruct

// Reference to the TabObject that hold TestSuiteExecutions
var detailedTestSuiteExecutionsUITabObjectRef *container.DocTabs

// Reference to the Exit-functions for the Tabs in 'detailedTestSuiteExecutionsUITabObjectRef'
var exitFunctionsForDetailedTestSuiteExecutionsUITabObjectPtr *map[*container.TabItem]func()

type openedDetailedTestSuiteExecutionsMapKeyType string // TestSuiteExecutionUuid + TestSuiteExecutionVersion

// Keeps track if a TestSuiteExecution is open in Tab and/or Window
type openedDetailedTestSuiteExecutionStruct struct {
	isTestSuiteExecutionOpenInTab                     bool
	TestSuiteInstructionPreViewObjectInTab            *TestSuiteInstructionPreViewStruct
	isTestSuiteExecutionOpenInExternalWindow          bool
	TestSuiteInstructionPreViewObjectInExternalWindow *TestSuiteInstructionPreViewStruct
	externalWindow                                    fyne.Window
	tabItem                                           *container.TabItem
}

// Map keeping track of all opened TestSuiteExecutions, in Tabs and/or Separate windows
var openedDetailedTestSuiteExecutionsMapPtr *map[openedDetailedTestSuiteExecutionsMapKeyType]*openedDetailedTestSuiteExecutionStruct

// From where is the opening of the TestSuiteExecution initiated; FromExecutionList, FromExternalWindow, FromTab
type openedTestSuiteExecutionFromType uint8

const (
	fromNotDefined    openedTestSuiteExecutionFromType = iota
	fromExecutionList openedTestSuiteExecutionFromType = iota
	fromExternalWindow
	fromTab
)

//var tempTestSuitePreviewTestInstructionExecutionLogSplitContainer

// Mutex for Attributes-map below
var testSuiteExecutionAttributesForPreviewMapMutex = &sync.Mutex{}

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

// TestSuiteExecutionListAndTestSuiteExecutionPreviewSplitContainer
// The Split container have both the TestSuiteExecutions-list and the Preview-container in it
var testSuiteExecutionListAndTestSuiteExecutionPreviewSplitContainer *container.Split

// Struct for holding all data for the Execution-list when data belongs  "One Execution per TestSuite"
type oneExecutionPerTestSuiteListObjectStruct struct {

	// The last TestSuiteExecution that was picked when we are in "One Execution per TestSuite
	lastSelectedTestSuiteExecutionForOneExecutionPerTestSuite string

	// The TestSuiteUuid for TestSuiteExecutions that is shown in Preview
	testSuiteUuidForTestSuiteExecutionThatIsShownInPreview string

	// The TestSuiteExecutions that is shown in Preview
	testSuiteExecutionUuidThatIsShownInPreview string

	// The TestSuiteExecutionVersion that is shown in Preview
	testSuiteExecutionVersionThatIsShownInPreview uint32

	// Is a row selected or not
	isAnyRowSelected bool

	// The current column that the TestSuiteExecutions-list is sorted on
	currentSortColumn int

	// The previous column that the TestSuiteExecution-list was sorted on
	previousSortColumn int

	currentHeader *sortableHeaderLabelStruct

	previousHeader *sortableHeaderLabelStruct

	// The current Columns SortDirect
	currentSortColumnsSortDirection SortingDirectionType
}

// Struct for holding all data for the Execution-list when data belongs "All Executions for one TestSuite"
type allExecutionsFoOneTestSuiteListObjectStruct struct {

	// The last TestSuiteExecution that was picked when we are in "All Executions for one TestSuite"
	lastSelectedTestSuiteExecutionForAllExecutionsForOneTestSuite string

	// The TestSuiteUuid for TestSuiteExecutions that is shown in Preview
	testSuiteUuidForTestSuiteExecutionThatIsShownInPreview string

	// The TestSuiteExecutionUuid that is shown in Preview
	testSuiteExecutionUuidThatIsShownInPreview string

	// The TestSuiteExecutionVersion that is shown in Preview
	testSuiteExecutionVersionThatIsShownInPreview uint32

	// Is a row selected or not
	isAnyRowSelected bool

	// The current column that the TestSuiteExecutions-list is sorted on
	currentSortColumn int

	// The previous column that the TestSuiteExecution-list was sorted on
	previousSortColumn int

	currentHeader *sortableHeaderLabelStruct

	previousHeader *sortableHeaderLabelStruct

	// The current Columns SortDirect
	currentSortColumnsSortDirection SortingDirectionType
}

// Object for keeping if a row is selected and which TestSuite/Execution that should be shown
var selectedTestSuiteExecutionObjected selectedTestSuiteExecutionStruct

// Struct for keeping if a row is selected and which TestSuite/Execution that should be shown
type selectedTestSuiteExecutionStruct struct {

	// Hold all data for the Execution-list when data belongs 'One Execution per TestSuite'
	oneExecutionPerTestSuiteListObject oneExecutionPerTestSuiteListObjectStruct

	// Hold all data for the Execution-list when data belongs "All Executions for one TestSuite"
	allExecutionsFoOneTestSuiteListObject allExecutionsFoOneTestSuiteListObjectStruct

	// ExecutionsInGuiIsOfType
	// The variable that keeps track on if TestSuitesExecutions in the GUI-list comes from
	// "One Execution per TestSuite" or "All Executions for one TestSuite"
	ExecutionsInGuiIsOfType CurrenExecutionListType
}

// CurrenExecutionListType
// The type that defines if TestSuitesExecutions in the GUI-list comes from
// "One Execution per TestSuite" or "All Executions for one TestSuite"
type CurrenExecutionListType uint8

// The constants that defines if TestSuitesExecutions in the GUI-list comes from
// "One Execution per TestSuite" or "All Executions for one TestSuite"
const (
	NotDefined CurrenExecutionListType = iota
	OneExecutionPerTestSuite
	AllExecutionsForOneTestSuite
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
