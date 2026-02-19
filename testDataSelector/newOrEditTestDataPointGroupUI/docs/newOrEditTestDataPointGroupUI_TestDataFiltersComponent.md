# newOrEditTestDataPointGroupUI_TestDataFiltersComponent.go

## File Overview
- Path: `testDataSelector/newOrEditTestDataPointGroupUI/newOrEditTestDataPointGroupUI_TestDataFiltersComponent.go`
- Package: `newOrEditTestDataPointGroupUI`
- Functions/Methods: `2`
- Imports: `8`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `FenixTesterGui/soundEngine`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/widget`
- `github.com/jlambert68/FenixScriptEngine/testDataEngine`
- `sort`
- `strings`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### generateTestDataSelectionsUIComponent
- Signature: `func generateTestDataSelectionsUIComponent(testDataModel *testDataEngine.TestDataModelStruct, testDataModelMap map[testDataEngine.TestDataDomainUuidType]*testDataEngine.TestDataDomainModelStruct, newOrEditTestDataPointGroupWindow *fyne.Window)`
- Exported: `false`
- Control-flow features: `if, for/range, go, defer`
- Doc: *** Create the selection boxes for selecting TestDataValues values
- Internal calls: `clearTestDataFilterCheckBoxesButtonFunction`, `searchTestDataButtonFunction`, `sortDataPointsList`, `string`, `testDataPointIntersectionOfTwoSlices`
- Selector calls: `allAvailablePointsList.Refresh`, `container.NewBorder`, `container.NewHBox`, `container.NewScroll`, `container.NewVBox`, `domainsSelect.Refresh`, `domainsSelect.SetSelected`, `fmt.Sprintf`

### sortDataPointsList
- Signature: `func sortDataPointsList(dataPointListToBeSorted []testDataEngine.DataPointTypeForGroupsStruct) []testDataEngine.DataPointTypeForGroupsStruct`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Sort a slice with DataPoints
- Internal calls: `string`
- Selector calls: `sort.Slice`, `strings.Split`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
