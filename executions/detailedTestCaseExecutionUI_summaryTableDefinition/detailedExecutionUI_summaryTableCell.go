package detailedTestCaseExecutionUI_summaryTableDefinition

import (
	"FenixTesterGui/executions/detailedExecutionsModel"
	"FenixTesterGui/resources"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

var _ fyne.Widget = (*TestCaseExecutionSummaryTableCellStruct)(nil)

type TestCaseExecutionSummaryTableCellStruct struct {
	widget.BaseWidget
	Label                             *widget.Label
	backgroundColorRectangle          *canvas.Rectangle
	showDetailedTestCaseExecution     *canvas.Image
	subrscriptionForTestCaseExecution *canvas.Image
	rowNumber                         int
	TestCaseExecutionMapKey           string
	testCaseExecutionsDetails         *detailedExecutionsModel.TestCaseExecutionsDetailsStruct
}

func (t *TestCaseExecutionSummaryTableCellStruct) Tapped(_ *fyne.PointEvent) {
	fmt.Println("I was clicked!!! ", t.Label)

}

func (t *TestCaseExecutionSummaryTableCellStruct) TappedSecondary(_ *fyne.PointEvent) {
	fmt.Println("I was Right clicked!!!")

}

func (testcaseExecutionSummaryTableCell *TestCaseExecutionSummaryTableCellStruct) DoubleTapped(_ *fyne.PointEvent) {

	fmt.Println("I was Double clicked!!!")

}

var backgroundRectangleBaseColor = color.RGBA{
	R: 0x33,
	G: 0x33,
	B: 0x33,
	A: 0x33,
}

var headerBackgroundRectangleBaseColor = color.RGBA{
	R: 0x33,
	G: 0x33,
	B: 0x33,
	A: 0x88,
}

func FlashAddedRow(testcaseExecutionSummaryTableCell *TestCaseExecutionSummaryTableCellStruct) {

	go func(testcaseExecutionSummaryTableCell *TestCaseExecutionSummaryTableCellStruct) {

		// Define how the Color-flash should look like
		rectangleColorAnimation := canvas.NewColorRGBAAnimation(backgroundRectangleBaseColor,
			color.RGBA{
				R: 0x00,
				G: 0xFF,
				B: 0x00,
				A: 0xAA,
			}, time.Millisecond*200, func(animationColorValue color.Color) {
				testcaseExecutionSummaryTableCell.backgroundColorRectangle.FillColor = animationColorValue
				canvas.Refresh(testcaseExecutionSummaryTableCell.backgroundColorRectangle)
			})

		// Initiate Color-flash
		rectangleColorAnimation.AutoReverse = true
		rectangleColorAnimation.Start()

	}(testcaseExecutionSummaryTableCell)

}

func FlashRowToBeRemoved(testcaseExecutionSummaryTableCell *TestCaseExecutionSummaryTableCellStruct) {
	go func(testcaseExecutionSummaryTableCell *TestCaseExecutionSummaryTableCellStruct) {

		// Define how the Color-flash should look like
		rectangleColorAnimation := canvas.NewColorRGBAAnimation(backgroundRectangleBaseColor,
			color.RGBA{
				R: 0xFF,
				G: 0x00,
				B: 0x00,
				A: 0xAA,
			}, time.Millisecond*200, func(animationColorValue color.Color) {
				testcaseExecutionSummaryTableCell.backgroundColorRectangle.FillColor = animationColorValue
				canvas.Refresh(testcaseExecutionSummaryTableCell.backgroundColorRectangle)
			})

		// Initiate Color-flash
		rectangleColorAnimation.AutoReverse = true
		rectangleColorAnimation.Start()

	}(testcaseExecutionSummaryTableCell)

}

func NewTestcaseExecutionSummaryTableCell(text string) *TestCaseExecutionSummaryTableCellStruct {
	newtestcaseExecutionSummaryTableCell := &TestCaseExecutionSummaryTableCellStruct{
		Label:                             widget.NewLabel(text),
		backgroundColorRectangle:          canvas.NewRectangle(backgroundRectangleBaseColor),
		showDetailedTestCaseExecution:     canvas.NewImageFromResource(resources.ResourceIcons8CheckMarkButton48Png),
		subrscriptionForTestCaseExecution: canvas.NewImageFromResource(resources.ResourceIcons8CheckMarkButton48Png),
	}

	// Hide the showDetailedTestCaseExecution-image and set it to fill its parent
	newtestcaseExecutionSummaryTableCell.showDetailedTestCaseExecution.FillMode = canvas.ImageFillContain
	newtestcaseExecutionSummaryTableCell.showDetailedTestCaseExecution.Hide()

	// Hide the subrscriptionForTestCaseExecution-image and set it to fill its parent
	newtestcaseExecutionSummaryTableCell.subrscriptionForTestCaseExecution.FillMode = canvas.ImageFillContain
	newtestcaseExecutionSummaryTableCell.subrscriptionForTestCaseExecution.Hide()

	newtestcaseExecutionSummaryTableCell.ExtendBaseWidget(newtestcaseExecutionSummaryTableCell)
	return newtestcaseExecutionSummaryTableCell

}

func (newtestcaseExecutionSummaryTableCell *TestCaseExecutionSummaryTableCellStruct) CreateRenderer() fyne.WidgetRenderer {

	// Generate slice of TestInstructionExecutions
	var testInstructionExecutionNames []fyne.CanvasObject
	for _, testInstructionName := range newtestcaseExecutionSummaryTableCell.testCaseExecutionsDetails.TestInstructionExecutionsStatusForSummaryTable {
		testInstructionExecutionNames = append(testInstructionExecutionNames,
			container.NewMax(widget.NewLabel(testInstructionName.TestInstructionExecutionUIName)))
	}

	// Use standard cell
	return &testcaseExecutionSummaryTableCellRenderer{
		testcaseExecutionSummaryTableCell: newtestcaseExecutionSummaryTableCell,
		container: container.NewMax(
			container.NewVBox(
				newtestcaseExecutionSummaryTableCell.Label,
				container.NewHBox(testInstructionExecutionNames...),
			),
			newtestcaseExecutionSummaryTableCell.backgroundColorRectangle), //newtestcaseExecutionSummaryTableCell.showDetailedTestCaseExecution,
	}
}

var _ fyne.WidgetRenderer = (*testcaseExecutionSummaryTableCellRenderer)(nil)

type testcaseExecutionSummaryTableCellRenderer struct {
	testcaseExecutionSummaryTableCell *TestCaseExecutionSummaryTableCellStruct
	container                         *fyne.Container
}

func (r *testcaseExecutionSummaryTableCellRenderer) MinSize() fyne.Size {
	return r.container.MinSize()
}

func (r *testcaseExecutionSummaryTableCellRenderer) Layout(size fyne.Size) {
	r.container.Resize(size)
}

func (r *testcaseExecutionSummaryTableCellRenderer) Refresh() {
	r.container.Refresh()
}

func (r *testcaseExecutionSummaryTableCellRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.container}
}

func (r *testcaseExecutionSummaryTableCellRenderer) Destroy() {}
