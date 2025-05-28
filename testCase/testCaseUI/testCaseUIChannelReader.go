package testCaseUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/soundEngine"
	"FenixTesterGui/testCase/testCaseModel"
	"FenixTesterGui/testCases/listTestCasesUI"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/sirupsen/logrus"
	"time"
)

// Channel reader which is used for reading out command to update GUI
func (testCasesUiCanvasObject *TestCasesUiModelStruct) startGUICommandChannelReader() {

	var incomingChannelCommandGraphicsUpdatedData sharedCode.ChannelCommandGraphicsUpdatedStruct

	for {
		// Wait for incoming trigger command over channel
		incomingChannelCommandGraphicsUpdatedData = <-*testCasesUiCanvasObject.GraphicsUpdateChannelReference

		switch incomingChannelCommandGraphicsUpdatedData.ChannelCommandGraphicsUpdate {

		case sharedCode.ChannelCommandGraphicsUpdatedNewTestCase:
			testCasesUiCanvasObject.updateTestCaseGraphics(incomingChannelCommandGraphicsUpdatedData)

		case sharedCode.ChannelCommandGraphicsUpdatedUpdateTestCaseGraphics:
			testCasesUiCanvasObject.updateTestCaseGraphics(incomingChannelCommandGraphicsUpdatedData)

		case sharedCode.ChannelCommandGraphicsUpdatedSelectTestInstruction:
			testCasesUiCanvasObject.selectTestInstructionInTestCaseGraphics(incomingChannelCommandGraphicsUpdatedData)

		case sharedCode.ChannelCommandGraphicsUpdatedSelectTestCaseTabBasedOnTestCaseUuid:
			testCasesUiCanvasObject.selectTestCaseTabBasedOnTestCaseUuid(incomingChannelCommandGraphicsUpdatedData)

		case sharedCode.ChannelCommandGraphicsUpdatedUpdateTestCaseTabName:
			testCasesUiCanvasObject.updatedUpdateTestCaseTabName(incomingChannelCommandGraphicsUpdatedData)

		case sharedCode.ChannelCommandGraphicsRemoveTestCaseTabBasedOnTestCaseUuid:
			testCasesUiCanvasObject.removeTestCaseTabBasedOnTestCaseUuid(incomingChannelCommandGraphicsUpdatedData)

		case sharedCode.ChannelCommandGraphicsCloseTestCaseTabBasedOnTestCaseUuiWithOutSaving:
			testCasesUiCanvasObject.closeTestCaseTabBasedOnTestCaseUuiWithOutSaving(incomingChannelCommandGraphicsUpdatedData)

		default:
			errorId := "388e2a87-1d0e-4db3-8dcf-18a69ac1faa4"
			err := errors.New(fmt.Sprintf("unknow 'incomingChannelCommandGraphicsUpdatedData', [ErrorID: %s]", errorId))

			//TODO Send error over error-channel
			fmt.Println(err) // TODO Send on Error-channel
			return

		}

	}

}

func (testCasesUiCanvasObject *TestCasesUiModelStruct) updateTestCaseGraphics(incomingChannelCommandGraphicsUpdatedData sharedCode.ChannelCommandGraphicsUpdatedStruct) {

	// Generate UI for New TestCase
	if incomingChannelCommandGraphicsUpdatedData.CreateNewTestCaseUI == true {

		err := testCasesUiCanvasObject.GenerateNewTestCaseTabObject(incomingChannelCommandGraphicsUpdatedData.ActiveTestCase)
		if err != nil {
			//TODO Send error over error-channel
			fmt.Println(err)

			return
		}
	}

	// Update Textual Representations, in UI-model, for TestCase
	err := testCasesUiCanvasObject.UpdateTextualStructuresForTestCase(
		incomingChannelCommandGraphicsUpdatedData.ActiveTestCase,
		incomingChannelCommandGraphicsUpdatedData.TextualTestCaseSimple,
		incomingChannelCommandGraphicsUpdatedData.TextualTestCaseComplex,
		incomingChannelCommandGraphicsUpdatedData.TextualTestCaseExtended)

	if err != nil {
		//TODO Send error over error-channel
		fmt.Println(err)

		return
	}

	// Update Graphical TestCase Representation
	err = testCasesUiCanvasObject.UpdateGraphicalRepresentationForTestCase(incomingChannelCommandGraphicsUpdatedData.ActiveTestCase)
	if err != nil {
		//TODO Send error over error-channel
		fmt.Println(err)

		return
	}

	// Set the active TestCase, from UI-perspective
	testCasesUiCanvasObject.TestCasesModelReference.CurrentActiveTestCaseUuid = incomingChannelCommandGraphicsUpdatedData.ActiveTestCase

}

