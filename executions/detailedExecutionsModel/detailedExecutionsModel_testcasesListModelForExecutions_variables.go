package detailedExecutionsModel

import (
	"FenixTesterGui/executions/detailedTestCaseExecutionUI_summaryTableDefinition"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
)

// DetailedExecutionsModelObjectStruct - struct to object that hold all parts to of TestCaseExecution-model together
type DetailedExecutionsModelObjectStruct struct {
}

// detailedExecutionsModelObject -  Object that hold all parts to of TestCaseExecution-model together
var detailedExecutionsModelObject DetailedExecutionsModelObjectStruct

// BLOCK START
// The block below is used to define the TestExecutionSummaryTable

// DetailedTestCaseExecutionsAdaptedForUiSummaryTableStruct
// Type for holding one row of data
type DetailedTestCaseExecutionsAdaptedForUiSummaryTableStruct struct {
	TestCaseUIName      string
	TestCaseStatusValue uint32
}

// DetailedTestCaseExecutionsSummaryTableOptions
// Defines the structure, and column order, for TestCaseExecutions-FinishedExecution-Table
var DetailedTestCaseExecutionsSummaryTableOptions = detailedTestCaseExecutionUI_summaryTableDefinition.DetailedTestCaseExecutionsSummaryTableOpts{
	Bindings: nil,
	ColAttrs: []detailedTestCaseExecutionUI_summaryTableDefinition.DetailedTestCaseExecutionsSummaryColumnsAttributes{
		{
			Name:         "TestCaseUIName",
			Header:       "TestCaseUIName",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},
		/*	{
			Name:         "TestCaseStatusValue",
			Header:       "TestCaseStatusValue",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},*/
	},
	OnDataCellSelect:                      nil,
	RefWidth:                              "reference width",
	HeaderLabel:                           "Detailed TestCaseExecutions Summary Table",
	TestCaseExecutionsDetailsMapReference: make(map[widget.TableCellID]*detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionSummaryTableCellStruct),
	//FlashingTableCellsReferenceMap: nil,
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
	ChannelCommandRemoveDetailedTestCaseExecution
	ChannelCommandRetrieveFullDetailedTestCaseExecution
	ChannelCommandDecideIfStatusUpdatesBelongsToDetailedTestCaseExecutionsMap
)

// ChannelCommandDetailedExecutionsStruct
// Definition for a channel message
type ChannelCommandDetailedExecutionsStruct struct {
	// Command
	ChannelCommandDetailedExecutionsStatus ChannelCommandTypeForDetailedExecutionStatus

	// TestCaseExecutionKey = TestCaseExecutionUuid + TestCaseExecutionVersion
	TestCaseExecutionKey string

	// The full TestCaseExecution-message with all relevant information about the execution
	FullTestCaseExecutionResponseMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionResponseMessage

	// Status updates received via the subscription enginge
	TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage
}

// BLOCK END

// TestCasesSummaryTable
// The variable holding the UI-summary-object
var TestCasesSummaryTable *fyne.Container
