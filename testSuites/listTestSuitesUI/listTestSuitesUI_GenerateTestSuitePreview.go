package listTestSuitesUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testCase/testCaseModel"
	"FenixTesterGui/testCases/listTestCasesUI"
	"FenixTesterGui/testSuites/listTestSuitesModel"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
)

// GenerateTestSuitePreviewContainer
// Generates the PreViewContainer for the TestSuite
func (listTestSuiteUIObject *ListTestSuiteUIStruct) GenerateTestSuitePreviewContainer(
	testSuiteUuid string,
	testCasesModel *testCaseModel.TestCasesModelsStruct) {

	var testSuitePreviewTopContainer *fyne.Container
	var testSuitePreviewBottomContainer *fyne.Container
	var testCasePreviewAreaForPreviewContainer *fyne.Container

	var tempTestSuitePreviewStructureMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestSuitePreviewStructureMessage

	// Get Data for the Preview
	tempTestSuitePreviewStructureMessage = listTestSuitesModel.TestSuitesThatCanBeEditedByUserMap[testSuiteUuid].TestSuitePreview.TestSuitePreview

	// Create the Top container
	testSuitePreviewTopContainer = container.New(layout.NewFormLayout())

	// Add TestSuiteName to Top container
	tempTestSuiteNameLabel := widget.NewLabel("TestSuiteName:")
	tempTestSuiteNameLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuitePreviewTopContainer.Add(tempTestSuiteNameLabel)
	testSuitePreviewTopContainer.Add(widget.NewLabel(tempTestSuitePreviewStructureMessage.GetTestSuiteName()))

	// Add TestSuiteOwner Domain Top container
	tempOwnerDomainLabel := widget.NewLabel("OwnerDomain:")
	tempOwnerDomainLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuitePreviewTopContainer.Add(tempOwnerDomainLabel)
	testSuitePreviewTopContainer.Add(widget.NewLabel(tempTestSuitePreviewStructureMessage.GetDomainNameThatOwnTheTestSuite()))

	// Add Description Top container
	tempTestSuiteDescription := widget.NewLabel("Description:")
	tempTestSuiteDescription.TextStyle = fyne.TextStyle{Bold: true}
	testSuitePreviewTopContainer.Add(tempTestSuiteDescription)
	testSuitePreviewTopContainer.Add(widget.NewRichTextWithText(tempTestSuitePreviewStructureMessage.GetTestSuiteDescription()))

	// Create Header for TestCase-list area
	testCaseListHeaderLabel := widget.NewLabel("TestCases in TestSuite")
	testCaseListHeaderLabel.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	testSuitePreviewTopContainer.Add(testCaseListHeaderLabel)
	testSuitePreviewTopContainer.Add(widget.NewLabel(""))

	// Create the Bottom container
	testSuitePreviewBottomContainer = container.New(layout.NewFormLayout())

	// Create Header for Bottom container
	bottomContainerHeaderLabel := widget.NewLabel("TestSuite Information")
	bottomContainerHeaderLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuitePreviewBottomContainer.Add(bottomContainerHeaderLabel)
	testSuitePreviewBottomContainer.Add(widget.NewLabel(""))

	// Add TestSuiteVersion to Bottom container
	tempTestSuiteVersionLabel := widget.NewLabel("TestSuiteVersion:")
	tempTestSuiteVersionLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuitePreviewBottomContainer.Add(tempTestSuiteVersionLabel)
	testSuitePreviewBottomContainer.Add(widget.NewLabel(tempTestSuitePreviewStructureMessage.GetTestSuiteVersion()))

	// Add LastSavedByUserOnComputer to Bottom container
	tempLastSavedByUserOnComputerLabel := widget.NewLabel("Last saved by user (on computer)::")
	tempLastSavedByUserOnComputerLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuitePreviewBottomContainer.Add(tempLastSavedByUserOnComputerLabel)
	testSuitePreviewBottomContainer.Add(widget.NewLabel(tempTestSuitePreviewStructureMessage.GetLastSavedByUserOnComputer()))

	// Add LastSavedByUserGCPAuthorization to Bottom container
	tempLastSavedByUserGCPAuthorizationLabel := widget.NewLabel("Last saved by GCP authenticated user:")
	tempLastSavedByUserGCPAuthorizationLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuitePreviewBottomContainer.Add(tempLastSavedByUserGCPAuthorizationLabel)
	testSuitePreviewBottomContainer.Add(widget.NewLabel(tempTestSuitePreviewStructureMessage.GetLastSavedByUserGCPAuthorization()))

	// Add LastSavedTimeStamp to Bottom container
	tempLastSavedTimeStamp := widget.NewLabel("Last saved TimeStamp:")
	tempLastSavedTimeStamp.TextStyle = fyne.TextStyle{Bold: true}
	testSuitePreviewBottomContainer.Add(tempLastSavedTimeStamp)
	testSuitePreviewBottomContainer.Add(widget.NewLabel(tempTestSuitePreviewStructureMessage.GetLastSavedTimeStamp()))

	// Create the Temporary container that should be shown
	temporaryContainer := container.NewCenter(widget.NewLabel("Select a TestCase to get the Preview"))
	testCasePreviewAreaForPreviewContainer = container.NewBorder(nil, nil, nil, nil, temporaryContainer)

	// Generate List with TestCases. User can click TestCase and its Preview is shown
	var testCasesWidgetList *widget.List

	// Create a short name to use in list
	testCasesList := tempTestSuitePreviewStructureMessage.TestCasesInTestSuite.TestCasesInTestSuite

	// Create and configure the list-component of all TestDataPoints
	testCasesWidgetList = widget.NewList(
		func() int { return len(testCasesList) },
		func() fyne.CanvasObject {

			return widget.NewLabel("")
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {

			obj.(*widget.Label).SetText(fmt.Sprintf(
				"%s [%s]",
				testCasesList[id].TestCaseName,
				sharedCode.GenerateShortUuidFromFullUuid(testCasesList[id].TestCaseUuid)))
		},
	)

	// When TestCase is selected then show Preview
	testCasesWidgetList.OnSelected = func(id widget.ListItemID) {

		listTestCasesUI.StandardListTesCasesUIObject.GenerateTestCasePreviewContainer(
			testCasesList[id].TestCaseUuid,
			testCasePreviewAreaForPreviewContainer,
			testCasesModel)

	}

	// Calculate a sensible minimum height for ~5 rows
	template := widget.NewLabel("template")
	rowH := template.MinSize().Height
	rows := float32(5.0)
	minH := rowH*rows + theme.Padding()*(rows-1)

	// Wrap the list so it can't shrink below ~5 rows
	wrap := container.New(layout.NewGridWrapLayout(fyne.NewSize(320, minH)), testCasesWidgetList)

	testSuiteMainAreaForPreviewBorderContainer := container.NewBorder(wrap, nil, nil, nil, testCasePreviewAreaForPreviewContainer)
	testSuiteMainAreaForPreviewScrollContainer := container.NewScroll(testSuiteMainAreaForPreviewBorderContainer)

	// Create Top header for Preview
	tempTopHeaderLabel := widget.NewLabel("TestSuite Preview")
	tempTopHeaderLabel.TextStyle = fyne.TextStyle{Bold: true}

	listTestSuiteUIObject.testSuitePreviewContainer.Objects[0] = container.NewBorder(
		container.NewVBox(container.NewCenter(tempTopHeaderLabel), testSuitePreviewTopContainer, widget.NewSeparator()),
		container.NewVBox(widget.NewSeparator(), testSuitePreviewBottomContainer), nil, nil,
		testSuiteMainAreaForPreviewScrollContainer)

	// Refresh the 'testSuitePreviewContainer'
	listTestSuiteUIObject.testSuitePreviewContainer.Refresh()

}
