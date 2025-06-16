package testSuiteUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testSuites/testSuitesModel"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
	"strings"
)

// Generate the TestCaseDescription Area for the TestCase
func (testSuiteUiModel *TestSuiteUiStruct) generateTestCaseDescriptionArea(
	testSuiteUuid string) (
	testCaseDescriptionAreaContainer *fyne.Container,
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
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":            "07ccaec5-3511-4ae0-a370-9cbd51f337ed",
			"testSuiteUuid": testSuiteUuid,
		}).Fatal("TestSuite doesn't exist in TestSuiteMap. This should not happen")
	}
	currentTestSuiteModel = *currentTestSuiteModelPtr

	// Create Form-layout container to be used for Description
	var testCaseDescriptionContainer *fyne.Container
	var testCaseDescriptionFormContainer *fyne.Container
	testCaseDescriptionContainer = container.New(layout.NewVBoxLayout())
	testCaseDescriptionFormContainer = container.New(layout.NewFormLayout())

	// Add Header to the Forms-container
	var headerLabel *widget.Label
	headerLabel = widget.NewLabel("TestCaseDescription")
	headerLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCaseDescriptionFormContainer.Add(headerLabel)

	// Add the Entry-widget for TestCaseDescription
	newTestCaseDescriptionEntry := widget.NewMultiLineEntry()
	newTestCaseDescriptionEntry.SetText(currentTestSuiteModel.TestSuiteUIModelBinding.TestSuiteDescription)

	// Change vertical Size of Entry-widget for TestCaseDescription
	newTestCaseDescriptionEntry.SetMinRowsVisible(5)

	// Change Description in model when UI-component-Description-value is changed
	newTestCaseDescriptionEntry.OnChanged = func(newValue string) {

		var trimmedValue string
		trimmedValue = strings.Trim(newValue, " ")

		// Get entryOnChangetestSuitesMap
		var entryOnChangetestSuitesMap map[string]*testSuitesModel.TestSuiteModelStruct
		entryOnChangetestSuitesMap = *testSuitesModel.TestSuitesModelPtr.TestSuitesMapPtr

		// Get a pointer to the TestSuite-model and the TestSuite-model itself
		var entryOnChangeCurrentTestSuiteModelPtr *testSuitesModel.TestSuiteModelStruct
		entryOnChangeCurrentTestSuiteModelPtr, existsInMap = entryOnChangetestSuitesMap[testSuiteUuid]

		if existsInMap == false {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":            "48285fad-09a3-4e52-8f34-a104cbcf358a",
				"testSuiteUuid": testSuiteUuid,
			}).Fatal("TestSuite doesn't exist in TestSuiteMap. This should not happen")
		}

		// Store the Description in the TestSuiteModel
		entryOnChangeCurrentTestSuiteModelPtr.TestSuiteUIModelBinding.TestSuiteDescription = trimmedValue

	}

	// Add the Entry-widget to the Forms-container
	testCaseDescriptionFormContainer.Add(newTestCaseDescriptionEntry)

	// Create the VBox-container that will be returned
	testCaseDescriptionContainer = container.NewVBox(testCaseDescriptionFormContainer)

	// Create and add Slider to control vertical size of the vertical size of Entry-widget for TestCaseDescription
	var verticalSizeSlider *widget.Slider
	verticalSizeSlider = &widget.Slider{Value: 5, Step: 1, Min: 3, Max: 30, OnChanged: func(f float64) {
		newTestCaseDescriptionEntry.SetMinRowsVisible(int(verticalSizeSlider.Value))
	}}

	testCaseDescriptionContainer.Add(verticalSizeSlider)

	return testCaseDescriptionContainer, err
}
