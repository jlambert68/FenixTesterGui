package listTestCaseExecutionsUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/soundEngine"
	"FenixTesterGui/testCaseExecutions/testCaseExecutionsModel"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"image/color"
	"log"
	"strings"
	"time"
)

const (
	notDefined labelTypeType = iota
	labelIsTestInstruction
	labelIsTestInstructionContainer
)

type labelTypeType uint8

// Definition of a clickable TestInstructionName or TestInstructionContainer Name label used in the TestCaseExecution-Preview
type clickableTInTICNameLabelInPreviewStruct struct {
	widget.Label
	TCExecutionKey     testCaseExecutionsModel.DetailedTestCaseExecutionMapKeyType
	TInTICExecutionKey testCaseExecutionsModel.
				TCEoTICoTIEAttributesContainerMapKeyType
	LeftClicked  func()
	RightClicked func()
	LabelType    labelTypeType
	lastTapTime  time.Time
}

// Used for creating a new TestInstructionName label
func newClickableTestInstructionNameLabelInPreview(
	tInTICName string,
	tCExecutionKey testCaseExecutionsModel.DetailedTestCaseExecutionMapKeyType,
	tInTICExecutionKey testCaseExecutionsModel.
		TCEoTICoTIEAttributesContainerMapKeyType,
	leftClicked func(),
	rightClicked func(),
	labelType labelTypeType,
) *clickableTInTICNameLabelInPreviewStruct {

	clickableTInTICNameLabelInPreview := &clickableTInTICNameLabelInPreviewStruct{
		Label:              widget.Label{Text: tInTICName},
		TCExecutionKey:     tCExecutionKey,
		TInTICExecutionKey: tInTICExecutionKey,
		LeftClicked:        leftClicked,
		RightClicked:       rightClicked,
		LabelType:          labelType,
		lastTapTime:        time.Now(),
	}

	clickableTInTICNameLabelInPreview.ExtendBaseWidget(clickableTInTICNameLabelInPreview)

	return clickableTInTICNameLabelInPreview
}

// CreateRenderer
// Renderer (required by fyne.Widget)
func (c *clickableTInTICNameLabelInPreviewStruct) CreateRenderer() fyne.WidgetRenderer {
	lbl := widget.NewLabel(c.Label.Text)
	return widget.NewSimpleRenderer(lbl)
}

