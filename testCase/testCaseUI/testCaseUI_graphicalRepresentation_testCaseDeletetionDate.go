package testCaseUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/soundEngine"
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
	"image/color"
	"time"
)

var tickerCountDownlabel *widget.Label
var tickerCountDownlabelDataBinding binding.String
var enableDeletionCheckbox *widget.Check
var tickerDoneChannel chan bool
var newTestCaseDeletionDateEntry *widget.Entry

func countDownTicker() {
	tickerDoneChannel = make(chan bool)
	countdown := 10 // Start from 5 seconds
	tickerCountDownlabelDataBinding.Set(fmt.Sprintf("Countdown: %d seconds remaining", countdown))
	tickerCountDownlabel.Show()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop() // Stop the ticker when we're done

	go func() {

		for {
			select {
			case <-tickerDoneChannel:
				return
			case <-ticker.C:
				tickerCountDownlabelDataBinding.Set(fmt.Sprintf("Countdown: %d seconds remaining", countdown))

				if countdown <= 0 {
					tickerDoneChannel <- true
				}
				countdown--
			}
		}
	}()

	// Wait for the countdown to finish
	<-tickerDoneChannel
	tickerCountDownlabel.Hide()
	enableDeletionCheckbox.SetChecked(false)

}

// Functions that hides the Fenix Screen and the flash the full screen
func flashScreen(mainApp fyne.App, mainWindow fyne.Window) {
	// Hide the main window
	mainWindow.Hide()

	// Create a new window for the red screen
	redWindow := mainApp.NewWindow("Red Screen")

	// Set the window to full-screen mode
	redWindow.SetFullScreen(true)

	// Create a red rectangle
	redRect := canvas.NewRectangle(color.NRGBA{R: 255, G: 0, B: 0, A: 255})

	// Use a Max layout to ensure the rectangle fills the window
	content := fyne.NewContainerWithLayout(
		layout.NewMaxLayout(),
		redRect,
	)

	// Set the content of the window
	redWindow.SetContent(content)

	// Show the red window
	redWindow.Show()

	tickerDoneChannel := make(chan bool)
	countdown := 10 // Start from 5 seconds

	var isRed bool
	isRed = true

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop() // Stop the ticker when we're done

	go func() {

		for {
			select {
			case <-tickerDoneChannel:
				return
			case <-ticker.C:
				if isRed == true {
					fmt.Println("isRed == true")
					redRect.FillColor = color.NRGBA{R: 255, G: 255, B: 0, A: 255}
					redRect.Refresh()
					isRed = false
				} else {
					fmt.Println("isRed == false")
					redRect.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
					redRect.Refresh()
					isRed = true
				}

				if countdown <= 0 {
					tickerDoneChannel <- true
				}
				countdown--
			}
		}
	}()

	// Wait for the countdown to finish
	<-tickerDoneChannel
	redWindow.Close()
	mainWindow.Show()

	// Notify That testCase is deleted per today

	// Trigger System Notification sound
	soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   "TestCase Deleted",
		Content: "The TestCase was set to Deleted in the Database!",
	})

}

