package commandAndRuleEngine

import (
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	"fmt"
	"github.com/google/uuid"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) verifySwapRuleAndConvertIntoMatureComponentElementModel(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct, ruleNameToVerify string) (matureElementToSwapIn testCaseModel.MatureElementStruct, err error) {

	// Get ElementType for first element to be swapped in
	elementToBeSwappedIn, existsInMap := immatureElementToSwapIn.ImmatureElementMap[immatureElementToSwapIn.FirstElementUuid]
	if existsInMap == false {

		errorId := "3f3fbc1a-a35b-4a89-b8d2-cef7fb52d70d"
		err = errors.New(fmt.Sprintf("element referenced by first element ('%s')  doesn't exist in Immature element-map to be swapped in for TestCase '%s' [ErrorID: %s]", immatureElementToSwapIn.FirstElementUuid, testCaseUuid, errorId))

		return testCaseModel.MatureElementStruct{}, err
	}

	elementTypeToSwapIn := elementToBeSwappedIn.TestCaseModelElementType

	// Verify Rules before start swapping
	canBeSwapped, _, matchedComplexRule, err := commandAndRuleEngine.verifyIfElementCanBeSwapped(testCaseUuid, uuidToSwapOut, elementTypeToSwapIn)

	// Can't be swapped in
	if canBeSwapped == false ||
		matchedComplexRule != ruleNameToVerify {

		// Verify if there is any error messafe
		if err == nil {
			err = errors.New("can't be swapped du to simple rule validation")
		}

		// So exit
		return testCaseModel.MatureElementStruct{}, err
	}

	// Extract the TestCaseModel
	currentTestCaseModel, existsInMap := commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid]
	if existsInMap == false {
		errorId := "37fcb025-f91b-4c51-aac3-b0f50fba7de5"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCasesMapPtr [ErrorID: %s]", testCaseUuid, errorId))
	}

	// If Cut-command is in progress then
	if currentTestCaseModel.CutCommandInitiated == true {

		//  Swap from cut-buffer then element already is in mature form
		matureElementToSwapIn := currentTestCaseModel.CutBuffer

		return matureElementToSwapIn, err

	} else {
		// Transform ImmatureElementModel into a MatureElementModel
		matureElementToSwapIn, err = commandAndRuleEngine.transformImmatureElementModelIntoMatureElementModel(immatureElementToSwapIn)

		return matureElementToSwapIn, err

	}
}

// TCRuleSwap101
//
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC(X)			B0					n 		B0								B1-n-B1					TCRuleSwap101
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleSwap101(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (matureElementToSwapIn testCaseModel.MatureElementStruct, err error) {

	matureElementToSwapIn, err = commandAndRuleEngine.verifySwapRuleAndConvertIntoMatureComponentElementModel(testCaseUuid, uuidToSwapOut, immatureElementToSwapIn, TCRuleSwap101)

	// Couldn't convert immature element component into mature element component
	if err != nil {
		return testCaseModel.MatureElementStruct{}, err
	}

	currentTestCase, existsInMap := commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all Testcases")

		return testCaseModel.MatureElementStruct{}, err
	}

	// Create the Bonds connecting the TIC
	newPreviousB1fBond := commandAndRuleEngine.createNewBondB1fElement("")
	newNextB1lBond := commandAndRuleEngine.createNewBondB1lElement("")

	// Extract TIC from new element model
	newTopElement := matureElementToSwapIn.MatureElementMap[matureElementToSwapIn.FirstElementUuid]

	// Connect the new structure
	newPreviousB1fBond.NextElementUuid = newTopElement.MatureElementUuid

	newTopElement.PreviousElementUuid = newPreviousB1fBond.MatureElementUuid
	newTopElement.NextElementUuid = newNextB1lBond.MatureElementUuid

	newNextB1lBond.PreviousElementUuid = newTopElement.MatureElementUuid

	// Add updated element back to 'matureElementToSwapIn'
	matureElementToSwapIn.MatureElementMap[matureElementToSwapIn.FirstElementUuid] = newTopElement

	// Add new Bonds to TestCase Element Model
	elementToBeAdded := currentTestCase.TestCaseModelMap[newPreviousB1fBond.MatureElementUuid]
	elementToBeAdded.MatureTestCaseModelElementMessage = newPreviousB1fBond
	currentTestCase.TestCaseModelMap[newPreviousB1fBond.MatureElementUuid] = elementToBeAdded

	elementToBeAdded = currentTestCase.TestCaseModelMap[newNextB1lBond.MatureElementUuid]
	elementToBeAdded.MatureTestCaseModelElementMessage = newNextB1lBond
	currentTestCase.TestCaseModelMap[newNextB1lBond.MatureElementUuid] = elementToBeAdded

	// Set First Element
	currentTestCase.FirstElementUuid = newPreviousB1fBond.MatureElementUuid

	// Add 'matureElementToSwapIn' to TestCase Element Model
	for elementUuid, element := range matureElementToSwapIn.MatureElementMap {

		elementToBeAdded := currentTestCase.TestCaseModelMap[elementUuid]
		elementToBeAdded.MatureTestCaseModelElementMessage = element

		/*
			// Add Color to first Top Element that was swapped in
			if elementUuid == matureElementToSwapIn.FirstElementUuid {
				elementToBeAdded.MatureTestCaseModelElementMetaData.ChosenDropZoneColorString = matureElementToSwapIn.ChosenDropZoneColor
			}
		*/
		// Add color to element that is of Typ TI or TIx
		if elementToBeAdded.MatureTestCaseModelElementMessage.
			TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			elementToBeAdded.MatureTestCaseModelElementMessage.
				TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE {
			elementToBeAdded.MatureTestCaseModelElementMetaData.ChosenDropZoneColorString = matureElementToSwapIn.ChosenDropZoneColor
			elementToBeAdded.MatureTestCaseModelElementMetaData.ChosenDropZoneColorString = matureElementToSwapIn.ChosenDropZoneColor
		}

		currentTestCase.TestCaseModelMap[elementUuid] = elementToBeAdded

	}

	// Delete old element to be swapped out
	delete(currentTestCase.TestCaseModelMap, uuidToSwapOut)

	// If there are no errors then save the TestCase back into map of all TestCasesMapPtr
	if err == nil {
		commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid] = currentTestCase
	}

	return matureElementToSwapIn, err
}

