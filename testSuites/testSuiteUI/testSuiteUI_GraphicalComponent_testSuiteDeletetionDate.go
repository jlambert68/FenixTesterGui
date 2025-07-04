package testSuiteUI

import (
	"FenixTesterGui/soundEngine"
	"FenixTesterGui/testSuites/testSuitesModel"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

var tickerCountDownlabel *widget.Label
var tickerCountDownlabelDataBinding binding.String
var enableDeletionCheckbox *widget.Check
var tickerDoneChannel chan bool
var newTestSuiteDeletionDateEntry *widget.Entry

func countDownTicker(testSuiteUiModel *TestSuiteUiStruct) {
	tickerDoneChannel = make(chan bool)
	countdown := 10 // Start from 5 seconds
	tickerCountDownlabelDataBinding.Set(fmt.Sprintf("Countdown: %d seconds remaining", countdown))
	fyne.Do(func() {
		tickerCountDownlabel.Show()
	})

	var tickerChannelValue bool

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop() // Stop the ticker when we're done

	go func() {

		for {
			select {
			case <-tickerDoneChannel:
				return
			case <-ticker.C:
				tickerCountDownlabelDataBinding.Set(fmt.
					Sprintf("Countdown: %d seconds remaining", countdown))

				if countdown <= 0 {
					tickerDoneChannel <- true
				}
				countdown--
			}
		}
	}()

	// Wait for the countdown to finish
	tickerChannelValue = <-tickerDoneChannel
	fyne.Do(func() {
		tickerCountDownlabel.Hide()
	})
	fyne.Do(func() {
		enableDeletionCheckbox.SetChecked(false)
	})

	// If The user ended the Ticker with the Delete-button then accept the value else clear text box
	if tickerChannelValue == true {
		fyne.Do(func() {
			newTestSuiteDeletionDateEntry.SetText("")
			testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestSuiteDeletionDate = ""

			// Trigger System Notification sound
			soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "TestSuite Deleted",
				Content: "Deleted Date is cleared in the Database!",
			})
		})
	}

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

	// Notify That TestSuite is deleted per today

	// Trigger System Notification sound
	soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   "TestSuite Deleted",
		Content: "The TestSuite was set to Deleted in the Database!",
	})

}

