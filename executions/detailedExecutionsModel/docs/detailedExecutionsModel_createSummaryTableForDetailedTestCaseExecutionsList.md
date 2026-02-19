# detailedExecutionsModel_createSummaryTableForDetailedTestCaseExecutionsList.go

## File Overview
- Path: `executions/detailedExecutionsModel/detailedExecutionsModel_createSummaryTableForDetailedTestCaseExecutionsList.go`
- Package: `detailedExecutionsModel`
- Generated: `2026-02-19T14:23:17+01:00`
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
### CreateSummaryTableForDetailedTestCaseExecutionsList
- Signature: `func CreateSummaryTableForDetailedTestCaseExecutionsList() testcaseExecutionsSummaryReturnTable *fyne.Container`
- Exported: `true`
- Control-flow features: `if, for/range`
- Internal calls: `createExecutionSummary`, `int32`
- External calls: `canvas.NewRectangle`, `container.New`, `layout.NewHBoxLayout`, `layout.NewMaxLayout`, `layout.NewVBoxLayout`, `log.Println`, `rowWithButtonsContainer.Add`, `testCaseRow.Add`

### MouseIn (method on `*spaceButton`)
- Signature: `func (*spaceButton) MouseIn(x *desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: MouseIn is called when a desktop pointer enters the widget
- External calls: `b.Refresh`, `b.SetText`, `fmt.Println`

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
- External calls: `b.Refresh`, `b.SetText`, `fmt.Println`

### createExecutionSummary
- Signature: `func createExecutionSummary(testInstructionName string, testInstructionExecutionStatus int32) (tempItemExecutionSummary *fyne.Container, err error)`
- Exported: `false`
- Control-flow features: `if, returns error`
- External calls: `canvas.NewRectangle`, `container.New`, `errors.New`, `fmt.Println`, `fmt.Sprintf`, `layout.NewMaxLayout`

### newSpaceButton
- Signature: `func newSpaceButton() *spaceButton`
- Exported: `false`
- Control-flow features: `none detected`
- External calls: `mySpaceButton.ExtendBaseWidget`, `mySpaceButton.SetText`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
