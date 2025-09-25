package testSuiteExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"sync"
)

var testSuiteExecutionsMapMutex = &sync.RWMutex{}

// InitiateTestSuiteExecutionsMap
// Add to the TestSuiteExecutions-Map
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) InitiateTestSuiteExecutionsMap() {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "a52f6fd8-8b67-493f-8d08-0d957ba35170",
	}).Debug("Incoming - 'InitiateTestSuiteExecutionsMap'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "fe601432-68fa-42ad-8390-391286d97893",
	}).Debug("Outgoing - 'InitiateTestSuiteExecutionsMap'")

	// Lock Map for Writing
	testSuiteExecutionsMapMutex.Lock()

	// Initiate map if it is not already done
	if testSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
		latestTestSuiteExecutionForEachTestSuiteUuidMap == nil {
		testSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
			latestTestSuiteExecutionForEachTestSuiteUuidMap = make(map[TestSuiteExecutionUuidType]*fenixExecutionServerGuiGrpcApi.
			TestSuiteExecutionsListMessage)
	}

	//UnLock Map
	testSuiteExecutionsMapMutex.Unlock()
}

// ReadFromTestSuiteExecutionsMap
// Read from the TestSuiteExecutions-Map
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) ReadFromTestSuiteExecutionsMap(
	testSuiteExecutionsMapKey TestSuiteExecutionUuidType) (
	testSuiteExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage,
	existInMap bool) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "01698693-b50e-4eeb-9057-326c70f52e5e",
	}).Debug("Incoming - 'ReadFromTestSuiteExecutionsMap'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "d654d90b-c5bc-428d-b1a9-013e95e4851d",
	}).Debug("Outgoing - 'ReadFromTestSuiteExecutionsMap'")

	// Lock Map for Reading
	testSuiteExecutionsMapMutex.RLock()

	//UnLock Map
	defer testSuiteExecutionsMapMutex.RUnlock()

	// Check if Map i nil
	if TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
		latestTestSuiteExecutionForEachTestSuiteUuidMap == nil {
		return nil, false
	}

	// Read Map
	testSuiteExecutionsListMessage, existInMap = TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
		latestTestSuiteExecutionForEachTestSuiteUuidMap[testSuiteExecutionsMapKey]

	return testSuiteExecutionsListMessage, existInMap
}

// ReadAllFromTestSuiteExecutionsMap
// Read all from the TestSuiteExecutions-Map
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) ReadAllFromTestSuiteExecutionsMap() (
	testSuiteExecutionsListMessage *[]*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "f11158ad-3495-42ab-85aa-68c7c2272c00",
	}).Debug("Incoming - 'ReadAllFromTestSuiteExecutionsMap'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "9589f7cc-5329-4bb8-b8ba-8582c214773e",
	}).Debug("Outgoing - 'ReadAllFromTestSuiteExecutionsMap'")

	// Lock Map for Reading
	testSuiteExecutionsMapMutex.RLock()

	//UnLock Map
	defer testSuiteExecutionsMapMutex.RUnlock()

	// Check if Map i nil
	if TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
		latestTestSuiteExecutionForEachTestSuiteUuidMap == nil {
		return nil
	}

	var tempTestSuiteExecutionsListMessage []*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage

	// Loop all items in map and add to response slice
	for _, testSuiteExecutionListMessage := range TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
		latestTestSuiteExecutionForEachTestSuiteUuidMap {
		tempTestSuiteExecutionsListMessage = append(tempTestSuiteExecutionsListMessage, testSuiteExecutionListMessage)
	}

	return &tempTestSuiteExecutionsListMessage
}

