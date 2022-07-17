package commandAndRuleEngine

import (
	"errors"
	"github.com/sirupsen/logrus"
)

// Verify if an element can be swapped or not, regarding swap rules
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) verifyIfElementCanBeSwapped(elementUuid string) (canBeSwapped bool, matchedSimpldRule string, matchedComplexRule string, err error) {

	// First verify towards simple rules
	canBeSwapped, matchedSimpldRule, err = commandAndRuleEngineObject.verifyIfComponentCanBeSwappedSimpleRules(elementUuid)

	// Only check complex rules if simple rules was OK for swapping
	if !(canBeSwapped == true &&
		err != nil) {
		return canBeSwapped, matchedSimpldRule, "", err
	}

	// Verify towards complex rules
	matchedComplexRule, err = commandAndRuleEngineObject.verifyIfComponentCanBeSwappedWithComplexRules(elementUuid)

	return canBeSwapped, matchedSimpldRule, matchedComplexRule, err
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

// Delete an element based on specific rule
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) executeSwapElementBasedOnRule(elementUuid string, matchedComplexRule string) (err error) {

	switch matchedComplexRule {
	case TCRuleSwap101:
		err = commandAndRuleEngineObject.executeTCRuleSwap101(elementUuid, immatureElementToSwapIn*immatureElementStruct)

	case TCRuleSwap102:
		err = commandAndRuleEngineObject.executeTCRuleSwap102(elementUuid)

	case TCRuleSwap103:
		err = commandAndRuleEngineObject.executeTCRuleSwap103(elementUuid)

	case TCRuleSwap104:
		err = commandAndRuleEngineObject.executeTCRuleSwap104(elementUuid)

	case TCRuleSwap105:
		commandAndRuleEngineObject.executeTCRuleSwap105(elementUuid)

	case TCRuleSwap106:
		err = commandAndRuleEngineObject.executeTCRuleSwap106(elementUuid)

	case TCRuleSwap107:
		err = commandAndRuleEngineObject.executeTCRuleSwap107(elementUuid)

	case TCRuleSwap108:
		err = commandAndRuleEngineObject.executeTCRuleSwap108(elementUuid)

	case TCRuleSwap109:
		err = commandAndRuleEngineObject.executeTCRuleSwap109(elementUuid)

	case TCRuleSwap110:
		err = commandAndRuleEngineObject.executeTCRuleSwap110(elementUuid)

	case TCRuleSwap111:
		err = commandAndRuleEngineObject.executeTCRuleSwap111(elementUuid)

	case TCRuleSwap112:
		err = commandAndRuleEngineObject.executeTCRuleSwap112(elementUuid)

	case TCRuleSwap113:
		err = commandAndRuleEngineObject.executeTCRuleSwap113(elementUuid)

	case TCRuleSwap114:
		err = commandAndRuleEngineObject.executeTCRuleSwap114(elementUuid)

	case TCRuleSwap115:
		err = commandAndRuleEngineObject.executeTCRuleSwap115(elementUuid)

	case TCRuleSwap116:
		err = commandAndRuleEngineObject.executeTCRuleSwap116(elementUuid)

	case TCRuleSwap117:
		err = commandAndRuleEngineObject.executeTCRuleSwap117(elementUuid)

	default:
		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":                 "eba1a213-aa42-4021-aaea-4b3107d5874c",
			"matchedComplexRule": matchedComplexRule,
		}).Error(" Unknown 'matchedComplexRule' was used when trying to swap")

		err = errors.New("'" + matchedComplexRule + "' is an unknown complex Swap rule")

	}

	return err

}
