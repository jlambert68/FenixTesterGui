package messageStreamEngine

import (
	"FenixTesterGui/common_code"
	"FenixTesterGui/gcp"
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
	subID = generatePubSubTopicSubscriptionNameForExecutionStatusUpdates(sharedCode.CurrentUserId)

	// Create a loop to be able to have a continuous PubSub Subscription Engine
	var numberOfMessagesInPullResponse int
	var err error

	for {

		// Pull a certain number of messages from Subscription
		numberOfMessagesInPullResponse, err = retrievePubSubMessagesViaRestApi(subID, gcp.GcpObject.GetGcpAccessTokenForAuthorizedAccountsPubSub())

		if err != nil {

			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":  "856533ec-5ba9-46ff-b8c5-af7f3a9da2ac",
				"err": err,
			}).Fatalln("PubSub receiver for TestInstructionExecutions ended, which is not intended")

		}

		// If there are more than zero messages then don't wait
		if numberOfMessagesInPullResponse == 0 {
			// Wait 15 seconds before looking for more PubSub-messages
			time.Sleep(5 * time.Second)
		}

	}
}
