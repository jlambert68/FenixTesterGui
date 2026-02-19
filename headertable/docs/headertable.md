# headertable.go

## File Overview
- Path: `headertable/headertable.go`
- Package: `headertable`
- Functions/Methods: `7`
- Imports: `6`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `CreateRenderer`
- `Destroy`
- `Layout`
- `MinSize`
- `NewHeaderTable`
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
- `HeaderTable`
- `headerTableRenderer`

## Declared Constants
- None

## Declared Variables
- `_`

## Functions and Methods
### NewHeaderTable
- Signature: `func NewHeaderTable(tableOpts *TableOpts) *HeaderTable`
- Exported: `true`
- Control-flow features: `if, for/range`
- Internal calls: `float32`
- Selector calls: `widget.NewTable`, `widget.NewLabel`, `l.SetText`, `l.Refresh`, `b.GetItem`, `log.Fatalf`, `t.ExtendBaseWidget`

### CreateRenderer (method on `*HeaderTable`)
- Signature: `func (*HeaderTable) CreateRenderer() fyne.WidgetRenderer`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `container.NewBorder`

### MinSize (method on `headerTableRenderer`)
- Signature: `func (headerTableRenderer) MinSize() fyne.Size`
- Exported: `true`
- Control-flow features: `none detected`
- Internal calls: `float32`, `float64`
- Selector calls: `fyne.NewSize`, `math.Max`

### Layout (method on `headerTableRenderer`)
- Signature: `func (headerTableRenderer) Layout(s fyne.Size)`
- Exported: `true`
- Control-flow features: `none detected`

### Destroy (method on `headerTableRenderer`)
- Signature: `func (headerTableRenderer) Destroy()`
- Exported: `true`
- Control-flow features: `none detected`

### Refresh (method on `headerTableRenderer`)
- Signature: `func (headerTableRenderer) Refresh()`
- Exported: `true`
- Control-flow features: `none detected`

### Objects (method on `headerTableRenderer`)
- Signature: `func (headerTableRenderer) Objects() []fyne.CanvasObject`
- Exported: `true`
- Control-flow features: `none detected`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
