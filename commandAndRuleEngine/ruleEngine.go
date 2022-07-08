package commandAndRuleEngine

import (
	"errors"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

// VerifyIfComponentCanBeDeleted - Verify if a component can be deleted or not
//	Element 	Can be deleted 		Rule
//	B0			False				TCRuleDeletion001
//	B1			False				TCRuleDeletion002
//	B10			False				TCRuleDeletion003
//	B11f		False				TCRuleDeletion004a
//	B11l		False				TCRuleDeletion004b
//	B12			False				TCRuleDeletion005
//	B10*x* 		False				TCRuleDeletion006
//	B10*x 		False				TCRuleDeletion007
//	B10x*		False				TCRuleDeletion008
//	B11fx		False				TCRuleDeletion009a
//	B11lx		False				TCRuleDeletion009b
//	B12x		False				TCRuleDeletion010
//	TI			True				TCRuleDeletion011
//	Tix			False				TCRuleDeletion012
//	TIC(X)		True				TCRuleDeletion013
//	TICx(X)		False				TCRuleDeletion014
func VerifyIfComponentCanBeDeleted(uuid string, testCaseModelMap *map[string]fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementMessage) (err error) {

	return err
}

// Verify the complex rules if a component can be deleted or not
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) verifyIfComponentCanBeDeletedSimpleRules(componentType fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum) (canBeDeleted bool, err error) {

	// Check simple rules of component can be deleted or not
	switch componentType {

	//	B0 - False - TCRuleDeletion001
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND:
		canBeDeleted = false

		//	B1 - False - TCRuleDeletion002
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1_BOND_NONE_SWAPPABLE:
		canBeDeleted = false

		//	B10 - False - TCRuleDeletion003
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND:
		canBeDeleted = false

		//	B11			False				TCRuleDeletion004
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND:
		canBeDeleted = false

		//	B12			False				TCRuleDeletion005
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND:
		canBeDeleted = false

		//	B10*x* 		False				TCRuleDeletion006
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10oxo_BOND:
		canBeDeleted = false

		//	B10*x 		False				TCRuleDeletion007
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10ox_BOND:
		canBeDeleted = false

		//	B10x*		False				TCRuleDeletion008
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10xo_BOND:

		//	B11x		False				TCRuleDeletion009
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE:
		canBeDeleted = false

		//	B12x		False				TCRuleDeletion010
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE:
		canBeDeleted = false

		//	TI			True				TCRuleDeletion011
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION:
		canBeDeleted = true

		//	Tix			False				TCRuleDeletion012
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE:
		canBeDeleted = false

		//	TIC(X)		True				TCRuleDeletion013
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER:
		canBeDeleted = true

		//	TICx(X)		False				TCRuleDeletion014
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE:
		canBeDeleted = false

	default:
		canBeDeleted = false
		err = errors.New(componentType.String() + " is an unknown componentType")

		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":            "3be29c9d-1db6-47dc-8ee1-8dcdfecda074",
			"componentType": componentType,
		}).Error(componentType.String() + " is an unknown componentType")

	}

	return canBeDeleted, err
}

