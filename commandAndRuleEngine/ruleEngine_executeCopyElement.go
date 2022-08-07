package commandAndRuleEngine

import (
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeCopyFullELementStructure(testCaseUuid string, uuidToCopy string) (err error) {

	var tempTestCase *testCaseModel.TestCaseModelStruct

	tempTestCaseModelMap := make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage)

	// Get current TestCase
	currentTestCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]

	if existsInMap == false {
		errorId := "64bb031b-88c7-4758-aade-7375816ac285"
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

	currentElement, existInMap := tempTestCase.TestCaseModelMap[uuidToCopy]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":         "b69296dd-7b73-46c9-b465-a8fb40a9a592",
			"uuidToCopy": uuidToCopy,
		}).Error(uuidToCopy + " could not be found in in map 'TestCaseModelMap'")

		errorId := "f4f32c48-a840-4fbc-807d-19bd4d4960d9"
		err = errors.New(fmt.Sprintf("'%s' could not be found in in currents TestCase-map [ErrorID: %s]", uuidToCopy, errorId))

		return err
	}

	// Extract data for Previous Element
	currentElementUuid := uuidToCopy

	// Remove references in currentElement to Previous- , Next- and Parent Elements
	currentElement.PreviousElementUuid = currentElement.MatureElementUuid
	currentElement.NextElementUuid = currentElement.MatureElementUuid
	currentElement.ParentElementUuid = currentElement.MatureElementUuid

	// Save updated element back into tempMap
	tempTestCase.TestCaseModelMap[currentElementUuid] = currentElement

	// Set up structure for copied element structure
	copiedStructure := testCaseModel.ImmatureElementStruct{
		FirstElementUuid:   currentElementUuid,
		ImmatureElementMap: nil,
	}
	// Initiate map for elements to be copied
	copiedStructure.ImmatureElementMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage)

	// Make the copying of current element and its children, if they exist
	err = commandAndRuleEngine.recursiveCopyingOfFullElementStructure(tempTestCase, currentElementUuid, &copiedStructure)
	if err != nil {

		errorId := "c2d1e439-ef76-4486-a7f5-3c7a54dc156c"
		err = errors.New(fmt.Sprintf("something went wrong when creating element structure from element '%s' and its children, in 'TestCaseModelMap', using recursive calls [ErrorID: %s]", currentElement.MatureElementUuid, errorId))

		return err
	}

	//Reload the TestCase - not needed
	// currentTestCase, existsInMap = commandAndRuleEngine.Testcases.TestCases[testCaseUuid]

	if existsInMap == false {
		errorId := "9d857471-7918-4762-be9b-729b82a961e2"
		err = errors.New(fmt.Sprintf("testcase with uuid '%s' doesn't exist in map with all Testcases [ErrorID: %s]", testCaseUuid, errorId))

		return err
	}

	// If there are no errors then save the copied Element Structure in Copy-buffer and then save the Updaed TestCase
	// Save Copy Buffer in TestCase
	currentTestCase.CopyBuffer = copiedStructure

	// Save TestCase
	commandAndRuleEngine.Testcases.TestCases[testCaseUuid] = currentTestCase

	return err

}

// Copy the full structure of all children, in TestCase-model, for specific Element
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) recursiveCopyingOfFullElementStructure(currentTestCase *testCaseModel.TestCaseModelStruct, elementsUuid string, copiedElementStructure *testCaseModel.ImmatureElementStruct) (err error) {

	// Extract current element
	currentElement, existInMap := currentTestCase.TestCaseModelMap[elementsUuid]

	// If the element doesn't exit then there is something really wrong
	if existInMap == false {
		errorId := "10bf1786-efc0-49f0-8434-232431c6072f"
		err = errors.New(fmt.Sprintf("'%s' could not be found in in currents TestCase-map [ErrorID: %s]", elementsUuid, errorId))

		return err

	}

	// Element has child-element then go that path
	if currentElement.FirstChildElementUuid != elementsUuid {
		err = commandAndRuleEngine.recursiveCopyingOfFullElementStructure(currentTestCase, currentElement.FirstChildElementUuid, copiedElementStructure)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return err
	}

	// If element has a next-element the go that path
	if currentElement.NextElementUuid != elementsUuid {
		err = commandAndRuleEngine.recursiveCopyingOfFullElementStructure(currentTestCase, currentElement.NextElementUuid, copiedElementStructure)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return err
	}

	// Copy the element
	newCopiedElement := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage{
		OriginalElementUuid:      currentElement.OriginalElementUuid,
		OriginalElementName:      currentElement.OriginalElementName,
		ImmatureElementUuid:      currentElement.MatureElementUuid,
		PreviousElementUuid:      currentElement.PreviousElementUuid,
		NextElementUuid:          currentElement.NextElementUuid,
		FirstChildElementUuid:    currentElement.FirstChildElementUuid,
		ParentElementUuid:        currentElement.ParentElementUuid,
		TestCaseModelElementType: currentElement.TestCaseModelElementType,
	}

	// Add the element to the referenced map
	copiedElementStructure.ImmatureElementMap[currentElement.MatureElementUuid] = newCopiedElement

	return err
}
