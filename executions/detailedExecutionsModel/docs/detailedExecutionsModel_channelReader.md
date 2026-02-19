# detailedExecutionsModel_channelReader.go

## File Overview
- Path: `executions/detailedExecutionsModel/detailedExecutionsModel_channelReader.go`
- Package: `detailedExecutionsModel`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `6`
- Imports: `4`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `InitiateCommandChannelReaderForDetailedStatusUpdates`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/executions/detailedTestCaseExecutionUI_summaryTableDefinition`
- `github.com/sirupsen/logrus`
- `time`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### InitiateCommandChannelReaderForDetailedStatusUpdates
- Signature: `func InitiateCommandChannelReaderForDetailedStatusUpdates()`
- Exported: `true`
- Control-flow features: `go`
- Internal calls: `newRefreshTestCasesSummaryTableThrottler`
- External calls: `detailedExecutionsModelObject.startCommandChannelReaderForDetailedStatusUpdates`

### startCommandChannelReaderForDetailedStatusUpdates (method on `*DetailedExecutionsModelObjectStruct`)
- Signature: `func (*DetailedExecutionsModelObjectStruct) startCommandChannelReaderForDetailedStatusUpdates()`
- Exported: `false`
- Control-flow features: `if, for/range, switch`
- Doc: Channel reader which is used secure that status updates are handled in a structured way
- Internal calls: `int32`
- External calls: `detailedExecutionsModelObject.triggerProcessFullDetailedExecutionsStatusUpdate`, `detailedExecutionsModelObject.triggerProcessRemoveDetailedTestCaseExecution`, `detailedExecutionsModelObject.triggerProcessRetrieveFullDetailedTestCaseExecution`, `detailedExecutionsModelObject.triggerProcessStatusUpdateOfDetailedExecutionsStatus`

### triggerProcessFullDetailedExecutionsStatusUpdate (method on `*DetailedExecutionsModelObjectStruct`)
- Signature: `func (*DetailedExecutionsModelObjectStruct) triggerProcessFullDetailedExecutionsStatusUpdate(incomingChannelCommandAndMessage ChannelCommandDetailedExecutionsStruct)`
- Exported: `false`
- Control-flow features: `if`
- Doc: Updates all Executions status with information received after direct gRPC-call to GUiExecutionServer
- Internal calls: `CreateSummaryTableForDetailedTestCaseExecutionsList`
- External calls: `detailedExecutionsModelObject.processFullDetailedTestCaseExecutionsStatusUpdate`, `testCasesSummaryTableRefreshThrottler.RequestRefreshTestCasesSummaryTable`

### triggerProcessRemoveDetailedTestCaseExecution (method on `*DetailedExecutionsModelObjectStruct`)
- Signature: `func (*DetailedExecutionsModelObjectStruct) triggerProcessRemoveDetailedTestCaseExecution(testCaseExecutionKey string)`
- Exported: `false`
- Control-flow features: `if`
- Doc: Remove the DetailedTestCaseExecution
- Internal calls: `CreateSummaryTableForDetailedTestCaseExecutionsList`
- External calls: `detailedExecutionsModelObject.processRemoveDetailedTestCaseExecution`, `testCasesSummaryTableRefreshThrottler.RequestRefreshTestCasesSummaryTable`

### triggerProcessRetrieveFullDetailedTestCaseExecution (method on `*DetailedExecutionsModelObjectStruct`)
- Signature: `func (*DetailedExecutionsModelObjectStruct) triggerProcessRetrieveFullDetailedTestCaseExecution(testCaseExecutionKey string)`
- Exported: `false`
- Control-flow features: `if`
- Doc: Retrieve a full Detailed TestCaseExecution from GuiExecutionServer
- Internal calls: `CreateSummaryTableForDetailedTestCaseExecutionsList`
- External calls: `detailedExecutionsModelObject.processRetrieveFullDetailedTestCaseExecution`, `testCasesSummaryTableRefreshThrottler.RequestRefreshTestCasesSummaryTable`

### triggerProcessStatusUpdateOfDetailedExecutionsStatus (method on `*DetailedExecutionsModelObjectStruct`)
- Signature: `func (*DetailedExecutionsModelObjectStruct) triggerProcessStatusUpdateOfDetailedExecutionsStatus(incomingChannelCommandAndMessage ChannelCommandDetailedExecutionsStruct)`
- Exported: `false`
- Control-flow features: `if`
- Doc: Updates specific status information based on subscriptions updates from GuiExecutionServer
- Internal calls: `CreateSummaryTableForDetailedTestCaseExecutionsList`
- External calls: `detailedExecutionsModelObject.processStatusUpdateOfDetailedExecutionsStatus`, `testCasesSummaryTableRefreshThrottler.RequestRefreshTestCasesSummaryTable`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
