package testSuiteUI

import (
	"FenixTesterGui/testCase/testCaseModel"
	"FenixTesterGui/testSuites/testSuitesModel"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

// GenerateBuildNewTestSuiteUI
// Create the UI used for creating new TestSuites
func (testSuiteUiModel TestSuiteUiStruct) GenerateBuildNewTestSuiteUI(
	testCasesModel *testCaseModel.TestCasesModelsStruct,
	newTestSuiteModel *testSuitesModel.TestSuiteModelStruct) (
	newTestSuiteUIContainer *fyne.Container,
	err error) {

	// Get the new TestSuite-UUID
	var testSuiteUuid string
	testSuiteUuid = newTestSuiteModel.GetTestSuiteUuid()

	// Check if TestSuitesModel needs to be initiated
	if testSuitesModel.TestSuitesModelPtr == nil {

		// Initiate 'TestSuitesMap'
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
	existingTestSuitesMap[testSuiteUuid] = newTestSuiteModel

	var buildTestSuiteSplitContainer *container.Split
	var leftSideBuildTestSuiteContainer *fyne.Container
	var rightSideBuildTestSuiteContainer *fyne.Container

	// Generate rightSideBuildTestSuiteContainer - MetaData filter for TestCases
	rightSideBuildTestSuiteContainer,
		testSuiteUiModel.preViewAndFilterTabsUsedForCreateTestSuite = generateRightSideBuildTestSuiteContainer(testCasesModel)

	// Generate leftSideBuildTestSuite - Main information for TestSuite
	leftSideBuildTestSuiteContainer, err = testSuiteUiModel.generateLeftSideBuildTestSuiteContainer(
		testSuiteUuid,
		testCasesModel,
		testSuiteUiModel.preViewAndFilterTabsUsedForCreateTestSuite)
	if err != nil {
		newTestSuiteUIContainer = container.NewVBox(widget.NewLabel(err.Error()))
		return newTestSuiteUIContainer, err

	}

	// make, LeftSide, hoverable transparent overlay, to stop mouse interference between to two sides of the split-container
	leftCreateTestSuiteOverlay := NewHoverableRect(color.Transparent, nil)
	leftCreateTestSuiteOverlay.OnMouseIn = func(ev *desktop.MouseEvent) {

		testSuiteUiModel.mouseHasLeftSideSplitContainer = false
		leftCreateTestSuiteOverlay.Hide()
		leftCreateTestSuiteOverlay.OtherHoverableRect.Show()
	}
	leftCreateTestSuiteOverlay.OnMouseOut = func() {

	}

	// make, RightSide, hoverable transparent overlay, to stop mouse interference between to two sides of the split-container
	rightCreateTestSuiteOverlay := NewHoverableRect(color.Transparent, nil)
	rightCreateTestSuiteOverlay.OnMouseIn = func(ev *desktop.MouseEvent) {

		testSuiteUiModel.mouseHasLeftSideSplitContainer = false
		rightCreateTestSuiteOverlay.Hide()
		rightCreateTestSuiteOverlay.OtherHoverableRect.Show()
	}
	rightCreateTestSuiteOverlay.OnMouseOut = func() {

	}

	// Cross connect the two overlays
	leftCreateTestSuiteOverlay.OtherHoverableRect = rightCreateTestSuiteOverlay
	rightCreateTestSuiteOverlay.OtherHoverableRect = leftCreateTestSuiteOverlay

	buildTestSuiteAndOverlayContainer := container.New(layout.NewStackLayout(),
		leftSideBuildTestSuiteContainer, leftCreateTestSuiteOverlay)
	rightSideAndOverlayContainer := container.New(layout.NewStackLayout(),
		rightSideBuildTestSuiteContainer, rightCreateTestSuiteOverlay)

	buildTestSuiteSplitContainer = container.NewHSplit(buildTestSuiteAndOverlayContainer, rightSideAndOverlayContainer)
	newTestSuiteUIContainer = container.NewBorder(nil, nil, nil, nil, buildTestSuiteSplitContainer)

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
