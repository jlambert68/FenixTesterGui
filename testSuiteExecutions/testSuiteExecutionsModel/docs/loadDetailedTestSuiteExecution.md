# loadDetailedTestSuiteExecution.go

## File Overview
- Path: `testSuiteExecutions/testSuiteExecutionsModel/loadDetailedTestSuiteExecution.go`
- Package: `testSuiteExecutionsModel`
- Functions/Methods: `1`
- Imports: `5`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `LoadDetailedTestSuiteExecutionFromDatabase`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/grpc_out_GuiExecutionServer`
- `github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `strconv`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### LoadDetailedTestSuiteExecutionFromDatabase
- Signature: `func LoadDetailedTestSuiteExecutionFromDatabase(testSuiteExecutionUuid string, testSuiteExecutionVersion uint32)`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: LoadDetailedTestSuiteExecutionFromDatabase Load all Detailed TestSuiteExecution-data for specific execution
- Internal calls: `DetailedTestSuiteExecutionMapKeyType`, `int`
- Selector calls: `TestSuiteExecutionsModel.AddToDetailedTestSuiteExecutionsMap`, `TestSuiteExecutionsModel.ClearFlagRefreshOngoingOfDetailedTestSuiteExecution`, `TestSuiteExecutionsModel.SetFlagRefreshOngoingOfDetailedTestSuiteExecution`, `fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum`, `getSingleTestSuiteExecutionResponse.GetAckNackResponse`, `getSingleTestSuiteExecutionResponse.GetTestSuiteExecutionResponse`, `grpc_out_GuiExecutionServer.GetHighestFenixGuiExecutionServerProtoFileVersion`, `strconv.Itoa`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
