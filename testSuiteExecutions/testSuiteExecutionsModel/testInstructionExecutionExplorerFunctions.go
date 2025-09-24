package testSuiteExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	"github.com/sirupsen/logrus"
	"sort"
)

// ListExecutionDataForTestInstructionExecutions
// List all Execution-data for supplied TestInstructionExecutions. Posts are sorted on UpdateTimeStamp
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) ListExecutionDataForTestInstructionExecutions(
	detailedTestCaseExecutionMapKey DetailedTestCaseExecutionMapKeyType,
	testInstructionExecutionDetailsMapKeys []TestInstructionExecutionDetailsMapKeyType) (
	testInstructionExecutionDetailsForExplorerPtr *[]*TestInstructionExecutionDetailsForExplorerStruct) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "1e3f1544-1f9b-448f-a629-3f6256bbe6b7",
	}).Debug("Incoming - 'ListExecutionDataForTestInstructionExecutions'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "b254edaf-6868-4014-9681-55b19e6a169f",
	}).Debug("Outgoing - 'ListExecutionDataForTestInstructionExecutions'")

	// Check input keys for values
	if len(detailedTestCaseExecutionMapKey) == 0 {
		return
	}
	for _, tempTestInstructionExecutionLogPostMapKey := range testInstructionExecutionDetailsMapKeys {
		if len(tempTestInstructionExecutionLogPostMapKey) == 0 {
			return
		}
	}

	// Lock Map for Writing
	detailedTestCaseExecutionsMapMutex.Lock()

	//UnLock Map
	defer detailedTestCaseExecutionsMapMutex.Unlock()

	var existInMap bool

	// Extract map with DetailedTestCaseExecutionsMapObjects
	var detailedTestCaseExecutionsObjectsMapPtr *map[DetailedTestCaseExecutionMapKeyType]*DetailedTestCaseExecutionsMapObjectStruct
	var detailedTestCaseExecutionsObjectsMap map[DetailedTestCaseExecutionMapKeyType]*DetailedTestCaseExecutionsMapObjectStruct

	detailedTestCaseExecutionsObjectsMapPtr = testCaseExecutionsModel.DetailedTestCaseExecutionsObjectsMapPtr
	detailedTestCaseExecutionsObjectsMap = *detailedTestCaseExecutionsObjectsMapPtr

	// Extract specificDetailedTestCaseExecutionsMapObject
	var detailedTestCaseExecutionsMapObjectPtr *DetailedTestCaseExecutionsMapObjectStruct
	var detailedTestCaseExecutionsMapObject DetailedTestCaseExecutionsMapObjectStruct
	detailedTestCaseExecutionsMapObjectPtr, existInMap = detailedTestCaseExecutionsObjectsMap[detailedTestCaseExecutionMapKey]

	// If it isn't initialized then exist
	if existInMap == false {
		return
	}

	detailedTestCaseExecutionsMapObject = *detailedTestCaseExecutionsMapObjectPtr

	// Check if TestInstructionnExecutionMap is nil, then exist
	if detailedTestCaseExecutionsMapObjectPtr.TestInstructionExecutionLogPostMapPtr == nil {
		return
	}

	// Get the TestInstructionExecutionDetailsMap
	var testInstructionExecutionDetailsMapPtr *map[TestInstructionExecutionDetailsMapKeyType]*TestInstructionExecutionDetailsStruct
	var testInstructionExecutionDetailsMap map[TestInstructionExecutionDetailsMapKeyType]*TestInstructionExecutionDetailsStruct
	testInstructionExecutionDetailsMapPtr = detailedTestCaseExecutionsMapObject.TestInstructionExecutionDetailsMapPtr
	testInstructionExecutionDetailsMap = *testInstructionExecutionDetailsMapPtr

	// Get TestInstructionExecutions from Map by loop the Map-keys and add the final result slice
	var testInstructionExecutionDetailsForExplorer []*TestInstructionExecutionDetailsForExplorerStruct
	var tempTestInstructionExecutionLogPostKey TestInstructionExecutionUuidType
	for _, testInstructionExecutionDetailsMapKey := range testInstructionExecutionDetailsMapKeys {

		// Convert to TestInstructionExecutionLogPostKey
		tempTestInstructionExecutionLogPostKey, existInMap = testCaseExecutionsModel.
			GetTestInstructionExecutionUuidFromTestInstructionUuid(
				TestCaseExecutionUuidType(detailedTestCaseExecutionMapKey),
				RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType(testInstructionExecutionDetailsMapKey))

		if existInMap == false {

			sharedCode.Logger.WithFields(logrus.Fields{
				"id":                                    "4ef4595b-5ddf-406e-b09c-7b2eb628400c",
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
				"id": "a37f358c-68a4-47e4-81ba-183c0f191f62",
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
