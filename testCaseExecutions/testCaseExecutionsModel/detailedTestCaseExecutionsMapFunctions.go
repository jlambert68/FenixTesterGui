package testCaseExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"sync"
)

var detailedTestCaseExecutionsMapMutex = &sync.RWMutex{}

// ReadFromDetailedTestCaseExecutionsMap
// Read from the DetailedTestCaseExecutions-Map
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) ReadFromDetailedTestCaseExecutionsMap(
	detailedTestCaseExecutionMapKey DetailedTestCaseExecutionMapKeyType) (
	detailedTestCaseExecution *fenixExecutionServerGuiGrpcApi.TestCaseExecutionResponseMessage,
	existInMap bool) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "4c333c7d-7df5-443d-b3cc-e155eba3d0e1",
	}).Debug("Incoming - 'ReadFromDetailedTestCaseExecutionsMap'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "7e329468-1335-42ce-b055-e6d5bf2e7e10",
	}).Debug("Outgoing - 'ReadFromDetailedTestCaseExecutionsMap'")

	// Lock Map for Writing
	detailedTestCaseExecutionsMapMutex.Lock()

	//UnLock Map
	defer detailedTestCaseExecutionsMapMutex.Unlock()

	// Check if Map i nil
	if TestCaseExecutionsModel.DetailedTestCaseExecutionsObjectsMapPtr == nil {

		return nil, false

	}

	// Get the map from the map pointer
	var tempDetailedTestCaseExecutionsObjectsMap map[DetailedTestCaseExecutionMapKeyType]*DetailedTestCaseExecutionsMapObjectStruct
	tempDetailedTestCaseExecutionsObjectsMap = *TestCaseExecutionsModel.DetailedTestCaseExecutionsObjectsMapPtr

	// Check of DetailedTestCaseExecution exist in Map
	var tempDetailedTestCaseExecutionsMapObjectPtr *DetailedTestCaseExecutionsMapObjectStruct
	var tempDetailedTestCaseExecutionsMapObject DetailedTestCaseExecutionsMapObjectStruct
	tempDetailedTestCaseExecutionsMapObjectPtr, existInMap = tempDetailedTestCaseExecutionsObjectsMap[detailedTestCaseExecutionMapKey]

	// If it exists then return the DetailedTestCaseExecution-object
	if existInMap == true {

		// Get the TestCaseExecutionObject from pointer
		tempDetailedTestCaseExecutionsMapObject = *tempDetailedTestCaseExecutionsMapObjectPtr

		if tempDetailedTestCaseExecutionsMapObject.DetailedTestCaseExecution != nil {

			// Exists so release Map-lock and then LockOn Update for this object
			if tempDetailedTestCaseExecutionsMapObject.WaitingForDatabaseUpdateMutex.TryLock() == false {

				// Object locked so unlock Map and the lock on to the object
				detailedTestCaseExecutionsMapMutex.Unlock()
				tempDetailedTestCaseExecutionsMapObject.WaitingForDatabaseUpdateMutex.Lock()

				// Lock Map when DetailedTestCaseExecution object has been released
				detailedTestCaseExecutionsMapMutex.Lock()

				// New update of DetailedTestCaseExecution has been received
				tempDetailedTestCaseExecutionsMapObjectPtr, _ = tempDetailedTestCaseExecutionsObjectsMap[detailedTestCaseExecutionMapKey]
				// Get the TestCaseExecutionObject from pointer
				tempDetailedTestCaseExecutionsMapObject = *tempDetailedTestCaseExecutionsMapObjectPtr

				// Unlock locked object
				defer tempDetailedTestCaseExecutionsMapObject.WaitingForDatabaseUpdateMutex.Unlock()
			}

			//return object
			return tempDetailedTestCaseExecutionsMapObject.DetailedTestCaseExecution, true

		} else {
			// Object is nil
			return nil, false

		}
	} else {
		// Doesn't exist in map
		return nil, false

	}

}

