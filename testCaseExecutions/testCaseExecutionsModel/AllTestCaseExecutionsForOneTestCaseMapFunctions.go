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
	if testCaseExecutionsModel.allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap == nil {

		testCaseExecutionsModel.allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap =
			make(map[TestCaseUuidType]*AllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUserMapType)
	}

	//UnLock Map
	allTestCaseExecutionsMapMutex.Unlock()
}

// ReadFromAllTestCaseExecutionsForOneTestCaseMap
// Read from the ReadFromAllTestCaseExecutionsForOneTestCase-Map
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) ReadFromAllTestCaseExecutionsForOneTestCaseMap(
	testCaseUuidMapKey TestCaseUuidType,
	testCaseExecutionUuidMapKey TestCaseExecutionUuidType) (
	testCaseExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage,
	existInMap bool) {

	// Lock Map for Reading
	allTestCaseExecutionsMapMutex.RLock()

	// Check if Outer Map i nil
	if TestCaseExecutionsModel.allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap == nil {
		return nil, false
	}

	// Get saved TestCaseExecutions for the TestCaseUuid
	var tempTestCaseExecutionsForOneTestCasePtr *AllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUserMapType
	var tempTestCaseExecutionsForOneTestCase AllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUserMapType
	tempTestCaseExecutionsForOneTestCasePtr, existInMap = testCaseExecutionsModel.
		allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap[testCaseUuidMapKey]

	if existInMap == false {

		// TestCaseUuid doesn't exist in map
		return nil, false
	}

	// Check if TestCaseExecutions exist in TestCaseUuid-map
	tempTestCaseExecutionsForOneTestCase = *tempTestCaseExecutionsForOneTestCasePtr
	testCaseExecutionsListMessage, existInMap = tempTestCaseExecutionsForOneTestCase[testCaseExecutionUuidMapKey]

	if existInMap == false {

		// TestCaseUuid doesn't exist in map
		return nil, false
	}

	//UnLock Map
	allTestCaseExecutionsMapMutex.RUnlock()

	return testCaseExecutionsListMessage, existInMap
}

// ReadAllFromAllTestCaseExecutionsForOneTestCaseMap
// Read all TestCaseExecutions from the ReadAllFromAllTestCaseExecutionsForOneTestCase-Map
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) ReadAllFromAllTestCaseExecutionsForOneTestCaseMap(
	testCaseUuidMapKey TestCaseUuidType) (
	testCaseExecutionsListMessage *[]*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage,
	existInMap bool) {

	// Lock Map for Reading
	allTestCaseExecutionsMapMutex.RLock()

	// Check if Outer Map i nil
	if TestCaseExecutionsModel.allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap == nil {
		return nil, false
	}

	// Get saved TestCaseExecutions for the TestCaseUuid
	var tempTestCaseExecutionsForOneTestCasePtr *AllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUserMapType
	var tempTestCaseExecutionsForOneTestCase AllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUserMapType
	tempTestCaseExecutionsForOneTestCasePtr, existInMap = testCaseExecutionsModel.
		allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap[testCaseUuidMapKey]

	if existInMap == false {

		// TestCaseUuid doesn't exist in map
		return nil, false
	}

	// Extract all TestCaseExecutions in TestCaseUuid-map for one TestCaseUuid
	var tempTestCaseExecutionsListMessage []*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage
	tempTestCaseExecutionsForOneTestCase = *tempTestCaseExecutionsForOneTestCasePtr

	for _, tempTestCaseExecution := range tempTestCaseExecutionsForOneTestCase {
		tempTestCaseExecutionsListMessage = append(tempTestCaseExecutionsListMessage, tempTestCaseExecution)
	}

	//UnLock Map
	allTestCaseExecutionsMapMutex.RUnlock()

	return &tempTestCaseExecutionsListMessage, existInMap
}

// AddToAllTestCaseExecutionsForOneTestCaseMap
// Add to the AddToAllTestCaseExecutionsForOneTestCase-Map
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) AddToAllTestCaseExecutionsForOneTestCaseMap(
	testCaseUuidMapKey TestCaseUuidType,
	testCaseExecutionUuidMapKey TestCaseExecutionUuidType,
	testCaseExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage) {

	var existInMap bool

	// Lock Map for Writing
	allTestCaseExecutionsMapMutex.Lock()

	// Check if Map i nil
	if TestCaseExecutionsModel.allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap == nil {

		testCaseExecutionsModel.allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap =
			make(map[TestCaseUuidType]*AllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUserMapType)
	}

	// Get saved TestCaseExecutions for the TestCaseUuid
	var tempTestCaseExecutionsForOneTestCasePtr *AllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUserMapType
	var tempTestCaseExecutionsForOneTestCase AllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUserMapType
	tempTestCaseExecutionsForOneTestCasePtr, existInMap = testCaseExecutionsModel.
		allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap[testCaseUuidMapKey]

	if existInMap == true {
		// Add to the existing TestCaseExecutions for the TestCaseUuid
		tempTestCaseExecutionsForOneTestCase = *tempTestCaseExecutionsForOneTestCasePtr
		tempTestCaseExecutionsForOneTestCase[testCaseExecutionUuidMapKey] = testCaseExecutionsListMessage

		testCaseExecutionsModel.allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap[testCaseUuidMapKey] =
			&tempTestCaseExecutionsForOneTestCase

	} else {
		// Initiate a new map for the TestCase
		tempTestCaseExecutionsForOneTestCase = make(map[TestCaseExecutionUuidType]*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage)

		// Add the TestCaseExecution
		tempTestCaseExecutionsForOneTestCase[testCaseExecutionUuidMapKey] = testCaseExecutionsListMessage

		// Add the TestCaseExecution to the TestCase-map
		testCaseExecutionsModel.
			allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap[testCaseUuidMapKey] = &tempTestCaseExecutionsForOneTestCase
	}

	//UnLock Map
	allTestCaseExecutionsMapMutex.Unlock()

}

/*
// DeleteFromAllTestCaseExecutionsForOneTestCaseMap
// Delete from the DeleteFromAllTestCaseExecutionsForOneTestCase-Map
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) DeleteFromAllTestCaseExecutionsForOneTestCaseMap(
	testCaseExecutionsMapKey TestCaseExecutionUuidType) {

	// Lock Map for Writing
	allTestCaseExecutionsMapMutex.Lock()

	// Check if Map i nil
	if TestCaseExecutionsModel.allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap == nil {

		testCaseExecutionsModel.allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap =
			make(map[TestCaseUuidType]*AllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUserMapType)

		return
	}

	// Save to TestCaseExecutions-Map
	delete(TestCaseExecutionsModel.AllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUserMap,
		testCaseExecutionsMapKey)

	//UnLock Map
	allTestCaseExecutionsMapMutex.Unlock()

}


*/
