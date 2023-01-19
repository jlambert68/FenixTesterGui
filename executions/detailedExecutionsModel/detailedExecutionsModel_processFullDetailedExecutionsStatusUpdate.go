package detailedExecutionsModel

import (
	"FenixTesterGui/executions/detailedTestCaseExecutionUI_summaryTableDefinition"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"strconv"
)

// Updates all Executions status with information received after direct gRPC-call to GUiExecutionServer
func (detailedExecutionsModelObject *DetailedExecutionsModelObjectStruct) processFullDetailedTestCaseExecutionsStatusUpdate(
	testCaseExecutionResponse *fenixExecutionServerGuiGrpcApi.TestCaseExecutionResponseMessage) {

	// Create the TestCaseExecutionMapJey
	var testCaseExecutionMapKey string
	testCaseExecutionMapKey = testCaseExecutionResponse.TestCaseExecutionBasicInformation.TestCaseExecutionUuid +
		strconv.Itoa(int(testCaseExecutionResponse.TestCaseExecutionBasicInformation.TestCaseExecutionVersion))

	// Check if TestCaseExecution already exist
	var existInMap bool
	var testCaseExecutionsDetails *detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsStruct
	testCaseExecutionsDetails, existInMap = detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap[testCaseExecutionMapKey]

	// If TestExecutionExecution doesn't exist in map then create a new instance
	if existInMap == false {

		var testInstructionExecutionsStatusMap map[string]*[]*fenixExecutionServerGuiGrpcApi.TestInstructionExecutionsInformationMessage
		testInstructionExecutionsStatusMap = make(map[string]*[]*fenixExecutionServerGuiGrpcApi.TestInstructionExecutionsInformationMessage)

		testCaseExecutionsDetails = &detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsStruct{
			TestCaseExecutionDatabaseResponseMessage:       testCaseExecutionResponse,
			TestCaseExecutionsStatusUpdates:                nil,
			TestInstructionExecutionsStatusUpdates:         nil,
			TestInstructionExecutionsStatusMap:             testInstructionExecutionsStatusMap,
			TestCaseExecutionsStatusForSummaryTable:        &detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsStatusForSummaryTableStruct{},
			TestInstructionExecutionsStatusForSummaryTable: nil,
		}

		// Add the TestCaseExecution to the Map
		detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap[testCaseExecutionMapKey] = testCaseExecutionsDetails
	}

	// TestCaseStatus
	// Add the TestCase Status for summary page to the Map by converting into the simpler summary page structure
	var testCaseExecutionsStatusForSummaryTableData detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsStatusForSummaryTableStruct
	for testCaseExecutionDetailsCounter, testCaseExecutionDetailsMessage := range testCaseExecutionResponse.TestCaseExecutionDetails {

		// When it's the first instance of status then use that as the base
		if testCaseExecutionDetailsCounter == 0 {
			testCaseExecutionsStatusForSummaryTableData = detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsStatusForSummaryTableStruct{
				TestCaseUIName:                 testCaseExecutionResponse.TestCaseExecutionBasicInformation.TestCaseName,
				TestCaseStatusValue:            uint32(testCaseExecutionDetailsMessage.TestCaseExecutionStatus),
				ExecutionStatusUpdateTimeStamp: testCaseExecutionDetailsMessage.ExecutionStatusUpdateTimeStamp.AsTime(),
				TestCaseExecutionUuid:          testCaseExecutionResponse.TestCaseExecutionBasicInformation.TestCaseExecutionUuid,
				TestCaseExecutionVersion:       strconv.Itoa(int(testCaseExecutionResponse.TestCaseExecutionBasicInformation.TestCaseExecutionVersion)),
			}
		} else {
			// Check if the new timestamp > existing timestamp, if so then use new instance
			if testCaseExecutionDetailsMessage.ExecutionStatusUpdateTimeStamp.AsTime().After(
				testCaseExecutionsStatusForSummaryTableData.ExecutionStatusUpdateTimeStamp) {

				testCaseExecutionsStatusForSummaryTableData = detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsStatusForSummaryTableStruct{
					TestCaseUIName:                 testCaseExecutionResponse.TestCaseExecutionBasicInformation.TestCaseName,
					TestCaseStatusValue:            uint32(testCaseExecutionDetailsMessage.TestCaseExecutionStatus),
					ExecutionStatusUpdateTimeStamp: testCaseExecutionDetailsMessage.ExecutionStatusUpdateTimeStamp.AsTime(),
					TestCaseExecutionUuid:          testCaseExecutionResponse.TestCaseExecutionBasicInformation.TestCaseExecutionUuid,
					TestCaseExecutionVersion:       strconv.Itoa(int(testCaseExecutionResponse.TestCaseExecutionBasicInformation.TestCaseExecutionVersion)),
				}
			}
		}
	}
	// Add the TestStatus for Summary page
	testCaseExecutionsDetails.TestCaseExecutionsStatusForSummaryTable = &testCaseExecutionsStatusForSummaryTableData

	// TestInstructionsStatus
	// Add the TestInstructions Statuses for summary page to the Map by converting into the simpler summary page structure
	for _, testInstructionExecutionDetailsMessage := range testCaseExecutionResponse.TestInstructionExecutions {
		var testInstructionExecutionsStatusForSummaryTableData detailedTestCaseExecutionUI_summaryTableDefinition.TestInstructionExecutionsStatusForSummaryTableStruct
		// Loop all status messages
		for testInstructionExecutionInformationCounter, testInstructionExecutionInformation := range testInstructionExecutionDetailsMessage.TestInstructionExecutionsInformation {

			// When it's the first instance of status then use that as the base
			if testInstructionExecutionInformationCounter == 0 {
				testInstructionExecutionsStatusForSummaryTableData = detailedTestCaseExecutionUI_summaryTableDefinition.TestInstructionExecutionsStatusForSummaryTableStruct{
					TestInstructionExecutionUIName: testInstructionExecutionDetailsMessage.TestInstructionExecutionBasicInformation.TestInstructionName,
					TestInstructionStatusValue:     uint32(testInstructionExecutionInformation.TestInstructionExecutionStatus),
					ExecutionStatusUpdateTimeStamp: testInstructionExecutionInformation.ExecutionStatusUpdateTimeStamp.AsTime(),
				}
			} else {
				// Check if the new timestamp > existing timestamp, if so then use new instance
				if testInstructionExecutionInformation.ExecutionStatusUpdateTimeStamp.AsTime().After(
					testInstructionExecutionsStatusForSummaryTableData.ExecutionStatusUpdateTimeStamp) {

					testInstructionExecutionsStatusForSummaryTableData = detailedTestCaseExecutionUI_summaryTableDefinition.TestInstructionExecutionsStatusForSummaryTableStruct{
						TestInstructionExecutionUIName: testInstructionExecutionDetailsMessage.TestInstructionExecutionBasicInformation.TestInstructionName,
						TestInstructionStatusValue:     uint32(testInstructionExecutionInformation.TestInstructionExecutionStatus),
						ExecutionStatusUpdateTimeStamp: testInstructionExecutionInformation.ExecutionStatusUpdateTimeStamp.AsTime(),
					}
				}
			}
		}
		// Append the TestInstructionsStatus for Summary page
		testCaseExecutionsDetails.TestInstructionExecutionsStatusForSummaryTable = append(
			testCaseExecutionsDetails.TestInstructionExecutionsStatusForSummaryTable,
			&testInstructionExecutionsStatusForSummaryTableData)
	}

	// Add reference for TestInstructionsStatus for Summary page to TestCaseStatus for Summary page
	testCaseExecutionsDetails.TestCaseExecutionsStatusForSummaryTable.TestInstructionExecutionsStatusForSummaryTable =
		&testCaseExecutionsDetails.TestInstructionExecutionsStatusForSummaryTable

	// Add reference for 'TestCaseExecutionsDetailsMap' to SummaryTableOpts
	//DetailedTestCaseExecutionsSummaryTableOptions.TestCaseExecutionsDetailsMapReference =
	//	&detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap

}
