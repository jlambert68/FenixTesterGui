package testSuiteUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Generate the TestSuiteInformation Area for the TestSuite
// Uuid, Created By, Created Date, Last Changed Date, Last Changed By
func (testSuiteUiModel *TestSuiteUiStruct) generateTestSuiteInformationFieldsArea() (
	testSuiteInformationScroll *container.Scroll,
	err error) {

	// Create container to be used for the Information parts
	var testSuiteInformationContainer *fyne.Container
	testSuiteInformationContainer = container.New(layout.NewHBoxLayout())

	// Add "Uuid" to the Information container ***********************
	var uuidLabel *widget.Label
	uuidLabel = widget.NewLabel("Uuid:")
	uuidLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuiteInformationContainer.Add(uuidLabel)

	var uuidCopyableLabel *copyableLabelStruct
	uuidCopyableLabel = newCopyableLabel(
		testSuiteUiModel.TestSuiteModelPtr.GetTestSuiteUuid(), true, testSuiteUiModel)
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
	createdByGcpLoginLabelCopyableLabel = newCopyableLabel(
		testSuiteUiModel.TestSuiteModelPtr.GetCreatedByGcpLogin(), true, testSuiteUiModel)
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
	createdByComputerLoginCopyableLabel = newCopyableLabel(
		testSuiteUiModel.TestSuiteModelPtr.GetCreatedByComputerLogin(), true, testSuiteUiModel)
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
	createdDateCopyableLabel = newCopyableLabel(
		testSuiteUiModel.TestSuiteModelPtr.GetCreatedDate(), true, testSuiteUiModel)
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
	lastChangedByGcpLoginCopyableLabel = newCopyableLabel(
		testSuiteUiModel.TestSuiteModelPtr.GetLastChangedByGcpLogin(), true, testSuiteUiModel)
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
	lastChangedByComputerLoginLabelCopyableLabel = newCopyableLabel(
		testSuiteUiModel.TestSuiteModelPtr.GetLastChangedByComputerLogin(), true, testSuiteUiModel)
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
	lastChangeDateCopyableLabel = newCopyableLabel(
		testSuiteUiModel.TestSuiteModelPtr.GetLastChangedDate(), true, testSuiteUiModel)
	lastChangeDateCopyableLabel.TextStyle = fyne.TextStyle{Italic: true}
	testSuiteInformationContainer.Add(lastChangeDateCopyableLabel)

	// Put forms container in VBox-container
	var testSuiteInformationVContainer *fyne.Container
	testSuiteInformationVContainer = container.NewVBox(testSuiteInformationContainer, widget.NewLabel(""), widget.NewSeparator())

	// Put the container in the Scroll container
	testSuiteInformationScroll = container.NewHScroll(testSuiteInformationVContainer)

	return testSuiteInformationScroll, err
}
