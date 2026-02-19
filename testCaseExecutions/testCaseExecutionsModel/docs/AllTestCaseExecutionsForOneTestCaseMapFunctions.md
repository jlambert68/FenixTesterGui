# AllTestCaseExecutionsForOneTestCaseMapFunctions.go

## File Overview
- Path: `testCaseExecutions/testCaseExecutionsModel/AllTestCaseExecutionsForOneTestCaseMapFunctions.go`
- Package: `testCaseExecutionsModel`
- Functions/Methods: `3`
- Imports: `4`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `AddTestCaseExecutionsForOneTestCaseUuid`
- `GetAllTestCaseExecutionsForOneTestCaseUuid`
- `GetSpecificTestCaseExecutionForOneTestCaseUuid`

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
- `allTestCaseExecutionsMapMutex`

## Functions and Methods
### AddTestCaseExecutionsForOneTestCaseUuid
- Signature: `func AddTestCaseExecutionsForOneTestCaseUuid(testCaseExecutionsModelRef *TestCaseExecutionsModelStruct, testCaseUuidMapKey TestCaseUuidType, testCaseExecutionUuidMapKey TestCaseExecutionUuidType, testCaseExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage, latestUniqueTestCaseExecutionDatabaseRowId int32, moreRowsExists bool)`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: AddTestCaseExecutionsForOneTestCaseUuid Add a TestCaseExecution to the map for TestCaseExecutions per TestCaseUuid
- Selector calls: `allTestCaseExecutionsMapMutex.Lock`, `allTestCaseExecutionsMapMutex.Unlock`

### GetAllTestCaseExecutionsForOneTestCaseUuid (method on `TestCaseExecutionsModelStruct`)
- Signature: `func (TestCaseExecutionsModelStruct) GetAllTestCaseExecutionsForOneTestCaseUuid(testCaseUuidMapKey TestCaseUuidType) (tempTestCaseExecutionsList *[]*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage, existInMap bool)`
- Exported: `true`
- Control-flow features: `if, for/range, defer`
- Doc: GetAllTestCaseExecutionsForOneTestCaseUuid Get all TestCaseExecutions for one TestCaseUuid
- Selector calls: `allTestCaseExecutionsMapMutex.RLock`, `allTestCaseExecutionsMapMutex.RUnlock`

### GetSpecificTestCaseExecutionForOneTestCaseUuid (method on `TestCaseExecutionsModelStruct`)
- Signature: `func (TestCaseExecutionsModelStruct) GetSpecificTestCaseExecutionForOneTestCaseUuid(testCaseUuidMapKey TestCaseUuidType, testCaseExecutionUuidMapKey TestCaseExecutionUuidType) (tempTestCaseExecution *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage, existInMap bool)`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: GetSpecificTestCaseExecutionForOneTestCaseUuid Get one specific TestCaseExecutions for one TestCaseUuid
- Selector calls: `allTestCaseExecutionsMapMutex.RLock`, `allTestCaseExecutionsMapMutex.RUnlock`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
