# detailedExecutionsModel_createSummaryTableForDetailedTestCaseExecutionsList.go

## File Overview
- Path: `executions/detailedExecutionsModel/detailedExecutionsModel_createSummaryTableForDetailedTestCaseExecutionsList.go`
- Package: `detailedExecutionsModel`
- Functions/Methods: `6`
- Imports: `11`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `CreateSummaryTableForDetailedTestCaseExecutionsList`
- `MouseIn`
- `MouseMoved`
- `MouseOut`

## Imports
- `FenixTesterGui/executions/detailedTestCaseExecutionUI_summaryTableDefinition`
- `errors`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/driver/desktop`
- `fyne.io/fyne/v2/layout`
- `fyne.io/fyne/v2/widget`
- `image/color`
- `log`

## Declared Types
- `ExecutionStatusColorMapStruct`
- `ExecutionStatusColorNameToNumberMapStruct`
- `spaceButton`
- `summaryTableForDetailedTestCaseExecutionsStruct`

## Declared Constants
- `backgroundStrokeWidth`

## Declared Variables
- `ExecutionStatusColorMap`
- `ExecutionStatusColorNameToNumberMap`
- `backgroundRectangleBaseColorForOddRows`
- `headerBackgroundRectangleBaseColorForEvenRows`
- `summaryTableForDetailedTestCaseExecutions`
- `testCaseRowBackgroundColorEvenRow`
- `testCaseRowBackgroundColorOddRow`

## Functions and Methods
### createExecutionSummary
- Signature: `func createExecutionSummary(testInstructionName string, testInstructionExecutionStatus int32) (tempItemExecutionSummary *fyne.Container, err error)`
- Exported: `false`
- Control-flow features: `if, returns error`
- Selector calls: `errors.New`, `fmt.Sprintf`, `fmt.Println`, `canvas.NewRectangle`, `container.New`, `layout.NewMaxLayout`

### CreateSummaryTableForDetailedTestCaseExecutionsList
- Signature: `func CreateSummaryTableForDetailedTestCaseExecutionsList() testcaseExecutionsSummaryReturnTable *fyne.Container`
- Exported: `true`
- Control-flow features: `if, for/range`
- Internal calls: `createExecutionSummary`, `int32`
- Selector calls: `layout.NewVBoxLayout`, `testcaseExecutionsSummaryTableContainer.Add`, `layout.NewHBoxLayout`, `testCaseRow.Add`, `testInstructionsForTestCase.Add`, `canvas.NewRectangle`, `log.Println`, `rowWithButtonsContainer.Add`

### newSpaceButton
- Signature: `func newSpaceButton() *spaceButton`
- Exported: `false`
- Control-flow features: `none detected`
- Selector calls: `mySpaceButton.ExtendBaseWidget`, `mySpaceButton.SetText`

### MouseIn (method on `*spaceButton`)
- Signature: `func (*spaceButton) MouseIn(x *desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: MouseIn is called when a desktop pointer enters the widget
- Selector calls: `fmt.Println`, `b.SetText`, `b.Refresh`

### MouseMoved (method on `*spaceButton`)
- Signature: `func (*spaceButton) MouseMoved(a *desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: MouseMoved is called when a desktop pointer hovers over the widget

### MouseOut (method on `*spaceButton`)
- Signature: `func (*spaceButton) MouseOut()`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: MouseOut is called when a desktop pointer exits the widget
- Selector calls: `fmt.Println`, `b.SetText`, `b.Refresh`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
