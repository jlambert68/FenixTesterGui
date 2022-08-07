package commandAndRuleEngine

import (
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeCutFullELementStructure(testCaseUuid string, uuidToBeCutOut string) (err error) {

	var tempTestCase *testCaseModel.TestCaseModelStruct

	tempTestCaseModelMap := make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage)

	// Get current TestCase
	currentTestCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]

	if existsInMap == false {
		errorId := "9ea79cce-e892-4e30-bbd5-7a7e13a1ff35"
		err = errors.New(fmt.Sprintf("testcase with uuid '%s' doesn't exist in map with all Testcases [ErrorID: %s]", testCaseUuid, errorId))

		return err
	}

	// Transform a copy of current TestCase to 'tempTestCase'
	for elemenUuid, tempElement := range currentTestCase.TestCaseModelMap {
		tempTestCaseModelMap[elemenUuid] = tempElement
	}

	tempTestCase = &testCaseModel.TestCaseModelStruct{
		FirstElementUuid: currentTestCase.FirstElementUuid,
		TestCaseModelMap: tempTestCaseModelMap,
	}

	currentElement, existInMap := tempTestCase.TestCaseModelMap[uuidToBeCutOut]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":             "1cf17c6b-d6bf-4045-be7d-07e03be25f6d",
			"uuidToBeCutOut": uuidToBeCutOut,
		}).Error(uuidToBeCutOut + " could not be found in in map 'TestCaseModelMap'")

		errorId := "ff4a18d8-b472-4673-a899-cb98217dab68"
		err = errors.New(fmt.Sprintf("'%s' could not be found in in currents TestCase-map [ErrorID: %s]", uuidToBeCutOut, errorId))

		return err
	}

	// Extract data for Previous Element
	currentElementUuid := uuidToBeCutOut

	// Remove references in currentElement to Previous- and Next- Elements
	currentElement.PreviousElementUuid = currentElement.MatureElementUuid
	currentElement.NextElementUuid = currentElement.MatureElementUuid

	// Save updated back into tempMap
	tempTestCase.TestCaseModelMap[currentElement.MatureElementUuid] = currentElement

	// Set up structure for copied element structure
	copiedStructure := testCaseModel.MatureElementStruct{
		FirstElementUuid: currentElementUuid,
		MatureElementMap: nil,
	}
	// Initiate map for elements to be copied
	copiedStructure.MatureElementMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage)

	// Make the copying of current element and its children, if they exist
	err = commandAndRuleEngine.recursiveCuttingOfFullElementStructure(tempTestCase, currentElementUuid, &copiedStructure)
	if err != nil {

		errorId := "4791e1e3-af61-4894-bc5d-ec7d0fef8d7b"
		err = errors.New(fmt.Sprintf("something went wrong when creating element structure from element '%s' and its children, in 'TestCaseModelMap', using recursive calls [ErrorID: %s]", currentElement.MatureElementUuid, errorId))

		return err
	}

	//

	// Remove the original structure from the TestCase
	err = commandAndRuleEngine.executeDeleteElement(testCaseUuid, currentElementUuid)

	if err != nil {

		return err
	}

	//Reload the TestCase
	currentTestCase, existsInMap = commandAndRuleEngine.Testcases.TestCases[testCaseUuid]

	if existsInMap == false {
		errorId := "9f8fe113-6980-4ad5-8ea6-ca9d56722145"
		err = errors.New(fmt.Sprintf("testcase with uuid '%s' doesn't exist in map with all Testcases [ErrorID: %s]", testCaseUuid, errorId))

		return err
	}

	// If there are no errors then save the copied Element Structure in Copy-buffer and then save the Updated TestCase
	// Save Copied element to Cut Buffer  in TestCase
	currentTestCase.CutBuffer = copiedStructure

	// Save TestCase
	commandAndRuleEngine.Testcases.TestCases[testCaseUuid] = currentTestCase

	return err

}

// Copy the full structure of all children, in TestCase-model, for specific Element
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) recursiveCuttingOfFullElementStructure(currentTestCase *testCaseModel.TestCaseModelStruct, elementsUuid string, copiedElementStructure *testCaseModel.MatureElementStruct) (err error) {

	// Extract current element
	currentElement, existInMap := currentTestCase.TestCaseModelMap[elementsUuid]

	// If the element doesn't exit then there is something really wrong
	if existInMap == false {
		errorId := "d06d753b-01dc-471b-8c97-c95e0e4562ae"
		err = errors.New(fmt.Sprintf("'%s' could not be found in in currents TestCase-map [ErrorID: %s]", elementsUuid, errorId))

		return err

	}

	// Element has child-element then go that path
	if currentElement.FirstChildElementUuid != elementsUuid {
		err = commandAndRuleEngine.recursiveCuttingOfFullElementStructure(currentTestCase, currentElement.FirstChildElementUuid, copiedElementStructure)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return err
	}

	// If element has a next-element the go that path
	if currentElement.NextElementUuid != elementsUuid {
		err = commandAndRuleEngine.recursiveCuttingOfFullElementStructure(currentTestCase, currentElement.NextElementUuid, copiedElementStructure)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return err
	}

	// Copy the element
	newCopiedElement := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      currentElement.OriginalElementUuid,
		OriginalElementName:      currentElement.OriginalElementName,
		MatureElementUuid:        currentElement.MatureElementUuid,
		PreviousElementUuid:      currentElement.PreviousElementUuid,
		NextElementUuid:          currentElement.NextElementUuid,
		FirstChildElementUuid:    currentElement.FirstChildElementUuid,
		ParentElementUuid:        currentElement.ParentElementUuid,
		TestCaseModelElementType: currentElement.TestCaseModelElementType,
	}

	// Add the element to the referenced map
	copiedElementStructure.MatureElementMap[currentElement.MatureElementUuid] = newCopiedElement

	return err
}
