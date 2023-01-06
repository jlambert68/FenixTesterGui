package executionsUIForExecutions

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func newThreePartAdaptiveSplit(top, middle, bottom fyne.CanvasObject) *fyne.Container {
	topAndMiddlesplit := container.NewVSplit(top, middle)
	topAndMiddlesplit.Offset = 0.33

	topAndBottomsplit := container.NewVSplit(topAndMiddlesplit, bottom)
	topAndBottomsplit.Offset = 0.33

	return container.New(&adaptiveLayout{split: topAndBottomsplit}, topAndBottomsplit)
}

type adaptiveLayout struct {
	split *container.Split
}

func (a *adaptiveLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	//dev := fyne.CurrentDevice()

	//a.split.Horizontal = !dev.IsMobile() || fyne.IsHorizontal(dev.Orientation())
	objects[0].Resize(size)
}

func (a *adaptiveLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return a.split.MinSize()
}
