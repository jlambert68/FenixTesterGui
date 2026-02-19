# copyableAndSelectableLabel.go

## File Overview
- Path: `testSuites/listTestSuitesUI/copyableAndSelectableLabel.go`
- Package: `listTestSuitesUI`
- Generated: `2026-02-19T14:23:17+01:00`
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
### CreateRenderer (method on `*copyableLabelStruct`)
- Signature: `func (*copyableLabelStruct) CreateRenderer() fyne.WidgetRenderer`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: CreateRenderer Renderer (required by fyne.Widget)
- External calls: `widget.NewLabel`, `widget.NewSimpleRenderer`

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
- External calls: `clipboard.SetContent`, `fenixMasterWindow.Clipboard`, `fmt.Sprintf`, `fyne.CurrentApp`

### newCopyableLabel
- Signature: `func newCopyableLabel(label string, isCopyable bool, testSuiteUiref *testSuiteUI.TestSuiteUiStruct) *copyableLabelStruct`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Used for creating a new copyable label
- External calls: `copyableLabel.ExtendBaseWidget`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
