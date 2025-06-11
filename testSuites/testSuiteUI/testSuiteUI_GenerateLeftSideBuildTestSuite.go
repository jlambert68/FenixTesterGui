package testSuiteUI

import (
	"FenixTesterGui/testCase/testCaseModel"
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

	leftSideBuildTestSuiteContainer = container.NewVBox(widget.NewLabel("leftSideBuildTestSuiteContainer"))

	/*

		var leftSideScrollContainer *fyne.Container
		var leftSideTopScrollContainer *fyne.Container
		var leftSideBottomScrollContainer *fyne.Container

		// Generate Delete TestSuite objec
		var testSuiteDeletionDateAreaContainer *fyne.Container
		testSuiteDeletionDateAreaContainer, err = testSuiteUiObject.generateTestSuiteDeletionDateArea(testSuiteUuid)

		if err != nil {

			errorId := "803beea6-a8db-420d-bf61-a1c16cb6fcad"

			leftSideBuildTestSuiteContainer = container.NewVBox(
				widget.NewLabel(fmt.Sprintf("couldn't generate 'Suite-delete-area, err=%s. [ErrorId = %s]",
					err.Error(),
					errorId)))




		}

	*/

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
