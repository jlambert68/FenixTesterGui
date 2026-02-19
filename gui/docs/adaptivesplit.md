# adaptivesplit.go

## File Overview
- Path: `gui/adaptivesplit.go`
- Package: `gui`
- Generated: `2026-02-19T14:23:17+01:00`
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
### Layout (method on `*adaptiveLayout`)
- Signature: `func (*adaptiveLayout) Layout(objects []fyne.CanvasObject, size fyne.Size)`
- Exported: `true`
- Control-flow features: `none detected`
- External calls: `dev.IsMobile`, `dev.Orientation`, `fyne.CurrentDevice`, `fyne.IsHorizontal`

### MinSize (method on `*adaptiveLayout`)
- Signature: `func (*adaptiveLayout) MinSize(_ []fyne.CanvasObject) fyne.Size`
- Exported: `true`
- Control-flow features: `none detected`

### newAdaptiveSplit
- Signature: `func newAdaptiveSplit(left, right fyne.CanvasObject) *fyne.Container`
- Exported: `false`
- Control-flow features: `none detected`
- External calls: `container.New`, `container.NewHSplit`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
