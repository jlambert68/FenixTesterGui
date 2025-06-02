package testCaseUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type customAttributeSelectComboBox struct {
	widget.BaseWidget
	rectangle      *canvas.Rectangle
	selectComboBox *widget.Select
}

func newCustomAttributeSelectComboBoxWidget(
	newSelect *widget.Select,
	attributeValueIsValidWarningBox *canvas.Rectangle) *customAttributeSelectComboBox {
	w := &customAttributeSelectComboBox{}

	tempEntry := widget.NewSelect([]string{"Hallo"}, func(s string) {})

	w.rectangle = attributeValueIsValidWarningBox
	w.rectangle.SetMinSize(fyne.NewSize(15, tempEntry.Size().Height))
	w.selectComboBox = newSelect
	w.ExtendBaseWidget(w)

	return w
}

func (w *customAttributeSelectComboBox) CreateRenderer() fyne.WidgetRenderer {
	return &customAttributeSelectComboBoxRenderer{
		widget:         w,
		rectangle:      w.rectangle,
		selectComboBox: w.selectComboBox,
	}
}

type customAttributeSelectComboBoxRenderer struct {
	widget         *customAttributeSelectComboBox
	rectangle      *canvas.Rectangle
	selectComboBox *widget.Select
}

func (r *customAttributeSelectComboBoxRenderer) MinSize() fyne.Size {
	return fyne.NewSize(r.rectangle.MinSize().Width+r.selectComboBox.MinSize().Width,
		fyne.Max(r.rectangle.MinSize().Height, r.selectComboBox.MinSize().Height))
}

func (r *customAttributeSelectComboBoxRenderer) Layout(size fyne.Size) {
	r.rectangle.Resize(fyne.NewSize(r.rectangle.MinSize().Width, size.Height))
	r.rectangle.Move(fyne.NewPos(0, 0))

	selectComboBoxSize := fyne.NewSize(size.Width-r.rectangle.Size().Width, size.Height)
	r.selectComboBox.Resize(selectComboBoxSize)
	r.selectComboBox.Move(fyne.NewPos(r.rectangle.Size().Width, 0))
}

func (r *customAttributeSelectComboBoxRenderer) Refresh() {
	canvas.Refresh(r.widget)
}

func (r *customAttributeSelectComboBoxRenderer) BackgroundColor() color.Color {
	return color.Transparent
}

func (r *customAttributeSelectComboBoxRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.rectangle, r.selectComboBox}
}

func (r *customAttributeSelectComboBoxRenderer) Destroy() {}
