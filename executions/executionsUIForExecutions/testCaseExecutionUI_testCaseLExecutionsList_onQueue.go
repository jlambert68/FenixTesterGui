package executionsUIForExecutions

import (
	"FenixTesterGui/executions/executionsModelForTestCaseExecutions"
	"FenixTesterGui/headertable"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
)

func CreateTableForTestCaseExecutionsOnQueue() *fyne.Container {
	var tableForTestCaseExecutionsOnQueueBindings []binding.DataMap

	// Create a binding for each TestExecutionOnQueueRow data
	for _, tempTestCaseExecutionsOnQueueDataAdaptedForUiTableReference := range executionsModelForExecutions.TestCaseExecutionsOnQueueMapAdaptedForUiTable {
		tableForTestCaseExecutionsOnQueueBindings = append(
			tableForTestCaseExecutionsOnQueueBindings,
			binding.BindStruct(tempTestCaseExecutionsOnQueueDataAdaptedForUiTableReference))
	}

	executionsModelForExecutions.TestCaseExecutionsOnQueueTableOptions.Bindings = tableForTestCaseExecutionsOnQueueBindings

	ht := headertable.NewSortingHeaderTable(&executionsModelForExecutions.TestCaseExecutionsOnQueueTableOptions)
	ExecutionsUIObject.OnQueueTable = ht

	mySortTable := container.NewMax(ht)

	/*
		first := container.NewHBox(widget.NewLabel("Första"), mySortTable, widget.NewLabel("Sista"))

		second := container.NewHBox(first)

		scrollableTable := container.NewScroll(second)

		scrollableTableContainer := container.NewMax(scrollableTable)
	*/
	//ht.Header.ScrollToTrailing()
	ht.Header.Refresh()
	//ht.Header.ScrollToLeading()

	return mySortTable

}
