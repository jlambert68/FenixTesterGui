# detaledTestCaseExecutionFunctions.go

## File Overview
- Path: `testCaseExecutions/testCaseExecutionsModel/detaledTestCaseExecutionFunctions.go`
- Package: `testCaseExecutionsModel`
- Functions/Methods: `1`
- Imports: `4`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `ExtractAndStoreRelationBetweenTestInstructionUuidAndTestCaseExecutionUuid`

## Imports
- `FenixTesterGui/common_code`
- `github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `strconv`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### ExtractAndStoreRelationBetweenTestInstructionUuidAndTestCaseExecutionUuid (method on `TestCaseExecutionsModelStruct`)
- Signature: `func (TestCaseExecutionsModelStruct) ExtractAndStoreRelationBetweenTestInstructionUuidAndTestCaseExecutionUuid(detailedTestCaseExecutionMapKey DetailedTestCaseExecutionMapKeyType) err error`
- Exported: `true`
- Control-flow features: `if, for/range, defer, returns error`
- Doc: ExtractAndStoreRelationBetweenTestInstructionUuidAndTestCaseExecutionUuid Extract relation between TestInstructionUuid and TestCaseExecutionUuid
- Internal calls: `RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType`, `TestInstructionExecutionUuidType`, `int`, `string`
- Selector calls: `testCaseExecutionsModel.ReadFromDetailedTestCaseExecutionsMap`, `detailedTestCaseExecutionsMapMutex.Lock`, `detailedTestCaseExecutionsMapMutex.Unlock`, `testInstructionExecution.GetTestInstructionExecutionBasicInformation`, `strconv.Itoa`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
