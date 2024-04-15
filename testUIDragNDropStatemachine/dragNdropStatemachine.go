package testUIDragNDropStatemachine

import (
	sharedCode "FenixTesterGui/common_code"
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
	commandChannelRef *sharedCode.CommandChannelType) {

	textRef = dragNDropText
	rectangleRef = dragNDropRectangleRef
	rectangle2Ref = dragNDropRectangle2Ref
	containerRef = dragNDropContainerRef
	commandChannelReference = commandChannelRef

	// Calculate label standard height
	tempLabel := widget.NewLabel("temp")
	tempLabel.Refresh()
	labelStandardHeight = tempLabel.MinSize().Height

}

//****************************************************

func (stateMachine *StateMachineDragAndDropStruct) NewDraggableLabel(uuid string) *DraggableLabel {
	draggableLabel := &DraggableLabel{}
	draggableLabel.ExtendBaseWidget(draggableLabel)

	draggableLabel.SourceUuid = uuid
	draggableLabel.Text = uuid

	return draggableLabel
}

func (stateMachine *StateMachineDragAndDropStruct) NewDroppableLabel(
	labelText string,
	nodeLevel float32,
	testCaseNodeRectangleSize int,
	uuid string, testCaseUuid string) *DroppableLabel {

	droppableLabel := &DroppableLabel{}
	droppableLabel.ExtendBaseWidget(droppableLabel)

	droppableLabel.TargetUuid = uuid
	droppableLabel.Text = labelText
	droppableLabel.nodeLevel = nodeLevel
	droppableLabel.testCaseNodeRectangleSize = testCaseNodeRectangleSize
	droppableLabel.CurrentTestCaseUuid = testCaseUuid

	droppableLabel.BackgroundRectangle = canvas.NewRectangle(color.RGBA{
		R: 0x33,
		G: 0x33,
		B: 0x33,
		A: 0x22,
	})

	droppableLabel.Refresh()
	droppableLabel.BackgroundRectangle.SetMinSize(fyne.NewSize(targetDropLabelRectangleWidth, labelStandardHeight)) //(droppableLabel.Size())
	//droppableLabel.BackgroundRectangle.Hide()

	stateMachineDragAndDrop.registeredDroppableTargetLabels = append(stateMachineDragAndDrop.registeredDroppableTargetLabels, droppableLabel)

	droppableLabel.labelStandardHeight = droppableLabel.MinSize().Height

	return droppableLabel
}

func (stateMachine *StateMachineDragAndDropStruct) NewDroppableRectangle(
	nodeLevel float32,
	testCaseNodeRectangleSize int,
	uuid string,
	testCaseUuid string) *DroppableRectangle {

	color := color.RGBA{
		R: 0x33,
		G: 0x33,
		B: 0x33,
		A: 0x22,
	}

	rect := canvas.NewRectangle(color)
	rect.SetMinSize(fyne.NewSize(targetDropRectangleWidth, targetDropRectangleHeight))
	droppableRectangle := &DroppableRectangle{
		rect: rect,
	}

	droppableRectangle.ExtendBaseWidget(droppableRectangle)

	droppableRectangle.TargetUuid = uuid
	droppableRectangle.nodeLevel = nodeLevel
	droppableRectangle.CurrentTestCaseUuid = testCaseUuid

	stateMachineDragAndDrop.registeredDroppableTargetRectangle = append(stateMachineDragAndDrop.registeredDroppableTargetRectangle, droppableRectangle)

	return droppableRectangle
}

func newNoneDroppableLabel(uuid string) *noneDroppableLabel {
	nonDroppableLabel := &noneDroppableLabel{}
	nonDroppableLabel.ExtendBaseWidget(nonDroppableLabel)

	nonDroppableLabel.uuid = uuid
	nonDroppableLabel.Text = uuid

	return nonDroppableLabel
}

//****************************************************

type StateMachineDragAndDropStruct struct {
	sourceStateMachine                 stateMachineSourceAndDestinationStruct
	targetStateMachine                 stateMachineSourceAndDestinationStruct
	registeredDroppableTargetLabels    []*DroppableLabel
	registeredDroppableTargetRectangle []*DroppableRectangle
	SourceUuid                         string
	SourceType                         int
	targetDroppableLabel               DroppableLabel
	targetDroppableRectangle           DroppableRectangle
}

// Structure for 'Drag-part of 'Drag-N-Drop' state machine
type stateMachineSourceAndDestinationStruct struct {
	currentState int
}

var stateMachineDragAndDrop StateMachineDragAndDropStruct

func switchStateForSource(newState int) {
	stateMachineDragAndDrop.sourceStateMachine.currentState = newState
}

func switchStateForTarget(newState int) {
	stateMachineDragAndDrop.targetStateMachine.currentState = newState
}

