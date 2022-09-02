package testCaseUI

import (
	sharedCode "FenixTesterGui/common_code"
	"fmt"
)

// Channel reader which is used for reading out command to update GUI
func (testCasesUiCanvasObject *TestCasesUiModelStruct) startGUICommandChannelReader() {

	var incomingChannelCommandGraphicsUpdatedData sharedCode.ChannelCommandGraphicsUpdatedStruct

	for {
		// Wait for incoming trigger command over channel
		incomingChannelCommandGraphicsUpdatedData = <-*testCasesUiCanvasObject.GraphicsUpdateChannelReference

		testCasesUiCanvasObject.updateTestCaseGraphics(incomingChannelCommandGraphicsUpdatedData)

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

}
