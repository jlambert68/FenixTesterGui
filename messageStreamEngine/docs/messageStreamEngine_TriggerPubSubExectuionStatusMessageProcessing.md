# messageStreamEngine_TriggerPubSubExectuionStatusMessageProcessing.go

## File Overview
- Path: `messageStreamEngine/messageStreamEngine_TriggerPubSubExectuionStatusMessageProcessing.go`
- Package: `messageStreamEngine`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `1`
- Imports: `6`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `FenixTesterGui/common_code`
- `github.com/golang/protobuf/ptypes/timestamp`
- `github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `google.golang.org/protobuf/encoding/protojson`
- `strings`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### triggerProcessPubSubExecutionStatusMessage
- Signature: `func triggerProcessPubSubExecutionStatusMessage(pubSubMessage []byte) err error`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Internal calls: `string`
- External calls: `executionStatusMessagesPubSubMessage.GetExecutionsStatus`, `fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum`, `fenixExecutionServerGuiGrpcApi.TestInstructionExecutionStatusEnum`, `protojson.Unmarshal`, `strings.ReplaceAll`, `tempGetTestCaseExecutionsStatusPubSubMessage.GetBroadcastTimeStamp`, `tempGetTestCaseExecutionsStatusPubSubMessage.GetPreviousBroadcastTimeStamp`, `tempGetTestCaseExecutionsStatusPubSubMessage.GetTestCaseExecutionDetails`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
