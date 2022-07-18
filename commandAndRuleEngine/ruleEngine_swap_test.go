package commandAndRuleEngine

import (
	"FenixTesterGui/testCase/testCaseModel"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"testing"

	"FenixTesterGui/gui/UnitTestTestData"
	"github.com/stretchr/testify/assert"
)

// Verify that a 'B0' can be swapped into 'B11f-TIC(B10)-B11l'
// TCRuleSwap101
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC(X)			B0					n 		B0								B1-n-B1					TCRuleSwap101
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

	// Initiate a TestCaseModel
	myTestCaseModel := testCaseModel.TestCaseModelStruct{
		LastLoadedTestCaseModelGRPCMessage: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage{},
		FirstElementUuid:                   "",
		TestCaseModelMap:                   nil,
	}

	testCaseModelMap := make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage)

	b0Bond := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "0a1c1266-f8ad-484a-a76c-59c5fd7fbda5",
		OriginalElementName:      "B0_BOND",
		MatureElementUuid:        "4b694f8c-f194-45af-a75e-f2bb3fd350e6",
		PreviousElementUuid:      "4b694f8c-f194-45af-a75e-f2bb3fd350e6",
		NextElementUuid:          "4b694f8c-f194-45af-a75e-f2bb3fd350e6",
		FirstChildElementUuid:    "4b694f8c-f194-45af-a75e-f2bb3fd350e6",
		ParentElementUuid:        "4b694f8c-f194-45af-a75e-f2bb3fd350e6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND,
	}

	// Add B0-bond to the TestCaseModel-map
	testCaseModelMap[b0Bond.MatureElementUuid] = b0Bond

	// Set the B0-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = b0Bond.MatureElementUuid

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := commandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		testcaseModel:     &myTestCaseModel,
	}

	// Create an Immature Element model for 'B11f-TIC(B10)-B11l'

	// Validate that an non-existing Building Block can't be pinned
	err := availableBuildingBlocksModel.pinTestInstructionOrTestInstructionContainer("NonExistingBuildingBlock")

	assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_PBB002_ExpectedResultInModel_001, fmt.Sprint(err))

}
