package detailedExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/executions/detailedTestCaseExecutionUI_summaryTableDefinition"
	"github.com/sirupsen/logrus"
)

// Removes a TestCaseExecution from both summary page and the Details page for the TestCaseExecution
func (detailedExecutionsModelObject *DetailedExecutionsModelObjectStruct) processRemoveDetailedTestCaseExecution(
	testCaseExecutionKey string) (err error) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id":                   "a766d9cc-9c57-47b6-be98-d384d420c16c",
		"testCaseExecutionKey": testCaseExecutionKey,
	}).Debug("Incoming 'processRemoveDetailedTestCaseExecution'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "22c98071-8d34-4a1e-b292-7a8e926a3bb6",
	}).Debug("Outgoing 'processRemoveDetailedTestCaseExecution'")

	// Remove the TestCaseExecution from the Map
	delete(detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap, testCaseExecutionKey)

	return err
}
