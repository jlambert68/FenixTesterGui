package commandAndRuleEngine

import (
	"errors"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"reflect"
)

//Rules if an element are allowed to be swapped into another element
//X = any allowed structure
//	Element			Element can be swapped out		Rule
//	B0				True							TCRuleSwap001
//	B1				False							TCRuleSwap002
//	B10				True							TCRuleSwap003
//	B11				True							TCRuleSwap004
//	B12				True							TCRuleSwap005
//	B10*x*			True							TCRuleSwap006
//	B10*x			True							TCRuleSwap007
//	B10x*			True							TCRuleSwap008
//	B11x			False							TCRuleSwap009
//	B12x			False							TCRuleSwap010
//	TI				False							TCRuleSwap011
//	Tix				False							TCRuleSwap012
//	TIC(X)			False							TCRuleSwap013
//	TICx(X)			False							TCRuleSwap014

// Verify the simple rules if a component can be Swapped or not
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) verifyIfComponentCanBeSwappedSimpleRules(elementUuid string) (canBeSwapped bool, matchedRule string, err error) {

	// Retrieve component to be verified for Swap
	element, existInMap := commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[elementUuid]
	if existInMap == false {
		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":          "9d8aebb2-4409-4236-8740-4ca396007088",
			"elementUuid": elementUuid,
		}).Error(elementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(elementUuid + " could not be found in in map 'TestCaseModelMap'")

		return false, "", err
	}

	// Extract component type to verify
	componentType := element.TestCaseModelElementType

	// Check simple rules of component can be Swapped or not
	switch componentType {

	//	B0 - False - TCRuleSwap001
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND:
		matchedRule = "TCRuleSwap001"
		canBeSwapped = false

		//	B1 - False - TCRuleSwap002
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1_BOND_NONE_SWAPPABLE:
		matchedRule = "TCRuleSwap002"

		//	B10 - False - TCRuleSwap003
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND:
		matchedRule = "TCRuleSwap002"
		canBeSwapped = false

		//	B11			False				TCRuleSwap004
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND:
		matchedRule = "TCRuleSwap004"
		canBeSwapped = false

		//	B12			False				TCRuleSwap005
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND:
		matchedRule = "TCRuleSwap005"
		canBeSwapped = false

		//	B10*x* 		False				TCRuleSwap006
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10oxo_BOND:
		matchedRule = "TCRuleSwap006"
		canBeSwapped = false

		//	B10*x 		False				TCRuleSwap007
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10ox_BOND:
		matchedRule = "TCRuleSwap007"
		canBeSwapped = false

		//	B10x*		False				TCRuleSwap008
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10xo_BOND:
		matchedRule = "TCRuleSwap008"
		canBeSwapped = false

		//	B11x		False				TCRuleSwap009
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE:
		matchedRule = "TCRuleSwap009"
		canBeSwapped = false

		//	B12x		False				TCRuleSwap010
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE:
		matchedRule = "TCRuleSwap010"
		canBeSwapped = false

		//	TI			True				TCRuleSwap011
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION:
		matchedRule = "TCRuleSwap011"
		canBeSwapped = true

		//	Tix			False				TCRuleSwap012
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE:
		matchedRule = "TCRuleSwap012"
		canBeSwapped = false

		//	TIC(X)		True				TCRuleSwap013
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER:
		matchedRule = "TCRuleSwap013"
		canBeSwapped = true

		//	TICx(X)		False				TCRuleSwap014
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE:
		matchedRule = "TCRuleSwap014"
		canBeSwapped = false

	default:
		matchedRule = ""
		canBeSwapped = false

		err = errors.New(componentType.String() + " is an unknown componentType")

		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":            "1e9846c0-9c79-4e6f-aabd-46d2b3d80053",
			"componentType": componentType,
		}).Error(componentType.String() + " is an unknown componentType")

	}

	return canBeSwapped, matchedRule, err
}

// Verify the complex rules if a component can be Swapped or not
// Rules how swapping out an element for another element
// X = any allowed structure
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC(X)			B0					n 		B0								B1-n-B1					TCRuleSwap101
//	n=TIC or TIC(X)		B10					n		TIC(B10)						TIC(B11f-n-B11l)		TCRuleSwap102
//	n=TIC or TIC(X)		B11f				n		TIC(B11f-X)						TIC(B11f-n-B12-X)		TCRuleSwap103
//	n=TIC or TIC(X)		B11l				n		TIC(X-B11l)						TIC(X-B12-n-B11l)		TCRuleSwap104
//	n=TIC or TIC(X)		B12					n		X-B12-X							X-B12-n-B12-X			TCRuleSwap105
//	n=TIC or TIC(X)		B10x*				n		TIC(B10*x*)						TIC(B11x-n-B11x)		TCRuleSwap106
//	n=TIC or TIC(X)		B10*x				n		TIC(B10*x)						TIC(B11x-n-B11)			TCRuleSwap107
//	n=TIC or TIC(X)		B10x*				n		TIC(B10x*)						TIC(B11-n-B11x)			TCRuleSwap108
//No other combinations of swapping elements are allowed

