package commandAndRuleEngine

import (
	"errors"
	"fmt"
)

// Verify if an element can be swapped for copy Buffer or not, regarding swap rules
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) verifyIfElementCanBeSwappedForCopyBuffer(testCaseUuid string, elementUuid string) (canBeSwapped bool, matchedSimpledRule string, matchedComplexRule string, err error) {

	// Get current TestCase
	currentTestCase, existsInMap := commandAndRuleEngine.testcases.TestCases[testCaseUuid]
	if existsInMap == false {

		errorId := "0b058d20-5ce3-4da7-b9cf-c71c47b5f72d"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))

		return false, "", "", err
	}

	// Verify that there are anything in Copy Buffer, use First Element as a proxy
	if currentTestCase.CopyBuffer.FirstElementUuid == "" {

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

	// Verify towards complex rules
	matchedComplexRule, err = commandAndRuleEngine.verifyIfComponentCanBeSwappedWithComplexRules(testCaseUuid, elementUuid)

	return canBeSwapped, matchedSimpledRule, matchedComplexRule, err
}

// Swap an element for content in Copy Buffer, but first ensure that rules for swapping are used
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeSwapElementForCopyBuffer(testCaseUuid string, elementToSwapOutUuid string) (err error) {

	// Verify that element is allowed, and can be swapped
	canBeSwapped, matchedSimpleRule, matchedComplexRule, err := commandAndRuleEngine.verifyIfElementCanBeSwapped(testCaseUuid, elementToSwapOutUuid)

	// If there was an error from swap verification then exit
	if err != nil {
		return err
	}

	// If the component couldn't be swapped then exit with error message
	if canBeSwapped == false {
		err = errors.New("element couldn't be swapped due to swap rule '" + matchedSimpleRule + "' or that complex rules aren't met")

		return err
	}

	// Get current TestCase
	currentTestCase, existsInMap := commandAndRuleEngine.testcases.TestCases[testCaseUuid]
	if existsInMap == false {

		errorId := "6d6fff2c-a007-485e-8bef-a67d58eac518"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))

		return err
	}

	// Extract element from Copy Buffer
	immatureElementToSwapInCopyBuffer := currentTestCase.CopyBuffer

	// Execute swap out element for copy buffer content
	err = commandAndRuleEngine.executeSwapElementBasedOnRule(testCaseUuid, elementToSwapOutUuid, &immatureElementToSwapInCopyBuffer, matchedComplexRule)

	return err
}
