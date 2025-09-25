package testSuiteExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	"fmt"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"sync"
)

var detailedTestSuiteExecutionsMapMutex = &sync.RWMutex{}

// ReadFromDetailedTestSuiteExecutionsMap
// Read from the DetailedTestSuiteExecutions-Map
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) ReadFromDetailedTestSuiteExecutionsMap(
	detailedTestSuiteExecutionMapKey DetailedTestSuiteExecutionMapKeyType) (
	detailedTestSuiteExecution *fenixExecutionServerGuiGrpcApi.TestSuiteExecutionResponseMessage,
	existInMap bool) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "c21250f0-6ad5-4bfa-9d3b-1485193d024e",
	}).Debug("Incoming - 'ReadFromDetailedTestSuiteExecutionsMap'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "a3e45792-7a32-42c0-9ad6-b781ff403702",
	}).Debug("Outgoing - 'ReadFromDetailedTestSuiteExecutionsMap'")

	// Lock Map for Writing
	detailedTestSuiteExecutionsMapMutex.Lock()

	//UnLock Map
	defer detailedTestSuiteExecutionsMapMutex.Unlock()

	// Check if Map i nil
	if TestSuiteExecutionsModel.DetailedTestSuiteExecutionsObjectsMapPtr == nil {

		return nil, false

	}

	// Get the map from the map pointer
	var tempDetailedTestSuiteExecutionsObjectsMap map[DetailedTestSuiteExecutionMapKeyType]*DetailedTestSuiteExecutionsMapObjectStruct
	tempDetailedTestSuiteExecutionsObjectsMap = *TestSuiteExecutionsModel.DetailedTestSuiteExecutionsObjectsMapPtr

	// Check of DetailedTestSuiteExecution exist in Map
	var tempDetailedTestSuiteExecutionsMapObjectPtr *DetailedTestSuiteExecutionsMapObjectStruct
	var tempDetailedTestSuiteExecutionsMapObject DetailedTestSuiteExecutionsMapObjectStruct
	tempDetailedTestSuiteExecutionsMapObjectPtr, existInMap = tempDetailedTestSuiteExecutionsObjectsMap[detailedTestSuiteExecutionMapKey]

	// If it exists then return the DetailedTestSuiteExecution-object
	if existInMap == true {

		// Get the TestSuiteExecutionObject from pointer
		tempDetailedTestSuiteExecutionsMapObject = *tempDetailedTestSuiteExecutionsMapObjectPtr

		if tempDetailedTestSuiteExecutionsMapObject.DetailedTestSuiteExecution != nil {

			// Exists so release Map-lock and then LockOn Update for this object
			if tempDetailedTestSuiteExecutionsMapObject.WaitingForDatabaseUpdateMutex.TryLock() == false {

				// Object locked so unlock Map and the lock on to the object
				detailedTestSuiteExecutionsMapMutex.Unlock()
				fmt.Println("tempDetailedTestSuiteExecutionsMapObject.WaitingForDatabaseUpdateMutex.Lock() - 01")
				tempDetailedTestSuiteExecutionsMapObject.WaitingForDatabaseUpdateMutex.Lock()

				// Lock Map when DetailedTestSuiteExecution object has been released
				detailedTestSuiteExecutionsMapMutex.Lock()

				// New update of DetailedTestSuiteExecution has been received
				tempDetailedTestSuiteExecutionsMapObjectPtr, _ = tempDetailedTestSuiteExecutionsObjectsMap[detailedTestSuiteExecutionMapKey]
				// Get the TestSuiteExecutionObject from pointer
				tempDetailedTestSuiteExecutionsMapObject = *tempDetailedTestSuiteExecutionsMapObjectPtr

				// Unlock locked object
				defer tempDetailedTestSuiteExecutionsMapObject.WaitingForDatabaseUpdateMutex.Unlock()
				defer fmt.Println("tempDetailedTestSuiteExecutionsMapObject.WaitingForDatabaseUpdateMutex.Lock() - 03")
			}

			//return object
			return tempDetailedTestSuiteExecutionsMapObject.DetailedTestSuiteExecution, true

		} else {
			// Object is nil
			return nil, false

		}
	} else {
		// Doesn't exist in map
		return nil, false

	}

}

