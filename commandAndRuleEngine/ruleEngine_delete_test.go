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
	commandAndRuleEngine := CommandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		Testcases:         nil,
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

	// Add myTestCaseModel to map of all Testcases
	allTestCases[testCaseUuid] = myTestCaseModel

	// Set Current User
	currentUser := "s41797"

	// Add AddTestCases to TestCases-model
	testCasesObject := testCaseModel.TestCasesModelsStruct{
		TestCases:   allTestCases,
		CurrentUser: currentUser}

	// Add reference to TestCAses in command and rule engine
	commandAndRuleEngine.Testcases = &testCasesObject

	// Execute command
	err := commandAndRuleEngine.executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel(testCaseUuid, uuidToBeDeleted)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Extract TestCase
	testCase := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]

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
	err = commandAndRuleEngine.Testcases.VerifyThatThereAreNoZombieElementsInTestCaseModel(testCaseUuid)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate Textual TestCase Presentation
	textualTestCaseSimple, textualTestCaseComplex, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)

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
	commandAndRuleEngine := CommandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		Testcases:         nil,
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
		OriginalElementName:      "B1f_BOND",
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

	// Add myTestCaseModel to map of all Testcases
	allTestCases[testCaseUuid] = myTestCaseModel

	// Set Current User
	currentUser := "s41797"

	// Add AddTestCases to TestCases-model
	testCasesObject := testCaseModel.TestCasesModelsStruct{
		TestCases:   allTestCases,
		CurrentUser: currentUser}

	// Add reference to TestCAses in command and rule engine
	commandAndRuleEngine.Testcases = &testCasesObject

	// Execute command
	err := commandAndRuleEngine.executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel(testCaseUuid, uuidToBeDeleted)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Extract TestCase
	testCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	assert.Equal(t, "true", fmt.Sprint(existsInMap))

	// Validate the result of the swap, 'B1f-TIC(B10)-B1l'
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

	// 4) Validate B10(2.1)
	testCaseModelElementUuid_2_1 := testCaseModelElement_2.FirstChildElementUuid
	testCaseModelElement_2_1 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_1]

	correctElement = testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2.FirstChildElementUuid &&
		testCaseModelElement_2_1.ParentElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.NextElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = commandAndRuleEngine.Testcases.VerifyThatThereAreNoZombieElementsInTestCaseModel(testCaseUuid)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate Textual TestCase Presentation
	textualTestCaseSimple, textualTestCaseComplex, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)

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
	assert.Equal(t, fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_REMOVE_ELEMENT, testCase.CommandStack[0].TestCaseCommandType)
	assert.Equal(t, fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_name[int32(fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_REMOVE_ELEMENT)], testCase.CommandStack[0].TestCaseCommandName)
	assert.Equal(t, uuidToBeDeleted, testCase.CommandStack[0].FirstParameter)
	assert.Equal(t, testCaseModel.NotApplicable, testCase.CommandStack[0].SecondParameter)
	assert.Equal(t, currentUser, testCase.CommandStack[0].UserId)

}

// TestTCRuleDeletion103
// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			B11fx-n-B11lx					B10*x*						TCRuleDeletion103
// Test to Delete an element from the TestCaseModel
func TestTCRuleDeletion103(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := CommandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		Testcases:         nil,
	}

	// Add needed data for availableBondsMap
	tempAvailableBondsMap := make(map[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage)

	// B0_BOND
	visibleBondAttributesMessage_AvaialbeBond_B10oxo_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage_VisibleBondAttributesMessage{
		BondUuid: "2858d47a-198c-43f3-abe8-abd2a36f6045",
		BondName: "B10oxo_BOND",
	}

	basicBondInformationMessage_AvaialbeBond_B10oxo_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage{
		VisibleBondAttributes: &visibleBondAttributesMessage_AvaialbeBond_B10oxo_BOND}

	immatureBondsMessage_ImmatureBondMessage_B10oxo_BOND := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage{
		BasicBondInformation: &basicBondInformationMessage_AvaialbeBond_B10oxo_BOND}

	tempAvailableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10oxo_BOND] = &immatureBondsMessage_ImmatureBondMessage_B10oxo_BOND

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
		OriginalElementName:      "B1f_BOND",
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

	tc_2_1_b11fx := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "f9df1310-0c47-4c0b-a5aa-5330c1589cac",
		OriginalElementName:      "B11fx_BOND",
		MatureElementUuid:        "beea7876-f566-44c7-9625-655d8d075c5a",
		PreviousElementUuid:      "beea7876-f566-44c7-9625-655d8d075c5a",
		NextElementUuid:          uuidToBeDeleted,
		FirstChildElementUuid:    "beea7876-f566-44c7-9625-655d8d075c5a",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE,
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

	tc_2_3_b11lx := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "66284917-6b5c-42a0-98d6-106a7bd269fa",
		OriginalElementName:      "B11lx_BOND",
		MatureElementUuid:        "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		PreviousElementUuid:      "a39133d9-d8de-4c96-b8f6-7115aa88bfa6",
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE,
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
	myTestCaseModel.TestCaseModelMap[tc_2_1_b11fx.MatureElementUuid] = tc_2_1_b11fx
	myTestCaseModel.TestCaseModelMap[tc_2_2_tic.MatureElementUuid] = tc_2_2_tic
	myTestCaseModel.TestCaseModelMap[tc_2_3_b11lx.MatureElementUuid] = tc_2_3_b11lx
	myTestCaseModel.TestCaseModelMap[tc_2_2_1_b10.MatureElementUuid] = tc_2_2_1_b10

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_1_b1f.MatureElementUuid

	// Initiate structure for all TestCases
	allTestCases := make(map[string]testCaseModel.TestCaseModelStruct)

	// Generate new UUID
	testCaseUuid := uuidGenerator.New().String()

	// Add myTestCaseModel to map of all Testcases
	allTestCases[testCaseUuid] = myTestCaseModel

	// Set Current User
	currentUser := "s41797"

	// Add AddTestCases to TestCases-model
	testCasesObject := testCaseModel.TestCasesModelsStruct{
		TestCases:   allTestCases,
		CurrentUser: currentUser}

	// Add reference to TestCAses in command and rule engine
	commandAndRuleEngine.Testcases = &testCasesObject

	// Execute command
	err := commandAndRuleEngine.executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel(testCaseUuid, uuidToBeDeleted)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Extract TestCase
	testCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	assert.Equal(t, "true", fmt.Sprint(existsInMap))

	// Validate the result of the swap, 'B1f-TIC(B10)-B1l'
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

	// 4) Validate B10oxo (2.1)
	testCaseModelElementUuid_2_1 := testCaseModelElement_2.FirstChildElementUuid
	testCaseModelElement_2_1 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_1]

	correctElement = testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2.FirstChildElementUuid &&
		testCaseModelElement_2_1.ParentElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.NextElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10oxo_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = commandAndRuleEngine.Testcases.VerifyThatThereAreNoZombieElementsInTestCaseModel(testCaseUuid)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate Textual TestCase Presentation
	textualTestCaseSimple, textualTestCaseComplex, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)

	textualTestCaseRepresentationSimple := "[B1-TIC(B10x)-B1]"
	textualTestCaseRepresentationComplex := "[B1f-TIC(B10oxo)-B1l]"

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

// TCRuleDeletion104
// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			B11f-n-B11lx					B10x*						TCRuleDeletion104
// Test to Delete an element from the TestCaseModel
func TestTCRuleDeletion104(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := CommandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		Testcases:         nil,
	}

	// Add needed data for availableBondsMap
	tempAvailableBondsMap := make(map[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage)

	// B0_BOND
	visibleBondAttributesMessage_AvaialbeBond_B10xo_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage_VisibleBondAttributesMessage{
		BondUuid: "2858d47a-198c-43f3-abe8-abd2a36f6045",
		BondName: "B10xo_BOND",
	}

	basicBondInformationMessage_AvaialbeBond_B10xo_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage{
		VisibleBondAttributes: &visibleBondAttributesMessage_AvaialbeBond_B10xo_BOND}

	immatureBondsMessage_ImmatureBondMessage_B10xo_BOND := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage{
		BasicBondInformation: &basicBondInformationMessage_AvaialbeBond_B10xo_BOND}

	tempAvailableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10xo_BOND] = &immatureBondsMessage_ImmatureBondMessage_B10xo_BOND

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
		OriginalElementName:      "B1f_BOND",
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
		OriginalElementName:      "B11f_BOND",
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

	tc_2_3_b11lx := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "66284917-6b5c-42a0-98d6-106a7bd269fa",
		OriginalElementName:      "B11lx_BOND",
		MatureElementUuid:        "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		PreviousElementUuid:      "a39133d9-d8de-4c96-b8f6-7115aa88bfa6",
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE,
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
	myTestCaseModel.TestCaseModelMap[tc_2_3_b11lx.MatureElementUuid] = tc_2_3_b11lx
	myTestCaseModel.TestCaseModelMap[tc_2_2_1_b10.MatureElementUuid] = tc_2_2_1_b10

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_1_b1f.MatureElementUuid

	// Initiate structure for all TestCases
	allTestCases := make(map[string]testCaseModel.TestCaseModelStruct)

	// Generate new UUID
	testCaseUuid := uuidGenerator.New().String()

	// Add myTestCaseModel to map of all Testcases
	allTestCases[testCaseUuid] = myTestCaseModel

	// Set Current User
	currentUser := "s41797"

	// Add AddTestCases to TestCases-model
	testCasesObject := testCaseModel.TestCasesModelsStruct{
		TestCases:   allTestCases,
		CurrentUser: currentUser}

	// Add reference to TestCAses in command and rule engine
	commandAndRuleEngine.Testcases = &testCasesObject

	// Execute command
	err := commandAndRuleEngine.executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel(testCaseUuid, uuidToBeDeleted)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Extract TestCase
	testCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	assert.Equal(t, "true", fmt.Sprint(existsInMap))

	// Validate the result of the swap, 'B1f-TIC(B10)-B1l'
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

	// 4) Validate B10xo (2.1)
	testCaseModelElementUuid_2_1 := testCaseModelElement_2.FirstChildElementUuid
	testCaseModelElement_2_1 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_1]

	correctElement = testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2.FirstChildElementUuid &&
		testCaseModelElement_2_1.ParentElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.NextElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10xo_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = commandAndRuleEngine.Testcases.VerifyThatThereAreNoZombieElementsInTestCaseModel(testCaseUuid)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate Textual TestCase Presentation
	textualTestCaseSimple, textualTestCaseComplex, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)

	textualTestCaseRepresentationSimple := "[B1-TIC(B10x)-B1]"
	textualTestCaseRepresentationComplex := "[B1f-TIC(B10xo)-B1l]"

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

