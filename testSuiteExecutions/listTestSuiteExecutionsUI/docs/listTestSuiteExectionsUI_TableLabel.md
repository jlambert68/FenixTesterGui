# listTestSuiteExectionsUI_TableLabel.go

## File Overview
- Path: `testSuiteExecutions/listTestSuiteExecutionsUI/listTestSuiteExectionsUI_TableLabel.go`
- Package: `listTestSuiteExecutionsUI`
- Functions/Methods: `6`
- Imports: `14`

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
- `FenixTesterGui/testCaseExecutions/listTestCaseExecutionsUI`
- `FenixTesterGui/testCaseExecutions/testCaseExecutionsModel`
- `FenixTesterGui/testSuiteExecutions/testSuiteExecutionsModel`
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
### newClickableTableLabel
- Signature: `func newClickableTableLabel(text string, onDoubleTap func(), tempIsClickable bool, testSuiteExecutionsModel *testSuiteExecutionsModel.TestSuiteExecutionsModelStruct) *clickableTableLabel`
- Exported: `false`
- Control-flow features: `none detected`
- Selector calls: `time.Now`, `canvas.NewRectangle`, `l.ExtendBaseWidget`, `theme.TextSize`

### Tapped (method on `*clickableTableLabel`)
- Signature: `func (*clickableTableLabel) Tapped(e *fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `if, switch`
- Selector calls: `time.Since`, `l.onDoubleTap`, `time.Now`, `fmt.Println`, `TestSuiteInstructionPreViewObject.GenerateTestSuiteExecutionPreviewContainer`, `testSuiteExecutionsListTable.Refresh`, `loadAllTestSuiteExecutionsForOneTestSuiteButtonReference.Enable`, `loadAllTestSuiteExecutionsForOneTestSuiteButtonReference.Refresh`

### TappedSecondary (method on `*clickableTableLabel`)
- Signature: `func (*clickableTableLabel) TappedSecondary(*fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `if`
- Doc: TappedSecondary Implement if you need right-click (secondary tap) actions.
- Selector calls: `fenixMasterWindow.Clipboard`, `clipboard.SetContent`, `fyne.CurrentApp`, `fmt.Sprintf`

### MouseIn (method on `*clickableTableLabel`)
- Signature: `func (*clickableTableLabel) MouseIn(*desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `if`
- Selector calls: `currentRowThatMouseIsHoveringAboveMutex.Lock`, `currentRowThatMouseIsHoveringAboveMutex.Unlock`, `l.Refresh`

### MouseMoved (method on `*clickableTableLabel`)
- Signature: `func (*clickableTableLabel) MouseMoved(*desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`

### MouseOut (method on `*clickableTableLabel`)
- Signature: `func (*clickableTableLabel) MouseOut()`
- Exported: `true`
- Control-flow features: `if`
- Selector calls: `currentRowThatMouseIsHoveringAboveMutex.Lock`, `currentRowThatMouseIsHoveringAboveMutex.Unlock`, `l.Refresh`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
