package commandAndRuleEngine

//

// NewTestCaseModel
// Used, mostly from GUI, to for creating a new TestCase-Model to be used within a new TestCase
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) NewTestCaseModel() (err error) {

	_, err = commandAndRuleEngine.executeCommandOnTestCaseModel_NewTestCaseModel()

	return err

}
