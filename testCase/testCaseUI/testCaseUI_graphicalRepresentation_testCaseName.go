package testCaseUI

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strings"
)

// Generate the TestCaseName Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateTestCaseNameArea(
	testCaseUuid string) (
	testCaseNameArea fyne.CanvasObject,
	err error) {

	// Extract the current TestCase UI model
	testCase_Model, existsInMap := testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]
	if existsInMap == false {
		errorId := "57d439e2-adf2-4abe-a0d0-f4afb98dd0a6"
		err := errors.New(fmt.Sprintf("testcase-model with TestCaseUuid '%s' is missing map for TestCases [ErrorID: %s]", testCaseUuid, errorId))

		//TODO Send ERRORS over error-channel
		fmt.Println(err)

		return nil, err

	}

	// Extract the current TestCase UI model
	//testCaseUIModel, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]

	// Create Form-layout container to be used for Name
	var testCaseNameContainer *fyne.Container
	var testCaseNameFormContainer *fyne.Container
	testCaseNameContainer = container.New(layout.NewVBoxLayout())
	testCaseNameFormContainer = container.New(layout.NewFormLayout())

	// Add Header to the Forms-container
	testCaseNameFormContainer.Add(widget.NewLabel("TestCaseName"))

	// Add the Entry-widget for TestCaseName
	newTestCaseNameEntry := widget.NewEntry() // testCasesUiCanvasObject.NewAttributeEntry(attributeItem.attributeUuid)
	newTestCaseNameEntry.SetText(testCase_Model.LocalTestCaseMessage.BasicTestCaseInformationMessageEditableInformation.TestCaseName)

	// Change Name in model and Tab-name when UI-component-Name-value is changed
	newTestCaseNameEntry.OnChanged = func(newValue string) {

		var trimmedValue string
		trimmedValue = strings.Trim(newValue, " ")

		testCase_Model.LocalTestCaseMessage.BasicTestCaseInformationMessageEditableInformation.TestCaseName = trimmedValue

		// Save TestCase back in Map
		//testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid] = testCase_Model

		// Generate short version of UUID to put in TestCase Tab-Name
		var shortUUid string
		var tabName string

		shortUUid = testCasesUiCanvasObject.TestCasesModelReference.GenerateShortUuidFromFullUuid(testCaseUuid)

		// Shorten Tab-name if name is longer then 'testCaseTabNameVisibleLenght'
		if len(trimmedValue) > testCaseTabNameVisibleLenght {
			tabName = trimmedValue[0:testCaseTabNameVisibleLenght] + " [" + shortUUid + "]"
		} else {
			tabName = trimmedValue + " [" + shortUUid + "]"
		}

		testCasesUiCanvasObject.TestCasesTabs.Selected().Text = tabName
		testCasesUiCanvasObject.TestCasesTabs.Refresh()

	}

	// Add the Entry-widget to the Forms-container
	testCaseNameFormContainer.Add(newTestCaseNameEntry)

	// Create the VBox-container that will be returned
	testCaseNameContainer = container.NewVBox(testCaseNameFormContainer)

	return testCaseNameContainer, err
}
