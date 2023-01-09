package detailedTestCaseExecutionsUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (detailedTestCaseExecutionsUIObject *DetailedTestCaseExecutionsUIModelStruct) CreateDetailedTestCaseExecutionsTabPage() (detailedTestCaseExecutionsTabPage *fyne.Container) {
	newLabel := widget.NewLabel("Detailed TestCaseExecution")

	detailedTestCaseExecutionsTabPage = container.New(layout.NewVBoxLayout(), layout.NewSpacer(), newLabel, layout.NewSpacer())

	return detailedTestCaseExecutionsTabPage
}
