package testSuitesTabsUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/soundEngine"
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
	testSuiteUuidToOpen string,
	testCasesModel *testCaseModel.TestCasesModelsStruct) {

	var err error

	// Verify that it is an Uuid that was received
	if len(testSuiteUuidToOpen) != 36 {

		// Notify the user

		// Trigger System Notification sound
		soundEngine.PlaySoundChannel <- soundEngine.InvalidNotificationSound

		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "Warning",
			Content: "TestSuite-uuid is not valid. It should be 36 characters long.",
		})

		return
	}

	// Check if TestSuite is already opened

	// Generate the new 'TestSuiteModel'
	var newTestSuiteModel *testSuitesModel.TestSuiteModelStruct
	newTestSuiteModel = testSuitesModel.GenerateNewTestSuiteModelObject(
		testSuiteUuidToOpen,
		testCasesModel)

	// Get TestSuiteUuid
	var shortTestSuiteUuid string
	shortTestSuiteUuid = sharedCode.GenerateShortUuidFromFullUuid(testSuiteUuidToOpen)

	// Load the new model with an existing TestSuite
	err = newTestSuiteModel.LoadFullTestSuiteFromDatabase(testSuiteUuidToOpen)

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
	testSuitesMap[testSuiteUuidToOpen] = newTestSuiteModel

	// Add New TestSuite-tab to all tabs
	testSuitesCommandEngine.TestSuiteTabsRef.Append(newTestSuiteUiObject.TestSuiteTabItem)

	// Set focus on the new TestSuite-tab
	testSuitesCommandEngine.TestSuiteTabsRef.Select(newTestSuiteUiObject.TestSuiteTabItem)

	// Refresh the new TestSuite-tab
	testSuitesCommandEngine.TestSuiteTabsRef.Refresh()
}
