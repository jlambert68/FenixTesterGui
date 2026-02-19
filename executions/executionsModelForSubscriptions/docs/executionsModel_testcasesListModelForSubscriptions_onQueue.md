# executionsModel_testcasesListModelForSubscriptions_onQueue.go

## File Overview
- Path: `executions/executionsModelForSubscriptions/executionsModel_testcasesListModelForSubscriptions_onQueue.go`
- Package: `executionsModelForSubscriptions`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `2`
- Imports: `7`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `InitiateSubscriptionModelForTestCaseOnExecutionQueue`
- `LoadAndCreateModelForTestCasesOnExecutionQueue`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/grpc_out_GuiExecutionServer`
- `errors`
- `fmt`
- `github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api`
- `strconv`
- `time`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### InitiateSubscriptionModelForTestCaseOnExecutionQueue
- Signature: `func InitiateSubscriptionModelForTestCaseOnExecutionQueue()`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: InitiateSubscriptionModelForTestCaseUnderExecution Initiate model and UI-model for Subscriptions for TestCase that waits on ExecutionQueue

### LoadAndCreateModelForTestCasesOnExecutionQueue (method on `*ExecutionsModelObjectStruct`)
- Signature: `func (*ExecutionsModelObjectStruct) LoadAndCreateModelForTestCasesOnExecutionQueue(domainsToInclude []string) err error`
- Exported: `true`
- Control-flow features: `if, for/range, returns error`
- Doc: LoadAndCreateModelForTestCasesOnExecutionQueue - Load TestCaseExecutions that waits on ExecutionQueue and transform them into model used
- Internal calls: `TestCaseExecutionMapKeyType`, `int`, `int32`
- External calls: `errors.New`, `fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum`, `fmt.Println`, `fmt.Sprintf`, `grpc_out_GuiExecutionServer.GetHighestFenixGuiExecutionServerProtoFileVersion`, `strconv.Itoa`, `time.Now`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
