package testCaseModel

import (
	"errors"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
)

// Verify all children, in TestCaseElement-model and remove the found element from 'allUuidKeys'
func (testCaseModel *TestCasesModelsStruct) recursiveZombieElementSearchInTestCaseModel(testCaseUuid string, elementsUuid string, allUuidKeys []string) (processedAllUuidKeys []string, err error) {

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all testcases")
		return nil, err
	}

	// Extract current element
	currentElement, existInMap := currentTestCase.TestCaseModelMap[elementsUuid]

	// If the element doesn't exit then there is something really wrong
	if existInMap == false {
		// This shouldn't happen
		err = errors.New(elementsUuid + " could not be found in in map 'testCaseModel.TestCaseModelMap'")

		return nil, err
	}

	// Element has child-element then go that path
	if currentElement.MatureTestCaseModelElementMessage.FirstChildElementUuid != elementsUuid {
		allUuidKeys, err = testCaseModel.recursiveZombieElementSearchInTestCaseModel(testCaseUuid, currentElement.MatureTestCaseModelElementMessage.FirstChildElementUuid, allUuidKeys)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return nil, err
	}

	// If element has a next-element the go that path
	if currentElement.MatureTestCaseModelElementMessage.NextElementUuid != elementsUuid {
		allUuidKeys, err = testCaseModel.recursiveZombieElementSearchInTestCaseModel(testCaseUuid, currentElement.MatureTestCaseModelElementMessage.NextElementUuid, allUuidKeys)
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

// Remove 'uuid' from slice
func findElementInSliceAndRemove(sliceToWorkOn *[]string, uuid string) (returnSlice *[]string) {

	var index int
	var uuidInSLice string

	// Find the index of the 'uuid'
	for index, uuidInSLice = range *sliceToWorkOn {
		if uuidInSLice == uuid {
			break
		}
	}

	// Create a temporary slice to work on
	tempSlice := *sliceToWorkOn

	// Remove the element in the slice
	tempSlice[index] = tempSlice[len(tempSlice)-1]
	tempSlice = tempSlice[:len(tempSlice)-1]

	returnSlice = &tempSlice

	return returnSlice
}

// Generate the slice with the elements in the TestCase. Order is the same as in the Textual Representation of the TestCase
func (testCaseModel *TestCasesModelsStruct) recursiveTextualTestCaseModelExtractor(testCaseUuid string, elementsUuid string, testCaseModelElementsIn []fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage) (testCaseModelElementsIOut []fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage, err error) {

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all testcases")
		return nil, err
	}

	// Extract current element
	currentElement, existInMap := currentTestCase.TestCaseModelMap[elementsUuid]

	// If the element doesn't exit then there is something really wrong
	if existInMap == false {
		// This shouldn't happen
		err = errors.New(elementsUuid + " could not be found in in map 'testCaseModel.TestCaseModelMap'")

		return nil, err
	}

	// Add element to slice
	testCaseModelElementsIOut = append(testCaseModelElementsIn, currentElement.MatureTestCaseModelElementMessage)

	// Element has child-element then go that path
	if currentElement.MatureTestCaseModelElementMessage.FirstChildElementUuid != elementsUuid {
		testCaseModelElementsIOut, err = testCaseModel.recursiveTextualTestCaseModelExtractor(testCaseUuid, currentElement.MatureTestCaseModelElementMessage.FirstChildElementUuid, testCaseModelElementsIOut)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return nil, err
	}

	// If element has a next-element the go that path
	if currentElement.MatureTestCaseModelElementMessage.NextElementUuid != elementsUuid {
		testCaseModelElementsIOut, err = testCaseModel.recursiveTextualTestCaseModelExtractor(testCaseUuid, currentElement.MatureTestCaseModelElementMessage.NextElementUuid, testCaseModelElementsIOut)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return nil, err
	}

	return testCaseModelElementsIOut, err
}

// Generate name to be used when presenting TestCase Element
func (testCaseModel *TestCasesModelsStruct) generateUINameForTestCaseElement(element *MatureTestCaseModelElementStruct) (elementUiName string) {

	elementUiName = element.MatureTestCaseModelElementMessage.OriginalElementName + " [" + element.MatureTestCaseModelElementMessage.MatureElementUuid[0:numberOfCharactersfromUuid-1] + "]"

	return elementUiName
}

