package detailedExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/executions/detailedTestCaseExecutionUI_summaryTableDefinition"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"sort"
	"strconv"
	"time"
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

		var TestTestInstructionExecutionsBaseInformationMap map[string]*detailedTestCaseExecutionUI_summaryTableDefinition.TestTestInstructionExecutionsBaseInformationStruct
		TestTestInstructionExecutionsBaseInformationMap = make(map[string]*detailedTestCaseExecutionUI_summaryTableDefinition.TestTestInstructionExecutionsBaseInformationStruct)

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
}

// Update the SummaryTable for TestInstructionExecutions
func (detailedExecutionsModelObject *DetailedExecutionsModelObjectStruct) updateTestInstructionExecutionsSummaryTable() {
	// Create new TestInstructionExecutionsSummaryTable-variable
	var tempTestInstructionExecutionsStatusForSummaryTable []*detailedTestCaseExecutionUI_summaryTableDefinition.TestInstructionExecutionsStatusForSummaryTableStruct

	// Loop all TestCases
	for _, tempTestCaseExecutionsDetail := range detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap {

		// Loop all TestInstructionExecutions
		for _, tempTestInstructionExecution := range tempTestCaseExecutionsDetail.TestInstructionExecutionsStatusMap {

			var latestTimeStamp time.Time
			var latestTestInstructionUpdateTimeStampMapKey string
			var firstTimeStampCheck bool
			// Loop all UpdateTimestamps and pick the last one and add to 'TestCaseExecutionsStatusForSummaryTable'
			for tempTestInstructionUpdateTimeStampMapKey, tempTestInstructionsExecutionsStatusUpdatesInformation := range tempTestInstructionExecution.AllTestInstructionsExecutionsStatusUpdatesInformationMap {

				// When first then use that timestamp as baseline
				if firstTimeStampCheck == false {
					firstTimeStampCheck = true
					latestTimeStamp = tempTestInstructionsExecutionsStatusUpdatesInformation.ExecutionStatusUpdateTimeStamp.AsTime()

				} else {
					// Check if this Timestamp is later than current timestamp, if so use that as latest TimeStamp
					if tempTestInstructionsExecutionsStatusUpdatesInformation.ExecutionStatusUpdateTimeStamp.AsTime().
						After(latestTimeStamp) {
						latestTimeStamp = tempTestInstructionsExecutionsStatusUpdatesInformation.ExecutionStatusUpdateTimeStamp.AsTime()
						latestTestInstructionUpdateTimeStampMapKey = tempTestInstructionUpdateTimeStampMapKey
					}
				}

				// Store StatusUpdate in SummaryTable
				if latestTestInstructionUpdateTimeStampMapKey != "" {

					// Create a Sort Order for The TestInstructionsExecution based on
					// 1) SortOrder; 2) TestInstructionExecutionUIName
					var sortOrderAsString string
					sortOrderAsString = strconv.Itoa(int(tempTestInstructionExecution.TestInstructionExecutionBasicInformation.TestInstructionExecutionOrder)) +
						tempTestInstructionExecution.TestInstructionExecutionBasicInformation.TestInstructionName

					// Create TestInstruction-UI-name for SummaryTable: TestInstructionName + part of UUID + VersionNumber
					var testInstructionExecutionUIName string
					testInstructionExecutionUIName = tempTestInstructionExecution.TestInstructionExecutionBasicInformation.TestInstructionName +
						"<" + sharedCode.GenerateShortUuidFromFullUuid(tempTestInstructionExecution.TestInstructionExecutionBasicInformation.TestInstructionExecutionUuid) +
						"..." + strconv.Itoa(int(tempTestInstructionExecution.TestInstructionExecutionBasicInformation.TestInstructionExecutionVersion)) +
						">"

					var tempTestInstructionExecutionForStatusForSummaryTable *detailedTestCaseExecutionUI_summaryTableDefinition.TestInstructionExecutionsStatusForSummaryTableStruct
					tempTestInstructionExecutionForStatusForSummaryTable = &detailedTestCaseExecutionUI_summaryTableDefinition.TestInstructionExecutionsStatusForSummaryTableStruct{
						TestInstructionExecutionUIName:  testInstructionExecutionUIName,
						TestInstructionName:             tempTestInstructionExecution.TestInstructionExecutionBasicInformation.TestInstructionName,
						TestInstructionExecutionUuid:    tempTestInstructionExecution.TestInstructionExecutionBasicInformation.TestInstructionExecutionUuid,
						TestInstructionExecutionVersion: tempTestInstructionExecution.TestInstructionExecutionBasicInformation.TestInstructionExecutionVersion,
						TestInstructionStatusValue:      uint32(tempTestInstructionExecution.AllTestInstructionsExecutionsStatusUpdatesInformationMap[latestTestInstructionUpdateTimeStampMapKey].TestInstructionExecutionStatus),
						ExecutionStatusUpdateTimeStamp:  tempTestInstructionExecution.AllTestInstructionsExecutionsStatusUpdatesInformationMap[latestTestInstructionUpdateTimeStampMapKey].ExecutionStatusUpdateTimeStamp.AsTime(),
						SortOrder:                       sortOrderAsString,
					}
					tempTestInstructionExecutionsStatusForSummaryTable = append(tempTestInstructionExecutionsStatusForSummaryTable,
						tempTestInstructionExecutionForStatusForSummaryTable)
				}
			}
		}

		// Sort New version of TestInstructionExecution-SummaryTable
		sort.Slice(tempTestInstructionExecutionsStatusForSummaryTable, func(i, j int) bool {
			return tempTestInstructionExecutionsStatusForSummaryTable[i].SortOrder < tempTestInstructionExecutionsStatusForSummaryTable[j].SortOrder
		})

		// Add New version of TestInstructionExecution-SummaryTable to TestCase
		tempTestCaseExecutionsDetail.TestInstructionExecutionsStatusForSummaryTable = tempTestInstructionExecutionsStatusForSummaryTable

	}
}
