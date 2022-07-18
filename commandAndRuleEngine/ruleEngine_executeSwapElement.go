package commandAndRuleEngine

import (
	"errors"
	"github.com/google/uuid"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) verifySwapRuleAndConvertIntoMatureComponentElementModel(uuidToSwapOut string, immatureElementToSwapIn *immatureElementStruct, ruleNameToVerify string) (matureElementToSwapIn matureElementStruct, err error) {

	// Verify Rules before start swapping
	canBeSwapped, _, matchedComplexRule, err := commandAndRuleEngine.verifyIfElementCanBeSwapped(uuidToSwapOut)

	// Can't be swapped in
	if canBeSwapped == false ||
		matchedComplexRule != ruleNameToVerify {

		// So exit
		return matureElementStruct{}, err
	}

	// Transform ImmatureElementModel into a MatureElementModel
	matureElementToSwapIn, err = commandAndRuleEngine.transformImmatureElementModelIntoMatureElementModel(immatureElementToSwapIn)

	return matureElementToSwapIn, err

}

// TCRuleSwap101
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC(X)			B0					n 		B0								B1-n-B1					TCRuleSwap101
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeTCRuleSwap101(uuidToSwapOut string, immatureElementToSwapIn *immatureElementStruct) (err error) {

	matureElementToSwapIn, err := commandAndRuleEngine.verifySwapRuleAndConvertIntoMatureComponentElementModel(uuidToSwapOut, immatureElementToSwapIn, TCRuleSwap101)

	// Couldn't convert immature eleObjectment component into mature element component
	if err != nil {
		return err
	}

	// Create the Bonds connecting the TIC
	newPreviousB1Bond := commandAndRuleEngine.createNewBondB11fElement("")
	newNextB1Bond := commandAndRuleEngine.createNewBondB11lElement("")

	// Extract TIC from new element model
	newTopElement := matureElementToSwapIn.matureElementMap[matureElementToSwapIn.firstElementUuid]

	// Connect the new structure
	newPreviousB1Bond.NextElementUuid = newTopElement.MatureElementUuid

	newTopElement.PreviousElementUuid = newPreviousB1Bond.MatureElementUuid
	newTopElement.NextElementUuid = newNextB1Bond.MatureElementUuid

	newNextB1Bond.PreviousElementUuid = newTopElement.MatureElementUuid

	// Add updated element back to 'matureElementToSwapIn'
	matureElementToSwapIn.matureElementMap[matureElementToSwapIn.firstElementUuid] = newTopElement

	// Add new Bonds to TestCase Element Model
	commandAndRuleEngine.testcaseModel.TestCaseModelMap[newPreviousB1Bond.MatureElementUuid] = newPreviousB1Bond
	commandAndRuleEngine.testcaseModel.TestCaseModelMap[newNextB1Bond.MatureElementUuid] = newNextB1Bond

	// Set First Element
	commandAndRuleEngine.testcaseModel.FirstElementUuid = newPreviousB1Bond.MatureElementUuid

	// Add 'matureElementToSwapIn' to TestCase Element Model
	for elementUuid, element := range matureElementToSwapIn.matureElementMap {

		commandAndRuleEngine.testcaseModel.TestCaseModelMap[elementUuid] = element

	}

	// Delete old element to be swapped out
	delete(commandAndRuleEngine.testcaseModel.TestCaseModelMap, uuidToSwapOut)

	return err
}

// TCRuleSwap102
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B10					n		TIC(B10)						TIC(B11f-n-B11l)		TCRuleSwap102
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeTCRuleSwap102(uuidToSwapOut string, immatureElementToSwapIn *immatureElementStruct) (err error) {

	return err
}

// TCRuleSwap103
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B11f				n		TIC(B11f-X)						TIC(B11f-n-B12-X)		TCRuleSwap103
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeTCRuleSwap103(uuidToSwapOut string, immatureElementToSwapIn *immatureElementStruct) (err error) {

	return err
}

// TCRuleSwap104
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B11l				n		TIC(X-B11l)						TIC(X-B12-n-B11l)		TCRuleSwap104
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeTCRuleSwap104(uuidToSwapOut string, immatureElementToSwapIn *immatureElementStruct) (err error) {

	return err
}

