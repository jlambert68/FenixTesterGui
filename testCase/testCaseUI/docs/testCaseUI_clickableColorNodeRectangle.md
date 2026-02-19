# testCaseUI_clickableColorNodeRectangle.go

## File Overview
- Path: `testCase/testCaseUI/testCaseUI_clickableColorNodeRectangle.go`
- Package: `testCaseUI`
- Functions/Methods: `5`
- Imports: `10`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `ForceClick`
- `NewClickableRectangle`
- `Tapped`
- `TappedSecondary`

## Imports
- `FenixTesterGui/testCase/testCaseModel`
- `bytes`
- `embed`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/widget`
- `image`
- `image/color`
- `image/png`
- `log`

## Declared Types
- `ClickableRectangle`
- `rectangleTypeType`

## Declared Constants
- `rectangleForParallelTestInstructionsContainer`
- `rectangleForSerialTestInstructionsContainer`
- `rectangleForTestInstruction`

## Declared Variables
- `imageData_tic_parallellImage`
- `imageData_tic_serialImage`
- `tic_parallellImage`
- `tic_serialImage`

## Functions and Methods
### ForceClick (method on `*ClickableRectangle`)
- Signature: `func (*ClickableRectangle) ForceClick()`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: Do a click on the Rectangle

### NewClickableRectangle (method on `*TestCasesUiModelStruct`)
- Signature: `func (*TestCasesUiModelStruct) NewClickableRectangle(rectangleColor color.Color, testCaseUuid string, testInstructionUuid string, rectangleType rectangleTypeType) *ClickableRectangle`
- Exported: `true`
- Control-flow features: `if, switch`
- Internal calls: `float32`
- Selector calls: `bytes.NewReader`, `canvas.NewImageFromImage`, `canvas.NewRectangle`, `fyne.NewSize`, `log.Fatalf`, `myClickableRectangle.ExtendBaseWidget`, `myClickableRectangle.SetText`, `myClickableRectangle.updateSelectedUINode`

### Tapped (method on `*ClickableRectangle`)
- Signature: `func (*ClickableRectangle) Tapped(_ *fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `if`
- Doc: Tapped - Single Click on colorRectangle
- Selector calls: `c.updateSelectedUINode`

### TappedSecondary (method on `*ClickableRectangle`)
- Signature: `func (*ClickableRectangle) TappedSecondary(_ *fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: TappedSecondary - Right Click on colorRectangle
- Selector calls: `c.updateSelectedUINode`

### updateSelectedUINode (method on `*ClickableRectangle`)
- Signature: `func (*ClickableRectangle) updateSelectedUINode()`
- Exported: `false`
- Control-flow features: `if`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
