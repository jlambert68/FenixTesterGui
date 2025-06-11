package testSuitesListUI

import (
	"FenixTesterGui/testCase/testCaseModel"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// GenerateListTestSuitesUI
// Create the UI used for list existing TestSuites
func GenerateListTestSuitesUI(testCasesModel *testCaseModel.TestCasesModelsStruct) (listTestSuiteUI fyne.CanvasObject) {

	listTestSuiteUI = container.NewVBox(widget.NewLabel("List of TestSuites"))

	return listTestSuiteUI
}
