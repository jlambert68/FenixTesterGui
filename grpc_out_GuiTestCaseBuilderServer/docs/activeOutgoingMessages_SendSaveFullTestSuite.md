# activeOutgoingMessages_SendSaveFullTestSuite.go

## File Overview
- Path: `grpc_out_GuiTestCaseBuilderServer/activeOutgoingMessages_SendSaveFullTestSuite.go`
- Package: `grpc_out_GuiTestCaseBuilderServer`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `1`
- Imports: `6`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `SendSaveFullTestSuite`

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
### SendSaveFullTestSuite (method on `*GRPCOutGuiTestCaseBuilderServerStruct`)
- Signature: `func (*GRPCOutGuiTestCaseBuilderServerStruct) SendSaveFullTestSuite(fullTestSuiteMessage *fenixGuiTestCaseBuilderServerGrpcApi.FullTestSuiteMessage) returnMessage *fenixGuiTestCaseBuilderServerGrpcApi.AckNackResponse`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: SendSaveFullTestSuite - Save full TestSuite to database
- Internal calls: `cancel`
- External calls: `context.Background`, `context.WithDeadline`, `err.Error`, `fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum`, `fenixGuiTestCaseCaseBuilderServerGrpcClient.SaveFullTestSuite`, `grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion`, `grpcOut.setConnectionToFenixGuiTestCaseBuilderServer_new`, `time.Now`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
