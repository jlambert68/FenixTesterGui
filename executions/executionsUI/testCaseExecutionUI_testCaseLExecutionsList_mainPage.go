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

// StartTableAddAndRemoveChannelReaders
// Start Channel readers for testCases OnQueue, UnderExecutions or Finished Executions
func StartTableAddAndRemoveChannelReaders() {

	// Start Channel-reader used for Adding and Deleting Execution items in OnQueue-table
	go StartOnQueueTableAddRemoveChannelReader()

}