// TCRuleDeletion105
// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			B11fx-n-B11l					B10*x						TCRuleDeletion105
func TestTCRuleDeletion105(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := CommandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		Testcases:         nil,
	}

	// Add needed data for availableBondsMap
	tempAvailableBondsMap := make(map[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage)

	// B0_BOND
	visibleBondAttributesMessage_AvaialbeBond_B10ox_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage_VisibleBondAttributesMessage{
		BondUuid: "2858d47a-198c-43f3-abe8-abd2a36f6045",
		BondName: "B10ox_BOND",
	}

	basicBondInformationMessage_AvaialbeBond_B10ox_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage{
		VisibleBondAttributes: &visibleBondAttributesMessage_AvaialbeBond_B10ox_BOND}

	immatureBondsMessage_ImmatureBondMessage_B10ox_BOND := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage{
		BasicBondInformation: &basicBondInformationMessage_AvaialbeBond_B10ox_BOND}

	tempAvailableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10ox_BOND] = &immatureBondsMessage_ImmatureBondMessage_B10ox_BOND

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
		OriginalElementName:      "B1f_BOND",
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

	tc_2_1_b11fx := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "f9df1310-0c47-4c0b-a5aa-5330c1589cac",
		OriginalElementName:      "B11fx_BOND",
		MatureElementUuid:        "beea7876-f566-44c7-9625-655d8d075c5a",
		PreviousElementUuid:      "beea7876-f566-44c7-9625-655d8d075c5a",
		NextElementUuid:          uuidToBeDeleted,
		FirstChildElementUuid:    "beea7876-f566-44c7-9625-655d8d075c5a",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE,
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
	myTestCaseModel.TestCaseModelMap[tc_2_1_b11fx.MatureElementUuid] = tc_2_1_b11fx
	myTestCaseModel.TestCaseModelMap[tc_2_2_tic.MatureElementUuid] = tc_2_2_tic
	myTestCaseModel.TestCaseModelMap[tc_2_3_b11l.MatureElementUuid] = tc_2_3_b11l
	myTestCaseModel.TestCaseModelMap[tc_2_2_1_b10.MatureElementUuid] = tc_2_2_1_b10

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_1_b1f.MatureElementUuid

	// Initiate structure for all TestCases
	allTestCases := make(map[string]testCaseModel.TestCaseModelStruct)

	// Generate new UUID
	testCaseUuid := uuidGenerator.New().String()

	// Add myTestCaseModel to map of all Testcases
	allTestCases[testCaseUuid] = myTestCaseModel

	// Set Current User
	currentUser := "s41797"

	// Add AddTestCases to TestCases-model
	testCasesObject := testCaseModel.TestCasesModelsStruct{
		TestCases:   allTestCases,
		CurrentUser: currentUser}

	// Add reference to TestCAses in command and rule engine
	commandAndRuleEngine.Testcases = &testCasesObject

	// Execute command
	err := commandAndRuleEngine.executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel(testCaseUuid, uuidToBeDeleted)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Extract TestCase
	testCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	assert.Equal(t, "true", fmt.Sprint(existsInMap))

	// Validate the result of the swap, 'B1f-TIC(B10)-B1l'
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

	// 4) Validate B10xo (2.1)
	testCaseModelElementUuid_2_1 := testCaseModelElement_2.FirstChildElementUuid
	testCaseModelElement_2_1 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_1]

	correctElement = testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2.FirstChildElementUuid &&
		testCaseModelElement_2_1.ParentElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.NextElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10ox_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = commandAndRuleEngine.Testcases.VerifyThatThereAreNoZombieElementsInTestCaseModel(testCaseUuid)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate Textual TestCase Presentation
	textualTestCaseSimple, textualTestCaseComplex, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)

	textualTestCaseRepresentationSimple := "[B1-TIC(B10x)-B1]"
	textualTestCaseRepresentationComplex := "[B1f-TIC(B10ox)-B1l]"

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

// TCRuleDeletion106
// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			B11f-n-B12-X					B11f-X						TCRuleDeletion106
func TestTCRuleDeletion106(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := CommandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		Testcases:         nil,
	}

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
		OriginalElementName:      "B1f_BOND",
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
		OriginalElementName:      "B11f_BOND",
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
		NextElementUuid:          "269a01da-7112-496a-82bf-da3cf30c03da",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_2_3_b12 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "6b9c5679-95f8-4e6b-87b6-e8083d47f66e",
		OriginalElementName:      "B12_BOND",
		MatureElementUuid:        "269a01da-7112-496a-82bf-da3cf30c03da",
		PreviousElementUuid:      uuidToBeDeleted,
		NextElementUuid:          "c690d7cd-1239-4bb5-b87f-f54be6a716ec",
		FirstChildElementUuid:    "269a01da-7112-496a-82bf-da3cf30c03da",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND,
	}

	tc_2_4_ti := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "f1a1294b-72a5-447d-af55-8d8ef746eec1",
		OriginalElementName:      "TI",
		MatureElementUuid:        "c690d7cd-1239-4bb5-b87f-f54be6a716ec",
		PreviousElementUuid:      "269a01da-7112-496a-82bf-da3cf30c03da",
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "c690d7cd-1239-4bb5-b87f-f54be6a716ec",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
	}

	tc_2_5_b11l := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "66284917-6b5c-42a0-98d6-106a7bd269fa",
		OriginalElementName:      "B11l_BOND",
		MatureElementUuid:        "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		PreviousElementUuid:      "c690d7cd-1239-4bb5-b87f-f54be6a716ec",
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND,
	}

	tc_2_2_1_b10 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "4f348a88-bb9a-4852-80de-3141687db11d",
		OriginalElementName:      "B10_BOND",
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
	myTestCaseModel.TestCaseModelMap[tc_2_3_b12.MatureElementUuid] = tc_2_3_b12
	myTestCaseModel.TestCaseModelMap[tc_2_4_ti.MatureElementUuid] = tc_2_4_ti
	myTestCaseModel.TestCaseModelMap[tc_2_5_b11l.MatureElementUuid] = tc_2_5_b11l
	myTestCaseModel.TestCaseModelMap[tc_2_2_1_b10.MatureElementUuid] = tc_2_2_1_b10

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_1_b1f.MatureElementUuid

	// Initiate structure for all TestCases
	allTestCases := make(map[string]testCaseModel.TestCaseModelStruct)

	// Generate new UUID
	testCaseUuid := uuidGenerator.New().String()

	// Add myTestCaseModel to map of all Testcases
	allTestCases[testCaseUuid] = myTestCaseModel

	// Set Current User
	currentUser := "s41797"

	// Add AddTestCases to TestCases-model
	testCasesObject := testCaseModel.TestCasesModelsStruct{
		TestCases:   allTestCases,
		CurrentUser: currentUser}

	// Add reference to TestCAses in command and rule engine
	commandAndRuleEngine.Testcases = &testCasesObject

	// Execute command
	err := commandAndRuleEngine.executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel(testCaseUuid, uuidToBeDeleted)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Extract TestCase
	testCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	assert.Equal(t, "true", fmt.Sprint(existsInMap))

	// Validate the result of the swap, 'B1f-TIC(B11f-TI-B11l)-B1l'
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

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_1.ParentElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 5) Validate TI (2.2)
	testCaseModelElementUuid_2_2 := testCaseModelElement_2_1.NextElementUuid
	testCaseModelElement_2_2 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_2.ParentElementUuid &&
		testCaseModelElement_2_2.MatureElementUuid == testCaseModelElement_2_1.NextElementUuid &&
		testCaseModelElement_2_2.MatureElementUuid == testCaseModelElement_2_2.FirstChildElementUuid &&
		testCaseModelElement_2_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 6) Validate B11l (2.3)
	testCaseModelElementUuid_2_3 := testCaseModelElement_2_2.NextElementUuid
	testCaseModelElement_2_3 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_3]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_3.ParentElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_2.NextElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_3.FirstChildElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_3.NextElementUuid &&
		testCaseModelElement_2_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = commandAndRuleEngine.Testcases.VerifyThatThereAreNoZombieElementsInTestCaseModel(testCaseUuid)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate Textual TestCase Presentation
	textualTestCaseSimple, textualTestCaseComplex, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)

	textualTestCaseRepresentationSimple := "[B1-TIC(B11-TI-B11)-B1]"
	textualTestCaseRepresentationComplex := "[B1f-TIC(B11f-TI-B11l)-B1l]"

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

// TCRuleDeletion107
// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			B11fx-n-B12x-X					B11fx-X						TCRuleDeletion107
func TestTCRuleDeletion107(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := CommandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		Testcases:         nil,
	}

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
		OriginalElementName:      "B1f_BOND",
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

	tc_2_1_b11fx := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "f9df1310-0c47-4c0b-a5aa-5330c1589cac",
		OriginalElementName:      "B11fx_BOND",
		MatureElementUuid:        "beea7876-f566-44c7-9625-655d8d075c5a",
		PreviousElementUuid:      "beea7876-f566-44c7-9625-655d8d075c5a",
		NextElementUuid:          uuidToBeDeleted,
		FirstChildElementUuid:    "beea7876-f566-44c7-9625-655d8d075c5a",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE,
	}

	tc_2_2_tic := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		MatureElementUuid:        uuidToBeDeleted,
		PreviousElementUuid:      "beea7876-f566-44c7-9625-655d8d075c5a",
		NextElementUuid:          "269a01da-7112-496a-82bf-da3cf30c03da",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_2_3_b12x := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "6b9c5679-95f8-4e6b-87b6-e8083d47f66e",
		OriginalElementName:      "B12x_BOND",
		MatureElementUuid:        "269a01da-7112-496a-82bf-da3cf30c03da",
		PreviousElementUuid:      uuidToBeDeleted,
		NextElementUuid:          "c690d7cd-1239-4bb5-b87f-f54be6a716ec",
		FirstChildElementUuid:    "269a01da-7112-496a-82bf-da3cf30c03da",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE,
	}

	tc_2_4_ti := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "f1a1294b-72a5-447d-af55-8d8ef746eec1",
		OriginalElementName:      "TI",
		MatureElementUuid:        "c690d7cd-1239-4bb5-b87f-f54be6a716ec",
		PreviousElementUuid:      "269a01da-7112-496a-82bf-da3cf30c03da",
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "c690d7cd-1239-4bb5-b87f-f54be6a716ec",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
	}

	tc_2_5_b11l := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "66284917-6b5c-42a0-98d6-106a7bd269fa",
		OriginalElementName:      "B11l_BOND",
		MatureElementUuid:        "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		PreviousElementUuid:      "c690d7cd-1239-4bb5-b87f-f54be6a716ec",
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND,
	}

	tc_2_2_1_b10 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "4f348a88-bb9a-4852-80de-3141687db11d",
		OriginalElementName:      "B10_BOND",
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
	myTestCaseModel.TestCaseModelMap[tc_2_1_b11fx.MatureElementUuid] = tc_2_1_b11fx
	myTestCaseModel.TestCaseModelMap[tc_2_2_tic.MatureElementUuid] = tc_2_2_tic
	myTestCaseModel.TestCaseModelMap[tc_2_3_b12x.MatureElementUuid] = tc_2_3_b12x
	myTestCaseModel.TestCaseModelMap[tc_2_4_ti.MatureElementUuid] = tc_2_4_ti
	myTestCaseModel.TestCaseModelMap[tc_2_5_b11l.MatureElementUuid] = tc_2_5_b11l
	myTestCaseModel.TestCaseModelMap[tc_2_2_1_b10.MatureElementUuid] = tc_2_2_1_b10

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_1_b1f.MatureElementUuid

	// Initiate structure for all TestCases
	allTestCases := make(map[string]testCaseModel.TestCaseModelStruct)

	// Generate new UUID
	testCaseUuid := uuidGenerator.New().String()

	// Add myTestCaseModel to map of all Testcases
	allTestCases[testCaseUuid] = myTestCaseModel

	// Set Current User
	currentUser := "s41797"

	// Add AddTestCases to TestCases-model
	testCasesObject := testCaseModel.TestCasesModelsStruct{
		TestCases:   allTestCases,
		CurrentUser: currentUser}

	// Add reference to TestCAses in command and rule engine
	commandAndRuleEngine.Testcases = &testCasesObject

	// Execute command
	err := commandAndRuleEngine.executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel(testCaseUuid, uuidToBeDeleted)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Extract TestCase
	testCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	assert.Equal(t, "true", fmt.Sprint(existsInMap))

	// Validate the result of the swap, 'B1f-TIC(B11fx-TI-B11l)-B1l'
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

	// 4) Validate B11fx (2.1)
	testCaseModelElementUuid_2_1 := testCaseModelElement_2.FirstChildElementUuid
	testCaseModelElement_2_1 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_1]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_1.ParentElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 5) Validate TI (2.2)
	testCaseModelElementUuid_2_2 := testCaseModelElement_2_1.NextElementUuid
	testCaseModelElement_2_2 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_2.ParentElementUuid &&
		testCaseModelElement_2_2.MatureElementUuid == testCaseModelElement_2_1.NextElementUuid &&
		testCaseModelElement_2_2.MatureElementUuid == testCaseModelElement_2_2.FirstChildElementUuid &&
		testCaseModelElement_2_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 6) Validate B11l (2.3)
	testCaseModelElementUuid_2_3 := testCaseModelElement_2_2.NextElementUuid
	testCaseModelElement_2_3 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_3]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_3.ParentElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_2.NextElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_3.FirstChildElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_3.NextElementUuid &&
		testCaseModelElement_2_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = commandAndRuleEngine.Testcases.VerifyThatThereAreNoZombieElementsInTestCaseModel(testCaseUuid)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate Textual TestCase Presentation
	textualTestCaseSimple, textualTestCaseComplex, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)

	textualTestCaseRepresentationSimple := "[B1-TIC(B11x-TI-B11)-B1]"
	textualTestCaseRepresentationComplex := "[B1f-TIC(B11fx-TI-B11l)-B1l]"

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

