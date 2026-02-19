# clickableImage.go

## File Overview
- Path: `memoryUsage/clickableImage.go`
- Package: `memoryUsage`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `7`
- Imports: `6`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `CreateRenderer`
- `MouseIn`
- `MouseMoved`
- `MouseOut`
- `NewClickableImage`
- `Tapped`
- `TappedSecondary`

## Imports
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/driver/desktop`
- `fyne.io/fyne/v2/widget`
- `image/color`

## Declared Types
- `ClickableImageStruct`

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### CreateRenderer (method on `*ClickableImageStruct`)
- Signature: `func (*ClickableImageStruct) CreateRenderer() fyne.WidgetRenderer`
- Exported: `true`
- Control-flow features: `none detected`
- External calls: `container.NewStack`, `widget.NewSimpleRenderer`

### MouseIn (method on `*ClickableImageStruct`)
- Signature: `func (*ClickableImageStruct) MouseIn(ev *desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: MouseIn is called when a desktop pointer enters the widget
- External calls: `c.Refresh`

### MouseMoved (method on `*ClickableImageStruct`)
- Signature: `func (*ClickableImageStruct) MouseMoved(ev *desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`

### MouseOut (method on `*ClickableImageStruct`)
- Signature: `func (*ClickableImageStruct) MouseOut()`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: MouseOut is called when a desktop pointer exits the widget
- External calls: `c.Refresh`

### NewClickableImage
- Signature: `func NewClickableImage(image *canvas.Image, tapped func(clickableContainer *ClickableImageStruct)) *ClickableImageStruct`
- Exported: `true`
- Control-flow features: `none detected`
- External calls: `canvas.NewRectangle`, `cc.ExtendBaseWidget`, `image.Size`, `rectangleOverLay.Resize`

### Tapped (method on `*ClickableImageStruct`)
- Signature: `func (*ClickableImageStruct) Tapped(_ *fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `if`
- External calls: `c.OnTapped`

### TappedSecondary (method on `*ClickableImageStruct`)
- Signature: `func (*ClickableImageStruct) TappedSecondary(_ *fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `none detected`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