// Select TestInstruction in Active TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) selectTestInstructionInTestCaseGraphics(incomingChannelCommandGraphicsUpdatedData sharedCode.ChannelCommandGraphicsUpdatedStruct) {

	// Select Latest dropped element
	var newPointEvent fyne.PointEvent
	testCasesUiCanvasObject.CurrentSelectedTestCaseUIElement.Tapped(&newPointEvent)

}

// Select tab that have TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) selectTestCaseTabBasedOnTestCaseUuid(incomingChannelCommandGraphicsUpdatedData sharedCode.ChannelCommandGraphicsUpdatedStruct) {

	var foundTestCase bool
	var tabReference *container.TabItem

	// Loop Map with TestCase-tabs to find relation between TabItem and UUID
	for _, tempTestCaseUITabRefToTestCaseUuidMapStructObject := range testCasesUiCanvasObject.TestCaseUITabRefToTestCaseUuidMap {

		// Is this the TestCaseUuid we are looking for
		if tempTestCaseUITabRefToTestCaseUuidMapStructObject.TestCaseUuid == incomingChannelCommandGraphicsUpdatedData.ActiveTestCase {
			foundTestCase = true
			tabReference = tempTestCaseUITabRefToTestCaseUuidMapStructObject.TestCaseUiTabRef
			break
		}
	}

	// When TestCase was found then switch tab
	if foundTestCase == true {
		testCasesUiCanvasObject.TestCasesTabs.Select(tabReference)
		//testCasesUiCanvasObject.TestCasesTabs.Refresh()

		return

	} else {
		// No TestCase was found
		//TODO Send error over error-channel
		fmt.Println("No Tab was found, but was expected for TestCase:", incomingChannelCommandGraphicsUpdatedData.ActiveTestCase)

		return
	}

}

// Update that tab name for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) updatedUpdateTestCaseTabName(incomingChannelCommandGraphicsUpdatedData sharedCode.ChannelCommandGraphicsUpdatedStruct) {

	testCasesUiCanvasObject.TestCasesTabs.Selected().Text = incomingChannelCommandGraphicsUpdatedData.TestCaseTabName
	testCasesUiCanvasObject.TestCasesTabs.Refresh()

}

