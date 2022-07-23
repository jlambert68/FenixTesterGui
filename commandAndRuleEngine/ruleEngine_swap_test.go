package commandAndRuleEngine

import (
	"FenixTesterGui/gui/UnitTestTestData"
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Verify that all UUIDs are correct in TestCaseModel. Meaning that no empty uuid is allowed and they all are correct
func verifyThatThereAreNoZombieElementsInTestCaseModel(testCaseModel *testCaseModel.TestCaseModelStruct) (err error) {

	var allUuidKeys []string

	// Extract all elements by key from TestCaseModel
	for _, elementKey := range testCaseModel.TestCaseModelMap {
		allUuidKeys = append(allUuidKeys, elementKey.MatureElementUuid)
	}

	// Follow the path from "first element and remove the found element from 'allUuidKeys'
	allUuidKeys, err = recursiveZombieElementSearchInTestCaseModel(testCaseModel.FirstElementUuid, allUuidKeys, testCaseModel)

	// If there are elements left in slice then there were zombie elements, which there shouldn't be
	if len(allUuidKeys) != 0 {
		err = errors.New("there existed Zombie elements in 'testCaseModel.TestCaseModelMap', for " + testCaseModel.FirstElementUuid)

		return err
	}

	return err
}

// Verify all children, in TestCaseElement-model and remove the found element from 'allUuidKeys'
func recursiveZombieElementSearchInTestCaseModel(elementsUuid string, allUuidKeys []string, testCaseModel *testCaseModel.TestCaseModelStruct) (processedAllUuidKeys []string, err error) {

	// Extract current element
	currentElement, existInMap := testCaseModel.TestCaseModelMap[elementsUuid]

	// If the element doesn't exit then there is something really wrong
	if existInMap == false {
		// This shouldn't happen
		err = errors.New(elementsUuid + " could not be found in in map 'testCaseModel.TestCaseModelMap'")

		return nil, err
	}

	// Element has child-element then go that path
	if currentElement.FirstChildElementUuid != elementsUuid {
		allUuidKeys, err = recursiveZombieElementSearchInTestCaseModel(currentElement.FirstChildElementUuid, allUuidKeys, testCaseModel)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return nil, err
	}

	// If element has a next-element the go that path
	if currentElement.NextElementUuid != elementsUuid {
		allUuidKeys, err = recursiveZombieElementSearchInTestCaseModel(currentElement.NextElementUuid, allUuidKeys, testCaseModel)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return nil, err
	}

	// Remove current element from "slice of all elements in map"
	tempAallUuidKeys := findElementInSliceAndRemove(&allUuidKeys, elementsUuid)

	processedAllUuidKeys = *tempAallUuidKeys

	return processedAllUuidKeys, err
}

// Verify that a 'B0' can be swapped into 'B1-TIC(B10)-B1'
// TCRuleSwap101
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC(X)			B0					n 		B0								B1-n-B1					TCRuleSwap101
func TestTCRuleSwap101(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate a TestCaseModel
	myTestCaseModel := testCaseModel.TestCaseModelStruct{
		LastLoadedTestCaseModelGRPCMessage: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage{},
		FirstElementUuid:                   "",
		TestCaseModelMap:                   nil,
	}

	myTestCaseModel.TestCaseModelMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage)

	uuidToSwapOut := "4b694f8c-f194-45af-a75e-f2bb3fd350e6"

	b0Bond := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "0a1c1266-f8ad-484a-a76c-59c5fd7fbda5",
		OriginalElementName:      "B0_BOND",
		MatureElementUuid:        uuidToSwapOut,
		PreviousElementUuid:      uuidToSwapOut,
		NextElementUuid:          uuidToSwapOut,
		FirstChildElementUuid:    uuidToSwapOut,
		ParentElementUuid:        uuidToSwapOut,
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND,
	}

	// Add B0-bond to the TestCaseModel-map
	myTestCaseModel.TestCaseModelMap[b0Bond.MatureElementUuid] = b0Bond

	// Set the B0-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = b0Bond.MatureElementUuid

	// Create an Immature Element model for 'TIC(B10)'
	immatureElementModel := immatureElementStruct{
		firstElementUuid:   "",
		immatureElementMap: nil,
	}

	immatureElementModel.immatureElementMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage)

	// Create TIC
	tic := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		ImmatureElementUuid:      "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		PreviousElementUuid:      "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		NextElementUuid:          "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		FirstChildElementUuid:    "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		ParentElementUuid:        "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	// Create B10 in TIC(x)
	b10Bond := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage{
		OriginalElementUuid:      "7be82e83-6048-4c30-b4aa-b68c11037c1d",
		OriginalElementName:      "B10-Bond",
		ImmatureElementUuid:      "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		PreviousElementUuid:      "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		NextElementUuid:          "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		FirstChildElementUuid:    "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		ParentElementUuid:        "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
	}

	// Add the Elements to the Immature Elements Model map
	immatureElementModel.immatureElementMap["d444b8d8-b2fb-4505-ad8e-36bfe89988ab"] = tic
	immatureElementModel.immatureElementMap["ff224d27-5c8a-48b9-ace9-af43245cb35d"] = b10Bond

	// Add first Element ti Immature Element Model
	immatureElementModel.firstElementUuid = "d444b8d8-b2fb-4505-ad8e-36bfe89988ab"

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := commandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		testcaseModel:     &myTestCaseModel,
	}

	// Add needed data for availableBondsMap
	tempAvailableBondsMap := make(map[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage)

	// B1f_BOND
	visibleBondAttributesMessage_AvaialbeBond_B1f_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage_VisibleBondAttributesMessage{
		BondUuid: "0d77690e-f8e2-4942-b532-6b3e26d0b81a",
		BondName: "B1f_BOND",
	}

	basicBondInformationMessage_AvaialbeBond_B1f_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage{
		VisibleBondAttributes: &visibleBondAttributesMessage_AvaialbeBond_B1f_BOND}

	immatureBondsMessage_ImmatureBondMessage_B1f_BOND := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage{
		BasicBondInformation: &basicBondInformationMessage_AvaialbeBond_B1f_BOND}

	tempAvailableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE] = &immatureBondsMessage_ImmatureBondMessage_B1f_BOND

	// B1l_BOND
	visibleBondAttributesMessage_AvaialbeBond_B1l_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage_VisibleBondAttributesMessage{
		BondUuid: "2858d47a-198c-43f3-abe8-abd2a36f6045",
		BondName: "B1l_BOND",
	}

	basicBondInformationMessage_AvaialbeBond_B1l_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage{
		VisibleBondAttributes: &visibleBondAttributesMessage_AvaialbeBond_B1l_BOND}

	immatureBondsMessage_ImmatureBondMessage_B1l_BOND := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage{
		BasicBondInformation: &basicBondInformationMessage_AvaialbeBond_B1l_BOND}

	tempAvailableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE] = &immatureBondsMessage_ImmatureBondMessage_B1l_BOND

	// Add bond-map to 'commandAndRuleEngine.availableBondsMap'
	commandAndRuleEngine.availableBondsMap = tempAvailableBondsMap

	// Execute Swap
	err := commandAndRuleEngine.executeTCRuleSwap101(uuidToSwapOut, &immatureElementModel)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate the result of the swap, 'B1f-TIC(B10)-B1l'
	// 1) Validate B1f
	firstTestCaseModelElementUuid := myTestCaseModel.FirstElementUuid
	firstTestCaseModelElement := myTestCaseModel.TestCaseModelMap[firstTestCaseModelElementUuid]

	correctElement := firstTestCaseModelElement.MatureElementUuid == firstTestCaseModelElement.ParentElementUuid &&
		firstTestCaseModelElement.MatureElementUuid == firstTestCaseModelElement.PreviousElementUuid &&
		firstTestCaseModelElement.MatureElementUuid == firstTestCaseModelElement.FirstChildElementUuid &&
		firstTestCaseModelElement.MatureElementUuid != firstTestCaseModelElement.NextElementUuid &&
		firstTestCaseModelElement.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 2) Validate TIC
	secondTestCaseModelElementUuid := firstTestCaseModelElement.NextElementUuid
	secondTestCaseModelElement := myTestCaseModel.TestCaseModelMap[secondTestCaseModelElementUuid]

	correctElement = secondTestCaseModelElement.MatureElementUuid == secondTestCaseModelElement.ParentElementUuid &&
		secondTestCaseModelElement.MatureElementUuid != secondTestCaseModelElement.PreviousElementUuid &&
		secondTestCaseModelElement.MatureElementUuid != secondTestCaseModelElement.FirstChildElementUuid &&
		secondTestCaseModelElement.MatureElementUuid != secondTestCaseModelElement.NextElementUuid &&
		secondTestCaseModelElement.PreviousElementUuid == firstTestCaseModelElement.MatureElementUuid &&
		secondTestCaseModelElement.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 3) Validate B1l
	thirdTestCaseModelElementUuid := secondTestCaseModelElement.NextElementUuid
	thirdTestCaseModelElement := myTestCaseModel.TestCaseModelMap[thirdTestCaseModelElementUuid]

	correctElement = thirdTestCaseModelElement.MatureElementUuid == thirdTestCaseModelElement.ParentElementUuid &&
		thirdTestCaseModelElement.MatureElementUuid != thirdTestCaseModelElement.PreviousElementUuid &&
		thirdTestCaseModelElement.MatureElementUuid == thirdTestCaseModelElement.FirstChildElementUuid &&
		thirdTestCaseModelElement.MatureElementUuid == thirdTestCaseModelElement.NextElementUuid &&
		thirdTestCaseModelElement.PreviousElementUuid == secondTestCaseModelElement.MatureElementUuid &&
		thirdTestCaseModelElement.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 4) Validate B10
	fourthTestCaseModelElementUuid := secondTestCaseModelElement.FirstChildElementUuid
	fourthTestCaseModelElement := myTestCaseModel.TestCaseModelMap[fourthTestCaseModelElementUuid]

	correctElement = fourthTestCaseModelElement.MatureElementUuid != fourthTestCaseModelElement.ParentElementUuid &&
		fourthTestCaseModelElement.MatureElementUuid == fourthTestCaseModelElement.PreviousElementUuid &&
		fourthTestCaseModelElement.MatureElementUuid == fourthTestCaseModelElement.FirstChildElementUuid &&
		fourthTestCaseModelElement.MatureElementUuid == fourthTestCaseModelElement.NextElementUuid &&
		fourthTestCaseModelElement.ParentElementUuid == secondTestCaseModelElement.MatureElementUuid &&
		fourthTestCaseModelElement.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = verifyThatThereAreNoZombieElementsInTestCaseModel(&myTestCaseModel)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

}

