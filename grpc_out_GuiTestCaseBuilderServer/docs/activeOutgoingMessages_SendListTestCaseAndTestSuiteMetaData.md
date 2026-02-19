# activeOutgoingMessages_SendListTestCaseAndTestSuiteMetaData.go

## File Overview
- Path: `grpc_out_GuiTestCaseBuilderServer/activeOutgoingMessages_SendListTestCaseAndTestSuiteMetaData.go`
- Package: `grpc_out_GuiTestCaseBuilderServer`
- Functions/Methods: `1`
- Imports: `6`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `SendListTestCaseAndTestSuiteMetaData`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/gcp`
- `context`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `time`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### SendListTestCaseAndTestSuiteMetaData (method on `*GRPCOutGuiTestCaseBuilderServerStruct`)
- Signature: `func (*GRPCOutGuiTestCaseBuilderServerStruct) SendListTestCaseAndTestSuiteMetaData() returnMessage *fenixGuiTestCaseBuilderServerGrpcApi.ListTestCaseAndTestSuiteMetaDataResponseMessage`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: SendListTestCaseAndTestSuiteMetaData - Load all TestCaseMetaData that the User can use when creating TestCasesMapPtr
- Internal calls: `cancel`
- Selector calls: `context.Background`, `context.WithTimeout`, `err.Error`, `fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum`, `fenixGuiTestCaseCaseBuilderServerGrpcClient.ListTestCaseAndTestSuiteMetaData`, `grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion`, `grpcOut.setConnectionToFenixGuiTestCaseBuilderServer_new`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