// TCRuleDeletion108
// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			B11f-n-B12x-X					B11fx-X						TCRuleDeletion108
func TestTCRuleDeletion108(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := CommandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		Testcases:         nil,
	}

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
		OriginalElementName:      "B1f_BOND",
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
		OriginalElementName:      "B11f_BOND",
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
		NextElementUuid:          "269a01da-7112-496a-82bf-da3cf30c03da",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_2_3_b12x := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "6b9c5679-95f8-4e6b-87b6-e8083d47f66e",
		OriginalElementName:      "B12x_BOND",
		MatureElementUuid:        "269a01da-7112-496a-82bf-da3cf30c03da",
		PreviousElementUuid:      uuidToBeDeleted,
		NextElementUuid:          "c690d7cd-1239-4bb5-b87f-f54be6a716ec",
		FirstChildElementUuid:    "269a01da-7112-496a-82bf-da3cf30c03da",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE,
	}

	tc_2_4_ti := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "f1a1294b-72a5-447d-af55-8d8ef746eec1",
		OriginalElementName:      "TI",
		MatureElementUuid:        "c690d7cd-1239-4bb5-b87f-f54be6a716ec",
		PreviousElementUuid:      "269a01da-7112-496a-82bf-da3cf30c03da",
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "c690d7cd-1239-4bb5-b87f-f54be6a716ec",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
	}

	tc_2_5_b11l := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "66284917-6b5c-42a0-98d6-106a7bd269fa",
		OriginalElementName:      "B11l_BOND",
		MatureElementUuid:        "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		PreviousElementUuid:      "c690d7cd-1239-4bb5-b87f-f54be6a716ec",
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND,
	}

	tc_2_2_1_b10 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "4f348a88-bb9a-4852-80de-3141687db11d",
		OriginalElementName:      "B10_BOND",
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
	myTestCaseModel.TestCaseModelMap[tc_2_3_b12x.MatureElementUuid] = tc_2_3_b12x
	myTestCaseModel.TestCaseModelMap[tc_2_4_ti.MatureElementUuid] = tc_2_4_ti
	myTestCaseModel.TestCaseModelMap[tc_2_5_b11l.MatureElementUuid] = tc_2_5_b11l
	myTestCaseModel.TestCaseModelMap[tc_2_2_1_b10.MatureElementUuid] = tc_2_2_1_b10

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_1_b1f.MatureElementUuid

	// Initiate structure for all TestCases
	allTestCases := make(map[string]testCaseModel.TestCaseModelStruct)

	// Generate new UUID
	testCaseUuid := uuidGenerator.New().String()

	// Add myTestCaseModel to map of all Testcases
	allTestCases[testCaseUuid] = myTestCaseModel

	// Set Current User
	currentUser := "s41797"

	// Add AddTestCases to TestCases-model
	testCasesObject := testCaseModel.TestCasesModelsStruct{
		TestCases:   allTestCases,
		CurrentUser: currentUser}

	// Add reference to TestCAses in command and rule engine
	commandAndRuleEngine.Testcases = &testCasesObject

	// Execute command
	err := commandAndRuleEngine.executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel(testCaseUuid, uuidToBeDeleted)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Extract TestCase
	testCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	assert.Equal(t, "true", fmt.Sprint(existsInMap))

	// Validate the result of the swap, 'B1f-TIC(B11fx-TI-B11l)-B1l'
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

	// 4) Validate B11fx (2.1)
	testCaseModelElementUuid_2_1 := testCaseModelElement_2.FirstChildElementUuid
	testCaseModelElement_2_1 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_1]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_1.ParentElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 5) Validate TI (2.2)
	testCaseModelElementUuid_2_2 := testCaseModelElement_2_1.NextElementUuid
	testCaseModelElement_2_2 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_2.ParentElementUuid &&
		testCaseModelElement_2_2.MatureElementUuid == testCaseModelElement_2_1.NextElementUuid &&
		testCaseModelElement_2_2.MatureElementUuid == testCaseModelElement_2_2.FirstChildElementUuid &&
		testCaseModelElement_2_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 6) Validate B11l (2.3)
	testCaseModelElementUuid_2_3 := testCaseModelElement_2_2.NextElementUuid
	testCaseModelElement_2_3 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_3]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_3.ParentElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_2.NextElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_3.FirstChildElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_3.NextElementUuid &&
		testCaseModelElement_2_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = commandAndRuleEngine.Testcases.VerifyThatThereAreNoZombieElementsInTestCaseModel(testCaseUuid)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate Textual TestCase Presentation
	textualTestCaseSimple, textualTestCaseComplex, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)

	textualTestCaseRepresentationSimple := "[B1-TIC(B11x-TI-B11)-B1]"
	textualTestCaseRepresentationComplex := "[B1f-TIC(B11fx-TI-B11l)-B1l]"

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

// TCRuleDeletion109
// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			B11fx-n-B12-X					B11fx-X						TCRuleDeletion109
func TestTCRuleDeletion109(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := CommandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		Testcases:         nil,
	}

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
		OriginalElementName:      "B1f_BOND",
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

	tc_2_1_b11fx := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "f9df1310-0c47-4c0b-a5aa-5330c1589cac",
		OriginalElementName:      "B11f_BOND",
		MatureElementUuid:        "beea7876-f566-44c7-9625-655d8d075c5a",
		PreviousElementUuid:      "beea7876-f566-44c7-9625-655d8d075c5a",
		NextElementUuid:          uuidToBeDeleted,
		FirstChildElementUuid:    "beea7876-f566-44c7-9625-655d8d075c5a",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE,
	}

	tc_2_2_tic := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		MatureElementUuid:        uuidToBeDeleted,
		PreviousElementUuid:      "beea7876-f566-44c7-9625-655d8d075c5a",
		NextElementUuid:          "269a01da-7112-496a-82bf-da3cf30c03da",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_2_3_b12 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "6b9c5679-95f8-4e6b-87b6-e8083d47f66e",
		OriginalElementName:      "B12x_BOND",
		MatureElementUuid:        "269a01da-7112-496a-82bf-da3cf30c03da",
		PreviousElementUuid:      uuidToBeDeleted,
		NextElementUuid:          "c690d7cd-1239-4bb5-b87f-f54be6a716ec",
		FirstChildElementUuid:    "269a01da-7112-496a-82bf-da3cf30c03da",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND,
	}

	tc_2_4_ti := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "f1a1294b-72a5-447d-af55-8d8ef746eec1",
		OriginalElementName:      "TI",
		MatureElementUuid:        "c690d7cd-1239-4bb5-b87f-f54be6a716ec",
		PreviousElementUuid:      "269a01da-7112-496a-82bf-da3cf30c03da",
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "c690d7cd-1239-4bb5-b87f-f54be6a716ec",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
	}

	tc_2_5_b11l := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "66284917-6b5c-42a0-98d6-106a7bd269fa",
		OriginalElementName:      "B11l_BOND",
		MatureElementUuid:        "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		PreviousElementUuid:      "c690d7cd-1239-4bb5-b87f-f54be6a716ec",
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND,
	}

	tc_2_2_1_b10 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "4f348a88-bb9a-4852-80de-3141687db11d",
		OriginalElementName:      "B10_BOND",
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
	myTestCaseModel.TestCaseModelMap[tc_2_1_b11fx.MatureElementUuid] = tc_2_1_b11fx
	myTestCaseModel.TestCaseModelMap[tc_2_2_tic.MatureElementUuid] = tc_2_2_tic
	myTestCaseModel.TestCaseModelMap[tc_2_3_b12.MatureElementUuid] = tc_2_3_b12
	myTestCaseModel.TestCaseModelMap[tc_2_4_ti.MatureElementUuid] = tc_2_4_ti
	myTestCaseModel.TestCaseModelMap[tc_2_5_b11l.MatureElementUuid] = tc_2_5_b11l
	myTestCaseModel.TestCaseModelMap[tc_2_2_1_b10.MatureElementUuid] = tc_2_2_1_b10

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_1_b1f.MatureElementUuid

	// Initiate structure for all TestCases
	allTestCases := make(map[string]testCaseModel.TestCaseModelStruct)

	// Generate new UUID
	testCaseUuid := uuidGenerator.New().String()

	// Add myTestCaseModel to map of all Testcases
	allTestCases[testCaseUuid] = myTestCaseModel

	// Set Current User
	currentUser := "s41797"

	// Add AddTestCases to TestCases-model
	testCasesObject := testCaseModel.TestCasesModelsStruct{
		TestCases:   allTestCases,
		CurrentUser: currentUser}

	// Add reference to TestCAses in command and rule engine
	commandAndRuleEngine.Testcases = &testCasesObject

	// Execute command
	err := commandAndRuleEngine.executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel(testCaseUuid, uuidToBeDeleted)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Extract TestCase
	testCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	assert.Equal(t, "true", fmt.Sprint(existsInMap))

	// Validate the result of the swap, 'B1f-TIC(B11fx-TI-B11l)-B1l'
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

	// 4) Validate B11fx (2.1)
	testCaseModelElementUuid_2_1 := testCaseModelElement_2.FirstChildElementUuid
	testCaseModelElement_2_1 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_1]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_1.ParentElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 5) Validate TI (2.2)
	testCaseModelElementUuid_2_2 := testCaseModelElement_2_1.NextElementUuid
	testCaseModelElement_2_2 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_2.ParentElementUuid &&
		testCaseModelElement_2_2.MatureElementUuid == testCaseModelElement_2_1.NextElementUuid &&
		testCaseModelElement_2_2.MatureElementUuid == testCaseModelElement_2_2.FirstChildElementUuid &&
		testCaseModelElement_2_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 6) Validate B11l (2.3)
	testCaseModelElementUuid_2_3 := testCaseModelElement_2_2.NextElementUuid
	testCaseModelElement_2_3 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_3]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_3.ParentElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_2.NextElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_3.FirstChildElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_3.NextElementUuid &&
		testCaseModelElement_2_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = commandAndRuleEngine.Testcases.VerifyThatThereAreNoZombieElementsInTestCaseModel(testCaseUuid)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate Textual TestCase Presentation
	textualTestCaseSimple, textualTestCaseComplex, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)

	textualTestCaseRepresentationSimple := "[B1-TIC(B11x-TI-B11)-B1]"
	textualTestCaseRepresentationComplex := "[B1f-TIC(B11fx-TI-B11l)-B1l]"

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

