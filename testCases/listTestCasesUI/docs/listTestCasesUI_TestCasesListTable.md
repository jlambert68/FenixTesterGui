# listTestCasesUI_TestCasesListTable.go

## File Overview
- Path: `testCases/listTestCasesUI/listTestCasesUI_TestCasesListTable.go`
- Package: `listTestCasesUI`
- Functions/Methods: `6`
- Imports: `20`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `RemoveTestCaseFromList`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/executions/detailedExecutionsModel`
- `FenixTesterGui/soundEngine`
- `FenixTesterGui/testCase/testCaseModel`
- `FenixTesterGui/testCases/listTestCasesModel`
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
### RemoveTestCaseFromList (method on `*ListTestCaseUIStruct`)
- Signature: `func (*ListTestCaseUIStruct) RemoveTestCaseFromList(testCaseUuidToBeRemoved string, testCasesModel *testCaseModel.TestCasesModelsStruct)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: RemoveTestCaseFromList Remove a TestCase from the List
- Selector calls: `listTestCaseUIObject.loadTestCaseListTableTable`, `listTestCaseUIObject.calculateAndSetCorrectColumnWidths`, `listTestCaseUIObject.updateTestCasesListTable`

### generateTestCasesListTable (method on `*ListTestCaseUIStruct`)
- Signature: `func (*ListTestCaseUIStruct) generateTestCasesListTable(testCasesModel *testCaseModel.TestCasesModelsStruct)`
- Exported: `false`
- Control-flow features: `if`
- Doc: Create the UI-list that holds the list of TestCasesMapPtr that the user can edit
- Internal calls: `newSortableHeaderLabel`
- Selector calls: `widget.NewTable`, `widget.NewLabel`, `png.Decode`, `bytes.NewReader`, `listTestCaseUIObject.updateTestCasesListTable`, `listTestCaseUIObject.calculateAndSetCorrectColumnWidths`

### updateTestCasesListTable (method on `*ListTestCaseUIStruct`)
- Signature: `func (*ListTestCaseUIStruct) updateTestCasesListTable(testCasesModel *testCaseModel.TestCasesModelsStruct)`
- Exported: `false`
- Control-flow features: `if, switch`
- Doc: Update the Table
- Internal calls: `newClickableTableLabel`, `int16`, `uint8`, `int32`, `int`
- Selector calls: `container.NewStack`, `canvas.NewRectangle`, `clickable.SetText`, `listTestCaseUIObject.openTestCase`, `fyne.CurrentApp`, `fmt.Sprintf`, `tempTestCasePtr.GetDomainUuid`, `tempTestCasePtr.GetDomainName`

### calculateAndSetCorrectColumnWidths (method on `*ListTestCaseUIStruct`)
- Signature: `func (*ListTestCaseUIStruct) calculateAndSetCorrectColumnWidths()`
- Exported: `false`
- Control-flow features: `if, for/range`
- Internal calls: `float32`
- Selector calls: `fyne.MeasureText`, `theme.TextSize`, `theme.Padding`

### loadTestCaseListTableTable (method on `*ListTestCaseUIStruct`)
- Signature: `func (*ListTestCaseUIStruct) loadTestCaseListTableTable(testCaseMetaDataFilterEntry *boolbits.Entry)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Internal calls: `int`, `int32`
- Selector calls: `boolbits.NewAllZerosEntry`, `fmt.Sprintf`, `err.Error`, `log.Fatalln`, `resultEntry.Or`, `tempTestCase.GetTestCaseUuid`, `resultEntry.And`, `resultEntry.Equals`

### sort2DStringSlice (method on `*ListTestCaseUIStruct`)
- Signature: `func (*ListTestCaseUIStruct) sort2DStringSlice(data [][]string, columnToSortOn int, sortingDirection SortingDirectionType)`
- Exported: `false`
- Control-flow features: `if, switch`
- Doc: Sort2DStringSlice sorts a 2D string slice by a specified column index. It assumes that the column index is valid for all rows in the slice.
- Selector calls: `sort.Slice`, `strconv.Atoi`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