// GetNumberOfDetailedTestSuiteExecutionsRetrievedFromDatabase
// Read all from the DetailedTestSuiteExecutions-Map
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) GetNumberOfDetailedTestSuiteExecutionsRetrievedFromDatabase() (
	numberOfDetailedTestSuiteExecutionsRetrievedFromDatabase int) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "5d2cdbfd-79ae-4e37-8829-8cedb9f9be74",
	}).Debug("Incoming - 'GetNumberOfTestSuiteExecutionsRetrievedFromDatabase'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "85593e4c-3b09-4c0d-870c-aba0bab8c1b5",
	}).Debug("Outgoing - 'GetNumberOfTestSuiteExecutionsRetrievedFromDatabase'")

	// Lock Map for Reading
	detailedTestSuiteExecutionsMapMutex.RLock()

	//UnLock Map
	defer detailedTestSuiteExecutionsMapMutex.RUnlock()

	// Check if Map i nil
	if TestSuiteExecutionsModel.DetailedTestSuiteExecutionsObjectsMapPtr == nil {

		return 0

	}

	// Get the map from the map pointer
	var tempDetailedTestSuiteExecutionsObjectsMap map[DetailedTestSuiteExecutionMapKeyType]*DetailedTestSuiteExecutionsMapObjectStruct
	tempDetailedTestSuiteExecutionsObjectsMap = *TestSuiteExecutionsModel.DetailedTestSuiteExecutionsObjectsMapPtr

	numberOfDetailedTestSuiteExecutionsRetrievedFromDatabase = len(tempDetailedTestSuiteExecutionsObjectsMap)

	return numberOfDetailedTestSuiteExecutionsRetrievedFromDatabase

}

// AddToDetailedTestSuiteExecutionsMap
// Add to the DetailedTestSuiteExecutions-Map
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) AddToDetailedTestSuiteExecutionsMap(
	detailedTestSuiteExecutionMapKey DetailedTestSuiteExecutionMapKeyType,
	detailedTestSuiteExecution *fenixExecutionServerGuiGrpcApi.TestSuiteExecutionResponseMessage) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "4271e32d-8a60-4e1a-83e3-309a9ff1a414",
	}).Debug("Incoming - 'AddToDetailedTestSuiteExecutionsMap'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "5ecb73a8-3cec-4319-b7aa-ab2bf7ac7709",
	}).Debug("Outgoing - 'AddToDetailedTestSuiteExecutionsMap'")

	var existInMap bool

	// Lock Map for Writing
	detailedTestSuiteExecutionsMapMutex.Lock()

	//UnLock Map
	defer detailedTestSuiteExecutionsMapMutex.Unlock()

	// Check if Map i nil
	if TestSuiteExecutionsModel.DetailedTestSuiteExecutionsObjectsMapPtr == nil {

		var tempDetailedTestSuiteExecutionsObjectsMap map[DetailedTestSuiteExecutionMapKeyType]*DetailedTestSuiteExecutionsMapObjectStruct
		tempDetailedTestSuiteExecutionsObjectsMap = make(map[DetailedTestSuiteExecutionMapKeyType]*DetailedTestSuiteExecutionsMapObjectStruct)

		TestSuiteExecutionsModel.DetailedTestSuiteExecutionsObjectsMapPtr = &tempDetailedTestSuiteExecutionsObjectsMap

	}

	// Get the map from the map pointer
	var tempDetailedTestSuiteExecutionsObjectsMap map[DetailedTestSuiteExecutionMapKeyType]*DetailedTestSuiteExecutionsMapObjectStruct
	tempDetailedTestSuiteExecutionsObjectsMap = *TestSuiteExecutionsModel.DetailedTestSuiteExecutionsObjectsMapPtr

	// Check of DetailedTestSuiteExecution exist in Map
	var tempDetailedTestSuiteExecutionsMapObjectPtr *DetailedTestSuiteExecutionsMapObjectStruct
	var tempDetailedTestSuiteExecutionsMapObject DetailedTestSuiteExecutionsMapObjectStruct
	tempDetailedTestSuiteExecutionsMapObjectPtr, existInMap = tempDetailedTestSuiteExecutionsObjectsMap[detailedTestSuiteExecutionMapKey]

	// Set that refresh of DetailedTestSuiteExecution is ongoing
	if existInMap == true {
		// Already exist,so use that object by Get the TestSuiteExecutionObject from pointer
		tempDetailedTestSuiteExecutionsMapObjectPtr.DetailedTestSuiteExecution = detailedTestSuiteExecution
		//tempDetailedTestSuiteExecutionsMapObject = *tempDetailedTestSuiteExecutionsMapObjectPtr
		//tempDetailedTestSuiteExecutionsMapObject.DetailedTestSuiteExecution = detailedTestSuiteExecution

	} else {
		// Object doesn't exist so create a new object and store in map
		tempDetailedTestSuiteExecutionsMapObject = DetailedTestSuiteExecutionsMapObjectStruct{
			DetailedTestSuiteExecution:    detailedTestSuiteExecution,
			WaitingForDatabaseUpdate:      false,
			WaitingForDatabaseUpdateMutex: &sync.RWMutex{},
		}

		tempDetailedTestSuiteExecutionsObjectsMap[detailedTestSuiteExecutionMapKey] = &tempDetailedTestSuiteExecutionsMapObject
	}

}

