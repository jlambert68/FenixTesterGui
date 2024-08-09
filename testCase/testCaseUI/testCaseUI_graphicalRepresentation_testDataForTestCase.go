package testCaseUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testCase/testCaseModel"
	"fmt"
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

	// The Select-items for Groups ans TestDataPoints for a Group
	var testDataPointGroupsSelect *widget.Select
	var testDataPointGroupsSelectSelected string
	var testDataPointsForAGroupSelect *widget.Select
	var testDataPointForAGroupSelectSelected string
	var testDataRowsForTestDataPointsSelect *widget.Select
	var testDataRowForTestDataPointsSelectSelected string

	// Accordion objects
	var testDataAccordionItem *widget.AccordionItem
	var testDataAccordion *widget.Accordion

	// The slices for Groups, TestDataPoints for a Group and the specific TestDataRows for a TestDataPoint
	var testDataPointGroups []string
	var testDataPointsForAGroup []string
	var testDataRowsForATestDataPoint []string

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
	testDataPointGroupsSelect = widget.NewSelect(getTestGroupsFromTestDataEngineFunction(), func(selected string) {

		testDataPointGroupsSelectSelected = selected

		// Select the correct TestDataPoint in the dropdown for TestDataPoints
		testDataPointsForAGroupSelect.SetOptions(testDataPointsToStringSliceFunction(selected))
		testDataPointsForAGroupSelect.Refresh()

		// UnSelect in DropDown- and List for TestDataPoints
		testDataPointsForAGroupSelect.ClearSelected()

	})

	// Create the Groups TestDataPoints dropdown - <Sub Custody/Main TestData Area/SEK/AccTest/SE/CRDT/CH/Switzerland/BBH/EUR/EUR/SEK>
	testDataPointsForAGroupSelect = widget.NewSelect(testDataPointsToStringSliceFunction(
		testDataPointGroupsSelectSelected), func(selected string) {

		testDataPointForAGroupSelectSelected = selected

		// Select the correct TestDataPoint in the dropdown for TestDataPoints
		testDataRowsForTestDataPointsSelect.SetOptions(testDataRowSliceToStringSliceFunction(
			testDataPointGroupsSelect.Selected, selected))
		testDataRowsForTestDataPointsSelect.Refresh()

		// UnSelect in DropDown- and List for Specific TestDataPoints
		testDataRowsForTestDataPointsSelect.ClearSelected()

	})

	// Create the Groups Specific TestDataPoint dropdown - <All the specific values>
	testDataRowsForTestDataPointsSelect = widget.NewSelect(testDataRowSliceToStringSliceFunction(
		testDataPointGroupsSelectSelected, testDataPointForAGroupSelectSelected), func(selected string) {

		testDataRowForTestDataPointsSelectSelected = selected
		fmt.Println(testDataRowForTestDataPointsSelectSelected) //TODO REMOVE

	})

	// Select TestData the TestCase
	selectTestDataButton := widget.NewButton("Select TestData", func() {

		var existInMap bool
		//var currentTestCase testCaseModel.TestCaseModelStruct

		_, existInMap = testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]
		if existInMap == false {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":           "a54bce68-fa84-4b29-aa62-5d47b8bdc7fb",
				"testCaseUuid": testCaseUuid,
			}).Fatal("TestCase doesn't exist in TestCaseMap. This should not happen")
		}

	})

	// Create UI component for selected TestData-selectors
	testDataSelectorsContainer := container.New(layout.NewFormLayout())
	testDataSelectorsContainer.Add(widget.NewLabel("TestData Group"))
	testDataSelectorsContainer.Add(testDataPointGroupsSelect)

	testDataSelectorsContainer.Add(widget.NewLabel("TestData Point"))
	testDataSelectorsContainer.Add(testDataPointsForAGroupSelect)

	testDataSelectorsContainer.Add(widget.NewLabel("TestData Row"))
	testDataSelectorsContainer.Add(testDataRowsForTestDataPointsSelect)

	// Create an Accordion item for the buttons
	buttonContainer := container.NewHBox(selectTestDataButton)

	selectorAndButtonContainer := container.NewBorder(nil, buttonContainer, nil, nil, testDataSelectorsContainer)
	selectorAndButtonContainer.Refresh()

	// Create an Accordion item for the list
	testDataAccordionItem = widget.NewAccordionItem("TestData", selectorAndButtonContainer)

	testDataAccordion = widget.NewAccordion(testDataAccordionItem) // widget.NewAccordion(tableAccordionItem)

	// Create the VBox-container that will be returned
	testDataArea := container.NewVBox(testDataAccordion)

	return testDataArea, nil

}
