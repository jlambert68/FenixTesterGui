package executionsModel

import (
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"time"
)

// Object, direct from database, holding TestCaseExecutions that is ongoing and belongs to all or some Domains
var allTestCaseExecutionFinishedExecutions allTestCaseExecutionsThatHaveBeenFinishedExecutedStruct

type allTestCaseExecutionsThatHaveBeenFinishedExecutedStruct struct {
	databaseReadTimeStamp                    time.Time
	testCaseExecutionsBelongsToTheseDomains  []string // When empty then there are no restrictions
	testCaseExecutionsWithFinishedExecutions []*fenixExecutionServerGuiGrpcApi.TestCaseWithFinishedExecutionMessage
}

// Object model for TestCaseExecutions that have finished their executions and belongs to all or some Domains
var allTestCaseExecutionsThatHaveBeenFinishedExecutedModel map[testCaseExecutionMapKeyType]*fenixExecutionServerGuiGrpcApi.TestCaseWithFinishedExecutionMessage
