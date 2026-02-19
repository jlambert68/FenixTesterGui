# listTestSuitesUI_TestSuitesListTable.go

## File Overview
- Path: `testSuites/listTestSuitesUI/listTestSuitesUI_TestSuitesListTable.go`
- Package: `listTestSuitesUI`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `6`
- Imports: `20`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `RemoveTestSuiteFromList`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/executions/detailedExecutionsModel`
- `FenixTesterGui/soundEngine`
- `FenixTesterGui/testCase/testCaseModel`
- `FenixTesterGui/testSuites/listTestSuitesModel`
- `bytes`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/theme`
- `fyne.io/fyne/v2/widget`
- `github.com/jlambert68/Fast_BitFilter_MetaData/boolbits/boolbits`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `image/color`
- `image/png`
- `log`
- `sort`
- `strconv`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### RemoveTestSuiteFromList (method on `*ListTestSuiteUIStruct`)
- Signature: `func (*ListTestSuiteUIStruct) RemoveTestSuiteFromList(testSuiteUuidToBeRemoved string, testCasesModel *testCaseModel.TestCasesModelsStruct)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: RemoveTestSuiteFromList Remove a TestSuite from the List
- External calls: `listTestSuiteUIObject.calculateAndSetCorrectColumnWidths`, `listTestSuiteUIObject.loadTestSuiteListTableTable`, `listTestSuiteUIObject.updateTestSuitesListTable`

### calculateAndSetCorrectColumnWidths (method on `*ListTestSuiteUIStruct`)
- Signature: `func (*ListTestSuiteUIStruct) calculateAndSetCorrectColumnWidths()`
- Exported: `false`
- Control-flow features: `if, for/range`
- Internal calls: `float32`
- External calls: `fyne.MeasureText`, `theme.Padding`, `theme.TextSize`

### generateTestSuitesListTable (method on `*ListTestSuiteUIStruct`)
- Signature: `func (*ListTestSuiteUIStruct) generateTestSuitesListTable(testCasesModel *testCaseModel.TestCasesModelsStruct)`
- Exported: `false`
- Control-flow features: `if`
- Doc: Create the UI-list that holds the list of TestSuitesMapPtr that the user can edit
- Internal calls: `newSortableHeaderLabel`
- External calls: `bytes.NewReader`, `listTestSuiteUIObject.calculateAndSetCorrectColumnWidths`, `listTestSuiteUIObject.updateTestSuitesListTable`, `png.Decode`, `widget.NewLabel`, `widget.NewTable`

### loadTestSuiteListTableTable (method on `*ListTestSuiteUIStruct`)
- Signature: `func (*ListTestSuiteUIStruct) loadTestSuiteListTableTable(testSuiteMetaDataFilterEntry *boolbits.Entry)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Internal calls: `int`, `int32`
- External calls: `boolbits.NewAllZerosEntry`, `err.Error`, `fmt.Sprintf`, `listTestSuiteUIObject.sort2DStringSlice`, `log.Fatalln`, `resultEntry.And`, `resultEntry.Equals`, `resultEntry.Or`

### sort2DStringSlice (method on `*ListTestSuiteUIStruct`)
- Signature: `func (*ListTestSuiteUIStruct) sort2DStringSlice(data [][]string, columnToSortOn int, sortingDirection SortingDirectionType)`
- Exported: `false`
- Control-flow features: `if, switch`
- Doc: Sort2DStringSlice sorts a 2D string slice by a specified column index. It assumes that the column index is valid for all rows in the slice.
- External calls: `sort.Slice`, `strconv.Atoi`

### updateTestSuitesListTable (method on `*ListTestSuiteUIStruct`)
- Signature: `func (*ListTestSuiteUIStruct) updateTestSuitesListTable(testCasesModel *testCaseModel.TestCasesModelsStruct)`
- Exported: `false`
- Control-flow features: `if, switch`
- Doc: Update the Table
- Internal calls: `int`, `int16`, `int32`, `newClickableTableLabel`, `uint8`
- External calls: `canvas.NewRectangle`, `clickable.Hide`, `clickable.SetText`, `clickable.Show`, `clickableContainer.Refresh`, `container.NewStack`, `fmt.Sprintf`, `fyne.CurrentApp`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