func expandDropAreas() {

	for _, targetLabel := range stateMachineDragAndDrop.registeredDroppableTargetLabels {

		targetLabel.BackgroundRectangle.StrokeWidth = 2

		targetLabel.Resize(fyne.NewSize(targetDropRectangleWidth, targetDropRectangleHeight))
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

	for _, targetRectangle := range stateMachineDragAndDrop.registeredDroppableTargetRectangle {

		//rectangleWidth := float32(500)

		targetRectangle.rect.StrokeWidth = 2

		//targetLabel.Resize(fyne.NewSize(rectangleWidth, 12))
		strokeColor := color.RGBA{
			R: 0xFF,
			G: 0x00,
			B: 0x00,
			A: 0xAA,
		}
		targetRectangle.rect.StrokeColor = strokeColor
		targetRectangle.rect.Show()
		//targetRectangle.Show()
	}

	/*
			go func(targetReferenceLabel *DroppableLabel) {
				rectangleColorAnimation := canvas.NewColorRGBAAnimation(color.RGBA{
					R: 0x00,
					G: 0x00,
					B: 0x00,
					A: 0x00,
				}, color.RGBA{
					R: 0xFF,
					G: 0x00,
					B: 0x00,
					A: 0xAA,
				}, time.Millisecond*300, func(c color.Color) {
					targetReferenceLabel.BackgroundRectangle.StrokeColor = c
					canvas.Refresh(targetReferenceLabel.BackgroundRectangle)
				})

				rectangleSizeAnimation := canvas.NewSizeAnimation(
					fyne.NewSize(rectangleWidth, targetReferenceLabel.labelStandardHeight/2),
					fyne.NewSize(rectangleWidth, targetReferenceLabel.labelStandardHeight),
					time.Millisecond*600,
					func(animationSize fyne.Size) {
						targetReferenceLabel.BackgroundRectangle.SetMinSize(animationSize)
						canvas.Refresh(targetReferenceLabel.BackgroundRectangle)
						//targetReferenceLabel.parrentAccordion.Refresh()
						//canvas.Refresh(DropFour)
						//canvas.Refresh(dropContainer)
					})

				rectangleColorAnimation.Start()
				rectangleSizeAnimation.Start()
			}(targetLabel)

		}

		go func() {
			time.Sleep(800 * time.Millisecond)
			for _, targetLabel := range stateMachineDragAndDrop.registeredDroppableTargetLabels {
				targetLabel.Show()
			}
		}()
	*/

}

func shrinkDropAreas() {
	for _, targetLabel := range stateMachineDragAndDrop.registeredDroppableTargetLabels {

		targetLabel.BackgroundRectangle.StrokeWidth = 2
		strokeColor := color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0x00,
		}
		targetLabel.BackgroundRectangle.StrokeColor = strokeColor
		//targetLabel.BackgroundRectangle.Show()
		//targetLabel.Hide()

	}

	for _, targetRectangle := range stateMachineDragAndDrop.registeredDroppableTargetRectangle {

		targetRectangle.rect.StrokeWidth = 2
		strokeColor := color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0x00,
		}
		targetRectangle.rect.StrokeColor = strokeColor
		//targetLabel.BackgroundRectangle.Show()
		//targetRectangle.rect.Hide()

	}
	/*
			go func(targetReferenceLabel *DroppableLabel) {
				rectangleColorAnimation := canvas.NewColorRGBAAnimation(color.RGBA{
					R: 0x00,
					G: 0x00,
					B: 0x00,
					A: 0x00,
				}, color.RGBA{
					R: 0xFF,
					G: 0x00,
					B: 0x00,
					A: 0xAA,
				}, time.Millisecond*300, func(c color.Color) {
					targetReferenceLabel.BackgroundRectangle.StrokeColor = c
					canvas.Refresh(targetReferenceLabel.BackgroundRectangle)
				})

				rectangleWidth := float32(500)

				rectangleSizeAnimation := canvas.NewSizeAnimation(
					fyne.NewSize(rectangleWidth, targetReferenceLabel.labelStandardHeight),
					fyne.NewSize(rectangleWidth, targetReferenceLabel.labelStandardHeight/2),
					time.Millisecond*300,
					func(animationSize fyne.Size) {
						targetReferenceLabel.BackgroundRectangle.SetMinSize(animationSize)
						canvas.Refresh(targetReferenceLabel.BackgroundRectangle)
						//canvas.Refresh(DropFour)
						//canvas.Refresh(dropContainer)
					})

				rectangleColorAnimation.Start()
				rectangleSizeAnimation.Start()
			}(targetLabel)

		}

		go func() {
			time.Sleep(400 * time.Millisecond)
			for _, targetLabel := range stateMachineDragAndDrop.registeredDroppableTargetLabels {
				targetLabel.BackgroundRectangle.Hide()
				targetLabel.BackgroundRectangle.Refresh()

			}
		}()

	*/

}

func executeDropAction() {
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

}
