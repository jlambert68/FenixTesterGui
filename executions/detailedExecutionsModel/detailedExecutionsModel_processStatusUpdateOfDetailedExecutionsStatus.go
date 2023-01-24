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
	for _, tempTestCaseExecutionStatusMessage := range testCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage.TestCaseExecutionsStatus {

		// Create TestCaseExecutionKey
		var tempTestCaseExecutionMapKey string
		tempTestCaseExecutionMapKey = tempTestCaseExecutionStatusMessage.TestCaseExecutionUuid +
			strconv.Itoa(int(tempTestCaseExecutionStatusMessage.TestCaseExecutionVersion))

		// Check if TestCaseExecution exist within the 'TestCaseExecutionsDetailsMap'
		_, existInMap = detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap[tempTestCaseExecutionMapKey]

		// If not then add it to the Map over TestCaseExecution to retrieve from the Database
		if existInMap == false {
			_, existInMap = testCaseExecutionKeysMap[tempTestCaseExecutionMapKey]
			// Has the tempTestCaseExecutionMapKey already been saved
			if existInMap == false {
				testCaseExecutionKeysMap[tempTestCaseExecutionMapKey] = tempTestCaseExecutionMapKey
			}
		}
	}

	// Process TestInstructionStatus-messages to check that all TestCases exist in 'detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap'
	for _, tempTestInstructionExecutionStatusMessage := range testCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage.TestInstructionExecutionsStatus {

		// Create TestCaseExecutionKey
		var tempTestCaseExecutionMapKey string
		tempTestCaseExecutionMapKey = tempTestInstructionExecutionStatusMessage.TestCaseExecutionUuid +
			strconv.Itoa(int(tempTestInstructionExecutionStatusMessage.TestCaseExecutionVersion))

		// Check if TestCaseExecution exist within the 'TestCaseExecutionsDetailsMap'
		var tempTestCaseExecutionsDetailsMap *detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsStruct
		tempTestCaseExecutionsDetailsMap, existInMap = detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap[tempTestCaseExecutionMapKey]

		// If not then add it to the Map over TestCaseExecution to retrieve from the Database
		if existInMap == false {

			_, existInMap = testCaseExecutionKeysMap[tempTestCaseExecutionMapKey]

			// Has the tempTestCaseExecutionMapKey already been saved
			if existInMap == false {
				testCaseExecutionKeysMap[tempTestCaseExecutionMapKey] = tempTestCaseExecutionMapKey

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

				_, existInMap = testCaseExecutionKeysMap[tempTestCaseExecutionMapKey]

				// Has the tempTestCaseExecutionMapKey already been saved
				if existInMap == false {
					testCaseExecutionKeysMap[tempTestCaseExecutionMapKey] = tempTestCaseExecutionMapKey

				}
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
		tempTestCaseExecutionsDetailsReference, existInMap = detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap[tempTestCaseExecutionMapKey]

		// If not then something is really wrong
		if existInMap == false {
			ErrorID := "b33739a9-e5c4-452c-891e-008b8c1a8a1d"
			err := errors.New(fmt.Sprintf("We shouldn't end up here. [ErrorID:'%s']", ErrorID))

			fmt.Println(err) // TODO Send on Error-channel

			return
		}

		// Append Incoming TestCaseExecutionStatus-message into stored Map-message
		tempTestCaseExecutionsDetailsReference.TestCaseExecutionsStatusUpdates = append(
			tempTestCaseExecutionsDetailsReference.TestCaseExecutionsStatusUpdates, tempTestCaseExecutionStatusMessage)

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

		// Append Incoming TestCaseExecutionStatus-message into stored Map-message
		tempTestCaseExecutionsDetailsReference.TestInstructionExecutionsStatusUpdates = append(
			tempTestCaseExecutionsDetailsReference.TestInstructionExecutionsStatusUpdates, tempTestInstructionExecutionStatusMessage)

		//*****

		var tempTestInstructionExecutionsStatusMapKey string
		tempTestInstructionExecutionsStatusMapKey = tempTestInstructionExecutionStatusMessage.TestInstructionExecutionUuid +
			strconv.Itoa(int(tempTestInstructionExecutionStatusMessage.TestInstructionExecutionVersion))

		// Check if TestInstructionExecution already exists within Map
		var tempTestTestInstructionExecutionsBaseInformation *detailedTestCaseExecutionUI_summaryTableDefinition.TestTestInstructionExecutionsBaseInformationStruct
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
}