// Remove tab that have the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) removeTestCaseTabBasedOnTestCaseUuid(
	incomingChannelCommandGraphicsUpdatedData sharedCode.ChannelCommandGraphicsUpdatedStruct) {

	var foundTestCaseTab bool
	var tabReference *container.TabItem
	var previousTabReference *container.TabItem
	var tabReferenceAsString string

	// Loop Map with TestCase-tabs to find relation between TabItem and UUID
	for tempTabReferenceAsString, tempTestCaseUITabRefToTestCaseUuidMapStructObject := range testCasesUiCanvasObject.TestCaseUITabRefToTestCaseUuidMap {

		// Is this the TestCaseUuid we are looking for
		if tempTestCaseUITabRefToTestCaseUuidMapStructObject.TestCaseUuid == incomingChannelCommandGraphicsUpdatedData.ActiveTestCase {
			foundTestCaseTab = true
			tabReference = tempTestCaseUITabRefToTestCaseUuidMapStructObject.TestCaseUiTabRef
			tabReferenceAsString = tempTabReferenceAsString
			break
		}

		// Save the previous tab reference
		previousTabReference = tempTestCaseUITabRefToTestCaseUuidMapStructObject.TestCaseUiTabRef
	}

	// When TestCase was found then remove tab
	if foundTestCaseTab == true {

		// Set TestCase to be deleted in Database

		// Check if The date is Today() or in the future
		var parseError error

		var validTodayDate string
		validTodayDate = time.Now().Format("2006-01-02")

		_, parseError = time.Parse("2006-01-02", newTestCaseDeletionDateEntry.Text)
		if parseError != nil {
			sharedCode.Logger.WithFields(logrus.Fields{
				"id": "fb41e410-cf07-489d-a993-20bb3e65305c",
				"incomingChannelCommandGraphicsUpdatedData.ActiveTestCase": incomingChannelCommandGraphicsUpdatedData.ActiveTestCase,
			}).Fatal(fmt.Sprintf("Couldn't parse time.Now(), should never happen:"))

		}

		var existInMap bool
		var currentTestCasePtr *testCaseModel.TestCaseModelStruct

		currentTestCasePtr, existInMap = testCasesUiCanvasObject.TestCasesModelReference.
			TestCasesMap[incomingChannelCommandGraphicsUpdatedData.ActiveTestCase]
		if existInMap == false {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":           "2e89c13f-0d3f-4dbd-86dd-c35e2a4b59e8",
				"testCaseUuid": incomingChannelCommandGraphicsUpdatedData.ActiveTestCase,
			}).Fatal("TestCase doesn't exist in TestCaseMap. This should not happen")
		}

		// This TestCase is New and not saved in Database
		if currentTestCasePtr.ThisIsANewTestCase == true {

			//If Delete date is equal to Today() then Delete everything
			if currentTestCasePtr.LocalTestCaseMessage.DeleteTimeStamp <= validTodayDate {

				// flash the window if Te
				go flashScreen(*sharedCode.FenixAppPtr, *sharedCode.FenixMasterWindowPtr)

				// Remove the Tab with the TestCase UI-objects
				testCasesUiCanvasObject.TestCasesTabs.Remove(tabReference)
				//testCasesUiCanvasObject.TestCasesTabs.Refresh()

				// Remove TestCase from UI-map
				delete(testCasesUiCanvasObject.TestCasesUiModelMap, incomingChannelCommandGraphicsUpdatedData.ActiveTestCase)

				// Remove TestCase TestCasesMap-model
				delete(testCasesUiCanvasObject.TestCasesModelReference.TestCasesMap, incomingChannelCommandGraphicsUpdatedData.ActiveTestCase)

				// Switch Active TestCase by Loop Map with TestCase-tabs to find relation between TabItem and UUID
				var testCaseUuidToSwitchTo string
				var foundPreviousTestCaseTab bool
				for _, tempTestCaseUITabRefToTestCaseUuidMapStructObject := range testCasesUiCanvasObject.TestCaseUITabRefToTestCaseUuidMap {

					// Is this the TestCaseUuid we are looking for
					if tempTestCaseUITabRefToTestCaseUuidMapStructObject.TestCaseUiTabRef == previousTabReference {
						foundPreviousTestCaseTab = true
						testCaseUuidToSwitchTo = tempTestCaseUITabRefToTestCaseUuidMapStructObject.TestCaseUuid
						break
					}

				}

				// Delete tab from TabReference-map
				delete(testCasesUiCanvasObject.TestCaseUITabRefToTestCaseUuidMap, tabReferenceAsString)

				// If no tab was found the active tab is "empty" base-tab, otherwise set the tabs TestCaseUuid to the active one
				if foundPreviousTestCaseTab == false {
					testCasesUiCanvasObject.TestCasesModelReference.CurrentActiveTestCaseUuid = ""
				} else {
					testCasesUiCanvasObject.TestCasesModelReference.CurrentActiveTestCaseUuid = testCaseUuidToSwitchTo
				}

				return
			} else {
				// Delete date is in the future so do nothing
			}

		} else {

			// Try to set Delete date in database for the TestCase
			var err error
			err = testCasesUiCanvasObject.TestCasesModelReference.DeleteTestCaseAtThisDate(incomingChannelCommandGraphicsUpdatedData.ActiveTestCase)

			if err != nil {

				errorId := "8daeb14c-2d73-4413-95ed-21033c7b46c6"
				err = errors.New(fmt.Sprintf("Couldn't Set delete date for TestCase '%s' in database [ErrorID: %s]. Error='%s'",
					incomingChannelCommandGraphicsUpdatedData.ActiveTestCase, errorId, err.Error()))

				fmt.Println(err) // TODO Send on Error-channel

				// Clear Date in Entry-box
				newTestCaseDeletionDateEntry.SetText("")
				enableDeletionCheckbox.Disable()

				return
			}

			//If Delete date is equal to Today() then Delete everything
			if currentTestCasePtr.LocalTestCaseMessage.DeleteTimeStamp <= validTodayDate {

				// flash the window if Te
				go flashScreen(*sharedCode.FenixAppPtr, *sharedCode.FenixMasterWindowPtr)

				// Remove the Tab with the TestCase UI-objects
				testCasesUiCanvasObject.TestCasesTabs.Remove(tabReference)
				//testCasesUiCanvasObject.TestCasesTabs.Refresh()

				// Remove TestCase from UI-map
				delete(testCasesUiCanvasObject.TestCasesUiModelMap,
					incomingChannelCommandGraphicsUpdatedData.ActiveTestCase)

				// Remove TestCase TestCasesMap-model
				delete(testCasesUiCanvasObject.TestCasesModelReference.TestCasesMap,
					incomingChannelCommandGraphicsUpdatedData.ActiveTestCase)

				// Switch Active TestCase by Loop Map with TestCase-tabs to find relation between TabItem and UUID
				var testCaseUuidToSwitchTo string
				var foundPreviousTestCaseTab bool
				for _, tempTestCaseUITabRefToTestCaseUuidMapStructObject := range testCasesUiCanvasObject.TestCaseUITabRefToTestCaseUuidMap {

					// Is this the TestCaseUuid we are looking for
					if tempTestCaseUITabRefToTestCaseUuidMapStructObject.TestCaseUiTabRef == previousTabReference {
						foundPreviousTestCaseTab = true
						testCaseUuidToSwitchTo = tempTestCaseUITabRefToTestCaseUuidMapStructObject.TestCaseUuid
						break
					}

				}

				// Delete tab from TabReference-map
				delete(testCasesUiCanvasObject.TestCaseUITabRefToTestCaseUuidMap, tabReferenceAsString)

				// If no tab was found the active tab is "emtpty" base-tab, otherwise set the tabs TestCaseUuid to the active one
				if foundPreviousTestCaseTab == false {
					testCasesUiCanvasObject.TestCasesModelReference.CurrentActiveTestCaseUuid = ""
				} else {
					testCasesUiCanvasObject.TestCasesModelReference.CurrentActiveTestCaseUuid = testCaseUuidToSwitchTo
				}

				// Remove TestCase from TestCase-List
				listTestCasesUI.RemoveTestCaseFromList(incomingChannelCommandGraphicsUpdatedData.ActiveTestCase,
					testCasesUiCanvasObject.TestCasesModelReference)

				return
			} else {
				// Delete date is in the future so only Notify That testCase is set to bed deleted in the future

				// Trigger System Notification sound
				soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

				fyne.CurrentApp().SendNotification(&fyne.Notification{
					Title: "TestCase Deleted",
					Content: fmt.Sprintf("The TestCase was set to Deleted in the future (%s)",
						currentTestCasePtr.LocalTestCaseMessage.DeleteTimeStamp),
				})
			}

		}

	} else {
		// No TestCase was found
		//TODO Send error over error-channel
		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "5c4319ec-7952-4713-8b3a-63c92fcf71f9",
			"incomingChannelCommandGraphicsUpdatedData.ActiveTestCase": incomingChannelCommandGraphicsUpdatedData.ActiveTestCase,
		}).Fatal(fmt.Sprintf("No Tab was found, but was expected for TestCase: '%s'", incomingChannelCommandGraphicsUpdatedData.ActiveTestCase))

		return
	}

}

