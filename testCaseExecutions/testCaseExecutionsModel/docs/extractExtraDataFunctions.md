# extractExtraDataFunctions.go

## File Overview
- Path: `testCaseExecutions/testCaseExecutionsModel/extractExtraDataFunctions.go`
- Package: `testCaseExecutionsModel`
- Functions/Methods: `1`
- Imports: `4`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `ExtractAndStoreLogPostsAndValuesFromDetailedTestCaseExecution`

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
### ExtractAndStoreLogPostsAndValuesFromDetailedTestCaseExecution (method on `TestCaseExecutionsModelStruct`)
- Signature: `func (TestCaseExecutionsModelStruct) ExtractAndStoreLogPostsAndValuesFromDetailedTestCaseExecution(detailedTestCaseExecutionMapKey DetailedTestCaseExecutionMapKeyType) err error`
- Exported: `true`
- Control-flow features: `if, for/range, defer, returns error`
- Doc: ExtractAndStoreLogPostsAndValuesFromDetailedTestCaseExecution Extracts all extra data that will be presented to the user in GUI, ie the explorer-tabs
- Internal calls: `AttributeNameMapKeyType`, `RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType`, `RunTimeUpdatedAttributeValueType`, `TestInstructionExecutionAttributeRunTimeUpdatedMapKeyType`, `TestInstructionExecutionDetailsMapKeyType`, `TestInstructionExecutionLogPostMapKeyType`, `TestInstructionExecutionUuidType`, `int`
- Selector calls: `detailedTestCaseExecutionsMapMutex.Lock`, `detailedTestCaseExecutionsMapMutex.Unlock`, `strconv.Itoa`, `testCaseExecutionsModel.ReadFromDetailedTestCaseExecutionsMap`, `testInstructionExecution.GetTestInstructionExecutionBasicInformation`, `testInstructionExecution.GetTestInstructionExecutionsInformation`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
