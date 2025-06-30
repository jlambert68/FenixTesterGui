package testSuiteUI

import (
	"FenixTesterGui/testCase/testCaseModel"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func generateRightSideBuildTestSuiteContainer(
	testCasesModel *testCaseModel.TestCasesModelsStruct) (
	rightSideBuildTestSuiteContainer *fyne.Container,
	preViewAndFilterTabsUsedForCreateTestSuite *container.AppTabs) {

	preViewAndFilterTabsUsedForCreateTestSuite = container.NewAppTabs()

	var rightSideScrollContainer *container.Scroll
	rightSideScrollContainer = container.NewScroll(preViewAndFilterTabsUsedForCreateTestSuite)

	rightSideBuildTestSuiteContainer = container.NewVBox(rightSideScrollContainer)

	return rightSideBuildTestSuiteContainer, preViewAndFilterTabsUsedForCreateTestSuite

}
