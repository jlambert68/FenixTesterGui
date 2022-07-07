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
	parentElementUuid := currentElement.ParentElementUuid

	// Create the structure after Delete

	// Create new B0-bond element
	newB0BondElement := commandAndRuleEngineObject.createNewBondB0Element(parentElementUuid)

	// Add New Elements to Map
	commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[newB0BondElement.MatureElementUuid] = newB0BondElement

	// Remove Old Elements from Map
	delete(commandAndRuleEngineObject.testcaseModel.TestCaseModelMap, previousElementUuid)
	delete(commandAndRuleEngineObject.testcaseModel.TestCaseModelMap, nextElementUuid)

	// Remove current element and children, if they exist, from map
	err = commandAndRuleEngineObject.recursiveDeleteOfChildElements(currentElementUuid)
	if err != nil {
		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":  "99abee1e-3078-42e7-a08a-a719e0a4ed8d",
			"err": err,
		}).Error(" Something went wrong when deleting children, in 'TestCaseModelMap', using recusive calls")

		err = errors.New("something went wrong when deleting children, in 'TestCaseModelMap', using recusive calls")

		return err
	}

	// Update Reference in parent TIC
	tempParentElement, existInMap := commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[parentElementUuid]
	if existInMap == false {
		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":                "5d0bca9d-86f7-448a-82d7-e0fc1a02a966",
			"parentElementUuid": parentElementUuid,
		}).Error(parentElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(parentElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	tempParentElement.FirstChildElementUuid = newB0BondElement.MatureElementUuid

	// Add updated parent-element back into TestCaseModelMap
	commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[parentElementUuid] = tempParentElement

	return err

}

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			B11f-n-B11l						B10							TCRuleDeletion102
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) executeTCRuleDeletion102(uuidToDelete string) (err error) {
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
			"id":           "e2773f12-3bf6-4ebe-856b-25e6bb51864f",
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
	parentElementUuid := currentElement.ParentElementUuid

	// Create the structure after Delete

	// Create new B10-bond element
	newB10BondElement := commandAndRuleEngineObject.createNewBondB10Element(parentElementUuid)

	// Add New Elements to Map
	commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[newB10BondElement.MatureElementUuid] = newB10BondElement

	// Remove Old Elements from Map
	delete(commandAndRuleEngineObject.testcaseModel.TestCaseModelMap, previousElementUuid)
	delete(commandAndRuleEngineObject.testcaseModel.TestCaseModelMap, nextElementUuid)

	// Remove current element and children, if they exist, from map
	err = commandAndRuleEngineObject.recursiveDeleteOfChildElements(currentElementUuid)
	if err != nil {
		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":  "43f9cfd5-5cfe-4ec0-b9e9-20271ab868e4",
			"err": err,
		}).Error(" Something went wrong when deleting children, in 'TestCaseModelMap', using recusive calls")

		err = errors.New("something went wrong when deleting children, in 'TestCaseModelMap', using recusive calls")

		return err
	}

	// Update Reference in parent TIC
	tempParentElement, existInMap := commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[parentElementUuid]
	if existInMap == false {
		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":                "b75f6401-0393-411d-bcb6-28005c03ac9d",
			"parentElementUuid": parentElementUuid,
		}).Error(parentElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(parentElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	tempParentElement.FirstChildElementUuid = newB10BondElement.MatureElementUuid

	// Add updated parent-element back into TestCaseModelMap
	commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[parentElementUuid] = tempParentElement

	return err

}

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			B11fx-n-B11lx					B10*x*						TCRuleDeletion103
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) executeTCRuleDeletion103(uuidToDelete string) (err error) {
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
			"id":           "83137d58-dd37-443c-bf1f-0b01b7a85a8b",
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
	parentElementUuid := currentElement.ParentElementUuid

	// Create the structure after Delete

	// Create new B10-bond element
	newB10oxoBondElement := commandAndRuleEngineObject.createNewBondB10oxoElement(parentElementUuid)

	// Add New Elements to Map
	commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[newB10oxoBondElement.MatureElementUuid] = newB10oxoBondElement

	// Remove Old Elements from Map
	delete(commandAndRuleEngineObject.testcaseModel.TestCaseModelMap, previousElementUuid)
	delete(commandAndRuleEngineObject.testcaseModel.TestCaseModelMap, nextElementUuid)

	// Remove current element and children, if they exist, from map
	err = commandAndRuleEngineObject.recursiveDeleteOfChildElements(currentElementUuid)
	if err != nil {
		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":  "f2683652-02ef-4260-8cb3-15cf627ddfa9",
			"err": err,
		}).Error(" Something went wrong when deleting children, in 'TestCaseModelMap', using recusive calls")

		err = errors.New("something went wrong when deleting children, in 'TestCaseModelMap', using recusive calls")

		return err
	}

	// Update Reference in parent TIC
	tempParentElement, existInMap := commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[parentElementUuid]
	if existInMap == false {
		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":                "68faa4b5-bf93-469e-b2b4-fe2dfb192650",
			"parentElementUuid": parentElementUuid,
		}).Error(parentElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(parentElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	tempParentElement.FirstChildElementUuid = newB10oxoBondElement.MatureElementUuid

	// Add updated parent-element back into TestCaseModelMap
	commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[parentElementUuid] = tempParentElement

	return err

}

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			B11f-n-B11lx					B10x*						TCRuleDeletion104
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) executeTCRuleDeletion104(uuidToDelete string) (err error) {
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
			"id":           "e7cf67a0-c8b9-44c6-920c-b0ef5d899d99",
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
	parentElementUuid := currentElement.ParentElementUuid

	// Create the structure after Delete

	// Create new B10-bond element
	newB10xoBondElement := commandAndRuleEngineObject.createNewBondB10xoElement(parentElementUuid)

	// Add New Elements to Map
	commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[newB10xoBondElement.MatureElementUuid] = newB10xoBondElement

	// Remove Old Elements from Map
	delete(commandAndRuleEngineObject.testcaseModel.TestCaseModelMap, previousElementUuid)
	delete(commandAndRuleEngineObject.testcaseModel.TestCaseModelMap, nextElementUuid)

	// Remove current element and children, if they exist, from map
	err = commandAndRuleEngineObject.recursiveDeleteOfChildElements(currentElementUuid)
	if err != nil {
		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":  "b464d0dc-86b6-405b-802d-b538a6c2c840",
			"err": err,
		}).Error(" Something went wrong when deleting children, in 'TestCaseModelMap', using recusive calls")

		err = errors.New("something went wrong when deleting children, in 'TestCaseModelMap', using recusive calls")

		return err
	}

	// Update Reference in parent TIC
	tempParentElement, existInMap := commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[parentElementUuid]
	if existInMap == false {
		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":                "6d7eae16-7b13-4f3e-9cbe-cc564307d86c",
			"parentElementUuid": parentElementUuid,
		}).Error(parentElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(parentElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	tempParentElement.FirstChildElementUuid = newB10xoBondElement.MatureElementUuid

	// Add updated parent-element back into TestCaseModelMap
	commandAndRuleEngineObject.testcaseModel.TestCaseModelMap[parentElementUuid] = tempParentElement

	return err

}

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			B11fx-n-B11l					B10*x						TCRuleDeletion105

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			B11f-n-B12-X					B11f-X						TCRuleDeletion106

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			B11fx-n-B12x-X					B11fx-X						TCRuleDeletion107

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			B11f-n-B12x-X					B11fx-X						TCRuleDeletion108

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			B11fx-n-B12-X					B11fx-X						TCRuleDeletion109

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-B12-n-B11l					X-B11l						TCRuleDeletion110

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-B12x-n-B11lx					X-B11lx						TCRuleDeletion111

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-B12-n-B11lx					X-B11lx						TCRuleDeletion112

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-B12x-n-B11l					X-B11lx						TCRuleDeletion113

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-B12-n-B12-X					X-B12-X						TCRuleDeletion114

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-B12x-n-B12x-X					X-B12x-X					TCRuleDeletion115

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-B12-n-B12x-X					X-B12x-X					TCRuleDeletion116

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-B12x-n-B12-X					X-B12x-X					TCRuleDeletion117

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
