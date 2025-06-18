package testSuiteUI

import (
	"FenixTesterGui/testSuites/testSuitesCommandEngine"
	"FenixTesterGui/testSuites/testSuitesModel"
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
func (testSuiteUiModel *TestSuiteUiStruct) generateTestEnvironmentForTestSuite() (
	testEnvironmentContainer *fyne.Container,
	customTestEnvironmentSelectComboBox *customSelectComboBox,
	err error) {

	// Get OwnerDomain
	var ownerDomainUuid string
	ownerDomainUuid = testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestSuiteOwnerDomainUuid

	var existsInMap bool

	// Only process when there is an OwnerDomain
	var tempCurrentOwnerDomainUuid string
	tempCurrentOwnerDomainUuid = testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestSuiteOwnerDomainUuid
	if tempCurrentOwnerDomainUuid == "" {
		testEnvironmentContainer = container.NewVBox(widget.NewLabel("No OwnerDomain specified for TestSuite"))

		return testEnvironmentContainer, nil, err
	}

	// Add TestEnvironment, if TestSuite-metadata exist *********************
	// Get the MetaDataGroups depending on Domain
	var metaDataGroupsMapPtr *map[string]*testSuitesModel.MetaDataGroupStruct
	var metaDataGroupsMap map[string]*testSuitesModel.MetaDataGroupStruct
	var testSuiteMetaDataForDomainsMap map[string]*testSuitesModel.TestSuiteMetaDataForDomainsForMapStruct
	var testSuiteMetaDataForDomainPtr *testSuitesModel.TestSuiteMetaDataForDomainsForMapStruct
	var testSuiteMetaDataForDomain testSuitesModel.TestSuiteMetaDataForDomainsForMapStruct
	testSuiteMetaDataForDomainsMap = testSuitesModel.TestSuitesModelPtr.TestSuiteMetaDataForDomains.
		TestSuiteMetaDataForDomainsMap
	testSuiteMetaDataForDomainPtr, existsInMap = testSuiteMetaDataForDomainsMap[ownerDomainUuid]
	if existsInMap == false {

		testEnvironmentContainer = container.NewVBox(widget.NewLabel("OwnerDomain doesn't have any TestSuite MetaData"))

		return testEnvironmentContainer,
			nil,
			err

	}

	testSuiteMetaDataForDomain = *testSuiteMetaDataForDomainPtr
	metaDataGroupsMapPtr, _ = testSuitesModel.ConvertTestSuiteMetaData(testSuiteMetaDataForDomain.TestSuiteMetaDataForDomainPtr)
	metaDataGroupsMap = *metaDataGroupsMapPtr

	// Get MetaDataGroup for "TestSuite"
	var tempMetaDataGroupForTestSuitePtr *testSuitesModel.MetaDataGroupStruct
	tempMetaDataGroupForTestSuitePtr, existsInMap = metaDataGroupsMap["TestSuite"]

	if existsInMap == false {
		errorId := "0a285304-2f4a-4c27-93b6-ed9cb58b55ff"
		err = errors.New(fmt.Sprintf("TestSuite MetaDataGroup '%s' doesn't exist in'metaDataGroupsMap'. Should never happen [ErrorID: %s]",
			"TestSuite", errorId))

		testEnvironmentContainer = container.NewVBox(widget.NewLabel(err.Error()))

		return testEnvironmentContainer,
			nil,
			err

	}

	// Get 'MetaDataInGroupMap'
	var tempMetaDataInGroupMapPtr *map[string]*testSuitesModel.MetaDataInGroupStruct
	var tempMetaDataInGroupMap map[string]*testSuitesModel.MetaDataInGroupStruct
	tempMetaDataInGroupMapPtr = tempMetaDataGroupForTestSuitePtr.MetaDataInGroupMapPtr
	tempMetaDataInGroupMap = *tempMetaDataInGroupMapPtr

	// Get the TestEnvironment for the TestSuite
	var tempMetaDataInGroupPtr *testSuitesModel.MetaDataInGroupStruct
	tempMetaDataInGroupPtr, existsInMap = tempMetaDataInGroupMap["TestEnvironment"]

	if existsInMap == false {
		errorId := "ad019e92-0f38-4e4a-8864-ceab5d701f46"
		err = errors.New(fmt.Sprintf("TestSuite MetaDataItem '%s' doesn't exist in'metaDataGroupsMap'. Should never happen [ErrorID: %s]",
			"TestEnvironment", errorId))

		testEnvironmentContainer = container.NewVBox(widget.NewLabel(err.Error()))

		return testEnvironmentContainer,
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
	metaDataItemPtr *testSuitesModel.MetaDataInGroupStruct) (
	testEnvironmentContainer *fyne.Container,
	customTestEnvironmentSelectComboBox *customSelectComboBox) {

	var metaDataItem testSuitesModel.MetaDataInGroupStruct
	metaDataItem = *metaDataItemPtr

	// append '*' to the label if it's mandatory
	label := metaDataItem.MetaDataName
	if metaDataItem.Mandatory {
		label += " *"
	}

	// Create correct widget depending on if the item is SingleSelect or MultiSelect
	switch metaDataItem.SelectType {

	case testSuitesModel.MetaDataSelectType_SingleSelect:

		var valueIsValidWarningBox *canvas.Rectangle

		// Generate Warnings-rectangle for valid value, or that value exist
		//var valueIsValidWarningBox *canvas.Rectangle
		var colorToUse color.NRGBA
		colorToUse = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
		valueIsValidWarningBox = canvas.NewRectangle(colorToUse)

		sel := widget.NewSelect(metaDataItem.AvailableMetaDataValues, func(val string) {

			var err error
			// store value in TestSuite-model
			testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestSuiteExecutionEnvironment = val

			// Trigger creation of a 'new' TestSuiteMetaData container for the TestSuite-UI ************************

			// Clear MetaData
			testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.
				TestSuiteMetaDataPtr = &testSuitesModel.TestSuiteMetaDataStruct{}

			// Generate TestSuite's ExecutionEnvironment
			var newTestSuiteMetaDataContainer *fyne.Container

			newTestSuiteMetaDataContainer, err = testSuiteUiModel.GenerateMetaDataAreaForTestCase()
			if err != nil {

				errorId := "f485384b-e1ae-4cf8-a6da-50468c365513"
				errorMessage := fmt.Sprintf("couldn't generate 'TestSuites MetaData-area', err=%s. [ErrorId = %s]",
					err.Error(),
					errorId)

				// Remove old 'testSuiteMetaDataContainer' from stack container
				testSuiteUiModel.testSuiteMetaDataStackContainer.Remove(testSuiteUiModel.testSuiteMetaDataContainer)

				// Create new
				testSuiteUiModel.testSuiteMetaDataContainer = container.NewVBox(widget.NewLabel(errorMessage))

				// Add new 'testSuiteMetaDataContainer' to stack container
				testSuiteUiModel.testSuiteMetaDataStackContainer.Add(testSuiteUiModel.testSuiteMetaDataContainer)

				// Refresh Tabs
				testSuitesCommandEngine.TestSuiteTabsRef.Refresh()

				return

			}

			// Remove old 'testSuiteMetaDataContainer' from stack container
			if testSuiteUiModel.testSuiteMetaDataStackContainer != nil {
				testSuiteUiModel.testSuiteMetaDataStackContainer.Remove(testSuiteUiModel.testSuiteMetaDataContainer)
			}

			// Add new 'testSuiteTestEnvironmentContainer' to stack container
			if testSuiteUiModel.testSuiteMetaDataStackContainer != nil {
				testSuiteUiModel.testSuiteMetaDataStackContainer.Add(newTestSuiteMetaDataContainer)
			} else {
				testSuiteUiModel.testSuiteMetaDataStackContainer = container.NewStack(newTestSuiteMetaDataContainer)

			}

			// Store 'newTestSuiteMetaDataContainer' in old onec place
			testSuiteUiModel.testSuiteMetaDataContainer = newTestSuiteMetaDataContainer

			// Refresh Tabs
			testSuitesCommandEngine.TestSuiteTabsRef.Refresh()

			// Set Warning box that value is not selected
			if len(val) == 0 && metaDataItem.Mandatory == true {
				valueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
			} else {
				valueIsValidWarningBox.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
			}

		})

		// Extract Selected values from TestSuite
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