// TCRuleSwap102
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B10					n		TIC(B10)						TIC(B11f-n-B11l)		TCRuleSwap102
func TestTCRuleSwap102(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate a TestCaseModel
	myTestCaseModel := testCaseModel.TestCaseModelStruct{
		LastLoadedTestCaseModelGRPCMessage: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage{},
		FirstElementUuid:                   "",
		TestCaseModelMap:                   nil,
	}

	myTestCaseModel.TestCaseModelMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage)

	uuidToSwapOut := "4b694f8c-f194-45af-a75e-f2bb3fd350e6"

	tc_b1f := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "cff02b54-ee53-47d1-94d6-48dc470073a3",
		OriginalElementName:      "B1l_BOND",
		MatureElementUuid:        "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		PreviousElementUuid:      "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		NextElementUuid:          "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		FirstChildElementUuid:    "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		ParentElementUuid:        "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE,
	}

	tc_tic := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		MatureElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		PreviousElementUuid:      "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		NextElementUuid:          "79a6702d-8370-446c-b001-d60494eca6fa",
		FirstChildElementUuid:    uuidToSwapOut,
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_b1l := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c10bfd8f-2786-480c-9a71-bad9ec7bc582",
		OriginalElementName:      "B1l_BOND",
		MatureElementUuid:        "79a6702d-8370-446c-b001-d60494eca6fa",
		PreviousElementUuid:      "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		NextElementUuid:          "79a6702d-8370-446c-b001-d60494eca6fa",
		FirstChildElementUuid:    "79a6702d-8370-446c-b001-d60494eca6fa",
		ParentElementUuid:        "79a6702d-8370-446c-b001-d60494eca6fa",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE,
	}

	tc_b10 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "4f348a88-bb9a-4852-80de-3141687db11d",
		OriginalElementName:      "B10_Bond",
		MatureElementUuid:        uuidToSwapOut,
		PreviousElementUuid:      uuidToSwapOut,
		NextElementUuid:          uuidToSwapOut,
		FirstChildElementUuid:    uuidToSwapOut,
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
	}

	// Set up TestCaseModel-map
	myTestCaseModel.TestCaseModelMap[tc_b1f.MatureElementUuid] = tc_b1f
	myTestCaseModel.TestCaseModelMap[tc_tic.MatureElementUuid] = tc_tic
	myTestCaseModel.TestCaseModelMap[tc_b1l.MatureElementUuid] = tc_b1l
	myTestCaseModel.TestCaseModelMap[tc_b10.MatureElementUuid] = tc_b10

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_b1f.MatureElementUuid

	// Create an Immature Element model for 'TIC(B10)'
	immatureElementModel := immatureElementStruct{
		firstElementUuid:   "",
		immatureElementMap: nil,
	}

	immatureElementModel.immatureElementMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage)

	// Create TIC
	tic := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		ImmatureElementUuid:      "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		PreviousElementUuid:      "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		NextElementUuid:          "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		FirstChildElementUuid:    "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		ParentElementUuid:        "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	// Create B10 in TIC(x)
	b10Bond := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage{
		OriginalElementUuid:      "7be82e83-6048-4c30-b4aa-b68c11037c1d",
		OriginalElementName:      "B10-Bond",
		ImmatureElementUuid:      "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		PreviousElementUuid:      "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		NextElementUuid:          "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		FirstChildElementUuid:    "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		ParentElementUuid:        "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
	}

	// Add the Elements to the Immature Elements Model map
	immatureElementModel.immatureElementMap["d444b8d8-b2fb-4505-ad8e-36bfe89988ab"] = tic
	immatureElementModel.immatureElementMap["ff224d27-5c8a-48b9-ace9-af43245cb35d"] = b10Bond

	// Add first Element ti Immature Element Model
	immatureElementModel.firstElementUuid = "d444b8d8-b2fb-4505-ad8e-36bfe89988ab"

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := commandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		testcaseModel:     &myTestCaseModel,
	}

	// Add needed data for availableBondsMap
	tempAvailableBondsMap := make(map[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage)

	// B11f_BOND
	visibleBondAttributesMessage_AvaialbeBond_B11f_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage_VisibleBondAttributesMessage{
		BondUuid: "0d77690e-f8e2-4942-b532-6b3e26d0b81a",
		BondName: "B11f_BOND",
	}

	basicBondInformationMessage_AvaialbeBond_B11f_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage{
		VisibleBondAttributes: &visibleBondAttributesMessage_AvaialbeBond_B11f_BOND}

	immatureBondsMessage_ImmatureBondMessage_B11f_BOND := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage{
		BasicBondInformation: &basicBondInformationMessage_AvaialbeBond_B11f_BOND}

	tempAvailableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND] = &immatureBondsMessage_ImmatureBondMessage_B11f_BOND

	// B11l_BOND
	visibleBondAttributesMessage_AvaialbeBond_B11l_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage_VisibleBondAttributesMessage{
		BondUuid: "61f98266-b9b1-4958-9f90-1d0d7f17aafc",
		BondName: "B11l_BOND",
	}

	basicBondInformationMessage_AvaialbeBond_B11l_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage{
		VisibleBondAttributes: &visibleBondAttributesMessage_AvaialbeBond_B11l_BOND}

	immatureBondsMessage_ImmatureBondMessage_B11l_BOND := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage{
		BasicBondInformation: &basicBondInformationMessage_AvaialbeBond_B11l_BOND}

	tempAvailableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND] = &immatureBondsMessage_ImmatureBondMessage_B11l_BOND

	// Add bond-map to 'commandAndRuleEngine.availableBondsMap'
	commandAndRuleEngine.availableBondsMap = tempAvailableBondsMap

	// Execute Swap
	err := commandAndRuleEngine.executeTCRuleSwap102(uuidToSwapOut, &immatureElementModel)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate the result of the swap, 'B1f-TIC(B11f-TIC(B10)-B11l)-B1l'
	// 1) Validate B1f (1)
	testCaseModelElementUuid_1 := myTestCaseModel.FirstElementUuid
	testCaseModelElement_1 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_1]

	correctElement := testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.ParentElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.PreviousElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.FirstChildElementUuid &&
		testCaseModelElement_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 2) Validate TIC (2)
	testCaseModelElementUuid_2 := testCaseModelElement_1.NextElementUuid
	testCaseModelElement_2 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2.ParentElementUuid &&
		testCaseModelElement_2.PreviousElementUuid == testCaseModelElement_1.MatureElementUuid &&
		testCaseModelElement_1.NextElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 3) Validate B1l (3)
	testCaseModelElementUuid_3 := testCaseModelElement_2.NextElementUuid
	testCaseModelElement_3 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_3]

	correctElement = testCaseModelElement_3.MatureElementUuid == testCaseModelElement_3.ParentElementUuid &&
		testCaseModelElement_3.NextElementUuid == testCaseModelElement_3.MatureElementUuid &&
		testCaseModelElement_3.FirstChildElementUuid == testCaseModelElement_3.MatureElementUuid &&
		testCaseModelElement_3.PreviousElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2.NextElementUuid == testCaseModelElement_3.MatureElementUuid &&
		testCaseModelElement_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 4) Validate B11f (2.1)
	testCaseModelElementUuid_2_1 := testCaseModelElement_2.FirstChildElementUuid
	testCaseModelElement_2_1 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_1]

	correctElement = testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2.FirstChildElementUuid &&
		testCaseModelElement_2_1.ParentElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 5) Validate TIC (2.2)
	testCaseModelElementUuid_2_2 := testCaseModelElement_2_1.NextElementUuid
	testCaseModelElement_2_2 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_2.ParentElementUuid &&
		testCaseModelElement_2_2.PreviousElementUuid == testCaseModelElement_2_1.MatureElementUuid &&
		testCaseModelElement_2_1.NextElementUuid == testCaseModelElement_2_2.MatureElementUuid &&
		testCaseModelElement_2_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 6) Validate B1ll (2.3)
	testCaseModelElementUuid_2_3 := testCaseModelElement_2_2.NextElementUuid
	testCaseModelElement_2_3 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_3]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_3.ParentElementUuid &&
		testCaseModelElement_2_3.FirstChildElementUuid == testCaseModelElement_2_3.MatureElementUuid &&
		testCaseModelElement_2_3.PreviousElementUuid == testCaseModelElement_2_2.MatureElementUuid &&
		testCaseModelElement_2_2.NextElementUuid == testCaseModelElement_2_3.MatureElementUuid &&
		testCaseModelElement_2_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 7) Validate B10 (2.2.1)
	testCaseModelElementUuid_2_2_1 := testCaseModelElement_2_2.FirstChildElementUuid
	testCaseModelElement_2_2_1 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_2_1]

	correctElement = testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2.FirstChildElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.PreviousElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.NextElementUuid &&
		testCaseModelElement_2_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = verifyThatThereAreNoZombieElementsInTestCaseModel(&myTestCaseModel)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

}

