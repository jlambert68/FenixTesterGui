package executionsUIForExecutions

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/headertable"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
)

// Remove item from the DataItem-slice and keep order
func remove(slice []binding.DataMap, s int) []binding.DataMap {
	return append(slice[:s], slice[s+1:]...)
}

const headerColumnExtraWidth float32 = 75

func ResizeTableColumns(t *headertable.SortingHeaderTable) {

	// Set Column widths
	bindings := t.TableOpts.Bindings
	numberOfRows, _ := t.Data.Length()
	var columnWidthToBeUsed float32
	var totalTableWidth float32

	// Loop columns
	for i, colAttr := range t.TableOpts.ColAttrs {

		// Loop Rows to get MaxTestDataWidth
		var currentColumnsMaxWidth float32
		var tempColumnWidth float32
		for rowCounter := 0; rowCounter < numberOfRows; rowCounter++ {
			b1 := bindings[rowCounter]
			d1, err := b1.GetItem(colAttr.Name)
			if err != nil {
				sharedCode.Logger.WithFields(logrus.Fields{
					"id":      "584ebb7e-9e35-4c60-b1a7-0713f973d838",
					"colAttr": colAttr,
					"b1":      b1,
				}).Fatalln("Couldn't get TestCaseExecution data due to no match")
			}

			cellData, err := d1.(binding.String).Get()
			if err != nil {
				sharedCode.Logger.WithFields(logrus.Fields{
					"id": "14e392a6-390e-4321-96c1-1da2bcbd33e1",
					"d1": d1,
				}).Fatalln("Couldn't get TestCaseExecution data due to no match")
			}

			// Check if width for row data is greater than previous max width for column
			tempColumnWidth = widget.NewLabel(cellData).MinSize().Width
			if tempColumnWidth > currentColumnsMaxWidth {
				currentColumnsMaxWidth = tempColumnWidth
			}
		}

		// Get HeaderWidth
		headerWidth := widget.NewLabel(colAttr.Header).MinSize().Width + headerColumnExtraWidth

		// Decide to used HeaderWidth or DataWidth
		if headerWidth > currentColumnsMaxWidth {
			columnWidthToBeUsed = headerWidth
		} else {
			columnWidthToBeUsed = currentColumnsMaxWidth
		}

		// Add to total Table Width
		totalTableWidth = totalTableWidth + columnWidthToBeUsed

		// Set Width for Header and data column
		t.Header.SetColumnWidth(i, float32(colAttr.WidthPercent)/100.0*columnWidthToBeUsed)
		t.Data.SetColumnWidth(i, float32(colAttr.WidthPercent)/100.0*columnWidthToBeUsed)

	}

	//t.Resize(fyne.NewSize(totalTableWidth, 200))
	t.Header.Resize(fyne.NewSize(totalTableWidth, t.Header.MinSize().Height))
	t.Data.Resize(fyne.NewSize(totalTableWidth, t.Data.MinSize().Height*float32(len(t.TableOpts.Bindings))))
}
