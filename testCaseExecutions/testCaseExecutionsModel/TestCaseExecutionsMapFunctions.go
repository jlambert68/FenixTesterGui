package testCaseExecutionsModel

import (
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"sync"
)

var testCaseExecutionsMapMutex = &sync.RWMutex{}

// InitiateTestCaseExecutionsMap
// Add to the TestCaseExecutions-Map
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) InitiateTestCaseExecutionsMap() {

	// Lock Map for Writing
	testCaseExecutionsMapMutex.Lock()

	// Initiate map if it is not already done
	if testCaseExecutionsModel.LatestTestCaseExecutionForEachTestCaseUuid.
		latestTestCaseExecutionForEachTestCaseUuidMap == nil {
		testCaseExecutionsModel.LatestTestCaseExecutionForEachTestCaseUuid.
			latestTestCaseExecutionForEachTestCaseUuidMap = make(map[TestCaseExecutionUuidType]*fenixExecutionServerGuiGrpcApi.
			TestCaseExecutionsListMessage)
	}

	//UnLock Map
	testCaseExecutionsMapMutex.Unlock()
}

// ReadFromTestCaseExecutionsMap
// Read from the TestCaseExecutions-Map
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) ReadFromTestCaseExecutionsMap(
	testCaseExecutionsMapKey TestCaseExecutionUuidType) (
	testCaseExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage,
	existInMap bool) {

	// Lock Map for Reading
	testCaseExecutionsMapMutex.RLock()

	// Check if Map i nil
	if TestCaseExecutionsModel.LatestTestCaseExecutionForEachTestCaseUuid.
		latestTestCaseExecutionForEachTestCaseUuidMap == nil {
		return nil, false
	}

	// Read Map
	testCaseExecutionsListMessage, existInMap = TestCaseExecutionsModel.LatestTestCaseExecutionForEachTestCaseUuid.
		latestTestCaseExecutionForEachTestCaseUuidMap[testCaseExecutionsMapKey]

	//UnLock Map
	testCaseExecutionsMapMutex.RUnlock()

	return testCaseExecutionsListMessage, existInMap
}

// ReadAllFromTestCaseExecutionsMap
// Read all from the TestCaseExecutions-Map
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) ReadAllFromTestCaseExecutionsMap() (
	testCaseExecutionsListMessage *[]*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage) {

	// Lock Map for Reading
	testCaseExecutionsMapMutex.RLock()

	// Check if Map i nil
	if TestCaseExecutionsModel.LatestTestCaseExecutionForEachTestCaseUuid.
		latestTestCaseExecutionForEachTestCaseUuidMap == nil {
		return nil
	}

	var tempTestCaseExecutionsListMessage []*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage

	// Loop all items in map and add to response slice
	for _, testCaseExecutionListMessage := range TestCaseExecutionsModel.LatestTestCaseExecutionForEachTestCaseUuid.
		latestTestCaseExecutionForEachTestCaseUuidMap {
		tempTestCaseExecutionsListMessage = append(tempTestCaseExecutionsListMessage, testCaseExecutionListMessage)
	}

	//UnLock Map
	testCaseExecutionsMapMutex.RUnlock()

	return &tempTestCaseExecutionsListMessage
}

// AddToTestCaseExecutionsMap
// Add to the TestCaseExecutions-Map
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) AddToTestCaseExecutionsMap(
	testCaseExecutionsMapKey TestCaseExecutionUuidType,
	testCaseExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage) {

	// Lock Map for Writing
	testCaseExecutionsMapMutex.Lock()

	// Check if Map i nil
	if TestCaseExecutionsModel.LatestTestCaseExecutionForEachTestCaseUuid.
		latestTestCaseExecutionForEachTestCaseUuidMap == nil {

		TestCaseExecutionsModel.LatestTestCaseExecutionForEachTestCaseUuid.
			latestTestCaseExecutionForEachTestCaseUuidMap = make(
			map[TestCaseExecutionUuidType]*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage)
	}

	// Save to TestCaseExecutions-Map
	TestCaseExecutionsModel.LatestTestCaseExecutionForEachTestCaseUuid.
		latestTestCaseExecutionForEachTestCaseUuidMap[testCaseExecutionsMapKey] = testCaseExecutionsListMessage

	//UnLock Map
	testCaseExecutionsMapMutex.Unlock()

}

// DeleteFromTestCaseExecutionsMap
// Delete from the TestCaseExecutions-Map
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) DeleteFromTestCaseExecutionsMap(
	testCaseExecutionsMapKey TestCaseExecutionUuidType) {

	// Lock Map for Writing
	testCaseExecutionsMapMutex.Lock()

	// Check if Map i nil
	if TestCaseExecutionsModel.LatestTestCaseExecutionForEachTestCaseUuid.latestTestCaseExecutionForEachTestCaseUuidMap == nil {

		TestCaseExecutionsModel.LatestTestCaseExecutionForEachTestCaseUuid.
			latestTestCaseExecutionForEachTestCaseUuidMap = make(
			map[TestCaseExecutionUuidType]*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage)

		return
	}

	// Save to TestCaseExecutions-Map
	delete(TestCaseExecutionsModel.LatestTestCaseExecutionForEachTestCaseUuid.
		latestTestCaseExecutionForEachTestCaseUuidMap,
		testCaseExecutionsMapKey)

	//UnLock Map
	testCaseExecutionsMapMutex.Unlock()

}
