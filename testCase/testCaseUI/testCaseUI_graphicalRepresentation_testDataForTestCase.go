package testCaseUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testCase/testCaseModel"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/jlambert68/FenixScriptEngine/testDataEngine"
	"github.com/sirupsen/logrus"
)

// Generate the TestData-table Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateSelectedTestDataForTestCaseArea(
	testCaseUuid string) (
	fyne.CanvasObject,
	error) {

	var existInMap bool
	var currentTestCase testCaseModel.TestCaseModelStruct

	currentTestCase, existInMap = testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]
	if existInMap == false {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":           "0bb2ebf8-fae9-4427-ad82-8fad3a73d6e9",
			"testCaseUuid": testCaseUuid,
		}).Fatal("TestCase doesn't exist in TestCaseMap. This should not happen")
	}

	// Initiate the TestData-object used for keeping Groups and their TestData in the TestCases
	currentTestCase.TestData = &testDataEngine.TestDataForGroupObjectStruct{}

	// Accordion objects
	var testDataAccordionItem *widget.AccordionItem
	var testDataAccordion *widget.Accordion

	// The slices for Groups, TestDataPoints for a Group and the specific TestDataRows for a TestDataPoint
	var testDataPointGroups []string
	var testDataPointsForAGroup []string
	var testDataRowsForATestDataPoint []string

	// Create UI component for selected TestData-selectors
	testDataSelectorsContainer := container.New(layout.NewFormLayout())

	// Create function that converts a GroupSlice into a string slice
	getTestGroupsFromTestDataEngineFunction := func() []string {

		testDataPointGroups = currentTestCase.TestData.ListTestDataGroups()

		return testDataPointGroups
	}

	// Create function that converts a TestDataPointsSlice into a string slice
	testDataPointsToStringSliceFunction := func(testDataGroup string) []string {

		if testDataGroup == "" {
			return []string{}
		}

		testDataPointsForAGroup = currentTestCase.TestData.ListTestDataGroupPointsForAGroup(testDataGroup)

		return testDataPointsForAGroup
	}

	// Create function that converts a slice with the specific TestDataPoints into a string slice
	testDataRowSliceToStringSliceFunction := func(testDataGroup string, testDataGroupPoint string) []string {

		if testDataGroup == "" || testDataGroupPoint == "" {
			return []string{}
		}

		testDataRowsForATestDataPoint = currentTestCase.TestData.ListTestDataRowsForAGroupPoint(
			testDataGroup, testDataGroupPoint)

		return testDataRowsForATestDataPoint
	}

	// Create the Group dropdown - <Name of the group>
	testDataPointGroupsSelectInMainTestCaseArea = widget.NewSelect(getTestGroupsFromTestDataEngineFunction(), func(selected string) {

		// Check that the selected Group name exist options
		if len(selected) > 0 {

			var foundGroupName bool
			for _, groupName := range getTestGroupsFromTestDataEngineFunction() {
				if groupName == selected {
					foundGroupName = true
					break
				}
			}

			if foundGroupName == false {
				selected = ""
			}
		}

		testDataPointGroupsSelectSelectedInMainTestCaseArea = selected

		// Store Selected value in 'TestCase'
		currentTestCase, existInMap = testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]
		if existInMap == false {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":           "ab542075-fd93-4ac8-bd4d-46668f4b131d",
				"testCaseUuid": testCaseUuid,
			}).Fatal("TestCase doesn't exist in TestCaseMap. This should not happen")
		}

		currentTestCase.TestData.SelectedTestDataGroup = selected
		testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid] = currentTestCase

		// Select the correct TestDataPoint in the dropdown for TestDataPoints
		testDataPointsForAGroupSelectInMainTestCaseArea.SetOptions(testDataPointsToStringSliceFunction(selected))
		testDataPointsForAGroupSelectInMainTestCaseArea.Refresh()

		// UnSelect in DropDown- and List for TestDataPoints
		testDataPointsForAGroupSelectInMainTestCaseArea.ClearSelected()

	})

	// Create the Groups TestDataPoints dropdown - <Sub Custody/Main TestData Area/SEK/AccTest/SE/CRDT/CH/Switzerland/BBH/EUR/EUR/SEK>
	testDataPointsForAGroupSelectInMainTestCaseArea = widget.NewSelect(testDataPointsToStringSliceFunction(
		testDataPointGroupsSelectSelectedInMainTestCaseArea), func(selected string) {

		testDataPointForAGroupSelectSelectedInMainTestCaseArea = selected

		// Store Selected value in 'TestCase'
		currentTestCase, existInMap = testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]
		if existInMap == false {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":           "ecb7e72b-8d7c-4e7f-b4f9-6cd3eec48ecf",
				"testCaseUuid": testCaseUuid,
			}).Fatal("TestCase doesn't exist in TestCaseMap. This should not happen")
		}

		currentTestCase.TestData.SelectedTestDataPoint = selected
		testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid] = currentTestCase

		// Select the correct TestDataPoint in the dropdown for TestDataPoints
		testDataRowsForTestDataPointsSelectInMainTestCaseArea.SetOptions(testDataRowSliceToStringSliceFunction(
			testDataPointGroupsSelectInMainTestCaseArea.Selected, selected))
		testDataRowsForTestDataPointsSelectInMainTestCaseArea.Refresh()

		// UnSelect in DropDown- and List for Specific TestDataPoints
		testDataRowsForTestDataPointsSelectInMainTestCaseArea.ClearSelected()

	})

	// Create the Groups Specific TestDataPoint dropdown - <All the specific values>
	testDataRowsForTestDataPointsSelectInMainTestCaseArea = widget.NewSelect(testDataRowSliceToStringSliceFunction(
		testDataPointGroupsSelectSelectedInMainTestCaseArea, testDataPointForAGroupSelectSelectedInMainTestCaseArea), func(selected string) {

		testDataRowForTestDataPointsSelectSelectedInMainTestCaseArea = selected

		// Store Selected value in 'TestCase'
		currentTestCase, existInMap = testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]
		if existInMap == false {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":           "ba37763b-0609-4cff-8f00-4f45b502feab",
				"testCaseUuid": testCaseUuid,
			}).Fatal("TestCase doesn't exist in TestCaseMap. This should not happen")
		}

		currentTestCase.TestData.SelectedTestDataPointRowSummary = selected

		// Extract Unique id, first column, for the TestDataRow

		var (
			domainUuid         string
			domainName         string
			domainTemplateName string
			testDataAreaUuid   string
			testDataAreaName   string
		)

		// Extract correlation between ColumnDataName and data value for DataRow
		var testDataColumnDataNameMap map[string]string // map[TestDataColumnDataNameType]TestDataValueType
		testDataColumnDataNameMap,
			domainUuid, domainName, domainTemplateName, testDataAreaUuid, testDataAreaName =
			currentTestCase.TestData.GetTestDataPointValuesMapBasedOnGroupPointNameAndSummaryValue(
				testDataPointGroupsSelectSelectedInMainTestCaseArea,
				testDataPointForAGroupSelectSelectedInMainTestCaseArea,
				testDataRowForTestDataPointsSelectSelectedInMainTestCaseArea)

		// Store Selected TestData on the TestCase
		currentTestCase.TestData.SelectedTestDataDomainUuid = domainUuid
		currentTestCase.TestData.SelectedTestDataDomainName = domainName
		currentTestCase.TestData.SelectedTestDataDomainTemplateName = domainTemplateName
		currentTestCase.TestData.SelectedTestDataTestDataAreaUuid = testDataAreaUuid
		currentTestCase.TestData.SelectedTestDataAreaName = testDataAreaName
		currentTestCase.TestData.TestDataColumnDataNameToValueMap = testDataColumnDataNameMap

		// Store back the TestCase
		testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid] = currentTestCase

	})

	// Select TestData the TestCase
	selectTestDataButton := widget.NewButton("Add TestData to TestCase", func() {

		currentTestCase, existInMap = testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]
		if existInMap == false {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":           "a54bce68-fa84-4b29-aa62-5d47b8bdc7fb",
				"testCaseUuid": testCaseUuid,
			}).Fatal("TestCase doesn't exist in TestCaseMap. This should not happen")
		}

		// Open up TestData Selector Window
		testCasesUiCanvasObject.MainTestDataSelector(
			*sharedCode.FenixAppPtr,
			*sharedCode.FenixMasterWindowPtr,
			&currentTestCase,
			testCaseUuid,
			testDataSelectorsContainer,
			testDataPointGroupsSelectInMainTestCaseArea,
			testDataPointsForAGroupSelectInMainTestCaseArea,
			testDataRowsForTestDataPointsSelectInMainTestCaseArea)

	})

	// Add the Select UI component for TestData-selectors
	testDataSelectorsContainer.Add(widget.NewLabel("TestData Group"))
	testDataSelectorsContainer.Add(testDataPointGroupsSelectInMainTestCaseArea)

	testDataSelectorsContainer.Add(widget.NewLabel("TestData Point"))
	testDataSelectorsContainer.Add(testDataPointsForAGroupSelectInMainTestCaseArea)

	testDataSelectorsContainer.Add(widget.NewLabel("TestData Row"))
	testDataSelectorsContainer.Add(testDataRowsForTestDataPointsSelectInMainTestCaseArea)

	currentTestCase, existInMap = testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]
	if existInMap == false {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":           "2aa83455-5d5f-42e0-9300-9b1eb55d53b8",
			"testCaseUuid": testCaseUuid,
		}).Fatal("TestCase doesn't exist in TestCaseMap. This should not happen")
	}

	// If there is no TestData then hide the "Select-boxes"
	if len(currentTestCase.TestData.ListTestDataGroups()) == 0 {
		testDataSelectorsContainer.Hide()
	} else {
		testDataPointGroupsSelectInMainTestCaseArea.SetOptions(getTestGroupsFromTestDataEngineFunction())
		testDataSelectorsContainer.Show()
		testDataPointGroupsSelectInMainTestCaseArea.Refresh()
	}

	// Create an Accordion item for the buttons
	buttonContainer := container.NewHBox(selectTestDataButton)

	selectorAndButtonContainer := container.NewBorder(nil, buttonContainer, nil, nil, testDataSelectorsContainer)
	selectorAndButtonContainer.Refresh()

	// Create an Accordion item for the list
	testDataAccordionItem = widget.NewAccordionItem("TestData", selectorAndButtonContainer)

	testDataAccordion = widget.NewAccordion(testDataAccordionItem) // widget.NewAccordion(tableAccordionItem)

	// Create the VBox-container that will be returned
	testDataArea := container.NewVBox(testDataAccordion, widget.NewLabel(""), widget.NewSeparator())

	return testDataArea, nil

}
