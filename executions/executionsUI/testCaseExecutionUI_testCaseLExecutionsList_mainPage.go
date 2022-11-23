package executionsUI

import (
	"fyne.io/fyne/v2"
)

func (executionsUIObject *ExecutionsUIModelStruct) CreateExecutionsListTabPage() (executionsListTabPage *fyne.Container) {

	executionsListTabPage = newThreePartAdaptiveSplit(CreateTableForTestCaseExecutionsOnQueue(), MySortTable(), MySortTable())

	return executionsListTabPage
}
