package messageStreamEngine

import sharedCode "FenixTesterGui/common_code"

func initiatePubSubFunctionality(tempGcpProject string) {
	gcpProject = tempGcpProject

	if sharedCode.UseNativeGcpPubSubClientLibrary == true {
		// Use Native GCP PubSub Client Library
		go PullPubSubTestInstructionExecutionMessagesGcpClientLib()

	} else {
		// Use REST to call GCP PubSub
		go PullPubSubTestInstructionExecutionMessagesGcpRestApi()

	}
}

// Create the PubSub-topic from TesterGui-ApplicationUuid
func generatePubSubTopicNameForExecutionStatusUpdates(testerGuiUserId string) (statusExecutionTopic string) {

	var pubSubTopicBase string
	pubSubTopicBase = sharedCode.TestExecutionStatusPubSubTopicBase

	// Build PubSub-topic
	statusExecutionTopic = pubSubTopicBase + "-" + testerGuiUserId

	return statusExecutionTopic
}

// Creates a Topic-Subscription-Name
func generatePubSubTopicSubscriptionNameForExecutionStatusUpdates(testerGuiUserId string) (topicSubscriptionName string) {

	const topicSubscriptionPostfix string = "-sub"

	// Get Topic-name
	var topicID string
	topicID = generatePubSubTopicNameForExecutionStatusUpdates(testerGuiUserId)

	// Create the Topic-Subscription-name
	topicSubscriptionName = topicID + topicSubscriptionPostfix

	return topicSubscriptionName
}