// TCRuleSwap105
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B12					n		X-B12-X							X-B12-n-B12-X			TCRuleSwap105
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeTCRuleSwap105(uuidToSwapOut string, immatureElementToSwapIn *immatureElementStruct) (err error) {

	return err
}

// TCRuleSwap106
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B10x*				n		TIC(B10*x*)						TIC(B11x-n-B11x)		TCRuleSwap106
//	n=TIC or TIC(X)		B10*x				n		TIC(B10*x)						TIC(B11x-n-B11)			TCRuleSwap107
//	n=TIC or TIC(X)		B10x*				n		TIC(B10x*)						TIC(B11-n-B11x)			TCRuleSwap108
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeTCRuleSwap106(uuidToSwapOut string, immatureElementToSwapIn *immatureElementStruct) (err error) {

	return err
}

// TCRuleSwap107
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B10*x				n		TIC(B10*x)						TIC(B11x-n-B11)			TCRuleSwap107
//	n=TIC or TIC(X)		B10x*				n		TIC(B10x*)						TIC(B11-n-B11x)			TCRuleSwap108
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeTCRuleSwap107(uuidToSwapOut string, immatureElementToSwapIn *immatureElementStruct) (err error) {

	return err
}

// TCRuleSwap108
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B10x*				n		TIC(B10x*)						TIC(B11-n-B11x)			TCRuleSwap108
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeTCRuleSwap108(uuidToSwapOut string, immatureElementToSwapIn *immatureElementStruct) (err error) {

	return err
}

// Transforms a immature element model into a mature element model.
// This means that new UUIDs are created for each element in the component
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) transformImmatureElementModelIntoMatureElementModel(immatureElementToSwapIn *immatureElementStruct) (matureElementModel matureElementStruct, err error) {

	var matureIndicator = "_mature"

	// Create the temp store for matureElementModel and initiate map
	tempMatureElementModel := matureElementStruct{}
	tempMatureElementModel.matureElementMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage)

	// Initiate the matureElementModel map
	matureElementModel.matureElementMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage)

	// Loop all ImmatureElements in Component and create a raw mature version of each immature element
	for immatureElementUuid, immatureElement := range immatureElementToSwapIn.immatureElementMap {

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
		tempMatureElementModel.matureElementMap[immatureElementUuid+matureIndicator] = newRawMatureComponent
	}

	// Loop all ImmatureElements in Component and copy references between a raw mature version of each immature element
	for immatureElementUuid, immatureElement := range immatureElementToSwapIn.immatureElementMap {

		// Extract 'raw' mature element
		rawMatureElement := tempMatureElementModel.matureElementMap[immatureElementUuid+matureIndicator]

		// Extract immature uuid's
		previousImmatureElementUuid := immatureElement.PreviousElementUuid
		nextImmatureElementUuid := immatureElement.NextElementUuid
		firstChildImmatureElementUuid := immatureElement.FirstChildElementUuid
		parentImmatureElementUuid := immatureElement.ParentElementUuid

		// Add UUIDs to rawMatureElement
		rawMatureElement.PreviousElementUuid = tempMatureElementModel.matureElementMap[previousImmatureElementUuid+matureIndicator].MatureElementUuid
		rawMatureElement.NextElementUuid = tempMatureElementModel.matureElementMap[nextImmatureElementUuid+matureIndicator].MatureElementUuid
		rawMatureElement.FirstChildElementUuid = tempMatureElementModel.matureElementMap[firstChildImmatureElementUuid+matureIndicator].MatureElementUuid
		rawMatureElement.ParentElementUuid = tempMatureElementModel.matureElementMap[parentImmatureElementUuid+matureIndicator].MatureElementUuid

		// Add the raw mature element to the map of mature elements model for the component
		matureElementModel.matureElementMap[rawMatureElement.MatureElementUuid] = rawMatureElement

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
				matureElementModel.firstElementUuid = rawMatureElement.MatureElementUuid
			}
		}
	}

	// If there is no top element then there is something wrong
	if matureElementModel.firstElementUuid == "" {
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":                      "d2fd3ecf-b7e7-4ddd-b8dd-e75993d7d3df",
			"immatureElementToSwapIn": immatureElementToSwapIn,
		}).Error("there is no top element 'immatureElementToSwapIn'")

		err = errors.New("there is no top element 'immatureElementToSwapIn'")
	}
	return matureElementModel, err
}
