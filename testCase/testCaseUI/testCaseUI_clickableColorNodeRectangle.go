package testCaseUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type ClickableRectangle struct {
	widget.Label
	colorRectangle    *canvas.Rectangle
	selectedRectangle *canvas.Rectangle

	testCaseUuid        string
	testInstructionUuid string

	testCasesUiModelStruct *TestCasesUiModelStruct
}

func (testCasesUiCanvasObject *TestCasesUiModelStruct) NewClickableRectangle(rectangleColor color.Color, testCaseUuid string, testInstructionUuid string) *ClickableRectangle {
	myClickableRectangle := &ClickableRectangle{}
	myClickableRectangle.ExtendBaseWidget(myClickableRectangle)
	myClickableRectangle.SetText("")

	myClickableRectangle.testCaseUuid = testCaseUuid
	myClickableRectangle.testInstructionUuid = testInstructionUuid

	myClickableRectangle.testCasesUiModelStruct = testCasesUiCanvasObject

	// Create colorRectangle to show TestInstruction-color
	myClickableRectangle.colorRectangle = canvas.NewRectangle(rectangleColor)

	myClickableRectangle.colorRectangle.StrokeColor = color.RGBA{
		R: 0xFF,
		G: 0x00,
		B: 0x00,
		A: 0x33,
	}
	myClickableRectangle.colorRectangle.StrokeWidth = 1
	myClickableRectangle.colorRectangle.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize), float32(testCaseNodeRectangleSize)))
	myClickableRectangle.colorRectangle.Resize(fyne.NewSize(float32(testCaseNodeRectangleSize), float32(testCaseNodeRectangleSize)))

	// Create selected rectangle to show when selected
	myClickableRectangle.selectedRectangle = canvas.NewRectangle(rectangleColor)

	myClickableRectangle.selectedRectangle.StrokeColor = color.RGBA{
		R: 0xFF,
		G: 0xFF,
		B: 0x00,
		A: 0x33,
	}
	myClickableRectangle.selectedRectangle.StrokeWidth = 3
	myClickableRectangle.selectedRectangle.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize-3), float32(testCaseNodeRectangleSize-3)))
	myClickableRectangle.selectedRectangle.Resize(fyne.NewSize(float32(testCaseNodeRectangleSize-3), float32(testCaseNodeRectangleSize-3)))

	myClickableRectangle.selectedRectangle.Hide()

	// Update Selected Node/Element to be selected
	myClickableRectangle.updateSelectedUINode()

	return myClickableRectangle
}

// Tapped - Single Click on colorRectangle
func (c *ClickableRectangle) Tapped(_ *fyne.PointEvent) {

	// Update Selected Node/Element
	c.updateSelectedUINode()

	// Update Attributes for TestInstruction
	c.testCasesUiModelStruct.generateTestCaseAttributesAreaForTestCase(c.testCaseUuid, c.testInstructionUuid)

}

// TappedSecondary - Right Click on colorRectangle
func (c *ClickableRectangle) TappedSecondary(_ *fyne.PointEvent) {

	// Update Selected Node/Element
	c.updateSelectedUINode()

}

func (c *ClickableRectangle) updateSelectedUINode() {

	// Clear graphics for previous selected
	previousSelectedClickableRectangle := c.testCasesUiModelStruct.CurrentSelectedTestCaseUIElement
	if previousSelectedClickableRectangle != nil {
		previousSelectedClickableRectangle.selectedRectangle.Hide()
	}

	// Set a reference to current UI-ClickableRectangle
	c.testCasesUiModelStruct.CurrentSelectedTestCaseUIElement = c

	// Show Graphics for Selected UI-node
	c.selectedRectangle.Show()

	// Set Node to selected
	currentTestCaseModel, _ := c.testCasesUiModelStruct.TestCasesModelReference.TestCases[c.testCaseUuid]
	currentTestCaseModel.CurrentSelectedTestCaseElement.CurrentSelectedTestCaseElementUuid = c.testInstructionUuid
	c.testCasesUiModelStruct.TestCasesModelReference.TestCases[c.testCaseUuid] = currentTestCaseModel

}

// Do a click on the Rectangle
func (c *ClickableRectangle) ForceClick() {

	// Update Selected Node/Element
	c.testCasesUiModelStruct.generateTestCaseAttributesAreaForTestCase(c.testCaseUuid, c.testInstructionUuid)

}
