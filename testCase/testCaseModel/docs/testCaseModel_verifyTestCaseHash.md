# testCaseModel_verifyTestCaseHash.go

## File Overview
- Path: `testCase/testCaseModel/testCaseModel_verifyTestCaseHash.go`
- Package: `testCaseModel`
- Functions/Methods: `3`
- Imports: `3`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `TestCaseHashIsChangedSinceLoadedOrSaved`
- `VerifyLatestLoadedOrSavedTestCaseHashTowardsDatabase`
- `VerifyTestCaseHashTowardsDatabase`

## Imports
- `errors`
- `fmt`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### VerifyTestCaseHashTowardsDatabase (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) VerifyTestCaseHashTowardsDatabase(testCaseUuid string) (hashIsTheSame bool, err error)`
- Exported: `true`
- Control-flow features: `if, returns error`
- Doc: VerifyTestCaseHashTowardsDatabase - Verify if the Hash for the TestCase is the same as the one in the database
- Selector calls: `errors.New`, `fmt.Sprintf`, `fmt.Println`, `testCaseModel.generateTestCaseForGrpcAndHash`

### VerifyLatestLoadedOrSavedTestCaseHashTowardsDatabase (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) VerifyLatestLoadedOrSavedTestCaseHashTowardsDatabase(testCaseUuid string) (hashIsTheSame bool, err error)`
- Exported: `true`
- Control-flow features: `if, returns error`
- Doc: VerifyLatestLoadedOrSavedTestCaseHashTowardsDatabase - Verify if the latest Loaded or Saved Hash for the TestCase is the same as the one in the database
- Selector calls: `errors.New`, `fmt.Sprintf`, `fmt.Println`

### TestCaseHashIsChangedSinceLoadedOrSaved (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) TestCaseHashIsChangedSinceLoadedOrSaved(testCaseUuid string) (hashIsChanged bool, err error)`
- Exported: `true`
- Control-flow features: `if, returns error`
- Doc: TestCaseHashIsChangedSinceLoadedOrSaved - Verify if the Hash for the TestCase is the same as the one when TestCasesMapPtr was last Loaded or Saved
- Selector calls: `errors.New`, `fmt.Sprintf`, `fmt.Println`, `testCaseModel.generateTestCaseForGrpcAndHash`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
