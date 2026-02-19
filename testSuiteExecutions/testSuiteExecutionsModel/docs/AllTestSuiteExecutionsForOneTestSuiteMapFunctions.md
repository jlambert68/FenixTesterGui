# AllTestSuiteExecutionsForOneTestSuiteMapFunctions.go

## File Overview
- Path: `testSuiteExecutions/testSuiteExecutionsModel/AllTestSuiteExecutionsForOneTestSuiteMapFunctions.go`
- Package: `testSuiteExecutionsModel`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `3`
- Imports: `4`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `AddTestSuiteExecutionsForOneTestSuiteUuid`
- `GetAllTestSuiteExecutionsForOneTestSuiteUuid`
- `GetSpecificTestSuiteExecutionForOneTestSuiteUuid`

## Imports
- `FenixTesterGui/common_code`
- `github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `sync`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- `allTestSuiteExecutionsMapMutex`

## Functions and Methods
### AddTestSuiteExecutionsForOneTestSuiteUuid
- Signature: `func AddTestSuiteExecutionsForOneTestSuiteUuid(testSuiteExecutionsModelRef *TestSuiteExecutionsModelStruct, testSuiteUuidMapKey TestSuiteUuidType, testSuiteExecutionUuidMapKey TestSuiteExecutionUuidType, testSuiteExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage, latestUniqueTestSuiteExecutionDatabaseRowId int32, moreRowsExists bool)`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: AddTestSuiteExecutionsForOneTestSuiteUuid Add a TestSuiteExecution to the map for TestSuiteExecutions per TestSuiteUuid
- External calls: `allTestSuiteExecutionsMapMutex.Lock`, `allTestSuiteExecutionsMapMutex.Unlock`

### GetAllTestSuiteExecutionsForOneTestSuiteUuid (method on `TestSuiteExecutionsModelStruct`)
- Signature: `func (TestSuiteExecutionsModelStruct) GetAllTestSuiteExecutionsForOneTestSuiteUuid(testSuiteUuidMapKey TestSuiteUuidType) (tempTestSuiteExecutionsList *[]*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage, existInMap bool)`
- Exported: `true`
- Control-flow features: `if, for/range, defer`
- Doc: GetAllTestSuiteExecutionsForOneTestSuiteUuid Get all TestSuiteExecutions for one TestSuiteUuid
- External calls: `allTestSuiteExecutionsMapMutex.RLock`, `allTestSuiteExecutionsMapMutex.RUnlock`

### GetSpecificTestSuiteExecutionForOneTestSuiteUuid (method on `TestSuiteExecutionsModelStruct`)
- Signature: `func (TestSuiteExecutionsModelStruct) GetSpecificTestSuiteExecutionForOneTestSuiteUuid(testSuiteUuidMapKey TestSuiteUuidType, testSuiteExecutionUuidMapKey TestSuiteExecutionUuidType) (tempTestSuiteExecution *fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage, existInMap bool)`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: GetSpecificTestSuiteExecutionForOneTestSuiteUuid Get one specific TestSuiteExecutions for one TestSuiteUuid
- External calls: `allTestSuiteExecutionsMapMutex.RLock`, `allTestSuiteExecutionsMapMutex.RUnlock`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
