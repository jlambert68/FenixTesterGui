package messageStreamEngine

import (
	"FenixTesterGui/common_code"
	"FenixTesterGui/gcp"
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

// PullPubSubTestInstructionExecutionMessagesGcpRestApi
// Use GCP RestApi to subscribe to a PubSub-Topic
func PullPubSubTestInstructionExecutionMessagesGcpRestApi() {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "7ad67b62-74cc-4ebf-806d-1dd0f5a08f98",
	}).Debug("Incoming 'PullPubSubTestInstructionExecutionMessagesGcpRestApi'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "2fbf4889-0a6f-48db-89f9-09b204d51162",
	}).Debug("Outgoing 'PullPubSubTestInstructionExecutionMessagesGcpRestApi'")

	// Generate Subscription name to use
	var subID string
	subID = generatePubSubTopicSubscriptionNameForExecutionStatusUpdates(sharedCode.CurrentUserIdLogedInOnComputer)

	// Create a loop to be able to have a continuous PubSub Subscription Engine
	var numberOfMessagesInPullResponse int
	var err error
	var returnAckNack bool
	var returnMessage string
	var ctx context.Context

	ctx = context.Background()

	for {

		// Generate a new token is needed
		_, returnAckNack, returnMessage = gcp.GcpObject.GenerateGCPAccessToken(ctx, gcp.TargetServerGuiExecutionServer)
		if returnAckNack == false {

			// Set to zero because we need some waiting time
			numberOfMessagesInPullResponse = 0

			sharedCode.Logger.WithFields(logrus.Fields{
				"id":            "4d4f1144-a905-4b3c-8d71-ef533eea514c",
				"returnMessage": returnMessage,
			}).Debug("Problem when generating a new token. Waiting some time before next try")

		} else {

			// Pull a certain number of messages from Subscription
			numberOfMessagesInPullResponse, err = retrievePubSubMessagesViaRestApi(subID, gcp.GcpObject.GetGcpAccessTokenForAuthorizedAccountsPubSub())

			if err != nil {

				sharedCode.Logger.WithFields(logrus.Fields{
					"ID":  "856533ec-5ba9-46ff-b8c5-af7f3a9da2ac",
					"err": err,
				}).Fatalln("PubSub receiver for TestInstructionExecutions ended, which is not intended")

			}

		}

		// If there are more than zero messages then don't wait
		if numberOfMessagesInPullResponse == 0 {
			// Wait 15 seconds before looking for more PubSub-messages
			time.Sleep(5 * time.Second)
		}

	}
}
