package commandAndRuleEngine

//

// NewTestCaseModel
// Used, mostly from GUI, to for creating a new TestCase-Model to be used within a new TestCase
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) NewTestCaseModel() (err error) {

	_, err = commandAndRuleEngine.executeCommandOnTestCaseModel_NewTestCaseModel()

	return err

}

// NewTestCaseModel
// Used, mostly from GUI, for Deleting an element from a TestCaseModel that is used within a TestCase
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) DeleteElementFromTestCaseModel(testCaseUuid string, elementId string) (err error) {

	err = commandAndRuleEngine.executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel(testCaseUuid, elementId)

	return err

}