// TCRuleSwap102
//
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B10					n		TIC(B10)						TIC(B11f-n-B11l)		TCRuleSwap102
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleSwap102(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (matureElementToSwapIn testCaseModel.MatureElementStruct, err error) {

	matureElementToSwapIn, err = commandAndRuleEngine.verifySwapRuleAndConvertIntoMatureComponentElementModel(testCaseUuid, uuidToSwapOut, immatureElementToSwapIn, TCRuleSwap102)

	// Couldn't convert immature element component into mature element component
	if err != nil {
		return testCaseModel.MatureElementStruct{}, err
	}

	currentTestCase, existsInMap := commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all Testcases")

		return testCaseModel.MatureElementStruct{}, err
	}

	// Extract parent-TIC to element to swap out
	elementToSwapOut, _ := currentTestCase.TestCaseModelMap[uuidToSwapOut]
	parentElementUuid := elementToSwapOut.MatureTestCaseModelElementMessage.ParentElementUuid
	parentElement := currentTestCase.TestCaseModelMap[parentElementUuid]

	// Create the Bonds connecting the TIC
	newPreviousB11fBond := commandAndRuleEngine.createNewBondB11fElement(parentElementUuid)
	newNextB11lBond := commandAndRuleEngine.createNewBondB11lElement(parentElementUuid)

	// Extract TIC/TI from new element model, same as first element
	topElementInModel := matureElementToSwapIn.MatureElementMap[matureElementToSwapIn.FirstElementUuid]

	// Connect the new structure
	newPreviousB11fBond.NextElementUuid = topElementInModel.MatureElementUuid

	topElementInModel.PreviousElementUuid = newPreviousB11fBond.MatureElementUuid
	topElementInModel.NextElementUuid = newNextB11lBond.MatureElementUuid
	topElementInModel.ParentElementUuid = parentElementUuid

	newNextB11lBond.PreviousElementUuid = topElementInModel.MatureElementUuid

	// Add updated element back to 'matureElementToSwapIn'
	matureElementToSwapIn.MatureElementMap[matureElementToSwapIn.FirstElementUuid] = topElementInModel

	// Update "first child" in parent element
	parentElement.MatureTestCaseModelElementMessage.FirstChildElementUuid = newPreviousB11fBond.MatureElementUuid

	// Save updated parent element back to TestCase model
	currentTestCase.TestCaseModelMap[parentElementUuid] = parentElement

	// Add new Bonds to TestCase Element Model
	elementToBeAdded := currentTestCase.TestCaseModelMap[newPreviousB11fBond.MatureElementUuid]
	elementToBeAdded.MatureTestCaseModelElementMessage = newPreviousB11fBond
	currentTestCase.TestCaseModelMap[newPreviousB11fBond.MatureElementUuid] = elementToBeAdded

	elementToBeAdded = currentTestCase.TestCaseModelMap[newNextB11lBond.MatureElementUuid]
	elementToBeAdded.MatureTestCaseModelElementMessage = newNextB11lBond
	currentTestCase.TestCaseModelMap[newNextB11lBond.MatureElementUuid] = elementToBeAdded

	// Add 'matureElementToSwapIn' to TestCase Element Model
	for elementUuid, element := range matureElementToSwapIn.MatureElementMap {

		elementToBeAdded := currentTestCase.TestCaseModelMap[elementUuid]
		elementToBeAdded.MatureTestCaseModelElementMessage = element

		/*
			// Add Color to first Top Element that was swapped in
			if elementUuid == matureElementToSwapIn.FirstElementUuid {
				elementToBeAdded.MatureTestCaseModelElementMetaData.ChosenDropZoneColorString = matureElementToSwapIn.ChosenDropZoneColor
			}
		*/
		// Add color to element that is of Typ TI or TIx
		if elementToBeAdded.MatureTestCaseModelElementMessage.
			TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			elementToBeAdded.MatureTestCaseModelElementMessage.
				TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE {
			elementToBeAdded.MatureTestCaseModelElementMetaData.ChosenDropZoneColorString = matureElementToSwapIn.ChosenDropZoneColor
		}

		currentTestCase.TestCaseModelMap[elementUuid] = elementToBeAdded

	}

	// Delete old element to be swapped out
	delete(currentTestCase.TestCaseModelMap, uuidToSwapOut)

	// If there are no errors then save the TestCase back into map of all TestCasesMapPtr
	if err == nil {
		commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid] = currentTestCase
	}

	return matureElementToSwapIn, err
}

