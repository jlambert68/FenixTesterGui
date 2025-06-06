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
	"image"
	"strconv"
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

// Create the UI used for list all TestCasesMapPtr that the User can edit
func GenerateListTestCasesUI(testCasesModel *testCaseModel.TestCasesModelsStruct) (listTestCasesUI fyne.CanvasObject) {

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
			time.Now().Add(-time.Hour*1000), time.Now().Add(-time.Hour*1000),
			testCasesModel)
		filterTestCasesButtonFunction()
	}

	// Define the 'loadTestCaseFromDataBaseButton'
	loadTestCaseFromDataBaseButton = widget.NewButton("Load TestCasesMapPtr from Database", loadTestCaseFromDataBaseFunction)

	// Define the function to be executed to filter TestCasesMapPtr that the user can edit
	filterTestCasesButtonFunction = func() {
		fmt.Println("'filterTestCasesButton' was pressed")
		loadTestCaseListTableTable(testCasesModel)
		calculateAndSetCorrectColumnWidths()
		updateTestCasesListTable(testCasesModel)

		// Update the number TestCasesMapPtr in the list
		var numberOfRowsAsString string
		numberOfRowsAsString = strconv.Itoa(len(testCaseListTableTable))
		numberOfTestCasesAfterLocalFilters.Set(
			fmt.Sprintf("Number of TestCasesMapPtr after local filters was applied: %s",
				numberOfRowsAsString))

		// Update the number TestCasesMapPtr retrieved from Database
		var numberOfRowsFromDatabaseAsString string
		numberOfRowsFromDatabaseAsString = strconv.Itoa(len(testCaseListTableTable))
		numberOfTestCasesInTheDatabaseSearch.Set(
			fmt.Sprintf("Number of TestCasesMapPtr retrieved from the Database: %s",
				numberOfRowsFromDatabaseAsString))

		sortableHeaderReference.sortImage.onTapped()

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
	generateTestCasesListTable(testCasesModel)
	testCaseTableContainer := container.NewBorder(nil, nil, nil, nil, testCaseListTable)

	// Create the Scroll container for the List
	testCasesListScrollContainer = container.NewScroll(testCaseTableContainer)

	// Create the label used for showing number of TestCasesMapPtr in the local filter
	numberOfTestCasesAfterLocalFilters = binding.NewString()
	_ = numberOfTestCasesAfterLocalFilters.Set("No TestCasesMapPtr in the List")
	numberOfTestCasesAfterLocalFilterLabel = widget.NewLabelWithData(numberOfTestCasesAfterLocalFilters)

	// Create the label used for showing number of TestCasesMapPtr retrieved from the Database
	numberOfTestCasesInTheDatabaseSearch = binding.NewString()
	_ = numberOfTestCasesInTheDatabaseSearch.Set("No TestCasesMapPtr retrieved from the Database")
	numberOfTestCasesRetrievedFromDatabaseLabel = widget.NewLabelWithData(numberOfTestCasesInTheDatabaseSearch)

	// Initiate 'statisticsContainer'
	statisticsContainer = container.NewHBox(numberOfTestCasesAfterLocalFilterLabel, numberOfTestCasesRetrievedFromDatabaseLabel)

	statisticsAndColorPaletteContainer = container.NewVBox(executionColorPaletteContainer, statisticsContainer)

	// Add 'testCasesListScrollContainer' to 'testCasesListContainer'
	testCasesListContainer = container.NewBorder(filterAndButtonsContainer, statisticsAndColorPaletteContainer, nil, nil, testCasesListScrollContainer)
	testCasesListScrollContainer2 := container.NewScroll(testCasesListContainer)

	// Create the Temporary container that should be shown
	temporaryContainer := container.NewCenter(widget.NewLabel("Select a TestCase to get the Preview"))

	testCasePreviewContainer = container.NewBorder(nil, nil, nil, nil, temporaryContainer)

	// Generate the container for the Preview, 'testCasePreviewContainer'
	testCasePreviewContainerScroll = container.NewScroll(testCasePreviewContainer)

	// Generate the Tab for the PreView
	preViewTab = container.NewTabItem(
		"PreView",
		testCasePreviewContainerScroll)

	// Generate the 'GenerateTestCaseMetaDataFilterContainer'
	var simpleAndAdvancedMetaDataFilter *container.AppTabs
	simpleAndAdvancedMetaDataFilter = GenerateTestCaseMetaDataFilterContainer(testCasesModel)

	// Generate Tab for TestCase-filter
	filterTab = container.NewTabItem(
		"TestCase-filter",
		simpleAndAdvancedMetaDataFilter)

	// Generate the AppTabsContainer
	preViewAndFilterTabs = container.NewAppTabs(filterTab, preViewTab)

	tempTestCaseListAndTestCasePreviewSplitContainer = container.NewHSplit(testCasesListScrollContainer2, preViewAndFilterTabs)
	tempTestCaseListAndTestCasePreviewSplitContainer.Offset = 0.75

	TestCaseListAndTestCasePreviewSplitContainer = tempTestCaseListAndTestCasePreviewSplitContainer

	return tempTestCaseListAndTestCasePreviewSplitContainer
}
