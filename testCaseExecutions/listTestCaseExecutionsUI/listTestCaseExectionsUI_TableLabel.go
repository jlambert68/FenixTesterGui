package listTestCaseExecutionsUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/soundEngine"
	"FenixTesterGui/testCaseExecutions/testCaseExecutionsModel"
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
	onDoubleTap                  func()
	lastTapTime                  time.Time
	isClickable                  bool
	currentRow                   int16
	currentTestCaseExecutionUuid string
	currentTestCaseUuid          string
	currentTestCaseName          string
	background                   *canvas.Rectangle
	testCaseExecutionsModel      *testCaseExecutionsModel.TestCaseExecutionsModelStruct
	textInsteadOfLabel           *canvas.Text
}

func newClickableTableLabel(text string, onDoubleTap func(), tempIsClickable bool,
	testCaseExecutionsModel *testCaseExecutionsModel.TestCaseExecutionsModelStruct) *clickableTableLabel {

	l := &clickableTableLabel{
		Label:       widget.Label{Text: text},
		onDoubleTap: onDoubleTap,
		lastTapTime: time.Now(),
		isClickable: tempIsClickable,
		currentRow:  -1}

	l.background = canvas.NewRectangle(color.Transparent)
	l.testCaseExecutionsModel = testCaseExecutionsModel
	l.currentTestCaseExecutionUuid = ""
	l.currentTestCaseUuid = ""
	l.currentTestCaseName = ""

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

	if time.Since(l.lastTapTime) < 500*time.Millisecond {
		if l.onDoubleTap != nil {
			l.onDoubleTap()

			l.lastTapTime = time.Now()

			return
		}
	}

	l.lastTapTime = time.Now()

	// Update TestCase Preview
	testCaseExecutionThatIsShownInPreview = l.currentTestCaseExecutionUuid
	GenerateTestCaseExecutionPreviewContainer(l.currentTestCaseExecutionUuid, l.testCaseExecutionsModel)
	testCaseExecutionsListTable.Refresh()

	// Update 'loadAllTestCaseExecutionsForOneTestCaseButton' with correct text
	loadAllTestCaseExecutionsForOneTestCaseButtonReference.
		Text = loadAllTestCaseExecutionsForOneTestCaseButtonTextPart1 +
		l.currentTestCaseName
	loadAllTestCaseExecutionsForOneTestCaseButtonReference.Refresh()

	// Save TestCaseUuid for TestCaseExecution shown in preview
	testCaseUuidForTestCaseExecutionThatIsShownInPreview = l.currentTestCaseUuid

}

// TappedSecondary
// Implement if you need right-click (secondary tap) actions.
func (l *clickableTableLabel) TappedSecondary(*fyne.PointEvent) {
	if l.isClickable == false {
		return
	}

	fenixMasterWindow := *sharedCode.FenixMasterWindowPtr
	clipboard := fenixMasterWindow.Clipboard()
	clipboard.SetContent(l.Text)

	// Notify the user

	// Trigger System Notification sound
	soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   "Clipboard",
		Content: fmt.Sprintf("'%s' copied to clipboard!", l.Text),
	})

}

func (l *clickableTableLabel) MouseIn(*desktop.MouseEvent) {

	if l.isClickable == false {
		return
	}

	// Hinder concurrent setting of variable
	currentRowThatMouseIsHoveringAboveMutex.Lock()

	// Set current TestCaseUuid to be highlighted
	currentRowThatMouseIsHoveringAbove = l.currentRow

	// Release variable
	currentRowThatMouseIsHoveringAboveMutex.Unlock()

	l.TextStyle = fyne.TextStyle{Bold: true}
	l.Refresh()
	testCaseExecutionsListTable.Refresh()

}
func (l *clickableTableLabel) MouseMoved(*desktop.MouseEvent) {}
func (l *clickableTableLabel) MouseOut() {

	if l.isClickable == false {
		return
	}

	// Hinder concurrent setting of variable
	currentRowThatMouseIsHoveringAboveMutex.Lock()

	// Set current TestCaseUuid to be highlighted
	currentRowThatMouseIsHoveringAbove = -1

	// Release variable
	currentRowThatMouseIsHoveringAboveMutex.Unlock()

	l.TextStyle = fyne.TextStyle{Bold: false}
	l.Refresh()
	testCaseExecutionsListTable.Refresh()

}
