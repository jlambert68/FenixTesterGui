package gui

import (
	"FenixTesterGui/grpc_out"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"

	"testing"

	"FenixTesterGui/gui/UnitTestTestData"
	"github.com/stretchr/testify/assert"
)

const printValues = false

//var availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct

// Checks that Available TestInstructions are put in Available Building Blocks-testCaseModel in a correct way
func TestLoadModelWithAvailableBuildingBlocksRegardingTestInstructions(t *testing.T) {

	var testInstructionsAndTestContainersMessage fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage
	//testInstructionsAndTestContainersMessage = &fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionsAndTestContainersMessage

	//res := fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionsAndTestContainersMessage{}
	//json.Unmarshal(byt, &res)
	mystring := UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_ABB001

	if err := jsonpb.UnmarshalString(mystring, &testInstructionsAndTestContainersMessage); err != nil {
		panic(err)
	}

	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Verify That TestData is using correct proto-version
	returnMessage := UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion(testInstructionsAndTestContainersMessage.AckNackResponse.ProtoFileVersionUsedByClient)
	if returnMessage != nil && returnMessage.AckNack == false {
		myLogger.WithFields(logrus.Fields{
			"Id": "c9453631-46b0-47a0-86fa-fe2e5b51ed92",
		}).Fatalln("Exiting because of wrong Proto-file version in TestData")
	}

	var availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct

	availableBuildingBlocksModel = &AvailableBuildingBlocksModelStruct{
		logger:                             myLogger,
		fenixGuiBuilderServerAddressToDial: "",
		fullDomainTestInstructionTypeTestInstructionRelationsMap:                   nil,
		fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap: nil,
		AvailableBuildingBlocksForUITreeNodes:                                      nil,
		grpcOut:                                                                    grpc_out.GRPCOutStruct{},
		availableBuildingBlockModelSuitedForFyneTreeView:                           nil,
	}

	// Initiate map
	availableBuildingBlocksModel.allBuildingBlocksTreeNameToUuid = make(map[string]uiTreeNodesNameToUuidStruct)

	availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocksRegardingTestInstructions(&testInstructionsAndTestContainersMessage)

	// fmt.Println(availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap)
	/*
		fmt.Println(availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes)
		b, err := json.Marshal(&availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes)

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b))

	*/

	// Validate 'availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes'
	assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_001, fmt.Sprint(availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes))
	if printValues {
		fmt.Println("TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_001")
		fmt.Println(availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes)
	}

	// Validate 'availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap'
	assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_002, fmt.Sprint(availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap))
	if printValues {
		fmt.Println("TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_002")
		fmt.Println(availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap)
	}
}

// Checks that Available TestInstructionContainers are put in Available Building Blocks-testCaseModel in a correct way
func TestLoadModelWithAvailableBuildingBlocksRegardingTestInstructionContainers(t *testing.T) {

	var testInstructionsAndTestContainersMessage fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage

	// Load data into gRPC-meTestLoadModelWithAvailableBuildingBlocksRegardingTestInstructionContainerssssage response (from DB)
	if err := jsonpb.UnmarshalString(UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_ABB001, &testInstructionsAndTestContainersMessage); err != nil {
		panic(err)
	}

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Verify That TestData is using correct proto-version
	returnMessage := UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion(testInstructionsAndTestContainersMessage.AckNackResponse.ProtoFileVersionUsedByClient)
	if returnMessage != nil && returnMessage.AckNack == false {
		myLogger.WithFields(logrus.Fields{
			"Id": "715600c1-9ad0-4c6e-adc1-ae79b25dd810",
		}).Fatalln("Exiting because of wrong Proto-file version in TestData")
	}

	// Initiate availableBuildingBlocksModel
	var availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct

	availableBuildingBlocksModel = &AvailableBuildingBlocksModelStruct{
		logger:                             myLogger,
		fenixGuiBuilderServerAddressToDial: "",
		fullDomainTestInstructionTypeTestInstructionRelationsMap:                   nil,
		fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap: nil,
		AvailableBuildingBlocksForUITreeNodes:                                      nil,
		grpcOut:                                                                    grpc_out.GRPCOutStruct{},
		availableBuildingBlockModelSuitedForFyneTreeView:                           nil,
	}

	// Initiate map
	availableBuildingBlocksModel.allBuildingBlocksTreeNameToUuid = make(map[string]uiTreeNodesNameToUuidStruct)

	// Load Available Building Blocks, in this case TestInstructionContainers
	availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocksRegardingTestInstructionContainers(&testInstructionsAndTestContainersMessage)

	//fmt.Println(availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap)
	/*
		fmt.Println(availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes)
		b, err := json.Marshal(&availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes)

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b))

	*/

	//fmt.Println(UnitTestTestData.TestInstructionsAndTestInstructionsRespons_ABB001_ExpectedResultInModel_003)
	//fmt.Println(UnitTestTestData.TestInstructionsAndTestInstructionsRespons_ABB001_ExpectedResultInModel_004)

	// Validate 'availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes'
	assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_003, fmt.Sprint(availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes))
	if printValues {
		fmt.Println("TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_003")
		fmt.Println(availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes)
	}

	// Validate 'availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap'
	assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_004, fmt.Sprint(availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap))
	if printValues {
		fmt.Println("TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_004")
		fmt.Println(availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap)
	}
}

