# listTestCaseExectionsUI_HeaderLabel.go

## File Overview
- Path: `testCaseExecutions/listTestCaseExecutionsUI/listTestCaseExectionsUI_HeaderLabel.go`
- Package: `listTestCaseExecutionsUI`
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
### newSortableHeaderLabel
- Signature: `func newSortableHeaderLabel(headerText string, tempIsSortable bool, tempColumnNumber int) *sortableHeaderLabelStruct`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Used for creating a new Header label
- Internal calls: `newClickableSortImage`, `SortOrReverseSortGuiTable`, `uint8`
- Selector calls: `widget.NewLabel`, `fmt.Println`, `tempSortableHeaderLabel.ExtendBaseWidget`

### Refresh (method on `*sortableHeaderLabelStruct`)
- Signature: `func (*sortableHeaderLabelStruct) Refresh()`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `fyne.Do`

### CreateRenderer (method on `*sortableHeaderLabelStruct`)
- Signature: `func (*sortableHeaderLabelStruct) CreateRenderer() fyne.WidgetRenderer`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `container.NewStack`, `container.NewHBox`, `widget.NewSimpleRenderer`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
