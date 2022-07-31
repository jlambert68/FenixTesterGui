package commandAndRuleEngine

import (
	"FenixTesterGui/gui/UnitTestTestData"
	"FenixTesterGui/testCase/testCaseModel"
	"fmt"
	uuidGenerator "github.com/google/uuid"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

// Test to create a New TestCaseModel
func TestNewTestCaseModelCommand(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate structure for all TestCases
	allTestCases := make(map[string]testCaseModel.TestCaseModelStruct)
	testCasesObject := testCaseModel.TestCaseModelsStruct{TestCases: allTestCases}

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := commandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		testcases:         &testCasesObject,
	}

	// Add needed data for availableBondsMap
	tempAvailableBondsMap := make(map[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage)

	// B0_BOND
	visibleBondAttributesMessage_AvaialbeBond_B0_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage_VisibleBondAttributesMessage{
		BondUuid: "2858d47a-198c-43f3-abe8-abd2a36f6045",
		BondName: "B0_BOND",
	}

	basicBondInformationMessage_AvaialbeBond_B0_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage{
		VisibleBondAttributes: &visibleBondAttributesMessage_AvaialbeBond_B0_BOND}

	immatureBondsMessage_ImmatureBondMessage_B0_BOND := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage{
		BasicBondInformation: &basicBondInformationMessage_AvaialbeBond_B0_BOND}

	tempAvailableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND] = &immatureBondsMessage_ImmatureBondMessage_B0_BOND

	// Add bond-map to 'commandAndRuleEngine.availableBondsMap'
	commandAndRuleEngine.availableBondsMap = tempAvailableBondsMap

	// Execute command
	testCaseUuid, err := commandAndRuleEngine.executeCommandOnTestCaseModel_NewTestCaseModel()

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Extract TestCase
	testCase := commandAndRuleEngine.testcases.TestCases[testCaseUuid]

	// Validate the result of the NewTestCaseModel-command, 'B0'
	// 1) Validate B0 (1)
	testCaseModelElementUuid_1 := testCase.FirstElementUuid
	testCaseModelElement_1 := testCase.TestCaseModelMap[testCaseModelElementUuid_1]

	correctElement := testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.ParentElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.PreviousElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.FirstChildElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.NextElementUuid &&
		testCaseModelElement_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = commandAndRuleEngine.testcases.VerifyThatThereAreNoZombieElementsInTestCaseModel(testCaseUuid)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate Textual TestCase Presentation
	textualTestCaseSimple, textualTestCaseComplex, err := commandAndRuleEngine.testcases.CreateTextualTestCase(testCaseUuid)

	textualTestCaseRepresentationSimple := "[B0]"
	textualTestCaseRepresentationComplex := "[B0]"

	assert.Equal(t, textualTestCaseRepresentationSimple, textualTestCaseSimple)
	assert.Equal(t, textualTestCaseRepresentationComplex, textualTestCaseComplex)

	// Validate Command stack, but fix timestamp
	commandTimeStamp := testCase.CommandStack[0].CommandExecutedTimeStamp
	commandTimeStampSecondsAsString := strconv.Itoa(int(commandTimeStamp.Seconds))
	commandTimeStampnanosAsString := strconv.Itoa(int(commandTimeStamp.Nanos))

	commandSliceToCompareWith := "[{{{} [] [] <nil>} 0 [] NEW_TESTCASE NEW_TESTCASE N/A N/A  seconds:" + commandTimeStampSecondsAsString + " nanos:" + commandTimeStampnanosAsString + "}]"

	assert.Equal(t, commandSliceToCompareWith, fmt.Sprint(testCase.CommandStack))

}

