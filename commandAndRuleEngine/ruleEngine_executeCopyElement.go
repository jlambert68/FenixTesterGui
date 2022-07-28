package commandAndRuleEngine

import (
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)
)


func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeCopyFullELementStructure(testCaseUuid string, uuidToCopy string) (err error) {

	// Get current TestCase
	currentTestCase, existsInMap := commandAndRuleEngine.testcases.TestCases[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all testcases")
		return err
	}

	currentElement, existInMap := currentTestCase.TestCaseModelMap[uuidToCopy]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":           "b69296dd-7b73-46c9-b465-a8fb40a9a592",
			"uuidToCopy": uuidToCopy,
		}).Error(uuidToCopy + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(uuidToCopy + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	// Extract data for Previous Element
	currentElementUuid := uuidToCopy


	// Remove references in currentElement to Previous- and Next- Elements
	currentElement.PreviousElementUuid = currentElement.MatureElementUuid
	currentElement.NextElementUuid = currentElement.MatureElementUuid


	// Save updated back into TestCase-map
	currentTestCase.TestCaseModelMap[currentElement.MatureElementUuid] = currentElement

	// Set up structure for copied element structure
	copiedStructure := testCaseModel.MatureElementStruct{
		FirstElementUuid: currentElementUuid,
		MatureElementMap: nil,
	}
	// Initiate map for elements to be copied
	copiedStructure.MatureElementMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage)

	// Make the copying of current element and its children, if they exist
	err = commandAndRuleEngine.recursiveCopyingOfFullElementStructure(&currentTestCase, currentElementUuid, &copiedStructure)
	if err != nil {

		errorId := "c2d1e439-ef76-4486-a7f5-3c7a54dc156c"
		err = errors.New(fmt.Sprintf("something went wrong when creating element structure from element '%s' and its children, in 'TestCaseModelMap', using recursive calls [ErrorID: %s]", currentElement.MatureElementUuid, errorId))

		return err
	}

	// If there are no errors then save the copied Element Structure in Copy-buffer
	if err == nil {
		//commandAndRuleEngine.testcases.TestCases[testCaseUuid] = currentTestCase
	}

	return err

}


// Copy the full structure of all children, in TestCase-model, for specific Element
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) recursiveCopyingOfFullElementStructure(currentTestCase *testCaseModel.TestCaseModelStruct, elementsUuid string, copiedElementStructure *testCaseModel.MatureElementStruct) (err error) {

	// Extract current element
	currentElement, existInMap := currentTestCase.TestCaseModelMap[elementsUuid]

	// If the element doesn't exit then there is something really wrong
	if existInMap == false {
		errorId := "10bf1786-efc0-49f0-8434-232431c6072f"
		err = errors.New(fmt.Sprintf("'%s' could not be found in in currents TestCase-map [ErrorID: %s]",elementsUuid, errorId))

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

	// Add the element to the reference map
	copiedElementStructure.MatureElementMap[currentElement.MatureElementUuid] = newCopiedElement

	return err
}
