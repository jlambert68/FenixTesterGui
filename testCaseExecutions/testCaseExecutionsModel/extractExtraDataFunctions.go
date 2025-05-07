package testCaseExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"strconv"
)

// ExtractAndStoreLogPostsAndValuesFromDetailedTestCaseExecution
// Extracts all extra data that will be presented to the user in GUI, ie the explorer-tabs
func (testCaseExecutionsModel TestCaseExecutionsModelStruct) ExtractAndStoreLogPostsAndValuesFromDetailedTestCaseExecution(
	detailedTestCaseExecutionMapKey DetailedTestCaseExecutionMapKeyType) (
	err error) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "d93737d7-3025-48c4-b683-bcc96a86024d",
	}).Debug("Incoming - 'ExtractAndStoreLogPostsAndValuesFromDetailedTestCaseExecution'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "000a3353-3fc2-44d8-9e22-8e7542fcad3f",
	}).Debug("Outgoing - 'ExtractAndStoreLogPostsAndValuesFromDetailedTestCaseExecution'")

	// Extract the raw detailedTestCaseExecution-message
	var detailedTestCaseExecution *fenixExecutionServerGuiGrpcApi.TestCaseExecutionResponseMessage
	var existInMap bool
	detailedTestCaseExecution, existInMap = testCaseExecutionsModel.
		ReadFromDetailedTestCaseExecutionsMap(detailedTestCaseExecutionMapKey)

	// Lock Map for Writing
	detailedTestCaseExecutionsMapMutex.Lock()

	//UnLock Map
	defer detailedTestCaseExecutionsMapMutex.Unlock()

	if existInMap == false {

		return err
	}

	// Extract map with DetailedTestCaseExecutionsMapObjects
	var detailedTestCaseExecutionsObjectsMapPtr *map[DetailedTestCaseExecutionMapKeyType]*DetailedTestCaseExecutionsMapObjectStruct
	var detailedTestCaseExecutionsObjectsMap map[DetailedTestCaseExecutionMapKeyType]*DetailedTestCaseExecutionsMapObjectStruct

	detailedTestCaseExecutionsObjectsMapPtr = testCaseExecutionsModel.DetailedTestCaseExecutionsObjectsMapPtr
	detailedTestCaseExecutionsObjectsMap = *detailedTestCaseExecutionsObjectsMapPtr

	// Extract specificDetailedTestCaseExecutionsMapObject
	var detailedTestCaseExecutionsMapObjectPtr *DetailedTestCaseExecutionsMapObjectStruct
	var detailedTestCaseExecutionsMapObject DetailedTestCaseExecutionsMapObjectStruct
	detailedTestCaseExecutionsMapObjectPtr, existInMap = detailedTestCaseExecutionsObjectsMap[detailedTestCaseExecutionMapKey]
	detailedTestCaseExecutionsMapObject = *detailedTestCaseExecutionsMapObjectPtr

	if existInMap == false {
		return err
	}

	// Always reInitialized RelationBetweenTestInstructionUuidAndTestInstructionExecutionUuidMapPtr
	var tempRelationBetweenTestInstructionUuidAndTestInstructionExectuionUuidMap map[RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType]TestInstructionExecutionUuidType
	//var tempRelationBetweenTestInstructionUuidAndTestInstructionExectuionUuidMapPtr *map[RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType]TestInstructionExecutionUuidType
	tempRelationBetweenTestInstructionUuidAndTestInstructionExectuionUuidMap = make(map[RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType]TestInstructionExecutionUuidType)

	detailedTestCaseExecutionsMapObjectPtr.RelationBetweenTestInstructionUuidAndTestInstructionExecutionUuidMapPtr = &tempRelationBetweenTestInstructionUuidAndTestInstructionExectuionUuidMap

	// Always reInitialized TestInstructionExecutionLogPostMapPtr
	var tempTestInstructionExecutionLogPostMap map[TestInstructionExecutionLogPostMapKeyType]*[]*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage
	tempTestInstructionExecutionLogPostMap = make(map[TestInstructionExecutionLogPostMapKeyType]*[]*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage)
	detailedTestCaseExecutionsMapObjectPtr.TestInstructionExecutionLogPostMapPtr = &tempTestInstructionExecutionLogPostMap

	// Always reInitialized RunTimeUpdatedAttributesMapPtr
	var tempRunTimeUpdatedAttributesMap map[TestInstructionExecutionAttributeRunTimeUpdatedMapKeyType]*map[AttributeNameMapKeyType]RunTimeUpdatedAttributeValueType
	tempRunTimeUpdatedAttributesMap = make(map[TestInstructionExecutionAttributeRunTimeUpdatedMapKeyType]*map[AttributeNameMapKeyType]RunTimeUpdatedAttributeValueType)
	detailedTestCaseExecutionsMapObjectPtr.RunTimeUpdatedAttributesMapPtr = &tempRunTimeUpdatedAttributesMap

	detailedTestCaseExecutionsMapObject = *detailedTestCaseExecutionsMapObjectPtr

	// Always reInitialized RelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMapPtr
	var tempRelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMap map[TestInstructionExecutionUuidType]RelationBetweenTestInstructionUuidAndTestInstructionExecutionStruct
	tempRelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMap = make(map[TestInstructionExecutionUuidType]RelationBetweenTestInstructionUuidAndTestInstructionExecutionStruct)

	detailedTestCaseExecutionsMapObjectPtr.RelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMapPtr = &tempRelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMap

	// Always reInitialized TestInstructionExecutionDetailsMapPtr
	var tempTestInstructionExecutionDetailsMap map[TestInstructionExecutionDetailsMapKeyType]*TestInstructionExecutionDetailsStruct
	tempTestInstructionExecutionDetailsMap = make(map[TestInstructionExecutionDetailsMapKeyType]*TestInstructionExecutionDetailsStruct)
	detailedTestCaseExecutionsMapObjectPtr.TestInstructionExecutionDetailsMapPtr = &tempTestInstructionExecutionDetailsMap

	// Get the LogPostMap
	var testInstructionExecutionLogPostMapPtr *map[TestInstructionExecutionLogPostMapKeyType]*[]*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage
	var testInstructionExecutionLogPostMap map[TestInstructionExecutionLogPostMapKeyType]*[]*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage
	testInstructionExecutionLogPostMapPtr = detailedTestCaseExecutionsMapObject.TestInstructionExecutionLogPostMapPtr
	testInstructionExecutionLogPostMap = *testInstructionExecutionLogPostMapPtr

	// Loop all TestInstructionExecutions and extract logPost-message, RuntTimeAttribute-data and TestInstructionExecution-data
	for _, testInstructionExecution := range detailedTestCaseExecution.TestInstructionExecutions {

		// Generate the TestInstructionExecutionLogPostMapKey
		var testInstructionExecutionLogPostMapKey TestInstructionExecutionLogPostMapKeyType
		testInstructionExecutionLogPostMapKey = TestInstructionExecutionLogPostMapKeyType(testInstructionExecution.GetTestInstructionExecutionBasicInformation().
			TestInstructionExecutionUuid +
			strconv.Itoa(int(testInstructionExecution.GetTestInstructionExecutionBasicInformation().GetTestInstructionExecutionVersion())))

		// Get the TestInstructionExecutionLogPOstAndValuesSLice
		var logPostAndValuesMessageSlicePtr *[]*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage
		logPostAndValuesMessageSlicePtr, existInMap = testInstructionExecutionLogPostMap[testInstructionExecutionLogPostMapKey]

		if existInMap == false {

			// No existing logs exist
			testInstructionExecutionLogPostMap[testInstructionExecutionLogPostMapKey] = &testInstructionExecution.ExecutionLogPostsAndValues

		} else {

			// Append to existing logs
			*logPostAndValuesMessageSlicePtr = append(*logPostAndValuesMessageSlicePtr, testInstructionExecution.ExecutionLogPostsAndValues...)
		}

		// Generate the relationBetweenTestInstructionUuidAndTestInstructionExecutionMapKey
		var relationBetweenTestInstructionUuidAndTestInstructionExecutionMapKey RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType
		relationBetweenTestInstructionUuidAndTestInstructionExecutionMapKey = RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType(
			testInstructionExecution.GetTestInstructionExecutionBasicInformation().TestInstructionUuid)
		var testInstructionExecutionUuid TestInstructionExecutionUuidType
		testInstructionExecutionUuid = TestInstructionExecutionUuidType(testInstructionExecution.GetTestInstructionExecutionBasicInformation().TestInstructionExecutionUuid +
			strconv.Itoa(int(testInstructionExecution.GetTestInstructionExecutionBasicInformation().GetTestInstructionExecutionVersion())))

		tempRelationBetweenTestInstructionUuidAndTestInstructionExectuionUuidMap[relationBetweenTestInstructionUuidAndTestInstructionExecutionMapKey] = testInstructionExecutionUuid

		tempRelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMap[testInstructionExecutionUuid] = RelationBetweenTestInstructionUuidAndTestInstructionExecutionStruct{
			TestInstructionUuid: relationBetweenTestInstructionUuidAndTestInstructionExecutionMapKey,
			TestInstructionName: testInstructionExecution.GetTestInstructionExecutionBasicInformation().TestInstructionName +
				" [" + string(relationBetweenTestInstructionUuidAndTestInstructionExecutionMapKey[:8]) + "]"}

		// Generate the relation between Attribute and runtime changed data for Attribute
		var testInstructionExecutionAttributeRunTimeUpdatedMapKey TestInstructionExecutionAttributeRunTimeUpdatedMapKeyType
		var attributeNameMapKey AttributeNameMapKeyType

		var tempRunTimeUpdatedAttributeMap map[AttributeNameMapKeyType]RunTimeUpdatedAttributeValueType
		var tempRunTimeUpdatedAttributeMapPtr *map[AttributeNameMapKeyType]RunTimeUpdatedAttributeValueType

		// Generate map key for TestInstructionExecution
		testInstructionExecutionAttributeRunTimeUpdatedMapKey = TestInstructionExecutionAttributeRunTimeUpdatedMapKeyType(testInstructionExecution.GetTestInstructionExecutionBasicInformation().TestInstructionExecutionUuid +
			strconv.Itoa(int(testInstructionExecution.GetTestInstructionExecutionBasicInformation().GetTestInstructionExecutionVersion())))

		// Check if TestInstructionUuid exist in Map
		tempRunTimeUpdatedAttributeMapPtr, existInMap = tempRunTimeUpdatedAttributesMap[testInstructionExecutionAttributeRunTimeUpdatedMapKey]
		if existInMap == false {

			// Doesn't exist so initiate
			tempRunTimeUpdatedAttributeMap = make(map[AttributeNameMapKeyType]RunTimeUpdatedAttributeValueType)

		} else {

			// Use the existing map
			tempRunTimeUpdatedAttributeMap = *tempRunTimeUpdatedAttributeMapPtr
		}

		// Loop all RunTimeChangedAttributes
		for _, runTimeUpdatedAttribute := range testInstructionExecution.RunTimeUpdatedAttributes {

			// Generate map key for Attribute
			attributeNameMapKey = AttributeNameMapKeyType(runTimeUpdatedAttribute.TestInstructionAttributeName)

			// Add Attributes RunTime-value to Attribute-map
			tempRunTimeUpdatedAttributeMap[attributeNameMapKey] = RunTimeUpdatedAttributeValueType(
				runTimeUpdatedAttribute.AttributeValueAsString)

		}

		// Add Attributes-map back to 'tempRunTimeUpdatedAttributesMap' if there were any attributes
		if len(testInstructionExecution.RunTimeUpdatedAttributes) > 0 {

			tempRunTimeUpdatedAttributesMap[testInstructionExecutionAttributeRunTimeUpdatedMapKey] = &tempRunTimeUpdatedAttributeMap

		}

		// Generate TestInstructionExecution-base information
		var tempTestInstructionExecutionDetailsStruct *TestInstructionExecutionDetailsStruct
		tempTestInstructionExecutionDetailsStruct = &TestInstructionExecutionDetailsStruct{
			TestInstructionExecutionDetails:          testInstructionExecution.GetTestInstructionExecutionsInformation(),
			TestInstructionExecutionBasicInformation: testInstructionExecution.GetTestInstructionExecutionBasicInformation(),
		}

		// Add TestInstructionExecution-map 'tempTestInstructionExecutionDetailsMap'
		tempTestInstructionExecutionDetailsMap[TestInstructionExecutionDetailsMapKeyType(testInstructionExecution.GetTestInstructionExecutionBasicInformation().TestInstructionExecutionUuid+
			strconv.Itoa(int(testInstructionExecution.GetTestInstructionExecutionBasicInformation().GetTestInstructionExecutionVersion())))] = tempTestInstructionExecutionDetailsStruct

	}

	// Add Maps back to 'detailedTestCaseExecutionsMapObject'
	detailedTestCaseExecutionsMapObjectPtr.RelationBetweenTestInstructionUuidAndTestInstructionExecutionUuidMapPtr = &tempRelationBetweenTestInstructionUuidAndTestInstructionExectuionUuidMap
	detailedTestCaseExecutionsMapObjectPtr.RelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMapPtr = &tempRelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMap
	detailedTestCaseExecutionsMapObjectPtr.RunTimeUpdatedAttributesMapPtr = &tempRunTimeUpdatedAttributesMap
	detailedTestCaseExecutionsMapObjectPtr.TestInstructionExecutionDetailsMapPtr = &tempTestInstructionExecutionDetailsMap

	return err
}
