package testCaseUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type customAttributeCheckBoxGroup struct {
	widget.BaseWidget
	rectangle     *canvas.Rectangle
	checkBoxGroup *widget.CheckGroup
}

func newCustomAttributeCheckBoxGroupWidget(
	newCheckGroup *widget.CheckGroup,
	attributeValueIsValidWarningBox *canvas.Rectangle) *customAttributeCheckBoxGroup {
	w := &customAttributeCheckBoxGroup{}

	tempEntry := widget.NewSelect([]string{"Hallo"}, func(s string) {})

	w.rectangle = attributeValueIsValidWarningBox
	w.rectangle.SetMinSize(fyne.NewSize(15, tempEntry.Size().Height))
	w.checkBoxGroup = newCheckGroup
	w.ExtendBaseWidget(w)

	return w
}

func (w *customAttributeCheckBoxGroup) CreateRenderer() fyne.WidgetRenderer {
	return &customAttributeCheckBoxGroupRenderer{
		widget:        w,
		rectangle:     w.rectangle,
		checkBoxGroup: w.checkBoxGroup,
	}
}

type customAttributeCheckBoxGroupRenderer struct {
	widget        *customAttributeCheckBoxGroup
	rectangle     *canvas.Rectangle
	checkBoxGroup *widget.CheckGroup
}

func (r *customAttributeCheckBoxGroupRenderer) MinSize() fyne.Size {
	return fyne.NewSize(r.rectangle.MinSize().Width+r.checkBoxGroup.MinSize().Width,
		fyne.Max(r.rectangle.MinSize().Height, r.checkBoxGroup.MinSize().Height))
}

func (r *customAttributeCheckBoxGroupRenderer) Layout(size fyne.Size) {
	r.rectangle.Resize(fyne.NewSize(r.rectangle.MinSize().Width, size.Height))
	r.rectangle.Move(fyne.NewPos(0, 0))

	entrySize := fyne.NewSize(size.Width-r.rectangle.Size().Width, size.Height)
	r.checkBoxGroup.Resize(entrySize)
	r.checkBoxGroup.Move(fyne.NewPos(r.rectangle.Size().Width, 0))
}

func (r *customAttributeCheckBoxGroupRenderer) Refresh() {
	canvas.Refresh(r.widget)
}

func (r *customAttributeCheckBoxGroupRenderer) BackgroundColor() color.Color {
	return color.Transparent
}

func (r *customAttributeCheckBoxGroupRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.rectangle, r.checkBoxGroup}
}

func (r *customAttributeCheckBoxGroupRenderer) Destroy() {}
