package commandAndRuleEngine

import (
	"FenixTesterGui/gui/UnitTestTestData"
	"FenixTesterGui/testCase/testCaseModel"
	"fmt"
	uuidGenerator "github.com/google/uuid"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TCRuleDeletion101
// What to remove			Remove in structure				Result after deletion		Rule
// n= TIC(X)				B1-n-B1							B0							TCRuleDeletion101
// Test to Delete an element from the TestCaseModel
func TestTCRuleDeletion101(t *testing.T) {

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

//TCRuleDeletion102
// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			B11f-n-B11l						B10							TCRuleDeletion102
// Test to Delete an element from the TestCaseModel
func TestTCRuleDeletion102(t *testing.T) {

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
	visibleBondAttributesMessage_AvaialbeBond_B10_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage_VisibleBondAttributesMessage{
		BondUuid: "2858d47a-198c-43f3-abe8-abd2a36f6045",
		BondName: "B10_BOND",
	}

	basicBondInformationMessage_AvaialbeBond_B10_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage{
		VisibleBondAttributes: &visibleBondAttributesMessage_AvaialbeBond_B10_BOND}

	immatureBondsMessage_ImmatureBondMessage_B10_BOND := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage{
		BasicBondInformation: &basicBondInformationMessage_AvaialbeBond_B10_BOND}

	tempAvailableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND] = &immatureBondsMessage_ImmatureBondMessage_B10_BOND

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

	tc_1_b1f := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "cff02b54-ee53-47d1-94d6-48dc470073a3",
		OriginalElementName:      "B1l_BOND",
		MatureElementUuid:        "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		PreviousElementUuid:      "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		NextElementUuid:          "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		FirstChildElementUuid:    "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		ParentElementUuid:        "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE,
	}

	tc_2_tic := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		MatureElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		PreviousElementUuid:      "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		NextElementUuid:          "79a6702d-8370-446c-b001-d60494eca6fa",
		FirstChildElementUuid:    "beea7876-f566-44c7-9625-655d8d075c5a",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_3_b1l := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c10bfd8f-2786-480c-9a71-bad9ec7bc582",
		OriginalElementName:      "B1l_BOND",
		MatureElementUuid:        "79a6702d-8370-446c-b001-d60494eca6fa",
		PreviousElementUuid:      "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		NextElementUuid:          "79a6702d-8370-446c-b001-d60494eca6fa",
		FirstChildElementUuid:    "79a6702d-8370-446c-b001-d60494eca6fa",
		ParentElementUuid:        "79a6702d-8370-446c-b001-d60494eca6fa",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE,
	}

	tc_2_1_b11f := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "f9df1310-0c47-4c0b-a5aa-5330c1589cac",
		OriginalElementName:      "B11l_BOND",
		MatureElementUuid:        "beea7876-f566-44c7-9625-655d8d075c5a",
		PreviousElementUuid:      "beea7876-f566-44c7-9625-655d8d075c5a",
		NextElementUuid:          uuidToBeDeleted,
		FirstChildElementUuid:    "beea7876-f566-44c7-9625-655d8d075c5a",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
	}

	tc_2_2_tic := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		MatureElementUuid:        uuidToBeDeleted,
		PreviousElementUuid:      "beea7876-f566-44c7-9625-655d8d075c5a",
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_2_3_b11l := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "66284917-6b5c-42a0-98d6-106a7bd269fa",
		OriginalElementName:      "B11l_BOND",
		MatureElementUuid:        "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		PreviousElementUuid:      "a39133d9-d8de-4c96-b8f6-7115aa88bfa6",
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND,
	}

	tc_2_2_1_b10 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "4f348a88-bb9a-4852-80de-3141687db11d",
		OriginalElementName:      "B10_Bond",
		MatureElementUuid:        "edaf757d-1205-4f2d-91a4-f053982f5ded",
		PreviousElementUuid:      "edaf757d-1205-4f2d-91a4-f053982f5ded",
		NextElementUuid:          "edaf757d-1205-4f2d-91a4-f053982f5ded",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        uuidToBeDeleted,
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
	}

	// Set up TestCaseModel-map
	myTestCaseModel.TestCaseModelMap[tc_1_b1f.MatureElementUuid] = tc_1_b1f
	myTestCaseModel.TestCaseModelMap[tc_2_tic.MatureElementUuid] = tc_2_tic
	myTestCaseModel.TestCaseModelMap[tc_3_b1l.MatureElementUuid] = tc_3_b1l
	myTestCaseModel.TestCaseModelMap[tc_2_1_b11f.MatureElementUuid] = tc_2_1_b11f
	myTestCaseModel.TestCaseModelMap[tc_2_2_tic.MatureElementUuid] = tc_2_2_tic
	myTestCaseModel.TestCaseModelMap[tc_2_3_b11l.MatureElementUuid] = tc_2_3_b11l
	myTestCaseModel.TestCaseModelMap[tc_2_2_1_b10.MatureElementUuid] = tc_2_2_1_b10

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_1_b1f.MatureElementUuid

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
	testCase, existsInMap := commandAndRuleEngine.testcases.TestCases[testCaseUuid]
	assert.Equal(t, "true", fmt.Sprint(existsInMap))

	// Validate the result of the swap, 'B1f-TIC(B11f-TIC(B10)-B1l'
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
		testCaseModelElement_2.PreviousElementUuid == testCaseModelElement_1.MatureElementUuid &&
		testCaseModelElement_1.NextElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 3) Validate B1l (3)
	testCaseModelElementUuid_3 := testCaseModelElement_2.NextElementUuid
	testCaseModelElement_3 := testCase.TestCaseModelMap[testCaseModelElementUuid_3]

	correctElement = testCaseModelElement_3.MatureElementUuid == testCaseModelElement_3.ParentElementUuid &&
		testCaseModelElement_3.NextElementUuid == testCaseModelElement_3.MatureElementUuid &&
		testCaseModelElement_3.FirstChildElementUuid == testCaseModelElement_3.MatureElementUuid &&
		testCaseModelElement_3.PreviousElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2.NextElementUuid == testCaseModelElement_3.MatureElementUuid &&
		testCaseModelElement_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 4) Validate B11f (2.1)
	testCaseModelElementUuid_2_1 := testCaseModelElement_2.FirstChildElementUuid
	testCaseModelElement_2_1 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_1]

	correctElement = testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2.FirstChildElementUuid &&
		testCaseModelElement_2_1.ParentElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 5) Validate TIC (2.2)
	testCaseModelElementUuid_2_2 := testCaseModelElement_2_1.NextElementUuid
	testCaseModelElement_2_2 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_2.ParentElementUuid &&
		testCaseModelElement_2_2.PreviousElementUuid == testCaseModelElement_2_1.MatureElementUuid &&
		testCaseModelElement_2_1.NextElementUuid == testCaseModelElement_2_2.MatureElementUuid &&
		testCaseModelElement_2_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 6) Validate B12 (2.3)
	testCaseModelElementUuid_2_3 := testCaseModelElement_2_2.NextElementUuid
	testCaseModelElement_2_3 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_3]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_3.ParentElementUuid &&
		testCaseModelElement_2_3.PreviousElementUuid == testCaseModelElement_2_2.MatureElementUuid &&
		testCaseModelElement_2_2.NextElementUuid == testCaseModelElement_2_3.MatureElementUuid &&
		testCaseModelElement_2_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND

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