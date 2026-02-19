# sortlabel.go

## File Overview
- Path: `headertable/sortlabel.go`
- Package: `headertable`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `10`
- Imports: `5`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `CreateRenderer`
- `Destroy`
- `Layout`
- `MinSize`
- `NewSortingLabel`
- `Objects`
- `OnTapped`
- `Refresh`
- `SetState`

## Imports
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/widget`
- `github.com/PaulWaldo/fyne-headertable/headertable/data`
- `log`

## Declared Types
- `SortFn`
- `SortState`
- `sortingLabel`
- `sortingLabelRenderer`

## Declared Constants
- `SortAscending`
- `SortDescending`
- `SortUnsorted`

## Declared Variables
- `_`

## Functions and Methods
### CreateRenderer (method on `*sortingLabel`)
- Signature: `func (*sortingLabel) CreateRenderer() fyne.WidgetRenderer`
- Exported: `true`
- Control-flow features: `none detected`
- External calls: `container.NewHBox`

### Destroy (method on `*sortingLabelRenderer`)
- Signature: `func (*sortingLabelRenderer) Destroy()`
- Exported: `true`
- Control-flow features: `none detected`

### Layout (method on `*sortingLabelRenderer`)
- Signature: `func (*sortingLabelRenderer) Layout(size fyne.Size)`
- Exported: `true`
- Control-flow features: `none detected`

### MinSize (method on `*sortingLabelRenderer`)
- Signature: `func (*sortingLabelRenderer) MinSize() fyne.Size`
- Exported: `true`
- Control-flow features: `none detected`

### NewSortingLabel
- Signature: `func NewSortingLabel(text string) *sortingLabel`
- Exported: `true`
- Control-flow features: `none detected`
- External calls: `sl.ExtendBaseWidget`, `sl.SetState`, `widget.NewButton`, `widget.NewLabel`

### Objects (method on `*sortingLabelRenderer`)
- Signature: `func (*sortingLabelRenderer) Objects() []fyne.CanvasObject`
- Exported: `true`
- Control-flow features: `none detected`

### OnTapped (method on `*sortingLabel`)
- Signature: `func (*sortingLabel) OnTapped()`
- Exported: `true`
- Control-flow features: `if`
- External calls: `s.OnAfterSort`, `s.SetState`, `s.Sorter`, `s.nextState`

### Refresh (method on `*sortingLabelRenderer`)
- Signature: `func (*sortingLabelRenderer) Refresh()`
- Exported: `true`
- Control-flow features: `none detected`

### SetState (method on `*sortingLabel`)
- Signature: `func (*sortingLabel) SetState(state SortState)`
- Exported: `true`
- Control-flow features: `switch`
- External calls: `log.Fatalf`

### nextState (method on `*sortingLabel`)
- Signature: `func (*sortingLabel) nextState() SortState`
- Exported: `false`
- Control-flow features: `switch`
- External calls: `log.Printf`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
