# updateTestCaseExecutionsList.go

## File Overview
- Path: `testCaseExecutions/listTestCaseExecutionsModel/updateTestCaseExecutionsList.go`
- Package: `listTestCaseExecutionsModel`
- Generated: `2026-02-19T14:23:17+01:00`
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
- External calls: `fmt.Println`, `listTestCaseExecutionsResponse.GetAckNackResponse`, `listTestCaseExecutionsResponse.GetLatestUniqueTestCaseExecutionDatabaseRowId`, `listTestCaseExecutionsResponse.GetMoreRowsExists`, `listTestCaseExecutionsResponse.GetTestCaseExecutionsList`

### storeAllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUser
- Signature: `func storeAllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUser(testCaseExecutionsList []*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage, testCaseExecutionsModelRef *testCaseExecutionsModel.TestCaseExecutionsModelStruct, latestUniqueTestCaseExecutionDatabaseRowId int32, moreRowsExists bool)`
- Exported: `false`
- Control-flow features: `for/range`
- Doc: Store All TestCaseExecutions for one TestCase, That Can Be Viewed By User
- External calls: `testCaseExecution.GetTestCaseExecutionUuid`, `testCaseExecution.GetTestCaseUuid`, `testCaseExecutionsModel.AddTestCaseExecutionsForOneTestCaseUuid`, `testCaseExecutionsModel.TestCaseExecutionUuidType`, `testCaseExecutionsModel.TestCaseUuidType`

### storeOneTestCaseExecutionPerTestCaseThatCanBeViewedByUser
- Signature: `func storeOneTestCaseExecutionPerTestCaseThatCanBeViewedByUser(testCaseExecutionsList []*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage, testCaseExecutionsModelRef *testCaseExecutionsModel.TestCaseExecutionsModelStruct, latestUniqueTestCaseExecutionDatabaseRowId int32, moreRowsExists bool)`
- Exported: `false`
- Control-flow features: `for/range`
- Doc: Store TestCaseExecutions That Can Be Viewed By User
- External calls: `testCaseExecutions.GetTestCaseExecutionUuid`, `testCaseExecutionsModel.TestCaseExecutionUuidType`, `testCaseExecutionsModelRef.AddToTestCaseExecutionsMap`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
