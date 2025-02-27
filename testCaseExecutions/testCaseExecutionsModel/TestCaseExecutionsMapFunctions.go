package testCaseExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"sync"
)

var testCaseExecutionsMapMutex = &sync.RWMutex{}

// InitiateTestCaseExecutionsMap
// Add to the TestCaseExecutions-Map
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) InitiateTestCaseExecutionsMap() {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "8c6e4db9-7acf-4d90-b4ed-0992c64c423a",
	}).Debug("Incoming - 'InitiateTestCaseExecutionsMap'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "37116430-aac9-44fd-8727-593741f06cd1",
	}).Debug("Outgoing - 'InitiateTestCaseExecutionsMap'")

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

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "0c34753b-cb62-401c-9c79-afb300b4939e",
	}).Debug("Incoming - 'ReadFromTestCaseExecutionsMap'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "da469a24-48a0-455d-8da1-540f277c1373",
	}).Debug("Outgoing - 'ReadFromTestCaseExecutionsMap'")

	// Lock Map for Reading
	testCaseExecutionsMapMutex.RLock()

	//UnLock Map
	defer testCaseExecutionsMapMutex.RUnlock()

	// Check if Map i nil
	if TestCaseExecutionsModel.LatestTestCaseExecutionForEachTestCaseUuid.
		latestTestCaseExecutionForEachTestCaseUuidMap == nil {
		return nil, false
	}

	// Read Map
	testCaseExecutionsListMessage, existInMap = TestCaseExecutionsModel.LatestTestCaseExecutionForEachTestCaseUuid.
		latestTestCaseExecutionForEachTestCaseUuidMap[testCaseExecutionsMapKey]

	return testCaseExecutionsListMessage, existInMap
}

// ReadAllFromTestCaseExecutionsMap
// Read all from the TestCaseExecutions-Map
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) ReadAllFromTestCaseExecutionsMap() (
	testCaseExecutionsListMessage *[]*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "bf82106c-b885-4858-84a8-f7fbdceb0ff9",
	}).Debug("Incoming - 'ReadAllFromTestCaseExecutionsMap'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "4cb0939c-6447-482b-9313-ae661c520927",
	}).Debug("Outgoing - 'ReadAllFromTestCaseExecutionsMap'")

	// Lock Map for Reading
	testCaseExecutionsMapMutex.RLock()

	//UnLock Map
	defer testCaseExecutionsMapMutex.RUnlock()

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

	return &tempTestCaseExecutionsListMessage
}

// GetNumberOfTestCaseExecutionsRetrievedFromDatabase
// Read all from the TestCaseExecutions-Map
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) GetNumberOfTestCaseExecutionsRetrievedFromDatabase() (
	numberOfTestCaseExecutionsRetrievedFromDatabase int) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "55c98a19-b55b-45af-8064-913062a64450",
	}).Debug("Incoming - 'GetNumberOfTestCaseExecutionsRetrievedFromDatabase'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "295cac3c-46b8-4bfb-a2c2-1ee2235ca913",
	}).Debug("Outgoing - 'GetNumberOfTestCaseExecutionsRetrievedFromDatabase'")

	// Lock Map for Reading
	testCaseExecutionsMapMutex.RLock()

	//UnLock Map
	defer testCaseExecutionsMapMutex.RUnlock()

	var testCaseExecutionsListMessage *[]*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage

	// Get all data retrieved and get hwo many they are
	testCaseExecutionsListMessage = testCaseExecutionsModel.ReadAllFromTestCaseExecutionsMap()

	numberOfTestCaseExecutionsRetrievedFromDatabase = len(*testCaseExecutionsListMessage)

	return numberOfTestCaseExecutionsRetrievedFromDatabase

}

// AddToTestCaseExecutionsMap
// Add to the TestCaseExecutions-Map
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) AddToTestCaseExecutionsMap(
	testCaseExecutionsMapKey TestCaseExecutionUuidType,
	testCaseExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "48ea063f-b8c9-4555-847b-6be5385fbf11",
	}).Debug("Incoming - 'AddToTestCaseExecutionsMap'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "fc1dbc22-48a4-45f2-9c42-6e503fd7af7e",
	}).Debug("Outgoing - 'AddToTestCaseExecutionsMap'")

	// Lock Map for Writing
	testCaseExecutionsMapMutex.Lock()

	//UnLock Map
	defer testCaseExecutionsMapMutex.Unlock()

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

}

// DeleteFromTestCaseExecutionsMap
// Delete from the TestCaseExecutions-Map
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) DeleteFromTestCaseExecutionsMap(
	testCaseExecutionsMapKey TestCaseExecutionUuidType) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "4c41aa9c-e99b-4a94-9449-eb7a7e516dfb",
	}).Debug("Incoming - 'DeleteFromTestCaseExecutionsMap'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "7b7ed8a8-ff56-4ab9-bf3b-b6c0456a82c6",
	}).Debug("Outgoing - 'DeleteFromTestCaseExecutionsMap'")

	// Lock Map for Writing
	testCaseExecutionsMapMutex.Lock()

	//UnLock Map
	defer testCaseExecutionsMapMutex.Unlock()

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

}