// TCRuleSwap103
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B11f				n		TIC(B11f-X)						TIC(B11f-n-B12-X)		TCRuleSwap103
func TestTCRuleSwap103(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate a TestCaseModel
	myTestCaseModel := testCaseModel.TestCaseModelStruct{
		LastLoadedTestCaseModelGRPCMessage: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage{},
		FirstElementUuid:                   "",
		TestCaseModelMap:                   nil,
	}

	myTestCaseModel.TestCaseModelMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage)

	uuidToSwapOut := "4b694f8c-f194-45af-a75e-f2bb3fd350e6"

	tc_1_b1f := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "cff02b54-ee53-47d1-94d6-48dc470073a3",
		OriginalElementName:      "B1l_BOND",
		MatureElementUuid:        "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		PreviousElementUuid:      "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		NextElementUuid:          "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		FirstChildElementUuid:    "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		ParentElementUuid:        "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE,
	}

	tc_2_tic := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		MatureElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		PreviousElementUuid:      "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		NextElementUuid:          "79a6702d-8370-446c-b001-d60494eca6fa",
		FirstChildElementUuid:    uuidToSwapOut,
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_3_b1l := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c10bfd8f-2786-480c-9a71-bad9ec7bc582",
		OriginalElementName:      "B1l_BOND",
		MatureElementUuid:        "79a6702d-8370-446c-b001-d60494eca6fa",
		PreviousElementUuid:      "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		NextElementUuid:          "79a6702d-8370-446c-b001-d60494eca6fa",
		FirstChildElementUuid:    "79a6702d-8370-446c-b001-d60494eca6fa",
		ParentElementUuid:        "79a6702d-8370-446c-b001-d60494eca6fa",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE,
	}

	tc_2_1_b11f := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "f9df1310-0c47-4c0b-a5aa-5330c1589cac",
		OriginalElementName:      "B11l_BOND",
		MatureElementUuid:        uuidToSwapOut,
		PreviousElementUuid:      uuidToSwapOut,
		NextElementUuid:          "a39133d9-d8de-4c96-b8f6-7115aa88bfa6",
		FirstChildElementUuid:    uuidToSwapOut,
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
	}

	tc_2_2_tic := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		MatureElementUuid:        "a39133d9-d8de-4c96-b8f6-7115aa88bfa6",
		PreviousElementUuid:      uuidToSwapOut,
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_2_3_b11l := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "66284917-6b5c-42a0-98d6-106a7bd269fa",
		OriginalElementName:      "B11l_BOND",
		MatureElementUuid:        "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		PreviousElementUuid:      "a39133d9-d8de-4c96-b8f6-7115aa88bfa6",
		NextElementUuid:          "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		FirstChildElementUuid:    "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND,
	}

	tc_2_2_1_b10 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "4f348a88-bb9a-4852-80de-3141687db11d",
		OriginalElementName:      "B10_Bond",
		MatureElementUuid:        "edaf757d-1205-4f2d-91a4-f053982f5ded",
		PreviousElementUuid:      "edaf757d-1205-4f2d-91a4-f053982f5ded",
		NextElementUuid:          "edaf757d-1205-4f2d-91a4-f053982f5ded",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "a39133d9-d8de-4c96-b8f6-7115aa88bfa6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
	}

	// Set up TestCaseModel-map
	myTestCaseModel.TestCaseModelMap[tc_1_b1f.MatureElementUuid] = tc_1_b1f
	myTestCaseModel.TestCaseModelMap[tc_2_tic.MatureElementUuid] = tc_2_tic
	myTestCaseModel.TestCaseModelMap[tc_3_b1l.MatureElementUuid] = tc_3_b1l
	myTestCaseModel.TestCaseModelMap[tc_2_1_b11f.MatureElementUuid] = tc_2_1_b11f
	myTestCaseModel.TestCaseModelMap[tc_2_2_tic.MatureElementUuid] = tc_2_2_tic
	myTestCaseModel.TestCaseModelMap[tc_2_3_b11l.MatureElementUuid] = tc_2_3_b11l
	myTestCaseModel.TestCaseModelMap[tc_2_2_1_b10.MatureElementUuid] = tc_2_2_1_b10

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_1_b1f.MatureElementUuid

	// Create an Immature Element model for 'TIC(B10)'
	immatureElementModel := immatureElementStruct{
		firstElementUuid:   "",
		immatureElementMap: nil,
	}

	immatureElementModel.immatureElementMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage)

	// Create TIC
	tic := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		ImmatureElementUuid:      "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		PreviousElementUuid:      "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		NextElementUuid:          "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		FirstChildElementUuid:    "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		ParentElementUuid:        "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	// Create B10 in TIC(x)
	b10Bond := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage{
		OriginalElementUuid:      "7be82e83-6048-4c30-b4aa-b68c11037c1d",
		OriginalElementName:      "B10-Bond",
		ImmatureElementUuid:      "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		PreviousElementUuid:      "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		NextElementUuid:          "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		FirstChildElementUuid:    "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		ParentElementUuid:        "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
	}

	// Add the Elements to the Immature Elements Model map
	immatureElementModel.immatureElementMap["d444b8d8-b2fb-4505-ad8e-36bfe89988ab"] = tic
	immatureElementModel.immatureElementMap["ff224d27-5c8a-48b9-ace9-af43245cb35d"] = b10Bond

	// Add first Element ti Immature Element Model
	immatureElementModel.firstElementUuid = "d444b8d8-b2fb-4505-ad8e-36bfe89988ab"

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := commandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		testcaseModel:     &myTestCaseModel,
	}

	// Add needed data for availableBondsMap
	tempAvailableBondsMap := make(map[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage)

	// B12_BOND
	visibleBondAttributesMessage_AvaialbeBond_B12_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage_VisibleBondAttributesMessage{
		BondUuid: "b7879e7f-4dcc-4324-ab85-67ea37b56bc9",
		BondName: "B12_BOND",
	}

	basicBondInformationMessage_AvaialbeBond_B12_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage{
		VisibleBondAttributes: &visibleBondAttributesMessage_AvaialbeBond_B12_BOND}

	immatureBondsMessage_ImmatureBondMessage_B12_BOND := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage{
		BasicBondInformation: &basicBondInformationMessage_AvaialbeBond_B12_BOND}

	tempAvailableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND] = &immatureBondsMessage_ImmatureBondMessage_B12_BOND

	// Add bond-map to 'commandAndRuleEngine.availableBondsMap'
	commandAndRuleEngine.availableBondsMap = tempAvailableBondsMap

	// Execute Swap
	err := commandAndRuleEngine.executeTCRuleSwap103(uuidToSwapOut, &immatureElementModel)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate the result of the swap, 'B1f-TIC(B11f-TIC(B10)-B12-TIC(B10)-B11l)-B1l'
	// 1) Validate B1f (1)
	testCaseModelElementUuid_1 := myTestCaseModel.FirstElementUuid
	testCaseModelElement_1 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_1]

	correctElement := testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.ParentElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.PreviousElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.FirstChildElementUuid &&
		testCaseModelElement_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 2) Validate TIC (2)
	testCaseModelElementUuid_2 := testCaseModelElement_1.NextElementUuid
	testCaseModelElement_2 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2.ParentElementUuid &&
		testCaseModelElement_2.PreviousElementUuid == testCaseModelElement_1.MatureElementUuid &&
		testCaseModelElement_1.NextElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 3) Validate B1l (3)
	testCaseModelElementUuid_3 := testCaseModelElement_2.NextElementUuid
	testCaseModelElement_3 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_3]

	correctElement = testCaseModelElement_3.MatureElementUuid == testCaseModelElement_3.ParentElementUuid &&
		testCaseModelElement_3.NextElementUuid == testCaseModelElement_3.MatureElementUuid &&
		testCaseModelElement_3.FirstChildElementUuid == testCaseModelElement_3.MatureElementUuid &&
		testCaseModelElement_3.PreviousElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2.NextElementUuid == testCaseModelElement_3.MatureElementUuid &&
		testCaseModelElement_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 4) Validate B11f (2.1)
	testCaseModelElementUuid_2_1 := testCaseModelElement_2.FirstChildElementUuid
	testCaseModelElement_2_1 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_1]

	correctElement = testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2.FirstChildElementUuid &&
		testCaseModelElement_2_1.ParentElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 5) Validate TIC (2.2)
	testCaseModelElementUuid_2_2 := testCaseModelElement_2_1.NextElementUuid
	testCaseModelElement_2_2 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_2.ParentElementUuid &&
		testCaseModelElement_2_2.PreviousElementUuid == testCaseModelElement_2_1.MatureElementUuid &&
		testCaseModelElement_2_1.NextElementUuid == testCaseModelElement_2_2.MatureElementUuid &&
		testCaseModelElement_2_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 6) Validate B12 (2.3)
	testCaseModelElementUuid_2_3 := testCaseModelElement_2_2.NextElementUuid
	testCaseModelElement_2_3 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_3]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_3.ParentElementUuid &&
		testCaseModelElement_2_3.PreviousElementUuid == testCaseModelElement_2_2.MatureElementUuid &&
		testCaseModelElement_2_2.NextElementUuid == testCaseModelElement_2_3.MatureElementUuid &&
		testCaseModelElement_2_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 7) Validate TIC (2.4)
	testCaseModelElementUuid_2_4 := testCaseModelElement_2_3.NextElementUuid
	testCaseModelElement_2_4 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_4]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_4.ParentElementUuid &&
		testCaseModelElement_2_4.PreviousElementUuid == testCaseModelElement_2_3.MatureElementUuid &&
		testCaseModelElement_2_3.NextElementUuid == testCaseModelElement_2_4.MatureElementUuid &&
		testCaseModelElement_2_4.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 8) Validate B1ll (2.3)
	testCaseModelElementUuid_2_5 := testCaseModelElement_2_4.NextElementUuid
	testCaseModelElement_2_5 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_5]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_5.ParentElementUuid &&
		testCaseModelElement_2_5.FirstChildElementUuid == testCaseModelElement_2_5.MatureElementUuid &&
		testCaseModelElement_2_5.PreviousElementUuid == testCaseModelElement_2_4.MatureElementUuid &&
		testCaseModelElement_2_4.NextElementUuid == testCaseModelElement_2_5.MatureElementUuid &&
		testCaseModelElement_2_5.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 9) Validate B10 (2.2.1)
	testCaseModelElementUuid_2_2_1 := testCaseModelElement_2_2.FirstChildElementUuid
	testCaseModelElement_2_2_1 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_2_1]

	correctElement = testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2.FirstChildElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.PreviousElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.NextElementUuid &&
		testCaseModelElement_2_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 10) Validate B10 (2.4.1)
	testCaseModelElementUuid_2_4_1 := testCaseModelElement_2_4.FirstChildElementUuid
	testCaseModelElement_2_4_1 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_4_1]

	correctElement = testCaseModelElement_2_4_1.MatureElementUuid == testCaseModelElement_2_4.FirstChildElementUuid &&
		testCaseModelElement_2_4_1.MatureElementUuid == testCaseModelElement_2_4_1.PreviousElementUuid &&
		testCaseModelElement_2_4_1.MatureElementUuid == testCaseModelElement_2_4_1.FirstChildElementUuid &&
		testCaseModelElement_2_4_1.MatureElementUuid == testCaseModelElement_2_4_1.NextElementUuid &&
		testCaseModelElement_2_4_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = verifyThatThereAreNoZombieElementsInTestCaseModel(&myTestCaseModel)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

}

