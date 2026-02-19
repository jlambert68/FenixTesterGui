# activeOutgoingMessages_SendListTestSuiteExecutionsThatCanBeViewed.go

## File Overview
- Path: `grpc_out_GuiExecutionServer/activeOutgoingMessages_SendListTestSuiteExecutionsThatCanBeViewed.go`
- Package: `grpc_out_GuiExecutionServer`
- Functions/Methods: `1`
- Imports: `7`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `SendListTestSuiteExecutionsThatCanBeViewed`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/gcp`
- `context`
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
### SendListTestSuiteExecutionsThatCanBeViewed (method on `*GRPCOutGuiExecutionServerStruct`)
- Signature: `func (*GRPCOutGuiExecutionServerStruct) SendListTestSuiteExecutionsThatCanBeViewed(latestUniqueTestSuiteExecutionDatabaseRowId int32, onlyRetrieveLimitedSizedBatch bool, batchSize int32, retrieveAllExecutionsForSpecificTestSuiteUuid bool, specificTestSuiteUuid string, testSuiteExecutionFromTimeStamp time.Time, testSuiteExecutionToTimeStamp time.Time) listTestSuiteExecutionsResponse *fenixExecutionServerGuiGrpcApi.ListTestSuiteExecutionsResponse`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: SendListTestSuiteExecutionsThatCanBeViewed - List all TestSuiteExecutions that can be views, used for producing a list that the user has access to
- Internal calls: `GetHighestFenixGuiExecutionServerProtoFileVersion`, `cancel`
- Selector calls: `context.Background`, `context.WithTimeout`, `err.Error`, `fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum`, `fenixGuiExecutionServerGrpcClient.ListTestSuiteExecutions`, `grpcOut.setConnectionToFenixGuiExecutionMessageServer_new`, `timestamppb.New`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