// Verify the Complex rules if a component can be Swapped or not
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) verifyIfComponentCanBeSwappedWithComplexRules(uuid string) (matchedRule string, err error) {

	var ruleName string
	var ruleCanBeProcessed bool

	ruleName = ""
	ruleCanBeProcessed = false

	// Extract data for Previous Elementfunc (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct)
	currentElementUuid := uuid
	currentElement, existInMap := commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[currentElementUuid]
	if existInMap == false {
		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":                 "d450e7e5-32f4-42e9-b371-279d5bfe9d14",
			"currentElementUuid": currentElementUuid,
		}).Error(currentElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(currentElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return "", err
	}
	currentElementType := currentElement.TestCaseModelElementType

	// Extract data for Previous Element
	previousElementUuid := commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[currentElementUuid].PreviousElementUuid
	previousElement, existInMap := commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[previousElementUuid]
	if existInMap == false {
		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":                  "42b050c6-3d63-45fc-9b6c-6b4a6e02516f",
			"previousElementUuid": previousElementUuid,
		}).Error(previousElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(previousElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return "", err
	}
	previousElementType := previousElement.TestCaseModelElementType

	// Extract data for Next Element
	nextElementUuid := commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[currentElementUuid].NextElementUuid
	nextElement, existInMap := commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[nextElementUuid]
	if existInMap == false {
		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":              "bf2ac32f-5edb-472f-af73-87d04400e132",
			"nextElementUuid": nextElementUuid,
		}).Error(nextElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(nextElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return "", err
	}
	nextlementType := nextElement.TestCaseModelElementType

	// TCRuleSwap101
	// What to remove			Remove in structure				Result after Swap		Rule
	// n= TIC(X)				B1-n-B1							B0							TCRuleSwap101
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1_BOND_NONE_SWAPPABLE &&
		currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1_BOND_NONE_SWAPPABLE {

		// Rule OK
		ruleName = "TCRuleSwap101"
		ruleCanBeProcessed = true

	}

	// TCRuleSwap102
	// What to remove			Remove in structure				Result after Swap		Rule
	// n=TI or TIC(X)			B11f-n-B11l						B10							TCRuleSwap102
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND {

		// Rule OK
		ruleName = "TCRuleSwap102"
		ruleCanBeProcessed = true

	}

	// TCRuleSwap103
	// What to remove			Remove in structure				Result after Swap		Rule
	// n=TI or TIC(X)			B11fx-n-B11lx					B10*x*						TCRuleSwap103
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE {

		// Rule OK
		ruleName = "TCRuleSwap103"
		ruleCanBeProcessed = true

	}

	// TCRuleSwap104
	// What to remove			Remove in structure				Result after Swap		Rule
	// n=TI or TIC(X)			B11f-n-B11lx					B10x*						TCRuleSwap104
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE {

		// Rule OK
		ruleName = "TCRuleSwap104"
		ruleCanBeProcessed = true

	}

	// TCRuleSwap105
	// What to remove			Remove in structure				Result after Swap		Rule
	// n=TI or TIC(X)			B11fx-n-B11l					B10*x						TCRuleSwap105
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND {

		// Rule OK
		ruleName = "TCRuleSwap105"
		ruleCanBeProcessed = true

	}

	// TCRuleSwap106
	// What to remove			Remove in structure				Result after Swap		Rule
	// n=TI or TIC(X)			B11f-n-B12-X					B11f-X						TCRuleSwap106
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND {

		// Rule OK
		ruleName = "TCRuleSwap106"
		ruleCanBeProcessed = true

	}

	// TCRuleSwap107
	// What to remove			Remove in structure				Result after Swap		Rule
	// n=TI or TIC(X)			B11fx-n-B12x-X					B11fx-X						TCRuleSwap107
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE {

		// Rule OK
		ruleName = "TCRuleSwap107"
		ruleCanBeProcessed = true

	}

	// TCRuleSwap108
	// What to remove			Remove in structure				Result after Swap		Rule
	// n=TI or TIC(X)			B11f-n-B12x-X					B11fx-X						TCRuleSwap108
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE {

		// Rule OK
		ruleName = "TCRuleSwap108"
		ruleCanBeProcessed = true

	}

	// TCRuleSwap109
	// What to remove			Remove in structure				Result after Swap		Rule
	// n=TI or TIC(X)			B11fx-n-B12-X					B11fx-X						TCRuleSwap109
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND {

		// Rule OK
		ruleName = "TCRuleSwap109"
		ruleCanBeProcessed = true

	}

	// TCRuleSwap110
	// What to remove			Remove in structure				Result after Swap		Rule
	// n=TI or TIC(X)			X-B12-n-B11l					X-B11l						TCRuleSwap110
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND {

		// Rule OK
		ruleName = "TCRuleSwap110"
		ruleCanBeProcessed = true

	}

	// TCRuleSwap111
	// What to remove			Remove in structure				Result after Swap		Rule
	// n=TI or TIC(X)			X-B12x-n-B11lx					X-B11lx						TCRuleSwap111
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE {

		// Rule OK
		ruleName = "TCRuleSwap111"
		ruleCanBeProcessed = true

	}

	// TCRuleSwap112
	// What to remove			Remove in structure				Result after Swap		Rule
	// n=TI or TIC(X)			X-B12-n-B11lx					X-B11lx						TCRuleSwap112
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE {

		// Rule OK
		ruleName = "TCRuleSwap112"
		ruleCanBeProcessed = true

	}

	// TCRuleSwap113
	// What to remove			Remove in structure				Result after Swap		Rule
	// n=TI or TIC(X)			X-B12x-n-B11l					X-B11lx						TCRuleSwap113
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND {

		// Rule OK
		ruleName = "TCRuleSwap113"
		ruleCanBeProcessed = true

	}

	// TCRuleSwap114
	// What to remove			Remove in structure				Result after Swap		Rule
	// n=TI or TIC(X)			X-B12-n-B12-X					X-B12-X						TCRuleSwap114
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND {

		// Rule OK
		ruleName = "TCRuleSwap114"
		ruleCanBeProcessed = true

	}

	// TCRuleSwap115
	// What to remove			Remove in structure				Result after Swap		Rule
	// n=TI or TIC(X)			X-B12x-n-B12x-X					X-B12x-X					TCRuleSwap115
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE {

		// Rule OK
		ruleName = "TCRuleSwap115"
		ruleCanBeProcessed = true

	}

	// TCRuleSwap116
	// What to remove			Remove in structure				Result after Swap		Rule
	// n=TI or TIC(X)			X-B12-n-B12x-X					X-B12x-X					TCRuleSwap116
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE {

		// Rule OK
		ruleName = "TCRuleSwap116"
		ruleCanBeProcessed = true

	}

	// TCRuleSwap117
	// What to remove			Remove in structure				Result after Swap		Rule
	// n=TI or TIC(X)			X-B12x-n-B12-X					X-B12x-X					TCRuleSwap117
	if ruleCanBeProcessed == false &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE &&
		(currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER) &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND {

		// Rule OK
		ruleName = "TCRuleSwap117"
		ruleCanBeProcessed = true

	}

	switch ruleName {

	case "TCRuleSwap101", "TCRuleSwap102", "TCRuleSwap103", "TCRuleSwap104",
		"TCRuleSwap105", "TCRuleSwap106", "TCRuleSwap107", "TCRuleSwap108",
		"TCRuleSwap109", "TCRuleSwap110", "TCRuleSwap111", "TCRuleSwap112",
		"TCRuleSwap113", "TCRuleSwap114", "TCRuleSwap115", "TCRuleSwap116", "TCRuleSwap117":

		// Swap can be execute by complex Swap rules
		err = nil

	default:
		// The criteria for Deleting is not met by complex Swap rule
		err = errors.New("The criteria for any complex Swap rule was not met")

		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":                  "62c2d049-c3aa-40be-b192-9a2a4c2ad42c",
			"previousElementType": previousElementType,
			"currentElementType":  currentElementType,
			"nextlementType":      nextlementType,
		}).Error("The criteria for any complex Swap rule was not met")

		return "", err

	}

	return ruleName, err

}

