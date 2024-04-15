package detailedExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/executions/detailedTestCaseExecutionUI_summaryTableDefinition"
	"github.com/sirupsen/logrus"
	"time"
)

func InitiateCommandChannelReaderForDetailedStatusUpdates() {

	// Initiate channel for Detailed Executions Status updates
	DetailedExecutionStatusCommandChannel = make(chan ChannelCommandDetailedExecutionsStruct, messageChannelMaxSizeDetailedExecutionStatus)

	// Initiate the map that stores all Detailed TestCaseExecutions-data
	detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap =
		make(map[string]*detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsStruct)

	// Start channel reader
	go detailedExecutionsModelObject.startCommandChannelReaderForDetailedStatusUpdates()

	// Initiate 'testCasesSummaryTableRefreshThrottler', used for only Refresh (repaint) the TestCasesSummaryTable at
	// regular interval. When many updates are made to it
	testCasesSummaryTableRefreshThrottler = newRefreshTestCasesSummaryTableThrottler(5 * time.Second)
	//defer throttler.Stop()

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
			// Triggered from 'RetrieveSingleTestCaseExecution(testCaseExecutionKey string)'
			detailedExecutionsModelObject.triggerProcessFullDetailedExecutionsStatusUpdate(incomingChannelCommandAndMessage)

		case ChannelCommandStatusUpdateOfDetailedExecutionsStatus:
			// Triggered from 'processFullDetailedTestCaseExecutionsStatusUpdate' -Updates all Executions status with information received after direct gRPC-call to GUiExecutionServer
			// Can have External call from (executionStatusUpdatesHandler_channelReader.go) 'Channel reader which is used for reading out Status messages that is sent from GuiExecutionServer'
			detailedExecutionsModelObject.triggerProcessStatusUpdateOfDetailedExecutionsStatus(incomingChannelCommandAndMessage)

		case ChannelCommandRemoveDetailedTestCaseExecution:
			// Triggered from 'RemoveTestCaseExecutionFromSummaryTable' - Removes a TestCaseExecution from both summary page and the Details page for the TestCaseExecution
			detailedExecutionsModelObject.triggerProcessRemoveDetailedTestCaseExecution(incomingChannelCommandAndMessage.TestCaseExecutionKey)

		case ChannelCommandRetrieveFullDetailedTestCaseExecution:
			// Triggered from 'processStatusUpdateOfDetailedExecutionsStatus' - Updates specific status information based on subscriptions updates from GuiExecutionServer
			detailedExecutionsModelObject.triggerProcessRetrieveFullDetailedTestCaseExecution(incomingChannelCommandAndMessage.TestCaseExecutionKey)

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

	sharedCode.Logger.WithFields(logrus.Fields{
		"Id": "5d167ac1-e44d-4300-833a-23651f92656a",
		"incomingChannelCommandAndMessage.TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage": incomingChannelCommandAndMessage.TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage,
	}).Debug("Incoming ProcessFullDetailedExecutionsStatusUpdate - 'ChannelCommandFullDetailedExecutionsStatusUpdate'")

	detailedExecutionsModelObject.processFullDetailedTestCaseExecutionsStatusUpdate(
		incomingChannelCommandAndMessage.FullTestCaseExecutionResponseMessage)

	// Recreate the Detailed Executions Summary Table
	*TestCasesSummaryTable = *CreateSummaryTableForDetailedTestCaseExecutionsList()

	// Only refresh (repaint) the table at a maximum refresh-rate
	testCasesSummaryTableRefreshThrottler.RequestRefreshTestCasesSummaryTable()

}

// Updates specific status information based on subscriptions updates from GuiExecutionServer
func (detailedExecutionsModelObject *DetailedExecutionsModelObjectStruct) triggerProcessStatusUpdateOfDetailedExecutionsStatus(
	incomingChannelCommandAndMessage ChannelCommandDetailedExecutionsStruct) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"Id": "5b27fadb-a9cd-46d5-8547-9c802352e2cd",
		"incomingChannelCommandAndMessage.TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage": incomingChannelCommandAndMessage.TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage,
	}).Debug("Incoming ProcessStatusUpdateOfDetailedExecutionsStatus - 'ChannelCommandStatusUpdateOfDetailedExecutionsStatus'")

	detailedExecutionsModelObject.processStatusUpdateOfDetailedExecutionsStatus(
		incomingChannelCommandAndMessage.TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage)

	// Recreate the Detailed Executions Summary Table
	*TestCasesSummaryTable = *CreateSummaryTableForDetailedTestCaseExecutionsList()

	// Only refresh (repaint) the table at a maximum refresh-rate
	testCasesSummaryTableRefreshThrottler.RequestRefreshTestCasesSummaryTable()

}

// Remove the DetailedTestCaseExecution
func (detailedExecutionsModelObject *DetailedExecutionsModelObjectStruct) triggerProcessRemoveDetailedTestCaseExecution(
	testCaseExecutionKey string) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"Id":                   "db78f927-07eb-419e-97d5-f4d5aed0d6ee",
		"testCaseExecutionKey": testCaseExecutionKey,
	}).Debug("Incoming RemoveDetailedTestCaseExecution - 'ChannelCommandRemoveDetailedTestCaseExecution'")

	err := detailedExecutionsModelObject.processRemoveDetailedTestCaseExecution(testCaseExecutionKey)

	if err != nil {
		return
	}

	// Recreate the Detailed Executions Summary Table
	*TestCasesSummaryTable = *CreateSummaryTableForDetailedTestCaseExecutionsList()

	// Only refresh (repaint) the table at a maximum refresh-rate
	testCasesSummaryTableRefreshThrottler.RequestRefreshTestCasesSummaryTable()

}

// Retrieve a full Detailed TestCaseExecution from GuiExecutionServer
func (detailedExecutionsModelObject *DetailedExecutionsModelObjectStruct) triggerProcessRetrieveFullDetailedTestCaseExecution(
	testCaseExecutionKey string) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"Id":                   "00d37e9d-cac5-4989-a5b8-97bcfbeb7275",
		"testCaseExecutionKey": testCaseExecutionKey,
	}).Debug("Incoming RetrieveFullDetailedTestCaseExecution - 'ChannelCommandRetrieveFullDetailedTestCaseExecution'")

	err := detailedExecutionsModelObject.processRetrieveFullDetailedTestCaseExecution(testCaseExecutionKey)

	if err != nil {
		return
	}

	// Recreate the Detailed Executions Summary Table
	*TestCasesSummaryTable = *CreateSummaryTableForDetailedTestCaseExecutionsList()

	// Only refresh (repaint) the table at a maximum refresh-rate
	testCasesSummaryTableRefreshThrottler.RequestRefreshTestCasesSummaryTable()

}
