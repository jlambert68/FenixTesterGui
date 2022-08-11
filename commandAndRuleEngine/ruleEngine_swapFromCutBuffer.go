package commandAndRuleEngine

import (
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	"fmt"
)

// Verify if anor element can be swapped or not, regarding swap rules
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) verifyIfElementCanBeSwappedForCutBuffer(testCaseUuid string, elementUuidToBeCutOut string) (canBeSwapped bool, matchedSimpledRule string, matchedComplexRule string, err error) {

	// Get current TestCase
	currentTestCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	if existsInMap == false {

		errorId := "c15ce96e-48ce-4061-9e29-e2d68d27151b"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))

		return false, "", "", err
	}

	// Verify that there are anything in Copy Buffer, use First Element as a proxy
	if currentTestCase.CutBuffer.FirstElementUuid == "" {

		errorId := "09b47520-14ff-48a4-8342-06ac58f07813"
		err = errors.New(fmt.Sprintf("there is no content in Cut Buffer for TestCase '%s' [ErrorID: %s]", testCaseUuid, errorId))

		return false, "", "", err
	}

	// Get ElementType for first element in Cut Buffer
	elementToBeSwappedIn, existsInMap := currentTestCase.CopyBuffer.ImmatureElementMap[currentTestCase.CopyBuffer.FirstElementUuid]
	if existsInMap == false {

		errorId := "52d593c3-ad7a-448d-b301-87fdedcf96b0"
		err = errors.New(fmt.Sprintf("element referenced by first element ('%s')  doesn't exist in element-map for CopyBuffer in TestCase '%s' [ErrorID: %s]", currentTestCase.CopyBuffer.FirstElementUuid, testCaseUuid, errorId))

		return false, "", "", err
	}

	elementTypeForElementToBeSwappedIn := elementToBeSwappedIn.TestCaseModelElementType

	// Forward command to swap command
	canBeSwapped, matchedSimpledRule, matchedComplexRule, err = commandAndRuleEngine.verifyIfElementCanBeSwapped(testCaseUuid, elementUuidToBeCutOut, elementTypeForElementToBeSwappedIn)

	return canBeSwapped, matchedSimpledRule, matchedComplexRule, err
}

// Swap an element, but first ensure that rules for swapping are used
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeSwapElementFromCutBuffer(testCaseUuid string, elementToSwapOutUuid string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (err error) {

	// Forward command to swap command
	err = commandAndRuleEngine.executeSwapElement(testCaseUuid, elementToSwapOutUuid, immatureElementToSwapIn)

	return err
}
