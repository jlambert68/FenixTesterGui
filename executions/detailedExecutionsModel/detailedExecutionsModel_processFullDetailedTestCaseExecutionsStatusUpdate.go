package detailedExecutionsModel

import (
	"FenixTesterGui/executions/detailedTestCaseExecutionUI_summaryTableDefinition"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"strconv"
)

// Updates all Executions status with information received after direct gRPC-call to GUiExecutionServer
func (detailedExecutionsModelObject *DetailedExecutionsModelObjectStruct) processFullDetailedTestCaseExecutionsStatusUpdate(
	testCaseExecutionResponse *fenixExecutionServerGuiGrpcApi.TestCaseExecutionResponseMessage) {

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
			TestCaseExecutionDatabaseResponseMessage:       testCaseExecutionResponse,
			TestCaseExecutionsStatusUpdates:                nil,
			TestInstructionExecutionsStatusUpdates:         nil,
			TestCaseExecutionsBaseInformation:              tempTestCaseExecutionsBaseInformation,
			TestInstructionExecutionsStatusMap:             TestTestInstructionExecutionsBaseInformationMap,
			TestInstructionExecutionsStatusForSummaryTable: nil,
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

	// Update the SummaryTable for TestInstructionExecutions
	detailedExecutionsModelObject.updateTestInstructionExecutionsSummaryTable()

	// Update the SummaryTable for TestCaseExecutions
	detailedExecutionsModelObject.updateTestCaseExecutionsSummaryTable()
}
