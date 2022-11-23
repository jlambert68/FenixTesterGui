package executionsUI

import (
	"FenixTesterGui/executions/executionsModel"
	"FenixTesterGui/headertable"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
)

func CreateTableForTestCaseExecutionsUnderExecution() *fyne.Container {
	var tableForTestCaseExecutionsUnderExecutionBindings []binding.DataMap

	// Create a binding for each TestExecutionUnderExecutionRow data
	for testDataRowCounter := 0; testDataRowCounter < len(executionsModel.TestCaseExecutionsUnderExecutionAdaptedForUiTable); testDataRowCounter++ {
		tableForTestCaseExecutionsUnderExecutionBindings = append(tableForTestCaseExecutionsUnderExecutionBindings, binding.BindStruct(&executionsModel.TestCaseExecutionsUnderExecutionAdaptedForUiTable[testDataRowCounter]))
	}
	executionsModel.TestCaseExecutionsUnderExecutionTableOptions.Bindings = tableForTestCaseExecutionsUnderExecutionBindings

	ht := headertable.NewSortingHeaderTable(&executionsModel.TestCaseExecutionsUnderExecutionTableOptions)
	mySortTable := container.NewMax(ht)

	return mySortTable

}
