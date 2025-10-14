package testSuiteExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	"github.com/sirupsen/logrus"
	"sort"
)

// ListExecutionDataForTestInstructionExecutions
// List all Execution-data for supplied TestInstructionExecutions. Posts are sorted on UpdateTimeStamp
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) ListExecutionDataForTestInstructionExecutions(
	detailedTestSuiteExecutionMapKey DetailedTestSuiteExecutionMapKeyType,
	testInstructionExecutionDetailsMapKeys []TestInstructionExecutionDetailsMapKeyType) (
	testInstructionExecutionDetailsForExplorerPtr *[]*TestInstructionExecutionDetailsForExplorerStruct) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "422a0e0e-4ab4-46d1-a74c-787758c454fa",
	}).Debug("Incoming - 'ListExecutionDataForTestInstructionExecutions'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "d3b00f2e-cb30-46dd-adae-96ce80e106b4",
	}).Debug("Outgoing - 'ListExecutionDataForTestInstructionExecutions'")

	// Check input keys for values
	if len(detailedTestSuiteExecutionMapKey) == 0 {
		return
	}
	for _, tempTestInstructionExecutionLogPostMapKey := range testInstructionExecutionDetailsMapKeys {
		if len(tempTestInstructionExecutionLogPostMapKey) == 0 {
			return
		}
	}

	// Lock Map for Writing
	detailedTestSuiteExecutionsMapMutex.Lock()

	//UnLock Map
	defer detailedTestSuiteExecutionsMapMutex.Unlock()

	var existInMap bool

	// Extract map with DetailedTestSuiteExecutionsMapObjects
	var detailedTestSuiteExecutionsObjectsMapPtr *map[DetailedTestSuiteExecutionMapKeyType]*DetailedTestSuiteExecutionsMapObjectStruct
	var detailedTestSuiteExecutionsObjectsMap map[DetailedTestSuiteExecutionMapKeyType]*DetailedTestSuiteExecutionsMapObjectStruct

	detailedTestSuiteExecutionsObjectsMapPtr = testSuiteExecutionsModel.DetailedTestSuiteExecutionsObjectsMapPtr
	detailedTestSuiteExecutionsObjectsMap = *detailedTestSuiteExecutionsObjectsMapPtr

	// Extract specificDetailedTestSuiteExecutionsMapObject
	var detailedTestSuiteExecutionsMapObjectPtr *DetailedTestSuiteExecutionsMapObjectStruct
	var detailedTestSuiteExecutionsMapObject DetailedTestSuiteExecutionsMapObjectStruct
	detailedTestSuiteExecutionsMapObjectPtr, existInMap = detailedTestSuiteExecutionsObjectsMap[detailedTestSuiteExecutionMapKey]

	// If it isn't initialized then exist
	if existInMap == false {
		return
	}

	detailedTestSuiteExecutionsMapObject = *detailedTestSuiteExecutionsMapObjectPtr

	// Check if TestInstructionnExecutionMap is nil, then exist
	if detailedTestSuiteExecutionsMapObjectPtr.TestInstructionExecutionLogPostMapPtr == nil {
		return
	}

	// Get the TestInstructionExecutionDetailsMap
	var testInstructionExecutionDetailsMapPtr *map[TestInstructionExecutionDetailsMapKeyType]*TestInstructionExecutionDetailsStruct
	var testInstructionExecutionDetailsMap map[TestInstructionExecutionDetailsMapKeyType]*TestInstructionExecutionDetailsStruct
	testInstructionExecutionDetailsMapPtr = detailedTestSuiteExecutionsMapObject.TestInstructionExecutionDetailsMapPtr
	testInstructionExecutionDetailsMap = *testInstructionExecutionDetailsMapPtr

	// Get TestInstructionExecutions from Map by loop the Map-keys and add the final result slice
	var testInstructionExecutionDetailsForExplorer []*TestInstructionExecutionDetailsForExplorerStruct
	var tempTestInstructionExecutionLogPostKey TestInstructionExecutionUuidType
	for _, testInstructionExecutionDetailsMapKey := range testInstructionExecutionDetailsMapKeys {

		// Convert to TestInstructionExecutionLogPostKey
		tempTestInstructionExecutionLogPostKey, existInMap = testSuiteExecutionsModel.
			GetTestInstructionExecutionUuidFromTestInstructionUuid(
				TestSuiteExecutionUuidType(detailedTestSuiteExecutionMapKey),
				RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType(testInstructionExecutionDetailsMapKey))

		if existInMap == false {

			sharedCode.Logger.WithFields(logrus.Fields{
				"id":                                    "31bc3351-7fe7-4c1b-8625-1c60c8627e6f",
				"testInstructionExecutionDetailsMapKey": testInstructionExecutionDetailsMapKey,
			}).Fatalln("Should never happen - Couldn't get tempTestInstructionExecutionLogPostKey from TestInstructionUuid")

			return nil
		}

		// Extract the TestInstructionExecution-data
		var tempTestInstructionExecutionDetailsPtr *TestInstructionExecutionDetailsStruct
		var tempTestInstructionExecutionDetails TestInstructionExecutionDetailsStruct
		tempTestInstructionExecutionDetailsPtr, existInMap = testInstructionExecutionDetailsMap[TestInstructionExecutionDetailsMapKeyType(tempTestInstructionExecutionLogPostKey)]
		if existInMap == false {

			sharedCode.Logger.WithFields(logrus.Fields{
				"id": "f284e78e-1b12-4834-86a4-2105f8865c6b",
			}).Fatalln("TestInstructionExecutionDetailsMapKey doesn't exist in 'testInstructionExecutionDetailsMap', should never happen")
		}

		tempTestInstructionExecutionDetails = *tempTestInstructionExecutionDetailsPtr

		// Loop 'TestInstructionExecutionDetails' and add to return slice

		for _, oneTestInstructionExecutionDetails := range tempTestInstructionExecutionDetails.TestInstructionExecutionDetails {

			// Create post for return slice
			var testInstructionExecutionDetails *TestInstructionExecutionDetailsForExplorerStruct
			testInstructionExecutionDetails = &TestInstructionExecutionDetailsForExplorerStruct{
				TestInstructionExecutionDetails:          oneTestInstructionExecutionDetails,
				TestInstructionExecutionBasicInformation: tempTestInstructionExecutionDetails.TestInstructionExecutionBasicInformation,
			}

			// Append to slice
			testInstructionExecutionDetailsForExplorer = append(testInstructionExecutionDetailsForExplorer, testInstructionExecutionDetails)
		}
	}

	testInstructionExecutionDetailsForExplorerPtr = &testInstructionExecutionDetailsForExplorer

	// sortLogPostsByTimestamp sorts LogPostAndValuesMessages slice by LogPostTimeStamp
	sortTestInstructionExecutionPostsByTimestamp(testInstructionExecutionDetailsForExplorerPtr, true)

	return testInstructionExecutionDetailsForExplorerPtr
}

// sortTestInstructionExecutionPostsByTimestamp sorts testInstructionExecutionDetailsForExplorer slice by ExecutionStatusUpdateTimeStamp
// Set ascending to true for ascending order, false for descending order
func sortTestInstructionExecutionPostsByTimestamp(
	testInstructionExecutionDetailsForExplorerPtr *[]*TestInstructionExecutionDetailsForExplorerStruct,
	ascending bool) {

	sort.SliceStable(*testInstructionExecutionDetailsForExplorerPtr, func(i, j int) bool {
		timeI := (*testInstructionExecutionDetailsForExplorerPtr)[i].TestInstructionExecutionDetails.ExecutionStatusUpdateTimeStamp.AsTime()
		timeJ := (*testInstructionExecutionDetailsForExplorerPtr)[j].TestInstructionExecutionDetails.ExecutionStatusUpdateTimeStamp.AsTime()

		if ascending {
			return timeI.Before(timeJ)
		}
		return timeI.After(timeJ)
	})
}
