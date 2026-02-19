# listTestSuitesUI_HeaderLabel.go

## File Overview
- Path: `testSuites/listTestSuitesUI/listTestSuitesUI_HeaderLabel.go`
- Package: `listTestSuitesUI`
- Functions/Methods: `3`
- Imports: `6`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `CreateRenderer`
- `Refresh`

## Imports
- `FenixTesterGui/common_code`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/widget`
- `github.com/sirupsen/logrus`

## Declared Types
- `sortableHeaderLabelStruct`

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### newSortableHeaderLabel
- Signature: `func newSortableHeaderLabel(headerText string, tempIsSortable bool, tempColumnNumber int, listTestSuiteUI *ListTestSuiteUIStruct) *sortableHeaderLabelStruct`
- Exported: `false`
- Control-flow features: `if, switch`
- Doc: Used for creating a new Header label
- Internal calls: `newClickableSortImage`
- Selector calls: `widget.NewLabel`, `fmt.Println`, `tempSortableHeaderLabel.Refresh`, `listTestSuiteUI.sort2DStringSlice`, `tempSortableHeaderLabel.ExtendBaseWidget`

### Refresh (method on `*sortableHeaderLabelStruct`)
- Signature: `func (*sortableHeaderLabelStruct) Refresh()`
- Exported: `true`
- Control-flow features: `none detected`

### CreateRenderer (method on `*sortableHeaderLabelStruct`)
- Signature: `func (*sortableHeaderLabelStruct) CreateRenderer() fyne.WidgetRenderer`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `container.NewStack`, `container.NewHBox`, `widget.NewSimpleRenderer`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
