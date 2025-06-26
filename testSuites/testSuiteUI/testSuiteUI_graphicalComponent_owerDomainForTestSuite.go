package testSuiteUI

import (
	"FenixTesterGui/testCase/testCaseModel"
	"FenixTesterGui/testSuites/testSuitesCommandEngine"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

// Generate the OwnerDomain Area for the TestSuite
func (testSuiteUiModel *TestSuiteUiStruct) generateOwnerDomainForTestSuiteArea(
	testCasesModel *testCaseModel.TestCasesModelsStruct) (
	testCaseOwnerDomainContainer *fyne.Container,
	testCaseOwnerDomainCustomSelectComboBox *customSelectComboBox,
	err error) {

	var tempCurrentOwnerDomainToBeChosenInDropDown string

	// If TestCase already has a chosen OwnerDomain then set that value
	var tempCurrentOwnerDomain string

	var testCaseHasOwnerDomain bool

	if len(testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestSuiteOwnerDomainUuid) > 0 {
		testCaseHasOwnerDomain = true
		tempCurrentOwnerDomain = testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestSuiteOwnerDomainUuid
	}

	// Load Domains that can own the TestSuite into options-array
	var options []string
	for _, tempDomainsThatCanOwnTheTestSuite := range testCasesModel.DomainsThatCanOwnTheTestCaseMap {
		options = append(options, tempDomainsThatCanOwnTheTestSuite.DomainNameShownInGui)

		// When TestCase has OwnerDomain find the one
		if testCaseHasOwnerDomain == true && tempDomainsThatCanOwnTheTestSuite.DomainUuid == tempCurrentOwnerDomain {
			tempCurrentOwnerDomainToBeChosenInDropDown = tempDomainsThatCanOwnTheTestSuite.DomainNameShownInGui
		}
	}

	// Create Form-layout container to be used for Name
	var testCaseOwnerDomainNameFormContainer *fyne.Container
	testCaseOwnerDomainContainer = container.New(layout.NewVBoxLayout())
	testCaseOwnerDomainNameFormContainer = container.New(layout.NewFormLayout())

	// Generate Warnings-rectangle for valid value, or that value exist
	//var valueIsValidWarningBox *canvas.Rectangle
	var colorToUse color.NRGBA
	var valueIsValidWarningBox *canvas.Rectangle
	var newOwnerDomainSelect *widget.Select

	colorToUse = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	valueIsValidWarningBox = canvas.NewRectangle(colorToUse)

	// Add the DropDown box with all domains that can own the TestCase
	newOwnerDomainSelect = widget.NewSelect(options,
		func(value string) {
			// This function is called when an option is selected.

			// Store Domain back to TestSuite-model
			testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestSuiteOwnerDomainUuid =
				testCasesModel.DomainsThatCanOwnTheTestCaseMap[value].DomainUuid
			testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestSuiteOwnerDomainName =
				testCasesModel.DomainsThatCanOwnTheTestCaseMap[value].DomainName

			// Set info in TestSuiteModel that OwnerDomain is set
			testSuiteUiModel.TestSuiteModelPtr.OwnerDomainHasValue(true)

			// Trigger creation of a 'new' TestEnvironment container for the TestSuite-UI ************************

			// Generate TestSuite's ExecutionEnvironment
			var newTestSuiteTestEnvironmentContainer *fyne.Container

			// Clear TestEnvironment, when OwnerDomain and Execution Environment hasn't been locked
			if testSuiteUiModel.TestSuiteModelPtr.HasLockButtonBeenClickedAndBothOwnerDomainAndTestEnvironmentHaveValues() == false {
				testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestSuiteExecutionEnvironment = ""

				// Clear info in TestSuiteModel about TestEnvironment is not selected
				testSuiteUiModel.TestSuiteModelPtr.TestEnvironmentHasValue(false)

				// Lock parts of UI that shouldn't be accessible before both OwnerDomain and TestEnvironment is Selected by user
				testSuiteUiModel.lockUIUntilOwnerDomainAndTestEnvironmenIsSelected()

			}

			newTestSuiteTestEnvironmentContainer,
				testSuiteUiModel.customTestEnvironmentSelectComboBox,
				err = testSuiteUiModel.
				generateTestEnvironmentForTestSuite()
			if err != nil {

				errorId := "533565fc-3851-4150-bf7c-364e8529f56c"
				errorMessage := fmt.Sprintf("couldn't generate 'TestSuites TestExecution environment-area', err=%s. [ErrorId = %s]",
					err.Error(),
					errorId)

				// Remove old 'testSuiteTestEnvironmentContainer' from stack container
				testSuiteUiModel.testSuiteTestEnvironmentStackContainer.
					Remove(testSuiteUiModel.testSuiteTestEnvironmentContainer)

				// Create new
				testSuiteUiModel.testSuiteTestEnvironmentContainer = container.NewVBox(widget.NewLabel(errorMessage))

				// Add new 'testSuiteTestEnvironmentContainer' to stack container
				testSuiteUiModel.testSuiteTestEnvironmentStackContainer.
					Add(testSuiteUiModel.testSuiteTestEnvironmentContainer)

				// Refresh Tabs
				testSuitesCommandEngine.TestSuiteTabsRef.Refresh()

				return

			}

			// Remove old 'testSuiteTestEnvironmentContainer' from stack container
			if testSuiteUiModel.testSuiteTestEnvironmentStackContainer != nil {
				testSuiteUiModel.testSuiteTestEnvironmentStackContainer.
					Remove(testSuiteUiModel.testSuiteTestEnvironmentContainer)
			}

			// Add new 'testSuiteTestEnvironmentContainer' to stack container
			if testSuiteUiModel.testSuiteTestEnvironmentStackContainer != nil {
				testSuiteUiModel.testSuiteTestEnvironmentStackContainer.Add(newTestSuiteTestEnvironmentContainer)
			} else {
				testSuiteUiModel.testSuiteTestEnvironmentStackContainer = container.
					NewStack(newTestSuiteTestEnvironmentContainer)

			}

			// Store 'newTestSuiteTestEnvironmentContainer' in old onec place
			testSuiteUiModel.testSuiteTestEnvironmentContainer = newTestSuiteTestEnvironmentContainer

			// Remove old 'testSuiteMetaDataContainer' from stack container
			testSuiteUiModel.testSuiteMetaDataStackContainer.Remove(testSuiteUiModel.testSuiteMetaDataContainer)

			// Create new
			testSuiteUiModel.testSuiteMetaDataContainer = container.NewVBox(widget.NewLabel("No OwnerDomain specified for TestSuite"))

			// Add new 'testSuiteMetaDataContainer' to stack container
			testSuiteUiModel.testSuiteMetaDataStackContainer.Add(testSuiteUiModel.testSuiteMetaDataContainer)

			// Refresh Tabs
			fyne.Do(func() {
				testSuitesCommandEngine.TestSuiteTabsRef.Refresh()
			})

			//var testCaseMetaDataArea fyne.CanvasObject
			/*
				var metaDataAccordion *widget.Accordion
				_, metaDataAccordion, err = testCasesUiCanvasObject.GenerateMetaDataAreaForTestCase(
					tempTestCasePtr,
					testCaseUuid,
					testCasesUiCanvasObject.TestCasesModelReference.DomainsThatCanOwnTheTestCaseMap[value].DomainUuid)

				if err != nil {
					log.Println(err, metaDataAccordion)

					return
				}



				metaDataAccordion.OpenAll()
			*/

			// Set Warning box that value is not selected
			if len(value) == 0 {
				valueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
			} else {
				valueIsValidWarningBox.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
			}
		})
	newOwnerDomainSelect.SetSelected(tempCurrentOwnerDomainToBeChosenInDropDown)

	// Create a custom SelectComboBox, with valueIsValidWarningBox
	var ownerDomainCustomSelectComboBox *customSelectComboBox
	ownerDomainCustomSelectComboBox = newCustomSelectComboBoxWidget(newOwnerDomainSelect, valueIsValidWarningBox)

	// Add Header to the Forms-container
	var headerLabel *widget.Label
	headerLabel = widget.NewLabel("Domain that 'Own' the TestSuite")
	headerLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCaseOwnerDomainNameFormContainer.Add(headerLabel)

	// Add the widget to the Forms-container
	testCaseOwnerDomainNameFormContainer.Add(ownerDomainCustomSelectComboBox)

	// Add Forms container to UI container
	testCaseOwnerDomainContainer.Add(testCaseOwnerDomainNameFormContainer)

	return testCaseOwnerDomainContainer,
		ownerDomainCustomSelectComboBox,
		err

}

// Sets the Selected value for the DropDown specifying the Owner-Domain of the TestSuite
func (testSuiteUiModel *TestSuiteUiStruct) setSelectedOwnerDomainForTestSuiteArea(
	tempCurrentOwnerDomainToBeChosenInDropDown string,
	newOwnerDomainSelect *widget.Select,
	valueIsValidWarningBox *canvas.Rectangle) {

	// Set the Visible value for DropDown, if there is any
	if len(tempCurrentOwnerDomainToBeChosenInDropDown) > 0 {
		newOwnerDomainSelect.SetSelected(tempCurrentOwnerDomainToBeChosenInDropDown)
	}

	// Set correct warning box color
	if len(newOwnerDomainSelect.Selected) == 0 {
		valueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	} else {
		valueIsValidWarningBox.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
	}
}

// calcSelectWidth returns the width needed to show the longest option
func calcSelectWidth(values []string) float32 {
	tmp := widget.NewSelect(values, nil)

	// Loop the values and check which has most characters
	var maxValue float32
	var indexForMaxValue int
	for valueIndex, value := range values {
		if maxValue < float32(len(value)) {
			maxValue = float32(len(value))
			indexForMaxValue = valueIndex
		}
	}

	tmp.PlaceHolder = values[indexForMaxValue] // ensure MinSize measures a non-empty string
	tmp.Refresh()                              // recalc cached size
	return tmp.MinSize().Width
}

// calcCheckGroupWidth returns the width needed to show the widest checkbox label
func calcCheckGroupWidth(values []string) float32 {
	tmp := widget.NewCheckGroup(values, nil)
	tmp.Refresh()
	return tmp.MinSize().Width
}
