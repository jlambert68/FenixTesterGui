package commandAndRuleEngine

import (
	"errors"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

// VerifyIfComponentcanBeCopied - Verify if a component can be copied or not
//	Element 	Can be copied 		Rule
//	B0			False				TCRuleCopy001
//	B1			False				TCRuleCopy002
//	B10			False				TCRuleCopy003
//	B11f		False				TCRuleCopy004a
//	B11l		False				TCRuleCopy004b
//	B12			False				TCRuleCopy005
//	B10*x* 		False				TCRuleCopy006
//	B10*x 		False				TCRuleCopy007
//	B10x*		False				TCRuleCopy008
//	B11fx		False				TCRuleCopy009a
//	B11lx		False				TCRuleCopy009b
//	B12x		False				TCRuleCopy010
//	TI			True				TCRuleCopy011
//	Tix			True				TCRuleCopy012
//	TIC(X)		True				TCRuleCopy013
//	TICx(X)		True				TCRuleCopy014

// Verify the simple rules if a component can be copied or not
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) verifyIfComponentCanBeCopiedSimpleRules(testCaseUuid string, elementUuid string) (canBeCopied bool, matchedRule string, err error) {

	// Get current TestCase
	currentTestCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	if existsInMap == false {

		errorId := "e60225ed-abcf-4339-b191-061b3084e92f"
		err = errors.New(fmt.Sprintf("testcase with uuid '%s' doesn't exist in map with all Testcases [ErrorID: %s]", elementUuid, errorId))

		return false, "", err
	}

	// Retrieve component to be verified for Copy
	element, existInMap := currentTestCase.TestCaseModelMap[elementUuid]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":          "a0c634b0-d55a-4b71-aede-968be81f153b",
			"elementUuid": elementUuid,
		}).Error(elementUuid + " could not be found in in map 'TestCaseModelMap'")

		errorId := "6d0f1456-e0e6-4e85-96b7-9a76a1838690"
		err = errors.New(fmt.Sprintf("%s could not be found in in map 'TestCaseModelMap' [ErrorID: %s]", elementUuid, errorId))

		return false, "", err
	}

	// Extract component type to verify
	componentType := element.TestCaseModelElementType

	// Check simple rules of component can be copied or not
	switch componentType {

	//	B0 - False - TCRuleCopy001
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND:
		matchedRule = "TCRuleCopy001"
		canBeCopied = false

		//	B1 - False - TCRuleCopy002
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE:

		matchedRule = "TCRuleCopy002"

		//	B10 - False - TCRuleCopy003
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND:
		matchedRule = "TCRuleCopy002"
		canBeCopied = false

		//	B11			False				TCRuleCopy004
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND:
		matchedRule = "TCRuleCopy004"
		canBeCopied = false

		//	B12			False				TCRuleCopy005
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND:
		matchedRule = "TCRuleCopy005"
		canBeCopied = false

		//	B10*x* 		False				TCRuleCopy006
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10oxo_BOND:
		matchedRule = "TCRuleCopy006"
		canBeCopied = false

		//	B10*x 		False				TCRuleCopy007
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10ox_BOND:
		matchedRule = "TCRuleCopy007"
		canBeCopied = false

		//	B10x*		False				TCRuleCopy008
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10xo_BOND:
		matchedRule = "TCRuleCopy008"
		canBeCopied = false

		//	B11x		False				TCRuleCopy009
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE:
		matchedRule = "TCRuleCopy009"
		canBeCopied = false

		//	B12x		False				TCRuleCopy010
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE:
		matchedRule = "TCRuleCopy010"
		canBeCopied = false

		//	TI			True				TCRuleCopy011
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION:
		matchedRule = "TCRuleCopy011"
		canBeCopied = true

		//	Tix			False				TCRuleCopy012
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE:
		matchedRule = "TCRuleCopy012"
		canBeCopied = true

		//	TIC(X)		True				TCRuleCopy013
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER:
		matchedRule = "TCRuleCopy013"
		canBeCopied = true

		//	TICx(X)		False				TCRuleCopy014
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE:
		matchedRule = "TCRuleCopy014"
		canBeCopied = true

	default:
		matchedRule = ""
		canBeCopied = false

		errorId := "53ec61f2-f50b-49e5-81af-aba27448056f"
		err = errors.New(fmt.Sprintf("%s is an unknown componentTypeelement [ErrorID: %s]", componentType.String(), errorId))

		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":            "4edf145f-feef-4d73-abf2-721a75b0a509",
			"componentType": componentType,
		}).Error(componentType.String() + " is an unknown componentType")

	}

	return canBeCopied, matchedRule, err
}
