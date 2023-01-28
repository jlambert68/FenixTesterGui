package detailedTestCaseExecutionUI_summaryTableDefinition

import (
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"time"
)

// BLOCK START
// The block below is used for storing all detailed data belonging to a TestCaseExecution and structures needed for reflecting status updates to the UI

// FullTestCaseExecutionUpdateWhenFirstExecutionStatusReceivedMaxSize
// Tha max size for the channel for StatusUpdate-messages
const FullExecutionUpdateWhenFirstExecutionStatusReceivedMaxSize int32 = 100

// TestCaseExecutionsDetailsStruct
// One TestCaseExecution and all of its data.
type TestCaseExecutionsDetailsStruct struct {

	// Waiting for of full TestCaseExecutionStatus update is retrieved
	WaitingForFullTestCaseExecutionUpdate bool

	// Waiting for of full TestCaseExecutionStatus update is retrieved, after first TestInstructionStatus was received
	WaitingForFullTestCaseExecutionUpdateAfterFirstTestInstructionExecutionStatusWasReceived bool

	// TestCaseExecution StatusMessages that are Waiting For a Full TestCaseExecutionUpdate are temporary stores in this channel
	TestCaseExecutionStatusMessagesWaitingForFullTestCaseExecutionUpdate chan *fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusMessage

	// TestInstructionExecution StatusMessages that are Waiting For a Full TestCaseExecutionUpdate are temporary stores in this channel
	TestInstructionExecutionStatusMessagesWaitingForFullTestCaseExecutionUpdate chan *fenixExecutionServerGuiGrpcApi.TestInstructionExecutionStatusMessage

	// A full TestCaseExecutionStatus will always be performed when first status update message is received
	FirstExecutionStatusReceived bool

	// Keeps track of previous Timestamp for incoming status message to secure that no messages are lost
	PreviousBroadcastTimeStamp time.Time

	// A full TestCaseExecutionStatus will always be performed when first TestInstruction-status update message is received
	FirstTestInstructionExecutionStatusReceived bool

	// The response message when a full TestCaseExecution is retrieved
	TestCaseExecutionDatabaseResponseMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionResponseMessage

	// The streamed status messages
	TestCaseExecutionsStatusUpdates        []*fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusMessage
	TestInstructionExecutionsStatusUpdates []*fenixExecutionServerGuiGrpcApi.TestInstructionExecutionStatusMessage

	// Holding all relevant executions information for the TestCaseExecution itself
	TestCaseExecutionsBaseInformation *TestCaseExecutionsBaseInformationStruct

	// A map holding all TestInstructions with their execution statuses.
	// map[TestInstructionExecutionKey]*TestInstructionExecutionsBaseInformationStruct
	// TestInstructionExecutionKey = TestInstructionExecutionUuid + TestInstructionExecutionVersion
	TestInstructionExecutionsStatusMap map[string]*TestInstructionExecutionsBaseInformationStruct

	// The slice of all TestInstructionExecution, for one TestCaseExecution, and their current status. The order is the same as it is presented on screen
	TestInstructionExecutionsStatusForSummaryTable []*TestInstructionExecutionsStatusForSummaryTableStruct
}

// TestCaseExecutionsDetailsMap
// map[TestCaseExecutionMapKey]*TestCaseExecutionsDetailsStruct, TestCaseExecutionMapKey = TestCaseExecutionUuid + TestCaseExecutionVersionNumber
var TestCaseExecutionsDetailsMap map[string]*TestCaseExecutionsDetailsStruct // m

// TestCaseExecutionsStatusForSummaryTable
// Holding the information to be show in the SummaryTable for all TestCaseExecutions
var TestCaseExecutionsStatusForSummaryTable []*TestCaseExecutionsStatusForSummaryTableStruct

// BLOCK END
