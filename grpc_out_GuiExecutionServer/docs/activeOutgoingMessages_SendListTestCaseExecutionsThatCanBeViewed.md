# activeOutgoingMessages_SendListTestCaseExecutionsThatCanBeViewed.go

## File Overview
- Path: `grpc_out_GuiExecutionServer/activeOutgoingMessages_SendListTestCaseExecutionsThatCanBeViewed.go`
- Package: `grpc_out_GuiExecutionServer`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `1`
- Imports: `8`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `SendListTestCaseExecutionsThatCanBeViewed`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/gcp`
- `context`
- `fmt`
- `github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api`
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
### SendListTestCaseExecutionsThatCanBeViewed (method on `*GRPCOutGuiExecutionServerStruct`)
- Signature: `func (*GRPCOutGuiExecutionServerStruct) SendListTestCaseExecutionsThatCanBeViewed(latestUniqueTestCaseExecutionDatabaseRowId int32, onlyRetrieveLimitedSizedBatch bool, batchSize int32, retrieveAllExecutionsForSpecificTestCaseUuid bool, specificTestCaseUuid string, testCaseExecutionFromTimeStamp time.Time, testCaseExecutionToTimeStamp time.Time) listTestCaseExecutionsResponse *fenixExecutionServerGuiGrpcApi.ListTestCaseExecutionsResponse`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: SendListTestCaseExecutionsThatCanBeViewed - List all TestCaseExecutions that can be views, used for producing a list that the user has access to
- Internal calls: `GetHighestFenixGuiExecutionServerProtoFileVersion`, `cancel`
- External calls: `context.Background`, `context.WithTimeout`, `err.Error`, `fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum`, `fenixGuiExecutionServerGrpcClient.ListTestCaseExecutions`, `fmt.Println`, `grpcOut.setConnectionToFenixGuiExecutionMessageServer_new`, `timestamppb.New`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
