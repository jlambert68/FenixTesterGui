package detailedExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	"github.com/sirupsen/logrus"
)

// RemoveTestCaseExecutionFromSummaryTable
// Removes a TestCaseExecution from both summary page and the Details page for the TestCaseExecution
func RemoveTestCaseExecutionFromSummaryTable(testCaseExecutionKey string) (err error) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id":                   "e8d677a8-e925-4494-8bd4-9600d95fd012",
		"testCaseExecutionKey": testCaseExecutionKey,
	}).Debug("Incoming 'RemoveTestCaseExecutionFromSummaryTable'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "961210a5-4fb5-487e-a0a5-1637ec1efcce",
	}).Debug("Outgoing 'RemoveTestCaseExecutionFromSummaryTable'")

	// Remove the TestCaseExecution-details via channelEngine
	var channelCommandDetailedExecutions ChannelCommandDetailedExecutionsStruct
	channelCommandDetailedExecutions = ChannelCommandDetailedExecutionsStruct{
		ChannelCommandDetailedExecutionsStatus:                            ChannelCommandRemoveDetailedTestCaseExecution,
		TestCaseExecutionKey:                                              testCaseExecutionKey,
		FullTestCaseExecutionResponseMessage:                              nil,
		TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage: nil,
	}

	// Send command ion channel
	DetailedExecutionStatusCommandChannel <- channelCommandDetailedExecutions

	return nil
}
