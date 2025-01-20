package testCaseExecutionsModel

import (
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
)

// TestCaseExecutionsModel
// Holding all data around Executions
var TestCaseExecutionsModel TestCaseExecutionsModelStruct

// TestCaseExecutionsModelStruct
// Type for holding all data around Executions
type TestCaseExecutionsModelStruct struct {
	TestCaseExecutionsThatCanBeViewedByUserMap map[string]*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage
	//TestCaseExecutionsThatCanBeViewedByUserSlice []*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage
	CurrentActiveTestCaseExecutionUuid string
}
