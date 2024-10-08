package testCaseUI

import (
	sharedCode "FenixTesterGui/common_code"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/sirupsen/logrus"
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

	// When TestCase was found then remove tab
	if foundTestCase == true {
		testCasesUiCanvasObject.TestCasesTabs.Remove(tabReference)
		//testCasesUiCanvasObject.TestCasesTabs.Refresh()

		// Remove TestCase from UI-map
		delete(testCasesUiCanvasObject.TestCasesUiModelMap, incomingChannelCommandGraphicsUpdatedData.ActiveTestCase)

		// Remove TestCase TestCases-model
		delete(testCasesUiCanvasObject.TestCasesModelReference.TestCases, incomingChannelCommandGraphicsUpdatedData.ActiveTestCase)

		return

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
