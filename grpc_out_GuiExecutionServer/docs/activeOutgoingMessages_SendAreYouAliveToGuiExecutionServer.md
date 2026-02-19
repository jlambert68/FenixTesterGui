# activeOutgoingMessages_SendAreYouAliveToGuiExecutionServer.go

## File Overview
- Path: `grpc_out_GuiExecutionServer/activeOutgoingMessages_SendAreYouAliveToGuiExecutionServer.go`
- Package: `grpc_out_GuiExecutionServer`
- Functions/Methods: `1`
- Imports: `6`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `SendAreYouAliveToGuiExecutionServer`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/gcp`
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
### SendAreYouAliveToGuiExecutionServer (method on `*GRPCOutGuiExecutionServerStruct`)
- Signature: `func (*GRPCOutGuiExecutionServerStruct) SendAreYouAliveToGuiExecutionServer() returnMessage *fenixExecutionServerGuiGrpcApi.AckNackResponse`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: SendAreYouAliveToGuiExecutionServer - Check if 'GuiExecutionServer' is alive
- Internal calls: `GetHighestFenixGuiExecutionServerProtoFileVersion`, `cancel`
- Selector calls: `context.Background`, `grpcOut.setConnectionToFenixGuiExecutionMessageServer_new`, `err.Error`, `fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum`, `context.WithTimeout`, `fenixGuiExecutionServerGrpcClient.AreYouAlive`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
