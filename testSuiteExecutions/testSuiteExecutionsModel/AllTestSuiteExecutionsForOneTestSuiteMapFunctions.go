package testSuiteExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"sync"
)

var allTestSuiteExecutionsMapMutex = &sync.RWMutex{}

/*
// InitiateAllTestSuiteExecutionsForOneTestSuiteMap
// Add to the InitiateAllTestSuiteExecutionsForOneTestSuiteMap-Map
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) InitiateAllTestSuiteExecutionsForOneTestSuiteMap() {

	// Lock Map for Writing
	allTestSuiteExecutionsMapMutex.Lock()

	// Initiate map if it is not already done
	if testSuiteExecutionsModel.allTestSuiteExecutionsForAllTestSuitesThatCanBeViewedByUserMap == nil {

		testSuiteExecutionsModel.allTestSuiteExecutionsForAllTestSuitesThatCanBeViewedByUserMap =
			make(map[TestSuiteUuidType]*AllTestSuiteExecutionsForOneTestSuiteThatCanBeViewedByUserMapType)
	}

	//UnLock Map
	allTestSuiteExecutionsMapMutex.Unlock()
}

*/

// GetAllTestSuiteExecutionsForOneTestSuiteUuid
// Get all TestSuiteExecutions for one TestSuiteUuid
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) GetAllTestSuiteExecutionsForOneTestSuiteUuid(
	testSuiteUuidMapKey TestSuiteUuidType) (
	tempTestSuiteExecutionsList *[]*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage,
	existInMap bool) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "1c2f2b96-1077-40ee-8615-07ff6b9d612c",
	}).Debug("Incoming - 'GetAllTestSuiteExecutionsForOneTestSuiteUuid'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "ebe287d0-953d-4b72-8dfc-b4c1dc763a06",
	}).Debug("Outgoing - 'GetAllTestSuiteExecutionsForOneTestSuiteUuid'")

	var testSuiteExecutions []*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage

	// Lock Map for Reading
	allTestSuiteExecutionsMapMutex.RLock()

	//UnLock Map
	defer allTestSuiteExecutionsMapMutex.RUnlock()

	// Check if Outer Map i nil, then no TestSuitesMapPtr with no TestSuiteExecutions
	if TestSuiteExecutionsModel.AllTestSuiteExecutionsForAllTestSuitesThatCanBeViewedByUserMap == nil {
		return nil, false
	}

	var tempTestSuiteExecutionsForOneTestSuitePtr *allTestSuiteExecutionsForOneTestSuiteUuidStruct
	var tempTestSuiteExecutionsForOneTestSuite allTestSuiteExecutionsForOneTestSuiteUuidStruct
	tempTestSuiteExecutionsForOneTestSuitePtr, existInMap = testSuiteExecutionsModel.
		AllTestSuiteExecutionsForAllTestSuitesThatCanBeViewedByUserMap[testSuiteUuidMapKey]

	if existInMap == true {
		// TestSuite exists in map
		tempTestSuiteExecutionsForOneTestSuite = *tempTestSuiteExecutionsForOneTestSuitePtr

	} else {

		// TestSuite doesn't exist within map
		return nil, false
	}

	// Check if 'allTestSuiteExecutionsForAllTestSuitesThatCanBeViewedByUserMap' is initiated
	// If nil then TestSuite has no TestSuiteExecutions
	if tempTestSuiteExecutionsForOneTestSuite.allTestSuiteExecutionsForAllTestSuitesThatCanBeViewedByUserMap == nil {

		return nil, false
	}

	// Extract all TestSuiteExecutions for the TestSuiteUuid
	for _, tempTestSuiteExecution := range tempTestSuiteExecutionsForOneTestSuite.
		allTestSuiteExecutionsForAllTestSuitesThatCanBeViewedByUserMap {

		testSuiteExecutions = append(testSuiteExecutions, tempTestSuiteExecution)

	}

	return &testSuiteExecutions, existInMap
}

