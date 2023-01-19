package detailedExecutionsModel

// Removes a TestCaseExecution from both summary page and the Details page for the TestCaseExecution
func (detailedExecutionsModelObject *DetailedExecutionsModelObjectStruct) processRetrieveFullDetailedTestCaseExecution(
	testCaseExecutionKey string) (err error) {

	// Do gRPC-call GuiExecutionServer to get a FullTestCaseExecution from the Database
	err = RetrieveSingleTestCaseExecution(testCaseExecutionKey)

	return err
}
