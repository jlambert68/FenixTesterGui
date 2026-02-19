# activeOutgoingMessages_SendListTestCasesThatCanBeEdited.go

## File Overview
- Path: `grpc_out_GuiTestCaseBuilderServer/activeOutgoingMessages_SendListTestCasesThatCanBeEdited.go`
- Package: `grpc_out_GuiTestCaseBuilderServer`
- Functions/Methods: `1`
- Imports: `7`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `ListTestCasesThatCanBeEditedResponseMessage`

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
### ListTestCasesThatCanBeEditedResponseMessage (method on `*GRPCOutGuiTestCaseBuilderServerStruct`)
- Signature: `func (*GRPCOutGuiTestCaseBuilderServerStruct) ListTestCasesThatCanBeEditedResponseMessage(testCaseUpdatedMinTimeStamp time.Time, testCaseExecutionUpdatedMinTimeStamp time.Time) returnMessage *fenixGuiTestCaseBuilderServerGrpcApi.ListTestCasesThatCanBeEditedResponseMessage`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: ListTestCasesThatCanBeEditedResponseMessage - List all TestCasesMapPtr that can be edited, used for producing a list that the used can chose TestCase to edit from
- Internal calls: `cancel`
- Selector calls: `context.Background`, `grpcOut.setConnectionToFenixGuiTestCaseBuilderServer_new`, `err.Error`, `fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum`, `grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion`, `timestamppb.New`, `context.WithTimeout`, `fenixGuiTestCaseCaseBuilderServerGrpcClient.ListTestCasesThatCanBeEdited`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