// Checks that Available TestInstructions & TestInstructionContainers are put in Available Building Blocks-testCaseModel in a correct way
func TestLoadModelWithAvailableBuildingBlocks(t *testing.T) {

	var testInstructionsAndTestContainersMessage fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage

	// Load data into gRPC-message response (from DB)
	if err := jsonpb.UnmarshalString(UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_ABB001, &testInstructionsAndTestContainersMessage); err != nil {
		panic(err)
	}

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Verify That TestData is using correct proto-version
	returnMessage := UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion(testInstructionsAndTestContainersMessage.AckNackResponse.ProtoFileVersionUsedByClient)
	if returnMessage != nil && returnMessage.AckNack == false {
		myLogger.WithFields(logrus.Fields{
			"Id": "c9453631-46b0-47a0-86fa-fe2e5b51ed92",
		}).Fatalln("Exiting because of wrong Proto-file version in TestData")
	}

	// Initiate availableBuildingBlocksModel
	var availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct

	availableBuildingBlocksModel = &AvailableBuildingBlocksModelStruct{
		logger:                             myLogger,
		fenixGuiBuilderServerAddressToDial: "",
		fullDomainTestInstructionTypeTestInstructionRelationsMap:                   nil,
		fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap: nil,
		AvailableBuildingBlocksForUITreeNodes:                                      nil,
		grpcOut:                                                                    grpc_out.GRPCOutStruct{},
		availableBuildingBlockModelSuitedForFyneTreeView:                           nil,
	}

	var assertResult bool

	// Load Available Building Blocks, in this case TestInstructionContainers
	availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocks(&testInstructionsAndTestContainersMessage)

	// Validate 'availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes'
	assertResult = assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_005, fmt.Sprint(availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes))
	if printValues && !assertResult {
		fmt.Println("TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_005")
		fmt.Println(availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes)
	}

	// Validate 'availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap'
	assertResult = assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_006, fmt.Sprint(availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap))
	if printValues && !assertResult {
		fmt.Println("TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_006")
		fmt.Println(availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap)
	}

	// Validate 'availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap'
	assertResult = assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_007, fmt.Sprint(availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap))
	if printValues && !assertResult {
		fmt.Println("TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_007")
		fmt.Println(availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap)
	}

}

