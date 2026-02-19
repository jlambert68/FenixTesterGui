# listTestSuiteExectionsUI_sortableImage.go

## File Overview
- Path: `testSuiteExecutions/listTestSuiteExecutionsUI/listTestSuiteExectionsUI_sortableImage.go`
- Package: `listTestSuiteExecutionsUI`
- Functions/Methods: `12`
- Imports: `6`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `BackgroundColor`
- `CreateRenderer`
- `Destroy`
- `Layout`
- `MinSize`
- `Objects`
- `Refresh`
- `Tapped`
- `TappedSecondary`

## Imports
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/widget`
- `image/color`
- `log`

## Declared Types
- `clickableSortImage`
- `clickableSortImageRenderer`

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### newClickableSortImage
- Signature: `func newClickableSortImage(onTapped func(), isSortable bool, headerColumnNumber int) *clickableSortImage`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: NewClickableSortImage creates a new ClickableSortImage with a given image path.
- Selector calls: `canvas.NewRectangle`, `canvas.NewImageFromImage`, `initialImageUnspecified.SetMinSize`, `fyne.NewSize`, `initialImageUnspecified.Resize`, `initialImageUnspecified.Refresh`, `container.NewStack`, `initialImageAscending.SetMinSize`

### Tapped (method on `*clickableSortImage`)
- Signature: `func (*clickableSortImage) Tapped(_ *fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `if`
- Doc: Tapped method handles click events.
- Selector calls: `log.Println`, `r.onTapped`, `r.updateImageVisibility`, `r.Refresh`

### TappedSecondary (method on `*clickableSortImage`)
- Signature: `func (*clickableSortImage) TappedSecondary(_ *fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: TappedSecondary method handles right-click events, can be ignored if not needed.

### updateImageVisibility (method on `*clickableSortImage`)
- Signature: `func (*clickableSortImage) updateImageVisibility()`
- Exported: `false`
- Control-flow features: `if, switch`

### CreateRenderer (method on `*clickableSortImage`)
- Signature: `func (*clickableSortImage) CreateRenderer() fyne.WidgetRenderer`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `container.NewMax`, `r.updateImageVisibility`

### Refresh (method on `*clickableSortImage`)
- Signature: `func (*clickableSortImage) Refresh()`
- Exported: `true`
- Control-flow features: `none detected`

### Layout (method on `*clickableSortImageRenderer`)
- Signature: `func (*clickableSortImageRenderer) Layout(size fyne.Size)`
- Exported: `true`
- Control-flow features: `none detected`

### MinSize (method on `*clickableSortImageRenderer`)
- Signature: `func (*clickableSortImageRenderer) MinSize() fyne.Size`
- Exported: `true`
- Control-flow features: `none detected`

### Refresh (method on `*clickableSortImageRenderer`)
- Signature: `func (*clickableSortImageRenderer) Refresh()`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `fyne.Do`

### BackgroundColor (method on `*clickableSortImageRenderer`)
- Signature: `func (*clickableSortImageRenderer) BackgroundColor() color.Color`
- Exported: `true`
- Control-flow features: `none detected`

### Objects (method on `*clickableSortImageRenderer`)
- Signature: `func (*clickableSortImageRenderer) Objects() []fyne.CanvasObject`
- Exported: `true`
- Control-flow features: `none detected`

### Destroy (method on `*clickableSortImageRenderer`)
- Signature: `func (*clickableSortImageRenderer) Destroy()`
- Exported: `true`
- Control-flow features: `none detected`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
