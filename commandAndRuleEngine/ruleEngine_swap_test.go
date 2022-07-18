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

// Verify that a 'B0' can be swapped into 'B1-TIC(B10)-B1'
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

	// Create an Immature Element model for 'B1-TIC(B10)-B1'
	immatureElementModel := immatureElementStruct{
		firstElementUuid:   "",
		immatureElementMap: nil,
	}

	// Create first 'B1'
	b1Bond_first := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage{
		OriginalElementUuid:      "a96539e8-79f8-4e9e-bff7-78263f265201",
		OriginalElementName:      "B1-Bond",
		ImmatureElementUuid:      "83cb94a6-be9b-4647-b9ec-d9ae36d0505c",
		PreviousElementUuid:      "83cb94a6-be9b-4647-b9ec-d9ae36d0505c",
		NextElementUuid:          "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		FirstChildElementUuid:    "83cb94a6-be9b-4647-b9ec-d9ae36d0505c",
		ParentElementUuid:        "83cb94a6-be9b-4647-b9ec-d9ae36d0505c",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1_BOND_NONE_SWAPPABLE,
	}

	// Create B10 in TIC(x)
	b10Bond := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage{
		OriginalElementUuid:      "7be82e83-6048-4c30-b4aa-b68c11037c1d",
		OriginalElementName:      "B10-Bond",
		ImmatureElementUuid:      "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		PreviousElementUuid:      "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		NextElementUuid:          "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		FirstChildElementUuid:    "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		ParentElementUuid:        "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
	}

	// Create TIC
	tIC := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		ImmatureElementUuid:      "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		PreviousElementUuid:      "83cb94a6-be9b-4647-b9ec-d9ae36d0505c",
		NextElementUuid:          "d7141e57-7fc8-4e21-84ab-3ef62df2ee90",
		FirstChildElementUuid:    "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		ParentElementUuid:        "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	// Create last 'B1'
	b1Bond_last := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage{
		OriginalElementUuid:      "a96539e8-79f8-4e9e-bff7-78263f265201",,
		OriginalElementName:      "B1-Bond",
		ImmatureElementUuid:      "d7141e57-7fc8-4e21-84ab-3ef62df2ee90",
		PreviousElementUuid:      "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		NextElementUuid:          "d7141e57-7fc8-4e21-84ab-3ef62df2ee90",
		FirstChildElementUuid:    "d7141e57-7fc8-4e21-84ab-3ef62df2ee90",
		ParentElementUuid:        "d7141e57-7fc8-4e21-84ab-3ef62df2ee90",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1_BOND_NONE_SWAPPABLE,
	}

	// Add the Elements to the Immature Elements Model map
	immatureElementModel.immatureElementMap["83cb94a6-be9b-4647-b9ec-d9ae36d0505c"] = b1Bond_first
	immatureElementModel.immatureElementMap["ff224d27-5c8a-48b9-ace9-af43245cb35d"] = b10Bond
	immatureElementModel.immatureElementMap["d444b8d8-b2fb-4505-ad8e-36bfe89988ab"] = tIC
	immatureElementModel.immatureElementMap["d7141e57-7fc8-4e21-84ab-3ef62df2ee90"] = b1Bond_last

	// Add first Element ti Immature Element Model
	immatureElementModel.firstElementUuid = "83cb94a6-be9b-4647-b9ec-d9ae36d0505c"

	// Very the swap
	xxx



	// Validate that an non-existing Building Block can't be pinned
	err := availableBuildingBlocksModel.pinTestInstructionOrTestInstructionContainer("NonExistingBuildingBlock")

	assert.Equal(t, UnitTestTestData.TestInstructionsAndTestInstructionsContainersRespons_PBB002_ExpectedResultInModel_001, fmt.Sprint(err))

}
