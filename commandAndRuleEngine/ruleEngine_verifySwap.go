package commandAndRuleEngine

import (
	"FenixTesterGui/testCase/testCaseModel"
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
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) verifyIfComponentCanBeSwappedSimpleRules(testCaseUuid string, elementUuid string) (canBeSwapped bool, matchedRule string, err error) {

	// Get current TestCase
	currentTestCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all Testcases")
		return false, "", err
	}

	// Retrieve component to be verified for Swap
	element, existInMap := currentTestCase.TestCaseModelMap[elementUuid]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":          "9d8aebb2-4409-4236-8740-4ca396007088",
			"elementUuid": elementUuid,
		}).Error(elementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(elementUuid + " could not be found in in map 'TestCaseModelMap'")

		return false, "", err
	}

	// Extract component type to verify
	componentType := element.MatureTestCaseModelElementMessage.TestCaseModelElementType

	// Check simple rules of component can be Swapped or not
	switch componentType {

	//	B0 - True - TCRuleSwap001
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND:
		matchedRule = "TCRuleSwap001"
		canBeSwapped = true

		//	B1 - False - TCRuleSwap002
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE:
		matchedRule = "TCRuleSwap002"
		canBeSwapped = false

		//	B10 - True - TCRuleSwap003
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND:
		matchedRule = "TCRuleSwap002"
		canBeSwapped = true

		//	B11			True				TCRuleSwap004
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND:
		matchedRule = "TCRuleSwap004"
		canBeSwapped = true

		//	B12			True				TCRuleSwap005
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND:
		matchedRule = "TCRuleSwap005"
		canBeSwapped = true

		//	B10*x* 		True				TCRuleSwap006
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10oxo_BOND:
		matchedRule = "TCRuleSwap006"
		canBeSwapped = true

		//	B10*x 		True				TCRuleSwap007
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10ox_BOND:
		matchedRule = "TCRuleSwap007"
		canBeSwapped = true

		//	B10x*		True				TCRuleSwap008
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10xo_BOND:
		matchedRule = "TCRuleSwap008"
		canBeSwapped = true

		//	B11x		False				TCRuleSwap009
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE:
		matchedRule = "TCRuleSwap009"
		canBeSwapped = true

		//	B12x		False				TCRuleSwap010
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE:
		matchedRule = "TCRuleSwap010"
		canBeSwapped = true

		//	TI			True				TCRuleSwap011
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION:
		matchedRule = "TCRuleSwap011"
		canBeSwapped = false

		//	Tix			False				TCRuleSwap012
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE:
		matchedRule = "TCRuleSwap012"
		canBeSwapped = false

		//	TIC(X)		True				TCRuleSwap013
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER:
		matchedRule = "TCRuleSwap013"
		canBeSwapped = false

		//	TICx(X)		False				TCRuleSwap014
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE:
		matchedRule = "TCRuleSwap014"
		canBeSwapped = false

	default:
		matchedRule = ""
		canBeSwapped = false

		err = errors.New(componentType.String() + " is an unknown componentType")

		commandAndRuleEngine.logger.WithFields(logrus.Fields{
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
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) verifyIfComponentCanBeSwappedWithComplexRules(testCaseUuid string, uuidToSwapOut string, elementTypeToBeSwappedIn fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum) (matchedRule string, err error) {

	var ruleName string
	var ruleCanBeProcessed bool

	ruleName = ""
	ruleCanBeProcessed = false

	// Get current TestCase
	currentTestCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all Testcases")
		return "", err
	}

	// Extract data for Previous Elementfunc (commandAndRuleEngine *CommandAndRuleEngineObjectStruct)
	currentElementUuid := uuidToSwapOut
	currentElement, existInMap := currentTestCase.TestCaseModelMap[currentElementUuid]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":                 "d450e7e5-32f4-42e9-b371-279d5bfe9d14",
			"currentElementUuid": currentElementUuid,
		}).Error(currentElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(currentElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return "", err
	}
	currentElementType := currentElement.MatureTestCaseModelElementMessage.TestCaseModelElementType

	// Extract data for Previous Element
	previousElementUuid := currentTestCase.TestCaseModelMap[currentElementUuid].MatureTestCaseModelElementMessage.PreviousElementUuid
	previousElement, existInMap := currentTestCase.TestCaseModelMap[previousElementUuid]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":                  "42b050c6-3d63-45fc-9b6c-6b4a6e02516f",
			"previousElementUuid": previousElementUuid,
		}).Error(previousElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(previousElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return "", err
	}
	previousElementType := previousElement.MatureTestCaseModelElementMessage.TestCaseModelElementType

	// Extract data for Next Element
	nextElementUuid := currentTestCase.TestCaseModelMap[currentElementUuid].MatureTestCaseModelElementMessage.NextElementUuid
	nextElement, existInMap := currentTestCase.TestCaseModelMap[nextElementUuid]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":              "bf2ac32f-5edb-472f-af73-87d04400e132",
			"nextElementUuid": nextElementUuid,
		}).Error(nextElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(nextElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return "", err
	}
	nextlementType := nextElement.MatureTestCaseModelElementMessage.TestCaseModelElementType

	// TCRuleSwap101
	//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
	//	n=TIC(X)			B0					n 		B0								B1-n-B1					TCRuleSwap101
	if ruleCanBeProcessed == false &&
		(elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER ||
			elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE) &&
		previousElementUuid == currentElementUuid &&
		currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND &&
		nextElementUuid == currentElementUuid {

		// Rule OK
		ruleName = "TCRuleSwap101"
		ruleCanBeProcessed = true

	}

	// TCRuleSwap102
	//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
	//	n=TIC or TIC(X)		B10					n		TIC(B10)						TIC(B11f-n-B11l)		TCRuleSwap102
	if ruleCanBeProcessed == false &&
		(elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE ||
			elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER ||
			elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE) &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND &&
		currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND {

		// Rule OK
		ruleName = "TCRuleSwap102"
		ruleCanBeProcessed = true

	}

	// TCRuleSwap103
	//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
	//	n=TIC or TIC(X)		B11f				n		TIC(B11f-X)						TIC(B11f-n-B12-X)		TCRuleSwap103
	if ruleCanBeProcessed == false &&
		(elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE ||
			elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER ||
			elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE) &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND &&
		currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND &&
		(nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE ||
			nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER ||
			nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE) {
		// Rule OK
		ruleName = "TCRuleSwap103"
		ruleCanBeProcessed = true

	}

	// TCRuleSwap104
	//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
	//	n=TIC or TIC(X)		B11l				n		TIC(X-B11l)						TIC(X-B12-n-B11l)		TCRuleSwap104
	if ruleCanBeProcessed == false &&
		(elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE ||
			elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER ||
			elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE) &&
		(previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE ||
			previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER ||
			previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE) &&
		currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND {

		// Rule OK
		ruleName = "TCRuleSwap104"
		ruleCanBeProcessed = true

	}

	// TCRuleSwap105
	//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
	//	n=TIC or TIC(X)		B12					n		X-B12-X							X-B12-n-B12-X			TCRuleSwap105
	if ruleCanBeProcessed == false &&
		(elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE ||
			elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER ||
			elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE) &&
		(previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE ||
			previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER ||
			previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE) &&
		currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND &&
		(nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE ||
			nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER ||
			nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE) {

		// Rule OK
		ruleName = "TCRuleSwap105"
		ruleCanBeProcessed = true

	}

	// TCRuleSwap106
	//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
	//	n=TIC or TIC(X)		B10x*				n		TIC(B10*x*)						TIC(B11x-n-B11x)		TCRuleSwap106
	if ruleCanBeProcessed == false &&
		(elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE ||
			elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER ||
			elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE) &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10oxo_BOND &&
		currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10oxo_BOND &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10oxo_BOND {

		// Rule OK
		ruleName = "TCRuleSwap106"
		ruleCanBeProcessed = true

	}

	// TCRuleSwap107
	//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
	//	n=TIC or TIC(X)		B10*x				n		TIC(B10*x)						TIC(B11x-n-B11)			TCRuleSwap107
	if ruleCanBeProcessed == false &&
		(elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE ||
			elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER ||
			elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE) &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10ox_BOND &&
		currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10ox_BOND &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10ox_BOND {

		// Rule OK
		ruleName = "TCRuleSwap107"
		ruleCanBeProcessed = true

	}

	// TCRuleSwap108
	//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
	//	n=TIC or TIC(X)		B10x*				n		TIC(B10x*)						TIC(B11-n-B11x)			TCRuleSwap108
	if ruleCanBeProcessed == false &&
		(elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE ||
			elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER ||
			elementTypeToBeSwappedIn == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE) &&
		previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10xo_BOND &&
		currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10xo_BOND &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10xo_BOND {

		// Rule OK
		ruleName = "TCRuleSwap108"
		ruleCanBeProcessed = true

	}

	switch ruleName {

	case "TCRuleSwap101", "TCRuleSwap102", "TCRuleSwap103", "TCRuleSwap104",
		"TCRuleSwap105", "TCRuleSwap106", "TCRuleSwap107", "TCRuleSwap108":

		// Swap can be execute by complex Swap rules
		err = nil

	default:
		// The criteria for Deleting is not met by complex Swap rule
		err = errors.New("The criteria for any complex Swap rule was not met")

		commandAndRuleEngine.logger.WithFields(logrus.Fields{
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
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) verifyThatThereAreNoZombieElementsInComponent(immatureElement testCaseModel.ImmatureElementStruct) (err error) {

	var allUuidKeys []string

	// Extract all elements by key from component
	for _, elementKey := range immatureElement.ImmatureElementMap {
		allUuidKeys = append(allUuidKeys, elementKey.ImmatureElementUuid)
	}

	// Follow the path from "first element and remove the found element from 'allUuidKeys'
	allUuidKeys, err = commandAndRuleEngine.recursiveZombieElementSearchInComponentModel(immatureElement.FirstElementUuid, allUuidKeys, &immatureElement)

	// If there are elements left in slice then there were zombie elements, which there shouldn't be
	if len(allUuidKeys) != 0 {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":                               "3e519b3c-367d-42d5-a8ce-ff507efd8972",
			"allUuidKeys":                      allUuidKeys,
			"Number of Zombie Elements":        len(allUuidKeys),
			"immatureElement.FirstElementUuid": immatureElement.FirstElementUuid,
		}).Error("There existed Zombie elements in 'immatureElement.ImmatureElementMap'")

		err = errors.New("there existed Zombie elements in 'immatureElement.ImmatureElementMap', for " + immatureElement.FirstElementUuid)

		return err
	}

	return err
}

// Verify all children, in ImmatureEleemnt-model and remove the found element from 'allUuidKeys'
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) recursiveZombieElementSearchInComponentModel(elementsUuid string, allUuidKeys []string, immatureElement *testCaseModel.ImmatureElementStruct) (processedAllUuidKeys []string, err error) {

	// Extract current element
	currentElement, existInMap := immatureElement.ImmatureElementMap[elementsUuid]

	// If the element doesn't exit then there is something really wrong
	if existInMap == false {
		// This shouldn't happen
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":           "9f628356-2ea2-48a6-8e6a-546a5f97f05b",
			"elementsUuid": elementsUuid,
		}).Error(elementsUuid + " could not be found in in map 'immatureElement.ImmatureElementMap'")

		err = errors.New(elementsUuid + " could not be found in in map 'immatureElement.ImmatureElementMap'")

		return nil, err
	}

	// Element has child-element then go that path
	if currentElement.FirstChildElementUuid != elementsUuid {
		allUuidKeys, err = commandAndRuleEngine.recursiveZombieElementSearchInComponentModel(currentElement.FirstChildElementUuid, allUuidKeys, immatureElement)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return nil, err
	}

	// If element has a next-element the go that path
	if currentElement.NextElementUuid != elementsUuid {
		allUuidKeys, err = commandAndRuleEngine.recursiveZombieElementSearchInComponentModel(currentElement.NextElementUuid, allUuidKeys, immatureElement)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return nil, err
	}

	// Remove current element from "slice of all elements in map"
	tempAallUuidKeys := FindElementInSliceAndRemove(&allUuidKeys, elementsUuid)

	processedAllUuidKeys = *tempAallUuidKeys

	return processedAllUuidKeys, err
}

