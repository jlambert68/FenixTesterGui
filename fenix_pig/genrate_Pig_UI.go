package fenix_pig

import (
	"FenixTesterGui/memoryUsage"
	_ "embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

//go:embed graphics/Fenix_pig_48x48.png
var fenixPig48x48 []byte

// GeneratePigUI
// Genrate the UI-component to be used at the bottom of the Fenix-UI, to show ongoing probes(pigs)
func GeneratePigUI() (
	pigMainContainer *fyne.Container) {

	// turn the raw bytes into a Fyne resource
	var fenixPig48x48Resource fyne.Resource
	fenixPig48x48Resource = fyne.NewStaticResource("picture.png", fenixPig48x48)

	// create an image from that resource
	var fenixPig48x48Image *canvas.Image
	fenixPig48x48Image = canvas.NewImageFromResource(fenixPig48x48Resource)

	fenixPig48x48Image.SetMinSize(fyne.NewSize(48, 48))

	var pigClickableImageContainer *memoryUsage.ClickableImageStruct
	pigClickableImageContainer = memoryUsage.NewClickableImage(
		fenixPig48x48Image, func(clickableContainer *memoryUsage.ClickableImageStruct) {
			fmt.Println("Pig clicked")
		})

	// center it in a container

	pigMainContainer = container.NewBorder(
		nil,
		nil,
		container.NewVBox(pigClickableImageContainer),
		nil,
		nil)

	return pigMainContainer
}
