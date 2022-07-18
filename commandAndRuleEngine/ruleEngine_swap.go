package commandAndRuleEngine

import (
	"errors"
	"github.com/sirupsen/logrus"
)

// Verify if an element can be swapped or not, regarding swap rules
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) verifyIfElementCanBeSwapped(elementUuid string) (canBeSwapped bool, matchedSimpledRule string, matchedComplexRule string, err error) {

	// First verify towards simple rules
	canBeSwapped, matchedSimpledRule, err = commandAndRuleEngine.verifyIfComponentCanBeSwappedSimpleRules(elementUuid)

	// Only check complex rules if simple rules was OK for swapping
	if !(canBeSwapped == true &&
		err == nil) {
		return canBeSwapped, matchedSimpledRule, "", err
	}

	// Verify towards complex rules
	matchedComplexRule, err = commandAndRuleEngine.verifyIfComponentCanBeSwappedWithComplexRules(elementUuid)

	return canBeSwapped, matchedSimpledRule, matchedComplexRule, err
}

// Swap an element, but first ensure that rules for swapping are used
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeSwapElement(elementUuid string) (err error) {

	// Verify that element is allowed, and can be swapped
	canBeSwapped, matchedSimpleRule, _, err := commandAndRuleEngine.verifyIfElementCanBeSwapped(elementUuid)

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
	err = commandAndRuleEngine.executeSwapElement(elementUuid)

	return err
}

// Execute a swap on an element based on specific rule
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeSwapElementBasedOnRule(elementUuid string, immatureElementToSwapIn *immatureElementStruct, matchedComplexRule string) (err error) {

	switch matchedComplexRule {
	case TCRuleSwap101:
		err = commandAndRuleEngine.executeTCRuleSwap101(elementUuid, immatureElementToSwapIn)

	case TCRuleSwap102:
		err = commandAndRuleEngine.executeTCRuleSwap102(elementUuid, immatureElementToSwapIn)

	case TCRuleSwap103:
		err = commandAndRuleEngine.executeTCRuleSwap103(elementUuid, immatureElementToSwapIn)

	case TCRuleSwap104:
		err = commandAndRuleEngine.executeTCRuleSwap104(elementUuid, immatureElementToSwapIn)

	case TCRuleSwap105:
		commandAndRuleEngine.executeTCRuleSwap105(elementUuid, immatureElementToSwapIn)

	case TCRuleSwap106:
		err = commandAndRuleEngine.executeTCRuleSwap106(elementUuid, immatureElementToSwapIn)

	case TCRuleSwap107:
		err = commandAndRuleEngine.executeTCRuleSwap107(elementUuid, immatureElementToSwapIn)

	case TCRuleSwap108:
		err = commandAndRuleEngine.executeTCRuleSwap108(elementUuid, immatureElementToSwapIn)

	default:
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":                 "eba1a213-aa42-4021-aaea-4b3107d5874c",
			"matchedComplexRule": matchedComplexRule,
		}).Error(" Unknown 'matchedComplexRule' was used when trying to swap")

		err = errors.New("'" + matchedComplexRule + "' is an unknown complex Swap rule")

	}

	return err

}