// Remove 'uuid' from slice
func FindElementInSliceAndRemove(sliceToWorkOn *[]string, uuid string) (returnSlice *[]string) {

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
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) verifyThatAllUuidsAreCorrectInComponent(immatureElement testCaseModel.ImmatureElementStruct) (err error) {

	// Loop all fields and find the ones defined as 'String'. Verify that content is a UUID
	e := reflect.ValueOf(&immatureElement.ImmatureElementMap).Elem()

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
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) recursiveVerifyAllUuidOfChildElements(testCaseUuid string, elementsUuid string) (err error) {

	// Get current TestCase
	currentTestCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all Testcases")
		return err
	}

	// Extract current element
	currentElement, existInMap := currentTestCase.TestCaseModelMap[elementsUuid]

	// If the element doesn't exit then there is something really wrong
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":           "9eae5791-88f0-481d-a7d9-21123b9eadfe",
			"elementsUuid": elementsUuid,
		}).Error(elementsUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(elementsUuid + " could not be found in in map 'immatureElement.ImmatureElementMap'")

		return err
	}

	// Element has child-element then go that path
	if currentElement.MatureTestCaseModelElementMessage.FirstChildElementUuid != elementsUuid {
		err = commandAndRuleEngine.recursiveDeleteOfChildElements(&currentTestCase, currentElement.MatureTestCaseModelElementMessage.FirstChildElementUuid)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return err
	}

	// If element has a next-element the go that path
	if currentElement.MatureTestCaseModelElementMessage.NextElementUuid != elementsUuid {
		err = commandAndRuleEngine.recursiveDeleteOfChildElements(&currentTestCase, currentElement.MatureTestCaseModelElementMessage.NextElementUuid)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return err
	}

	// Remove current element from Map
	delete(currentTestCase.TestCaseModelMap, elementsUuid)

	return err
}
