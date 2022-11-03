package gui

import (
	"FenixTesterGui/grpc_out_GuiTestCaseBuilderServer"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"

	"testing"

	"FenixTesterGui/gui/UnitTestTestData"
	"github.com/stretchr/testify/assert"
)

//const printValues = false

//var availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct

// Checks that a non-existing Building Block can't be pinned
func TestThatNonExistingBuildBlockCanBePinned(t *testing.T) {

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
			"Id": "4757ab26-60ef-4c59-b299-71e32ae74ed7",
		}).Fatalln("Exiting because of wrong Proto-file version in TestData")
	}

	returnMessage = UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion(pinnedTestInstructionsAndTestContainersMessage.AckNackResponse.ProtoFileVersionUsedByClient)
	if returnMessage != nil && returnMessage.AckNack == false {
		myLogger.WithFields(logrus.Fields{
			"Id": "2c22eadb-cf09-495d-8b89-00443d16f4ce",
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
		grpcOut:                                                                    grpc_out_GuiTestCaseBuilderServer.GRPCOutGuiTestCaseBuilderServerStruct{},
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

	// Validate that an non-existing Building Block can't be pinned
	err := availableBuildingBlocksModel.pinTestInstructionOrTestInstructionContainer("NonExistingBuildingBlock")

	assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_PBB002_ExpectedResultInModel_001, fmt.Sprint(err))

}

// Checks that an already pinned Building Block can be pinned
func TestThatAlreadyPinnedBuildingBlockNotCanBePinned(t *testing.T) {

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
			"Id": "df1285df-db49-404f-9e8b-a9549aff0f74",
		}).Fatalln("Exiting because of wrong Proto-file version in TestData")
	}

	returnMessage = UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion(pinnedTestInstructionsAndTestContainersMessage.AckNackResponse.ProtoFileVersionUsedByClient)
	if returnMessage != nil && returnMessage.AckNack == false {
		myLogger.WithFields(logrus.Fields{
			"Id": "03d024d1-3fdf-4d2d-84e6-56bb09fcb736",
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
		grpcOut:                                                                    grpc_out_GuiTestCaseBuilderServer.GRPCOutGuiTestCaseBuilderServerStruct{},
		availableBuildingBlockModelSuitedForFyneTreeView:                           nil,
	}

	// Clear and initiate variables
	availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes = make(map[string]uiTreeNodesNameToUuidStruct)

	// Initiate map
	availableBuildingBlocksModel.allBuildingBlocksTreeNameToUuid = make(map[string]uiTreeNodesNameToUuidStruct)

	// Load Available Building Blocks, in this case TestInstructionContainers
	availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocks(&testInstructionsAndTestContainersMessage)

	// Load Pinned Building Blocks
	availableBuildingBlocksModel.loadModelWithPinnedBuildingBlocks(&pinnedTestInstructionsAndTestContainersMessage)

	// Validate that an already pinned  Building Block can't be pinned again
	pinnedBuildingBlockName := availableBuildingBlocksModel.getPinnedBuildingBlocksTreeNamesFromModel()[0]
	pinnedBuildingBlockUuid := availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes[pinnedBuildingBlockName].uuid
	buildingBlockName := availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[pinnedBuildingBlockUuid].nameInUITree
	err := availableBuildingBlocksModel.pinTestInstructionOrTestInstructionContainer(buildingBlockName)

	assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_PBB002_ExpectedResultInModel_002, fmt.Sprint(err))

}

// Checks that a non-pinned Building Block can be pinned
func TestToPinBuildingBlockCanBePinned(t *testing.T) {

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
			"Id": "df1285df-db49-404f-9e8b-a9549aff0f74",
		}).Fatalln("Exiting because of wrong Proto-file version in TestData")
	}

	returnMessage = UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion(pinnedTestInstructionsAndTestContainersMessage.AckNackResponse.ProtoFileVersionUsedByClient)
	if returnMessage != nil && returnMessage.AckNack == false {
		myLogger.WithFields(logrus.Fields{
			"Id": "03d024d1-3fdf-4d2d-84e6-56bb09fcb736",
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
		grpcOut:                                                                    grpc_out_GuiTestCaseBuilderServer.GRPCOutGuiTestCaseBuilderServerStruct{},
		availableBuildingBlockModelSuitedForFyneTreeView:                           nil,
	}

	// Clear and initiate variables
	availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes = make(map[string]uiTreeNodesNameToUuidStruct)

	// Initiate map
	availableBuildingBlocksModel.allBuildingBlocksTreeNameToUuid = make(map[string]uiTreeNodesNameToUuidStruct)

	// Load Available Building Blocks, in this case TestInstructionContainers
	availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocks(&testInstructionsAndTestContainersMessage)

	// Load Pinned Building Blocks
	availableBuildingBlocksModel.loadModelWithPinnedBuildingBlocks(&pinnedTestInstructionsAndTestContainersMessage)

	// Validate that a not pinned Building Block can be pinned
	tempTreeName := "Emtpy parallelled processed Turbo TestInstructionsContainer [aa1b973]"
	err := availableBuildingBlocksModel.pinTestInstructionOrTestInstructionContainer(tempTreeName)

	assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_PBB002_ExpectedResultInModel_003, fmt.Sprint(err))

	// Validate that node got pinned
	assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_PBB002_ExpectedResultInModel_004, fmt.Sprint(availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes))

}

