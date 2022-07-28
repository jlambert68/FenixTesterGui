package commandAndRuleEngine

import (
	"errors"
	"fmt"
)

// Verify if an element can be copied or not, regarding copy rules
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) verifyIfElementCanBeCopied(testCaseUuid string, elementUuid string) (canBeCopied bool, matchedSimpldRule string, err error) {

	// First verify towards simple rules
	canBeCopied, matchedSimpldRule, err = commandAndRuleEngine.verifyIfComponentCanBeCopiedSimpleRules(testCaseUuid, elementUuid)

	return canBeCopied, matchedSimpldRule, err
}

// Copy an element, but first ensure that rules for copying are used
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeCopyElement(testCaseUuid string, elementUuid string) (err error) {

	// Verify that element is allowed, and can be copied
	canBeCopyied, matchedSimpleRule, err := commandAndRuleEngine.verifyIfElementCanBeCopied(testCaseUuid, elementUuid)

	// If there was an error from copy verification then exit
	if err != nil {
		return err
	}

	// If the component couldn't be copied then exit with error message
	if canBeCopyied == false {
		errorId := "e016b80b-572d-475f-b887-84044b1a0227"
		err = errors.New(fmt.Sprintf("element couldn't be copied due to copy rule '%s' is not met [ErrorID: %s]", matchedSimpleRule, errorId))

		return err
	}

	// Execute copy element
	err = commandAndRuleEngine.executeCopyFullELementStructure(testCaseUuid, elementUuid)

	return err
}
