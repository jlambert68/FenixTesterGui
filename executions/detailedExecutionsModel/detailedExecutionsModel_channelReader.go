package detailedExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/executions/detailedTestCaseExecutionUI_summaryTableDefinition"
	"github.com/sirupsen/logrus"
)

func InitiateCommandChannelReaderForDetailedStatusUpdates() {

	// Initiate channel for Detailed Executions Status updates
	DetailedExecutionStatusCommandChannel = make(chan ChannelCommandDetailedExecutionsStruct, messageChannelMaxSizeDetailedExecutionStatus)

	// Initiate the map that stores all Detailed TestCaseExecutions-data
	detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap =
		make(map[string]*detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsStruct)

	// Start channel reader
	go detailedExecutionsModelObject.startCommandChannelReaderForDetailedStatusUpdates()

}

// Channel reader which is used secure that status updates are handled in a structured way
func (detailedExecutionsModelObject *DetailedExecutionsModelObjectStruct) startCommandChannelReaderForDetailedStatusUpdates() {

	var incomingChannelCommandAndMessage ChannelCommandDetailedExecutionsStruct
	var currentChannelSize int32

	for {
		// Wait for incoming command over channel
		incomingChannelCommandAndMessage = <-DetailedExecutionStatusCommandChannel

		// If channel is near its maximum size then create error message
		currentChannelSize = int32(len(DetailedExecutionStatusCommandChannel))
		if currentChannelSize > messageChannelMaxSizeDetailedExecutionStatus-10 {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":                 "8986fa1f-1d63-4441-b416-91955b9a351a",
				"currentChannelSize": currentChannelSize,
				"messageChannelMaxSizeDetailedExecutionStatus": messageChannelMaxSizeDetailedExecutionStatus,
			}).Warning("Number of messages on 'DetailedExecutionStatusCommandChannel' is close to its maximum")
		}

		// Act correctly depending on what message that was received
		switch incomingChannelCommandAndMessage.ChannelCommandDetailedExecutionsStatus {

		case ChannelCommandFullDetailedExecutionsStatusUpdate:
			detailedExecutionsModelObject.triggerProcessFullDetailedExecutionsStatusUpdate(incomingChannelCommandAndMessage)

		case ChannelCommandStatusUpdateOfDetailedExecutionsStatus:
			detailedExecutionsModelObject.triggerProcessStatusUpdateOfDetailedExecutionsStatus(incomingChannelCommandAndMessage)

		case ChannelCommandRemoveDetailedTestCaseExecution:
			detailedExecutionsModelObject.triggerProcessRemoveDetailedTestCaseExecution(incomingChannelCommandAndMessage.TestCaseExecutionKey)

		// No other command is supported
		default:
			sharedCode.Logger.WithFields(logrus.Fields{
				"Id":                               "76caff3b-6cb2-4907-bf19-88bf51122d6a",
				"incomingChannelCommandAndMessage": incomingChannelCommandAndMessage,
			}).Fatalln("Unknown command in CommandChannel for MessageStreamEngine")
		}
	}

}

// Updates all Executions status with information received after direct gRPC-call to GUiExecutionServer
func (detailedExecutionsModelObject *DetailedExecutionsModelObjectStruct) triggerProcessFullDetailedExecutionsStatusUpdate(
	incomingChannelCommandAndMessage ChannelCommandDetailedExecutionsStruct) {

	detailedExecutionsModelObject.processFullDetailedTestCaseExecutionsStatusUpdate(
		incomingChannelCommandAndMessage.FullTestCaseExecutionResponseMessage)

	// Recreate the Detailed Executions Summary Table
	*TestCasesSummaryTable = *CreateSummaryTableForDetailedTestCaseExecutionsList()

	TestCasesSummaryTable.Refresh()

}

// Updates specific status information based on subscriptions updates from GuiExecutionServer
func (detailedExecutionsModelObject *DetailedExecutionsModelObjectStruct) triggerProcessStatusUpdateOfDetailedExecutionsStatus(
	incomingChannelCommandAndMessage ChannelCommandDetailedExecutionsStruct) {

	detailedExecutionsModelObject.processStatusUpdateOfDetailedExecutionsStatus(
		incomingChannelCommandAndMessage.TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage)

	// Recreate the Detailed Executions Summary Table
	*TestCasesSummaryTable = *CreateSummaryTableForDetailedTestCaseExecutionsList()

	TestCasesSummaryTable.Refresh()

}

// Remove the DetailedTestCaseExecution
func (detailedExecutionsModelObject *DetailedExecutionsModelObjectStruct) triggerProcessRemoveDetailedTestCaseExecution(
	testCaseExecutionKey string) {

	detailedExecutionsModelObject.processRemoveDetailedTestCaseExecution(testCaseExecutionKey)

	// Recreate the Detailed Executions Summary Table
	*TestCasesSummaryTable = *CreateSummaryTableForDetailedTestCaseExecutionsList()

	TestCasesSummaryTable.Refresh()

}
