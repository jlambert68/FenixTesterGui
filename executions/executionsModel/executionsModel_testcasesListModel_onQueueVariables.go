package executionsModel

import (
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"time"
)

// Object, direct from database,  holding TestCaseExecutions that exists on the TestCaseExecutionQueue and belongs to all or some Domains
var allTestCaseExecutionsOnQueue allTestCaseExecutionsOnQueueStruct

type allTestCaseExecutionsOnQueueStruct struct {
	databaseReadTimeStamp                   time.Time
	testCaseExecutionsBelongsToTheseDomains []string // When empty then there are no restrictions
	testCaseExecutionsOnQueue               []*fenixExecutionServerGuiGrpcApi.TestCaseExecutionBasicInformationMessage
}

// Object model for TestCaseExecutions that exists on the TestCaseExecutionQueue and belongs to all or some Domains
var allTestCaseExecutionsOnQueueModel map[testCaseExecutionMapKeyType]*fenixExecutionServerGuiGrpcApi.TestCaseExecutionBasicInformationMessage
