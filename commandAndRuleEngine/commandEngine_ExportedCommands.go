package commandAndRuleEngine

import "FenixTesterGui/testCase/testCaseModel"

//

// NewTestCaseModel
// Used, mostly from GUI, to for creating a new TestCase-Model to be used within a new TestCase
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) NewTestCaseModel() (err error) {

	_, err = commandAndRuleEngine.executeCommandOnTestCaseModel_NewTestCaseModel()

	return err

}

// DeleteElementFromTestCaseModel
// Used, mostly from GUI, for Deleting an element from a TestCaseModel that is used within a TestCase
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) DeleteElementFromTestCaseModel(testCaseUuid string, elementId string) (err error) {

	err = commandAndRuleEngine.executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel(testCaseUuid, elementId)

	return err

}

// SwapElementsInTestCaseModel
// Used, mostly from GUI, for Swapping out an element, and in an element structure, from a TestCaseModel that is used within a TestCase
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) SwapElementsInTestCaseModel(testcaseUuid string, elementUuidTobeSwappedOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (err error) {

	err = commandAndRuleEngine.executeCommandOnTestCaseModel_SwapOutElemenAndInNewElementInTestCaseModel(testcaseUuid, elementUuidTobeSwappedOut, immatureElementToSwapIn)

	return err

}
