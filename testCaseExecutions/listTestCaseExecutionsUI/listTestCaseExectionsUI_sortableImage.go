package listTestCaseExecutionsUI

import (
	//sharedCode "FenixTesterGui/common_code"
	//"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	//"github.com/sirupsen/logrus"
	"image/color"
	"log"
)

// ClickableSortImage represents a custom clickable image widget.
type clickableSortImage struct {
	widget.BaseWidget
	unspecifiedImageContainer *fyne.Container
	ascendingImageContainer   *fyne.Container
	descendingImageContainer  *fyne.Container
	imagesContainer           *fyne.Container
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

	/*
		// Define the Image
		r := &clickableSortImage{
			onTapped:                  onTapped,
			unspecifiedImageContainer: imageContainerUnspecified,
			ascendingImageContainer:   imageContainerAscending,
			descendingImageContainer:  imageContainerDescending,
			imageContainerToRender:    imageContainerUnspecified,
			isSortable:                isSortable,
			headerColumnNumber:        headerColumnNumber,
		}

		r.ExtendBaseWidget(r) // Necessary to extend the widget properly
		return r


	*/
	// Define the Image
	r := &clickableSortImage{
		onTapped:                  onTapped,
		unspecifiedImageContainer: imageContainerUnspecified,
		ascendingImageContainer:   imageContainerAscending,
		descendingImageContainer:  imageContainerDescending,
		isSortable:                isSortable,
		headerColumnNumber:        headerColumnNumber,
		latestSelectedSortOrder:   SortingDirectionUnSpecified, // Set the initial sort order
	}

	r.ExtendBaseWidget(r) // Necessary to extend the widget properly
	return r
}

// Tapped method handles click events.
func (r *clickableSortImage) Tapped(_ *fyne.PointEvent) {
	log.Println("Image clicked", r.headerColumnNumber)

	if !r.isSortable {
		return
	}

	if r.onTapped != nil {
		r.onTapped()
	}

	// Update image visibility
	r.updateImageVisibility()

	// Refresh the widget to update the UI
	r.Refresh()

	/*

		if r.isSortable == false {
			return
		}

		if r.onTapped != nil {
			r.onTapped()
		}

		fmt.Println("CreateRenderer")
		fmt.Println("headerColumnNumber", r.headerColumnNumber)
		fmt.Println("currentSortColumn", currentSortColumn)

		// Depending on if this is the active sort column or not, then chose correct sort-icon
		if r.headerColumnNumber == currentSortColumn {

			switch r.latestSelectedSortOrder {

			case SortingDirectionUnSpecified:
				r.imageContainerToRender = r.unspecifiedImageContainer
				fmt.Println("Should show: 'SortingDirectionUnSpecified'")

			case SortingDirectionAscending:
				r.imageContainerToRender = r.ascendingImageContainer
				fmt.Println("Should show: 'SortingDirectionAscending'")

			case SortingDirectionDescending:
				r.imageContainerToRender = r.descendingImageContainer
				fmt.Println("Should show: 'SortingDirectionDescending'")

			default:
				sharedCode.Logger.WithFields(logrus.Fields{
					"Id":                              "faec7eb8-a960-4221-ae94-a48acdfd215b",
					"currentSortColumnsSortDirection": currentSortColumnsSortDirection,
				}).Fatalln("Unhandled Sorting direction")

			}

		} else {
			r.imageContainerToRender = r.unspecifiedImageContainer
		}

		// Should the sort-icon be visible or not
		if r.isSortable == true {
			r.imageContainerToRender.Show()
		} else {
			r.imageContainerToRender.Hide()
		}


	*/
}

// TappedSecondary method handles right-click events, can be ignored if not needed.
func (r *clickableSortImage) TappedSecondary(_ *fyne.PointEvent) {}

/*
// CreateRenderer returns the renderer for the image.
func (r *clickableSortImage) CreateRenderer() fyne.WidgetRenderer {

	return widget.NewSimpleRenderer(r.imageContainerToRender)
}

*/

func (r *clickableSortImage) updateImageVisibility() {

	// Set that Previous Header only show 'unspecifiedImageContainer'
	//if previousHeader != nil {
	//	previousHeader.sortImage.unspecifiedImageContainer.Show()
	//	previousHeader.sortImage.ascendingImageContainer.Hide()
	//	previousHeader.sortImage.descendingImageContainer.Hide()
	//}

	// Hide all images first
	r.unspecifiedImageContainer.Hide()
	r.ascendingImageContainer.Hide()
	r.descendingImageContainer.Hide()

	// Show the appropriate image
	if r.isSortable {
		switch r.latestSelectedSortOrder {
		case SortingDirectionUnSpecified:
			r.unspecifiedImageContainer.Show()
		case SortingDirectionAscending:
			r.ascendingImageContainer.Show()
		case SortingDirectionDescending:
			r.descendingImageContainer.Show()
		default:
			// Handle unexpected cases
		}
	}
}

func (r *clickableSortImage) CreateRenderer() fyne.WidgetRenderer {
	// Create a container with all three image containers
	r.imagesContainer = container.NewMax(
		r.unspecifiedImageContainer,
		r.ascendingImageContainer,
		r.descendingImageContainer,
	)
	// Initially, show the appropriate image
	r.updateImageVisibility()

	return &clickableSortImageRenderer{
		clickableSortImage: r,
		imageContainer:     r.imagesContainer,
	}
}

func (r *clickableSortImage) Refresh() {
	r.BaseWidget.Refresh()
}

// Custom renderer struct
type clickableSortImageRenderer struct {
	clickableSortImage *clickableSortImage
	imageContainer     *fyne.Container
}

func (renderer *clickableSortImageRenderer) Layout(size fyne.Size) {
	renderer.imageContainer.Resize(size)
}

func (renderer *clickableSortImageRenderer) MinSize() fyne.Size {
	return renderer.imageContainer.MinSize()
}

func (renderer *clickableSortImageRenderer) Refresh() {
	renderer.clickableSortImage.updateImageVisibility()
	renderer.imageContainer.Refresh()
}

func (renderer *clickableSortImageRenderer) BackgroundColor() color.Color {
	return color.Transparent
}

func (renderer *clickableSortImageRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{renderer.imageContainer}
}

func (renderer *clickableSortImageRenderer) Destroy() {}