// DeleteFromDetailedTestSuiteExecutionsMap
// Delete from the DetailedTestSuiteExecutions-Map
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) DeleteFromDetailedTestSuiteExecutionsMap(
	detailedTestSuiteExecutionMapKey DetailedTestSuiteExecutionMapKeyType) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "fbf30dc5-14d0-4fa2-ab66-2969cfaf50f0",
	}).Debug("Incoming - 'DeleteFromDetailedTestSuiteExecutionsMap'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "a53464f7-6c38-4d06-b2a9-9982e60860d3",
	}).Debug("Outgoing - 'DeleteFromDetailedTestSuiteExecutionsMap'")

	var existInMap bool

	// Lock Map for Writing
	detailedTestSuiteExecutionsMapMutex.Lock()

	//UnLock Map
	defer detailedTestSuiteExecutionsMapMutex.Unlock()

	// Check if Map i nil
	if TestSuiteExecutionsModel.DetailedTestSuiteExecutionsObjectsMapPtr == nil {

		// it doesn't exist so just return
		return

	}

	// Get the map from the map pointer
	var tempDetailedTestSuiteExecutionsObjectsMap map[DetailedTestSuiteExecutionMapKeyType]*DetailedTestSuiteExecutionsMapObjectStruct
	tempDetailedTestSuiteExecutionsObjectsMap = *TestSuiteExecutionsModel.DetailedTestSuiteExecutionsObjectsMapPtr

	// Check of DetailedTestSuiteExecution exist in Map
	_, existInMap = tempDetailedTestSuiteExecutionsObjectsMap[detailedTestSuiteExecutionMapKey]

	// Delete DetailedTestSuiteExecution from Map, if it exist
	if existInMap == true {

		// It exists,so just Delete the TestSuiteExecutionObject from map
		delete(tempDetailedTestSuiteExecutionsObjectsMap, detailedTestSuiteExecutionMapKey)

	} else {

		// it doesn't exist so just return
		return
	}

}

