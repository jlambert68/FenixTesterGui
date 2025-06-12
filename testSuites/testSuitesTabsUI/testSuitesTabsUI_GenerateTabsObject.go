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
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
)

func GenerateTestSuiteTabObject(testCasesModel *testCaseModel.TestCasesModelsStruct) *container.DocTabs {

	// Initiate TestSuiteTabsRef-object
	testSuitesCommandEngine.TestSuiteTabsRef = container.NewDocTabs()

	// Add intercept when closing to be able to NOT close the TestSuite-Home-tab
	testSuitesCommandEngine.TestSuiteTabsRef.CloseIntercept = func(tab *container.TabItem) {

		// Never close Home-tab for TestSuites
		if tab == testSuiteHomeTabItem {
			// Notify the user that Home-tab for TestSuites can't be closed

			// Trigger System Notification sound
			soundEngine.PlaySoundChannel <- soundEngine.InvalidNotificationSound

			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Information",
				Content: fmt.Sprintf("TestSuite Home-tab can't be closed!"),
			})

		} else {

			var existInMap bool

			// Get 'TestSuitesMapPtr' and 'TestSuitesMap'
			var testSuitesMapPtr *map[string]*testSuitesModel.TestSuiteModelStruct
			var testSuitesMap map[string]*testSuitesModel.TestSuiteModelStruct

			testSuitesMapPtr = testSuitesModel.TestSuitesModelPtr.TestSuitesMapPtr
			testSuitesMap = *testSuitesMapPtr

			// Get 'TestSuiteUiMap'
			var testSuiteUiMap map[*container.TabItem]*testSuiteUI.TestSuiteUiStruct
			testSuiteUiMap = *TestSuiteUiMapPtr

			// Get the TestSuiteUIModel
			var testSuiteUiModelPtr *testSuiteUI.TestSuiteUiStruct
			var testSuiteUiModel testSuiteUI.TestSuiteUiStruct
			testSuiteUiModelPtr, existInMap = testSuiteUiMap[tab]
			if existInMap == false {
				sharedCode.Logger.WithFields(logrus.Fields{
					"ID": "ee4fcb9d-8c9b-477c-b730-7bef1e7e3539",
				}).Fatal("TestSuiteUiModel doesn't exist in TestSuiteUiMap. This should not happen")
			}

			testSuiteUiModel = *testSuiteUiModelPtr

			// Get TestSuiteUuid
			var testSuiteUuid string
			testSuiteUuid = testSuiteUiModel.TestSuiteModelPtr.GetTestSuiteUuid()

			// Get the TestSuiteModel
			var testSuitesModelPtr *testSuitesModel.TestSuiteModelStruct
			var testSuitesModel testSuitesModel.TestSuiteModelStruct

			testSuitesModelPtr, existInMap = testSuitesMap[testSuiteUuid]
			if existInMap == false {
				sharedCode.Logger.WithFields(logrus.Fields{
					"ID": "33bae085-1a7f-45e4-8b05-2500a9ed310b",
				}).Fatal("TestSuiteUiModel doesn't exist in TestSuiteMap. This should not happen")
			}

			testSuitesModel = *testSuitesModelPtr

			// Check if TestSuite has changed data
			var testSuiteHasChangedData bool
			testSuiteHasChangedData = testSuitesModel.IsTestSuiteChanged()

			if testSuiteHasChangedData == true {
				question := widget.NewLabel("Do you want to close TestSuite when there are unsaved changes?")
				dialog.ShowCustomConfirm(
					"Close TestSuite",
					"Just close it",
					"I want to save the changes",
					question,
					func(close bool) {
						if close {
							// Close TestSuite
							testSuitesCommandEngine.TestSuiteTabsRef.Remove(tab)

							// remove Tab from Tab-map
							delete(testSuiteUiMap, tab)

						} else {
							// user cancelled
						}
					},
					*sharedCode.FenixMasterWindowPtr,
				)
			} else {
				// No changes so, just close the Tab

				// Close TestSuite
				testSuitesCommandEngine.TestSuiteTabsRef.Remove(tab)

				// remove Tab from Tab-map
				delete(testSuiteUiMap, tab)
			}
		}
	}

	// Generate TestSuite-Home page
	GenerateTestSuiteHomeTab(testCasesModel)

	return testSuitesCommandEngine.TestSuiteTabsRef
}
