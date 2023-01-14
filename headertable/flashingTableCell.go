package headertable

import (
	executionsModelForExecutions "FenixTesterGui/executions/detailedExecutionsModel"
	"FenixTesterGui/resources"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

var _ fyne.Widget = (*FlashingTableCellStruct)(nil)

type FlashCellWhenRemoveFromTableFunctionType func(ascending bool)

type FlashCellWhenAddToTableFunctionType func(ascending bool)

// Type used to define that this is TestCaseExecutionKey for model-maps
type TestCaseExecutionMapKeyType string // Should consist of 'TestCaseExecutionUuid' + 'TestCaseExecutionVersion'

type FlashingTableCellStruct struct {
	widget.BaseWidget
	Label                                *widget.Label
	backgroundColorRectangle             *canvas.Rectangle
	showDetailedTestCaseExecution        *canvas.Image
	rowNumber                            int
	TestCaseExecutionMapKey              TestCaseExecutionMapKeyType
	FlashCellWhenRemoveFromTableFunction FlashCellWhenRemoveFromTableFunctionType
	FlashCellWhenAddToTableFunction      FlashCellWhenAddToTableFunctionType
}

func (t *FlashingTableCellStruct) Tapped(_ *fyne.PointEvent) {
	fmt.Println("I was clicked!!! ", t.Label)

}

func (t *FlashingTableCellStruct) TappedSecondary(_ *fyne.PointEvent) {
	fmt.Println("I was Right clicked!!!")

}

func (flashingTableCell *FlashingTableCellStruct) DoubleTapped(_ *fyne.PointEvent) {

	// Switch True/ false for ShowDetailedTestCaseExecution
	if flashingTableCell.showDetailedTestCaseExecution.Hidden == true {

		// Send message Executions Details handler to retrieve full TestCaseExecutions details
		err := executionsModelForExecutions.RetrieveSingleTestCaseExecution(string(flashingTableCell.TestCaseExecutionMapKey))

		// Only Switch if there was no error when doing the gRPC-call to GuiExecutionServer
		if err != nil {
			flashingTableCell.Label.Text = "true"
			flashingTableCell.showDetailedTestCaseExecution.Show()
		}

	} else {
		flashingTableCell.Label.Text = "false"
		flashingTableCell.showDetailedTestCaseExecution.Hide()
	}

	flashingTableCell.showDetailedTestCaseExecution.Refresh()
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

func FlashAddedRow(flashingTableCell *FlashingTableCellStruct) {

	go func(flashingTableCell *FlashingTableCellStruct) {

		// Define how the Color-flash should look like
		rectangleColorAnimation := canvas.NewColorRGBAAnimation(backgroundRectangleBaseColor,
			color.RGBA{
				R: 0x00,
				G: 0xFF,
				B: 0x00,
				A: 0xAA,
			}, time.Millisecond*200, func(animationColorValue color.Color) {
				flashingTableCell.backgroundColorRectangle.FillColor = animationColorValue
				canvas.Refresh(flashingTableCell.backgroundColorRectangle)
			})

		// Initiate Color-flash
		rectangleColorAnimation.AutoReverse = true
		rectangleColorAnimation.Start()

	}(flashingTableCell)

}

func FlashRowToBeRemoved(flashingTableCell *FlashingTableCellStruct) {
	go func(flashingTableCell *FlashingTableCellStruct) {

		// Define how the Color-flash should look like
		rectangleColorAnimation := canvas.NewColorRGBAAnimation(backgroundRectangleBaseColor,
			color.RGBA{
				R: 0xFF,
				G: 0x00,
				B: 0x00,
				A: 0xAA,
			}, time.Millisecond*200, func(animationColorValue color.Color) {
				flashingTableCell.backgroundColorRectangle.FillColor = animationColorValue
				canvas.Refresh(flashingTableCell.backgroundColorRectangle)
			})

		// Initiate Color-flash
		rectangleColorAnimation.AutoReverse = true
		rectangleColorAnimation.Start()

	}(flashingTableCell)

}

func NewFlashingTableCell(text string) *FlashingTableCellStruct {
	newFlashingTableCell := &FlashingTableCellStruct{
		Label:                    widget.NewLabel(text),
		backgroundColorRectangle: canvas.NewRectangle(backgroundRectangleBaseColor),
		//FlashCellWhenRemoveFromTableFunction: widget.NewButton("", func() {}),
		//FlashCellWhenAddToTableFunction:  SortUnsorted,
		showDetailedTestCaseExecution: canvas.NewImageFromResource(resources.ResourceIcons8CheckMarkButton48Png),
	}

	// Hide the image and set it to fill its parent
	newFlashingTableCell.showDetailedTestCaseExecution.FillMode = canvas.ImageFillContain
	newFlashingTableCell.showDetailedTestCaseExecution.Hide()

	newFlashingTableCell.ExtendBaseWidget(newFlashingTableCell)
	return newFlashingTableCell

}

func (newflashingTableCell *FlashingTableCellStruct) CreateRenderer() fyne.WidgetRenderer {

	// Use standard cell
	return &flashingTableCellRenderer{
		flashingTableCell: newflashingTableCell,
		container:         container.NewMax(newflashingTableCell.Label, newflashingTableCell.showDetailedTestCaseExecution, newflashingTableCell.backgroundColorRectangle), //newflashingTableCell.showDetailedTestCaseExecution,
	}
}

var _ fyne.WidgetRenderer = (*flashingTableCellRenderer)(nil)

type flashingTableCellRenderer struct {
	flashingTableCell *FlashingTableCellStruct
	container         *fyne.Container
}

func (r *flashingTableCellRenderer) MinSize() fyne.Size {
	return r.container.MinSize()
}

func (r *flashingTableCellRenderer) Layout(size fyne.Size) {
	r.container.Resize(size)
}

func (r *flashingTableCellRenderer) Refresh() {
	r.container.Refresh()
}

func (r *flashingTableCellRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.container}
}

func (r *flashingTableCellRenderer) Destroy() {}