// GetSpecificTestSuiteExecutionForOneTestSuiteUuid
// Get one specific TestSuiteExecutions for one TestSuiteUuid
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) GetSpecificTestSuiteExecutionForOneTestSuiteUuid(
	testSuiteUuidMapKey TestSuiteUuidType,
	testSuiteExecutionUuidMapKey TestSuiteExecutionUuidType) (
	tempTestSuiteExecution *fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage,
	existInMap bool) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "f95a4685-34f9-4c10-996f-d9220a8fb7db",
	}).Debug("Incoming - 'GetSpecificTestSuiteExecutionForOneTestSuiteUuid'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "a984570b-81fb-45ca-b186-4d70ce3f1817",
	}).Debug("Outgoing - 'GetSpecificTestSuiteExecutionForOneTestSuiteUuid'")

	// Lock Map for Reading
	allTestSuiteExecutionsMapMutex.RLock()

	//UnLock Map
	defer allTestSuiteExecutionsMapMutex.RUnlock()

	// Check if Outer Map i nil, then no TestSuitesMapPtr with no TestSuiteExecutions
	if TestSuiteExecutionsModel.AllTestSuiteExecutionsForAllTestSuitesThatCanBeViewedByUserMap == nil {
		return nil, false
	}

	var tempTestSuiteExecutionsForOneTestSuitePtr *allTestSuiteExecutionsForOneTestSuiteUuidStruct
	var tempTestSuiteExecutionsForOneTestSuite allTestSuiteExecutionsForOneTestSuiteUuidStruct
	tempTestSuiteExecutionsForOneTestSuitePtr, existInMap = testSuiteExecutionsModel.
		AllTestSuiteExecutionsForAllTestSuitesThatCanBeViewedByUserMap[testSuiteUuidMapKey]

	if existInMap == true {
		// TestSuite exists in map
		tempTestSuiteExecutionsForOneTestSuite = *tempTestSuiteExecutionsForOneTestSuitePtr

	} else {

		// TestSuite doesn't exist within map
		return nil, false
	}

	// Check if 'allTestSuiteExecutionsForAllTestSuitesThatCanBeViewedByUserMap' is initiated
	// If nil then TestSuite has no TestSuiteExecutions
	if tempTestSuiteExecutionsForOneTestSuite.allTestSuiteExecutionsForAllTestSuitesThatCanBeViewedByUserMap == nil {

		return nil, false
	}

	// Extract the specific TestSuiteExecution for the TestSuiteUuid
	tempTestSuiteExecution, existInMap = tempTestSuiteExecutionsForOneTestSuite.
		allTestSuiteExecutionsForAllTestSuitesThatCanBeViewedByUserMap[testSuiteExecutionUuidMapKey]

	return tempTestSuiteExecution, existInMap
}