// TCRuleSwap103
//
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B11f				n		TIC(B11f-X)						TIC(B11f-n-B12-X)		TCRuleSwap103
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleSwap103(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (matureElementToSwapIn testCaseModel.MatureElementStruct, err error) {

	matureElementToSwapIn, err = commandAndRuleEngine.verifySwapRuleAndConvertIntoMatureComponentElementModel(testCaseUuid, uuidToSwapOut, immatureElementToSwapIn, TCRuleSwap103)

	// Couldn't convert immature element component into mature element component
	if err != nil {
		return testCaseModel.MatureElementStruct{}, err
	}

	currentTestCase, existsInMap := commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all Testcases")

		return testCaseModel.MatureElementStruct{}, err
	}

	// Extract parent-TIC to element to swap out
	elementToSwapOut, _ := currentTestCase.TestCaseModelMap[uuidToSwapOut]
	parentElementUuid := elementToSwapOut.MatureTestCaseModelElementMessage.ParentElementUuid
	nextElement := currentTestCase.TestCaseModelMap[elementToSwapOut.MatureTestCaseModelElementMessage.NextElementUuid]

	// Create the new Bonds
	newB12Bond := commandAndRuleEngine.createNewBondB12Element(parentElementUuid)

	// Extract TIC/TI from new element model, same as first element
	topElementInModelToBeSwappedIn := matureElementToSwapIn.MatureElementMap[matureElementToSwapIn.FirstElementUuid]

	// Connect the new structure
	elementToSwapOut.MatureTestCaseModelElementMessage.NextElementUuid = topElementInModelToBeSwappedIn.MatureElementUuid

	topElementInModelToBeSwappedIn.PreviousElementUuid = elementToSwapOut.MatureTestCaseModelElementMessage.MatureElementUuid
	topElementInModelToBeSwappedIn.NextElementUuid = newB12Bond.MatureElementUuid
	topElementInModelToBeSwappedIn.ParentElementUuid = parentElementUuid

	newB12Bond.PreviousElementUuid = topElementInModelToBeSwappedIn.MatureElementUuid
	newB12Bond.NextElementUuid = nextElement.MatureTestCaseModelElementMessage.MatureElementUuid

	nextElement.MatureTestCaseModelElementMessage.PreviousElementUuid = newB12Bond.MatureElementUuid

	// Add updated element back to 'matureElementToSwapIn'
	matureElementToSwapIn.MatureElementMap[topElementInModelToBeSwappedIn.MatureElementUuid] = topElementInModelToBeSwappedIn

	// Update elements in TestCaseModel
	currentTestCase.TestCaseModelMap[elementToSwapOut.MatureTestCaseModelElementMessage.MatureElementUuid] = elementToSwapOut
	currentTestCase.TestCaseModelMap[nextElement.MatureTestCaseModelElementMessage.MatureElementUuid] = nextElement

	// Add new Bonds to TestCase Element Model
	currentTestCase.TestCaseModelMap[newB12Bond.MatureElementUuid] = testCaseModel.MatureTestCaseModelElementStruct{
		MatureTestCaseModelElementMessage:  newB12Bond,
		MatureTestCaseModelElementMetaData: testCaseModel.MatureTestCaseModelElementMetaDataStruct{},
	}

	// Add 'matureElementToSwapIn' to TestCase Element Model
	for elementUuid, element := range matureElementToSwapIn.MatureElementMap {

		elementToBeAdded := currentTestCase.TestCaseModelMap[elementUuid]
		elementToBeAdded.MatureTestCaseModelElementMessage = element

		/*
			// Add Color to first Top Element that was swapped in
			if elementUuid == matureElementToSwapIn.FirstElementUuid {
				elementToBeAdded.MatureTestCaseModelElementMetaData.ChosenDropZoneColorString = matureElementToSwapIn.ChosenDropZoneColor
			}
		*/
		// Add color to element that is of Typ TI or TIx
		if elementToBeAdded.MatureTestCaseModelElementMessage.
			TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			elementToBeAdded.MatureTestCaseModelElementMessage.
				TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE {
			elementToBeAdded.MatureTestCaseModelElementMetaData.ChosenDropZoneColorString = matureElementToSwapIn.ChosenDropZoneColor
		}

		currentTestCase.TestCaseModelMap[elementUuid] = elementToBeAdded

	}

	// If there are no errors then save the TestCase back into map of all TestCasesMapPtr
	if err == nil {
		commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid] = currentTestCase
	}

	return matureElementToSwapIn, err
}

