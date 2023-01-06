package executionsUIForExecutions

import (
	"FenixTesterGui/executions/executionsModelForSubscriptions"
	"FenixTesterGui/headertable"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
)

func CreateTableForTestCaseExecutionsOnQueue() *fyne.Container {
	var tableForTestCaseExecutionsOnQueueBindings []binding.DataMap

	// Create a binding for each TestExecutionOnQueueRow data
	for _, tempTestCaseExecutionsOnQueueDataAdaptedForUiTableReference := range executionsModelForSubscriptions.TestCaseExecutionsOnQueueMapAdaptedForUiTable {
		tableForTestCaseExecutionsOnQueueBindings = append(
			tableForTestCaseExecutionsOnQueueBindings,
			binding.BindStruct(tempTestCaseExecutionsOnQueueDataAdaptedForUiTableReference))
	}

	executionsModelForSubscriptions.TestCaseExecutionsOnQueueTableOptions.Bindings = tableForTestCaseExecutionsOnQueueBindings

	ht := headertable.NewSortingHeaderTable(&executionsModelForSubscriptions.TestCaseExecutionsOnQueueTableOptions)
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