// Checks that Pinned TestInstructions are put in Pinned Building Blocks-testCaseModel in a correct way
func TestLoadModelWithPinnedBuildingBlocksRegardingTestInstructions(t *testing.T) {

	var pinnedTestInstructionsAndTestContainersMessage fenixGuiTestCaseBuilderServerGrpcApi.AvailablePinnedTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage

	//  Needed to support test
	var testInstructionsAndTestContainersMessage fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage

	// Load data into gRPC-message response (from DB) - needed to support test
	if err := jsonpb.UnmarshalString(UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_ABB001, &testInstructionsAndTestContainersMessage); err != nil {
		panic(err)
	}

	// Load data into gRPC-message response (from DB)
	if err := jsonpb.UnmarshalString(UnitTestTestData.PinnedTestInstructionsAndTestInstructionsContainersRespons_PBB001, &pinnedTestInstructionsAndTestContainersMessage); err != nil {
		panic(err)
	}

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Verify That TestData is using correct proto-version
	returnMessage := UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion(testInstructionsAndTestContainersMessage.AckNackResponse.ProtoFileVersionUsedByClient)
	if returnMessage != nil && returnMessage.AckNack == false {
		myLogger.WithFields(logrus.Fields{
			"Id": "c2c66337-65e8-472b-86a3-a1f4628dcccc",
		}).Fatalln("Exiting because of wrong Proto-file version in TestData")
	}

	returnMessage = UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion(pinnedTestInstructionsAndTestContainersMessage.AckNackResponse.ProtoFileVersionUsedByClient)
	if returnMessage != nil && returnMessage.AckNack == false {
		myLogger.WithFields(logrus.Fields{
			"Id": "ec2fe8d6-1ce4-4923-a074-3c2f9482c4a6",
		}).Fatalln("Exiting because of wrong Proto-file version in TestData")
	}

	// Initiate availableBuildingBlocksModel
	var availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct

	availableBuildingBlocksModel = &AvailableBuildingBlocksModelStruct{
		logger:                             myLogger,
		fenixGuiBuilderServerAddressToDial: "",
		fullDomainTestInstructionTypeTestInstructionRelationsMap:                   nil,
		fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap: nil,
		AvailableBuildingBlocksForUITreeNodes:                                      nil,
		grpcOut:                                                                    grpc_out.GRPCOutStruct{},
		availableBuildingBlockModelSuitedForFyneTreeView:                           nil,
		pinnedBuildingBlocksForUITreeNodes:                                         nil,
	}

	// Clear and initiate variables
	availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes = make(map[string]uiTreeNodesNameToUuidStruct)

	// Initiate map
	availableBuildingBlocksModel.allBuildingBlocksTreeNameToUuid = make(map[string]uiTreeNodesNameToUuidStruct)

	// Load Available Building Blocks - needed for the test should crash
	availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocksRegardingTestInstructions(&testInstructionsAndTestContainersMessage)

	// Load Pinned Building Blocks
	availableBuildingBlocksModel.loadModelWithPinnedBuildingBlocksRegardingTestInstructions(&pinnedTestInstructionsAndTestContainersMessage)

	// Validate 'availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes'
	assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_PBB001_ExpectedResultInModel_001, fmt.Sprint(availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes))

}

// Checks that Pinned TestInstructionContainers are put in Pinned Building Blocks-testCaseModel in a correct way
func TestLoadModelWithPinnedBuildingBlocksRegardingTestInstructionContainers(t *testing.T) {

	var pinnedTestInstructionsAndTestContainersMessage fenixGuiTestCaseBuilderServerGrpcApi.AvailablePinnedTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage

	//  Needed to support test
	var testInstructionsAndTestContainersMessage fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage

	// Load data into gRPC-message response (from DB) - needed to support test
	if err := jsonpb.UnmarshalString(UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_ABB001, &testInstructionsAndTestContainersMessage); err != nil {
		panic(err)
	}

	// Load data into gRPC-message response (from DB)
	if err := jsonpb.UnmarshalString(UnitTestTestData.PinnedTestInstructionsAndTestInstructionsContainersRespons_PBB001, &pinnedTestInstructionsAndTestContainersMessage); err != nil {
		panic(err)
	}

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Verify That TestData is using correct proto-version
	returnMessage := UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion(testInstructionsAndTestContainersMessage.AckNackResponse.ProtoFileVersionUsedByClient)
	if returnMessage != nil && returnMessage.AckNack == false {
		myLogger.WithFields(logrus.Fields{
			"Id": "70481f24-cc0a-4e5e-b4e3-cc3c74985eb3",
		}).Fatalln("Exiting because of wrong Proto-file version in TestData")
	}

	returnMessage = UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion(pinnedTestInstructionsAndTestContainersMessage.AckNackResponse.ProtoFileVersionUsedByClient)
	if returnMessage != nil && returnMessage.AckNack == false {
		myLogger.WithFields(logrus.Fields{
			"Id": "8148ab5e-a31f-492a-8f23-ff442563aaf6",
		}).Fatalln("Exiting because of wrong Proto-file version in TestData")
	}

	// Initiate availableBuildingBlocksModel
	var availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct

	availableBuildingBlocksModel = &AvailableBuildingBlocksModelStruct{
		logger:                             myLogger,
		fenixGuiBuilderServerAddressToDial: "",
		fullDomainTestInstructionTypeTestInstructionRelationsMap:                   nil,
		fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap: nil,
		AvailableBuildingBlocksForUITreeNodes:                                      nil,
		grpcOut:                                                                    grpc_out.GRPCOutStruct{},
		availableBuildingBlockModelSuitedForFyneTreeView:                           nil,
		pinnedBuildingBlocksForUITreeNodes:                                         nil,
	}

	// Clear and initiate variables
	availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes = make(map[string]uiTreeNodesNameToUuidStruct)

	// Initiate map
	availableBuildingBlocksModel.allBuildingBlocksTreeNameToUuid = make(map[string]uiTreeNodesNameToUuidStruct)

	// Load Available Building Blocks - needed for the test should crash
	availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocksRegardingTestInstructionContainers(&testInstructionsAndTestContainersMessage)

	// Load Pinned Building Blocks
	availableBuildingBlocksModel.loadModelWithPinnedBuildingBlocksRegardingTestInstructionContainers(&pinnedTestInstructionsAndTestContainersMessage)

	// Validate 'availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes'
	assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_PBB001_ExpectedResultInModel_002, fmt.Sprint(availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes))

}

