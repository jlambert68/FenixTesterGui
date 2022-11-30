package executionsModel

import (
	"FenixTesterGui/headertable"
	"fyne.io/fyne/v2"
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

// AllTestCaseExecutionsOnQueueModel
// Object model for TestCaseExecutions that exists on the TestCaseExecutionQueue and belongs to all or some Domains
var AllTestCaseExecutionsOnQueueModel map[TestCaseExecutionMapKeyType]*fenixExecutionServerGuiGrpcApi.TestCaseExecutionBasicInformationMessage

// TestCaseExecutionsOnQueueMapAdaptedForUiTable
// Object holding the testdata to be shown in UI regarding which TestCaseExecutions that is waiting on ExecutionQueue
var TestCaseExecutionsOnQueueMapAdaptedForUiTable map[TestCaseExecutionMapKeyType]*TestCaseExecutionsOnQueueAdaptedForUiTableStruct

// TestCaseExecutionsOnQueueAdaptedForUiTableStruct
// Type for holding one row of data
type TestCaseExecutionsOnQueueAdaptedForUiTableStruct struct {
	// TestCaseExecutionBasicInformation
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
}

// TestCaseExecutionsOnQueueTableOptions
// Defines the structure, and column order, for TestCaseExecutions-OnQueue-Table
var TestCaseExecutionsOnQueueTableOptions = headertable.TableOpts{
	RefWidth:    "reference width",
	HeaderLabel: "TestExecutions waiting on Queue to be Executed",
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
