# gRPCInHelpers.go

## File Overview
- Path: `grpc_in/gRPCInHelpers.go`
- Package: `grpc_in`
- Functions/Methods: `3`
- Imports: `3`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `IsClientUsingCorrectTestDataProtoFileVersion`
- `SetLogger`

## Imports
- `FenixTesterGui/common_code`
- `github.com/jlambert68/FenixGrpcApi/FenixUserGui/fenixUserGuiGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- `highestFenixProtoFileVersion`

## Functions and Methods
### IsClientUsingCorrectTestDataProtoFileVersion
- Signature: `func IsClientUsingCorrectTestDataProtoFileVersion(callingClientUuid string, usedProtoFileVersion fenixUserGuiGrpcApi.CurrentFenixUserGuiProtoFileVersionEnum) returnMessage *fenixUserGuiGrpcApi.AckNackResponse`
- Exported: `true`
- Control-flow features: `if`
- Doc: IsClientUsingCorrectTestDataProtoFileVersion ******************************************************************************************************************** Check if Calling Client is using correct proto-file version
- Internal calls: `getHighestFenixUserGuiServerProtoFileVersion`
- Selector calls: `fenixUserGuiGrpcApi.CurrentFenixUserGuiProtoFileVersionEnum`, `protoFileExpected.String`, `protoFileUsed.String`

### getHighestFenixUserGuiServerProtoFileVersion
- Signature: `func getHighestFenixUserGuiServerProtoFileVersion() int32`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: ******************************************************************************************************************* Get the highest FenixProtoFileVersionEnumeration

### SetLogger (method on `*GRPCInStruct`)
- Signature: `func (*GRPCInStruct) SetLogger(logger *logrus.Logger)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: SetLogger Set to use the same logger reference as is used by central part of system

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
