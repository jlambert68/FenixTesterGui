package listTestSuiteExecutionsUI

import (
	"FenixTesterGui/testCaseExecutions/listTestCaseExecutionsUI"
	"FenixTesterGui/testCaseExecutions/testCaseExecutionsModel"
	"FenixTesterGui/testSuiteExecutions/testSuiteExecutionsModel"

	"fyne.io/fyne/v2"
)

func (testSuiteInstructionPreViewObjectRef *TestSuiteInstructionPreViewStruct) GenerateTestSuiteExecutionPreviewContainer(
	testSuiteExecutionUuid string,
	testSuiteExecutionVersion uint32,
	testSuiteExecutionsModelRef *testSuiteExecutionsModel.TestSuiteExecutionsModelStruct,
	openedTestSuiteExecutionFrom listTestCaseExecutionsUI.OpenedTestCaseExecutionOrTestSuiteExecutionFromType,
	currentWindowPtr *fyne.Window,
	testCaseExecutionsModelRef *testCaseExecutionsModel.TestCaseExecutionsModelStruct) {

}

// ClearTestSuiteExecutionPreviewContainer
// Clears the preview area for the TestSuiteExecution
func (testCaseInstructionPreViewObjectRef *TestSuiteInstructionPreViewStruct) ClearTestSuiteExecutionPreviewContainer() {
	//	testCaseInstructionPreViewObjectRef.testSuiteExecutionPreviewContainer.Objects[0] = container.NewCenter(widget.NewLabel("Select a TestSuiteExecution to get the Preview"))
}