// TCRuleDeletion110
// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-B12-n-B11l					X-B11l						TCRuleDeletion110
func TestTCRuleDeletion110(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := CommandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		Testcases:         nil,
	}

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
		OriginalElementName:      "B1f_BOND",
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
		OriginalElementName:      "B11f_BOND",
		MatureElementUuid:        "beea7876-f566-44c7-9625-655d8d075c5a",
		PreviousElementUuid:      "beea7876-f566-44c7-9625-655d8d075c5a",
		NextElementUuid:          "1780b088-ac62-4d13-ab55-b370017690ba",
		FirstChildElementUuid:    "beea7876-f566-44c7-9625-655d8d075c5a",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
	}

	tc_2_2_tic := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		MatureElementUuid:        "1780b088-ac62-4d13-ab55-b370017690ba",
		PreviousElementUuid:      "beea7876-f566-44c7-9625-655d8d075c5a",
		NextElementUuid:          "269a01da-7112-496a-82bf-da3cf30c03da",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_2_3_b12 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "6b9c5679-95f8-4e6b-87b6-e8083d47f66e",
		OriginalElementName:      "B12_BOND",
		MatureElementUuid:        "269a01da-7112-496a-82bf-da3cf30c03da",
		PreviousElementUuid:      "1780b088-ac62-4d13-ab55-b370017690ba",
		NextElementUuid:          uuidToBeDeleted,
		FirstChildElementUuid:    "269a01da-7112-496a-82bf-da3cf30c03da",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND,
	}

	tc_2_4_ti := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "f1a1294b-72a5-447d-af55-8d8ef746eec1",
		OriginalElementName:      "TI",
		MatureElementUuid:        uuidToBeDeleted,
		PreviousElementUuid:      "269a01da-7112-496a-82bf-da3cf30c03da",
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    uuidToBeDeleted,
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
	}

	tc_2_5_b11l := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "66284917-6b5c-42a0-98d6-106a7bd269fa",
		OriginalElementName:      "B11l_BOND",
		MatureElementUuid:        "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		PreviousElementUuid:      uuidToBeDeleted,
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND,
	}

	tc_2_2_1_b10 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "4f348a88-bb9a-4852-80de-3141687db11d",
		OriginalElementName:      "B10_BOND",
		MatureElementUuid:        "edaf757d-1205-4f2d-91a4-f053982f5ded",
		PreviousElementUuid:      "edaf757d-1205-4f2d-91a4-f053982f5ded",
		NextElementUuid:          "edaf757d-1205-4f2d-91a4-f053982f5ded",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "1780b088-ac62-4d13-ab55-b370017690ba",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
	}

	// Set up TestCaseModel-map
	myTestCaseModel.TestCaseModelMap[tc_1_b1f.MatureElementUuid] = tc_1_b1f
	myTestCaseModel.TestCaseModelMap[tc_2_tic.MatureElementUuid] = tc_2_tic
	myTestCaseModel.TestCaseModelMap[tc_3_b1l.MatureElementUuid] = tc_3_b1l
	myTestCaseModel.TestCaseModelMap[tc_2_1_b11f.MatureElementUuid] = tc_2_1_b11f
	myTestCaseModel.TestCaseModelMap[tc_2_2_tic.MatureElementUuid] = tc_2_2_tic
	myTestCaseModel.TestCaseModelMap[tc_2_3_b12.MatureElementUuid] = tc_2_3_b12
	myTestCaseModel.TestCaseModelMap[tc_2_4_ti.MatureElementUuid] = tc_2_4_ti
	myTestCaseModel.TestCaseModelMap[tc_2_5_b11l.MatureElementUuid] = tc_2_5_b11l
	myTestCaseModel.TestCaseModelMap[tc_2_2_1_b10.MatureElementUuid] = tc_2_2_1_b10

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_1_b1f.MatureElementUuid

	// Initiate structure for all TestCases
	allTestCases := make(map[string]testCaseModel.TestCaseModelStruct)

	// Generate new UUID
	testCaseUuid := uuidGenerator.New().String()

	// Add myTestCaseModel to map of all Testcases
	allTestCases[testCaseUuid] = myTestCaseModel

	// Set Current User
	currentUser := "s41797"

	// Add AddTestCases to TestCases-model
	testCasesObject := testCaseModel.TestCasesModelsStruct{
		TestCases:   allTestCases,
		CurrentUser: currentUser}

	// Add reference to TestCAses in command and rule engine
	commandAndRuleEngine.Testcases = &testCasesObject

	// Execute command
	err := commandAndRuleEngine.executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel(testCaseUuid, uuidToBeDeleted)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Extract TestCase
	testCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	assert.Equal(t, "true", fmt.Sprint(existsInMap))

	// Validate the result of the swap, 'B1f-TIC(B11f-TIC(B10)-B11l)-B1l'
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

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_1.ParentElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 5) Validate TIC (2.2)
	testCaseModelElementUuid_2_2 := testCaseModelElement_2_1.NextElementUuid
	testCaseModelElement_2_2 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_2.ParentElementUuid &&
		testCaseModelElement_2_2.PreviousElementUuid == testCaseModelElement_2_1.MatureElementUuid &&
		testCaseModelElement_2_2.MatureElementUuid == testCaseModelElement_2_1.NextElementUuid &&
		testCaseModelElement_2_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 6) Validate B11l (2.3)
	testCaseModelElementUuid_2_3 := testCaseModelElement_2_2.NextElementUuid
	testCaseModelElement_2_3 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_3]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_3.ParentElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_2.NextElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_3.FirstChildElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_3.NextElementUuid &&
		testCaseModelElement_2_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 6) Validate B11l (2.2.1)
	testCaseModelElementUuid_2_2_1 := testCaseModelElement_2_2.FirstChildElementUuid
	testCaseModelElement_2_2_1 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_2_1]

	correctElement = testCaseModelElement_2_2.MatureElementUuid == testCaseModelElement_2_2_1.ParentElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.PreviousElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.NextElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = commandAndRuleEngine.Testcases.VerifyThatThereAreNoZombieElementsInTestCaseModel(testCaseUuid)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate Textual TestCase Presentation
	textualTestCaseSimple, textualTestCaseComplex, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)

	textualTestCaseRepresentationSimple := "[B1-TIC(B11-TIC(B10)-B11)-B1]"
	textualTestCaseRepresentationComplex := "[B1f-TIC(B11f-TIC(B10)-B11l)-B1l]"

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

// TCRuleDeletion111
// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-B12x-n-B11lx					X-B11lx						TCRuleDeletion111
func TestTCRuleDeletion111(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := CommandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		Testcases:         nil,
	}

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
		OriginalElementName:      "B1f_BOND",
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
		OriginalElementName:      "B11f_BOND",
		MatureElementUuid:        "beea7876-f566-44c7-9625-655d8d075c5a",
		PreviousElementUuid:      "beea7876-f566-44c7-9625-655d8d075c5a",
		NextElementUuid:          "1780b088-ac62-4d13-ab55-b370017690ba",
		FirstChildElementUuid:    "beea7876-f566-44c7-9625-655d8d075c5a",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
	}

	tc_2_2_tic := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		MatureElementUuid:        "1780b088-ac62-4d13-ab55-b370017690ba",
		PreviousElementUuid:      "beea7876-f566-44c7-9625-655d8d075c5a",
		NextElementUuid:          "269a01da-7112-496a-82bf-da3cf30c03da",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_2_3_b12x := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "6b9c5679-95f8-4e6b-87b6-e8083d47f66e",
		OriginalElementName:      "B12x_BOND",
		MatureElementUuid:        "269a01da-7112-496a-82bf-da3cf30c03da",
		PreviousElementUuid:      "1780b088-ac62-4d13-ab55-b370017690ba",
		NextElementUuid:          uuidToBeDeleted,
		FirstChildElementUuid:    "269a01da-7112-496a-82bf-da3cf30c03da",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE,
	}

	tc_2_4_ti := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "f1a1294b-72a5-447d-af55-8d8ef746eec1",
		OriginalElementName:      "TI",
		MatureElementUuid:        uuidToBeDeleted,
		PreviousElementUuid:      "269a01da-7112-496a-82bf-da3cf30c03da",
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    uuidToBeDeleted,
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
	}

	tc_2_5_b11lx := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "66284917-6b5c-42a0-98d6-106a7bd269fa",
		OriginalElementName:      "B11lx_BOND",
		MatureElementUuid:        "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		PreviousElementUuid:      uuidToBeDeleted,
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE,
	}

	tc_2_2_1_b10 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "4f348a88-bb9a-4852-80de-3141687db11d",
		OriginalElementName:      "B10_BOND",
		MatureElementUuid:        "edaf757d-1205-4f2d-91a4-f053982f5ded",
		PreviousElementUuid:      "edaf757d-1205-4f2d-91a4-f053982f5ded",
		NextElementUuid:          "edaf757d-1205-4f2d-91a4-f053982f5ded",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "1780b088-ac62-4d13-ab55-b370017690ba",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
	}

	// Set up TestCaseModel-map
	myTestCaseModel.TestCaseModelMap[tc_1_b1f.MatureElementUuid] = tc_1_b1f
	myTestCaseModel.TestCaseModelMap[tc_2_tic.MatureElementUuid] = tc_2_tic
	myTestCaseModel.TestCaseModelMap[tc_3_b1l.MatureElementUuid] = tc_3_b1l
	myTestCaseModel.TestCaseModelMap[tc_2_1_b11f.MatureElementUuid] = tc_2_1_b11f
	myTestCaseModel.TestCaseModelMap[tc_2_2_tic.MatureElementUuid] = tc_2_2_tic
	myTestCaseModel.TestCaseModelMap[tc_2_3_b12x.MatureElementUuid] = tc_2_3_b12x
	myTestCaseModel.TestCaseModelMap[tc_2_4_ti.MatureElementUuid] = tc_2_4_ti
	myTestCaseModel.TestCaseModelMap[tc_2_5_b11lx.MatureElementUuid] = tc_2_5_b11lx
	myTestCaseModel.TestCaseModelMap[tc_2_2_1_b10.MatureElementUuid] = tc_2_2_1_b10

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_1_b1f.MatureElementUuid

	// Initiate structure for all TestCases
	allTestCases := make(map[string]testCaseModel.TestCaseModelStruct)

	// Generate new UUID
	testCaseUuid := uuidGenerator.New().String()

	// Add myTestCaseModel to map of all Testcases
	allTestCases[testCaseUuid] = myTestCaseModel

	// Set Current User
	currentUser := "s41797"

	// Add AddTestCases to TestCases-model
	testCasesObject := testCaseModel.TestCasesModelsStruct{
		TestCases:   allTestCases,
		CurrentUser: currentUser}

	// Add reference to TestCAses in command and rule engine
	commandAndRuleEngine.Testcases = &testCasesObject

	// Execute command
	err := commandAndRuleEngine.executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel(testCaseUuid, uuidToBeDeleted)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Extract TestCase
	testCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	assert.Equal(t, "true", fmt.Sprint(existsInMap))

	// Validate the result of the swap, 'B1f-TIC(B11f-TIC(B10)-B11lx)-B1l'
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

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_1.ParentElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 5) Validate TIC (2.2)
	testCaseModelElementUuid_2_2 := testCaseModelElement_2_1.NextElementUuid
	testCaseModelElement_2_2 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_2.ParentElementUuid &&
		testCaseModelElement_2_2.PreviousElementUuid == testCaseModelElement_2_1.MatureElementUuid &&
		testCaseModelElement_2_2.MatureElementUuid == testCaseModelElement_2_1.NextElementUuid &&
		testCaseModelElement_2_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 6) Validate B11l (2.3)
	testCaseModelElementUuid_2_3 := testCaseModelElement_2_2.NextElementUuid
	testCaseModelElement_2_3 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_3]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_3.ParentElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_2.NextElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_3.FirstChildElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_3.NextElementUuid &&
		testCaseModelElement_2_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 6) Validate B11l (2.2.1)
	testCaseModelElementUuid_2_2_1 := testCaseModelElement_2_2.FirstChildElementUuid
	testCaseModelElement_2_2_1 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_2_1]

	correctElement = testCaseModelElement_2_2.MatureElementUuid == testCaseModelElement_2_2_1.ParentElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.PreviousElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.NextElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = commandAndRuleEngine.Testcases.VerifyThatThereAreNoZombieElementsInTestCaseModel(testCaseUuid)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate Textual TestCase Presentation
	textualTestCaseSimple, textualTestCaseComplex, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)

	textualTestCaseRepresentationSimple := "[B1-TIC(B11-TIC(B10)-B11x)-B1]"
	textualTestCaseRepresentationComplex := "[B1f-TIC(B11f-TIC(B10)-B11lx)-B1l]"

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

