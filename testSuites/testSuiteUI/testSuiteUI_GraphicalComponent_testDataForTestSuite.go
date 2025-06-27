package testSuiteUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testDataSelector/testDataSelectorForTestSuite"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/jlambert68/FenixScriptEngine/testDataEngine"
	"log"
	"strings"
)

// Generate the TestData-table Area for the TestSuite
func (testSuiteUiModel *TestSuiteUiStruct) generateSelectedTestDataForTestSuiteArea(
	testCaseUuid string) (
	*fyne.Container,
	error) {

	// Initiate the TestData-object used for keeping Groups and their TestData in the TestSuiteModelPtr
	// Will only been done when a TestSuite hasn't been locked down due to user selected OwnerDomain and Environment
	if testSuiteUiModel.TestSuiteModelPtr.HasLockButtonBeenClickedAndBothOwnerDomainAndTestEnvironmentHaveValues() == false {
		testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestDataPtr = &testDataEngine.TestDataForGroupObjectStruct{}
	}

	// Accordion objects
	var testDataAccordionItem *widget.AccordionItem
	var testDataAccordion *widget.Accordion

	// The slices for Groups, TestDataPoints for a Group and the specific TestDataRows for a TestDataPoint
	var testDataPointGroups []string
	var testDataPointsForAGroup []string
	var testDataRowsForATestDataPoint []string

	var generateTestDataAsRichTextFunction func(bool)
	var generateTestDataAsRichTextCallBackFunction func()

	// Create UI component for selected TestData-selectors
	testDataSingleSelectFormContainer := container.New(layout.NewFormLayout())

	// Create the Container holding the RichText
	richTextContainer := container.NewVBox()

	// Create the TestContainer containing both single select and RichText containers
	testDataContainer := container.NewVBox()

	// Add the Single Select container and the RichText container for TestData
	testDataContainer.Add(testDataSingleSelectFormContainer)
	testDataContainer.Add(richTextContainer)

	var selectorAndButtonContainer *fyne.Container
	var testDataPointRadioGroupContainer *fyne.Container

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
	testDataPointGroupsSelect = widget.NewSelect(
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
			testDataPointsForAGroupSelect.SetOptions(testDataPointsToStringSliceFunction(selected))
			testDataPointsForAGroupSelect.Refresh()

			// UnSelect in DropDown- and List for TestDataPoints
			testDataPointsForAGroupSelect.ClearSelected()

		})

	// Create the Groups TestDataPoints dropdown - <Sub Custody/Main TestData Area/SEK/AccTest/SE/CRDT/CH/Switzerland/BBH/EUR/EUR/SEK>
	testDataPointsForAGroupSelect = widget.NewSelect(testDataPointsToStringSliceFunction(
		testDataPointGroupsSelectSelectedInMainTestSuiteArea), func(selected string) {

		testDataPointForAGroupSelectSelectedInMainTestSuiteArea = selected

		// Store Selected value in 'TestSuite'
		testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestDataPtr.SelectedTestDataPoint = selected

		// Select the correct TestDataPoint in the dropdown for TestDataPoints
		testDataRowsForTestDataPointsSelect.SetOptions(testDataRowSliceToStringSliceFunction(
			testDataPointGroupsSelect.Selected, selected))
		testDataRowsForTestDataPointsSelect.Refresh()

		// UnSelect in DropDown- and List for Specific TestDataPoints
		testDataRowsForTestDataPointsSelect.ClearSelected()

	})

	// Create the Groups Specific TestDataPoint dropdown - <All the specific values>
	testDataRowsForTestDataPointsSelect = widget.NewSelect(testDataRowSliceToStringSliceFunction(
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
			testDataContainer,
			testDataPointRadioGroupContainer,
			generateTestDataAsRichTextCallBackFunction,
			testDataPointGroupsSelect,
			testDataPointsForAGroupSelect,
			testDataRowsForTestDataPointsSelect)

	})

	// Define the callback function to be able to trigger correct way of showing TestDataPointRows or not
	generateTestDataAsRichTextCallBackFunction = func() {
		generateTestDataAsRichTextFunction(showTestDataPointRows)
	}

	// Define the function that creates the RichText for TestDataGroups, TestDataPoints and TestDataPointRows
	generateTestDataAsRichTextFunction = func(showTestDataPointRows bool) {

		// Generate All TestData as RichText component
		var testDataAsRichTextBuilder strings.Builder
		// Loop all TestDataGroups
		for testDataGroupIndex, testDataGroup := range testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestDataPtr.
			ListTestDataGroups() {

			// Add extra empty line for each of every extra TestDataGroup
			if testDataGroupIndex > 0 {
				testDataAsRichTextBuilder.WriteString(fmt.Sprintf(" \n\n "))
			}

			// Add TestDataGroup to StringBuilder
			testDataAsRichTextBuilder.WriteString(fmt.Sprintf("## %s (TestDataGroup)\n\n ", testDataGroup))

			// For each TestDataGroup loop all its TestDataPoints
			for _, testDataPoint := range testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestDataPtr.
				ListTestDataGroupPointsForAGroup(testDataGroup) {

				// Add extra empty line for each of every extra TestDataPoints
				//if testDataPointIndex > 0 {
				//	testDataAsRichTextBuilder.WriteString(fmt.Sprintf(" \n\n "))
				//}

				// Add TestDataPoint to StringBuilder
				testDataAsRichTextBuilder.WriteString(fmt.Sprintf("### %s (TestDataPoint)\n\n ", testDataPoint))

				// Only add TestDataPointRow when they should be visible
				if showTestDataPointRows == true {

					// For each TestDataPoint loop all its TestDataRows
					for _, testDataRow := range testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestDataPtr.
						ListTestDataRowsForAGroupPoint(testDataGroup, testDataPoint) {

						// Add TestDataRow to StringBuilder
						testDataAsRichTextBuilder.WriteString(fmt.Sprintf("%s\n\n ", testDataRow))
					}
				}
			}
		}

		// Create the RichText component for the TestData
		testDataAsRichText = widget.NewRichTextFromMarkdown(testDataAsRichTextBuilder.String())
		testDataAsRichTextBorderContainer := container.NewBorder(
			widget.NewLabel("All TestData for this TestSuite"),
			nil, nil, nil,
			testDataAsRichText)
		//testDataAsRichTextScrollContainer := container.NewHScroll(testDataAsRichTextBorderContainer)

		testDataContainer.Remove(richTextContainer)
		testDataContainer.Add(testDataAsRichTextBorderContainer)

		// Store new 'TestDataAsRichTextBorderContainer' back into 'TestDataContainer'
		richTextContainer = testDataAsRichTextBorderContainer
		testDataContainer.Refresh()

	}

	// Generate the Radio group for showing/Hiding TestDataPointRows
	generateRichTextTestDataRadioGroup := widget.NewRadioGroup([]string{"Show TestDataPointRows", "Hide TestDataPointRows"}, func(selected string) {

		switch selected {
		case "Show TestDataPointRows":
			showTestDataPointRows = true
			generateTestDataAsRichTextFunction(showTestDataPointRows)

		case "Hide TestDataPointRows":
			showTestDataPointRows = false
			generateTestDataAsRichTextFunction(showTestDataPointRows)

		case "":
			// No change so do nothing

		default:
			errorId := "a5536546-fd85-4f13-9071-509437ff0d7e"
			errorMsg := fmt.Sprintf("RadioGroup selected value not recognized (%s) [ErrorId = %s]",
				selected,
				errorId)

			log.Fatalln(errorMsg)
		}
	})
	generateRichTextTestDataRadioGroup.Horizontal = true
	generateRichTextTestDataRadioGroup.SetSelected("Hide TestDataPointRows")
	showTestDataPointRows = false

	// Add RadioGroup to its container
	testDataPointRadioGroupContainer = container.NewVBox(
		generateRichTextTestDataRadioGroup)

	// Header regarding the Select-components
	testDataContainer.Add(widget.NewLabel("TestData to be used for manual single TestSuite Execution"))

	// Add the Select UI component for TestData-selectors
	testDataSingleSelectFormContainer.Add(widget.NewLabel("TestData Group"))
	testDataSingleSelectFormContainer.Add(testDataPointGroupsSelect)

	testDataSingleSelectFormContainer.Add(widget.NewLabel("TestData Point"))
	testDataSingleSelectFormContainer.Add(testDataPointsForAGroupSelect)

	testDataSingleSelectFormContainer.Add(widget.NewLabel("TestData Row"))
	testDataSingleSelectFormContainer.Add(testDataRowsForTestDataPointsSelect)

	// Add SingleSelect  container to the overall TestDataContainer
	testDataContainer.Add(testDataSingleSelectFormContainer)

	// Add separator
	testDataContainer.Add(widget.NewSeparator())

	// Create the RichText component for the TestData
	richTextContainer = container.NewBorder(
		widget.NewLabel("No TestData added to this TestSuite yet."),
		nil, nil, nil, nil)

	//richTextContainer.Add(widget.NewLabel("All TestData for this TestSuite"))
	richTextContainer.Add(testDataAsRichText)

	// Add RichText container to the overall TestDataContainer
	testDataContainer.Add(richTextContainer)

	// If there is no TestData then hide the "TestData-container"
	if len(testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestDataPtr.ListTestDataGroups()) == 0 {
		testDataContainer.Hide()
		testDataPointRadioGroupContainer.Hide()
	} else {
		testDataPointGroupsSelect.SetOptions(getTestGroupsFromTestDataEngineFunction())
		testDataContainer.Show()
		testDataPointRadioGroupContainer.Show()
		testDataPointGroupsSelect.Refresh()
	}

	// Create an Accordion item for the buttons
	buttonContainer := container.NewHBox(selectTestDataButton, testDataPointRadioGroupContainer)

	selectorAndButtonContainer = container.NewBorder(buttonContainer, nil, nil, nil, testDataContainer)
	selectorAndButtonContainer.Refresh()

	// Create an Accordion item for the list
	testDataAccordionItem = widget.NewAccordionItem("TestData", selectorAndButtonContainer)

	testDataAccordion = widget.NewAccordion(testDataAccordionItem) // widget.NewAccordion(tableAccordionItem)

	// Create the VBox-container that will be returned
	testDataArea := container.NewVBox(testDataAccordion, widget.NewLabel(""), widget.NewSeparator())

	return testDataArea, nil

}
