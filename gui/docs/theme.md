# theme.go

## File Overview
- Path: `gui/theme.go`
- Package: `gui`
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
- Selector calls: `theme.DefaultTheme`

### Icon (method on `*myTheme`)
- Signature: `func (*myTheme) Icon(n fyne.ThemeIconName) fyne.Resource`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: func (m *myTheme) Font(s fyne.TextStyle) fyne.Resource { return resourceGochiHandTtf
- Selector calls: `theme.DefaultTheme`

### Size (method on `*myTheme`)
- Signature: `func (*myTheme) Size(n fyne.ThemeSizeName) float32`
- Exported: `true`
- Control-flow features: `if`
- Selector calls: `theme.DefaultTheme`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
