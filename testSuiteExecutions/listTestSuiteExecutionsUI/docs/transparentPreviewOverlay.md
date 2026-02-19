# transparentPreviewOverlay.go

## File Overview
- Path: `testSuiteExecutions/listTestSuiteExecutionsUI/transparentPreviewOverlay.go`
- Package: `listTestSuiteExecutionsUI`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `5`
- Imports: `5`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `CreateRenderer`
- `MouseIn`
- `MouseMoved`
- `MouseOut`
- `NewHoverableRect`

## Imports
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/driver/desktop`
- `fyne.io/fyne/v2/widget`
- `image/color`

## Declared Types
- `HoverableRect`

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### CreateRenderer (method on `*HoverableRect`)
- Signature: `func (*HoverableRect) CreateRenderer() fyne.WidgetRenderer`
- Exported: `true`
- Control-flow features: `none detected`
- External calls: `widget.NewSimpleRenderer`

### MouseIn (method on `*HoverableRect`)
- Signature: `func (*HoverableRect) MouseIn(ev *desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `if`
- Doc: implement desktop.Hoverable:
- External calls: `h.OnMouseIn`

### MouseMoved (method on `*HoverableRect`)
- Signature: `func (*HoverableRect) MouseMoved(ev *desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`

### MouseOut (method on `*HoverableRect`)
- Signature: `func (*HoverableRect) MouseOut()`
- Exported: `true`
- Control-flow features: `if`
- External calls: `h.OnMouseOut`

### NewHoverableRect
- Signature: `func NewHoverableRect(color color.Color, otherHoverableRect *HoverableRect) *HoverableRect`
- Exported: `true`
- Control-flow features: `none detected`
- External calls: `canvas.NewRectangle`, `h.ExtendBaseWidget`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
