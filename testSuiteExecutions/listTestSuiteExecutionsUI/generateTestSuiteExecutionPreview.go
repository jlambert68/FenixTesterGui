package listTestSuiteExecutionsUI

import (
	"FenixTesterGui/testSuiteExecutions/testSuiteExecutionsModel"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (testSuiteInstructionPreViewObjectRef *TestSuiteInstructionPreViewStruct) GenerateTestSuiteExecutionPreviewContainer(
	testSuiteExecutionUuid string,
	testSuiteExecutionVersion uint32,
	testSuiteExecutionsModelRef *testSuiteExecutionsModel.TestSuiteExecutionsModelStruct,
	openedTestSuiteExecutionFrom openedTestSuiteExecutionFromType,
	currentWindowPtr *fyne.Window) {

	return
}

// ClearTestSuiteExecutionPreviewContainer
// Clears the preview area for the TestSuiteExecution
func (testCaseInstructionPreViewObjectRef *TestSuiteInstructionPreViewStruct) ClearTestSuiteExecutionPreviewContainer() {
	testCaseInstructionPreViewObjectRef.testSuiteExecutionPreviewContainer.Objects[0] = container.NewCenter(widget.NewLabel("Select a TestSuiteExecution to get the Preview"))
}
