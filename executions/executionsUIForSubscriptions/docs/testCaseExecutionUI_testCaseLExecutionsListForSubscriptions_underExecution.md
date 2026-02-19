# testCaseExecutionUI_testCaseLExecutionsListForSubscriptions_underExecution.go

## File Overview
- Path: `executions/executionsUIForSubscriptions/testCaseExecutionUI_testCaseLExecutionsListForSubscriptions_underExecution.go`
- Package: `executionsUIForSubscriptions`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `5`
- Imports: `11`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `AddTestCaseExecutionUnderExecutionTable`
- `CreateTableForTestCaseExecutionsUnderExecution`
- `MoveTestCaseExecutionFromOnQueueToUnderExecution`
- `RemoveTestCaseExecutionFromUnderExecutionTable`
- `StartUnderExecutionTableAddRemoveChannelReader`

## Imports
- `FenixTesterGui/executions/executionsModelForSubscriptions`
- `FenixTesterGui/headertable`
- `errors`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/data/binding`
- `fyne.io/fyne/v2/widget`
- `github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api`
- `strconv`
- `time`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### AddTestCaseExecutionUnderExecutionTable
- Signature: `func AddTestCaseExecutionUnderExecutionTable(testCaseExecutionBasicInformation *fenixExecutionServerGuiGrpcApi.TestCaseExecutionBasicInformationMessage) err error`
- Exported: `true`
- Control-flow features: `if, returns error`
- Internal calls: `int`, `int32`, `verifyThatTestCaseExecutionIsNotInUse`
- External calls: `binding.BindStruct`, `executionsModelForSubscriptions.TestCaseExecutionMapKeyType`, `strconv.Itoa`

### CreateTableForTestCaseExecutionsUnderExecution
- Signature: `func CreateTableForTestCaseExecutionsUnderExecution() *fyne.Container`
- Exported: `true`
- Control-flow features: `for/range`
- External calls: `binding.BindStruct`, `container.NewMax`, `headertable.NewSortingHeaderTable`

### MoveTestCaseExecutionFromOnQueueToUnderExecution
- Signature: `func MoveTestCaseExecutionFromOnQueueToUnderExecution(testCaseExecutionsOnQueueDataRowAdaptedForUiTableReference *executionsModelForSubscriptions.TestCaseExecutionsOnQueueAdaptedForUiTableStruct, testCaseExecutionDetails *fenixExecutionServerGuiGrpcApi.TestCaseExecutionDetailsMessage) err error`
- Exported: `true`
- Control-flow features: `if, for/range, returns error`
- Doc: MoveTestCaseExecutionFromOnQueueToUnderExecution Move TestCaseExecution from OnQueue-table to UnderExecution-table
- Internal calls: `ResizeTableColumns`, `verifyThatTestCaseExecutionIsNotInUse`
- External calls: `binding.BindStruct`, `errors.New`, `executionsModelForSubscriptions.TestCaseExecutionMapKeyType`, `fmt.Println`, `fmt.Sprintf`, `headertable.FlashAddedRow`, `headertable.LoadFromFlashingTableCellsReferenceMap`, `strconv.FormatBool`

### RemoveTestCaseExecutionFromUnderExecutionTable
- Signature: `func RemoveTestCaseExecutionFromUnderExecutionTable(testCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference *executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionAdaptedForUiTableStruct, underExecutionTableChannelCommand executionsModelForSubscriptions.UnderExecutionTableChannelCommandType) err error`
- Exported: `true`
- Control-flow features: `if, for/range, switch, go, returns error`
- Doc: RemoveTestCaseExecutionFromUnderExecutionTable Remove from both table-slice and from Map that Table-slice got its data from
- Internal calls: `ResizeTableColumns`, `remove`
- External calls: `dataMapBinding.GetItem`, `errors.New`, `executionsModelForSubscriptions.TestCaseExecutionMapKeyType`, `fmt.Println`, `fmt.Sprintf`, `headertable.FlashRowToBeRemoved`, `headertable.LoadFromFlashingTableCellsReferenceMap`, `time.Sleep`

### StartUnderExecutionTableAddRemoveChannelReader
- Signature: `func StartUnderExecutionTableAddRemoveChannelReader()`
- Exported: `true`
- Control-flow features: `for/range, switch`
- Doc: StartUnderExecutionTableAddRemoveChannelReader Start the channel reader and process messages from the channel
- Internal calls: `MoveTestCaseExecutionFromOnQueueToUnderExecution`, `RemoveTestCaseExecutionFromUnderExecutionTable`
- External calls: `errors.New`, `fmt.Println`, `fmt.Sprintf`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
