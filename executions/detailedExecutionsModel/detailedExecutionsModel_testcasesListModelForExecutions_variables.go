package detailedExecutionsModel

import (
	"FenixTesterGui/executions/detailedTestCaseExecutionUI_summaryTableDefinition"
	"fyne.io/fyne/v2"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"time"
)

// DetailedExecutionsModelObjectStruct - struct to object that hold all parts to of TestCaseExecution-model together
type DetailedExecutionsModelObjectStruct struct {
}

// detailedExecutionsModelObject -  Object that hold all parts to of TestCaseExecution-model together
var detailedExecutionsModelObject DetailedExecutionsModelObjectStruct

// BLOCK START
// The block below is used to define the TestExecutionSummaryTable

// TestCaseExecutionsFinishedExecutionAdaptedForUiTableStruct
// Type for holding one row of data
type TestCaseExecutionsFinishedExecutionAdaptedForUiTableStruct___ struct {
	TestCaseUIName      string
	TestCaseStatusValue uint32
}

// TestCaseExecutionsFinishedExecutionTableOptions
// Defines the structure, and column order, for TestCaseExecutions-FinishedExecution-Table
var TestCaseExecutionsFinishedExecutionTableOptions = detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsSummaryTableOpts{
	Bindings: nil,
	ColAttrs: []detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsColAttr{
		{
			Name:         "ShowDetailedTestCaseExecution",
			Header:       "ShowDetailedTestCaseExecution",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},
		{
			Name:         "DomainUuid",
			Header:       "DomainUuid",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},
		{
			Name:         "DomainName",
			Header:       "DomainName",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},
		{
			Name:         "TestCaseUuid",
			Header:       "TestCaseUuid",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},
		{
			Name:         "TestCaseName",
			Header:       "TestCaseName",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},
		{
			Name:         "TestCaseVersion",
			Header:       "TestCaseVersion",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},
		{
			Name:         "TestCaseExecutionUuid",
			Header:       "TestCaseExecutionUuid",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},
		{
			Name:         "TestCaseExecutionVersion",
			Header:       "TestCaseExecutionVersion",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},
		{
			Name:         "PlacedOnTestExecutionQueueTimeStamp",
			Header:       "PlacedOnTestExecutionQueueTimeStamp",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},
		{
			Name:         "ExecutionPriority",
			Header:       "ExecutionPriority",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},
		{
			Name:         "ExecutionStartTimeStamp",
			Header:       "ExecutionStartTimeStamp",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},
		{
			Name:         "ExecutionStopTimeStamp",
			Header:       "ExecutionStopTimeStamp",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},

		{
			Name:         "TestCaseExecutionStatus",
			Header:       "TestCaseExecutionStatus",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},

		{
			Name:         "ExecutionHasFinished",
			Header:       "ExecutionHasFinished",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},
		{
			Name:         "ExecutionStatusUpdateTimeStamp",
			Header:       "ExecutionStatusUpdateTimeStamp",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},
		{
			Name:         "TestSuiteUuid",
			Header:       "TestSuiteUuid",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},
		{
			Name:         "TestSuiteName",
			Header:       "TestSuiteName",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},
		{
			Name:         "TestSuiteVersion",
			Header:       "TestSuiteVersion",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},
		{
			Name:         "TestSuiteExecutionUuid",
			Header:       "TestSuiteExecutionUuid",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},
		{
			Name:         "TestSuiteExecutionVersion",
			Header:       "TestSuiteExecutionVersion",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},
	},
	OnDataCellSelect:               nil,
	RefWidth:                       "reference width",
	HeaderLabel:                    "TestExecutions that are finished",
	FlashingTableCellsReferenceMap: nil,
}

// BLOCK END

// TestCaseExecutionMapKeyType
// Type used to define that this is TestCaseExecutionKey for model-maps
type TestCaseExecutionMapKeyType string // Should consist of 'TestCaseExecutionUuid' + 'TestCaseExecutionVersion'

// BLOCK START
// The block below  is used when checking if  a TestCaseExecution exists in any of the tables (OnQueue, UnderExecution, FinishedExecutions)

// SubscriptionTableType
// The type for subscription tables for TestCaseExecutions
type SubscriptionTableType uint8

