package testCaseUI

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Generate the OwnerDomain Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateOwnerDomainForTestCaseArea(
	testCaseUuid string) (
	ownerDomainArea fyne.CanvasObject,
	err error) {

	// Extract the current TestCase UI model
	testCase_Model, existsInMap := testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]
	if existsInMap == false {
		errorId := "bb7fe228-2079-481f-89d3-8cf07a4da26a"
		err := errors.New(fmt.Sprintf("testcase-model with TestCaseUuid '%s' is missing map for TestCases [ErrorID: %s]", testCaseUuid, errorId))

		//TODO Send ERRORS over error-channel
		fmt.Println(err)

		return nil, err

	}

	// If TestCase already has a chosen OwnerDomain then set that value
	var tempCurrentOwnerDomain string
	var tempCurrentOwnerDomainToBeChosenInDropDown string
	var testCaseHasOwnerDomain bool

	if len(testCase_Model.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.DomainUuid) > 0 {
		testCaseHasOwnerDomain = true
		tempCurrentOwnerDomain = testCase_Model.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.DomainUuid
	}

	// Load Domains that can own the TestCase into options-array
	var options []string
	for _, tempDomainsThatCanOwnTheTestCase := range testCasesUiCanvasObject.TestCasesModelReference.DomainsThatCanOwnTheTestCaseMap {
		options = append(options, tempDomainsThatCanOwnTheTestCase.DomainNameShownInGui)

		// When TestCase has OwnerDomain find the one
		if testCaseHasOwnerDomain == true && tempDomainsThatCanOwnTheTestCase.DomainUuid == tempCurrentOwnerDomain {
			tempCurrentOwnerDomainToBeChosenInDropDown = tempDomainsThatCanOwnTheTestCase.DomainNameShownInGui
		}
	}

	// Create Form-layout container to be used for Name
	var testCaseOwnerDomainontainer *fyne.Container
	var testCaseOwnerDomainNameFormContainer *fyne.Container
	testCaseOwnerDomainontainer = container.New(layout.NewVBoxLayout())
	testCaseOwnerDomainNameFormContainer = container.New(layout.NewFormLayout())

	// Add Header to the Forms-container
	testCaseOwnerDomainNameFormContainer.Add(widget.NewLabel("Domain that 'Own' the TestCase"))

	// Add the DropDown box with all domains that can own the TestCase
	newOwnerDomainDropDown := widget.NewSelect(options,
		func(value string) {
			// This function is called when an option is selected.
			// You can handle the selection here.
			fmt.Println("Selected:", value)

			// Save TestCase back in Map
			// Get the latest version of TestCase
			tempTestCase, _ := testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]

			// Store Domain in LocalTestCase in TestCase-model
			tempTestCase.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.DomainUuid =
				testCasesUiCanvasObject.TestCasesModelReference.DomainsThatCanOwnTheTestCaseMap[value].DomainUuid
			tempTestCase.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.DomainName =
				testCasesUiCanvasObject.TestCasesModelReference.DomainsThatCanOwnTheTestCaseMap[value].DomainName

			// Store back TestCase-model in Map
			testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid] = tempTestCase
		})

	// Set the Visible value for DropDown, if there is any
	if len(tempCurrentOwnerDomainToBeChosenInDropDown) > 0 {
		newOwnerDomainDropDown.SetSelected(tempCurrentOwnerDomainToBeChosenInDropDown)
	}

	// Add the Entry-widget to the Forms-container
	testCaseOwnerDomainNameFormContainer.Add(newOwnerDomainDropDown)

	// Create the VBox-container that will be returned
	testCaseOwnerDomainontainer = container.NewVBox(testCaseOwnerDomainNameFormContainer)

	return testCaseOwnerDomainontainer, err
}