// Checks that a non-existing Building Block can't be unpinned
func TestThatNonExistingBuildingBlockCanNotBeUnPinned(t *testing.T) {

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
			"Id": "4757ab26-60ef-4c59-b299-71e32ae74ed7",
		}).Fatalln("Exiting because of wrong Proto-file version in TestData")
	}

	returnMessage = UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion(pinnedTestInstructionsAndTestContainersMessage.AckNackResponse.ProtoFileVersionUsedByClient)
	if returnMessage != nil && returnMessage.AckNack == false {
		myLogger.WithFields(logrus.Fields{
			"Id": "2c22eadb-cf09-495d-8b89-00443d16f4ce",
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
		grpcOut:                                                                    grpc_out_GuiTestCaseBuilderServer.GRPCOutGuiTestCaseBuilderServerStruct{},
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

	// Validate that an non-existing Building Block can't be Unpinned
	err := availableBuildingBlocksModel.unPinTestInstructionOrTestInstructionContainer("NonExistingBuildingBlock")

	assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_PBB003_ExpectedResultInModel_001, fmt.Sprint(err))

}

/*
// Checks that an pinned Building Block exist among pinned building blocks
func TestThatBuildingBlockToBeUnPinnedExistsAmongPinned(t *testing.T) {

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
			"Id": "df1285df-db49-404f-9e8b-a9549aff0f74",
		}).Fatalln("Exiting because of wrong Proto-file version in TestData")
	}

	returnMessage = UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion(pinnedTestInstructionsAndTestContainersMessage.AckNackResponse.ProtoFileVersionUsedByClient)
	if returnMessage != nil && returnMessage.AckNack == false {
		myLogger.WithFields(logrus.Fields{
			"Id": "03d024d1-3fdf-4d2d-84e6-56bb09fcb736",
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
		grpcOut:                                                                    grpc_out_GuiTestCaseBuilderServer.GRPCOutGuiTestCaseBuilderServerStruct{},
		availableBuildingBlockModelSuitedForFyneTreeView:                           nil,
	}

	// Clear and initiate variables
	availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes = make(map[string]uiTreeNodesNameToUuidStruct)

	// Initiate map
	availableBuildingBlocksModel.allBuildingBlocksTreeNameToUuid = make(map[string]uiTreeNodesNameToUuidStruct)

	// Load Available Building Blocks, in this case TestInstructionContainers
	availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocks(&testInstructionsAndTestContainersMessage)

	// Load Pinned Building Blocks
	availableBuildingBlocksModel.loadModelWithPinnedBuildingBlocks(&pinnedTestInstructionsAndTestContainersMessage)

	// Validate that an already pinned  Building Block can't be pinned again
	pinnedBuildingBlockName := availableBuildingBlocksModel.getPinnedBuildingBlocksTreeNamesFromModel()[0]
	pinnedBuildingBlockUuid := availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes[pinnedBuildingBlockName].uuid
	buildingBlockName := availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[pinnedBuildingBlockUuid].nameInUITree
	err := availableBuildingBlocksModel.pinTestInstructionOrTestInstructionContainer(buildingBlockName)

	assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_PBB003_ExpectedResultInModel_002, fmt.Sprint(err))

}
*/
// Checks that a pinned Building Block can be unpinned
func TestUnPinBuildingBlockThatIsPinned(t *testing.T) {

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
			"Id": "df1285df-db49-404f-9e8b-a9549aff0f74",
		}).Fatalln("Exiting because of wrong Proto-file version in TestData")
	}

	returnMessage = UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion(pinnedTestInstructionsAndTestContainersMessage.AckNackResponse.ProtoFileVersionUsedByClient)
	if returnMessage != nil && returnMessage.AckNack == false {
		myLogger.WithFields(logrus.Fields{
			"Id": "03d024d1-3fdf-4d2d-84e6-56bb09fcb736",
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
		grpcOut:                                                                    grpc_out_GuiTestCaseBuilderServer.GRPCOutGuiTestCaseBuilderServerStruct{},
		availableBuildingBlockModelSuitedForFyneTreeView:                           nil,
	}

	// Clear and initiate variables
	availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes = make(map[string]uiTreeNodesNameToUuidStruct)

	// Initiate map
	availableBuildingBlocksModel.allBuildingBlocksTreeNameToUuid = make(map[string]uiTreeNodesNameToUuidStruct)

	// Load Available Building Blocks, in this case TestInstructionContainers
	availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocks(&testInstructionsAndTestContainersMessage)

	// Load Pinned Building Blocks
	availableBuildingBlocksModel.loadModelWithPinnedBuildingBlocks(&pinnedTestInstructionsAndTestContainersMessage)

	// Validate that a not pinned Building Block can be pinned
	tempTreeName := "Just the name (Custody Arrangement) [2f130d7]"
	err := availableBuildingBlocksModel.unPinTestInstructionOrTestInstructionContainer(tempTreeName)

	assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_PBB003_ExpectedResultInModel_003, fmt.Sprint(err))

	// Validate that node got unpinned
	assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_PBB003_ExpectedResultInModel_004, fmt.Sprint(availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes))

}