// Subscription tables for TestCaseExecutions
const (
	SubscriptionTableForTestCaseExecutionOnQueueTable SubscriptionTableType = iota
	SubscriptionTableForTestCaseExecutionUnderExecutionTable
	SubscriptionTableForTestCaseExecutionFinishedExecutionsTable
)

// ExecutionsForTestCaseExecutionMapOverallType
// Map holding all information about all 'TestCaseExecutionMapKey'
type ExecutionsForTestCaseExecutionMapOverallType map[TestCaseExecutionMapKeyType]ExecutionsForTestCaseExecutionMapDetailsType

// ExecutionsForTestCaseExecutionMapDetailsType
// Map holding all information about one 'TestCaseExecutionMapKey' if the TestCaseExecution should exist in specific table (OnQueue, UnderExecution, FinishedExecutions)
type ExecutionsForTestCaseExecutionMapDetailsType map[SubscriptionTableType]ExecutionsForTestCaseExecutionMapDetailsStruct

// ExecutionsForTestCaseExecutionMapDetailsStruct
// Hold information about if a TestCaseExecution should exist in specific table (OnQueue, UnderExecution, FinishedExecutions)
type ExecutionsForTestCaseExecutionMapDetailsStruct struct {
	ShouldExistInTable bool
}

// BLOCK END

// BLOCK START
// The block below is used for securing that status updates of TestCaseExecutions and TestInstructionExecutions are handled in a controlled way

// DetailedExecutionStatusCommandChannel
// Parameters used for channel to update status on Executions for TestCases and TestInstructions
var DetailedExecutionStatusCommandChannel DetailedExecutionStatusChannelType

// DetailedExecutionStatusChannelType
// The definition for the channel
type DetailedExecutionStatusChannelType chan ChannelCommandDetailedExecutionsStruct

// Number of messages that the channel can queue
const messageChannelMaxSizeDetailedExecutionStatus int32 = 100

// ChannelCommandTypeForDetailedExecutionStatus
// The type used for ChannelCommands
type ChannelCommandTypeForDetailedExecutionStatus uint8

// The channel commands
const (
	ChannelCommandFullDetailedExecutionsStatusUpdate ChannelCommandTypeForDetailedExecutionStatus = iota
	ChannelCommandStatusUpdateOfDetailedExecutionsStatus
)

// ChannelCommandDetailedExecutionsStruct
// Definition for a channel message
type ChannelCommandDetailedExecutionsStruct struct {
	// Command
	ChannelCommandDetailedExecutionsStatus ChannelCommandTypeForDetailedExecutionStatus

	// The full TestCaseExecution-message with all relevant information about the execution
	FullTestCaseExecutionResponseMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionResponseMessage

	// Status updates received via the subscription enginge
	TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage
}

// BLOCK END

// BLOCK START
// The block below is used for storing all detailed data belonging to a TestCaseExecution and structures needed for reflecting status updates to the UI

// TestCaseExecutionsStatusForSummaryTableStruct
// The definition used in SummaryTable to represent one TestCaseExecution and its current execution status
type TestCaseExecutionsStatusForSummaryTableStruct struct {
	TestCaseUIName                 string
	TestCaseStatusValue            uint32
	ExecutionStatusUpdateTimeStamp time.Time
}

// TestInstructionExecutionsStatusForSummaryTableStruct
// The definition used in SummaryTable to represent one TestInstructionExecution and its current execution status
type TestInstructionExecutionsStatusForSummaryTableStruct struct {
	TestInstructionExecutionUIName string
	TestInstructionStatusValue     uint32
	ExecutionStatusUpdateTimeStamp time.Time
}

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
	TestCaseExecutionsStatusForSummaryTable TestCaseExecutionsStatusForSummaryTableStruct

	// The slice of all TestInstructionExecution, for one TestCaseExecution, and their current status. The order is the same as it is presented on screen
	TestInstructionExecutionsStatusForSummaryTable []TestInstructionExecutionsStatusForSummaryTableStruct
}

// TestCaseExecutionsDetailsMap
// map[TestCaseExecutionMapKey]*TestCaseExecutionsDetailsStruct, TestCaseExecutionMapKey = TestCaseExecutionUuid + TestCaseExecutionVersionNumber
var TestCaseExecutionsDetailsMap map[string]*TestCaseExecutionsDetailsStruct // m

// BLOCK END
