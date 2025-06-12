package testSuiteUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testDataSelector/testDataSelectorForTestSuite"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/jlambert68/FenixScriptEngine/testDataEngine"
)

// Generate the TestData-table Area for the TestSuite
func (testSuiteUiModel *TestSuiteUiStruct) generateSelectedTestDataForTestSuiteArea(
	testCaseUuid string) (
	fyne.CanvasObject,
	error) {

	// Initiate the TestData-object used for keeping Groups and their TestData in the TestSuiteModelPtr
	testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestDataPtr = &testDataEngine.TestDataForGroupObjectStruct{}

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

		testDataPointGroups = testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestDataPtr.ListTestDataGroups()

		return testDataPointGroups
	}

	// Create function that converts a TestDataPointsSlice into a string slice
	testDataPointsToStringSliceFunction := func(testDataGroup string) []string {

		if testDataGroup == "" {
			return []string{}
		}

		testDataPointsForAGroup = testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestDataPtr.
			ListTestDataGroupPointsForAGroup(testDataGroup)

		return testDataPointsForAGroup
	}

	// Create function that converts a slice with the specific TestDataPoints into a string slice
	testDataRowSliceToStringSliceFunction := func(testDataGroup string, testDataGroupPoint string) []string {

		if testDataGroup == "" || testDataGroupPoint == "" {
			return []string{}
		}

		testDataRowsForATestDataPoint = testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestDataPtr.
			ListTestDataRowsForAGroupPoint(
				testDataGroup, testDataGroupPoint)

		return testDataRowsForATestDataPoint
	}

	// Create the Group dropdown - <Name of the group>
	testDataPointGroupsSelectInMainTestSuiteArea = widget.NewSelect(
		getTestGroupsFromTestDataEngineFunction(), func(selected string) {

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

			testDataPointGroupsSelectSelectedInMainTestSuiteArea = selected

			// Store Selected value in 'TestSuite'
			testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestDataPtr.SelectedTestDataGroup = selected

			// Select the correct TestDataPoint in the dropdown for TestDataPoints
			testDataPointsForAGroupSelectInMainTestSuiteArea.SetOptions(testDataPointsToStringSliceFunction(selected))
			testDataPointsForAGroupSelectInMainTestSuiteArea.Refresh()

			// UnSelect in DropDown- and List for TestDataPoints
			testDataPointsForAGroupSelectInMainTestSuiteArea.ClearSelected()

		})

	// Create the Groups TestDataPoints dropdown - <Sub Custody/Main TestData Area/SEK/AccTest/SE/CRDT/CH/Switzerland/BBH/EUR/EUR/SEK>
	testDataPointsForAGroupSelectInMainTestSuiteArea = widget.NewSelect(testDataPointsToStringSliceFunction(
		testDataPointGroupsSelectSelectedInMainTestSuiteArea), func(selected string) {

		testDataPointForAGroupSelectSelectedInMainTestSuiteArea = selected

		// Store Selected value in 'TestCase'
		testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestDataPtr.SelectedTestDataPoint = selected

		// Select the correct TestDataPoint in the dropdown for TestDataPoints
		testDataRowsForTestDataPointsSelectInMainTestSuiteArea.SetOptions(testDataRowSliceToStringSliceFunction(
			testDataPointGroupsSelectInMainTestSuiteArea.Selected, selected))
		testDataRowsForTestDataPointsSelectInMainTestSuiteArea.Refresh()

		// UnSelect in DropDown- and List for Specific TestDataPoints
		testDataRowsForTestDataPointsSelectInMainTestSuiteArea.ClearSelected()

	})

	// Create the Groups Specific TestDataPoint dropdown - <All the specific values>
	testDataRowsForTestDataPointsSelectInMainTestSuiteArea = widget.NewSelect(testDataRowSliceToStringSliceFunction(
		testDataPointGroupsSelectSelectedInMainTestSuiteArea, testDataPointForAGroupSelectSelectedInMainTestSuiteArea), func(selected string) {

		testDataRowForTestDataPointsSelectSelectedInMainTestSuiteArea = selected

		// Store Selected value in 'TestSuite'
		testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestDataPtr.SelectedTestDataPointRowSummary = selected

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
			testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestDataPtr.
				GetTestDataPointValuesMapBasedOnGroupPointNameAndSummaryValue(
					testDataPointGroupsSelectSelectedInMainTestSuiteArea,
					testDataPointForAGroupSelectSelectedInMainTestSuiteArea,
					testDataRowForTestDataPointsSelectSelectedInMainTestSuiteArea)

		// Store Selected TestData on the TestSuite
		testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestDataPtr.SelectedTestDataDomainUuid = domainUuid
		testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestDataPtr.SelectedTestDataDomainName = domainName
		testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestDataPtr.SelectedTestDataDomainTemplateName = domainTemplateName
		testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestDataPtr.SelectedTestDataTestDataAreaUuid = testDataAreaUuid
		testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestDataPtr.SelectedTestDataAreaName = testDataAreaName
		testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestDataPtr.TestDataColumnDataNameToValueMap = testDataColumnDataNameMap

	})

	// Select TestData the TestSuite
	selectTestDataButton := widget.NewButton("Add TestData to TestSuite", func() {

		// Open up TestData Selector Window
		testDataSelectorForTestSuite.MainTestDataSelector(
			*sharedCode.FenixAppPtr,
			*sharedCode.FenixMasterWindowPtr,
			testSuiteUiModel.TestSuiteModelPtr,
			testCaseUuid,
			testDataSelectorsContainer,
			testDataPointGroupsSelectInMainTestSuiteArea,
			testDataPointsForAGroupSelectInMainTestSuiteArea,
			testDataRowsForTestDataPointsSelectInMainTestSuiteArea)

	})

	// Add the Select UI component for TestData-selectors
	testDataSelectorsContainer.Add(widget.NewLabel("TestData Group"))
	testDataSelectorsContainer.Add(testDataPointGroupsSelectInMainTestSuiteArea)

	testDataSelectorsContainer.Add(widget.NewLabel("TestData Point"))
	testDataSelectorsContainer.Add(testDataPointsForAGroupSelectInMainTestSuiteArea)

	testDataSelectorsContainer.Add(widget.NewLabel("TestData Row"))
	testDataSelectorsContainer.Add(testDataRowsForTestDataPointsSelectInMainTestSuiteArea)

	// If there is no TestData then hide the "Select-boxes"
	if len(testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestDataPtr.ListTestDataGroups()) == 0 {
		testDataSelectorsContainer.Hide()
	} else {
		testDataPointGroupsSelectInMainTestSuiteArea.SetOptions(getTestGroupsFromTestDataEngineFunction())
		testDataSelectorsContainer.Show()
		testDataPointGroupsSelectInMainTestSuiteArea.Refresh()
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