// Test to Delete an element from the TestCaseModel
func TestRemoveElementCommandOnTestCaseModel(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := commandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		testcases:         nil,
	}

	// Add needed data for availableBondsMap
	tempAvailableBondsMap := make(map[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage)

	// B0_BOND
	visibleBondAttributesMessage_AvaialbeBond_B0_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage_VisibleBondAttributesMessage{
		BondUuid: "2858d47a-198c-43f3-abe8-abd2a36f6045",
		BondName: "B0_BOND",
	}

	basicBondInformationMessage_AvaialbeBond_B0_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage{
		VisibleBondAttributes: &visibleBondAttributesMessage_AvaialbeBond_B0_BOND}

	immatureBondsMessage_ImmatureBondMessage_B0_BOND := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage{
		BasicBondInformation: &basicBondInformationMessage_AvaialbeBond_B0_BOND}

	tempAvailableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND] = &immatureBondsMessage_ImmatureBondMessage_B0_BOND

	// Add bond-map to 'commandAndRuleEngine.availableBondsMap'
	commandAndRuleEngine.availableBondsMap = tempAvailableBondsMap

	// Initiate a TestCaseModel
	myTestCaseModel := testCaseModel.TestCaseModelStruct{
		LastLoadedTestCaseModelGRPCMessage: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage{},
		FirstElementUuid:                   "",
		TestCaseModelMap:                   nil,
	}

	myTestCaseModel.TestCaseModelMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage)

	uuidToBeDeleted := "4b694f8c-f194-45af-a75e-f2bb3fd350e6"

	tc_b1f := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "cff02b54-ee53-47d1-94d6-48dc470073a3",
		OriginalElementName:      "B1l_BOND",
		MatureElementUuid:        "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		PreviousElementUuid:      "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		NextElementUuid:          uuidToBeDeleted,
		FirstChildElementUuid:    "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		ParentElementUuid:        "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE,
	}

	tc_tic := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		MatureElementUuid:        uuidToBeDeleted,
		PreviousElementUuid:      "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		NextElementUuid:          "79a6702d-8370-446c-b001-d60494eca6fa",
		FirstChildElementUuid:    "b1d535a7-e0b4-4a67-9581-f7d157f7ba1e",
		ParentElementUuid:        uuidToBeDeleted,
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_b1l := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c10bfd8f-2786-480c-9a71-bad9ec7bc582",
		OriginalElementName:      "B1l_BOND",
		MatureElementUuid:        "79a6702d-8370-446c-b001-d60494eca6fa",
		PreviousElementUuid:      uuidToBeDeleted,
		NextElementUuid:          "79a6702d-8370-446c-b001-d60494eca6fa",
		FirstChildElementUuid:    "79a6702d-8370-446c-b001-d60494eca6fa",
		ParentElementUuid:        "79a6702d-8370-446c-b001-d60494eca6fa",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE,
	}

	tc_b10 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "4f348a88-bb9a-4852-80de-3141687db11d",
		OriginalElementName:      "B10_Bond",
		MatureElementUuid:        "b1d535a7-e0b4-4a67-9581-f7d157f7ba1e",
		PreviousElementUuid:      "b1d535a7-e0b4-4a67-9581-f7d157f7ba1e",
		NextElementUuid:          "b1d535a7-e0b4-4a67-9581-f7d157f7ba1e",
		FirstChildElementUuid:    "b1d535a7-e0b4-4a67-9581-f7d157f7ba1e",
		ParentElementUuid:        uuidToBeDeleted,
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
	}

	// Set up TestCaseModel-map
	myTestCaseModel.TestCaseModelMap[tc_b1f.MatureElementUuid] = tc_b1f
	myTestCaseModel.TestCaseModelMap[tc_tic.MatureElementUuid] = tc_tic
	myTestCaseModel.TestCaseModelMap[tc_b1l.MatureElementUuid] = tc_b1l
	myTestCaseModel.TestCaseModelMap[tc_b10.MatureElementUuid] = tc_b10

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_b1f.MatureElementUuid

	// Initiate structure for all TestCases
	allTestCases := make(map[string]testCaseModel.TestCaseModelStruct)

	// Generate new UUID
	testCaseUuid := uuidGenerator.New().String()

	// Add myTestCaseModel to map of all testcases
	allTestCases[testCaseUuid] = myTestCaseModel

	// Set Current User
	currentUser := "s41797"

	// Add AddTestCases to TestCases-model
	testCasesObject := testCaseModel.TestCaseModelsStruct{
		TestCases:   allTestCases,
		CurrentUser: currentUser}

	// Add reference to TestCAses in command and rule engine
	commandAndRuleEngine.testcases = &testCasesObject

	// Execute command
	err := commandAndRuleEngine.executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel(testCaseUuid, uuidToBeDeleted)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Extract TestCase
	testCase := commandAndRuleEngine.testcases.TestCases[testCaseUuid]

	// Validate the result of the NewTestCaseModel-command, 'B0'
	// 1) Validate B0 (1)
	testCaseModelElementUuid_1 := testCase.FirstElementUuid
	testCaseModelElement_1 := testCase.TestCaseModelMap[testCaseModelElementUuid_1]

	correctElement := testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.ParentElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.PreviousElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.FirstChildElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.NextElementUuid &&
		testCaseModelElement_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = commandAndRuleEngine.testcases.VerifyThatThereAreNoZombieElementsInTestCaseModel(testCaseUuid)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate Textual TestCase Presentation
	textualTestCaseSimple, textualTestCaseComplex, err := commandAndRuleEngine.testcases.CreateTextualTestCase(testCaseUuid)

	textualTestCaseRepresentationSimple := "[B0]"
	textualTestCaseRepresentationComplex := "[B0]"

	assert.Equal(t, textualTestCaseRepresentationSimple, textualTestCaseSimple)
	assert.Equal(t, textualTestCaseRepresentationComplex, textualTestCaseComplex)

	// Validate Simple- And Complex- Textual Representation Stack - length
	lenghtIsOneSimple := fmt.Sprint(len(testCase.TextualTestCaseRepresentationSimpleStack) == 1)
	assert.Equal(t, "true", lenghtIsOneSimple)
	lenghtIsOneComplex := fmt.Sprint(len(testCase.TextualTestCaseRepresentationComplexStack) == 1)
	assert.Equal(t, "true", lenghtIsOneComplex)

	// Validate Simple Textual Representation Stack - Content
	textualRepresentationStackContentSimple := testCase.TextualTestCaseRepresentationSimpleStack[0]
	assert.Equal(t, textualTestCaseRepresentationSimple, textualRepresentationStackContentSimple)

	// Validate Complex Textual Representation Stack - Content
	textualRepresentationStackContentSComplex := testCase.TextualTestCaseRepresentationComplexStack[0]
	assert.Equal(t, textualTestCaseRepresentationComplex, textualRepresentationStackContentSComplex)

	// Validate Command stack length
	lenghtIsOne := fmt.Sprint(len(testCase.CommandStack) == 1)
	assert.Equal(t, "true", lenghtIsOne)

	// Validate Command Stack content
	assert.Equal(t, fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_REMOVE_ELEMENT, testCase.CommandStack[0].TestCaseCommandType)
	assert.Equal(t, fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_name[int32(fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_REMOVE_ELEMENT)], testCase.CommandStack[0].TestCaseCommandName)
	assert.Equal(t, uuidToBeDeleted, testCase.CommandStack[0].FirstParameter)
	assert.Equal(t, testCaseModel.NotApplicable, testCase.CommandStack[0].SecondParameter)
	assert.Equal(t, currentUser, testCase.CommandStack[0].UserId)

}

