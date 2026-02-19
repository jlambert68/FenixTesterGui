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
- Selector calls: `testCaseExecutionsModel.TestCaseExecutionUuidType`, `testCaseExecutionsModelRef.DeleteFromTestCaseExecutionsMap`

### SortGuiTableOnCurrentColumnAndSorting
- Signature: `func SortGuiTableOnCurrentColumnAndSorting()`
- Exported: `true`
- Control-flow features: `switch`
- Doc: Sort the matrix, for GUI table, update the Gui. Use current table sorting and column, if exist
- Internal calls: `sortGuiTableAscendingOnTestCaseExecutionTimeStamp`, `sortGuiTableOnColumn`, `uint8`

### SortOrReverseSortGuiTable
- Signature: `func SortOrReverseSortGuiTable(sortInThisColumn uint8)`
- Exported: `true`
- Control-flow features: `if, switch`
- Doc: Sort the matrix, ascending, for GUI table, update the Gui for 'latestTestCaseExecutionTimeStampColumnNumber'
- Internal calls: `int`, `sortGuiTableOnColumn`

### calculateAndSetCorrectColumnWidths
- Signature: `func calculateAndSetCorrectColumnWidths()`
- Exported: `false`
- Control-flow features: `if, for/range`
- Internal calls: `float32`
- Selector calls: `fyne.Do`, `fyne.MeasureText`, `testCaseExecutionsListTable.Refresh`, `testCaseExecutionsListTable.SetColumnWidth`, `theme.Padding`, `theme.TextSize`

### generateTestCaseExecutionsListTable
- Signature: `func generateTestCaseExecutionsListTable(testCaseExecutionsModel *testCaseExecutionsModel.TestCaseExecutionsModelStruct)`
- Exported: `false`
- Control-flow features: `if`
- Doc: Create the UI-list that holds the list of TestCasesMapPtr that the user can edit
- Internal calls: `calculateAndSetCorrectColumnWidths`, `newSortableHeaderLabel`, `updateTestCaseExecutionsListTable`
- Selector calls: `bytes.NewReader`, `png.Decode`, `widget.NewLabel`, `widget.NewTable`

### loadTestCaseExecutionListTableTable
- Signature: `func loadTestCaseExecutionListTableTable(testCaseExecutionsModelObject *testCaseExecutionsModel.TestCaseExecutionsModelStruct, retrieveAllExecutionsForSpecificTestCaseUuid bool, specificTestCaseUuid string)`
- Exported: `false`
- Control-flow features: `if, for/range, switch, defer`
- Internal calls: `int`, `int32`, `sort2DStringSlice`
- Selector calls: `fmt.Sprintf`, `loadTestCaseExecutionListTableTableMutex.Lock`, `loadTestCaseExecutionListTableTableMutex.Unlock`, `log.Fatalln`, `sharedCode.ConvertGrpcTimeStampToStringForDB`, `strconv.Itoa`, `tempTestCaseExecution.GetDomainName`, `tempTestCaseExecution.GetDomainUUID`

### sort2DStringSlice
- Signature: `func sort2DStringSlice(data [][]string, columnToSortOn int, sortingDirection SortingDirectionType)`
- Exported: `false`
- Control-flow features: `if, switch`
- Doc: Sort2DStringSlice sorts a 2D string slice by a specified column index. It assumes that the column index is valid for all rows in the slice.
- Selector calls: `sort.Slice`, `strconv.Atoi`

### sortGuiTableAscendingOnTestCaseExecutionTimeStamp
- Signature: `func sortGuiTableAscendingOnTestCaseExecutionTimeStamp()`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Sort the matrix, ascending, for GUI table, update the Gui for 'latestTestCaseExecutionTimeStampColumnNumber'
- Internal calls: `sortGuiTableOnColumn`

### sortGuiTableOnColumn
- Signature: `func sortGuiTableOnColumn(columnNumber uint8, sortDirection SortingDirectionType)`
- Exported: `false`
- Control-flow features: `if, for/range, switch`
- Doc: Sort the matrix for GUI table, update the Gui and Set correct Sort-icon for sorted Header
- Internal calls: `calculateAndSetCorrectColumnWidths`, `int`, `loadTestCaseExecutionListTableTable`, `sort2DStringSlice`, `updateTestCaseExecutionsListTable`
- Selector calls: `fyne.Do`, `testCaseExecutionsListTable.Refresh`

### updateTestCaseExecutionsListTable
- Signature: `func updateTestCaseExecutionsListTable(testCaseExecutionsModel *testCaseExecutionsModel.TestCaseExecutionsModelStruct)`
- Exported: `false`
- Control-flow features: `if, switch, defer`
- Doc: Update the Table Update the Table
- Internal calls: `int`, `int16`, `int32`, `newClickableTableLabel`, `uint8`
- Selector calls: `alt.Hide`, `alt.Show`, `canvas.NewRectangle`, `cont.Refresh`, `container.NewStack`, `h.Refresh`, `lbl.Hide`, `lbl.SetText`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
