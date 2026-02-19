# listTestCaseExectionsUI_TableLabel.go

## File Overview
- Path: `testCaseExecutions/listTestCaseExecutionsUI/listTestCaseExectionsUI_TableLabel.go`
- Package: `listTestCaseExecutionsUI`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `6`
- Imports: `12`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `MouseIn`
- `MouseMoved`
- `MouseOut`
- `Tapped`
- `TappedSecondary`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/soundEngine`
- `FenixTesterGui/testCaseExecutions/testCaseExecutionsModel`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/driver/desktop`
- `fyne.io/fyne/v2/theme`
- `fyne.io/fyne/v2/widget`
- `github.com/sirupsen/logrus`
- `image/color`
- `time`

## Declared Types
- `clickableTableLabel`

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### MouseIn (method on `*clickableTableLabel`)
- Signature: `func (*clickableTableLabel) MouseIn(*desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `if`
- External calls: `currentRowThatMouseIsHoveringAboveMutex.Lock`, `currentRowThatMouseIsHoveringAboveMutex.Unlock`, `l.Refresh`

### MouseMoved (method on `*clickableTableLabel`)
- Signature: `func (*clickableTableLabel) MouseMoved(*desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`

### MouseOut (method on `*clickableTableLabel`)
- Signature: `func (*clickableTableLabel) MouseOut()`
- Exported: `true`
- Control-flow features: `if`
- External calls: `currentRowThatMouseIsHoveringAboveMutex.Lock`, `currentRowThatMouseIsHoveringAboveMutex.Unlock`, `l.Refresh`

### Tapped (method on `*clickableTableLabel`)
- Signature: `func (*clickableTableLabel) Tapped(e *fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `if, switch`
- External calls: `TestCaseInstructionPreViewObject.GenerateTestCaseExecutionPreviewContainer`, `l.onDoubleTap`, `loadAllTestCaseExecutionsForOneTestCaseButtonReference.Enable`, `loadAllTestCaseExecutionsForOneTestCaseButtonReference.Refresh`, `testCaseExecutionsListTable.Refresh`, `testCaseExecutionsModel.LoadDetailedTestCaseExecutionFromDatabase`, `time.Now`, `time.Since`

### TappedSecondary (method on `*clickableTableLabel`)
- Signature: `func (*clickableTableLabel) TappedSecondary(*fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `if`
- Doc: TappedSecondary Implement if you need right-click (secondary tap) actions.
- External calls: `clipboard.SetContent`, `fenixMasterWindow.Clipboard`, `fmt.Sprintf`, `fyne.CurrentApp`

### newClickableTableLabel
- Signature: `func newClickableTableLabel(text string, onDoubleTap func(), tempIsClickable bool, testCaseExecutionsModel *testCaseExecutionsModel.TestCaseExecutionsModelStruct) *clickableTableLabel`
- Exported: `false`
- Control-flow features: `none detected`
- External calls: `canvas.NewRectangle`, `l.ExtendBaseWidget`, `theme.TextSize`, `time.Now`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
