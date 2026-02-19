# testCaseModel_loadTestCase.go

## File Overview
- Path: `testCase/testCaseModel/testCaseModel_loadTestCase.go`
- Package: `testCaseModel`
- Functions/Methods: `1`
- Imports: `5`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `LoadFullTestCaseFromDatabase`

## Imports
- `FenixTesterGui/importFilesFromGitHub`
- `errors`
- `fmt`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `github.com/jlambert68/FenixScriptEngine/testDataEngine`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### LoadFullTestCaseFromDatabase (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) LoadFullTestCaseFromDatabase(testCaseUuid string) err error`
- Exported: `true`
- Control-flow features: `if, for/range, returns error`
- Doc: LoadFullTestCaseFromDatabase - Load the TestCase from the Database into model
- Internal calls: `int`, `MetaDataSelectType`
- Selector calls: `errors.New`, `fmt.Sprintf`, `fmt.Println`, `detailedTestCaseResponse.GetDetailedTestCase`, `testDataEngine.TestDataPointRowUuidType`, `testDataPointRowValueSummaryGrpc.GetTestDataPointRowUuid`, `testDataEngine.TestDataPointRowValuesSummaryType`, `testDataPointRowValueSummaryGrpc.GetTestDataPointRowValuesSummary`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