// Test to Swap out an element and in another element in the TestCaseModel
func TestSwapElementCommandOnTestCaseModel(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := commandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		testcases:         nil,
	}

	// Add needed data for availableBondsMap
	tempAvailableBondsMap := make(map[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage)

	// B1f_BOND
	visibleBondAttributesMessage_AvaialbeBond_B1f_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage_VisibleBondAttributesMessage{
		BondUuid: "2858d47a-198c-43f3-abe8-abd2a36f6045",
		BondName: "B1f_BOND",
	}

	basicBondInformationMessage_AvaialbeBond_B1f_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage{
		VisibleBondAttributes: &visibleBondAttributesMessage_AvaialbeBond_B1f_BOND}

	immatureBondsMessage_ImmatureBondMessage_B1f_BOND := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage{
		BasicBondInformation: &basicBondInformationMessage_AvaialbeBond_B1f_BOND}

	tempAvailableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE] = &immatureBondsMessage_ImmatureBondMessage_B1f_BOND

	// B1l_BOND
	visibleBondAttributesMessage_AvaialbeBond_B1l_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage_VisibleBondAttributesMessage{
		BondUuid: "95cbb203-1bb3-4ab4-84b7-c2a27a2e40fb",
		BondName: "B1l_BOND",
	}

	basicBondInformationMessage_AvaialbeBond_B1l_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage{
		VisibleBondAttributes: &visibleBondAttributesMessage_AvaialbeBond_B1l_BOND}

	immatureBondsMessage_ImmatureBondMessage_B1l_BOND := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage{
		BasicBondInformation: &basicBondInformationMessage_AvaialbeBond_B1l_BOND}

	tempAvailableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE] = &immatureBondsMessage_ImmatureBondMessage_B1l_BOND

	// Add bond-map to 'commandAndRuleEngine.availableBondsMap'
	commandAndRuleEngine.availableBondsMap = tempAvailableBondsMap

	// Initiate a TestCaseModel
	myTestCaseModel := testCaseModel.TestCaseModelStruct{
		LastLoadedTestCaseModelGRPCMessage: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage{},
		FirstElementUuid:                   "",
		TestCaseModelMap:                   nil,
	}

	myTestCaseModel.TestCaseModelMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage)

	uuidToBeDeleted := "4b694f8c-f194-45af-a75e-f2bb3fd350e6"

	tc_b0 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "B0_BOND",
		MatureElementUuid:        uuidToBeDeleted,
		PreviousElementUuid:      uuidToBeDeleted,
		NextElementUuid:          uuidToBeDeleted,
		FirstChildElementUuid:    uuidToBeDeleted,
		ParentElementUuid:        uuidToBeDeleted,
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND,
	}

	// Set up TestCaseModel-map
	myTestCaseModel.TestCaseModelMap[tc_b0.MatureElementUuid] = tc_b0

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_b0.MatureElementUuid

	// Initiate structure for all TestCases
	allTestCases := make(map[string]testCaseModel.TestCaseModelStruct)

	// Generate new UUID
	testCaseUuid := uuidGenerator.New().String()

	// Add myTestCaseModel to map of all testcases
	allTestCases[testCaseUuid] = myTestCaseModel

	// Set Current User
	currentUser := "s41797"

	// Add AddTestCases to TestCases-model
	testCasesObject := testCaseModel.TestCaseModelsStruct{
		TestCases:   allTestCases,
		CurrentUser: currentUser}

	// Add reference to TestCAses in command and rule engine
	commandAndRuleEngine.testcases = &testCasesObject

	// Create an Immature Element model for 'TIC(B10)'
	immatureElementModel := testCaseModel.ImmatureElementStruct{
		FirstElementUuid:   "",
		ImmatureElementMap: nil,
	}

	immatureElementModel.ImmatureElementMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage)

	// Create TIC
	tic := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		ImmatureElementUuid:      "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		PreviousElementUuid:      "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		NextElementUuid:          "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		FirstChildElementUuid:    "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		ParentElementUuid:        "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
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

	// Add the Elements to the Immature Elements Model map
	immatureElementModel.ImmatureElementMap["d444b8d8-b2fb-4505-ad8e-36bfe89988ab"] = tic
	immatureElementModel.ImmatureElementMap["ff224d27-5c8a-48b9-ace9-af43245cb35d"] = b10Bond

	// Add first Element ti Immature Element Model
	immatureElementModel.FirstElementUuid = "d444b8d8-b2fb-4505-ad8e-36bfe89988ab"

	// Execute command
	err := commandAndRuleEngine.executeCommandOnTestCaseModel_SwapOutElemenAndInNewElementInTestCaseModel(testCaseUuid, uuidToBeDeleted, &immatureElementModel)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Extract TestCase
	testCase := commandAndRuleEngine.testcases.TestCases[testCaseUuid]

	// Validate the result of the Swap-Element-command, 'B1f-TIC(B10)-B1l'
	// 1) Validate B1f (1)
	testCaseModelElementUuid_1 := testCase.FirstElementUuid
	testCaseModelElement_1 := testCase.TestCaseModelMap[testCaseModelElementUuid_1]

	correctElement := testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.ParentElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.PreviousElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.FirstChildElementUuid &&
		testCaseModelElement_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 2) Validate TIC (2)
	testCaseModelElementUuid_2 := testCaseModelElement_1.NextElementUuid
	testCaseModelElement_2 := testCase.TestCaseModelMap[testCaseModelElementUuid_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2.ParentElementUuid &&
		testCaseModelElement_2.MatureElementUuid == testCaseModelElement_1.NextElementUuid &&
		testCaseModelElement_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 3) Validate B1l (3)
	testCaseModelElementUuid_3 := testCaseModelElement_2.NextElementUuid
	testCaseModelElement_3 := testCase.TestCaseModelMap[testCaseModelElementUuid_3]

	correctElement = testCaseModelElement_3.MatureElementUuid == testCaseModelElement_3.ParentElementUuid &&
		testCaseModelElement_3.MatureElementUuid == testCaseModelElement_2.NextElementUuid &&
		testCaseModelElement_3.MatureElementUuid == testCaseModelElement_3.NextElementUuid &&
		testCaseModelElement_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 4) Validate B10 (2.1)
	testCaseModelElementUuid_2_1 := testCaseModelElement_2.FirstChildElementUuid
	testCaseModelElement_2_1 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_1]

	correctElement = testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2.FirstChildElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.NextElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = commandAndRuleEngine.testcases.VerifyThatThereAreNoZombieElementsInTestCaseModel(testCaseUuid)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate Textual TestCase Presentation
	textualTestCaseSimple, textualTestCaseComplex, err := commandAndRuleEngine.testcases.CreateTextualTestCase(testCaseUuid)

	textualTestCaseRepresentationSimple := "[B1-TIC(B10)-B1]"
	textualTestCaseRepresentationComplex := "[B1f-TIC(B10)-B1l]"

	assert.Equal(t, textualTestCaseRepresentationSimple, textualTestCaseSimple)
	assert.Equal(t, textualTestCaseRepresentationComplex, textualTestCaseComplex)

	// Validate Simple- And Complex- Textual Representation Stack - length
	lenghtIsOneSimple := fmt.Sprint(len(testCase.TextualTestCaseRepresentationSimpleStack) == 1)
	assert.Equal(t, "true", lenghtIsOneSimple)
	lenghtIsOneComplex := fmt.Sprint(len(testCase.TextualTestCaseRepresentationComplexStack) == 1)
	assert.Equal(t, "true", lenghtIsOneComplex)

	// Validate Simple Textual Representation Stack - Content
	textualRepresentationStackContentSimple := testCase.TextualTestCaseRepresentationSimpleStack[0]
	assert.Equal(t, textualTestCaseRepresentationSimple, textualRepresentationStackContentSimple)

	// Validate Complex Textual Representation Stack - Content
	textualRepresentationStackContentSComplex := testCase.TextualTestCaseRepresentationComplexStack[0]
	assert.Equal(t, textualTestCaseRepresentationComplex, textualRepresentationStackContentSComplex)

	// Validate Command stack length
	lenghtIsOne := fmt.Sprint(len(testCase.CommandStack) == 1)
	assert.Equal(t, "true", lenghtIsOne)

	// Validate Command Stack content
	assert.Equal(t, fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_NEW_ELEMENT, testCase.CommandStack[0].TestCaseCommandType)
	assert.Equal(t, fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_name[int32(fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_NEW_ELEMENT)], testCase.CommandStack[0].TestCaseCommandName)
	assert.Equal(t, uuidToBeDeleted, testCase.CommandStack[0].FirstParameter)
	assert.Equal(t, immatureElementModel.FirstElementUuid, testCase.CommandStack[0].SecondParameter)
	assert.Equal(t, currentUser, testCase.CommandStack[0].UserId)

}

