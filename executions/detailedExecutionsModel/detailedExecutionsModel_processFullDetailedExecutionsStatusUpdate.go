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

		var TestTestInstructionExecutionsBaseInformationMap map[string]*[]*detailedTestCaseExecutionUI_summaryTableDefinition.TestTestInstructionExecutionsBaseInformationStruct
		TestTestInstructionExecutionsBaseInformationMap = make(map[string]*[]*detailedTestCaseExecutionUI_summaryTableDefinition.TestTestInstructionExecutionsBaseInformationStruct)

		testCaseExecutionsDetails = &detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsStruct{
			TestCaseExecutionDatabaseResponseMessage:       testCaseExecutionResponse,
			TestCaseExecutionsStatusUpdates:                nil,
			TestInstructionExecutionsStatusUpdates:         nil,
			TestInstructionExecutionsStatusMap:             TestTestInstructionExecutionsBaseInformationMap,
			TestCaseExecutionsStatusForSummaryTable:        &detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsStatusForSummaryTableStruct{},
			TestInstructionExecutionsStatusForSummaryTable: nil,
		}

		// Add the TestCaseExecution to the Map
		detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap[testCaseExecutionMapKey] = testCaseExecutionsDetails
	} else {
		// Replace TestCaseExecutionResponse from Database
		testCaseExecutionsDetails.TestCaseExecutionDatabaseResponseMessage = testCaseExecutionResponse
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
	// Add the TestInstructions Statuses for summary page to the Map by converting into structure
	for _, testInstructionExecutionDetailsMessage := range testCaseExecutionResponse.TestInstructionExecutions {

		var tempTestInstructionExecutionsStatusMapKey string
		tempTestInstructionExecutionsStatusMapKey = testInstructionExecutionDetailsMessage.TestInstructionExecutionBasicInformation.TestInstructionExecutionUuid +
			strconv.Itoa(int(testInstructionExecutionDetailsMessage.TestInstructionExecutionBasicInformation.TestInstructionExecutionVersion))

		// Check if TestInstructionExecution already exists within Map
		var tempTestTestInstructionExecutionsBaseInformation *detailedTestCaseExecutionUI_summaryTableDefinition.TestTestInstructionExecutionsBaseInformationStruct
		tempTestTestInstructionExecutionsBaseInformation, existInMap = testCaseExecutionsDetails.TestInstructionExecutionsStatusMap[tempTestInstructionExecutionsStatusMapKey]

		if existInMap == false {
			// TestInstructionExecution doesn't exist, so create the TestInstructionExecution-object

			var tempAllTestInstructionsExecutionsStatusUpdatesInformationMap map[string]*fenixExecutionServerGuiGrpcApi.TestInstructionExecutionsInformationMessage
			tempAllTestInstructionsExecutionsStatusUpdatesInformationMap = make(map[string]*fenixExecutionServerGuiGrpcApi.TestInstructionExecutionsInformationMessage)

			tempTestTestInstructionExecutionsBaseInformation = &detailedTestCaseExecutionUI_summaryTableDefinition.TestTestInstructionExecutionsBaseInformationStruct{
				TestInstructionExecutionBasicInformation:                 testInstructionExecutionDetailsMessage.TestInstructionExecutionBasicInformation,
				AllTestInstructionsExecutionsStatusUpdatesInformationMap: tempAllTestInstructionsExecutionsStatusUpdatesInformationMap,
				CurrentTestInstructionExecutionsStatusForSummaryTable:    nil,
			}

			// Save TestInstructionExecution-object back into map
			testCaseExecutionsDetails.TestInstructionExecutionsStatusMap[tempTestInstructionExecutionsStatusMapKey] = tempTestTestInstructionExecutionsBaseInformation

		}

		// Check if TestInstructionStatus exist in 'AllTestInstructionsExecutionsStatusUpdatesInformationMap'
		// Loop all TestInstructionStatus-messages
		for _, tempTestInstructionExecutionsInformation := range testInstructionExecutionDetailsMessage.TestInstructionExecutionsInformation {
			var  tempExecutionStatusUpdateTimeStampMapKey string
			tempExecutionStatusUpdateTimeStampMapKey = tempTestInstructionExecutionsInformation.ExecutionStatusUpdateTimeStamp.AsTime().String()

			// Verify if this UpdateTimeStamp exist within 'AllTestInstructionsExecutionsStatusUpdatesInformationMap'
			_, existInMap = tempTestTestInstructionExecutionsBaseInformation.AllTestInstructionsExecutionsStatusUpdatesInformationMap[tempExecutionStatusUpdateTimeStampMapKey]

			// If it doesn't exist then add it to the 'AllTestInstructionsExecutionsStatusUpdatesInformationMap'
			if existInMap == false {
				tempTestTestInstructionExecutionsBaseInformation.AllTestInstructionsExecutionsStatusUpdatesInformationMap[tempExecutionStatusUpdateTimeStampMapKey] = tempTestInstructionExecutionsInformation

			}
		}

		// Loop all UpdateTimestamps and pick the last one and add to 'TestCaseExecutionsStatusForSummaryTable'
		for



		var tempTestInstructionExecutionsInformationMessages []*fenixExecutionServerGuiGrpcApi.TestInstructionExecutionsInformationMessage
		tempTestInstructionExecutionsInformationMessages = append(tempTestInstructionExecutionsInformationMessages, testInstructionExecutionDetailsMessage.TestInstructionExecutionsInformation...)

		// Initiate Map with TestInstructionStatusMessages with UpdateTimeStamp as Map-key
		if testInstructionExecutionDetailsMessage.
		var  tempAllTestInstructionsExecutionsStatusUpdatesInformationMap map[string]*fenixExecutionServerGuiGrpcApi.TestInstructionExecutionsInformationMessage
		tempAllTestInstructionsExecutionsStatusUpdatesInformationMap map[string]*fenixExecutionServerGuiGrpcApi.TestInstructionExecutionsInformationMessage

		tempTestTestInstructionExecutionsBaseInformation = &detailedTestCaseExecutionUI_summaryTableDefinition.TestTestInstructionExecutionsBaseInformationStruct{
			TestInstructionExecutionBasicInformation:              testInstructionExecutionDetailsMessage.TestInstructionExecutionBasicInformation,
			AllTestInstructiosExecutionsStatusUpdatesInformation:  tempTestInstructionExecutionsInformationMessages,
			CurrentTestInstructionExecutionsStatusForSummaryTable: nil,
		}

		// Loop

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
