# listTestSuiteExectionsUI_TestSuiteExectionsListTable.go

## File Overview
- Path: `testSuiteExecutions/listTestSuiteExecutionsUI/listTestSuiteExectionsUI_TestSuiteExectionsListTable.go`
- Package: `listTestSuiteExecutionsUI`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `10`
- Imports: `18`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `RemoveTestSuiteExecutionFromList`
- `SortGuiTableOnCurrentColumnAndSorting`
- `SortOrReverseSortGuiTable`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/executions/detailedExecutionsModel`
- `FenixTesterGui/testSuiteExecutions/testSuiteExecutionsModel`
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
- `loadTestSuiteExecutionListTableTableMutex`
- `updateTestSuiteExecutionsListTableMutex`

## Functions and Methods
### RemoveTestSuiteExecutionFromList
- Signature: `func RemoveTestSuiteExecutionFromList(testSuiteExecutionUuidToBeRemoved string, testSuiteExecutionsModelRef *testSuiteExecutionsModel.TestSuiteExecutionsModelStruct)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: RemoveTestSuiteExecutionFromList Remove a TestSuiteExecution from the List
- External calls: `testSuiteExecutionsModel.TestSuiteExecutionUuidType`, `testSuiteExecutionsModelRef.DeleteFromTestSuiteExecutionsMap`

### SortGuiTableOnCurrentColumnAndSorting
- Signature: `func SortGuiTableOnCurrentColumnAndSorting()`
- Exported: `true`
- Control-flow features: `switch`
- Doc: Sort the matrix, for GUI table, update the Gui. Use current table sorting and column, if exist
- Internal calls: `sortGuiTableAscendingOnTestSuiteExecutionTimeStamp`, `sortGuiTableOnColumn`, `uint8`

### SortOrReverseSortGuiTable
- Signature: `func SortOrReverseSortGuiTable(sortInThisColumn uint8)`
- Exported: `true`
- Control-flow features: `if, switch`
- Doc: Sort the matrix, ascending, for GUI table, update the Gui for 'latestTestSuiteExecutionTimeStampColumnNumber'
- Internal calls: `int`, `sortGuiTableOnColumn`

### calculateAndSetCorrectColumnWidths
- Signature: `func calculateAndSetCorrectColumnWidths()`
- Exported: `false`
- Control-flow features: `if, for/range`
- Internal calls: `float32`
- External calls: `fyne.Do`, `fyne.MeasureText`, `testSuiteExecutionsListTable.Refresh`, `testSuiteExecutionsListTable.SetColumnWidth`, `theme.Padding`, `theme.TextSize`

### generateTestSuiteExecutionsListTable
- Signature: `func generateTestSuiteExecutionsListTable(testSuiteExecutionsModel *testSuiteExecutionsModel.TestSuiteExecutionsModelStruct)`
- Exported: `false`
- Control-flow features: `if`
- Doc: Create the UI-list that holds the list of TestSuitesMapPtr that the user can edit
- Internal calls: `calculateAndSetCorrectColumnWidths`, `newSortableHeaderLabel`, `updateTestSuiteExecutionsListTable`
- External calls: `bytes.NewReader`, `png.Decode`, `widget.NewLabel`, `widget.NewTable`

### loadTestSuiteExecutionListTableTable
- Signature: `func loadTestSuiteExecutionListTableTable(testSuiteExecutionsModelObject *testSuiteExecutionsModel.TestSuiteExecutionsModelStruct, retrieveAllExecutionsForSpecificTestSuiteUuid bool, specificTestSuiteUuid string)`
- Exported: `false`
- Control-flow features: `if, for/range, switch, defer`
- Internal calls: `int`, `int32`, `sort2DStringSlice`
- External calls: `fmt.Sprintf`, `loadTestSuiteExecutionListTableTableMutex.Lock`, `loadTestSuiteExecutionListTableTableMutex.Unlock`, `log.Fatalln`, `sharedCode.ConvertGrpcTimeStampToStringForDB`, `strconv.Itoa`, `tempTestSuiteExecution.GetDomainName`, `tempTestSuiteExecution.GetDomainUUID`

### sort2DStringSlice
- Signature: `func sort2DStringSlice(data [][]string, columnToSortOn int, sortingDirection SortingDirectionType)`
- Exported: `false`
- Control-flow features: `if, switch`
- Doc: Sort2DStringSlice sorts a 2D string slice by a specified column index. It assumes that the column index is valid for all rows in the slice.
- External calls: `sort.Slice`, `strconv.Atoi`

### sortGuiTableAscendingOnTestSuiteExecutionTimeStamp
- Signature: `func sortGuiTableAscendingOnTestSuiteExecutionTimeStamp()`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Sort the matrix, ascending, for GUI table, update the Gui for 'latestTestSuiteExecutionTimeStampColumnNumber'
- Internal calls: `sortGuiTableOnColumn`

### sortGuiTableOnColumn
- Signature: `func sortGuiTableOnColumn(columnNumber uint8, sortDirection SortingDirectionType)`
- Exported: `false`
- Control-flow features: `if, for/range, switch`
- Doc: Sort the matrix for GUI table, update the Gui and Set correct Sort-icon for sorted Header
- Internal calls: `calculateAndSetCorrectColumnWidths`, `int`, `loadTestSuiteExecutionListTableTable`, `sort2DStringSlice`, `updateTestSuiteExecutionsListTable`
- External calls: `fyne.Do`, `testSuiteExecutionsListTable.Refresh`

### updateTestSuiteExecutionsListTable
- Signature: `func updateTestSuiteExecutionsListTable(testSuiteExecutionsModel *testSuiteExecutionsModel.TestSuiteExecutionsModelStruct)`
- Exported: `false`
- Control-flow features: `if, switch, defer`
- Doc: Update the Table
- Internal calls: `int`, `int16`, `int32`, `newClickableTableLabel`, `uint8`
- External calls: `alt.Hide`, `alt.Show`, `canvas.NewRectangle`, `cont.Refresh`, `container.NewStack`, `fyne.Do`, `h.Refresh`, `lbl.Hide`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
