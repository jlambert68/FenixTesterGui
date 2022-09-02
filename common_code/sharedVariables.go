package sharedCode

// CommandChannel Channel used for triggering command to CommandEngine

var CommandChannel CommandChannelType

type CommandChannelType chan ChannelCommandStruct

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
	ElementType     BuildingBlock
}

// CommandChannelGraphicsUpdate - Channel for updating TestCase Graphics
var CommandChannelGraphicsUpdate CommandChannelGraphicsUpdateType

type CommandChannelGraphicsUpdateType chan ChannelCommandGraphicsUpdatedStruct

type ChannelCommandGraphicsUpdatedStruct struct {
	CreateNewTestCaseUI     bool
	ActiveTestCase          string
	TextualTestCaseSimple   string
	TextualTestCaseComplex  string
	TextualTestCaseExtended string
}

// BuildingBlock - Used for defining which type of element that user dragged from available building blocks tree
type BuildingBlock int

const (
	Undefined BuildingBlock = iota
	TestInstruction
	TestInstructionContainer
)
