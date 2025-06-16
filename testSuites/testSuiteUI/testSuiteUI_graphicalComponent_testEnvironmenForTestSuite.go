package testSuiteUI

import (
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

// Generate the OwnerDomain Area for the TestSuite
func (testSuiteUiModel *TestSuiteUiStruct) generateTestEnvironmentForTestSuite(
	testSuiteUuid string,
	testCasesModel *testCaseModel.TestCasesModelsStruct) (
	testEnvironmentContainer *fyne.Container,
	customTestEnvironmentSelectComboBox *customSelectComboBox,
	err error) {

	var existsInMap bool

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
			nil,
			err

	}

	var testEnvironmentFormContainer *fyne.Container
	testEnvironmentFormContainer = container.New(layout.NewVBoxLayout())

	// Add Header to the Forms-container regarding TestExecution Environment
	var testEnvironmentLabel *widget.Label
	testEnvironmentLabel = widget.NewLabel("TestExecution Environment")
	testEnvironmentLabel.TextStyle = fyne.TextStyle{Bold: true}
	testEnvironmentFormContainer.Add(testEnvironmentLabel)

	var testEnvironmentUIContainer *fyne.Container
	testEnvironmentUIContainer, customTestEnvironmentSelectComboBox = testSuiteUiModel.
		buildTestEnvironmentGUIContainer(tempMetaDataInGroupPtr)
	testEnvironmentFormContainer.Add(testEnvironmentUIContainer)

	testEnvironmentContainer = container.NewVBox(testEnvironmentFormContainer)

	return testEnvironmentContainer,
		customTestEnvironmentSelectComboBox,

		err
}

// Sets the Selected value for the DropDown specifying the Owner-Domain of the TestSuite
func (testSuiteUiModel *TestSuiteUiStruct) setSelectedTestEnvironmentForTestSuite(
	selectedTestEnvironment string,
	testEnvironmentCustomSelectComboBox *customSelectComboBox) {

	// Set the Visible value for DropDown, if there is any
	if len(selectedTestEnvironment) > 0 {
		testEnvironmentCustomSelectComboBox.selectComboBox.SetSelected(selectedTestEnvironment)
	}

	// Set correct warning box color
	if len(testEnvironmentCustomSelectComboBox.selectComboBox.Selected) == 0 {
		testEnvironmentCustomSelectComboBox.rectangle.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	} else {
		testEnvironmentCustomSelectComboBox.rectangle.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
	}
}

// Generates the TestEnvironment container in the TestSuite
func (testSuiteUiModel *TestSuiteUiStruct) buildTestEnvironmentGUIContainer(
	metaDataItemPtr *testCaseModel.MetaDataInGroupStruct) (
	testEnvironmentContainer *fyne.Container,
	customTestEnvironmentSelectComboBox *customSelectComboBox) {

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

		return testEnvironmentContainer, customTestEnvironmentSelectComboBox

	}

	// top‐level grid: each card cell is 220×180
	return testEnvironmentContainer, customTestEnvironmentSelectComboBox

}
