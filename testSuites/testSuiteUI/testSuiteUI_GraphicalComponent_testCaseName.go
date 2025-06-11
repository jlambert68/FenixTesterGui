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

// Generate the TestSuiteName Area for the TestSuite
func (testSuiteUiModel *TestSuiteUiStruct) generateTestSuiteNameArea(
	testSuiteUuid string) (
	testSuiteNameArea fyne.CanvasObject,
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
			"ID":            "48285fad-09a3-4e52-8f34-a104cbcf358a",
			"testSuiteUuid": testSuiteUuid,
		}).Fatal("TestSuite doesn't exist in TestSuiteMap. This should not happen")
	}
	currentTestSuiteModel = *currentTestSuiteModelPtr

	// Create Form-layout container to be used for Name
	var testSuiteNameContainer *fyne.Container
	var testSuiteNameFormContainer *fyne.Container
	testSuiteNameContainer = container.New(layout.NewVBoxLayout())
	testSuiteNameFormContainer = container.New(layout.NewFormLayout())

	// Add Header to the Forms-container
	var headerLabel *widget.Label
	headerLabel = widget.NewLabel("TestSuiteName")
	headerLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuiteNameFormContainer.Add(headerLabel)

	// Add the Entry-widget for TestSuiteName
	newTestSuiteNameEntry := widget.NewEntry()
	newTestSuiteNameEntry.SetText(currentTestSuiteModel.TestSuiteUIModelBinding.TestSuiteName)

	// Change Name in model and Tab-name when UI-component-Name-value is changed
	newTestSuiteNameEntry.OnChanged = func(newValue string) {

		var trimmedValue string
		trimmedValue = strings.Trim(newValue, " ")

		// Save TestSuite back in Map
		currentTestSuiteModel.TestSuiteUIModelBinding.TestSuiteName = trimmedValue

		// Generate short version of UUID to put in TestSuite Tab-Name
		var shortUUid string
		var tabName string

		shortUUid = sharedCode.GenerateShortUuidFromFullUuid(testSuiteUuid)

		// Shorten Tab-name if name is longer then 'TestSuiteTabNameVisibleLength'
		if len(trimmedValue) > sharedCode.TestSuiteTabNameVisibleLength {
			tabName = trimmedValue[0:sharedCode.TestSuiteTabNameVisibleLength] + " [" + shortUUid + "] (*)"
		} else {
			tabName = trimmedValue + " [" + shortUUid + "] (*)"
		}

		testSuiteUiModel.TestSuiteTabItem.Text = tabName
		//Test.TestSuitesTabs.Refresh()

	}

	// Add the Entry-widget to the Forms-container
	testSuiteNameFormContainer.Add(newTestSuiteNameEntry)

	// Create the VBox-container that will be returned
	testSuiteNameContainer = container.NewVBox(testSuiteNameFormContainer)

	return testSuiteNameContainer, err
}
