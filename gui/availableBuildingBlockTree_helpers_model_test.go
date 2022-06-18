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

/*
func testlist() *notelist {
	a := test.NewApp()
	n := &notelist{pref: a.Preferences()}

	return n
}

*/


//var availableBuildingBlocksModel *availableBuildingBlocksModelStruct

// Checks that Available TestInstructions are put in Available Building Blocks-model in a correct way
func TestLoadModelWithAvailableBuildingBlocksRegardingTestInstructions(t *testing.T) {

	var testInstructionsAndTestContainersMessage fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionsAndTestContainersMessage
	//testInstructionsAndTestContainersMessage = &fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionsAndTestContainersMessage

	//res := fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionsAndTestContainersMessage{}
	//json.Unmarshal(byt, &res)
	mystring := UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_ABB001

	if err := jsonpb.UnmarshalString(mystring, &testInstructionsAndTestContainersMessage); err != nil {
		panic(err)
	}

	myLogger := UnitTestTestData.InitLoggerForTest("")

	returnMessage := UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion(testInstructionsAndTestContainersMessage.)
	if returnMessage.AckNack == false {
		myLogger.WithFields(logrus.Fields{
			"Id": "c9453631-46b0-47a0-86fa-fe2e5b51ed92",
		}).Fatalln("Exiting because of wrong Proto-file version in TestData")
	}
	myLogger.WithFields(logrus.Fields{
		"Id": "c9453631-46b0-47a0-86fa-fe2e5b51ed92",
	}).Info("Clean up and shut down servers")

	var availableBuildingBlocksModel *availableBuildingBlocksModelStruct

	availableBuildingBlocksModel = &availableBuildingBlocksModelStruct{
		logger:                             myLogger,
		fenixGuiBuilderServerAddressToDial: "",
		fullDomainTestInstructionTypeTestInstructionRelationsMap:                   nil,
		fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap: nil,
		availableBuildingBlocksForUITreeNodes:                                      nil,
		grpcOut:                                                                    grpc_out.GRPCOutStruct{},
		availableBuildingBlockModelSuitedForFyneTreeView:                           nil,
	}

	availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocksRegardingTestInstructions(&testInstructionsAndTestContainersMessage)

	// fmt.Println(availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap)
	/*
		fmt.Println(availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes)
		b, err := json.Marshal(&availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes)

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b))

	*/

	// Validate 'availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes'
	assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_001, fmt.Sprint(availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes))

	// Validate 'availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap'
	assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_002, fmt.Sprint(availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap))

}

// Checks that Available TestInstructionContainers are put in Available Building Blocks-model in a correct way
func TestLoadModelWithAvailableBuildingBlocksRegardingTestInstructionContainerss(t *testing.T) {

	var testInstructionsAndTestContainersMessage fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionsAndTestContainersMessage

	// Load data into gRPC-meTestLoadModelWithAvailableBuildingBlocksRegardingTestInstructionContainerssssage response (from DB)
	if err := jsonpb.UnmarshalString(UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_ABB001, &testInstructionsAndTestContainersMessage); err != nil {
		panic(err)
	}

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate availableBuildingBlocksModel
	var availableBuildingBlocksModel *availableBuildingBlocksModelStruct

	availableBuildingBlocksModel = &availableBuildingBlocksModelStruct{
		logger:                             myLogger,
		fenixGuiBuilderServerAddressToDial: "",
		fullDomainTestInstructionTypeTestInstructionRelationsMap:                   nil,
		fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap: nil,
		availableBuildingBlocksForUITreeNodes:                                      nil,
		grpcOut:                                                                    grpc_out.GRPCOutStruct{},
		availableBuildingBlockModelSuitedForFyneTreeView:                           nil,
	}

	// Load Available Building Blocks, in this case TestInstructionContainers
	availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocksRegardingTestInstructionContainers(&testInstructionsAndTestContainersMessage)

	//fmt.Println(availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap)
	/*
		fmt.Println(availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes)
		b, err := json.Marshal(&availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes)

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b))

	*/

	//fmt.Println(UnitTestTestData.TestInstructionsAndTestInstructionsRespons_ABB001_ExpectedResultInModel_003)
	//fmt.Println(UnitTestTestData.TestInstructionsAndTestInstructionsRespons_ABB001_ExpectedResultInModel_004)

	// Validate 'availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes'
	assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_003, fmt.Sprint(availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes))

	// Validate 'availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap'
	assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_004, fmt.Sprint(availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap))

}

// Checks that Available TestInstructions & TestInstructionContainers are put in Available Building Blocks-model in a correct way
func TestLoadModelWithAvailableBuildingBlocks(t *testing.T) {

	var testInstructionsAndTestContainersMessage fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionsAndTestContainersMessage

	// Load data into gRPC-message response (from DB)
	if err := jsonpb.UnmarshalString(UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_ABB001, &testInstructionsAndTestContainersMessage); err != nil {
		panic(err)
	}

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate availableBuildingBlocksModel
	var availableBuildingBlocksModel *availableBuildingBlocksModelStruct

	availableBuildingBlocksModel = &availableBuildingBlocksModelStruct{
		logger:                             myLogger,
		fenixGuiBuilderServerAddressToDial: "",
		fullDomainTestInstructionTypeTestInstructionRelationsMap:                   nil,
		fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap: nil,
		availableBuildingBlocksForUITreeNodes:                                      nil,
		grpcOut:                                                                    grpc_out.GRPCOutStruct{},
		availableBuildingBlockModelSuitedForFyneTreeView:                           nil,
	}

	// Load Available Building Blocks, in this case TestInstructionContainers
	availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocks(&testInstructionsAndTestContainersMessage)

	//fmt.Println(UnitTestTestData.TestInstructionsAndTestInstructionsRespons_ABB001_ExpectedResultInModel_003)
	//fmt.Println(UnitTestTestData.TestInstructionsAndTestInstructionsRespons_ABB001_ExpectedResultInModel_004)

	// Validate 'availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes'
	assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_005, fmt.Sprint(availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes))

	// Validate 'availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap'
	assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_006, fmt.Sprint(availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap))

	// Validate 'availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap'
	assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_007, fmt.Sprint(availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap))

}
