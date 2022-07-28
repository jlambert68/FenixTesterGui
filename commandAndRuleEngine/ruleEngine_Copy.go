package commandAndRuleEngine

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
)

// Verify if an element can be copied or not, regarding copy rules
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) verifyIfElementCanBeCopied(testCaseUuid string, elementUuid string) (canBeDeleted bool, matchedSimpldRule string, err error) {

	// First verify towards simple rules
	canBeDeleted, matchedSimpldRule, err = commandAndRuleEngine.verifyIfComponentCanBeCopiedSimpleRules(testCaseUuid, elementUuid)

	return canBeDeleted, matchedSimpldRule, err
}

// Copy an element, but first ensure that rules for copying are used
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeCopyElement(testCaseUuid string, elementUuid string) (err error) {

	// Verify that element is allowed, and can be deleted
	canBeDeleted, matchedSimpleRule, err := commandAndRuleEngine.verifyIfElementCanBeCopied(testCaseUuid, elementUuid)

	// If there was an error from delete verification then exit
	if err != nil {
		return err
	}

	// If the component couldn't be copied then exit with error message
	if canBeDeleted == false {
		errorId := "e016b80b-572d-475f-b887-84044b1a0227"
		err = errors.New(fmt.Sprintf("element couldn't be copied due to copy rule '%s' is not met [ErrorID: %s]", matchedSimpleRule, errorId))

		return err
	}

	// Execute deletion of element based on rule-number
	err = commandAndRuleEngine.executeCopyElementBasedOnRule(testCaseUuid, elementUuid)

	return err
}

// Delete an element based on specific rule
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeCopyElementBasedOnRule(testCaseUuid string, elementUuid string) (err error) {

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
			"id":                 "128f51cd-2e0b-4ecc-93e3-de75eab90c5d",
			"matchedComplexRule": matchedComplexRule,
		}).Error(" Unknown 'matchedComplexRule' was used when trying to delete")

		errorId := "b106fc72-50de-40b7-bacc-ae99cf7ca725"
		err = errors.New(fmt.Sprintf("'%s' is an unknown complex deletion rule [ErrorID: %s]", matchedComplexRule, errorId))

	}

	return err

}
