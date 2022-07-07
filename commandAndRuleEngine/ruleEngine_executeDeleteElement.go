package commandAndRuleEngine

import (
	"errors"
	"github.com/sirupsen/logrus"
)

// What to remove			Remove in structure				Result after deletion		Rule
// n= TIC(X)				B1-n-B1							B0							TCRuleDeletion101
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) executeTCRuleDeletion101(uuidToDelete string) (err error) {
	/*
		// One TestCaseElement, that is not a TestInstructionContainer, that hold type, reference to previous, next and parent elements
		message TestCaseModelElementMessage {
		  string OriginalElementUuid = 1; // The original elements UUID, e.g. a TestInstruction unique UUID set by client system
		  string OriginalElementName = 2; // The original elements Name, e.g. a TestInstruction Name set by client system
		  string MatureElementUuid = 3; // The UUID that is created in the TestCase to give it a unique id
		  string PreviousElementUuid = 4;  // The UUID of the previous element. When there are no previous element then this field is populated with current element UUID
		  string NextElementUuid = 5;  // The UUID of the previous element. When there are no next element then this field is populated with current element UUID
		  string FirstChildElementUuid = 6; // The UUID of the first child element. Only applicable when this is a TestInstructionContainer. When there are no child element then this field is populated with current element UUID
		  string ParentElementUuid = 7; // The UUID of the parent, TestInstructionContainer. Only applicable when this is the last element inside a TestInstructionContainer. When there are no parent element then this field is populated with current element UUID
		  TestCaseModelElementTypeEnum TestCaseModelElementType = 8; // The specific type of this TestCase-element
		  string CurrentElementModelElement = 9; // The UUID of the element that this data act on, e.g. For TI & TIC the it is the same as 'OriginalElementUuid' but for BONDs then it is the BONDs UUID
		}
	*/

	currentElement, existInMap := commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[uuidToDelete]
	if existInMap == false {
		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":           "4c328a0d-aaa1-4820-8fa2-e3067456faff",
			"uuidToDelete": uuidToDelete,
		}).Error(uuidToDelete + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(uuidToDelete + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	// Extract data for Previous Elementfunc (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct)
	currentElementUuid := uuidToDelete

	// Extract data for Previous Element
	previousElementUuid := currentElement.PreviousElementUuid

	// Extract data for Next Element
	nextElementUuid := currentElement.NextElementUuid

	// Extract dta for Parent Element
	parrentElementUuid := currentElement.ParentElementUuid

	// Create the structure after Delete

	// Create new B0-bond element
	newB0Element := commandAndRuleEngineObject.createNewBondB0Element(parrentElementUuid)

	// Add New Elements to Map
	commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[newB0Element.MatureElementUuid] = newB0Element

	// Remove Old Elements from Map
	delete(commandAndRuleEngineObject.testcaseModel.TestCaseModelMap, previousElementUuid)
	delete(commandAndRuleEngineObject.testcaseModel.TestCaseModelMap, currentElementUuid)
	delete(commandAndRuleEngineObject.testcaseModel.TestCaseModelMap, nextElementUuid)

	// Remove all children from map
	err = commandAndRuleEngineObject.recursiveDeleteOfChildElements(currentElementUuid)
	if err != nil {
		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":  "99abee1e-3078-42e7-a08a-a719e0a4ed8d",
			"err": err,
		}).Error(" Something went wrong when deleting children, in 'TestCaseModelMap', using recusive calls")

		err = errors.New("Something went wrong when deleting children, in 'TestCaseModelMap', using recusive calls")

		return err
	}

	// Update Reference in parent TIC
	tempParrentElement := commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[parrentElementUuid]
	tempParrentElement.FirstChildElementUuid = newB0Element.MatureElementUuid
	commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[parrentElementUuid] = tempParrentElement

	return err

}

// Remove all children, in TestCase-model, for specific Element
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) recursiveDeleteOfChildElements(elementsUuid string) (err error) {

	// Extract current element
	currentElement, existInMap := commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[elementsUuid]

	// If the element doesn't exit then there is something really wrong
	if existInMap == false {
		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":           "9eae5791-88f0-481d-a7d9-21123b9eadfe",
			"elementsUuid": elementsUuid,
		}).Error(elementsUuid + " could not be found in in map 'TestCaseModelMap'")
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