// TCRuleDeletion112
// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-B12-n-B11lx					X-B11lx						TCRuleDeletion112
func TestTCRuleDeletion112(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := CommandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		Testcases:         nil,
	}

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
		OriginalElementName:      "B1f_BOND",
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
		OriginalElementName:      "B11f_BOND",
		MatureElementUuid:        "beea7876-f566-44c7-9625-655d8d075c5a",
		PreviousElementUuid:      "beea7876-f566-44c7-9625-655d8d075c5a",
		NextElementUuid:          "1780b088-ac62-4d13-ab55-b370017690ba",
		FirstChildElementUuid:    "beea7876-f566-44c7-9625-655d8d075c5a",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
	}

	tc_2_2_tic := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		MatureElementUuid:        "1780b088-ac62-4d13-ab55-b370017690ba",
		PreviousElementUuid:      "beea7876-f566-44c7-9625-655d8d075c5a",
		NextElementUuid:          "269a01da-7112-496a-82bf-da3cf30c03da",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_2_3_b12 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "6b9c5679-95f8-4e6b-87b6-e8083d47f66e",
		OriginalElementName:      "B12_BOND",
		MatureElementUuid:        "269a01da-7112-496a-82bf-da3cf30c03da",
		PreviousElementUuid:      "1780b088-ac62-4d13-ab55-b370017690ba",
		NextElementUuid:          uuidToBeDeleted,
		FirstChildElementUuid:    "269a01da-7112-496a-82bf-da3cf30c03da",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND,
	}

	tc_2_4_ti := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "f1a1294b-72a5-447d-af55-8d8ef746eec1",
		OriginalElementName:      "TI",
		MatureElementUuid:        uuidToBeDeleted,
		PreviousElementUuid:      "269a01da-7112-496a-82bf-da3cf30c03da",
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    uuidToBeDeleted,
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
	}

	tc_2_5_b11lx := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "66284917-6b5c-42a0-98d6-106a7bd269fa",
		OriginalElementName:      "B11lx_BOND",
		MatureElementUuid:        "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		PreviousElementUuid:      uuidToBeDeleted,
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE,
	}

	tc_2_2_1_b10 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "4f348a88-bb9a-4852-80de-3141687db11d",
		OriginalElementName:      "B10_BOND",
		MatureElementUuid:        "edaf757d-1205-4f2d-91a4-f053982f5ded",
		PreviousElementUuid:      "edaf757d-1205-4f2d-91a4-f053982f5ded",
		NextElementUuid:          "edaf757d-1205-4f2d-91a4-f053982f5ded",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "1780b088-ac62-4d13-ab55-b370017690ba",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
	}

	// Set up TestCaseModel-map
	myTestCaseModel.TestCaseModelMap[tc_1_b1f.MatureElementUuid] = tc_1_b1f
	myTestCaseModel.TestCaseModelMap[tc_2_tic.MatureElementUuid] = tc_2_tic
	myTestCaseModel.TestCaseModelMap[tc_3_b1l.MatureElementUuid] = tc_3_b1l
	myTestCaseModel.TestCaseModelMap[tc_2_1_b11f.MatureElementUuid] = tc_2_1_b11f
	myTestCaseModel.TestCaseModelMap[tc_2_2_tic.MatureElementUuid] = tc_2_2_tic
	myTestCaseModel.TestCaseModelMap[tc_2_3_b12.MatureElementUuid] = tc_2_3_b12
	myTestCaseModel.TestCaseModelMap[tc_2_4_ti.MatureElementUuid] = tc_2_4_ti
	myTestCaseModel.TestCaseModelMap[tc_2_5_b11lx.MatureElementUuid] = tc_2_5_b11lx
	myTestCaseModel.TestCaseModelMap[tc_2_2_1_b10.MatureElementUuid] = tc_2_2_1_b10

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_1_b1f.MatureElementUuid

	// Initiate structure for all TestCases
	allTestCases := make(map[string]testCaseModel.TestCaseModelStruct)

	// Generate new UUID
	testCaseUuid := uuidGenerator.New().String()

	// Add myTestCaseModel to map of all Testcases
	allTestCases[testCaseUuid] = myTestCaseModel

	// Set Current User
	currentUser := "s41797"

	// Add AddTestCases to TestCases-model
	testCasesObject := testCaseModel.TestCasesModelsStruct{
		TestCases:   allTestCases,
		CurrentUser: currentUser}

	// Add reference to TestCAses in command and rule engine
	commandAndRuleEngine.Testcases = &testCasesObject

	// Execute command
	err := commandAndRuleEngine.executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel(testCaseUuid, uuidToBeDeleted)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Extract TestCase
	testCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	assert.Equal(t, "true", fmt.Sprint(existsInMap))

	// Validate the result of the swap, 'B1f-TIC(B11f-TIC(B10)-B11lx)-B1l'
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

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_1.ParentElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 5) Validate TIC (2.2)
	testCaseModelElementUuid_2_2 := testCaseModelElement_2_1.NextElementUuid
	testCaseModelElement_2_2 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_2.ParentElementUuid &&
		testCaseModelElement_2_2.PreviousElementUuid == testCaseModelElement_2_1.MatureElementUuid &&
		testCaseModelElement_2_2.MatureElementUuid == testCaseModelElement_2_1.NextElementUuid &&
		testCaseModelElement_2_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 6) Validate B11l (2.3)
	testCaseModelElementUuid_2_3 := testCaseModelElement_2_2.NextElementUuid
	testCaseModelElement_2_3 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_3]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_3.ParentElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_2.NextElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_3.FirstChildElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_3.NextElementUuid &&
		testCaseModelElement_2_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 6) Validate B11l (2.2.1)
	testCaseModelElementUuid_2_2_1 := testCaseModelElement_2_2.FirstChildElementUuid
	testCaseModelElement_2_2_1 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_2_1]

	correctElement = testCaseModelElement_2_2.MatureElementUuid == testCaseModelElement_2_2_1.ParentElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.PreviousElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.NextElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = commandAndRuleEngine.Testcases.VerifyThatThereAreNoZombieElementsInTestCaseModel(testCaseUuid)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate Textual TestCase Presentation
	textualTestCaseSimple, textualTestCaseComplex, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)

	textualTestCaseRepresentationSimple := "[B1-TIC(B11-TIC(B10)-B11x)-B1]"
	textualTestCaseRepresentationComplex := "[B1f-TIC(B11f-TIC(B10)-B11lx)-B1l]"

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

// TCRuleDeletion113
// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-B12x-n-B11l					X-B11lx						TCRuleDeletion113
func TestTCRuleDeletion113(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := CommandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		Testcases:         nil,
	}

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
		OriginalElementName:      "B1f_BOND",
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
		OriginalElementName:      "B11f_BOND",
		MatureElementUuid:        "beea7876-f566-44c7-9625-655d8d075c5a",
		PreviousElementUuid:      "beea7876-f566-44c7-9625-655d8d075c5a",
		NextElementUuid:          "1780b088-ac62-4d13-ab55-b370017690ba",
		FirstChildElementUuid:    "beea7876-f566-44c7-9625-655d8d075c5a",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
	}

	tc_2_2_tic := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		MatureElementUuid:        "1780b088-ac62-4d13-ab55-b370017690ba",
		PreviousElementUuid:      "beea7876-f566-44c7-9625-655d8d075c5a",
		NextElementUuid:          "269a01da-7112-496a-82bf-da3cf30c03da",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_2_3_b12x := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "6b9c5679-95f8-4e6b-87b6-e8083d47f66e",
		OriginalElementName:      "B12x_BOND",
		MatureElementUuid:        "269a01da-7112-496a-82bf-da3cf30c03da",
		PreviousElementUuid:      "1780b088-ac62-4d13-ab55-b370017690ba",
		NextElementUuid:          uuidToBeDeleted,
		FirstChildElementUuid:    "269a01da-7112-496a-82bf-da3cf30c03da",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE,
	}

	tc_2_4_ti := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "f1a1294b-72a5-447d-af55-8d8ef746eec1",
		OriginalElementName:      "TI",
		MatureElementUuid:        uuidToBeDeleted,
		PreviousElementUuid:      "269a01da-7112-496a-82bf-da3cf30c03da",
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    uuidToBeDeleted,
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
	}

	tc_2_5_b11l := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "66284917-6b5c-42a0-98d6-106a7bd269fa",
		OriginalElementName:      "B11l_BOND",
		MatureElementUuid:        "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		PreviousElementUuid:      uuidToBeDeleted,
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND,
	}

	tc_2_2_1_b10 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "4f348a88-bb9a-4852-80de-3141687db11d",
		OriginalElementName:      "B10_BOND",
		MatureElementUuid:        "edaf757d-1205-4f2d-91a4-f053982f5ded",
		PreviousElementUuid:      "edaf757d-1205-4f2d-91a4-f053982f5ded",
		NextElementUuid:          "edaf757d-1205-4f2d-91a4-f053982f5ded",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "1780b088-ac62-4d13-ab55-b370017690ba",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
	}

	// Set up TestCaseModel-map
	myTestCaseModel.TestCaseModelMap[tc_1_b1f.MatureElementUuid] = tc_1_b1f
	myTestCaseModel.TestCaseModelMap[tc_2_tic.MatureElementUuid] = tc_2_tic
	myTestCaseModel.TestCaseModelMap[tc_3_b1l.MatureElementUuid] = tc_3_b1l
	myTestCaseModel.TestCaseModelMap[tc_2_1_b11f.MatureElementUuid] = tc_2_1_b11f
	myTestCaseModel.TestCaseModelMap[tc_2_2_tic.MatureElementUuid] = tc_2_2_tic
	myTestCaseModel.TestCaseModelMap[tc_2_3_b12x.MatureElementUuid] = tc_2_3_b12x
	myTestCaseModel.TestCaseModelMap[tc_2_4_ti.MatureElementUuid] = tc_2_4_ti
	myTestCaseModel.TestCaseModelMap[tc_2_5_b11l.MatureElementUuid] = tc_2_5_b11l
	myTestCaseModel.TestCaseModelMap[tc_2_2_1_b10.MatureElementUuid] = tc_2_2_1_b10

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_1_b1f.MatureElementUuid

	// Initiate structure for all TestCases
	allTestCases := make(map[string]testCaseModel.TestCaseModelStruct)

	// Generate new UUID
	testCaseUuid := uuidGenerator.New().String()

	// Add myTestCaseModel to map of all Testcases
	allTestCases[testCaseUuid] = myTestCaseModel

	// Set Current User
	currentUser := "s41797"

	// Add AddTestCases to TestCases-model
	testCasesObject := testCaseModel.TestCasesModelsStruct{
		TestCases:   allTestCases,
		CurrentUser: currentUser}

	// Add reference to TestCAses in command and rule engine
	commandAndRuleEngine.Testcases = &testCasesObject

	// Execute command
	err := commandAndRuleEngine.executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel(testCaseUuid, uuidToBeDeleted)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Extract TestCase
	testCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	assert.Equal(t, "true", fmt.Sprint(existsInMap))

	// Validate the result of the swap, 'B1f-TIC(B11f-TIC(B10)-B11lx)-B1l'
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

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_1.ParentElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 5) Validate TIC (2.2)
	testCaseModelElementUuid_2_2 := testCaseModelElement_2_1.NextElementUuid
	testCaseModelElement_2_2 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_2.ParentElementUuid &&
		testCaseModelElement_2_2.PreviousElementUuid == testCaseModelElement_2_1.MatureElementUuid &&
		testCaseModelElement_2_2.MatureElementUuid == testCaseModelElement_2_1.NextElementUuid &&
		testCaseModelElement_2_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 6) Validate B11l (2.3)
	testCaseModelElementUuid_2_3 := testCaseModelElement_2_2.NextElementUuid
	testCaseModelElement_2_3 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_3]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_3.ParentElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_2.NextElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_3.FirstChildElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_3.NextElementUuid &&
		testCaseModelElement_2_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 6) Validate B11l (2.2.1)
	testCaseModelElementUuid_2_2_1 := testCaseModelElement_2_2.FirstChildElementUuid
	testCaseModelElement_2_2_1 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_2_1]

	correctElement = testCaseModelElement_2_2.MatureElementUuid == testCaseModelElement_2_2_1.ParentElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.PreviousElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.NextElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = commandAndRuleEngine.Testcases.VerifyThatThereAreNoZombieElementsInTestCaseModel(testCaseUuid)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate Textual TestCase Presentation
	textualTestCaseSimple, textualTestCaseComplex, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)

	textualTestCaseRepresentationSimple := "[B1-TIC(B11-TIC(B10)-B11x)-B1]"
	textualTestCaseRepresentationComplex := "[B1f-TIC(B11f-TIC(B10)-B11lx)-B1l]"

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