// TCRuleSwap104
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B11l				n		TIC(X-B11l)						TIC(X-B12-n-B11l)		TCRuleSwap104
func TestTCRuleSwap104(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate a TestCaseModel
	myTestCaseModel := testCaseModel.TestCaseModelStruct{
		LastLoadedTestCaseModelGRPCMessage: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage{},
		FirstElementUuid:                   "",
		TestCaseModelMap:                   nil,
	}

	myTestCaseModel.TestCaseModelMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage)

	uuidToSwapOut := "4b694f8c-f194-45af-a75e-f2bb3fd350e6"

	tc_1_b1f := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "cff02b54-ee53-47d1-94d6-48dc470073a3",
		OriginalElementName:      "B1l_BOND",
		MatureElementUuid:        "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		PreviousElementUuid:      "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		NextElementUuid:          "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		FirstChildElementUuid:    "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		ParentElementUuid:        "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE,
	}

	tc_2_tic := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		MatureElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		PreviousElementUuid:      "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		NextElementUuid:          "79a6702d-8370-446c-b001-d60494eca6fa",
		FirstChildElementUuid:    "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_3_b1l := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c10bfd8f-2786-480c-9a71-bad9ec7bc582",
		OriginalElementName:      "B1l_BOND",
		MatureElementUuid:        "79a6702d-8370-446c-b001-d60494eca6fa",
		PreviousElementUuid:      "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		NextElementUuid:          "79a6702d-8370-446c-b001-d60494eca6fa",
		FirstChildElementUuid:    "79a6702d-8370-446c-b001-d60494eca6fa",
		ParentElementUuid:        "79a6702d-8370-446c-b001-d60494eca6fa",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE,
	}

	tc_2_1_b11f := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "f9df1310-0c47-4c0b-a5aa-5330c1589cac",
		OriginalElementName:      "B11l_BOND",
		MatureElementUuid:        "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		PreviousElementUuid:      "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		NextElementUuid:          "a39133d9-d8de-4c96-b8f6-7115aa88bfa6",
		FirstChildElementUuid:    "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
	}

	tc_2_2_tic := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		MatureElementUuid:        "a39133d9-d8de-4c96-b8f6-7115aa88bfa6",
		PreviousElementUuid:      "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		NextElementUuid:          uuidToSwapOut,
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_2_3_b11l := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "66284917-6b5c-42a0-98d6-106a7bd269fa",
		OriginalElementName:      "B11l_BOND",
		MatureElementUuid:        uuidToSwapOut,
		PreviousElementUuid:      "a39133d9-d8de-4c96-b8f6-7115aa88bfa6",
		NextElementUuid:          uuidToSwapOut,
		FirstChildElementUuid:    uuidToSwapOut,
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND,
	}

	tc_2_2_1_b10 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "4f348a88-bb9a-4852-80de-3141687db11d",
		OriginalElementName:      "B10_Bond",
		MatureElementUuid:        "edaf757d-1205-4f2d-91a4-f053982f5ded",
		PreviousElementUuid:      "edaf757d-1205-4f2d-91a4-f053982f5ded",
		NextElementUuid:          "edaf757d-1205-4f2d-91a4-f053982f5ded",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "a39133d9-d8de-4c96-b8f6-7115aa88bfa6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
	}

	// Set up TestCaseModel-map
	myTestCaseModel.TestCaseModelMap[tc_1_b1f.MatureElementUuid] = tc_1_b1f
	myTestCaseModel.TestCaseModelMap[tc_2_tic.MatureElementUuid] = tc_2_tic
	myTestCaseModel.TestCaseModelMap[tc_3_b1l.MatureElementUuid] = tc_3_b1l
	myTestCaseModel.TestCaseModelMap[tc_2_1_b11f.MatureElementUuid] = tc_2_1_b11f
	myTestCaseModel.TestCaseModelMap[tc_2_2_tic.MatureElementUuid] = tc_2_2_tic
	myTestCaseModel.TestCaseModelMap[tc_2_3_b11l.MatureElementUuid] = tc_2_3_b11l
	myTestCaseModel.TestCaseModelMap[tc_2_2_1_b10.MatureElementUuid] = tc_2_2_1_b10

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_1_b1f.MatureElementUuid

	// Create an Immature Element model for 'TIC(B10)'
	immatureElementModel := immatureElementStruct{
		firstElementUuid:   "",
		immatureElementMap: nil,
	}

	immatureElementModel.immatureElementMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage)

	// Create TIC
	tic := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		ImmatureElementUuid:      "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		PreviousElementUuid:      "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		NextElementUuid:          "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		FirstChildElementUuid:    "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		ParentElementUuid:        "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	// Create B10 in TIC(x)
	b10Bond := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage{
		OriginalElementUuid:      "7be82e83-6048-4c30-b4aa-b68c11037c1d",
		OriginalElementName:      "B10-Bond",
		ImmatureElementUuid:      "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		PreviousElementUuid:      "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		NextElementUuid:          "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		FirstChildElementUuid:    "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		ParentElementUuid:        "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
	}

	// Add the Elements to the Immature Elements Model map
	immatureElementModel.immatureElementMap["d444b8d8-b2fb-4505-ad8e-36bfe89988ab"] = tic
	immatureElementModel.immatureElementMap["ff224d27-5c8a-48b9-ace9-af43245cb35d"] = b10Bond

	// Add first Element ti Immature Element Model
	immatureElementModel.firstElementUuid = "d444b8d8-b2fb-4505-ad8e-36bfe89988ab"

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := commandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		testcaseModel:     &myTestCaseModel,
	}

	// Add needed data for availableBondsMap
	tempAvailableBondsMap := make(map[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage)

	// B12_BOND
	visibleBondAttributesMessage_AvaialbeBond_B12_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage_VisibleBondAttributesMessage{
		BondUuid: "b7879e7f-4dcc-4324-ab85-67ea37b56bc9",
		BondName: "B12_BOND",
	}

	basicBondInformationMessage_AvaialbeBond_B12_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage{
		VisibleBondAttributes: &visibleBondAttributesMessage_AvaialbeBond_B12_BOND}

	immatureBondsMessage_ImmatureBondMessage_B12_BOND := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage{
		BasicBondInformation: &basicBondInformationMessage_AvaialbeBond_B12_BOND}

	tempAvailableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND] = &immatureBondsMessage_ImmatureBondMessage_B12_BOND

	// Add bond-map to 'commandAndRuleEngine.availableBondsMap'
	commandAndRuleEngine.availableBondsMap = tempAvailableBondsMap

	// Execute Swap
	err := commandAndRuleEngine.executeTCRuleSwap104(uuidToSwapOut, &immatureElementModel)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate the result of the swap, 'B1f-TIC(B11f-TIC(B10)-B12-TIC(B10)-B11l)-B1l'
	// 1) Validate B1f (1)
	testCaseModelElementUuid_1 := myTestCaseModel.FirstElementUuid
	testCaseModelElement_1 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_1]

	correctElement := testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.ParentElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.PreviousElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.FirstChildElementUuid &&
		testCaseModelElement_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 2) Validate TIC (2)
	testCaseModelElementUuid_2 := testCaseModelElement_1.NextElementUuid
	testCaseModelElement_2 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2.ParentElementUuid &&
		testCaseModelElement_2.PreviousElementUuid == testCaseModelElement_1.MatureElementUuid &&
		testCaseModelElement_1.NextElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 3) Validate B1l (3)
	testCaseModelElementUuid_3 := testCaseModelElement_2.NextElementUuid
	testCaseModelElement_3 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_3]

	correctElement = testCaseModelElement_3.MatureElementUuid == testCaseModelElement_3.ParentElementUuid &&
		testCaseModelElement_3.NextElementUuid == testCaseModelElement_3.MatureElementUuid &&
		testCaseModelElement_3.FirstChildElementUuid == testCaseModelElement_3.MatureElementUuid &&
		testCaseModelElement_3.PreviousElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2.NextElementUuid == testCaseModelElement_3.MatureElementUuid &&
		testCaseModelElement_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 4) Validate B11f (2.1)
	testCaseModelElementUuid_2_1 := testCaseModelElement_2.FirstChildElementUuid
	testCaseModelElement_2_1 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_1]

	correctElement = testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2.FirstChildElementUuid &&
		testCaseModelElement_2_1.ParentElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 5) Validate TIC (2.2)
	testCaseModelElementUuid_2_2 := testCaseModelElement_2_1.NextElementUuid
	testCaseModelElement_2_2 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_2.ParentElementUuid &&
		testCaseModelElement_2_2.PreviousElementUuid == testCaseModelElement_2_1.MatureElementUuid &&
		testCaseModelElement_2_1.NextElementUuid == testCaseModelElement_2_2.MatureElementUuid &&
		testCaseModelElement_2_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 6) Validate B12 (2.3)
	testCaseModelElementUuid_2_3 := testCaseModelElement_2_2.NextElementUuid
	testCaseModelElement_2_3 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_3]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_3.ParentElementUuid &&
		testCaseModelElement_2_3.PreviousElementUuid == testCaseModelElement_2_2.MatureElementUuid &&
		testCaseModelElement_2_2.NextElementUuid == testCaseModelElement_2_3.MatureElementUuid &&
		testCaseModelElement_2_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 7) Validate TIC (2.4)
	testCaseModelElementUuid_2_4 := testCaseModelElement_2_3.NextElementUuid
	testCaseModelElement_2_4 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_4]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_4.ParentElementUuid &&
		testCaseModelElement_2_4.PreviousElementUuid == testCaseModelElement_2_3.MatureElementUuid &&
		testCaseModelElement_2_3.NextElementUuid == testCaseModelElement_2_4.MatureElementUuid &&
		testCaseModelElement_2_4.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 8) Validate B1ll (2.3)
	testCaseModelElementUuid_2_5 := testCaseModelElement_2_4.NextElementUuid
	testCaseModelElement_2_5 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_5]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_5.ParentElementUuid &&
		testCaseModelElement_2_5.FirstChildElementUuid == testCaseModelElement_2_5.MatureElementUuid &&
		testCaseModelElement_2_5.PreviousElementUuid == testCaseModelElement_2_4.MatureElementUuid &&
		testCaseModelElement_2_4.NextElementUuid == testCaseModelElement_2_5.MatureElementUuid &&
		testCaseModelElement_2_5.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 9) Validate B10 (2.2.1)
	testCaseModelElementUuid_2_2_1 := testCaseModelElement_2_2.FirstChildElementUuid
	testCaseModelElement_2_2_1 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_2_1]

	correctElement = testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2.FirstChildElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.PreviousElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.NextElementUuid &&
		testCaseModelElement_2_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 10) Validate B10 (2.4.1)
	testCaseModelElementUuid_2_4_1 := testCaseModelElement_2_4.FirstChildElementUuid
	testCaseModelElement_2_4_1 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_4_1]

	correctElement = testCaseModelElement_2_4_1.MatureElementUuid == testCaseModelElement_2_4.FirstChildElementUuid &&
		testCaseModelElement_2_4_1.MatureElementUuid == testCaseModelElement_2_4_1.PreviousElementUuid &&
		testCaseModelElement_2_4_1.MatureElementUuid == testCaseModelElement_2_4_1.FirstChildElementUuid &&
		testCaseModelElement_2_4_1.MatureElementUuid == testCaseModelElement_2_4_1.NextElementUuid &&
		testCaseModelElement_2_4_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = verifyThatThereAreNoZombieElementsInTestCaseModel(&myTestCaseModel)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

}

