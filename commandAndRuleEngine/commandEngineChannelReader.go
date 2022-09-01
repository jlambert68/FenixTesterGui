package commandAndRuleEngine

import (
	sharedCode "FenixTesterGui/common_code"
	"fmt"
)

// Channel reader which is used for reading out commands to CommandEngine
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) startCommandChannelReader() {

	var incomingChannelCommand sharedCode.ChannelCommandStruct

	for {
		// Wait for incomming command over channel
		incomingChannelCommand = <-sharedCode.CommandChannel

		switch incomingChannelCommand.ChannelCommand {

		case sharedCode.ChannelCommandNewTestCase:
			commandAndRuleEngine.createNewTestCase()

		case sharedCode.ChannelCommandSwapElement:
			commandAndRuleEngine.

		case sharedCode.ChannelCommandRemoveElement:
			commandAndRuleEngine.channelCommandRemove()

		// No other command is supported
		default:
			//TODO Send Error over ERROR-channel
		}
	}

}

// Execute command 'New TestCase' received from Channel reader
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) createNewTestCase() {

	// Create New TestCase
	testCaseUuid, err := commandAndRuleEngine.NewTestCaseModel()

	if err != nil {
		fmt.Println(err)

		return

	}

	// Update UI with TestCase Textual Representation
	textualTestCaseSimple, textualTestCaseComplex, textualTestCaseExtended, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)
	if err != nil {
		fmt.Println(err)

		return
	}

	// Generate UI for New TestCase
	err = commandAndRuleEngine.uiServer.testCasesUiModel.GenerateNewTestCaseTabObject(testCaseUuid)
	if err != nil {
		fmt.Println(err)

		return
	}

	// Update Textual Representations, in UI-model, for TestCase
	err = uiServer.testCasesUiModel.UpdateTextualStructuresForTestCase(
		testCaseUuid,
		textualTestCaseSimple,
		textualTestCaseComplex,
		textualTestCaseExtended)

	if err != nil {
		fmt.Println(err)

		return
	}

}

// Execute command 'Remove Element' received from Channel reader
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) channelCommandRemove(currentTestCaseUuid string, elementToRemove string) {

	// Delete Element from TestCase
	err := commandAndRuleEngine.DeleteElementFromTestCaseModel(currentTestCaseUuid, elementToRemove)
	if err != nil {
		//TODO Send ERRORS over error-channel
		fmt.Println(err)

		return
	}



	// Update UI with TestCase Textual Representation
	textualTestCaseSimple, textualTestCaseComplex, textualTestCaseExtended, err := uiServer.commandAndRuleEngine.Testcases.CreateTextualTestCase(availableTestCasesSelectWidget.Selected)
	if err != nil {
		fmt.Println(err)
	}

	// Update Textual Representations, in UI-model, for TestCase
	err = uiServer.testCasesUiModel.UpdateTextualStructuresForTestCase(
		testCaseUuid,
		textualTestCaseSimple,
		textualTestCaseComplex,
		textualTestCaseExtended)

	if err != nil {
		fmt.Println(err)

		return
	}

	// Update Graphical TestCase Representation
	err = uiServer.testCasesUiModel.UpdateGraphicalRepresentationForTestCase(testCaseUuid)
	if err != nil {
		fmt.Println(err)

		return
	}
}


// Execute command 'Swap Element' received from Channel reader
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) channelCommandSwap(currentTestCaseUuid string, elementToRemove string) {

	// Execute Swap of Elements
	err := commandAndRuleEngine.SwapElementsInTestCaseModel(currentTestCaseUuid, elementUuidTobeSwappedOut, &immatureElementToSwapInTestCaseFormat)
	if err != nil {
		fmt.Println(err)

		return
	}

	// Update UI with TestCase Textual Representation
	textualTestCaseSimple, textualTestCaseComplex, textualTestCaseExtended, err := uiServer.commandAndRuleEngine.Testcases.CreateTextualTestCase(availableTestCasesSelectWidget.Selected)
	if err != nil {
		fmt.Println(err)
	}

	// Update Textual Representations, in UI-model, for TestCase
	err = uiServer.testCasesUiModel.UpdateTextualStructuresForTestCase(
		testCaseUuid,
		textualTestCaseSimple,
		textualTestCaseComplex,
		textualTestCaseExtended)

	if err != nil {
		fmt.Println(err)

		return
	}

	// Update Graphical TestCase Representation
	err = uiServer.testCasesUiModel.UpdateGraphicalRepresentationForTestCase(testCaseUuid)
	if err != nil {
		fmt.Println(err)

		return
	}

}