// Generate the slice with the elements in the TestCase. Order is the same as in the Textual Representation of the TestCase
func (testCaseModel *TestCasesModelsStruct) recursiveGraphicalTestCaseTreeModelExtractor(
	testCaseUuid string,
	currentElementsUuid string, treeViewNodeChildrenIn []TestCaseModelAdaptedForUiTreeDataStruct) (
	treeViewNodeChildrenOut []TestCaseModelAdaptedForUiTreeDataStruct, err error) {

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {
		errorId := "68f37aee-0b93-4d4f-9225-31ea5ccd8f8a"
		err = errors.New(fmt.Sprintf("testcase with uuid '%s' doesn't exist in map with all testcases [ErrorID: %s]", testCaseUuid, errorId))

		return nil, err
	}

	// Extract current element
	currentElement, existInMap := currentTestCase.TestCaseModelMap[currentElementsUuid]

	// If the element doesn't exit then there is something really wrong
	if existInMap == false {
		// This shouldn't happen
		errorId := "2397cd98-e5dd-4a42-9d7a-46de3f0286d2"
		err = errors.New(fmt.Sprintf("element '%s', in testcase '%s' doesn't exist in map with all testcases [ErrorID: %s]", currentElementsUuid, testCaseUuid, errorId))

		return nil, err
	}

	// Element has child-element then go that path
	if currentElement.MatureTestCaseModelElementMessage.FirstChildElementUuid != currentElementsUuid {
		treeViewNodeChildrenOut, err = testCaseModel.recursiveGraphicalTestCaseTreeModelExtractor(testCaseUuid, currentElement.MatureTestCaseModelElementMessage.FirstChildElementUuid, []TestCaseModelAdaptedForUiTreeDataStruct{})

		// reverse Slice to get correct order in Tree-view
		treeViewNodeChildrenToBeSaved := testCaseModel.reverseSliceOfNodeObjects(treeViewNodeChildrenOut)

		// Save children under currentUUid
		currentTestCase.testCaseModelAdaptedForUiTree[currentElementsUuid] = treeViewNodeChildrenToBeSaved

	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return nil, err
	}

	// If element has a next-element the go that path
	if currentElement.MatureTestCaseModelElementMessage.NextElementUuid != currentElementsUuid {
		treeViewNodeChildrenIn, err = testCaseModel.recursiveGraphicalTestCaseTreeModelExtractor(testCaseUuid, currentElement.MatureTestCaseModelElementMessage.NextElementUuid, treeViewNodeChildrenIn)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return nil, err
	}

	// Data to extract
	var (
		nodeColor                string
		testInstructionTypeColor string
		canBeDeleted             bool
		canBeSwappedOut          bool
		isBond                   bool
	)

	switch currentElement.MatureTestCaseModelElementMessage.TestCaseModelElementType {
	// B0-Bond
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND:
		nodeColor = nodeColor_Bond_B0
		isBond = true

	// B1-Bonds
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE:
		nodeColor = nodeColor_Bond_B1
		isBond = true

		// B10, B11, B12
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND:
		nodeColor = nodeColor_Swappeble_Bonds
		isBond = true

		// TI
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION:
		if currentTestCase.MatureTestInstructionMap[currentElementsUuid].MatureBasicTestInstructionInformation.ChosenDropZoneUuid == "No DropZone exists" {
			nodeColor = currentTestCase.MatureTestInstructionMap[currentElementsUuid].BasicTestInstructionInformation_NonEditableInformation.TestInstructionColor

		} else {
			nodeColor = currentTestCase.MatureTestInstructionMap[currentElementsUuid].MatureBasicTestInstructionInformation.ChosenDropZoneColor //currentElement.MatureTestCaseModelElementMetaData.ChosenDropZoneColorString //nodeColor_TI_TIC
			isBond = false
		}

		// TIC
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER:

		nodeColor = currentElement.MatureTestCaseModelElementMetaData.ChosenDropZoneColorString //nodeColor_TI_TIC
		isBond = false

	// B11x, B12x
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE:
		nodeColor = nodeColor_X_Bonds
		isBond = true

	// B10x
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10xo_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10oxo_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10ox_BOND:
		nodeColor = nodeColor_B10X_Bonds
		isBond = true

	// TIx
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE:
		nodeColor = currentTestCase.MatureTestInstructionMap[currentElementsUuid].MatureBasicTestInstructionInformation.ChosenDropZoneColor
		isBond = false

	// TICx
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE:
		nodeColor = currentElement.MatureTestCaseModelElementMetaData.ChosenDropZoneColorString // nodeColor = nodeColor_TIx_TICx
		isBond = false

	default:
		errorId := "97e0d6e6-791e-4228-8cb4-e58f059dad1e"
		err = errors.New(fmt.Sprintf("unknown Element Type '%s' in Element '%s' in TestCase '%s' [ErrorID: %s]", currentElement.MatureTestCaseModelElementMessage.TestCaseModelElementType, currentElementsUuid, testCaseUuid, errorId))

		return nil, err

	}

	// Check if it is a Bond-element
	if isBond {
		// The Element is a Bond so extract it
		currentImmatureBond, existInMap := testCaseModel.AvailableBondsMap[currentElement.MatureTestCaseModelElementMessage.TestCaseModelElementType]

		// If the element doesn't exit then there is something really wrong
		if existInMap == false {
			// This shouldn't happen
			errorId := "6c3522bb-b3fc-4b65-acb0-0090df6970b9"
			err = errors.New(fmt.Sprintf("bond element '%s', doesn't exist in map with all Bonds [ErrorID: %s]", currentElement.MatureTestCaseModelElementMessage.TestCaseModelElementType, errorId))

			return nil, err
		}

		// Extract the data
		canBeDeleted = currentImmatureBond.BasicBondInformation.VisibleBondAttributes.CanBeDeleted
		canBeSwappedOut = currentImmatureBond.BasicBondInformation.VisibleBondAttributes.CanBeSwappedOut

	}

	// Set TestInstruction type color
	if currentElement.MatureTestCaseModelElementMessage.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
		currentElement.MatureTestCaseModelElementMessage.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE {
		testInstructionTypeColor = "#FFFF0033"
	} else {
		// Use transparent, alfa channel added,  if no TestInstruction
		testInstructionTypeColor = "#00000000"
	}

	// Add element to slice
	elementDataToAdd := TestCaseModelAdaptedForUiTreeDataStruct{
		Uuid:                     currentElementsUuid,
		OriginalUuid:             currentElement.MatureTestCaseModelElementMessage.OriginalElementUuid,
		NodeName:                 currentElement.MatureTestCaseModelElementMessage.OriginalElementName,
		NodeColor:                nodeColor,
		TestInstructionTypeColor: testInstructionTypeColor,
		NodeTypeEnum:             currentElement.MatureTestCaseModelElementMessage.TestCaseModelElementType,
		CanBeDeleted:             canBeDeleted,
		CanBeSwappedOut:          canBeSwappedOut,
	}
	treeViewNodeChildrenIn = append(treeViewNodeChildrenIn, elementDataToAdd)

	// Save the element with itself as original-uuid-key, to be able to find its data
	currentTestCase.testCaseModelAdaptedForUiTree[currentElementsUuid+"_originalUuid"] = []TestCaseModelAdaptedForUiTreeDataStruct{elementDataToAdd}

	// Save Top-Left-Children in Map with ParentElementUuid as map-key
	if currentElementsUuid == currentElement.MatureTestCaseModelElementMessage.ParentElementUuid &&
		currentElementsUuid == currentElement.MatureTestCaseModelElementMessage.PreviousElementUuid {
		// reverse Slice to get correct order in Tree-view
		treeViewNodeChildrenToBeSaved := testCaseModel.reverseSliceOfNodeObjects(treeViewNodeChildrenIn)

		// Children has no defined parent, put them under "standard" Fyne UI Tree-component Top Node, ("")
		currentTestCase.testCaseModelAdaptedForUiTree[""] = treeViewNodeChildrenToBeSaved

	}

	return treeViewNodeChildrenIn, err
}

// Reverse a slice of strings
func (testCaseModel *TestCasesModelsStruct) reverseSliceOfNodeObjects(inSlice []TestCaseModelAdaptedForUiTreeDataStruct) (outSlice []TestCaseModelAdaptedForUiTreeDataStruct) {

	numberOfElementsInSlice := len(inSlice)

	for positionCounter := numberOfElementsInSlice - 1; positionCounter >= 0; positionCounter-- {
		outSlice = append(outSlice, inSlice[positionCounter])
	}
	return outSlice
}
