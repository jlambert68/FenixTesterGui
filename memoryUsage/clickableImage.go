package memoryUsage

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type ClickableImageStruct struct {
	widget.BaseWidget
	Image            *canvas.Image
	OnTapped         func(clickableContainer *ClickableImageStruct)
	AlreadyOpen      bool
	RectangleOverLay *canvas.Rectangle
}

func NewClickableImage(
	image *canvas.Image,
	tapped func(clickableContainer *ClickableImageStruct)) *ClickableImageStruct {

	// Create the rectangle semitransparent overlay
	var rectangleOverLay *canvas.Rectangle
	rectangleOverLay = canvas.NewRectangle(color.Transparent)
	rectangleOverLay.Resize(image.Size())

	cc := &ClickableImageStruct{
		Image:            image,
		OnTapped:         tapped,
		AlreadyOpen:      false,
		RectangleOverLay: rectangleOverLay}

	cc.ExtendBaseWidget(cc)
	return cc
}

func (c *ClickableImageStruct) CreateRenderer() fyne.WidgetRenderer {

	var widgetContainer *fyne.Container
	widgetContainer = container.NewStack(c.Image, c.RectangleOverLay)

	return widget.NewSimpleRenderer(widgetContainer)
}

func (c *ClickableImageStruct) Tapped(_ *fyne.PointEvent) {
	if c.OnTapped != nil {
		c.OnTapped(c)
	}
}

func (c *ClickableImageStruct) TappedSecondary(_ *fyne.PointEvent) {
	// optional right-click behavior
}

// MouseIn is called when a desktop pointer enters the widget
func (c *ClickableImageStruct) MouseIn(ev *desktop.MouseEvent) {

	// Change transparency level for rectangle
	c.RectangleOverLay.FillColor = color.RGBA{R: 0, G: 0, B: 0, A: 50}
	c.Refresh()
}

// MouseOut is called when a desktop pointer exits the widget
func (c *ClickableImageStruct) MouseOut() {

	// Change transparency level to be fully transparent
	c.RectangleOverLay.FillColor = color.Transparent
	c.Refresh()

}

func (c *ClickableImageStruct) MouseMoved(ev *desktop.MouseEvent) {
	// not used, but required to fully satisfy Hoverable if you want it
}