// TCRuleSwap105
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B12					n		X-B12-X							X-B12-n-B12-X			TCRuleSwap105
func TestTCRuleSwap105(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate a TestCaseModel
	myTestCaseModel := testCaseModel.TestCaseModelStruct{
		LastLoadedTestCaseModelGRPCMessage: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage{},
		FirstElementUuid:                   "",
		TestCaseModelMap:                   nil,
	}

	myTestCaseModel.TestCaseModelMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage)

	uuidToSwapOut := "4b694f8c-f194-45af-a75e-f2bb3fd350e6"

	tc_1_b1f := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "cff02b54-ee53-47d1-94d6-48dc470073a3",
		OriginalElementName:      "B1l_BOND",
		MatureElementUuid:        "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		PreviousElementUuid:      "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		NextElementUuid:          "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		FirstChildElementUuid:    "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		ParentElementUuid:        "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE,
	}

	tc_2_tic := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		MatureElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		PreviousElementUuid:      "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		NextElementUuid:          "79a6702d-8370-446c-b001-d60494eca6fa",
		FirstChildElementUuid:    "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_3_b1l := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c10bfd8f-2786-480c-9a71-bad9ec7bc582",
		OriginalElementName:      "B1l_BOND",
		MatureElementUuid:        "79a6702d-8370-446c-b001-d60494eca6fa",
		PreviousElementUuid:      "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		NextElementUuid:          "79a6702d-8370-446c-b001-d60494eca6fa",
		FirstChildElementUuid:    "79a6702d-8370-446c-b001-d60494eca6fa",
		ParentElementUuid:        "79a6702d-8370-446c-b001-d60494eca6fa",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE,
	}

	tc_2_1_b11f := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "f9df1310-0c47-4c0b-a5aa-5330c1589cac",
		OriginalElementName:      "B11l_BOND",
		MatureElementUuid:        "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		PreviousElementUuid:      "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		NextElementUuid:          "a39133d9-d8de-4c96-b8f6-7115aa88bfa6",
		FirstChildElementUuid:    "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
	}

	tc_2_2_tic := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		MatureElementUuid:        "a39133d9-d8de-4c96-b8f6-7115aa88bfa6",
		PreviousElementUuid:      "8c5611ad-715f-449a-bf1c-33df40ae9a25",
		NextElementUuid:          uuidToSwapOut,
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_2_3_b12 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "b7879e7f-4dcc-4324-ab85-67ea37b56bc9",
		OriginalElementName:      "B12_BOND",
		MatureElementUuid:        uuidToSwapOut,
		PreviousElementUuid:      "a39133d9-d8de-4c96-b8f6-7115aa88bfa6",
		NextElementUuid:          "6b295ad2-0ff3-4c54-a7cf-ebb4d6016d26",
		FirstChildElementUuid:    uuidToSwapOut,
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND,
	}

	tc_2_4_ti := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "0f5132d4-8b36-4f1a-83ff-36402a210307",
		OriginalElementName:      "TI",
		MatureElementUuid:        "6b295ad2-0ff3-4c54-a7cf-ebb4d6016d26",
		PreviousElementUuid:      uuidToSwapOut,
		NextElementUuid:          "baebdb3d-c8b6-42f6-b93f-be16f5e344b0",
		FirstChildElementUuid:    "6b295ad2-0ff3-4c54-a7cf-ebb4d6016d26",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
	}

	tc_2_5_b11l := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "66284917-6b5c-42a0-98d6-106a7bd269fa",
		OriginalElementName:      "B11l_BOND",
		MatureElementUuid:        "baebdb3d-c8b6-42f6-b93f-be16f5e344b0",
		PreviousElementUuid:      "6b295ad2-0ff3-4c54-a7cf-ebb4d6016d26",
		NextElementUuid:          "baebdb3d-c8b6-42f6-b93f-be16f5e344b0",
		FirstChildElementUuid:    "baebdb3d-c8b6-42f6-b93f-be16f5e344b0",
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND,
	}

	tc_2_2_1_b10 := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "4f348a88-bb9a-4852-80de-3141687db11d",
		OriginalElementName:      "B10_Bond",
		MatureElementUuid:        "edaf757d-1205-4f2d-91a4-f053982f5ded",
		PreviousElementUuid:      "edaf757d-1205-4f2d-91a4-f053982f5ded",
		NextElementUuid:          "edaf757d-1205-4f2d-91a4-f053982f5ded",
		FirstChildElementUuid:    "edaf757d-1205-4f2d-91a4-f053982f5ded",
		ParentElementUuid:        "a39133d9-d8de-4c96-b8f6-7115aa88bfa6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
	}

	// Set up TestCaseModel-map
	myTestCaseModel.TestCaseModelMap[tc_1_b1f.MatureElementUuid] = tc_1_b1f
	myTestCaseModel.TestCaseModelMap[tc_2_tic.MatureElementUuid] = tc_2_tic
	myTestCaseModel.TestCaseModelMap[tc_3_b1l.MatureElementUuid] = tc_3_b1l
	myTestCaseModel.TestCaseModelMap[tc_2_1_b11f.MatureElementUuid] = tc_2_1_b11f
	myTestCaseModel.TestCaseModelMap[tc_2_2_tic.MatureElementUuid] = tc_2_2_tic
	myTestCaseModel.TestCaseModelMap[tc_2_3_b12.MatureElementUuid] = tc_2_3_b12
	myTestCaseModel.TestCaseModelMap[tc_2_4_ti.MatureElementUuid] = tc_2_4_ti
	myTestCaseModel.TestCaseModelMap[tc_2_5_b11l.MatureElementUuid] = tc_2_5_b11l
	myTestCaseModel.TestCaseModelMap[tc_2_2_1_b10.MatureElementUuid] = tc_2_2_1_b10

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_1_b1f.MatureElementUuid

	// Create an Immature Element model for 'TI'
	immatureElementModel := immatureElementStruct{
		firstElementUuid:   "",
		immatureElementMap: nil,
	}

	immatureElementModel.immatureElementMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage)

	// Create TI
	ti := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage{
		OriginalElementUuid:      "0f5132d4-8b36-4f1a-83ff-36402a210307",
		OriginalElementName:      "TI",
		ImmatureElementUuid:      "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		PreviousElementUuid:      "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		NextElementUuid:          "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		FirstChildElementUuid:    "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		ParentElementUuid:        "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
	}

	// Add the Elements to the Immature Elements Model map
	immatureElementModel.immatureElementMap["d444b8d8-b2fb-4505-ad8e-36bfe89988ab"] = ti

	// Add first Element to Immature Element Model
	immatureElementModel.firstElementUuid = "d444b8d8-b2fb-4505-ad8e-36bfe89988ab"

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := commandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		testcaseModel:     &myTestCaseModel,
	}

	// Add needed data for availableBondsMap
	tempAvailableBondsMap := make(map[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage)

	// B12_BOND
	visibleBondAttributesMessage_AvaialbeBond_B12_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage_VisibleBondAttributesMessage{
		BondUuid: "b7879e7f-4dcc-4324-ab85-67ea37b56bc9",
		BondName: "B12_BOND",
	}

	basicBondInformationMessage_AvaialbeBond_B12_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage{
		VisibleBondAttributes: &visibleBondAttributesMessage_AvaialbeBond_B12_BOND}

	immatureBondsMessage_ImmatureBondMessage_B12_BOND := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage{
		BasicBondInformation: &basicBondInformationMessage_AvaialbeBond_B12_BOND}

	tempAvailableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND] = &immatureBondsMessage_ImmatureBondMessage_B12_BOND

	// Add bond-map to 'commandAndRuleEngine.availableBondsMap'
	commandAndRuleEngine.availableBondsMap = tempAvailableBondsMap

	// Execute Swap
	err := commandAndRuleEngine.executeTCRuleSwap105(uuidToSwapOut, &immatureElementModel)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate the result of the swap, 'B1f-TIC(B11f-TIC(B10)-B12-TI-B12-TI-B11l)-B1l'
	// 1) Validate B1f (1)
	testCaseModelElementUuid_1 := myTestCaseModel.FirstElementUuid
	testCaseModelElement_1 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_1]

	correctElement := testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.ParentElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.PreviousElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.FirstChildElementUuid &&
		testCaseModelElement_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 2) Validate TIC (2)
	testCaseModelElementUuid_2 := testCaseModelElement_1.NextElementUuid
	testCaseModelElement_2 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2.ParentElementUuid &&
		testCaseModelElement_2.PreviousElementUuid == testCaseModelElement_1.MatureElementUuid &&
		testCaseModelElement_1.NextElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 3) Validate B1l (3)
	testCaseModelElementUuid_3 := testCaseModelElement_2.NextElementUuid
	testCaseModelElement_3 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_3]

	correctElement = testCaseModelElement_3.MatureElementUuid == testCaseModelElement_3.ParentElementUuid &&
		testCaseModelElement_3.NextElementUuid == testCaseModelElement_3.MatureElementUuid &&
		testCaseModelElement_3.FirstChildElementUuid == testCaseModelElement_3.MatureElementUuid &&
		testCaseModelElement_3.PreviousElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2.NextElementUuid == testCaseModelElement_3.MatureElementUuid &&
		testCaseModelElement_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 4) Validate B11f (2.1)
	testCaseModelElementUuid_2_1 := testCaseModelElement_2.FirstChildElementUuid
	testCaseModelElement_2_1 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_1]

	correctElement = testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2.FirstChildElementUuid &&
		testCaseModelElement_2_1.ParentElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 5) Validate TIC (2.2)
	testCaseModelElementUuid_2_2 := testCaseModelElement_2_1.NextElementUuid
	testCaseModelElement_2_2 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_2.ParentElementUuid &&
		testCaseModelElement_2_2.PreviousElementUuid == testCaseModelElement_2_1.MatureElementUuid &&
		testCaseModelElement_2_1.NextElementUuid == testCaseModelElement_2_2.MatureElementUuid &&
		testCaseModelElement_2_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 6) Validate B12 (2.3)
	testCaseModelElementUuid_2_3 := testCaseModelElement_2_2.NextElementUuid
	testCaseModelElement_2_3 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_3]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_3.ParentElementUuid &&
		testCaseModelElement_2_3.PreviousElementUuid == testCaseModelElement_2_2.MatureElementUuid &&
		testCaseModelElement_2_2.NextElementUuid == testCaseModelElement_2_3.MatureElementUuid &&
		testCaseModelElement_2_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 7) Validate TI (2.4)
	testCaseModelElementUuid_2_4 := testCaseModelElement_2_3.NextElementUuid
	testCaseModelElement_2_4 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_4]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_4.ParentElementUuid &&
		testCaseModelElement_2_4.PreviousElementUuid == testCaseModelElement_2_3.MatureElementUuid &&
		testCaseModelElement_2_3.NextElementUuid == testCaseModelElement_2_4.MatureElementUuid &&
		testCaseModelElement_2_4.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 8) Validate B12 (2.5)
	testCaseModelElementUuid_2_5 := testCaseModelElement_2_4.NextElementUuid
	testCaseModelElement_2_5 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_5]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_5.ParentElementUuid &&
		testCaseModelElement_2_5.PreviousElementUuid == testCaseModelElement_2_4.MatureElementUuid &&
		testCaseModelElement_2_4.NextElementUuid == testCaseModelElement_2_5.MatureElementUuid &&
		testCaseModelElement_2_5.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 9) Validate TI (2.6)
	testCaseModelElementUuid_2_6 := testCaseModelElement_2_5.NextElementUuid
	testCaseModelElement_2_6 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_6]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_6.ParentElementUuid &&
		testCaseModelElement_2_6.PreviousElementUuid == testCaseModelElement_2_5.MatureElementUuid &&
		testCaseModelElement_2_5.NextElementUuid == testCaseModelElement_2_6.MatureElementUuid &&
		testCaseModelElement_2_6.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 10) Validate B1ll (2.7)
	testCaseModelElementUuid_2_7 := testCaseModelElement_2_6.NextElementUuid
	testCaseModelElement_2_7 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_7]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_7.ParentElementUuid &&
		testCaseModelElement_2_7.FirstChildElementUuid == testCaseModelElement_2_7.MatureElementUuid &&
		testCaseModelElement_2_7.PreviousElementUuid == testCaseModelElement_2_6.MatureElementUuid &&
		testCaseModelElement_2_6.NextElementUuid == testCaseModelElement_2_7.MatureElementUuid &&
		testCaseModelElement_2_7.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 11) Validate B10 (2.2.1)
	testCaseModelElementUuid_2_2_1 := testCaseModelElement_2_2.FirstChildElementUuid
	testCaseModelElement_2_2_1 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_2_1]

	correctElement = testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2.FirstChildElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.PreviousElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.NextElementUuid &&
		testCaseModelElement_2_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = verifyThatThereAreNoZombieElementsInTestCaseModel(&myTestCaseModel)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

}

