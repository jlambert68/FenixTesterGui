package testCaseUI

import (
	"FenixTesterGui/testCase/testCaseModel"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

// TestCaseUiStruct
// This structure holds all UI-objects for all the TestCases
type TestCasesObjectForUiStruct struct {
	TestCaseToolUIBar       *widget.Toolbar                          // Toolbar used copy, cut, paste Building Blocks
	TestCasesTabs           *container.AppTabs                       // The Tab-structure where each TestCase has its own Tab
	TestCasesUiModelMap     map[string]*testCaseGraphicalAreasStruct // Holds all UI sub-parts for a TestCase
	TestCasesModelReference *testCaseModel.TestCasesModelsStruct     // A reference to the model for all TestCases

}

// This structure holds the UI-objects for one TestCase
type testCaseGraphicalAreasStruct struct {
	currentTestCaseTextualStructure currentTestCaseTextualStructureStruct

	TestCaseTextualModelArea    *fyne.CanvasObject
	TestCaseGraphicalModelArea  *fyne.CanvasObject
	TestCaseBaseInformationArea *fyne.CanvasObject
	TestCaseMetaDataArea        *fyne.CanvasObject
	TestCaseAttributesArea      *fyne.CanvasObject
}

// Keeps track of the latest Textual Representations for the TestCase
type currentTestCaseTextualStructureStruct struct {
	currentTestCaseTextualStructureSimple   binding.String
	currentTestCaseTextualStructureComplex  binding.String
	currentTestCaseTextualStructureExtended binding.String
}
