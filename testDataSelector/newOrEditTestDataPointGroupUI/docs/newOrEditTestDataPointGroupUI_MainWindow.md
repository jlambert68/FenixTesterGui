# newOrEditTestDataPointGroupUI_MainWindow.go

## File Overview
- Path: `testDataSelector/newOrEditTestDataPointGroupUI/newOrEditTestDataPointGroupUI_MainWindow.go`
- Package: `newOrEditTestDataPointGroupUI`
- Functions/Methods: `2`
- Imports: `4`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `ShowNewOrEditGroupWindow`

## Imports
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/widget`
- `github.com/jlambert68/FenixScriptEngine/testDataEngine`

## Declared Types
- None

## Declared Constants
- `testDataDomainLabelText`
- `testDataTestAreaLabelText`

## Declared Variables
- None

## Functions and Methods
### ShowNewOrEditGroupWindow
- Signature: `func ShowNewOrEditGroupWindow(app fyne.App, parent fyne.Window, isNew bool, responseChannel *chan testDataEngine.ResponseChannelStruct, incomingGroupName testDataEngine.TestDataPointGroupNameType, newOrEditedChosenTestDataPointsThisGroupMapPtr *map[testDataEngine.TestDataPointGroupNameType]*testDataEngine.TestDataPointNameMapType, testDataForGroupObject *testDataEngine.TestDataForGroupObjectStruct)`
- Exported: `true`
- Control-flow features: `if`
- Internal calls: `generateAllAvailablePointsListUIComponent`, `generateTestDataSelectionsUIComponent`, `generateSelectedPointsListUIComponent`
- Selector calls: `parent.Hide`, `app.NewWindow`, `newOrEditTestDataPointGroupWindow.SetTitle`, `newOrEditTestDataPointGroupWindow.Resize`, `fyne.NewSize`, `newOrEditTestDataPointGroupWindow.CenterOnScreen`, `newOrEditTestDataPointGroupWindow.SetOnClosed`, `parent.Show`

### testDataPointIntersectionOfTwoSlices
- Signature: `func testDataPointIntersectionOfTwoSlices(firstSlice, secondSlice []testDataEngine.TestDataPointRowUuidType) []testDataEngine.TestDataPointRowUuidType`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: testDataPointIntersectionOfTwoSlices returns a new slice containing only the elements that appear in both a and b.

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
