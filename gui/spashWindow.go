package gui

import (
	"FenixTesterGui/resources"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"time"
)

type customSplashWindow struct {
	fyne.Window
}

func (splashWindow *customSplashWindow) Focused() {
	fmt.Println("Window gained focus!")
}

func (splashWindow *customSplashWindow) Unfocused() {
	fmt.Println("Window lost focus!")
}

func createSplashWindow(splashWindow *fyne.Window, splashWindowProlongedVisibleChannelPtr *chan time.Duration) {

	//var splashWindow fyne.Window
	var splashWindowProlongedVisibleChannel chan time.Duration
	splashWindowProlongedVisibleChannel = *splashWindowProlongedVisibleChannelPtr

	if drv, ok := fyne.CurrentApp().Driver().(desktop.Driver); ok {
		(*splashWindow) = drv.CreateSplashWindow()

		// Fenix Header
		fenixHeaderText := canvas.Text{
			Alignment: fyne.TextAlignCenter,
			Color:     nil,
			Text:      "Fenix Inception - SaaS",
			TextSize:  20,
			TextStyle: fyne.TextStyle{Bold: true},
		}

		// Text Footer
		//halFinney := widget.NewLabel("\"If you want to change the world, don't protest. Write code!\" - Hal Finney (1994)")
		halFinneyText := canvas.Text{
			Alignment: fyne.TextAlignCenter,
			Color:     nil,
			Text:      "\"If you want to change the world, don't protest. Write code!\" - Hal Finney (1994)",
			TextSize:  20,
			TextStyle: fyne.TextStyle{Italic: true},
		}

		// Fenix picture
		//image := canvas.NewImageFromResource(resources.ResourceFenix61Png)
		image := canvas.NewImageFromResource(resources.ResourceFenixdalle3v1512Png)

		image.FillMode = canvas.ImageFillOriginal

		// Container holding Header, picture and Footer
		splashContainer := container.New(layout.NewVBoxLayout(), &fenixHeaderText, image, &halFinneyText)

		(*splashWindow).SetContent(splashContainer)
		(*splashWindow).CenterOnScreen()
		(*splashWindow).Show()

		go func() {
			var sleepTime time.Duration
			sleepTime = <-splashWindowProlongedVisibleChannel
			(*splashWindow).Show()
			time.Sleep(sleepTime)
			(*splashWindow).Close()

		}()
	}

}
