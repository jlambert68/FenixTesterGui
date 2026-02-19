# activeOutgoingMessages_SendTesterGuiIsStartingUp.go

## File Overview
- Path: `grpc_out_GuiExecutionServer/activeOutgoingMessages_SendTesterGuiIsStartingUp.go`
- Package: `grpc_out_GuiExecutionServer`
- Functions/Methods: `1`
- Imports: `7`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `SendTesterGuiIsStartingUp`

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
### SendTesterGuiIsStartingUp (method on `*GRPCOutGuiExecutionServerStruct`)
- Signature: `func (*GRPCOutGuiExecutionServerStruct) SendTesterGuiIsStartingUp() ackNackResponse *fenixExecutionServerGuiGrpcApi.AckNackResponse`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: SendTesterGuiIsStartingUp - Inform GuiExecutionServer that this TesterGui is starting up
- Internal calls: `GetHighestFenixGuiExecutionServerProtoFileVersion`, `cancel`
- Selector calls: `fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum`, `context.Background`, `grpcOut.setConnectionToFenixGuiExecutionMessageServer_new`, `err.Error`, `context.WithTimeout`, `fenixGuiExecutionServerGrpcClient.TesterGuiIsStartingUp`, `fmt.Sprintf`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
