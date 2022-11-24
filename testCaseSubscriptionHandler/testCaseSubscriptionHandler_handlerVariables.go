package testCaseSubscriptionHandler

import (
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"time"
)

// TestCaseExecutionStatusSubscriptionHandlerStruct - Holds every together around Subscriptions
type TestCaseExecutionStatusSubscriptionHandlerStruct struct {
}

// TestCaseExecutionStatusSubscriptionHandlerObject - Main object for Subscriptions
var TestCaseExecutionStatusSubscriptionHandlerObject TestCaseExecutionStatusSubscriptionHandlerStruct

// TestCaseExecutionExecutionStatusSubscriptionMap - Keeps track of all TestCaseExecution Subscriptions that the TesterGui has.
// Map['TestCaseExecutionUuid' + 'TestCaseExecutionVersion']*TestCaseExecutionStatusSubscriptionStruct
var TestCaseExecutionExecutionStatusSubscriptionMap map[TestCaseExecutionStatusSubscriptionMapKeyType]*TestCaseExecutionStatusSubscriptionStruct

// TestCaseExecutionStatusSubscriptionMapKeyType - Defines the key to the Subscription-Map map['TestCaseExecutionUuid' + 'TestCaseExecutionVersion']
type TestCaseExecutionStatusSubscriptionMapKeyType string

// TestCaseExecutionStatusSubscriptionStruct - Hold all information about a Subscription
type TestCaseExecutionStatusSubscriptionStruct struct {
	TestCaseExecutionUuid            string
	TestCaseExecutionVersion         string
	AddedForSubscriptionTimeStamp    time.Time
	ReceivedExecutionsStatusMessages *[]*ReceivedExecutionsStatusMessagesStruct
}

// ReceivedExecutionsStatusMessagesStruct - Struct to hold one received ExecutionStatus.message and its TimeStamps
type ReceivedExecutionsStatusMessagesStruct struct {
	OriginalMessageCreationTimeStamp          time.Time
	ReceivedExecutionsStatusMessagesTimeStamp time.Time
	ReceivedExecutionsStatusMessages          *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage
}
