package executionsUIForExecutions

import (
	"FenixTesterGui/executions/executionsModelForTestCaseExecutions"
	"FenixTesterGui/headertable"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
)

// CreateTableForTestCaseExecutionsWithFinishedExecution
// Create bindings to the data used by the table and then create the UI-table itself
func CreateTableForTestCaseExecutionsWithFinishedExecution() *fyne.Container {
	var tableForTestCaseExecutionsWithFinishedExecutionBindings []binding.DataMap

	// Create a binding for each TestExecutionWithFinishedExecutionRow data
	for _, tempTestCaseExecutionsFinishedExecutionDataAdaptedForUiTableReference := range executionsModelForExecutions.TestCaseExecutionsFinishedExecutionMapAdaptedForUiTable {
		tableForTestCaseExecutionsWithFinishedExecutionBindings = append(
			tableForTestCaseExecutionsWithFinishedExecutionBindings,
			binding.BindStruct(tempTestCaseExecutionsFinishedExecutionDataAdaptedForUiTableReference))
	}

	executionsModelForExecutions.TestCaseExecutionsFinishedExecutionTableOptions.Bindings = tableForTestCaseExecutionsWithFinishedExecutionBindings

	ht := headertable.NewSortingHeaderTable(&executionsModelForExecutions.TestCaseExecutionsFinishedExecutionTableOptions)
	ExecutionsUIObject.FinishedExecutionTable = ht

	mySortTable := container.NewMax(ht)

	return mySortTable

}
