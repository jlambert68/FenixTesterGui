package commandAndRuleEngine

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
