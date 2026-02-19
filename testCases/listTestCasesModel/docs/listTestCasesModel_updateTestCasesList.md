# listTestCasesModel_updateTestCasesList.go

## File Overview
- Path: `testCases/listTestCasesModel/listTestCasesModel_updateTestCasesList.go`
- Package: `listTestCasesModel`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `2`
- Imports: `8`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `LoadTestCaseThatCanBeEditedByUser`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/testCase/testCaseModel`
- `fmt`
- `github.com/jlambert68/Fast_BitFilter_MetaData/boolbits/boolbits`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `log`
- `time`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### LoadTestCaseThatCanBeEditedByUser
- Signature: `func LoadTestCaseThatCanBeEditedByUser(testCasesModeReference *testCaseModel.TestCasesModelsStruct, testCaseUpdatedMinTimeStamp time.Time, testCaseExecutionUpdatedMinTimeStamp time.Time)`
- Exported: `true`
- Control-flow features: `if`
- Doc: LoadTestCaseThatCanBeEditedByUser Load list with TestCasesMapPtr that the user can edit
- Internal calls: `storeTestCaseThatCanBeEditedByUser`
- External calls: `listTestCasesThatCanBeEditedResponseMessage.GetAckNackResponse`, `listTestCasesThatCanBeEditedResponseMessage.GetTestCasesThatCanBeEditedByUser`

### storeTestCaseThatCanBeEditedByUser
- Signature: `func storeTestCaseThatCanBeEditedByUser(testCasesThatCanBeEditedByUserAsSlice []*fenixGuiTestCaseBuilderServerGrpcApi.TestCaseThatCanBeEditedByUserMessage, testCasesModeReference *testCaseModel.TestCasesModelsStruct)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Store TestCasesMapPtr That Can Be Edited By User
- External calls: `boolbits.NewAllZerosEntry`, `boolbits.NewEntry`, `err.Error`, `fmt.Sprintf`, `latestTestCaseExecutionUpdatedMinTimeStamp.Before`, `latestTestCaseUpdatedMinTimeStamp.Before`, `log.Fatalln`, `resultsEntry.Or`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
