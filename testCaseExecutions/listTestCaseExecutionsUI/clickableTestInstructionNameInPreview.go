package listTestCaseExecutionsUI

import (
	"FenixTesterGui/testCaseExecutions/testCaseExecutionsModel"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Definition of a clickable TestInstructionName label used in the TestCaseExecution-Preview
type clickableTestInstructionNameLabelInPreviewStruct struct {
	widget.Label
	TestInstructionExecutionKey testCaseExecutionsModel.
					TestInstructionExecutionAttributesContainerMapKeyType
	LeftClicked  func()
	RightClicked func()
}

// Used for creating a new TestInstructionName label
func newClickableTestInstructionNameLabelInPreview(
	testInstructionName string,
	testInstructionExecutionKey testCaseExecutionsModel.
		TestInstructionExecutionAttributesContainerMapKeyType,
	leftClicked func(),
	rightClicked func(),
) *clickableTestInstructionNameLabelInPreviewStruct {

	clickableTestInstructionNameLabelInPreview := &clickableTestInstructionNameLabelInPreviewStruct{
		Label:                       widget.Label{Text: testInstructionName},
		TestInstructionExecutionKey: testInstructionExecutionKey,
		LeftClicked:                 leftClicked,
		RightClicked:                rightClicked,
	}

	clickableTestInstructionNameLabelInPreview.ExtendBaseWidget(clickableTestInstructionNameLabelInPreview)

	return clickableTestInstructionNameLabelInPreview
}

// Renderer (required by fyne.Widget)
func (c *clickableTestInstructionNameLabelInPreviewStruct) CreateRenderer() fyne.WidgetRenderer {
	lbl := widget.NewLabel(c.Label.Text)
	return widget.NewSimpleRenderer(lbl)
}

// Tapped interface implementation
func (c *clickableTestInstructionNameLabelInPreviewStruct) Tapped(*fyne.PointEvent) {

	fmt.Println("LeftClicked")

	var existInMap bool
	var attributesContainerPtr *fyne.Container
	var testCaseExecutionAttributesForPreviewMap map[testCaseExecutionsModel.TestInstructionExecutionAttributesContainerMapKeyType]*fyne.Container
	testCaseExecutionAttributesForPreviewMap = *testCaseExecutionAttributesForPreviewMapPtr
	attributesContainerPtr, existInMap = testCaseExecutionAttributesForPreviewMap[testCaseExecutionsModel.
		TestInstructionExecutionAttributesContainerMapKeyType(c.TestInstructionExecutionKey)]

	if existInMap == false {
		fmt.Println("existInMap:", existInMap)
		return
	}

	switch attributesContainerPtr.Visible() {
	case true:
		attributesContainerPtr.Hide()

	case false:
		attributesContainerPtr.Show()
	}

	testCaseExecutionPreviewContainerScroll.Refresh()
	attributesContainerPtr.Refresh()

	if c.LeftClicked != nil {

		//c.LeftClicked()
	}
}

// Optional: Handle secondary tap (right-click)
func (c *clickableTestInstructionNameLabelInPreviewStruct) TappedSecondary(*fyne.PointEvent) {

	fmt.Println("RightClicked")

	// handle secondary tap if needed
	if c.RightClicked != nil {
		//c.RightClicked()
	}
}
