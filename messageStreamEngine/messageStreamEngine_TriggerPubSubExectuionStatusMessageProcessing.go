package messageStreamEngine

import (
	sharedCode "FenixTesterGui/common_code"
	"github.com/golang/protobuf/ptypes/timestamp"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/encoding/protojson"
	"strings"
)

func triggerProcessPubSubExecutionStatusMessage(pubSubMessage []byte) (err error) {

	// Remove any unwanted characters
	// Remove '\n'
	var cleanedMessage string
	var cleanedMessageAsByteArray []byte
	var pubSubMessageAsString string

	pubSubMessageAsString = string(pubSubMessage)
	cleanedMessage = strings.ReplaceAll(pubSubMessageAsString, "\n", "")

	// Replace '\"' with '"'
	cleanedMessage = strings.ReplaceAll(cleanedMessage, "\\\"", "\"")

	cleanedMessage = strings.ReplaceAll(cleanedMessage, " ", "")

	// Convert back into byte-array
	cleanedMessageAsByteArray = []byte(cleanedMessage)

	// Convert PubSub-message back into proto-message
	var executionStatusMessagesPubSubMessage fenixExecutionServerGuiGrpcApi.ExecutionStatusMessagesPubSubSchema
	err = protojson.Unmarshal(cleanedMessageAsByteArray, &executionStatusMessagesPubSubMessage)
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"Id":                         "f623c727-1c8d-477d-aa16-8f0a4967ba28",
			"Error":                      err,
			"string(pubSubMessage.Data)": string(pubSubMessage),
		}).Error("Something went wrong when converting 'PubSub-message into proto-message")

		// Drop this message, without sending 'Ack'
		return err
	}

	// Convert from PubSub-message into proto-message used by TesterGui
	var subscribeToMessagesStreamResponse fenixExecutionServerGuiGrpcApi.SubscribeToMessagesStreamResponse
	subscribeToMessagesStreamResponse = fenixExecutionServerGuiGrpcApi.SubscribeToMessagesStreamResponse{
		ProtoFileVersionUsedByClient: 0,
		OriginalMessageCreationTimeStamp: &timestamp.Timestamp{
			Seconds: 0,
			Nanos:   0,
		},
		IsKeepAliveMessage: false,
		ExecutionsStatus: &fenixExecutionServerGuiGrpcApi.TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage{
			ProtoFileVersionUsedByClient:    0,
			TestCaseExecutionsStatus:        nil,
			TestInstructionExecutionsStatus: nil,
		},
	}

	// Convert 'TestCaseExecutionsStatus'
	var testCaseExecutionsStatusSlice []*fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusMessage

	// Loop PubSub-message for GetTestCaseExecutionsStatus-messages
	for _, tempGetTestCaseExecutionsStatusPubSubMessage := range executionStatusMessagesPubSubMessage.GetExecutionsStatus().GetTestCaseExecutionsStatus() {

		var testCaseExecutionsStatus *fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusMessage
		testCaseExecutionsStatus = &fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusMessage{
			TestCaseExecutionUuid:    tempGetTestCaseExecutionsStatusPubSubMessage.GetTestCaseExecutionUuid(),
			TestCaseExecutionVersion: tempGetTestCaseExecutionsStatusPubSubMessage.GetTestCaseExecutionVersion(),
			BroadcastTimeStamp: &timestamp.Timestamp{
				Seconds: tempGetTestCaseExecutionsStatusPubSubMessage.GetBroadcastTimeStamp().GetSeconds(),
				Nanos:   tempGetTestCaseExecutionsStatusPubSubMessage.GetBroadcastTimeStamp().GetNanos(),
			},
			PreviousBroadcastTimeStamp: &timestamp.Timestamp{
				Seconds: tempGetTestCaseExecutionsStatusPubSubMessage.GetPreviousBroadcastTimeStamp().GetSeconds(),
				Nanos:   tempGetTestCaseExecutionsStatusPubSubMessage.GetPreviousBroadcastTimeStamp().GetNanos(),
			},
			TestCaseExecutionDetails: &fenixExecutionServerGuiGrpcApi.TestCaseExecutionDetailsMessage{
				ExecutionStartTimeStamp: &timestamp.Timestamp{
					Seconds: tempGetTestCaseExecutionsStatusPubSubMessage.GetTestCaseExecutionDetails().
						GetExecutionStartTimeStamp().GetSeconds(),
					Nanos: tempGetTestCaseExecutionsStatusPubSubMessage.GetTestCaseExecutionDetails().
						GetExecutionStartTimeStamp().GetNanos(),
				},
				ExecutionStopTimeStamp: &timestamp.Timestamp{
					Seconds: tempGetTestCaseExecutionsStatusPubSubMessage.GetTestCaseExecutionDetails().
						GetExecutionStopTimeStamp().GetSeconds(),
					Nanos: tempGetTestCaseExecutionsStatusPubSubMessage.GetTestCaseExecutionDetails().
						GetExecutionStopTimeStamp().GetNanos(),
				},
				TestCaseExecutionStatus: fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum(
					tempGetTestCaseExecutionsStatusPubSubMessage.GetTestCaseExecutionDetails().
						GetTestCaseExecutionStatus()),
				ExecutionHasFinished: tempGetTestCaseExecutionsStatusPubSubMessage.TestCaseExecutionDetails.
					GetExecutionHasFinished(),
				ExecutionStatusUpdateTimeStamp: &timestamp.Timestamp{
					Seconds: tempGetTestCaseExecutionsStatusPubSubMessage.GetTestCaseExecutionDetails().
						GetExecutionStatusUpdateTimeStamp().GetSeconds(),
					Nanos: tempGetTestCaseExecutionsStatusPubSubMessage.GetTestCaseExecutionDetails().
						GetExecutionStatusUpdateTimeStamp().GetNanos(),
				},
				UniqueDatabaseRowCounter: tempGetTestCaseExecutionsStatusPubSubMessage.GetTestCaseExecutionDetails().
					GetUniqueDatabaseRowCounter(),
			},
		}

		// Append to slice of messages
		testCaseExecutionsStatusSlice = append(testCaseExecutionsStatusSlice, testCaseExecutionsStatus)

	}

	// Add 'TestCaseExecutionsStatus' to proto-message
	subscribeToMessagesStreamResponse.ExecutionsStatus.TestCaseExecutionsStatus = testCaseExecutionsStatusSlice

	// Convert 'TestInstructionExecutionsStatus'
	var testInstructionExecutionsStatusSlice []*fenixExecutionServerGuiGrpcApi.TestInstructionExecutionStatusMessage

	// Loop PubSub-message for GetTestCaseExecutionsStatus-messages
	for _, tempTestInstructionExecutionsStatus := range executionStatusMessagesPubSubMessage.GetExecutionsStatus().
		GetTestInstructionExecutionsStatus() {

		var testInstructionExecutionStatus *fenixExecutionServerGuiGrpcApi.TestInstructionExecutionStatusMessage
		testInstructionExecutionStatus = &fenixExecutionServerGuiGrpcApi.TestInstructionExecutionStatusMessage{
			TestCaseExecutionUuid:           tempTestInstructionExecutionsStatus.GetTestCaseExecutionUuid(),
			TestCaseExecutionVersion:        tempTestInstructionExecutionsStatus.GetTestCaseExecutionVersion(),
			TestInstructionExecutionUuid:    tempTestInstructionExecutionsStatus.GetTestInstructionExecutionUuid(),
			TestInstructionExecutionVersion: tempTestInstructionExecutionsStatus.GetTestInstructionExecutionVersion(),
			TestInstructionExecutionStatus: fenixExecutionServerGuiGrpcApi.TestInstructionExecutionStatusEnum(
				tempTestInstructionExecutionsStatus.GetTestInstructionExecutionStatus()),
			BroadcastTimeStamp: &timestamp.Timestamp{
				Seconds: tempTestInstructionExecutionsStatus.GetBroadcastTimeStamp().GetSeconds(),
				Nanos:   tempTestInstructionExecutionsStatus.GetBroadcastTimeStamp().GetNanos(),
			},
			PreviousBroadcastTimeStamp: &timestamp.Timestamp{
				Seconds: tempTestInstructionExecutionsStatus.GetPreviousBroadcastTimeStamp().GetSeconds(),
				Nanos:   tempTestInstructionExecutionsStatus.GetPreviousBroadcastTimeStamp().GetNanos(),
			},
			TestInstructionExecutionsStatusInformation: &fenixExecutionServerGuiGrpcApi.TestInstructionExecutionsInformationMessage{
				SentTimeStamp: &timestamp.Timestamp{
					Seconds: tempTestInstructionExecutionsStatus.GetTestInstructionExecutionsStatusInformation().
						GetSentTimeStamp().GetSeconds(),
					Nanos: tempTestInstructionExecutionsStatus.GetTestInstructionExecutionsStatusInformation().
						GetSentTimeStamp().GetNanos(),
				},
				ExpectedExecutionEndTimeStamp: &timestamp.Timestamp{
					Seconds: tempTestInstructionExecutionsStatus.GetTestInstructionExecutionsStatusInformation().
						GetExpectedExecutionEndTimeStamp().GetSeconds(),
					Nanos: tempTestInstructionExecutionsStatus.GetTestInstructionExecutionsStatusInformation().
						GetExpectedExecutionEndTimeStamp().GetNanos(),
				},
				TestInstructionExecutionStatus: fenixExecutionServerGuiGrpcApi.TestInstructionExecutionStatusEnum(

					tempTestInstructionExecutionsStatus.GetTestInstructionExecutionsStatusInformation().
						GetTestInstructionExecutionStatus()),
				TestInstructionExecutionEndTimeStamp: &timestamp.Timestamp{
					Seconds: tempTestInstructionExecutionsStatus.GetTestInstructionExecutionsStatusInformation().
						GetTestInstructionExecutionEndTimeStamp().GetSeconds(),
					Nanos: tempTestInstructionExecutionsStatus.GetTestInstructionExecutionsStatusInformation().
						GetTestInstructionExecutionEndTimeStamp().GetNanos(),
				},
				TestInstructionExecutionHasFinished: tempTestInstructionExecutionsStatus.GetTestInstructionExecutionsStatusInformation().
					GetTestInstructionExecutionHasFinished(),
				UniqueDatabaseRowCounter: tempTestInstructionExecutionsStatus.GetTestInstructionExecutionsStatusInformation().
					GetUniqueDatabaseRowCounter(),
				TestInstructionCanBeReExecuted: tempTestInstructionExecutionsStatus.GetTestInstructionExecutionsStatusInformation().
					GetTestInstructionCanBeReExecuted(),
				ExecutionStatusUpdateTimeStamp: &timestamp.Timestamp{
					Seconds: tempTestInstructionExecutionsStatus.GetTestInstructionExecutionsStatusInformation().
						GetExecutionStatusUpdateTimeStamp().GetSeconds(),
					Nanos: tempTestInstructionExecutionsStatus.GetTestInstructionExecutionsStatusInformation().
						GetExecutionStatusUpdateTimeStamp().GetNanos(),
				},
			},
		}

		// Append to slice of messages
		testInstructionExecutionsStatusSlice = append(testInstructionExecutionsStatusSlice, testInstructionExecutionStatus)
	}

	// Add 'TestCaseExecutionsStatus' to proto-message
	subscribeToMessagesStreamResponse.ExecutionsStatus.TestInstructionExecutionsStatus = testInstructionExecutionsStatusSlice

	// Put message containing ExecutionStatus-updates on 'executionStatusCommandChannel' to be processed
	if subscribeToMessagesStreamResponse.ExecutionsStatus != nil {
		var channelCommandAndMessage ChannelCommandStruct
		channelCommandAndMessage = ChannelCommandStruct{
			ChannelCommand:          ChannelCommandExecutionsStatusesHaveBeUpdated,
			ExecutionsStatusMessage: subscribeToMessagesStreamResponse.ExecutionsStatus,
		}

		// Send message on channel
		executionStatusCommandChannel <- channelCommandAndMessage
	}

	return err

}