// Generate the TestSuiteDeletionDate Area for the TestSuite
func (testSuiteUiModel *TestSuiteUiStruct) generateTestSuiteDeletionDateArea(
	testSuiteUuid string) (
	testSuiteDeletionDateAreaContainer *fyne.Container,
	err error) {

	// Verify that there is a TestSuiteModel
	if testSuiteUiModel.TestSuiteModelPtr == nil {

		errorId := "d2338d89-979b-44a4-b3fb-21fbada2dee9"
		err = errors.New(fmt.Sprintf("TestSuiteModelPtr is nil. TestSuiteUuid=%s [ErrorID: %s]",
			testSuiteUuid,
			errorId))

		//TODO Send ERRORS over error-channel
		fmt.Println(err)

		return nil, err

	}

	// Get current TestSuiteModel
	var currentTestSuiteModel testSuitesModel.TestSuiteModelStruct
	currentTestSuiteModel = *testSuiteUiModel.TestSuiteModelPtr

	var deleteTestSuiteButton *widget.Button

	// Create Form-layout container to be used for DeletionDate
	var testSuiteDeletionDateContainer *fyne.Container
	var testSuiteDeletionDateFormContainer *fyne.Container
	testSuiteDeletionDateContainer = container.New(layout.NewVBoxLayout())
	testSuiteDeletionDateFormContainer = container.New(layout.NewFormLayout())

	// Add Ticker count down label
	tickerCountDownlabelDataBinding = binding.NewString()
	tickerCountDownlabel = widget.NewLabelWithData(tickerCountDownlabelDataBinding)
	tickerCountDownlabel.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	tickerCountDownlabel.Hide()

	// Add Header to the Forms-container
	var headerLabel *widget.Label
	headerLabel = widget.NewLabel("TestSuite will be deleted by this date")

	headerLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuiteDeletionDateFormContainer.Add(headerLabel)

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
			if dateToValidate > validTodayDate {
				return true // Valid date
			} else {
				return false
			}

		}
	}

	// Add the Entry-widget for testSuiteDeletionDate
	newTestSuiteDeletionDateEntry = widget.NewEntry()

	newTestSuiteDeletionDateEntry.SetPlaceHolder("Date when TestSuite should be removed: YYYY-MM-DD")
	//newTestSuiteDeletionDateEntry.Disable()
	newTestSuiteDeletionDateEntry.SetText(currentTestSuiteModel.TestSuiteUIModelBinding.TestSuiteDeletionDate)
	newTestSuiteDeletionDateEntry.OnChanged = func(s string) {
		var dateIsValid bool

		dateIsValid = dateValidatorFunction(s)
		// Set Warning box that value is not selected
		if dateIsValid == false {
			if len(newTestSuiteDeletionDateEntry.Text) == 0 {
				valueIsValidWarningBox.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
			} else {
				valueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
			}

			enableDeletionCheckbox.Disable()
			enableDeletionCheckbox.SetChecked(false)
			deleteTestSuiteButton.Disable()

		} else {
			valueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 255, B: 0, A: 255}
			enableDeletionCheckbox.Enable()

		}

		// Check if values has changed, if so the trigger 'ticker count down'
		if dateIsValid == true && s != currentTestSuiteModel.TestSuiteUIModelBinding.TestSuiteDeletionDate &&
			time.Now().Format("2006-01-02") == s {

			enableDeletionCheckbox.SetChecked(true)

			// Trigger System Notification sound
			soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Delete TestSuite",
				Content: fmt.Sprintf("You must accept the new Delete date"),
			})

		}

	}

	valueIsValidWarningBox.SetMinSize(fyne.NewSize(15, newTestSuiteDeletionDateEntry.Size().Height))

	// The button that activates the Deletion of the TestSuite
	deleteTestSuiteButton = widget.NewButton("Auto-Delete TestSuite at specified date", func() {
		fmt.Println("DELETE")
		tickerDoneChannel <- false

		// Which type of Delete should be performed?
		//var dateIsInTheFuture bool

		// This TestSuite is not saved in Database
		//if currentTestSuitePtr.ThisIsANewTestSuite == true {

		// Check if The date is Today() or in the future
		var parseError error

		//var validTodayDate string
		//validTodayDate = time.Now().Format("2006-01-02")

		_, parseError = time.Parse("2006-01-02", newTestSuiteDeletionDateEntry.Text)
		if parseError != nil {

			newTestSuiteDeletionDateEntry.SetText("")
			enableDeletionCheckbox.Disable()

			return

		} else {

			/*
				// Date must be equal or bigger then Today()
				if newTestSuiteDeletionDateEntry.Text > validTodayDate {
					dateIsInTheFuture = true
				} else {
					dateIsInTheFuture = false
				}

			*/

		}

		// Change color to Green for the color-box
		valueIsValidWarningBox.FillColor = color.NRGBA{R: 0, G: 255, B: 0, A: 255}

		//if dateIsInTheFuture == false {

		// Store the Delete date in the TestSuiteModel
		testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestSuiteDeletionDate = newTestSuiteDeletionDateEntry.Text

		// If TestSuite hase been saved then Trigger 'DeleteTestSuiteAtThisDate'
		// Otherwise normal save takes care of saving the Delete-date

		if testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestSuiteIsNew == false {

			// Send Delete to Database
			err = testSuiteUiModel.TestSuiteModelPtr.DeleteTestSuiteAtThisDate(
				testSuiteUiModel.TestSuiteModelPtr.GetTestSuiteUuid())

			if err != nil {

				// Trigger System Notification sound
				soundEngine.PlaySoundChannel <- soundEngine.InvalidNotificationSound

				fyne.CurrentApp().SendNotification(&fyne.Notification{
					Title:   "Delete TestSuite",
					Content: fmt.Sprintf("Got some error when trying to Delete TestSuite"),
				})

				return

			} else {
				// Trigger System Notification sound
				soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

				fyne.CurrentApp().SendNotification(&fyne.Notification{
					Title:   "Delete TestSuite",
					Content: fmt.Sprintf("Delete Date was successfuly updated for TestSuite"),
				})
			}
		}

		// This TestSuite is saved in Database and Delete date is Today()

		// This TestSuite is saved in Database and Delete date is in the future

	})
	deleteTestSuiteButton.Disable()

	enableDeletionCheckbox = widget.NewCheck("Enable Deletion of TestSuite", func(b bool) {

		// Switch button for the actual deletion
		if b == true {
			deleteTestSuiteButton.Enable()
			fyne.Do(func() {
				go countDownTicker(testSuiteUiModel)
			})
		} else {
			deleteTestSuiteButton.Disable()
		}
	})
	enableDeletionCheckbox.Disable()

	// Set correct warning box color from the beginning
	var dateIsValid bool
	dateIsValid = dateValidatorFunction(newTestSuiteDeletionDateEntry.Text)
	// Set Warning box that value is not selected
	if dateIsValid == false {
		if len(newTestSuiteDeletionDateEntry.Text) == 0 {
			valueIsValidWarningBox.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
		} else {
			valueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
		}

		enableDeletionCheckbox.Disable()
		enableDeletionCheckbox.SetChecked(false)
		deleteTestSuiteButton.Disable()
	} else {
		valueIsValidWarningBox.FillColor = color.NRGBA{R: 0, G: 255, B: 0, A: 255}
		enableDeletionCheckbox.Enable()
	}

	// Deletion logic container
	var deletionLogicContainer *fyne.Container
	deletionLogicContainer = container.NewHBox(enableDeletionCheckbox, deleteTestSuiteButton)

	var entryContainer *fyne.Container
	entryContainer = container.NewBorder(nil, nil, valueIsValidWarningBox, deletionLogicContainer, newTestSuiteDeletionDateEntry)

	// Add the Entry-widget to the Forms-container
	testSuiteDeletionDateFormContainer.Add(entryContainer)

	// Create the VBox-container that will be returned
	testSuiteDeletionDateContainer = container.NewVBox(tickerCountDownlabel, testSuiteDeletionDateFormContainer)

	// Create an Accordion item for the list
	testSuiteDeletionAccordionItem := widget.NewAccordionItem("Delete TestSuite", testSuiteDeletionDateContainer)

	testSuiteDeletionAccordion := widget.NewAccordion(testSuiteDeletionAccordionItem)

	// Create the VBox-container that will be returned
	testSuiteDeletionDateAreaContainer = container.NewVBox(testSuiteDeletionAccordion, widget.NewSeparator())

	return testSuiteDeletionDateAreaContainer, err
}
