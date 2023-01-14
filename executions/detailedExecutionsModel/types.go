package detailedExecutionsModel

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type TestCaseExecutionsColAttr struct {
	Alignment          fyne.TextAlign
	Header             string
	Name               string
	TextStyle          fyne.TextStyle
	WidthPercent       int
	Wrapping           fyne.TextWrap
	HeaderWidth        float32
	MaxColumnDataWidth float32
}

type TestCaseExecutionsSummaryTableOpts struct {
	Bindings                       []binding.DataMap
	ColAttrs                       []TestCaseExecutionsColAttr
	OnDataCellSelect               func(cellID widget.TableCellID)
	RefWidth                       string
	HeaderLabel                    string
	FlashingTableCellsReferenceMap map[widget.TableCellID]*FlashingTableCellStruct
}

type Header struct {
	widget.Table
}
