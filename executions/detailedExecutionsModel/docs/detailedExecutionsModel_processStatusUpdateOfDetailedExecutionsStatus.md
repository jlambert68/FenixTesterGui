# detailedExecutionsModel_processStatusUpdateOfDetailedExecutionsStatus.go

## File Overview
- Path: `executions/detailedExecutionsModel/detailedExecutionsModel_processStatusUpdateOfDetailedExecutionsStatus.go`
- Package: `detailedExecutionsModel`
- Functions/Methods: `2`
- Imports: `8`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `IsTestCaseExecutionStatusAnEndStatus`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/executions/detailedTestCaseExecutionUI_summaryTableDefinition`
- `errors`
- `fmt`
- `github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `strconv`
- `time`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### IsTestCaseExecutionStatusAnEndStatus
- Signature: `func IsTestCaseExecutionStatusAnEndStatus(testCaseExecutionStatus fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum) isTestCaseEndStatus bool`
- Exported: `true`
- Control-flow features: `switch`
- Doc: Check of the TestCaseExecutionStatus is classified as an EndStatus

### processStatusUpdateOfDetailedExecutionsStatus (method on `*DetailedExecutionsModelObjectStruct`)
- Signature: `func (*DetailedExecutionsModelObjectStruct) processStatusUpdateOfDetailedExecutionsStatus(testCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage *fenixExecutionServerGuiGrpcApi.
	TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage)`
- Exported: `false`
- Control-flow features: `if, for/range, defer`
- Doc: Updates specific status information based on subscriptions updates from GuiExecutionServer
- Internal calls: `IsTestCaseExecutionStatusAnEndStatus`, `int`, `int32`
- Selector calls: `detailedExecutionsModelObject.updateTestCaseExecutionsSummaryTable`, `detailedExecutionsModelObject.updateTestInstructionExecutionsSummaryTable`, `errors.New`, `fmt.Println`, `fmt.Sprintf`, `strconv.Itoa`, `tempExecutionStatusUpdateTimeStampMapKey.After`, `tempExecutionStatusUpdateTimeStampMapKey.String`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
