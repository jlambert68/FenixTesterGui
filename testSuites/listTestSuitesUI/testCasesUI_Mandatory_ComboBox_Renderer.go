package listTestSuitesUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type customMandatorySelectComboBox struct {
	widget.BaseWidget
	rectangle                           *canvas.Rectangle
	selectComboBox                      *widget.Select
	dataValueRepresentingVisualizedData string
}

func newCustomMandatorySelectComboBoxWidget(
	newSelect *widget.Select,
	attributeValueIsValidWarningBox *canvas.Rectangle) *customMandatorySelectComboBox {
	w := &customMandatorySelectComboBox{}

	tempEntry := widget.NewSelect([]string{"Hallo"}, func(s string) {})

	w.rectangle = attributeValueIsValidWarningBox
	w.rectangle.SetMinSize(fyne.NewSize(15, tempEntry.Size().Height))
	w.selectComboBox = newSelect
	w.ExtendBaseWidget(w)

	return w
}

func (w *customMandatorySelectComboBox) CreateRenderer() fyne.WidgetRenderer {
	return &customMandatorySelectComboBoxRenderer{
		widget:         w,
		rectangle:      w.rectangle,
		selectComboBox: w.selectComboBox,
	}
}

type customMandatorySelectComboBoxRenderer struct {
	widget         *customMandatorySelectComboBox
	rectangle      *canvas.Rectangle
	selectComboBox *widget.Select
}

func (r *customMandatorySelectComboBoxRenderer) MinSize() fyne.Size {
	return fyne.NewSize(r.rectangle.MinSize().Width+r.selectComboBox.MinSize().Width,
		fyne.Max(r.rectangle.MinSize().Height, r.selectComboBox.MinSize().Height))
}

func (r *customMandatorySelectComboBoxRenderer) Layout(size fyne.Size) {
	r.rectangle.Resize(fyne.NewSize(r.rectangle.MinSize().Width, size.Height))
	r.rectangle.Move(fyne.NewPos(0, 0))

	selectComboBoxSize := fyne.NewSize(size.Width-r.rectangle.Size().Width, size.Height)
	r.selectComboBox.Resize(selectComboBoxSize)
	r.selectComboBox.Move(fyne.NewPos(r.rectangle.Size().Width, 0))
}

func (r *customMandatorySelectComboBoxRenderer) Refresh() {
	canvas.Refresh(r.widget)
}

func (r *customMandatorySelectComboBoxRenderer) BackgroundColor() color.Color {
	return color.Transparent
}

func (r *customMandatorySelectComboBoxRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.rectangle, r.selectComboBox}
}

func (r *customMandatorySelectComboBoxRenderer) Destroy() {}
