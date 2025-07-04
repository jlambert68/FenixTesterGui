package testSuiteUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

// HoverableRect is a widget that shows a Rectangle and lets you hook MouseIn/Out
type HoverableRect struct {
	widget.BaseWidget
	Rect               *canvas.Rectangle
	OnMouseIn          func(*desktop.MouseEvent)
	OnMouseOut         func()
	OtherHoverableRect *HoverableRect
}

func NewHoverableRect(color color.Color, otherHoverableRect *HoverableRect) *HoverableRect {
	h := &HoverableRect{
		Rect:               canvas.NewRectangle(color),
		OtherHoverableRect: otherHoverableRect,
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