// SetFlagRefreshOngoingOfDetailedTestSuiteExecution
// Set the flag there is an ongoing refresh of the DetailedTestSuiteExecution-data
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) SetFlagRefreshOngoingOfDetailedTestSuiteExecution(
	detailedTestSuiteExecutionMapKey DetailedTestSuiteExecutionMapKeyType) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "31b89222-53a7-4533-b046-5af2df15532b",
	}).Debug("Incoming - 'SetFlagRefreshOngoingOfDetailedTestSuiteExecution'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "2ffba728-c017-4abd-afe0-5a3ce19b6ee9",
	}).Debug("Outgoing - 'SetFlagRefreshOngoingOfDetailedTestSuiteExecution'")

	var existInMap bool

	// Lock Map for Writing
	detailedTestSuiteExecutionsMapMutex.Lock()

	//UnLock Map
	defer detailedTestSuiteExecutionsMapMutex.Unlock()

	// Check if Map i nil
	if TestSuiteExecutionsModel.DetailedTestSuiteExecutionsObjectsMapPtr == nil {

		var tempDetailedTestSuiteExecutionsObjectsMap map[DetailedTestSuiteExecutionMapKeyType]*DetailedTestSuiteExecutionsMapObjectStruct
		tempDetailedTestSuiteExecutionsObjectsMap = make(map[DetailedTestSuiteExecutionMapKeyType]*DetailedTestSuiteExecutionsMapObjectStruct)

		TestSuiteExecutionsModel.DetailedTestSuiteExecutionsObjectsMapPtr = &tempDetailedTestSuiteExecutionsObjectsMap

	}

	// Get the map from the map pointer
	var tempDetailedTestSuiteExecutionsObjectsMap map[DetailedTestSuiteExecutionMapKeyType]*DetailedTestSuiteExecutionsMapObjectStruct
	tempDetailedTestSuiteExecutionsObjectsMap = *TestSuiteExecutionsModel.DetailedTestSuiteExecutionsObjectsMapPtr

	// Check of DetailedTestSuiteExecution exist in Map
	var tempDetailedTestSuiteExecutionsMapObjectPtr *DetailedTestSuiteExecutionsMapObjectStruct
	var tempDetailedTestSuiteExecutionsMapObject DetailedTestSuiteExecutionsMapObjectStruct
	tempDetailedTestSuiteExecutionsMapObjectPtr, existInMap = tempDetailedTestSuiteExecutionsObjectsMap[detailedTestSuiteExecutionMapKey]

	// Set that refresh of DetailedTestSuiteExecution is ongoing
	if existInMap == true {

		// Get the Map
		tempDetailedTestSuiteExecutionsMapObject = *tempDetailedTestSuiteExecutionsMapObjectPtr

		// Already exist,so use that object by Get the TestSuiteExecutionObject from pointer
		tempDetailedTestSuiteExecutionsMapObject.WaitingForDatabaseUpdate = true
		fmt.Println("tempDetailedTestSuiteExecutionsMapObject.WaitingForDatabaseUpdateMutex.Lock() - 02")
		tempDetailedTestSuiteExecutionsMapObject.WaitingForDatabaseUpdateMutex.Lock()

	} else {
		// Object doesn't exist so create a new object and store in map
		tempDetailedTestSuiteExecutionsMapObject = DetailedTestSuiteExecutionsMapObjectStruct{
			DetailedTestSuiteExecution:    nil,
			WaitingForDatabaseUpdate:      true,
			WaitingForDatabaseUpdateMutex: &sync.RWMutex{},
		}

		fmt.Println("tempDetailedTestSuiteExecutionsMapObject.WaitingForDatabaseUpdateMutex.Lock() - 03")
		tempDetailedTestSuiteExecutionsMapObject.WaitingForDatabaseUpdateMutex.Lock()

		tempDetailedTestSuiteExecutionsObjectsMap[detailedTestSuiteExecutionMapKey] = &tempDetailedTestSuiteExecutionsMapObject
	}

}

