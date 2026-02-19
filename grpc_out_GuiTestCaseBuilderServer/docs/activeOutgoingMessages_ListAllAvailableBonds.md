# activeOutgoingMessages_ListAllAvailableBonds.go

## File Overview
- Path: `grpc_out_GuiTestCaseBuilderServer/activeOutgoingMessages_ListAllAvailableBonds.go`
- Package: `grpc_out_GuiTestCaseBuilderServer`
- Functions/Methods: `1`
- Imports: `6`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `ListAllAvailableBonds`

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
### ListAllAvailableBonds (method on `*GRPCOutGuiTestCaseBuilderServerStruct`)
- Signature: `func (*GRPCOutGuiTestCaseBuilderServerStruct) ListAllAvailableBonds(userId string) returnMessage *fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: ListAllAvailableBonds - Get all Bonds that can be used within a TestCase
- Internal calls: `cancel`
- Selector calls: `context.Background`, `grpcOut.setConnectionToFenixGuiTestCaseBuilderServer_new`, `err.Error`, `fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum`, `grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion`, `context.WithTimeout`, `fenixGuiTestCaseCaseBuilderServerGrpcClient.ListAllAvailableBonds`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
