# theme.go

## File Overview
- Path: `gui/theme.go`
- Package: `gui`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `3`
- Imports: `3`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `Color`
- `Icon`
- `Size`

## Imports
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/theme`
- `image/color`

## Declared Types
- `myTheme`

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### Color (method on `*myTheme`)
- Signature: `func (*myTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color`
- Exported: `true`
- Control-flow features: `if, switch`
- External calls: `theme.DefaultTheme`

### Icon (method on `*myTheme`)
- Signature: `func (*myTheme) Icon(n fyne.ThemeIconName) fyne.Resource`
- Exported: `true`
- Control-flow features: `none detected`
- External calls: `theme.DefaultTheme`

### Size (method on `*myTheme`)
- Signature: `func (*myTheme) Size(n fyne.ThemeSizeName) float32`
- Exported: `true`
- Control-flow features: `if`
- External calls: `theme.DefaultTheme`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