// AddTestSuiteExecutionsForOneTestSuiteUuid
// Add a TestSuiteExecution to the map for TestSuiteExecutions per TestSuiteUuid
func AddTestSuiteExecutionsForOneTestSuiteUuid(
	testSuiteExecutionsModelRef *TestSuiteExecutionsModelStruct,
	testSuiteUuidMapKey TestSuiteUuidType,
	testSuiteExecutionUuidMapKey TestSuiteExecutionUuidType,
	testSuiteExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage,
	latestUniqueTestSuiteExecutionDatabaseRowId int32,
	moreRowsExists bool) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "b5393735-1b8b-4112-8196-6ea478c0781a",
	}).Debug("Incoming - 'AddTestSuiteExecutionsForOneTestSuiteUuid'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "b2689f4c-f09f-4c8d-9cf2-c0075eb69008",
	}).Debug("Outgoing - 'AddTestSuiteExecutionsForOneTestSuiteUuid'")

	var existInMap bool

	// Lock Map for Writing
	allTestSuiteExecutionsMapMutex.Lock()

	//UnLock Map
	defer allTestSuiteExecutionsMapMutex.Unlock()

	// Check if Map i nil
	if testSuiteExecutionsModelRef.AllTestSuiteExecutionsForAllTestSuitesThatCanBeViewedByUserMap == nil {

		testSuiteExecutionsModelRef.AllTestSuiteExecutionsForAllTestSuitesThatCanBeViewedByUserMap =
			make(map[TestSuiteUuidType]*allTestSuiteExecutionsForOneTestSuiteUuidStruct)
	}

	// Get saved TestSuiteExecutions for the TestSuiteUuid
	var tempTestSuiteExecutionsForOneTestSuitePtr *allTestSuiteExecutionsForOneTestSuiteUuidStruct
	var tempTestSuiteExecutionsForOneTestSuite allTestSuiteExecutionsForOneTestSuiteUuidStruct
	tempTestSuiteExecutionsForOneTestSuitePtr, existInMap = testSuiteExecutionsModelRef.
		AllTestSuiteExecutionsForAllTestSuitesThatCanBeViewedByUserMap[testSuiteUuidMapKey]

	if existInMap == true {
		// Add to the existing TestSuiteExecutions for the TestSuiteUuid
		tempTestSuiteExecutionsForOneTestSuite = *tempTestSuiteExecutionsForOneTestSuitePtr

	} else {
		// Initiate a new 'allTestSuiteExecutionsForOneTestSuiteUuidStruct'
		tempTestSuiteExecutionsForOneTestSuite = allTestSuiteExecutionsForOneTestSuiteUuidStruct{}
	}

	// Check if 'allTestSuiteExecutionsForAllTestSuitesThatCanBeViewedByUserMap' is initated
	if tempTestSuiteExecutionsForOneTestSuite.allTestSuiteExecutionsForAllTestSuitesThatCanBeViewedByUserMap == nil {
		// Initiate the map
		tempTestSuiteExecutionsForOneTestSuite.allTestSuiteExecutionsForAllTestSuitesThatCanBeViewedByUserMap =
			make(map[TestSuiteExecutionUuidType]*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage)
	}

	// Save 'testSuiteExecutionsListMessage' for TestSuite
	tempTestSuiteExecutionsForOneTestSuite.
		allTestSuiteExecutionsForAllTestSuitesThatCanBeViewedByUserMap[testSuiteExecutionUuidMapKey] =
		testSuiteExecutionsListMessage

	// Save 'latestUniqueTestSuiteExecutionDatabaseRowId' & 'moreRowsExists'
	tempTestSuiteExecutionsForOneTestSuite.latestUniqueTestSuiteExecutionDatabaseRowId =
		latestUniqueTestSuiteExecutionDatabaseRowId
	tempTestSuiteExecutionsForOneTestSuite.moreRowsExists = moreRowsExists

	// Store the execution back into TestSuiteUuid-map
	testSuiteExecutionsModelRef.
		AllTestSuiteExecutionsForAllTestSuitesThatCanBeViewedByUserMap[testSuiteUuidMapKey] =
		&tempTestSuiteExecutionsForOneTestSuite

}

/*
// DeleteFromAllTestSuiteExecutionsForOneTestSuiteMap
// Delete from the DeleteFromAllTestSuiteExecutionsForOneTestSuite-Map
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) DeleteFromAllTestSuiteExecutionsForOneTestSuiteMap(
	testSuiteExecutionsMapKey TestSuiteExecutionUuidType) {

	// Lock Map for Writing
	allTestSuiteExecutionsMapMutex.Lock()

	// Check if Map i nil
	if TestSuiteExecutionsModel.allTestSuiteExecutionsForAllTestSuitesThatCanBeViewedByUserMap == nil {

		testSuiteExecutionsModel.allTestSuiteExecutionsForAllTestSuitesThatCanBeViewedByUserMap =
			make(map[TestSuiteUuidType]*AllTestSuiteExecutionsForOneTestSuiteThatCanBeViewedByUserMapType)

		return
	}

	// Save to TestSuiteExecutions-Map
	delete(TestSuiteExecutionsModel.AllTestSuiteExecutionsForOneTestSuiteThatCanBeViewedByUserMap,
		testSuiteExecutionsMapKey)

	//UnLock Map
	allTestSuiteExecutionsMapMutex.Unlock()

}


*/
