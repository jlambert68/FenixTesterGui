# activeOutgoingMessages_SendInitiateTestSuiteExecutionWithAllTestDataSets.go

## File Overview
- Path: `grpc_out_GuiExecutionServer/activeOutgoingMessages_SendInitiateTestSuiteExecutionWithAllTestDataSets.go`
- Package: `grpc_out_GuiExecutionServer`
- Functions/Methods: `1`
- Imports: `7`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `SendInitiateTestSuiteExecutionWithAllTestDataSets`

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
### SendInitiateTestSuiteExecutionWithAllTestDataSets (method on `*GRPCOutGuiExecutionServerStruct`)
- Signature: `func (*GRPCOutGuiExecutionServerStruct) SendInitiateTestSuiteExecutionWithAllTestDataSets(initiateTestSuiteExecutionWithAllTestDataSetsRequestMessage *fenixExecutionServerGuiGrpcApi.InitiateTestSuiteExecutionWithAllTestDataSetsRequestMessage) initiateSingleTestSuiteExecutionResponseMessage *fenixExecutionServerGuiGrpcApi.InitiateSingleTestSuiteExecutionResponseMessage`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: SendInitiateTestSuiteExecutionWithAllTestDataSets - Initiate a TestSuiteExecution with all its TestDataSets
- Internal calls: `GetHighestFenixGuiExecutionServerProtoFileVersion`, `cancel`
- Selector calls: `context.Background`, `context.WithTimeout`, `err.Error`, `fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum`, `fenixGuiExecutionServerGrpcClient.InitiateTestSuiteExecutionWithAllTestDataSets`, `fmt.Sprintf`, `grpcOut.setConnectionToFenixGuiExecutionMessageServer_new`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