// Test to Copy an element from the TestCaseModel
func TestCopyElementCommandOnTestCaseModel(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := commandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		testcases:         nil,
	}

	// Initiate a TestCaseModel
	myTestCaseModel := testCaseModel.TestCaseModelStruct{
		LastLoadedTestCaseModelGRPCMessage: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage{},
		FirstElementUuid:                   "",
		TestCaseModelMap:                   nil,
	}

	myTestCaseModel.TestCaseModelMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage)

	uuidToBeCopied := "4b694f8c-f194-45af-a75e-f2bb3fd350e6"

	tc_b1f := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "cff02b54-ee53-47d1-94d6-48dc470073a3",
		OriginalElementName:      "B1l_BOND",
		MatureElementUuid:        "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		PreviousElementUuid:      "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		NextElementUuid:          uuidToBeCopied,
		FirstChildElementUuid:    "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		ParentElementUuid:        "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE,
	}

	tc_tic := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		MatureElementUuid:        uuidToBeCopied,
		PreviousElementUuid:      "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		NextElementUuid:          "79a6702d-8370-446c-b001-d60494eca6fa",
		FirstChildElementUuid:    "b1d535a7-e0b4-4a67-9581-f7d157f7ba1e",
		ParentElementUuid:        uuidToBeCopied,
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_b1l := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c10bfd8f-2786-480c-9a71-bad9ec7bc582",
		OriginalElementName:      "B1l_BOND",
		MatureElementUuid:        "79a6702d-8370-446c-b001-d60494eca6fa",
		PreviousElementUuid:      uuidToBeCopied,
		NextElementUuid:          "79a6702d-8370-446c-b001-d60494eca6fa",
		FirstChildElementUuid:    "79a6702d-8370-446c-b001-d60494eca6fa",
		ParentElementUuid:        "79a6702d-8370-446c-b001-d60494eca6fa",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE,
	}

	tc_b10 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "4f348a88-bb9a-4852-80de-3141687db11d",
		OriginalElementName:      "B10_Bond",
		MatureElementUuid:        "b1d535a7-e0b4-4a67-9581-f7d157f7ba1e",
		PreviousElementUuid:      "b1d535a7-e0b4-4a67-9581-f7d157f7ba1e",
		NextElementUuid:          "b1d535a7-e0b4-4a67-9581-f7d157f7ba1e",
		FirstChildElementUuid:    "b1d535a7-e0b4-4a67-9581-f7d157f7ba1e",
		ParentElementUuid:        uuidToBeCopied,
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
	}

	// Set up TestCaseModel-map
	myTestCaseModel.TestCaseModelMap[tc_b1f.MatureElementUuid] = tc_b1f
	myTestCaseModel.TestCaseModelMap[tc_tic.MatureElementUuid] = tc_tic
	myTestCaseModel.TestCaseModelMap[tc_b1l.MatureElementUuid] = tc_b1l
	myTestCaseModel.TestCaseModelMap[tc_b10.MatureElementUuid] = tc_b10

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_b1f.MatureElementUuid

	// Initiate structure for all TestCases
	allTestCases := make(map[string]testCaseModel.TestCaseModelStruct)

	// Generate new UUID
	testCaseUuid := uuidGenerator.New().String()

	// Add myTestCaseModel to map of all testcases
	allTestCases[testCaseUuid] = myTestCaseModel

	// Set Current User
	currentUser := "s41797"

	// Add AddTestCases to TestCases-model
	testCasesObject := testCaseModel.TestCaseModelsStruct{
		TestCases:   allTestCases,
		CurrentUser: currentUser}

	// Add reference to TestCAses in command and rule engine
	commandAndRuleEngine.testcases = &testCasesObject

	// Execute command
	err := commandAndRuleEngine.executeCommandOnTestCaseModel_CopyElementInTestCaseModel(testCaseUuid, uuidToBeCopied)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Extract TestCase
	testCase := commandAndRuleEngine.testcases.TestCases[testCaseUuid]

	// Validate the result of the Copy-Element-command, 'B1f-TIC(B10)-B1l'
	// 1) Validate B1f (1)
	testCaseModelElementUuid_1 := testCase.FirstElementUuid
	testCaseModelElement_1 := testCase.TestCaseModelMap[testCaseModelElementUuid_1]

	correctElement := testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.ParentElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.PreviousElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.FirstChildElementUuid &&
		testCaseModelElement_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 2) Validate TIC (2)
	testCaseModelElementUuid_2 := testCaseModelElement_1.NextElementUuid
	testCaseModelElement_2 := testCase.TestCaseModelMap[testCaseModelElementUuid_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2.ParentElementUuid &&
		testCaseModelElement_2.MatureElementUuid == testCaseModelElement_1.NextElementUuid &&
		testCaseModelElement_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 3) Validate B1l (3)
	testCaseModelElementUuid_3 := testCaseModelElement_2.NextElementUuid
	testCaseModelElement_3 := testCase.TestCaseModelMap[testCaseModelElementUuid_3]

	correctElement = testCaseModelElement_3.MatureElementUuid == testCaseModelElement_3.ParentElementUuid &&
		testCaseModelElement_3.MatureElementUuid == testCaseModelElement_2.NextElementUuid &&
		testCaseModelElement_3.MatureElementUuid == testCaseModelElement_3.NextElementUuid &&
		testCaseModelElement_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 4) Validate B10 (2.1)
	testCaseModelElementUuid_2_1 := testCaseModelElement_2.FirstChildElementUuid
	testCaseModelElement_2_1 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_1]

	correctElement = testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2.FirstChildElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.NextElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = commandAndRuleEngine.testcases.VerifyThatThereAreNoZombieElementsInTestCaseModel(testCaseUuid)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate Textual TestCase Presentation
	textualTestCaseSimple, textualTestCaseComplex, err := commandAndRuleEngine.testcases.CreateTextualTestCase(testCaseUuid)

	textualTestCaseRepresentationSimple := "[B1-TIC(B10)-B1]"
	textualTestCaseRepresentationComplex := "[B1f-TIC(B10)-B1l]"

	assert.Equal(t, textualTestCaseRepresentationSimple, textualTestCaseSimple)
	assert.Equal(t, textualTestCaseRepresentationComplex, textualTestCaseComplex)

	// Validate Command stack length
	lenghtIsOne := fmt.Sprint(len(testCase.CommandStack) == 1)
	assert.Equal(t, "true", lenghtIsOne)

	// Validate Command Stack content
	assert.Equal(t, fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_COPY_ELEMENT, testCase.CommandStack[0].TestCaseCommandType)
	assert.Equal(t, fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_name[int32(fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_COPY_ELEMENT)], testCase.CommandStack[0].TestCaseCommandName)
	assert.Equal(t, uuidToBeCopied, testCase.CommandStack[0].FirstParameter)
	assert.Equal(t, testCaseModel.NotApplicable, testCase.CommandStack[0].SecondParameter)
	assert.Equal(t, currentUser, testCase.CommandStack[0].UserId)

	// Validate Copy Buffer content
	// 1) Validate TIC(1)
	copyBufferElementUuid_1 := testCase.CopyBuffer.FirstElementUuid
	copyBufferElementElement_1 := testCase.CopyBuffer.ImmatureElementMap[copyBufferElementUuid_1]

	correctElement = copyBufferElementElement_1.ImmatureElementUuid == copyBufferElementElement_1.ParentElementUuid &&
		copyBufferElementElement_1.ImmatureElementUuid == copyBufferElementElement_1.PreviousElementUuid &&
		copyBufferElementElement_1.ImmatureElementUuid == copyBufferElementElement_1.NextElementUuid &&
		copyBufferElementElement_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 2) Validate TIC(1.1)
	copyBufferElementUuid_1_1 := copyBufferElementElement_1.FirstChildElementUuid
	copyBufferElementElement_1_1 := testCase.CopyBuffer.ImmatureElementMap[copyBufferElementUuid_1_1]

	correctElement = copyBufferElementElement_1_1.ImmatureElementUuid == copyBufferElementElement_1.FirstChildElementUuid &&
		copyBufferElementElement_1_1.ImmatureElementUuid == copyBufferElementElement_1_1.PreviousElementUuid &&
		copyBufferElementElement_1_1.ImmatureElementUuid == copyBufferElementElement_1_1.NextElementUuid &&
		copyBufferElementElement_1_1.ImmatureElementUuid == copyBufferElementElement_1_1.FirstChildElementUuid &&
		copyBufferElementElement_1_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

}

