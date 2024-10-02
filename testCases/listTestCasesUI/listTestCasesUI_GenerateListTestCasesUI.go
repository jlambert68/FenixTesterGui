package listTestCasesUI

import (
	"FenixTesterGui/testCase/testCaseModel"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

// Create the UI used for list all TestCases that the User can edit
func GenerateListTestCasesUI(testCasesModel testCaseModel.TestCasesModelsStruct) (listTestCasesUI fyne.CanvasObject) {

	//var testCaseTable *widget.Table

	var testCasesListContainer *fyne.Container
	var testCasesListScrollContainer *container.Scroll
	var statisticsContainer *fyne.Container

	var listTestCasesButton *widget.Button
	var listTestCasesButtonFunction func()
	var clearFiltersButton *widget.Button
	var clearFiltersButtonFunction func()
	var buttonsContainer *fyne.Container

	var numberOfTestCasesInTheSearchLabel *widget.Label

	var filterAndButtonsContainer *fyne.Container

	// Define the function to be executed to list TestCases that the user can edit
	listTestCasesButtonFunction = func() {
		fmt.Println("'listTestCasesButtonFunction' was pressed")
		loadTestCaseListTableTable(testCasesModel)
		calculateAndSetCorrectColumnWidths()
		updateTestCasesListTable()

		// Update the number TestCases in the list
		numberOfRowsAsString := strconv.Itoa(len(testCaseListTableTable))
		numberOfTestCasesInTheSearch.Set(
			fmt.Sprintf("Number of TestCases in the Search: %s",
				numberOfRowsAsString))

	}

	// Define the 'listTestCasesButton'
	listTestCasesButton = widget.NewButton("List TestCases", listTestCasesButtonFunction)

	// Define the function to be executed to list TestCases that the user can edit
	clearFiltersButtonFunction = func() {
		fmt.Println("'clearFiltersButtonFunction' was pressed")
	}

	// Define the 'listTestCasesButton'
	clearFiltersButton = widget.NewButton("Clear all search filters", clearFiltersButtonFunction)

	// Add the buttons to the buttonsContainer
	buttonsContainer = container.NewHBox(listTestCasesButton, clearFiltersButton)

	// Add objects to the 'filterAndButtonsContainer'
	filterAndButtonsContainer = container.NewVBox(buttonsContainer)

	// Initiate the Table
	generateTestCasesListTable()
	testCaseTableContainer := container.NewBorder(nil, nil, nil, nil, testCaseListTable)

	// Create the Scroll container for the List
	testCasesListScrollContainer = container.NewScroll(testCaseTableContainer)

	// Create the label used for showing number of TestCases in the Search
	numberOfTestCasesInTheSearch = binding.NewString()
	err := numberOfTestCasesInTheSearch.Set("No TestCases in the List")
	fmt.Println(err)
	numberOfTestCasesInTheSearchLabel = widget.NewLabelWithData(numberOfTestCasesInTheSearch)

	// Initiate 'statisticsContainer'
	statisticsContainer = container.NewHBox(numberOfTestCasesInTheSearchLabel)

	// Add 'testCasesListScrollContainer' to 'testCasesListContainer'
	testCasesListContainer = container.NewBorder(filterAndButtonsContainer, statisticsContainer, nil, nil, testCasesListScrollContainer)

	return testCasesListContainer
}
