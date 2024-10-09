package testCaseUI

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

// Generate the TestCaseDeletionDate Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateTestCaseDeletionDateArea(
	testCaseUuid string) (
	testCaseDeletionDateArea fyne.CanvasObject,
	err error) {

	/*
		// Extract the current TestCase UI model
		testCase_Model, existsInMap := testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]
		if existsInMap == false {
			errorId := "e98a46a6-ddb8-4f00-b68b-74d9132863e6"
			err := errors.New(fmt.Sprintf("testcase-model with TestCaseUuid '%s' is missing map for TestCases [ErrorID: %s]", testCaseUuid, errorId))

			//TODO Send ERRORS over error-channel
			fmt.Println(err)

			return nil, err

		}

	*/

	var enableDeletionCheckbox *widget.Check
	var deleteTestCaseButton *widget.Button

	// Create Form-layout container to be used for DeletionDate
	var testCaseDeletionDateContainer *fyne.Container
	var testCaseDeletionDateFormContainer *fyne.Container
	testCaseDeletionDateContainer = container.New(layout.NewVBoxLayout())
	testCaseDeletionDateFormContainer = container.New(layout.NewFormLayout())

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
	var newTestCaseDeletionDateEntry *widget.Entry
	newTestCaseDeletionDateEntry = widget.NewEntry()
	newTestCaseDeletionDateEntry.SetPlaceHolder("Enter Date when TestCase should be removed, use format: YYYY-MM-DD")
	//newTestCaseDeletionDateEntry.Disable()
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
	})
	deleteTestCaseButton.Disable()

	enableDeletionCheckbox = widget.NewCheck("Enable Deletion of TestCase", func(b bool) {

		// Switch button for the actual deletion
		if b == true {
			deleteTestCaseButton.Enable()
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
	testCaseDeletionDateContainer = container.NewVBox(testCaseDeletionDateFormContainer)

	// Create an Accordion item for the list
	testCaseDeletionAccordionItem := widget.NewAccordionItem("Delete TestCase", testCaseDeletionDateContainer)

	testCaseDeletionAccordion := widget.NewAccordion(testCaseDeletionAccordionItem)

	// Create the VBox-container that will be returned
	testCaseDeletionArea := container.NewVBox(testCaseDeletionAccordion, widget.NewSeparator())

	return testCaseDeletionArea, err
}
