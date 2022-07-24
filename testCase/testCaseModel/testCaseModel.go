package testCaseModel

import (
	"errors"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"strings"
)

func (testCaseModel *TestCaseModelStruct) CreateTextualTestCase() (textualTestCase string, err error) {

	// Create slice with all elementTypes in presentation order
	testCaseModelElements, err := testCaseModel.recursiveTextualTestCaseModelExtractor(testCaseModel.FirstElementUuid, []fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum{})

	// Something wrong happen
	if err != nil {
		return "", err
	}

	// If there are no elements in TestCaseModel then return empty Textual description
	if len(testCaseModelElements) == 0 {
		return "[]", nil
	}

	// Loop all elements and convert element type into presentation representation
	for _, testCaseModelElementType := range testCaseModelElements {
		presentationName := fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_name[int32(testCaseModelElementType)]

		separatorIndex := strings.Index(presentationName, "_")
		presentationName = presentationName[:separatorIndex]

		switch testCaseModelElementType {

		// First element in TC
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE:

			presentationName = "[" + presentationName

			// Last element in TC
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE:

			presentationName = "-" + presentationName + "]"

			// The only element in TC
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND:

			presentationName = "[" + presentationName + "]"

		// First element child in TIC(x)
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE:

			presentationName = "(" + presentationName

		// Last element child in TIC(x)
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE:

			presentationName = "-" + presentationName + ")"

			// The only element in TIC
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10ox_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10oxo_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10xo_BOND:

			presentationName = "(" + presentationName + ")"

			// Element surrounded with other elements
		case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND,
			fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE:

			presentationName = "-" + presentationName

			// No match in element
		default:

			err = errors.New("no match in element type for: " + testCaseModelElementType.String())
			return "", err

		}

		// Add presentation name to Textual TestCase
		textualTestCase = textualTestCase + presentationName

	}

	return textualTestCase, nil
}

// Verify all children, in TestCaseElement-model and remove the found element from 'allUuidKeys'
func (testCaseModel *TestCaseModelStruct) recursiveTextualTestCaseModelExtractor(elementsUuid string, testCaseModelElementsIn []fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum) (testCaseModelElementsIOut []fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum, err error) {

	var

	// Extract current element
	currentElement, existInMap = testCaseModel.TestCaseModelMap[elementsUuid]

	// If the element doesn't exit then there is something really wrong
	if existInMap == false {
		// This shouldn't happen
		err = errors.New(elementsUuid + " could not be found in in map 'testCaseModel.TestCaseModelMap'")

		return nil, err
	}

	// Add elementType to slice
	testCaseModelElementsIOut = append(testCaseModelElementsIn, currentElement.TestCaseModelElementType)

	// Element has child-element then go that path
	if currentElement.FirstChildElementUuid != elementsUuid {
		testCaseModelElementsIOut, err = testCaseModel.recursiveTextualTestCaseModelExtractor(currentElement.FirstChildElementUuid, testCaseModelElementsIOut)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return nil, err
	}

	// If element has a next-element the go that path
	if currentElement.NextElementUuid != elementsUuid {
		testCaseModelElementsIOut, err = testCaseModel.recursiveTextualTestCaseModelExtractor(currentElement.NextElementUuid, testCaseModelElementsIOut)
	}

	// If we got an error back then something wrong happen, so just back out
	if err != nil {
		return nil, err
	}

	return testCaseModelElementsIOut, err
}
