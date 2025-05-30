package commandAndRuleEngine

import (
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

// What to remove			Remove in structure				Result after deletion		Rule
// n= TIC(X)				B1-n-B1							B0							TCRuleDeletion101
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleDeletion101(testCaseUuid string, uuidToDelete string) (err error) {
	/*
		// One TestCaseElement, that is not a TestInstructionContainer, that hold type, reference to previous, next and parent elements
		message MatureTestCaseModelElementMessage {
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

	// Get current TestCase
	var currentTestCasePtr *testCaseModel.TestCaseModelStruct
	var existsInMap bool

	// Get the TestCasesMap
	var testCasesMap map[string]*testCaseModel.TestCaseModelStruct
	testCasesMap = *commandAndRuleEngine.Testcases.TestCasesMapPtr

	currentTestCasePtr, existsInMap = testCasesMap[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all Testcases")
		return err
	}

	currentElement, existInMap := currentTestCasePtr.TestCaseModelMap[uuidToDelete]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":           "4c328a0d-aaa1-4820-8fa2-e3067456faff",
			"uuidToDelete": uuidToDelete,
		}).Error(uuidToDelete + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(uuidToDelete + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	// Extract data for Previous Element
	currentElementUuid := uuidToDelete

	// Extract data for Previous Element
	previousElementUuid := currentElement.MatureTestCaseModelElementMessage.PreviousElementUuid

	// Extract data for Next Element
	nextElementUuid := currentElement.MatureTestCaseModelElementMessage.NextElementUuid

	// Create the structure after Delete

	// Create new B0-bond element
	newB0BondElement := commandAndRuleEngine.createNewBondB0Element()

	// Add New Elements to Map
	currentTestCasePtr.TestCaseModelMap[newB0BondElement.MatureElementUuid] = testCaseModel.MatureTestCaseModelElementStruct{
		MatureTestCaseModelElementMessage:  newB0BondElement,
		MatureTestCaseModelElementMetaData: testCaseModel.MatureTestCaseModelElementMetaDataStruct{},
	}

	// Remove Old Elements from Map
	delete(currentTestCasePtr.TestCaseModelMap, previousElementUuid)
	delete(currentTestCasePtr.TestCaseModelMap, nextElementUuid)

	// Remove references in currentElement to already removed Previous and Next Elements
	currentElement.MatureTestCaseModelElementMessage.PreviousElementUuid = currentElement.MatureTestCaseModelElementMessage.MatureElementUuid
	currentElement.MatureTestCaseModelElementMessage.NextElementUuid = currentElement.MatureTestCaseModelElementMessage.MatureElementUuid

	// Save updated currentElement back into TestCase-map
	currentTestCasePtr.TestCaseModelMap[currentElement.MatureTestCaseModelElementMessage.MatureElementUuid] = currentElement

	// Remove current element and children, if they exist, from map
	err = commandAndRuleEngine.recursiveDeleteOfChildElements(currentTestCasePtr, currentElementUuid)
	if err != nil {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":  "99abee1e-3078-42e7-a08a-a719e0a4ed8d",
			"err": err,
		}).Error(" Something went wrong when deleting children, in 'TestCaseModelMap', using recusive calls")

		err = errors.New("something went wrong when deleting children, in 'TestCaseModelMap', using recusive calls")

		return err
	}

	// Update first element of this TestCase
	currentTestCasePtr.FirstElementUuid = newB0BondElement.MatureElementUuid

	// If there are no errors then save the TestCase back into map of all TestCasesMapPtr
	//if err == nil {
	//	commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid] = currentTestCasePtr
	//}

	return err

}

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			B11f-n-B11l						B10							TCRuleDeletion102
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleDeletion102(testCaseUuid string, uuidToDelete string) (err error) {
	/*
		// One TestCaseElement, that is not a TestInstructionContainer, that hold type, reference to previous, next and parent elements
		message MatureTestCaseModelElementMessage {
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

	var currentTestCasePtr *testCaseModel.TestCaseModelStruct
	var existsInMap bool

	// Get TestCasesMap
	var testCasesMap map[string]*testCaseModel.TestCaseModelStruct
	testCasesMap = *commandAndRuleEngine.Testcases.TestCasesMapPtr

	// Get current TestCase
	currentTestCasePtr, existsInMap = testCasesMap[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all Testcases")
		return err
	}

	currentElement, existInMap := currentTestCasePtr.TestCaseModelMap[uuidToDelete]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":           "e2773f12-3bf6-4ebe-856b-25e6bb51864f",
			"uuidToDelete": uuidToDelete,
		}).Error(uuidToDelete + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(uuidToDelete + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	// Extract data for Previous Element
	currentElementUuid := uuidToDelete

	// Extract data for Previous Element
	previousElementUuid := currentElement.MatureTestCaseModelElementMessage.PreviousElementUuid

	// Extract data for Next Element
	nextElementUuid := currentElement.MatureTestCaseModelElementMessage.NextElementUuid

	// Extract dta for Parent Element
	parentElementUuid := currentElement.MatureTestCaseModelElementMessage.ParentElementUuid

	// Create the structure after Delete

	// Create new B10-bond element
	newB10BondElement := commandAndRuleEngine.createNewBondB10Element(parentElementUuid)

	// Add New Elements to Map
	currentTestCasePtr.TestCaseModelMap[newB10BondElement.MatureElementUuid] = testCaseModel.MatureTestCaseModelElementStruct{
		MatureTestCaseModelElementMessage:  newB10BondElement,
		MatureTestCaseModelElementMetaData: testCaseModel.MatureTestCaseModelElementMetaDataStruct{},
	}

	// Remove Old Elements from Map
	delete(currentTestCasePtr.TestCaseModelMap, previousElementUuid)
	delete(currentTestCasePtr.TestCaseModelMap, nextElementUuid)

	// Remove references in currentElement to already removed Previous and Next Elements
	currentElement.MatureTestCaseModelElementMessage.PreviousElementUuid = currentElement.MatureTestCaseModelElementMessage.MatureElementUuid
	currentElement.MatureTestCaseModelElementMessage.NextElementUuid = currentElement.MatureTestCaseModelElementMessage.MatureElementUuid

	// Save updated currentElement back into TestCase-map
	currentTestCasePtr.TestCaseModelMap[currentElement.MatureTestCaseModelElementMessage.MatureElementUuid] = currentElement

	// Remove current element and children, if they exist, from map
	err = commandAndRuleEngine.recursiveDeleteOfChildElements(currentTestCasePtr, currentElementUuid)
	if err != nil {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":  "43f9cfd5-5cfe-4ec0-b9e9-20271ab868e4",
			"err": err,
		}).Error(" Something went wrong when deleting children, in 'TestCaseModelMap', using recusive calls")

		err = errors.New("something went wrong when deleting children, in 'TestCaseModelMap', using recusive calls")

		return err
	}

	// Update Reference in parent TIC
	tempParentElement, existInMap := currentTestCasePtr.TestCaseModelMap[parentElementUuid]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":                "b75f6401-0393-411d-bcb6-28005c03ac9d",
			"parentElementUuid": parentElementUuid,
		}).Error(parentElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(parentElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	tempParentElement.MatureTestCaseModelElementMessage.FirstChildElementUuid = newB10BondElement.MatureElementUuid

	// Add updated parent-element back into TestCaseModelMap
	currentTestCasePtr.TestCaseModelMap[parentElementUuid] = tempParentElement

	// If there are no errors then save the TestCase back into map of all TestCasesMapPtr
	//if err == nil {
	//	commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid] = currentTestCasePtr
	//}

	return err

}

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			B11fx-n-B11lx					B10*x*						TCRuleDeletion103
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleDeletion103(testCaseUuid string, uuidToDelete string) (err error) {
	/*
		// One TestCaseElement, that is not a TestInstructionContainer, that hold type, reference to previous, next and parent elements
		message MatureTestCaseModelElementMessage {
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

	var existsInMap bool

	// Get TestCasesMap
	var testCasesMap map[string]*testCaseModel.TestCaseModelStruct
	testCasesMap = *commandAndRuleEngine.Testcases.TestCasesMapPtr

	var currentTestCasePtr *testCaseModel.TestCaseModelStruct

	// Get current TestCase
	currentTestCasePtr, existsInMap = testCasesMap[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all Testcases")
		return err
	}

	currentElement, existInMap := currentTestCasePtr.TestCaseModelMap[uuidToDelete]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":           "83137d58-dd37-443c-bf1f-0b01b7a85a8b",
			"uuidToDelete": uuidToDelete,
		}).Error(uuidToDelete + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(uuidToDelete + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	// Extract data for Previous Element
	currentElementUuid := uuidToDelete

	// Extract data for Previous Element
	previousElementUuid := currentElement.MatureTestCaseModelElementMessage.PreviousElementUuid

	// Extract data for Next Element
	nextElementUuid := currentElement.MatureTestCaseModelElementMessage.NextElementUuid

	// Extract dta for Parent Element
	parentElementUuid := currentElement.MatureTestCaseModelElementMessage.ParentElementUuid

	// Create the structure after Delete

	// Create new B10-bond element
	newB10oxoBondElement := commandAndRuleEngine.createNewBondB10oxoElement(parentElementUuid)

	// Add New Elements to Map
	currentTestCasePtr.TestCaseModelMap[newB10oxoBondElement.MatureElementUuid] = testCaseModel.MatureTestCaseModelElementStruct{
		MatureTestCaseModelElementMessage:  newB10oxoBondElement,
		MatureTestCaseModelElementMetaData: testCaseModel.MatureTestCaseModelElementMetaDataStruct{},
	}

	// Remove Old Elements from Map
	delete(currentTestCasePtr.TestCaseModelMap, previousElementUuid)
	delete(currentTestCasePtr.TestCaseModelMap, nextElementUuid)

	// Remove references in currentElement to already removed Previous and Next Elements
	currentElement.MatureTestCaseModelElementMessage.PreviousElementUuid = currentElement.MatureTestCaseModelElementMessage.MatureElementUuid
	currentElement.MatureTestCaseModelElementMessage.NextElementUuid = currentElement.MatureTestCaseModelElementMessage.MatureElementUuid

	// Save updated currentElement back into TestCase-map
	currentTestCasePtr.TestCaseModelMap[currentElement.MatureTestCaseModelElementMessage.MatureElementUuid] = currentElement

	// Remove current element and children, if they exist, from map
	err = commandAndRuleEngine.recursiveDeleteOfChildElements(currentTestCasePtr, currentElementUuid)
	if err != nil {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":  "f2683652-02ef-4260-8cb3-15cf627ddfa9",
			"err": err,
		}).Error(" Something went wrong when deleting children, in 'TestCaseModelMap', using recusive calls")

		err = errors.New("something went wrong when deleting children, in 'TestCaseModelMap', using recusive calls")

		return err
	}

	// Update Reference in parent TIC
	tempParentElement, existInMap := currentTestCasePtr.TestCaseModelMap[parentElementUuid]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":                "68faa4b5-bf93-469e-b2b4-fe2dfb192650",
			"parentElementUuid": parentElementUuid,
		}).Error(parentElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(parentElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	tempParentElement.MatureTestCaseModelElementMessage.FirstChildElementUuid = newB10oxoBondElement.MatureElementUuid

	// Add updated parent-element back into TestCaseModelMap
	currentTestCasePtr.TestCaseModelMap[parentElementUuid] = tempParentElement

	// If there are no errors then save the TestCase back into map of all TestCasesMapPtr
	//if err == nil {
	//	commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid] = currentTestCase
	//}

	return err

}

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			B11f-n-B11lx					B10x*						TCRuleDeletion104
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleDeletion104(testCaseUuid string, uuidToDelete string) (err error) {
	/*
		// One TestCaseElement, that is not a TestInstructionContainer, that hold type, reference to previous, next and parent elements
		message MatureTestCaseModelElementMessage {
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

	var existsInMap bool

	// Get TestCasesMap
	var testCasesMap map[string]*testCaseModel.TestCaseModelStruct
	testCasesMap = *commandAndRuleEngine.Testcases.TestCasesMapPtr

	var currentTestCasePtr *testCaseModel.TestCaseModelStruct

	// Get current TestCase
	currentTestCasePtr, existsInMap = testCasesMap[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all Testcases")
		return err
	}

	currentElement, existInMap := currentTestCasePtr.TestCaseModelMap[uuidToDelete]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":           "e7cf67a0-c8b9-44c6-920c-b0ef5d899d99",
			"uuidToDelete": uuidToDelete,
		}).Error(uuidToDelete + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(uuidToDelete + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	// Extract data for Previous Element
	currentElementUuid := uuidToDelete

	// Extract data for Previous Element
	previousElementUuid := currentElement.MatureTestCaseModelElementMessage.PreviousElementUuid

	// Extract data for Next Element
	nextElementUuid := currentElement.MatureTestCaseModelElementMessage.NextElementUuid

	// Extract dta for Parent Element
	parentElementUuid := currentElement.MatureTestCaseModelElementMessage.ParentElementUuid

	// Create the structure after Delete

	// Create new B10-bond element
	newB10xoBondElement := commandAndRuleEngine.createNewBondB10xoElement(parentElementUuid)

	// Add New Elements to Map
	currentTestCasePtr.TestCaseModelMap[newB10xoBondElement.MatureElementUuid] = testCaseModel.MatureTestCaseModelElementStruct{
		MatureTestCaseModelElementMessage:  newB10xoBondElement,
		MatureTestCaseModelElementMetaData: testCaseModel.MatureTestCaseModelElementMetaDataStruct{},
	}

	// Remove Old Elements from Map
	delete(currentTestCasePtr.TestCaseModelMap, previousElementUuid)
	delete(currentTestCasePtr.TestCaseModelMap, nextElementUuid)

	// Remove references in currentElement to already removed Previous and Next Elements
	currentElement.MatureTestCaseModelElementMessage.PreviousElementUuid = currentElement.MatureTestCaseModelElementMessage.MatureElementUuid
	currentElement.MatureTestCaseModelElementMessage.NextElementUuid = currentElement.MatureTestCaseModelElementMessage.MatureElementUuid

	// Save updated currentElement back into TestCase-map
	currentTestCasePtr.TestCaseModelMap[currentElement.MatureTestCaseModelElementMessage.MatureElementUuid] = currentElement

	// Remove current element and children, if they exist, from map
	err = commandAndRuleEngine.recursiveDeleteOfChildElements(currentTestCasePtr, currentElementUuid)
	if err != nil {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":  "b464d0dc-86b6-405b-802d-b538a6c2c840",
			"err": err,
		}).Error(" Something went wrong when deleting children, in 'TestCaseModelMap', using recusive calls")

		err = errors.New("something went wrong when deleting children, in 'TestCaseModelMap', using recusive calls")

		return err
	}

	// Update Reference in parent TIC
	tempParentElement, existInMap := currentTestCasePtr.TestCaseModelMap[parentElementUuid]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":                "6d7eae16-7b13-4f3e-9cbe-cc564307d86c",
			"parentElementUuid": parentElementUuid,
		}).Error(parentElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(parentElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	tempParentElement.MatureTestCaseModelElementMessage.FirstChildElementUuid = newB10xoBondElement.MatureElementUuid

	// Add updated parent-element back into TestCaseModelMap
	currentTestCasePtr.TestCaseModelMap[parentElementUuid] = tempParentElement

	// If there are no errors then save the TestCase back into map of all TestCasesMapPtr
	//if err == nil {
	//	commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid] = currentTestCase
	//}

	return err

}

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			B11fx-n-B11l					B10*x						TCRuleDeletion105
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleDeletion105(testCaseUuid string, uuidToDelete string) (err error) {
	/*
		// One TestCaseElement, that is not a TestInstructionContainer, that hold type, reference to previous, next and parent elements
		message MatureTestCaseModelElementMessage {
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

	var existsInMap bool

	// Get TestCasesMap
	var testCasesMap map[string]*testCaseModel.TestCaseModelStruct
	testCasesMap = *commandAndRuleEngine.Testcases.TestCasesMapPtr

	var currentTestCasePtr *testCaseModel.TestCaseModelStruct

	// Get current TestCase
	currentTestCasePtr, existsInMap = testCasesMap[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all Testcases")
		return err
	}

	currentElement, existInMap := currentTestCasePtr.TestCaseModelMap[uuidToDelete]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":           "b69296dd-7b73-46c9-b465-a8fb40a9a592",
			"uuidToDelete": uuidToDelete,
		}).Error(uuidToDelete + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(uuidToDelete + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	// Extract data for Previous Element
	currentElementUuid := uuidToDelete

	// Extract data for Previous Element
	previousElementUuid := currentElement.MatureTestCaseModelElementMessage.PreviousElementUuid

	// Extract data for Next Element
	nextElementUuid := currentElement.MatureTestCaseModelElementMessage.NextElementUuid

	// Extract dta for Parent Element
	parentElementUuid := currentElement.MatureTestCaseModelElementMessage.ParentElementUuid

	// Create the structure after Delete

	// Create new B10-bond element
	newB10oxBondElement := commandAndRuleEngine.createNewBondB10oxElement(parentElementUuid)

	// Add New Elements to Map
	currentTestCasePtr.TestCaseModelMap[newB10oxBondElement.MatureElementUuid] = testCaseModel.MatureTestCaseModelElementStruct{
		MatureTestCaseModelElementMessage:  newB10oxBondElement,
		MatureTestCaseModelElementMetaData: testCaseModel.MatureTestCaseModelElementMetaDataStruct{},
	}

	// Remove Old Elements from Map
	delete(currentTestCasePtr.TestCaseModelMap, previousElementUuid)
	delete(currentTestCasePtr.TestCaseModelMap, nextElementUuid)

	// Remove references in currentElement to already removed Previous and Next Elements
	currentElement.MatureTestCaseModelElementMessage.PreviousElementUuid = currentElement.MatureTestCaseModelElementMessage.MatureElementUuid
	currentElement.MatureTestCaseModelElementMessage.NextElementUuid = currentElement.MatureTestCaseModelElementMessage.MatureElementUuid

	// Save updated currentElement back into TestCase-map
	currentTestCasePtr.TestCaseModelMap[currentElement.MatureTestCaseModelElementMessage.MatureElementUuid] = currentElement

	// Remove current element and children, if they exist, from map
	err = commandAndRuleEngine.recursiveDeleteOfChildElements(currentTestCasePtr, currentElementUuid)
	if err != nil {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":  "ccda61cf-5c53-4248-aaa2-9be53bd7f46b",
			"err": err,
		}).Error(" Something went wrong when deleting children, in 'TestCaseModelMap', using recursive calls")

		err = errors.New("something went wrong when deleting children, in 'TestCaseModelMap', using recursive calls")

		return err
	}

	// Update Reference in parent TIC
	tempParentElement, existInMap := currentTestCasePtr.TestCaseModelMap[parentElementUuid]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":                "5117aa1e-6382-428a-a15d-bafa2528748c",
			"parentElementUuid": parentElementUuid,
		}).Error(parentElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(parentElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	tempParentElement.MatureTestCaseModelElementMessage.FirstChildElementUuid = newB10oxBondElement.MatureElementUuid

	// Add updated parent-element back into TestCaseModelMap
	currentTestCasePtr.TestCaseModelMap[parentElementUuid] = tempParentElement

	// If there are no errors then save the TestCase back into map of all TestCasesMapPtr
	//if err == nil {
	//	commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid] = currentTestCasePtr
	//}

	return err

}

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			B11f-n-B12-X					B11f-X						TCRuleDeletion106
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleDeletion106(testCaseUuid string, uuidToDelete string) (err error) {

	// Call shared delete function
	err = commandAndRuleEngine.executeTCRuleDeletion_106_107_108_109(testCaseUuid, uuidToDelete)

	return err
}

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			B11fx-n-B12x-X					B11fx-X						TCRuleDeletion107
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleDeletion107(testCaseUuid string, uuidToDelete string) (err error) {

	// Call shared delete function
	err = commandAndRuleEngine.executeTCRuleDeletion_106_107_108_109(testCaseUuid, uuidToDelete)

	return err
}

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			B11f-n-B12x-X					B11fx-X						TCRuleDeletion108
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleDeletion108(testCaseUuid string, uuidToDelete string) (err error) {

	// Call shared delete function
	err = commandAndRuleEngine.executeTCRuleDeletion_106_107_108_109(testCaseUuid, uuidToDelete)

	return err
}

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			B11fx-n-B12-X					B11fx-X						TCRuleDeletion109
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleDeletion109(testCaseUuid string, uuidToDelete string) (err error) {

	// Call shared delete function
	err = commandAndRuleEngine.executeTCRuleDeletion_106_107_108_109(testCaseUuid, uuidToDelete)

	return err
}

// Multi-deletion-rule function
// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			FirstBond-n-Bond-X				FirstBond-X					executeTCRuleDeletion_106_107_108_109
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleDeletion_106_107_108_109(testCaseUuid string, uuidToDelete string) (err error) {
	/*
		// One TestCaseElement, that is not a TestInstructionContainer, that hold type, reference to previous, next and parent elements
		message MatureTestCaseModelElementMessage {
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

	var existsInMap bool

	// Get TestCasesMap
	var testCasesMap map[string]*testCaseModel.TestCaseModelStruct
	testCasesMap = *commandAndRuleEngine.Testcases.TestCasesMapPtr

	var currentTestCasePtr *testCaseModel.TestCaseModelStruct

	// Get current TestCase
	currentTestCasePtr, existsInMap = testCasesMap[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all Testcases")
		return err
	}

	currentElement, existInMap := currentTestCasePtr.TestCaseModelMap[uuidToDelete]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":           "b69296dd-7b73-46c9-b465-a8fb40a9a592",
			"uuidToDelete": uuidToDelete,
		}).Error(uuidToDelete + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(uuidToDelete + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	// Extract data for Previous Element
	currentElementUuid := uuidToDelete

	// Extract data for Previous Element
	previousElementUuid := currentElement.MatureTestCaseModelElementMessage.PreviousElementUuid
	previousElement, existInMap := currentTestCasePtr.TestCaseModelMap[previousElementUuid]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":                "b8837af2-5578-4fbc-9513-d438ebf7af2c",
			"parentElementUuid": previousElementUuid,
		}).Error(previousElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(previousElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	// Extract data for Next Element
	nextElementUuid := currentElement.MatureTestCaseModelElementMessage.NextElementUuid
	nextElement, existInMap := currentTestCasePtr.TestCaseModelMap[nextElementUuid]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":                "06ee4383-d249-4a42-b3e6-9327b3b2b1ef",
			"parentElementUuid": nextElementUuid,
		}).Error(nextElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(nextElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	// Extract data for Next-Next Element
	nextNextElementUuid := nextElement.MatureTestCaseModelElementMessage.NextElementUuid
	nextNextElement, existInMap := currentTestCasePtr.TestCaseModelMap[nextNextElementUuid]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":                "7ef92315-adfd-4cc9-808b-ed4ac2537752",
			"parentElementUuid": nextNextElementUuid,
		}).Error(nextNextElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(nextNextElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	// Connect new structure
	previousElement.MatureTestCaseModelElementMessage.NextElementUuid = nextNextElementUuid
	nextNextElement.MatureTestCaseModelElementMessage.PreviousElementUuid = previousElementUuid

	// Remove Old Elements from Map
	delete(currentTestCasePtr.TestCaseModelMap, nextElementUuid)

	// Remove references in currentElement to already removed Next Elements
	currentElement.MatureTestCaseModelElementMessage.PreviousElementUuid = currentElement.MatureTestCaseModelElementMessage.MatureElementUuid
	currentElement.MatureTestCaseModelElementMessage.NextElementUuid = currentElement.MatureTestCaseModelElementMessage.MatureElementUuid

	// Handle special case for switching 'b11f' into 'b11fx' when there is a 'b12x' that is auto-deleted
	if previousElement.MatureTestCaseModelElementMessage.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND &&
		nextElement.MatureTestCaseModelElementMessage.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE {

		previousElement.MatureTestCaseModelElementMessage.TestCaseModelElementType = fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE
	}

	// Save updated back into TestCase-map
	currentTestCasePtr.TestCaseModelMap[previousElement.MatureTestCaseModelElementMessage.MatureElementUuid] = previousElement
	currentTestCasePtr.TestCaseModelMap[nextNextElement.MatureTestCaseModelElementMessage.MatureElementUuid] = nextNextElement
	currentTestCasePtr.TestCaseModelMap[currentElement.MatureTestCaseModelElementMessage.MatureElementUuid] = currentElement

	// Remove current element and children, if they exist, from map
	err = commandAndRuleEngine.recursiveDeleteOfChildElements(currentTestCasePtr, currentElementUuid)
	if err != nil {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":  "ccda61cf-5c53-4248-aaa2-9be53bd7f46b",
			"err": err,
		}).Error(" Something went wrong when deleting children, in 'TestCaseModelMap', using recursive calls")

		err = errors.New("something went wrong when deleting children, in 'TestCaseModelMap', using recursive calls")

		return err
	}

	// If there are no errors then save the TestCase back into map of all TestCasesMapPtr
	//if err == nil {
	//	commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid] = currentTestCasePtr
	//}

	return err

}

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-B12-n-B11l					X-B11l						TCRuleDeletion110
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleDeletion110(testCaseUuid string, uuidToDelete string) (err error) {

	// Call shared delete function
	err = commandAndRuleEngine.executeTCRuleDeletion_110_111_112_113(testCaseUuid, uuidToDelete)

	return err
}

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-B12x-n-B11lx					X-B11lx						TCRuleDeletion111
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleDeletion111(testCaseUuid string, uuidToDelete string) (err error) {

	// Call shared delete function
	err = commandAndRuleEngine.executeTCRuleDeletion_110_111_112_113(testCaseUuid, uuidToDelete)

	return err
}

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-B12-n-B11lx					X-B11lx						TCRuleDeletion112
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleDeletion112(testCaseUuid string, uuidToDelete string) (err error) {

	// Call shared delete function
	err = commandAndRuleEngine.executeTCRuleDeletion_110_111_112_113(testCaseUuid, uuidToDelete)

	return err
}

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-B12x-n-B11l					X-B11lx						TCRuleDeletion113
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleDeletion113(testCaseUuid string, uuidToDelete string) (err error) {

	// Call shared delete function
	err = commandAndRuleEngine.executeTCRuleDeletion_110_111_112_113(testCaseUuid, uuidToDelete)

	return err
}

// Multi-deletion-rule function
// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-Bond-n-LastBond				X-LastBond					executeTCRuleDeletion_110_111_112_113
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleDeletion_110_111_112_113(testCaseUuid string, uuidToDelete string) (err error) {
	/*
		// One TestCaseElement, that is not a TestInstructionContainer, that hold type, reference to previous, next and parent elements
		message MatureTestCaseModelElementMessage {
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

	var existsInMap bool

	// Get TestCasesMap
	var testCasesMap map[string]*testCaseModel.TestCaseModelStruct
	testCasesMap = *commandAndRuleEngine.Testcases.TestCasesMapPtr

	var currentTestCasePtr *testCaseModel.TestCaseModelStruct

	// Get current TestCase
	currentTestCasePtr, existsInMap = testCasesMap[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all Testcases")
		return err
	}

	currentElement, existInMap := currentTestCasePtr.TestCaseModelMap[uuidToDelete]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":           "b69296dd-7b73-46c9-b465-a8fb40a9a592",
			"uuidToDelete": uuidToDelete,
		}).Error(uuidToDelete + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(uuidToDelete + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	// Extract data for Previous Element
	currentElementUuid := uuidToDelete

	// Extract data for Previous Element
	previousElementUuid := currentElement.MatureTestCaseModelElementMessage.PreviousElementUuid
	previousElement, existInMap := currentTestCasePtr.TestCaseModelMap[previousElementUuid]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":                "c7629a23-53b6-4db3-a43f-ca1fde8c1ed0",
			"parentElementUuid": previousElementUuid,
		}).Error(previousElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(previousElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	// Extract data for Next Element
	nextElementUuid := currentElement.MatureTestCaseModelElementMessage.NextElementUuid
	nextElement, existInMap := currentTestCasePtr.TestCaseModelMap[nextElementUuid]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":                "9949b78e-7ef2-4946-ae02-1b1fabff9877",
			"parentElementUuid": nextElementUuid,
		}).Error(nextElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(nextElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	// Extract data for Previous-Previous Element
	previousPreviousElementUuid := previousElement.MatureTestCaseModelElementMessage.PreviousElementUuid
	previousPreviousElement, existInMap := currentTestCasePtr.TestCaseModelMap[previousPreviousElementUuid]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":                "d9d652aa-90d4-4e5a-8851-f89d6b970091",
			"parentElementUuid": previousPreviousElementUuid,
		}).Error(previousPreviousElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(previousPreviousElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	// Connect new structure
	nextElement.MatureTestCaseModelElementMessage.PreviousElementUuid = previousPreviousElementUuid
	previousPreviousElement.MatureTestCaseModelElementMessage.NextElementUuid = nextElementUuid

	// Remove Old Elements from Map
	delete(currentTestCasePtr.TestCaseModelMap, previousElementUuid)

	// Remove references in currentElement to already removed Next Elements
	currentElement.MatureTestCaseModelElementMessage.NextElementUuid = currentElement.MatureTestCaseModelElementMessage.MatureElementUuid

	// Handle special case for switching 'b11l' into 'b11lx' when there is a 'b12x' that is auto-deleted
	if previousElement.MatureTestCaseModelElementMessage.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE &&
		nextElement.MatureTestCaseModelElementMessage.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND {

		nextElement.MatureTestCaseModelElementMessage.TestCaseModelElementType = fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE
	}

	// Save updated back into TestCase-map
	currentTestCasePtr.TestCaseModelMap[previousPreviousElement.MatureTestCaseModelElementMessage.MatureElementUuid] = previousPreviousElement
	currentTestCasePtr.TestCaseModelMap[nextElement.MatureTestCaseModelElementMessage.MatureElementUuid] = nextElement
	currentTestCasePtr.TestCaseModelMap[currentElement.MatureTestCaseModelElementMessage.MatureElementUuid] = currentElement

	// Remove current element and children, if they exist, from map
	err = commandAndRuleEngine.recursiveDeleteOfChildElements(currentTestCasePtr, currentElementUuid)
	if err != nil {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":  "73d5c570-d845-4685-a9f8-158ef782eee3",
			"err": err,
		}).Error(" Something went wrong when deleting children, in 'TestCaseModelMap', using recursive calls")

		err = errors.New("something went wrong when deleting children, in 'TestCaseModelMap', using recursive calls")

		return err
	}

	// If there are no errors then save the TestCase back into map of all TestCasesMapPtr
	//if err == nil {
	//	commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid] = currentTestCasePtr
	//}

	return err

}

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-B12-n-B12-X					X-B12-X						TCRuleDeletion114
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleDeletion114(testCaseUuid string, uuidToDelete string) (err error) {

	// Call shared delete function
	err = commandAndRuleEngine.executeTCRuleDeletion_114_115_116_117(testCaseUuid, uuidToDelete)

	return err
}

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-B12x-n-B12x-X					X-B12x-X					TCRuleDeletion115
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleDeletion115(testCaseUuid string, uuidToDelete string) (err error) {

	// Call shared delete function
	err = commandAndRuleEngine.executeTCRuleDeletion_114_115_116_117(testCaseUuid, uuidToDelete)

	return err
}

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-B12-n-B12x-X					X-B12x-X					TCRuleDeletion116
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleDeletion116(testCaseUuid string, uuidToDelete string) (err error) {

	// Call shared delete function
	err = commandAndRuleEngine.executeTCRuleDeletion_114_115_116_117(testCaseUuid, uuidToDelete)

	return err
}

// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-B12x-n-B12-X					X-B12x-X					TCRuleDeletion117
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleDeletion117(testCaseUuid string, uuidToDelete string) (err error) {

	// Call shared delete function
	err = commandAndRuleEngine.executeTCRuleDeletion_114_115_116_117(testCaseUuid, uuidToDelete)

	return err
}

// Multi-deletion-rule function
// What to remove			Remove in structure				Result after deletion		Rule
// n=TI or TIC(X)			X-Bond-n-Bond-X					X-Bond-X					executeTCRuleDeletion_114_115_116_117
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleDeletion_114_115_116_117(testCaseUuid string, uuidToDelete string) (err error) {
	/*
		// One TestCaseElement, that is not a TestInstructionContainer, that hold type, reference to previous, next and parent elements
		message MatureTestCaseModelElementMessage {
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

	var existsInMap bool

	// Get TestCasesMap
	var testCasesMap map[string]*testCaseModel.TestCaseModelStruct
	testCasesMap = *commandAndRuleEngine.Testcases.TestCasesMapPtr

	var currentTestCasePtr *testCaseModel.TestCaseModelStruct

	// Get current TestCase
	currentTestCasePtr, existsInMap = testCasesMap[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all Testcases")
		return err
	}

	currentElement, existInMap := currentTestCasePtr.TestCaseModelMap[uuidToDelete]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":           "1ceef411-16ce-4af4-a0ef-ff7caef2e06c",
			"uuidToDelete": uuidToDelete,
		}).Error(uuidToDelete + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(uuidToDelete + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	// Extract data for Previous Element
	currentElementUuid := uuidToDelete

	// Extract data for Previous Element
	previousElementUuid := currentElement.MatureTestCaseModelElementMessage.PreviousElementUuid
	previousElement, existInMap := currentTestCasePtr.TestCaseModelMap[previousElementUuid]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":                "5f7d484f-41ec-43af-9ec1-5698ef345832",
			"parentElementUuid": previousElementUuid,
		}).Error(previousElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(previousElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	// Extract data for Next Element
	nextElementUuid := currentElement.MatureTestCaseModelElementMessage.NextElementUuid
	nextElement, existInMap := currentTestCasePtr.TestCaseModelMap[nextElementUuid]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":                "0142571b-60ee-4699-8639-7a658640dcd9",
			"parentElementUuid": nextElementUuid,
		}).Error(nextElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(nextElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	// Decide which of the Bond (Previous or Next to keep)
	var elementToKeep *testCaseModel.MatureTestCaseModelElementStruct
	var elementToRemove *testCaseModel.MatureTestCaseModelElementStruct

	if previousElement.MatureTestCaseModelElementMessage.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND &&
		nextElement.MatureTestCaseModelElementMessage.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE {
		// Keep Next Element
		elementToKeep = &nextElement
		elementToRemove = &previousElement

	} else {
		// Keep previous Element
		elementToKeep = &previousElement
		elementToRemove = &nextElement
	}

	// Extract data for Previous-Previous Element
	previousPreviousElementUuid := previousElement.MatureTestCaseModelElementMessage.PreviousElementUuid
	previousPreviousElement, existInMap := currentTestCasePtr.TestCaseModelMap[previousPreviousElementUuid]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":                          "f989a5f4-3bcf-4f3b-90a8-e7d5427f96d4",
			"previousPreviousElementUuid": previousPreviousElementUuid,
		}).Error(previousPreviousElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(previousPreviousElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	// Extract data for Next-Next Element
	nextNextElementUuid := nextElement.MatureTestCaseModelElementMessage.NextElementUuid
	nextNextElement, existInMap := currentTestCasePtr.TestCaseModelMap[nextNextElementUuid]
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":                  "f989a5f4-3bcf-4f3b-90a8-e7d5427f96d4",
			"nextNextElementUuid": nextNextElementUuid,
		}).Error(nextNextElementUuid + " could not be found in in map 'TestCaseModelMap'")

		err = errors.New(nextNextElementUuid + " could not be found in in map 'TestCaseModelMap'")

		return err
	}

	// Connect new structure
	previousPreviousElement.MatureTestCaseModelElementMessage.NextElementUuid = elementToKeep.MatureTestCaseModelElementMessage.MatureElementUuid
	elementToKeep.MatureTestCaseModelElementMessage.PreviousElementUuid = previousPreviousElement.MatureTestCaseModelElementMessage.MatureElementUuid

	elementToKeep.MatureTestCaseModelElementMessage.NextElementUuid = nextNextElement.MatureTestCaseModelElementMessage.MatureElementUuid
	nextNextElement.MatureTestCaseModelElementMessage.PreviousElementUuid = elementToKeep.MatureTestCaseModelElementMessage.MatureElementUuid

	// Remove Old Elements from Map
	delete(currentTestCasePtr.TestCaseModelMap, elementToRemove.MatureTestCaseModelElementMessage.MatureElementUuid)

	// Remove references in currentElement to already removed Next Elements
	currentElement.MatureTestCaseModelElementMessage.PreviousElementUuid = currentElement.MatureTestCaseModelElementMessage.MatureElementUuid
	currentElement.MatureTestCaseModelElementMessage.NextElementUuid = currentElement.MatureTestCaseModelElementMessage.MatureElementUuid

	// Save updated back into TestCase-map
	currentTestCasePtr.TestCaseModelMap[previousPreviousElement.MatureTestCaseModelElementMessage.MatureElementUuid] = previousPreviousElement
	currentTestCasePtr.TestCaseModelMap[nextNextElement.MatureTestCaseModelElementMessage.MatureElementUuid] = nextNextElement
	currentTestCasePtr.TestCaseModelMap[elementToKeep.MatureTestCaseModelElementMessage.MatureElementUuid] = *elementToKeep
	currentTestCasePtr.TestCaseModelMap[currentElement.MatureTestCaseModelElementMessage.MatureElementUuid] = currentElement

	// Remove current element and children, if they exist, from map
	err = commandAndRuleEngine.recursiveDeleteOfChildElements(currentTestCasePtr, currentElementUuid)
	if err != nil {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":  "43492f86-71ee-4a72-bfa4-1a68f84fcbed",
			"err": err,
		}).Error(" Something went wrong when deleting children, in 'TestCaseModelMap', using recursive calls")

		err = errors.New("something went wrong when deleting children, in 'TestCaseModelMap', using recursive calls")

		return err
	}

	// If there are no errors then save the TestCase back into map of all TestCasesMapPtr
	//if err == nil {
	//	commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid] = currentTestCasePtr
	//}

	return err

}

// Remove all children, in TestCase-model, for specific Element
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) recursiveDeleteOfChildElements(currentTestCase *testCaseModel.TestCaseModelStruct, elementsUuid string) (err error) {

	// Get current TestCase
	/*
		currentTestCase, existsInMap := commandAndRuleEngine.Testcases.TestCasesMapPtr[testCase]
		if existsInMap == false {
			err = errors.New("testcase with uuid '" + testCase + "' doesn't exist in map with all Testcases")
			return err
		}
	*/

	// Extract current element
	currentElement, existInMap := currentTestCase.TestCaseModelMap[elementsUuid]

	// If the element doesn't exit then there is something really wrong
	if existInMap == false {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":           "9eae5791-88f0-481d-a7d9-21123b9eadfe",
			"elementsUuid": elementsUuid,
		}).Error(elementsUuid + " could not be found in in map 'TestCaseModelMap'")
	}

	// Element has child-element then go that path
	if currentElement.MatureTestCaseModelElementMessage.FirstChildElementUuid != elementsUuid {
		err = commandAndRuleEngine.recursiveDeleteOfChildElements(currentTestCase, currentElement.MatureTestCaseModelElementMessage.FirstChildElementUuid)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return err
	}

	// If element has a next-element the go that path
	if currentElement.MatureTestCaseModelElementMessage.NextElementUuid != elementsUuid {
		err = commandAndRuleEngine.recursiveDeleteOfChildElements(currentTestCase, currentElement.MatureTestCaseModelElementMessage.NextElementUuid)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return err
	}

	// Remove current element from Map
	delete(currentTestCase.TestCaseModelMap, elementsUuid)

	return err
}
