package testSuiteExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"sync"
)

var testCaseExecutionsMapMutex = &sync.RWMutex{}

// InitiateTestCaseExecutionsMap
// Add to the TestCaseExecutions-Map
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) InitiateTestCaseExecutionsMap() {

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
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) ReadFromTestCaseExecutionsMap(
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
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) ReadAllFromTestCaseExecutionsMap() (
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
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) GetNumberOfTestCaseExecutionsRetrievedFromDatabase() (
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
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) AddToTestCaseExecutionsMap(
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
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) DeleteFromTestCaseExecutionsMap(
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

// GetTestInstructionExecutionUuidFromTestInstructionUuid
// Read from the TestCaseExecutions-Map and get the TestInstructionExecutionUuid + Version (mpKey) based on TestInstructionUuid
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) GetTestInstructionExecutionUuidFromTestInstructionUuid(
	testCaseExecutionsMapKey TestCaseExecutionUuidType,
	testInstructionUuid RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType) (
	testInstructionExecutionUuid TestInstructionExecutionUuidType,
	existInMap bool) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "0c34753b-cb62-401c-9c79-afb300b4939e",
	}).Debug("Incoming - 'GetTestInstructionExecutionUuidFromTestInstructionUuid'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "da469a24-48a0-455d-8da1-540f277c1373",
	}).Debug("Outgoing - 'GetTestInstructionExecutionUuidFromTestInstructionUuid'")

	// Lock Map for Reading
	//testCaseExecutionsMapMutex.RLock()

	//UnLock Map
	//defer testCaseExecutionsMapMutex.RUnlock()

	// Check if Map i nil
	if TestCaseExecutionsModel.LatestTestCaseExecutionForEachTestCaseUuid.
		latestTestCaseExecutionForEachTestCaseUuidMap == nil {
		return "", false
	}

	// Read Map
	var tempDetailedTestCaseExecutionsObjectsMapPtr *map[DetailedTestCaseExecutionMapKeyType]*DetailedTestCaseExecutionsMapObjectStruct
	var tempDetailedTestCaseExecutionsObjectsMap map[DetailedTestCaseExecutionMapKeyType]*DetailedTestCaseExecutionsMapObjectStruct
	tempDetailedTestCaseExecutionsObjectsMapPtr = TestCaseExecutionsModel.DetailedTestCaseExecutionsObjectsMapPtr
	tempDetailedTestCaseExecutionsObjectsMap = *tempDetailedTestCaseExecutionsObjectsMapPtr

	var tempdetailedTestCaseExecutionsMapObjectPtr *DetailedTestCaseExecutionsMapObjectStruct
	var tempdetailedTestCaseExecutionsMapObject DetailedTestCaseExecutionsMapObjectStruct
	tempdetailedTestCaseExecutionsMapObjectPtr, existInMap = tempDetailedTestCaseExecutionsObjectsMap[DetailedTestCaseExecutionMapKeyType(testCaseExecutionsMapKey)]
	if existInMap == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                       "8fb40f4d-dad9-49e8-8d97-01c0fa2dde6c",
			"testCaseExecutionsMapKey": testCaseExecutionsMapKey,
		}).Error("Should never happen - 'testCaseExecutionsMapKey' could not be found in 'DetailedTestCaseExecutionsObjectsMap")

		return "", false
	}

	tempdetailedTestCaseExecutionsMapObject = *tempdetailedTestCaseExecutionsMapObjectPtr

	// Extract map with relation between TestInstruction and TestInstructionExecution for TestCaseExecution
	var tempRelationBetweenTestInstructionUuidAndTestInstructionExectuionUuidMapPtr *map[RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType]TestInstructionExecutionUuidType
	var tempRelationBetweenTestInstructionUuidAndTestInstructionExectuionUuidMap map[RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType]TestInstructionExecutionUuidType

	tempRelationBetweenTestInstructionUuidAndTestInstructionExectuionUuidMapPtr = tempdetailedTestCaseExecutionsMapObject.RelationBetweenTestInstructionUuidAndTestInstructionExecutionUuidMapPtr
	tempRelationBetweenTestInstructionUuidAndTestInstructionExectuionUuidMap = *tempRelationBetweenTestInstructionUuidAndTestInstructionExectuionUuidMapPtr

	testInstructionExecutionUuid, existInMap = tempRelationBetweenTestInstructionUuidAndTestInstructionExectuionUuidMap[testInstructionUuid]
	if existInMap == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                       "d49656ad-fed7-4c1c-a667-c191ab11dcda",
			"testInstructionUuid":      testInstructionUuid,
			"testCaseExecutionsMapKey": testCaseExecutionsMapKey,
		}).Error("Should never happen - 'testInstructionUuid' could not be found in 'RelationBetweenTestInstructionUuidAndTestInstructionExectuionUuidMap")

		return "", false
	}

	return testInstructionExecutionUuid, existInMap
}

