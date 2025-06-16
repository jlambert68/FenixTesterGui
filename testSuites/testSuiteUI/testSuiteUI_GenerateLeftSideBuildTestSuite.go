package testSuiteUI

import (
	"FenixTesterGui/testCase/testCaseModel"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Generate leftSideBuildTestSuite - Main information for TestSuite
func (testSuiteUiModel TestSuiteUiStruct) generateLeftSideBuildTestSuiteContainer(
	testSuiteUuid string,
	testCasesModel *testCaseModel.TestCasesModelsStruct) (
	leftSideBuildTestSuiteContainer *fyne.Container,
	err error) {

	var leftTopSideBuildTestSuiteContainer *fyne.Container

	leftTopSideBuildTestSuiteContainer = container.NewVBox()

	var testSuiteDeleteDateAreaContainer *fyne.Container
	var testSuiteNameAreaContainer *fyne.Container
	var testSuiteDescriptionAreaContainer *fyne.Container
	var testSuiteInformationAreaContainer *fyne.Container
	var testSuiteOwnerDomainContainer *fyne.Container
	var testSuiteTestEnvironmentContainer *fyne.Container
	var testSuiteTestDataAreaContainer *fyne.Container

	// Generate TestSuite-DeleteDate area
	testSuiteDeleteDateAreaContainer, err = testSuiteUiModel.generateTestSuiteDeletionDateArea(testSuiteUuid)
	if err != nil {

		errorId := "8f46fe50-51c8-411c-ad81-30d036083e8f"
		errorMessage := fmt.Sprintf("couldn't generate 'TestSuite-DeleteDate-area, err=%s. [ErrorId = %s]",
			err.Error(),
			errorId)

		leftSideBuildTestSuiteContainer = container.NewVBox(widget.NewLabel(errorMessage))

		return leftSideBuildTestSuiteContainer, nil

	}
	leftTopSideBuildTestSuiteContainer.Add(testSuiteDeleteDateAreaContainer)

	// Generate TestSuite-Name area
	testSuiteNameAreaContainer, err = testSuiteUiModel.generateTestSuiteNameArea(testSuiteUuid)
	if err != nil {

		errorId := "b965b768-3eab-49d0-bc24-c33e2a85e7fe"
		errorMessage := fmt.Sprintf("couldn't generate 'TestSuite-name-area, err=%s. [ErrorId = %s]",
			err.Error(),
			errorId)

		leftSideBuildTestSuiteContainer = container.NewVBox(widget.NewLabel(errorMessage))

		return leftSideBuildTestSuiteContainer, nil

	}
	leftTopSideBuildTestSuiteContainer.Add(testSuiteNameAreaContainer)

	// Generate TestSuite-Description-area
	testSuiteDescriptionAreaContainer, err = testSuiteUiModel.generateTestCaseDescriptionArea(testSuiteUuid)
	if err != nil {

		errorId := "84aae5d4-d081-48f6-aee9-c25bdae7ddd4"
		errorMessage := fmt.Sprintf("couldn't generate 'TestSuite-Description-area, err=%s. [ErrorId = %s]",
			err.Error(),
			errorId)

		leftSideBuildTestSuiteContainer = container.NewVBox(widget.NewLabel(errorMessage))

		return leftSideBuildTestSuiteContainer, nil

	}
	leftTopSideBuildTestSuiteContainer.Add(testSuiteDescriptionAreaContainer)

	// Generate TestSuite-Information-area
	testSuiteInformationAreaContainer, err = testSuiteUiModel.generateTestSuiteInformationFieldsArea(testSuiteUuid)
	if err != nil {

		errorId := "a3fa30d8-df0e-408c-8e4c-99aabb62cfab"
		errorMessage := fmt.Sprintf("couldn't generate 'TestSuite-Information-area, err=%s. [ErrorId = %s]",
			err.Error(),
			errorId)

		leftSideBuildTestSuiteContainer = container.NewVBox(widget.NewLabel(errorMessage))

		return leftSideBuildTestSuiteContainer, nil

	}
	leftTopSideBuildTestSuiteContainer.Add(testSuiteInformationAreaContainer)

	// Generate TestSuite-Owner area
	var testCaseOwnerDomainCustomSelectComboBox *customSelectComboBox
	testSuiteOwnerDomainContainer, testCaseOwnerDomainCustomSelectComboBox, err = testSuiteUiModel.generateOwnerDomainForTestSuiteArea(
		testSuiteUuid, testCasesModel)
	if err != nil {

		errorId := "46cab305-6172-4e06-aa19-3c891c5b8d57"
		errorMessage := fmt.Sprintf("couldn't generate 'TestSuite-Owner area', err=%s. [ErrorId = %s]",
			err.Error(),
			errorId)

		leftSideBuildTestSuiteContainer = container.NewVBox(widget.NewLabel(errorMessage))

		return leftSideBuildTestSuiteContainer, nil

	}
	leftTopSideBuildTestSuiteContainer.Add(testSuiteOwnerDomainContainer)
	fmt.Println(testCaseOwnerDomainCustomSelectComboBox)

	// Generate TestSuite-Owner and Execution environment-area
	var customTestEnvironmentSelectComboBox *customSelectComboBox
	testSuiteTestEnvironmentContainer, customTestEnvironmentSelectComboBox, err = testSuiteUiModel.
		generateTestEnvironmentForTestSuite(
			testSuiteUuid, testCasesModel)
	if err != nil {

		errorId := "4f059e29-1e5c-46e1-a766-f2f95d7a5c36"
		errorMessage := fmt.Sprintf("couldn't generate 'TestSuites TestExecution environment-area', err=%s. [ErrorId = %s]",
			err.Error(),
			errorId)

		leftSideBuildTestSuiteContainer = container.NewVBox(widget.NewLabel(errorMessage))

		return leftSideBuildTestSuiteContainer, nil

	}
	leftTopSideBuildTestSuiteContainer.Add(testSuiteTestEnvironmentContainer)
	fmt.Println(customTestEnvironmentSelectComboBox)

	// Generate TestSuite-TestData-area
	testSuiteTestDataAreaContainer, err = testSuiteUiModel.generateSelectedTestDataForTestSuiteArea(testSuiteUuid)
	if err != nil {

		errorId := "d1f4abd1-95bb-4935-9403-8fe4cc360a57"
		errorMessage := fmt.Sprintf("couldn't generate 'TestSuite-TestData-area, err=%s. [ErrorId = %s]",
			err.Error(),
			errorId)

		leftSideBuildTestSuiteContainer = container.NewVBox(widget.NewLabel(errorMessage))

		return leftSideBuildTestSuiteContainer, nil

	}
	leftTopSideBuildTestSuiteContainer.Add(testSuiteTestDataAreaContainer)

	// Create the Left side Container
	leftSideBuildTestSuiteContainer = container.NewBorder(
		leftTopSideBuildTestSuiteContainer,
		nil,
		nil,
		nil,
		nil)

	return leftSideBuildTestSuiteContainer, err

}

/*
// Build top part of the build TestSuite
func generateLeftTopSideBuildTestSuiteContainer
	testCasesModel *testCaseModel.TestCasesModelsStruct) (
	leftTopSideBuildTestSuiteContainer *fyne.Container) {

	return leftTopSideBuildTestSuiteContainer
}


*/
