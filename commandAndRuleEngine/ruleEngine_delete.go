package commandAndRuleEngine

import (
	"errors"
	"github.com/sirupsen/logrus"
)

// Verify if an element can be deleted or not, regarding deletion rules
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) verifyIfElementCanBeDeleted(elementUuid string) (canBeDeleted bool, matchedSimpldRule string, matchedComplexRule string, err error) {

	// First verify towards simple rules
	canBeDeleted, matchedSimpldRule, err = commandAndRuleEngine.verifyIfComponentCanBeDeletedSimpleRules(elementUuid)

	// Only check complex rules if simple rules was OK for deletion
	if !(canBeDeleted == true &&
		err != nil) {
		return canBeDeleted, matchedSimpldRule, "", err
	}

	// Verify towards complex rules
	matchedComplexRule, err = commandAndRuleEngine.verifyIfComponentCanBeDeletedWithComplexRules(elementUuid)

	return canBeDeleted, matchedSimpldRule, matchedComplexRule, err
}

// Delete an element, but first ensure that rules for deletion are used
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeDeleteElement(elementUuid string) (err error) {

	// Verify that element is allowed, and can be deleted
	canBeDeleted, matcheSimpleRule, matchedComplexRule, err := commandAndRuleEngine.verifyIfElementCanBeDeleted(elementUuid)

	// If there was an error from delete verification then exit
	if err != nil {
		return err
	}

	// If the component couldn't be deleted then exit with error message
	if canBeDeleted == false {
		err = errors.New("element couldn't be deleted due to deletion rule '" + matcheSimpleRule + "' or that complex rules aren't met")

		return err
	}

	// Execute deletion of element based on rule-number
	err = commandAndRuleEngine.executeDeleteElementBasedOnRule(elementUuid, matchedComplexRule)

	return err
}

// Delete an element based on specific rule
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeDeleteElementBasedOnRule(elementUuid string, matchedComplexRule string) (err error) {

	switch matchedComplexRule {
	case TCRuleDeletion101:
		err = commandAndRuleEngine.executeTCRuleDeletion101(elementUuid)

	case TCRuleDeletion102:
		err = commandAndRuleEngine.executeTCRuleDeletion102(elementUuid)

	case TCRuleDeletion103:
		err = commandAndRuleEngine.executeTCRuleDeletion103(elementUuid)

	case TCRuleDeletion104:
		err = commandAndRuleEngine.executeTCRuleDeletion104(elementUuid)

	case TCRuleDeletion105:
		err = commandAndRuleEngine.executeTCRuleDeletion105(elementUuid)

	case TCRuleDeletion106:
		err = commandAndRuleEngine.executeTCRuleDeletion106(elementUuid)

	case TCRuleDeletion107:
		err = commandAndRuleEngine.executeTCRuleDeletion107(elementUuid)

	case TCRuleDeletion108:
		err = commandAndRuleEngine.executeTCRuleDeletion108(elementUuid)

	case TCRuleDeletion109:
		err = commandAndRuleEngine.executeTCRuleDeletion109(elementUuid)

	case TCRuleDeletion110:
		err = commandAndRuleEngine.executeTCRuleDeletion110(elementUuid)

	case TCRuleDeletion111:
		err = commandAndRuleEngine.executeTCRuleDeletion111(elementUuid)

	case TCRuleDeletion112:
		err = commandAndRuleEngine.executeTCRuleDeletion112(elementUuid)

	case TCRuleDeletion113:
		err = commandAndRuleEngine.executeTCRuleDeletion113(elementUuid)

	case TCRuleDeletion114:
		err = commandAndRuleEngine.executeTCRuleDeletion114(elementUuid)

	case TCRuleDeletion115:
		err = commandAndRuleEngine.executeTCRuleDeletion115(elementUuid)

	case TCRuleDeletion116:
		err = commandAndRuleEngine.executeTCRuleDeletion116(elementUuid)

	case TCRuleDeletion117:
		err = commandAndRuleEngine.executeTCRuleDeletion117(elementUuid)

	default:
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":                 "928e8983-7477-426e-89a8-73b04846c251",
			"matchedComplexRule": matchedComplexRule,
		}).Error(" Unknown 'matchedComplexRule' was used when trying to delete")

		err = errors.New("'" + matchedComplexRule + "' is an unknown complex deletion rule")

	}

	return err

}
