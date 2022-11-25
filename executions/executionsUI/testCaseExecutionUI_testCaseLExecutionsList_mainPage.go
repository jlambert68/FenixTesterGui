package executionsUI

import (
	"fyne.io/fyne/v2"
)

func (executionsUIObject *ExecutionsUIModelStruct) CreateExecutionsListTabPage() (executionsListTabPage *fyne.Container) {

	underExecutionTableObject := CreateTableForTestCaseExecutionsUnderExecution()

	executionsListTabPage = newThreePartAdaptiveSplit(
		CreateTableForTestCaseExecutionsOnQueue(),
		underExecutionTableObject,
		CreateTableForTestCaseExecutionsWithFinishedExecution())

	executionsListTabPage.Refresh()

	return executionsListTabPage
}
