package testSuiteExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"sort"
)

// ListLogPostsAndValuesForTestInstructionExecutions
// List all LogPosts and Values for supplied TestInstructionExecutions. Log-posts are sorted on Logging DateTime
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) ListLogPostsAndValuesForTestInstructionExecutions(
	detailedTestCaseExecutionMapKey DetailedTestCaseExecutionMapKeyType,
	testInstructionLogPostMapKeys []TestInstructionExecutionLogPostMapKeyType) (
	logPostAndValuesMessagesPtr *[]*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "ff572b43-297f-410d-931b-f5e5ea2aae53",
	}).Debug("Incoming - 'ListLogPostsAndValuesForTestInstructionExecutions'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "422691a7-9e68-447a-85f1-dfca07567606",
	}).Debug("Outgoing - 'ListLogPostsAndValuesForTestInstructionExecutions'")

	// Check input keys for values
	if len(detailedTestCaseExecutionMapKey) == 0 {
		return
	}
	//if logPostAndValuesMessagesPtr == nil || len(*logPostAndValuesMessagesPtr) == 0 {
	//	return
	//}
	for _, tempTestInstructionExecutionLogPostMapKey := range testInstructionLogPostMapKeys {
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

	// Check if TestInstructionNExecutionMap is nil, then exist
	if detailedTestCaseExecutionsMapObjectPtr.TestInstructionExecutionLogPostMapPtr == nil {
		return
	}

	// Get the LogPostMap
	var testInstructionExecutionLogPostMapPtr *map[TestInstructionExecutionLogPostMapKeyType]*[]*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage
	var testInstructionExecutionLogPostMap map[TestInstructionExecutionLogPostMapKeyType]*[]*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage
	testInstructionExecutionLogPostMapPtr = detailedTestCaseExecutionsMapObject.TestInstructionExecutionLogPostMapPtr
	testInstructionExecutionLogPostMap = *testInstructionExecutionLogPostMapPtr

	// Get logPost from Map by loop the Map-keys and add the final result slice
	var logPostAndValuesMessages []*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage
	var tempTestInstructionExecutionLogPostKey TestInstructionExecutionUuidType
	for _, tempTestInstructionLogPostMapKey := range testInstructionLogPostMapKeys {

		// Convert to TestInstructionExecutionLogPostKey
		tempTestInstructionExecutionLogPostKey, existInMap = testCaseExecutionsModel.
			GetTestInstructionExecutionUuidFromTestInstructionUuid(
				TestCaseExecutionUuidType(detailedTestCaseExecutionMapKey),
				RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType(tempTestInstructionLogPostMapKey))

		if existInMap == false {
			// TestInstructionContainer or that there are not attributes

			continue
			//sharedCode.Logger.WithFields(logrus.Fields{
			//	"id": "6859a4fb-33b2-47f7-9de0-0a26c2f73b1f",
			//}).Debug("Should never happen - Couldn't get tempTestInstructionExecutionLogPostKey from TestInstructionUuid")

			//return nil
		}

		// Get LogPosts for this Key
		var templogPostAndValuesMessagesPtr *[]*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage
		templogPostAndValuesMessagesPtr, existInMap = testInstructionExecutionLogPostMap[TestInstructionExecutionLogPostMapKeyType(tempTestInstructionExecutionLogPostKey)]
		if existInMap == true {
			// Add the logPosts to the logPost-slice
			logPostAndValuesMessages = append(logPostAndValuesMessages, *templogPostAndValuesMessagesPtr...)
		}

	}

	logPostAndValuesMessagesPtr = &logPostAndValuesMessages

	// sortLogPostsByTimestamp sorts LogPostAndValuesMessages slice by LogPostTimeStamp
	sortLogPostsByTimestamp(logPostAndValuesMessagesPtr, true)

	return logPostAndValuesMessagesPtr
}

// sortLogPostsByTimestamp sorts LogPostAndValuesMessages slice by LogPostTimeStamp
// Set ascending to true for ascending order, false for descending order
func sortLogPostsByTimestamp(logPostAndValuesMessagesPtr *[]*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage, ascending bool) {
	sort.SliceStable(*logPostAndValuesMessagesPtr, func(i, j int) bool {
		timeI := (*logPostAndValuesMessagesPtr)[i].LogPostTimeStamp.AsTime()
		timeJ := (*logPostAndValuesMessagesPtr)[j].LogPostTimeStamp.AsTime()

		if ascending {
			return timeI.Before(timeJ)
		}
		return timeI.After(timeJ)
	})
}
