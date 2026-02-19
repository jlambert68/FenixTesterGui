# testCaseModel_loadTestCase.go

## File Overview
- Path: `testCase/testCaseModel/testCaseModel_loadTestCase.go`
- Package: `testCaseModel`
- Generated: `2026-02-19T14:23:17+01:00`
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
- Internal calls: `MetaDataSelectType`, `int`
- External calls: `detailedTestCaseResponse.GetDetailedTestCase`, `errors.New`, `fmt.Println`, `fmt.Sprintf`, `tempMetaDataGroupFromGrpc.GetMetaDataInGroupMap`, `testCaseModel.generateTestCaseForGrpcAndHash`, `testDataEngine.TestDataAreaNameType`, `testDataEngine.TestDataAreaUuidType`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
