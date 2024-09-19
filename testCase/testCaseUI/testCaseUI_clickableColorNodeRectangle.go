package testCaseUI

import (
	"bytes"
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image"
	"image/color"
	"image/png"
	"log"
)

//go:embed resources/TIC-Horizontal_32x32.png
var tic_parallellImage []byte
var imageData_tic_parallellImage image.Image

//go:embed resources/TIC-Vertical_32x32.png
var tic_serialImage []byte
var imageData_tic_serialImage image.Image

type ClickableRectangle struct {
	widget.Label
	colorRectangle    *canvas.Rectangle
	selectedRectangle *canvas.Rectangle

	rectangleType  rectangleTypeType
	rectangleImage *canvas.Image

	testCaseUuid        string
	testInstructionUuid string

	testCasesUiModelStruct *TestCasesUiModelStruct
}

// Type for rectangle in the TestCase-UI-tree
type rectangleTypeType uint8

// Defines what kind of rectangle should be created in the TestCase-UI-tree
const (
	rectangleForTestInstruction rectangleTypeType = iota
	rectangleForParallelTestInstructionsContainer
	rectangleForSerialTestInstructionsContainer
)

func (testCasesUiCanvasObject *TestCasesUiModelStruct) NewClickableRectangle(
	rectangleColor color.Color,
	testCaseUuid string,
	testInstructionUuid string,
	rectangleType rectangleTypeType) *ClickableRectangle {

	var err error

	myClickableRectangle := &ClickableRectangle{}
	myClickableRectangle.ExtendBaseWidget(myClickableRectangle)
	myClickableRectangle.SetText("")

	myClickableRectangle.testCaseUuid = testCaseUuid
	myClickableRectangle.testInstructionUuid = testInstructionUuid

	myClickableRectangle.testCasesUiModelStruct = testCasesUiCanvasObject

	// Create colorRectangle to show TestInstruction-color
	myClickableRectangle.colorRectangle = canvas.NewRectangle(rectangleColor)

	// What should the ClickableRectangle be used as
	myClickableRectangle.rectangleType = rectangleType
	switch rectangleType {

	// Do nothing
	case rectangleForTestInstruction:

	// Add picture to indicate parallel processed TestInstructionContainer
	case rectangleForParallelTestInstructionsContainer:

		// Convert the byte slice into an image.Image object
		if imageData_tic_parallellImage == nil {
			imageData_tic_parallellImage, err = png.Decode(bytes.NewReader(tic_parallellImage))
			if err != nil {
				log.Fatalf("Failed to decode image: %v", err)
			}
		}

		myClickableRectangle.rectangleImage = canvas.NewImageFromImage(imageData_tic_parallellImage)
		myClickableRectangle.rectangleImage.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize-4), float32(testCaseNodeRectangleSize-4)))
		myClickableRectangle.rectangleImage.Resize(fyne.NewSize(float32(testCaseNodeRectangleSize-4), float32(testCaseNodeRectangleSize-4)))

	// Add picture to indicate serial processed TestInstructionContainer
	case rectangleForSerialTestInstructionsContainer:

		// Convert the byte slice into an image.Image object
		if imageData_tic_serialImage == nil {
			imageData_tic_serialImage, err = png.Decode(bytes.NewReader(tic_serialImage))
			if err != nil {
				log.Fatalf("Failed to decode image: %v", err)
			}
		}

		myClickableRectangle.rectangleImage = canvas.NewImageFromImage(imageData_tic_serialImage)
		myClickableRectangle.rectangleImage.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize-4), float32(testCaseNodeRectangleSize-4)))
		myClickableRectangle.rectangleImage.Resize(fyne.NewSize(float32(testCaseNodeRectangleSize-4), float32(testCaseNodeRectangleSize-4)))

	}

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
		A: 0xAA,
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
	if c.rectangleType == rectangleForTestInstruction {
		c.testCasesUiModelStruct.generateTestCaseAttributesAreaForTestCase(c.testCaseUuid, c.testInstructionUuid)

	}

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
