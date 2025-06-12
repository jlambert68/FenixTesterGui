package testSuitesCommandEngine

import "fyne.io/fyne/v2/container"

// ***************** TestSuite Shared Variables *****************
// TestSuiteTabsRef
// The Tab-object holding all TestSuiteTabsRef
var TestSuiteTabs *container.DocTabs

// ***************** TestSuite CommandChannel *****************
var TestSuiteCommandChannel TestSuiteCommandChannelType

type TestSuiteCommandChannelType chan CommandTestSuiteChannelStruct

type TestSuiteChannelCommandType uint8

// The differerent commands supported when one need send commands upwards in the dependency chain
const (
	TestSuiteChannelCommandRefreshTestSuiteTabsObject TestSuiteChannelCommandType = iota
)

// CommandTestSuiteChannelStruct
// The command structure that is sent when one need send commands upwards in the dependency chain
type CommandTestSuiteChannelStruct struct {
	ChannelCommand TestSuiteChannelCommandType
}

// TestSuiteTabsRef
// The Tab-object holding all TestSuiteTabsRef
var TestSuiteTabsRef *container.DocTabs
