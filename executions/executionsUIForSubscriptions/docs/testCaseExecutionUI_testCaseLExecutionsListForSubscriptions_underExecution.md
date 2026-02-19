# testCaseExecutionUI_testCaseLExecutionsListForSubscriptions_underExecution.go

## File Overview
- Path: `executions/executionsUIForSubscriptions/testCaseExecutionUI_testCaseLExecutionsListForSubscriptions_underExecution.go`
- Package: `executionsUIForSubscriptions`
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
### CreateTableForTestCaseExecutionsUnderExecution
- Signature: `func CreateTableForTestCaseExecutionsUnderExecution() *fyne.Container`
- Exported: `true`
- Control-flow features: `for/range`
- Selector calls: `binding.BindStruct`, `headertable.NewSortingHeaderTable`, `container.NewMax`

### RemoveTestCaseExecutionFromUnderExecutionTable
- Signature: `func RemoveTestCaseExecutionFromUnderExecutionTable(testCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference *executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionAdaptedForUiTableStruct, underExecutionTableChannelCommand executionsModelForSubscriptions.UnderExecutionTableChannelCommandType) err error`
- Exported: `true`
- Control-flow features: `if, for/range, switch, go, returns error`
- Doc: RemoveTestCaseExecutionFromUnderExecutionTable Remove from both table-slice and from Map that Table-slice got its data from
- Internal calls: `remove`, `ResizeTableColumns`
- Selector calls: `executionsModelForSubscriptions.TestCaseExecutionMapKeyType`, `fmt.Println`, `dataMapBinding.GetItem`, `errors.New`, `fmt.Sprintf`, `headertable.LoadFromFlashingTableCellsReferenceMap`, `headertable.FlashRowToBeRemoved`, `time.Sleep`

### MoveTestCaseExecutionFromOnQueueToUnderExecution
- Signature: `func MoveTestCaseExecutionFromOnQueueToUnderExecution(testCaseExecutionsOnQueueDataRowAdaptedForUiTableReference *executionsModelForSubscriptions.TestCaseExecutionsOnQueueAdaptedForUiTableStruct, testCaseExecutionDetails *fenixExecutionServerGuiGrpcApi.TestCaseExecutionDetailsMessage) err error`
- Exported: `true`
- Control-flow features: `if, for/range, returns error`
- Doc: MoveTestCaseExecutionFromOnQueueToUnderExecution Move TestCaseExecution from OnQueue-table to UnderExecution-table
- Internal calls: `verifyThatTestCaseExecutionIsNotInUse`, `ResizeTableColumns`
- Selector calls: `executionsModelForSubscriptions.TestCaseExecutionMapKeyType`, `errors.New`, `fmt.Sprintf`, `fmt.Println`, `strconv.FormatBool`, `binding.BindStruct`, `headertable.LoadFromFlashingTableCellsReferenceMap`, `headertable.FlashAddedRow`

### AddTestCaseExecutionUnderExecutionTable
- Signature: `func AddTestCaseExecutionUnderExecutionTable(testCaseExecutionBasicInformation *fenixExecutionServerGuiGrpcApi.TestCaseExecutionBasicInformationMessage) err error`
- Exported: `true`
- Control-flow features: `if, returns error`
- Internal calls: `int`, `int32`, `verifyThatTestCaseExecutionIsNotInUse`
- Selector calls: `strconv.Itoa`, `executionsModelForSubscriptions.TestCaseExecutionMapKeyType`, `binding.BindStruct`

### StartUnderExecutionTableAddRemoveChannelReader
- Signature: `func StartUnderExecutionTableAddRemoveChannelReader()`
- Exported: `true`
- Control-flow features: `for/range, switch`
- Doc: StartUnderExecutionTableAddRemoveChannelReader Start the channel reader and process messages from the channel
- Internal calls: `MoveTestCaseExecutionFromOnQueueToUnderExecution`, `RemoveTestCaseExecutionFromUnderExecutionTable`
- Selector calls: `errors.New`, `fmt.Sprintf`, `fmt.Println`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