// GetNumberOfDetailedTestCaseExecutionsRetrievedFromDatabase
// Read all from the DetailedTestCaseExecutions-Map
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) GetNumberOfDetailedTestCaseExecutionsRetrievedFromDatabase() (
	numberOfDetailedTestCaseExecutionsRetrievedFromDatabase int) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "305226a6-e3d2-40bb-b48e-012f6725c19a",
	}).Debug("Incoming - 'GetNumberOfTestCaseExecutionsRetrievedFromDatabase'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "34d7764b-a4e8-4985-942e-f2b0e5efd029",
	}).Debug("Outgoing - 'GetNumberOfTestCaseExecutionsRetrievedFromDatabase'")

	// Lock Map for Reading
	detailedTestCaseExecutionsMapMutex.RLock()

	//UnLock Map
	defer detailedTestCaseExecutionsMapMutex.RUnlock()

	// Check if Map i nil
	if TestCaseExecutionsModel.DetailedTestCaseExecutionsObjectsMapPtr == nil {

		return 0

	}

	// Get the map from the map pointer
	var tempDetailedTestCaseExecutionsObjectsMap map[DetailedTestCaseExecutionMapKeyType]*DetailedTestCaseExecutionsMapObjectStruct
	tempDetailedTestCaseExecutionsObjectsMap = *TestCaseExecutionsModel.DetailedTestCaseExecutionsObjectsMapPtr

	numberOfDetailedTestCaseExecutionsRetrievedFromDatabase = len(tempDetailedTestCaseExecutionsObjectsMap)

	return numberOfDetailedTestCaseExecutionsRetrievedFromDatabase

}

// AddToDetailedTestCaseExecutionsMap
// Add to the DetailedTestCaseExecutions-Map
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) AddToDetailedTestCaseExecutionsMap(
	detailedTestCaseExecutionMapKey DetailedTestCaseExecutionMapKeyType,
	detailedTestCaseExecution *fenixExecutionServerGuiGrpcApi.TestCaseExecutionResponseMessage) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "17bfb3bf-8d78-4da3-8554-00507b33629e",
	}).Debug("Incoming - 'AddToDetailedTestCaseExecutionsMap'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "996562c0-8307-44f6-97cb-d0b94d71d266",
	}).Debug("Outgoing - 'AddToDetailedTestCaseExecutionsMap'")

	var existInMap bool

	// Lock Map for Writing
	detailedTestCaseExecutionsMapMutex.Lock()

	//UnLock Map
	defer detailedTestCaseExecutionsMapMutex.Unlock()

	// Check if Map i nil
	if TestCaseExecutionsModel.DetailedTestCaseExecutionsObjectsMapPtr == nil {

		var tempDetailedTestCaseExecutionsObjectsMap map[DetailedTestCaseExecutionMapKeyType]*DetailedTestCaseExecutionsMapObjectStruct
		tempDetailedTestCaseExecutionsObjectsMap = make(map[DetailedTestCaseExecutionMapKeyType]*DetailedTestCaseExecutionsMapObjectStruct)

		TestCaseExecutionsModel.DetailedTestCaseExecutionsObjectsMapPtr = &tempDetailedTestCaseExecutionsObjectsMap

	}

	// Get the map from the map pointer
	var tempDetailedTestCaseExecutionsObjectsMap map[DetailedTestCaseExecutionMapKeyType]*DetailedTestCaseExecutionsMapObjectStruct
	tempDetailedTestCaseExecutionsObjectsMap = *TestCaseExecutionsModel.DetailedTestCaseExecutionsObjectsMapPtr

	// Check of DetailedTestCaseExecution exist in Map
	var tempDetailedTestCaseExecutionsMapObjectPtr *DetailedTestCaseExecutionsMapObjectStruct
	var tempDetailedTestCaseExecutionsMapObject DetailedTestCaseExecutionsMapObjectStruct
	tempDetailedTestCaseExecutionsMapObjectPtr, existInMap = tempDetailedTestCaseExecutionsObjectsMap[detailedTestCaseExecutionMapKey]

	// Set that refresh of DetailedTestCaseExecution is ongoing
	if existInMap == true {
		// Already exist,so use that object by Get the TestCaseExecutionObject from pointer
		tempDetailedTestCaseExecutionsMapObjectPtr.DetailedTestCaseExecution = detailedTestCaseExecution
		//tempDetailedTestCaseExecutionsMapObject = *tempDetailedTestCaseExecutionsMapObjectPtr
		//tempDetailedTestCaseExecutionsMapObject.DetailedTestCaseExecution = detailedTestCaseExecution

	} else {
		// Object doesn't exist so create a new object and store in map
		tempDetailedTestCaseExecutionsMapObject = DetailedTestCaseExecutionsMapObjectStruct{
			DetailedTestCaseExecution:     detailedTestCaseExecution,
			WaitingForDatabaseUpdate:      false,
			WaitingForDatabaseUpdateMutex: &sync.RWMutex{},
		}

		tempDetailedTestCaseExecutionsObjectsMap[detailedTestCaseExecutionMapKey] = &tempDetailedTestCaseExecutionsMapObject
	}

}