// GetNumberOfTestSuiteExecutionsRetrievedFromDatabase
// Read all from the TestSuiteExecutions-Map
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) GetNumberOfTestSuiteExecutionsRetrievedFromDatabase() (
	numberOfTestSuiteExecutionsRetrievedFromDatabase int) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "29fa838d-d52e-48bf-b399-7abf73cf4457",
	}).Debug("Incoming - 'GetNumberOfTestSuiteExecutionsRetrievedFromDatabase'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "27f1a1a9-bdbd-4fed-9ee4-764801be229a",
	}).Debug("Outgoing - 'GetNumberOfTestSuiteExecutionsRetrievedFromDatabase'")

	// Lock Map for Reading
	testSuiteExecutionsMapMutex.RLock()

	//UnLock Map
	defer testSuiteExecutionsMapMutex.RUnlock()

	var testSuiteExecutionsListMessage *[]*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage

	// Get all data retrieved and get hwo many they are
	testSuiteExecutionsListMessage = testSuiteExecutionsModel.ReadAllFromTestSuiteExecutionsMap()

	numberOfTestSuiteExecutionsRetrievedFromDatabase = len(*testSuiteExecutionsListMessage)

	return numberOfTestSuiteExecutionsRetrievedFromDatabase

}

// AddToTestSuiteExecutionsMap
// Add to the TestSuiteExecutions-Map
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) AddToTestSuiteExecutionsMap(
	testSuiteExecutionsMapKey TestSuiteExecutionUuidType,
	testSuiteExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "9984f6c6-996f-4039-b431-9a221ea36fa2",
	}).Debug("Incoming - 'AddToTestSuiteExecutionsMap'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "11d9fa55-7bbe-4d8b-9115-848dc37774fb",
	}).Debug("Outgoing - 'AddToTestSuiteExecutionsMap'")

	// Lock Map for Writing
	testSuiteExecutionsMapMutex.Lock()

	//UnLock Map
	defer testSuiteExecutionsMapMutex.Unlock()

	// Check if Map i nil
	if TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
		latestTestSuiteExecutionForEachTestSuiteUuidMap == nil {

		TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
			latestTestSuiteExecutionForEachTestSuiteUuidMap = make(
			map[TestSuiteExecutionUuidType]*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage)
	}

	// Save to TestSuiteExecutions-Map
	TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
		latestTestSuiteExecutionForEachTestSuiteUuidMap[testSuiteExecutionsMapKey] = testSuiteExecutionsListMessage

}

// DeleteFromTestSuiteExecutionsMap
// Delete from the TestSuiteExecutions-Map
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) DeleteFromTestSuiteExecutionsMap(
	testSuiteExecutionsMapKey TestSuiteExecutionUuidType) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "85632715-c3a4-435c-bd5e-3a9870f3b690",
	}).Debug("Incoming - 'DeleteFromTestSuiteExecutionsMap'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "d1f707ac-8965-4119-98d2-930aad1cea78",
	}).Debug("Outgoing - 'DeleteFromTestSuiteExecutionsMap'")

	// Lock Map for Writing
	testSuiteExecutionsMapMutex.Lock()

	//UnLock Map
	defer testSuiteExecutionsMapMutex.Unlock()

	// Check if Map i nil
	if TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.latestTestSuiteExecutionForEachTestSuiteUuidMap == nil {

		TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
			latestTestSuiteExecutionForEachTestSuiteUuidMap = make(
			map[TestSuiteExecutionUuidType]*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage)

		return
	}

	// Save to TestSuiteExecutions-Map
	delete(TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
		latestTestSuiteExecutionForEachTestSuiteUuidMap,
		testSuiteExecutionsMapKey)

}

