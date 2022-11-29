package headertable

import (
	sharedCode "FenixTesterGui/common_code"
	"fyne.io/fyne/v2/container"
	"github.com/sirupsen/logrus"
	"log"
	"math"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

const headerColumnExtraWidth float32 = 75

var _ fyne.Widget = &HeaderTable{}

type SortingHeaderTable struct {
	widget.BaseWidget
	TableOpts *TableOpts
	Header    *widget.Table
	Data      *widget.Table
	//MagicTable *widget.Table
	sortLabels  []*sortingLabel
	HeaderLabel *widget.Label
}

func NewSortingHeaderTable(tableOpts *TableOpts) *SortingHeaderTable {
	sortLabels := make([]*sortingLabel, len(tableOpts.ColAttrs))

	dataTable := widget.NewTable(
		// Dimensions (rows, cols)
		func() (int, int) { return len(tableOpts.Bindings), len(tableOpts.ColAttrs) },

		// Default value
		func() fyne.CanvasObject { return widget.NewLabel("wide content") },

		// Cell values
		func(cellID widget.TableCellID, cnvObj fyne.CanvasObject) {
			b := tableOpts.Bindings[cellID.Row]
			itemKey := tableOpts.ColAttrs[cellID.Col].Name
			d, err := b.GetItem(itemKey)
			if err != nil {
				log.Fatalf("Data table Update Cell callback, GetItem(%s): %s", itemKey, err)
			}
			str, err := d.(binding.String).Get()
			if err != nil {
				log.Fatalf("Data table Update Cell callback, Get: %s", err)
			}
			l := cnvObj.(*widget.Label)
			l.SetText(str)
		},
	)
	headerTable := widget.NewTable(
		// Dimensions (rows, cols)
		func() (int, int) { return 1, len(tableOpts.ColAttrs) },

		// Default value
		func() fyne.CanvasObject { return NewSortingLabel("the content") },

		// Cell values
		func(cellID widget.TableCellID, o fyne.CanvasObject) {
			l := o.(*sortingLabel)
			sortLabels[cellID.Col] = l
			col := cellID.Col
			opts := tableOpts.ColAttrs[col]
			l.Sorter = stringSort(tableOpts, col)
			l.OnAfterSort = func() {
				dataTable.Refresh()
				// Set all but this column to unsorted
				for i, sl := range sortLabels {
					if i != cellID.Col {
						sl.SetState(SortUnsorted)
					}
				}
			}
			l.Col = col
			l.Label.SetText(opts.Header)
			l.Label.TextStyle = opts.TextStyle
			l.Label.Alignment = opts.Alignment
			l.Label.Wrapping = opts.Wrapping
			l.Refresh()
		},
	)
	/*
		myMagicTable := widget.NewTable(
			// Dimensions (rows, cols)
			func() (int, int) { return 1 + len(tableOpts.Bindings), len(tableOpts.ColAttrs) },

			// Default value
			func() fyne.CanvasObject { return widget.NewLabel("magic content") },
			func() fyne.CanvasObject { return NewSortingLabel("the content") },

			func(cellID widget.TableCellID, cnvObj fyne.CanvasObject) {
				if cellID.Row == 0 {
					// Header

					l := cnvObj.(*sortingLabel)
					sortLabels[cellID.Col] = l
					col := cellID.Col
					opts := tableOpts.ColAttrs[col]
					l.Sorter = stringSort(tableOpts, col)
					l.OnAfterSort = func() {
						dataTable.Refresh()
						// Set all but this column to unsorted
						for i, sl := range sortLabels {
							if i != cellID.Col {
								sl.SetState(SortUnsorted)
							}
						}
					}
					l.Col = col
					l.Label.SetText(opts.Header)
					l.Label.TextStyle = opts.TextStyle
					l.Label.Alignment = opts.Alignment
					l.Label.Wrapping = opts.Wrapping
					l.Refresh()

				} else {
					// Data

					b := tableOpts.Bindings[cellID.Row-1]
					itemKey := tableOpts.ColAttrs[cellID.Col].Name
					d, err := b.GetItem(itemKey)
					if err != nil {
						log.Fatalf("Data table Update Cell callback, GetItem(%s): %s", itemKey, err)
					}
					str, err := d.(binding.String).Get()
					if err != nil {
						log.Fatalf("Data table Update Cell callback, Get: %s", err)
					}
					l := cnvObj.(*widget.Label)
					l.SetText(str)
				}

			},
		)
	*/
	headerLabel := widget.NewLabel(tableOpts.HeaderLable)

	t := &SortingHeaderTable{
		sortLabels:  sortLabels,
		TableOpts:   tableOpts,
		Header:      headerTable,
		Data:        dataTable,
		HeaderLabel: headerLabel,
		//MagicTable: myMagicTable,
	}
	t.ExtendBaseWidget(t)

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

	t.Hide()
	t.Show()

	return t
}

func stringSort(tableOpts *TableOpts, col int) SortFn {
	return func(ascending bool) {
		log.Printf("Request to sort column %d ascending: %t\n", col, ascending)
		bindings := tableOpts.Bindings
		sort.Slice(bindings, func(i int, j int) bool {
			b1 := bindings[i]
			b2 := bindings[j]
			d1, err := b1.GetItem(tableOpts.ColAttrs[col].Name)
			if err != nil {
				panic(err)
			}
			d2, err := b2.GetItem(tableOpts.ColAttrs[col].Name)
			if err != nil {
				panic(err)
			}
			str1, err := d1.(binding.String).Get()
			if err != nil {
				panic(err)
			}
			str2, err := d2.(binding.String).Get()
			if err != nil {
				panic(err)
			}
			if ascending {
				return str1 < str2
			} else {
				return str1 > str2
			}
		})
	}
}

func (h *SortingHeaderTable) CreateRenderer() fyne.WidgetRenderer {

	topContainer := container.NewVBox(h.HeaderLabel, h.Header)

	return sortingHeaderTableRenderer{
		headerTable: h,
		container:   container.NewBorder(topContainer, nil, nil, nil, h.Data),
		//container: container.NewVBox(h.Header, h.Data),
	}
}

//*******************************************************************************

var _ fyne.WidgetRenderer = sortingHeaderTableRenderer{}

type sortingHeaderTableRenderer struct {
	headerTable *SortingHeaderTable
	container   *fyne.Container
}

func (r sortingHeaderTableRenderer) MinSize() fyne.Size {
	return fyne.NewSize(
		float32(math.Max(float64(r.headerTable.Data.MinSize().Width), float64(r.headerTable.Header.MinSize().Width))),
		r.headerTable.Data.MinSize().Height+r.headerTable.Header.MinSize().Height)
}

func (r sortingHeaderTableRenderer) Layout(s fyne.Size) {
	r.container.Resize(s)
}

func (r sortingHeaderTableRenderer) Destroy() {
}

func (r sortingHeaderTableRenderer) Refresh() {
	r.container.Refresh()
}

func (r sortingHeaderTableRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.container}
}
