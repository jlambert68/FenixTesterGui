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
	myClickableRectangle.rectangle.StrokeColor = color.Black
	myClickableRectangle.rectangle.StrokeWidth = 0
	myClickableRectangle.rectangle.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize), float32(testCaseNodeRectangleSize)))
	myClickableRectangle.rectangle.Resize(fyne.NewSize(float32(testCaseNodeRectangleSize), float32(testCaseNodeRectangleSize)))

	return myClickableRectangle
}

func (t *clickableRectangle) Tapped(_ *fyne.PointEvent) {

	t.testCasesUiModelStruct.generateTestCaseAttributesAreaForTestCase(t.testCaseUuid, t.testInstructionUuid)

}