// GetTestInstructionExecutionUuidFromTestInstructionUuid
// Read from the TestSuiteExecutions-Map and get the TestInstructionExecutionUuid + Version (mpKey) based on TestInstructionUuid
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) GetTestInstructionExecutionUuidFromTestInstructionUuid(
	testSuiteExecutionsMapKey TestSuiteExecutionUuidType,
	testInstructionUuid RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType) (
	testInstructionExecutionUuid TestInstructionExecutionUuidType,
	existInMap bool) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "8e69f6fe-f077-4403-832a-041abe690724",
	}).Debug("Incoming - 'GetTestInstructionExecutionUuidFromTestInstructionUuid'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "94b61a49-6506-4081-bcc6-b9776dd99b8c",
	}).Debug("Outgoing - 'GetTestInstructionExecutionUuidFromTestInstructionUuid'")

	// Lock Map for Reading
	//testSuiteExecutionsMapMutex.RLock()

	//UnLock Map
	//defer testSuiteExecutionsMapMutex.RUnlock()

	// Check if Map i nil
	if TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
		latestTestSuiteExecutionForEachTestSuiteUuidMap == nil {
		return "", false
	}

	// Read Map
	var tempDetailedTestSuiteExecutionsObjectsMapPtr *map[DetailedTestSuiteExecutionMapKeyType]*DetailedTestSuiteExecutionsMapObjectStruct
	var tempDetailedTestSuiteExecutionsObjectsMap map[DetailedTestSuiteExecutionMapKeyType]*DetailedTestSuiteExecutionsMapObjectStruct
	tempDetailedTestSuiteExecutionsObjectsMapPtr = TestSuiteExecutionsModel.DetailedTestSuiteExecutionsObjectsMapPtr
	tempDetailedTestSuiteExecutionsObjectsMap = *tempDetailedTestSuiteExecutionsObjectsMapPtr

	var tempdetailedTestSuiteExecutionsMapObjectPtr *DetailedTestSuiteExecutionsMapObjectStruct
	var tempdetailedTestSuiteExecutionsMapObject DetailedTestSuiteExecutionsMapObjectStruct
	tempdetailedTestSuiteExecutionsMapObjectPtr, existInMap = tempDetailedTestSuiteExecutionsObjectsMap[DetailedTestSuiteExecutionMapKeyType(testSuiteExecutionsMapKey)]
	if existInMap == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                        "89affc07-9efe-42ab-b7cf-d39a01525365",
			"testSuiteExecutionsMapKey": testSuiteExecutionsMapKey,
		}).Error("Should never happen - 'testSuiteExecutionsMapKey' could not be found in 'DetailedTestSuiteExecutionsObjectsMap")

		return "", false
	}

	tempdetailedTestSuiteExecutionsMapObject = *tempdetailedTestSuiteExecutionsMapObjectPtr

	// Extract map with relation between TestInstruction and TestInstructionExecution for TestSuiteExecution
	var tempRelationBetweenTestInstructionUuidAndTestInstructionExectuionUuidMapPtr *map[RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType]TestInstructionExecutionUuidType
	var tempRelationBetweenTestInstructionUuidAndTestInstructionExectuionUuidMap map[RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType]TestInstructionExecutionUuidType

	tempRelationBetweenTestInstructionUuidAndTestInstructionExectuionUuidMapPtr = tempdetailedTestSuiteExecutionsMapObject.RelationBetweenTestInstructionUuidAndTestInstructionExecutionUuidMapPtr
	tempRelationBetweenTestInstructionUuidAndTestInstructionExectuionUuidMap = *tempRelationBetweenTestInstructionUuidAndTestInstructionExectuionUuidMapPtr

	testInstructionExecutionUuid, existInMap = tempRelationBetweenTestInstructionUuidAndTestInstructionExectuionUuidMap[testInstructionUuid]
	if existInMap == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                        "4630f29e-b9b6-4ca1-a977-21114663ba88",
			"testInstructionUuid":       testInstructionUuid,
			"testSuiteExecutionsMapKey": testSuiteExecutionsMapKey,
		}).Error("Should never happen - 'testInstructionUuid' could not be found in 'RelationBetweenTestInstructionUuidAndTestInstructionExectuionUuidMap")

		return "", false
	}

	return testInstructionExecutionUuid, existInMap
}

