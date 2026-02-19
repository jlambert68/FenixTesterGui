# genrate_memoryUsage_UI.go

## File Overview
- Path: `memoryUsage/genrate_memoryUsage_UI.go`
- Package: `memoryUsage`
- Functions/Methods: `2`
- Imports: `10`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `GenerateMemoryUsageIcon`

## Imports
- `FenixTesterGui/common_code`
- `embed`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/theme`
- `github.com/skoona/sknlinechart`
- `runtime`
- `sync`
- `time`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- `memoryUsage`
- `memoryUsageWindow`

## Functions and Methods
### GenerateMemoryUsageIcon
- Signature: `func GenerateMemoryUsageIcon() memoryUsageContainer *ClickableImageStruct`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: GenerateMemoryUsageIcon Genrate the UI-component to be used at the bottom of the Fenix-UI, to show ongoing probes(pigs)
- Internal calls: `NewClickableImage`
- Selector calls: `fyne.NewStaticResource`, `canvas.NewImageFromResource`, `memoryUsageImage.SetMinSize`, `fyne.NewSize`

### openStatisticsWindowFunction
- Signature: `func openStatisticsWindowFunction(clickableContainer *ClickableImageStruct)`
- Exported: `false`
- Control-flow features: `if, for/range, select, go, defer`
- Doc: Open up the statistics window
- Internal calls: `float32`
- Selector calls: `memoryUsageWindow.RequestFocus`, `wg.Add`, `memoryUsageWindow.SetCloseIntercept`, `wg.Wait`, `memoryUsageWindow.Close`, `runtime.ReadMemStats`, `time.Now`, `sknlinechart.NewChartDatapoint`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
