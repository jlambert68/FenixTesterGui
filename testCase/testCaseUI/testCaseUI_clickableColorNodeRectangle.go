package testCaseUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type clickableRectangle struct {
	widget.Label
	rectangle *canvas.Rectangle

	testCaseUuid        string
	testInstructionUuid string

	testCasesUiModelStruct *TestCasesUiModelStruct
}

func (testCasesUiCanvasObject *TestCasesUiModelStruct) NewClickableRectangle(rectangleColor color.Color, testCaseUuid string, testInstructionUuid string) *clickableRectangle {
	myClickableRectangle := &clickableRectangle{}
	myClickableRectangle.ExtendBaseWidget(myClickableRectangle)
	myClickableRectangle.SetText("")

	myClickableRectangle.testCaseUuid = testCaseUuid
	myClickableRectangle.testInstructionUuid = testInstructionUuid

	myClickableRectangle.testCasesUiModelStruct = testCasesUiCanvasObject

	// Create rectangle to show TestInstruction-color
	myClickableRectangle.rectangle = canvas.NewRectangle(rectangleColor)

	myClickableRectangle.rectangle.StrokeColor = color.RGBA{
		R: 0xFF,
		G: 0x00,
		B: 0x00,
		A: 0x33,
	}
	myClickableRectangle.rectangle.StrokeWidth = 1
	myClickableRectangle.rectangle.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize), float32(testCaseNodeRectangleSize)))
	myClickableRectangle.rectangle.Resize(fyne.NewSize(float32(testCaseNodeRectangleSize), float32(testCaseNodeRectangleSize)))

	return myClickableRectangle
}

// Tapped - Single Click on rectangle
func (c *clickableRectangle) Tapped(_ *fyne.PointEvent) {

	// Set Node to selected
	currentTestCaseModel, _ := c.testCasesUiModelStruct.TestCasesModelReference.TestCases[c.testCaseUuid]
	currentTestCaseModel.CurrentSelectedTestCaseElement.CurrentSelectedTestCaseElementUuid = c.testInstructionUuid
	c.testCasesUiModelStruct.TestCasesModelReference.TestCases[c.testCaseUuid] = currentTestCaseModel

	// Update Attributes for TestInstruction
	c.testCasesUiModelStruct.generateTestCaseAttributesAreaForTestCase(c.testCaseUuid, c.testInstructionUuid)

}

// TappedSecondary - Right Click on rectangle
func (c *clickableRectangle) TappedSecondary(_ *fyne.PointEvent) {

	currentTestCaseModel, _ := c.testCasesUiModelStruct.TestCasesModelReference.TestCases[c.testCaseUuid]
	currentTestCaseModel.CurrentSelectedTestCaseElement.CurrentSelectedTestCaseElementUuid = c.testInstructionUuid
	c.testCasesUiModelStruct.TestCasesModelReference.TestCases[c.testCaseUuid] = currentTestCaseModel

}
