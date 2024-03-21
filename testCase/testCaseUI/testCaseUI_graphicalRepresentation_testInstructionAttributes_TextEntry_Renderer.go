package testCaseUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type customAttributeEntryWidget struct {
	widget.BaseWidget
	rectangle *canvas.Rectangle
	entry     *widget.Entry
}

func newCustomAttributeEntryWidget(
	newEntry *widget.Entry,
	attributeValueIsValidWarningBox *canvas.Rectangle) *customAttributeEntryWidget {
	w := &customAttributeEntryWidget{}

	tempEntry := widget.NewEntry()
	tempEntry.SetText("Hallo")

	w.rectangle = attributeValueIsValidWarningBox
	w.rectangle.SetMinSize(fyne.NewSize(15, tempEntry.Size().Height))
	w.entry = newEntry
	w.ExtendBaseWidget(w)
	return w
}

func (w *customAttributeEntryWidget) CreateRenderer() fyne.WidgetRenderer {
	return &customEntryWidgetRenderer{
		widget:    w,
		rectangle: w.rectangle,
		entry:     w.entry,
	}
}

type customEntryWidgetRenderer struct {
	widget    *customAttributeEntryWidget
	rectangle *canvas.Rectangle
	entry     *widget.Entry
}

func (r *customEntryWidgetRenderer) MinSize() fyne.Size {
	return fyne.NewSize(r.rectangle.MinSize().Width+r.entry.MinSize().Width,
		fyne.Max(r.rectangle.MinSize().Height, r.entry.MinSize().Height))
}

func (r *customEntryWidgetRenderer) Layout(size fyne.Size) {
	r.rectangle.Resize(fyne.NewSize(r.rectangle.MinSize().Width, size.Height))
	r.rectangle.Move(fyne.NewPos(0, 0))

	entrySize := fyne.NewSize(size.Width-r.rectangle.Size().Width, size.Height)
	r.entry.Resize(entrySize)
	r.entry.Move(fyne.NewPos(r.rectangle.Size().Width, 0))
}

func (r *customEntryWidgetRenderer) Refresh() {
	canvas.Refresh(r.widget)
}

func (r *customEntryWidgetRenderer) BackgroundColor() color.Color {
	return color.Transparent
}

func (r *customEntryWidgetRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.rectangle, r.entry}
}

func (r *customEntryWidgetRenderer) Destroy() {}
