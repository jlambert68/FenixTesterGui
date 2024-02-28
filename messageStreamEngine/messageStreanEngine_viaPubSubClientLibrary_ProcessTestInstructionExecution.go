package messageStreamEngine

import (
	sharedCode "FenixTesterGui/common_code"
	"cloud.google.com/go/pubsub"
	"context"
	"crypto/tls"
	"fmt"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/encoding/protojson"
	"strings"
)

// PullPubSubTestInstructionExecutionMessagesGcpClientLib
// Use GCP Client Library to subscribe to a PubSub-Topic
func PullPubSubTestInstructionExecutionMessagesGcpClientLib() {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "1ed3f12b-65fb-41bc-b12d-ca4af21e8a36",
	}).Debug("Incoming 'PullPubSubTestInstructionExecutionMessagesGcpClientLib'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "1a6fc28d-2523-44c7-9f02-8a7392c4966c",
	}).Debug("Outgoing 'PullPubSubTestInstructionExecutionMessagesGcpClientLib'")

	// Get the GCP-project
	var projectID string
	projectID = sharedCode.GcpProject

	// Extract PubSub-Id and remove characters that are not allowed
	var subID string
	subID = generatePubSubTopicSubscriptionNameForExecutionStatusUpdates(sharedCode.CurrentUserAuthenticatedTowardsGCP)
	subID = strings.ReplaceAll(subID, "@", ".")

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
	if true { //sharedCode.ExecutionLocationForWorker == sharedCode.GCP {

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
			"ID":  "a606a500-7e76-43a5-876d-c51ff8a683ec",
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

	err = clientSubscription.Receive(ctx, func(_ context.Context, pubSubMessage *pubsub.Message) {

		sharedCode.Logger.WithFields(logrus.Fields{
			"ID": "8e75e797-7d75-45fa-93b2-1190c48dd0af",
		}).Debug(fmt.Printf("Got message: %q", string(pubSubMessage.Data)))

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
				"Id":                         "df899f00-5e6a-49c3-a272-ecb3e4bc19f2",
				"Error":                      err,
				"string(pubSubMessage.Data)": string(pubSubMessage.Data),
			}).Error("Something went wrong when converting 'PubSub-message into proto-message")

			// Drop this message, without sending 'Ack'
			return
		}

		// Trigger TestInstruction in parallel while processing next message
		go func() {

		}()

		go func() {
			err = triggerProcessPubSubExecutionStatusMessage(pubSubMessage.Data)
			if err == nil {

				// Acknowledge the message
				// Send 'Ack' back to PubSub-system that message has taken care of
				pubSubMessage.Ack()

			} else {

				sharedCode.Logger.WithFields(logrus.Fields{
					"ID": "2d74199d-a434-4658-a085-46a83c14c8fb",
				}).Error("Failed to Process TestInstructionExecution")

			}
		}()

	})

	// Shouldn't happen
	if err != nil {

		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":  "8acdd7af-d71e-41e0-a7e0-0b36c79c952f",
			"err": err,
		}).Fatalln("PubSub receiver for TestInstructionExecutions ended, which is not intended")

	}

}
