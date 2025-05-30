package detailedExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/executions/detailedTestCaseExecutionUI_summaryTableDefinition"
	"errors"
	"fmt"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

// Updates specific status information based on subscriptions updates from GuiExecutionServer
func (detailedExecutionsModelObject *DetailedExecutionsModelObjectStruct) processStatusUpdateOfDetailedExecutionsStatus(
	testCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage *fenixExecutionServerGuiGrpcApi.
		TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "f92a832b-4020-432e-a7bc-8eb630b24bbf",
		"testCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage": testCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage,
	}).Debug("Incoming 'processStatusUpdateOfDetailedExecutionsStatus'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "fd279ee5-0c5b-4e72-8a62-da353b63f95f",
	}).Debug("Outgoing 'processStatusUpdateOfDetailedExecutionsStatus'")

	// Process TestCaseStatus-messages to check that all TestCasesMapPtr exist in 'detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap'
	var testCaseExecutionKeysMap map[string]string // map[tempTestCaseExecutionMapKey]tempTestCaseExecutionMapKey
	testCaseExecutionKeysMap = make(map[string]string)
	var existInMap bool
	var existInTestCaseExecutionsDetailsMap bool
	var existsInTestCaseExecutionKeysMap bool

	// Loop all TestCaseStatus-messages
	for _, tempTestCaseExecutionStatusMessage := range testCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage.TestCaseExecutionsStatus {

		// Create TestCaseExecutionKey
		var tempTestCaseExecutionToBeFullyRetrievedMapKey string
		tempTestCaseExecutionToBeFullyRetrievedMapKey = tempTestCaseExecutionStatusMessage.TestCaseExecutionUuid +
			strconv.Itoa(int(tempTestCaseExecutionStatusMessage.TestCaseExecutionVersion))

		// Check if TestCaseExecution exist within the 'TestCaseExecutionsDetailsMap'
		var tempTestCaseExecutionsDetails *detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsStruct
		tempTestCaseExecutionsDetails, existInTestCaseExecutionsDetailsMap = detailedTestCaseExecutionUI_summaryTableDefinition.
			TestCaseExecutionsDetailsMap[tempTestCaseExecutionToBeFullyRetrievedMapKey]

		// If not then add it to the Map over TestCaseExecution to retrieve from the Database
		if existInTestCaseExecutionsDetailsMap == false {

			_, existsInTestCaseExecutionKeysMap = testCaseExecutionKeysMap[tempTestCaseExecutionToBeFullyRetrievedMapKey]
			// Has the tempTestCaseExecutionToBeFullyRetrievedMapKey already been saved
			if existsInTestCaseExecutionKeysMap == false {
				testCaseExecutionKeysMap[tempTestCaseExecutionToBeFullyRetrievedMapKey] = tempTestCaseExecutionToBeFullyRetrievedMapKey

				// Initiate map TestInstructionExecutions
				var TestTestInstructionExecutionsBaseInformationMap map[string]*detailedTestCaseExecutionUI_summaryTableDefinition.
					TestInstructionExecutionsBaseInformationStruct
				TestTestInstructionExecutionsBaseInformationMap = make(map[string]*detailedTestCaseExecutionUI_summaryTableDefinition.
					TestInstructionExecutionsBaseInformationStruct)

				// Initiate structure for Execution Summary page
				var tempTestCaseExecutionsBaseInformation *detailedTestCaseExecutionUI_summaryTableDefinition.
					TestCaseExecutionsBaseInformationStruct
				tempTestCaseExecutionsBaseInformation = &detailedTestCaseExecutionUI_summaryTableDefinition.
					TestCaseExecutionsBaseInformationStruct{
					TestCaseExecutionBasicInformation: nil,
					AllTestCaseExecutionsStatusUpdatesInformationMap: make(map[string]*fenixExecutionServerGuiGrpcApi.
						TestCaseExecutionDetailsMessage),
				}

				tempTestCaseExecutionsDetails = &detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsStruct{
					WaitingForFullTestCaseExecutionUpdate:                                                    true,
					WaitingForFullTestCaseExecutionUpdateAfterFirstTestInstructionExecutionStatusWasReceived: false,
					TestCaseExecutionStatusMessagesWaitingForFullTestCaseExecutionUpdate: make(
						chan *fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusMessage,
						detailedTestCaseExecutionUI_summaryTableDefinition.
							FullExecutionUpdateWhenFirstExecutionStatusReceivedMaxSize),
					TestInstructionExecutionStatusMessagesWaitingForFullTestCaseExecutionUpdate: make(
						chan *fenixExecutionServerGuiGrpcApi.TestInstructionExecutionStatusMessage,
						detailedTestCaseExecutionUI_summaryTableDefinition.
							FullExecutionUpdateWhenFirstExecutionStatusReceivedMaxSize),
					FirstExecutionStatusReceived:                   true,
					PreviousBroadcastTimeStamp:                     time.Time{},
					FirstTestInstructionExecutionStatusReceived:    false,
					TestCaseExecutionDatabaseResponseMessage:       nil,
					TestCaseExecutionsStatusUpdates:                nil,
					TestInstructionExecutionsStatusUpdates:         nil,
					TestCaseExecutionsBaseInformation:              tempTestCaseExecutionsBaseInformation,
					TestInstructionExecutionsStatusMap:             TestTestInstructionExecutionsBaseInformationMap,
					TestInstructionExecutionsStatusForSummaryTable: nil,
				}

				// Add the TestCaseExecution to the Map
				detailedTestCaseExecutionUI_summaryTableDefinition.
					TestCaseExecutionsDetailsMap[tempTestCaseExecutionToBeFullyRetrievedMapKey] = tempTestCaseExecutionsDetails

				// Add TestCaseStatusMessage to Waiting-channel
				tempTestCaseExecutionsDetails.
					TestCaseExecutionStatusMessagesWaitingForFullTestCaseExecutionUpdate <- tempTestCaseExecutionStatusMessage

				// Set that the first TestExecutionStatusUpdate was received
				tempTestCaseExecutionsDetails.FirstExecutionStatusReceived = true
			}
		} else {

			// Add TestCaseStatusMessage to Waiting-channel when TestCaseExecution is withing for full TestCaseExecution-StatusUpdate
			if tempTestCaseExecutionsDetails.WaitingForFullTestCaseExecutionUpdate == true {
				tempTestCaseExecutionsDetails.
					TestCaseExecutionStatusMessagesWaitingForFullTestCaseExecutionUpdate <- tempTestCaseExecutionStatusMessage
			}

			// Set that the first TestExecutionStatusUpdate was received
			tempTestCaseExecutionsDetails.FirstExecutionStatusReceived = true

		}
	}

	// Process TestInstructionStatus-messages to check that all TestCasesMapPtr exist in 'detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap'
	for _, tempTestInstructionExecutionStatusMessage := range testCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage.
		TestInstructionExecutionsStatus {

		// Create TestCaseExecutionKey
		var tempTestCaseExecutionToBeFullyRetrievedMapKey string
		tempTestCaseExecutionToBeFullyRetrievedMapKey = tempTestInstructionExecutionStatusMessage.TestCaseExecutionUuid +
			strconv.Itoa(int(tempTestInstructionExecutionStatusMessage.TestCaseExecutionVersion))

		// Check if TestCaseExecution exist within the 'TestCaseExecutionsDetailsMap'
		var tempTestCaseExecutionsDetails *detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsStruct
		tempTestCaseExecutionsDetails, existInTestCaseExecutionsDetailsMap = detailedTestCaseExecutionUI_summaryTableDefinition.
			TestCaseExecutionsDetailsMap[tempTestCaseExecutionToBeFullyRetrievedMapKey]

		// If not then add it to the Map over TestCaseExecution to retrieve from the Database
		if existInTestCaseExecutionsDetailsMap == false {

			_, existsInTestCaseExecutionKeysMap = testCaseExecutionKeysMap[tempTestCaseExecutionToBeFullyRetrievedMapKey]

			// Has the tempTestCaseExecutionToBeFullyRetrievedMapKey already been saved
			if existsInTestCaseExecutionKeysMap == false {
				// Force to do full TestCaseExecutionDetails-update
				testCaseExecutionKeysMap[tempTestCaseExecutionToBeFullyRetrievedMapKey] = tempTestCaseExecutionToBeFullyRetrievedMapKey

				//_, existsInTestCaseExecutionKeysMap = testCaseExecutionKeysMap[tempTestCaseExecutionToBeFullyRetrievedMapKey]

				// Has the tempTestCaseExecutionToBeFullyRetrievedMapKey already been saved
				//if existsInTestCaseExecutionKeysMap == false {
				//	testCaseExecutionKeysMap[tempTestCaseExecutionToBeFullyRetrievedMapKey] = tempTestCaseExecutionToBeFullyRetrievedMapKey

				// Initiate map TestInstructionExecutions
				var TestTestInstructionExecutionsBaseInformationMap map[string]*detailedTestCaseExecutionUI_summaryTableDefinition.TestInstructionExecutionsBaseInformationStruct
				TestTestInstructionExecutionsBaseInformationMap = make(map[string]*detailedTestCaseExecutionUI_summaryTableDefinition.TestInstructionExecutionsBaseInformationStruct)

				// Initiate structure for Execution Summary page
				var tempTestCaseExecutionsBaseInformation *detailedTestCaseExecutionUI_summaryTableDefinition.
					TestCaseExecutionsBaseInformationStruct
				tempTestCaseExecutionsBaseInformation = &detailedTestCaseExecutionUI_summaryTableDefinition.
					TestCaseExecutionsBaseInformationStruct{
					TestCaseExecutionBasicInformation: nil,
					AllTestCaseExecutionsStatusUpdatesInformationMap: make(map[string]*fenixExecutionServerGuiGrpcApi.
						TestCaseExecutionDetailsMessage),
				}

				tempTestCaseExecutionsDetails = &detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsStruct{
					WaitingForFullTestCaseExecutionUpdate:                                                    true,
					WaitingForFullTestCaseExecutionUpdateAfterFirstTestInstructionExecutionStatusWasReceived: true,
					TestCaseExecutionStatusMessagesWaitingForFullTestCaseExecutionUpdate: make(
						chan *fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusMessage,
						detailedTestCaseExecutionUI_summaryTableDefinition.FullExecutionUpdateWhenFirstExecutionStatusReceivedMaxSize),
					TestInstructionExecutionStatusMessagesWaitingForFullTestCaseExecutionUpdate: make(
						chan *fenixExecutionServerGuiGrpcApi.TestInstructionExecutionStatusMessage,
						detailedTestCaseExecutionUI_summaryTableDefinition.FullExecutionUpdateWhenFirstExecutionStatusReceivedMaxSize),
					FirstExecutionStatusReceived:                   true,
					PreviousBroadcastTimeStamp:                     time.Time{},
					FirstTestInstructionExecutionStatusReceived:    true,
					TestCaseExecutionDatabaseResponseMessage:       nil,
					TestCaseExecutionsStatusUpdates:                nil,
					TestInstructionExecutionsStatusUpdates:         nil,
					TestCaseExecutionsBaseInformation:              tempTestCaseExecutionsBaseInformation,
					TestInstructionExecutionsStatusMap:             TestTestInstructionExecutionsBaseInformationMap,
					TestInstructionExecutionsStatusForSummaryTable: nil,
				}

				// Add the TestCaseExecution to the Map
				detailedTestCaseExecutionUI_summaryTableDefinition.
					TestCaseExecutionsDetailsMap[tempTestCaseExecutionToBeFullyRetrievedMapKey] = tempTestCaseExecutionsDetails

				//}

				// Add TestInstructionStatusMessage to Waiting-channel
				tempTestCaseExecutionsDetails.
					TestInstructionExecutionStatusMessagesWaitingForFullTestCaseExecutionUpdate <- tempTestInstructionExecutionStatusMessage

			} else {
				// Add TestInstructionStatusMessage to Waiting-channel when TestCaseExecution is withing for full TestCaseExecution-StatusUpdate
				if tempTestCaseExecutionsDetails.WaitingForFullTestCaseExecutionUpdate == true ||
					tempTestCaseExecutionsDetails.
						WaitingForFullTestCaseExecutionUpdateAfterFirstTestInstructionExecutionStatusWasReceived == true {
					tempTestCaseExecutionsDetails.
						TestInstructionExecutionStatusMessagesWaitingForFullTestCaseExecutionUpdate <- tempTestInstructionExecutionStatusMessage
				}
			}
		} else {

			// Is this the first TestInstructionStatusUpdate-message
			if tempTestCaseExecutionsDetails.FirstTestInstructionExecutionStatusReceived == false {

				// Set that the first TestInstructionExecutionStatusUpdate was received
				tempTestCaseExecutionsDetails.FirstExecutionStatusReceived = true
				tempTestCaseExecutionsDetails.FirstTestInstructionExecutionStatusReceived = true

				// Set that we wait for first FullStatusUpdate after first TestInstructionExecutionStatusUpdate
				tempTestCaseExecutionsDetails.
					WaitingForFullTestCaseExecutionUpdateAfterFirstTestInstructionExecutionStatusWasReceived = true

				// TestCaseExecution Exist in map so now check if we should wait for FullStatusExecutionUpdate
				if tempTestCaseExecutionsDetails.WaitingForFullTestCaseExecutionUpdate == true ||
					tempTestCaseExecutionsDetails.
						WaitingForFullTestCaseExecutionUpdateAfterFirstTestInstructionExecutionStatusWasReceived == true {

					// Set that the first TestInstructionExecutionStatusUpdate was received
					tempTestCaseExecutionsDetails.FirstExecutionStatusReceived = true
					tempTestCaseExecutionsDetails.FirstTestInstructionExecutionStatusReceived = true

					// Put Status Message on WaitQueue
					tempTestCaseExecutionsDetails.
						TestInstructionExecutionStatusMessagesWaitingForFullTestCaseExecutionUpdate <- tempTestInstructionExecutionStatusMessage

					// Check if a FullTestCaseStatusUpdate already is to be done
					_, existsInTestCaseExecutionKeysMap = testCaseExecutionKeysMap[tempTestCaseExecutionToBeFullyRetrievedMapKey]

					// Has the tempTestCaseExecutionToBeFullyRetrievedMapKey already been saved
					if existsInTestCaseExecutionKeysMap == false {

						// Add TestCaseExecutionKey for a FullTestCaseStatusUpdate
						testCaseExecutionKeysMap[tempTestCaseExecutionToBeFullyRetrievedMapKey] = tempTestCaseExecutionToBeFullyRetrievedMapKey
					}
				} else {
					tempTestCaseExecutionsDetails.
						WaitingForFullTestCaseExecutionUpdateAfterFirstTestInstructionExecutionStatusWasReceived = true
					// Has the tempTestCaseExecutionToBeFullyRetrievedMapKey already been saved
					if existsInTestCaseExecutionKeysMap == false {

						// Add TestCaseExecutionKey for a FullTestCaseStatusUpdate
						testCaseExecutionKeysMap[tempTestCaseExecutionToBeFullyRetrievedMapKey] = tempTestCaseExecutionToBeFullyRetrievedMapKey

						tempTestCaseExecutionsDetails.
							TestInstructionExecutionStatusMessagesWaitingForFullTestCaseExecutionUpdate <- tempTestInstructionExecutionStatusMessage

					}
				}
			} else {

				// TestCaseExecution Exist in map so now check if we should wait for FullStatusExecutionUpdate
				if tempTestCaseExecutionsDetails.WaitingForFullTestCaseExecutionUpdate == true ||
					tempTestCaseExecutionsDetails.
						WaitingForFullTestCaseExecutionUpdateAfterFirstTestInstructionExecutionStatusWasReceived == true {

					// Put Status Message on WaitQueue
					tempTestCaseExecutionsDetails.
						TestInstructionExecutionStatusMessagesWaitingForFullTestCaseExecutionUpdate <- tempTestInstructionExecutionStatusMessage

				}
			}
		}
	}

	// If there are only TestCaseExecution that is not already within the TestCaseExecutionsMap then get them first before process the updates
	// This is also done when this is the first Status-message that is received
	if len(testCaseExecutionKeysMap) > 0 {

		// Loop all TestCaseExecutionKeys and send commands to retrieve the TestCaseExecutions
		for _, testCaseExecutionKey := range testCaseExecutionKeysMap {

			// Create command to retrieve missing TestCaseExecutions, via channelEngine
			var channelCommandDetailedExecutions ChannelCommandDetailedExecutionsStruct
			channelCommandDetailedExecutions = ChannelCommandDetailedExecutionsStruct{
				ChannelCommandDetailedExecutionsStatus:                            ChannelCommandRetrieveFullDetailedTestCaseExecution,
				TestCaseExecutionKey:                                              testCaseExecutionKey,
				FullTestCaseExecutionResponseMessage:                              nil,
				TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage: nil,
			}

			// Don't put on Channel if more than 9 items from max capacity
			var currentChannelSize int32
			currentChannelSize = int32(len(DetailedExecutionStatusCommandChannel))
			if currentChannelSize > MessageChannelMaxSizeDetailedExecutionStatus-9 {
				for {
					time.Sleep(5 * time.Second)

					currentChannelSize = int32(len(DetailedExecutionStatusCommandChannel))
					if currentChannelSize < MessageChannelMaxSizeDetailedExecutionStatus-9 {
						break
					}
				}
			}

			// Send command ion channel
			DetailedExecutionStatusCommandChannel <- channelCommandDetailedExecutions
		}

	}

	// All TestCaseExecutions exist within 'TestCaseExecutionsDetailsMap' so start process the status updates

	// Loop all TestCaseStatus-messages and update TestCaseExecutionStatus for each TestCase
	for _, tempTestCaseExecutionStatusMessage := range testCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage.
		TestCaseExecutionsStatus {

		// Create TestCaseExecutionKey
		var tempTestCaseExecutionMapKey string
		tempTestCaseExecutionMapKey = tempTestCaseExecutionStatusMessage.TestCaseExecutionUuid +
			strconv.Itoa(int(tempTestCaseExecutionStatusMessage.TestCaseExecutionVersion))

		// Check if TestCaseExecution exist within the 'TestCaseExecutionsDetailsMap'
		var tempTestCaseExecutionsDetailsReference *detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsStruct
		tempTestCaseExecutionsDetailsReference, existInTestCaseExecutionsDetailsMap = detailedTestCaseExecutionUI_summaryTableDefinition.
			TestCaseExecutionsDetailsMap[tempTestCaseExecutionMapKey]

		// If not then something is really wrong
		if existInTestCaseExecutionsDetailsMap == false {
			ErrorID := "b33739a9-e5c4-452c-891e-008b8c1a8a1d"
			err := errors.New(fmt.Sprintf("We shouldn't end up here. [ErrorID:'%s']", ErrorID))

			fmt.Println(err) // TODO Send on Error-channel

			return
		}

		// Only Process TestInstructionExecutionStatus-message if there is no waiting for a Full TestCaseExecution-status-update
		if tempTestCaseExecutionsDetailsReference.WaitingForFullTestCaseExecutionUpdate == true {
			continue
		}

		// Append Incoming TestCaseExecutionStatus-message into stored Map-message
		tempTestCaseExecutionsDetailsReference.TestCaseExecutionsStatusUpdates = append(
			tempTestCaseExecutionsDetailsReference.TestCaseExecutionsStatusUpdates, tempTestCaseExecutionStatusMessage)

		// Extract UpdateStatusTimeStamp to be used as MapKey
		var tempExecutionStatusUpdateTimeStampMapKey string
		tempExecutionStatusUpdateTimeStampMapKey = tempTestCaseExecutionStatusMessage.
			TestCaseExecutionDetails.ExecutionStatusUpdateTimeStamp.AsTime().String()

		// Verify if this UpdateTimeStamp exist within 'AllTestCaseExecutionsStatusUpdatesInformationMap'
		_, existInMap = tempTestCaseExecutionsDetailsReference.TestCaseExecutionsBaseInformation.
			AllTestCaseExecutionsStatusUpdatesInformationMap[tempExecutionStatusUpdateTimeStampMapKey]

		// If it doesn't exist then add it to the 'AllTestInstructionsExecutionsStatusUpdatesInformationMap'
		if existInMap == false {
			tempTestCaseExecutionsDetailsReference.TestCaseExecutionsBaseInformation.
				AllTestCaseExecutionsStatusUpdatesInformationMap[tempExecutionStatusUpdateTimeStampMapKey] =
				tempTestCaseExecutionStatusMessage.TestCaseExecutionDetails

			// Is there a mismatch between locally store Previous Timestamp Status-message and incoming Previous Timestamp Status-message
			// When BroadcastTimeStamp == PreviousBroadcastTimeStamp in incoming status message then that message is the first for that TestCaseExecution, for this ExecutionServerInstance
			if tempTestCaseExecutionStatusMessage.BroadcastTimeStamp != tempTestCaseExecutionStatusMessage.PreviousBroadcastTimeStamp &&
				tempTestCaseExecutionStatusMessage.PreviousBroadcastTimeStamp.AsTime() != tempTestCaseExecutionsDetailsReference.
					PreviousBroadcastTimeStamp &&
				tempTestCaseExecutionsDetailsReference.PreviousBroadcastTimeStamp.IsZero() != true {

				// Create command to retrieve missing TestCaseExecutions, via channelEngine
				var channelCommandDetailedExecutions ChannelCommandDetailedExecutionsStruct
				channelCommandDetailedExecutions = ChannelCommandDetailedExecutionsStruct{
					ChannelCommandDetailedExecutionsStatus:                            ChannelCommandRetrieveFullDetailedTestCaseExecution,
					TestCaseExecutionKey:                                              tempTestCaseExecutionMapKey,
					FullTestCaseExecutionResponseMessage:                              nil,
					TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage: nil,
				}

				// Don't put on Channel if more than 9 items from max capacity
				var currentChannelSize int32
				currentChannelSize = int32(len(DetailedExecutionStatusCommandChannel))
				if currentChannelSize > MessageChannelMaxSizeDetailedExecutionStatus-9 {
					for {
						time.Sleep(5 * time.Second)

						currentChannelSize = int32(len(DetailedExecutionStatusCommandChannel))
						if currentChannelSize < MessageChannelMaxSizeDetailedExecutionStatus-9 {
							break
						}
					}
				}

				// Send command ion channel
				DetailedExecutionStatusCommandChannel <- channelCommandDetailedExecutions

			} else {

				// Save new BroadcastTimestamp as PreviousBroadcastTimestamp
				tempTestCaseExecutionsDetailsReference.PreviousBroadcastTimeStamp = tempTestCaseExecutionStatusMessage.
					BroadcastTimeStamp.AsTime()
			}

			// Is the 'TestCaseExecutionStatus' an end-status?
			var isEndStatus bool
			isEndStatus = isTestCaseExecutionStatusAnEndStatus(tempTestCaseExecutionStatusMessage.
				TestCaseExecutionDetails.GetTestCaseExecutionStatus())

			// When TestCaseExecutionStatus is an EndStatus then retrieve Full TestCaseExecution Result

			if tempTestCaseExecutionStatusMessage.GetTestCaseExecutionDetails().GetExecutionHasFinished() == true ||
				isEndStatus {
				// Create command to retrieve Full TestCaseExecutions, via channelEngine
				var channelCommandDetailedExecutions ChannelCommandDetailedExecutionsStruct
				channelCommandDetailedExecutions = ChannelCommandDetailedExecutionsStruct{
					ChannelCommandDetailedExecutionsStatus:                            ChannelCommandRetrieveFullDetailedTestCaseExecution,
					TestCaseExecutionKey:                                              tempTestCaseExecutionMapKey,
					FullTestCaseExecutionResponseMessage:                              nil,
					TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage: nil,
				}

				// Don't put on Channel if more than 9 items from max capacity
				var currentChannelSize int32
				currentChannelSize = int32(len(DetailedExecutionStatusCommandChannel))
				if currentChannelSize > MessageChannelMaxSizeDetailedExecutionStatus-9 {
					for {
						time.Sleep(5 * time.Second)

						currentChannelSize = int32(len(DetailedExecutionStatusCommandChannel))
						if currentChannelSize < MessageChannelMaxSizeDetailedExecutionStatus-9 {
							break
						}
					}
				}

				// Send command ion channel
				DetailedExecutionStatusCommandChannel <- channelCommandDetailedExecutions

			}
		}
	}

	// Loop all TestInstructionStatus-messages and update TestCaseExecutionStatus for each TestCase
	for _, tempTestInstructionExecutionStatusMessage := range testCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage.
		TestInstructionExecutionsStatus {

		// Create TestCaseExecutionKey
		var tempTestCaseExecutionMapKey string
		tempTestCaseExecutionMapKey = tempTestInstructionExecutionStatusMessage.TestCaseExecutionUuid +
			strconv.Itoa(int(tempTestInstructionExecutionStatusMessage.TestCaseExecutionVersion))

		// Check if TestCaseExecution exist within the 'TestCaseExecutionsDetailsMap'
		var tempTestCaseExecutionsDetailsReference *detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsStruct
		tempTestCaseExecutionsDetailsReference, existInMap = detailedTestCaseExecutionUI_summaryTableDefinition.
			TestCaseExecutionsDetailsMap[tempTestCaseExecutionMapKey]

		// If not then something is really wrong
		if existInMap == false {
			ErrorID := "e69d14fd-c571-4f76-9bd5-dc3e303e967d"
			err := errors.New(fmt.Sprintf("We shouldn't end up here. [ErrorID:'%s']", ErrorID))

			fmt.Println(err) // TODO Send on Error-channel

			return
		}

		// Only Process TestInstructionExecutionStatus-message if there is no waiting for a Full TestCaseExecution-status-update
		if tempTestCaseExecutionsDetailsReference.WaitingForFullTestCaseExecutionUpdate == true ||
			tempTestCaseExecutionsDetailsReference.
				WaitingForFullTestCaseExecutionUpdateAfterFirstTestInstructionExecutionStatusWasReceived == true {

			continue
		}

		// Append Incoming TestCaseExecutionStatus-message into stored Map-message
		tempTestCaseExecutionsDetailsReference.TestInstructionExecutionsStatusUpdates = append(
			tempTestCaseExecutionsDetailsReference.TestInstructionExecutionsStatusUpdates, tempTestInstructionExecutionStatusMessage)

		// Extract TestInstructionExecutionsStatusMapKey
		var tempTestInstructionExecutionsStatusMapKey string
		tempTestInstructionExecutionsStatusMapKey = tempTestInstructionExecutionStatusMessage.TestInstructionExecutionUuid +
			strconv.Itoa(int(tempTestInstructionExecutionStatusMessage.TestInstructionExecutionVersion))

		// Check if TestInstructionExecution already exists within Map
		var tempTestTestInstructionExecutionsBaseInformation *detailedTestCaseExecutionUI_summaryTableDefinition.TestInstructionExecutionsBaseInformationStruct
		tempTestTestInstructionExecutionsBaseInformation, existInMap = tempTestCaseExecutionsDetailsReference.
			TestInstructionExecutionsStatusMap[tempTestInstructionExecutionsStatusMapKey]

		if existInMap == false {
			// TestInstructionExecution doesn't exist, then some is really wrong

			ErrorID := "d29c18aa-086d-4c45-bc06-724564ccc893"
			err := errors.New(fmt.Sprintf("We shouldn't end up here. [ErrorID:'%s']", ErrorID))

			fmt.Println(err) // TODO Send on Error-channel

			return

		}

		// Check if TestInstructionStatus exist in 'AllTestInstructionsExecutionsStatusUpdatesInformationMap'
		var tempExecutionStatusUpdateTimeStampMapKey time.Time
		var tempExecutionStatusUpdateTimeStampMapKeyAsString string

		// Create ExecutionStatusUpdateTimeStampMapKey
		tempExecutionStatusUpdateTimeStampMapKey = tempTestInstructionExecutionStatusMessage.
			TestInstructionExecutionsStatusInformation.ExecutionStatusUpdateTimeStamp.AsTime()

		tempExecutionStatusUpdateTimeStampMapKeyAsString = tempExecutionStatusUpdateTimeStampMapKey.String()

		// Verify if this UpdateTimeStamp exist within 'AllTestInstructionsExecutionsStatusUpdatesInformationMap'
		_, existInMap = tempTestTestInstructionExecutionsBaseInformation.
			AllTestInstructionsExecutionsStatusUpdatesInformationMap[tempExecutionStatusUpdateTimeStampMapKeyAsString]

		// If it doesn't exist then add it to the 'AllTestInstructionsExecutionsStatusUpdatesInformationMap'

		if existInMap == false {

			// Do check against the latest know UpdateTimeStamp regarding if new TimeStamp is later than existing latest one
			if tempTestTestInstructionExecutionsBaseInformation.LatestStatusUpdateTimeStamp.IsZero() ||
				tempExecutionStatusUpdateTimeStampMapKey.After(tempTestTestInstructionExecutionsBaseInformation.LatestStatusUpdateTimeStamp) {

				tempTestTestInstructionExecutionsBaseInformation.LatestStatusUpdateTimeStamp = tempExecutionStatusUpdateTimeStampMapKey
			}

			// Add ExecutionStatusUpdateTimeStamp to Status-Map
			tempTestTestInstructionExecutionsBaseInformation.
				AllTestInstructionsExecutionsStatusUpdatesInformationMap[tempExecutionStatusUpdateTimeStampMapKeyAsString] =
				tempTestInstructionExecutionStatusMessage.GetTestInstructionExecutionsStatusInformation()

			// Check that no status update-messages are missing
			// If there is a mismatch between locally store Previous Timestamp Status-message and
			// incoming Previous Timestamp Status-message then get full TestCaseExecution
			var tempestCaseExecutionsDetails *detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsStruct

			tempestCaseExecutionsDetails, existInTestCaseExecutionsDetailsMap = detailedTestCaseExecutionUI_summaryTableDefinition.
				TestCaseExecutionsDetailsMap[tempTestCaseExecutionMapKey]

			// Is there a mismatch between locally store Previous Timestamp Status-message and incoming Previous Timestamp Status-message
			// When BroadcastTimeStamp == PreviousBroadcastTimeStamp in incoming status message then that message is the first for that TestCaseExecution, for this ExecutionServerInstance
			if tempTestInstructionExecutionStatusMessage.BroadcastTimeStamp != tempTestInstructionExecutionStatusMessage.
				PreviousBroadcastTimeStamp &&
				tempTestInstructionExecutionStatusMessage.PreviousBroadcastTimeStamp.AsTime() != tempestCaseExecutionsDetails.
					PreviousBroadcastTimeStamp &&
				tempestCaseExecutionsDetails.PreviousBroadcastTimeStamp.IsZero() != true {

				// Create command to retrieve missing TestCaseExecutions, via channelEngine
				var channelCommandDetailedExecutions ChannelCommandDetailedExecutionsStruct
				channelCommandDetailedExecutions = ChannelCommandDetailedExecutionsStruct{
					ChannelCommandDetailedExecutionsStatus:                            ChannelCommandRetrieveFullDetailedTestCaseExecution,
					TestCaseExecutionKey:                                              tempTestCaseExecutionMapKey,
					FullTestCaseExecutionResponseMessage:                              nil,
					TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage: nil,
				}

				// Don't put on Channel if more than 9 items from max capacity
				var currentChannelSize int32
				currentChannelSize = int32(len(DetailedExecutionStatusCommandChannel))
				if currentChannelSize > MessageChannelMaxSizeDetailedExecutionStatus-9 {
					for {
						time.Sleep(5 * time.Second)

						currentChannelSize = int32(len(DetailedExecutionStatusCommandChannel))
						if currentChannelSize < MessageChannelMaxSizeDetailedExecutionStatus-9 {
							break
						}
					}
				}

				// Send command ion channel
				DetailedExecutionStatusCommandChannel <- channelCommandDetailedExecutions

			} else {
				// Save new BroadcastTimestamp as PreviousBroadcastTimestamp
				tempTestCaseExecutionsDetailsReference.PreviousBroadcastTimeStamp = tempTestInstructionExecutionStatusMessage.
					BroadcastTimeStamp.AsTime()

			}
		}
	}
	//****

	// Update the SummaryTable for TestInstructionExecutions
	detailedExecutionsModelObject.updateTestInstructionExecutionsSummaryTable()

	// Update the SummaryTable for TestCaseExecutions
	detailedExecutionsModelObject.updateTestCaseExecutionsSummaryTable()
}

// Check of the TestCaseExecutionStatus is classified as an EndStatus
func isTestCaseExecutionStatusAnEndStatus(testCaseExecutionStatus fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum) (isTestCaseEndStatus bool) {
	switch testCaseExecutionStatus {

	case fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum_TCE_CONTROLLED_INTERRUPTION,
		fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum_TCE_CONTROLLED_INTERRUPTION_CAN_BE_RERUN,
		fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum_TCE_FINISHED_OK,
		fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum_TCE_FINISHED_OK_CAN_BE_RERUN,
		fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum_TCE_FINISHED_NOT_OK,
		fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum_TCE_FINISHED_NOT_OK_CAN_BE_RERUN,
		fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum_TCE_UNEXPECTED_INTERRUPTION,
		fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum_TCE_UNEXPECTED_INTERRUPTION_CAN_BE_RERUN,
		fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum_TCE_TIMEOUT_INTERRUPTION,
		fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum_TCE_TIMEOUT_INTERRUPTION_CAN_BE_RERUN:

		isTestCaseEndStatus = true

	default:
		isTestCaseEndStatus = false
	}

	return isTestCaseEndStatus
}
