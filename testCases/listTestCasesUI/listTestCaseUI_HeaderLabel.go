package listTestCasesUI

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Defenition of a Header label
type sortableHeaderLabelStruct struct {
	widget.BaseWidget
	label                      *widget.Label
	isSortable                 bool
	sortImage                  *clickableSortImage
	columnNumber               int
	latestSelectedSortOrder    SortingDirectionType
	updateColumnNumberFunction func()
}

// Used for creating a new Header label
func newSortableHeaderLabel(headerText string, tempIsSortable bool, tempColumnNumber int) *sortableHeaderLabelStruct {

	tempSortableHeaderLabel := &sortableHeaderLabelStruct{
		label:        widget.NewLabel(headerText),
		isSortable:   tempIsSortable,
		columnNumber: tempColumnNumber,
	}

	// Add ClickableSortImage
	tempSortableHeaderLabel.sortImage = newClickableSortImage(
		func() {
			fmt.Println("SortIcon was Clicked!!!")
			currentSortColumn = tempColumnNumber
		},
		true,
		tempColumnNumber)

	tempSortableHeaderLabel.updateColumnNumberFunction = func() {

		fmt.Println("updateColumnNumberFunction")
		fmt.Println(tempSortableHeaderLabel.label.Text)
		fmt.Println(tempSortableHeaderLabel.columnNumber)

		tempSortableHeaderLabel.sortImage.headerColumnNumber = tempSortableHeaderLabel.columnNumber
		tempSortableHeaderLabel.sortImage.latestSelectedSortOrder = tempSortableHeaderLabel.latestSelectedSortOrder

	}

	tempSortableHeaderLabel.ExtendBaseWidget(tempSortableHeaderLabel)

	return tempSortableHeaderLabel
}

func (s *sortableHeaderLabelStruct) Refresh() {
	s.label.Refresh()
	s.sortImage.Refresh()
	s.BaseWidget.Refresh()
}

func (s *sortableHeaderLabelStruct) CreateRenderer() fyne.WidgetRenderer {

	// Create the Sort Icons container
	var newSortIconsContainer *fyne.Container
	newSortIconsContainer = container.NewStack(s.sortImage.unspecifiedImageContainer, s.sortImage)

	var newSortableHeaderLabelContainer *fyne.Container
	newSortableHeaderLabelContainer = container.NewHBox(
		s.label, newSortIconsContainer) //canvas.NewImageFromImage(sortImageUnspecifiedAsImage)) //tempNewSortableHeaderLabel.sortImage)

	fmt.Println("CreateRenderer")
	s.updateColumnNumberFunction()

	return widget.NewSimpleRenderer(newSortableHeaderLabelContainer)
}
