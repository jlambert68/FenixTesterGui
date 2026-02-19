# activeOutgoingMessages_SendListAllAvailableTestInstructionsAndTestInstructionContainers.go

## File Overview
- Path: `grpc_out_GuiTestCaseBuilderServer/activeOutgoingMessages_SendListAllAvailableTestInstructionsAndTestInstructionContainers.go`
- Package: `grpc_out_GuiTestCaseBuilderServer`
- Functions/Methods: `1`
- Imports: `6`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `SendListAllAvailableTestInstructionsAndTestInstructionContainers`

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
### SendListAllAvailableTestInstructionsAndTestInstructionContainers (method on `*GRPCOutGuiTestCaseBuilderServerStruct`)
- Signature: `func (*GRPCOutGuiTestCaseBuilderServerStruct) SendListAllAvailableTestInstructionsAndTestInstructionContainers(gCPAuthenticatedUser string) returnMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: SendListAllAvailableTestInstructionsAndTestInstructionContainers - Get available TestInstructions and TestInstructionContainers
- Internal calls: `cancel`
- Selector calls: `fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum`, `grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion`, `context.Background`, `grpcOut.setConnectionToFenixGuiTestCaseBuilderServer_new`, `err.Error`, `context.WithTimeout`, `fenixGuiTestCaseCaseBuilderServerGrpcClient.ListAllAvailableTestInstructionsAndTestInstructionContainers`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
