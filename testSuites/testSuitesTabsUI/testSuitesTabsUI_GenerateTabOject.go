package testSuitesTabsUI

import (
	"FenixTesterGui/testCase/testCaseModel"
	"fyne.io/fyne/v2/container"
)

func GenerateTestSuiteTabObject(testCasesModel *testCaseModel.TestCasesModelsStruct) *container.DocTabs {

	// Initiate TestSuiteTabs-object
	TestSuiteTabs = container.NewDocTabs()

	// Generate TestSuite-Home page
	GenerateTestSuiteHomeTab(testCasesModel)

	return TestSuiteTabs
}
