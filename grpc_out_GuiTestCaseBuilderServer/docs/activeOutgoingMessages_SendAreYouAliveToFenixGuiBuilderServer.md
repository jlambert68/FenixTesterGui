# activeOutgoingMessages_SendAreYouAliveToFenixGuiBuilderServer.go

## File Overview
- Path: `grpc_out_GuiTestCaseBuilderServer/activeOutgoingMessages_SendAreYouAliveToFenixGuiBuilderServer.go`
- Package: `grpc_out_GuiTestCaseBuilderServer`
- Functions/Methods: `1`
- Imports: `6`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `SendAreYouAliveToFenixGuiBuilderServer`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/gcp`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
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
### SendAreYouAliveToFenixGuiBuilderServer (method on `*GRPCOutGuiTestCaseBuilderServerStruct`)
- Signature: `func (*GRPCOutGuiTestCaseBuilderServerStruct) SendAreYouAliveToFenixGuiBuilderServer() returnMessage *fenixGuiTestCaseBuilderServerGrpcApi.AckNackResponse`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: SendAreYouAliveToFenixGuiBuilderServer - Check if FenixGuiBuilderServer is alive
- Internal calls: `cancel`
- Selector calls: `context.Background`, `grpcOut.setConnectionToFenixGuiTestCaseBuilderServer_new`, `err.Error`, `fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum`, `grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion`, `context.WithTimeout`, `fenixGuiTestCaseCaseBuilderServerGrpcClient.AreYouAlive`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