// TCRuleSwap104
//
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B11l				n		TIC(X-B11l)						TIC(X-B12-n-B11l)		TCRuleSwap104
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleSwap104(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (matureElementToSwapIn testCaseModel.MatureElementStruct, err error) {

	matureElementToSwapIn, err = commandAndRuleEngine.verifySwapRuleAndConvertIntoMatureComponentElementModel(testCaseUuid, uuidToSwapOut, immatureElementToSwapIn, TCRuleSwap104)

	// Couldn't convert immature element component into mature element component
	if err != nil {
		return testCaseModel.MatureElementStruct{}, err
	}

	currentTestCase, existsInMap := commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all Testcases")

		return testCaseModel.MatureElementStruct{}, err
	}

	// Extract parent-TIC to element to swap out
	elementToSwapOut, _ := currentTestCase.TestCaseModelMap[uuidToSwapOut]
	parentElementUuid := elementToSwapOut.MatureTestCaseModelElementMessage.ParentElementUuid
	previousElement := currentTestCase.TestCaseModelMap[elementToSwapOut.MatureTestCaseModelElementMessage.PreviousElementUuid]

	// Create the new Bonds
	newB12Bond := commandAndRuleEngine.createNewBondB12Element(parentElementUuid)

	// Extract TIC/TI from new element model, same as first element
	topElementInModelToBeSwappedIn := matureElementToSwapIn.MatureElementMap[matureElementToSwapIn.FirstElementUuid]

	// Connect the new structure
	previousElement.MatureTestCaseModelElementMessage.NextElementUuid = newB12Bond.MatureElementUuid

	newB12Bond.PreviousElementUuid = previousElement.MatureTestCaseModelElementMessage.MatureElementUuid
	newB12Bond.NextElementUuid = topElementInModelToBeSwappedIn.MatureElementUuid

	topElementInModelToBeSwappedIn.PreviousElementUuid = newB12Bond.MatureElementUuid
	topElementInModelToBeSwappedIn.NextElementUuid = elementToSwapOut.MatureTestCaseModelElementMessage.MatureElementUuid
	topElementInModelToBeSwappedIn.ParentElementUuid = parentElementUuid

	elementToSwapOut.MatureTestCaseModelElementMessage.PreviousElementUuid = topElementInModelToBeSwappedIn.MatureElementUuid

	// Add updated element back to 'matureElementToSwapIn'
	matureElementToSwapIn.MatureElementMap[topElementInModelToBeSwappedIn.MatureElementUuid] = topElementInModelToBeSwappedIn

	// Update elements in TestCaseModel
	currentTestCase.TestCaseModelMap[elementToSwapOut.MatureTestCaseModelElementMessage.MatureElementUuid] = elementToSwapOut
	currentTestCase.TestCaseModelMap[previousElement.MatureTestCaseModelElementMessage.MatureElementUuid] = previousElement

	// Add new Bonds to TestCase Element Model
	currentTestCase.TestCaseModelMap[newB12Bond.MatureElementUuid] = testCaseModel.MatureTestCaseModelElementStruct{
		MatureTestCaseModelElementMessage:  newB12Bond,
		MatureTestCaseModelElementMetaData: testCaseModel.MatureTestCaseModelElementMetaDataStruct{},
	}

	// Add 'matureElementToSwapIn' to TestCase Element Model
	for elementUuid, element := range matureElementToSwapIn.MatureElementMap {

		elementToBeAdded := currentTestCase.TestCaseModelMap[elementUuid]
		elementToBeAdded.MatureTestCaseModelElementMessage = element

		/*
			// Add Color to first Top Element that was swapped in
			if elementUuid == matureElementToSwapIn.FirstElementUuid {
				elementToBeAdded.MatureTestCaseModelElementMetaData.ChosenDropZoneColorString = matureElementToSwapIn.ChosenDropZoneColor
			}
		*/
		// Add color to element that is of Typ TI or TIx
		if elementToBeAdded.MatureTestCaseModelElementMessage.
			TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			elementToBeAdded.MatureTestCaseModelElementMessage.
				TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE {
			elementToBeAdded.MatureTestCaseModelElementMetaData.ChosenDropZoneColorString = matureElementToSwapIn.ChosenDropZoneColor
		}

		currentTestCase.TestCaseModelMap[elementUuid] = elementToBeAdded

	}

	// If there are no errors then save the TestCase back into map of all TestCasesMapPtr
	if err == nil {
		commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid] = currentTestCase
	}

	return matureElementToSwapIn, err
}

