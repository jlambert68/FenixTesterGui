package testCaseUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testCase/testCaseModel"
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

	// Create Form-layout container to be used for Name
	var testCaseNameContainer *fyne.Container
	var testCaseNameFormContainer *fyne.Container
	testCaseNameContainer = container.New(layout.NewVBoxLayout())
	testCaseNameFormContainer = container.New(layout.NewFormLayout())

	// Add Header to the Forms-container
	var headerLabel *widget.Label
	headerLabel = widget.NewLabel("TestCaseName")
	headerLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCaseNameFormContainer.Add(headerLabel)

	// Add the Entry-widget for TestCaseName
	newTestCaseNameEntry := widget.NewEntry()
	newTestCaseNameEntry.SetText(currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageEditableInformation.TestCaseName)

	// Change Name in model and Tab-name when UI-component-Name-value is changed
	newTestCaseNameEntry.OnChanged = func(newValue string) {

		var trimmedValue string
		trimmedValue = strings.Trim(newValue, " ")

		// Save TestCase back in Map
		// Get the latest version of TestCase
		tempTestCasePtr, _ := testCasesMap[testCaseUuid]

		tempTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageEditableInformation.TestCaseName = trimmedValue
		//testCasesUiCanvasObject.TestCasesModelReference.TestCasesMapPtr[testCaseUuid] = tempTestCasePtr

		// Generate short version of UUID to put in TestCase Tab-Name
		var shortUUid string
		var tabName string

		shortUUid = testCasesUiCanvasObject.TestCasesModelReference.GenerateShortUuidFromFullUuid(testCaseUuid)

		// Shorten Tab-name if name is longer then 'TestCaseTabNameVisibleLength'
		if len(trimmedValue) > sharedCode.TestCaseTabNameVisibleLength {
			tabName = trimmedValue[0:sharedCode.TestCaseTabNameVisibleLength] + " [" + shortUUid + "] (*)"
		} else {
			tabName = trimmedValue + " [" + shortUUid + "] (*)"
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
