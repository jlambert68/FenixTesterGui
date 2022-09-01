package sharedCode

// CommandChannel Channel used for triggering command to CommandEngine
var CommandChannel chan ChannelCommandStruct

type ChannelCommandType uint8

const (
	ChannelCommandNewTestCase ChannelCommandType = iota
	ChannelCommandSwapElement
	ChannelCommandRemoveElement
)

type ChannelCommandStruct struct {
	ChannelCommand  ChannelCommandType
	FirstParameter  string
	SecondParameter string
	ActiveTestCase  string
}

// CommandChannelGraphicsUpdate - Channel for updating TestCase Graphics
var CommandChannelGraphicsUpdate chan ChannelCommandGraphicsUpdatedStruct

type ChannelCommandGraphicsUpdatedStruct struct {
	ActiveTestCase          string
	TextualTestCaseSimple   string
	TextualTestCaseComplex  string
	TextualTestCaseExtended string
}
