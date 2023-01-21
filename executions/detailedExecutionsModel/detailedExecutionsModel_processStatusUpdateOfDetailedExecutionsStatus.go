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
			if tempTestCaseExecutionsDetailsMap.FullTestCaseExecutionUpdateWhenFirstStatusReceived == false {
				tempTestCaseExecutionsDetailsMap.FullTestCaseExecutionUpdateWhenFirstStatusReceived = true

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

	}

	/*

	   // TestCaseStatus
	   // Add to existing structure for the TestCase-Status on summary page
	   var testCaseExecutionsStatusForSummaryTableData TestCaseExecutionsStatusForSummaryTableStruct

	   // First get the latest data from DB-content
	   for testCaseExecutionDetailsCounter, testCaseExecutionDetailsMessage := range testCaseExecutionResponse.TestCaseExecutionDetails {

	   // When it's the first instance of status then use that as the base
	   if testCaseExecutionDetailsCounter == 0 {
	   testCaseExecutionsStatusForSummaryTableData = TestCaseExecutionsStatusForSummaryTableStruct{
	   TestCaseUIName:                 testCaseExecutionResponse.TestCaseExecutionBasicInformation.TestCaseName,
	   TestCaseStatusValue:            uint32(testCaseExecutionDetailsMessage.TestCaseExecutionStatus),
	   ExecutionStatusUpdateTimeStamp: testCaseExecutionDetailsMessage.ExecutionStatusUpdateTimeStamp.AsTime(),
	   }
	   } else {
	   // Check if the new timestamp > existing timestamp, if so then use new instance
	   if testCaseExecutionDetailsMessage.ExecutionStatusUpdateTimeStamp.AsTime().After(
	   testCaseExecutionsStatusForSummaryTableData.ExecutionStatusUpdateTimeStamp) {

	   testCaseExecutionsStatusForSummaryTableData = TestCaseExecutionsStatusForSummaryTableStruct{
	   TestCaseUIName:                 testCaseExecutionResponse.TestCaseExecutionBasicInformation.TestCaseName,
	   TestCaseStatusValue:            uint32(testCaseExecutionDetailsMessage.TestCaseExecutionStatus),
	   ExecutionStatusUpdateTimeStamp: testCaseExecutionDetailsMessage.ExecutionStatusUpdateTimeStamp.AsTime(),
	   }
	   }
	   }
	   }

	   // Second get the latest data from status updates
	   for _, testCaseExecutionDetailsMessage := range testCaseExecutionsDetails.TestCaseExecutionsStatusUpdates {

	   // Check if the new timestamp > existing timestamp, if so then use new instance
	   if testCaseExecutionDetailsMessage.TestCaseExecutionDetails.ExecutionStatusUpdateTimeStamp.AsTime().After(
	   testCaseExecutionsStatusForSummaryTableData.ExecutionStatusUpdateTimeStamp) {

	   testCaseExecutionsStatusForSummaryTableData = TestCaseExecutionsStatusForSummaryTableStruct{
	   TestCaseUIName:                 testCaseExecutionResponse.TestCaseExecutionBasicInformation.
	   TestCaseName,
	   TestCaseStatusValue:            uint32(testCaseExecutionDetailsMessage.TestCaseExecutionDetails.
	   TestCaseExecutionStatus),
	   ExecutionStatusUpdateTimeStamp: testCaseExecutionDetailsMessage.TestCaseExecutionDetails.
	   ExecutionStatusUpdateTimeStamp.AsTime(),
	   }
	   }
	   }

	   // Add the TestStatus for Summary page
	   testCaseExecutionsDetails.TestCaseExecutionsStatusForSummaryTable = testCaseExecutionsStatusForSummaryTableData

	   // TestInstructionsStatus

	   // First get the latest data from DB-content
	   for _, testInstructionExecutionDetailsMessage := range testCaseExecutionResponse.TestInstructionExecutions {
	   var testInstructionExecutionsStatusForSummaryTableData TestInstructionExecutionsStatusForSummaryTable
	   // Loop all status messages
	   for testInstructionExecutionInformationCounter, testInstructionExecutionInformation := range testInstructionExecutionDetailsMessage.
	   TestInstructionExecutionsInformation {

	   // When it's the first instance of status then use that as the base
	   if testInstructionExecutionInformationCounter == 0 {
	   testInstructionExecutionsStatusForSummaryTableData = TestInstructionExecutionsStatusForSummaryTable{
	   TestInstructionExecutionUIName: testInstructionExecutionDetailsMessage.TestInstructionExecutionBasicInformation.TestInstructionName,
	   TestInstructionStatusValue:     uint32(testInstructionExecutionInformation.TestInstructionExecutionStatus),
	   ExecutionStatusUpdateTimeStamp: testInstructionExecutionInformation.ExecutionStatusUpdateTimeStamp.AsTime(),
	   }
	   } else {
	   // Check if the new timestamp > existing timestamp, if so then use new instance
	   if testInstructionExecutionInformation.ExecutionStatusUpdateTimeStamp.AsTime().After(
	   testInstructionExecutionsStatusForSummaryTableData.ExecutionStatusUpdateTimeStamp) {

	   testInstructionExecutionsStatusForSummaryTableData = TestInstructionExecutionsStatusForSummaryTable{
	   TestInstructionExecutionUIName: testInstructionExecutionDetailsMessage.TestInstructionExecutionBasicInformation.TestInstructionName,
	   TestInstructionStatusValue: uint32(testInstructionExecutionInformation.TestInstructionExecutionStatus),
	   ExecutionStatusUpdateTimeStamp: testInstructionExecutionInformation.ExecutionStatusUpdateTimeStamp.AsTime(),
	   }
	   }
	   }
	   }

	   // Append the TestInstructionsStatus for Summary page
	   testCaseExecutionsDetails.TestInstructionExecutionsStatusForSummaryTable = append(
	   testCaseExecutionsDetails.TestInstructionExecutionsStatusForSummaryTable,
	   testInstructionExecutionsStatusForSummaryTableData)
	   }

	   hmm hur matchar man dessa två slice:ar?

	   }

	   }
	*/

}
