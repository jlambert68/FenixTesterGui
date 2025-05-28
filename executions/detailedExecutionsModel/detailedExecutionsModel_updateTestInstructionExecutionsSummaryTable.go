package detailedExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/executions/detailedTestCaseExecutionUI_summaryTableDefinition"
	"github.com/sirupsen/logrus"
	"sort"
	"strconv"
	"time"
)

// Update the SummaryTable for TestInstructionExecutions
func (detailedExecutionsModelObject *DetailedExecutionsModelObjectStruct) updateTestInstructionExecutionsSummaryTable() {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "279ac1b9-1c69-441a-816e-d95c5a5926c0",
	}).Debug("Incoming 'updateTestInstructionExecutionsSummaryTable'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "a94931cd-117f-44ce-8e1b-792f2f3fa38c",
	}).Debug("Outgoing 'updateTestInstructionExecutionsSummaryTable'")

	// Loop all TestCasesMap
	for _, tempTestCaseExecutionsDetail := range detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap {

		// Create new TestInstructionExecutionsSummaryTable-variable for every TestCase
		var tempTestInstructionExecutionsStatusForSummaryTable []*detailedTestCaseExecutionUI_summaryTableDefinition.TestInstructionExecutionsStatusForSummaryTableStruct

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
					latestTestInstructionUpdateTimeStampMapKey = tempTestInstructionUpdateTimeStampMapKey

				} else {

					// Check if this Timestamp is later than current timestamp, if so use that as latest TimeStamp
					if tempTestInstructionsExecutionsStatusUpdatesInformation.ExecutionStatusUpdateTimeStamp.AsTime().
						After(latestTimeStamp) {
						latestTimeStamp = tempTestInstructionsExecutionsStatusUpdatesInformation.ExecutionStatusUpdateTimeStamp.AsTime()
						latestTestInstructionUpdateTimeStampMapKey = tempTestInstructionUpdateTimeStampMapKey
					}
				}
			}

			// Store StatusUpdate in SummaryTable
			if latestTestInstructionUpdateTimeStampMapKey != "" {

				// Create a Sort Order for The TestInstructionsExecution based on
				// 1) SortOrder; 2) TestInstructionExecutionUIName

				// Create TestInstruction-UI-name for SummaryTable: TestInstructionName + part of UUID + VersionNumber
				var testInstructionExecutionUIName string
				testInstructionExecutionUIName = tempTestInstructionExecution.TestInstructionExecutionBasicInformation.TestInstructionName +
					" [" + sharedCode.GenerateShortUuidFromFullUuid(tempTestInstructionExecution.TestInstructionExecutionBasicInformation.TestInstructionExecutionUuid) +
					"..." + strconv.Itoa(int(tempTestInstructionExecution.TestInstructionExecutionBasicInformation.TestInstructionExecutionVersion)) +
					"]"

				var sortOrderAsString string
				sortOrderAsString = strconv.Itoa(int(tempTestInstructionExecution.TestInstructionExecutionBasicInformation.TestInstructionExecutionOrder)) +
					testInstructionExecutionUIName

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

		// Sort New version of TestInstructionExecution-SummaryTable
		sort.Slice(tempTestInstructionExecutionsStatusForSummaryTable, func(i, j int) bool {
			return tempTestInstructionExecutionsStatusForSummaryTable[i].SortOrder < tempTestInstructionExecutionsStatusForSummaryTable[j].SortOrder
		})

		// Add New version of TestInstructionExecution-SummaryTable to TestCase
		tempTestCaseExecutionsDetail.TestInstructionExecutionsStatusForSummaryTable = tempTestInstructionExecutionsStatusForSummaryTable

	}
}
