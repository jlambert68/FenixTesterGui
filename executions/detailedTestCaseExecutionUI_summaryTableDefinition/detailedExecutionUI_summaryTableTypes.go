package detailedTestCaseExecutionUI_summaryTableDefinition

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type DetailedTestCaseExecutionsSummaryColumnsAttributes struct {
	Alignment          fyne.TextAlign
	Header             string
	Name               string
	TextStyle          fyne.TextStyle
	WidthPercent       int
	Wrapping           fyne.TextWrap
	HeaderWidth        float32
	MaxColumnDataWidth float32
}

type DetailedTestCaseExecutionsSummaryTableOpts struct {
	Bindings                       []binding.DataMap
	ColAttrs                       []DetailedTestCaseExecutionsSummaryColumnsAttributes
	OnDataCellSelect               func(cellID widget.TableCellID)
	RefWidth                       string
	HeaderLabel                    string
	FlashingTableCellsReferenceMap map[widget.TableCellID]*FlashingTableCellStruct
}

type Header struct {
	widget.Table
}
