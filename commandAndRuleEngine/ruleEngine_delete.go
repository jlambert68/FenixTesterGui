package commandAndRuleEngine

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
)

// Verify if an element can be deleted or not, regarding deletion rules
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) verifyIfElementCanBeDeleted(testCaseUuid string, elementUuid string) (canBeDeleted bool, matchedSimpldRule string, matchedComplexRule string, err error) {

	// First verify towards simple rules
	canBeDeleted, matchedSimpldRule, err = commandAndRuleEngine.verifyIfComponentCanBeDeletedSimpleRules(testCaseUuid, elementUuid)

	// Only check complex rules if simple rules was OK for deletion
	if !(canBeDeleted == true &&
		err == nil) {
		return canBeDeleted, matchedSimpldRule, "", err
	}

	// Verify towards complex rules
	matchedComplexRule, err = commandAndRuleEngine.verifyIfComponentCanBeDeletedWithComplexRules(testCaseUuid, elementUuid)

	return canBeDeleted, matchedSimpldRule, matchedComplexRule, err
}

// Delete an element, but first ensure that rules for deletion are used
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeDeleteElement(testCaseUuid string, elementUuid string) (err error) {

	// Verify that element is allowed, and can be deleted
	canBeDeleted, matcheSimpleRule, matchedComplexRule, err := commandAndRuleEngine.verifyIfElementCanBeDeleted(testCaseUuid, elementUuid)

	// If there was an error from delete verification then exit
	if err != nil {
		return err
	}

	// If the component couldn't be deleted then exit with error message
	if canBeDeleted == false {
		errorId := "6319ec76-d471-46dc-841b-16c488c6e728"
		err = errors.New(fmt.Sprintf("element couldn't be deleted due to deletion rule '%s' or that complex rules aren't met [ErrorID: %s]", matcheSimpleRule, errorId))

		return err
	}

	// Execute deletion of element based on rule-number
	err = commandAndRuleEngine.executeDeleteElementBasedOnRule(testCaseUuid, elementUuid, matchedComplexRule)

	return err
}

// Delete an element based on specific rule
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeDeleteElementBasedOnRule(testCaseUuid string, elementUuid string, matchedComplexRule string) (err error) {

	switch matchedComplexRule {
	case TCRuleDeletion101:
		err = commandAndRuleEngine.executeTCRuleDeletion101(testCaseUuid, elementUuid)

	case TCRuleDeletion102:
		err = commandAndRuleEngine.executeTCRuleDeletion102(testCaseUuid, elementUuid)

	case TCRuleDeletion103:
		err = commandAndRuleEngine.executeTCRuleDeletion103(testCaseUuid, elementUuid)

	case TCRuleDeletion104:
		err = commandAndRuleEngine.executeTCRuleDeletion104(testCaseUuid, elementUuid)

	case TCRuleDeletion105:
		err = commandAndRuleEngine.executeTCRuleDeletion105(testCaseUuid, elementUuid)

	case TCRuleDeletion106:
		err = commandAndRuleEngine.executeTCRuleDeletion106(testCaseUuid, elementUuid)

	case TCRuleDeletion107:
		err = commandAndRuleEngine.executeTCRuleDeletion107(testCaseUuid, elementUuid)

	case TCRuleDeletion108:
		err = commandAndRuleEngine.executeTCRuleDeletion108(testCaseUuid, elementUuid)

	case TCRuleDeletion109:
		err = commandAndRuleEngine.executeTCRuleDeletion109(testCaseUuid, elementUuid)

	case TCRuleDeletion110:
		err = commandAndRuleEngine.executeTCRuleDeletion110(testCaseUuid, elementUuid)

	case TCRuleDeletion111:
		err = commandAndRuleEngine.executeTCRuleDeletion111(testCaseUuid, elementUuid)

	case TCRuleDeletion112:
		err = commandAndRuleEngine.executeTCRuleDeletion112(testCaseUuid, elementUuid)

	case TCRuleDeletion113:
		err = commandAndRuleEngine.executeTCRuleDeletion113(testCaseUuid, elementUuid)

	case TCRuleDeletion114:
		err = commandAndRuleEngine.executeTCRuleDeletion114(testCaseUuid, elementUuid)

	case TCRuleDeletion115:
		err = commandAndRuleEngine.executeTCRuleDeletion115(testCaseUuid, elementUuid)

	case TCRuleDeletion116:
		err = commandAndRuleEngine.executeTCRuleDeletion116(testCaseUuid, elementUuid)

	case TCRuleDeletion117:
		err = commandAndRuleEngine.executeTCRuleDeletion117(testCaseUuid, elementUuid)

	default:
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":                 "928e8983-7477-426e-89a8-73b04846c251",
			"matchedComplexRule": matchedComplexRule,
		}).Error(" Unknown 'matchedComplexRule' was used when trying to delete")

		errorId := "b106fc72-50de-40b7-bacc-ae99cf7ca725"
		err = errors.New(fmt.Sprintf("'%s' is an unknown complex deletion rule [ErrorID: %s]", matchedComplexRule, errorId))

	}

	return err

}
