package commandAndRuleEngine

import "errors"

// Verify if an element can be deleted or not, regarding deletion rules
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) verifyIfElementCanBeDeleted(elementUuid string) (canBeDeleted bool, matcheSimpldRule string, matchedComplexRule string, err error) {

	// First verify towards simple rules
	canBeDeleted, matcheSimpldRule, err = commandAndRuleEngineObject.verifyIfComponentCanBeDeletedSimpleRules(elementUuid)

	// Only check complex rules if simple rules was OK for deletion
	if !(canBeDeleted == true &&
		err != nil) {
		return canBeDeleted, matcheSimpldRule, "", err
	}

	// Verify towards complex rules
	matchedComplexRule, err = commandAndRuleEngineObject.verifyIfComponentCanBeDeletedWithComplexRules(elementUuid)

	return canBeDeleted, matcheSimpldRule, matchedComplexRule, err
}

// Delete an element, but first ensure that rules for deletion are used
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) executeDeleteElement(elementUuid string) (err error) {

	// Verify that element is allowed, and can be deleted
	canBeDeleted, matcheSimpldRule, _, err := commandAndRuleEngineObject.verifyIfElementCanBeDeleted(elementUuid)

	// If there was an error from delete verification then exit
	if err != nil {
		return err
	}

	// If the component couldn't be deleted then exit with error message
	if canBeDeleted == false {
		err = errors.New("Element couldn't be deleted due to deletion rule '" + matcheSimpldRule + "'")

		return err
	}

	// Execute deletion of element
	err = commandAndRuleEngineObject.executeDeleteElement(elementUuid)

	return err
}
