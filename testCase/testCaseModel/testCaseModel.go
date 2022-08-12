package testCaseModel

import (
	"errors"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"strings"
)

// Verify that all UUIDs are correct in TestCaseModel. Meaning that no empty uuid is allowed and they all are correct
func (testCaseModel *TestCasesModelsStruct) VerifyThatThereAreNoZombieElementsInTestCaseModel(testCaseUuid string) (err error) {

	var allUuidKeys []string

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {
		err = errors.New("testcase with uuid '" + testCaseUuid + "' doesn't exist in map with all testcases")
		return err
	}

	// Extract all elements by key from TestCaseModel
	for _, elementKey := range currentTestCase.TestCaseModelMap {
		allUuidKeys = append(allUuidKeys, elementKey.MatureElementUuid)
	}

	// Follow the path from "first element and remove the found element from 'allUuidKeys'
	allUuidKeys, err = testCaseModel.recursiveZombieElementSearchInTestCaseModel(testCaseUuid, currentTestCase.FirstElementUuid, allUuidKeys)

	// If there are elements left in slice then there were zombie elements, which there shouldn't be
	if len(allUuidKeys) != 0 {
		err = errors.New("there existed Zombie elements in 'testCaseModel.TestCaseModelMap', for " + currentTestCase.FirstElementUuid)

		return err
	}

	return err
}

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
	if currentElement.FirstChildElementUuid != elementsUuid {
		allUuidKeys, err = testCaseModel.recursiveZombieElementSearchInTestCaseModel(testCaseUuid, currentElement.FirstChildElementUuid, allUuidKeys)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return nil, err
	}

	// If element has a next-element the go that path
	if currentElement.NextElementUuid != elementsUuid {
		allUuidKeys, err = testCaseModel.recursiveZombieElementSearchInTestCaseModel(testCaseUuid, currentElement.NextElementUuid, allUuidKeys)
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

func (testCaseModel *TestCasesModelsStruct) CreateTextualTestCase(testCaseUuid string) (textualTestCaseSimple string, textualTestCaseComplex string, textualTestCaseExtended string, err error) {

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {

		errorId := "591afb7e-a372-45d4-88c0-332535642a3b"
		err = errors.New(fmt.Sprintf("testcase with uuid '%s'  doesn't exist in map with all testcases [ErrorID: %s]", testCaseUuid, errorId))

		return "", "", "", err
	}

	// Create slice with all elementTypes in presentation order
	testCaseModelElements, err := testCaseModel.recursiveTextualTestCaseModelExtractor(testCaseUuid, currentTestCase.FirstElementUuid, []fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{})

	// Something wrong happen
	if err != nil {
		return "", "", "", err
	}

	// If there are no elements in TestCaseModel then return empty Textual description
	if len(testCaseModelElements) == 0 {
		return "{}", "{}", "{}", nil
	}

	// Loop all elements and convert element type into presentation representation
	for _, testCaseModelElement := range testCaseModelElements {

		// Get short UUID for
		shourtUuid := testCaseModel.GenerateShortUuidFromFullUuid(testCaseModelElement.MatureElementUuid)

		// Simple presentation style, like 'B10x' for "B10oxo"
		presentationNameSimple := fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementPresentationNameEnum_name[int32(testCaseModelElement.TestCaseModelElementType)]
		separatorIndexSimple := strings.Index(presentationNameSimple, "_")
		presentationNameSimple = presentationNameSimple[:separatorIndexSimple]

		// Complex presentation style, like 'B10oxo' for "B10oxo"
		presentationNameComplex := fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_name[int32(testCaseModelElement.TestCaseModelElementType)]
		separatorIndexComplex := strings.Index(presentationNameComplex, "_")
		presentationNameComplex = presentationNameComplex[:separatorIndexComplex]

		// Extended presentation style, like 'B10oxo[0291462]' for "B10oxo"
		presentationNameExtended := fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_name[int32(testCaseModelElement.TestCaseModelElementType)]
		separatorIndexExtended := strings.Index(presentationNameExtended, "_")
		presentationNameExtended = presentationNameExtended[:separatorIndexExtended] + "[" + shourtUuid + "]"

		switch testCaseModelElement.TestCaseModelElementType {

		// First element in TC
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE:

			presentationNameSimple = "{" + presentationNameSimple
			presentationNameComplex = "{" + presentationNameComplex
			presentationNameExtended = "{" + presentationNameExtended

			// Last element in TC
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE:

			presentationNameSimple = "-" + presentationNameSimple + "}"
			presentationNameComplex = "-" + presentationNameComplex + "}"
			presentationNameExtended = "-" + presentationNameExtended + "}"

			// The only element in TC
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND:

			presentationNameSimple = "{" + presentationNameSimple + "}"
			presentationNameComplex = "{" + presentationNameComplex + "}"
			presentationNameExtended = "{" + presentationNameExtended + "}"

		// First element child in TIC(x)
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE:

			presentationNameSimple = "(" + presentationNameSimple
			presentationNameComplex = "(" + presentationNameComplex
			presentationNameExtended = "(" + presentationNameExtended

		// Last element child in TIC(x)
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE:

			presentationNameSimple = "-" + presentationNameSimple + ")"
			presentationNameComplex = "-" + presentationNameComplex + ")"
			presentationNameExtended = "-" + presentationNameExtended + ")"

			// The only element in TIC
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10ox_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10oxo_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10xo_BOND:

			presentationNameSimple = "(" + presentationNameSimple + ")"
			presentationNameComplex = "(" + presentationNameComplex + ")"
			presentationNameExtended = "(" + presentationNameExtended + ")"

			// Element surrounded with other elements
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE:

			presentationNameSimple = "-" + presentationNameSimple
			presentationNameComplex = "-" + presentationNameComplex
			presentationNameExtended = "-" + presentationNameExtended

			// No match in element
		default:

			errorId := "46d49c79-93a4-4e56-a74c-9c08a43b26e8"
			err = errors.New(fmt.Sprintf("no match in element type for: '%s' [ErrorID: %s]", testCaseModelElement.TestCaseModelElementType.String(), errorId))

			return "", "", "", err

		}

		// Add presentation name to Textual TestCase
		textualTestCaseSimple = textualTestCaseSimple + presentationNameSimple
		textualTestCaseComplex = textualTestCaseComplex + presentationNameComplex
		textualTestCaseExtended = textualTestCaseExtended + presentationNameExtended

	}

	return textualTestCaseSimple, textualTestCaseComplex, textualTestCaseExtended, err
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
	testCaseModelElementsIOut = append(testCaseModelElementsIn, currentElement)

	// Element has child-element then go that path
	if currentElement.FirstChildElementUuid != elementsUuid {
		testCaseModelElementsIOut, err = testCaseModel.recursiveTextualTestCaseModelExtractor(testCaseUuid, currentElement.FirstChildElementUuid, testCaseModelElementsIOut)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return nil, err
	}

	// If element has a next-element the go that path
	if currentElement.NextElementUuid != elementsUuid {
		testCaseModelElementsIOut, err = testCaseModel.recursiveTextualTestCaseModelExtractor(testCaseUuid, currentElement.NextElementUuid, testCaseModelElementsIOut)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return nil, err
	}

	return testCaseModelElementsIOut, err
}

// List ALL Building Blocks in TestCase
func (testCaseModel *TestCasesModelsStruct) ListAllAvailableBuildingBlocksInTestCase(testCaseUuid string) (availableBuidlingBlocksInTestCaseList []string, err error) {

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]

	if existsInMap == false {
		errorId := "02914625-46a8-4174-800a-f519f4cf0532"
		err = errors.New(fmt.Sprintf("testcase with uuid '%s' doesn't exist in map with all testcases [ErrorID: %s]", testCaseUuid, errorId))

		return nil, err
	}

	// Loop all available building blocks and create list to be used in DropDown
	for _, element := range currentTestCase.TestCaseModelMap {

		elementUiName := testCaseModel.generateUINameForTestCaseElement(&element)

		switch element.TestCaseModelElementType {

		// TestInstructions
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE:

			availableBuidlingBlocksInTestCaseList = append(availableBuidlingBlocksInTestCaseList, elementUiName+" [TI]")

			// TestInstructionContainers
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE:

			availableBuidlingBlocksInTestCaseList = append(availableBuidlingBlocksInTestCaseList, elementUiName+" [TIC]")

			// Bonds
		default:
			availableBuidlingBlocksInTestCaseList = append(availableBuidlingBlocksInTestCaseList, elementUiName+" [BOND]")
		}
	}

	return availableBuidlingBlocksInTestCaseList, err
}

