package listTestSuitesUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/soundEngine"
	"FenixTesterGui/testCase/testCaseModel"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

type clickableTableLabel struct {
	widget.Label
	onDoubleTap          func()
	lastTapTime          time.Time
	isClickable          bool
	currentRow           int16
	currentTestSuiteUuid string
	background           *canvas.Rectangle
	testCasesModel       *testCaseModel.TestCasesModelsStruct
	textInsteadOfLabel   *canvas.Text
	listTestSuiteUIPtr   *ListTestSuiteUIStruct
}

func newClickableTableLabel(text string, onDoubleTap func(), tempIsClickable bool,
	testCasesModel *testCaseModel.TestCasesModelsStruct, listTestSuiteUI *ListTestSuiteUIStruct) *clickableTableLabel {

	l := &clickableTableLabel{
		Label:              widget.Label{Text: text},
		onDoubleTap:        onDoubleTap,
		lastTapTime:        time.Now(),
		isClickable:        tempIsClickable,
		currentRow:         -1,
		listTestSuiteUIPtr: listTestSuiteUI}

	l.background = canvas.NewRectangle(color.Transparent)
	l.testCasesModel = testCasesModel
	l.currentTestSuiteUuid = ""

	l.ExtendBaseWidget(l)

	l.textInsteadOfLabel = &canvas.Text{
		Alignment: fyne.TextAlignCenter,
		Color: color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0xFF,
		},
		Text:     "",
		TextSize: theme.TextSize(),
		TextStyle: fyne.TextStyle{
			Bold:      false,
			Italic:    false,
			Monospace: false,
			Symbol:    false,
			TabWidth:  0,
		},
	}
	return l

}

func (l *clickableTableLabel) Tapped(e *fyne.PointEvent) {
	if l.isClickable == false {
		return
	}

	if mouseHasLeftTable == true {
		return
	}

	if time.Since(l.lastTapTime) < 500*time.Millisecond {
		if l.onDoubleTap != nil {
			l.onDoubleTap()

			l.lastTapTime = time.Now()

			return
		}
	}

	l.lastTapTime = time.Now()

	// Update TestSuite Preview
	l.listTestSuiteUIPtr.preViewAndFilterTabs.Select(l.listTestSuiteUIPtr.preViewTab)
	l.listTestSuiteUIPtr.GenerateTestSuitePreviewContainer(l.currentTestSuiteUuid, l.testCasesModel)
	l.listTestSuiteUIPtr.testSuiteThatIsShownInPreview = l.currentTestSuiteUuid
	l.listTestSuiteUIPtr.testSuiteListTable.Refresh()

}

// TappedSecondary
// Implement if you need right-click (secondary tap) actions.
func (l *clickableTableLabel) TappedSecondary(*fyne.PointEvent) {
	if l.isClickable == false {
		return
	}

	if mouseHasLeftTable == true {
		return
	}

	fenixMasterWindow := *sharedCode.FenixMasterWindowPtr
	clipboard := fenixMasterWindow.Clipboard()
	clipboard.SetContent(l.Text)

	// Notify the user

	// Trigger System Notification sound
	soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

	// If the message is too long, then cut it
	var textToShowInNotification string
	if len(textToShowInNotification) > sharedCode.TextToShowInNotificationLength {
		textToShowInNotification = l.Text[:sharedCode.TextToShowInNotificationLength] + "..."
	} else {
		textToShowInNotification = l.Text
	}

	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   "Clipboard",
		Content: fmt.Sprintf("'%s' copied to clipboard!", textToShowInNotification),
	})

}

func (l *clickableTableLabel) MouseIn(*desktop.MouseEvent) {

	if l.isClickable == false {
		return
	}

	if mouseHasLeftTable == true {
		return
	}

	// Hinder concurrent setting of variable
	l.listTestSuiteUIPtr.currentRowThatMouseIsHoveringAboveMutex.Lock()

	// Set current TestSuiteUuid to be highlighted
	l.listTestSuiteUIPtr.currentRowThatMouseIsHoveringAbove = l.currentRow

	// Release variable
	l.listTestSuiteUIPtr.currentRowThatMouseIsHoveringAboveMutex.Unlock()

	l.TextStyle = fyne.TextStyle{Bold: true}
	l.Refresh()
	//l.listTestSuiteUIPtr.testSuiteListTable.Refresh()

}
func (l *clickableTableLabel) MouseMoved(*desktop.MouseEvent) {}
func (l *clickableTableLabel) MouseOut() {

	if l.isClickable == false {
		return
	}

	if mouseHasLeftTable == true {
		return
	}

	// Hinder concurrent setting of variable
	l.listTestSuiteUIPtr.currentRowThatMouseIsHoveringAboveMutex.Lock()

	// Set current TestSuiteUuid to be highlighted
	l.listTestSuiteUIPtr.currentRowThatMouseIsHoveringAbove = -1

	// Release variable
	l.listTestSuiteUIPtr.currentRowThatMouseIsHoveringAboveMutex.Unlock()

	l.TextStyle = fyne.TextStyle{Bold: false}
	l.Refresh()
	//l.listTestSuiteUIPtr.testSuiteListTable.Refresh()

}