// Close tab that have the TestCase without saving the TestCse
func (testCasesUiCanvasObject *TestCasesUiModelStruct) closeTestCaseTabBasedOnTestCaseUuiWithOutSaving(
	incomingChannelCommandGraphicsUpdatedData sharedCode.ChannelCommandGraphicsUpdatedStruct) {

	var foundTestCaseTab bool
	var tabReference *container.TabItem
	var previousTabReference *container.TabItem
	var tabReferenceAsString string

	// Loop Map with TestCase-tabs to find relation between TabItem and UUID
	for tempTabReferenceAsString, tempTestCaseUITabRefToTestCaseUuidMapStructObject := range testCasesUiCanvasObject.TestCaseUITabRefToTestCaseUuidMap {

		// Is this the TestCaseUuid we are looking for
		if tempTestCaseUITabRefToTestCaseUuidMapStructObject.TestCaseUuid == incomingChannelCommandGraphicsUpdatedData.ActiveTestCase {
			foundTestCaseTab = true
			tabReference = tempTestCaseUITabRefToTestCaseUuidMapStructObject.TestCaseUiTabRef
			tabReferenceAsString = tempTabReferenceAsString
			break
		}

		// Save the previous tab reference
		previousTabReference = tempTestCaseUITabRefToTestCaseUuidMapStructObject.TestCaseUiTabRef
	}

	// When TestCase was found then remove tab
	if foundTestCaseTab == true {

		var existInMap bool

		_, existInMap = testCasesUiCanvasObject.TestCasesModelReference.
			TestCasesMap[incomingChannelCommandGraphicsUpdatedData.ActiveTestCase]

		if existInMap == false {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":           "b93f6a4d-7d35-4ed8-9b80-e11ba29c62a2",
				"testCaseUuid": incomingChannelCommandGraphicsUpdatedData.ActiveTestCase,
			}).Fatal("TestCase doesn't exist in TestCaseMap. This should not happen")
		}

		fmt.Println("TabIndex: ", testCasesUiCanvasObject.TestCasesTabs.SelectedIndex())

		// Remove the Tab with the TestCase UI-objects
		testCasesUiCanvasObject.TestCasesTabs.Remove(tabReference)
		//testCasesUiCanvasObject.TestCasesTabs.Refresh()

		fmt.Println("TabIndex: ", testCasesUiCanvasObject.TestCasesTabs.SelectedIndex())

		// Remove TestCase from UI-map
		delete(testCasesUiCanvasObject.TestCasesUiModelMap,
			incomingChannelCommandGraphicsUpdatedData.ActiveTestCase)

		// Remove TestCase TestCasesMap-model
		delete(testCasesUiCanvasObject.TestCasesModelReference.TestCasesMap,
			incomingChannelCommandGraphicsUpdatedData.ActiveTestCase)

		// Switch Active TestCase by Loop Map with TestCase-tabs to find relation between TabItem and UUID
		var testCaseUuidToSwitchTo string
		var foundPreviousTestCaseTab bool
		for _, tempTestCaseUITabRefToTestCaseUuidMapStructObject := range testCasesUiCanvasObject.TestCaseUITabRefToTestCaseUuidMap {

			// When 'previousTabReference' is nil then there only "Home"-tab left
			if previousTabReference == nil {
				testCaseUuidToSwitchTo = ""
				break

			} else {

				// Is this the TestCaseUuid we are looking for
				if tempTestCaseUITabRefToTestCaseUuidMapStructObject.TestCaseUiTabRef == previousTabReference {
					foundPreviousTestCaseTab = true
					testCaseUuidToSwitchTo = tempTestCaseUITabRefToTestCaseUuidMapStructObject.TestCaseUuid
					break
				}
			}
		}

		// Delete tab from TabReference-map
		delete(testCasesUiCanvasObject.TestCaseUITabRefToTestCaseUuidMap, tabReferenceAsString)

		// If no tab was found the active tab is "emtpty" base-tab, otherwise set the tabs TestCaseUuid to the active one
		if foundPreviousTestCaseTab == false {
			testCasesUiCanvasObject.TestCasesModelReference.CurrentActiveTestCaseUuid = ""
		} else {
			testCasesUiCanvasObject.TestCasesModelReference.CurrentActiveTestCaseUuid = testCaseUuidToSwitchTo
		}

		// Delete date is in the future so only Notify That testCase is set to bed deleted in the future

		// Trigger System Notification sound
		soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "TestCase Closed",
			Content: "The TestCase was closed without to be saved",
		})

	} else {
		// No TestCase was found
		//TODO Send error over error-channel
		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "9ef7a49d-531d-4d98-9c38-5f6b735598b9",
			"incomingChannelCommandGraphicsUpdatedData.ActiveTestCase": incomingChannelCommandGraphicsUpdatedData.ActiveTestCase,
		}).Fatal(fmt.Sprintf("No Tab was found, but was expected for TestCase: '%s'", incomingChannelCommandGraphicsUpdatedData.ActiveTestCase))

		return
	}

}
