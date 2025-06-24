package testSuitesTabsUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testCase/testCaseModel"
	"FenixTesterGui/testSuites/testSuiteUI"
	"FenixTesterGui/testSuites/testSuitesCommandEngine"
	"FenixTesterGui/testSuites/testSuitesModel"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func GenerateTestSuiteTabFromExistingTestSuite(
	testCasesModel *testCaseModel.TestCasesModelsStruct) {

	var err error

	// Check if TestSuite is already opened

	// Generate the new 'TestSuiteModel'
	var newTestSuiteModel *testSuitesModel.TestSuiteModelStruct
	newTestSuiteModel = testSuitesModel.GenerateNewTestSuiteModelObject(testCasesModel)

	// Generate a new TestSuiteUI-object
	var newTestSuiteUiObject *testSuiteUI.TestSuiteUiStruct
	newTestSuiteUiObject = &testSuiteUI.TestSuiteUiStruct{
		TestSuiteTabItem:  nil,
		TestSuiteModelPtr: newTestSuiteModel,
	}

	// Create the Tab-UI-object
	newTestSuiteUiObject.TestSuiteTabItem = container.NewTabItem(
		fmt.Sprintf("<New TestSuite> [%s]", "not created yet"),
		nil)

	var newTestSuiteUiObjectContainer *fyne.Container
	var newTestSuiteUiObjectBorderContainer *fyne.Container

	newTestSuiteUiObjectContainer, err = newTestSuiteUiObject.GenerateBuildNewTestSuiteUI(testCasesModel, newTestSuiteModel)

	if err != nil {
		newTestSuiteUiObjectContainer = container.NewVBox(
			widget.NewLabel(fmt.Sprintf("couldn't generate a new 'TestSuite', err=%s'",
				err.Error())))

		// Generate the BorderContainer that is going to be placed on the Tab
		newTestSuiteUiObjectBorderContainer = container.NewBorder(
			nil, nil, nil, nil, newTestSuiteUiObjectContainer)

	} else {

		var testSuiteUiTabToolbar *widget.Toolbar

		// Create toolbar for a new TestSuite page
		testSuiteUiTabToolbar = widget.NewToolbar(

			// New TestSuite
			widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {

				// Function to Save Suite
				newTestSuiteModel.SaveTestSuite()

			}),
		)

		// Generate the BorderContainer that is going to be placed on the Tab
		newTestSuiteUiObjectBorderContainer = container.NewBorder(
			testSuiteUiTabToolbar, nil, nil, nil, newTestSuiteUiObjectContainer)
	}

	// Get TestSuiteUuid
	var testSuiteUuid string
	var shortTestSuiteUuid string

	testSuiteUuid = newTestSuiteModel.GetTestSuiteUuid()
	shortTestSuiteUuid = sharedCode.GenerateShortUuidFromFullUuid(testSuiteUuid)

	// Add content to the Tab-UI-object
	newTestSuiteUiObject.TestSuiteTabItem.Text = fmt.Sprintf("<New TestSuite> [%s]", shortTestSuiteUuid)
	newTestSuiteUiObject.TestSuiteTabItem.Content = newTestSuiteUiObjectBorderContainer

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

	// Get TestSuitesMap
	var testSuitesMap map[string]*testSuitesModel.TestSuiteModelStruct
	testSuitesMap = *testSuitesModel.TestSuitesModelPtr.TestSuitesMapPtr

	// Add TestSuiteModel to the TestSuiteModels-map
	testSuitesMap[testSuiteUuid] = newTestSuiteModel

	// Add New TestSuite-tab to all tabs
	testSuitesCommandEngine.TestSuiteTabsRef.Append(newTestSuiteUiObject.TestSuiteTabItem)

	// Set focus on the new TestSuite-tab
	testSuitesCommandEngine.TestSuiteTabsRef.Select(newTestSuiteUiObject.TestSuiteTabItem)

	// Refresh the new TestSuite-tab
	testSuitesCommandEngine.TestSuiteTabsRef.Refresh()
}
