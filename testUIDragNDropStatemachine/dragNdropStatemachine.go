package testUIDragNDropStatemachine

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

// State for handling Drag from Source object
const (
	sourceStateSearching              = iota // 0
	sourceStateFinds                         // 1
	sourceStateGrabs                         // 2
	sourceStateDragging                      // 3
	sourceStateReleasingWithOutTarget        // 4
	sourceStateEnteringTarget                // 5
	sourceStateReleasingOnTarget             // 6
	sourceStateReleasedOnTarget              // 7
)

// State for handling Drop-targetDroppableLabel object
const (
	targetStateWaitingForSourceToEnteringTarget = iota // 0
	targetStateSourceIsDraggingObject                  // 1
	targetStateSourceEnteredTargetWithObject           // 2
	targetStateSourceReleasingOnTarget                 // 3
	targetStateSourceReleasedOnTarget                  // 4
)

const (
	targetDropLabelRectangleWidth  = 500
	targetDropLabelRectangleHeight = 12
	targetDropRectangleWidth       = 500
	targetDropRectangleHeight      = 12
)

type targetDroppedTypeType uint8

const (
	unspecifiedType targetDroppedTypeType = iota
	droppableLabelType
	droppableRectangleType
)

// Local variables for the Drag n Drop object
var textRef *canvas.Text
var rectangleRef *canvas.Rectangle
var rectangle2Ref *canvas.Rectangle
var containerRef *fyne.Container
var labelStandardHeight float32
var commandChannelReference *sharedCode.CommandChannelType

// ****************************************************

type noneDroppableLabel struct {
	widget.Label
	uuid string
}

// InitiateStateStateMachine
// InitiateState State machine
func (stateMachine *StateMachineDragAndDropStruct) InitiateStateStateMachine(
	dragNDropText *canvas.Text,
	dragNDropRectangleRef *canvas.Rectangle,
	dragNDropRectangle2Ref *canvas.Rectangle,
	dragNDropContainerRef *fyne.Container,
	commandChannelRef *sharedCode.CommandChannelType,
	testCasesRef *testCaseModel.TestCasesModelsStruct) {

	textRef = dragNDropText
	rectangleRef = dragNDropRectangleRef
	rectangle2Ref = dragNDropRectangle2Ref
	containerRef = dragNDropContainerRef
	commandChannelReference = commandChannelRef

	// Calculate label standard height
	tempLabel := widget.NewLabel("temp")
	tempLabel.Refresh()
	labelStandardHeight = tempLabel.MinSize().Height

	// Make map which keeps track of TestInstructionContainers Bond-belongings map
	stateMachineDragAndDrop.testInstructionContainerBondBelongingRectangleMap = make(map[string]*canvas.Rectangle)

	// Maps that keeps track of all registered TargetLabels and TargetRectangles
	stateMachineDragAndDrop.registeredDroppableTargetLabelsMap = make(map[string]*[]*DroppableLabel)
	stateMachineDragAndDrop.registeredDroppableTargetRectangleMap = make(map[string]*[]*DroppableRectangle)

	// Store reference to TestCases
	stateMachine.testCasesRef = testCasesRef
	stateMachineDragAndDrop.testCasesRef = testCasesRef

}

//****************************************************

func newNoneDroppableLabel(uuid string) *noneDroppableLabel {
	nonDroppableLabel := &noneDroppableLabel{}
	nonDroppableLabel.ExtendBaseWidget(nonDroppableLabel)

	nonDroppableLabel.uuid = uuid
	nonDroppableLabel.Text = uuid

	return nonDroppableLabel
}

//****************************************************

type StateMachineDragAndDropStruct struct {
	testCasesRef                                      *testCaseModel.TestCasesModelsStruct
	sourceStateMachine                                stateMachineSourceAndDestinationStruct
	targetStateMachine                                stateMachineSourceAndDestinationStruct
	registeredDroppableTargetLabelsMap                map[string]*[]*DroppableLabel
	registeredDroppableTargetRectangleMap             map[string]*[]*DroppableRectangle
	SourceUuid                                        string
	SourceType                                        int
	targetDroppedType                                 targetDroppedTypeType
	targetDroppableLabel                              DroppableLabel
	targetDroppableRectangle                          DroppableRectangle
	testInstructionContainerBondBelongingRectangleMap map[string]*canvas.Rectangle //map[TestInstructionContainerUuid]*canvas.Rectangle

}

// Structure for 'Drag-part of 'Drag-N-Drop' state machine
type stateMachineSourceAndDestinationStruct struct {
	currentState int
}

var stateMachineDragAndDrop StateMachineDragAndDropStruct

// Add testInstructionContainerBondBelongingRectangle to map
func TestInstructionContainerBondBelongingRectangleToMap(testInstructionContainerUuid string, rectangle *canvas.Rectangle) {
	stateMachineDragAndDrop.testInstructionContainerBondBelongingRectangleMap[testInstructionContainerUuid] = rectangle
}

func switchStateForSource(newState int) {
	stateMachineDragAndDrop.sourceStateMachine.currentState = newState
}

func switchStateForTarget(newState int) {
	stateMachineDragAndDrop.targetStateMachine.currentState = newState
}

