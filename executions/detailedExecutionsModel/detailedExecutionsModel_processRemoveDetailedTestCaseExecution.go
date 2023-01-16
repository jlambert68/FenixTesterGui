package detailedExecutionsModel

// Removes a TestCaseExecution from both summary page and the Details page for the TestCaseExecution
func (detailedExecutionsModelObject *DetailedExecutionsModelObjectStruct) processRemoveDetailedTestCaseExecution(testCaseExecutionKey string) (err error) {

	// Remove the TestCaseExecution from the Map
	delete(TestCaseExecutionsDetailsMap, testCaseExecutionKey)

	return err
}
