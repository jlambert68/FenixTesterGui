package commandAndRuleEngine

import (
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	"github.com/sirupsen/logrus"
)

// Verify if anor element can be swapped or not, regarding swap rules
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) verifyIfElementCanBeSwapped(testCaseUuid string, elementUuid string) (canBeSwapped bool, matchedSimpledRule string, matchedComplexRule string, err error) {

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

// Swap an element, but first ensure that rules for swapping are used
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeSwapElement(testCaseUuid string, elementToSwapOutUuid string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (err error) {

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

	// Execute deletion of element
	err = commandAndRuleEngine.executeSwapElementBasedOnRule(testCaseUuid, elementToSwapOutUuid, immatureElementToSwapIn, matchedComplexRule)

	return err
}

// Execute a swap on an element based on specific rule
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeSwapElementBasedOnRule(testCaseUuid string, elementUuid string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct, matchedComplexRule string) (err error) {

	switch matchedComplexRule {
	case TCRuleSwap101:
		err = commandAndRuleEngine.executeTCRuleSwap101(testCaseUuid, elementUuid, immatureElementToSwapIn)

	case TCRuleSwap102:
		err = commandAndRuleEngine.executeTCRuleSwap102(testCaseUuid, elementUuid, immatureElementToSwapIn)

	case TCRuleSwap103:
		err = commandAndRuleEngine.executeTCRuleSwap103(testCaseUuid, elementUuid, immatureElementToSwapIn)

	case TCRuleSwap104:
		err = commandAndRuleEngine.executeTCRuleSwap104(testCaseUuid, elementUuid, immatureElementToSwapIn)

	case TCRuleSwap105:
		commandAndRuleEngine.executeTCRuleSwap105(testCaseUuid, elementUuid, immatureElementToSwapIn)

	case TCRuleSwap106:
		err = commandAndRuleEngine.executeTCRuleSwap106(testCaseUuid, elementUuid, immatureElementToSwapIn)

	case TCRuleSwap107:
		err = commandAndRuleEngine.executeTCRuleSwap107(testCaseUuid, elementUuid, immatureElementToSwapIn)

	case TCRuleSwap108:
		err = commandAndRuleEngine.executeTCRuleSwap108(testCaseUuid, elementUuid, immatureElementToSwapIn)

	default:
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":                 "eba1a213-aa42-4021-aaea-4b3107d5874c",
			"matchedComplexRule": matchedComplexRule,
		}).Error(" Unknown 'matchedComplexRule' was used when trying to swap")

		err = errors.New("'" + matchedComplexRule + "' is an unknown complex Swap rule")

	}

	return err

}