// TCRuleDeletion114
// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-B12-n-B12-X					X-B12-X						TCRuleDeletion114
func TestTCRuleDeletion114(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := CommandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		Testcases:         nil,
	}

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
		OriginalElementName:      "B1f_BOND",
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
		OriginalElementName:      "B11f_BOND",
		MatureElementUuid:        "beea7876-f566-44c7-9625-655d8d075c5a",
		PreviousElementUuid:      "beea7876-f566-44c7-9625-655d8d075c5a",
		NextElementUuid:          "1780b088-ac62-4d13-ab55-b370017690ba",
		FirstChildElementUuid:    "beea7876-f566-44c7-9625-655d8d075c5a",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
	}

	tc_2_2_tic := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		MatureElementUuid:        "1780b088-ac62-4d13-ab55-b370017690ba",
		PreviousElementUuid:      "beea7876-f566-44c7-9625-655d8d075c5a",
		NextElementUuid:          "269a01da-7112-496a-82bf-da3cf30c03da",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_2_3_b12 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "6b9c5679-95f8-4e6b-87b6-e8083d47f66e",
		OriginalElementName:      "B12_BOND",
		MatureElementUuid:        "269a01da-7112-496a-82bf-da3cf30c03da",
		PreviousElementUuid:      "1780b088-ac62-4d13-ab55-b370017690ba",
		NextElementUuid:          uuidToBeDeleted,
		FirstChildElementUuid:    "269a01da-7112-496a-82bf-da3cf30c03da",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND,
	}

	tc_2_4_ti := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "f1a1294b-72a5-447d-af55-8d8ef746eec1",
		OriginalElementName:      "TI",
		MatureElementUuid:        uuidToBeDeleted,
		PreviousElementUuid:      "269a01da-7112-496a-82bf-da3cf30c03da",
		NextElementUuid:          "b022365c-1107-40ec-9c95-60d6fd73745f",
		FirstChildElementUuid:    uuidToBeDeleted,
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
	}

	tc_2_5_b12 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "6b9c5679-95f8-4e6b-87b6-e8083d47f66e",
		OriginalElementName:      "B12_BOND",
		MatureElementUuid:        "b022365c-1107-40ec-9c95-60d6fd73745f",
		PreviousElementUuid:      uuidToBeDeleted,
		NextElementUuid:          "99ec6586-19cb-4988-9444-d5518a84f632",
		FirstChildElementUuid:    "b022365c-1107-40ec-9c95-60d6fd73745f",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND,
	}

	tc_2_6_ti := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "f1a1294b-72a5-447d-af55-8d8ef746eec1",
		OriginalElementName:      "TI",
		MatureElementUuid:        "99ec6586-19cb-4988-9444-d5518a84f632",
		PreviousElementUuid:      "b022365c-1107-40ec-9c95-60d6fd73745f",
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "99ec6586-19cb-4988-9444-d5518a84f632",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
	}

	tc_2_7_b11l := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "66284917-6b5c-42a0-98d6-106a7bd269fa",
		OriginalElementName:      "B11l_BOND",
		MatureElementUuid:        "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		PreviousElementUuid:      "99ec6586-19cb-4988-9444-d5518a84f632",
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND,
	}

	tc_2_2_1_b10 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "4f348a88-bb9a-4852-80de-3141687db11d",
		OriginalElementName:      "B10_BOND",
		MatureElementUuid:        "edaf757d-1205-4f2d-91a4-f053982f5ded",
		PreviousElementUuid:      "edaf757d-1205-4f2d-91a4-f053982f5ded",
		NextElementUuid:          "edaf757d-1205-4f2d-91a4-f053982f5ded",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "1780b088-ac62-4d13-ab55-b370017690ba",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
	}

	// Set up TestCaseModel-map
	myTestCaseModel.TestCaseModelMap[tc_1_b1f.MatureElementUuid] = tc_1_b1f
	myTestCaseModel.TestCaseModelMap[tc_2_tic.MatureElementUuid] = tc_2_tic
	myTestCaseModel.TestCaseModelMap[tc_3_b1l.MatureElementUuid] = tc_3_b1l
	myTestCaseModel.TestCaseModelMap[tc_2_1_b11f.MatureElementUuid] = tc_2_1_b11f
	myTestCaseModel.TestCaseModelMap[tc_2_2_tic.MatureElementUuid] = tc_2_2_tic
	myTestCaseModel.TestCaseModelMap[tc_2_3_b12.MatureElementUuid] = tc_2_3_b12
	myTestCaseModel.TestCaseModelMap[tc_2_4_ti.MatureElementUuid] = tc_2_4_ti
	myTestCaseModel.TestCaseModelMap[tc_2_5_b12.MatureElementUuid] = tc_2_5_b12
	myTestCaseModel.TestCaseModelMap[tc_2_6_ti.MatureElementUuid] = tc_2_6_ti
	myTestCaseModel.TestCaseModelMap[tc_2_7_b11l.MatureElementUuid] = tc_2_7_b11l
	myTestCaseModel.TestCaseModelMap[tc_2_2_1_b10.MatureElementUuid] = tc_2_2_1_b10

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_1_b1f.MatureElementUuid

	// Initiate structure for all TestCases
	allTestCases := make(map[string]testCaseModel.TestCaseModelStruct)

	// Generate new UUID
	testCaseUuid := uuidGenerator.New().String()

	// Add myTestCaseModel to map of all Testcases
	allTestCases[testCaseUuid] = myTestCaseModel

	// Set Current User
	currentUser := "s41797"

	// Add AddTestCases to TestCases-model
	testCasesObject := testCaseModel.TestCasesModelsStruct{
		TestCases:   allTestCases,
		CurrentUser: currentUser}

	// Add reference to TestCAses in command and rule engine
	commandAndRuleEngine.Testcases = &testCasesObject

	// Execute command
	err := commandAndRuleEngine.executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel(testCaseUuid, uuidToBeDeleted)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Extract TestCase
	testCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	assert.Equal(t, "true", fmt.Sprint(existsInMap))

	// Validate the result of the swap, 'B1f-TIC(B11f-TIC(B10)-B12-TI-B11l)-B1l'
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

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_1.ParentElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 5) Validate TIC (2.2)
	testCaseModelElementUuid_2_2 := testCaseModelElement_2_1.NextElementUuid
	testCaseModelElement_2_2 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_2.ParentElementUuid &&
		testCaseModelElement_2_2.PreviousElementUuid == testCaseModelElement_2_1.MatureElementUuid &&
		testCaseModelElement_2_2.MatureElementUuid == testCaseModelElement_2_1.NextElementUuid &&
		testCaseModelElement_2_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 6) Validate B12 (2.3)
	testCaseModelElementUuid_2_3 := testCaseModelElement_2_2.NextElementUuid
	testCaseModelElement_2_3 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_3]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_3.ParentElementUuid &&
		testCaseModelElement_2_3.PreviousElementUuid == testCaseModelElement_2_2.MatureElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_2.NextElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_3.FirstChildElementUuid &&
		testCaseModelElement_2_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 7) Validate TI (2.4)
	testCaseModelElementUuid_2_4 := testCaseModelElement_2_3.NextElementUuid
	testCaseModelElement_2_4 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_4]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_4.ParentElementUuid &&
		testCaseModelElement_2_4.PreviousElementUuid == testCaseModelElement_2_3.MatureElementUuid &&
		testCaseModelElement_2_4.MatureElementUuid == testCaseModelElement_2_3.NextElementUuid &&
		testCaseModelElement_2_4.MatureElementUuid == testCaseModelElement_2_4.FirstChildElementUuid &&
		testCaseModelElement_2_4.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 7) Validate B11l (2.5)
	testCaseModelElementUuid_2_5 := testCaseModelElement_2_4.NextElementUuid
	testCaseModelElement_2_5 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_5]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_5.ParentElementUuid &&
		testCaseModelElement_2_5.MatureElementUuid == testCaseModelElement_2_4.NextElementUuid &&
		testCaseModelElement_2_5.MatureElementUuid == testCaseModelElement_2_5.FirstChildElementUuid &&
		testCaseModelElement_2_5.MatureElementUuid == testCaseModelElement_2_5.NextElementUuid &&
		testCaseModelElement_2_5.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 8) Validate B11l (2.2.1)
	testCaseModelElementUuid_2_2_1 := testCaseModelElement_2_2.FirstChildElementUuid
	testCaseModelElement_2_2_1 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_2_1]

	correctElement = testCaseModelElement_2_2.MatureElementUuid == testCaseModelElement_2_2_1.ParentElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.PreviousElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.NextElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = commandAndRuleEngine.Testcases.VerifyThatThereAreNoZombieElementsInTestCaseModel(testCaseUuid)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate Textual TestCase Presentation
	textualTestCaseSimple, textualTestCaseComplex, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)

	textualTestCaseRepresentationSimple := "[B1-TIC(B11-TIC(B10)-B12-TI-B11)-B1]"
	textualTestCaseRepresentationComplex := "[B1f-TIC(B11f-TIC(B10)-B12-TI-B11l)-B1l]"

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

