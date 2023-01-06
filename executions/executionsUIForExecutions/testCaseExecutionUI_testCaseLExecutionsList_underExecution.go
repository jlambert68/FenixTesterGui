package executionsUIForExecutions

import (
	"FenixTesterGui/executions/executionsModelForSubscriptions"
	"FenixTesterGui/headertable"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
)

func CreateTableForTestCaseExecutionsUnderExecution() *fyne.Container {
	var tableForTestCaseExecutionsUnderExecutionBindings []binding.DataMap

	// Create a binding for each TestExecutionUnderExecutionRow data
	/*
		for testDataRowCounter := 0; testDataRowCounter < len(executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionMapAdaptedForUiTable); testDataRowCounter++ {
			tableForTestCaseExecutionsUnderExecutionBindings = append(
				tableForTestCaseExecutionsUnderExecutionBindings,
				binding.BindStruct(&executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionMapAdaptedForUiTable[testDataRowCounter]))
		}
	*/

	for _, tempTestCaseExecutionsUnderExecutionDataAdaptedForUiTableReference := range executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionMapAdaptedForUiTable {
		tableForTestCaseExecutionsUnderExecutionBindings = append(
			tableForTestCaseExecutionsUnderExecutionBindings,
			binding.BindStruct(tempTestCaseExecutionsUnderExecutionDataAdaptedForUiTableReference))

	}

	executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionTableOptions.Bindings = tableForTestCaseExecutionsUnderExecutionBindings

	ht := headertable.NewSortingHeaderTable(&executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionTableOptions)
	ExecutionsUIObject.UnderExecutionTable = ht

	mySortTable := container.NewMax(ht)

	return mySortTable

}