// TCRuleSwap105
//
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B12					n		X-B12-X							X-B12-n-B12-X			TCRuleSwap105
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleSwap105(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (matureElementToSwapIn testCaseModel.MatureElementStruct, err error) {

	matureElementToSwapIn, err = commandAndRuleEngine.verifySwapRuleAndConvertIntoMatureComponentElementModel(testCaseUuid, uuidToSwapOut, immatureElementToSwapIn, TCRuleSwap105)

	// Couldn't convert immature element component into mature element component
	if err != nil {
		return testCaseModel.MatureElementStruct{}, err
	}

	currentTestCase, existsInMap := commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all Testcases")

		return testCaseModel.MatureElementStruct{}, err
	}

	// Extract parent-TIC to element to swap out
	elementToSwapOut, _ := currentTestCase.TestCaseModelMap[uuidToSwapOut]
	parentElementUuid := elementToSwapOut.MatureTestCaseModelElementMessage.ParentElementUuid
	previousElement := currentTestCase.TestCaseModelMap[elementToSwapOut.MatureTestCaseModelElementMessage.PreviousElementUuid]

	// Create the new Bonds
	newB12Bond := commandAndRuleEngine.createNewBondB12Element(parentElementUuid)

	// Extract TIC/TI from new element model, same as first element
	topElementInModelToBeSwappedIn := matureElementToSwapIn.MatureElementMap[matureElementToSwapIn.FirstElementUuid]

	// Connect the new structure
	previousElement.MatureTestCaseModelElementMessage.NextElementUuid = newB12Bond.MatureElementUuid

	newB12Bond.PreviousElementUuid = previousElement.MatureTestCaseModelElementMessage.MatureElementUuid
	newB12Bond.NextElementUuid = topElementInModelToBeSwappedIn.MatureElementUuid

	topElementInModelToBeSwappedIn.PreviousElementUuid = newB12Bond.MatureElementUuid
	topElementInModelToBeSwappedIn.NextElementUuid = elementToSwapOut.MatureTestCaseModelElementMessage.MatureElementUuid
	topElementInModelToBeSwappedIn.ParentElementUuid = parentElementUuid

	elementToSwapOut.MatureTestCaseModelElementMessage.PreviousElementUuid = topElementInModelToBeSwappedIn.MatureElementUuid

	// Add updated element back to 'matureElementToSwapIn'
	matureElementToSwapIn.MatureElementMap[topElementInModelToBeSwappedIn.MatureElementUuid] = topElementInModelToBeSwappedIn

	// Update elements in TestCaseModel
	currentTestCase.TestCaseModelMap[elementToSwapOut.MatureTestCaseModelElementMessage.MatureElementUuid] = elementToSwapOut
	currentTestCase.TestCaseModelMap[previousElement.MatureTestCaseModelElementMessage.MatureElementUuid] = previousElement

	// Add new Bonds to TestCase Element Model
	currentTestCase.TestCaseModelMap[newB12Bond.MatureElementUuid] = testCaseModel.MatureTestCaseModelElementStruct{
		MatureTestCaseModelElementMessage:  newB12Bond,
		MatureTestCaseModelElementMetaData: testCaseModel.MatureTestCaseModelElementMetaDataStruct{},
	}

	// Add 'matureElementToSwapIn' to TestCase Element Model
	for elementUuid, element := range matureElementToSwapIn.MatureElementMap {

		elementToBeAdded := currentTestCase.TestCaseModelMap[elementUuid]
		elementToBeAdded.MatureTestCaseModelElementMessage = element

		/*
			// Add Color to first Top Element that was swapped in
			if elementUuid == matureElementToSwapIn.FirstElementUuid {
				elementToBeAdded.MatureTestCaseModelElementMetaData.ChosenDropZoneColorString = matureElementToSwapIn.ChosenDropZoneColor
			}
		*/
		// Add color to element that is of Typ TI or TIx
		if elementToBeAdded.MatureTestCaseModelElementMessage.
			TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			elementToBeAdded.MatureTestCaseModelElementMessage.
				TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE {
			elementToBeAdded.MatureTestCaseModelElementMetaData.ChosenDropZoneColorString = matureElementToSwapIn.ChosenDropZoneColor
		}

		currentTestCase.TestCaseModelMap[elementUuid] = elementToBeAdded

	}

	// If there are no errors then save the TestCase back into map of all TestCasesMapPtr
	if err == nil {
		commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid] = currentTestCase
	}

	return matureElementToSwapIn, err
}

// TCRuleSwap106
//
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B10x*				n		TIC(B10*x*)						TIC(B11x-n-B11x)		TCRuleSwap106
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleSwap106(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (matureElementToSwapIn testCaseModel.MatureElementStruct, err error) {

	matureElementToSwapIn, err = commandAndRuleEngine.verifySwapRuleAndConvertIntoMatureComponentElementModel(testCaseUuid, uuidToSwapOut, immatureElementToSwapIn, TCRuleSwap106)

	// Couldn't convert immature element component into mature element component
	if err != nil {
		return testCaseModel.MatureElementStruct{}, err
	}

	currentTestCase, existsInMap := commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all Testcases")

		return testCaseModel.MatureElementStruct{}, err
	}

	// Extract parent-TIC to element to swap out
	elementToSwapOut, _ := currentTestCase.TestCaseModelMap[uuidToSwapOut]
	parentElementUuid := elementToSwapOut.MatureTestCaseModelElementMessage.ParentElementUuid
	parentElement := currentTestCase.TestCaseModelMap[parentElementUuid]

	// Create the Bonds connecting the TIC
	newPreviousB11fxBond := commandAndRuleEngine.createNewBondB11fxElement(parentElementUuid)
	newNextB11lxBond := commandAndRuleEngine.createNewBondB11lxElement(parentElementUuid)

	// Extract TIC/TI from new element model, same as first element
	topElementInModel := matureElementToSwapIn.MatureElementMap[matureElementToSwapIn.FirstElementUuid]

	// Connect the new structure
	newPreviousB11fxBond.NextElementUuid = topElementInModel.MatureElementUuid

	topElementInModel.PreviousElementUuid = newPreviousB11fxBond.MatureElementUuid
	topElementInModel.NextElementUuid = newNextB11lxBond.MatureElementUuid
	topElementInModel.ParentElementUuid = parentElementUuid

	newNextB11lxBond.PreviousElementUuid = topElementInModel.MatureElementUuid

	// Add updated element back to 'matureElementToSwapIn'
	matureElementToSwapIn.MatureElementMap[matureElementToSwapIn.FirstElementUuid] = topElementInModel

	// Update "first child" in parent element
	parentElement.MatureTestCaseModelElementMessage.FirstChildElementUuid = newPreviousB11fxBond.MatureElementUuid

	// Save updated parent element back to TestCase model
	currentTestCase.TestCaseModelMap[parentElementUuid] = parentElement

	// Add new Bonds to TestCase Element Model
	currentTestCase.TestCaseModelMap[newPreviousB11fxBond.MatureElementUuid] = testCaseModel.MatureTestCaseModelElementStruct{
		MatureTestCaseModelElementMessage:  newPreviousB11fxBond,
		MatureTestCaseModelElementMetaData: testCaseModel.MatureTestCaseModelElementMetaDataStruct{},
	}

	currentTestCase.TestCaseModelMap[newNextB11lxBond.MatureElementUuid] = testCaseModel.MatureTestCaseModelElementStruct{
		MatureTestCaseModelElementMessage:  newNextB11lxBond,
		MatureTestCaseModelElementMetaData: testCaseModel.MatureTestCaseModelElementMetaDataStruct{},
	}

	// Add 'matureElementToSwapIn' to TestCase Element Model
	for elementUuid, element := range matureElementToSwapIn.MatureElementMap {

		elementToBeAdded := currentTestCase.TestCaseModelMap[elementUuid]
		elementToBeAdded.MatureTestCaseModelElementMessage = element

		/*
			// Add Color to first Top Element that was swapped in
			if elementUuid == matureElementToSwapIn.FirstElementUuid {
				elementToBeAdded.MatureTestCaseModelElementMetaData.ChosenDropZoneColorString = matureElementToSwapIn.ChosenDropZoneColor
			}
		*/
		// Add color to element that is of Typ TI or TIx
		if elementToBeAdded.MatureTestCaseModelElementMessage.
			TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			elementToBeAdded.MatureTestCaseModelElementMessage.
				TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE {
			elementToBeAdded.MatureTestCaseModelElementMetaData.ChosenDropZoneColorString = matureElementToSwapIn.ChosenDropZoneColor
		}

		currentTestCase.TestCaseModelMap[elementUuid] = elementToBeAdded

	}

	// If there are no errors then save the TestCase back into map of all TestCasesMapPtr
	if err == nil {
		commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid] = currentTestCase
	}

	// Delete old element to be swapped out
	delete(currentTestCase.TestCaseModelMap, uuidToSwapOut)

	return matureElementToSwapIn, err
}

