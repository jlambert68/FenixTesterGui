package memoryUsage

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type ClickableContainerStruct struct {
	widget.BaseWidget
	Content     fyne.CanvasObject
	OnTapped    func(clickableContainer *ClickableContainerStruct)
	AlreadyOpen bool
}

func NewClickableContainer(
	obj fyne.CanvasObject,
	tapped func(clickableContainer *ClickableContainerStruct)) *ClickableContainerStruct {
	cc := &ClickableContainerStruct{
		Content:  obj,
		OnTapped: tapped}

	cc.ExtendBaseWidget(cc)
	return cc
}

func (c *ClickableContainerStruct) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(c.Content)
}

func (c *ClickableContainerStruct) Tapped(_ *fyne.PointEvent) {
	if c.OnTapped != nil {

		if c.AlreadyOpen == true {
			return
		}

		c.AlreadyOpen = true
		c.OnTapped(c)
	}
}

func (c *ClickableContainerStruct) TappedSecondary(_ *fyne.PointEvent) {
	// optional right-click behavior
}
