package testCaseExecutionsModel

import (
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"sync"
)

var allTestCaseExecutionsMapMutex = &sync.RWMutex{}

// InitiateAllTestCaseExecutionsForOneTestCaseMap
// Add to the InitiateAllTestCaseExecutionsForOneTestCaseMap-Map
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) InitiateAllTestCaseExecutionsForOneTestCaseMap() {

	// Lock Map for Writing
	allTestCaseExecutionsMapMutex.Lock()

	// Initiate map if it is not already done
	if testCaseExecutionsModel.AllTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap == nil {

		testCaseExecutionsModel.AllTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap =
			make(map[AllTestCaseExecutionsForAllTestCasesUuidType]*AllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUserMapType)
	}

	//UnLock Map
	allTestCaseExecutionsMapMutex.Unlock()
}

// ReadFromAllTestCaseExecutionsForOneTestCaseMap
// Read from the ReadFromAllTestCaseExecutionsForOneTestCase-Map
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) ReadFromAllTestCaseExecutionsForOneTestCaseMap(
	testCaseExecutionsMapKey TestCaseExecutionUuidType) (
	testCaseExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage,
	existInMap bool) {

	// Lock Map for Reading
	allTestCaseExecutionsMapMutex.RLock()

	// Check if Outer Map i nil
	if TestCaseExecutionsModel.AllTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap == nil {
		return nil, false
	}

	// Check if Inner Map i nil
	if TestCaseExecutionsModel.AllTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap == nil {
		return nil, false
	}

	// Read Map
	testCaseExecutionsListMessage, existInMap = TestCaseExecutionsModel.
		AllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUserMap[testCaseExecutionsMapKey]

	//UnLock Map
	allTestCaseExecutionsMapMutex.RUnlock()

	return testCaseExecutionsListMessage, existInMap
}

// AddToAllTestCaseExecutionsForOneTestCaseMap
// Add to the AddToAllTestCaseExecutionsForOneTestCase-Map
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) AddToAllTestCaseExecutionsForOneTestCaseMap(
	testCaseExecutionsMapKey TestCaseExecutionUuidType,
	testCaseExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage) {

	// Lock Map for Writing
	allTestCaseExecutionsMapMutex.Lock()

	// Check if Map i nil
	if TestCaseExecutionsModel.AllTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap == nil {

		testCaseExecutionsModel.AllTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap =
			make(map[AllTestCaseExecutionsForAllTestCasesUuidType]*AllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUserMapType)
	}

	// Save to TestCaseExecutions-Map
	TestCaseExecutionsModel.
		AllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUserMap[testCaseExecutionsMapKey] = testCaseExecutionsListMessage

	//UnLock Map
	allTestCaseExecutionsMapMutex.Unlock()

}

// DeleteFromAllTestCaseExecutionsForOneTestCaseMap
// Delete from the DeleteFromAllTestCaseExecutionsForOneTestCase-Map
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) DeleteFromAllTestCaseExecutionsForOneTestCaseMap(
	testCaseExecutionsMapKey TestCaseExecutionUuidType) {

	// Lock Map for Writing
	allTestCaseExecutionsMapMutex.Lock()

	// Check if Map i nil
	if TestCaseExecutionsModel.AllTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap == nil {

		testCaseExecutionsModel.AllTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap =
			make(map[AllTestCaseExecutionsForAllTestCasesUuidType]*AllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUserMapType)

		return
	}

	// Save to TestCaseExecutions-Map
	delete(TestCaseExecutionsModel.AllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUserMap,
		testCaseExecutionsMapKey)

	//UnLock Map
	allTestCaseExecutionsMapMutex.Unlock()

}
