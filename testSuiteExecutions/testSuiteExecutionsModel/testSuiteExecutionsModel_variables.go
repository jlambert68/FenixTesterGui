package testSuiteExecutionsModel

import (
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"sync"
	"time"
)

// StandardTestSuiteExecutionsBatchSize
// The number if rows to be retrieved in one batch
const StandardTestSuiteExecutionsBatchSize = 100

// NullTimeStampForTestSuiteExecutionsSearch
// Null timestamp used in Search/Load of TestSuiteExecutions
var NullTimeStampForTestSuiteExecutionsSearch time.Time = time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)

// TestSuiteExecutionsModel
// Holding all data around Executions
var TestSuiteExecutionsModel TestSuiteExecutionsModelStruct

// TestSuiteExecutionUuidType
// The type for key of the 'TestSuiteExecutionsThatCanBeViewedByUserMap'
type TestSuiteExecutionUuidType string

type latestTestSuiteExecutionForEachTestSuiteUuidStruct struct {
	latestTestSuiteExecutionForEachTestSuiteUuidMap map[TestSuiteExecutionUuidType]*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage
	LatestUniqueTestSuiteExecutionDatabaseRowId     int32
	MoreRowsExists                                  bool
}

// TestSuiteUuidType
// The type for key of the 'AllTestSuiteExecutionsForOneTestSuiteThatCanBeViewedByUserMap'
type TestSuiteUuidType string

type AllTestSuiteExecutionsForOneTestSuiteThatCanBeViewedByUserMapType map[TestSuiteExecutionUuidType]*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage

type allTestSuiteExecutionsForOneTestSuiteUuidStruct struct {
	allTestSuiteExecutionsForAllTestSuitesThatCanBeViewedByUserMap map[TestSuiteExecutionUuidType]*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage
	latestUniqueTestSuiteExecutionDatabaseRowId                    int32
	moreRowsExists                                                 bool
}

type DetailedTestSuiteExecutionMapKeyType string // TestSuiteExecutionUuid + TestSuiteExecutionVersion

type TestInstructionExecutionLogPostMapKeyType string // TestInstructionExecutionUuid + TestInstructionExecutionVersion

type TCEoTICoTIEAttributesContainerMapKeyType string // TestInstructionExecutionUuid || TestInstructionContainerUuid || TestSuiteExecutionUuid

type RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType string // TestInstructionUuid

type TestInstructionExecutionAttributeRunTimeUpdatedMapKeyType string // TestInstructionExecutionUuid + TestInstructionExecutionVersion

type AttributeNameMapKeyType string          // AttributeName
type RunTimeUpdatedAttributeValueType string // The attributes value after it was changed during run time

type TestInstructionExecutionDetailsMapKeyType string // TestInstructionExecutionUuid + TestInstructionExecutionVersion

// TestInstructionExecutionDetailsStruct
// Used to keep the base-data for TestInstructionExecution-explorer
type TestInstructionExecutionDetailsStruct struct {
	TestInstructionExecutionDetails          []*fenixExecutionServerGuiGrpcApi.TestInstructionExecutionsInformationMessage
	TestInstructionExecutionBasicInformation *fenixExecutionServerGuiGrpcApi.TestInstructionExecutionBasicInformationMessage
}

// TestInstructionExecutionDetailsForExplorerStruct
// Used in TestInstructionExecution-explorer to present data regarding the TestInstructionExecutions
type TestInstructionExecutionDetailsForExplorerStruct struct {
	TestInstructionExecutionDetails          *fenixExecutionServerGuiGrpcApi.TestInstructionExecutionsInformationMessage
	TestInstructionExecutionBasicInformation *fenixExecutionServerGuiGrpcApi.TestInstructionExecutionBasicInformationMessage
}

type RelationBetweenTestInstructionUuidAndTestInstructionExecutionStruct struct {
	TestInstructionUuid RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType
	TestInstructionName string
}

type TestInstructionExecutionUuidType string // TestInstructionExecutionUuid

type DetailedTestSuiteExecutionsMapObjectStruct struct {
	DetailedTestSuiteExecution                                              *fenixExecutionServerGuiGrpcApi.TestSuiteExecutionResponseMessage
	TestInstructionExecutionLogPostMapPtr                                   *map[TestInstructionExecutionLogPostMapKeyType]*[]*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage
	RelationBetweenTestInstructionUuidAndTestInstructionExecutionUuidMapPtr *map[RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType]TestInstructionExecutionUuidType
	RelationBetweenTestInstructionExecutionUuidAndTestInstructionUuidMapPtr *map[TestInstructionExecutionUuidType]RelationBetweenTestInstructionUuidAndTestInstructionExecutionStruct
	RunTimeUpdatedAttributesMapPtr                                          *map[TestInstructionExecutionAttributeRunTimeUpdatedMapKeyType]*map[AttributeNameMapKeyType]RunTimeUpdatedAttributeValueType
	TestInstructionExecutionDetailsMapPtr                                   *map[TestInstructionExecutionDetailsMapKeyType]*TestInstructionExecutionDetailsStruct

	WaitingForDatabaseUpdate      bool
	WaitingForDatabaseUpdateMutex *sync.RWMutex
}

// TestSuiteExecutionsModelStruct
// Type for holding all data around Executions
type TestSuiteExecutionsModelStruct struct {
	LatestTestSuiteExecutionForEachTestSuiteUuid                   latestTestSuiteExecutionForEachTestSuiteUuidStruct
	AllTestSuiteExecutionsForAllTestSuitesThatCanBeViewedByUserMap map[TestSuiteUuidType]*allTestSuiteExecutionsForOneTestSuiteUuidStruct
	DetailedTestSuiteExecutionsObjectsMapPtr                       *map[DetailedTestSuiteExecutionMapKeyType]*DetailedTestSuiteExecutionsMapObjectStruct

	StandardTestSuiteExecutionsBatchSize      int32 // The number if rows to be retrieved in one batch
	NullTimeStampForTestSuiteExecutionsSearch time.Time
}
