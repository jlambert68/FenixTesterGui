package testCaseUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func newAdaptiveSplit(leftTop, leftBottom, rightTop, rightBottom fyne.CanvasObject) *fyne.Container {

	leftSplit := container.NewVSplit(leftTop, leftBottom)
	leftSplit.Offset = 0.8

	rightSplit := container.NewVSplit(rightTop, rightBottom)
	rightSplit.Offset = 0.8

	split := container.NewHSplit(leftSplit, rightSplit)
	split.Offset = 0.33
	return container.New(&adaptiveLayout{split: split}, split)
}

type adaptiveLayout struct {
	split *container.Split
}

func (a *adaptiveLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	dev := fyne.CurrentDevice()

	a.split.Horizontal = !dev.IsMobile() || fyne.IsHorizontal(dev.Orientation())
	objects[0].Resize(size)
}

func (a *adaptiveLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return a.split.MinSize()
}
