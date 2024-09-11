package detailedExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	"github.com/sirupsen/logrus"
	"time"
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

	// Don't put on Channel if more than 9 items from max capacity
	var currentChannelSize int32
	currentChannelSize = int32(len(DetailedExecutionStatusCommandChannel))
	if currentChannelSize > MessageChannelMaxSizeDetailedExecutionStatus-9 {
		for {
			time.Sleep(5 * time.Second)

			currentChannelSize = int32(len(DetailedExecutionStatusCommandChannel))
			if currentChannelSize < MessageChannelMaxSizeDetailedExecutionStatus-9 {
				break
			}
		}
	}

	// Send command ion channel
	DetailedExecutionStatusCommandChannel <- channelCommandDetailedExecutions

	return nil
}
