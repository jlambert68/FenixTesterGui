# testSuiteModel_deleteTestSuiteFromDatabase.go

## File Overview
- Path: `testSuites/testSuitesModel/testSuiteModel_deleteTestSuiteFromDatabase.go`
- Package: `testSuitesModel`
- Functions/Methods: `1`
- Imports: `3`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `DeleteTestSuiteAtThisDate`

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
### DeleteTestSuiteAtThisDate (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) DeleteTestSuiteAtThisDate(testSuiteUuid string) err error`
- Exported: `true`
- Control-flow features: `if, returns error`
- Doc: DeleteTestSuiteAtThisDate - Mark the TestSuite as deleted by this date, in the Database
- Selector calls: `errors.New`, `fmt.Println`, `fmt.Sprintf`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
