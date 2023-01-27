package detailedExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/executions/detailedTestCaseExecutionUI_summaryTableDefinition"
	"github.com/sirupsen/logrus"
	"sort"
	"strconv"
	"time"
)

// Update the SummaryTable for TestCaseExecutions
func (detailedExecutionsModelObject *DetailedExecutionsModelObjectStruct) updateTestCaseExecutionsSummaryTable() {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "5962f52e-c516-4a8c-a008-02a26ec181cc",
	}).Debug("Incoming 'updateTestCaseExecutionsSummaryTable'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "cb9e7b80-c78f-4236-a67c-8de4fced4443",
	}).Debug("Outgoing 'updateTestCaseExecutionsSummaryTable'")

	// Create new TestCaseExecutionsStatusForSummaryTable-variable
	var tempCaseExecutionsStatusForSummaryTable []*detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsStatusForSummaryTableStruct

	// Loop all TestCases
	for _, tempTestCaseExecutionsDetail := range detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap {

		var latestTimeStamp time.Time
		var latestTestCaseUpdateTimeStampMapKey string
		var firstTimeStampCheck bool

		// Loop all UpdateTimestamps and pick the one with latest UpdateTimestamp one and add to 'TestCaseExecutionsStatusForSummaryTable'
		for tempTestCaseUpdateTimeStampMapKey, tempTestCaseExecutionStatusUpdatesInformation := range tempTestCaseExecutionsDetail.TestCaseExecutionsBaseInformation.AllTestCaseExecutionsStatusUpdatesInformationMap {

			// When first then use that timestamp as baseline
			if firstTimeStampCheck == false {
				firstTimeStampCheck = true
				latestTimeStamp = tempTestCaseExecutionStatusUpdatesInformation.ExecutionStatusUpdateTimeStamp.AsTime()
				latestTestCaseUpdateTimeStampMapKey = tempTestCaseUpdateTimeStampMapKey

			} else {

				// Check if this Timestamp is later than current timestamp, if so use that as latest TimeStamp
				if tempTestCaseExecutionStatusUpdatesInformation.ExecutionStatusUpdateTimeStamp.AsTime().
					After(latestTimeStamp) {
					latestTimeStamp = tempTestCaseExecutionStatusUpdatesInformation.ExecutionStatusUpdateTimeStamp.AsTime()
					latestTestCaseUpdateTimeStampMapKey = tempTestCaseUpdateTimeStampMapKey
				}
			}
		}

		// Store StatusUpdate in SummaryTable
		if latestTestCaseUpdateTimeStampMapKey != "" {

			// Create a Sort Order for The TestCassExecutions based on
			// 1) PlacedOnTestExecutionQueueTimeStamp; 2) TestCaseUIName

			// Create TestCase-UI-name for SummaryTable: TestCaseName + part of UUID + VersionNumber
			var testCaseExecutionUIName string
			testCaseExecutionUIName = tempTestCaseExecutionsDetail.TestCaseExecutionsBaseInformation.
				TestCaseExecutionBasicInformation.TestCaseName + " [" +
				sharedCode.GenerateShortUuidFromFullUuid(tempTestCaseExecutionsDetail.
					TestCaseExecutionsBaseInformation.TestCaseExecutionBasicInformation.TestCaseExecutionUuid) +
				"..." +
				strconv.Itoa(int(tempTestCaseExecutionsDetail.TestCaseExecutionsBaseInformation.
					TestCaseExecutionBasicInformation.TestCaseExecutionVersion)) + "]"

			var sortOrderAsString string
			sortOrderAsString = tempTestCaseExecutionsDetail.TestCaseExecutionsBaseInformation.
				TestCaseExecutionBasicInformation.PlacedOnTestExecutionQueueTimeStamp.AsTime().String() +
				testCaseExecutionUIName

			var tempCaseExecutionStatusForSummaryTable *detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsStatusForSummaryTableStruct
			tempCaseExecutionStatusForSummaryTable = &detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsStatusForSummaryTableStruct{
				TestCaseUIName: testCaseExecutionUIName,
				TestCaseStatusValue: uint32(tempTestCaseExecutionsDetail.TestCaseExecutionsBaseInformation.
					AllTestCaseExecutionsStatusUpdatesInformationMap[latestTestCaseUpdateTimeStampMapKey].TestCaseExecutionStatus),
				ExecutionStatusUpdateTimeStamp: tempTestCaseExecutionsDetail.TestCaseExecutionsBaseInformation.
					AllTestCaseExecutionsStatusUpdatesInformationMap[latestTestCaseUpdateTimeStampMapKey].
					ExecutionStatusUpdateTimeStamp.AsTime(),
				TestCaseExecutionUuid: tempTestCaseExecutionsDetail.TestCaseExecutionsBaseInformation.
					TestCaseExecutionBasicInformation.TestCaseExecutionUuid,
				TestCaseExecutionVersion: tempTestCaseExecutionsDetail.TestCaseExecutionsBaseInformation.
					TestCaseExecutionBasicInformation.TestCaseExecutionVersion,
				SortOrder: sortOrderAsString,
				TestInstructionExecutionsStatusForSummaryTableReference: &tempTestCaseExecutionsDetail.TestInstructionExecutionsStatusForSummaryTable,
			}

			tempCaseExecutionsStatusForSummaryTable = append(tempCaseExecutionsStatusForSummaryTable,
				tempCaseExecutionStatusForSummaryTable)

		}
	}

	// Sort New version of TestCaseExecution-SummaryTable
	sort.Slice(tempCaseExecutionsStatusForSummaryTable, func(i, j int) bool {
		return tempCaseExecutionsStatusForSummaryTable[i].SortOrder < tempCaseExecutionsStatusForSummaryTable[j].SortOrder
	})

	// Save the New version of TestCaseExecution-SummaryTable to TestCase
	detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsStatusForSummaryTable = tempCaseExecutionsStatusForSummaryTable

}
