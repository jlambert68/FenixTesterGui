package messageStreamEngine

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/executions/executionsModel"
	"FenixTesterGui/executions/executionsUI"
	"errors"
	"fmt"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

//Channel reader which is used for reading out Status messages that is sent from GuiExecutionServer
func (messageStreamEngineObject *MessageStreamEngineStruct) startCommandChannelReader() {

	var incomingChannelCommandAndMessage ChannelCommandStruct
	var currentChannelSize int32

	for {
		// Wait for incoming command over channel
		incomingChannelCommandAndMessage = <-executionStatusCommandChannel

		// If channel is near its maximum size then create error message
		currentChannelSize = int32(len(executionStatusCommandChannel))
		if currentChannelSize > messageChannelMaxSize-10 {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":                    "b4164d7a-a485-411d-ad18-feb50ed98566",
				"currentChannelSize":    currentChannelSize,
				"messageChannelMaxSize": messageChannelMaxSize,
			}).Error("Number of messages on 'executionStatusCommandChannel' is close to its maximum")
		}

		// Act correctly depending of what message that was received
		switch incomingChannelCommandAndMessage.ChannelCommand {

		case ChannelCommandExecutionsStatusesHaveBeUpdated:
			// TestCaseExecutionStatus or TestInstructionExecutionStatus has been updated
			fmt.Println(incomingChannelCommandAndMessage)
			messageStreamEngineObject.processTestExecutionStatusChange(incomingChannelCommandAndMessage.ExecutionsStatusMessage)

		case ChannelCommandTriggerRequestForTestInstructionExecutionToProcess:
			messageStreamEngineObject.initiateOpenMessageStreamToGuiExecutionServer()

		case ChannelCommandTriggerRequestForTestInstructionExecutionToProcessIn1Second:
			messageStreamEngineObject.initiateOpenMessageStreamToGuiExecutionServerInXSeconds(1 * 1)

		// No other command is supported
		default:
			sharedCode.Logger.WithFields(logrus.Fields{
				"Id":                               "6bf37452-da99-4e7e-aa6a-4627b05d1bdb",
				"incomingChannelCommandAndMessage": incomingChannelCommandAndMessage,
			}).Fatalln("Unknown command in CommandChannel for MessageStreamEngine")
		}
	}

}

// Call Worker to get TestInstructions to Execute, which is done as a message stream in the response from the Worker
func (messageStreamEngineObject *MessageStreamEngineStruct) initiateOpenMessageStreamToGuiExecutionServer() {

	// Call RequestForProcessTestInstructionExecution with parameter set to zero sleep before do the gPRC-call
	messageStreamEngineObject.initiateOpenMessageStreamToGuiExecutionServerInXSeconds(0)

}

// Call Worker in X seconds, due to some connection error, to get TestInstructions to Execute, which is done as a message stream in the response from the Worker
func (messageStreamEngineObject *MessageStreamEngineStruct) initiateOpenMessageStreamToGuiExecutionServerInXSeconds(waitTimeInSeconds uint) {

	// Only trigger wait time if there are no ongoing message connection to GuiExecutionServer
	if messageStreamEngineObject.ongoingTimerOrConnectionForCallingWorkerForTestInstructionsToExecute == true {
		return
	}

	// Run it as a go-routine
	go func() {

		// Set that there is an ongoing timer
		messageStreamEngineObject.ongoingTimerOrConnectionForCallingWorkerForTestInstructionsToExecute = true

		// Wait x minutes/second before triggering
		sleepDuration := time.Duration(waitTimeInSeconds) * time.Second
		if waitTimeInSeconds != 0 {
			time.Sleep(sleepDuration)
		}
		// Call GuiExecutionServer to open up for stream of messages
		messageStreamEngineObject.initiateGuiExecutionServerRequestForMessages()

		// Connection broke down
		messageStreamEngineObject.ongoingTimerOrConnectionForCallingWorkerForTestInstructionsToExecute = false

		// Create Message for CommandChannel to retry to connect in 1 second
		var channelCommandAndMessage ChannelCommandStruct
		channelCommandAndMessage = ChannelCommandStruct{
			ChannelCommand: ChannelCommandTriggerRequestForTestInstructionExecutionToProcessIn1Second,
		}

		// Send message on channel
		executionStatusCommandChannel <- channelCommandAndMessage

	}()
}

