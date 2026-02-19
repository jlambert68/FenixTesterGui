# copyableAndSelectableLabel.go

## File Overview
- Path: `testSuites/listTestSuitesUI/copyableAndSelectableLabel.go`
- Package: `listTestSuitesUI`
- Functions/Methods: `4`
- Imports: `6`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `CreateRenderer`
- `Tapped`
- `TappedSecondary`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/soundEngine`
- `FenixTesterGui/testSuites/testSuiteUI`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/widget`

## Declared Types
- `copyableLabelStruct`

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### newCopyableLabel
- Signature: `func newCopyableLabel(label string, isCopyable bool, testSuiteUiref *testSuiteUI.TestSuiteUiStruct) *copyableLabelStruct`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Used for creating a new copyable label
- Selector calls: `copyableLabel.ExtendBaseWidget`

### CreateRenderer (method on `*copyableLabelStruct`)
- Signature: `func (*copyableLabelStruct) CreateRenderer() fyne.WidgetRenderer`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: CreateRenderer Renderer (required by fyne.Widget)
- Selector calls: `widget.NewLabel`, `widget.NewSimpleRenderer`

### Tapped (method on `*copyableLabelStruct`)
- Signature: `func (*copyableLabelStruct) Tapped(*fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: Tapped Tapped interface clickableAttributeInPreviewStruct

### TappedSecondary (method on `*copyableLabelStruct`)
- Signature: `func (*copyableLabelStruct) TappedSecondary(*fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `if`
- Doc: TappedSecondary Optional: Handle secondary tap (right-click)
- Selector calls: `fenixMasterWindow.Clipboard`, `clipboard.SetContent`, `fyne.CurrentApp`, `fmt.Sprintf`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
