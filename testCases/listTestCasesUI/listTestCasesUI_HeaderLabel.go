package listTestCasesUI

import (
	sharedCode "FenixTesterGui/common_code"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
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
	listTestCaseUIPtr          *ListTestCaseUIStruct
}

// Used for creating a new Header label
func newSortableHeaderLabel(headerText string, tempIsSortable bool, tempColumnNumber int, listTestCaseUI *ListTestCaseUIStruct) *sortableHeaderLabelStruct {

	tempSortableHeaderLabel := &sortableHeaderLabelStruct{
		label:             widget.NewLabel(headerText),
		isSortable:        tempIsSortable,
		columnNumber:      tempColumnNumber,
		listTestCaseUIPtr: listTestCaseUI,
	}

	// Add ClickableSortImage
	tempSortableHeaderLabel.sortImage = newClickableSortImage(
		func() {
			fmt.Println("SortIcon was Clicked!!!", tempSortableHeaderLabel.columnNumber)

			// Check if there is a new that the table should be sorted on
			if listTestCaseUI.currentSortColumn != tempSortableHeaderLabel.columnNumber {

				// New column to sort on
				previousSortColumn = listTestCaseUI.currentSortColumn
				listTestCaseUI.currentSortColumn = tempSortableHeaderLabel.columnNumber

				if listTestCaseUI.currentHeader != nil {
					listTestCaseUI.previousHeader = listTestCaseUI.currentHeader

					// Reset the previous header's sort order to Unspecified
					listTestCaseUI.previousHeader.latestSelectedSortOrder = SortingDirectionUnSpecified
					listTestCaseUI.previousHeader.sortImage.latestSelectedSortOrder = SortingDirectionUnSpecified

					// Refresh the previous header's sort image and widget
					listTestCaseUI.previousHeader.sortImage.Refresh()
					listTestCaseUI.previousHeader.Refresh()

					//previousHeader.sortImage.unspecifiedImageContainer.Show()
					//previousHeader.sortImage.ascendingImageContainer.Hide()
					//previousHeader.sortImage.descendingImageContainer.Hide()
					//previousHeader.sortImage.Refresh()

				}
				listTestCaseUI.currentHeader = tempSortableHeaderLabel

				// New column so use the previous  sort-direction if that existed
				switch tempSortableHeaderLabel.latestSelectedSortOrder {

				case SortingDirectionUnSpecified:
					tempSortableHeaderLabel.latestSelectedSortOrder = SortingDirectionAscending

				case SortingDirectionAscending:
					// Do nothing

				case SortingDirectionDescending:
					// Do nothing

				default:
					sharedCode.Logger.WithFields(logrus.Fields{
						"Id": "f6c3f4ec-91c3-4b2a-bdab-0aef96453a2a",
						"tempSortableHeaderLabel.latestSelectedSortOrder": tempSortableHeaderLabel.latestSelectedSortOrder,
						"ColumnNumber": tempSortableHeaderLabel.columnNumber,
					}).Fatalln("Unhandled SortOrder")
				}

			} else {

				// Same column so switch sort-direction
				switch tempSortableHeaderLabel.latestSelectedSortOrder {

				case SortingDirectionUnSpecified:
					tempSortableHeaderLabel.latestSelectedSortOrder = SortingDirectionAscending

				case SortingDirectionAscending:
					tempSortableHeaderLabel.latestSelectedSortOrder = SortingDirectionDescending

				case SortingDirectionDescending:
					tempSortableHeaderLabel.latestSelectedSortOrder = SortingDirectionAscending

				default:
					sharedCode.Logger.WithFields(logrus.Fields{
						"Id": "f6c3f4ec-91c3-4b2a-bdab-0aef96453a2a",
						"tempSortableHeaderLabel.latestSelectedSortOrder": tempSortableHeaderLabel.latestSelectedSortOrder,
						"ColumnNumber": tempSortableHeaderLabel.columnNumber,
					}).Fatalln("Unhandled SortOrder")
				}
			}

			// Update the current header's sort image
			tempSortableHeaderLabel.sortImage.latestSelectedSortOrder = tempSortableHeaderLabel.latestSelectedSortOrder
			tempSortableHeaderLabel.sortImage.Refresh()
			tempSortableHeaderLabel.Refresh()

			listTestCaseUI.sort2DStringSlice(listTestCaseUI.testCasesListTableTable, listTestCaseUI.currentSortColumn, tempSortableHeaderLabel.latestSelectedSortOrder)
			listTestCaseUI.testCaseListTable.Refresh()

		},
		true,
		tempColumnNumber)

	tempSortableHeaderLabel.updateColumnNumberFunction = func() {

		fmt.Println("updateColumnNumberFunction")
		fmt.Println(tempSortableHeaderLabel.label.Text)
		fmt.Println(tempSortableHeaderLabel.columnNumber)

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
	newSortIconsContainer = container.NewStack(s.sortImage)

	var newSortableHeaderLabelContainer *fyne.Container
	newSortableHeaderLabelContainer = container.NewHBox(
		s.label, newSortIconsContainer)

	return widget.NewSimpleRenderer(newSortableHeaderLabelContainer)
}
