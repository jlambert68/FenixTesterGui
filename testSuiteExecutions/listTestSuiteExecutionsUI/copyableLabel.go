package listTestSuiteExecutionsUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/soundEngine"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Definition of a simple widget label that supports to be copied
type copyableLabelStruct struct {
	widget.Label
	IsCopyable bool
}

// Used for creating a new copyable label
func newCopyableLabel(
	label string,
	isCopyable bool,
) *copyableLabelStruct {

	copyableLabel := &copyableLabelStruct{
		Label:      widget.Label{Text: label},
		IsCopyable: isCopyable,
	}

	copyableLabel.ExtendBaseWidget(copyableLabel)

	return copyableLabel
}

// CreateRenderer
// Renderer (required by fyne.Widget)
func (c *copyableLabelStruct) CreateRenderer() fyne.WidgetRenderer {
	lbl := widget.NewLabel(c.Label.Text)
	return widget.NewSimpleRenderer(lbl)
}

// Tapped
// Tapped interface clickableAttributeInPreviewStruct
func (c *copyableLabelStruct) Tapped(*fyne.PointEvent) {

}

// TappedSecondary
// Optional: Handle secondary tap (right-click)
func (c *copyableLabelStruct) TappedSecondary(*fyne.PointEvent) {

	// Check if mouse has left TestCaseExecutionTree-container
	if mouseHasLeftTestCaseExecutionPreviewTree == true {
		return
	}

	// Check if value can be copied
	if c.IsCopyable == false {
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

}
