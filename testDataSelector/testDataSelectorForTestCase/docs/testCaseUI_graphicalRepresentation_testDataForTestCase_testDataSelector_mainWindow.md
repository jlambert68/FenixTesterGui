# testCaseUI_graphicalRepresentation_testDataForTestCase_testDataSelector_mainWindow.go

## File Overview
- Path: `testDataSelector/testDataSelectorForTestCase/testCaseUI_graphicalRepresentation_testDataForTestCase_testDataSelector_mainWindow.go`
- Package: `testDataSelectorForTestCase`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `1`
- Imports: `8`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `MainTestDataSelector`

## Imports
- `FenixTesterGui/testCase/testCaseModel`
- `FenixTesterGui/testDataSelector/newOrEditTestDataPointGroupUI`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/dialog`
- `fyne.io/fyne/v2/widget`
- `github.com/jlambert68/FenixScriptEngine/testDataEngine`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### MainTestDataSelector
- Signature: `func MainTestDataSelector(app fyne.App, parent fyne.Window, currentTestCasePtr *testCaseModel.TestCaseModelStruct, testCaseUuid string, testDataSelectorsContainer *fyne.Container, testDataPointGroupsSelectInMainTestCaseArea *widget.Select, testDataPointsForAGroupSelectInMainTestCaseArea *widget.Select, testDataRowsForTestDataPointsSelectInMainTestCaseArea *widget.Select)`
- Exported: `true`
- Control-flow features: `if, for/range, go`
- Internal calls: `string`, `testDataPointGroupsToStringSliceFunction`, `testDataPointsToStringSliceFunction`, `updateTestDataPointsForAGroupList`
- External calls: `app.NewWindow`, `container.NewBorder`, `container.NewHBox`, `container.NewHSplit`, `dialog.ShowConfirm`, `dialog.ShowInformation`, `fmt.Sprintf`, `fyne.NewSize`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
