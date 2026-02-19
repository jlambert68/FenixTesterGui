# detailedExecutionUI_summaryTable.go

## File Overview
- Path: `executions/detailedTestCaseExecutionUI_summaryTableDefinition/detailedExecutionUI_summaryTable.go`
- Package: `detailedTestCaseExecutionUI_summaryTableDefinition`
- Functions/Methods: `7`
- Imports: `6`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `CreateRenderer`
- `Destroy`
- `Layout`
- `MinSize`
- `NewTestCaseExecutionsSummaryTable`
- `Objects`
- `Refresh`

## Imports
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/data/binding`
- `fyne.io/fyne/v2/widget`
- `log`
- `math`

## Declared Types
- `TestCaseExecutionsSummaryTableStruct`
- `testCaseExecutionSummaryTableRenderer`

## Declared Constants
- None

## Declared Variables
- `TestCaseExecutionsSummaryTable`
- `_`

## Functions and Methods
### NewTestCaseExecutionsSummaryTable
- Signature: `func NewTestCaseExecutionsSummaryTable(tableOpts *DetailedTestCaseExecutionsSummaryTableOpts) *TestCaseExecutionsSummaryTableStruct`
- Exported: `true`
- Control-flow features: `if, for/range`
- Internal calls: `NewTestcaseExecutionSummaryTableCell`, `string`, `float32`
- Selector calls: `widget.NewTable`, `b.GetItem`, `log.Fatalf`, `l.SetText`, `t.ExtendBaseWidget`, `widget.NewLabel`

### CreateRenderer (method on `*TestCaseExecutionsSummaryTableStruct`)
- Signature: `func (*TestCaseExecutionsSummaryTableStruct) CreateRenderer() fyne.WidgetRenderer`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `container.NewMax`

### MinSize (method on `testCaseExecutionSummaryTableRenderer`)
- Signature: `func (testCaseExecutionSummaryTableRenderer) MinSize() fyne.Size`
- Exported: `true`
- Control-flow features: `none detected`
- Internal calls: `float32`, `float64`
- Selector calls: `fyne.NewSize`, `math.Max`, `math.Min`

### Layout (method on `testCaseExecutionSummaryTableRenderer`)
- Signature: `func (testCaseExecutionSummaryTableRenderer) Layout(s fyne.Size)`
- Exported: `true`
- Control-flow features: `none detected`

### Destroy (method on `testCaseExecutionSummaryTableRenderer`)
- Signature: `func (testCaseExecutionSummaryTableRenderer) Destroy()`
- Exported: `true`
- Control-flow features: `none detected`

### Refresh (method on `testCaseExecutionSummaryTableRenderer`)
- Signature: `func (testCaseExecutionSummaryTableRenderer) Refresh()`
- Exported: `true`
- Control-flow features: `none detected`

### Objects (method on `testCaseExecutionSummaryTableRenderer`)
- Signature: `func (testCaseExecutionSummaryTableRenderer) Objects() []fyne.CanvasObject`
- Exported: `true`
- Control-flow features: `none detected`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
