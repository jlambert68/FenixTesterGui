package headertable

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

var _ fyne.Widget = (*flashingTableCellStruct)(nil)

type FlashCellWhenRemoveFromTableFunctionType func(ascending bool)

type FlashCellWhenAddToTableFunctionType func(ascending bool)

// Type used to define that this is TestCaseExecutionKey for model-maps
type TestCaseExecutionMapKeyType string // Should consist of 'TestCaseExecutionUuid' + 'TestCaseExecutionVersion'

type flashingTableCellStruct struct {
	widget.BaseWidget
	Label                                *widget.Label
	backgroundColorRectangle             *canvas.Rectangle
	TestCaseExecutionMapKey              TestCaseExecutionMapKeyType
	FlashCellWhenRemoveFromTableFunction FlashCellWhenRemoveFromTableFunctionType
	FlashCellWhenAddToTableFunction      FlashCellWhenAddToTableFunctionType
}

func NewFlashingTableCell(text string) *flashingTableCellStruct {
	newFlashingTableCell := &flashingTableCellStruct{
		Label: widget.NewLabel(text),
		backgroundColorRectangle: canvas.NewRectangle(color.RGBA{
			R: 0x33,
			G: 0x33,
			B: 0x33,
			A: 0x33,
		}),
		//FlashCellWhenRemoveFromTableFunction: widget.NewButton("", func() {}),
		//FlashCellWhenAddToTableFunction:  SortUnsorted,
	}

	newFlashingTableCell.ExtendBaseWidget(newFlashingTableCell)
	return newFlashingTableCell

}

func (newflashingTableCell *flashingTableCellStruct) CreateRenderer() fyne.WidgetRenderer {
	return &flashingTableCellRenderer{
		flashingTableCell: newflashingTableCell,
		container:         container.NewMax(newflashingTableCell.Label, newflashingTableCell.backgroundColorRectangle),
	}
}

var _ fyne.WidgetRenderer = (*flashingTableCellRenderer)(nil)

type flashingTableCellRenderer struct {
	flashingTableCell *flashingTableCellStruct
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
