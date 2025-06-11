package testSuitesTabsUI

import (
	"FenixTesterGui/testCase/testCaseModel"
	"FenixTesterGui/testSuites/testSuiteUI"
	"FenixTesterGui/testSuites/testSuitesModel"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func GenerateTestSuiteHomeTab(testCasesModel *testCaseModel.TestCasesModelsStruct) {

	var err error

	var testSuiteHomePageTabContainer *fyne.Container
	var testSuiteHomePageTabToolbar *widget.Toolbar

	// Create toolbar for TestSuite Home page
	testSuiteHomePageTabToolbar = widget.NewToolbar(

		// New TestSuite
		widget.NewToolbarAction(theme.DocumentIcon(), func() {

			// Generate the new 'TestSuiteModel'
			var newTestSuiteModel *testSuitesModel.TestSuiteModelStruct
			newTestSuiteModel = testSuitesModel.GenerateNewTestSuiteModelObject()

			// Generate a new TestSuiteUI-object
			var newTestSuiteUiObject *testSuiteUI.TestSuiteUiStruct
			newTestSuiteUiObject = &testSuiteUI.TestSuiteUiStruct{
				TestSuiteTabItem:  nil,
				TestSuiteModelPtr: newTestSuiteModel,
			}

			var newTestSuiteUiObjectContainer *fyne.Container

			newTestSuiteUiObjectContainer, err = newTestSuiteUiObject.GenerateBuildNewTestSuiteUI(testCasesModel)

			if err != nil {
				newTestSuiteUiObjectContainer = container.NewVBox(
					widget.NewLabel(fmt.Sprintf("couldn't generate a new 'TestSuite', err=%s'",
						err.Error())))

			}

			// Create the Tab-UI-object
			newTestSuiteUiObject.TestSuiteTabItem = container.NewTabItem("<New TestSuite>", newTestSuiteUiObjectContainer)

			// Get the TestSuiteUiMap from the map-pointer
			var testSuiteUiMap map[*container.TabItem]*testSuiteUI.TestSuiteUiStruct

			// Check if Map needs to be initialized
			if TestSuiteUiMapPtr == nil {
				testSuiteUiMap = make(map[*container.TabItem]*testSuiteUI.TestSuiteUiStruct)
				TestSuiteUiMapPtr = &testSuiteUiMap
			}

			testSuiteUiMap = *TestSuiteUiMapPtr

			// Store the 'newTestSuiteUiObject' in the map
			testSuiteUiMap[newTestSuiteUiObject.TestSuiteTabItem] = newTestSuiteUiObject

			// Add New TestSuite-tab to all tabs
			TestSuiteTabs.Append(newTestSuiteUiObject.TestSuiteTabItem)

			// Set focus on the new TestSuite-tab
			TestSuiteTabs.Select(newTestSuiteUiObject.TestSuiteTabItem)

			// Refresh the new TestSuite-tab
			TestSuiteTabs.Refresh()

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
