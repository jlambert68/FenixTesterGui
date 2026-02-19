# updateTestSuiteExecutionsList.go

## File Overview
- Path: `testSuiteExecutions/listTestSuiteExecutionsModel/updateTestSuiteExecutionsList.go`
- Package: `listTestSuiteExecutionsModel`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `3`
- Imports: `6`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `LoadTestSuiteExecutionsThatCanBeViewedByUser`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/grpc_out_GuiExecutionServer`
- `FenixTesterGui/testSuiteExecutions/testSuiteExecutionsModel`
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
### LoadTestSuiteExecutionsThatCanBeViewedByUser
- Signature: `func LoadTestSuiteExecutionsThatCanBeViewedByUser(latestUniqueTestSuiteExecutionDatabaseRowId int32, onlyRetrieveLimitedSizedBatch bool, batchSize int32, retrieveAllExecutionsForSpecificTestSuiteUuid bool, specificTestSuiteUuid string, testSuiteExecutionFromTimeStamp time.Time, testSuiteExecutionToTimeStamp time.Time, loadAllDataFromDatabase bool, updateGuiChannel *chan bool)`
- Exported: `true`
- Control-flow features: `if, go`
- Doc: LoadTestSuiteExecutionsThatCanBeViewedByUser Load list with TestSuiteExecutions that the user can view
- Internal calls: `storeAllTestSuiteExecutionsForOneTestSuiteThatCanBeViewedByUser`, `storeOneTestSuiteExecutionPerTestSuiteThatCanBeViewedByUser`
- External calls: `listTestSuiteExecutionsResponse.GetAckNackResponse`, `listTestSuiteExecutionsResponse.GetLatestUniqueTestSuiteExecutionDatabaseRowId`, `listTestSuiteExecutionsResponse.GetMoreRowsExists`, `listTestSuiteExecutionsResponse.GetTestSuiteExecutionsList`

### storeAllTestSuiteExecutionsForOneTestSuiteThatCanBeViewedByUser
- Signature: `func storeAllTestSuiteExecutionsForOneTestSuiteThatCanBeViewedByUser(testSuiteExecutionsList []*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage, testSuiteExecutionsModelRef *testSuiteExecutionsModel.TestSuiteExecutionsModelStruct, latestUniqueTestSuiteExecutionDatabaseRowId int32, moreRowsExists bool)`
- Exported: `false`
- Control-flow features: `for/range`
- Doc: Store All TestSuiteExecutions for one TestSuite, That Can Be Viewed By User
- External calls: `testSuiteExecution.GetTestSuiteExecutionUuid`, `testSuiteExecution.GetTestSuiteUuid`, `testSuiteExecutionsModel.AddTestSuiteExecutionsForOneTestSuiteUuid`, `testSuiteExecutionsModel.TestSuiteExecutionUuidType`, `testSuiteExecutionsModel.TestSuiteUuidType`

### storeOneTestSuiteExecutionPerTestSuiteThatCanBeViewedByUser
- Signature: `func storeOneTestSuiteExecutionPerTestSuiteThatCanBeViewedByUser(testSuiteExecutionsList []*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage, testSuiteExecutionsModelRef *testSuiteExecutionsModel.TestSuiteExecutionsModelStruct, latestUniqueTestSuiteExecutionDatabaseRowId int32, moreRowsExists bool)`
- Exported: `false`
- Control-flow features: `for/range`
- Doc: Store TestSuiteExecutions That Can Be Viewed By User
- External calls: `testSuiteExecutions.GetTestSuiteExecutionUuid`, `testSuiteExecutionsModel.TestSuiteExecutionUuidType`, `testSuiteExecutionsModelRef.AddToTestSuiteExecutionsMap`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
