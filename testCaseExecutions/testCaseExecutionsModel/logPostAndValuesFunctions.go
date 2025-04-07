package testCaseExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"sort"
	"strconv"
)

// ListLogPostsAndValuesForTestInstructionExecutions
// List all LogPosts and Values for supplied TestInstructionExecutions. Log-posts are sorted on Logging DateTime
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) ListLogPostsAndValuesForTestInstructionExecutions(
	detailedTestCaseExecutionMapKey DetailedTestCaseExecutionMapKeyType,
	testInstructionExecutionLogPostMapKeys []TestInstructionExecutionLogPostMapKeyType) (
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
	if len(*logPostAndValuesMessagesPtr) == 0 {
		return
	}
	for _, tempTestInstructionExecutionLogPostMapKey := range testInstructionExecutionLogPostMapKeys {
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
	detailedTestCaseExecutionsMapObject = *detailedTestCaseExecutionsMapObjectPtr

	// If it isn't initialized then exist
	if existInMap == false {
		return
	}

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
	for _, tempTestInstructionExecutionLogPostMapKey := range testInstructionExecutionLogPostMapKeys {

		// Get LogPosts for this Key
		var templogPostAndValuesMessagesPtr *[]*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage
		templogPostAndValuesMessagesPtr, existInMap = testInstructionExecutionLogPostMap[tempTestInstructionExecutionLogPostMapKey]
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

// ExtractAndStoreLogPostsAndValuesFromDetailedTestCaseExecution
// Extracts all LogPost-messages from a TestCaseExecution and store them in a map per TInTICExecutionKey.
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) ExtractAndStoreLogPostsAndValuesFromDetailedTestCaseExecution(
	detailedTestCaseExecutionMapKey DetailedTestCaseExecutionMapKeyType) (
	err error) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "d93737d7-3025-48c4-b683-bcc96a86024d",
	}).Debug("Incoming - 'ExtractAndStoreLogPostsAndValuesFromDetailedTestCaseExecution'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "000a3353-3fc2-44d8-9e22-8e7542fcad3f",
	}).Debug("Outgoing - 'ExtractAndStoreLogPostsAndValuesFromDetailedTestCaseExecution'")

	// Extract the raw detailedTestCaseExecution-message
	var detailedTestCaseExecution *fenixExecutionServerGuiGrpcApi.TestCaseExecutionResponseMessage
	var existInMap bool
	detailedTestCaseExecution, existInMap = testCaseExecutionsModel.
		ReadFromDetailedTestCaseExecutionsMap(detailedTestCaseExecutionMapKey)

	// Lock Map for Writing
	detailedTestCaseExecutionsMapMutex.Lock()

	//UnLock Map
	defer detailedTestCaseExecutionsMapMutex.Unlock()

	if existInMap == false {

		return err
	}

	// Extract map with DetailedTestCaseExecutionsMapObjects
	var detailedTestCaseExecutionsObjectsMapPtr *map[DetailedTestCaseExecutionMapKeyType]*DetailedTestCaseExecutionsMapObjectStruct
	var detailedTestCaseExecutionsObjectsMap map[DetailedTestCaseExecutionMapKeyType]*DetailedTestCaseExecutionsMapObjectStruct

	detailedTestCaseExecutionsObjectsMapPtr = testCaseExecutionsModel.DetailedTestCaseExecutionsObjectsMapPtr
	detailedTestCaseExecutionsObjectsMap = *detailedTestCaseExecutionsObjectsMapPtr

	// Extract specificDetailedTestCaseExecutionsMapObject
	var detailedTestCaseExecutionsMapObjectPtr *DetailedTestCaseExecutionsMapObjectStruct
	var detailedTestCaseExecutionsMapObject DetailedTestCaseExecutionsMapObjectStruct
	detailedTestCaseExecutionsMapObjectPtr, existInMap = detailedTestCaseExecutionsObjectsMap[detailedTestCaseExecutionMapKey]
	detailedTestCaseExecutionsMapObject = *detailedTestCaseExecutionsMapObjectPtr

	if existInMap == false {
		return err
	}

	// Always reInitialized TestInstructionExecutionLogPostMapPtr
	var tempTestInstructionExecutionLogPostMap map[TestInstructionExecutionLogPostMapKeyType]*[]*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage
	tempTestInstructionExecutionLogPostMap = make(map[TestInstructionExecutionLogPostMapKeyType]*[]*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage)
	detailedTestCaseExecutionsMapObjectPtr.TestInstructionExecutionLogPostMapPtr = &tempTestInstructionExecutionLogPostMap
	detailedTestCaseExecutionsMapObject = *detailedTestCaseExecutionsMapObjectPtr

	// Get the LogPostMap
	var testInstructionExecutionLogPostMapPtr *map[TestInstructionExecutionLogPostMapKeyType]*[]*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage
	var testInstructionExecutionLogPostMap map[TestInstructionExecutionLogPostMapKeyType]*[]*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage
	testInstructionExecutionLogPostMapPtr = detailedTestCaseExecutionsMapObject.TestInstructionExecutionLogPostMapPtr
	testInstructionExecutionLogPostMap = *testInstructionExecutionLogPostMapPtr

	// Loop all TestInstructionExecutions and extract logPost-message
	for _, testInstructionExecution := range detailedTestCaseExecution.TestInstructionExecutions {

		// Generate the TestInstructionExecutionLogPostMapKey
		var testInstructionExecutionLogPostMapKey TestInstructionExecutionLogPostMapKeyType
		testInstructionExecutionLogPostMapKey = TestInstructionExecutionLogPostMapKeyType(testInstructionExecution.GetTestInstructionExecutionBasicInformation().
			TestInstructionExecutionUuid +
			strconv.Itoa(int(testInstructionExecution.GetTestInstructionExecutionBasicInformation().GetTestInstructionExecutionVersion())))

		// Get the TestInstructionExecutionLogPOstAndValuesSLice
		var logPostAndValuesMessageSlicePtr *[]*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage
		logPostAndValuesMessageSlicePtr, existInMap = testInstructionExecutionLogPostMap[testInstructionExecutionLogPostMapKey]

		if existInMap == false {

			// No existing logs exist
			testInstructionExecutionLogPostMap[testInstructionExecutionLogPostMapKey] = &testInstructionExecution.ExecutionLogPostsAndValues

		} else {

			// Append to existing logs
			*logPostAndValuesMessageSlicePtr = append(*logPostAndValuesMessageSlicePtr, testInstructionExecution.ExecutionLogPostsAndValues...)
		}

	}

	return err
}
