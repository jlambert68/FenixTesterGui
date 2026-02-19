# testCaseExecutionUI_testCaseLExecutionsListForSubscriptions_finishedExecution.go

## File Overview
- Path: `executions/executionsUIForSubscriptions/testCaseExecutionUI_testCaseLExecutionsListForSubscriptions_finishedExecution.go`
- Package: `executionsUIForSubscriptions`
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
- Selector calls: `binding.BindStruct`, `headertable.NewSortingHeaderTable`, `container.NewMax`

### RemoveTestCaseExecutionFromFinishedTable
- Signature: `func RemoveTestCaseExecutionFromFinishedTable(testCaseExecutionsFinishedDataRowAdaptedForUiTableReference *executionsModelForSubscriptions.TestCaseExecutionsFinishedExecutionAdaptedForUiTableStruct, finishedExecutionsTableChannelCommand executionsModelForSubscriptions.FinishedExecutionsTableChannelCommandType) err error`
- Exported: `true`
- Control-flow features: `if, for/range, switch, go, returns error`
- Internal calls: `remove`, `ResizeTableColumns`
- Selector calls: `executionsModelForSubscriptions.TestCaseExecutionMapKeyType`, `fmt.Println`, `dataMapBinding.GetItem`, `errors.New`, `fmt.Sprintf`, `headertable.LoadFromFlashingTableCellsReferenceMap`, `headertable.FlashRowToBeRemoved`, `time.Sleep`

### MoveTestCaseExecutionFromUnderExecutionToFinishedExecution
- Signature: `func MoveTestCaseExecutionFromUnderExecutionToFinishedExecution(testCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference *executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionAdaptedForUiTableStruct, testCaseExecutionDetails *fenixExecutionServerGuiGrpcApi.TestCaseExecutionDetailsMessage) err error`
- Exported: `true`
- Control-flow features: `if, for/range, returns error`
- Doc: MoveTestCaseExecutionFromUnderExecutionToFinishedExecution Move TestCaseExecution from UnderExecution-table to FinishedExecutions-table
- Internal calls: `verifyThatTestCaseExecutionIsNotInUse`, `ResizeTableColumns`
- Selector calls: `executionsModelForSubscriptions.TestCaseExecutionMapKeyType`, `fmt.Println`, `errors.New`, `fmt.Sprintf`, `strconv.FormatBool`, `binding.BindStruct`, `headertable.LoadFromFlashingTableCellsReferenceMap`, `headertable.FlashAddedRow`

### StartFinishedExecutionsTableAddRemoveChannelReader
- Signature: `func StartFinishedExecutionsTableAddRemoveChannelReader()`
- Exported: `true`
- Control-flow features: `for/range, switch`
- Doc: StartFinishedExecutionsTableAddRemoveChannelReader Start the channel reader and process messages from the channel
- Internal calls: `MoveTestCaseExecutionFromUnderExecutionToFinishedExecution`, `RemoveTestCaseExecutionFromFinishedTable`
- Selector calls: `errors.New`, `fmt.Sprintf`, `fmt.Println`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
