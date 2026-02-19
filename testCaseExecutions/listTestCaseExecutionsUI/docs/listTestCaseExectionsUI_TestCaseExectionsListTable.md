# listTestCaseExectionsUI_TestCaseExectionsListTable.go

## File Overview
- Path: `testCaseExecutions/listTestCaseExecutionsUI/listTestCaseExectionsUI_TestCaseExectionsListTable.go`
- Package: `listTestCaseExecutionsUI`
- Functions/Methods: `10`
- Imports: `18`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `RemoveTestCaseExecutionFromList`
- `SortGuiTableOnCurrentColumnAndSorting`
- `SortOrReverseSortGuiTable`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/executions/detailedExecutionsModel`
- `FenixTesterGui/testCaseExecutions/testCaseExecutionsModel`
- `bytes`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/theme`
- `fyne.io/fyne/v2/widget`
- `github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `image/color`
- `image/png`
- `log`
- `sort`
- `strconv`
- `sync`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- `loadTestCaseExecutionListTableTableMutex`
- `updateTestCaseExecutionsListTableMutex`

## Functions and Methods
### RemoveTestCaseExecutionFromList
- Signature: `func RemoveTestCaseExecutionFromList(testCaseExecutionUuidToBeRemoved string, testCaseExecutionsModelRef *testCaseExecutionsModel.TestCaseExecutionsModelStruct)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: RemoveTestCaseExecutionFromList Remove a TestCaseExecution from the List
- Selector calls: `testCaseExecutionsModelRef.DeleteFromTestCaseExecutionsMap`, `testCaseExecutionsModel.TestCaseExecutionUuidType`

### generateTestCaseExecutionsListTable
- Signature: `func generateTestCaseExecutionsListTable(testCaseExecutionsModel *testCaseExecutionsModel.TestCaseExecutionsModelStruct)`
- Exported: `false`
- Control-flow features: `if`
- Doc: Create the UI-list that holds the list of TestCasesMapPtr that the user can edit
- Internal calls: `newSortableHeaderLabel`, `updateTestCaseExecutionsListTable`, `calculateAndSetCorrectColumnWidths`
- Selector calls: `widget.NewTable`, `widget.NewLabel`, `png.Decode`, `bytes.NewReader`

### updateTestCaseExecutionsListTable
- Signature: `func updateTestCaseExecutionsListTable(testCaseExecutionsModel *testCaseExecutionsModel.TestCaseExecutionsModelStruct)`
- Exported: `false`
- Control-flow features: `if, switch, defer`
- Doc: Update the Table Update the Table
- Internal calls: `newClickableTableLabel`, `int16`, `int`, `uint8`, `int32`
- Selector calls: `updateTestCaseExecutionsListTableMutex.Lock`, `updateTestCaseExecutionsListTableMutex.Unlock`, `container.NewStack`, `canvas.NewRectangle`, `lbl.Show`, `alt.Hide`, `lbl.SetText`, `lbl.Hide`

### calculateAndSetCorrectColumnWidths
- Signature: `func calculateAndSetCorrectColumnWidths()`
- Exported: `false`
- Control-flow features: `if, for/range`
- Internal calls: `float32`
- Selector calls: `fyne.MeasureText`, `theme.TextSize`, `fyne.Do`, `testCaseExecutionsListTable.SetColumnWidth`, `theme.Padding`, `testCaseExecutionsListTable.Refresh`

### loadTestCaseExecutionListTableTable
- Signature: `func loadTestCaseExecutionListTableTable(testCaseExecutionsModelObject *testCaseExecutionsModel.TestCaseExecutionsModelStruct, retrieveAllExecutionsForSpecificTestCaseUuid bool, specificTestCaseUuid string)`
- Exported: `false`
- Control-flow features: `if, for/range, switch, defer`
- Internal calls: `int`, `int32`, `sort2DStringSlice`
- Selector calls: `loadTestCaseExecutionListTableTableMutex.Lock`, `loadTestCaseExecutionListTableTableMutex.Unlock`, `testCaseExecutionsModelObject.ReadAllFromTestCaseExecutionsMap`, `testCaseExecutionsModelObject.GetAllTestCaseExecutionsForOneTestCaseUuid`, `testCaseExecutionsModel.TestCaseUuidType`, `fmt.Sprintf`, `tempTestCaseExecution.GetDomainName`, `tempTestCaseExecution.GetDomainUUID`

### sortGuiTableOnColumn
- Signature: `func sortGuiTableOnColumn(columnNumber uint8, sortDirection SortingDirectionType)`
- Exported: `false`
- Control-flow features: `if, for/range, switch`
- Doc: Sort the matrix for GUI table, update the Gui and Set correct Sort-icon for sorted Header
- Internal calls: `loadTestCaseExecutionListTableTable`, `sort2DStringSlice`, `int`, `calculateAndSetCorrectColumnWidths`, `updateTestCaseExecutionsListTable`
- Selector calls: `fyne.Do`, `testCaseExecutionsListTable.Refresh`

### sortGuiTableAscendingOnTestCaseExecutionTimeStamp
- Signature: `func sortGuiTableAscendingOnTestCaseExecutionTimeStamp()`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Sort the matrix, ascending, for GUI table, update the Gui for 'latestTestCaseExecutionTimeStampColumnNumber'
- Internal calls: `sortGuiTableOnColumn`

### SortOrReverseSortGuiTable
- Signature: `func SortOrReverseSortGuiTable(sortInThisColumn uint8)`
- Exported: `true`
- Control-flow features: `if, switch`
- Doc: Sort the matrix, ascending, for GUI table, update the Gui for 'latestTestCaseExecutionTimeStampColumnNumber'
- Internal calls: `int`, `sortGuiTableOnColumn`

### SortGuiTableOnCurrentColumnAndSorting
- Signature: `func SortGuiTableOnCurrentColumnAndSorting()`
- Exported: `true`
- Control-flow features: `switch`
- Doc: Sort the matrix, for GUI table, update the Gui. Use current table sorting and column, if exist
- Internal calls: `sortGuiTableOnColumn`, `uint8`, `sortGuiTableAscendingOnTestCaseExecutionTimeStamp`

### sort2DStringSlice
- Signature: `func sort2DStringSlice(data [][]string, columnToSortOn int, sortingDirection SortingDirectionType)`
- Exported: `false`
- Control-flow features: `if, switch`
- Doc: Sort2DStringSlice sorts a 2D string slice by a specified column index. It assumes that the column index is valid for all rows in the slice.
- Selector calls: `sort.Slice`, `strconv.Atoi`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
