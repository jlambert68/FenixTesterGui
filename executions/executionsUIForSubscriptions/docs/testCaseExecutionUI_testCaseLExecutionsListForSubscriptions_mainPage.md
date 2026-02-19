# testCaseExecutionUI_testCaseLExecutionsListForSubscriptions_mainPage.go

## File Overview
- Path: `executions/executionsUIForSubscriptions/testCaseExecutionUI_testCaseLExecutionsListForSubscriptions_mainPage.go`
- Package: `executionsUIForSubscriptions`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `2`
- Imports: `1`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `CreateExecutionsListTabPageForSubsacriptions`
- `StartTableAddAndRemoveChannelReaders`

## Imports
- `fyne.io/fyne/v2`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### CreateExecutionsListTabPageForSubsacriptions (method on `*ExecutionsUIModelStruct`)
- Signature: `func (*ExecutionsUIModelStruct) CreateExecutionsListTabPageForSubsacriptions() executionsListTabPage *fyne.Container`
- Exported: `true`
- Control-flow features: `none detected`
- Internal calls: `CreateTableForTestCaseExecutionsOnQueue`, `CreateTableForTestCaseExecutionsUnderExecution`, `CreateTableForTestCaseExecutionsWithFinishedExecution`, `newThreePartAdaptiveSplit`
- External calls: `executionsListTabPage.Refresh`

### StartTableAddAndRemoveChannelReaders
- Signature: `func StartTableAddAndRemoveChannelReaders()`
- Exported: `true`
- Control-flow features: `go`
- Doc: StartTableAddAndRemoveChannelReaders Start Channel readers for testCases OnQueue, UnderExecutions or Finished Executions
- Internal calls: `StartFinishedExecutionsTableAddRemoveChannelReader`, `StartOnQueueTableAddRemoveChannelReader`, `StartUnderExecutionTableAddRemoveChannelReader`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
