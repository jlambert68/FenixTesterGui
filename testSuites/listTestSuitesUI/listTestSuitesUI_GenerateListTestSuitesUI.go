package listTestSuitesUI

import (
	detailedTestCaseExecutionsUI "FenixTesterGui/executions/detailedExecutionsUI"
	"FenixTesterGui/testCase/testCaseModel"
	"FenixTesterGui/testSuites/listTestSuitesModel"
	_ "embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"image"
	"image/color"
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

func InitiateListTestSuiteUIObject(
	tempHowShouldItBeUsed UsedForTestSuitesListType,
	selectedTestSuitesPtr *map[string]*fenixGuiTestCaseBuilderServerGrpcApi.TestCaseInTestSuiteMessage) (
	listTestCaseUIObject *ListTestSuiteUIStruct) {

	listTestCaseUIObject = &ListTestSuiteUIStruct{
		howShouldItBeUsed:                                tempHowShouldItBeUsed,
		testSuiteListTable:                               nil,
		testSuitesListTableTable:                         nil,
		numberOfTestSuitesAfterLocalFilters:              nil,
		numberOfTestSuitesInTheDatabaseSearch:            nil,
		sortableHeaderReference:                          nil,
		currentRowThatMouseIsHoveringAbove:               -1,
		currentRowThatMouseIsHoveringAboveMutex:          sync.Mutex{},
		testSuitePreviewContainerScroll:                  nil,
		testSuitePreviewContainer:                        nil,
		preViewAndFilterTabs:                             nil,
		preViewTab:                                       nil,
		filterTab:                                        nil,
		simpleTestSuiteMetaDataSelectedDomainUuid:        "",
		simpleTestSuiteMetaDataSelectedDomainDisplayName: "",
		testSuiteFullMetaDataFilterContainer:             nil,
		testSuiteMainAreaForMetaDataFilterContainer:      nil,
		newMandatoryOwnerDomainSelect:                    nil,
		filterOnMetaDataFunction:                         nil,
		calculateMetaDataFilterFunction:                  nil,
		useAutoFilter:                                    false,
		simpleMetaDataFilterEntryMap:                     nil,
		testSuiteThatIsShownInPreview:                    "",
		currentSortColumn:                                -1,
		currentHeader:                                    nil,
		previousHeader:                                   nil,
		selectedTestCasesPtr:                             selectedTestSuitesPtr,
		//testSuiteListTableHeader:                       ???
	}

	// Handle Headers for TestCaseList by adding first column for when TestCase is selected
	if tempHowShouldItBeUsed == UsedForTestSuitesList {
		listTestCaseUIObject.testSuiteListTableHeader = testSuiteListTableHeaderForTestSuitesList

	} else {
		listTestCaseUIObject.testSuiteListTableHeader = testSuiteListTableHeaderForTestSuiteBuilder
	}

	return listTestCaseUIObject

}

// Create the UI used for list all TestSuitesMapPtr that the User can edit
func (listTestSuiteUIObject *ListTestSuiteUIStruct) GenerateListTestSuitesUI(
	testCasesModel *testCaseModel.TestCasesModelsStruct,
	preViewAndFilterTabsUsedForCreateTestSuite *container.AppTabs) (
	_ *fyne.Container) {

	//var testSuiteTable *widget.Table

	var tempTestSuiteListAndTestSuitePreviewSplitContainer *container.Split

	var testSuitesListContainer *fyne.Container
	var testSuitesListScrollContainer *container.Scroll
	var statisticsContainer *fyne.Container
	var executionColorPaletteContainer *fyne.Container
	var statisticsAndColorPaletteContainer *fyne.Container

	var loadTestSuiteFromDataBaseButton *widget.Button
	var loadTestSuiteFromDataBaseFunction func()
	var filterTestSuitesButton *widget.Button
	var filterTestSuitesButtonFunction func()
	var clearFiltersButton *widget.Button
	var clearFiltersButtonFunction func()
	var buttonsContainer *fyne.Container

	var numberOfTestSuitesAfterLocalFilterLabel *widget.Label
	var numberOfTestSuitesRetrievedFromDatabaseLabel *widget.Label

	var filterAndButtonsContainer *fyne.Container

	// Define the function to be executed to load TestSuitesMapPtr from that Database that the user can edit
	loadTestSuiteFromDataBaseFunction = func() {
		fmt.Println("'loadTestSuiteFromDataBaseButton' was pressed")
		listTestSuitesModel.LoadtestSuiteThatCanBeEditedByUser(
			testCasesModel,
			time.Now().Add(-time.Hour*10000), time.Now().Add(-time.Hour*10000))
		filterTestSuitesButtonFunction()
		//sortTestSuitesTable()
	}

	// Define the 'loadTestSuiteFromDataBaseButton'
	loadTestSuiteFromDataBaseButton = widget.NewButton("Load TestSuitesMapPtr from Database", loadTestSuiteFromDataBaseFunction)

	// Define the function to be executed to filter TestSuitesMapPtr that the user can edit
	filterTestSuitesButtonFunction = func() {
		fmt.Println("'filterTestSuitesButton' was pressed")
		listTestSuiteUIObject.loadTestSuiteListTableTable(nil)
		listTestSuiteUIObject.calculateAndSetCorrectColumnWidths()
		listTestSuiteUIObject.updateTestSuitesListTable(testCasesModel)

		// Update the number TestSuitesMapPtr in the list
		var numberOfRowsAsString string
		numberOfRowsAsString = strconv.Itoa(len(listTestSuiteUIObject.testSuitesListTableTable))
		listTestSuiteUIObject.numberOfTestSuitesAfterLocalFilters.Set(
			fmt.Sprintf("Number of TestSuitesMapPtr after local filters was applied: %s",
				numberOfRowsAsString))

		// Update the number TestSuitesMapPtr retrieved from Database
		var numberOfRowsFromDatabaseAsString string
		numberOfRowsFromDatabaseAsString = strconv.Itoa(len(listTestSuiteUIObject.testSuitesListTableTable))
		listTestSuiteUIObject.numberOfTestSuitesInTheDatabaseSearch.Set(
			fmt.Sprintf("Number of TestSuitesMapPtr retrieved from the Database: %s",
				numberOfRowsFromDatabaseAsString))

		listTestSuiteUIObject.sortableHeaderReference.sortImage.onTapped()

	}

	// Define the 'filterTestSuitesButton'
	filterTestSuitesButton = widget.NewButton("Filter TestSuitesMapPtr", filterTestSuitesButtonFunction)

	// Define the function to be executed to list TestSuitesMapPtr that the user can edit
	clearFiltersButtonFunction = func() {
		fmt.Println("'clearFiltersButtonFunction' was pressed")
	}

	// Define the 'filterTestSuitesButton'
	clearFiltersButton = widget.NewButton("Clear all search filters", clearFiltersButtonFunction)

	// Add the buttons to the buttonsContainer
	buttonsContainer = container.NewHBox(loadTestSuiteFromDataBaseButton, filterTestSuitesButton, clearFiltersButton)

	// Add objects to the 'filterAndButtonsContainer'
	filterAndButtonsContainer = container.NewVBox(buttonsContainer)

	// Generate the ExecutionColorPaletteContainer
	executionColorPaletteContainer = detailedTestCaseExecutionsUI.GenerateExecutionColorPalette()

	// Initiate the Table
	listTestSuiteUIObject.generateTestSuitesListTable(testCasesModel)
	testSuiteTableContainer := container.NewBorder(nil, nil, nil, nil, listTestSuiteUIObject.testSuiteListTable)

	// Create the Scroll container for the List
	testSuitesListScrollContainer = container.NewScroll(testSuiteTableContainer)

	// Create the label used for showing number of TestSuitesMapPtr in the local filter
	listTestSuiteUIObject.numberOfTestSuitesAfterLocalFilters = binding.NewString()
	_ = listTestSuiteUIObject.numberOfTestSuitesAfterLocalFilters.Set("No TestSuitesMapPtr in the List")
	numberOfTestSuitesAfterLocalFilterLabel = widget.NewLabelWithData(listTestSuiteUIObject.numberOfTestSuitesAfterLocalFilters)

	// Create the label used for showing number of TestSuitesMapPtr retrieved from the Database
	listTestSuiteUIObject.numberOfTestSuitesInTheDatabaseSearch = binding.NewString()
	_ = listTestSuiteUIObject.numberOfTestSuitesInTheDatabaseSearch.Set("No TestSuitesMapPtr retrieved from the Database")
	numberOfTestSuitesRetrievedFromDatabaseLabel = widget.NewLabelWithData(listTestSuiteUIObject.numberOfTestSuitesInTheDatabaseSearch)

	// Initiate 'statisticsContainer'
	statisticsContainer = container.NewHBox(numberOfTestSuitesAfterLocalFilterLabel, numberOfTestSuitesRetrievedFromDatabaseLabel)

	statisticsAndColorPaletteContainer = container.NewVBox(executionColorPaletteContainer, statisticsContainer)

	// Add 'testSuitesListScrollContainer' to 'testSuitesListContainer'
	testSuitesListContainer = container.NewBorder(filterAndButtonsContainer, statisticsAndColorPaletteContainer, nil, nil, testSuitesListScrollContainer)
	testSuitesListScrollContainer2 := container.NewScroll(testSuitesListContainer)

	// Create the Temporary container that should be shown
	temporaryContainer := container.NewCenter(widget.NewLabel("Select a TestSuite to get the Preview"))

	listTestSuiteUIObject.testSuitePreviewContainer = container.NewBorder(nil, nil, nil, nil, temporaryContainer)

	// Generate the container for the Preview, 'testSuitePreviewContainer'
	listTestSuiteUIObject.testSuitePreviewContainerScroll = container.NewScroll(listTestSuiteUIObject.testSuitePreviewContainer)

	// Generate the Tab for the PreView
	listTestSuiteUIObject.preViewTab = container.NewTabItem(
		"PreView",
		listTestSuiteUIObject.testSuitePreviewContainerScroll)

	// Generate the 'GenerateTestSuiteMetaDataFilterContainer'
	var simpleAndAdvancedMetaDataFilter *container.AppTabs
	simpleAndAdvancedMetaDataFilter = listTestSuiteUIObject.GenerateTestSuiteMetaDataFilterContainer(testCasesModel)

	// Generate Tab for TestSuite-filter
	listTestSuiteUIObject.filterTab = container.NewTabItem(
		"TestSuite-filter",
		simpleAndAdvancedMetaDataFilter)

	// Generate the AppTabsContainer, depending on from where it was initiated
	if listTestSuiteUIObject.howShouldItBeUsed == UsedForTestSuitesList {
		// We are in standard List TestSuites
		listTestSuiteUIObject.preViewAndFilterTabs = container.NewAppTabs(listTestSuiteUIObject.filterTab, listTestSuiteUIObject.preViewTab)

		// make a hoverable transparent PreView-overlay, to stop table-row hovering in left Table
		preViewOverlay := NewHoverableRect(color.Transparent, nil)
		preViewOverlay.OnMouseIn = func(ev *desktop.MouseEvent) {

			mouseHasLeftTable = true
			preViewOverlay.Hide()
			preViewOverlay.OtherHoverableRect.Show()
		}
		preViewOverlay.OnMouseOut = func() {

		}

		// make a hoverable transparent TestCaseList-overlay, to stop table-row hovering in left Table
		testSuiteListViewOverlay := NewHoverableRect(color.Transparent, nil)
		testSuiteListViewOverlay.OnMouseIn = func(ev *desktop.MouseEvent) {

			mouseHasLeftTable = false
			testSuiteListViewOverlay.Hide()
			testSuiteListViewOverlay.OtherHoverableRect.Show()
		}
		testSuiteListViewOverlay.OnMouseOut = func() {

		}

		// Cross connect the two overlays
		preViewOverlay.OtherHoverableRect = testSuiteListViewOverlay
		testSuiteListViewOverlay.OtherHoverableRect = preViewOverlay

		preViewAndOverlayContainer := container.New(layout.NewStackLayout(), listTestSuiteUIObject.preViewAndFilterTabs, preViewOverlay)
		testSuiteListAndOverlayContainer := container.New(layout.NewStackLayout(), testSuitesListScrollContainer2, testSuiteListViewOverlay)

		tempTestSuiteListAndTestSuitePreviewSplitContainer = container.NewHSplit(
			testSuiteListAndOverlayContainer,
			preViewAndOverlayContainer)
		tempTestSuiteListAndTestSuitePreviewSplitContainer.Offset = 0.75

		//TestSuiteListAndTestSuitePreviewSplitContainer = tempTestSuiteListAndTestSuitePreviewSplitContainer

		return container.NewBorder(nil, nil, nil, nil, tempTestSuiteListAndTestSuitePreviewSplitContainer)

	} else {
		// We are in Create TestSuite
		listTestSuiteUIObject.preViewAndFilterTabs = preViewAndFilterTabsUsedForCreateTestSuite
		listTestSuiteUIObject.preViewAndFilterTabs.Append(listTestSuiteUIObject.filterTab)
		listTestSuiteUIObject.preViewAndFilterTabs.Append(listTestSuiteUIObject.preViewTab)

		testCasesListContainer2 := container.NewBorder(nil, nil, nil, nil, testSuitesListScrollContainer2)

		return testCasesListContainer2

	}

}
