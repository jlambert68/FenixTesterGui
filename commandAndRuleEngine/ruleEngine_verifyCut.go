package commandAndRuleEngine

import (
	"errors"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

// VerifyIfComponentcanBeCut - Verify if a component can be Cut or not
//	Element 	Can be Cut 		Rule
//	B0			False				TCRuleCut001
//	B1			False				TCRuleCut002
//	B10			False				TCRuleCut003
//	B11f		False				TCRuleCut004a
//	B11l		False				TCRuleCut004b
//	B12			False				TCRuleCut005
//	B10*x* 		False				TCRuleCut006
//	B10*x 		False				TCRuleCut007
//	B10x*		False				TCRuleCut008
//	B11fx		False				TCRuleCut009a
//	B11lx		False				TCRuleCut009b
//	B12x		False				TCRuleCut010
//	TI			True				TCRuleCut011
//	Tix			False				TCRuleCut012
//	TIC(X)		True				TCRuleCut013
//	TICx(X)		False				TCRuleCut014

// Verify the simple rules if a component can be Cut or not
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) verifyIfComponentCanBeCutSimpleRules(testCaseUuid string, elementUuid string) (canBeCut bool, matchedRule string, err error) {

	// Get current TestCase
	currentTestCase, existsInMap := commandAndRuleEngine.testcases.TestCases[testCaseUuid]
	if existsInMap == false {

		errorId := "6efc5eca-dfd4-47a8-9599-9f946bcd43e0"
		err = errors.New(fmt.Sprintf("testcase with uuid '%s' doesn't exist in map with all testcases [ErrorID: %s]", elementUuid, errorId))

		return false, "", err
	}

	// Retrieve component to be verified for Cut
	element, existInMap := currentTestCase.TestCaseModelMap[elementUuid]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":          "4dd42abe-1798-46ad-b682-9713134a31f6",
			"elementUuid": elementUuid,
		}).Error(elementUuid + " could not be found in in map 'TestCaseModelMap'")

		errorId := "d399b328-1b9a-4022-b98a-206f39f3acef"
		err = errors.New(fmt.Sprintf("%s could not be found in in map 'TestCaseModelMap' [ErrorID: %s]", elementUuid, errorId))

		return false, "", err
	}

	// Extract component type to verify
	componentType := element.TestCaseModelElementType

	// Check simple rules of component can be Cut or not
	switch componentType {

	//	B0 - False - TCRuleCut001
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND:
		matchedRule = "TCRuleCut001"
		canBeCut = false

		//	B1 - False - TCRuleCut002
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE:

		matchedRule = "TCRuleCut002"

		//	B10 - False - TCRuleCut003
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND:
		matchedRule = "TCRuleCut002"
		canBeCut = false

		//	B11			False				TCRuleCut004
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND:
		matchedRule = "TCRuleCut004"
		canBeCut = false

		//	B12			False				TCRuleCut005
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND:
		matchedRule = "TCRuleCut005"
		canBeCut = false

		//	B10*x* 		False				TCRuleCut006
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10oxo_BOND:
		matchedRule = "TCRuleCut006"
		canBeCut = false

		//	B10*x 		False				TCRuleCut007
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10ox_BOND:
		matchedRule = "TCRuleCut007"
		canBeCut = false

		//	B10x*		False				TCRuleCut008
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10xo_BOND:
		matchedRule = "TCRuleCut008"
		canBeCut = false

		//	B11x		False				TCRuleCut009
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE:
		matchedRule = "TCRuleCut009"
		canBeCut = false

		//	B12x		False				TCRuleCut010
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE:
		matchedRule = "TCRuleCut010"
		canBeCut = false

		//	TI			True				TCRuleCut011
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION:
		matchedRule = "TCRuleCut011"
		canBeCut = true

		//	Tix			False				TCRuleCut012
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE:
		matchedRule = "TCRuleCut012"
		canBeCut = false

		//	TIC(X)		True				TCRuleCut013
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER:
		matchedRule = "TCRuleCut013"
		canBeCut = true

		//	TICx(X)		False				TCRuleCut014
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE:
		matchedRule = "TCRuleCut014"
		canBeCut = false

	default:
		matchedRule = ""
		canBeCut = false

		errorId := "7ce45c3f-53d5-4104-a164-8ecca190bd2c"
		err = errors.New(fmt.Sprintf("%s is an unknown componentTypeelement [ErrorID: %s]", componentType.String(), errorId))

		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":            "6d6f78da-809b-4a7f-a297-de39bf28409f",
			"componentType": componentType,
		}).Error(componentType.String() + " is an unknown componentType")

	}

	return canBeCut, matchedRule, err
}
