package detailedTestCaseExecutionUI_summaryTableDefinition

import (
	"log"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Widget = &TestCaseExecutionsSummaryTableStruct{}

type TestCaseExecutionsSummaryTableStruct struct {
	widget.BaseWidget
	TableOpts *DetailedTestCaseExecutionsSummaryTableOpts
	//Header    *widget.Table
	Data *widget.Table
}

var TestCaseExecutionsSummaryTable *TestCaseExecutionsSummaryTableStruct

func NewTestCaseExecutionsSummaryTable(tableOpts *DetailedTestCaseExecutionsSummaryTableOpts) *TestCaseExecutionsSummaryTableStruct {
	t := &TestCaseExecutionsSummaryTableStruct{
		TableOpts: tableOpts,

		Data: widget.NewTable(
			// Dimensions (rows, cols)
			func() (int, int) { return len(tableOpts.Bindings), len(tableOpts.ColAttrs) },

			// Default value
			func() fyne.CanvasObject { return NewTestcaseExecutionSummaryTableCell("wide content") },

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
				l := cnvObj.(*TestCaseExecutionSummaryTableCellStruct).Label
				l.SetText(str)

				// Update row number for "FlashingTableCell"
				cnvObj.(*TestCaseExecutionSummaryTableCellStruct).rowNumber = cellID.Row

				var mylables []string
				mylables = append(mylables, "First")
				mylables = append(mylables, "Second")
				mylables = append(mylables, "Third")
				cnvObj.(*TestCaseExecutionSummaryTableCellStruct).MyLabelsTextValues = mylables

				cnvObj.(*TestCaseExecutionSummaryTableCellStruct).MyLabels[0](*widget.Label)

				if cellID.Row == 0 {
					backgroundRectangle := cnvObj.(*TestCaseExecutionSummaryTableCellStruct).backgroundColorRectangle
					backgroundRectangle.FillColor = headerBackgroundRectangleBaseColor
					cnvObj.(*TestCaseExecutionSummaryTableCellStruct).Label.SetText("Hej")
				}

				// If  there are no 'TestCaseExecutionMapKey' then update it
				if cnvObj.(*TestCaseExecutionSummaryTableCellStruct).TestCaseExecutionMapKey == "" {
					testcaseExecutionUuidReference, err := b.GetItem("TestCaseExecutionUuid")
					if err != nil {
						log.Fatalf("Data table Update Cell callback, GetItem(%s): %s", itemKey, err)
					}
					testcaseExecutionVersionReference, err := b.GetItem("TestCaseExecutionVersion")
					if err != nil {
						log.Fatalf("Data table Update Cell callback, GetItem(%s): %s", itemKey, err)
					}

					testcaseExecutionUuidValue, err := testcaseExecutionUuidReference.(binding.String).Get()
					if err != nil {
						log.Fatalf("Data table Update Cell callback, Get: %s", err)
					}
					testcaseExecutionVersionValue, err := testcaseExecutionVersionReference.(binding.String).Get()
					if err != nil {
						log.Fatalf("Data table Update Cell callback, Get: %s", err)
					}

					var tempTestCaseExecutionMapKey string
					tempTestCaseExecutionMapKey = testcaseExecutionUuidValue + testcaseExecutionVersionValue

					cnvObj.(*TestCaseExecutionSummaryTableCellStruct).TestCaseExecutionMapKey = string(tempTestCaseExecutionMapKey)
					cnvObj.(*TestCaseExecutionSummaryTableCellStruct).TestCaseExecutionMapKeyLabel.SetText(string(tempTestCaseExecutionMapKey))

					// Add reference to 'testCaseExecutionsDetails'
					var testCaseExecutionsDetails *TestCaseExecutionsDetailsStruct
					var existInMap bool
					testCaseExecutionsDetails, existInMap = TestCaseExecutionsDetailsMap[tempTestCaseExecutionMapKey]

					// If TestExecutionExecution doesn't exist in map then create a new instance
					if existInMap == false {
						log.Fatalf("Shouldn't be like this")
					}
					cnvObj.(*TestCaseExecutionSummaryTableCellStruct).testCaseExecutionsDetails = testCaseExecutionsDetails

					//testCaseExecutionsDetailsMap := *tableOpts.TestCaseExecutionsDetailsMapReference
					//testCaseExecutionsDetails, existInMap = testCaseExecutionsDetailsMap[tempTestCaseExecutionMapKey]

					// If TestExecutionExecution doesn't exist in map then create a new instance
					if existInMap == false {
						log.Fatalf("Shouldn't be like this")
					}
					cnvObj.(*TestCaseExecutionSummaryTableCellStruct).testCaseExecutionsDetails = testCaseExecutionsDetails
					/*
						testInstructionExecutionsStatusForSummaryTableReference, err := b.GetItem("TestInstructionExecutionsStatusForSummaryTable")
						if err != nil {
							log.Fatalf("Data table Update Cell callback, GetItem(%s): %s", "TestInstructionExecutionsStatusForSummaryTable", err)
						}
						testInstructionExecutionsStatusForSummaryTableValue, err := testInstructionExecutionsStatusForSummaryTableReference.(binding.String).Get()
						if err != nil {
							log.Fatalf("Data table Update Cell callback, Get: %s", err)
						}
						testInstructionExecutionsStatusForSummaryTableValue

						cnvObj.(*TestCaseExecutionSummaryTableCellStruct).testCaseExecutionsStatusForSummaryTable = testInstructionExecutionsStatusForSummaryTableValue
					*/
					var testInstructionExecutionNames []fyne.CanvasObject
					var testInstructionExecutionsStatusForSummaryTableSReference *[]*TestInstructionExecutionsStatusForSummaryTableStruct
					testInstructionExecutionsStatusForSummaryTableSReference = testCaseExecutionsDetails.TestCaseExecutionsStatusForSummaryTable.TestInstructionExecutionsStatusForSummaryTable

					var testInstructionExecutionsStatusForSummaryTable []*TestInstructionExecutionsStatusForSummaryTableStruct
					testInstructionExecutionsStatusForSummaryTable = *testInstructionExecutionsStatusForSummaryTableSReference

					for _, testInstructionNameRef := range testInstructionExecutionsStatusForSummaryTable {
						testInstructionExecutionNames = append(testInstructionExecutionNames,
							container.NewMax(widget.NewLabel(testInstructionNameRef.TestInstructionExecutionUIName)))
					}
					cnvObj.(*TestCaseExecutionSummaryTableCellStruct).testInstructionExecutionNames = testInstructionExecutionNames

				}

				// Add reference to 'flashingTableCell'
				tableOpts.TestCaseExecutionsDetailsMapReference[cellID] = cnvObj.(*TestCaseExecutionSummaryTableCellStruct)
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

var _ fyne.WidgetRenderer = testCaseExecutionSummaryTableRenderer{}

type testCaseExecutionSummaryTableRenderer struct {
	testCaseExecutionSummaryTable *TestCaseExecutionsSummaryTableStruct
	container                     *fyne.Container
}

func (h *TestCaseExecutionsSummaryTableStruct) CreateRenderer() fyne.WidgetRenderer {
	return testCaseExecutionSummaryTableRenderer{
		testCaseExecutionSummaryTable: h,
		container:                     container.NewMax(h.Data), // container.NewBorder(h.Data, nil, nil, nil),
	}
}

func (r testCaseExecutionSummaryTableRenderer) MinSize() fyne.Size {

	return fyne.NewSize(
		float32(math.Max(float64(r.testCaseExecutionSummaryTable.Data.MinSize().Width),
			float64(r.testCaseExecutionSummaryTable.Data.MinSize().Width))),
		float32(math.Min(
			float64(r.testCaseExecutionSummaryTable.Data.MinSize().Height*float32(10)), // Minimum is 10 rows
			float64(r.testCaseExecutionSummaryTable.Data.MinSize().Height*float32(len(r.testCaseExecutionSummaryTable.TableOpts.Bindings))))))
	//r.testCaseExecutionSummaryTable.Data.MinSize().Height+r.testCaseExecutionSummaryTable.Data.MinSize().Height)
}

func (r testCaseExecutionSummaryTableRenderer) Layout(s fyne.Size) {
	r.container.Resize(s)
}

func (r testCaseExecutionSummaryTableRenderer) Destroy() {
}

func (r testCaseExecutionSummaryTableRenderer) Refresh() {
	r.container.Refresh()
}

func (r testCaseExecutionSummaryTableRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.container}
}
