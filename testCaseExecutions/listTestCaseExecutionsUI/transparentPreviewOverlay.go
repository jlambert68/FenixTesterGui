package listTestCaseExecutionsUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// HoverableRect is a widget that shows a Rectangle and lets you hook MouseIn/Out
type HoverableRect struct {
	widget.BaseWidget
	Rect       *canvas.Rectangle
	OnMouseIn  func(*desktop.MouseEvent)
	OnMouseOut func()
}

func NewHoverableRect(col color.Color) *HoverableRect {
	h := &HoverableRect{
		Rect: canvas.NewRectangle(col),
	}
	h.ExtendBaseWidget(h)
	return h
}

func (h *HoverableRect) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(h.Rect)
}

// implement desktop.Hoverable:
func (h *HoverableRect) MouseIn(ev *desktop.MouseEvent) {
	if h.OnMouseIn != nil {
		h.OnMouseIn(ev)
	}
}
func (h *HoverableRect) MouseOut() {
	if h.OnMouseOut != nil {
		h.OnMouseOut()
	}
}
func (h *HoverableRect) MouseMoved(ev *desktop.MouseEvent) {
	// no-op
}
