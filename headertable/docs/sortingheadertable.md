# sortingheadertable.go

## File Overview
- Path: `headertable/sortingheadertable.go`
- Package: `headertable`
- Functions/Methods: `8`
- Imports: `11`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `CreateRenderer`
- `Destroy`
- `Layout`
- `MinSize`
- `NewSortingHeaderTable`
- `Objects`
- `Refresh`

## Imports
- `FenixTesterGui/common_code`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/data/binding`
- `fyne.io/fyne/v2/widget`
- `github.com/sirupsen/logrus`
- `log`
- `math`
- `sort`
- `sync`

## Declared Types
- `SortingHeaderTable`
- `sortingHeaderTableRenderer`

## Declared Constants
- `headerColumnExtraWidth`

## Declared Variables
- `_`

## Functions and Methods
### CreateRenderer (method on `*SortingHeaderTable`)
- Signature: `func (*SortingHeaderTable) CreateRenderer() fyne.WidgetRenderer`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `container.NewBorder`, `container.NewVBox`

### Destroy (method on `sortingHeaderTableRenderer`)
- Signature: `func (sortingHeaderTableRenderer) Destroy()`
- Exported: `true`
- Control-flow features: `none detected`

### Layout (method on `sortingHeaderTableRenderer`)
- Signature: `func (sortingHeaderTableRenderer) Layout(s fyne.Size)`
- Exported: `true`
- Control-flow features: `none detected`

### MinSize (method on `sortingHeaderTableRenderer`)
- Signature: `func (sortingHeaderTableRenderer) MinSize() fyne.Size`
- Exported: `true`
- Control-flow features: `none detected`
- Internal calls: `float32`, `float64`
- Selector calls: `fyne.NewSize`, `math.Max`

### NewSortingHeaderTable
- Signature: `func NewSortingHeaderTable(tableOpts *TableOpts) *SortingHeaderTable`
- Exported: `true`
- Control-flow features: `if, for/range`
- Internal calls: `NewFlashingTableCell`, `NewSortingLabel`, `SaveToFlashingTableCellsReferenceMap`, `TestCaseExecutionMapKeyType`, `float32`, `stringSort`
- Selector calls: `b.GetItem`, `b1.GetItem`, `dataTable.Refresh`, `fmt.Println`, `fyne.NewSize`, `l.Hide`, `l.Refresh`, `l.SetText`

### Objects (method on `sortingHeaderTableRenderer`)
- Signature: `func (sortingHeaderTableRenderer) Objects() []fyne.CanvasObject`
- Exported: `true`
- Control-flow features: `none detected`

### Refresh (method on `sortingHeaderTableRenderer`)
- Signature: `func (sortingHeaderTableRenderer) Refresh()`
- Exported: `true`
- Control-flow features: `none detected`

### stringSort
- Signature: `func stringSort(tableOpts *TableOpts, col int) SortFn`
- Exported: `false`
- Control-flow features: `if`
- Selector calls: `b1.GetItem`, `b2.GetItem`, `log.Printf`, `sort.Slice`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
