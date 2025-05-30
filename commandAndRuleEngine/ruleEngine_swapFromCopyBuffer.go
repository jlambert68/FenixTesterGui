package commandAndRuleEngine

import (
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	"fmt"
)

// Verify if an element can be swapped for copy Buffer or not, regarding swap rules
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) verifyIfElementCanBeSwappedForCopyBuffer(testCaseUuid string, elementUuid string) (canBeSwapped bool, matchedSimpledRule string, matchedComplexRule string, err error) {

	var existsInMap bool

	// Get current TestCase
	// Get TestCasesMap
	var testCasesMap map[string]*testCaseModel.TestCaseModelStruct
	testCasesMap = *commandAndRuleEngine.Testcases.TestCasesMapPtr

	// Get current TestCase
	var currentTestCasePtr *testCaseModel.TestCaseModelStruct
	currentTestCasePtr, existsInMap = testCasesMap[testCaseUuid]

	if existsInMap == false {

		errorId := "0b058d20-5ce3-4da7-b9cf-c71c47b5f72d"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCasesMapPtr [ErrorID: %s]", testCaseUuid, errorId))

		return false, "", "", err
	}

	// Verify that there are anything in Copy Buffer, use First Element as a proxy
	if currentTestCasePtr.CopyBuffer.FirstElementUuid == "" {

		errorId := "cdb36ada-03c6-4dbb-be4b-4617cfd4f383"
		err = errors.New(fmt.Sprintf("there is no content in Copy Buffer for TestCase '%s' [ErrorID: %s]", testCaseUuid, errorId))

		return false, "", "", err
	}

	// First verify towards simple rules
	canBeSwapped, matchedSimpledRule, err = commandAndRuleEngine.verifyIfComponentCanBeSwappedSimpleRules(testCaseUuid, elementUuid)

	// Only check complex rules if simple rules was OK for swapping
	if !(canBeSwapped == true &&
		err == nil) {
		return canBeSwapped, matchedSimpledRule, "", err
	}

	// Get ElementType for first element in Copy Buffer
	elementToBeSwappedIn, existsInMap := currentTestCasePtr.CopyBuffer.ImmatureElementMap[currentTestCasePtr.CopyBuffer.FirstElementUuid]
	if existsInMap == false {

		errorId := "52d593c3-ad7a-448d-b301-87fdedcf96b0"
		err = errors.New(fmt.Sprintf("element referenced by first element ('%s')  doesn't exist in element-map for CopyBuffer in TestCase '%s' [ErrorID: %s]", currentTestCasePtr.CopyBuffer.FirstElementUuid, testCaseUuid, errorId))

		return false, "", "", err
	}

	elementTypeForElementToBeSwappedIn := elementToBeSwappedIn.TestCaseModelElementType

	// Verify towards complex rules
	matchedComplexRule, err = commandAndRuleEngine.verifyIfComponentCanBeSwappedWithComplexRules(testCaseUuid, elementUuid, elementTypeForElementToBeSwappedIn)

	return canBeSwapped, matchedSimpledRule, matchedComplexRule, err
}

// Swap an element for content in Copy Buffer, but first ensure that rules for swapping are used
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeSwapElementForCopyBuffer(testCaseUuid string, elementToSwapOutUuid string) (err error) {

	var existsInMap bool

	// Get current TestCase
	// Get TestCasesMap
	var testCasesMap map[string]*testCaseModel.TestCaseModelStruct
	testCasesMap = *commandAndRuleEngine.Testcases.TestCasesMapPtr

	// Get current TestCase
	var currentTestCasePtr *testCaseModel.TestCaseModelStruct
	currentTestCasePtr, existsInMap = testCasesMap[testCaseUuid]

	if existsInMap == false {

		errorId := "91de8132-1950-41be-b567-12fe388b0440"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCasesMapPtr [ErrorID: %s]", testCaseUuid, errorId))

		return err
	}

	// Verify that there are anything in Copy Buffer, use First Element as a proxy
	if currentTestCasePtr.CopyBuffer.FirstElementUuid == "" {

		errorId := "fc59e5d6-880e-42e7-b7a9-0b221a00825b"
		err = errors.New(fmt.Sprintf("there is no content in Copy Buffer for TestCase '%s' [ErrorID: %s]", testCaseUuid, errorId))

		return err
	}

	// Get ElementType for first element in Copy Buffer
	elementToBeSwappedIn, existsInMap := currentTestCasePtr.CopyBuffer.ImmatureElementMap[currentTestCasePtr.CopyBuffer.FirstElementUuid]
	if existsInMap == false {

		errorId := "52d593c3-ad7a-448d-b301-87fdedcf96b0"
		err = errors.New(fmt.Sprintf("element referenced by first element ('%s')  doesn't exist in element-map for CopyBuffer in TestCase '%s' [ErrorID: %s]", currentTestCasePtr.CopyBuffer.FirstElementUuid, testCaseUuid, errorId))

		return err
	}

	elementTypeToSwapIn := elementToBeSwappedIn.TestCaseModelElementType

	// Verify that element is allowed, and can be swapped
	canBeSwapped, matchedSimpleRule, matchedComplexRule, err := commandAndRuleEngine.verifyIfElementCanBeSwapped(testCaseUuid, elementToSwapOutUuid, elementTypeToSwapIn)

	// If there was an error from swap verification then exit
	if err != nil {
		return err
	}

	// If the component couldn't be swapped then exit with error message
	if canBeSwapped == false {
		err = errors.New("element couldn't be swapped due to swap rule '" + matchedSimpleRule + "' or that complex rules aren't met")

		return err
	}

	// Extract element from Copy Buffer
	immatureElementToSwapInFromCopyBuffer := currentTestCasePtr.CopyBuffer

	// Execute swap out element for copy buffer content
	err = commandAndRuleEngine.executeSwapElementBasedOnRule(testCaseUuid, elementToSwapOutUuid, &immatureElementToSwapInFromCopyBuffer, matchedComplexRule)

	return err
}
