package executionsModel

import (
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"time"
)

// Object, direct from database, holding TestCaseExecutions that is ongoing and belongs to all or some Domains
var allTestCaseExecutionsUnderExecution allTestCaseExecutionsOngoingUnderExecutionStruct

type allTestCaseExecutionsOngoingUnderExecutionStruct struct {
	databaseReadTimeStamp                   time.Time
	testCaseExecutionsBelongsToTheseDomains []string // When empty then there are no restrictions
	testCaseExecutionsUnderExecution        []*fenixExecutionServerGuiGrpcApi.TestCaseUnderExecutionMessage
}

// Object model for TestCaseExecutions that is ongoing and belongs to all or some Domains
var allTestCaseExecutionsUnderExecutionModel map[testCaseExecutionMapKeyType]*fenixExecutionServerGuiGrpcApi.TestCaseUnderExecutionMessage
