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
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_NewTestCaseModel() (testCaseUuid string, err error) {

	// Create new B0-Bind
	b0Bond := commandAndRuleEngine.createNewBondB0Element()

	// Initiate a TestCaseModel
	newTestCaseModel := testCaseModel.TestCaseModelStruct{
		LastLoadedTestCaseModelGRPCMessage: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage{},
		FirstElementUuid:                   "",
		TestCaseModelMap:                   nil,
	}

	// Initiate TestCaseModel-map
	newTestCaseModel.TestCaseModelMap = make(map[string]testCaseModel.MatureTestCaseModelElementStruct)

	// Add B0-bond to the TestCaseModel-map
	newTestCaseModel.TestCaseModelMap[b0Bond.MatureElementUuid] = testCaseModel.MatureTestCaseModelElementStruct{
		MatureTestCaseModelElementMessage:  b0Bond,
		MatureTestCaseModelElementMetaData: testCaseModel.MatureTestCaseModelElementMetaDataStruct{},
	}

	// Set the B0-bond as first element in TestCaseModel
	newTestCaseModel.FirstElementUuid = b0Bond.MatureElementUuid

	// Create a new Command
	newCommandEntry := fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage_TestCaseModelCommandMessage{
		TestCaseCommandType:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_NEW_TESTCASE,
		TestCaseCommandName:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_name[int32(fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_NEW_TESTCASE)],
		FirstParameter:           testCaseModel.NotApplicable,
		SecondParameter:          testCaseModel.NotApplicable,
		UserId:                   commandAndRuleEngine.Testcases.CurrentUser,
		CommandExecutedTimeStamp: timestamppb.Now(),
	}

	// Add command to command stack
	newTestCaseModel.CommandStack = append(newTestCaseModel.CommandStack, newCommandEntry)

	// Generate new TestCase-UUID
	testCaseUuid = uuidGenerator.New().String()

	// If TestCases-map is not initialized then do that
	if commandAndRuleEngine.Testcases.TestCases == nil {
		commandAndRuleEngine.Testcases.TestCases = make(map[string]testCaseModel.TestCaseModelStruct)
	}

	// Add the TestCaseModel into map of all TestCaseModels
	commandAndRuleEngine.Testcases.TestCases[testCaseUuid] = newTestCaseModel

	// Add command Textual representations to Textual Representation Stack
	textualRepresentationSimple, textualRepresentationComplex, textualRepresentationExtended, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)

	if err == nil {
		newTestCaseModel.TextualTestCaseRepresentationSimpleStack = append(newTestCaseModel.TextualTestCaseRepresentationSimpleStack, textualRepresentationSimple)
		newTestCaseModel.TextualTestCaseRepresentationComplexStack = append(newTestCaseModel.TextualTestCaseRepresentationComplexStack, textualRepresentationComplex)
		newTestCaseModel.TextualTestCaseRepresentationExtendedStack = append(newTestCaseModel.TextualTestCaseRepresentationExtendedStack, textualRepresentationExtended)

		// Add the TestCaseModel back into map of all TestCaseModels
		commandAndRuleEngine.Testcases.TestCases[testCaseUuid] = newTestCaseModel

	}

	return testCaseUuid, err

}

// TestCaseCommandTypeEnum_REMOVE_ELEMENT
// Used for Deleting an element from a TestCaseModel that is used within a TestCase
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel(testCaseUuid string, elementId string) (err error) {

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
		UserId:                   commandAndRuleEngine.Testcases.CurrentUser,
		CommandExecutedTimeStamp: timestamppb.Now(),
	}

	// Extract the TestCaseModel
	currentTestCaseModel, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	if existsInMap == false {
		errorId := "9e42e135-e5c3-479c-8a09-0e33213a68d1"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))
	}

	// Add command to command stack
	currentTestCaseModel.CommandStack = append(currentTestCaseModel.CommandStack, newCommandEntry)

	// Add the TestCaseModel back into map of all TestCaseModels
	commandAndRuleEngine.Testcases.TestCases[testCaseUuid] = currentTestCaseModel

	// Add command Textual representations to Textual Representation Stack
	textualRepresentationSimple, textualRepresentationComplex, textualRepresentationExtended, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)

	if err == nil {
		currentTestCaseModel.TextualTestCaseRepresentationSimpleStack = append(currentTestCaseModel.TextualTestCaseRepresentationSimpleStack, textualRepresentationSimple)
		currentTestCaseModel.TextualTestCaseRepresentationComplexStack = append(currentTestCaseModel.TextualTestCaseRepresentationComplexStack, textualRepresentationComplex)
		currentTestCaseModel.TextualTestCaseRepresentationExtendedStack = append(currentTestCaseModel.TextualTestCaseRepresentationExtendedStack, textualRepresentationExtended)

		// Add the TestCaseModel back into map of all TestCaseModels
		commandAndRuleEngine.Testcases.TestCases[testCaseUuid] = currentTestCaseModel
	}

	return err
}

// TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_NEW_ELEMENT
// Used for Swapping out an element, and in an element structure, from a TestCaseModel that is used within a TestCase
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_SwapOutElemenAndInNewElementInTestCaseModel(testCaseUuid string, elementToSwapOutUuid string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (err error) {

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
		SecondParameter:          immatureElementToSwapIn.FirstElementUuid,
		UserId:                   commandAndRuleEngine.Testcases.CurrentUser,
		CommandExecutedTimeStamp: timestamppb.Now(),
	}

	// Extract the TestCaseModel
	currentTestCaseModel, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	if existsInMap == false {
		errorId := "73cf671c-79e7-4a5e-8f42-d39cd86d94c9"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))
	}

	// Add command to command stack
	currentTestCaseModel.CommandStack = append(currentTestCaseModel.CommandStack, newCommandEntry)

	// Add the TestCaseModel back into map of all TestCaseModels
	// commandAndRuleEngine.Testcases.TestCases[testCaseUuid] = currentTestCaseModel

	// Add command Textual representations to Textual Representation Stack
	textualRepresentationSimple, textualRepresentationComplex, textualRepresentationExtended, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)

	// If no errors then add Simple and Complex Textual Representation to their stacks
	if err == nil {
		currentTestCaseModel.TextualTestCaseRepresentationSimpleStack = append(currentTestCaseModel.TextualTestCaseRepresentationSimpleStack, textualRepresentationSimple)
		currentTestCaseModel.TextualTestCaseRepresentationComplexStack = append(currentTestCaseModel.TextualTestCaseRepresentationComplexStack, textualRepresentationComplex)
		currentTestCaseModel.TextualTestCaseRepresentationExtendedStack = append(currentTestCaseModel.TextualTestCaseRepresentationExtendedStack, textualRepresentationExtended)

		// Add the TestCaseModel back into map of all TestCaseModels
		commandAndRuleEngine.Testcases.TestCases[testCaseUuid] = currentTestCaseModel
	}

	return err
}

// TestCaseCommandTypeEnum_COPY_ELEMENT
// Used for copying an element  in a TestCaseModel that is used within a TestCase
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_CopyElementInTestCaseModel(testCaseUuid string, elementIdToCopy string) (err error) {

	// Try to Copy element to Copy Buffer
	err = commandAndRuleEngine.executeCopyElement(testCaseUuid, elementIdToCopy)

	// Exit if the element couldn't be deleted
	if err != nil {
		return err
	}

	// Create a new Command
	newCommandEntry := fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage_TestCaseModelCommandMessage{
		TestCaseCommandType:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_COPY_ELEMENT,
		TestCaseCommandName:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_name[int32(fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_COPY_ELEMENT)],
		FirstParameter:           elementIdToCopy,
		SecondParameter:          testCaseModel.NotApplicable,
		UserId:                   commandAndRuleEngine.Testcases.CurrentUser,
		CommandExecutedTimeStamp: timestamppb.Now(),
	}

	// Extract the TestCaseModel
	currentTestCaseModel, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	if existsInMap == false {
		errorId := "2d6af5bd-5a1b-4cc0-b3e7-da21b5928c4f"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))
	}

	// Add command to command stack
	currentTestCaseModel.CommandStack = append(currentTestCaseModel.CommandStack, newCommandEntry)

	// If no errors then add the TestCaseModel back into map of all TestCaseModels
	if err == nil {
		commandAndRuleEngine.Testcases.TestCases[testCaseUuid] = currentTestCaseModel
	}

	return err
}

// TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_COPY_BUFFER_ELEMENT
// Used for Swapping in an element from Copy Buffer in a TestCaseModel that is used within a TestCase
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_SwapInElementFromCopyBufferInTestCaseModel(testCaseUuid string, elementIdToBeReplacedByCopyBuffer string) (err error) {

	// Try to Swap Element From Copy Buffer
	err = commandAndRuleEngine.executeSwapElementForCopyBuffer(testCaseUuid, elementIdToBeReplacedByCopyBuffer)

	// Exit if the element couldn't be deleted
	if err != nil {
		return err
	}

	// Create a new Command
	newCommandEntry := fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage_TestCaseModelCommandMessage{
		TestCaseCommandType:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_COPY_BUFFER_ELEMENT,
		TestCaseCommandName:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_name[int32(fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_COPY_BUFFER_ELEMENT)],
		FirstParameter:           elementIdToBeReplacedByCopyBuffer,
		SecondParameter:          testCaseModel.NotApplicable,
		UserId:                   commandAndRuleEngine.Testcases.CurrentUser,
		CommandExecutedTimeStamp: timestamppb.Now(),
	}

	// Extract the TestCaseModel
	currentTestCaseModel, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	if existsInMap == false {
		errorId := "10ef5496-d92e-4e35-af41-e16c51c7df71"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))
	}

	// Add command to command stack
	currentTestCaseModel.CommandStack = append(currentTestCaseModel.CommandStack, newCommandEntry)

	// Indicate that Cut command has been initiated
	currentTestCaseModel.CutCommandInitiated = true

	// If no errors then add the TestCaseModel back into map of all TestCaseModels
	if err == nil {
		commandAndRuleEngine.Testcases.TestCases[testCaseUuid] = currentTestCaseModel
	}

	return err
}

// TestCaseCommandTypeEnum_CUT_ELEMENT
// Used for cutting an element in a TestCaseModel that is used within a TestCase
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_CutElementInTestCaseModel(testCaseUuid string, elementIdToCut string) (err error) {

	// Try to Copy element to Cut Buffer
	err = commandAndRuleEngine.executeCutElement(testCaseUuid, elementIdToCut)

	// Exit if the element couldn't be deleted
	if err != nil {
		return err
	}

	// Create a new Command
	newCommandEntry := fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage_TestCaseModelCommandMessage{
		TestCaseCommandType:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_CUT_ELEMENT,
		TestCaseCommandName:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_name[int32(fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_CUT_ELEMENT)],
		FirstParameter:           elementIdToCut,
		SecondParameter:          testCaseModel.NotApplicable,
		UserId:                   commandAndRuleEngine.Testcases.CurrentUser,
		CommandExecutedTimeStamp: timestamppb.Now(),
	}

	// Extract the TestCaseModel
	currentTestCaseModel, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	if existsInMap == false {
		errorId := "dc1cd5d3-e809-4465-aeda-cdf6ec44070f"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))
	}

	// Add command to command stack
	currentTestCaseModel.CommandStack = append(currentTestCaseModel.CommandStack, newCommandEntry)

	// If no errors then add the TestCaseModel back into map of all TestCaseModels
	if err == nil {
		commandAndRuleEngine.Testcases.TestCases[testCaseUuid] = currentTestCaseModel
	}

	return err
}

// TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_CUT_BUFFER_ELEMENT
// Used for Swapping in an element from Cut opy Buffer in a TestCaseModel that is used within a TestCase
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_SwapInElementFromCutBufferInTestCaseModel(testCaseUuid string, uuidToReplacedByCutBufferContent string) (err error) {

	// Try to Swap Element From Cut Buffer
	err = commandAndRuleEngine.executeSwapElementFromCutBuffer(testCaseUuid, uuidToReplacedByCutBufferContent, nil)

	// Exit if the element couldn't be Swapped in from Cut
	if err != nil {
		return err
	}

	// Create a new Command
	newCommandEntry := fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage_TestCaseModelCommandMessage{
		TestCaseCommandType:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_CUT_BUFFER_ELEMENT,
		TestCaseCommandName:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_name[int32(fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_CUT_BUFFER_ELEMENT)],
		FirstParameter:           uuidToReplacedByCutBufferContent,
		SecondParameter:          testCaseModel.NotApplicable,
		UserId:                   commandAndRuleEngine.Testcases.CurrentUser,
		CommandExecutedTimeStamp: timestamppb.Now(),
	}

	// Extract the TestCaseModel
	currentTestCaseModel, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	if existsInMap == false {
		errorId := "e1f7b09a-1867-4c0d-a02a-2b513788d711"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))
	}

	// Add command to command stack
	currentTestCaseModel.CommandStack = append(currentTestCaseModel.CommandStack, newCommandEntry)

	// Indicate that Cut command has been finished, or is not active anymore
	currentTestCaseModel.CutCommandInitiated = false

	// Clear Cut Buffer
	newEmptyCutBufferContet := testCaseModel.MatureElementStruct{
		FirstElementUuid: "",
		MatureElementMap: nil,
	}
	currentTestCaseModel.CutBuffer = newEmptyCutBufferContet

	// If no errors then add the TestCaseModel back into map of all TestCaseModels
	if err == nil {
		commandAndRuleEngine.Testcases.TestCases[testCaseUuid] = currentTestCaseModel
	}

	return err
}