// Test to Swap in element from Copy Buffer on the TestCaseModel
func TestSwapElementFromCopyBufferCommandOnTestCaseModel(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := commandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		testcases:         nil,
	}

	// Add needed data for availableBondsMap
	tempAvailableBondsMap := make(map[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage)

	// B11f_BOND
	visibleBondAttributesMessage_AvaialbeBond_B11f_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage_VisibleBondAttributesMessage{
		BondUuid: "2858d47a-198c-43f3-abe8-abd2a36f6045",
		BondName: "B11f_BOND",
	}

	basicBondInformationMessage_AvaialbeBond_B11f_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage{
		VisibleBondAttributes: &visibleBondAttributesMessage_AvaialbeBond_B11f_BOND}

	immatureBondsMessage_ImmatureBondMessage_B11f_BOND := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage{
		BasicBondInformation: &basicBondInformationMessage_AvaialbeBond_B11f_BOND}

	tempAvailableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND] = &immatureBondsMessage_ImmatureBondMessage_B11f_BOND

	// B11l_BOND
	visibleBondAttributesMessage_AvaialbeBond_B11l_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage_VisibleBondAttributesMessage{
		BondUuid: "95cbb203-1bb3-4ab4-84b7-c2a27a2e40fb",
		BondName: "B11l_BOND",
	}

	basicBondInformationMessage_AvaialbeBond_B11l_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage{
		VisibleBondAttributes: &visibleBondAttributesMessage_AvaialbeBond_B11l_BOND}

	immatureBondsMessage_ImmatureBondMessage_B11l_BOND := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage{
		BasicBondInformation: &basicBondInformationMessage_AvaialbeBond_B11l_BOND}

	tempAvailableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND] = &immatureBondsMessage_ImmatureBondMessage_B11l_BOND

	// Add bond-map to 'commandAndRuleEngine.availableBondsMap'
	commandAndRuleEngine.availableBondsMap = tempAvailableBondsMap

	// Initiate a TestCaseModel
	myTestCaseModel := testCaseModel.TestCaseModelStruct{
		LastLoadedTestCaseModelGRPCMessage: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage{},
		FirstElementUuid:                   "",
		TestCaseModelMap:                   nil,
	}

	myTestCaseModel.TestCaseModelMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage)

	uuidToBeCopied := "4b694f8c-f194-45af-a75e-f2bb3fd350e6"
	uuidToReplacedByCopyBufferConten := "b1d535a7-e0b4-4a67-9581-f7d157f7ba1e"

	tc_b1f := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "cff02b54-ee53-47d1-94d6-48dc470073a3",
		OriginalElementName:      "B1l_BOND",
		MatureElementUuid:        "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		PreviousElementUuid:      "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		NextElementUuid:          uuidToBeCopied,
		FirstChildElementUuid:    "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		ParentElementUuid:        "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE,
	}

	tc_tic := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		MatureElementUuid:        uuidToBeCopied,
		PreviousElementUuid:      "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		NextElementUuid:          "79a6702d-8370-446c-b001-d60494eca6fa",
		FirstChildElementUuid:    uuidToReplacedByCopyBufferConten,
		ParentElementUuid:        uuidToBeCopied,
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_b1l := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c10bfd8f-2786-480c-9a71-bad9ec7bc582",
		OriginalElementName:      "B1l_BOND",
		MatureElementUuid:        "79a6702d-8370-446c-b001-d60494eca6fa",
		PreviousElementUuid:      uuidToBeCopied,
		NextElementUuid:          "79a6702d-8370-446c-b001-d60494eca6fa",
		FirstChildElementUuid:    "79a6702d-8370-446c-b001-d60494eca6fa",
		ParentElementUuid:        "79a6702d-8370-446c-b001-d60494eca6fa",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE,
	}

	tc_b10 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "4f348a88-bb9a-4852-80de-3141687db11d",
		OriginalElementName:      "B10_Bond",
		MatureElementUuid:        uuidToReplacedByCopyBufferConten,
		PreviousElementUuid:      uuidToReplacedByCopyBufferConten,
		NextElementUuid:          uuidToReplacedByCopyBufferConten,
		FirstChildElementUuid:    uuidToReplacedByCopyBufferConten,
		ParentElementUuid:        uuidToBeCopied,
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
	}

	// Set up TestCaseModel-map
	myTestCaseModel.TestCaseModelMap[tc_b1f.MatureElementUuid] = tc_b1f
	myTestCaseModel.TestCaseModelMap[tc_tic.MatureElementUuid] = tc_tic
	myTestCaseModel.TestCaseModelMap[tc_b1l.MatureElementUuid] = tc_b1l
	myTestCaseModel.TestCaseModelMap[tc_b10.MatureElementUuid] = tc_b10

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_b1f.MatureElementUuid

	// Initiate structure for all TestCases
	allTestCases := make(map[string]testCaseModel.TestCaseModelStruct)

	// Generate new UUID
	testCaseUuid := uuidGenerator.New().String()

	// Add myTestCaseModel to map of all testcases
	allTestCases[testCaseUuid] = myTestCaseModel

	// Set Current User
	currentUser := "s41797"

	// Add AddTestCases to TestCases-model
	testCasesObject := testCaseModel.TestCaseModelsStruct{
		TestCases:   allTestCases,
		CurrentUser: currentUser}

	// Add reference to TestCAses in command and rule engine
	commandAndRuleEngine.testcases = &testCasesObject

	// Execute command Copy
	err := commandAndRuleEngine.executeCommandOnTestCaseModel_CopyElementInTestCaseModel(testCaseUuid, uuidToBeCopied)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Execute command Swap from Copy Buffer
	err = commandAndRuleEngine.executeCommandOnTestCaseModel_SwapInElementFromCopyBufferInTestCaseModel(testCaseUuid, uuidToReplacedByCopyBufferConten)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Extract TestCase
	testCase := commandAndRuleEngine.testcases.TestCases[testCaseUuid]

	// Validate the result of the Copy-Element-command, 'B1f-TIC(B10)-B1l'
	// 1) Validate B1f (1)
	testCaseModelElementUuid_1 := testCase.FirstElementUuid
	testCaseModelElement_1 := testCase.TestCaseModelMap[testCaseModelElementUuid_1]

	correctElement := testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.ParentElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.PreviousElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.FirstChildElementUuid &&
		testCaseModelElement_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 2) Validate TIC (2)
	testCaseModelElementUuid_2 := testCaseModelElement_1.NextElementUuid
	testCaseModelElement_2 := testCase.TestCaseModelMap[testCaseModelElementUuid_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2.ParentElementUuid &&
		testCaseModelElement_2.MatureElementUuid == testCaseModelElement_1.NextElementUuid &&
		testCaseModelElement_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 3) Validate B1l (3)
	testCaseModelElementUuid_3 := testCaseModelElement_2.NextElementUuid
	testCaseModelElement_3 := testCase.TestCaseModelMap[testCaseModelElementUuid_3]

	correctElement = testCaseModelElement_3.MatureElementUuid == testCaseModelElement_3.ParentElementUuid &&
		testCaseModelElement_3.MatureElementUuid == testCaseModelElement_2.NextElementUuid &&
		testCaseModelElement_3.MatureElementUuid == testCaseModelElement_3.NextElementUuid &&
		testCaseModelElement_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 4) Validate B11f(2.1)
	testCaseModelElementUuid_2_1 := testCaseModelElement_2.FirstChildElementUuid
	testCaseModelElement_2_1 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_1]

	correctElement = testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2.FirstChildElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 5) Validate TIC(2.2)
	testCaseModelElementUuid_2_2 := testCaseModelElement_2_1.NextElementUuid
	testCaseModelElement_2_2 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_2]

	correctElement = testCaseModelElement_2_2.ParentElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2_2.PreviousElementUuid == testCaseModelElement_2_1.MatureElementUuid &&
		testCaseModelElement_2_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 6) Validate B11l(2.3)
	testCaseModelElementUuid_2_3 := testCaseModelElement_2_2.NextElementUuid
	testCaseModelElement_2_3 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_3]

	correctElement = testCaseModelElement_2_3.ParentElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2_3.PreviousElementUuid == testCaseModelElement_2_2.MatureElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_3.NextElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_3.FirstChildElementUuid &&
		testCaseModelElement_2_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 7) Validate B10 (2.2.1)
	testCaseModelElementUuid_2_2_1 := testCaseModelElement_2_2.FirstChildElementUuid
	testCaseModelElement_2_2_1 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_2_1]

	correctElement = testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2.FirstChildElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.PreviousElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.NextElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = commandAndRuleEngine.testcases.VerifyThatThereAreNoZombieElementsInTestCaseModel(testCaseUuid)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate Textual TestCase Presentation
	textualTestCaseSimple, textualTestCaseComplex, err := commandAndRuleEngine.testcases.CreateTextualTestCase(testCaseUuid)

	textualTestCaseRepresentationSimple := "[B1-TIC(B11-TIC(B10)-B11)-B1]"
	textualTestCaseRepresentationComplex := "[B1f-TIC(B11f-TIC(B10)-B11l)-B1l]"

	assert.Equal(t, textualTestCaseRepresentationSimple, textualTestCaseSimple)
	assert.Equal(t, textualTestCaseRepresentationComplex, textualTestCaseComplex)

	// Validate Command stack length
	lenghtIsOne := fmt.Sprint(len(testCase.CommandStack) == 2)
	assert.Equal(t, "true", lenghtIsOne)

	// Validate Command Stack content - Copy
	assert.Equal(t, fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_COPY_ELEMENT, testCase.CommandStack[0].TestCaseCommandType)
	assert.Equal(t, fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_name[int32(fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_COPY_ELEMENT)], testCase.CommandStack[0].TestCaseCommandName)
	assert.Equal(t, uuidToBeCopied, testCase.CommandStack[0].FirstParameter)
	assert.Equal(t, testCaseModel.NotApplicable, testCase.CommandStack[0].SecondParameter)
	assert.Equal(t, currentUser, testCase.CommandStack[0].UserId)

	// Validate Command Stack content - Swap from Copy Buffer
	assert.Equal(t, fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_COPY_BUFFER_ELEMENT, testCase.CommandStack[1].TestCaseCommandType)
	assert.Equal(t, fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_name[int32(fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_COPY_BUFFER_ELEMENT)], testCase.CommandStack[1].TestCaseCommandName)
	assert.Equal(t, uuidToReplacedByCopyBufferConten, testCase.CommandStack[1].FirstParameter)
	assert.Equal(t, testCaseModel.NotApplicable, testCase.CommandStack[1].SecondParameter)
	assert.Equal(t, currentUser, testCase.CommandStack[1].UserId)

	// Validate Copy Buffer content
	// 1) Validate TIC(1)
	copyBufferElementUuid_1 := testCase.CopyBuffer.FirstElementUuid
	copyBufferElementElement_1 := testCase.CopyBuffer.ImmatureElementMap[copyBufferElementUuid_1]

	correctElement = copyBufferElementElement_1.ImmatureElementUuid == copyBufferElementElement_1.ParentElementUuid &&
		copyBufferElementElement_1.ImmatureElementUuid == copyBufferElementElement_1.PreviousElementUuid &&
		copyBufferElementElement_1.ImmatureElementUuid == copyBufferElementElement_1.NextElementUuid &&
		copyBufferElementElement_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 2) Validate TIC(1.1)
	copyBufferElementUuid_1_1 := copyBufferElementElement_1.FirstChildElementUuid
	copyBufferElementElement_1_1 := testCase.CopyBuffer.ImmatureElementMap[copyBufferElementUuid_1_1]

	correctElement = copyBufferElementElement_1_1.ImmatureElementUuid == copyBufferElementElement_1.FirstChildElementUuid &&
		copyBufferElementElement_1_1.ImmatureElementUuid == copyBufferElementElement_1_1.PreviousElementUuid &&
		copyBufferElementElement_1_1.ImmatureElementUuid == copyBufferElementElement_1_1.NextElementUuid &&
		copyBufferElementElement_1_1.ImmatureElementUuid == copyBufferElementElement_1_1.FirstChildElementUuid &&
		copyBufferElementElement_1_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

}