// GetTestInstructionFromTestInstructionExecutionUuid
// Read from the TestSuiteExecutions-Map and get the TestInstruction Uuid and Name based on TestInstructionExecutionUuid + Version (mapKey)
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) GetTestInstructionFromTestInstructionExecutionUuid(
	testSuiteExecutionsMapKey TestSuiteExecutionUuidType,
	testInstructionExecutionUuid TestInstructionExecutionUuidType,
	lockMap bool) (
	testInstructionUuid RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType,
	testInstructionName string,
	existInMap bool) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "8c70226a-f65c-4c77-a04c-2c044235961e",
	}).Debug("Incoming - 'GetTestInstructionFromTestInstructionExecutionUuid'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "488c3fd4-90c4-4b5b-9224-f4f35f4cde68",
	}).Debug("Outgoing - 'GetTestInstructionFromTestInstructionExecutionUuid'")

	if lockMap == true {
		// Lock Map for Reading
		testSuiteExecutionsMapMutex.RLock()

		//UnLock Map
		defer testSuiteExecutionsMapMutex.RUnlock()
	}

	// Check if Map i nil
	if TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
		latestTestSuiteExecutionForEachTestSuiteUuidMap == nil {
		return "", "", false
	}

	// Read Map
	var tempDetailedTestSuiteExecutionsObjectsMapPtr *map[DetailedTestSuiteExecutionMapKeyType]*DetailedTestSuiteExecutionsMapObjectStruct
	var tempDetailedTestSuiteExecutionsObjectsMap map[DetailedTestSuiteExecutionMapKeyType]*DetailedTestSuiteExecutionsMapObjectStruct
	tempDetailedTestSuiteExecutionsObjectsMapPtr = TestSuiteExecutionsModel.DetailedTestSuiteExecutionsObjectsMapPtr
	tempDetailedTestSuiteExecutionsObjectsMap = *tempDetailedTestSuiteExecutionsObjectsMapPtr

	var tempdetailedTestSuiteExecutionsMapObjectPtr *DetailedTestSuiteExecutionsMapObjectStruct
	var tempdetailedTestSuiteExecutionsMapObject DetailedTestSuiteExecutionsMapObjectStruct
	tempdetailedTestSuiteExecutionsMapObjectPtr, existInMap = tempDetailedTestSuiteExecutionsObjectsMap[DetailedTestSuiteExecutionMapKeyType(testSuiteExecutionsMapKey)]
	if existInMap == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                        "7d35b326-73d0-4d04-9352-9c10281026ea",
			"testSuiteExecutionsMapKey": testSuiteExecutionsMapKey,
		}).Error("Should never happen - 'testSuiteExecutionsMapKey' could not be found in 'DetailedTestSuiteExecutionsObjectsMap")

		return "", "", false
	}

	tempdetailedTestSuiteExecutionsMapObject = *tempdetailedTestSuiteExecutionsMapObjectPtr

	// Extract map with relation between TestInstructionExecution and TestInstruction for TestSuiteExecution
	var tempRelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMapPtr *map[TestInstructionExecutionUuidType]RelationBetweenTestInstructionUuidAndTestInstructionExecutionStruct
	var tempRelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMap map[TestInstructionExecutionUuidType]RelationBetweenTestInstructionUuidAndTestInstructionExecutionStruct

	tempRelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMapPtr = tempdetailedTestSuiteExecutionsMapObject.RelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMapPtr
	tempRelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMap = *tempRelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMapPtr

	var testInstructionObject RelationBetweenTestInstructionUuidAndTestInstructionExecutionStruct

	testInstructionObject, existInMap = tempRelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMap[testInstructionExecutionUuid]
	if existInMap == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                           "fad82e9d-470d-4b8e-9cbf-746e45029362",
			"testSuiteExecutionsMapKey":    testSuiteExecutionsMapKey,
			"testInstructionExecutionUuid": testInstructionExecutionUuid,
		}).Error("Should never happen - 'testInstructionExecutionUuid' could not be found in 'RelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMapPtr")

		return "", "", false
	}

	return testInstructionObject.TestInstructionUuid, testInstructionObject.TestInstructionName, existInMap
}
