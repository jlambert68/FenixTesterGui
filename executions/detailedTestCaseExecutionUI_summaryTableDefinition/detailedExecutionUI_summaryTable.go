package detailedTestCaseExecutionUI_summaryTableDefinition

import (
	"log"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Widget = &TestCaseExecutionsSummaryTable{}

type TestCaseExecutionsSummaryTable struct {
	widget.BaseWidget
	TableOpts *TestCaseExecutionsSummaryTableOpts
	//Header    *widget.Table
	Data *widget.Table
}

func NewTestCaseExecutionsSummaryTable(tableOpts *TestCaseExecutionsSummaryTableOpts) *TestCaseExecutionsSummaryTable {
	t := &TestCaseExecutionsSummaryTable{
		TableOpts: tableOpts,
		/*Header: widget.NewTable(
			// Dimensions (rows, cols)
			func() (int, int) { return 1, len(tableOpts.ColAttrs) },
			// Default value
			func() fyne.CanvasObject { return widget.NewLabel("the content") },
			// Cell values
			func(cellID widget.TableCellID, o fyne.CanvasObject) {
				l := o.(*widget.Label)
				opts := tableOpts.ColAttrs[cellID.Col]
				l.SetText(opts.Header)
				l.TextStyle = opts.TextStyle
				l.Alignment = opts.Alignment
				l.Wrapping = opts.Wrapping
				l.Refresh()
			},
		),*/
		Data: widget.NewTable(
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
		),
	}
	t.ExtendBaseWidget(t)

	// Set Column widths
	refWidth := widget.NewLabel(t.TableOpts.RefWidth).MinSize().Width
	for i, colAttr := range t.TableOpts.ColAttrs {
		t.Data.SetColumnWidth(i, float32(colAttr.WidthPercent)/100.0*refWidth)
		//t.Header.SetColumnWidth(i, float32(colAttr.WidthPercent)/100.0*refWidth)
	}

	return t
}

//*******************************************************************************

var _ fyne.WidgetRenderer = headerTableRenderer{}

type headerTableRenderer struct {
	headerTable *TestCaseExecutionsSummaryTable
	container   *fyne.Container
}

func (h *TestCaseExecutionsSummaryTable) CreateRenderer() fyne.WidgetRenderer {
	return headerTableRenderer{
		headerTable: h,
		container:   container.NewBorder(h.Data, nil, nil, nil),
	}
}

func (r headerTableRenderer) MinSize() fyne.Size {
	return fyne.NewSize(
		float32(math.Max(float64(r.headerTable.Data.MinSize().Width), float64(r.headerTable.Data.MinSize().Width))),
		r.headerTable.Data.MinSize().Height+r.headerTable.Data.MinSize().Height)
}

func (r headerTableRenderer) Layout(s fyne.Size) {
	r.container.Resize(s)
}

func (r headerTableRenderer) Destroy() {
}

func (r headerTableRenderer) Refresh() {
	r.container.Refresh()
}

func (r headerTableRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.container}
}