// TCRuleDeletion115
// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-B12x-n-B12x-X					X-B12x-X					TCRuleDeletion115
func TestTCRuleDeletion115(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := CommandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		Testcases:         nil,
	}

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
		OriginalElementName:      "B1f_BOND",
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
		OriginalElementName:      "B11f_BOND",
		MatureElementUuid:        "beea7876-f566-44c7-9625-655d8d075c5a",
		PreviousElementUuid:      "beea7876-f566-44c7-9625-655d8d075c5a",
		NextElementUuid:          "1780b088-ac62-4d13-ab55-b370017690ba",
		FirstChildElementUuid:    "beea7876-f566-44c7-9625-655d8d075c5a",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
	}

	tc_2_2_tic := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		MatureElementUuid:        "1780b088-ac62-4d13-ab55-b370017690ba",
		PreviousElementUuid:      "beea7876-f566-44c7-9625-655d8d075c5a",
		NextElementUuid:          "269a01da-7112-496a-82bf-da3cf30c03da",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_2_3_b12x := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "6b9c5679-95f8-4e6b-87b6-e8083d47f66e",
		OriginalElementName:      "B12x_BOND",
		MatureElementUuid:        "269a01da-7112-496a-82bf-da3cf30c03da",
		PreviousElementUuid:      "1780b088-ac62-4d13-ab55-b370017690ba",
		NextElementUuid:          uuidToBeDeleted,
		FirstChildElementUuid:    "269a01da-7112-496a-82bf-da3cf30c03da",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE,
	}

	tc_2_4_ti := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "f1a1294b-72a5-447d-af55-8d8ef746eec1",
		OriginalElementName:      "TI",
		MatureElementUuid:        uuidToBeDeleted,
		PreviousElementUuid:      "269a01da-7112-496a-82bf-da3cf30c03da",
		NextElementUuid:          "b022365c-1107-40ec-9c95-60d6fd73745f",
		FirstChildElementUuid:    uuidToBeDeleted,
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
	}

	tc_2_5_b12x := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "6b9c5679-95f8-4e6b-87b6-e8083d47f66e",
		OriginalElementName:      "B12x_BOND",
		MatureElementUuid:        "b022365c-1107-40ec-9c95-60d6fd73745f",
		PreviousElementUuid:      uuidToBeDeleted,
		NextElementUuid:          "99ec6586-19cb-4988-9444-d5518a84f632",
		FirstChildElementUuid:    "b022365c-1107-40ec-9c95-60d6fd73745f",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE,
	}

	tc_2_6_ti := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "f1a1294b-72a5-447d-af55-8d8ef746eec1",
		OriginalElementName:      "TI",
		MatureElementUuid:        "99ec6586-19cb-4988-9444-d5518a84f632",
		PreviousElementUuid:      "b022365c-1107-40ec-9c95-60d6fd73745f",
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "99ec6586-19cb-4988-9444-d5518a84f632",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
	}

	tc_2_7_b11l := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "66284917-6b5c-42a0-98d6-106a7bd269fa",
		OriginalElementName:      "B11l_BOND",
		MatureElementUuid:        "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		PreviousElementUuid:      "99ec6586-19cb-4988-9444-d5518a84f632",
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND,
	}

	tc_2_2_1_b10 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "4f348a88-bb9a-4852-80de-3141687db11d",
		OriginalElementName:      "B10_BOND",
		MatureElementUuid:        "edaf757d-1205-4f2d-91a4-f053982f5ded",
		PreviousElementUuid:      "edaf757d-1205-4f2d-91a4-f053982f5ded",
		NextElementUuid:          "edaf757d-1205-4f2d-91a4-f053982f5ded",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "1780b088-ac62-4d13-ab55-b370017690ba",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
	}

	// Set up TestCaseModel-map
	myTestCaseModel.TestCaseModelMap[tc_1_b1f.MatureElementUuid] = tc_1_b1f
	myTestCaseModel.TestCaseModelMap[tc_2_tic.MatureElementUuid] = tc_2_tic
	myTestCaseModel.TestCaseModelMap[tc_3_b1l.MatureElementUuid] = tc_3_b1l
	myTestCaseModel.TestCaseModelMap[tc_2_1_b11f.MatureElementUuid] = tc_2_1_b11f
	myTestCaseModel.TestCaseModelMap[tc_2_2_tic.MatureElementUuid] = tc_2_2_tic
	myTestCaseModel.TestCaseModelMap[tc_2_3_b12x.MatureElementUuid] = tc_2_3_b12x
	myTestCaseModel.TestCaseModelMap[tc_2_4_ti.MatureElementUuid] = tc_2_4_ti
	myTestCaseModel.TestCaseModelMap[tc_2_5_b12x.MatureElementUuid] = tc_2_5_b12x
	myTestCaseModel.TestCaseModelMap[tc_2_6_ti.MatureElementUuid] = tc_2_6_ti
	myTestCaseModel.TestCaseModelMap[tc_2_7_b11l.MatureElementUuid] = tc_2_7_b11l
	myTestCaseModel.TestCaseModelMap[tc_2_2_1_b10.MatureElementUuid] = tc_2_2_1_b10

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_1_b1f.MatureElementUuid

	// Initiate structure for all TestCases
	allTestCases := make(map[string]testCaseModel.TestCaseModelStruct)

	// Generate new UUID
	testCaseUuid := uuidGenerator.New().String()

	// Add myTestCaseModel to map of all Testcases
	allTestCases[testCaseUuid] = myTestCaseModel

	// Set Current User
	currentUser := "s41797"

	// Add AddTestCases to TestCases-model
	testCasesObject := testCaseModel.TestCasesModelsStruct{
		TestCases:   allTestCases,
		CurrentUser: currentUser}

	// Add reference to TestCAses in command and rule engine
	commandAndRuleEngine.Testcases = &testCasesObject

	// Execute command
	err := commandAndRuleEngine.executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel(testCaseUuid, uuidToBeDeleted)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Extract TestCase
	testCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	assert.Equal(t, "true", fmt.Sprint(existsInMap))

	// Validate the result of the swap, 'B1f-TIC(B11f-TIC(B10)-B12x-TI-B11l)-B1l'
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

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_1.ParentElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 5) Validate TIC (2.2)
	testCaseModelElementUuid_2_2 := testCaseModelElement_2_1.NextElementUuid
	testCaseModelElement_2_2 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_2.ParentElementUuid &&
		testCaseModelElement_2_2.PreviousElementUuid == testCaseModelElement_2_1.MatureElementUuid &&
		testCaseModelElement_2_2.MatureElementUuid == testCaseModelElement_2_1.NextElementUuid &&
		testCaseModelElement_2_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 6) Validate B12x (2.3)
	testCaseModelElementUuid_2_3 := testCaseModelElement_2_2.NextElementUuid
	testCaseModelElement_2_3 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_3]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_3.ParentElementUuid &&
		testCaseModelElement_2_3.PreviousElementUuid == testCaseModelElement_2_2.MatureElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_2.NextElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_3.FirstChildElementUuid &&
		testCaseModelElement_2_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 7) Validate TI (2.4)
	testCaseModelElementUuid_2_4 := testCaseModelElement_2_3.NextElementUuid
	testCaseModelElement_2_4 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_4]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_4.ParentElementUuid &&
		testCaseModelElement_2_4.PreviousElementUuid == testCaseModelElement_2_3.MatureElementUuid &&
		testCaseModelElement_2_4.MatureElementUuid == testCaseModelElement_2_3.NextElementUuid &&
		testCaseModelElement_2_4.MatureElementUuid == testCaseModelElement_2_4.FirstChildElementUuid &&
		testCaseModelElement_2_4.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 7) Validate B11l (2.5)
	testCaseModelElementUuid_2_5 := testCaseModelElement_2_4.NextElementUuid
	testCaseModelElement_2_5 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_5]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_5.ParentElementUuid &&
		testCaseModelElement_2_5.MatureElementUuid == testCaseModelElement_2_4.NextElementUuid &&
		testCaseModelElement_2_5.MatureElementUuid == testCaseModelElement_2_5.FirstChildElementUuid &&
		testCaseModelElement_2_5.MatureElementUuid == testCaseModelElement_2_5.NextElementUuid &&
		testCaseModelElement_2_5.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 8) Validate B11l (2.2.1)
	testCaseModelElementUuid_2_2_1 := testCaseModelElement_2_2.FirstChildElementUuid
	testCaseModelElement_2_2_1 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_2_1]

	correctElement = testCaseModelElement_2_2.MatureElementUuid == testCaseModelElement_2_2_1.ParentElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.PreviousElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.NextElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = commandAndRuleEngine.Testcases.VerifyThatThereAreNoZombieElementsInTestCaseModel(testCaseUuid)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate Textual TestCase Presentation
	textualTestCaseSimple, textualTestCaseComplex, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)

	textualTestCaseRepresentationSimple := "[B1-TIC(B11-TIC(B10)-B12x-TI-B11)-B1]"
	textualTestCaseRepresentationComplex := "[B1f-TIC(B11f-TIC(B10)-B12x-TI-B11l)-B1l]"

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

// TCRuleDeletion116
// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-B12-n-B12x-X					X-B12x-X					TCRuleDeletion116
func TestTCRuleDeletion116(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := CommandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		Testcases:         nil,
	}

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
		OriginalElementName:      "B1f_BOND",
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
		OriginalElementName:      "B11f_BOND",
		MatureElementUuid:        "beea7876-f566-44c7-9625-655d8d075c5a",
		PreviousElementUuid:      "beea7876-f566-44c7-9625-655d8d075c5a",
		NextElementUuid:          "1780b088-ac62-4d13-ab55-b370017690ba",
		FirstChildElementUuid:    "beea7876-f566-44c7-9625-655d8d075c5a",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
	}

	tc_2_2_tic := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		MatureElementUuid:        "1780b088-ac62-4d13-ab55-b370017690ba",
		PreviousElementUuid:      "beea7876-f566-44c7-9625-655d8d075c5a",
		NextElementUuid:          "269a01da-7112-496a-82bf-da3cf30c03da",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_2_3_b12 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "6b9c5679-95f8-4e6b-87b6-e8083d47f66e",
		OriginalElementName:      "B12_BOND",
		MatureElementUuid:        "269a01da-7112-496a-82bf-da3cf30c03da",
		PreviousElementUuid:      "1780b088-ac62-4d13-ab55-b370017690ba",
		NextElementUuid:          uuidToBeDeleted,
		FirstChildElementUuid:    "269a01da-7112-496a-82bf-da3cf30c03da",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND,
	}

	tc_2_4_ti := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "f1a1294b-72a5-447d-af55-8d8ef746eec1",
		OriginalElementName:      "TI",
		MatureElementUuid:        uuidToBeDeleted,
		PreviousElementUuid:      "269a01da-7112-496a-82bf-da3cf30c03da",
		NextElementUuid:          "b022365c-1107-40ec-9c95-60d6fd73745f",
		FirstChildElementUuid:    uuidToBeDeleted,
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
	}

	tc_2_5_b12x := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "6b9c5679-95f8-4e6b-87b6-e8083d47f66e",
		OriginalElementName:      "B12x_BOND",
		MatureElementUuid:        "b022365c-1107-40ec-9c95-60d6fd73745f",
		PreviousElementUuid:      uuidToBeDeleted,
		NextElementUuid:          "99ec6586-19cb-4988-9444-d5518a84f632",
		FirstChildElementUuid:    "b022365c-1107-40ec-9c95-60d6fd73745f",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE,
	}

	tc_2_6_ti := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "f1a1294b-72a5-447d-af55-8d8ef746eec1",
		OriginalElementName:      "TI",
		MatureElementUuid:        "99ec6586-19cb-4988-9444-d5518a84f632",
		PreviousElementUuid:      "b022365c-1107-40ec-9c95-60d6fd73745f",
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "99ec6586-19cb-4988-9444-d5518a84f632",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
	}

	tc_2_7_b11l := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "66284917-6b5c-42a0-98d6-106a7bd269fa",
		OriginalElementName:      "B11l_BOND",
		MatureElementUuid:        "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		PreviousElementUuid:      "99ec6586-19cb-4988-9444-d5518a84f632",
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND,
	}

	tc_2_2_1_b10 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "4f348a88-bb9a-4852-80de-3141687db11d",
		OriginalElementName:      "B10_BOND",
		MatureElementUuid:        "edaf757d-1205-4f2d-91a4-f053982f5ded",
		PreviousElementUuid:      "edaf757d-1205-4f2d-91a4-f053982f5ded",
		NextElementUuid:          "edaf757d-1205-4f2d-91a4-f053982f5ded",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "1780b088-ac62-4d13-ab55-b370017690ba",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
	}

	// Set up TestCaseModel-map
	myTestCaseModel.TestCaseModelMap[tc_1_b1f.MatureElementUuid] = tc_1_b1f
	myTestCaseModel.TestCaseModelMap[tc_2_tic.MatureElementUuid] = tc_2_tic
	myTestCaseModel.TestCaseModelMap[tc_3_b1l.MatureElementUuid] = tc_3_b1l
	myTestCaseModel.TestCaseModelMap[tc_2_1_b11f.MatureElementUuid] = tc_2_1_b11f
	myTestCaseModel.TestCaseModelMap[tc_2_2_tic.MatureElementUuid] = tc_2_2_tic
	myTestCaseModel.TestCaseModelMap[tc_2_3_b12.MatureElementUuid] = tc_2_3_b12
	myTestCaseModel.TestCaseModelMap[tc_2_4_ti.MatureElementUuid] = tc_2_4_ti
	myTestCaseModel.TestCaseModelMap[tc_2_5_b12x.MatureElementUuid] = tc_2_5_b12x
	myTestCaseModel.TestCaseModelMap[tc_2_6_ti.MatureElementUuid] = tc_2_6_ti
	myTestCaseModel.TestCaseModelMap[tc_2_7_b11l.MatureElementUuid] = tc_2_7_b11l
	myTestCaseModel.TestCaseModelMap[tc_2_2_1_b10.MatureElementUuid] = tc_2_2_1_b10

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_1_b1f.MatureElementUuid

	// Initiate structure for all TestCases
	allTestCases := make(map[string]testCaseModel.TestCaseModelStruct)

	// Generate new UUID
	testCaseUuid := uuidGenerator.New().String()

	// Add myTestCaseModel to map of all Testcases
	allTestCases[testCaseUuid] = myTestCaseModel

	// Set Current User
	currentUser := "s41797"

	// Add AddTestCases to TestCases-model
	testCasesObject := testCaseModel.TestCasesModelsStruct{
		TestCases:   allTestCases,
		CurrentUser: currentUser}

	// Add reference to TestCAses in command and rule engine
	commandAndRuleEngine.Testcases = &testCasesObject

	// Execute command
	err := commandAndRuleEngine.executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel(testCaseUuid, uuidToBeDeleted)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Extract TestCase
	testCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	assert.Equal(t, "true", fmt.Sprint(existsInMap))

	// Validate the result of the swap, 'B1f-TIC(B11f-TIC(B10)-B12x-TI-B11l)-B1l'
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

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_1.ParentElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 5) Validate TIC (2.2)
	testCaseModelElementUuid_2_2 := testCaseModelElement_2_1.NextElementUuid
	testCaseModelElement_2_2 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_2.ParentElementUuid &&
		testCaseModelElement_2_2.PreviousElementUuid == testCaseModelElement_2_1.MatureElementUuid &&
		testCaseModelElement_2_2.MatureElementUuid == testCaseModelElement_2_1.NextElementUuid &&
		testCaseModelElement_2_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 6) Validate B12x (2.3)
	testCaseModelElementUuid_2_3 := testCaseModelElement_2_2.NextElementUuid
	testCaseModelElement_2_3 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_3]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_3.ParentElementUuid &&
		testCaseModelElement_2_3.PreviousElementUuid == testCaseModelElement_2_2.MatureElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_2.NextElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_3.FirstChildElementUuid &&
		testCaseModelElement_2_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 7) Validate TI (2.4)
	testCaseModelElementUuid_2_4 := testCaseModelElement_2_3.NextElementUuid
	testCaseModelElement_2_4 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_4]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_4.ParentElementUuid &&
		testCaseModelElement_2_4.PreviousElementUuid == testCaseModelElement_2_3.MatureElementUuid &&
		testCaseModelElement_2_4.MatureElementUuid == testCaseModelElement_2_3.NextElementUuid &&
		testCaseModelElement_2_4.MatureElementUuid == testCaseModelElement_2_4.FirstChildElementUuid &&
		testCaseModelElement_2_4.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 7) Validate B11l (2.5)
	testCaseModelElementUuid_2_5 := testCaseModelElement_2_4.NextElementUuid
	testCaseModelElement_2_5 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_5]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_5.ParentElementUuid &&
		testCaseModelElement_2_5.MatureElementUuid == testCaseModelElement_2_4.NextElementUuid &&
		testCaseModelElement_2_5.MatureElementUuid == testCaseModelElement_2_5.FirstChildElementUuid &&
		testCaseModelElement_2_5.MatureElementUuid == testCaseModelElement_2_5.NextElementUuid &&
		testCaseModelElement_2_5.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 8) Validate B11l (2.2.1)
	testCaseModelElementUuid_2_2_1 := testCaseModelElement_2_2.FirstChildElementUuid
	testCaseModelElement_2_2_1 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_2_1]

	correctElement = testCaseModelElement_2_2.MatureElementUuid == testCaseModelElement_2_2_1.ParentElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.PreviousElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.NextElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = commandAndRuleEngine.Testcases.VerifyThatThereAreNoZombieElementsInTestCaseModel(testCaseUuid)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate Textual TestCase Presentation
	textualTestCaseSimple, textualTestCaseComplex, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)

	textualTestCaseRepresentationSimple := "[B1-TIC(B11-TIC(B10)-B12x-TI-B11)-B1]"
	textualTestCaseRepresentationComplex := "[B1f-TIC(B11f-TIC(B10)-B12x-TI-B11l)-B1l]"

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

