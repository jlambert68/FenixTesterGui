package listTestCasesUI

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Create the UI used for list all TestCases that the User can edit
func LoadListTestCasesUI() (listTestCasesUI fyne.CanvasObject) {

	//var testCaseTable *widget.Table

	var testCasesListContainer *fyne.Container
	var testCasesListScrollContainer *container.Scroll

	var listTestCasesButton *widget.Button
	var listTestCasesButtonFunction func()
	var clearFiltersButton *widget.Button
	var clearFiltersButtonFunction func()
	var buttonsContainer *fyne.Container

	var filterAndButtonsContainer *fyne.Container

	// Define the function to be executed to list TestCases that the user can edit
	listTestCasesButtonFunction = func() {
		fmt.Println("'listTestCasesButtonFunction' was pressed")
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
	//testCaseTable = widget.NewTable()
	testCaseTableContainer := container.NewVBox(widget.NewLabel("My table"))

	// Create the Scroll container for the List
	testCasesListScrollContainer = container.NewScroll(testCaseTableContainer)

	// Add 'testCasesListScrollContainer' to 'testCasesListContainer'
	testCasesListContainer = container.NewBorder(filterAndButtonsContainer, nil, nil, nil, testCasesListScrollContainer)

	return testCasesListContainer
}
