# activeOutgoingMessages_SendTesterGuiIsClosingDown.go

## File Overview
- Path: `grpc_out_GuiExecutionServer/activeOutgoingMessages_SendTesterGuiIsClosingDown.go`
- Package: `grpc_out_GuiExecutionServer`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `1`
- Imports: `7`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `SendTesterGuiIsClosingDown`

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
### SendTesterGuiIsClosingDown (method on `*GRPCOutGuiExecutionServerStruct`)
- Signature: `func (*GRPCOutGuiExecutionServerStruct) SendTesterGuiIsClosingDown() ackNackResponse *fenixExecutionServerGuiGrpcApi.AckNackResponse`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: SendTesterGuiIsClosingDown - Inform GuiExecutionServer that this TesterGui is closing down
- Internal calls: `GetHighestFenixGuiExecutionServerProtoFileVersion`, `cancel`
- External calls: `context.Background`, `context.WithTimeout`, `err.Error`, `fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum`, `fenixGuiExecutionServerGrpcClient.TesterGuiIsClosingDown`, `fmt.Sprintf`, `grpcOut.setConnectionToFenixGuiExecutionMessageServer_new`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
