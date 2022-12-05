package executionsModelForSubscriptions

import (
	"FenixTesterGui/headertable"
	"fyne.io/fyne/v2"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"time"
)

// MaximumNumberOfItemsForUnderExecutionTableAddRemoveChannel - Maximum number of items that can be put on Channel before they
// needed to read out of channel
const MaximumNumberOfItemsForUnderExecutionTableAddRemoveChannel = 100

// UnderExecutionTableAddRemoveChannel - Used to secure that Add and Remove don't conflict
var UnderExecutionTableAddRemoveChannel UnderExecutionTableAddRemoveChannelType

// UnderExecutionTableAddRemoveChannelType - Type for 'UnderExecutionTableAddRemoveChannel'
type UnderExecutionTableAddRemoveChannelType chan UnderExecutionTableAddRemoveChannelStruct

// UnderExecutionTableChannelCommandType - Type for the channelCommand enumeration
type UnderExecutionTableChannelCommandType uint8

// Enumeration for the channel command
const (
	UnderExecutionTableAddRemoveChannelAddCommand_MoveFromOnQueueToUnderExecution UnderExecutionTableChannelCommandType = iota
	UnderExecutionTableAddRemoveChannelRemoveCommand_Flash
	UnderExecutionTableAddRemoveChannelRemoveCommand_Remove
)

// UnderExecutionTableAddRemoveChannelStruct - The channel message structure
type UnderExecutionTableAddRemoveChannelStruct struct {
	ChannelCommand    UnderExecutionTableChannelCommandType
	AddCommandData    UnderExecutionAddCommandDataStruct
	RemoveCommandData UnderExecutionRemoveCommandDataStruct
}

// UnderExecutionAddCommandDataStruct -The data used when a row should be added to the UnderExecution-table
type UnderExecutionAddCommandDataStruct struct {
	TestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference *TestCaseExecutionsOnQueueAdaptedForUiTableStruct
	TestCaseExecutionDetails                                   *fenixExecutionServerGuiGrpcApi.TestCaseExecutionDetailsMessage
}

// UnderExecutionRemoveCommandDataStruct -The data used when a row should be deleted from the UnderExecution-table
type UnderExecutionRemoveCommandDataStruct struct {
	TestCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference *TestCaseExecutionsUnderExecutionAdaptedForUiTableStruct
}

// Object, direct from database, holding TestCaseExecutions that is ongoing and belongs to all or some Domains
var allTestCaseExecutionsUnderExecution allTestCaseExecutionsOngoingUnderExecutionStruct

type allTestCaseExecutionsOngoingUnderExecutionStruct struct {
	databaseReadTimeStamp                   time.Time
	testCaseExecutionsBelongsToTheseDomains []string // When empty then there are no restrictions
	testCaseExecutionsUnderExecution        []*fenixExecutionServerGuiGrpcApi.TestCaseUnderExecutionMessage
}

// AllTestCaseExecutionsUnderExecutionModel
// Object model for TestCaseExecutions that is ongoing and belongs to all or some Domains
var AllTestCaseExecutionsUnderExecutionModel map[TestCaseExecutionMapKeyType]*fenixExecutionServerGuiGrpcApi.TestCaseUnderExecutionMessage

// TestCaseExecutionsUnderExecutionMapAdaptedForUiTable
// Object holding the testdata to be shown in UI regarding which TestCaseExecutions that is UnderExecution
var TestCaseExecutionsUnderExecutionMapAdaptedForUiTable map[TestCaseExecutionMapKeyType]*TestCaseExecutionsUnderExecutionAdaptedForUiTableStruct

// TestCaseExecutionsUnderExecutionAdaptedForUiTableStruct
// Type for holding one row of data
type TestCaseExecutionsUnderExecutionAdaptedForUiTableStruct struct {
	//TestCaseExecutionBasicInformation
	DomainUuid                          string // The Domain, UUID, where the TestCase 'has its home'
	DomainName                          string // The Domain, Name, where the TestCase 'has its home'
	TestSuiteUuid                       string // The TestSuite, UUID, that the TestCase was executed from
	TestSuiteName                       string // The TestSuite, Name, that the TestCase was executed from
	TestSuiteVersion                    string // The TestSuites version number
	TestSuiteExecutionUuid              string // The Unique UUID for the TestSuite Execution
	TestSuiteExecutionVersion           string // The TestSuites execution version
	TestCaseUuid                        string // The TestCase, UUID, set by TestCase-builder
	TestCaseName                        string // The TestCase, Name, set in TestCase-builder
	TestCaseVersion                     string // Each time a TestCase is saved then the 'TestCaseVersion' will be incremented by +1
	TestCaseExecutionUuid               string // The Unique UUID for the TestCase Execution
	TestCaseExecutionVersion            string // The Unique UUID for the TestCase Execution
	PlacedOnTestExecutionQueueTimeStamp string // The timestamp when the TestCase was placed on queue for execution
	ExecutionPriority                   string // The priority for the execution. Depends on who started it and if it belongs to a suite, scheduled or not

	// TestCaseExecutionDetails
	ExecutionStartTimeStamp        string // The timestamp when the execution was put for execution, not on queue for execution
	ExecutionStopTimeStamp         string // The timestamp when the execution was ended, in any way
	TestCaseExecutionStatus        string // The status of the ongoing  TestCase execution
	ExecutionHasFinished           string // A simple status telling if the execution has ended or not
	ExecutionStatusUpdateTimeStamp string // The timestamp when the status was last updated
}

// TestCaseExecutionsUnderExecutionTableOptions
// Defines the structure, and column order, for TestCaseExecutions-UnderExecution-Table
var TestCaseExecutionsUnderExecutionTableOptions = headertable.TableOpts{
	RefWidth:    "reference width",
	HeaderLabel: "TestExecutions under Execution",
	ColAttrs: []headertable.ColAttr{
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
}
