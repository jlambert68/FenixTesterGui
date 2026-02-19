# testCaseExecutionUI_testCaseLExecutionsListForSubscriptions_finishedExecution.go

## File Overview
- Path: `executions/executionsUIForSubscriptions/testCaseExecutionUI_testCaseLExecutionsListForSubscriptions_finishedExecution.go`
- Package: `executionsUIForSubscriptions`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `4`
- Imports: `11`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `CreateTableForTestCaseExecutionsWithFinishedExecution`
- `MoveTestCaseExecutionFromUnderExecutionToFinishedExecution`
- `RemoveTestCaseExecutionFromFinishedTable`
- `StartFinishedExecutionsTableAddRemoveChannelReader`

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
### CreateTableForTestCaseExecutionsWithFinishedExecution
- Signature: `func CreateTableForTestCaseExecutionsWithFinishedExecution() *fyne.Container`
- Exported: `true`
- Control-flow features: `for/range`
- Doc: CreateTableForTestCaseExecutionsWithFinishedExecution Create bindings to the data used by the table and then create the UI-table itself
- External calls: `binding.BindStruct`, `container.NewMax`, `headertable.NewSortingHeaderTable`

### MoveTestCaseExecutionFromUnderExecutionToFinishedExecution
- Signature: `func MoveTestCaseExecutionFromUnderExecutionToFinishedExecution(testCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference *executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionAdaptedForUiTableStruct, testCaseExecutionDetails *fenixExecutionServerGuiGrpcApi.TestCaseExecutionDetailsMessage) err error`
- Exported: `true`
- Control-flow features: `if, for/range, returns error`
- Doc: MoveTestCaseExecutionFromUnderExecutionToFinishedExecution Move TestCaseExecution from UnderExecution-table to FinishedExecutions-table
- Internal calls: `ResizeTableColumns`, `verifyThatTestCaseExecutionIsNotInUse`
- External calls: `binding.BindStruct`, `errors.New`, `executionsModelForSubscriptions.TestCaseExecutionMapKeyType`, `fmt.Println`, `fmt.Sprintf`, `headertable.FlashAddedRow`, `headertable.LoadFromFlashingTableCellsReferenceMap`, `strconv.FormatBool`

### RemoveTestCaseExecutionFromFinishedTable
- Signature: `func RemoveTestCaseExecutionFromFinishedTable(testCaseExecutionsFinishedDataRowAdaptedForUiTableReference *executionsModelForSubscriptions.TestCaseExecutionsFinishedExecutionAdaptedForUiTableStruct, finishedExecutionsTableChannelCommand executionsModelForSubscriptions.FinishedExecutionsTableChannelCommandType) err error`
- Exported: `true`
- Control-flow features: `if, for/range, switch, go, returns error`
- Internal calls: `ResizeTableColumns`, `remove`
- External calls: `dataMapBinding.GetItem`, `errors.New`, `executionsModelForSubscriptions.TestCaseExecutionMapKeyType`, `fmt.Println`, `fmt.Sprintf`, `headertable.FlashRowToBeRemoved`, `headertable.LoadFromFlashingTableCellsReferenceMap`, `time.Sleep`

### StartFinishedExecutionsTableAddRemoveChannelReader
- Signature: `func StartFinishedExecutionsTableAddRemoveChannelReader()`
- Exported: `true`
- Control-flow features: `for/range, switch`
- Doc: StartFinishedExecutionsTableAddRemoveChannelReader Start the channel reader and process messages from the channel
- Internal calls: `MoveTestCaseExecutionFromUnderExecutionToFinishedExecution`, `RemoveTestCaseExecutionFromFinishedTable`
- External calls: `errors.New`, `fmt.Println`, `fmt.Sprintf`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
