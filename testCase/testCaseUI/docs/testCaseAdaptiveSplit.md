# testCaseAdaptiveSplit.go

## File Overview
- Path: `testCase/testCaseUI/testCaseAdaptiveSplit.go`
- Package: `testCaseUI`
- Functions/Methods: `3`
- Imports: `2`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `Layout`
- `MinSize`

## Imports
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/container`

## Declared Types
- `adaptiveLayout`

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### newAdaptiveSplit
- Signature: `func newAdaptiveSplit(leftTop, leftBottom, rightTop, rightBottom fyne.CanvasObject) *fyne.Container`
- Exported: `false`
- Control-flow features: `none detected`
- Selector calls: `container.NewVSplit`, `container.NewHSplit`, `container.New`

### Layout (method on `*adaptiveLayout`)
- Signature: `func (*adaptiveLayout) Layout(objects []fyne.CanvasObject, size fyne.Size)`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `fyne.CurrentDevice`, `dev.IsMobile`, `fyne.IsHorizontal`, `dev.Orientation`

### MinSize (method on `*adaptiveLayout`)
- Signature: `func (*adaptiveLayout) MinSize(_ []fyne.CanvasObject) fyne.Size`
- Exported: `true`
- Control-flow features: `none detected`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