// Generate the TestCaseDeletionDate Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateTestCaseDeletionDateArea(
	testCaseUuid string) (
	testCaseDeletionDateArea fyne.CanvasObject,
	err error) {

	var existsInMap bool

	// Get TestCasesMap
	var testCasesMap map[string]*testCaseModel.TestCaseModelStruct
	testCasesMap = *testCasesUiCanvasObject.TestCasesModelReference.TestCasesMapPtr

	// Get current TestCase
	var currentTestCasePtr *testCaseModel.TestCaseModelStruct
	currentTestCasePtr, existsInMap = testCasesMap[testCaseUuid]

	if existsInMap == false {
		errorId := "5c27a6eb-21b2-4719-ab26-04c43cb70f5a"
		err := errors.New(fmt.Sprintf("testcase-model with TestCaseUuid '%s' is missing map for TestCasesMapPtr [ErrorID: %s]", testCaseUuid, errorId))

		//TODO Send ERRORS over error-channel
		fmt.Println(err)

		return nil, err

	}

	var deleteTestCaseButton *widget.Button

	// Create Form-layout container to be used for DeletionDate
	var testCaseDeletionDateContainer *fyne.Container
	var testCaseDeletionDateFormContainer *fyne.Container
	testCaseDeletionDateContainer = container.New(layout.NewVBoxLayout())
	testCaseDeletionDateFormContainer = container.New(layout.NewFormLayout())

	// Add Ticker count down label
	tickerCountDownlabelDataBinding = binding.NewString()
	tickerCountDownlabel = widget.NewLabelWithData(tickerCountDownlabelDataBinding)
	tickerCountDownlabel.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	tickerCountDownlabel.Hide()

	// Add Header to the Forms-container
	var headerLabel *widget.Label
	headerLabel = widget.NewLabel("TestCase will be deleted by this date")

	headerLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCaseDeletionDateFormContainer.Add(headerLabel)

	// Generate Warnings-rectangle for valid value, or that value exist
	var valueIsValidWarningBox *canvas.Rectangle
	var colorToUse color.NRGBA
	colorToUse = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	valueIsValidWarningBox = canvas.NewRectangle(colorToUse)

	// Date validator function for Entry
	dateValidatorFunction := func(dateToValidate string) bool {
		var parseError error

		var validTodayDate string
		validTodayDate = time.Now().Format("2006-01-02")

		_, parseError = time.Parse("2006-01-02", dateToValidate)
		if parseError != nil {
			return false

		} else {

			// Date must be equal or bigger then Today()
			if dateToValidate >= validTodayDate {
				return true // Valid date
			} else {
				return false
			}

		}
	}

	// Add the Entry-widget for testCaseDeletionDate
	newTestCaseDeletionDateEntry = widget.NewEntry()
	newTestCaseDeletionDateEntry.SetPlaceHolder("Date when TestCase should be removed: YYYY-MM-DD")
	//newTestCaseDeletionDateEntry.Disable()
	newTestCaseDeletionDateEntry.SetText(currentTestCasePtr.LocalTestCaseMessage.DeleteTimeStamp)
	newTestCaseDeletionDateEntry.OnChanged = func(s string) {
		var dateIsValid bool

		dateIsValid = dateValidatorFunction(s)
		// Set Warning box that value is not selected
		if dateIsValid == false {
			valueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
			enableDeletionCheckbox.Disable()
			enableDeletionCheckbox.SetChecked(false)
			deleteTestCaseButton.Disable()

		} else {
			valueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 255, B: 0, A: 255}
			enableDeletionCheckbox.Enable()

		}

	}

	valueIsValidWarningBox.SetMinSize(fyne.NewSize(15, newTestCaseDeletionDateEntry.Size().Height))

	// The button that activates the Deletion of the TestCase
	deleteTestCaseButton = widget.NewButton("Auto-Delete TestCase at specified date", func() {
		fmt.Println("DELETE")
		tickerDoneChannel <- true

		var existInMap bool

		currentTestCasePtr, existsInMap = testCasesMap[testCaseUuid]

		if existInMap == false {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":           "879e46d2-4439-404b-ac6d-d99a3307b6f6",
				"testCaseUuid": testCaseUuid,
			}).Fatal("TestCase doesn't exist in TestCaseMap. This should not happen")
		}

		// Which type of Delete should be performed?
		//var dateIsInTheFuture bool

		// This TestCase is not saved in Database
		//if currentTestCasePtr.ThisIsANewTestCase == true {

		// Check if The date is Today() or in the future
		var parseError error

		//var validTodayDate string
		//validTodayDate = time.Now().Format("2006-01-02")

		_, parseError = time.Parse("2006-01-02", newTestCaseDeletionDateEntry.Text)
		if parseError != nil {

			newTestCaseDeletionDateEntry.SetText("")
			enableDeletionCheckbox.Disable()

			return

		} else {

			/*
				// Date must be equal or bigger then Today()
				if newTestCaseDeletionDateEntry.Text > validTodayDate {
					dateIsInTheFuture = true
				} else {
					dateIsInTheFuture = false
				}

			*/

		}

		//if dateIsInTheFuture == false {

		// Save the Delete date in the local version of the TestCase
		currentTestCasePtr.LocalTestCaseMessage.DeleteTimeStamp = newTestCaseDeletionDateEntry.Text

		// Save back the updated TestCase
		//testCasesUiCanvasObject.TestCasesModelReference.TestCasesMapPtr[testCaseUuid] = currentTestCasePtr

		// Remove TestCase from TestCase model and the UI-model
		commandEngineChannelMessage := sharedCode.ChannelCommandStruct{
			ChannelCommand:  sharedCode.ChannelCommandRemoveTestCaseWithOutSaving,
			FirstParameter:  testCaseUuid,
			SecondParameter: "",
			ActiveTestCase:  "",
			ElementType:     sharedCode.Undefined,
		}

		// Send command message over channel to Command and Rule Engine
		*testCasesUiCanvasObject.CommandChannelReference <- commandEngineChannelMessage

		//}
		//}

		// This TestCase is saved in Database and Delete date is Today()

		// This TestCase is saved in Database and Delete date is in the future

	})
	deleteTestCaseButton.Disable()

	enableDeletionCheckbox = widget.NewCheck("Enable Deletion of TestCase", func(b bool) {

		// Switch button for the actual deletion
		if b == true {
			deleteTestCaseButton.Enable()
			go countDownTicker()
		} else {
			deleteTestCaseButton.Disable()
		}
	})
	enableDeletionCheckbox.Disable()

	// Set correct warning box color from the beginning
	var dateIsValid bool
	dateIsValid = dateValidatorFunction(newTestCaseDeletionDateEntry.Text)
	// Set Warning box that value is not selected
	if dateIsValid == false {
		valueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
		enableDeletionCheckbox.Disable()
		enableDeletionCheckbox.SetChecked(false)
		deleteTestCaseButton.Disable()
	} else {
		valueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 255, B: 0, A: 255}
		enableDeletionCheckbox.Enable()
	}

	// Deletion logic container
	var deletionLogicContainer *fyne.Container
	deletionLogicContainer = container.NewHBox(enableDeletionCheckbox, deleteTestCaseButton)

	var entryContainer *fyne.Container
	entryContainer = container.NewBorder(nil, nil, valueIsValidWarningBox, deletionLogicContainer, newTestCaseDeletionDateEntry)

	// Add the Entry-widget to the Forms-container
	testCaseDeletionDateFormContainer.Add(entryContainer)

	// Create the VBox-container that will be returned
	testCaseDeletionDateContainer = container.NewVBox(tickerCountDownlabel, testCaseDeletionDateFormContainer)

	// Create an Accordion item for the list
	testCaseDeletionAccordionItem := widget.NewAccordionItem("Delete TestCase", testCaseDeletionDateContainer)

	testCaseDeletionAccordion := widget.NewAccordion(testCaseDeletionAccordionItem)

	// Create the VBox-container that will be returned
	testCaseDeletionArea := container.NewVBox(testCaseDeletionAccordion, widget.NewSeparator())

	return testCaseDeletionArea, err
}
