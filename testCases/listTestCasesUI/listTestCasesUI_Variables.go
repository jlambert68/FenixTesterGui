package listTestCasesUI

import (
	"FenixTesterGui/testCase/testCaseModel"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/jlambert68/Fast_BitFilter_MetaData/boolbits/boolbits"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"sync"
)

var StandardListTesCasesUIObject *ListTestCaseUIStruct

type ListTestCaseUIStruct struct {
	howShouldItBeUsed UsedForTestCasesListType

	// The UI-table for the List with TestCase
	testCaseListTable *widget.Table

	// The data source used to produce the UI-table for the List with TestCase
	testCasesListTableTable [][]string

	// Keeps the number of TestCase that is shown in the list, after local filter is applied
	numberOfTestCasesAfterLocalFilters binding.String

	// Keeps the number of TestCase that have been retrieved from the Database
	numberOfTestCasesInTheDatabaseSearch binding.String

	sortableHeaderReference *sortableHeaderLabelStruct

	// The row that the mouse is currently hovering above. Used for highlight that row in the UI-Table
	currentRowThatMouseIsHoveringAbove int16

	// Use a mutex to synchronize access to 'currentRowThatMouseIsHoveringAbove'
	currentRowThatMouseIsHoveringAboveMutex sync.Mutex

	// The TestCase Preview-container
	testCasePreviewContainerScroll *container.Scroll
	testCasePreviewContainer       *fyne.Container

	preViewAndFilterTabs *container.AppTabs
	preViewTab           *container.TabItem
	filterTab            *container.TabItem

	simpleTestCaseMetaDataSelectedDomainUuid        string
	simpleTestCaseMetaDataSelectedDomainDisplayName string
	testCaseFullMetaDataFilterContainer             *fyne.Container
	testCaseMainAreaForMetaDataFilterContainer      *fyne.Container
	newMandatoryOwnerDomainSelect                   *customMandatorySelectComboBox

	filterOnMetaDataFunction        func(*boolbits.Entry, *testCaseModel.TestCasesModelsStruct)
	calculateMetaDataFilterFunction func()
	useAutoFilter                   bool

	// Holding all separate MetaDataEntries used in the Simple MetaData-filter
	simpleMetaDataFilterEntryMap map[string]simpleMetaDataFilterEntryMapStruct // Key = DomainUuid.GroupName.GroupItemName

	// The TestCase that is shown in Preview
	testCaseThatIsShownInPreview string

	// The current column that the TestCase-list is sorted on
	currentSortColumn int

	currentHeader *sortableHeaderLabelStruct

	previousHeader *sortableHeaderLabelStruct

	// Map holding a pointer to the object in the TestSuite having the selected TestCases for the TestSuite
	selectedTestCasesPtr *map[string]*fenixGuiTestCaseBuilderServerGrpcApi.TestCaseInTestSuiteMessage

	// Header for the TestCases listings and also deciding number of columns
	testCaseListTableHeader []string
}

type UsedForTestCasesListType uint8

const (
	UsedForTestCasesList UsedForTestCasesListType = iota
	UsedForTestSuiteBuilder
)

var testCaseListTableHeaderForTestSuiteBuilder = []string{"Selected", "DomainName", "TestCaseName", "TestCaseUuid"}

var testCaseListTableHeaderForTestCasesList = []string{
	"DomainName", "TestCaseName", "TestCaseUuid", "TestCaseVersion", "Latest TestCaseExecution Status",
	"Latest TestCaseExecution TimeStamp", "Latest OK Finished TestCaseExecution TimeStamp", "Last Saved TimeStamp", "DomainUuid"}

// The number of visible columns in UI-table for TestCasesMapPtr
const numberColumnsInTestCasesListUIForTestCasesList int = 9
const numberColumnsInTestCasesListUIForTestSuiteBuilder int = 4

// Keeps track of the in which column the TestCaseUUID exist in the data source for the UI-table
const testCaseUuidColumnNumber uint8 = 2

// Keeps track of the in which column the "Latest TestCaseExecution Status" exist in the data source for the UI-table
const latestTestCaseExecutionStatusColumnNumber uint8 = 4

// Keeps track of the in which column the ""Latest TestCaseExecution TimeStamp"" exist in the data source for the UI-table
const latestTestCaseExecutionTimeStampColumnNumber uint8 = 6

// Keeps track of the in which column the "Latest TestCaseExecution TimeStamp" exist in the data source for the UI-table
const latestOkFinishedTestCaseExecutionTimeStamp uint8 = 6

// Reference to Header for column for "Latest TestCaseExecution TimeStamp"

// The size of the rectangles used for TestInstructionContainers-processing type and the color of the TestInstruction
const testCaseNodeRectangleSize = 40

type simpleMetaDataFilterEntryMapStruct struct {
	valueEntryListToBeProcessedWithBooleanOrSlice []*boolbits.Entry
}

// The Split container have both the TestCase-list and the Preview-container in it
var TestCaseListAndTestCasePreviewSplitContainer *container.Split

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
	initialColumnToSortOnForTestCasesList        int                  = 7
	initialColumnToSortOnForTestSuiteBuilder     int                  = 2
)

// The previous column that the TestCase-list was sorted on
var previousSortColumn int

// The current Columns SortDirect
var currentSortColumnsSortDirection SortingDirectionType

const (
	autoFilterRadioGroupOn  = "AutoFilter - On"
	autoFilterRadioGroupOff = "AutoFilter - Off"
)