func expandDropAreas() {

	// Get registeredDroppableTargetLabels for currentTestCase
	var droppableLabelsRef *[]*DroppableLabel
	var droppableLabels []*DroppableLabel
	droppableLabelsRef = stateMachineDragAndDrop.registeredDroppableTargetLabelsMap[stateMachineDragAndDrop.testCasesRef.CurrentActiveTestCaseUuid]

	if droppableLabelsRef != nil {
		droppableLabels = *droppableLabelsRef

		for _, targetLabel := range droppableLabels {

			targetLabel.BackgroundRectangle.StrokeWidth = 2

			targetLabel.Resize(fyne.NewSize(targetDropRectangleWidth, labelStandardHeight))
			strokeColor := color.RGBA{
				R: 0xFF,
				G: 0x00,
				B: 0x00,
				A: 0xAA,
			}
			targetLabel.BackgroundRectangle.StrokeColor = strokeColor
			//targetLabel.BackgroundRectangle.Show()
			//targetLabel.Show()
		}
	}

	// Get registeredDroppableTargetLabels for currentTestCase
	var droppableRectangleRef *[]*DroppableRectangle
	var droppableRectangle []*DroppableRectangle
	droppableRectangleRef = stateMachineDragAndDrop.registeredDroppableTargetRectangleMap[stateMachineDragAndDrop.testCasesRef.CurrentActiveTestCaseUuid]

	if droppableRectangleRef != nil {
		droppableRectangle = *droppableRectangleRef

		for _, targetRectangle := range droppableRectangle {

			//rectangleWidth := float32(500)

			targetRectangle.Rectangle.StrokeWidth = 2

			strokeColor := color.RGBA{
				R: 0xFF,
				G: 0x00,
				B: 0x00,
				A: 0xAA,
			}
			targetRectangle.Rectangle.StrokeColor = strokeColor

		}
	}
}

func shrinkDropAreas() {

	// Get registeredDroppableTargetLabels for currentTestCase
	var droppableLabelsRef *[]*DroppableLabel
	var droppableLabels []*DroppableLabel
	droppableLabelsRef = stateMachineDragAndDrop.registeredDroppableTargetLabelsMap[stateMachineDragAndDrop.testCasesRef.CurrentActiveTestCaseUuid]

	if droppableLabelsRef != nil {
		droppableLabels = *droppableLabelsRef

		for _, targetLabel := range droppableLabels {

			targetLabel.BackgroundRectangle.StrokeWidth = 2
			strokeColor := color.RGBA{
				R: 0x00,
				G: 0x00,
				B: 0x00,
				A: 0x00,
			}
			targetLabel.BackgroundRectangle.StrokeColor = strokeColor
		}
	}

	// Get registeredDroppableTargetLabels for currentTestCase
	var droppableRectangleRef *[]*DroppableRectangle
	var droppableRectangle []*DroppableRectangle
	droppableRectangleRef = stateMachineDragAndDrop.registeredDroppableTargetRectangleMap[stateMachineDragAndDrop.testCasesRef.CurrentActiveTestCaseUuid]

	if droppableRectangleRef != nil {
		droppableRectangle = *droppableRectangleRef

		for _, targetRectangle := range droppableRectangle {

			targetRectangle.Rectangle.StrokeWidth = 2
			strokeColor := color.RGBA{
				R: 0x00,
				G: 0x00,
				B: 0x00,
				A: 0x00,
			}
			targetRectangle.Rectangle.StrokeColor = strokeColor

		}
	}
}

func executeDropAction() {

	// In what target was it dropped
	switch stateMachineDragAndDrop.targetDroppedType {

	case droppableLabelType:

		fmt.Println(fmt.Sprintf("'%s' was dropped in '%s'. Current TestCase is '%s'", stateMachineDragAndDrop.SourceUuid, stateMachineDragAndDrop.targetDroppableLabel.TargetUuid, stateMachineDragAndDrop.targetDroppableLabel.CurrentTestCaseUuid))

		commandEngineChannelMessage := sharedCode.ChannelCommandStruct{
			ChannelCommand:  sharedCode.ChannelCommandSwapElement,
			FirstParameter:  stateMachineDragAndDrop.targetDroppableLabel.TargetUuid,
			SecondParameter: stateMachineDragAndDrop.SourceUuid,
			ActiveTestCase:  stateMachineDragAndDrop.targetDroppableLabel.CurrentTestCaseUuid,
			ElementType:     sharedCode.BuildingBlock(stateMachineDragAndDrop.SourceType),
		}

		// Send command message over channel to Command and Rule Engine
		*commandChannelReference <- commandEngineChannelMessage

	case droppableRectangleType:

		fmt.Println(fmt.Sprintf("'%s' was dropped in '%s'. Current TestCase is '%s'", stateMachineDragAndDrop.SourceUuid, stateMachineDragAndDrop.targetDroppableRectangle.TargetUuid, stateMachineDragAndDrop.targetDroppableRectangle.CurrentTestCaseUuid))

		commandEngineChannelMessage := sharedCode.ChannelCommandStruct{
			ChannelCommand:  sharedCode.ChannelCommandSwapElement,
			FirstParameter:  stateMachineDragAndDrop.targetDroppableRectangle.TargetUuid,
			SecondParameter: stateMachineDragAndDrop.SourceUuid,
			ActiveTestCase:  stateMachineDragAndDrop.targetDroppableRectangle.CurrentTestCaseUuid,
			ElementType:     sharedCode.BuildingBlock(stateMachineDragAndDrop.SourceType),
		}

		// Send command message over channel to Command and Rule Engine
		*commandChannelReference <- commandEngineChannelMessage

	default:
		errorID := "37b775a6-6a06-4022-94a0-22631a6c9286"
		err := errors.New(fmt.Sprintf("Unknown or unhandled targetDroppedType: '%d' . [ErrorID:'%s']", stateMachineDragAndDrop.targetDroppedType, errorID))

		fmt.Println(err) // TODO Send on Error-channel

	}

}
