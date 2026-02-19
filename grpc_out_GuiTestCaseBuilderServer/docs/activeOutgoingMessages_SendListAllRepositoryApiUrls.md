# activeOutgoingMessages_SendListAllRepositoryApiUrls.go

## File Overview
- Path: `grpc_out_GuiTestCaseBuilderServer/activeOutgoingMessages_SendListAllRepositoryApiUrls.go`
- Package: `grpc_out_GuiTestCaseBuilderServer`
- Functions/Methods: `1`
- Imports: `6`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `SendListAllRepositoryApiUrls`

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
### SendListAllRepositoryApiUrls (method on `*GRPCOutGuiTestCaseBuilderServerStruct`)
- Signature: `func (*GRPCOutGuiTestCaseBuilderServerStruct) SendListAllRepositoryApiUrls() returnMessage *fenixGuiTestCaseBuilderServerGrpcApi.ListAllRepositoryApiUrlsResponseMessage`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: SendListAllRepositoryApiUrls - Load all Repository Api Urls for where Templates are stored
- Internal calls: `cancel`
- Selector calls: `context.Background`, `context.WithTimeout`, `err.Error`, `fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum`, `fenixGuiTestCaseCaseBuilderServerGrpcClient.ListAllRepositoryApiUrls`, `grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion`, `grpcOut.setConnectionToFenixGuiTestCaseBuilderServer_new`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
