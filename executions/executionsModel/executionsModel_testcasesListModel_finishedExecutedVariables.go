package executionsModel

import (
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"time"
)

// Object, direct from database, holding TestCaseExecutions that is ongoing and belongs to all or some Domains
var allTestCaseExecutionFinsiedExecutions allTestCaseExecutionsThatHaveBeenFinishedExecutedStruct

type allTestCaseExecutionsThatHaveBeenFinishedExecutedStruct struct {
	databaseReadTimeStamp                   time.Time
	testCaseExecutionsBelongsToTheseDomains []domainsStruct // When empty then there are no restrictions
	testCaseExecutionsOngoingExecutions     []*fenixExecutionServerGuiGrpcApi.TestCaseUnderExecutionMessage
}

// Object model for TestCaseExecutions that have finished their executions and belongs to all or some Domains
var allTestCaseExecutionsThatHaveBeenFinishedExecutedModel map[testCaseExecutionMapKey]*fenixExecutionServerGuiGrpcApi.TestCaseUnderExecutionMessage
