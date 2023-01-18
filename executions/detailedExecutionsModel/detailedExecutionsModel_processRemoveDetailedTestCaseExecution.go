package detailedExecutionsModel

import "FenixTesterGui/executions/detailedTestCaseExecutionUI_summaryTableDefinition"

// Removes a TestCaseExecution from both summary page and the Details page for the TestCaseExecution
func (detailedExecutionsModelObject *DetailedExecutionsModelObjectStruct) processRemoveDetailedTestCaseExecution(testCaseExecutionKey string) (err error) {

	// Remove the TestCaseExecution from the Map
	delete(detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap, testCaseExecutionKey)

	return err
}
