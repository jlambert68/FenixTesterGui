# logPostAndValuesFunctions.go

## File Overview
- Path: `testCaseExecutions/testCaseExecutionsModel/logPostAndValuesFunctions.go`
- Package: `testCaseExecutionsModel`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `2`
- Imports: `4`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `ListLogPostsAndValuesForTestInstructionExecutions`

## Imports
- `FenixTesterGui/common_code`
- `github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `sort`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### ListLogPostsAndValuesForTestInstructionExecutions (method on `TestCaseExecutionsModelStruct`)
- Signature: `func (TestCaseExecutionsModelStruct) ListLogPostsAndValuesForTestInstructionExecutions(detailedTestCaseExecutionMapKey DetailedTestCaseExecutionMapKeyType, testInstructionLogPostMapKeys []TestInstructionExecutionLogPostMapKeyType) logPostAndValuesMessagesPtr *[]*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage`
- Exported: `true`
- Control-flow features: `if, for/range, defer`
- Doc: ListLogPostsAndValuesForTestInstructionExecutions List all LogPosts and Values for supplied TestInstructionExecutions. Log-posts are sorted on Logging DateTime
- Internal calls: `RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType`, `TestCaseExecutionUuidType`, `TestInstructionExecutionLogPostMapKeyType`, `sortLogPostsByTimestamp`
- External calls: `detailedTestCaseExecutionsMapMutex.Lock`, `detailedTestCaseExecutionsMapMutex.Unlock`, `testCaseExecutionsModel.GetTestInstructionExecutionUuidFromTestInstructionUuid`

### sortLogPostsByTimestamp
- Signature: `func sortLogPostsByTimestamp(logPostAndValuesMessagesPtr *[]*fenixExecutionServerGuiGrpcApi.LogPostAndValuesMessage, ascending bool)`
- Exported: `false`
- Control-flow features: `if`
- Doc: sortLogPostsByTimestamp sorts LogPostAndValuesMessages slice by LogPostTimeStamp Set ascending to true for ascending order, false for descending order
- External calls: `sort.SliceStable`, `timeI.After`, `timeI.Before`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