// DeleteFromDetailedTestCaseExecutionsMap
// Delete from the DetailedTestCaseExecutions-Map
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) DeleteFromDetailedTestCaseExecutionsMap(
	detailedTestCaseExecutionMapKey DetailedTestCaseExecutionMapKeyType) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "04d642b9-079a-46e0-a67f-6295a2c2efd9",
	}).Debug("Incoming - 'DeleteFromDetailedTestCaseExecutionsMap'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "6a00605b-7c5d-4a37-8184-e781a23e527f",
	}).Debug("Outgoing - 'DeleteFromDetailedTestCaseExecutionsMap'")

	var existInMap bool

	// Lock Map for Writing
	detailedTestCaseExecutionsMapMutex.Lock()

	//UnLock Map
	defer detailedTestCaseExecutionsMapMutex.Unlock()

	// Check if Map i nil
	if TestCaseExecutionsModel.DetailedTestCaseExecutionsObjectsMapPtr == nil {

		// it doesn't exist so just return
		return

	}

	// Get the map from the map pointer
	var tempDetailedTestCaseExecutionsObjectsMap map[DetailedTestCaseExecutionMapKeyType]*DetailedTestCaseExecutionsMapObjectStruct
	tempDetailedTestCaseExecutionsObjectsMap = *TestCaseExecutionsModel.DetailedTestCaseExecutionsObjectsMapPtr

	// Check of DetailedTestCaseExecution exist in Map
	_, existInMap = tempDetailedTestCaseExecutionsObjectsMap[detailedTestCaseExecutionMapKey]

	// Delete DetailedTestCaseExecution from Map, if it exist
	if existInMap == true {

		// It exists,so just Delete the TestCaseExecutionObject from map
		delete(tempDetailedTestCaseExecutionsObjectsMap, detailedTestCaseExecutionMapKey)

	} else {

		// it doesn't exist so just return
		return
	}

}

// SetFlagRefreshOngoingOfDetailedTestCaseExecution
// Set the flag there is an ongoing refresh of the DetailedTestCaseExecution-data
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) SetFlagRefreshOngoingOfDetailedTestCaseExecution(
	detailedTestCaseExecutionMapKey DetailedTestCaseExecutionMapKeyType) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "528eda7e-69e7-48a1-9926-0fa154490e3e",
	}).Debug("Incoming - 'SetFlagRefreshOngoingOfDetailedTestCaseExecution'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "131b423d-0350-496b-8f5c-117bcdff7b3c",
	}).Debug("Outgoing - 'SetFlagRefreshOngoingOfDetailedTestCaseExecution'")

	var existInMap bool

	// Lock Map for Writing
	detailedTestCaseExecutionsMapMutex.Lock()

	//UnLock Map
	defer detailedTestCaseExecutionsMapMutex.Unlock()

	// Check if Map i nil
	if TestCaseExecutionsModel.DetailedTestCaseExecutionsObjectsMapPtr == nil {

		var tempDetailedTestCaseExecutionsObjectsMap map[DetailedTestCaseExecutionMapKeyType]*DetailedTestCaseExecutionsMapObjectStruct
		tempDetailedTestCaseExecutionsObjectsMap = make(map[DetailedTestCaseExecutionMapKeyType]*DetailedTestCaseExecutionsMapObjectStruct)

		TestCaseExecutionsModel.DetailedTestCaseExecutionsObjectsMapPtr = &tempDetailedTestCaseExecutionsObjectsMap

	}

	// Get the map from the map pointer
	var tempDetailedTestCaseExecutionsObjectsMap map[DetailedTestCaseExecutionMapKeyType]*DetailedTestCaseExecutionsMapObjectStruct
	tempDetailedTestCaseExecutionsObjectsMap = *TestCaseExecutionsModel.DetailedTestCaseExecutionsObjectsMapPtr

	// Check of DetailedTestCaseExecution exist in Map
	var tempDetailedTestCaseExecutionsMapObjectPtr *DetailedTestCaseExecutionsMapObjectStruct
	var tempDetailedTestCaseExecutionsMapObject DetailedTestCaseExecutionsMapObjectStruct
	tempDetailedTestCaseExecutionsMapObjectPtr, existInMap = tempDetailedTestCaseExecutionsObjectsMap[detailedTestCaseExecutionMapKey]

	// Set that refresh of DetailedTestCaseExecution is ongoing
	if existInMap == true {

		// Get the Map
		tempDetailedTestCaseExecutionsMapObject = *tempDetailedTestCaseExecutionsMapObjectPtr

		// Already exist,so use that object by Get the TestCaseExecutionObject from pointer
		tempDetailedTestCaseExecutionsMapObject.WaitingForDatabaseUpdate = true
		tempDetailedTestCaseExecutionsMapObject.WaitingForDatabaseUpdateMutex.Lock()

	} else {
		// Object doesn't exist so create a new object and store in map
		tempDetailedTestCaseExecutionsMapObject = DetailedTestCaseExecutionsMapObjectStruct{
			DetailedTestCaseExecution:     nil,
			WaitingForDatabaseUpdate:      true,
			WaitingForDatabaseUpdateMutex: &sync.RWMutex{},
		}

		tempDetailedTestCaseExecutionsMapObject.WaitingForDatabaseUpdateMutex.Lock()

		tempDetailedTestCaseExecutionsObjectsMap[detailedTestCaseExecutionMapKey] = &tempDetailedTestCaseExecutionsMapObject
	}

}

