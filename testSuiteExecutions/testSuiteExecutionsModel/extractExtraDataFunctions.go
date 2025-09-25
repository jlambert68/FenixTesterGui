package testSuiteExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"strconv"
)

// ExtractAndStoreLogPostsAndValuesFromDetailedtestSuiteExecution
// Extracts all extra data that will be presented to the user in GUI, ie the explorer-tabs
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) ExtractAndStoreLogPostsAndValuesFromDetailedtestSuiteExecution(
	detailedTestSuiteExecutionMapKey DetailedTestSuiteExecutionMapKeyType) (
	err error) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "64c02558-b28f-4dba-a085-1c52b42d320f",
	}).Debug("Incoming - 'ExtractAndStoreLogPostsAndValuesFromDetailedtestSuiteExecution'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "3895b37e-e8f4-4e98-b9cc-9e111d4523d5",
	}).Debug("Outgoing - 'ExtractAndStoreLogPostsAndValuesFromDetailedtestSuiteExecution'")

	// Extract the raw detailedTestSuiteExecution-message
	var detailedTestSuiteExecution *fenixExecutionServerGuiGrpcApi.testSuiteExecutionResponseMessage
	var existInMap bool
	detailedTestSuiteExecution, existInMap = testSuiteExecutionsModel.
		ReadFromDetailedTestSuiteExecutionsMap(detailedTestSuiteExecutionMapKey)

	// Lock Map for Writing
	detailedTestSuiteExecutionsMapMutex.Lock()

	//UnLock Map
	defer detailedTestSuiteExecutionsMapMutex.Unlock()

	if existInMap == false {

		return err
	}

	// Extract map with DetailedTestSuiteExecutionsMapObjects
	var detailedTestSuiteExecutionsObjectsMapPtr *map[DetailedTestSuiteExecutionMapKeyType]*DetailedTestSuiteExecutionsMapObjectStruct
	var detailedTestSuiteExecutionsObjectsMap map[DetailedTestSuiteExecutionMapKeyType]*DetailedTestSuiteExecutionsMapObjectStruct

	detailedTestSuiteExecutionsObjectsMapPtr = testSuiteExecutionsModel.DetailedTestSuiteExecutionsObjectsMapPtr
	detailedTestSuiteExecutionsObjectsMap = *detailedTestSuiteExecutionsObjectsMapPtr

	// Extract specificDetailedTestSuiteExecutionsMapObject
	var detailedTestSuiteExecutionsMapObjectPtr *DetailedTestSuiteExecutionsMapObjectStruct
	var detailedTestSuiteExecutionsMapObject DetailedTestSuiteExecutionsMapObjectStruct
	detailedTestSuiteExecutionsMapObjectPtr, existInMap = detailedTestSuiteExecutionsObjectsMap[detailedTestSuiteExecutionMapKey]
	detailedTestSuiteExecutionsMapObject = *detailedTestSuiteExecutionsMapObjectPtr

	if existInMap == false {
		return err
	}

	// Always reInitialized RelationBetweenTestInstructionUuidAndTestInstructionExecutionUuidMapPtr
	var tempRelationBetweenTestInstructionUuidAndTestInstructionExectuionUuidMap map[RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType]TestInstructionExecutionUuidType
	//var tempRelationBetweenTestInstructionUuidAndTestInstructionExectuionUuidMapPtr *map[RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType]TestInstructionExecutionUuidType
	tempRelationBetweenTestInstructionUuidAndTestInstructionExectuionUuidMap = make(map[RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType]TestInstructionExecutionUuidType)

	detailedTestSuiteExecutionsMapObjectPtr.RelationBetweenTestInstructionUuidAndTestInstructionExecutionUuidMapPtr = &tempRelationBetweenTestInstructionUuidAndTestInstructionExectuionUuidMap

	// Always reInitialized TestInstructionExecutionLogPostMapPtr
	var tempTestInstructionExecutionLogPostMap map[TestInstructionExecutionLogPostMapKeyType]*[]*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage
	tempTestInstructionExecutionLogPostMap = make(map[TestInstructionExecutionLogPostMapKeyType]*[]*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage)
	detailedTestSuiteExecutionsMapObjectPtr.TestInstructionExecutionLogPostMapPtr = &tempTestInstructionExecutionLogPostMap

	// Always reInitialized RunTimeUpdatedAttributesMapPtr
	var tempRunTimeUpdatedAttributesMap map[TestInstructionExecutionAttributeRunTimeUpdatedMapKeyType]*map[AttributeNameMapKeyType]RunTimeUpdatedAttributeValueType
	tempRunTimeUpdatedAttributesMap = make(map[TestInstructionExecutionAttributeRunTimeUpdatedMapKeyType]*map[AttributeNameMapKeyType]RunTimeUpdatedAttributeValueType)
	detailedTestSuiteExecutionsMapObjectPtr.RunTimeUpdatedAttributesMapPtr = &tempRunTimeUpdatedAttributesMap

	detailedTestSuiteExecutionsMapObject = *detailedTestSuiteExecutionsMapObjectPtr

	// Always reInitialized RelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMapPtr
	var tempRelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMap map[TestInstructionExecutionUuidType]RelationBetweenTestInstructionUuidAndTestInstructionExecutionStruct
	tempRelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMap = make(map[TestInstructionExecutionUuidType]RelationBetweenTestInstructionUuidAndTestInstructionExecutionStruct)

	detailedTestSuiteExecutionsMapObjectPtr.RelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMapPtr = &tempRelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMap

	// Always reInitialized TestInstructionExecutionDetailsMapPtr
	var tempTestInstructionExecutionDetailsMap map[TestInstructionExecutionDetailsMapKeyType]*TestInstructionExecutionDetailsStruct
	tempTestInstructionExecutionDetailsMap = make(map[TestInstructionExecutionDetailsMapKeyType]*TestInstructionExecutionDetailsStruct)
	detailedTestSuiteExecutionsMapObjectPtr.TestInstructionExecutionDetailsMapPtr = &tempTestInstructionExecutionDetailsMap

	// Get the LogPostMap
	var testInstructionExecutionLogPostMapPtr *map[TestInstructionExecutionLogPostMapKeyType]*[]*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage
	var testInstructionExecutionLogPostMap map[TestInstructionExecutionLogPostMapKeyType]*[]*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage
	testInstructionExecutionLogPostMapPtr = detailedTestSuiteExecutionsMapObject.TestInstructionExecutionLogPostMapPtr
	testInstructionExecutionLogPostMap = *testInstructionExecutionLogPostMapPtr

	// Loop all TestInstructionExecutions and extract logPost-message, RuntTimeAttribute-data and TestInstructionExecution-data
	for _, testInstructionExecution := range detailedTestSuiteExecution.TestInstructionExecutions {

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

	// Add Maps back to 'detailedTestSuiteExecutionsMapObject'
	detailedTestSuiteExecutionsMapObjectPtr.RelationBetweenTestInstructionUuidAndTestInstructionExecutionUuidMapPtr = &tempRelationBetweenTestInstructionUuidAndTestInstructionExectuionUuidMap
	detailedTestSuiteExecutionsMapObjectPtr.RelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMapPtr = &tempRelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMap
	detailedTestSuiteExecutionsMapObjectPtr.RunTimeUpdatedAttributesMapPtr = &tempRunTimeUpdatedAttributesMap
	detailedTestSuiteExecutionsMapObjectPtr.TestInstructionExecutionDetailsMapPtr = &tempTestInstructionExecutionDetailsMap

	return err
}
