package executionsUI

import (
	"fyne.io/fyne/v2"
)

func (executionsUIObject *ExecutionsUIModelStruct) CreateExecutionsListTabPage() (executionsListTabPage *fyne.Container) {

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
