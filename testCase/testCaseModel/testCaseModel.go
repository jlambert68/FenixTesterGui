package testCaseModel

import (
	"errors"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
)

// Verify all children, in TestCaseElement-model and remove the found element from 'allUuidKeys'
func (testCaseModel *TestCasesModelsStruct) recursiveZombieElementSearchInTestCaseModel(testCaseUuid string, elementsUuid string, allUuidKeys []string) (processedAllUuidKeys []string, err error) {

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
	tempAallUuidKeys := findElementInSliceAndRemove(&allUuidKeys, elementsUuid)

	processedAllUuidKeys = *tempAallUuidKeys

	return processedAllUuidKeys, err
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

// Generate the slice with the elements in the TestCase. Order is the same as in the Textual Representation of the TestCase
func (testCaseModel *TestCasesModelsStruct) recursiveTextualTestCaseModelExtractor(testCaseUuid string, elementsUuid string, testCaseModelElementsIn []fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage) (testCaseModelElementsIOut []fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage, err error) {

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

	// Add element to slice
	testCaseModelElementsIOut = append(testCaseModelElementsIn, currentElement)

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

// Generate name to be used when presenting TestCase Element
func (testCaseModel *TestCasesModelsStruct) generateUINameForTestCaseElement(element *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage) (elementUiName string) {

	elementUiName = element.OriginalElementName + " [" + element.MatureElementUuid[0:numberOfCharactersfromUuid-1] + "]"

	return elementUiName
}

// Generate the slice with the elements in the TestCase. Order is the same as in the Textual Representation of the TestCase
func (testCaseModel *TestCasesModelsStruct) recursiveGraphicalTestCaseTreeModelExtractor(testCaseUuid string, currentElementsUuid string, treeViewNodeChildrenIn []string) (err error) {

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {
		errorId := "68f37aee-0b93-4d4f-9225-31ea5ccd8f8a"
		err = errors.New(fmt.Sprintf("testcase with uuid '%s' doesn't exist in map with all testcases [ErrorID: %s]", testCaseUuid, errorId))

		return err
	}

	// Extract current elementnil
	currentElement, existInMap := currentTestCase.TestCaseModelMap[currentElementsUuid]

	// If the element doesn't exit then there is something really wrong
	if existInMap == false {
		// This shouldn't happen
		errorId := "2397cd98-e5dd-4a42-9d7a-46de3f0286d2"
		err = errors.New(fmt.Sprintf("element '%s', in testcase '%s' doesn't exist in map with all testcases [ErrorID: %s]", currentElementsUuid, testCaseUuid, errorId))

		return err
	}

	// Add element to slice
	treeViewNodeChildrenIn = append(treeViewNodeChildrenIn, currentElementsUuid)

	// Element has child-element then go that path
	if currentElement.FirstChildElementUuid != currentElementsUuid {
		err = testCaseModel.recursiveGraphicalTestCaseTreeModelExtractor(testCaseUuid, currentElement.FirstChildElementUuid, []string{})
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return err
	}

	// If element has a next-element the go that path
	if currentElement.NextElementUuid != currentElementsUuid {
		err = testCaseModel.recursiveGraphicalTestCaseTreeModelExtractor(testCaseUuid, currentElement.NextElementUuid, treeViewNodeChildrenIn)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return err
	}

	// Add Children to Map with currenElementUuid as map-key
	currentTestCase.testCaseModelAdaptedForUiTree[currentElementsUuid] = treeViewNodeChildrenIn

	return err
}
