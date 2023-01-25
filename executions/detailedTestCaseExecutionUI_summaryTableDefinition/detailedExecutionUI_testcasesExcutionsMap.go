package detailedTestCaseExecutionUI_summaryTableDefinition

import fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"

// BLOCK START
// The block below is used for storing all detailed data belonging to a TestCaseExecution and structures needed for reflecting status updates to the UI

// TestCaseExecutionsDetailsStruct
// One TestCaseExecution and all of its data.
type TestCaseExecutionsDetailsStruct struct {
	// A full TestCaseExecutionStatus will always be performed when first status update message is received
	FullTestCaseExecutionUpdateWhenFirstExecutionStatusReceived bool

	// A full TestCaseExecutionStatus will always be performed when first TestInstruction-status update message is received
	FullTestCaseExecutionUpdateWhenFirstTestInstructionExecutionStatusReceived bool

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

// Holding the information to be show in the SummaryTable for all TestCaseExecutions
var TestCaseExecutionsStatusForSummaryTable []*TestCaseExecutionsStatusForSummaryTableStruct

// BLOCK END
