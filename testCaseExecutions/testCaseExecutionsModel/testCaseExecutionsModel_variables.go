package testCaseExecutionsModel

import (
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"time"
)

// StandardTestCaseExecutionsBatchSize
// The number if rows to be retrieved in one batch
const StandardTestCaseExecutionsBatchSize = 100

// NullTimeStampForTestCaseExecutionsSearch
// Null timestamp used in Search/Load of TestCaseExecutions
var NullTimeStampForTestCaseExecutionsSearch time.Time = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)

// TestCaseExecutionsModel
// Holding all data around Executions
var TestCaseExecutionsModel TestCaseExecutionsModelStruct

// TestCaseExecutionsModelStruct
// Type for holding all data around Executions
type TestCaseExecutionsModelStruct struct {
	TestCaseExecutionsThatCanBeViewedByUserMap map[string]*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage
	//TestCaseExecutionsThatCanBeViewedByUserSlice []*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage
	LatestUniqueTestCaseExecutionDatabaseRowId int32
	MoreRowsExists                             bool
	CurrentActiveTestCaseExecutionUuid         string
	StandardTestCaseExecutionsBatchSize        int32 // The number if rows to be retrieved in one batch
	NullTimeStampForTestCaseExecutionsSearch   time.Time
}
