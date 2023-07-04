package sharedCode

import "github.com/sirupsen/logrus"

// CommandChannel Channel used for triggering command to CommandEngine

var CommandChannel CommandChannelType

type CommandChannelType chan ChannelCommandStruct

type ChannelCommandType uint8

const (
	ChannelCommandNewTestCase ChannelCommandType = iota
	ChannelCommandSwapElement
	ChannelCommandRemoveElement
	ChannelCommandSaveTestCase
	ChannelCommandExecuteTestCase
	ChannelCommandChangeActiveTestCase
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
	ChannelCommandGraphicsUpdate ChannelCommandGraphicsUpdatedType
	CreateNewTestCaseUI          bool
	ActiveTestCase               string
	TextualTestCaseSimple        string
	TextualTestCaseComplex       string
	TextualTestCaseExtended      string
	TestInstructionUuid          string
}

type ChannelCommandGraphicsUpdatedType uint8

const (
	ChannelCommandGraphicsUpdatedNewTestCase ChannelCommandGraphicsUpdatedType = iota
	ChannelCommandGraphicsUpdatedUpdateTestCaseGraphics
	ChannelCommandGraphicsUpdatedSelectTestInstruction
)

// BuildingBlock - Used for defining which type of element that user dragged from available building blocks tree
type BuildingBlock int

const (
	Undefined BuildingBlock = iota
	TestInstruction
	TestInstructionContainer
)

// Logger that can be used in every part of the code
var Logger *logrus.Logger

// Unique 'Uuid' for this running instance. Created at start up. Used as identification
var ApplicationRunTimeUuid string

// The current user of the TesterGui
var CurrentUserId string

const ZeroUuid = "00000000-0000-0000-0000-000000000000"
