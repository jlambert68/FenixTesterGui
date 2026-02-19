# newOrEditTestDataPointGroupUI_AvailableTestDataComponent.go

## File Overview
- Path: `testDataSelector/newOrEditTestDataPointGroupUI/newOrEditTestDataPointGroupUI_AvailableTestDataComponent.go`
- Package: `newOrEditTestDataPointGroupUI`
- Functions/Methods: `2`
- Imports: `6`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/widget`
- `github.com/jlambert68/FenixScriptEngine/testDataEngine`
- `regexp`
- `strings`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### generateAllAvailablePointsListUIComponent
- Signature: `func generateAllAvailablePointsListUIComponent(newOrEditTestDataPointGroupWindow *fyne.Window, testDataModel *testDataEngine.TestDataModelStruct)`
- Exported: `false`
- Control-flow features: `none detected`
- Internal calls: `string`, `buildPopUpTableDataFromTestDataPointName`, `showTable`
- Selector calls: `widget.NewList`, `widget.NewLabel`, `fmt.Sprintf`, `allAvailablePointsList.UnselectAll`, `allAvailablePointsList.Refresh`, `selectedPointsList.Refresh`

### filterToRemoveNumberOfSimilarTestDataPointsInName
- Signature: `func filterToRemoveNumberOfSimilarTestDataPointsInName(dataPointNameToClean string) cleanedName string`
- Exported: `false`
- Control-flow features: `if`
- Doc: Removes the part of the name that specifies the number similar TestDataPoints, i.e. Sub Custody/Main TestData Area/SEK/AccTest/SE/CRDT/CH/Switzerland/BBH/EUR/EUR/SEK [2] ->
- Selector calls: `regexp.MustCompile`, `re.FindStringSubmatch`, `fmt.Println`, `strings.Trim`, `fmt.Sprintf`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
