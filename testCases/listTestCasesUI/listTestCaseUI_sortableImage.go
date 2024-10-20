package listTestCasesUI

import (
	sharedCode "FenixTesterGui/common_code"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
	"image/color"
	"log"
)

// ClickableSortImage represents a custom clickable image widget.
type clickableSortImage struct {
	widget.BaseWidget
	unspecifiedImageContainer *fyne.Container
	ascendingImageContainer   *fyne.Container
	descendingImageContainer  *fyne.Container
	onTapped                  func() // Function to call when the image is clicked
	isSortable                bool
	headerColumnNumber        int
	latestSelectedSortOrder   SortingDirectionType
}

// NewClickableSortImage creates a new ClickableSortImage with a given image path.
func newClickableSortImage(
	onTapped func(),
	isSortable bool,
	headerColumnNumber int) *clickableSortImage {

	// Initial Setup for 'clickableSortImage' - Unspecified
	initialImageBackgroundUnspecified := canvas.NewRectangle(color.RGBA{
		R: 0x88,
		G: 0x88,
		B: 0x88,
		A: 0xFF,
	})

	initialImageUnspecified := canvas.NewImageFromImage(sortImageUnspecifiedAsImage)
	initialImageUnspecified.SetMinSize(fyne.NewSize(15, 12)) //(32, 51))
	initialImageUnspecified.Resize(fyne.NewSize(15, 12))     //(32, 51))
	initialImageUnspecified.Refresh()

	imageContainerUnspecified := container.NewStack(initialImageBackgroundUnspecified, initialImageUnspecified)

	// Initial Setup for 'clickableSortImage' - Ascending
	initialImageBackgroundAscending := canvas.NewRectangle(color.RGBA{
		R: 0x88,
		G: 0x88,
		B: 0x88,
		A: 0xFF,
	})

	initialImageAscending := canvas.NewImageFromImage(sortImageAscendingAsImage)
	initialImageAscending.SetMinSize(fyne.NewSize(15, 12)) //(32, 51))
	initialImageAscending.Resize(fyne.NewSize(15, 12))     //(32, 51))
	initialImageAscending.Refresh()

	imageContainerAscending := container.NewStack(initialImageBackgroundAscending, initialImageAscending)

	// Initial Setup for 'clickableSortImage' - Descending
	initialImageBackgroundDescending := canvas.NewRectangle(color.RGBA{
		R: 0x88,
		G: 0x88,
		B: 0x88,
		A: 0xFF,
	})

	initialImageDescending := canvas.NewImageFromImage(sortImageDescendingAsImage)
	initialImageDescending.SetMinSize(fyne.NewSize(15, 12)) //(32, 51))
	initialImageDescending.Resize(fyne.NewSize(15, 12))     //(32, 51))
	initialImageDescending.Refresh()

	imageContainerDescending := container.NewStack(initialImageBackgroundDescending, initialImageDescending)

	// Define the Image
	r := &clickableSortImage{
		onTapped:                  onTapped,
		unspecifiedImageContainer: imageContainerUnspecified,
		ascendingImageContainer:   imageContainerAscending,
		descendingImageContainer:  imageContainerDescending,
		isSortable:                isSortable,
		headerColumnNumber:        headerColumnNumber,
	}

	r.ExtendBaseWidget(r) // Necessary to extend the widget properly
	return r
}

// Tapped method handles click events.
func (r *clickableSortImage) Tapped(_ *fyne.PointEvent) {
	log.Println("Image clicked", r.headerColumnNumber)

	if r.isSortable == false {
		return
	}

	if r.onTapped != nil {
		r.onTapped()
	}
}

// TappedSecondary method handles right-click events, can be ignored if not needed.
func (r *clickableSortImage) TappedSecondary(_ *fyne.PointEvent) {}

// CreateRenderer returns the renderer for the image.
func (r *clickableSortImage) CreateRenderer() fyne.WidgetRenderer {

	var layOutContainerToReturn *fyne.Container

	// Depending on if this is the active sort column or not, then chose correct sort-icon
	if r.headerColumnNumber == currentSortColumn {

		switch r.latestSelectedSortOrder {

		case SortingDirectionUnSpecified:
			layOutContainerToReturn = container.NewWithoutLayout(r.unspecifiedImageContainer)

		case SortingDirectionAscending:
			layOutContainerToReturn = container.NewWithoutLayout(r.ascendingImageContainer)

		case SortingDirectionDescending:
			layOutContainerToReturn = container.NewWithoutLayout(r.descendingImageContainer)

		default:
			sharedCode.Logger.WithFields(logrus.Fields{
				"Id":                              "faec7eb8-a960-4221-ae94-a48acdfd215b",
				"currentSortColumnsSortDirection": currentSortColumnsSortDirection,
			}).Fatalln("Unhandled Sorting direction")

		}

	} else {
		layOutContainerToReturn = container.NewWithoutLayout(r.unspecifiedImageContainer)
	}

	// Should the sort-icon be visible or not
	if r.isSortable == true {
		layOutContainerToReturn.Show()
	} else {
		layOutContainerToReturn.Hide()
	}

	return widget.NewSimpleRenderer(layOutContainerToReturn)
}