// TCRuleSwap107
//
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B10*x				n		TIC(B10*x)						TIC(B11x-n-B11)			TCRuleSwap107
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleSwap107(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (matureElementToSwapIn testCaseModel.MatureElementStruct, err error) {

	matureElementToSwapIn, err = commandAndRuleEngine.verifySwapRuleAndConvertIntoMatureComponentElementModel(testCaseUuid, uuidToSwapOut, immatureElementToSwapIn, TCRuleSwap107)

	// Couldn't convert immature element component into mature element component
	if err != nil {
		return testCaseModel.MatureElementStruct{}, err
	}

	currentTestCase, existsInMap := commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all Testcases")

		return testCaseModel.MatureElementStruct{}, err
	}

	// Extract parent-TIC to element to swap out
	elementToSwapOut, _ := currentTestCase.TestCaseModelMap[uuidToSwapOut]
	parentElementUuid := elementToSwapOut.MatureTestCaseModelElementMessage.ParentElementUuid
	parentElement := currentTestCase.TestCaseModelMap[parentElementUuid]

	// Create the Bonds connecting the TIC
	newPreviousB11fxBond := commandAndRuleEngine.createNewBondB11fxElement(parentElementUuid)
	newNextB11lBond := commandAndRuleEngine.createNewBondB11lElement(parentElementUuid)

	// Extract TIC/TI from new element model, same as first element
	topElementInModel := matureElementToSwapIn.MatureElementMap[matureElementToSwapIn.FirstElementUuid]

	// Connect the new structure
	newPreviousB11fxBond.NextElementUuid = topElementInModel.MatureElementUuid

	topElementInModel.PreviousElementUuid = newPreviousB11fxBond.MatureElementUuid
	topElementInModel.NextElementUuid = newNextB11lBond.MatureElementUuid
	topElementInModel.ParentElementUuid = parentElementUuid

	newNextB11lBond.PreviousElementUuid = topElementInModel.MatureElementUuid

	// Add updated element back to 'matureElementToSwapIn'
	matureElementToSwapIn.MatureElementMap[matureElementToSwapIn.FirstElementUuid] = topElementInModel

	// Update "first child" in parent element
	parentElement.MatureTestCaseModelElementMessage.FirstChildElementUuid = newPreviousB11fxBond.MatureElementUuid

	// Save updated parent element back to TestCase model
	currentTestCase.TestCaseModelMap[parentElementUuid] = parentElement

	// Add new Bonds to TestCase Element Model
	currentTestCase.TestCaseModelMap[newPreviousB11fxBond.MatureElementUuid] = testCaseModel.MatureTestCaseModelElementStruct{
		MatureTestCaseModelElementMessage:  newPreviousB11fxBond,
		MatureTestCaseModelElementMetaData: testCaseModel.MatureTestCaseModelElementMetaDataStruct{},
	}

	currentTestCase.TestCaseModelMap[newNextB11lBond.MatureElementUuid] = testCaseModel.MatureTestCaseModelElementStruct{
		MatureTestCaseModelElementMessage:  newNextB11lBond,
		MatureTestCaseModelElementMetaData: testCaseModel.MatureTestCaseModelElementMetaDataStruct{},
	}

	// Add 'matureElementToSwapIn' to TestCase Element Model
	for elementUuid, element := range matureElementToSwapIn.MatureElementMap {

		elementToBeAdded := currentTestCase.TestCaseModelMap[elementUuid]
		elementToBeAdded.MatureTestCaseModelElementMessage = element

		/*
			// Add Color to first Top Element that was swapped in
			if elementUuid == matureElementToSwapIn.FirstElementUuid {
				elementToBeAdded.MatureTestCaseModelElementMetaData.ChosenDropZoneColorString = matureElementToSwapIn.ChosenDropZoneColor
			}
		*/
		// Add color to element that is of Typ TI or TIx
		if elementToBeAdded.MatureTestCaseModelElementMessage.
			TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			elementToBeAdded.MatureTestCaseModelElementMessage.
				TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE {
			elementToBeAdded.MatureTestCaseModelElementMetaData.ChosenDropZoneColorString = matureElementToSwapIn.ChosenDropZoneColor
		}

		currentTestCase.TestCaseModelMap[elementUuid] = elementToBeAdded

	}

	// Delete old element to be swapped out
	delete(currentTestCase.TestCaseModelMap, uuidToSwapOut)

	// If there are no errors then save the TestCase back into map of all TestCasesMapPtr
	if err == nil {
		commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid] = currentTestCase
	}

	return matureElementToSwapIn, err
}

