package listTestCaseExecutionsUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testCaseExecutions/testCaseExecutionsModel"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
)

const (
	notDefined labelTypeType = iota
	labelIsTestInstruction
	labelIsTestInstructionContainer
)

type labelTypeType uint8

// Definition of a clickable TestInstructionName or TestInstructionContainer Name label used in the TestCaseExecution-Preview
type clickableTInTICNameLabelInPreviewStruct struct {
	widget.Label
	TInTICExecutionKey testCaseExecutionsModel.
				TCEoTICoTIEAttributesContainerMapKeyType
	LeftClicked  func()
	RightClicked func()
	LabelType    labelTypeType
}

// Used for creating a new TestInstructionName label
func newClickableTestInstructionNameLabelInPreview(
	tInTICName string,
	tInTICExecutionKey testCaseExecutionsModel.
		TCEoTICoTIEAttributesContainerMapKeyType,
	leftClicked func(),
	rightClicked func(),
	labelType labelTypeType,
) *clickableTInTICNameLabelInPreviewStruct {

	clickableTInTICNameLabelInPreview := &clickableTInTICNameLabelInPreviewStruct{
		Label:              widget.Label{Text: tInTICName},
		TInTICExecutionKey: tInTICExecutionKey,
		LeftClicked:        leftClicked,
		RightClicked:       rightClicked,
		LabelType:          labelType,
	}

	clickableTInTICNameLabelInPreview.ExtendBaseWidget(clickableTInTICNameLabelInPreview)

	return clickableTInTICNameLabelInPreview
}

// CreateRenderer
// Renderer (required by fyne.Widget)
func (c *clickableTInTICNameLabelInPreviewStruct) CreateRenderer() fyne.WidgetRenderer {
	lbl := widget.NewLabel(c.Label.Text)
	return widget.NewSimpleRenderer(lbl)
}

// Tapped
// Tapped interface implementation
func (c *clickableTInTICNameLabelInPreviewStruct) Tapped(*fyne.PointEvent) {

	fmt.Println("LeftClicked")

	testCaseExecutionAttributesForPreviewMapMutex.Lock()

	var existInMap bool
	var attributesContainerPtr *fyne.Container
	var testCaseExecutionAttributesForPreviewObjectPtr *testCaseExecutionAttributesForPreviewStruct
	var testCaseExecutionAttributesForPreviewMap map[testCaseExecutionsModel.TCEoTICoTIEAttributesContainerMapKeyType]*testCaseExecutionAttributesForPreviewStruct
	testCaseExecutionAttributesForPreviewMap = *testCaseExecutionAttributesForPreviewMapPtr
	testCaseExecutionAttributesForPreviewObjectPtr, existInMap = testCaseExecutionAttributesForPreviewMap[c.TInTICExecutionKey]

	if existInMap == false {

		// Should never happen
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":                   "765dd5e9-0f00-4494-8128-33c986c5b13d",
			"c.TInTICExecutionKey": c.TInTICExecutionKey,
		}).Error("Couldn't find object in  for 'testCaseExecutionAttributesForPreviewMap', should never happen")

		testCaseExecutionAttributesForPreviewMapMutex.Unlock()
		return
	}

	attributesContainerPtr = testCaseExecutionAttributesForPreviewObjectPtr.testInstructionExecutionAttributesContainer

	switch testCaseExecutionAttributesForPreviewObjectPtr.attributesContainerShouldBeVisible {

	// Hide attributes
	case true:

		testCaseExecutionAttributesForPreviewObjectPtr.attributesContainerShouldBeVisible = false

		// Hide attributes for object that was clicked on
		if testCaseExecutionAttributesForPreviewObjectPtr.testInstructionExecutionAttributesContainer != nil {

			attributesContainerPtr.Hide()
		}

		// Hide attributes for child-objects to object that was clicked on
		for _, childObjectKey := range testCaseExecutionAttributesForPreviewObjectPtr.childObjectsWithAttributes {

			var childTestCaseExecutionAttributesForPreviewObjectPtr *testCaseExecutionAttributesForPreviewStruct
			childTestCaseExecutionAttributesForPreviewObjectPtr, existInMap = testCaseExecutionAttributesForPreviewMap[childObjectKey]
			if existInMap == false {

				// Should never happen
				sharedCode.Logger.WithFields(logrus.Fields{
					"ID":             "222087c5-1c93-4e8f-8862-af6baf1ae2ae",
					"childObjectKey": childObjectKey,
				}).Error("Couldn't find child for TestInstructionAttributes, should never happen")

				testCaseExecutionAttributesForPreviewMapMutex.Unlock()

				return
			}

			childTestCaseExecutionAttributesForPreviewObjectPtr.attributesContainerShouldBeVisible = false
			if childTestCaseExecutionAttributesForPreviewObjectPtr.testInstructionExecutionAttributesContainer != nil {
				childTestCaseExecutionAttributesForPreviewObjectPtr.testInstructionExecutionAttributesContainer.Hide()
			}
		}

	// Show attributes
	case false:

		testCaseExecutionAttributesForPreviewObjectPtr.attributesContainerShouldBeVisible = true

		// Show attributes for object that was clicked on
		if testCaseExecutionAttributesForPreviewObjectPtr.testInstructionExecutionAttributesContainer != nil {

			attributesContainerPtr.Show()
		}

		// Hide attributes for child-objects to object that was clicked on
		for _, childObjectKey := range testCaseExecutionAttributesForPreviewObjectPtr.childObjectsWithAttributes {

			var childTestCaseExecutionAttributesForPreviewObjectPtr *testCaseExecutionAttributesForPreviewStruct
			childTestCaseExecutionAttributesForPreviewObjectPtr, existInMap = testCaseExecutionAttributesForPreviewMap[childObjectKey]
			if existInMap == false {

				// Should never happen
				sharedCode.Logger.WithFields(logrus.Fields{
					"ID":             "75369175-4406-4127-95e6-6171a73aae27",
					"childObjectKey": childObjectKey,
				}).Error("Couldn't find child for TestInstructionAttributes, should never happen")

				testCaseExecutionAttributesForPreviewMapMutex.Unlock()

				return
			}

			childTestCaseExecutionAttributesForPreviewObjectPtr.attributesContainerShouldBeVisible = true
			if childTestCaseExecutionAttributesForPreviewObjectPtr.testInstructionExecutionAttributesContainer != nil {
				childTestCaseExecutionAttributesForPreviewObjectPtr.testInstructionExecutionAttributesContainer.Show()
			}

		}

	}

	testCaseExecutionPreviewContainerScroll.Refresh()

	if testCaseExecutionAttributesForPreviewObjectPtr.testInstructionExecutionAttributesContainer != nil {
		attributesContainerPtr.Refresh()
	}

	testCaseExecutionAttributesForPreviewMapMutex.Unlock()

	if c.LeftClicked != nil {

		//c.LeftClicked()
	}
}

// TappedSecondary
// Optional: Handle secondary tap (right-click)
func (c *clickableTInTICNameLabelInPreviewStruct) TappedSecondary(*fyne.PointEvent) {

	fmt.Println("RightClicked")

	// handle secondary tap if needed
	if c.RightClicked != nil {
		//c.RightClicked()
	}
}
