package commandAndRuleEngine

import (
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	"fmt"
	"github.com/google/uuid"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) verifySwapRuleAndConvertIntoMatureComponentElementModel(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct, ruleNameToVerify string) (matureElementToSwapIn testCaseModel.MatureElementStruct, err error) {

	// Verify Rules before start swapping
	canBeSwapped, _, matchedComplexRule, err := commandAndRuleEngine.verifyIfElementCanBeSwapped(testCaseUuid, uuidToSwapOut)

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
	currentTestCaseModel, existsInMap := commandAndRuleEngine.testcases.TestCases[testCaseUuid]
	if existsInMap == false {
		errorId := "37fcb025-f91b-4c51-aac3-b0f50fba7de5"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))
	}

	// If Cut-command is in progress then
	if currentTestCaseModel.CutCommandInitiated == false {

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
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC(X)			B0					n 		B0								B1-n-B1					TCRuleSwap101
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeTCRuleSwap101(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (err error) {

	matureElementToSwapIn, err := commandAndRuleEngine.verifySwapRuleAndConvertIntoMatureComponentElementModel(testCaseUuid, uuidToSwapOut, immatureElementToSwapIn, TCRuleSwap101)

	// Couldn't convert immature element component into mature element component
	if err != nil {
		return err
	}

	currentTestCase, existsInMap := commandAndRuleEngine.testcases.TestCases[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all testcases")

		return err
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
	currentTestCase.TestCaseModelMap[newPreviousB1fBond.MatureElementUuid] = newPreviousB1fBond
	currentTestCase.TestCaseModelMap[newNextB1lBond.MatureElementUuid] = newNextB1lBond

	// Set First Element
	currentTestCase.FirstElementUuid = newPreviousB1fBond.MatureElementUuid

	// Add 'matureElementToSwapIn' to TestCase Element Model
	for elementUuid, element := range matureElementToSwapIn.MatureElementMap {

		currentTestCase.TestCaseModelMap[elementUuid] = element

	}

	// Delete old element to be swapped out
	delete(currentTestCase.TestCaseModelMap, uuidToSwapOut)

	// If there are no errors then save the TestCase back into map of all TestCases
	if err == nil {
		commandAndRuleEngine.testcases.TestCases[testCaseUuid] = currentTestCase
	}

	return err
}

// TCRuleSwap102
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B10					n		TIC(B10)						TIC(B11f-n-B11l)		TCRuleSwap102
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeTCRuleSwap102(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (err error) {

	matureElementToSwapIn, err := commandAndRuleEngine.verifySwapRuleAndConvertIntoMatureComponentElementModel(testCaseUuid, uuidToSwapOut, immatureElementToSwapIn, TCRuleSwap102)

	// Couldn't convert immature element component into mature element component
	if err != nil {
		return err
	}

	currentTestCase, existsInMap := commandAndRuleEngine.testcases.TestCases[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all testcases")

		return err
	}

	// Extract parent-TIC to element to swap out
	elementToSwapOut, _ := currentTestCase.TestCaseModelMap[uuidToSwapOut]
	parentElementUuid := elementToSwapOut.ParentElementUuid
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
	parentElement.FirstChildElementUuid = newPreviousB11fBond.MatureElementUuid

	// Save updated parent element back to TestCase model
	currentTestCase.TestCaseModelMap[parentElementUuid] = parentElement

	// Add new Bonds to TestCase Element Model
	currentTestCase.TestCaseModelMap[newPreviousB11fBond.MatureElementUuid] = newPreviousB11fBond
	currentTestCase.TestCaseModelMap[newNextB11lBond.MatureElementUuid] = newNextB11lBond

	// Add 'matureElementToSwapIn' to TestCase Element Model
	for elementUuid, element := range matureElementToSwapIn.MatureElementMap {

		currentTestCase.TestCaseModelMap[elementUuid] = element

	}

	// Delete old element to be swapped out
	delete(currentTestCase.TestCaseModelMap, uuidToSwapOut)

	// If there are no errors then save the TestCase back into map of all TestCases
	if err == nil {
		commandAndRuleEngine.testcases.TestCases[testCaseUuid] = currentTestCase
	}

	return err
}

// TCRuleSwap103
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B11f				n		TIC(B11f-X)						TIC(B11f-n-B12-X)		TCRuleSwap103
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeTCRuleSwap103(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (err error) {

	matureElementToSwapIn, err := commandAndRuleEngine.verifySwapRuleAndConvertIntoMatureComponentElementModel(testCaseUuid, uuidToSwapOut, immatureElementToSwapIn, TCRuleSwap103)

	// Couldn't convert immature element component into mature element component
	if err != nil {
		return err
	}

	currentTestCase, existsInMap := commandAndRuleEngine.testcases.TestCases[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all testcases")

		return err
	}

	// Extract parent-TIC to element to swap out
	elementToSwapOut, _ := currentTestCase.TestCaseModelMap[uuidToSwapOut]
	parentElementUuid := elementToSwapOut.ParentElementUuid
	nextElement := currentTestCase.TestCaseModelMap[elementToSwapOut.NextElementUuid]

	// Create the new Bonds
	newB12Bond := commandAndRuleEngine.createNewBondB12Element(parentElementUuid)

	// Extract TIC/TI from new element model, same as first element
	topElementInModelToBeSwappedIn := matureElementToSwapIn.MatureElementMap[matureElementToSwapIn.FirstElementUuid]

	// Connect the new structure
	elementToSwapOut.NextElementUuid = topElementInModelToBeSwappedIn.MatureElementUuid

	topElementInModelToBeSwappedIn.PreviousElementUuid = elementToSwapOut.MatureElementUuid
	topElementInModelToBeSwappedIn.NextElementUuid = newB12Bond.MatureElementUuid
	topElementInModelToBeSwappedIn.ParentElementUuid = parentElementUuid

	newB12Bond.PreviousElementUuid = topElementInModelToBeSwappedIn.MatureElementUuid
	newB12Bond.NextElementUuid = nextElement.MatureElementUuid

	nextElement.PreviousElementUuid = newB12Bond.MatureElementUuid

	// Add updated element back to 'matureElementToSwapIn'
	matureElementToSwapIn.MatureElementMap[topElementInModelToBeSwappedIn.MatureElementUuid] = topElementInModelToBeSwappedIn

	// Update elements in TestCaseModel
	currentTestCase.TestCaseModelMap[elementToSwapOut.MatureElementUuid] = elementToSwapOut
	currentTestCase.TestCaseModelMap[nextElement.MatureElementUuid] = nextElement

	// Add new Bonds to TestCase Element Model
	currentTestCase.TestCaseModelMap[newB12Bond.MatureElementUuid] = newB12Bond

	// Add 'matureElementToSwapIn' to TestCase Element Model
	for elementUuid, element := range matureElementToSwapIn.MatureElementMap {

		currentTestCase.TestCaseModelMap[elementUuid] = element

	}

	// If there are no errors then save the TestCase back into map of all TestCases
	if err == nil {
		commandAndRuleEngine.testcases.TestCases[testCaseUuid] = currentTestCase
	}

	return err
}

// TCRuleSwap104
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B11l				n		TIC(X-B11l)						TIC(X-B12-n-B11l)		TCRuleSwap104
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeTCRuleSwap104(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (err error) {

	matureElementToSwapIn, err := commandAndRuleEngine.verifySwapRuleAndConvertIntoMatureComponentElementModel(testCaseUuid, uuidToSwapOut, immatureElementToSwapIn, TCRuleSwap104)

	// Couldn't convert immature element component into mature element component
	if err != nil {
		return err
	}

	currentTestCase, existsInMap := commandAndRuleEngine.testcases.TestCases[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all testcases")

		return err
	}

	// Extract parent-TIC to element to swap out
	elementToSwapOut, _ := currentTestCase.TestCaseModelMap[uuidToSwapOut]
	parentElementUuid := elementToSwapOut.ParentElementUuid
	previousElement := currentTestCase.TestCaseModelMap[elementToSwapOut.PreviousElementUuid]

	// Create the new Bonds
	newB12Bond := commandAndRuleEngine.createNewBondB12Element(parentElementUuid)

	// Extract TIC/TI from new element model, same as first element
	topElementInModelToBeSwappedIn := matureElementToSwapIn.MatureElementMap[matureElementToSwapIn.FirstElementUuid]

	// Connect the new structure
	previousElement.NextElementUuid = newB12Bond.MatureElementUuid

	newB12Bond.PreviousElementUuid = previousElement.MatureElementUuid
	newB12Bond.NextElementUuid = topElementInModelToBeSwappedIn.MatureElementUuid

	topElementInModelToBeSwappedIn.PreviousElementUuid = newB12Bond.MatureElementUuid
	topElementInModelToBeSwappedIn.NextElementUuid = elementToSwapOut.MatureElementUuid
	topElementInModelToBeSwappedIn.ParentElementUuid = parentElementUuid

	elementToSwapOut.PreviousElementUuid = topElementInModelToBeSwappedIn.MatureElementUuid

	// Add updated element back to 'matureElementToSwapIn'
	matureElementToSwapIn.MatureElementMap[topElementInModelToBeSwappedIn.MatureElementUuid] = topElementInModelToBeSwappedIn

	// Update elements in TestCaseModel
	currentTestCase.TestCaseModelMap[elementToSwapOut.MatureElementUuid] = elementToSwapOut
	currentTestCase.TestCaseModelMap[previousElement.MatureElementUuid] = previousElement

	// Add new Bonds to TestCase Element Model
	currentTestCase.TestCaseModelMap[newB12Bond.MatureElementUuid] = newB12Bond

	// Add 'matureElementToSwapIn' to TestCase Element Model
	for elementUuid, element := range matureElementToSwapIn.MatureElementMap {

		currentTestCase.TestCaseModelMap[elementUuid] = element

	}

	// If there are no errors then save the TestCase back into map of all TestCases
	if err == nil {
		commandAndRuleEngine.testcases.TestCases[testCaseUuid] = currentTestCase
	}

	return err
}

// TCRuleSwap105
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B12					n		X-B12-X							X-B12-n-B12-X			TCRuleSwap105
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeTCRuleSwap105(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (err error) {

	matureElementToSwapIn, err := commandAndRuleEngine.verifySwapRuleAndConvertIntoMatureComponentElementModel(testCaseUuid, uuidToSwapOut, immatureElementToSwapIn, TCRuleSwap105)

	// Couldn't convert immature element component into mature element component
	if err != nil {
		return err
	}

	currentTestCase, existsInMap := commandAndRuleEngine.testcases.TestCases[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all testcases")

		return err
	}

	// Extract parent-TIC to element to swap out
	elementToSwapOut, _ := currentTestCase.TestCaseModelMap[uuidToSwapOut]
	parentElementUuid := elementToSwapOut.ParentElementUuid
	previousElement := currentTestCase.TestCaseModelMap[elementToSwapOut.PreviousElementUuid]

	// Create the new Bonds
	newB12Bond := commandAndRuleEngine.createNewBondB12Element(parentElementUuid)

	// Extract TIC/TI from new element model, same as first element
	topElementInModelToBeSwappedIn := matureElementToSwapIn.MatureElementMap[matureElementToSwapIn.FirstElementUuid]

	// Connect the new structure
	previousElement.NextElementUuid = newB12Bond.MatureElementUuid

	newB12Bond.PreviousElementUuid = previousElement.MatureElementUuid
	newB12Bond.NextElementUuid = topElementInModelToBeSwappedIn.MatureElementUuid

	topElementInModelToBeSwappedIn.PreviousElementUuid = newB12Bond.MatureElementUuid
	topElementInModelToBeSwappedIn.NextElementUuid = elementToSwapOut.MatureElementUuid
	topElementInModelToBeSwappedIn.ParentElementUuid = parentElementUuid

	elementToSwapOut.PreviousElementUuid = topElementInModelToBeSwappedIn.MatureElementUuid

	// Add updated element back to 'matureElementToSwapIn'
	matureElementToSwapIn.MatureElementMap[topElementInModelToBeSwappedIn.MatureElementUuid] = topElementInModelToBeSwappedIn

	// Update elements in TestCaseModel
	currentTestCase.TestCaseModelMap[elementToSwapOut.MatureElementUuid] = elementToSwapOut
	currentTestCase.TestCaseModelMap[previousElement.MatureElementUuid] = previousElement

	// Add new Bonds to TestCase Element Model
	currentTestCase.TestCaseModelMap[newB12Bond.MatureElementUuid] = newB12Bond

	// Add 'matureElementToSwapIn' to TestCase Element Model
	for elementUuid, element := range matureElementToSwapIn.MatureElementMap {

		currentTestCase.TestCaseModelMap[elementUuid] = element

	}

	// If there are no errors then save the TestCase back into map of all TestCases
	if err == nil {
		commandAndRuleEngine.testcases.TestCases[testCaseUuid] = currentTestCase
	}

	return err
}

// TCRuleSwap106
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B10x*				n		TIC(B10*x*)						TIC(B11x-n-B11x)		TCRuleSwap106
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeTCRuleSwap106(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (err error) {

	matureElementToSwapIn, err := commandAndRuleEngine.verifySwapRuleAndConvertIntoMatureComponentElementModel(testCaseUuid, uuidToSwapOut, immatureElementToSwapIn, TCRuleSwap106)

	// Couldn't convert immature element component into mature element component
	if err != nil {
		return err
	}

	currentTestCase, existsInMap := commandAndRuleEngine.testcases.TestCases[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all testcases")

		return err
	}

	// Extract parent-TIC to element to swap out
	elementToSwapOut, _ := currentTestCase.TestCaseModelMap[uuidToSwapOut]
	parentElementUuid := elementToSwapOut.ParentElementUuid
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
	parentElement.FirstChildElementUuid = newPreviousB11fxBond.MatureElementUuid

	// Save updated parent element back to TestCase model
	currentTestCase.TestCaseModelMap[parentElementUuid] = parentElement

	// Add new Bonds to TestCase Element Model
	currentTestCase.TestCaseModelMap[newPreviousB11fxBond.MatureElementUuid] = newPreviousB11fxBond
	currentTestCase.TestCaseModelMap[newNextB11lxBond.MatureElementUuid] = newNextB11lxBond

	// Add 'matureElementToSwapIn' to TestCase Element Model
	for elementUuid, element := range matureElementToSwapIn.MatureElementMap {

		currentTestCase.TestCaseModelMap[elementUuid] = element

	}

	// If there are no errors then save the TestCase back into map of all TestCases
	if err == nil {
		commandAndRuleEngine.testcases.TestCases[testCaseUuid] = currentTestCase
	}

	// Delete old element to be swapped out
	delete(currentTestCase.TestCaseModelMap, uuidToSwapOut)

	return err
}

// TCRuleSwap107
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B10*x				n		TIC(B10*x)						TIC(B11x-n-B11)			TCRuleSwap107
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeTCRuleSwap107(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (err error) {

	matureElementToSwapIn, err := commandAndRuleEngine.verifySwapRuleAndConvertIntoMatureComponentElementModel(testCaseUuid, uuidToSwapOut, immatureElementToSwapIn, TCRuleSwap107)

	// Couldn't convert immature element component into mature element component
	if err != nil {
		return err
	}

	currentTestCase, existsInMap := commandAndRuleEngine.testcases.TestCases[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all testcases")

		return err
	}

	// Extract parent-TIC to element to swap out
	elementToSwapOut, _ := currentTestCase.TestCaseModelMap[uuidToSwapOut]
	parentElementUuid := elementToSwapOut.ParentElementUuid
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
	parentElement.FirstChildElementUuid = newPreviousB11fxBond.MatureElementUuid

	// Save updated parent element back to TestCase model
	currentTestCase.TestCaseModelMap[parentElementUuid] = parentElement

	// Add new Bonds to TestCase Element Model
	currentTestCase.TestCaseModelMap[newPreviousB11fxBond.MatureElementUuid] = newPreviousB11fxBond
	currentTestCase.TestCaseModelMap[newNextB11lBond.MatureElementUuid] = newNextB11lBond

	// Add 'matureElementToSwapIn' to TestCase Element Model
	for elementUuid, element := range matureElementToSwapIn.MatureElementMap {

		currentTestCase.TestCaseModelMap[elementUuid] = element

	}

	// Delete old element to be swapped out
	delete(currentTestCase.TestCaseModelMap, uuidToSwapOut)

	// If there are no errors then save the TestCase back into map of all TestCases
	if err == nil {
		commandAndRuleEngine.testcases.TestCases[testCaseUuid] = currentTestCase
	}

	return err
}

// TCRuleSwap108
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B10x*				n		TIC(B10x*)						TIC(B11-n-B11x)			TCRuleSwap108
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeTCRuleSwap108(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (err error) {

	matureElementToSwapIn, err := commandAndRuleEngine.verifySwapRuleAndConvertIntoMatureComponentElementModel(testCaseUuid, uuidToSwapOut, immatureElementToSwapIn, TCRuleSwap108)

	// Couldn't convert immature element component into mature element component
	if err != nil {
		return err
	}

	currentTestCase, existsInMap := commandAndRuleEngine.testcases.TestCases[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all testcases")

		return err
	}

	// Extract parent-TIC to element to swap out
	elementToSwapOut, _ := currentTestCase.TestCaseModelMap[uuidToSwapOut]
	parentElementUuid := elementToSwapOut.ParentElementUuid
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
	parentElement.FirstChildElementUuid = newPreviousB11fBond.MatureElementUuid

	// Save updated parent element back to TestCase model
	currentTestCase.TestCaseModelMap[parentElementUuid] = parentElement

	// Add new Bonds to TestCase Element Model
	currentTestCase.TestCaseModelMap[newPreviousB11fBond.MatureElementUuid] = newPreviousB11fBond
	currentTestCase.TestCaseModelMap[newNextB11lxBond.MatureElementUuid] = newNextB11lxBond

	// Add 'matureElementToSwapIn' to TestCase Element Model
	for elementUuid, element := range matureElementToSwapIn.MatureElementMap {

		currentTestCase.TestCaseModelMap[elementUuid] = element

	}

	// Delete old element to be swapped out
	delete(currentTestCase.TestCaseModelMap, uuidToSwapOut)

	// If there are no errors then save the TestCase back into map of all TestCases
	if err == nil {
		commandAndRuleEngine.testcases.TestCases[testCaseUuid] = currentTestCase
	}

	return err
}

// Transforms a immature element model into a mature element model.
// This means that new UUIDs are created for each element in the component
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) transformImmatureElementModelIntoMatureElementModel(immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (matureElementModel testCaseModel.MatureElementStruct, err error) {

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
	return matureElementModel, err
}
