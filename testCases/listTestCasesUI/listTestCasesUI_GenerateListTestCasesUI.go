package listTestCasesUI

import (
	detailedTestCaseExecutionsUI "FenixTesterGui/executions/detailedExecutionsUI"
	"FenixTesterGui/testCase/testCaseModel"
	"FenixTesterGui/testCases/listTestCasesModel"
	_ "embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"image"
	"strconv"
	"sync"
	"time"
)

//go:embed resources/TIC-Horizontal_32x32.png
var tic_parallellImage []byte
var imageData_tic_parallellImage image.Image

//go:embed resources/TIC-Vertical_32x32.png
var tic_serialImage []byte
var imageData_tic_serialImage image.Image

//go:embed resources/sort_cropped_32x51.png
var sortUnspecifiedImageAsByteArray []byte
var sortImageUnspecifiedAsImage image.Image

//go:embed resources/sort-down-descending_cropped_32x51.png
var sortImageAscendingAsByteArray []byte
var sortImageAscendingAsImage image.Image

//go:embed resources/sort-up-ascending_cropped_32x51.png
var sortImageDescendingAsByteArray []byte
var sortImageDescendingAsImage image.Image

func InitiateListTestCaseUIObject(
	tempHowShouldItBeUsed UsedForTestCasesListType,
	selectedTestCasesPtr *map[string]*fenixGuiTestCaseBuilderServerGrpcApi.TestCaseInTestSuiteMessage) (
	listTestCaseUIObject *ListTestCaseUIStruct) {

	listTestCaseUIObject = &ListTestCaseUIStruct{
		howShouldItBeUsed:                               tempHowShouldItBeUsed,
		testCaseListTable:                               nil,
		testCasesListTableTable:                         nil,
		numberOfTestCasesAfterLocalFilters:              nil,
		numberOfTestCasesInTheDatabaseSearch:            nil,
		sortableHeaderReference:                         nil,
		currentRowThatMouseIsHoveringAbove:              -1,
		currentRowThatMouseIsHoveringAboveMutex:         sync.Mutex{},
		testCasePreviewContainerScroll:                  nil,
		testCasePreviewContainer:                        nil,
		preViewAndFilterTabs:                            nil,
		preViewTab:                                      nil,
		filterTab:                                       nil,
		simpleTestCaseMetaDataSelectedDomainUuid:        "",
		simpleTestCaseMetaDataSelectedDomainDisplayName: "",
		testCaseFullMetaDataFilterContainer:             nil,
		testCaseMainAreaForMetaDataFilterContainer:      nil,
		newMandatoryOwnerDomainSelect:                   nil,
		filterOnMetaDataFunction:                        nil,
		calculateMetaDataFilterFunction:                 nil,
		useAutoFilter:                                   false,
		simpleMetaDataFilterEntryMap:                    nil,
		testCaseThatIsShownInPreview:                    "",
		currentSortColumn:                               -1,
		currentHeader:                                   nil,
		previousHeader:                                  nil,
		selectedTestCasesPtr:                            selectedTestCasesPtr,
		//testCaseListTableHeader:                       ???
	}

	// Handle Headers for TestCaseList by adding first column for when TestCase is selected
	if tempHowShouldItBeUsed == UsedForTestCasesList {
		listTestCaseUIObject.testCaseListTableHeader = testCaseListTableHeaderForTestCasesList

	} else {
		listTestCaseUIObject.testCaseListTableHeader = testCaseListTableHeaderForTestSuiteBuilder
	}

	return listTestCaseUIObject

}