// TCRuleSwap108
//
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B10x*				n		TIC(B10x*)						TIC(B11-n-B11x)			TCRuleSwap108
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeTCRuleSwap108(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (matureElementToSwapIn testCaseModel.MatureElementStruct, err error) {

	matureElementToSwapIn, err = commandAndRuleEngine.verifySwapRuleAndConvertIntoMatureComponentElementModel(testCaseUuid, uuidToSwapOut, immatureElementToSwapIn, TCRuleSwap108)

	// Couldn't convert immature element component into mature element component
	if err != nil {
		return testCaseModel.MatureElementStruct{}, err
	}

	currentTestCase, existsInMap := commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all Testcases")

		return testCaseModel.MatureElementStruct{}, err
	}

	// Extract parent-TIC to element to swap out
	elementToSwapOut, _ := currentTestCase.TestCaseModelMap[uuidToSwapOut]
	parentElementUuid := elementToSwapOut.MatureTestCaseModelElementMessage.ParentElementUuid
	parentElement := currentTestCase.TestCaseModelMap[parentElementUuid]

	// Create the Bonds connecting the TIC
	newPreviousB11fBond := commandAndRuleEngine.createNewBondB11fElement(parentElementUuid)
	newNextB11lxBond := commandAndRuleEngine.createNewBondB11lxElement(parentElementUuid)

	// Extract TIC/TI from new element model, same as first element
	topElementInModel := matureElementToSwapIn.MatureElementMap[matureElementToSwapIn.FirstElementUuid]

	// Connect the new structure
	newPreviousB11fBond.NextElementUuid = topElementInModel.MatureElementUuid

	topElementInModel.PreviousElementUuid = newPreviousB11fBond.MatureElementUuid
	topElementInModel.NextElementUuid = newNextB11lxBond.MatureElementUuid
	topElementInModel.ParentElementUuid = parentElementUuid

	newNextB11lxBond.PreviousElementUuid = topElementInModel.MatureElementUuid

	// Add updated element back to 'matureElementToSwapIn'
	matureElementToSwapIn.MatureElementMap[matureElementToSwapIn.FirstElementUuid] = topElementInModel

	// Update "first child" in parent element
	parentElement.MatureTestCaseModelElementMessage.FirstChildElementUuid = newPreviousB11fBond.MatureElementUuid

	// Save updated parent element back to TestCase model
	currentTestCase.TestCaseModelMap[parentElementUuid] = parentElement

	// Add new Bonds to TestCase Element Model
	currentTestCase.TestCaseModelMap[newPreviousB11fBond.MatureElementUuid] = testCaseModel.MatureTestCaseModelElementStruct{
		MatureTestCaseModelElementMessage:  newPreviousB11fBond,
		MatureTestCaseModelElementMetaData: testCaseModel.MatureTestCaseModelElementMetaDataStruct{},
	}
	currentTestCase.TestCaseModelMap[newNextB11lxBond.MatureElementUuid] = testCaseModel.MatureTestCaseModelElementStruct{
		MatureTestCaseModelElementMessage:  newNextB11lxBond,
		MatureTestCaseModelElementMetaData: testCaseModel.MatureTestCaseModelElementMetaDataStruct{},
	}

	// Add 'matureElementToSwapIn' to TestCase Element Model
	for elementUuid, element := range matureElementToSwapIn.MatureElementMap {

		elementToBeAdded := currentTestCase.TestCaseModelMap[elementUuid]
		elementToBeAdded.MatureTestCaseModelElementMessage = element

		/*
			// Add Color to first Top Element that was swapped in
			if elementUuid == matureElementToSwapIn.FirstElementUuid {
				elementToBeAdded.MatureTestCaseModelElementMetaData.ChosenDropZoneColorString = matureElementToSwapIn.ChosenDropZoneColor
			}
		*/
		// Add color to element that is of Typ TI or TIx
		if elementToBeAdded.MatureTestCaseModelElementMessage.
			TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			elementToBeAdded.MatureTestCaseModelElementMessage.
				TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE {
			elementToBeAdded.MatureTestCaseModelElementMetaData.ChosenDropZoneColorString = matureElementToSwapIn.ChosenDropZoneColor
		}

		currentTestCase.TestCaseModelMap[elementUuid] = elementToBeAdded

	}

	// Delete old element to be swapped out
	delete(currentTestCase.TestCaseModelMap, uuidToSwapOut)

	// If there are no errors then save the TestCase back into map of all TestCasesMapPtr
	if err == nil {
		commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid] = currentTestCase
	}

	return matureElementToSwapIn, err
}

