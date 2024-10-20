package listTestCasesUI

import (
	"fmt"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// Defenition of a Header label
type sortableHeaderLabel struct {
	widget.Label
	isSortable                 bool
	background                 *canvas.Rectangle
	sortImage                  *clickableSortImage
	ColumnNumber               int
	LatestSelectedSortOrder    SortingDirectionType
	UpdateColumnNumberFunction func()
}

// Used for creating a new Header label
func newSortableHeaderLabel(headerText string, tempIsSortable bool, tempColumnNumber int) *sortableHeaderLabel {

	tempSortableHeaderLabel := &sortableHeaderLabel{
		Label:        widget.Label{Text: headerText},
		isSortable:   tempIsSortable,
		ColumnNumber: tempColumnNumber,
	}

	// Add ClickableSortImage
	tempSortableHeaderLabel.sortImage = newClickableSortImage(
		func() {
			fmt.Println("SortIcon was Clicked!!!")
			currentSortColumn = tempColumnNumber
		},
		true,
		tempColumnNumber)

	tempSortableHeaderLabel.UpdateColumnNumberFunction = func() {

		tempSortableHeaderLabel.sortImage.headerColumnNumber = tempSortableHeaderLabel.ColumnNumber
		tempSortableHeaderLabel.sortImage.latestSelectedSortOrder = tempSortableHeaderLabel.LatestSelectedSortOrder

	}

	tempSortableHeaderLabel.ExtendBaseWidget(tempSortableHeaderLabel)

	return tempSortableHeaderLabel
}
