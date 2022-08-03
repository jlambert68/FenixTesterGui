package commandAndRuleEngine

import (
	"FenixTesterGui/testCase/testCaseModel"
)

// Verify if anor element can be swapped or not, regarding swap rules
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) verifyIfElementCanBeSwappedForCutBuffer(testCaseUuid string, elementUuid string) (canBeSwapped bool, matchedSimpledRule string, matchedComplexRule string, err error) {

	// Forward command to swap command
	canBeSwapped, matchedSimpledRule, matchedComplexRule, err = commandAndRuleEngine.verifyIfElementCanBeSwapped(testCaseUuid, elementUuid)

	return canBeSwapped, matchedSimpledRule, matchedComplexRule, err
}

// Swap an element, but first ensure that rules for swapping are used
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeSwapElementFromCutBuffer(testCaseUuid string, elementToSwapOutUuid string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (err error) {

	// Forward command to swap command
	err = commandAndRuleEngine.executeSwapElement(testCaseUuid, elementToSwapOutUuid, immatureElementToSwapIn)

	return err
}