// Tapped
// Tapped interface implementation
func (c *clickableTInTICNameLabelInPreviewStruct) Tapped(*fyne.PointEvent) {

	fmt.Println("LeftClicked")

	testCaseExecutionAttributesForPreviewMapMutex.Lock()

	var existInMap bool
	var attributesContainerPtr *fyne.Container
	var testCaseExecutionAttributesForPreviewObjectPtr *testCaseExecutionAttributesForPreviewStruct
	var testCaseExecutionAttributesForPreviewMap map[testCaseExecutionsModel.TCEoTICoTIEAttributesContainerMapKeyType]*testCaseExecutionAttributesForPreviewStruct
	var testInstructionExecutionLogPostMapKeys []testCaseExecutionsModel.TestInstructionExecutionLogPostMapKeyType

	testCaseExecutionAttributesForPreviewMap = *testCaseExecutionAttributesForPreviewMapPtr
	testCaseExecutionAttributesForPreviewObjectPtr, existInMap = testCaseExecutionAttributesForPreviewMap[c.TInTICExecutionKey]

	if existInMap == false {

		// Should never happen
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                   "765dd5e9-0f00-4494-8128-33c986c5b13d",
			"c.TInTICExecutionKey": c.TInTICExecutionKey,
		}).Error("Couldn't find object in  for 'testCaseExecutionAttributesForPreviewMap', should never happen")

		testCaseExecutionAttributesForPreviewMapMutex.Unlock()
		return
	}

	// Add this TI or TIC execution-key to slice of execution key to extract log-post data for
	//testInstructionExecutionLogPostMapKeys = append(testInstructionExecutionLogPostMapKeys,
	//	testCaseExecutionsModel.TestInstructionExecutionLogPostMapKeyType(c.TInTICExecutionKey))

	attributesContainerPtr = testCaseExecutionAttributesForPreviewObjectPtr.testInstructionExecutionAttributesContainer

	switch testCaseExecutionAttributesForPreviewObjectPtr.attributesContainerShouldBeVisible {

	// Hide attributes
	case true:

		testCaseExecutionAttributesForPreviewObjectPtr.attributesContainerShouldBeVisible = false

		// Hide attributes for object that was clicked on
		if testCaseExecutionAttributesForPreviewObjectPtr.testInstructionExecutionAttributesContainer != nil {

			attributesContainerPtr.Hide()
		}

		// Hide attributes for child-objects to object that was clicked on
		for _, childObjectKey := range testCaseExecutionAttributesForPreviewObjectPtr.childObjectsWithAttributes {

			var childTestCaseExecutionAttributesForPreviewObjectPtr *testCaseExecutionAttributesForPreviewStruct
			childTestCaseExecutionAttributesForPreviewObjectPtr, existInMap = testCaseExecutionAttributesForPreviewMap[childObjectKey]
			if existInMap == false {

				// Should never happen
				sharedCode.Logger.WithFields(logrus.Fields{
					"ID":             "222087c5-1c93-4e8f-8862-af6baf1ae2ae",
					"childObjectKey": childObjectKey,
				}).Error("Couldn't find child for TestInstructionAttributes, should never happen")

				testCaseExecutionAttributesForPreviewMapMutex.Unlock()

				return
			}

			childTestCaseExecutionAttributesForPreviewObjectPtr.attributesContainerShouldBeVisible = false
			if childTestCaseExecutionAttributesForPreviewObjectPtr.testInstructionExecutionAttributesContainer != nil {
				childTestCaseExecutionAttributesForPreviewObjectPtr.testInstructionExecutionAttributesContainer.Hide()
			}
		}

	// Show attributes
	case false:

		testCaseExecutionAttributesForPreviewObjectPtr.attributesContainerShouldBeVisible = true

		// Show attributes for object that was clicked on
		if testCaseExecutionAttributesForPreviewObjectPtr.testInstructionExecutionAttributesContainer != nil {

			attributesContainerPtr.Show()
		}

		// Hide attributes for child-objects to object that was clicked on
		for _, childObjectKey := range testCaseExecutionAttributesForPreviewObjectPtr.childObjectsWithAttributes {

			var childTestCaseExecutionAttributesForPreviewObjectPtr *testCaseExecutionAttributesForPreviewStruct
			childTestCaseExecutionAttributesForPreviewObjectPtr, existInMap = testCaseExecutionAttributesForPreviewMap[childObjectKey]
			if existInMap == false {

				// Should never happen
				sharedCode.Logger.WithFields(logrus.Fields{
					"ID":             "75369175-4406-4127-95e6-6171a73aae27",
					"childObjectKey": childObjectKey,
				}).Error("Couldn't find child for TestInstructionAttributes, should never happen")

				testCaseExecutionAttributesForPreviewMapMutex.Unlock()

				return
			}

			childTestCaseExecutionAttributesForPreviewObjectPtr.attributesContainerShouldBeVisible = true
			if childTestCaseExecutionAttributesForPreviewObjectPtr.testInstructionExecutionAttributesContainer != nil {
				childTestCaseExecutionAttributesForPreviewObjectPtr.testInstructionExecutionAttributesContainer.Show()
			}

			// Add this TI or TIC execution-key to slice of execution key to extract log-post data for
			//testInstructionExecutionLogPostMapKeys = append(testInstructionExecutionLogPostMapKeys,
			//	testCaseExecutionsModel.TestInstructionExecutionLogPostMapKeyType(childObjectKey))

		}

	}

	testCaseExecutionPreviewContainerScroll.Refresh()

	if testCaseExecutionAttributesForPreviewObjectPtr.testInstructionExecutionAttributesContainer != nil {
		attributesContainerPtr.Refresh()
	}

	// Loop TI and TIC and check if they are expanded or not, and add to slice for objects to get ExecutionLog data from
	for tempChildObjectKey, tempChildTestCaseExecutionAttributesForPreviewObjectPtr := range testCaseExecutionAttributesForPreviewMap {

		// If TI or TIC is visible/expanded then add to slice
		if tempChildTestCaseExecutionAttributesForPreviewObjectPtr.attributesContainerShouldBeVisible == true &&
			tempChildTestCaseExecutionAttributesForPreviewObjectPtr.LabelType == labelIsTestInstruction {
			// Add this TI execution-key to slice of execution key to extract log-post data for
			testInstructionExecutionLogPostMapKeys = append(testInstructionExecutionLogPostMapKeys,
				testCaseExecutionsModel.TestInstructionExecutionLogPostMapKeyType(tempChildObjectKey))
		}
	}

	// Extract log-data from clicked object and its children
	var logPostAndValuesMessagesPtr *[]*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage
	logPostAndValuesMessagesPtr = testCaseExecutionsModel.TestCaseExecutionsModel.
		ListLogPostsAndValuesForTestInstructionExecutions(
			c.TCExecutionKey,
			testInstructionExecutionLogPostMapKeys)

	if logPostAndValuesMessagesPtr != nil && len(*logPostAndValuesMessagesPtr) > 0 {

		// Variables for extracted data from log-post-message
		var logpostStatus fenixExecutionServerGuiGrpcApi.LogPostStatusEnum
		var logpostStatusAsText string
		var logPostTimeStamp time.Time
		var logPostTimeStampAsText string
		var testInstructionExecutionUuid string
		var testInstructionExecutionMapKey string
		var testInstructionExecutionVersion int32
		var logPostText string
		//var testInstructionUuid testCaseExecutionsModel.RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType,
		var testInstructionNameForLog string
		var foundVersusExpectedValue []*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage_FoundVersusExpectedValueMessage

		// Variables for the log-post-container
		var logPostFormContainer *fyne.Container

		logPostFormContainer = container.New(layout.NewFormLayout())

		// Loop over all log-post messages and create the RichText instance for each log-post and add to a Form-container
		for _, logPostAndValuesMessage := range *logPostAndValuesMessagesPtr {

			var timStampAndStatusRichText *widget.RichText
			var logMessageRichText *widget.RichText

			//var timStampAndTextStatusContainer *fyne.Container
			var timStampAndStatusContainer *fyne.Container
			//var logMessageContainerScroll *container.Scroll
			var logMessageBorderContainer *fyne.Container

			//timStampAndTextStatusContainer =  container.New(layout.NewVBoxLayout())
			timStampAndStatusContainer = container.New(layout.NewHBoxLayout())

			// Build a markdown-formatted string dynamically
			var timeStampStringBuilder strings.Builder
			var logMessageStringBuilder strings.Builder

			// Extract the data from 'logPostAndValuesMessage'
			logpostStatus = logPostAndValuesMessage.GetLogPostStatus()
			logpostStatusAsText = logpostStatus.String()
			logPostTimeStamp = logPostAndValuesMessage.GetLogPostTimeStamp().AsTime()
			logPostTimeStampAsText = logPostTimeStamp.Format(time.RFC3339)
			testInstructionExecutionUuid = logPostAndValuesMessage.GetTestInstructionExecutionUuid()
			testInstructionExecutionVersion = logPostAndValuesMessage.GetTestInstructionExecutionVersion()
			logPostText = logPostAndValuesMessage.GetLogPostText()
			foundVersusExpectedValue = logPostAndValuesMessage.GetFoundVersusExpectedValue()
			testInstructionExecutionMapKey = fmt.Sprintf("%s%d", testInstructionExecutionUuid, testInstructionExecutionVersion)
			_, testInstructionNameForLog, existInMap =
				testCaseExecutionsModel.TestCaseExecutionsModel.GetTestInstructionFromTestInstructionExecutionUuid(
					testCaseExecutionsModel.TestCaseExecutionUuidType(c.TCExecutionKey),
					testCaseExecutionsModel.TestInstructionExecutionUuidType(testInstructionExecutionMapKey),
					true)

			if existInMap == false {
				testInstructionNameForLog = "couldn't find test instruction"
			}

			/*
				richText := widget.NewRichText(
					&widget.TextSegment{
						Text:  "Normal text segment. ",
						Style: widget.RichTextStyleInline,
					},
					&widget.TextSegment{
						Text: "Red colored text. ",
						Style: widget.RichTextStyle{
							Inline: true,
							ColorName: theme.ColorNameError, // Built-in theme color (red-ish)
						},
					},
					&widget.TextSegment{
						Text: "Custom blue color text.",
						Style: widget.RichTextStyle{
							Inline: true,
							ColorName: theme.ColorGreen, // Custom Blue
						},
					},
				)
			*/
			// Format TimeStamp information using Markdown syntax
			timeStampStringBuilder.WriteString(fmt.Sprintf("%s\n\n%s", logPostTimeStampAsText, logpostStatusAsText))

			// Create RichText widget from the generated markdown content for TimeStamp and Status
			timStampAndStatusRichText = widget.NewRichTextFromMarkdown(timeStampStringBuilder.String())

			// Add TimeStamp And Status-text to container
			timStampAndStatusContainer.Add(timStampAndStatusRichText)

			// Create Status-color-box
			var logStatusRectangle *canvas.Rectangle
			logStatusRectangle = canvas.NewRectangle(color.Transparent)

			// Resize the ExecutionStatus rectangle
			logStatusRectangle.SetMinSize(fyne.Size{
				Width:  logStatusRectangleWidth,
				Height: logStatusRectangleHeight,
			})
			logStatusRectangle.Resize(fyne.Size{
				Width:  logStatusRectangleWidth,
				Height: logStatusRectangleHeight,
			})

			// Set border width
			logStatusRectangle.StrokeWidth = 4

			// Set correct color for the Status-color-box
			switch logpostStatus {

			case fenixExecutionServerGuiGrpcApi.LogPostStatusEnum_INFO:
				// Transparent box
				logStatusRectangle.StrokeColor = color.Transparent
				logStatusRectangle.FillColor = color.Transparent

			case fenixExecutionServerGuiGrpcApi.LogPostStatusEnum_WARNING:
				// Yellow box
				logStatusRectangle.StrokeColor = color.RGBA{
					R: 0xFF,
					G: 0xFF,
					B: 0x00,
					A: 0xFF,
				}
				logStatusRectangle.FillColor = color.Transparent

			case fenixExecutionServerGuiGrpcApi.LogPostStatusEnum_EXECUTION_OK:
				logStatusRectangle.StrokeColor = color.Transparent
				logStatusRectangle.FillColor = color.Transparent

			case fenixExecutionServerGuiGrpcApi.LogPostStatusEnum_EXECUTION_ERROR:
				// Filled Red box
				logStatusRectangle.StrokeColor = color.RGBA{
					R: 0xFF,
					G: 0x00,
					B: 0x00,
					A: 0xFF,
				}
				logStatusRectangle.FillColor = color.RGBA{
					R: 0xFF,
					G: 0x00,
					B: 0x00,
					A: 0xFF,
				}

			case fenixExecutionServerGuiGrpcApi.LogPostStatusEnum_VALIDATION_OK:
				// Transparent box with Green border
				logStatusRectangle.StrokeColor = color.RGBA{
					R: 0x00,
					G: 0xFF,
					B: 0x00,
					A: 0xFF,
				}
				logStatusRectangle.FillColor = color.Transparent

			case fenixExecutionServerGuiGrpcApi.LogPostStatusEnum_VALIDATION_ERROR:
				// Transparent box with Red border
				logStatusRectangle.StrokeColor = color.RGBA{
					R: 0xFF,
					G: 0x00,
					B: 0x00,
					A: 0xFF,
				}
				logStatusRectangle.FillColor = color.Transparent

			case fenixExecutionServerGuiGrpcApi.LogPostStatusEnum_LogPostStatusEnum_DEFAULT_NOT_SET:

				log.Fatalf("LogPostStatusEnum was not set: %v, should never happen", logpostStatus)

			default:
				log.Fatalf("Unknown LogPostStatusEnum: %v, should never happen", logpostStatus)

			}

			// Add spacer
			timStampAndStatusContainer.Add(layout.NewSpacer())

			// Add Status-box  to container
			timStampAndStatusContainer.Add(logStatusRectangle)

			// Format message information using Markdown syntax
			logMessageStringBuilder.WriteString(fmt.Sprintf("## %s\n\n", testInstructionNameForLog))
			logMessageStringBuilder.WriteString(fmt.Sprintf("*TIEUuid: %s(%d)* \n\n ", testInstructionExecutionUuid, testInstructionExecutionVersion))
			logMessageStringBuilder.WriteString(fmt.Sprintf("LogText: *%s*", logPostText))

			// Add Found vs Expected values, if exist
			if len(foundVersusExpectedValue) > 0 {

				var variableName string
				var variableDescription string
				var foundValue string
				var expectedValue string

				// Loop all found vs expected value pars
				if len(foundVersusExpectedValue) > 0 {
					logMessageStringBuilder.WriteString(" \n\n *** \n\n **Expected vs Found** \n\n ")
				}
				for _, foundVersusExpectedValueMessage := range foundVersusExpectedValue {

					foundValue = foundVersusExpectedValueMessage.GetFoundValue()
					expectedValue = foundVersusExpectedValueMessage.GetExpectedValue()
					variableName = foundVersusExpectedValueMessage.GetVariableName()
					variableDescription = foundVersusExpectedValueMessage.GetVariableDescription()

					logMessageStringBuilder.WriteString(" \n\n *** \n\n ")
					logMessageStringBuilder.WriteString(fmt.Sprintf("**%s:** \n\n ", variableName))
					logMessageStringBuilder.WriteString(fmt.Sprintf("Description: *%s* \n\n ", variableDescription))
					logMessageStringBuilder.WriteString(fmt.Sprintf("Found:      *%s* \n\n ", foundValue))
					logMessageStringBuilder.WriteString(fmt.Sprintf("Expected: *%s*", expectedValue))
				}

			}

			// Create RichText widget from the generated markdown content for Log-message
			logMessageRichText = widget.NewRichTextFromMarkdown(logMessageStringBuilder.String())

			// Make the log do soft line breaks
			logMessageRichText.Wrapping = fyne.TextWrapWord

			// Put Log-message-richText into a scrollable container
			//logMessageContainerScroll = container.NewScroll(
			//	container.NewBorder(nil, nil, nil, nil, logMessageRichText))
			logMessageBorderContainer = container.NewBorder(nil, nil, nil, nil, logMessageRichText)

			// Add 'timStampAndStatusContainer' to 'logPostFormContainer'
			logPostFormContainer.Add(timStampAndStatusContainer)

			// Add 'logMessageContainerScroll' to 'logPostFormContainer'
			logPostFormContainer.Add(logMessageBorderContainer)

			// Add the updated Scroll container to the Border container for the logs
			testInstructionsExecutionLogContainer.Objects[0] = container.NewBorder(nil, nil, nil, nil, logPostFormContainer)

			// Update GUI for logs
			testInstructionsExecutionLogContainer.Refresh()

		}
	} else {

		// Create a new temporary container for the logs
		testInstructionsExecutionLogContainer.Objects[0] = container.NewCenter(
			widget.NewLabel("Select a TestInstructionExecution or a TesInstructionContainer to get the Logs"))

		// Update GUI for logs
		testInstructionsExecutionLogContainer.Refresh()
	}

	testCaseExecutionAttributesForPreviewMapMutex.Unlock()

	if c.LeftClicked != nil {

		//c.LeftClicked()
	}
}

// TappedSecondary
// Optional: Handle secondary tap (right-click)
func (c *clickableTInTICNameLabelInPreviewStruct) TappedSecondary(*fyne.PointEvent) {

	fmt.Println("RightClicked")

	fenixMasterWindow := *sharedCode.FenixMasterWindowPtr
	clipboard := fenixMasterWindow.Clipboard()
	clipboard.SetContent(c.Text)

	// Notify the user

	// Trigger System Notification sound
	soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   "Clipboard",
		Content: fmt.Sprintf("'%s' copied to clipboard!", c.Text),
	})

	// handle secondary tap if needed
	if c.RightClicked != nil {
		//c.RightClicked()
	}
}