// TCRuleSwap106
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B10x*				n		TIC(B10*x*)						TIC(B11x-n-B11x)		TCRuleSwap106
func TestTCRuleSwap106(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate a TestCaseModel
	myTestCaseModel := testCaseModel.TestCaseModelStruct{
		LastLoadedTestCaseModelGRPCMessage: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage{},
		FirstElementUuid:                   "",
		TestCaseModelMap:                   nil,
	}

	myTestCaseModel.TestCaseModelMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage)

	uuidToSwapOut := "4b694f8c-f194-45af-a75e-f2bb3fd350e6"

	tc_1_b1f := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "cff02b54-ee53-47d1-94d6-48dc470073a3",
		OriginalElementName:      "B1l_BOND",
		MatureElementUuid:        "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		PreviousElementUuid:      "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		NextElementUuid:          "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		FirstChildElementUuid:    "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		ParentElementUuid:        "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE,
	}

	tc_2_tic := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		MatureElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		PreviousElementUuid:      "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		NextElementUuid:          "79a6702d-8370-446c-b001-d60494eca6fa",
		FirstChildElementUuid:    uuidToSwapOut,
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_3_b1l := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c10bfd8f-2786-480c-9a71-bad9ec7bc582",
		OriginalElementName:      "B1l_BOND",
		MatureElementUuid:        "79a6702d-8370-446c-b001-d60494eca6fa",
		PreviousElementUuid:      "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		NextElementUuid:          "79a6702d-8370-446c-b001-d60494eca6fa",
		FirstChildElementUuid:    "79a6702d-8370-446c-b001-d60494eca6fa",
		ParentElementUuid:        "79a6702d-8370-446c-b001-d60494eca6fa",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE,
	}

	tc_2_1_b10oxo := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "4f348a88-bb9a-4852-80de-3141687db11d",
		OriginalElementName:      "B10_Bond",
		MatureElementUuid:        uuidToSwapOut,
		PreviousElementUuid:      uuidToSwapOut,
		NextElementUuid:          uuidToSwapOut,
		FirstChildElementUuid:    uuidToSwapOut,
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10oxo_BOND,
	}

	// Set up TestCaseModel-map
	myTestCaseModel.TestCaseModelMap[tc_1_b1f.MatureElementUuid] = tc_1_b1f
	myTestCaseModel.TestCaseModelMap[tc_2_tic.MatureElementUuid] = tc_2_tic
	myTestCaseModel.TestCaseModelMap[tc_3_b1l.MatureElementUuid] = tc_3_b1l
	myTestCaseModel.TestCaseModelMap[tc_2_1_b10oxo.MatureElementUuid] = tc_2_1_b10oxo

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_1_b1f.MatureElementUuid

	// Create an Immature Element model for 'TIC(B10)'
	immatureElementModel := immatureElementStruct{
		firstElementUuid:   "",
		immatureElementMap: nil,
	}

	immatureElementModel.immatureElementMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage)

	// Create TIC
	tic := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		ImmatureElementUuid:      "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		PreviousElementUuid:      "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		NextElementUuid:          "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		FirstChildElementUuid:    "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		ParentElementUuid:        "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	// Create B10 in TIC(x)
	b10Bond := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage{
		OriginalElementUuid:      "7be82e83-6048-4c30-b4aa-b68c11037c1d",
		OriginalElementName:      "B10-Bond",
		ImmatureElementUuid:      "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		PreviousElementUuid:      "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		NextElementUuid:          "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		FirstChildElementUuid:    "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		ParentElementUuid:        "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
	}

	// Add the Elements to the Immature Elements Model map
	immatureElementModel.immatureElementMap["d444b8d8-b2fb-4505-ad8e-36bfe89988ab"] = tic
	immatureElementModel.immatureElementMap["ff224d27-5c8a-48b9-ace9-af43245cb35d"] = b10Bond

	// Add first Element ti Immature Element Model
	immatureElementModel.firstElementUuid = "d444b8d8-b2fb-4505-ad8e-36bfe89988ab"

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := commandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		testcaseModel:     &myTestCaseModel,
	}

	// Add needed data for availableBondsMap
	tempAvailableBondsMap := make(map[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage)

	// B11f_BOND
	visibleBondAttributesMessage_AvaialbeBond_B11fx_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage_VisibleBondAttributesMessage{
		BondUuid: "0d77690e-f8e2-4942-b532-6b3e26d0b81a",
		BondName: "B11fx_BOND",
	}

	basicBondInformationMessage_AvaialbeBond_B11fx_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage{
		VisibleBondAttributes: &visibleBondAttributesMessage_AvaialbeBond_B11fx_BOND}

	immatureBondsMessage_ImmatureBondMessage_B11fx_BOND := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage{
		BasicBondInformation: &basicBondInformationMessage_AvaialbeBond_B11fx_BOND}

	tempAvailableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE] = &immatureBondsMessage_ImmatureBondMessage_B11fx_BOND

	// B11l_BOND
	visibleBondAttributesMessage_AvaialbeBond_B11lx_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage_VisibleBondAttributesMessage{
		BondUuid: "61f98266-b9b1-4958-9f90-1d0d7f17aafc",
		BondName: "B11lx_BOND",
	}

	basicBondInformationMessage_AvaialbeBond_B11lx_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage{
		VisibleBondAttributes: &visibleBondAttributesMessage_AvaialbeBond_B11lx_BOND}

	immatureBondsMessage_ImmatureBondMessage_B11lx_BOND := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage{
		BasicBondInformation: &basicBondInformationMessage_AvaialbeBond_B11lx_BOND}

	tempAvailableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE] = &immatureBondsMessage_ImmatureBondMessage_B11lx_BOND

	// Add bond-map to 'commandAndRuleEngine.availableBondsMap'
	commandAndRuleEngine.availableBondsMap = tempAvailableBondsMap

	// Execute Swap
	err := commandAndRuleEngine.executeTCRuleSwap106(uuidToSwapOut, &immatureElementModel)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate the result of the swap, 'B1f-TIC(B11fx-TIC(B10)-B11lx)-B1l'
	// 1) Validate B1f (1)
	testCaseModelElementUuid_1 := myTestCaseModel.FirstElementUuid
	testCaseModelElement_1 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_1]

	correctElement := testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.ParentElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.PreviousElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.FirstChildElementUuid &&
		testCaseModelElement_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 2) Validate TIC (2)
	testCaseModelElementUuid_2 := testCaseModelElement_1.NextElementUuid
	testCaseModelElement_2 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2.ParentElementUuid &&
		testCaseModelElement_2.PreviousElementUuid == testCaseModelElement_1.MatureElementUuid &&
		testCaseModelElement_1.NextElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 3) Validate B1l (3)
	testCaseModelElementUuid_3 := testCaseModelElement_2.NextElementUuid
	testCaseModelElement_3 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_3]

	correctElement = testCaseModelElement_3.MatureElementUuid == testCaseModelElement_3.ParentElementUuid &&
		testCaseModelElement_3.NextElementUuid == testCaseModelElement_3.MatureElementUuid &&
		testCaseModelElement_3.FirstChildElementUuid == testCaseModelElement_3.MatureElementUuid &&
		testCaseModelElement_3.PreviousElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2.NextElementUuid == testCaseModelElement_3.MatureElementUuid &&
		testCaseModelElement_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 4) Validate B11fx (2.1)
	testCaseModelElementUuid_2_1 := testCaseModelElement_2.FirstChildElementUuid
	testCaseModelElement_2_1 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_1]

	correctElement = testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2.FirstChildElementUuid &&
		testCaseModelElement_2_1.ParentElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 5) Validate TIC (2.2)
	testCaseModelElementUuid_2_2 := testCaseModelElement_2_1.NextElementUuid
	testCaseModelElement_2_2 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_2.ParentElementUuid &&
		testCaseModelElement_2_2.PreviousElementUuid == testCaseModelElement_2_1.MatureElementUuid &&
		testCaseModelElement_2_1.NextElementUuid == testCaseModelElement_2_2.MatureElementUuid &&
		testCaseModelElement_2_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 6) Validate B11lx (2.3)
	testCaseModelElementUuid_2_3 := testCaseModelElement_2_2.NextElementUuid
	testCaseModelElement_2_3 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_3]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_3.ParentElementUuid &&
		testCaseModelElement_2_3.FirstChildElementUuid == testCaseModelElement_2_3.MatureElementUuid &&
		testCaseModelElement_2_3.PreviousElementUuid == testCaseModelElement_2_2.MatureElementUuid &&
		testCaseModelElement_2_2.NextElementUuid == testCaseModelElement_2_3.MatureElementUuid &&
		testCaseModelElement_2_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 7) Validate B10 (2.2.1)
	testCaseModelElementUuid_2_2_1 := testCaseModelElement_2_2.FirstChildElementUuid
	testCaseModelElement_2_2_1 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_2_1]

	correctElement = testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2.FirstChildElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.PreviousElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.NextElementUuid &&
		testCaseModelElement_2_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = verifyThatThereAreNoZombieElementsInTestCaseModel(&myTestCaseModel)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

}

