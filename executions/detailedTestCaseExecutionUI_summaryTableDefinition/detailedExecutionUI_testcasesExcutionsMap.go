package detailedTestCaseExecutionUI_summaryTableDefinition

import fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"

// BLOCK START
// The block below is used for storing all detailed data belonging to a TestCaseExecution and structures needed for reflecting status updates to the UI

// TestCaseExecutionsDetailsStruct
// One TestCaseExecution and all of its data.
type TestCaseExecutionsDetailsStruct struct {
	// The response message when a full TestCaseExecution is retrieved
	TestCaseExecutionDatabaseResponseMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionResponseMessage

	// The streamed status messages
	TestCaseExecutionsStatusUpdates        []*fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusMessage
	TestInstructionExecutionsStatusUpdates []*fenixExecutionServerGuiGrpcApi.TestInstructionExecutionStatusMessage

	// A map holding all TestInstructions with their execution status. Each slice is sorted by 'UniqueDatabaseRowCounter' ASC order
	// The slice data is used to show execution status and the last item in the slice is the one that has the current status
	// map[TestInstructionExecutionKey]*[]*fenixExecutionServerGuiGrpcApi.TestInstructionExecutionsInformationMessage
	// TestInstructionExecutionKey = TestInstructionExecutionUuid + TestInstructionExecutionVersion
	TestInstructionExecutionsStatusMap map[string]*[]*fenixExecutionServerGuiGrpcApi.TestInstructionExecutionsInformationMessage

	// Holding the information to be show in the SummaryTable for one TestCaseExecution
	TestCaseExecutionsStatusForSummaryTable *TestCaseExecutionsStatusForSummaryTableStruct

	// The slice of all TestInstructionExecution, for one TestCaseExecution, and their current status. The order is the same as it is presented on screen
	TestInstructionExecutionsStatusForSummaryTable []*TestInstructionExecutionsStatusForSummaryTableStruct
}

// TestCaseExecutionsDetailsMap
// map[TestCaseExecutionMapKey]*TestCaseExecutionsDetailsStruct, TestCaseExecutionMapKey = TestCaseExecutionUuid + TestCaseExecutionVersionNumber
var TestCaseExecutionsDetailsMap map[string]*TestCaseExecutionsDetailsStruct // m

// BLOCK END
