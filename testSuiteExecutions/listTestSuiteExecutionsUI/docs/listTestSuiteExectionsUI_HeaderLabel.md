# listTestSuiteExectionsUI_HeaderLabel.go

## File Overview
- Path: `testSuiteExecutions/listTestSuiteExecutionsUI/listTestSuiteExectionsUI_HeaderLabel.go`
- Package: `listTestSuiteExecutionsUI`
- Functions/Methods: `3`
- Imports: `4`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `CreateRenderer`
- `Refresh`

## Imports
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/widget`

## Declared Types
- `sortableHeaderLabelStruct`

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### CreateRenderer (method on `*sortableHeaderLabelStruct`)
- Signature: `func (*sortableHeaderLabelStruct) CreateRenderer() fyne.WidgetRenderer`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `container.NewHBox`, `container.NewStack`, `widget.NewSimpleRenderer`

### Refresh (method on `*sortableHeaderLabelStruct`)
- Signature: `func (*sortableHeaderLabelStruct) Refresh()`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `fyne.Do`

### newSortableHeaderLabel
- Signature: `func newSortableHeaderLabel(headerText string, tempIsSortable bool, tempColumnNumber int) *sortableHeaderLabelStruct`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Used for creating a new Header label
- Internal calls: `SortOrReverseSortGuiTable`, `newClickableSortImage`, `uint8`
- Selector calls: `fmt.Println`, `tempSortableHeaderLabel.ExtendBaseWidget`, `widget.NewLabel`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