// Checks that Pinned TestInstruction And TestInstructionContainers are put in Pinned Building Blocks-testCaseModel in a correct way
func TestLoadModelWithPinnedBuildingBlocksRegardingTestInstructionAndTestInstructionContainers(t *testing.T) {

	var pinnedTestInstructionsAndTestContainersMessage fenixGuiTestCaseBuilderServerGrpcApi.AvailablePinnedTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage

	//  Needed to support test
	var testInstructionsAndTestContainersMessage fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage

	// Load data into gRPC-message response (from DB) - needed to support test
	if err := jsonpb.UnmarshalString(UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_ABB001, &testInstructionsAndTestContainersMessage); err != nil {
		panic(err)
	}

	// Load data into gRPC-message response (from DB)
	if err := jsonpb.UnmarshalString(UnitTestTestData.PinnedTestInstructionsAndTestInstructionsContainersRespons_PBB001, &pinnedTestInstructionsAndTestContainersMessage); err != nil {
		panic(err)
	}

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Verify That TestData is using correct proto-version
	returnMessage := UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion(testInstructionsAndTestContainersMessage.AckNackResponse.ProtoFileVersionUsedByClient)
	if returnMessage != nil && returnMessage.AckNack == false {
		myLogger.WithFields(logrus.Fields{
			"Id": "e6a6812f-9132-4f40-9478-8a7c226ab84f",
		}).Fatalln("Exiting because of wrong Proto-file version in TestData")
	}

	returnMessage = UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion(pinnedTestInstructionsAndTestContainersMessage.AckNackResponse.ProtoFileVersionUsedByClient)
	if returnMessage != nil && returnMessage.AckNack == false {
		myLogger.WithFields(logrus.Fields{
			"Id": "a5869787-8b8e-4248-9100-e2fd7a3eeb1a",
		}).Fatalln("Exiting because of wrong Proto-file version in TestData")
	}

	// Initiate availableBuildingBlocksModel
	var availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct

	availableBuildingBlocksModel = &AvailableBuildingBlocksModelStruct{
		logger:                             myLogger,
		fenixGuiBuilderServerAddressToDial: "",
		fullDomainTestInstructionTypeTestInstructionRelationsMap:                   nil,
		fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap: nil,
		AvailableBuildingBlocksForUITreeNodes:                                      nil,
		grpcOut:                                                                    grpc_out.GRPCOutStruct{},
		availableBuildingBlockModelSuitedForFyneTreeView:                           nil,
		pinnedBuildingBlocksForUITreeNodes:                                         nil,
	}

	// Load Available Building Blocks - needed for the test should crash
	availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocks(&testInstructionsAndTestContainersMessage)

	// Load Pinned Building Blocks
	availableBuildingBlocksModel.loadModelWithPinnedBuildingBlocks(&pinnedTestInstructionsAndTestContainersMessage)

	// Validate 'availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes'
	assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_PBB001_ExpectedResultInModel_003, fmt.Sprint(availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes))

}
