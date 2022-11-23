package executionsUI

import (
	"fyne.io/fyne/v2"
)

func (executionsUIObject *ExecutionsUIModelStruct) CreateExecutionsListTabPage() (executionsListTabPage *fyne.Container) {

	executionsListTabPage = newThreePartAdaptiveSplit(
		CreateTableForTestCaseExecutionsOnQueue(),
		CreateTableForTestCaseExecutionsUnderExecution(),
		MySortTable())

	return executionsListTabPage
}
