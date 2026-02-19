# testSuiteUI_graphicalRepresentation_testDataForTestSuite_testDataSelector_mainWindow.go

## File Overview
- Path: `testDataSelector/testDataSelectorForTestSuite/testSuiteUI_graphicalRepresentation_testDataForTestSuite_testDataSelector_mainWindow.go`
- Package: `testDataSelectorForTestSuite`
- Functions/Methods: `1`
- Imports: `8`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `MainTestDataSelector`

## Imports
- `FenixTesterGui/testDataSelector/newOrEditTestDataPointGroupUI`
- `FenixTesterGui/testSuites/testSuitesModel`
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
- Signature: `func MainTestDataSelector(app fyne.App, parent fyne.Window, currentTestSuitePtr *testSuitesModel.TestSuiteModelStruct, testCaseUuid string, testDataSelectorsContainer *fyne.Container, testDataPointRadioGroupContainer *fyne.Container, generateTestDataAsRichTextCallBackFunctionRef func(), testDataPointGroupsSelectInMainTestSuiteArea *widget.Select, testDataPointsForAGroupSelectInMainTestSuiteArea *widget.Select, testDataRowsForTestDataPointsSelectInMainTestSuiteArea *widget.Select)`
- Exported: `true`
- Control-flow features: `if, for/range, go`
- Internal calls: `generateTestDataAsRichTextCallBackFunctionRef`, `string`, `testDataPointGroupsToStringSliceFunction`, `testDataPointsToStringSliceFunction`, `updateTestDataPointsForAGroupList`
- Selector calls: `app.NewWindow`, `container.NewBorder`, `container.NewHBox`, `container.NewHSplit`, `dialog.ShowConfirm`, `dialog.ShowInformation`, `fmt.Sprintf`, `fyne.Do`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
