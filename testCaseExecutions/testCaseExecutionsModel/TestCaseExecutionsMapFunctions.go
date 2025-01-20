package testCaseExecutionsModel

import (
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"sync"
)

var testCaseExecutionsMapMutex = &sync.RWMutex{}

// ReadFromTestCaseExecutionsMap
// Read from the TestCaseExecutions-Map
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) ReadFromTestCaseExecutionsMap(
	testCaseExecutionsMapKey string) (
	testCaseExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage,
	existInMap bool) {

	// Lock Map for Reading
	testCaseExecutionsMapMutex.RLock()

	// Check if Map i nil
	if TestCaseExecutionsModel.TestCaseExecutionsThatCanBeViewedByUserMap == nil {
		return nil, false
	}

	// Read Map
	testCaseExecutionsListMessage, existInMap = TestCaseExecutionsModel.TestCaseExecutionsThatCanBeViewedByUserMap[testCaseExecutionsMapKey]

	//UnLock Map
	testCaseExecutionsMapMutex.RUnlock()

	return testCaseExecutionsListMessage, existInMap
}

// AddToTestCaseExecutionsMap
// Add to the TestCaseExecutions-Map
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) AddToTestCaseExecutionsMap(
	testCaseExecutionsMapKey string,
	testCaseExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage) {

	// Lock Map for Writing
	testCaseExecutionsMapMutex.Lock()

	// Check if Map i nil
	if TestCaseExecutionsModel.TestCaseExecutionsThatCanBeViewedByUserMap == nil {

		TestCaseExecutionsModel.TestCaseExecutionsThatCanBeViewedByUserMap = make(
			map[string]*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage)
	}

	// Save to TestCaseExecutions-Map
	TestCaseExecutionsModel.TestCaseExecutionsThatCanBeViewedByUserMap[testCaseExecutionsMapKey] = testCaseExecutionsListMessage

	//UnLock Map
	testCaseExecutionsMapMutex.Unlock()

}

// DeleteFromTestCaseExecutionsMap
// Delete from the TestCaseExecutions-Map
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) DeleteFromTestCaseExecutionsMap(
	testCaseExecutionsMapKey string) {

	// Lock Map for Writing
	testCaseExecutionsMapMutex.Lock()

	// Check if Map i nil
	if TestCaseExecutionsModel.TestCaseExecutionsThatCanBeViewedByUserMap == nil {

		TestCaseExecutionsModel.TestCaseExecutionsThatCanBeViewedByUserMap = make(
			map[string]*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage)

		return
	}

	// Save to TestCaseExecutions-Map
	delete(TestCaseExecutionsModel.TestCaseExecutionsThatCanBeViewedByUserMap, testCaseExecutionsMapKey)

	//UnLock Map
	testCaseExecutionsMapMutex.Unlock()

}
