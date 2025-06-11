package testSuiteUI

import (
	"FenixTesterGui/testCase/testCaseModel"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func generateRightSideBuildTestSuiteContainer(
	testCasesModel *testCaseModel.TestCasesModelsStruct) (
	rightSideBuildTestSuiteContainer *fyne.Container) {

	var rightSideScrollContainer *fyne.Container
	rightSideScrollContainer = container.NewCenter(widget.NewLabel("rightSideScrollContainer"))

	rightSideBuildTestSuiteContainer = container.NewVBox(rightSideScrollContainer)

	return rightSideBuildTestSuiteContainer

}
