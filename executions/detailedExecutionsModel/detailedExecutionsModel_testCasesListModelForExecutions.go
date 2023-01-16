package detailedExecutionsModel

import (
	"FenixTesterGui/executions/detailedTestCaseExecutionUI_summaryTableDefinition"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
)

func CreateSummaryTableForDetailedTestCaseExecutionsList() *fyne.Container {
	var tableForTestCaseExecutionsSummaryBindings []binding.DataMap

	// Create a binding for each TestCaseExecutionsSummary data
	for _, tempDetailedTestCaseExecutionReference := range TestCaseExecutionsDetailsMap {

		tableForTestCaseExecutionsSummaryBindings = append(
			tableForTestCaseExecutionsSummaryBindings,
			binding.BindStruct(tempDetailedTestCaseExecutionReference.TestCaseExecutionsStatusForSummaryTable))
	}

	DetailedTestCaseExecutionsSummaryTableOptions.Bindings = tableForTestCaseExecutionsSummaryBindings

	ht := detailedTestCaseExecutionUI_summaryTableDefinition.NewTestCaseExecutionsSummaryTable(&DetailedTestCaseExecutionsSummaryTableOptions)
	detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsSummaryTable = ht

	mySortTable := container.NewMax(ht, layout.NewSpacer())
	mySortTable.Resize(ht.Size())
	//ht.Header.ScrollToTrailing()
	ht.Data.Refresh()
	//ht.Header.ScrollToLeading()

	return mySortTable

}
