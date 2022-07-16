package commandAndRuleEngine

import (
	"errors"
	"github.com/google/uuid"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

// TCRuleSwap101
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC(X)			B0					n 		B0								B1-n-B1					TCRuleSwap101
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) executeTCRuleSwap101(uuidToSwapOut string, immatureElementToSwapIn *immatureElementStruct) (err error) {

	// Transform ImmatureElementModel into a MatureElementModel
	matureElementToSwapIn, err := commandAndRuleEngineObject.transformImmatureElementModelIntoMatureElementModel(immatureElementToSwapIn)

	return err
}

// TCRuleSwap102
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B10					n		TIC(B10)						TIC(B11f-n-B11l)		TCRuleSwap102
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) executeTCRuleSwap102(uuidToSwapOut string, immatureElementToSwapIn immatureElementStruct) (err error) {

	return err
}

// TCRuleSwap103
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B11f				n		TIC(B11f-X)						TIC(B11f-n-B12-X)		TCRuleSwap103
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) executeTCRuleSwap103(uuidToSwapOut string, immatureElementToSwapIn immatureElementStruct) (err error) {

	return err
}

// TCRuleSwap104
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B11l				n		TIC(X-B11l)						TIC(X-B12-n-B11l)		TCRuleSwap104
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) executeTCRuleSwap104(uuidToSwapOut string, immatureElementToSwapIn immatureElementStruct) (err error) {

	return err
}

// TCRuleSwap105
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B12					n		X-B12-X							X-B12-n-B12-X			TCRuleSwap105
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) executeTCRuleSwap105(uuidToSwapOut string, immatureElementToSwapIn immatureElementStruct) (err error) {

	return err
}

// TCRuleSwap106
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B10x*				n		TIC(B10*x*)						TIC(B11x-n-B11x)		TCRuleSwap106
//	n=TIC or TIC(X)		B10*x				n		TIC(B10*x)						TIC(B11x-n-B11)			TCRuleSwap107
//	n=TIC or TIC(X)		B10x*				n		TIC(B10x*)						TIC(B11-n-B11x)			TCRuleSwap108
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) executeTCRuleSwap106(uuidToSwapOut string, immatureElementToSwapIn immatureElementStruct) (err error) {

	return err
}

// TCRuleSwap107
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B10*x				n		TIC(B10*x)						TIC(B11x-n-B11)			TCRuleSwap107
//	n=TIC or TIC(X)		B10x*				n		TIC(B10x*)						TIC(B11-n-B11x)			TCRuleSwap108
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) executeTCRuleSwap107(uuidToSwapOut string, immatureElementToSwapIn immatureElementStruct) (err error) {

	return err
}

// TCRuleSwap108
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B10x*				n		TIC(B10x*)						TIC(B11-n-B11x)			TCRuleSwap108
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) executeTCRuleSwap108(uuidToSwapOut string, immatureElementToSwapIn immatureElementStruct) (err error) {

	return err
}

// Transforms a immature element model into a mature element model.
// This means that new UUIDs are created for each element in the component
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) transformImmatureElementModelIntoMatureElementModel(immatureElementToSwapIn *immatureElementStruct) (matureElementModel matureElementStruct, err error) {

	var matureIndicator = "_mature"

	// Create the temp store for matureElementModel
	tempMatureElementModel := matureElementStruct{}

	// Initiate the matureElementModel
	matureElementModel = matureElementStruct{}

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
			TestCaseModelElementType: 0,
		}

		// Add the raw mature element to the map of mature elements model for the component. Use Immature UUID+"_mature" as temp-id in Map
		tempMatureElementModel.matureElementMap[immatureElementUuid+matureIndicator] = newRawMatureComponent
	}

	// Loop all ImmatureElements in Component and copy references between a raw mature version of each immature element
	for immatureElementUuid, immatureElement := range immatureElementToSwapIn.immatureElementMap {

		// Extract 'raw' mature element
		rawMatureElement := matureElementModel.matureElementMap[immatureElementUuid+matureIndicator]

		// Extract immature uuid's
		previousImmatureElementUuid := immatureElement.PreviousElementUuid
		nextImmatureElementUuid := immatureElement.NextElementUuid
		firstChildImmatureElementUuid := immatureElement.FirstChildElementUuid
		parentImmatureElementUuid := immatureElement.FirstChildElementUuid

		// Extract immature type
		immatureTestCaseModelElementType := immatureElement.TestCaseModelElementType

		// Add UUIDs to rawMatureElement
		rawMatureElement.PreviousElementUuid = tempMatureElementModel.matureElementMap[previousImmatureElementUuid+matureIndicator].MatureElementUuid
		rawMatureElement.NextElementUuid = tempMatureElementModel.matureElementMap[nextImmatureElementUuid+matureIndicator].MatureElementUuid
		rawMatureElement.FirstChildElementUuid = tempMatureElementModel.matureElementMap[firstChildImmatureElementUuid+matureIndicator].MatureElementUuid
		rawMatureElement.ParentElementUuid = tempMatureElementModel.matureElementMap[parentImmatureElementUuid+matureIndicator].MatureElementUuid

		// Add the element type to rawMatureElement
		rawMatureElement.TestCaseModelElementType = immatureTestCaseModelElementType

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
	commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
		"id":                      "d2fd3ecf-b7e7-4ddd-b8dd-e75993d7d3df",
		"immatureElementToSwapIn": immatureElementToSwapIn,
	}).Error("there is no top element 'immatureElementToSwapIn'")

	err = errors.New("there is no top element 'immatureElementToSwapIn'")

	return matureElementModel, err
}

// Follow all children, in ImmatureElement-model and create the Mature Element Model instead
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) recursiveImmatureToMatureElementComponentModelCreator(elementsUuid string, immatureElement *immatureElementStruct, matureElement *matureElementStruct) (err error) {

	// Extract current element
	currentElement, existInMap := immatureElement.immatureElementMap[elementsUuid]

	// If the element doesn't exit then there is something really wrong
	if existInMap == false {
		// This shouldn't happen
		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":           "e4215635-9770-4142-b2af-cc12b851c79f",
			"elementsUuid": elementsUuid,
		}).Error(elementsUuid + " could not be found in in map 'immatureElement.immatureElementMap'")

		err = errors.New(elementsUuid + " could not be found in in map 'immatureElement.immatureElementMap'")

		return err
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

	// Create the 'raw 'MatureElement from an ImmatureElement. This element is not yet connected
	newRawMatureComponent := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      currentElement.OriginalElementUuid,
		OriginalElementName:      currentElement.OriginalElementName,
		MatureElementUuid:        uuid.New().String(),
		PreviousElementUuid:      "",
		NextElementUuid:          "",
		FirstChildElementUuid:    "",
		ParentElementUuid:        "",
		TestCaseModelElementType: 0,
	}

	// Add

	return err
}
