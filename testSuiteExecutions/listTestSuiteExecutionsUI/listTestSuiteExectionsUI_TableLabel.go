package listTestSuiteExecutionsUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/soundEngine"
	"FenixTesterGui/testSuiteExecutions/testSuiteExecutionsModel"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
	"image/color"
	"time"
)

type clickableTableLabel struct {
	widget.Label
	onDoubleTap                      func()
	lastTapTime                      time.Time
	isClickable                      bool
	currentRow                       int16
	currentTestSuiteExecutionUuid    string
	currentTestSuiteExecutionVersion uint32
	currentTestSuiteUuid             string
	currentTestSuiteName             string
	background                       *canvas.Rectangle
	testSuiteExecutionsModel         *testSuiteExecutionsModel.TestSuiteExecutionsModelStruct
	textInsteadOfLabel               *canvas.Text
}

func newClickableTableLabel(text string, onDoubleTap func(), tempIsClickable bool,
	testSuiteExecutionsModel *testSuiteExecutionsModel.TestSuiteExecutionsModelStruct) *clickableTableLabel {

	l := &clickableTableLabel{
		Label:       widget.Label{Text: text},
		onDoubleTap: onDoubleTap,
		lastTapTime: time.Now(),
		isClickable: tempIsClickable,
		currentRow:  -1}

	l.background = canvas.NewRectangle(color.Transparent)
	l.testSuiteExecutionsModel = testSuiteExecutionsModel
	l.currentTestSuiteExecutionUuid = ""
	l.currentTestSuiteExecutionVersion = 0
	l.currentTestSuiteUuid = ""
	l.currentTestSuiteName = ""

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

	// Only execute 'clicks', if it is allowed
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

	// Raw data for DetailedTestSuiteExecutionMapKey
	var testSuiteExecutionUuid string
	var testSuiteExecutionVersion uint32

	// Decide mode 'AllExecutionsForOneTestSuite' or 'OneExecutionPerTestSuite'
	switch selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType {

	case AllExecutionsForOneTestSuite:

		// Save TestSuiteExecutionUuid for TestSuiteExecution shown in preview
		selectedTestSuiteExecutionObjected.allExecutionsFoOneTestSuiteListObject.isAnyRowSelected = true
		selectedTestSuiteExecutionObjected.allExecutionsFoOneTestSuiteListObject.
			testSuiteExecutionUuidThatIsShownInPreview = l.currentTestSuiteExecutionUuid
		selectedTestSuiteExecutionObjected.allExecutionsFoOneTestSuiteListObject.
			testSuiteExecutionVersionThatIsShownInPreview = l.currentTestSuiteExecutionVersion

		testSuiteExecutionUuid = selectedTestSuiteExecutionObjected.allExecutionsFoOneTestSuiteListObject.
			testSuiteExecutionUuidThatIsShownInPreview
		testSuiteExecutionVersion = selectedTestSuiteExecutionObjected.allExecutionsFoOneTestSuiteListObject.
			testSuiteExecutionVersionThatIsShownInPreview

		// Save TestSuiteUuid for TestSuiteExecution shown in preview
		selectedTestSuiteExecutionObjected.allExecutionsFoOneTestSuiteListObject.
			testSuiteUuidForTestSuiteExecutionThatIsShownInPreview = l.currentTestSuiteUuid

	case OneExecutionPerTestSuite:

		// Save TestSuiteExecutionUuid for TestSuiteExecution shown in preview
		selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.isAnyRowSelected = true
		selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.
			testSuiteExecutionUuidThatIsShownInPreview = l.currentTestSuiteExecutionUuid
		selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.
			testSuiteExecutionVersionThatIsShownInPreview = l.currentTestSuiteExecutionVersion

		testSuiteExecutionUuid = selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.
			testSuiteExecutionUuidThatIsShownInPreview
		testSuiteExecutionVersion = selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.
			testSuiteExecutionVersionThatIsShownInPreview

		// Save TestSuiteUuid for TestSuiteExecution shown in preview
		selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.
			testSuiteUuidForTestSuiteExecutionThatIsShownInPreview = l.currentTestSuiteUuid

	case NotDefined:

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "4136c74f-65c9-4ebc-977f-1c727b3a47a9",
			"selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType": selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType,
		}).Error("Unhandled 'selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType', should not happen")

		return

	}

	// Load Detailed TestSuiteExecution from Database
	testSuiteExecutionsModel.LoadDetailedTestSuiteExecutionFromDatabase(testSuiteExecutionUuid, testSuiteExecutionVersion)

	// Update TestSuite Preview
	TestSuiteInstructionPreViewObject.GenerateTestSuiteExecutionPreviewContainer(
		l.currentTestSuiteExecutionUuid,
		l.currentTestSuiteExecutionVersion,
		l.testSuiteExecutionsModel,
		fromExecutionList,
		sharedCode.FenixMasterWindowPtr)

	testSuiteExecutionsListTable.Refresh()

	// Update 'loadAllTestSuiteExecutionsForOneTestSuiteButton' with correct text, if we are in 'OneExecutionPerTestSuite'
	if selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType == OneExecutionPerTestSuite {
		loadAllTestSuiteExecutionsForOneTestSuiteButtonReference.
			Text = loadAllTestSuiteExecutionsForOneTestSuiteButtonTextPart1 +
			l.currentTestSuiteName
		loadAllTestSuiteExecutionsForOneTestSuiteButtonReference.Enable()
		loadAllTestSuiteExecutionsForOneTestSuiteButtonReference.Refresh()

	}

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
	currentRowThatMouseIsHoveringAboveMutex.Lock()

	// Set current TestSuiteUuid to be highlighted
	currentRowThatMouseIsHoveringAbove = l.currentRow

	// Release variable
	currentRowThatMouseIsHoveringAboveMutex.Unlock()

	l.TextStyle = fyne.TextStyle{Bold: true}
	l.Refresh()
	//testSuiteExecutionsListTable.Refresh()

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
	currentRowThatMouseIsHoveringAboveMutex.Lock()

	// Set current TestSuiteUuid to be highlighted
	currentRowThatMouseIsHoveringAbove = -1

	// Release variable
	currentRowThatMouseIsHoveringAboveMutex.Unlock()

	l.TextStyle = fyne.TextStyle{Bold: false}
	l.Refresh()
	//testSuiteExecutionsListTable.Refresh()

}
