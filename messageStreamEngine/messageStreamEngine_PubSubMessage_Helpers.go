package messageStreamEngine

import sharedCode "FenixTesterGui/common_code"

func initiatePubSubFunctionality(tempGcpProject string) {
	gcpProject = tempGcpProject
	go PullPubSubTestInstructionExecutionMessagessages()
}

// Create the PubSub-topic from TesterGui-ApplicationUuid
func generatePubSubTopicForExecutionStatusUpdates(testerGuiUserId string) (statusExecutionTopic string) {

	var pubSubTopicBase string
	pubSubTopicBase = sharedCode.TestExecutionStatusPubSubTopicBase

	// Build PubSub-topic
	statusExecutionTopic = pubSubTopicBase + "-" + testerGuiUserId

	return statusExecutionTopic
}
