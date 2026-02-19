# executionsModel_testcasesListModelForSubscriptions_underExecution.go

## File Overview
- Path: `executions/executionsModelForSubscriptions/executionsModel_testcasesListModelForSubscriptions_underExecution.go`
- Package: `executionsModelForSubscriptions`
- Functions/Methods: `2`
- Imports: `7`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `InitiateSubscriptionModelForTestCaseUnderExecution`
- `LoadAndCreateModelForTestCaseUnderExecutions`

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
### InitiateSubscriptionModelForTestCaseUnderExecution
- Signature: `func InitiateSubscriptionModelForTestCaseUnderExecution()`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: Initiate model and UI-model for Subscriptions for TestCase Under Execution

### LoadAndCreateModelForTestCaseUnderExecutions (method on `*ExecutionsModelObjectStruct`)
- Signature: `func (*ExecutionsModelObjectStruct) LoadAndCreateModelForTestCaseUnderExecutions(domainsToInclude []string) err error`
- Exported: `true`
- Control-flow features: `if, for/range, returns error`
- Doc: LoadAndCreateModelForTestCaseUnderExecutions - Load TestCaseExecutions that are Under Execution and transform them into model used
- Internal calls: `TestCaseExecutionMapKeyType`, `int`, `int32`
- Selector calls: `errors.New`, `fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum`, `fmt.Println`, `fmt.Sprintf`, `grpc_out_GuiExecutionServer.GetHighestFenixGuiExecutionServerProtoFileVersion`, `strconv.FormatBool`, `strconv.Itoa`, `time.Now`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