// Generate name to be used when presenting TestCase Element
func (testCaseModel *TestCasesModelsStruct) generateUINameForTestCaseElement(element *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage) (elementUiName string) {

	elementUiName = element.OriginalElementName + " [" + element.MatureElementUuid[0:numberOfCharactersfromUuid-1] + "]"

	return elementUiName
}

// ListAvailableTestCases
// List all available TestCase in TestCasesModel
func (testCaseModel *TestCasesModelsStruct) ListAvailableTestCases() (availableTestCasesAsList []string) {

	// Loop all available TestCases and append  UUID for TestCase to list
	for testCaseUuid, _ := range testCaseModel.TestCases {

		availableTestCasesAsList = append(availableTestCasesAsList, testCaseUuid)

	}

	return availableTestCasesAsList
}

// GetUuidFromUiName
// Finds the UUID for from a UI-name like ' B0_BOND [3c8a3bc] [BOND] to live forever..'
func (testCaseModel *TestCasesModelsStruct) GetUuidFromUiName(testCaseUuid string, uiName string) (elementUuid string, err error) {

	// Get first square brackets, for part of UUID
	firstSquareBracketStart := strings.Index(uiName, "[")
	firstSquareBracketEnd := strings.Index(uiName, "]")

	// Get second square brackets, for type
	secondSquareBracketStart := strings.Index(uiName[firstSquareBracketEnd+1:], "[")
	secondSquareBracketEnd := strings.Index(uiName[firstSquareBracketEnd+1:], "]")

	// Extract UUID-part
	uuidPart := uiName[firstSquareBracketStart+1 : firstSquareBracketEnd]

	// Extract Type
	elementTypeFromName := uiName[firstSquareBracketEnd+1:][secondSquareBracketStart+1 : secondSquareBracketEnd]

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]

	if existsInMap == false {
		errorId := "b04c16dc-ff83-4f53-908c-4b2483cfb01a"
		err = errors.New(fmt.Sprintf("testcase with uuid '%s' doesn't exist in map with all testcases [ErrorID: %s]", testCaseUuid, errorId))

		return "", err
	}

	// Loop all available building blocks and create list to be used in DropDown
	var element fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage
	for elementUuid, element = range currentTestCase.TestCaseModelMap {

		switch elementTypeFromName {

		// TestInstructions
		case "TI":
			if (element.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
				element.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE) &&
				element.MatureElementUuid[:len(uuidPart)] == uuidPart {

				return elementUuid, nil
			}

			// TestInstructionContainers
		case "TIC":
			if (element.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER ||
				element.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE) &&
				element.MatureElementUuid[:len(uuidPart)] == uuidPart {

				return elementUuid, nil
			}

			// Bonds
		default:
			if element.MatureElementUuid[:len(uuidPart)] == uuidPart {

				return elementUuid, nil
			}

		}
	}

	errorId := "19144095-966d-4974-be34-2a33d6309758"
	err = errors.New(fmt.Sprintf("couldn't find element with UI-name '%s' in testcase '%s' [ErrorID: %s]", uiName, testCaseUuid, errorId))

	return "", err

	return "elementUuid", err
}

// GenerateShortUuidFromFullUuid
// Generate a short version of the UUID to be used in GUI
func (testCaseModel *TestCasesModelsStruct) GenerateShortUuidFromFullUuid(fullUuid string) (shortUuid string) {

	shortUuid = fullUuid[0 : numberOfCharactersfromUuid-1]

	return shortUuid
}

// Generate a short version of the UUID to be used in GUI
func (testCaseModel *TestCasesModelsStruct) GetTestCaseNameUuid(testCaseUuid string) (testCaseName string, err error) {

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]

	if existsInMap == false {
		errorId := "97198543-7717-4925-8643-240ad34bb205"
		err = errors.New(fmt.Sprintf("testcase with uuid '%s' doesn't exist in map with all testcases [ErrorID: %s]", testCaseUuid, errorId))

		return "", err
	}

	testCaseName = currentTestCase.LocalTestCaseMessage.BasicTestCaseInformationMessageEditableInformation.TestCaseName

	return testCaseName, err
}
