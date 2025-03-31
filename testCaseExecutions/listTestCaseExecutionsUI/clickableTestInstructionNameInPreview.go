package listTestCaseExecutionsUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testCaseExecutions/testCaseExecutionsModel"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
)

// Definition of a clickable TestInstructionName label used in the TestCaseExecution-Preview
type clickableTestInstructionNameLabelInPreviewStruct struct {
	widget.Label
	LeftClicked  func()
	RightClicked func()
}

// Used for creating a new TestInstructionName label
func newClickableTestInstructionNameLabelInPreview(
	testInstructionName string,
	leftClicked func(),
	rightClicked func(),
) *clickableTestInstructionNameLabelInPreviewStruct {

	clickableTestInstructionNameLabelInPreview := &clickableTestInstructionNameLabelInPreviewStruct{
		Label:        widget.Label{Text: testInstructionName},
		LeftClicked:  leftClicked,
		RightClicked: rightClicked,
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

	var testCaseExecutionUuid string
	var testCaseExecutionVersion uint32

	switch selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType {

	case OneExecutionPerTestCase:
		testCaseExecutionUuid = selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.
			testCaseExecutionUuidThatIsShownInPreview
		testCaseExecutionVersion = selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.
			testCaseExecutionVersionThatIsShownInPreview

	case AllExecutionsForOneTestCase:
		testCaseExecutionUuid = selectedTestCaseExecutionObjected.allExecutionsFoOneTestCaseListObject.
			testCaseExecutionUuidThatIsShownInPreview
		testCaseExecutionVersion = selectedTestCaseExecutionObjected.allExecutionsFoOneTestCaseListObject.
			testCaseExecutionVersionThatIsShownInPreview

	default:

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "3f371409-aa81-4c8e-a63f-361f3870cd58",
			"selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType": selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType,
		}).Error("Unhandled 'selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType', should not happen")

		return

	}

	testCaseExecutionsModel.LoadDetailedTestCaseExecutionFromDatabase(testCaseExecutionUuid, testCaseExecutionVersion)

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
