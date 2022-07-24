package commandAndRuleEngine

import (
	"FenixTesterGui/testCase/testCaseModel"
	//"errors"
	uuidGenerator "github.com/google/uuid"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
)

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

	// Generate new TestCase-UUID
	testCaseUuid = uuidGenerator.New().String()

	// Add the TestCaseModel into map of all TestCaseModels
	commandAndRuleEngine.testcases.TestCases[testCaseUuid] = newTestCaseModel

	return testCaseUuid, nil

}