// Verify the complex rules if a component can be deleted or not
// Rules how deletion of an element is done
// X = any allowed structure
// What to remove			Remove in structure				Result after deletion		Rule
// n= TIC(X)				B1-n-B1							B0							TCRuleDeletion101
// n=TI or TIC(X)			B11f-n-B11l						B10							TCRuleDeletion102
// n=TI or TIC(X)			B11fx-n-B11lx					B10*x*						TCRuleDeletion103
// n=TI or TIC(X)			B11f-n-B11lx					B10x*						TCRuleDeletion104
// n=TI or TIC(X)			B11fx-n-B11l					B10*x						TCRuleDeletion105
// n=TI or TIC(X)			B11f-n-B12-X					B11f-X						TCRuleDeletion106
// n=TI or TIC(X)			B11fx-n-B12x-X					B11fx-X						TCRuleDeletion107
// n=TI or TIC(X)			B11f-n-B12x-X					B11fx-X						TCRuleDeletion108
// n=TI or TIC(X)			B11fx-n-B12-X					B11fx-X						TCRuleDeletion109
// n=TI or TIC(X)			X-B12-n-B11l					X-B11l						TCRuleDeletion110
// n=TI or TIC(X)			X-B12x-n-B11lx					X-B11lx						TCRuleDeletion111
// n=TI or TIC(X)			X-B12-n-B11lx					X-B11lx						TCRuleDeletion112
// n=TI or TIC(X)			X-B12x-n-B11l					X-B11lx						TCRuleDeletion113
// n=TI or TIC(X)			X-B12-n-B12-X					X-B12-X						TCRuleDeletion114
// n=TI or TIC(X)			X-B12x-n-B12x-X					X-B12x-X					TCRuleDeletion115
// n=TI or TIC(X)			X-B12-n-B12x-X					X-B12x-X					TCRuleDeletion116
// n=TI or TIC(X)			X-B12x-n-B12-X					X-B12x-X					TCRuleDeletion117
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) verifyIfComponentCanBeDeletedWithComplexRules(uuid string) (err error) {

	var ruleName string
	var ruleCanBeProcessed bool

	ruleName = ""
	ruleCanBeProcessed = false

	// Extract data for Previous Elementfunc (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct)
	currentElementUuid := uuid
	currentElement, existInMap := commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[currentElementUuid]
	if existInMap == false {
		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":                 "8c69112a-31ea-4606-89a5-54b80789e691",
			"currentElementUuid": currentElementUuid,
		}).Error(currentElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(currentElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return err
	}
	currentElementType := currentElement.TestCaseModelElementType

	// Extract data for Previous Element
	previousElementUuid := commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[currentElementUuid].PreviousElementUuid
	previousElement, existInMap := commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[previousElementUuid]
	if existInMap == false {
		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":                  "d801356c-5ab6-48d7-bcd5-73d820b86d1e",
			"previousElementUuid": previousElementUuid,
		}).Error(previousElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(previousElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return err
	}
	previousElementType := previousElement.TestCaseModelElementType

	// Extract data for Next Element
	nextElementUuid := commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[currentElementUuid].NextElementUuid
	nextElement, existInMap := commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[nextElementUuid]
	if existInMap == false {
		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":              "6c8c7382-48c7-4041-9f19-0c9b11298bbf",
			"nextElementUuid": nextElementUuid,
		}).Error(nextElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(nextElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return err
	}
	nextlementType := nextElement.TestCaseModelElementType

	// TCRuleDeletion101
	// What to remove			Remove in structure				Result after deletion		Rule
	// n= TIC(X)				B1-n-B1							B0							TCRuleDeletion101
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1_BOND_NONE_SWAPPABLE &&
		currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1_BOND_NONE_SWAPPABLE {

		// Rule OK
		ruleName = "TCRuleDeletion101"
		ruleCanBeProcessed = true

	}

	// TCRuleDeletion102
	// What to remove			Remove in structure				Result after deletion		Rule
	// n=TI or TIC(X)			B11f-n-B11l						B10							TCRuleDeletion102
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND {

		// Rule OK
		ruleName = "TCRuleDeletion102"
		ruleCanBeProcessed = true

	}

	// TCRuleDeletion103
	// What to remove			Remove in structure				Result after deletion		Rule
	// n=TI or TIC(X)			B11fx-n-B11lx					B10*x*						TCRuleDeletion103
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE {

		// Rule OK
		ruleName = "TCRuleDeletion103"
		ruleCanBeProcessed = true

	}

	// TCRuleDeletion104
	// What to remove			Remove in structure				Result after deletion		Rule
	// n=TI or TIC(X)			B11f-n-B11lx					B10x*						TCRuleDeletion104
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE {

		// Rule OK
		ruleName = "TCRuleDeletion104"
		ruleCanBeProcessed = true

	}

	// TCRuleDeletion105
	// What to remove			Remove in structure				Result after deletion		Rule
	// n=TI or TIC(X)			B11fx-n-B11l					B10*x						TCRuleDeletion105
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND {

		// Rule OK
		ruleName = "TCRuleDeletion105"
		ruleCanBeProcessed = true

	}

	// TCRuleDeletion106
	// What to remove			Remove in structure				Result after deletion		Rule
	// n=TI or TIC(X)			B11f-n-B12-X					B11f-X						TCRuleDeletion106
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND {

		// Rule OK
		ruleName = "TCRuleDeletion106"
		ruleCanBeProcessed = true

	}

	// TCRuleDeletion107
	// What to remove			Remove in structure				Result after deletion		Rule
	// n=TI or TIC(X)			B11fx-n-B12x-X					B11fx-X						TCRuleDeletion107
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE {

		// Rule OK
		ruleName = "TCRuleDeletion107"
		ruleCanBeProcessed = true

	}

	// TCRuleDeletion108
	// What to remove			Remove in structure				Result after deletion		Rule
	// n=TI or TIC(X)			B11f-n-B12x-X					B11fx-X						TCRuleDeletion108
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE {

		// Rule OK
		ruleName = "TCRuleDeletion108"
		ruleCanBeProcessed = true

	}

	// TCRuleDeletion109
	// What to remove			Remove in structure				Result after deletion		Rule
	// n=TI or TIC(X)			B11fx-n-B12-X					B11fx-X						TCRuleDeletion109
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND {

		// Rule OK
		ruleName = "TCRuleDeletion109"
		ruleCanBeProcessed = true

	}

	// TCRuleDeletion110
	// What to remove			Remove in structure				Result after deletion		Rule
	// n=TI or TIC(X)			X-B12-n-B11l					X-B11l						TCRuleDeletion110
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND {

		// Rule OK
		ruleName = "TCRuleDeletion110"
		ruleCanBeProcessed = true

	}

	// TCRuleDeletion111
	// What to remove			Remove in structure				Result after deletion		Rule
	// n=TI or TIC(X)			X-B12x-n-B11lx					X-B11lx						TCRuleDeletion111
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE {

		// Rule OK
		ruleName = "TCRuleDeletion111"
		ruleCanBeProcessed = true

	}

	// TCRuleDeletion112
	// What to remove			Remove in structure				Result after deletion		Rule
	// n=TI or TIC(X)			X-B12-n-B11lx					X-B11lx						TCRuleDeletion112
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE {

		// Rule OK
		ruleName = "TCRuleDeletion112"
		ruleCanBeProcessed = true

	}

	// TCRuleDeletion113
	// What to remove			Remove in structure				Result after deletion		Rule
	// n=TI or TIC(X)			X-B12x-n-B11l					X-B11lx						TCRuleDeletion113
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND {

		// Rule OK
		ruleName = "TCRuleDeletion113"
		ruleCanBeProcessed = true

	}

	// TCRuleDeletion114
	// What to remove			Remove in structure				Result after deletion		Rule
	// n=TI or TIC(X)			X-B12-n-B12-X					X-B12-X						TCRuleDeletion114
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND {

		// Rule OK
		ruleName = "TCRuleDeletion114"
		ruleCanBeProcessed = true

	}

	// TCRuleDeletion115
	// What to remove			Remove in structure				Result after deletion		Rule
	// n=TI or TIC(X)			X-B12x-n-B12x-X					X-B12x-X					TCRuleDeletion115
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE {

		// Rule OK
		ruleName = "TCRuleDeletion115"
		ruleCanBeProcessed = true

	}

	// TCRuleDeletion116
	// What to remove			Remove in structure				Result after deletion		Rule
	// n=TI or TIC(X)			X-B12-n-B12x-X					X-B12x-X					TCRuleDeletion116
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE {

		// Rule OK
		ruleName = "TCRuleDeletion116"
		ruleCanBeProcessed = true

	}

	// TCRuleDeletion117
	// What to remove			Remove in structure				Result after deletion		Rule
	// n=TI or TIC(X)			X-B12x-n-B12-X					X-B12x-X					TCRuleDeletion117
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND {

		// Rule OK
		ruleName = "TCRuleDeletion117"
		ruleCanBeProcessed = true

	}

	switch ruleName {

	case "TCRuleDeletion101", "TCRuleDeletion102", "TCRuleDeletion103", "TCRuleDeletion104",
		"TCRuleDeletion105", "TCRuleDeletion106", "TCRuleDeletion107", "TCRuleDeletion108",
		"TCRuleDeletion109", "TCRuleDeletion110", "TCRuleDeletion111", "TCRuleDeletion112",
		"TCRuleDeletion113", "TCRuleDeletion114", "TCRuleDeletion115", "TCRuleDeletion116", "TCRuleDeletion117":

		// Deletion can be execute by complex deletion rules
		err = nil

	default:
		// The criteria for Deleting is not met by complex deletion rule
		err = errors.New("The criteria for any complex deletion rule was not met")

		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":                  "4e62df40-4192-4f45-ac4d-adbf1a687ad2",
			"previousElementType": previousElementType,
			"currentElementType":  currentElementType,
			"nextlementType":      nextlementType,
		}).Error("The criteria for any complex deletion rule was not met")

		return err

	}

	return err

}
