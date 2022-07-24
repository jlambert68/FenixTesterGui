package testCaseModel

import (
	"FenixTesterGui/commandAndRuleEngine"
	"errors"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"strings"
)

// Verify that all UUIDs are correct in TestCaseModel. Meaning that no empty uuid is allowed and they all are correct
func (testCaseModel *TestCaseModelsStruct) VerifyThatThereAreNoZombieElementsInTestCaseModel(testCaseUuid string) (err error) {

	var allUuidKeys []string

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all testcases")
		return err
	}

	// Extract all elements by key from TestCaseModel
	for _, elementKey := range currentTestCase.TestCaseModelMap {
		allUuidKeys = append(allUuidKeys, elementKey.MatureElementUuid)
	}

	// Follow the path from "first element and remove the found element from 'allUuidKeys'
	allUuidKeys, err = testCaseModel.recursiveZombieElementSearchInTestCaseModel(testCaseUuid, currentTestCase.FirstElementUuid, allUuidKeys)

	// If there are elements left in slice then there were zombie elements, which there shouldn't be
	if len(allUuidKeys) != 0 {
		err = errors.New("there existed Zombie elements in 'testCaseModel.TestCaseModelMap', for " + currentTestCase.FirstElementUuid)

		return err
	}

	return err
}

// Verify all children, in TestCaseElement-model and remove the found element from 'allUuidKeys'
func (testCaseModel *TestCaseModelsStruct) recursiveZombieElementSearchInTestCaseModel(testCaseUuid string, elementsUuid string, allUuidKeys []string) (processedAllUuidKeys []string, err error) {

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all testcases")
		return nil, err
	}

	// Extract current element
	currentElement, existInMap := currentTestCase.TestCaseModelMap[elementsUuid]

	// If the element doesn't exit then there is something really wrong
	if existInMap == false {
		// This shouldn't happen
		err = errors.New(elementsUuid + " could not be found in in map 'testCaseModel.TestCaseModelMap'")

		return nil, err
	}

	// Element has child-element then go that path
	if currentElement.FirstChildElementUuid != elementsUuid {
		allUuidKeys, err = testCaseModel.recursiveZombieElementSearchInTestCaseModel(testCaseUuid, currentElement.FirstChildElementUuid, allUuidKeys)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return nil, err
	}

	// If element has a next-element the go that path
	if currentElement.NextElementUuid != elementsUuid {
		allUuidKeys, err = testCaseModel.recursiveZombieElementSearchInTestCaseModel(testCaseUuid, currentElement.NextElementUuid, allUuidKeys)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return nil, err
	}

	// Remove current element from "slice of all elements in map"
	tempAallUuidKeys := commandAndRuleEngine.FindElementInSliceAndRemove(&allUuidKeys, elementsUuid)

	processedAllUuidKeys = *tempAallUuidKeys

	return processedAllUuidKeys, err
}

func (testCaseModel *TestCaseModelsStruct) CreateTextualTestCase(testCaseUuid string) (textualTestCaseSimple string, textualTestCaseComplex string, err error) {

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all testcases")
		return "", "", err
	}

	// Create slice with all elementTypes in presentation order
	testCaseModelElements, err := testCaseModel.recursiveTextualTestCaseModelExtractor(testCaseUuid, currentTestCase.FirstElementUuid, []fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum{})

	// Something wrong happen
	if err != nil {
		return "", "", err
	}

	// If there are no elements in TestCaseModel then return empty Textual description
	if len(testCaseModelElements) == 0 {
		return "[]", "[]", nil
	}

	// Loop all elements and convert element type into presentation representation
	for _, testCaseModelElementType := range testCaseModelElements {

		// Simple presentation style, like 'B10x' for "B10oxo"
		presentationNameSimple := fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementPresentationNameEnum_name[int32(testCaseModelElementType)]
		separatorIndexSimple := strings.Index(presentationNameSimple, "_")
		presentationNameSimple = presentationNameSimple[:separatorIndexSimple]

		// Complex presentation style, like 'B10oxo' for "B10oxo"
		presentationNameComplex := fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_name[int32(testCaseModelElementType)]
		separatorIndexComplex := strings.Index(presentationNameComplex, "_")
		presentationNameComplex = presentationNameComplex[:separatorIndexComplex]

		switch testCaseModelElementType {

		// First element in TC
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE:

			presentationNameSimple = "[" + presentationNameSimple
			presentationNameComplex = "[" + presentationNameComplex

			// Last element in TC
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE:

			presentationNameSimple = "-" + presentationNameSimple + "]"
			presentationNameComplex = "-" + presentationNameComplex + "]"

			// The only element in TC
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND:

			presentationNameSimple = "[" + presentationNameSimple + "]"
			presentationNameComplex = "[" + presentationNameComplex + "]"

		// First element child in TIC(x)
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE:

			presentationNameSimple = "(" + presentationNameSimple
			presentationNameComplex = "(" + presentationNameComplex

		// Last element child in TIC(x)
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE:

			presentationNameSimple = "-" + presentationNameSimple + ")"
			presentationNameComplex = "-" + presentationNameComplex + ")"

			// The only element in TIC
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10ox_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10oxo_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10xo_BOND:

			presentationNameSimple = "(" + presentationNameSimple + ")"
			presentationNameComplex = "(" + presentationNameComplex + ")"

			// Element surrounded with other elements
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE:

			presentationNameSimple = "-" + presentationNameSimple
			presentationNameComplex = "-" + presentationNameComplex

			// No match in element
		default:

			err = errors.New("no match in element type for: " + testCaseModelElementType.String())
			return "", "", err

		}

		// Add presentation name to Textual TestCase
		textualTestCaseSimple = textualTestCaseSimple + presentationNameSimple
		textualTestCaseComplex = textualTestCaseComplex + presentationNameComplex

	}

	return textualTestCaseSimple, textualTestCaseComplex, err
}

// Verify all children, in TestCaseElement-model and remove the found element from 'allUuidKeys'
func (testCaseModel *TestCaseModelsStruct) recursiveTextualTestCaseModelExtractor(testCaseUuid string, elementsUuid string, testCaseModelElementsIn []fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum) (testCaseModelElementsIOut []fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum, err error) {

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all testcases")
		return nil, err
	}

	// Extract current element
	currentElement, existInMap := currentTestCase.TestCaseModelMap[elementsUuid]

	// If the element doesn't exit then there is something really wrong
	if existInMap == false {
		// This shouldn't happen
		err = errors.New(elementsUuid + " could not be found in in map 'testCaseModel.TestCaseModelMap'")

		return nil, err
	}

	// Add elementType to slice
	testCaseModelElementsIOut = append(testCaseModelElementsIn, currentElement.TestCaseModelElementType)

	// Element has child-element then go that path
	if currentElement.FirstChildElementUuid != elementsUuid {
		testCaseModelElementsIOut, err = testCaseModel.recursiveTextualTestCaseModelExtractor(testCaseUuid, currentElement.FirstChildElementUuid, testCaseModelElementsIOut)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return nil, err
	}

	// If element has a next-element the go that path
	if currentElement.NextElementUuid != elementsUuid {
		testCaseModelElementsIOut, err = testCaseModel.recursiveTextualTestCaseModelExtractor(testCaseUuid, currentElement.NextElementUuid, testCaseModelElementsIOut)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return nil, err
	}

	return testCaseModelElementsIOut, err
}
