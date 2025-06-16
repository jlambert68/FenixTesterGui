package testSuiteUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testSuites/testSuitesModel"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
)

// Generate the TestSuiteInformation Area for the TestSuite
// Uuid, Created By, Created Date, Last Changed Date, Last Changed By
func (testSuiteUiModel *TestSuiteUiStruct) generateTestSuiteInformationFieldsArea(
	testSuiteUuid string) (
	testSuiteInformationStackContainer *fyne.Container,
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
			"ID":            "48285fad-09a3-4e52-8f34-a104check358a",
			"testSuiteUuid": testSuiteUuid,
		}).Fatal("TestSuite doesn't exist in TestSuiteMap. This should not happen")
	}
	currentTestSuiteModel = *currentTestSuiteModelPtr

	// Create container to be used for the Information parts
	var testSuiteInformationContainer *fyne.Container
	testSuiteInformationContainer = container.New(layout.NewHBoxLayout())

	// Add "Uuid" to the Information container ***********************
	var uuidLabel *widget.Label
	uuidLabel = widget.NewLabel("Uuid:")
	uuidLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuiteInformationContainer.Add(uuidLabel)

	var uuidCopyableLabel *copyableLabelStruct
	uuidCopyableLabel = newCopyableLabel(currentTestSuiteModel.GetTestSuiteUuid(), true)
	uuidCopyableLabel.TextStyle = fyne.TextStyle{Italic: true}
	testSuiteInformationContainer.Add(uuidCopyableLabel)

	// Add "space" between Information parts
	testSuiteInformationContainer.Add(widget.NewLabel(spaceBetweenInformationBoxes))

	// Add "Created By Gcp-Login" to the Information container ***********************
	var createdByGcpLoginLabel *widget.Label
	createdByGcpLoginLabel = widget.NewLabel("Created By - GCP Login:")
	createdByGcpLoginLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuiteInformationContainer.Add(createdByGcpLoginLabel)

	var createdByGcpLoginLabelCopyableLabel *copyableLabelStruct
	createdByGcpLoginLabelCopyableLabel = newCopyableLabel(currentTestSuiteModel.GetCreatedByGcpLogin(), true)
	createdByGcpLoginLabelCopyableLabel.TextStyle = fyne.TextStyle{Italic: true}
	testSuiteInformationContainer.Add(createdByGcpLoginLabelCopyableLabel)

	// Add "space" between Information parts
	testSuiteInformationContainer.Add(widget.NewLabel(spaceBetweenInformationBoxes))

	// Add "Created By Computer Login" to the Information container ***********************
	var createdByComputerLoginLabel *widget.Label
	createdByComputerLoginLabel = widget.NewLabel("Created By - Computer Login:")
	createdByComputerLoginLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuiteInformationContainer.Add(createdByComputerLoginLabel)

	var createdByComputerLoginCopyableLabel *copyableLabelStruct
	createdByComputerLoginCopyableLabel = newCopyableLabel(currentTestSuiteModel.GetCreatedByComputerLogin(), true)
	createdByComputerLoginCopyableLabel.TextStyle = fyne.TextStyle{Italic: true}
	testSuiteInformationContainer.Add(createdByComputerLoginCopyableLabel)

	// Add "space" between Information parts
	testSuiteInformationContainer.Add(widget.NewLabel(spaceBetweenInformationBoxes))

	// Add "Created Date" to the Information container ***********************
	var createdDateLabel *widget.Label
	createdDateLabel = widget.NewLabel("Created Date:")
	createdDateLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuiteInformationContainer.Add(createdDateLabel)

	var createdDateCopyableLabel *copyableLabelStruct
	createdDateCopyableLabel = newCopyableLabel(currentTestSuiteModel.GetCreatedDate(), true)
	createdDateCopyableLabel.TextStyle = fyne.TextStyle{Italic: true}
	testSuiteInformationContainer.Add(createdDateCopyableLabel)

	// Add "space" between Information parts
	testSuiteInformationContainer.Add(widget.NewLabel(spaceBetweenInformationBoxes))

	// Add "Last Changed  By Gcp-Login" to the Information container ***********************
	var lastChangedByGcpLoginLabel *widget.Label
	lastChangedByGcpLoginLabel = widget.NewLabel("Last Changed By - GCP Login:")
	lastChangedByGcpLoginLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuiteInformationContainer.Add(lastChangedByGcpLoginLabel)

	var lastChangedByGcpLoginCopyableLabel *copyableLabelStruct
	lastChangedByGcpLoginCopyableLabel = newCopyableLabel(currentTestSuiteModel.GetLastChangedByGcpLogin(), true)
	lastChangedByGcpLoginCopyableLabel.TextStyle = fyne.TextStyle{Italic: true}
	testSuiteInformationContainer.Add(lastChangedByGcpLoginCopyableLabel)

	// Add "space" between Information parts
	testSuiteInformationContainer.Add(widget.NewLabel(spaceBetweenInformationBoxes))

	// Add "Last Changed By Computer Login" to the Information container ***********************
	var lastChangedByComputerLoginLabel *widget.Label
	lastChangedByComputerLoginLabel = widget.NewLabel("Last Changed By - Computer Login:")
	lastChangedByComputerLoginLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuiteInformationContainer.Add(lastChangedByComputerLoginLabel)

	var lastChangedByComputerLoginLabelCopyableLabel *copyableLabelStruct
	lastChangedByComputerLoginLabelCopyableLabel = newCopyableLabel(currentTestSuiteModel.GetLastChangedByComputerLogin(), true)
	lastChangedByComputerLoginLabelCopyableLabel.TextStyle = fyne.TextStyle{Italic: true}
	testSuiteInformationContainer.Add(lastChangedByComputerLoginLabelCopyableLabel)

	// Add "space" between Information parts
	testSuiteInformationContainer.Add(widget.NewLabel(spaceBetweenInformationBoxes))

	// Add "Last Changed  Date" to the Information container ***********************
	var lastChangeDateLabel *widget.Label
	lastChangeDateLabel = widget.NewLabel("Last Change Date:")
	lastChangeDateLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuiteInformationContainer.Add(lastChangeDateLabel)

	var lastChangeDateCopyableLabel *copyableLabelStruct
	lastChangeDateCopyableLabel = newCopyableLabel(currentTestSuiteModel.GetLastChangedDate(), true)
	lastChangeDateCopyableLabel.TextStyle = fyne.TextStyle{Italic: true}
	testSuiteInformationContainer.Add(lastChangeDateCopyableLabel)

	// Put the container in the Scroll container
	var testSuiteInformationScroll *container.Scroll
	testSuiteInformationScroll = container.NewScroll(testSuiteInformationContainer)

	testSuiteInformationStackContainer = container.NewStack(testSuiteInformationScroll)

	return testSuiteInformationStackContainer, err
}
