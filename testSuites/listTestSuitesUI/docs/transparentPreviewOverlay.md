# transparentPreviewOverlay.go

## File Overview
- Path: `testSuites/listTestSuitesUI/transparentPreviewOverlay.go`
- Package: `listTestSuitesUI`
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
### NewHoverableRect
- Signature: `func NewHoverableRect(color color.Color, otherHoverableRect *HoverableRect) *HoverableRect`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `canvas.NewRectangle`, `h.ExtendBaseWidget`

### CreateRenderer (method on `*HoverableRect`)
- Signature: `func (*HoverableRect) CreateRenderer() fyne.WidgetRenderer`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `widget.NewSimpleRenderer`

### MouseIn (method on `*HoverableRect`)
- Signature: `func (*HoverableRect) MouseIn(ev *desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `if`
- Doc: implement desktop.Hoverable:
- Selector calls: `h.OnMouseIn`

### MouseOut (method on `*HoverableRect`)
- Signature: `func (*HoverableRect) MouseOut()`
- Exported: `true`
- Control-flow features: `if`
- Selector calls: `h.OnMouseOut`

### MouseMoved (method on `*HoverableRect`)
- Signature: `func (*HoverableRect) MouseMoved(ev *desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
