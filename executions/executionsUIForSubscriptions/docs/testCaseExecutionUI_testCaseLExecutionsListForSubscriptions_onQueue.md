# testCaseExecutionUI_testCaseLExecutionsListForSubscriptions_onQueue.go

## File Overview
- Path: `executions/executionsUIForSubscriptions/testCaseExecutionUI_testCaseLExecutionsListForSubscriptions_onQueue.go`
- Package: `executionsUIForSubscriptions`
- Functions/Methods: `4`
- Imports: `11`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `AddTestCaseExecutionToOnQueueTable`
- `CreateTableForTestCaseExecutionsOnQueue`
- `RemoveTestCaseExecutionFromOnQueueTable`
- `StartOnQueueTableAddRemoveChannelReader`

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
### CreateTableForTestCaseExecutionsOnQueue
- Signature: `func CreateTableForTestCaseExecutionsOnQueue() *fyne.Container`
- Exported: `true`
- Control-flow features: `for/range`
- Selector calls: `binding.BindStruct`, `headertable.NewSortingHeaderTable`, `container.NewMax`

### RemoveTestCaseExecutionFromOnQueueTable
- Signature: `func RemoveTestCaseExecutionFromOnQueueTable(testCaseExecutionsOnQueueDataRowAdaptedForUiTableReference *executionsModelForSubscriptions.TestCaseExecutionsOnQueueAdaptedForUiTableStruct, onQueueTableChannelCommand executionsModelForSubscriptions.OnQueueTableChannelCommandType) err error`
- Exported: `true`
- Control-flow features: `if, for/range, switch, go, returns error`
- Internal calls: `remove`, `ResizeTableColumns`
- Selector calls: `executionsModelForSubscriptions.TestCaseExecutionMapKeyType`, `fmt.Println`, `dataMapBinding.GetItem`, `errors.New`, `fmt.Sprintf`, `headertable.LoadFromFlashingTableCellsReferenceMap`, `headertable.FlashRowToBeRemoved`, `time.Sleep`

### AddTestCaseExecutionToOnQueueTable
- Signature: `func AddTestCaseExecutionToOnQueueTable(testCaseExecutionBasicInformation *fenixExecutionServerGuiGrpcApi.TestCaseExecutionBasicInformationMessage) err error`
- Exported: `true`
- Control-flow features: `if, for/range, returns error`
- Internal calls: `int`, `int32`, `verifyThatTestCaseExecutionIsNotInUse`, `ResizeTableColumns`
- Selector calls: `strconv.Itoa`, `executionsModelForSubscriptions.TestCaseExecutionMapKeyType`, `binding.BindStruct`, `headertable.LoadFromFlashingTableCellsReferenceMap`, `headertable.FlashAddedRow`

### StartOnQueueTableAddRemoveChannelReader
- Signature: `func StartOnQueueTableAddRemoveChannelReader()`
- Exported: `true`
- Control-flow features: `for/range, switch`
- Doc: StartOnQueueTableAddRemoveChannelReader Start the channel reader and process messages from the channel
- Internal calls: `AddTestCaseExecutionToOnQueueTable`, `RemoveTestCaseExecutionFromOnQueueTable`
- Selector calls: `errors.New`, `fmt.Sprintf`, `fmt.Println`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
