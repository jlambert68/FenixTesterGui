package testSuiteUI

import (
	"FenixTesterGui/testCase/testCaseModel"
	"FenixTesterGui/testSuites/testSuitesModel"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
)

// Generate the OwnerDomain Area for the TestSuite
func (testSuiteUiModel *TestSuiteUiStruct) generateOwnerDomainForTestSuiteArea(
	testSuiteUuid string,
	testCasesModel *testCaseModel.TestCasesModelsStruct) (
	ownerDomainArea fyne.CanvasObject,
	tempCurrentOwnerDomainToBeChosenInDropDown string,
	newOwnerDomainSelect *widget.Select,
	valueIsValidWarningBox *canvas.Rectangle,
	err error) {

	var existsInMap bool

	// Get testSuitesMap
	var testSuitesMap map[string]*testSuitesModel.TestSuiteModelStruct
	testSuitesMap = *testSuitesModel.TestSuitesModelPtr.TestSuitesMapPtr

	// Get a pointer to the TestSuite-model and the TestSuite-model itself
	var currentTestSuiteModelPtr *testSuitesModel.TestSuiteModelStruct
	var currentTestSuiteModel testSuitesModel.TestSuiteModelStruct
	currentTestSuiteModelPtr, existsInMap = testSuitesMap[testSuiteUuid]

	if existsInMap == false {
		errorId := "47f70e15-17dc-4dac-86b1-8545de829461"
		err := errors.New(fmt.Sprintf("TestSuite, %s, doesn't exist in TestSuiteMap. This should not happen [ErrorID: %s]",
			testSuiteUuid,
			errorId))

		//TODO Send ERRORS over error-channel
		fmt.Println(err)

		return nil,
			"",
			nil,
			nil,
			err

	}
	currentTestSuiteModel = *currentTestSuiteModelPtr

	// If TestCase already has a chosen OwnerDomain then set that value
	var tempCurrentOwnerDomain string

	var testCaseHasOwnerDomain bool

	if len(currentTestSuiteModel.TestSuiteUIModelBinding.TestSuiteOwnerDomainUuid) > 0 {
		testCaseHasOwnerDomain = true
		tempCurrentOwnerDomain = currentTestSuiteModel.TestSuiteUIModelBinding.TestSuiteOwnerDomainUuid
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
	var testCaseOwnerDomainContainer *fyne.Container
	var testCaseOwnerDomainNameFormContainer *fyne.Container
	testCaseOwnerDomainContainer = container.New(layout.NewVBoxLayout())
	testCaseOwnerDomainNameFormContainer = container.New(layout.NewFormLayout())

	// Add Header to the Forms-container
	var headerLabel *widget.Label
	headerLabel = widget.NewLabel("Domain that 'Own' the TestSuite")
	headerLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCaseOwnerDomainNameFormContainer.Add(headerLabel)

	// Generate Warnings-rectangle for valid value, or that value exist
	//var valueIsValidWarningBox *canvas.Rectangle
	var colorToUse color.NRGBA
	colorToUse = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	valueIsValidWarningBox = canvas.NewRectangle(colorToUse)

	// Add the DropDown box with all domains that can own the TestCase
	newOwnerDomainSelect = widget.NewSelect(options,
		func(value string) {
			// This function is called when an option is selected.

			// Store Domain back to TestSuite-model
			currentTestSuiteModel.TestSuiteUIModelBinding.TestSuiteOwnerDomainUuid =
				testCasesModel.DomainsThatCanOwnTheTestCaseMap[value].DomainUuid
			currentTestSuiteModel.TestSuiteUIModelBinding.TestSuiteOwnerDomainName =
				testCasesModel.DomainsThatCanOwnTheTestCaseMap[value].DomainName

			// Get the TestCaseUI-model
			var testCasesUiModelMap map[string]*testCaseGraphicalAreasStruct
			var testCaseGraphicalAreas *testCaseGraphicalAreasStruct
			testCasesUiModelMap = testCasesUiCanvasObject.TestCasesUiModelMap
			testCaseGraphicalAreas, _ = testCasesUiModelMap[testCaseUuid]

			//var testCaseMetaDataArea fyne.CanvasObject
			var metaDataAccordion *widget.Accordion
			_, metaDataAccordion, err = testCasesUiCanvasObject.GenerateMetaDataAreaForTestCase(
				tempTestCasePtr,
				testCaseUuid,
				testCasesUiCanvasObject.TestCasesModelReference.DomainsThatCanOwnTheTestCaseMap[value].DomainUuid)

			if err != nil {
				log.Println(err, metaDataAccordion)

				return
			}

			//var testCaseMetaDataAreaContainer fyne.Container
			//var ok binding.BoolList
			var TestCaseMetaDataAreaCanvasObject fyne.CanvasObject
			TestCaseMetaDataAreaCanvasObject = *testCaseGraphicalAreas.TestCaseMetaDataArea
			if testCaseMetaDataAreaContainer, ok := TestCaseMetaDataAreaCanvasObject.(*fyne.Container); ok {
				if ok != true {
					log.Fatalln("couldn't cast to fyne.Container")
				}

				metaDataAccordion.OpenAll()

				testCaseMetaDataAreaContainer.Objects[0] = metaDataAccordion

				testCaseMetaDataAreaContainer.Refresh()
			}

			// Save back TestCaseUI
			testCasesUiModelMap[testCaseUuid] = testCaseGraphicalAreas

			testCasesUiCanvasObject.TestCasesUiModelMap = testCasesUiModelMap

			// Clear the set MetaData in the TestCase
			//tempTestCasePtr.TestCaseMetaDataPtr = nil

			//testCaseMetaDataArea.Refresh()

			// Store back TestCase-model in Map
			//testCasesUiCanvasObject.TestCasesModelReference.TestCasesMapPtr[testCaseUuid] = tempTestCasePtr

			// Set Warning box that value is not selected
			if len(value) == 0 {
				valueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
			} else {
				valueIsValidWarningBox.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
			}
		})

	// Create a custom SelectComboBox, with valueIsValidWarningBox
	var ownerDomainCustomSelectComboBox *customSelectComboBox
	ownerDomainCustomSelectComboBox = newCustomSelectComboBoxWidget(newOwnerDomainSelect, valueIsValidWarningBox)

	// Add the widget to the Forms-container
	testCaseOwnerDomainNameFormContainer.Add(ownerDomainCustomSelectComboBox)

	// Add TestEnvironment, if TestSuite-metadata exist *********************
	// Get the MetaDataGroups depending on Domain
	var metaDataGroupsMapPtr *map[string]*testCaseModel.MetaDataGroupStruct
	var metaDataGroupsMap map[string]*testCaseModel.MetaDataGroupStruct
	var testCaseMetaDataForDomainsMap map[string]*testCaseModel.TestCaseMetaDataForDomainsForMapStruct
	var testCaseMetaDataForDomainPtr *testCaseModel.TestCaseMetaDataForDomainsForMapStruct
	var testCaseMetaDataForDomain testCaseModel.TestCaseMetaDataForDomainsForMapStruct
	testCaseMetaDataForDomainsMap = testCasesModel.TestCaseMetaDataForDomains.TestCaseMetaDataForDomainsMap
	testCaseMetaDataForDomainPtr, existsInMap = testCaseMetaDataForDomainsMap[testSuiteUuid]
	if existsInMap == false {

		errorId := "91e00fd7-e9bc-4172-b7ad-0f6684514e2f"
		err = errors.New(fmt.Sprintf("Domain with Uuid '%s' doesn't exist in'testCaseMetaDataForDomainsMap'. Should never happen [ErrorID: %s]",
			testSuiteUuid, errorId))

		return nil,
			"",
			nil,
			nil,
			err

	}

	testCaseMetaDataForDomain = *testCaseMetaDataForDomainPtr
	metaDataGroupsMapPtr, _ = testCaseModel.ConvertTestCaseMetaData(testCaseMetaDataForDomain.TestCaseMetaDataForDomainPtr)
	metaDataGroupsMap = *metaDataGroupsMapPtr

	// Get MetaDataGroup for "TestSuite"
	var tempMetaDataGroupForTestSuitePtr *testCaseModel.MetaDataGroupStruct
	tempMetaDataGroupForTestSuitePtr, existsInMap = metaDataGroupsMap["TestSuite"]

	if existsInMap == false {
		errorId := "0a285304-2f4a-4c27-93b6-ed9cb58b55ff"
		err = errors.New(fmt.Sprintf("TestSuite MetaDataGroup '%s' doesn't exist in'metaDataGroupsMap'. Should never happen [ErrorID: %s]",
			"TestSuite", errorId))

		return nil,
			"",
			nil,
			nil,
			err

	}

	// Get 'MetaDataInGroupMap'
	var tempMetaDataInGroupMapPtr *map[string]*testCaseModel.MetaDataInGroupStruct
	var tempMetaDataInGroupMap map[string]*testCaseModel.MetaDataInGroupStruct
	tempMetaDataInGroupMapPtr = tempMetaDataGroupForTestSuitePtr.MetaDataInGroupMapPtr
	tempMetaDataInGroupMap = *tempMetaDataInGroupMapPtr

	// Get the TestENvironment for the TestSuite
	var tempMetaDataInGroupPtr *testCaseModel.MetaDataInGroupStruct
	tempMetaDataInGroupPtr, existsInMap = tempMetaDataInGroupMap["TestEnvironment"]

	if existsInMap == false {
		errorId := "ad019e92-0f38-4e4a-8864-ceab5d701f46"
		err = errors.New(fmt.Sprintf("TestSuite MetaDataItem '%s' doesn't exist in'metaDataGroupsMap'. Should never happen [ErrorID: %s]",
			"TestEnvironment", errorId))

		return nil,
			"",
			nil,
			nil,
			err

	}

	// Add Header to the Forms-container regarding TestExecution Environment
	var testEnvironMentLabel *widget.Label
	testEnvironMentLabel = widget.NewLabel("TestExecution Environment")
	testEnvironMentLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCaseOwnerDomainNameFormContainer.Add(testEnvironMentLabel)

	var testEnvironmentContainer *fyne.Container
	testEnvironmentContainer = testSuiteUiModel.buildTestEnvironmentGUIContainer(tempMetaDataInGroupPtr)
	testCaseOwnerDomainNameFormContainer.Add(testEnvironmentContainer)

	return testCaseOwnerDomainContainer,
		tempCurrentOwnerDomainToBeChosenInDropDown,
		newOwnerDomainSelect,
		valueIsValidWarningBox,
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

// Generates the TestEnvironment container in the TestSuite
func (testSuiteUiModel *TestSuiteUiStruct) buildTestEnvironmentGUIContainer(
	metaDataItemPtr *testCaseModel.MetaDataInGroupStruct) (
	testEnvironmentContainer *fyne.Container) {

	var metaDataItem testCaseModel.MetaDataInGroupStruct
	metaDataItem = *metaDataItemPtr

	// append '*' to the label if it's mandatory
	label := metaDataItem.MetaDataName
	if metaDataItem.Mandatory {
		label += " *"
	}

	// Create correct widget depending on if the item is SingleSelect or MultiSelect
	switch metaDataItem.SelectType {

	case testCaseModel.MetaDataSelectType_SingleSelect:

		var valueIsValidWarningBox *canvas.Rectangle

		// Generate Warnings-rectangle for valid value, or that value exist
		//var valueIsValidWarningBox *canvas.Rectangle
		var colorToUse color.NRGBA
		colorToUse = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
		valueIsValidWarningBox = canvas.NewRectangle(colorToUse)

		sel := widget.NewSelect(metaDataItem.AvailableMetaDataValues, func(val string) {

			// store value in TestSuite-model
			testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestSuiteExecutionEnvironment = val

			// Set Warning box that value is not selected
			if len(val) == 0 && metaDataItem.Mandatory == true {
				valueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
			} else {
				valueIsValidWarningBox.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
			}

		})

		// Extract Selected values from TestCase
		var selectedValue string
		for _, availableValue := range metaDataItem.AvailableMetaDataValues {
			if testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.
				TestSuiteExecutionEnvironment == availableValue {
				selectedValue = availableValue
				break
			}
		}

		sel.PlaceHolder = "Choose..."
		// apply the existing selection if any
		if selectedValue != "" {
			sel.SetSelected(selectedValue)
		}

		// Resize the DropDown

		// Create a custom SelectComboBox, with valueIsValidWarningBox
		var customTestEnvironmentSelectComboBox *customSelectComboBox
		customTestEnvironmentSelectComboBox = newCustomSelectComboBoxWidget(sel, valueIsValidWarningBox)

		// wrap in a 1-cell grid to force width
		w := calcSelectWidth(metaDataItem.AvailableMetaDataValues)
		wrapper := container.New(
			layout.NewGridWrapLayout(fyne.NewSize(w, sel.MinSize().Height)),
			customTestEnvironmentSelectComboBox,
		)

		// Set Warning box that value is not selected
		if len(customTestEnvironmentSelectComboBox.selectComboBox.Selected) == 0 && metaDataItem.Mandatory == true {
			valueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
		} else {
			valueIsValidWarningBox.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
		}

		testEnvironmentContainer = container.NewVBox(
			widget.NewLabel("   "+label),
			wrapper)

	default:

		errorId := "9e7f27a0-2622-40b5-a09e-fb76ee47bd66"
		err := errors.New(fmt.Sprintf("Unhandled 'metaDataItem.SelectType' '%d'. [ErrorID: %s]",
			metaDataItem.SelectType,
			errorId))

		//TODO Send ERRORS over error-channel
		fmt.Println(err)

		testEnvironmentContainer = container.NewVBox(widget.NewLabel(err.Error()))

		return testEnvironmentContainer

	}

	// top‐level grid: each card cell is 220×180
	return testEnvironmentContainer

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
