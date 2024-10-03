package listTestCasesUI

import (
	detailedTestCaseExecutionsUI "FenixTesterGui/executions/detailedExecutionsUI"
	"FenixTesterGui/testCase/testCaseModel"
	"FenixTesterGui/testCases/listTestCasesModel"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

// Create the UI used for list all TestCases that the User can edit
func GenerateListTestCasesUI(testCasesModel *testCaseModel.TestCasesModelsStruct) (listTestCasesUI fyne.CanvasObject) {

	//var testCaseTable *widget.Table

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

	// Define the function to be executed to load TestCases from that Database that the user can edit
	loadTestCaseFromDataBaseFunction = func() {
		fmt.Println("'loadTestCaseFromDataBaseButton' was pressed")
		listTestCasesModel.LoadTestCaseThatCanBeEditedByUser(testCasesModel)
		filterTestCasesButtonFunction()
	}

	// Define the 'loadTestCaseFromDataBaseButton'
	loadTestCaseFromDataBaseButton = widget.NewButton("Load TestCases from Database", loadTestCaseFromDataBaseFunction)

	// Define the function to be executed to filter TestCases that the user can edit
	filterTestCasesButtonFunction = func() {
		fmt.Println("'filterTestCasesButton' was pressed")
		loadTestCaseListTableTable(testCasesModel)
		calculateAndSetCorrectColumnWidths()
		updateTestCasesListTable(testCasesModel)

		// Update the number TestCases in the list
		var numberOfRowsAsString string
		numberOfRowsAsString = strconv.Itoa(len(testCaseListTableTable))
		numberOfTestCasesAfterLocalFilters.Set(
			fmt.Sprintf("Number of TestCases after local filters was applied: %s",
				numberOfRowsAsString))

		// Update the number TestCases retrieved from Database
		var numberOfRowsFromDatabaseAsString string
		numberOfRowsFromDatabaseAsString = strconv.Itoa(len(testCaseListTableTable))
		numberOfTestCasesInTheDatabaseSearch.Set(
			fmt.Sprintf("Number of TestCases retrieved from the Database: %s",
				numberOfRowsFromDatabaseAsString))

	}

	// Define the 'filterTestCasesButton'
	filterTestCasesButton = widget.NewButton("Filter TestCases", filterTestCasesButtonFunction)

	// Define the function to be executed to list TestCases that the user can edit
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

	// Create the label used for showing number of TestCases in the local filter
	numberOfTestCasesAfterLocalFilters = binding.NewString()
	_ = numberOfTestCasesAfterLocalFilters.Set("No TestCases in the List")
	numberOfTestCasesAfterLocalFilterLabel = widget.NewLabelWithData(numberOfTestCasesAfterLocalFilters)

	// Create the label used for showing number of TestCases retrieved from the Database
	numberOfTestCasesInTheDatabaseSearch = binding.NewString()
	_ = numberOfTestCasesInTheDatabaseSearch.Set("No TestCases retrieved from the Database")
	numberOfTestCasesRetrievedFromDatabaseLabel = widget.NewLabelWithData(numberOfTestCasesInTheDatabaseSearch)

	// Initiate 'statisticsContainer'
	statisticsContainer = container.NewHBox(numberOfTestCasesAfterLocalFilterLabel, numberOfTestCasesRetrievedFromDatabaseLabel)

	statisticsAndColorPaletteContainer = container.NewVBox(executionColorPaletteContainer, statisticsContainer)

	// Add 'testCasesListScrollContainer' to 'testCasesListContainer'
	testCasesListContainer = container.NewBorder(filterAndButtonsContainer, statisticsAndColorPaletteContainer, nil, nil, testCasesListScrollContainer)

	return testCasesListContainer
}