// ClearFlagRefreshOngoingOfDetailedTestSuiteExecution
// Clear the flag there is an ongoing refresh of the DetailedTestSuiteExecution-data
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) ClearFlagRefreshOngoingOfDetailedTestSuiteExecution(
	detailedTestSuiteExecutionMapKey DetailedTestSuiteExecutionMapKeyType) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "2a1073bb-82c3-441c-800b-415a73434fde",
	}).Debug("Incoming - 'ClearFlagRefreshOngoingOfDetailedTestSuiteExecution'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "9e81f4f4-9519-4102-9aa3-a756e5d2c6e6",
	}).Debug("Outgoing - 'ClearFlagRefreshOngoingOfDetailedTestSuiteExecution'")

	var existInMap bool

	// Lock Map for Writing
	detailedTestSuiteExecutionsMapMutex.Lock()

	//UnLock Map
	defer detailedTestSuiteExecutionsMapMutex.Unlock()

	// Check if Map i nil
	if TestSuiteExecutionsModel.DetailedTestSuiteExecutionsObjectsMapPtr == nil {

		var tempDetailedTestSuiteExecutionsObjectsMap map[DetailedTestSuiteExecutionMapKeyType]*DetailedTestSuiteExecutionsMapObjectStruct
		tempDetailedTestSuiteExecutionsObjectsMap = make(map[DetailedTestSuiteExecutionMapKeyType]*DetailedTestSuiteExecutionsMapObjectStruct)

		TestSuiteExecutionsModel.DetailedTestSuiteExecutionsObjectsMapPtr = &tempDetailedTestSuiteExecutionsObjectsMap

	}

	// Get the map from the map pointer
	var tempDetailedTestSuiteExecutionsObjectsMap map[DetailedTestSuiteExecutionMapKeyType]*DetailedTestSuiteExecutionsMapObjectStruct
	tempDetailedTestSuiteExecutionsObjectsMap = *TestSuiteExecutionsModel.DetailedTestSuiteExecutionsObjectsMapPtr

	// Check of DetailedTestSuiteExecution exist in Map
	var tempDetailedTestSuiteExecutionsMapObjectPtr *DetailedTestSuiteExecutionsMapObjectStruct
	var tempDetailedTestSuiteExecutionsMapObject DetailedTestSuiteExecutionsMapObjectStruct
	tempDetailedTestSuiteExecutionsMapObjectPtr, existInMap = tempDetailedTestSuiteExecutionsObjectsMap[detailedTestSuiteExecutionMapKey]

	// Set that refresh of DetailedTestSuiteExecution is ongoing
	if existInMap == true {
		// Already exist,so use that object by Get the TestSuiteExecutionObject from pointer
		// Get the Map
		tempDetailedTestSuiteExecutionsMapObject = *tempDetailedTestSuiteExecutionsMapObjectPtr

		// Already exist,so use that object by Get the TestSuiteExecutionObject from pointer
		tempDetailedTestSuiteExecutionsMapObject.WaitingForDatabaseUpdate = false
		fmt.Println("tempDetailedTestSuiteExecutionsMapObject.WaitingForDatabaseUpdateMutex.UnLock() - 01")
		tempDetailedTestSuiteExecutionsMapObject.WaitingForDatabaseUpdateMutex.Unlock()
	} else {
		// Object doesn't exist so create a new object and store in map
		tempDetailedTestSuiteExecutionsMapObject = DetailedTestSuiteExecutionsMapObjectStruct{
			DetailedTestSuiteExecution:    nil,
			WaitingForDatabaseUpdate:      true,
			WaitingForDatabaseUpdateMutex: &sync.RWMutex{},
		}

		fmt.Println("tempDetailedTestSuiteExecutionsMapObject.WaitingForDatabaseUpdateMutex.Lock() - 02")
		tempDetailedTestSuiteExecutionsMapObject.WaitingForDatabaseUpdateMutex.Unlock()

		tempDetailedTestSuiteExecutionsObjectsMap[detailedTestSuiteExecutionMapKey] = &tempDetailedTestSuiteExecutionsMapObject
	}

}
