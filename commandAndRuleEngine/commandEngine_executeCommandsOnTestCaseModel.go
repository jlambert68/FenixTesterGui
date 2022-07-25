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
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases", testCaseUuid))
	}

	// Add command to command stack
	currentTestCaseModel.CommandStack = append(currentTestCaseModel.CommandStack, newCommandEntry)

	// Add the TestCaseModel back into map of all TestCaseModels
	commandAndRuleEngine.testcases.TestCases[testCaseUuid] = currentTestCaseModel

	return err
}

// TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_NEW_ELEMENT
// Used for Swapping out an element, and in an element structure, from a TestCaseModel that is used within a TestCase
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_SwapOutElementFromTestCaseModel() (err error) {

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
