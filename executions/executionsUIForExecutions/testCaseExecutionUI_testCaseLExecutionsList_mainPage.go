package executionsUIForExecutions

import (
	"fyne.io/fyne/v2"
)

func (executionsUIObject *ExecutionsUIModelStruct) CreateExecutionsListTabPageForSubsacriptions() (executionsListTabPage *fyne.Container) {

	onQueueTableObject := CreateTableForTestCaseExecutionsOnQueue()
	underExecutionTableObject := CreateTableForTestCaseExecutionsUnderExecution()
	finishedExecutionTableObject := CreateTableForTestCaseExecutionsWithFinishedExecution()

	executionsListTabPage = newThreePartAdaptiveSplit(
		onQueueTableObject,
		underExecutionTableObject,
		finishedExecutionTableObject)

	executionsListTabPage.Refresh()

	return executionsListTabPage
}
