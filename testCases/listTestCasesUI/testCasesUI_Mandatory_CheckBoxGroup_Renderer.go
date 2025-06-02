package listTestCasesUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type customMandatoryCheckBoxGroup struct {
	widget.BaseWidget
	rectangle     *canvas.Rectangle
	checkBoxGroup *widget.CheckGroup
}

func newCustomMandatoryCheckBoxGroupWidget(
	newCheckGroup *widget.CheckGroup,
	attributeValueIsValidWarningBox *canvas.Rectangle) *customMandatoryCheckBoxGroup {
	w := &customMandatoryCheckBoxGroup{}

	tempEntry := widget.NewSelect([]string{"Hallo"}, func(s string) {})

	w.rectangle = attributeValueIsValidWarningBox
	w.rectangle.SetMinSize(fyne.NewSize(15, tempEntry.Size().Height))
	w.checkBoxGroup = newCheckGroup
	w.ExtendBaseWidget(w)

	return w
}

func (w *customMandatoryCheckBoxGroup) CreateRenderer() fyne.WidgetRenderer {
	return &customMandatoryCheckBoxGroupRenderer{
		widget:        w,
		rectangle:     w.rectangle,
		checkBoxGroup: w.checkBoxGroup,
	}
}

type customMandatoryCheckBoxGroupRenderer struct {
	widget        *customMandatoryCheckBoxGroup
	rectangle     *canvas.Rectangle
	checkBoxGroup *widget.CheckGroup
}

func (r *customMandatoryCheckBoxGroupRenderer) MinSize() fyne.Size {
	return fyne.NewSize(r.rectangle.MinSize().Width+r.checkBoxGroup.MinSize().Width,
		fyne.Max(r.rectangle.MinSize().Height, r.checkBoxGroup.MinSize().Height))
}

func (r *customMandatoryCheckBoxGroupRenderer) Layout(size fyne.Size) {
	r.rectangle.Resize(fyne.NewSize(r.rectangle.MinSize().Width, size.Height))
	r.rectangle.Move(fyne.NewPos(0, 0))

	entrySize := fyne.NewSize(size.Width-r.rectangle.Size().Width, size.Height)
	r.checkBoxGroup.Resize(entrySize)
	r.checkBoxGroup.Move(fyne.NewPos(r.rectangle.Size().Width, 0))
}

func (r *customMandatoryCheckBoxGroupRenderer) Refresh() {
	canvas.Refresh(r.widget)
}

func (r *customMandatoryCheckBoxGroupRenderer) BackgroundColor() color.Color {
	return color.Transparent
}

func (r *customMandatoryCheckBoxGroupRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.rectangle, r.checkBoxGroup}
}

func (r *customMandatoryCheckBoxGroupRenderer) Destroy() {}
