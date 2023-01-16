package detailedExecutionsModel

// RemoveTestCaseExecutionFromSummaryTable
// Removes a TestCaseExecution from both summary page and the Details page for the TestCaseExecution
func RemoveTestCaseExecutionFromSummaryTable(testCaseExecutionKey string) (err error) {

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
