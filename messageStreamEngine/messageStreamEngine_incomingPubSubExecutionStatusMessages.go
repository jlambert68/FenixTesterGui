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
		
		// Create 'TestCaseExecutionsStatus'
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
						Nanos:   tempGetTestCaseExecutionsStatusPubSubMessage.GetTestCaseExecutionDetails().
							GetExecutionStartTimeStamp().GetNanos(),
					},
					ExecutionStopTimeStamp: &timestamp.Timestamp{
						Seconds: tempGetTestCaseExecutionsStatusPubSubMessage.GetTestCaseExecutionDetails().
							GetExecutionStopTimeStamp().GetSeconds(),
						Nanos:   tempGetTestCaseExecutionsStatusPubSubMessage.GetTestCaseExecutionDetails().
							GetExecutionStopTimeStamp().GetNanos(),
					},
					TestCaseExecutionStatus: fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum(
						tempGetTestCaseExecutionsStatusPubSubMessage.GetTestCaseExecutionDetails().
							GetTestCaseExecutionStatus()),
					ExecutionHasFinished:    tempGetTestCaseExecutionsStatusPubSubMessage.TestCaseExecutionDetails.
						GetExecutionHasFinished(),
					ExecutionStatusUpdateTimeStamp: &timestamp.Timestamp{
						Seconds: tempGetTestCaseExecutionsStatusPubSubMessage.GetTestCaseExecutionDetails().
							GetExecutionStatusUpdateTimeStamp().GetSeconds(),
						Nanos:   tempGetTestCaseExecutionsStatusPubSubMessage.GetTestCaseExecutionDetails().
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

	}
		
		
		// Create 'TestInstructionExecutionsStatus'
		


		// Split incoming message from GuiExecutionServer into its individual messages and put on channel to MessagechannelEngine
		if subscribeToMessagesStreamResponse.ExecutionsStatus != nil {
			var channelCommandAndMessage ChannelCommandStruct
			channelCommandAndMessage = ChannelCommandStruct{
				ChannelCommand:          ChannelCommandExecutionsStatusesHaveBeUpdated,
				ExecutionsStatusMessage: subscribeToMessagesStreamResponse.ExecutionsStatus,
			}

			// Send message on channel
			executionStatusCommandChannel <- channelCommandAndMessage
		

		// Convert into Message used by converter which is the message from reversed request service
		var processTestInstructionExecutionReveredRequest *fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionReveredRequest
		processTestInstructionExecutionReveredRequest = &fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionReveredRequest{
			ProtoFileVersionUsedByClient: fenixExecutionWorkerGrpcApi.CurrentFenixExecutionWorkerProtoFileVersionEnum(
				processTestInstructionExecutionPubSubRequest.GetProtoFileVersionUsedByClient()),
			TestInstruction: &fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionReveredRequest_TestInstructionExecutionMessage{
				TestInstructionExecutionUuid: processTestInstructionExecutionPubSubRequest.TestInstruction.GetTestInstructionExecutionUuid(),
				TestInstructionUuid:          processTestInstructionExecutionPubSubRequest.TestInstruction.GetTestInstructionUuid(),
				TestInstructionName:          processTestInstructionExecutionPubSubRequest.TestInstruction.GetTestInstructionName(),
				MajorVersionNumber:           processTestInstructionExecutionPubSubRequest.TestInstruction.GetMajorVersionNumber(),
				MinorVersionNumber:           processTestInstructionExecutionPubSubRequest.TestInstruction.GetMinorVersionNumber(),
				TestInstructionAttributes:    nil, // Converted below
			},
			TestData: &fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionReveredRequest_TestDataMessage{
				TestDataSetUuid:           processTestInstructionExecutionPubSubRequest.TestData.GetTestDataSetUuid(),
				ManualOverrideForTestData: nil, // Converted below
			},
		}

		// Convert 'TestInstruction:TestInstructionAttributes'
		var tempTestInstructionAttributes []*fenixExecutionWorkerGrpcApi.
			ProcessTestInstructionExecutionReveredRequest_TestInstructionAttributeMessage

		// Loop 'TestInstructionAttributes' from PubSub-message
		for _, pubSubTestInstructionAttribute := range processTestInstructionExecutionPubSubRequest.TestInstruction.TestInstructionAttributes {
			var tempTestInstructionAttribute *fenixExecutionWorkerGrpcApi.
				ProcessTestInstructionExecutionReveredRequest_TestInstructionAttributeMessage
			tempTestInstructionAttribute = &fenixExecutionWorkerGrpcApi.
				ProcessTestInstructionExecutionReveredRequest_TestInstructionAttributeMessage{
				TestInstructionAttributeType: fenixExecutionWorkerGrpcApi.TestInstructionAttributeTypeEnum(
					pubSubTestInstructionAttribute.GetTestInstructionAttributeType()),
				TestInstructionAttributeUuid:     pubSubTestInstructionAttribute.GetTestInstructionAttributeUuid(),
				TestInstructionAttributeName:     pubSubTestInstructionAttribute.GetTestInstructionAttributeName(),
				AttributeValueAsString:           pubSubTestInstructionAttribute.GetAttributeValueAsString(),
				AttributeValueUuid:               pubSubTestInstructionAttribute.GetTestInstructionAttributeUuid(),
				TestInstructionAttributeTypeUuid: pubSubTestInstructionAttribute.GetTestInstructionAttributeTypeUuid(),
				TestInstructionAttributeTypeName: pubSubTestInstructionAttribute.GetTestInstructionAttributeTypeName(),
			}

			// Append to slice of 'TestInstructionAttributes'
			tempTestInstructionAttributes = append(tempTestInstructionAttributes, tempTestInstructionAttribute)
		}

		processTestInstructionExecutionReveredRequest.TestInstruction.TestInstructionAttributes = tempTestInstructionAttributes

		// Convert 'TestData:ManualOverrideForTestData'
		var tempManualOverrideForTestDataSlice []*fenixExecutionWorkerGrpcApi.
			ProcessTestInstructionExecutionReveredRequest_TestDataMessage_ManualOverrideForTestDataMessage

		// Loop 'TestInstructionAttributes' from PubSub-message
		for _, pubSubManualOverrideForTestData := range processTestInstructionExecutionPubSubRequest.TestData.ManualOverrideForTestData {
			var tempManualOverrideForTestDataMessage *fenixExecutionWorkerGrpcApi.
				ProcessTestInstructionExecutionReveredRequest_TestDataMessage_ManualOverrideForTestDataMessage
			tempManualOverrideForTestDataMessage = &fenixExecutionWorkerGrpcApi.
				ProcessTestInstructionExecutionReveredRequest_TestDataMessage_ManualOverrideForTestDataMessage{
				TestDataSetAttributeUuid:  pubSubManualOverrideForTestData.GetTestDataSetAttributeUuid(),
				TestDataSetAttributeName:  pubSubManualOverrideForTestData.GetTestDataSetAttributeName(),
				TestDataSetAttributeValue: pubSubManualOverrideForTestData.GetTestDataSetAttributeValue(),
			}

			// Append to slice of 'TestInstructionAttributes'
			tempManualOverrideForTestDataSlice = append(tempManualOverrideForTestDataSlice, tempManualOverrideForTestDataMessage)
		}

		processTestInstructionExecutionReveredRequest.TestData.ManualOverrideForTestData = tempManualOverrideForTestDataSlice

		// Call 'CA' backend to convert 'TestInstruction' into useful structure later to be used by FangEngine
		var tempProcessTestInstructionExecutionResponse *fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionResponse
		var fangEngineRestApiMessageValues *restCallsToCAEngine.FangEngineRestApiMessageStruct
		_, tempProcessTestInstructionExecutionResponse, fangEngineRestApiMessageValues =
			messagesToExecutionWorkerServer.ConvertTestInstructionIntoFangEngineStructure(
				processTestInstructionExecutionReveredRequest)

		// Send 'ProcessTestInstructionExecutionPubSubRequest-response' back to worker over direct gRPC-call
		couldSend, returnMessage := connectorEngine.TestInstructionExecutionEngine.
			MessagesToExecutionWorkerObjectReference.
			SendConnectorProcessTestInstructionExecutionResponse(tempProcessTestInstructionExecutionResponse)

		if couldSend == false {
			common_config.Logger.WithFields(logrus.Fields{
				"ID":            "55820706-bd18-41a6-be0a-c7d3b649e0e2",
				"returnMessage": returnMessage,
			}).Error("Couldn't send response to Worker")

		} else {

			// Send TestInstruction to FangEngine using RestCall
			var finalTestInstructionExecutionResultMessage *fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage
			finalTestInstructionExecutionResultMessage = messagesToExecutionWorkerServer.SendTestInstructionToFangEngineUsingRestCall(
				fangEngineRestApiMessageValues, processTestInstructionExecutionReveredRequest)

			// Send 'ProcessTestInstructionExecutionReversedResponse' back to worker over direct gRPC-call
			couldSend, returnMessage := connectorEngine.TestInstructionExecutionEngine.MessagesToExecutionWorkerObjectReference.
				SendReportCompleteTestInstructionExecutionResultToFenixWorkerServer(finalTestInstructionExecutionResultMessage)

			if couldSend == false {
				common_config.Logger.WithFields(logrus.Fields{
					"ID": "1ce93ee2-5542-4437-9c05-d7f9d19313fa",
					"finalTestInstructionExecutionResultMessage": finalTestInstructionExecutionResultMessage,
					"returnMessage": returnMessage,
				}).Error("Couldn't send response to Worker")

			} else {

				// Send 'Ack' back to PubSub-system that message has taken care of
				pubSubMessage.Ack()
			}
		}

	})
	if err != nil {
		common_config.Logger.WithFields(logrus.Fields{
			"ID":  "2410eaa0-dce7-458b-ad9b-28d53680f995",
			"err": err,
		}).Fatalln("PubSub receiver for TestInstructionExecutions ended, which is not intended")
	}

}