// TCRuleDeletion117
// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-B12x-n-B12-X					X-B12x-X					TCRuleDeletion117
func TestTCRuleDeletion117(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := CommandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		Testcases:         nil,
	}

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
		OriginalElementName:      "B1f_BOND",
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
		OriginalElementName:      "B11f_BOND",
		MatureElementUuid:        "beea7876-f566-44c7-9625-655d8d075c5a",
		PreviousElementUuid:      "beea7876-f566-44c7-9625-655d8d075c5a",
		NextElementUuid:          "1780b088-ac62-4d13-ab55-b370017690ba",
		FirstChildElementUuid:    "beea7876-f566-44c7-9625-655d8d075c5a",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
	}

	tc_2_2_tic := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		MatureElementUuid:        "1780b088-ac62-4d13-ab55-b370017690ba",
		PreviousElementUuid:      "beea7876-f566-44c7-9625-655d8d075c5a",
		NextElementUuid:          "269a01da-7112-496a-82bf-da3cf30c03da",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_2_3_b12x := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "6b9c5679-95f8-4e6b-87b6-e8083d47f66e",
		OriginalElementName:      "B12x_BOND",
		MatureElementUuid:        "269a01da-7112-496a-82bf-da3cf30c03da",
		PreviousElementUuid:      "1780b088-ac62-4d13-ab55-b370017690ba",
		NextElementUuid:          uuidToBeDeleted,
		FirstChildElementUuid:    "269a01da-7112-496a-82bf-da3cf30c03da",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE,
	}

	tc_2_4_ti := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "f1a1294b-72a5-447d-af55-8d8ef746eec1",
		OriginalElementName:      "TI",
		MatureElementUuid:        uuidToBeDeleted,
		PreviousElementUuid:      "269a01da-7112-496a-82bf-da3cf30c03da",
		NextElementUuid:          "b022365c-1107-40ec-9c95-60d6fd73745f",
		FirstChildElementUuid:    uuidToBeDeleted,
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
	}

	tc_2_5_b12 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "6b9c5679-95f8-4e6b-87b6-e8083d47f66e",
		OriginalElementName:      "B12_BOND",
		MatureElementUuid:        "b022365c-1107-40ec-9c95-60d6fd73745f",
		PreviousElementUuid:      uuidToBeDeleted,
		NextElementUuid:          "99ec6586-19cb-4988-9444-d5518a84f632",
		FirstChildElementUuid:    "b022365c-1107-40ec-9c95-60d6fd73745f",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND,
	}

	tc_2_6_ti := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "f1a1294b-72a5-447d-af55-8d8ef746eec1",
		OriginalElementName:      "TI",
		MatureElementUuid:        "99ec6586-19cb-4988-9444-d5518a84f632",
		PreviousElementUuid:      "b022365c-1107-40ec-9c95-60d6fd73745f",
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "99ec6586-19cb-4988-9444-d5518a84f632",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
	}

	tc_2_7_b11l := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "66284917-6b5c-42a0-98d6-106a7bd269fa",
		OriginalElementName:      "B11l_BOND",
		MatureElementUuid:        "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		PreviousElementUuid:      "99ec6586-19cb-4988-9444-d5518a84f632",
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND,
	}

	tc_2_2_1_b10 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "4f348a88-bb9a-4852-80de-3141687db11d",
		OriginalElementName:      "B10_BOND",
		MatureElementUuid:        "edaf757d-1205-4f2d-91a4-f053982f5ded",
		PreviousElementUuid:      "edaf757d-1205-4f2d-91a4-f053982f5ded",
		NextElementUuid:          "edaf757d-1205-4f2d-91a4-f053982f5ded",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "1780b088-ac62-4d13-ab55-b370017690ba",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
	}

	// Set up TestCaseModel-map
	myTestCaseModel.TestCaseModelMap[tc_1_b1f.MatureElementUuid] = tc_1_b1f
	myTestCaseModel.TestCaseModelMap[tc_2_tic.MatureElementUuid] = tc_2_tic
	myTestCaseModel.TestCaseModelMap[tc_3_b1l.MatureElementUuid] = tc_3_b1l
	myTestCaseModel.TestCaseModelMap[tc_2_1_b11f.MatureElementUuid] = tc_2_1_b11f
	myTestCaseModel.TestCaseModelMap[tc_2_2_tic.MatureElementUuid] = tc_2_2_tic
	myTestCaseModel.TestCaseModelMap[tc_2_3_b12x.MatureElementUuid] = tc_2_3_b12x
	myTestCaseModel.TestCaseModelMap[tc_2_4_ti.MatureElementUuid] = tc_2_4_ti
	myTestCaseModel.TestCaseModelMap[tc_2_5_b12.MatureElementUuid] = tc_2_5_b12
	myTestCaseModel.TestCaseModelMap[tc_2_6_ti.MatureElementUuid] = tc_2_6_ti
	myTestCaseModel.TestCaseModelMap[tc_2_7_b11l.MatureElementUuid] = tc_2_7_b11l
	myTestCaseModel.TestCaseModelMap[tc_2_2_1_b10.MatureElementUuid] = tc_2_2_1_b10

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_1_b1f.MatureElementUuid

	// Initiate structure for all TestCases
	allTestCases := make(map[string]testCaseModel.TestCaseModelStruct)

	// Generate new UUID
	testCaseUuid := uuidGenerator.New().String()

	// Add myTestCaseModel to map of all Testcases
	allTestCases[testCaseUuid] = myTestCaseModel

	// Set Current User
	currentUser := "s41797"

	// Add AddTestCases to TestCases-model
	testCasesObject := testCaseModel.TestCasesModelsStruct{
		TestCases:   allTestCases,
		CurrentUser: currentUser}

	// Add reference to TestCAses in command and rule engine
	commandAndRuleEngine.Testcases = &testCasesObject

	// Execute command
	err := commandAndRuleEngine.executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel(testCaseUuid, uuidToBeDeleted)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Extract TestCase
	testCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	assert.Equal(t, "true", fmt.Sprint(existsInMap))

	// Validate the result of the swap, 'B1f-TIC(B11f-TIC(B10)-B12x-TI-B11l)-B1l'
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

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_1.ParentElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 5) Validate TIC (2.2)
	testCaseModelElementUuid_2_2 := testCaseModelElement_2_1.NextElementUuid
	testCaseModelElement_2_2 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_2.ParentElementUuid &&
		testCaseModelElement_2_2.PreviousElementUuid == testCaseModelElement_2_1.MatureElementUuid &&
		testCaseModelElement_2_2.MatureElementUuid == testCaseModelElement_2_1.NextElementUuid &&
		testCaseModelElement_2_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 6) Validate B12x (2.3)
	testCaseModelElementUuid_2_3 := testCaseModelElement_2_2.NextElementUuid
	testCaseModelElement_2_3 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_3]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_3.ParentElementUuid &&
		testCaseModelElement_2_3.PreviousElementUuid == testCaseModelElement_2_2.MatureElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_2.NextElementUuid &&
		testCaseModelElement_2_3.MatureElementUuid == testCaseModelElement_2_3.FirstChildElementUuid &&
		testCaseModelElement_2_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 7) Validate TI (2.4)
	testCaseModelElementUuid_2_4 := testCaseModelElement_2_3.NextElementUuid
	testCaseModelElement_2_4 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_4]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_4.ParentElementUuid &&
		testCaseModelElement_2_4.PreviousElementUuid == testCaseModelElement_2_3.MatureElementUuid &&
		testCaseModelElement_2_4.MatureElementUuid == testCaseModelElement_2_3.NextElementUuid &&
		testCaseModelElement_2_4.MatureElementUuid == testCaseModelElement_2_4.FirstChildElementUuid &&
		testCaseModelElement_2_4.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 7) Validate B11l (2.5)
	testCaseModelElementUuid_2_5 := testCaseModelElement_2_4.NextElementUuid
	testCaseModelElement_2_5 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_5]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_5.ParentElementUuid &&
		testCaseModelElement_2_5.MatureElementUuid == testCaseModelElement_2_4.NextElementUuid &&
		testCaseModelElement_2_5.MatureElementUuid == testCaseModelElement_2_5.FirstChildElementUuid &&
		testCaseModelElement_2_5.MatureElementUuid == testCaseModelElement_2_5.NextElementUuid &&
		testCaseModelElement_2_5.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 8) Validate B11l (2.2.1)
	testCaseModelElementUuid_2_2_1 := testCaseModelElement_2_2.FirstChildElementUuid
	testCaseModelElement_2_2_1 := testCase.TestCaseModelMap[testCaseModelElementUuid_2_2_1]

	correctElement = testCaseModelElement_2_2.MatureElementUuid == testCaseModelElement_2_2_1.ParentElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.PreviousElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.NextElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = commandAndRuleEngine.Testcases.VerifyThatThereAreNoZombieElementsInTestCaseModel(testCaseUuid)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate Textual TestCase Presentation
	textualTestCaseSimple, textualTestCaseComplex, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)

	textualTestCaseRepresentationSimple := "[B1-TIC(B11-TIC(B10)-B12x-TI-B11)-B1]"
	textualTestCaseRepresentationComplex := "[B1f-TIC(B11f-TIC(B10)-B12x-TI-B11l)-B1l]"

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