// Verify that all UUIDs are correct in component to be swapped in. Means that no empty uuid is allowed and they all are correct
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) verifyThatThereAreNoZombieElementsInComponent(immatureElement immatureElementStruct) (err error) {

	var allUuidKeys []string

	// Extract all elements by key from component
	for _, elementKey := range immatureElement.immatureElementMap {
		allUuidKeys = append(allUuidKeys, elementKey.ImmatureElementUuid)
	}

	// Follow the path from "first element and remove the found element from 'allUuidKeys'
	commandAndRuleEngineObject.recursiveZombieElementSearchInComponentModel(immatureElement.firstElementUuid, &allUuidKeys, &immatureElement)

	// If there are elements left in slice then there were zombie elements, which there shouldn't be
	if len(allUuidKeys) != 0 {
		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":                               "3e519b3c-367d-42d5-a8ce-ff507efd8972",
			"allUuidKeys":                      allUuidKeys,
			"Number of Zombie Elements":        len(allUuidKeys),
			"immatureElement.firstElementUuid": immatureElement.firstElementUuid,
		}).Error("There existed Zombie elements in 'immatureElement.immatureElementMap'")

		err = errors.New("there existed Zombie elements in 'immatureElement.immatureElementMap', for " + immatureElement.firstElementUuid)

		return err
	}

	return err
}

