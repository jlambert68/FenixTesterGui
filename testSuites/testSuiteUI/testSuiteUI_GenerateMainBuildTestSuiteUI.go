package testSuiteUI

import (
	"FenixTesterGui/testCase/testCaseModel"
	"FenixTesterGui/testSuites/testSuitesModel"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	uuidGenerator "github.com/google/uuid"
)

// GenerateBuildNewTestSuiteUI
// Create the UI used for creating new TestSuites
func (testSuiteUiModel TestSuiteUiStruct) GenerateBuildNewTestSuiteUI(
	testCasesModel *testCaseModel.TestCasesModelsStruct) (
	newTestSuiteUIContainer *fyne.Container,
	err error) {

	// Generate new TestSuite-UUID
	var testSuiteUuid string
	testSuiteUuid = uuidGenerator.New().String()

	// Generate the new TestSuiteModelStruct
	var newTestSuiteModel testSuitesModel.TestSuiteModelStruct
	newTestSuiteModel = testSuitesModel.TestSuiteModelStruct{
		TestSuiteUIModelBinding: testSuitesModel.TestSuiteUIModelBindingStruct{
			TestSuiteDeletionDate:    "",
			TestSuiteName:            "",
			TestSuiteDescription:     "",
			TestSuiteOwnerDomainUuid: "",
			TestSuiteIsNew:           true,
		},
	}

	// Check if TestSuitesModel needs to be initiated
	if testSuitesModel.TestSuitesModelPtr == nil {

		// Initate 'TestSuitesMap'
		var tempTestSuitesMap map[string]*testSuitesModel.TestSuiteModelStruct
		tempTestSuitesMap = make(map[string]*testSuitesModel.TestSuiteModelStruct)

		var tempTestSuitesModel testSuitesModel.TestSuitesModelStruct
		tempTestSuitesModel = testSuitesModel.TestSuitesModelStruct{
			TestSuitesMapPtr: &tempTestSuitesMap}

		// Store the initiated Object
		testSuitesModel.TestSuitesModelPtr = &tempTestSuitesModel
	}

	var existingTestSuitesMap map[string]*testSuitesModel.TestSuiteModelStruct
	existingTestSuitesMap = *testSuitesModel.TestSuitesModelPtr.TestSuitesMapPtr

	// Save new TestSuiteModel in TestSuitesModel
	existingTestSuitesMap[testSuiteUuid] = &newTestSuiteModel

	var buildTestSuiteSplitContainer *container.Split
	var leftSideBuildTestSuiteContainer *fyne.Container
	var rightSideBuildTestSuiteContainer *fyne.Container

	// Generate leftSideBuildTestSuite - Main information for TestSuite
	leftSideBuildTestSuiteContainer, err = testSuiteUiModel.generateLeftSideBuildTestSuiteContainer(
		testSuiteUuid,
		testCasesModel)
	if err != nil {
		newTestSuiteUIContainer = container.NewVBox(widget.NewLabel(err.Error()))
		return newTestSuiteUIContainer, err

	}

	// Generate rightSideBuildTestSuiteContainer - MetaData filter for TestCases
	rightSideBuildTestSuiteContainer = generateRightSideBuildTestSuiteContainer(testCasesModel)

	buildTestSuiteSplitContainer = container.NewHSplit(leftSideBuildTestSuiteContainer, rightSideBuildTestSuiteContainer)
	newTestSuiteUIContainer = container.NewVBox(buildTestSuiteSplitContainer)

	return newTestSuiteUIContainer, err
}

/*
// GenerateBuildExistingTestSuiteUI
// Create the UI used for creating an existing TestSuite
func GenerateBuildExistingTestSuiteUI(
	testSuiteUuid string,
	testCasesModel *testCaseModel.TestCasesModelsStruct) (
	_ fyne.CanvasObject) {

	var buildTestSuiteSplitContainer *container.Split
	var leftSideBuildTestSuiteContainer *fyne.Container
	var rightSideBuildTestSuiteContainer *fyne.Container

	// Generate a new TestSuiteUI-object
	var newTestSuiteUiObject *TestSuiteUiStruct
	newTestSuiteUiObject = &TestSuiteUiStruct{
		testSuiteTabItem:  nil,
		TestSuiteModelPtr: nil,
	}

	// Generate leftSideBuildTestSuite - Main information for TestSuite
	leftSideBuildTestSuiteContainer = generateLeftSideBuildTestSuiteContainer(testCasesModel)

	// Generate rightSideBuildTestSuiteContainer - MetaData filter for TestCases
	rightSideBuildTestSuiteContainer = generateRightSideBuildTestSuiteContainer(testCasesModel)

	buildTestSuiteSplitContainer = container.NewHSplit(leftSideBuildTestSuiteContainer, rightSideBuildTestSuiteContainer)

	return buildTestSuiteSplitContainer
}


*/