// GetTestInstructionFromTestInstructionExecutionUuid
// Read from the TestCaseExecutions-Map and get the TestInstruction Uuid and Name based on TestInstructionExecutionUuid + Version (mapKey)
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) GetTestInstructionFromTestInstructionExecutionUuid(
	testCaseExecutionsMapKey TestCaseExecutionUuidType,
	testInstructionExecutionUuid TestInstructionExecutionUuidType,
	lockMap bool) (
	testInstructionUuid RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType,
	testInstructionName string,
	existInMap bool) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "b0eb55bd-80be-4609-bc12-5388d1387fd4",
	}).Debug("Incoming - 'GetTestInstructionFromTestInstructionExecutionUuid'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "de80f3dd-4c15-4bdc-9b3c-098d4cdcddf1",
	}).Debug("Outgoing - 'GetTestInstructionFromTestInstructionExecutionUuid'")

	if lockMap == true {
		// Lock Map for Reading
		testCaseExecutionsMapMutex.RLock()

		//UnLock Map
		defer testCaseExecutionsMapMutex.RUnlock()
	}

	// Check if Map i nil
	if TestCaseExecutionsModel.LatestTestCaseExecutionForEachTestCaseUuid.
		latestTestCaseExecutionForEachTestCaseUuidMap == nil {
		return "", "", false
	}

	// Read Map
	var tempDetailedTestCaseExecutionsObjectsMapPtr *map[DetailedTestCaseExecutionMapKeyType]*DetailedTestCaseExecutionsMapObjectStruct
	var tempDetailedTestCaseExecutionsObjectsMap map[DetailedTestCaseExecutionMapKeyType]*DetailedTestCaseExecutionsMapObjectStruct
	tempDetailedTestCaseExecutionsObjectsMapPtr = TestCaseExecutionsModel.DetailedTestCaseExecutionsObjectsMapPtr
	tempDetailedTestCaseExecutionsObjectsMap = *tempDetailedTestCaseExecutionsObjectsMapPtr

	var tempdetailedTestCaseExecutionsMapObjectPtr *DetailedTestCaseExecutionsMapObjectStruct
	var tempdetailedTestCaseExecutionsMapObject DetailedTestCaseExecutionsMapObjectStruct
	tempdetailedTestCaseExecutionsMapObjectPtr, existInMap = tempDetailedTestCaseExecutionsObjectsMap[DetailedTestCaseExecutionMapKeyType(testCaseExecutionsMapKey)]
	if existInMap == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                       "8934b0f3-e835-472e-8913-99a426eee2b8",
			"testCaseExecutionsMapKey": testCaseExecutionsMapKey,
		}).Error("Should never happen - 'testCaseExecutionsMapKey' could not be found in 'DetailedTestCaseExecutionsObjectsMap")

		return "", "", false
	}

	tempdetailedTestCaseExecutionsMapObject = *tempdetailedTestCaseExecutionsMapObjectPtr

	// Extract map with relation between TestInstructionExecution and TestInstruction for TestCaseExecution
	var tempRelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMapPtr *map[TestInstructionExecutionUuidType]RelationBetweenTestInstructionUuidAndTestInstructionExecutionStruct
	var tempRelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMap map[TestInstructionExecutionUuidType]RelationBetweenTestInstructionUuidAndTestInstructionExecutionStruct

	tempRelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMapPtr = tempdetailedTestCaseExecutionsMapObject.RelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMapPtr
	tempRelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMap = *tempRelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMapPtr

	var testInstructionObject RelationBetweenTestInstructionUuidAndTestInstructionExecutionStruct

	testInstructionObject, existInMap = tempRelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMap[testInstructionExecutionUuid]
	if existInMap == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                           "e93a6e66-20f8-45ac-8e0f-80c2bfb6e68f",
			"testCaseExecutionsMapKey":     testCaseExecutionsMapKey,
			"testInstructionExecutionUuid": testInstructionExecutionUuid,
		}).Error("Should never happen - 'testInstructionExecutionUuid' could not be found in 'RelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMapPtr")

		return "", "", false
	}

	return testInstructionObject.TestInstructionUuid, testInstructionObject.TestInstructionName, existInMap
}
