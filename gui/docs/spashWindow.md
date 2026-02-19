# spashWindow.go

## File Overview
- Path: `gui/spashWindow.go`
- Package: `gui`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `3`
- Imports: `9`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `Focused`
- `Unfocused`

## Imports
- `FenixTesterGui/resources`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/driver/desktop`
- `fyne.io/fyne/v2/layout`
- `image/color`
- `time`

## Declared Types
- `customSplashWindow`

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### Focused (method on `*customSplashWindow`)
- Signature: `func (*customSplashWindow) Focused()`
- Exported: `true`
- Control-flow features: `none detected`
- External calls: `fmt.Println`

### Unfocused (method on `*customSplashWindow`)
- Signature: `func (*customSplashWindow) Unfocused()`
- Exported: `true`
- Control-flow features: `none detected`
- External calls: `fmt.Println`

### createSplashWindow
- Signature: `func createSplashWindow(splashWindow *fyne.Window, splashWindowProlongedVisibleChannelPtr *chan time.Duration)`
- Exported: `false`
- Control-flow features: `if, go`
- External calls: `canvas.NewImageFromResource`, `container.New`, `drv.CreateSplashWindow`, `fyne.CurrentApp`, `fyne.Do`, `layout.NewVBoxLayout`, `time.Sleep`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