// TCRuleSwap107
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B10*x				n		TIC(B10*x)						TIC(B11x-n-B11)			TCRuleSwap107
func TestTCRuleSwap107(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate a TestCaseModel
	myTestCaseModel := testCaseModel.TestCaseModelStruct{
		LastLoadedTestCaseModelGRPCMessage: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage{},
		FirstElementUuid:                   "",
		TestCaseModelMap:                   nil,
	}

	myTestCaseModel.TestCaseModelMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage)

	uuidToSwapOut := "4b694f8c-f194-45af-a75e-f2bb3fd350e6"

	tc_1_b1f := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "cff02b54-ee53-47d1-94d6-48dc470073a3",
		OriginalElementName:      "B1l_BOND",
		MatureElementUuid:        "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		PreviousElementUuid:      "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		NextElementUuid:          "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		FirstChildElementUuid:    "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		ParentElementUuid:        "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE,
	}

	tc_2_tic := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		MatureElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		PreviousElementUuid:      "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		NextElementUuid:          "79a6702d-8370-446c-b001-d60494eca6fa",
		FirstChildElementUuid:    uuidToSwapOut,
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_3_b1l := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c10bfd8f-2786-480c-9a71-bad9ec7bc582",
		OriginalElementName:      "B1l_BOND",
		MatureElementUuid:        "79a6702d-8370-446c-b001-d60494eca6fa",
		PreviousElementUuid:      "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		NextElementUuid:          "79a6702d-8370-446c-b001-d60494eca6fa",
		FirstChildElementUuid:    "79a6702d-8370-446c-b001-d60494eca6fa",
		ParentElementUuid:        "79a6702d-8370-446c-b001-d60494eca6fa",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE,
	}

	tc_2_1_b10ox := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "4f348a88-bb9a-4852-80de-3141687db11d",
		OriginalElementName:      "B10_Bond",
		MatureElementUuid:        uuidToSwapOut,
		PreviousElementUuid:      uuidToSwapOut,
		NextElementUuid:          uuidToSwapOut,
		FirstChildElementUuid:    uuidToSwapOut,
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10ox_BOND,
	}

	// Set up TestCaseModel-map
	myTestCaseModel.TestCaseModelMap[tc_1_b1f.MatureElementUuid] = tc_1_b1f
	myTestCaseModel.TestCaseModelMap[tc_2_tic.MatureElementUuid] = tc_2_tic
	myTestCaseModel.TestCaseModelMap[tc_3_b1l.MatureElementUuid] = tc_3_b1l
	myTestCaseModel.TestCaseModelMap[tc_2_1_b10ox.MatureElementUuid] = tc_2_1_b10ox

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_1_b1f.MatureElementUuid

	// Create an Immature Element model for 'TIC(B10)'
	immatureElementModel := immatureElementStruct{
		firstElementUuid:   "",
		immatureElementMap: nil,
	}

	immatureElementModel.immatureElementMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage)

	// Create TIC
	tic := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		ImmatureElementUuid:      "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		PreviousElementUuid:      "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		NextElementUuid:          "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		FirstChildElementUuid:    "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		ParentElementUuid:        "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	// Create B10 in TIC(x)
	b10Bond := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage{
		OriginalElementUuid:      "7be82e83-6048-4c30-b4aa-b68c11037c1d",
		OriginalElementName:      "B10-Bond",
		ImmatureElementUuid:      "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		PreviousElementUuid:      "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		NextElementUuid:          "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		FirstChildElementUuid:    "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		ParentElementUuid:        "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
	}

	// Add the Elements to the Immature Elements Model map
	immatureElementModel.immatureElementMap["d444b8d8-b2fb-4505-ad8e-36bfe89988ab"] = tic
	immatureElementModel.immatureElementMap["ff224d27-5c8a-48b9-ace9-af43245cb35d"] = b10Bond

	// Add first Element ti Immature Element Model
	immatureElementModel.firstElementUuid = "d444b8d8-b2fb-4505-ad8e-36bfe89988ab"

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := commandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		testcaseModel:     &myTestCaseModel,
	}

	// Add needed data for availableBondsMap
	tempAvailableBondsMap := make(map[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage)

	// B11f_BOND
	visibleBondAttributesMessage_AvaialbeBond_B11fx_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage_VisibleBondAttributesMessage{
		BondUuid: "0d77690e-f8e2-4942-b532-6b3e26d0b81a",
		BondName: "B11fx_BOND",
	}

	basicBondInformationMessage_AvaialbeBond_B11fx_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage{
		VisibleBondAttributes: &visibleBondAttributesMessage_AvaialbeBond_B11fx_BOND}

	immatureBondsMessage_ImmatureBondMessage_B11fx_BOND := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage{
		BasicBondInformation: &basicBondInformationMessage_AvaialbeBond_B11fx_BOND}

	tempAvailableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE] = &immatureBondsMessage_ImmatureBondMessage_B11fx_BOND

	// B11l_BOND
	visibleBondAttributesMessage_AvaialbeBond_B11l_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage_VisibleBondAttributesMessage{
		BondUuid: "61f98266-b9b1-4958-9f90-1d0d7f17aafc",
		BondName: "B11l_BOND",
	}

	basicBondInformationMessage_AvaialbeBond_B11l_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage{
		VisibleBondAttributes: &visibleBondAttributesMessage_AvaialbeBond_B11l_BOND}

	immatureBondsMessage_ImmatureBondMessage_B11l_BOND := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage{
		BasicBondInformation: &basicBondInformationMessage_AvaialbeBond_B11l_BOND}

	tempAvailableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND] = &immatureBondsMessage_ImmatureBondMessage_B11l_BOND

	// Add bond-map to 'commandAndRuleEngine.availableBondsMap'
	commandAndRuleEngine.availableBondsMap = tempAvailableBondsMap

	// Execute Swap
	err := commandAndRuleEngine.executeTCRuleSwap107(uuidToSwapOut, &immatureElementModel)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate the result of the swap, 'B1f-TIC(B11fx-TIC(B10)-B11l)-B1l'
	// 1) Validate B1f (1)
	testCaseModelElementUuid_1 := myTestCaseModel.FirstElementUuid
	testCaseModelElement_1 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_1]

	correctElement := testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.ParentElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.PreviousElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.FirstChildElementUuid &&
		testCaseModelElement_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 2) Validate TIC (2)
	testCaseModelElementUuid_2 := testCaseModelElement_1.NextElementUuid
	testCaseModelElement_2 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2.ParentElementUuid &&
		testCaseModelElement_2.PreviousElementUuid == testCaseModelElement_1.MatureElementUuid &&
		testCaseModelElement_1.NextElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 3) Validate B1l (3)
	testCaseModelElementUuid_3 := testCaseModelElement_2.NextElementUuid
	testCaseModelElement_3 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_3]

	correctElement = testCaseModelElement_3.MatureElementUuid == testCaseModelElement_3.ParentElementUuid &&
		testCaseModelElement_3.NextElementUuid == testCaseModelElement_3.MatureElementUuid &&
		testCaseModelElement_3.FirstChildElementUuid == testCaseModelElement_3.MatureElementUuid &&
		testCaseModelElement_3.PreviousElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2.NextElementUuid == testCaseModelElement_3.MatureElementUuid &&
		testCaseModelElement_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 4) Validate B11fx (2.1)
	testCaseModelElementUuid_2_1 := testCaseModelElement_2.FirstChildElementUuid
	testCaseModelElement_2_1 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_1]

	correctElement = testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2.FirstChildElementUuid &&
		testCaseModelElement_2_1.ParentElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 5) Validate TIC (2.2)
	testCaseModelElementUuid_2_2 := testCaseModelElement_2_1.NextElementUuid
	testCaseModelElement_2_2 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_2.ParentElementUuid &&
		testCaseModelElement_2_2.PreviousElementUuid == testCaseModelElement_2_1.MatureElementUuid &&
		testCaseModelElement_2_1.NextElementUuid == testCaseModelElement_2_2.MatureElementUuid &&
		testCaseModelElement_2_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 6) Validate B11lx (2.3)
	testCaseModelElementUuid_2_3 := testCaseModelElement_2_2.NextElementUuid
	testCaseModelElement_2_3 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_3]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_3.ParentElementUuid &&
		testCaseModelElement_2_3.FirstChildElementUuid == testCaseModelElement_2_3.MatureElementUuid &&
		testCaseModelElement_2_3.PreviousElementUuid == testCaseModelElement_2_2.MatureElementUuid &&
		testCaseModelElement_2_2.NextElementUuid == testCaseModelElement_2_3.MatureElementUuid &&
		testCaseModelElement_2_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 7) Validate B10 (2.2.1)
	testCaseModelElementUuid_2_2_1 := testCaseModelElement_2_2.FirstChildElementUuid
	testCaseModelElement_2_2_1 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_2_1]

	correctElement = testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2.FirstChildElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.PreviousElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.NextElementUuid &&
		testCaseModelElement_2_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = verifyThatThereAreNoZombieElementsInTestCaseModel(&myTestCaseModel)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

}

