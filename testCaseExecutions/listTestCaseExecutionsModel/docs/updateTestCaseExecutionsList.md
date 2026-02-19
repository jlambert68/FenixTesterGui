# updateTestCaseExecutionsList.go

## File Overview
- Path: `testCaseExecutions/listTestCaseExecutionsModel/updateTestCaseExecutionsList.go`
- Package: `listTestCaseExecutionsModel`
- Functions/Methods: `3`
- Imports: `7`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `LoadTestCaseExecutionsThatCanBeViewedByUser`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/grpc_out_GuiExecutionServer`
- `FenixTesterGui/testCaseExecutions/testCaseExecutionsModel`
- `fmt`
- `github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `time`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### LoadTestCaseExecutionsThatCanBeViewedByUser
- Signature: `func LoadTestCaseExecutionsThatCanBeViewedByUser(latestUniqueTestCaseExecutionDatabaseRowId int32, onlyRetrieveLimitedSizedBatch bool, batchSize int32, retrieveAllExecutionsForSpecificTestCaseUuid bool, specificTestCaseUuid string, testCaseExecutionFromTimeStamp time.Time, testCaseExecutionToTimeStamp time.Time, loadAllDataFromDatabase bool, updateGuiChannel *chan bool)`
- Exported: `true`
- Control-flow features: `if, go`
- Doc: LoadTestCaseExecutionsThatCanBeViewedByUser Load list with TestCaseExecutions that the user can view
- Internal calls: `storeAllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUser`, `storeOneTestCaseExecutionPerTestCaseThatCanBeViewedByUser`
- Selector calls: `fmt.Println`, `listTestCaseExecutionsResponse.GetAckNackResponse`, `listTestCaseExecutionsResponse.GetTestCaseExecutionsList`, `listTestCaseExecutionsResponse.GetLatestUniqueTestCaseExecutionDatabaseRowId`, `listTestCaseExecutionsResponse.GetMoreRowsExists`

### storeOneTestCaseExecutionPerTestCaseThatCanBeViewedByUser
- Signature: `func storeOneTestCaseExecutionPerTestCaseThatCanBeViewedByUser(testCaseExecutionsList []*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage, testCaseExecutionsModelRef *testCaseExecutionsModel.TestCaseExecutionsModelStruct, latestUniqueTestCaseExecutionDatabaseRowId int32, moreRowsExists bool)`
- Exported: `false`
- Control-flow features: `for/range`
- Doc: Store TestCaseExecutions That Can Be Viewed By User
- Selector calls: `testCaseExecutionsModelRef.AddToTestCaseExecutionsMap`, `testCaseExecutionsModel.TestCaseExecutionUuidType`, `testCaseExecutions.GetTestCaseExecutionUuid`

### storeAllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUser
- Signature: `func storeAllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUser(testCaseExecutionsList []*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage, testCaseExecutionsModelRef *testCaseExecutionsModel.TestCaseExecutionsModelStruct, latestUniqueTestCaseExecutionDatabaseRowId int32, moreRowsExists bool)`
- Exported: `false`
- Control-flow features: `for/range`
- Doc: Store All TestCaseExecutions for one TestCase, That Can Be Viewed By User
- Selector calls: `testCaseExecution.GetTestCaseUuid`, `testCaseExecution.GetTestCaseExecutionUuid`, `testCaseExecutionsModel.AddTestCaseExecutionsForOneTestCaseUuid`, `testCaseExecutionsModel.TestCaseUuidType`, `testCaseExecutionsModel.TestCaseExecutionUuidType`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
