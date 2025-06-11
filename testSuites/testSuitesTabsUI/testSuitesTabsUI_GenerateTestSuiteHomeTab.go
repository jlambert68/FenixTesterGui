package testSuitesTabsUI

import (
	"FenixTesterGui/testCase/testCaseModel"
	"FenixTesterGui/testSuites/testSuiteUI"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func GenerateTestSuiteHomeTab(testCasesModel *testCaseModel.TestCasesModelsStruct) {

	var testSuiteHomePageTabContainer *fyne.Container
	var testSuiteHomePageTabToolbar *widget.Toolbar

	// Create toolbar for TestSuite Home page
	testSuiteHomePageTabToolbar = widget.NewToolbar(

		// New TestSuite
		widget.NewToolbarAction(theme.DocumentIcon(), func() {

			GenerateNewTestSuiteTab(testCasesModel)

		}),

		// Open TestSuite
		widget.NewToolbarAction(theme.FolderOpenIcon(), func() {

			fmt.Println("Open TestSuite")
		}),
	)

	// Create the Container that will be placed on the HomaPage-tab
	testSuiteHomePageTabContainer = container.NewBorder(
		testSuiteHomePageTabToolbar, nil, nil, nil, nil)

	// Create the HomePage-tab
	testSuiteHomeTabItem = &container.TabItem{
		Text:    "Build TestSuites",
		Icon:    theme.HomeIcon(),
		Content: testSuiteHomePageTabContainer,
	}

	// Add HomePage-tab to all TestSuite-Tabs
	TestSuiteTabs.Append(testSuiteHomeTabItem)

	// Add HomePage-tab 'open' Tab-items for TestSuites
	var testSuiteUiMap map[*container.TabItem]*testSuiteUI.TestSuiteUiStruct

	if TestSuiteUiMapPtr == nil {
		testSuiteUiMap = make(map[*container.TabItem]*testSuiteUI.TestSuiteUiStruct)
		TestSuiteUiMapPtr = &testSuiteUiMap
	}
	testSuiteUiMap = *TestSuiteUiMapPtr

	var testSuiteUIForHomeTab *testSuiteUI.TestSuiteUiStruct
	testSuiteUIForHomeTab = &testSuiteUI.TestSuiteUiStruct{
		TestSuiteTabItem:  testSuiteHomeTabItem,
		TestSuiteModelPtr: nil,
	}

	testSuiteUiMap[testSuiteHomeTabItem] = testSuiteUIForHomeTab

}
