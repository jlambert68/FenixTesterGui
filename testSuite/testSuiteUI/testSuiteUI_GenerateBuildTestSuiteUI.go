package testSuiteUI

import (
	"FenixTesterGui/testCase/testCaseModel"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// GenerateBuildTestSuiteUI
// Create the UI used for creating new TestSuites
func GenerateBuildTestSuiteUI(testCasesModel *testCaseModel.TestCasesModelsStruct) (listTestSuiteUI fyne.CanvasObject) {

	listTestSuiteUI = container.NewVBox(widget.NewLabel("List of TestSuites"))

	return listTestSuiteUI
}
