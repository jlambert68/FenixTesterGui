package messageStreamEngine

import (
	sharedCode "FenixTesterGui/common_code"
	"fmt"
	"github.com/sirupsen/logrus"
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
