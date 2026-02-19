# activeOutgoingMessages_SendListTestSuitesThatCanBeEdited.go

## File Overview
- Path: `grpc_out_GuiTestCaseBuilderServer/activeOutgoingMessages_SendListTestSuitesThatCanBeEdited.go`
- Package: `grpc_out_GuiTestCaseBuilderServer`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `1`
- Imports: `7`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `ListTestSuitesThatCanBeEditedResponseMessage`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/gcp`
- `context`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `google.golang.org/protobuf/types/known/timestamppb`
- `time`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### ListTestSuitesThatCanBeEditedResponseMessage (method on `*GRPCOutGuiTestCaseBuilderServerStruct`)
- Signature: `func (*GRPCOutGuiTestCaseBuilderServerStruct) ListTestSuitesThatCanBeEditedResponseMessage(testCaseUpdatedMinTimeStamp time.Time, testCaseExecutionUpdatedMinTimeStamp time.Time) returnMessage *fenixGuiTestCaseBuilderServerGrpcApi.ListTestSuitesResponseMessage`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: ListTestSuitesThatCanBeEditedResponseMessage - List all TestSuitesMapPtr that can be edited, used for producing a list that the used can chose TestCase to edit from
- Internal calls: `cancel`
- External calls: `context.Background`, `context.WithTimeout`, `err.Error`, `fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum`, `fenixGuiTestCaseCaseBuilderServerGrpcClient.ListTestSuitesThatCanBeEdited`, `grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion`, `grpcOut.setConnectionToFenixGuiTestCaseBuilderServer_new`, `timestamppb.New`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