// ClearFlagRefreshOngoingOfDetailedTestCaseExecution
// Clear the flag there is an ongoing refresh of the DetailedTestCaseExecution-data
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) ClearFlagRefreshOngoingOfDetailedTestCaseExecution(
	detailedTestCaseExecutionMapKey DetailedTestCaseExecutionMapKeyType) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "0e504515-d6f8-47f1-8981-4b681f4199c2",
	}).Debug("Incoming - 'ClearFlagRefreshOngoingOfDetailedTestCaseExecution'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "5c4c66a1-295f-4b53-8cf8-39225452416c",
	}).Debug("Outgoing - 'ClearFlagRefreshOngoingOfDetailedTestCaseExecution'")

	var existInMap bool

	// Lock Map for Writing
	detailedTestCaseExecutionsMapMutex.Lock()

	//UnLock Map
	defer detailedTestCaseExecutionsMapMutex.Unlock()

	// Check if Map i nil
	if TestCaseExecutionsModel.DetailedTestCaseExecutionsObjectsMapPtr == nil {

		var tempDetailedTestCaseExecutionsObjectsMap map[DetailedTestCaseExecutionMapKeyType]*DetailedTestCaseExecutionsMapObjectStruct
		tempDetailedTestCaseExecutionsObjectsMap = make(map[DetailedTestCaseExecutionMapKeyType]*DetailedTestCaseExecutionsMapObjectStruct)

		TestCaseExecutionsModel.DetailedTestCaseExecutionsObjectsMapPtr = &tempDetailedTestCaseExecutionsObjectsMap

	}

	// Get the map from the map pointer
	var tempDetailedTestCaseExecutionsObjectsMap map[DetailedTestCaseExecutionMapKeyType]*DetailedTestCaseExecutionsMapObjectStruct
	tempDetailedTestCaseExecutionsObjectsMap = *TestCaseExecutionsModel.DetailedTestCaseExecutionsObjectsMapPtr

	// Check of DetailedTestCaseExecution exist in Map
	var tempDetailedTestCaseExecutionsMapObjectPtr *DetailedTestCaseExecutionsMapObjectStruct
	var tempDetailedTestCaseExecutionsMapObject DetailedTestCaseExecutionsMapObjectStruct
	tempDetailedTestCaseExecutionsMapObjectPtr, existInMap = tempDetailedTestCaseExecutionsObjectsMap[detailedTestCaseExecutionMapKey]

	// Set that refresh of DetailedTestCaseExecution is ongoing
	if existInMap == true {
		// Already exist,so use that object by Get the TestCaseExecutionObject from pointer
		// Get the Map
		tempDetailedTestCaseExecutionsMapObject = *tempDetailedTestCaseExecutionsMapObjectPtr

		// Already exist,so use that object by Get the TestCaseExecutionObject from pointer
		tempDetailedTestCaseExecutionsMapObject.WaitingForDatabaseUpdate = false
		tempDetailedTestCaseExecutionsMapObject.WaitingForDatabaseUpdateMutex.Unlock()
	} else {
		// Object doesn't exist so create a new object and store in map
		tempDetailedTestCaseExecutionsMapObject = DetailedTestCaseExecutionsMapObjectStruct{
			DetailedTestCaseExecution:     nil,
			WaitingForDatabaseUpdate:      true,
			WaitingForDatabaseUpdateMutex: &sync.RWMutex{},
		}

		tempDetailedTestCaseExecutionsMapObject.WaitingForDatabaseUpdateMutex.Unlock()

		tempDetailedTestCaseExecutionsObjectsMap[detailedTestCaseExecutionMapKey] = &tempDetailedTestCaseExecutionsMapObject
	}

}
