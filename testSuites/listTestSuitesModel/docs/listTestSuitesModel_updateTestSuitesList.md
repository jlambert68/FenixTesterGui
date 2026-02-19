# listTestSuitesModel_updateTestSuitesList.go

## File Overview
- Path: `testSuites/listTestSuitesModel/listTestSuitesModel_updateTestSuitesList.go`
- Package: `listTestSuitesModel`
- Functions/Methods: `2`
- Imports: `9`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `LoadtestSuiteThatCanBeEditedByUser`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/testCase/testCaseModel`
- `FenixTesterGui/testSuites/testSuitesModel`
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
### LoadtestSuiteThatCanBeEditedByUser
- Signature: `func LoadtestSuiteThatCanBeEditedByUser(testCasesModeReference *testCaseModel.TestCasesModelsStruct, testSuiteUpdatedMinTimeStamp time.Time, testSuiteExecutionUpdatedMinTimeStamp time.Time)`
- Exported: `true`
- Control-flow features: `if`
- Doc: LoadTestSuiteThatCanBeEditedByUser Load list with TestSuitesMapPtr that the user can edit
- Internal calls: `storeTestSuiteThatCanBeEditedByUser`
- Selector calls: `listTestSuitesThatCanBeEditedResponseMessage.GetAckNackResponse`, `listTestSuitesThatCanBeEditedResponseMessage.GetBasicTestSuiteInformation`

### storeTestSuiteThatCanBeEditedByUser
- Signature: `func storeTestSuiteThatCanBeEditedByUser(TestSuitesThatCanBeEditedByUserAsSlice []*fenixGuiTestCaseBuilderServerGrpcApi.BasicTestSuiteInformationMessage, testCasesModeReference *testCaseModel.TestCasesModelsStruct)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Store TestSuitesMapPtr That Can Be Edited By User
- Selector calls: `testSuiteThatCanBeEditedByUser.GetTestSuitePreview`, `tempTestSuitePreview.GetSelectedTestSuiteMetaDataValuesMap`, `boolbits.NewAllZerosEntry`, `fmt.Sprintf`, `err.Error`, `log.Fatalln`, `boolbits.NewEntry`, `resultsEntry.Or`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