// Process TestExecutionStatus-change
func (messageStreamEngineObject *MessageStreamEngineStruct) processTestExecutionStatusChange(executionsStatusMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage) {

	var err error

	// Process TestCaseExecutionStatus-change
	if executionsStatusMessage.TestCaseExecutionsStatus != nil {
		// Loop TestExecutionStatusMessage
		for _, testCaseExecutionStatusMessage := range executionsStatusMessage.TestCaseExecutionsStatus {

			// Convert TestCaseExecutionVersion into string
			var testCaseExecutionVersionAsString string
			testCaseExecutionVersionAsString = strconv.Itoa(int(testCaseExecutionStatusMessage.TestCaseExecutionVersion))

			//Depending on TestCase-status then act differently
			switch testCaseExecutionStatusMessage.TestCaseExecutionDetails.TestCaseExecutionStatus {

			case fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum_TCE_INITIATED,
				fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum_TCE_EXECUTING:

				// Remove TestCaseInstructionExecution to OnQueue-table
				var testCaseExecutionsOnQueueDataRowAdaptedForUiTableReference *executionsModel.TestCaseExecutionsOnQueueAdaptedForUiTableStruct
				testCaseExecutionsOnQueueDataRowAdaptedForUiTableReference = &executionsModel.TestCaseExecutionsOnQueueAdaptedForUiTableStruct{
					TestCaseExecutionUuid:    testCaseExecutionStatusMessage.TestCaseExecutionUuid,
					TestCaseExecutionVersion: testCaseExecutionVersionAsString,
				}

				// Move TestCaseInstructionExecution from OnQueue-table to UnderExecution-table
				err = executionsUI.MoveTestCaseInstructionExecutionFromOnQueueToUnderExecution(testCaseExecutionsOnQueueDataRowAdaptedForUiTableReference, testCaseExecutionStatusMessage.TestCaseExecutionDetails)
				if err != nil {
					// There were some error som continue to next item in slice
					continue
				}

				//err = executionsUI.RemoveTestCaseExecutionFromOnQueueTable(testCaseExecutionsOnQueueDataRowAdaptedForUiTableReference)
				//if err != nil {
				// There were some error som continue to next item in slice
				//	continue
				//}

			case fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum_TCE_CONTROLLED_INTERRUPTION,
				fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum_TCE_CONTROLLED_INTERRUPTION_CAN_BE_RERUN,
				fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum_TCE_FINISHED_OK,
				fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum_TCE_FINISHED_OK_CAN_BE_RERUN,
				fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum_TCE_FINISHED_NOT_OK,
				fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum_TCE_FINISHED_NOT_OK_CAN_BE_RERUN,
				fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum_TCE_UNEXPECTED_INTERRUPTION,
				fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum_TCE_UNEXPECTED_INTERRUPTION_CAN_BE_RERUN:
				// Remove TestCaseInstructionExecution to UnderExecution-table
				var testCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference *executionsModel.TestCaseExecutionsUnderExecutionAdaptedForUiTableStruct
				testCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference = &executionsModel.TestCaseExecutionsUnderExecutionAdaptedForUiTableStruct{
					TestCaseExecutionUuid:    testCaseExecutionStatusMessage.TestCaseExecutionUuid,
					TestCaseExecutionVersion: testCaseExecutionVersionAsString,
				}
				err = executionsUI.RemoveTestCaseExecutionFromUnderExecutionTable(testCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference)
				if err != nil {
					// There were some error som continue to next item in slice
					continue
				}

				// Add TestCaseInstructionExecution to FinishedExecution-table
				if err != nil {
					// There were some error som continue to next item in slice
					continue
				}

			default:
				// Unknown TestCaseExecutionStatus
				sharedCode.Logger.WithFields(logrus.Fields{
					"ID": "b4164d7a-a485-411d-ad18-feb50ed98566",
					"testCaseExecutionStatusMessage.TestCaseExecutionDetails.TestCaseExecutionStatus": testCaseExecutionStatusMessage.TestCaseExecutionDetails.TestCaseExecutionStatus,
					"executionsStatusMessage": executionsStatusMessage,
				}).Error("Unknown TestCaseExecutionStatus")

				errorId := "6650d51a-787d-48ef-a596-67d7fe9c49cc"
				err := errors.New(fmt.Sprintf("unknown TestCaseExecutionStatus, '%s', in executionsStatusMessage: '%s' [ErrorID: %s]", testCaseExecutionStatusMessage.TestCaseExecutionDetails.TestCaseExecutionStatus, executionsStatusMessage, errorId))

				fmt.Println(err) //TODO send on Error channel

				return

			}

		}
	}

	// Process TestInstructionExecutionStatus-change
	if executionsStatusMessage.TestCaseExecutionsStatus != nil {

	}

}
