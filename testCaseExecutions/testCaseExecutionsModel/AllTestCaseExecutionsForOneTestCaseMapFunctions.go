package testCaseExecutionsModel

import (
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"sync"
)

var allTestCaseExecutionsMapMutex = &sync.RWMutex{}

/*
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

*/

// GetAllTestCaseExecutionsForOneTestCaseUuid
// Get all TestCaseExecutions for one TestCaseUuid
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) GetAllTestCaseExecutionsForOneTestCaseUuid(
	testCaseUuidMapKey TestCaseUuidType) (
	tempTestCaseExecutionsList *[]*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage,
	existInMap bool) {

	var testCaseExecutions []*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage

	// Lock Map for Reading
	allTestCaseExecutionsMapMutex.RLock()

	// Check if Outer Map i nil, then no TestCases with no TestCaseExecutions
	if TestCaseExecutionsModel.AllTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap == nil {
		return nil, false
	}

	var tempTestCaseExecutionsForOneTestCasePtr *allTestCaseExecutionsForOneTestCaseUuidStruct
	var tempTestCaseExecutionsForOneTestCase allTestCaseExecutionsForOneTestCaseUuidStruct
	tempTestCaseExecutionsForOneTestCasePtr, existInMap = testCaseExecutionsModel.
		AllTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap[testCaseUuidMapKey]

	if existInMap == true {
		// TestCase exists in map
		tempTestCaseExecutionsForOneTestCase = *tempTestCaseExecutionsForOneTestCasePtr

	} else {

		// TestCase doesn't exist within map
		return nil, false
	}

	// Check if 'allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap' is initiated
	// If nil then TestCase has no TestCaseExecutions
	if tempTestCaseExecutionsForOneTestCase.allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap == nil {

		return nil, false
	}

	// Extract all TestCaseExecutions for the TestCaseUuid
	for _, tempTestCaseExecution := range tempTestCaseExecutionsForOneTestCase.
		allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap {

		testCaseExecutions = append(testCaseExecutions, tempTestCaseExecution)

	}

	//UnLock Map
	allTestCaseExecutionsMapMutex.RUnlock()

	return &testCaseExecutions, existInMap
}

// GetSpecificTestCaseExecutionForOneTestCaseUuid
// Get one specific TestCaseExecutions for one TestCaseUuid
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) GetSpecificTestCaseExecutionForOneTestCaseUuid(
	testCaseUuidMapKey TestCaseUuidType,
	testCaseExecutionUuidMapKey TestCaseExecutionUuidType) (
	tempTestCaseExecution *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage,
	existInMap bool) {

	// Lock Map for Reading
	allTestCaseExecutionsMapMutex.RLock()

	// Check if Outer Map i nil, then no TestCases with no TestCaseExecutions
	if TestCaseExecutionsModel.AllTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap == nil {
		return nil, false
	}

	var tempTestCaseExecutionsForOneTestCasePtr *allTestCaseExecutionsForOneTestCaseUuidStruct
	var tempTestCaseExecutionsForOneTestCase allTestCaseExecutionsForOneTestCaseUuidStruct
	tempTestCaseExecutionsForOneTestCasePtr, existInMap = testCaseExecutionsModel.
		AllTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap[testCaseUuidMapKey]

	if existInMap == true {
		// TestCase exists in map
		tempTestCaseExecutionsForOneTestCase = *tempTestCaseExecutionsForOneTestCasePtr

	} else {

		// TestCase doesn't exist within map
		return nil, false
	}

	// Check if 'allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap' is initiated
	// If nil then TestCase has no TestCaseExecutions
	if tempTestCaseExecutionsForOneTestCase.allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap == nil {

		return nil, false
	}

	// Extract the specific TestCaseExecution for the TestCaseUuid
	tempTestCaseExecution, existInMap = tempTestCaseExecutionsForOneTestCase.
		allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap[testCaseExecutionUuidMapKey]

	//UnLock Map
	allTestCaseExecutionsMapMutex.RUnlock()

	return tempTestCaseExecution, existInMap
}

// AddTestCaseExecutionsForOneTestCaseUuid
// Add a TestCaseExecution to the map for TestCaseExecutions per TestCaseUuid
func AddTestCaseExecutionsForOneTestCaseUuid(
	testCaseExecutionsModelRef *TestCaseExecutionsModelStruct,
	testCaseUuidMapKey TestCaseUuidType,
	testCaseExecutionUuidMapKey TestCaseExecutionUuidType,
	testCaseExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage,
	latestUniqueTestCaseExecutionDatabaseRowId int32,
	moreRowsExists bool) {

	var existInMap bool

	// Lock Map for Writing
	allTestCaseExecutionsMapMutex.Lock()

	// Check if Map i nil
	if testCaseExecutionsModelRef.AllTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap == nil {

		testCaseExecutionsModelRef.AllTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap =
			make(map[TestCaseUuidType]*allTestCaseExecutionsForOneTestCaseUuidStruct)
	}

	// Get saved TestCaseExecutions for the TestCaseUuid
	var tempTestCaseExecutionsForOneTestCasePtr *allTestCaseExecutionsForOneTestCaseUuidStruct
	var tempTestCaseExecutionsForOneTestCase allTestCaseExecutionsForOneTestCaseUuidStruct
	tempTestCaseExecutionsForOneTestCasePtr, existInMap = testCaseExecutionsModelRef.
		AllTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap[testCaseUuidMapKey]

	if existInMap == true {
		// Add to the existing TestCaseExecutions for the TestCaseUuid
		tempTestCaseExecutionsForOneTestCase = *tempTestCaseExecutionsForOneTestCasePtr

	} else {
		// Initiate a new 'allTestCaseExecutionsForOneTestCaseUuidStruct'
		tempTestCaseExecutionsForOneTestCase = allTestCaseExecutionsForOneTestCaseUuidStruct{}
	}

	// Check if 'allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap' is initated
	if tempTestCaseExecutionsForOneTestCase.allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap == nil {
		// Initiate the map
		tempTestCaseExecutionsForOneTestCase.allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap =
			make(map[TestCaseExecutionUuidType]*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage)
	}

	// Save 'testCaseExecutionsListMessage' for TestCase
	tempTestCaseExecutionsForOneTestCase.
		allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap[testCaseExecutionUuidMapKey] =
		testCaseExecutionsListMessage

	// Save 'latestUniqueTestCaseExecutionDatabaseRowId' & 'moreRowsExists'
	tempTestCaseExecutionsForOneTestCase.latestUniqueTestCaseExecutionDatabaseRowId =
		latestUniqueTestCaseExecutionDatabaseRowId
	tempTestCaseExecutionsForOneTestCase.moreRowsExists = moreRowsExists

	// Store the execution back into TestCaseUuid-map
	testCaseExecutionsModelRef.
		AllTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap[testCaseUuidMapKey] =
		&tempTestCaseExecutionsForOneTestCase

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