// Transforms a immature element model into a mature element model.
// This means that new UUIDs are created for each element in the component
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) transformImmatureElementModelIntoMatureElementModel(
	immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (
	matureElementModel testCaseModel.MatureElementStruct, err error) {

	var matureIndicator = "_mature"

	// Create the temp store for matureElementModel and initiate map
	tempMatureElementModel := testCaseModel.MatureElementStruct{}
	tempMatureElementModel.MatureElementMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage)

	// Initiate the matureElementModel map
	matureElementModel.MatureElementMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage)

	// Loop all ImmatureElements in Component and create a raw mature version of each immature element
	for immatureElementUuid, immatureElement := range immatureElementToSwapIn.ImmatureElementMap {

		// Create the 'raw 'MatureElement from an ImmatureElement. This element is not yet connected
		newRawMatureComponent := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
			OriginalElementUuid:      immatureElement.OriginalElementUuid,
			OriginalElementName:      immatureElement.OriginalElementName,
			MatureElementUuid:        uuid.New().String(),
			PreviousElementUuid:      "",
			NextElementUuid:          "",
			FirstChildElementUuid:    "",
			ParentElementUuid:        "",
			TestCaseModelElementType: immatureElement.TestCaseModelElementType,
		}

		// Add the raw mature element to the map of mature elements model for the component. Use Immature UUID+"_mature" as temp-id in Map
		tempMatureElementModel.MatureElementMap[immatureElementUuid+matureIndicator] = newRawMatureComponent
	}

	// Loop all ImmatureElements in Component and copy references between a raw mature version of each immature element
	for immatureElementUuid, immatureElement := range immatureElementToSwapIn.ImmatureElementMap {

		// Extract 'raw' mature element
		rawMatureElement := tempMatureElementModel.MatureElementMap[immatureElementUuid+matureIndicator]

		// Extract immature uuid's
		previousImmatureElementUuid := immatureElement.PreviousElementUuid
		nextImmatureElementUuid := immatureElement.NextElementUuid
		firstChildImmatureElementUuid := immatureElement.FirstChildElementUuid
		parentImmatureElementUuid := immatureElement.ParentElementUuid

		// Add UUIDs to rawMatureElement
		rawMatureElement.PreviousElementUuid = tempMatureElementModel.MatureElementMap[previousImmatureElementUuid+matureIndicator].MatureElementUuid
		rawMatureElement.NextElementUuid = tempMatureElementModel.MatureElementMap[nextImmatureElementUuid+matureIndicator].MatureElementUuid
		rawMatureElement.FirstChildElementUuid = tempMatureElementModel.MatureElementMap[firstChildImmatureElementUuid+matureIndicator].MatureElementUuid
		rawMatureElement.ParentElementUuid = tempMatureElementModel.MatureElementMap[parentImmatureElementUuid+matureIndicator].MatureElementUuid

		// Add the raw mature element to the map of mature elements model for the component
		matureElementModel.MatureElementMap[rawMatureElement.MatureElementUuid] = rawMatureElement

		// Add DropZone colors to Mature model when it is a TestInstruction
		if rawMatureElement.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			rawMatureElement.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE {

			matureElementModel.ChosenDropZoneColor = immatureElementToSwapIn.ChosenDropZoneColor

		}

		// If this element is the top element then this element's uuid as first element
		// A top element must be a TI or TIC
		if rawMatureElement.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			rawMatureElement.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE ||
			rawMatureElement.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER ||
			rawMatureElement.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE {

			// A top element doesn't have any previous, next or parent element
			if rawMatureElement.PreviousElementUuid == rawMatureElement.NextElementUuid &&
				rawMatureElement.NextElementUuid == rawMatureElement.ParentElementUuid &&
				rawMatureElement.ParentElementUuid == rawMatureElement.MatureElementUuid {

				// Top element
				matureElementModel.FirstElementUuid = rawMatureElement.MatureElementUuid
			}
		}
	}

	// If there is no top element then there is something wrong
	if matureElementModel.FirstElementUuid == "" {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":                      "d2fd3ecf-b7e7-4ddd-b8dd-e75993d7d3df",
			"immatureElementToSwapIn": immatureElementToSwapIn,
		}).Error("there is no top element 'immatureElementToSwapIn'")

		err = errors.New("there is no top element 'immatureElementToSwapIn'")
	}

	// Move UUID and Color for Element
	matureElementModel.ChosenDropZoneUuid = immatureElementToSwapIn.ChosenDropZoneUuid
	matureElementModel.ChosenDropZoneName = immatureElementToSwapIn.ChosenDropZoneName
	matureElementModel.ChosenDropZoneColor = immatureElementToSwapIn.ChosenDropZoneColor

	return matureElementModel, err
}
