package detailedExecutionsModel

import (
	"FenixTesterGui/executions/detailedTestCaseExecutionUI_summaryTableDefinition"
	"errors"
	"fmt"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"strconv"
)

// Updates specific status information based on subscriptions updates from GuiExecutionServer
func (detailedExecutionsModelObject *DetailedExecutionsModelObjectStruct) processStatusUpdateOfDetailedExecutionsStatus(
	testCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage) {

	// Process TestCaseStatus-messages to check that all TestCases exist in 'detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap'
	var testCaseExecutionKeysMap map[string]string // map[tempTestCaseExecutionMapKey]tempTestCaseExecutionMapKey
	testCaseExecutionKeysMap = make(map[string]string)
	var existInMap bool
	var existInTestCaseExecutionsDetailsMap bool
	var existsInTestCaseExecutionKeysMap bool

	for _, tempTestCaseExecutionStatusMessage := range testCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage.TestCaseExecutionsStatus {

		// Create TestCaseExecutionKey
		var tempTestCaseExecutionToBeFullyRetrievedMapKey string
		tempTestCaseExecutionToBeFullyRetrievedMapKey = tempTestCaseExecutionStatusMessage.TestCaseExecutionUuid +
			strconv.Itoa(int(tempTestCaseExecutionStatusMessage.TestCaseExecutionVersion))

		// Check if TestCaseExecution exist within the 'TestCaseExecutionsDetailsMap'
		_, existInTestCaseExecutionsDetailsMap = detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap[tempTestCaseExecutionToBeFullyRetrievedMapKey]

		// If not then add it to the Map over TestCaseExecution to retrieve from the Database
		if existInTestCaseExecutionsDetailsMap == false {
			_, existsInTestCaseExecutionKeysMap = testCaseExecutionKeysMap[tempTestCaseExecutionToBeFullyRetrievedMapKey]
			// Has the tempTestCaseExecutionToBeFullyRetrievedMapKey already been saved
			if existsInTestCaseExecutionKeysMap == false {
				testCaseExecutionKeysMap[tempTestCaseExecutionToBeFullyRetrievedMapKey] = tempTestCaseExecutionToBeFullyRetrievedMapKey
			}
		}

		// If there is a mismatch between locally store Previous Timestamp Status-message and incoming Previous Timestamp Status-message then get full TestCaseExecution
		if existsInTestCaseExecutionKeysMap == false && existInTestCaseExecutionsDetailsMap == true {

			var tempestCaseExecutionsDetails *detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsStruct
			tempestCaseExecutionsDetails, existInTestCaseExecutionsDetailsMap = detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap[tempTestCaseExecutionToBeFullyRetrievedMapKey]

			// Is there a mismatch between locally store Previous Timestamp Status-message and incoming Previous Timestamp Status-message
			// When BroadcastTimeStamp == PreviousBroadcastTimeStamp in incoming status message then that message is the first for that TestCaseExecution, for this ExecutionServerInstance
			if tempTestCaseExecutionStatusMessage.BroadcastTimeStamp != tempTestCaseExecutionStatusMessage.PreviousBroadcastTimeStamp &&
				tempTestCaseExecutionStatusMessage.PreviousBroadcastTimeStamp.AsTime() != tempestCaseExecutionsDetails.PreviousBroadcastTimeStamp {

				// We have a mismatch, so retrieve full TestCaseExecution
				testCaseExecutionKeysMap[tempTestCaseExecutionToBeFullyRetrievedMapKey] = tempTestCaseExecutionToBeFullyRetrievedMapKey
			} else {

			}
		}

	}

	// Process TestInstructionStatus-messages to check that all TestCases exist in 'detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap'
	for _, tempTestInstructionExecutionStatusMessage := range testCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage.TestInstructionExecutionsStatus {

		// Create TestCaseExecutionKey
		var tempTestCaseExecutionToBeFullyRetrievedMapKey string
		tempTestCaseExecutionToBeFullyRetrievedMapKey = tempTestInstructionExecutionStatusMessage.TestCaseExecutionUuid +
			strconv.Itoa(int(tempTestInstructionExecutionStatusMessage.TestCaseExecutionVersion))

		// Check if TestCaseExecution exist within the 'TestCaseExecutionsDetailsMap'
		var tempTestCaseExecutionsDetailsMap *detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsStruct
		tempTestCaseExecutionsDetailsMap, existInTestCaseExecutionsDetailsMap = detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap[tempTestCaseExecutionToBeFullyRetrievedMapKey]

		// If not then add it to the Map over TestCaseExecution to retrieve from the Database
		if existInTestCaseExecutionsDetailsMap == false {

			_, existsInTestCaseExecutionKeysMap = testCaseExecutionKeysMap[tempTestCaseExecutionToBeFullyRetrievedMapKey]

			// Has the tempTestCaseExecutionToBeFullyRetrievedMapKey already been saved
			if existsInTestCaseExecutionKeysMap == false {
				testCaseExecutionKeysMap[tempTestCaseExecutionToBeFullyRetrievedMapKey] = tempTestCaseExecutionToBeFullyRetrievedMapKey

			}
		} else {
			// Has a first full TestCaseExecutionStatus been retrieved
			// Or this is the first TestInstructionExecutionUpdate for TestCaseExecution
			if tempTestCaseExecutionsDetailsMap.FullTestCaseExecutionUpdateWhenFirstExecutionStatusReceived == false ||
				tempTestCaseExecutionsDetailsMap.FullTestCaseExecutionUpdateWhenFirstTestInstructionExecutionStatusReceived == false &&
					len(testCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage.TestInstructionExecutionsStatus) > 0 {

				// Set flag that keeps track if first full TestCaseExecution is retrieved for first TestExecution-message
				tempTestCaseExecutionsDetailsMap.FullTestCaseExecutionUpdateWhenFirstExecutionStatusReceived = true

				// Set flag that keeps track if first full TestCaseExecution is retrieved for first TestInstructionsExecution-message
				if tempTestCaseExecutionsDetailsMap.FullTestCaseExecutionUpdateWhenFirstTestInstructionExecutionStatusReceived == false &&
					len(testCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage.TestInstructionExecutionsStatus) > 0 {
					tempTestCaseExecutionsDetailsMap.FullTestCaseExecutionUpdateWhenFirstTestInstructionExecutionStatusReceived = true
				}

				_, existsInTestCaseExecutionKeysMap = testCaseExecutionKeysMap[tempTestCaseExecutionToBeFullyRetrievedMapKey]

				// Has the tempTestCaseExecutionToBeFullyRetrievedMapKey already been saved
				if existsInTestCaseExecutionKeysMap == false {
					testCaseExecutionKeysMap[tempTestCaseExecutionToBeFullyRetrievedMapKey] = tempTestCaseExecutionToBeFullyRetrievedMapKey

				}
			}
		}

		// If there is a mismatch between locally store Previous Timestamp Status-message and incoming Previous Timestamp Status-message then get full TestCaseExecution
		if existsInTestCaseExecutionKeysMap == false && existInTestCaseExecutionsDetailsMap == true {

			var tempestCaseExecutionsDetails *detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsStruct
			tempestCaseExecutionsDetails, existInTestCaseExecutionsDetailsMap = detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap[tempTestCaseExecutionToBeFullyRetrievedMapKey]

			// Is there a mismatch between locally store Previous Timestamp Status-message and incoming Previous Timestamp Status-message
			// When BroadcastTimeStamp == PreviousBroadcastTimeStamp in incoming status message then that message is the first for that TestCaseExecution, for this ExecutionServerInstance
			if tempTestInstructionExecutionStatusMessage.BroadcastTimeStamp != tempTestInstructionExecutionStatusMessage.PreviousBroadcastTimeStamp &&
				tempTestInstructionExecutionStatusMessage.PreviousBroadcastTimeStamp.AsTime() != tempestCaseExecutionsDetails.PreviousBroadcastTimeStamp {

				// We have a mismatch, so retrieve full TestCaseExecution
				testCaseExecutionKeysMap[tempTestCaseExecutionToBeFullyRetrievedMapKey] = tempTestCaseExecutionToBeFullyRetrievedMapKey
			} else {

			}
		}
	}

	// If there are ony TestCaseExecution that is not already within the TestCaseExecutionsMap then get them first before process the updates
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

			// Send command ion channel
			DetailedExecutionStatusCommandChannel <- channelCommandDetailedExecutions
		}

		// Resend the original message that was processed
		var channelCommandDetailedExecutions ChannelCommandDetailedExecutionsStruct
		channelCommandDetailedExecutions = ChannelCommandDetailedExecutionsStruct{
			ChannelCommandDetailedExecutionsStatus:                            ChannelCommandStatusUpdateOfDetailedExecutionsStatus,
			TestCaseExecutionKey:                                              "",
			FullTestCaseExecutionResponseMessage:                              nil,
			TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage: testCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage,
		}
		// Send command on channel
		DetailedExecutionStatusCommandChannel <- channelCommandDetailedExecutions

		// Exit
		return
	}

	// All TestCaseExecutions exist within 'TestCaseExecutionsDetailsMap' so start process the status updates

	// Loop all TestCaseStatus-messages and update TestCaseExecutionStatus for each TestCase
	for _, tempTestCaseExecutionStatusMessage := range testCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage.TestCaseExecutionsStatus {

		// Create TestCaseExecutionKey
		var tempTestCaseExecutionMapKey string
		tempTestCaseExecutionMapKey = tempTestCaseExecutionStatusMessage.TestCaseExecutionUuid +
			strconv.Itoa(int(tempTestCaseExecutionStatusMessage.TestCaseExecutionVersion))

		// Check if TestCaseExecution exist within the 'TestCaseExecutionsDetailsMap'
		var tempTestCaseExecutionsDetailsReference *detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsStruct
		tempTestCaseExecutionsDetailsReference, existInTestCaseExecutionsDetailsMap = detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap[tempTestCaseExecutionMapKey]

		// If not then something is really wrong
		if existInTestCaseExecutionsDetailsMap == false {
			ErrorID := "b33739a9-e5c4-452c-891e-008b8c1a8a1d"
			err := errors.New(fmt.Sprintf("We shouldn't end up here. [ErrorID:'%s']", ErrorID))

			fmt.Println(err) // TODO Send on Error-channel

			return
		}

		// Save new BroadcastTimestamp as PreviousBroadcastTimestamp
		tempTestCaseExecutionsDetailsReference.PreviousBroadcastTimeStamp = tempTestCaseExecutionStatusMessage.BroadcastTimeStamp.AsTime()

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
		}

	}

	// Loop all TestInstructionStatus-messages and update TestCaseExecutionStatus for each TestCase
	for _, tempTestInstructionExecutionStatusMessage := range testCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage.TestInstructionExecutionsStatus {

		// Create TestCaseExecutionKey
		var tempTestCaseExecutionMapKey string
		tempTestCaseExecutionMapKey = tempTestInstructionExecutionStatusMessage.TestCaseExecutionUuid +
			strconv.Itoa(int(tempTestInstructionExecutionStatusMessage.TestCaseExecutionVersion))

		// Check if TestCaseExecution exist within the 'TestCaseExecutionsDetailsMap'
		var tempTestCaseExecutionsDetailsReference *detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsStruct
		tempTestCaseExecutionsDetailsReference, existInMap = detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap[tempTestCaseExecutionMapKey]

		// If not then something is really wrong
		if existInMap == false {
			ErrorID := "e69d14fd-c571-4f76-9bd5-dc3e303e967d"
			err := errors.New(fmt.Sprintf("We shouldn't end up here. [ErrorID:'%s']", ErrorID))

			fmt.Println(err) // TODO Send on Error-channel

			return
		}

		// Save new BroadcastTimestamp as PreviousBroadcastTimestamp
		tempTestCaseExecutionsDetailsReference.PreviousBroadcastTimeStamp = tempTestInstructionExecutionStatusMessage.BroadcastTimeStamp.AsTime()

		// Append Incoming TestCaseExecutionStatus-message into stored Map-message
		tempTestCaseExecutionsDetailsReference.TestInstructionExecutionsStatusUpdates = append(
			tempTestCaseExecutionsDetailsReference.TestInstructionExecutionsStatusUpdates, tempTestInstructionExecutionStatusMessage)

		//*****

		var tempTestInstructionExecutionsStatusMapKey string
		tempTestInstructionExecutionsStatusMapKey = tempTestInstructionExecutionStatusMessage.TestInstructionExecutionUuid +
			strconv.Itoa(int(tempTestInstructionExecutionStatusMessage.TestInstructionExecutionVersion))

		// Check if TestInstructionExecution already exists within Map
		var tempTestTestInstructionExecutionsBaseInformation *detailedTestCaseExecutionUI_summaryTableDefinition.TestInstructionExecutionsBaseInformationStruct
		tempTestTestInstructionExecutionsBaseInformation, existInMap = tempTestCaseExecutionsDetailsReference.TestInstructionExecutionsStatusMap[tempTestInstructionExecutionsStatusMapKey]

		if existInMap == false {
			// TestInstructionExecution doesn't exist, then some is really wrong

			ErrorID := "d29c18aa-086d-4c45-bc06-724564ccc893"
			err := errors.New(fmt.Sprintf("We shouldn't end up here. [ErrorID:'%s']", ErrorID))

			fmt.Println(err) // TODO Send on Error-channel

			return

		}

		// Check if TestInstructionStatus exist in 'AllTestInstructionsExecutionsStatusUpdatesInformationMap'
		var tempExecutionStatusUpdateTimeStampMapKey string
		tempExecutionStatusUpdateTimeStampMapKey = tempTestInstructionExecutionStatusMessage.TestInstructionExecutionsStatusInformation.ExecutionStatusUpdateTimeStamp.AsTime().String()

		// Verify if this UpdateTimeStamp exist within 'AllTestInstructionsExecutionsStatusUpdatesInformationMap'
		_, existInMap = tempTestTestInstructionExecutionsBaseInformation.AllTestInstructionsExecutionsStatusUpdatesInformationMap[tempExecutionStatusUpdateTimeStampMapKey]

		// If it doesn't exist then add it to the 'AllTestInstructionsExecutionsStatusUpdatesInformationMap'
		if existInMap == false {
			tempTestTestInstructionExecutionsBaseInformation.AllTestInstructionsExecutionsStatusUpdatesInformationMap[tempExecutionStatusUpdateTimeStampMapKey] = tempTestInstructionExecutionStatusMessage.GetTestInstructionExecutionsStatusInformation()

		}
	}
	//****

	// Update the SummaryTable for TestInstructionExecutions
	detailedExecutionsModelObject.updateTestInstructionExecutionsSummaryTable()

	// Update the SummaryTable for TestCaseExecutions
	detailedExecutionsModelObject.updateTestCaseExecutionsSummaryTable()
}
