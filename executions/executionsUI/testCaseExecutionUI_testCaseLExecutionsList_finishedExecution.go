package executionsUI

import (
	"FenixTesterGui/executions/executionsModel"
	"FenixTesterGui/headertable"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
)

func CreateTableForTestCaseExecutionsWithFinishedExecution() *fyne.Container {
	var tableForTestCaseExecutionsWithFinishedExecutionBindings []binding.DataMap

	// Create a binding for each TestExecutionWithFinishedExecutionRow data
	for testDataRowCounter := 0; testDataRowCounter < len(executionsModel.TestCaseExecutionsFinishedExecutionAdaptedForUiTable); testDataRowCounter++ {
		tableForTestCaseExecutionsWithFinishedExecutionBindings = append(tableForTestCaseExecutionsWithFinishedExecutionBindings, binding.BindStruct(&executionsModel.TestCaseExecutionsFinishedExecutionAdaptedForUiTable[testDataRowCounter]))
	}
	executionsModel.TestCaseExecutionsFinishedExecutionTableOptions.Bindings = tableForTestCaseExecutionsWithFinishedExecutionBindings

	ht := headertable.NewSortingHeaderTable(&executionsModel.TestCaseExecutionsFinishedExecutionTableOptions)
	mySortTable := container.NewMax(ht)

	return mySortTable

}