// TCRuleSwap108
//	What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
//	n=TIC or TIC(X)		B10x*				n		TIC(B10x*)						TIC(B11-n-B11x)			TCRuleSwap108
func TestTCRuleSwap108(t *testing.T) {

	// Initiate logger used when testing
	myLogger := UnitTestTestData.InitLoggerForTest("")

	// Initiate a TestCaseModel
	myTestCaseModel := testCaseModel.TestCaseModelStruct{
		LastLoadedTestCaseModelGRPCMessage: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage{},
		FirstElementUuid:                   "",
		TestCaseModelMap:                   nil,
	}

	myTestCaseModel.TestCaseModelMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage)

	uuidToSwapOut := "4b694f8c-f194-45af-a75e-f2bb3fd350e6"

	tc_1_b1f := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "cff02b54-ee53-47d1-94d6-48dc470073a3",
		OriginalElementName:      "B1l_BOND",
		MatureElementUuid:        "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		PreviousElementUuid:      "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		NextElementUuid:          "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		FirstChildElementUuid:    "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		ParentElementUuid:        "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE,
	}

	tc_2_tic := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		MatureElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		PreviousElementUuid:      "bfe9c2ba-05db-4a75-bc07-db110a0a73ef",
		NextElementUuid:          "79a6702d-8370-446c-b001-d60494eca6fa",
		FirstChildElementUuid:    uuidToSwapOut,
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	tc_3_b1l := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c10bfd8f-2786-480c-9a71-bad9ec7bc582",
		OriginalElementName:      "B1l_BOND",
		MatureElementUuid:        "79a6702d-8370-446c-b001-d60494eca6fa",
		PreviousElementUuid:      "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		NextElementUuid:          "79a6702d-8370-446c-b001-d60494eca6fa",
		FirstChildElementUuid:    "79a6702d-8370-446c-b001-d60494eca6fa",
		ParentElementUuid:        "79a6702d-8370-446c-b001-d60494eca6fa",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE,
	}

	tc_2_1_b10ox := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      "4f348a88-bb9a-4852-80de-3141687db11d",
		OriginalElementName:      "B10_Bond",
		MatureElementUuid:        uuidToSwapOut,
		PreviousElementUuid:      uuidToSwapOut,
		NextElementUuid:          uuidToSwapOut,
		FirstChildElementUuid:    uuidToSwapOut,
		ParentElementUuid:        "f00393b6-fb69-4c77-93c5-94674dd8f2b6",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10xo_BOND,
	}

	// Set up TestCaseModel-map
	myTestCaseModel.TestCaseModelMap[tc_1_b1f.MatureElementUuid] = tc_1_b1f
	myTestCaseModel.TestCaseModelMap[tc_2_tic.MatureElementUuid] = tc_2_tic
	myTestCaseModel.TestCaseModelMap[tc_3_b1l.MatureElementUuid] = tc_3_b1l
	myTestCaseModel.TestCaseModelMap[tc_2_1_b10ox.MatureElementUuid] = tc_2_1_b10ox

	// Set the B1f-bond as first element in TestCaseModel
	myTestCaseModel.FirstElementUuid = tc_1_b1f.MatureElementUuid

	// Create an Immature Element model for 'TIC(B10)'
	immatureElementModel := immatureElementStruct{
		firstElementUuid:   "",
		immatureElementMap: nil,
	}

	immatureElementModel.immatureElementMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage)

	// Create TIC
	tic := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage{
		OriginalElementUuid:      "c072d6bf-e349-4730-9b04-4949368f50ea",
		OriginalElementName:      "TIC",
		ImmatureElementUuid:      "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		PreviousElementUuid:      "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		NextElementUuid:          "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		FirstChildElementUuid:    "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		ParentElementUuid:        "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
	}

	// Create B10 in TIC(x)
	b10Bond := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage{
		OriginalElementUuid:      "7be82e83-6048-4c30-b4aa-b68c11037c1d",
		OriginalElementName:      "B10-Bond",
		ImmatureElementUuid:      "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		PreviousElementUuid:      "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		NextElementUuid:          "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		FirstChildElementUuid:    "ff224d27-5c8a-48b9-ace9-af43245cb35d",
		ParentElementUuid:        "d444b8d8-b2fb-4505-ad8e-36bfe89988ab",
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
	}

	// Add the Elements to the Immature Elements Model map
	immatureElementModel.immatureElementMap["d444b8d8-b2fb-4505-ad8e-36bfe89988ab"] = tic
	immatureElementModel.immatureElementMap["ff224d27-5c8a-48b9-ace9-af43245cb35d"] = b10Bond

	// Add first Element ti Immature Element Model
	immatureElementModel.firstElementUuid = "d444b8d8-b2fb-4505-ad8e-36bfe89988ab"

	// Initiate CommandAndRule-engine
	commandAndRuleEngine := commandAndRuleEngineObjectStruct{
		logger:            myLogger,
		availableBondsMap: nil,
		testcaseModel:     &myTestCaseModel,
	}

	// Add needed data for availableBondsMap
	tempAvailableBondsMap := make(map[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage)

	// B11f_BOND
	visibleBondAttributesMessage_AvaialbeBond_B11f_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage_VisibleBondAttributesMessage{
		BondUuid: "0d77690e-f8e2-4942-b532-6b3e26d0b81a",
		BondName: "B11fx_BOND",
	}

	basicBondInformationMessage_AvaialbeBond_B11f_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage{
		VisibleBondAttributes: &visibleBondAttributesMessage_AvaialbeBond_B11f_BOND}

	immatureBondsMessage_ImmatureBondMessage_B11f_BOND := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage{
		BasicBondInformation: &basicBondInformationMessage_AvaialbeBond_B11f_BOND}

	tempAvailableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND] = &immatureBondsMessage_ImmatureBondMessage_B11f_BOND

	// B11lx_BOND
	visibleBondAttributesMessage_AvaialbeBond_B11lx_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage_VisibleBondAttributesMessage{
		BondUuid: "61f98266-b9b1-4958-9f90-1d0d7f17aafc",
		BondName: "B11lx_BOND",
	}

	basicBondInformationMessage_AvaialbeBond_B11lx_BOND := fenixGuiTestCaseBuilderServerGrpcApi.BasicBondInformationMessage{
		VisibleBondAttributes: &visibleBondAttributesMessage_AvaialbeBond_B11lx_BOND}

	immatureBondsMessage_ImmatureBondMessage_B11lx_BOND := fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage{
		BasicBondInformation: &basicBondInformationMessage_AvaialbeBond_B11lx_BOND}

	tempAvailableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE] = &immatureBondsMessage_ImmatureBondMessage_B11lx_BOND

	// Add bond-map to 'commandAndRuleEngine.availableBondsMap'
	commandAndRuleEngine.availableBondsMap = tempAvailableBondsMap

	// Execute Swap
	err := commandAndRuleEngine.executeTCRuleSwap108(uuidToSwapOut, &immatureElementModel)

	// Validate that there were no errors
	assert.Equal(t, "<nil>", fmt.Sprint(err))

	// Validate the result of the swap, 'B1f-TIC(B11fx-TIC(B10)-B11l)-B1l'
	// 1) Validate B1f (1)
	testCaseModelElementUuid_1 := myTestCaseModel.FirstElementUuid
	testCaseModelElement_1 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_1]

	correctElement := testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.ParentElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.PreviousElementUuid &&
		testCaseModelElement_1.MatureElementUuid == testCaseModelElement_1.FirstChildElementUuid &&
		testCaseModelElement_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 2) Validate TIC (2)
	testCaseModelElementUuid_2 := testCaseModelElement_1.NextElementUuid
	testCaseModelElement_2 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2.ParentElementUuid &&
		testCaseModelElement_2.PreviousElementUuid == testCaseModelElement_1.MatureElementUuid &&
		testCaseModelElement_1.NextElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 3) Validate B1l (3)
	testCaseModelElementUuid_3 := testCaseModelElement_2.NextElementUuid
	testCaseModelElement_3 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_3]

	correctElement = testCaseModelElement_3.MatureElementUuid == testCaseModelElement_3.ParentElementUuid &&
		testCaseModelElement_3.NextElementUuid == testCaseModelElement_3.MatureElementUuid &&
		testCaseModelElement_3.FirstChildElementUuid == testCaseModelElement_3.MatureElementUuid &&
		testCaseModelElement_3.PreviousElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2.NextElementUuid == testCaseModelElement_3.MatureElementUuid &&
		testCaseModelElement_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 4) Validate B11fx (2.1)
	testCaseModelElementUuid_2_1 := testCaseModelElement_2.FirstChildElementUuid
	testCaseModelElement_2_1 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_1]

	correctElement = testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2.FirstChildElementUuid &&
		testCaseModelElement_2_1.ParentElementUuid == testCaseModelElement_2.MatureElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.PreviousElementUuid &&
		testCaseModelElement_2_1.MatureElementUuid == testCaseModelElement_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 5) Validate TIC (2.2)
	testCaseModelElementUuid_2_2 := testCaseModelElement_2_1.NextElementUuid
	testCaseModelElement_2_2 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_2]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_2.ParentElementUuid &&
		testCaseModelElement_2_2.PreviousElementUuid == testCaseModelElement_2_1.MatureElementUuid &&
		testCaseModelElement_2_1.NextElementUuid == testCaseModelElement_2_2.MatureElementUuid &&
		testCaseModelElement_2_2.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 6) Validate B11lx (2.3)
	testCaseModelElementUuid_2_3 := testCaseModelElement_2_2.NextElementUuid
	testCaseModelElement_2_3 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_3]

	correctElement = testCaseModelElement_2.MatureElementUuid == testCaseModelElement_2_3.ParentElementUuid &&
		testCaseModelElement_2_3.FirstChildElementUuid == testCaseModelElement_2_3.MatureElementUuid &&
		testCaseModelElement_2_3.PreviousElementUuid == testCaseModelElement_2_2.MatureElementUuid &&
		testCaseModelElement_2_2.NextElementUuid == testCaseModelElement_2_3.MatureElementUuid &&
		testCaseModelElement_2_3.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// 7) Validate B10 (2.2.1)
	testCaseModelElementUuid_2_2_1 := testCaseModelElement_2_2.FirstChildElementUuid
	testCaseModelElement_2_2_1 := myTestCaseModel.TestCaseModelMap[testCaseModelElementUuid_2_2_1]

	correctElement = testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2.FirstChildElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.PreviousElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.FirstChildElementUuid &&
		testCaseModelElement_2_2_1.MatureElementUuid == testCaseModelElement_2_2_1.NextElementUuid &&
		testCaseModelElement_2_2_1.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND

	assert.Equal(t, "true", fmt.Sprint(correctElement))

	// Validate that there are no zombie elements in TestCaseModel
	err = verifyThatThereAreNoZombieElementsInTestCaseModel(&myTestCaseModel)

	assert.Equal(t, "<nil>", fmt.Sprint(err))

}
