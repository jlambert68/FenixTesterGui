package testCaseUI

import (
	"FenixTesterGui/commandAndRuleEngine"
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testCase/testCaseModel"
	"FenixTesterGui/testUIDragNDropStatemachine"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

const testCaseNodeRectangleSize = 40

// TestCaseUiStruct
// This structure holds all UI-objects for all the TestCases
type TestCasesUiModelStruct struct {
	TestCaseToolUIBar       *widget.Toolbar                          // Toolbar used copy, cut, paste Building Blocks
	TestCasesTabs           *container.AppTabs                       // The Tab-structure where each TestCase has its own Tab
	TestCasesUiModelMap     map[string]*testCaseGraphicalAreasStruct // Holds all UI sub-parts for a TestCase
	TestCasesModelReference *testCaseModel.TestCasesModelsStruct     // A reference to the model for all TestCases

	DragNDropObject       DragNDropObjectStruct // The object used for Dragging TI and TIC into the TestCase
	DragNDropStateMachine testUIDragNDropStatemachine.StateMachineDragAndDropStruct

	//TODO - Remove, used for Testing only
	DragNDropText                    *canvas.Text // Text used for Drag n Drop of TI and TIC into TextCase //TODO, is this used?
	DragNDropRectangle               *canvas.Rectangle
	DragNDropRectangleTextBackground *canvas.Rectangle
	DragNDropContainer               *fyne.Container

	CommandAndRuleEngineReference *commandAndRuleEngine.CommandAndRuleEngineObjectStruct

	CommandChannelReference        *sharedCode.CommandChannelType
	GraphicsUpdateChannelReference *sharedCode.CommandChannelGraphicsUpdateType
}

// This object hold references to the object that is used when user drags a new TestInstruction or
// TestInstructionContainer from Available Building Blocks and drops it into the TestCase
type DragNDropObjectStruct struct {
	DragNDropText                    *canvas.Text      // Text used for Drag n Drop of TI and TIC into TextCase //TODO, is this used?
	DragNDropRectangle               *canvas.Rectangle // Outer rectangle
	DragNDropRectangleTextBackground *canvas.Rectangle // The text background for text to be more readable, regarding text and background colors
	DragNDropContainer               *fyne.Container   // The Container holding the above objects together
}

// This structure holds the UI-objects for one TestCase
type testCaseGraphicalAreasStruct struct {
	currentTestCaseTextualStructure   testCaseTextualStructureStruct   // Keeps track of the latest Textual Representations for the TestCase
	currentTestCaseGraphicalStructure testCaseGraphicalStructureStruct // Keeps track of important objects for Graphical Representation for TestCase

	TestCaseTextualModelArea    fyne.CanvasObject
	TestCaseGraphicalModelArea  fyne.CanvasObject
	TestCaseBaseInformationArea fyne.CanvasObject
	TestCaseMetaDataArea        fyne.CanvasObject
	TestCaseAttributesArea      fyne.CanvasObject
}

// Keeps track of the latest Textual Representations for the TestCase
type testCaseTextualStructureStruct struct {
	currentTestCaseTextualStructureSimple   binding.String
	currentTestCaseTextualStructureComplex  binding.String
	currentTestCaseTextualStructureExtended binding.String
	currentTestCaseGraphicalAccordionObject *widget.Accordion
}

// Keeps track of important object for the Graphical Representations for the TestCase
type testCaseGraphicalStructureStruct struct {
	currentTestCaseGraphicalModelArea       fyne.CanvasObject
	currentTestCaseGraphicalTreeComponent   *widget.Tree
	currentTestCaseGraphicalObject          *fyne.CanvasObject
	currentTestCaseGraphicalAccordionObject *widget.Accordion
}
