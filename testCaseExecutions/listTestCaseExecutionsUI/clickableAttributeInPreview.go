package listTestCaseExecutionsUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/soundEngine"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strings"
	"time"
)

const (
	attributeTypeNotDefined attributeTypeType = iota
	attributeIsOriginal
	attributeIsRunTimeChanged
)

type attributeTypeType uint8

// Definition of a clickable Attribute used in the TestCaseExecution-Preview
type clickableAttributeInPreviewStruct struct {
	widget.Label
	AttributeName       string
	TestInstructionName string
	LeftClicked         func()
	RightClicked        func()
	lastTapTime         time.Time
	AttributeType       attributeTypeType
}

// Used for creating a new Attribute label
func newClickableAttributeInPreview(
	attributeValue string,
	attributeName string,
	testInstructionName string,
	leftClicked func(),
	rightClicked func(),
	attributeType attributeTypeType,
) *clickableAttributeInPreviewStruct {

	clickableAttributeInPreview := &clickableAttributeInPreviewStruct{
		Label:               widget.Label{Text: attributeValue},
		AttributeName:       attributeName,
		TestInstructionName: testInstructionName,
		LeftClicked:         leftClicked,
		RightClicked:        rightClicked,
		lastTapTime:         time.Now(),
		AttributeType:       attributeType,
	}

	clickableAttributeInPreview.ExtendBaseWidget(clickableAttributeInPreview)

	return clickableAttributeInPreview
}

// CreateRenderer
// Renderer (required by fyne.Widget)
func (c *clickableAttributeInPreviewStruct) CreateRenderer() fyne.WidgetRenderer {
	lbl := widget.NewLabel(c.Label.Text)
	return widget.NewSimpleRenderer(lbl)
}

// Tapped
// Tapped interface clickableAttributeInPreviewStruct
func (c *clickableAttributeInPreviewStruct) Tapped(*fyne.PointEvent) {

	// Create the Attribute information to present in Attribute-exploret
	var attributeMessageStringBuilder strings.Builder
	var attributeMessageRichText *widget.RichText
	var attributeMessageBorderContainer *fyne.Container
	var attributeMessageContainer *fyne.Container
	var attributeMessageScrollContainer *container.Scroll
	var copyAttributeClipBoardContainer *fyne.Container

	// Create the 'copyLogToClipBoardContainer'
	copyAttributeClipBoardContainer = container.NewVBox(
		widget.NewButton("Copy attribute to clipboard", func() {

			var textToCopy strings.Builder

			fenixMasterWindow := *sharedCode.FenixMasterWindowPtr
			clipboard := fenixMasterWindow.Clipboard()

			textToCopy.WriteString(fmt.Sprintf("TestInstruction: %s \n\n", c.TestInstructionName))
			textToCopy.WriteString(fmt.Sprintf("Attribute: %s \n\n", c.AttributeName))
			textToCopy.WriteString(fmt.Sprintf("%s", c.Text))

			clipboard.SetContent(textToCopy.String())

			// Notify the user

			// Trigger System Notification sound
			soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Clipboard",
				Content: fmt.Sprintf("Attribute was copied to clipboard!"),
			})
		}))

	// Format message information using Markdown syntax
	attributeMessageStringBuilder.WriteString(fmt.Sprintf("## TestInstruction: *%s* \n\n", c.TestInstructionName))
	attributeMessageStringBuilder.WriteString(fmt.Sprintf("### Attribute: *%s* \n\n ", c.AttributeName))

	//attributeMessageStringBuilder.WriteString(fmt.Sprintf("%s", c.Text))

	// Create RichText widget from the generated markdown content for Attribute
	attributeMessageRichText = widget.NewRichTextFromMarkdown(attributeMessageStringBuilder.String())

	// Make the log do soft line breaks
	attributeMessageRichText.Wrapping = fyne.TextWrapOff

	// Create attributeMessageContainer containing the attribute information
	attributeMessageContainer = container.NewVBox(attributeMessageRichText, widget.NewLabel(c.Text))

	// Add 'attributeMessageRichText' to scroll container
	attributeMessageScrollContainer = container.NewScroll(attributeMessageContainer)

	// Build container containing all attribute-explorer-objects
	attributeMessageBorderContainer = container.NewBorder(copyAttributeClipBoardContainer, nil, nil, nil, attributeMessageScrollContainer)

	// Add attribute-explorer to tab
	attributeExplorerTab.Content = attributeMessageBorderContainer
	attributeMessageBorderContainer.Refresh()
	preViewTabs.Refresh()

	if c.LeftClicked != nil {

		c.LeftClicked()
	}
}

// TappedSecondary
// Optional: Handle secondary tap (right-click)
func (c *clickableAttributeInPreviewStruct) TappedSecondary(*fyne.PointEvent) {

	fenixMasterWindow := *sharedCode.FenixMasterWindowPtr
	clipboard := fenixMasterWindow.Clipboard()
	clipboard.SetContent(c.Text)

	// Notify the user

	// Trigger System Notification sound
	soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

	// If the message is too long, then cut it
	var textToShowInNotification string
	if len(textToShowInNotification) > sharedCode.TextToShowInNotificationLength {
		textToShowInNotification = c.Text[:sharedCode.TextToShowInNotificationLength] + "..."
	} else {
		textToShowInNotification = c.Text
	}

	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   "Clipboard",
		Content: fmt.Sprintf("'%s' copied to clipboard!", textToShowInNotification),
	})

	// handle secondary tap if needed
	if c.RightClicked != nil {
		//c.RightClicked()
	}
}
