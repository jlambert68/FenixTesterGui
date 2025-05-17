package memoryUsage

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type ClickableContainerStruct struct {
	widget.BaseWidget
	Content  fyne.CanvasObject
	OnTapped func()
}

func NewClickableContainer(obj fyne.CanvasObject, tapped func()) *ClickableContainerStruct {
	cc := &ClickableContainerStruct{Content: obj, OnTapped: tapped}
	cc.ExtendBaseWidget(cc)
	return cc
}

func (c *ClickableContainerStruct) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(c.Content)
}

func (c *ClickableContainerStruct) Tapped(_ *fyne.PointEvent) {
	if c.OnTapped != nil {
		c.OnTapped()
	}
}

func (c *ClickableContainerStruct) TappedSecondary(_ *fyne.PointEvent) {
	// optional right-click behavior
}
