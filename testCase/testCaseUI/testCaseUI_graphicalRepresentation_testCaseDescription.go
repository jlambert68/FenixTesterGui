package testCaseUI

import (
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strings"
)

// Generate the TestCaseDescription Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateTestCaseDescriptionArea(
	testCaseUuid string) (
	testCaseDescriptionArea fyne.CanvasObject,
	err error) {

	var existsInMap bool

	// Get TestCasesMap
	var testCasesMap map[string]*testCaseModel.TestCaseModelStruct
	testCasesMap = *testCasesUiCanvasObject.TestCasesModelReference.TestCasesMapPtr

	// Get current TestCase
	var currentTestCasePtr *testCaseModel.TestCaseModelStruct
	currentTestCasePtr, existsInMap = testCasesMap[testCaseUuid]

	if existsInMap == false {
		errorId := "57d439e2-adf2-4abe-a0d0-f4afb98dd0a6"
		err := errors.New(fmt.Sprintf("testcase-model with TestCaseUuid '%s' is missing map for TestCasesMapPtr [ErrorID: %s]", testCaseUuid, errorId))

		//TODO Send ERRORS over error-channel
		fmt.Println(err)

		return nil, err

	}

	// Create Form-layout container to be used for Description
	var testCaseDescriptionContainer *fyne.Container
	var testCaseDescriptionFormContainer *fyne.Container
	testCaseDescriptionContainer = container.New(layout.NewVBoxLayout())
	testCaseDescriptionFormContainer = container.New(layout.NewFormLayout())

	// Add Header to the Forms-container
	var headerLabel *widget.Label
	headerLabel = widget.NewLabel("TestCaseDescription")
	headerLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCaseDescriptionFormContainer.Add(headerLabel)

	// Add the Entry-widget for TestCaseDescription
	newTestCaseDescriptionEntry := widget.NewMultiLineEntry()
	newTestCaseDescriptionEntry.SetText(currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageEditableInformation.TestCaseDescription)

	// Change vertical Size of Entry-widget for TestCaseDescription
	newTestCaseDescriptionEntry.SetMinRowsVisible(5)

	// Change Description in model when UI-component-Description-value is changed
	newTestCaseDescriptionEntry.OnChanged = func(newValue string) {

		var trimmedValue string
		trimmedValue = strings.Trim(newValue, " ")

		// Save TestCase back in Map
		// Get the latest version of TestCase
		tempTestCase, _ := testCasesMap[testCaseUuid]

		tempTestCase.LocalTestCaseMessage.BasicTestCaseInformationMessageEditableInformation.TestCaseDescription = trimmedValue
		//testCasesUiCanvasObject.TestCasesModelReference.TestCasesMapPtr[testCaseUuid] = tempTestCase

	}

	// Add the Entry-widget to the Forms-container
	testCaseDescriptionFormContainer.Add(newTestCaseDescriptionEntry)

	// Create the VBox-container that will be returned
	testCaseDescriptionContainer = container.NewVBox(testCaseDescriptionFormContainer)

	// Create and add Slider to control vertical size of the vertical size of Entry-widget for TestCaseDescription
	var verticalSizeSlider *widget.Slider
	verticalSizeSlider = &widget.Slider{Value: 5, Step: 1, Min: 3, Max: 30, OnChanged: func(f float64) {
		newTestCaseDescriptionEntry.SetMinRowsVisible(int(verticalSizeSlider.Value))
	}}

	testCaseDescriptionContainer.Add(verticalSizeSlider)

	return testCaseDescriptionContainer, err
}
