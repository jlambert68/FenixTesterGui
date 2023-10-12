package messageStreamEngine

import (
	"FenixTesterGui/common_code"
	"cloud.google.com/go/pubsub"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/encoding/protojson"
	"strings"
	"sync/atomic"
)

func PullPubSubTestInstructionExecutionMessagessages() {
	projectID := sharedCode.GcpProject
	subID := generatePubSubTopicForExecutionStatusUpdates(sharedCode.CurrentUserId)

	var pubSubClient *pubsub.Client
	var err error
	var opts []grpc.DialOption

	ctx := context.Background()

	// Add Access token
	//var returnMessageAckNack bool
	//var returnMessageString string

	//ctx, returnMessageAckNack, returnMessageString = gcp.Gcp.GenerateGCPAccessToken(ctx, gcp.GenerateTokenForPubSub)
	//if returnMessageAckNack == false {
	//	return errors.New(returnMessageString)
	//}

	if len(sharedCode.LocalServiceAccountPath) != 0 {
		//ctx = context.Background()
		pubSubClient, err = pubsub.NewClient(ctx, projectID)
	} else {

	}
	//When running on GCP then use credential otherwise not
	if true { //common_config.ExecutionLocationForWorker == common_config.GCP {

		var creds credentials.TransportCredentials
		creds = credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})

		opts = []grpc.DialOption{
			grpc.WithTransportCredentials(creds),
		}

		pubSubClient, err = pubsub.NewClient(ctx, projectID, option.WithGRPCDialOption(opts[0]))

	}

	if err != nil {

		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":  "8b9f2a07-2931-4585-b2d8-deb6838ada0b",
			"err": err,
		}).Error("Got some problem when creating 'pubsub.NewClient'")

		return
	}
	defer pubSubClient.Close()

	clientSubscription := pubSubClient.Subscription(subID)

	// Receive messages for 10 seconds, which simplifies testing.
	// Comment this out in production, since `Receive` should
	// be used as a long running operation.
	//ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	//defer cancel()

	var received int32
	err = clientSubscription.Receive(ctx, func(_ context.Context, pubSubMessage *pubsub.Message) {

		sharedCode.Logger.WithFields(logrus.Fields{
			"ID": "1995ff58-61c3-418f-a3e8-b9ad3e9e0216",
		}).Debug(fmt.Printf("Got message: %q", string(pubSubMessage.Data)))

		atomic.AddInt32(&received, 1)

		// Remove any unwanted characters
		// Remove '\n'
		var cleanedMessage string
		var cleanedMessageAsByteArray []byte
		var pubSubMessageAsString string

		pubSubMessageAsString = string(pubSubMessage.Data)
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
				"string(pubSubMessage.Data)": string(pubSubMessage.Data),
			}).Error("Something went wrong when converting 'PubSub-message into proto-message")

			// Drop this message, without sending 'Ack'
			return
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

		// Send 'Ack' back to PubSub-system that message has taken care of
		pubSubMessage.Ack()

	})
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":  "12265393-e646-45d4-81b2-e6f69a47de05",
			"err": err,
		}).Fatalln("PubSub receiver for TestInstructionExecutions ended, which is not intended")
	}

}
