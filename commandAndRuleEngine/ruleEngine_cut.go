package commandAndRuleEngine

import (
	"errors"
	"fmt"
)

// Verify if an element can be cut out or not, regarding cut rules
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) verifyIfElementCanBeCutOut(testCaseUuid string, elementUuid string) (canBeCut bool, matchedSimpldRule string, err error) {

	// First verify towards simple rules
	canBeCut, matchedSimpldRule, err = commandAndRuleEngine.verifyIfComponentCanBeCutSimpleRules(testCaseUuid, elementUuid)

	return canBeCut, matchedSimpldRule, err
}

// Cut an element, but first ensure that rules for cutting are used
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) executeCutElement(testCaseUuid string, elementUuid string) (err error) {

	// Verify that element is allowed, and can be cut
	canBeCut, matchedSimpleRule, err := commandAndRuleEngine.verifyIfElementCanBeCutOut(testCaseUuid, elementUuid)

	// If there was an error from cut verification then exit
	if err != nil {
		return err
	}

	// If the component couldn't be cut then exit with error message
	if canBeCut == false {
		errorId := "1b1ef073-ba1d-4614-939f-0f685cb1c489"
		err = errors.New(fmt.Sprintf("element couldn't be cut due to cut rule '%s' is not met [ErrorID: %s]", matchedSimpleRule, errorId))

		return err
	}

	// Execute cut element
	err = commandAndRuleEngine.executeCutFullELementStructure(testCaseUuid, elementUuid)

	return err
}
