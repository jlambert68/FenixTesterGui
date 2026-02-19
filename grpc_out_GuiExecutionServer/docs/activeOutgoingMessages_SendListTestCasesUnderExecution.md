# activeOutgoingMessages_SendListTestCasesUnderExecution.go

## File Overview
- Path: `grpc_out_GuiExecutionServer/activeOutgoingMessages_SendListTestCasesUnderExecution.go`
- Package: `grpc_out_GuiExecutionServer`
- Functions/Methods: `1`
- Imports: `7`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `SendListTestCasesUnderExecution`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/gcp`
- `fmt`
- `github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `golang.org/x/net/context`
- `time`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### SendListTestCasesUnderExecution (method on `*GRPCOutGuiExecutionServerStruct`)
- Signature: `func (*GRPCOutGuiExecutionServerStruct) SendListTestCasesUnderExecution(listTestCasesUnderExecutionRequest *fenixExecutionServerGuiGrpcApi.ListTestCasesUnderExecutionRequest) listTestCasesUnderExecutionResponse *fenixExecutionServerGuiGrpcApi.ListTestCasesUnderExecutionResponse`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: SendListTestCasesUnderExecution - Load all TestCaseExecutions that is under execution
- Internal calls: `GetHighestFenixGuiExecutionServerProtoFileVersion`, `cancel`
- Selector calls: `context.Background`, `grpcOut.setConnectionToFenixGuiExecutionMessageServer_new`, `err.Error`, `fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum`, `context.WithTimeout`, `fenixGuiExecutionServerGrpcClient.ListTestCasesUnderExecution`, `fmt.Sprintf`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
