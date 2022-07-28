package commandAndRuleEngine

import (
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"

	//"errors"
	uuidGenerator "github.com/google/uuid"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
)

// TestCaseCommandTypeEnum_NEW_TESTCASE
// Used for creating a new TestCase-Model to be used within a new TestCase
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_NewTestCaseModel() (testCaseUuid string, err error) {

	// Create new B0-Bind
	b0Bond := commandAndRuleEngine.createNewBondB0Element()

	// Initiate a TestCaseModel
	newTestCaseModel := testCaseModel.TestCaseModelStruct{
		LastLoadedTestCaseModelGRPCMessage: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage{},
		FirstElementUuid:                   "",
		TestCaseModelMap:                   nil,
	}

	// Initiate TestCaseModel-map
	newTestCaseModel.TestCaseModelMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage)

	// Add B0-bond to the TestCaseModel-map
	newTestCaseModel.TestCaseModelMap[b0Bond.MatureElementUuid] = b0Bond

	// Set the B0-bond as first element in TestCaseModel
	newTestCaseModel.FirstElementUuid = b0Bond.MatureElementUuid

	// Create a new Command
	newCommandEntry := fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage_TestCaseModelCommandMessage{
		TestCaseCommandType:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_NEW_TESTCASE,
		TestCaseCommandName:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_name[int32(fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_NEW_TESTCASE)],
		FirstParameter:           testCaseModel.NotApplicable,
		SecondParameter:          testCaseModel.NotApplicable,
		UserId:                   commandAndRuleEngine.testcases.CurrentUser,
		CommandExecutedTimeStamp: timestamppb.Now(),
	}

	// Add command to command stack
	newTestCaseModel.CommandStack = append(newTestCaseModel.CommandStack, newCommandEntry)

	// Generate new TestCase-UUID
	testCaseUuid = uuidGenerator.New().String()

	// Add the TestCaseModel into map of all TestCaseModels
	commandAndRuleEngine.testcases.TestCases[testCaseUuid] = newTestCaseModel

	// Add command Textual representations to Textual Representation Stack
	textualRepresentationSimple, textualRepresentationComplex, err := commandAndRuleEngine.testcases.CreateTextualTestCase(testCaseUuid)

	if err == nil {
		newTestCaseModel.TextualTestCaseRepresentationSimpleStack = append(newTestCaseModel.TextualTestCaseRepresentationSimpleStack, textualRepresentationSimple)
		newTestCaseModel.TextualTestCaseRepresentationComplexStack = append(newTestCaseModel.TextualTestCaseRepresentationComplexStack, textualRepresentationComplex)

		// Add the TestCaseModel back into map of all TestCaseModels
		commandAndRuleEngine.testcases.TestCases[testCaseUuid] = newTestCaseModel

	}

	return testCaseUuid, err

}

// TestCaseCommandTypeEnum_REMOVE_ELEMENT
// Used for Deleting an element from a TestCaseModel that is used within a TestCase
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel(testCaseUuid string, elementId string) (err error) {

	// Try to Delete element
	err = commandAndRuleEngine.executeDeleteElement(testCaseUuid, elementId)

	// Exit if the element couldn't be deleted
	if err != nil {
		return err
	}

	// Create a new Command
	newCommandEntry := fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage_TestCaseModelCommandMessage{
		TestCaseCommandType:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_REMOVE_ELEMENT,
		TestCaseCommandName:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_name[int32(fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_REMOVE_ELEMENT)],
		FirstParameter:           elementId,
		SecondParameter:          testCaseModel.NotApplicable,
		UserId:                   commandAndRuleEngine.testcases.CurrentUser,
		CommandExecutedTimeStamp: timestamppb.Now(),
	}

	// Extract the TestCaseModel
	currentTestCaseModel, existsInMap := commandAndRuleEngine.testcases.TestCases[testCaseUuid]
	if existsInMap == false {
		errorId := "9e42e135-e5c3-479c-8a09-0e33213a68d1"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))
	}

	// Add command to command stack
	currentTestCaseModel.CommandStack = append(currentTestCaseModel.CommandStack, newCommandEntry)

	// Add the TestCaseModel back into map of all TestCaseModels
	commandAndRuleEngine.testcases.TestCases[testCaseUuid] = currentTestCaseModel

	// Add command Textual representations to Textual Representation Stack
	textualRepresentationSimple, textualRepresentationComplex, err := commandAndRuleEngine.testcases.CreateTextualTestCase(testCaseUuid)

	if err == nil {
		currentTestCaseModel.TextualTestCaseRepresentationSimpleStack = append(currentTestCaseModel.TextualTestCaseRepresentationSimpleStack, textualRepresentationSimple)
		currentTestCaseModel.TextualTestCaseRepresentationComplexStack = append(currentTestCaseModel.TextualTestCaseRepresentationComplexStack, textualRepresentationComplex)

		// Add the TestCaseModel back into map of all TestCaseModels
		commandAndRuleEngine.testcases.TestCases[testCaseUuid] = currentTestCaseModel
	}

	return err
}

// TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_NEW_ELEMENT
// Used for Swapping out an element, and in an element structure, from a TestCaseModel that is used within a TestCase
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_SwapOutElemenAndInNewElementInTestCaseModel(testCaseUuid string, elementToSwapOutUuid string, immatureElementToSwapIn *immatureElementStruct) (err error) {

	// Try to Swap out element
	err = commandAndRuleEngine.executeSwapElement(testCaseUuid, elementToSwapOutUuid, immatureElementToSwapIn)

	// Exit if the element couldn't be deleted
	if err != nil {
		return err
	}

	// Create a new Command
	newCommandEntry := fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage_TestCaseModelCommandMessage{
		TestCaseCommandType:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_NEW_ELEMENT,
		TestCaseCommandName:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_name[int32(fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_NEW_ELEMENT)],
		FirstParameter:           elementToSwapOutUuid,
		SecondParameter:          immatureElementToSwapIn.firstElementUuid,
		UserId:                   commandAndRuleEngine.testcases.CurrentUser,
		CommandExecutedTimeStamp: timestamppb.Now(),
	}

	// Extract the TestCaseModel
	currentTestCaseModel, existsInMap := commandAndRuleEngine.testcases.TestCases[testCaseUuid]
	if existsInMap == false {
		errorId := "73cf671c-79e7-4a5e-8f42-d39cd86d94c9"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))
	}

	// Add command to command stack
	currentTestCaseModel.CommandStack = append(currentTestCaseModel.CommandStack, newCommandEntry)

	// Add the TestCaseModel back into map of all TestCaseModels
	// commandAndRuleEngine.testcases.TestCases[testCaseUuid] = currentTestCaseModel

	// Add command Textual representations to Textual Representation Stack
	textualRepresentationSimple, textualRepresentationComplex, err := commandAndRuleEngine.testcases.CreateTextualTestCase(testCaseUuid)

	// If no errors then add Simple and Complex Textual Representation to their stacks
	if err == nil {
		currentTestCaseModel.TextualTestCaseRepresentationSimpleStack = append(currentTestCaseModel.TextualTestCaseRepresentationSimpleStack, textualRepresentationSimple)
		currentTestCaseModel.TextualTestCaseRepresentationComplexStack = append(currentTestCaseModel.TextualTestCaseRepresentationComplexStack, textualRepresentationComplex)

		// Add the TestCaseModel back into map of all TestCaseModels
		commandAndRuleEngine.testcases.TestCases[testCaseUuid] = currentTestCaseModel
	}

	return err
}

// TestCaseCommandTypeEnum_COPY_ELEMENT
// Used for copying an element  in a TestCaseModel that is used within a TestCase
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_CopyElementInTestCaseModel() (err error) {

	return err
}

// TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_COPY_BUFFER_ELEMENT
// Used for Swapping in an element from Copy Buffer in a TestCaseModel that is used within a TestCase
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_SwapInElementFromCopyBufferInTestCaseModel() (err error) {

	return err
}

// TestCaseCommandTypeEnum_CUT_ELEMENT
// Used for cutting an element in a TestCaseModel that is used within a TestCase
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_CutElementInTestCaseModel() (err error) {

	return err
}

// TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_CUT_BUFFER_ELEMENT
// Used for Swapping in an element from Cut opy Buffer in a TestCaseModel that is used within a TestCase
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_SwapInElementFromCutBufferInTestCaseModel() (err error) {

	return err
}