// Verify all children, in ImmatureEleemnt-model and remove the found element from 'allUuidKeys'
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) recursiveZombieElementSearchInComponentModel(elementsUuid string, allUuidKeys *[]string, immatureElement *immatureElementStruct) (err error) {

	// Extract current element
	currentElement, existInMap := immatureElement.immatureElementMap[elementsUuid]

	// If the element doesn't exit then there is something really wrong
	if existInMap == false {
		// This shouldn't happen
		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":           "9f628356-2ea2-48a6-8e6a-546a5f97f05b",
			"elementsUuid": elementsUuid,
		}).Error(elementsUuid + " could not be found in in map 'immatureElement.immatureElementMap'")

		err = errors.New(elementsUuid + " could not be found in in map 'immatureElement.immatureElementMap'")

		return err
	}

	// Element has child-element then go that path
	if currentElement.FirstChildElementUuid != elementsUuid {
		err = commandAndRuleEngineObject.recursiveDeleteOfChildElements(currentElement.FirstChildElementUuid)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return err
	}

	// If element has a next-element the go that path
	if currentElement.NextElementUuid != elementsUuid {
		err = commandAndRuleEngineObject.recursiveDeleteOfChildElements(currentElement.NextElementUuid)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return err
	}

	// Remove current element from "slice of all elements in map"
	allUuidKeys = findElementInSliceAndRemove(allUuidKeys, elementsUuid)

	return err
}

// Remove 'uuid' from slice
func findElementInSliceAndRemove(sliceToWorkOn *[]string, uuid string) (returnSlice *[]string) {

	var index int
	var uuidInSLice string

	// Find the index of the 'uuid'
	for index, uuidInSLice = range *sliceToWorkOn {
		if uuidInSLice == uuid {
			break
		}
	}

	// Create a temporary slice to work on
	tempSlice := *sliceToWorkOn

	// Remove the element in the slice
	tempSlice[index] = tempSlice[len(tempSlice)-1]
	tempSlice = tempSlice[:len(tempSlice)-1]

	returnSlice = &tempSlice

	return returnSlice
}

// Verify that all UUIDs are correct in component to be swapped in. Means that no empty uuid is allowed and they all are correct
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) verifyThatAllUuidsAreCorrectInComponent(immatureElement immatureElementStruct) (err error) {

	// Loop all fields and find the ones defined as 'String'. Verify that content is a UUID
	e := reflect.ValueOf(&immatureElement.immatureElementMap).Elem()

	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varType := e.Type().Field(i).Type
		varValue := e.Field(i).Interface()

		// If Type is 'String' then verify that it's a correct UUID
		if varType.Kind() == reflect.String {
			//TODO Implement UUID validation instead
			fmt.Printf("****************************REPLACE WITH CORRECT CHECK ON UUID-type***********************************************")
		}
		fmt.Printf("%v %v %v\n", varName, varType, varValue)
	}

	return err
}

// Verify all children, in new Element-model to be swapped in, that they contain correct UUIDs
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) recursiveVerifyAllUuidOfChildElements(elementsUuid string) (err error) {

	// Extract current element
	currentElement, existInMap := commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[elementsUuid]

	// If the element doesn't exit then there is something really wrong
	if existInMap == false {
		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":           "9eae5791-88f0-481d-a7d9-21123b9eadfe",
			"elementsUuid": elementsUuid,
		}).Error(elementsUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(elementsUuid + " could not be found in in map 'immatureElement.immatureElementMap'")

		return err
	}

	// Element has child-element then go that path
	if currentElement.FirstChildElementUuid != elementsUuid {
		err = commandAndRuleEngineObject.recursiveDeleteOfChildElements(currentElement.FirstChildElementUuid)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return err
	}

	// If element has a next-element the go that path
	if currentElement.NextElementUuid != elementsUuid {
		err = commandAndRuleEngineObject.recursiveDeleteOfChildElements(currentElement.NextElementUuid)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return err
	}

	// Remove current element from Map
	delete(commandAndRuleEngineObject.testcaseModel.TestCaseModelMap, elementsUuid)

	return err
}