// Create the UI used for list all TestCasesMapPtr that the User can edit
func (listTestCaseUIObject *ListTestCaseUIStruct) GenerateListTestCasesUI(
	testCasesModel *testCaseModel.TestCasesModelsStruct,
	preViewAndFilterTabsUsedForCreateTestSuite *container.AppTabs) (
	_ *fyne.Container) {

	//var testCaseTable *widget.Table

	var tempTestCaseListAndTestCasePreviewSplitContainer *container.Split

	var testCasesListContainer *fyne.Container
	var testCasesListScrollContainer *container.Scroll
	var statisticsContainer *fyne.Container
	var executionColorPaletteContainer *fyne.Container
	var statisticsAndColorPaletteContainer *fyne.Container

	var loadTestCaseFromDataBaseButton *widget.Button
	var loadTestCaseFromDataBaseFunction func()
	var filterTestCasesButton *widget.Button
	var filterTestCasesButtonFunction func()
	var clearFiltersButton *widget.Button
	var clearFiltersButtonFunction func()
	var buttonsContainer *fyne.Container

	var numberOfTestCasesAfterLocalFilterLabel *widget.Label
	var numberOfTestCasesRetrievedFromDatabaseLabel *widget.Label

	var filterAndButtonsContainer *fyne.Container

	// Define the function to be executed to load TestCasesMapPtr from that Database that the user can edit
	loadTestCaseFromDataBaseFunction = func() {
		fmt.Println("'loadTestCaseFromDataBaseButton' was pressed")
		listTestCasesModel.LoadTestCaseThatCanBeEditedByUser(
			testCasesModel,
			time.Now().Add(-time.Hour*10000), time.Now().Add(-time.Hour*10000))
		filterTestCasesButtonFunction()
		//sortTestCasesTable()
	}

	// Define the 'loadTestCaseFromDataBaseButton'
	loadTestCaseFromDataBaseButton = widget.NewButton("Load TestCasesMapPtr from Database", loadTestCaseFromDataBaseFunction)

	// Define the function to be executed to filter TestCasesMapPtr that the user can edit
	filterTestCasesButtonFunction = func() {
		fmt.Println("'filterTestCasesButton' was pressed")
		listTestCaseUIObject.loadTestCaseListTableTable(nil)
		listTestCaseUIObject.calculateAndSetCorrectColumnWidths()
		listTestCaseUIObject.updateTestCasesListTable(testCasesModel)

		// Update the number TestCasesMapPtr in the list
		var numberOfRowsAsString string
		numberOfRowsAsString = strconv.Itoa(len(listTestCaseUIObject.testCasesListTableTable))
		listTestCaseUIObject.numberOfTestCasesAfterLocalFilters.Set(
			fmt.Sprintf("Number of TestCasesMapPtr after local filters was applied: %s",
				numberOfRowsAsString))

		// Update the number TestCasesMapPtr retrieved from Database
		var numberOfRowsFromDatabaseAsString string
		numberOfRowsFromDatabaseAsString = strconv.Itoa(len(listTestCaseUIObject.testCasesListTableTable))
		listTestCaseUIObject.numberOfTestCasesInTheDatabaseSearch.Set(
			fmt.Sprintf("Number of TestCasesMapPtr retrieved from the Database: %s",
				numberOfRowsFromDatabaseAsString))

		listTestCaseUIObject.sortableHeaderReference.sortImage.onTapped()

	}

	// Define the 'filterTestCasesButton'
	filterTestCasesButton = widget.NewButton("Filter TestCasesMapPtr", filterTestCasesButtonFunction)

	// Define the function to be executed to list TestCasesMapPtr that the user can edit
	clearFiltersButtonFunction = func() {
		fmt.Println("'clearFiltersButtonFunction' was pressed")
	}

	// Define the 'filterTestCasesButton'
	clearFiltersButton = widget.NewButton("Clear all search filters", clearFiltersButtonFunction)

	// Add the buttons to the buttonsContainer
	buttonsContainer = container.NewHBox(loadTestCaseFromDataBaseButton, filterTestCasesButton, clearFiltersButton)

	// Add objects to the 'filterAndButtonsContainer'
	filterAndButtonsContainer = container.NewVBox(buttonsContainer)

	// Generate the ExecutionColorPaletteContainer
	executionColorPaletteContainer = detailedTestCaseExecutionsUI.GenerateExecutionColorPalette()

	// Initiate the Table
	listTestCaseUIObject.generateTestCasesListTable(testCasesModel)
	testCaseTableContainer := container.NewBorder(nil, nil, nil, nil, listTestCaseUIObject.testCaseListTable)

	// Create the Scroll container for the List
	testCasesListScrollContainer = container.NewScroll(testCaseTableContainer)

	// Create the label used for showing number of TestCasesMapPtr in the local filter
	listTestCaseUIObject.numberOfTestCasesAfterLocalFilters = binding.NewString()
	_ = listTestCaseUIObject.numberOfTestCasesAfterLocalFilters.Set("No TestCasesMapPtr in the List")
	numberOfTestCasesAfterLocalFilterLabel = widget.NewLabelWithData(listTestCaseUIObject.numberOfTestCasesAfterLocalFilters)

	// Create the label used for showing number of TestCasesMapPtr retrieved from the Database
	listTestCaseUIObject.numberOfTestCasesInTheDatabaseSearch = binding.NewString()
	_ = listTestCaseUIObject.numberOfTestCasesInTheDatabaseSearch.Set("No TestCasesMapPtr retrieved from the Database")
	numberOfTestCasesRetrievedFromDatabaseLabel = widget.NewLabelWithData(listTestCaseUIObject.numberOfTestCasesInTheDatabaseSearch)

	// Initiate 'statisticsContainer'
	statisticsContainer = container.NewHBox(numberOfTestCasesAfterLocalFilterLabel, numberOfTestCasesRetrievedFromDatabaseLabel)

	statisticsAndColorPaletteContainer = container.NewVBox(executionColorPaletteContainer, statisticsContainer)

	// Add 'testCasesListScrollContainer' to 'testCasesListContainer'
	testCasesListContainer = container.NewBorder(filterAndButtonsContainer, statisticsAndColorPaletteContainer, nil, nil, testCasesListScrollContainer)
	testCasesListScrollContainer2 := container.NewScroll(testCasesListContainer)

	// Create the Temporary container that should be shown
	temporaryContainer := container.NewCenter(widget.NewLabel("Select a TestCase to get the Preview"))

	listTestCaseUIObject.testCasePreviewContainer = container.NewBorder(nil, nil, nil, nil, temporaryContainer)

	// Generate the container for the Preview, 'testCasePreviewContainer'
	listTestCaseUIObject.testCasePreviewContainerScroll = container.NewScroll(listTestCaseUIObject.testCasePreviewContainer)

	// Generate the Tab for the PreView
	listTestCaseUIObject.preViewTab = container.NewTabItem(
		"PreView",
		listTestCaseUIObject.testCasePreviewContainerScroll)

	// Generate the 'GenerateTestSuiteMetaDataFilterContainer'
	var simpleAndAdvancedMetaDataFilter *container.AppTabs
	simpleAndAdvancedMetaDataFilter = listTestCaseUIObject.GenerateTestCaseMetaDataFilterContainer(testCasesModel)

	// Generate Tab for TestCase-filter
	listTestCaseUIObject.filterTab = container.NewTabItem(
		"TestCase-filter",
		simpleAndAdvancedMetaDataFilter)

	// Generate the AppTabsContainer, depending on from where it was initiated
	if listTestCaseUIObject.howShouldItBeUsed == UsedForTestCasesList {
		// We are in standard List TestCases
		listTestCaseUIObject.preViewAndFilterTabs = container.NewAppTabs(listTestCaseUIObject.filterTab, listTestCaseUIObject.preViewTab)

		tempTestCaseListAndTestCasePreviewSplitContainer = container.NewHSplit(testCasesListScrollContainer2, listTestCaseUIObject.preViewAndFilterTabs)
		tempTestCaseListAndTestCasePreviewSplitContainer.Offset = 0.75

		//TestCaseListAndTestCasePreviewSplitContainer = tempTestCaseListAndTestCasePreviewSplitContainer

		return container.NewBorder(nil, nil, nil, nil, tempTestCaseListAndTestCasePreviewSplitContainer)

	} else {
		// We are in Create TestSuite
		listTestCaseUIObject.preViewAndFilterTabs = preViewAndFilterTabsUsedForCreateTestSuite
		listTestCaseUIObject.preViewAndFilterTabs.Append(listTestCaseUIObject.filterTab)
		listTestCaseUIObject.preViewAndFilterTabs.Append(listTestCaseUIObject.preViewTab)

		testCasesListContainer := container.NewBorder(nil, nil, nil, nil, testCasesListScrollContainer2)

		return testCasesListContainer

	}

}
