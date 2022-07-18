package commandAndRuleEngine

import (
	"errors"
	"github.com/sirupsen/logrus"
)

// Verify if an element can be swapped or not, regarding swap rules
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) verifyIfElementCanBeSwapped(elementUuid string) (canBeSwapped bool, matchedSimpledRule string, matchedComplexRule string, err error) {

	// First verify towards simple rules
	canBeSwapped, matchedSimpledRule, err = commandAndRuleEngineObject.verifyIfComponentCanBeSwappedSimpleRules(elementUuid)

	// Only check complex rules if simple rules was OK for swapping
	if !(canBeSwapped == true &&
		err != nil) {
		return canBeSwapped, matchedSimpledRule, "", err
	}

	// Verify towards complex rules
	matchedComplexRule, err = commandAndRuleEngineObject.verifyIfComponentCanBeSwappedWithComplexRules(elementUuid)

	return canBeSwapped, matchedSimpledRule, matchedComplexRule, err
}

// Swap an element, but first ensure that rules for swapping are used
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) executeSwapElement(elementUuid string) (err error) {

	// Verify that element is allowed, and can be swapped
	canBeSwapped, matchedSimpleRule, _, err := commandAndRuleEngineObject.verifyIfElementCanBeSwapped(elementUuid)

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
	err = commandAndRuleEngineObject.executeSwapElement(elementUuid)

	return err
}

// Execute a swap on an element based on specific rule
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) executeSwapElementBasedOnRule(elementUuid string, immatureElementToSwapIn *immatureElementStruct, matchedComplexRule string) (err error) {

	switch matchedComplexRule {
	case TCRuleSwap101:
		err = commandAndRuleEngineObject.executeTCRuleSwap101(elementUuid, immatureElementToSwapIn)

	case TCRuleSwap102:
		err = commandAndRuleEngineObject.executeTCRuleSwap102(elementUuid, immatureElementToSwapIn)

	case TCRuleSwap103:
		err = commandAndRuleEngineObject.executeTCRuleSwap103(elementUuid, immatureElementToSwapIn)

	case TCRuleSwap104:
		err = commandAndRuleEngineObject.executeTCRuleSwap104(elementUuid, immatureElementToSwapIn)

	case TCRuleSwap105:
		commandAndRuleEngineObject.executeTCRuleSwap105(elementUuid, immatureElementToSwapIn)

	case TCRuleSwap106:
		err = commandAndRuleEngineObject.executeTCRuleSwap106(elementUuid, immatureElementToSwapIn)

	case TCRuleSwap107:
		err = commandAndRuleEngineObject.executeTCRuleSwap107(elementUuid, immatureElementToSwapIn)

	case TCRuleSwap108:
		err = commandAndRuleEngineObject.executeTCRuleSwap108(elementUuid, immatureElementToSwapIn)

	default:
		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":                 "eba1a213-aa42-4021-aaea-4b3107d5874c",
			"matchedComplexRule": matchedComplexRule,
		}).Error(" Unknown 'matchedComplexRule' was used when trying to swap")

		err = errors.New("'" + matchedComplexRule + "' is an unknown complex Swap rule")

	}

	return err

}
