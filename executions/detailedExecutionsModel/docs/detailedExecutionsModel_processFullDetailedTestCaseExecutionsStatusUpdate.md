# detailedExecutionsModel_processFullDetailedTestCaseExecutionsStatusUpdate.go

## File Overview
- Path: `executions/detailedExecutionsModel/detailedExecutionsModel_processFullDetailedTestCaseExecutionsStatusUpdate.go`
- Package: `detailedExecutionsModel`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `1`
- Imports: `6`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/executions/detailedTestCaseExecutionUI_summaryTableDefinition`
- `github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `strconv`
- `time`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### processFullDetailedTestCaseExecutionsStatusUpdate (method on `*DetailedExecutionsModelObjectStruct`)
- Signature: `func (*DetailedExecutionsModelObjectStruct) processFullDetailedTestCaseExecutionsStatusUpdate(testCaseExecutionResponse *fenixExecutionServerGuiGrpcApi.TestCaseExecutionResponseMessage)`
- Exported: `false`
- Control-flow features: `if, for/range, defer`
- Doc: Updates all Executions status with information received after direct gRPC-call to GUiExecutionServer
- Internal calls: `int`, `int32`
- External calls: `detailedExecutionsModelObject.updateTestCaseExecutionsSummaryTable`, `detailedExecutionsModelObject.updateTestInstructionExecutionsSummaryTable`, `strconv.Itoa`, `time.Sleep`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
