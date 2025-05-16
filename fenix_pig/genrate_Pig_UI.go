package fenix_pig

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

//go:embed graphics/Fenix_pig_48x48.png
var fenixPig48x48 []byte

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

	// center it in a container

	pigMainContainer = container.NewBorder(
		nil,
		nil,
		container.NewVBox(fenixPig48x48Image),
		nil,
		nil)

	return pigMainContainer
}
