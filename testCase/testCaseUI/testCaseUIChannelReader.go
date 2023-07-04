package testCaseUI

import (
	sharedCode "FenixTesterGui/common_code"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
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
