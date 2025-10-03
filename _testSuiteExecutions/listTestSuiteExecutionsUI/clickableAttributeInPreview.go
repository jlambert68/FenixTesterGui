package listTestSuiteExecutionsUI

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

const (
	numberOfVisibleRow        = 10
	numberOfVisibleCharacters = 100
	visibleCharacterOfType    = "A"
)

type attributeTypeType uint8

// Definition of a clickable Attribute used in the TestCaseExecution-Preview
type clickableAttributeInPreviewStruct struct {
	widget.Label
	AttributeName                       string
	TestInstructionName                 string
	LeftClicked                         func()
	RightClicked                        func()
	lastTapTime                         time.Time
	AttributeType                       attributeTypeType
	testCaseInstructionPreViewObjectRef *TestSuiteInstructionPreViewStruct
}

// Used for creating a new Attribute label
func newClickableAttributeInPreview(
	attributeValue string,
	attributeName string,
	testInstructionName string,
	leftClicked func(),
	rightClicked func(),
	attributeType attributeTypeType,
	testCaseInstructionPreViewObject *TestSuiteInstructionPreViewStruct,
) *clickableAttributeInPreviewStruct {

	clickableAttributeInPreview := &clickableAttributeInPreviewStruct{
		Label:                               widget.Label{Text: attributeValue},
		AttributeName:                       attributeName,
		TestInstructionName:                 testInstructionName,
		LeftClicked:                         leftClicked,
		RightClicked:                        rightClicked,
		lastTapTime:                         time.Now(),
		AttributeType:                       attributeType,
		testCaseInstructionPreViewObjectRef: testCaseInstructionPreViewObject,
	}

	clickableAttributeInPreview.ExtendBaseWidget(clickableAttributeInPreview)

	return clickableAttributeInPreview
}

// CreateRenderer
// Renderer (required by fyne.Widget)
func (c *clickableAttributeInPreviewStruct) CreateRenderer() fyne.WidgetRenderer {

	var labelScrollContainer *container.Scroll
	var labelBorderContainer *fyne.Container

	lbl := widget.NewLabel(c.Label.Text)

	// Get Size and MaxSize
	var existingLabelSize fyne.Size
	var newLabelSize fyne.Size
	var labelMaxSize fyne.Size

	// Get existing size of label
	existingLabelSize = lbl.MinSize()
	newLabelSize = existingLabelSize

	// Generate max label size
	var maxlabelStringBuilder strings.Builder

	// Generate rows
	for row := 0; row < numberOfVisibleRow; row++ {

		// When not the first then add 'new row'
		if row != 0 {
			maxlabelStringBuilder.WriteString("\n")
		}

		// Generate one row
		for characterIndex := 0; characterIndex < numberOfVisibleCharacters; characterIndex++ {
			maxlabelStringBuilder.WriteString(visibleCharacterOfType)

		}
	}

	// Extract max size
	tempMaxWidget := widget.NewLabel(maxlabelStringBuilder.String())
	tempMaxWidget.Refresh()
	labelMaxSize = tempMaxWidget.MinSize()

	// Check if current text is bigger than Max, if so then create a smaller size-object
	if existingLabelSize.Height > labelMaxSize.Height {
		newLabelSize.Height = labelMaxSize.Height
	}
	if existingLabelSize.Width > labelMaxSize.Width {
		newLabelSize.Width = labelMaxSize.Width
	}

	// Create the border container for the label
	labelBorderContainer = container.NewBorder(nil, nil, nil, nil, lbl)
	labelBorderContainer.Resize(newLabelSize)
	labelBorderContainer.Refresh()

	// Create a scroll container and set then new size
	labelScrollContainer = container.NewScroll(labelBorderContainer)
	labelScrollContainer.SetMinSize(newLabelSize)
	labelScrollContainer.Resize(newLabelSize)
	labelScrollContainer.Refresh()

	//limitedScroll := container.New(layout.NewStackLayout(), labelScrollContainer)
	//limitedScroll.Resize(newLabelSize)
	//labelScrollContainer.Refresh()

	return widget.NewSimpleRenderer(labelScrollContainer)
}

// Tapped
// Tapped interface clickableAttributeInPreviewStruct
func (c *clickableAttributeInPreviewStruct) Tapped(*fyne.PointEvent) {

	// Check if mouse has left TestCaseExecutionTree-container
	if mouseHasLeftTestSuiteExecutionPreviewTree == true {
		return
	}

	// Create the Attribute information to present in Attribute-exploret
	var attributeMessageStringBuilder strings.Builder
	var attributeMessageRichText *widget.RichText
	var attributeMessageBorderContainer *fyne.Container
	var attributeMessageContainer *fyne.Container
	var attributeMessageScrollContainer *container.Scroll
	var copyAttributeClipBoardContainer *fyne.Container

	// Create the 'copyAttributeClipBoardContainer'
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
	c.testCaseInstructionPreViewObjectRef.attributeExplorerTab.Content = attributeMessageBorderContainer
	attributeMessageBorderContainer.Refresh()
	c.testCaseInstructionPreViewObjectRef.preViewTabs.Refresh()

	if c.LeftClicked != nil {

		c.LeftClicked()
	}
}

// TappedSecondary
// Optional: Handle secondary tap (right-click)
func (c *clickableAttributeInPreviewStruct) TappedSecondary(*fyne.PointEvent) {

	// Check if mouse has left TestCaseExecutionTree-container
	if mouseHasLeftTestSuiteExecutionPreviewTree == true {
		return
	}

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
