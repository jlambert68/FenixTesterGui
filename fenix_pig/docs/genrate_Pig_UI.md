# genrate_Pig_UI.go

## File Overview
- Path: `fenix_pig/genrate_Pig_UI.go`
- Package: `fenix_pig`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `1`
- Imports: `6`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `GeneratePigUI`

## Imports
- `FenixTesterGui/memoryUsage`
- `embed`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/container`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- `fenixPig48x48`

## Functions and Methods
### GeneratePigUI
- Signature: `func GeneratePigUI() pigMainContainer *fyne.Container`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: GeneratePigUI Genrate the UI-component to be used at the bottom of the Fenix-UI, to show ongoing probes(pigs)
- External calls: `canvas.NewImageFromResource`, `container.NewBorder`, `container.NewVBox`, `fenixPig48x48Image.SetMinSize`, `fmt.Println`, `fyne.NewSize`, `fyne.NewStaticResource`, `memoryUsage.NewClickableImage`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
