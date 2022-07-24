package commandAndRuleEngine

import (
	"FenixTesterGui/gui/UnitTestTestData"
	"FenixTesterGui/testCase/testCaseModel"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTestCaseModel(t *testing.T) {

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

}
