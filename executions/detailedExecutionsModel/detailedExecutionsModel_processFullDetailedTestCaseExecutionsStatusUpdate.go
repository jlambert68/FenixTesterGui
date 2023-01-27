package detailedExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/executions/detailedTestCaseExecutionUI_summaryTableDefinition"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

// Updates all Executions status with information received after direct gRPC-call to GUiExecutionServer
func (detailedExecutionsModelObject *DetailedExecutionsModelObjectStruct) processFullDetailedTestCaseExecutionsStatusUpdate(
	testCaseExecutionResponse *fenixExecutionServerGuiGrpcApi.TestCaseExecutionResponseMessage) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id":                        "c9bb2cf2-3c1f-42a6-b078-57748ab2b6ff",
		"testCaseExecutionResponse": testCaseExecutionResponse,
	}).Debug("Incoming 'processFullDetailedTestCaseExecutionsStatusUpdate'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "6f0f3a50-3250-426c-9431-698136e93b40",
	}).Debug("Outgoing 'processFullDetailedTestCaseExecutionsStatusUpdate'")

	// Create the TestCaseExecutionMapkey
	var testCaseExecutionMapKey string
	testCaseExecutionMapKey = testCaseExecutionResponse.TestCaseExecutionBasicInformation.TestCaseExecutionUuid +
		strconv.Itoa(int(testCaseExecutionResponse.TestCaseExecutionBasicInformation.TestCaseExecutionVersion))

	// Check if TestCaseExecution already exist
	var existInMap bool
	var testCaseExecutionsDetails *detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsStruct
	testCaseExecutionsDetails, existInMap = detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap[testCaseExecutionMapKey]

	// If TestExecutionExecution doesn't exist in map then create a new instance
	if existInMap == false {

		// Initiate map TestInstructionExecutions
		var TestTestInstructionExecutionsBaseInformationMap map[string]*detailedTestCaseExecutionUI_summaryTableDefinition.TestInstructionExecutionsBaseInformationStruct
		TestTestInstructionExecutionsBaseInformationMap = make(map[string]*detailedTestCaseExecutionUI_summaryTableDefinition.TestInstructionExecutionsBaseInformationStruct)

		// Initiate structure for Execution Summary page
		var tempTestCaseExecutionsBaseInformation *detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsBaseInformationStruct
		tempTestCaseExecutionsBaseInformation = &detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsBaseInformationStruct{
			TestCaseExecutionBasicInformation:                testCaseExecutionResponse.TestCaseExecutionBasicInformation,
			AllTestCaseExecutionsStatusUpdatesInformationMap: make(map[string]*fenixExecutionServerGuiGrpcApi.TestCaseExecutionDetailsMessage),
		}

		testCaseExecutionsDetails = &detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsStruct{
			WaitingForFullTestCaseExecutionUpdate: false,
			TestCaseExecutionStatusMessagesWaitingForFullTestCaseExecutionUpdate: make(chan *fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusMessage,
				detailedTestCaseExecutionUI_summaryTableDefinition.FullExecutionUpdateWhenFirstExecutionStatusReceivedMaxSize),
			TestInstructionExecutionStatusMessagesWaitingForFullTestCaseExecutionUpdate: make(chan *fenixExecutionServerGuiGrpcApi.TestInstructionExecutionStatusMessage,
				detailedTestCaseExecutionUI_summaryTableDefinition.FullExecutionUpdateWhenFirstExecutionStatusReceivedMaxSize),
			FullTestCaseExecutionUpdateWhenFirstExecutionStatusReceived:                false,
			PreviousBroadcastTimeStamp:                                                 time.Time{},
			FullTestCaseExecutionUpdateWhenFirstTestInstructionExecutionStatusReceived: false,
			TestCaseExecutionDatabaseResponseMessage:                                   testCaseExecutionResponse,
			TestCaseExecutionsStatusUpdates:                                            nil,
			TestInstructionExecutionsStatusUpdates:                                     nil,
			TestCaseExecutionsBaseInformation:                                          tempTestCaseExecutionsBaseInformation,
			TestInstructionExecutionsStatusMap:                                         TestTestInstructionExecutionsBaseInformationMap,
			TestInstructionExecutionsStatusForSummaryTable:                             nil,
		}

		// Add the TestCaseExecution to the Map
		detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap[testCaseExecutionMapKey] = testCaseExecutionsDetails

	} else {
		// Replace TestCaseExecutionResponse from Database
		testCaseExecutionsDetails.TestCaseExecutionDatabaseResponseMessage = testCaseExecutionResponse
	}

	// Check if TestCaseStatus exist in 'AllTestCaseExecutionsStatusUpdatesInformationMap'
	// Loop all TestCaseStatus-messages
	for _, tempTestCaseExecutionDetails := range testCaseExecutionResponse.TestCaseExecutionDetails {

		var tempExecutionStatusUpdateTimeStampMapKey string
		tempExecutionStatusUpdateTimeStampMapKey =
			tempTestCaseExecutionDetails.ExecutionStatusUpdateTimeStamp.AsTime().String()

		// Verify if this UpdateTimeStamp exist within 'AllTestCaseExecutionsStatusUpdatesInformationMap'
		_, existInMap = testCaseExecutionsDetails.TestCaseExecutionsBaseInformation.
			AllTestCaseExecutionsStatusUpdatesInformationMap[tempExecutionStatusUpdateTimeStampMapKey]

		// If it doesn't exist then add it to the 'AllTestCaseExecutionsStatusUpdatesInformationMap'
		if existInMap == false {
			testCaseExecutionsDetails.TestCaseExecutionsBaseInformation.
				AllTestCaseExecutionsStatusUpdatesInformationMap[tempExecutionStatusUpdateTimeStampMapKey] = tempTestCaseExecutionDetails

		}
	}

	// TestInstructionsStatus
	// Add the TestInstructions Statuses for summary page to the Map by converting into structure
	for _, testInstructionExecutionDetailsMessage := range testCaseExecutionResponse.TestInstructionExecutions {

		var tempTestInstructionExecutionsStatusMapKey string
		tempTestInstructionExecutionsStatusMapKey = testInstructionExecutionDetailsMessage.TestInstructionExecutionBasicInformation.TestInstructionExecutionUuid +
			strconv.Itoa(int(testInstructionExecutionDetailsMessage.TestInstructionExecutionBasicInformation.TestInstructionExecutionVersion))

		// Check if TestInstructionExecution already exists within Map
		var tempTestTestInstructionExecutionsBaseInformation *detailedTestCaseExecutionUI_summaryTableDefinition.TestInstructionExecutionsBaseInformationStruct
		tempTestTestInstructionExecutionsBaseInformation, existInMap = testCaseExecutionsDetails.TestInstructionExecutionsStatusMap[tempTestInstructionExecutionsStatusMapKey]

		if existInMap == false {
			// TestInstructionExecution doesn't exist, so create the TestInstructionExecution-object
			var tempAllTestInstructionsExecutionsStatusUpdatesInformationMap map[string]*fenixExecutionServerGuiGrpcApi.TestInstructionExecutionsInformationMessage
			tempAllTestInstructionsExecutionsStatusUpdatesInformationMap = make(map[string]*fenixExecutionServerGuiGrpcApi.TestInstructionExecutionsInformationMessage)

			tempTestTestInstructionExecutionsBaseInformation = &detailedTestCaseExecutionUI_summaryTableDefinition.TestInstructionExecutionsBaseInformationStruct{
				TestInstructionExecutionBasicInformation:                 testInstructionExecutionDetailsMessage.TestInstructionExecutionBasicInformation,
				AllTestInstructionsExecutionsStatusUpdatesInformationMap: tempAllTestInstructionsExecutionsStatusUpdatesInformationMap,
			}

			// Save TestInstructionExecution-object back into map
			testCaseExecutionsDetails.TestInstructionExecutionsStatusMap[tempTestInstructionExecutionsStatusMapKey] = tempTestTestInstructionExecutionsBaseInformation

		}

		// Check if TestInstructionStatus exist in 'AllTestInstructionsExecutionsStatusUpdatesInformationMap'
		// Loop all TestInstructionStatus-messages
		for _, tempTestInstructionExecutionsInformation := range testInstructionExecutionDetailsMessage.TestInstructionExecutionsInformation {
			var tempExecutionStatusUpdateTimeStampMapKey string
			tempExecutionStatusUpdateTimeStampMapKey = tempTestInstructionExecutionsInformation.ExecutionStatusUpdateTimeStamp.AsTime().String()

			// Verify if this UpdateTimeStamp exist within 'AllTestInstructionsExecutionsStatusUpdatesInformationMap'
			_, existInMap = tempTestTestInstructionExecutionsBaseInformation.AllTestInstructionsExecutionsStatusUpdatesInformationMap[tempExecutionStatusUpdateTimeStampMapKey]

			// If it doesn't exist then add it to the 'AllTestInstructionsExecutionsStatusUpdatesInformationMap'
			if existInMap == false {
				tempTestTestInstructionExecutionsBaseInformation.AllTestInstructionsExecutionsStatusUpdatesInformationMap[tempExecutionStatusUpdateTimeStampMapKey] = tempTestInstructionExecutionsInformation

			}
		}
	}
	// Process TestCaseStatusUpdate-messages that are waiting on wait-channel
	var numberOfWaitingMessages int

	numberOfWaitingMessages = len(testCaseExecutionsDetails.TestCaseExecutionStatusMessagesWaitingForFullTestCaseExecutionUpdate)

	if numberOfWaitingMessages > 0 {
		for messageCounter := 0; messageCounter < numberOfWaitingMessages; numberOfWaitingMessages++ {

			// Extract one message from wait-channel
			var tempTestCaseExecutionStatusMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusMessage
			tempTestCaseExecutionStatusMessage = <-testCaseExecutionsDetails.TestCaseExecutionStatusMessagesWaitingForFullTestCaseExecutionUpdate

			// Add the message to slice of messages
			var tempTestCaseExecutionStatusMessages []*fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusMessage
			tempTestCaseExecutionStatusMessages = []*fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusMessage{tempTestCaseExecutionStatusMessage}

			// Create message to be sent on channel for processing
			var tempTestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage
			tempTestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage = &fenixExecutionServerGuiGrpcApi.TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage{
				ProtoFileVersionUsedByClient:    0,
				TestCaseExecutionsStatus:        tempTestCaseExecutionStatusMessages,
				TestInstructionExecutionsStatus: nil,
			}

			// Resend the message so it can be processed
			var channelCommandDetailedExecutions ChannelCommandDetailedExecutionsStruct
			channelCommandDetailedExecutions = ChannelCommandDetailedExecutionsStruct{
				ChannelCommandDetailedExecutionsStatus:                            ChannelCommandStatusUpdateOfDetailedExecutionsStatus,
				TestCaseExecutionKey:                                              "",
				FullTestCaseExecutionResponseMessage:                              nil,
				TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage: tempTestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage,
			}
			// Send command on channel
			DetailedExecutionStatusCommandChannel <- channelCommandDetailedExecutions

		}
	}

	// Process TestInstructionStatusUpdate-messages that are waiting on wait-channel
	numberOfWaitingMessages = len(testCaseExecutionsDetails.TestInstructionExecutionStatusMessagesWaitingForFullTestCaseExecutionUpdate)

	if numberOfWaitingMessages > 0 {
		for messageCounter := 0; messageCounter < numberOfWaitingMessages; numberOfWaitingMessages++ {

			// Extract one message from wait-channel
			var tempTestInstructionExecutionStatusMessage *fenixExecutionServerGuiGrpcApi.TestInstructionExecutionStatusMessage
			tempTestInstructionExecutionStatusMessage = <-testCaseExecutionsDetails.TestInstructionExecutionStatusMessagesWaitingForFullTestCaseExecutionUpdate

			// Add the message to slice of messages
			var tempTestInstructionExecutionStatusMessages []*fenixExecutionServerGuiGrpcApi.TestInstructionExecutionStatusMessage
			tempTestInstructionExecutionStatusMessages = []*fenixExecutionServerGuiGrpcApi.TestInstructionExecutionStatusMessage{tempTestInstructionExecutionStatusMessage}

			// Create message to be sent on channel for processing
			var tempTestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage
			tempTestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage = &fenixExecutionServerGuiGrpcApi.TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage{
				ProtoFileVersionUsedByClient:    0,
				TestCaseExecutionsStatus:        nil,
				TestInstructionExecutionsStatus: tempTestInstructionExecutionStatusMessages,
			}

			// Resend the message so it can be procesed
			var channelCommandDetailedExecutions ChannelCommandDetailedExecutionsStruct
			channelCommandDetailedExecutions = ChannelCommandDetailedExecutionsStruct{
				ChannelCommandDetailedExecutionsStatus:                            ChannelCommandStatusUpdateOfDetailedExecutionsStatus,
				TestCaseExecutionKey:                                              "",
				FullTestCaseExecutionResponseMessage:                              nil,
				TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage: tempTestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage,
			}
			// Send command on channel
			DetailedExecutionStatusCommandChannel <- channelCommandDetailedExecutions

		}
	}

	// Update so 'nothing' needs to wait for a FullDetailedStatusUpdate needs to be performed
	testCaseExecutionsDetails.WaitingForFullTestCaseExecutionUpdate = false

	// Update the SummaryTable for TestInstructionExecutions
	detailedExecutionsModelObject.updateTestInstructionExecutionsSummaryTable()

	// Update the SummaryTable for TestCaseExecutions
	detailedExecutionsModelObject.updateTestCaseExecutionsSummaryTable()
}
