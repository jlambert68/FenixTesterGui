package testSuiteExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"sync"
)

var allTestCaseExecutionsMapMutex = &sync.RWMutex{}

/*
// InitiateAllTestCaseExecutionsForOneTestCaseMap
// Add to the InitiateAllTestCaseExecutionsForOneTestCaseMap-Map
func (testSuiteExecutionsModel TestCaseExecutionsModelStruct) InitiateAllTestCaseExecutionsForOneTestCaseMap() {

	// Lock Map for Writing
	allTestCaseExecutionsMapMutex.Lock()

	// Initiate map if it is not already done
	if testSuiteExecutionsModel.allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap == nil {

		testSuiteExecutionsModel.allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap =
			make(map[TestCaseUuidType]*AllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUserMapType)
	}

	//UnLock Map
	allTestCaseExecutionsMapMutex.Unlock()
}

*/

// GetAllTestCaseExecutionsForOneTestCaseUuid
// Get all TestCaseExecutions for one TestCaseUuid
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) GetAllTestCaseExecutionsForOneTestCaseUuid(
	testCaseUuidMapKey TestCaseUuidType) (
	tempTestCaseExecutionsList *[]*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage,
	existInMap bool) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "09bd03e2-843a-4b33-a735-a44c1c4d7076",
	}).Debug("Incoming - 'GetAllTestCaseExecutionsForOneTestCaseUuid'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "b7dc1fac-62bd-4f0e-b363-02ac7d9ae4af",
	}).Debug("Outgoing - 'GetAllTestCaseExecutionsForOneTestCaseUuid'")

	var testCaseExecutions []*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage

	// Lock Map for Reading
	allTestCaseExecutionsMapMutex.RLock()

	//UnLock Map
	defer allTestCaseExecutionsMapMutex.RUnlock()

	// Check if Outer Map i nil, then no TestCasesMapPtr with no TestCaseExecutions
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

	return &testCaseExecutions, existInMap
}

// GetSpecificTestCaseExecutionForOneTestCaseUuid
// Get one specific TestCaseExecutions for one TestCaseUuid
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) GetSpecificTestCaseExecutionForOneTestCaseUuid(
	testCaseUuidMapKey TestCaseUuidType,
	testCaseExecutionUuidMapKey TestCaseExecutionUuidType) (
	tempTestCaseExecution *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage,
	existInMap bool) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "25a958cd-1301-44c2-9724-3749d3b1dc9f",
	}).Debug("Incoming - 'GetSpecificTestCaseExecutionForOneTestCaseUuid'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "76d25456-3da0-4e7e-bbec-251e49c66d70",
	}).Debug("Outgoing - 'GetSpecificTestCaseExecutionForOneTestCaseUuid'")

	// Lock Map for Reading
	allTestCaseExecutionsMapMutex.RLock()

	//UnLock Map
	defer allTestCaseExecutionsMapMutex.RUnlock()

	// Check if Outer Map i nil, then no TestCasesMapPtr with no TestCaseExecutions
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

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "0ec6baee-3788-4634-82b6-c6f3e83af25e",
	}).Debug("Incoming - 'AddTestCaseExecutionsForOneTestCaseUuid'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "7add4c25-89ae-46ec-9f4d-89df07c59a36",
	}).Debug("Outgoing - 'AddTestCaseExecutionsForOneTestCaseUuid'")

	var existInMap bool

	// Lock Map for Writing
	allTestCaseExecutionsMapMutex.Lock()

	//UnLock Map
	defer allTestCaseExecutionsMapMutex.Unlock()

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

}

/*
// DeleteFromAllTestCaseExecutionsForOneTestCaseMap
// Delete from the DeleteFromAllTestCaseExecutionsForOneTestCase-Map
func (testSuiteExecutionsModel TestCaseExecutionsModelStruct) DeleteFromAllTestCaseExecutionsForOneTestCaseMap(
	testCaseExecutionsMapKey TestCaseExecutionUuidType) {

	// Lock Map for Writing
	allTestCaseExecutionsMapMutex.Lock()

	// Check if Map i nil
	if TestCaseExecutionsModel.allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap == nil {

		testSuiteExecutionsModel.allTestCaseExecutionsForAllTestCasesThatCanBeViewedByUserMap =
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
