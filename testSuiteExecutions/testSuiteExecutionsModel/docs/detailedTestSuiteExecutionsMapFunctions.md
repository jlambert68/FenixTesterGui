# detailedTestSuiteExecutionsMapFunctions.go

## File Overview
- Path: `testSuiteExecutions/testSuiteExecutionsModel/detailedTestSuiteExecutionsMapFunctions.go`
- Package: `testSuiteExecutionsModel`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `6`
- Imports: `5`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `AddToDetailedTestSuiteExecutionsMap`
- `ClearFlagRefreshOngoingOfDetailedTestSuiteExecution`
- `DeleteFromDetailedTestSuiteExecutionsMap`
- `GetNumberOfDetailedTestSuiteExecutionsRetrievedFromDatabase`
- `ReadFromDetailedTestSuiteExecutionsMap`
- `SetFlagRefreshOngoingOfDetailedTestSuiteExecution`

## Imports
- `FenixTesterGui/common_code`
- `fmt`
- `github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `sync`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- `detailedTestSuiteExecutionsMapMutex`

## Functions and Methods
### AddToDetailedTestSuiteExecutionsMap (method on `TestSuiteExecutionsModelStruct`)
- Signature: `func (TestSuiteExecutionsModelStruct) AddToDetailedTestSuiteExecutionsMap(detailedTestSuiteExecutionMapKey DetailedTestSuiteExecutionMapKeyType, detailedTestSuiteExecution *fenixExecutionServerGuiGrpcApi.TestSuiteExecutionResponseMessage)`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: AddToDetailedTestSuiteExecutionsMap Add to the DetailedTestSuiteExecutions-Map
- External calls: `detailedTestSuiteExecutionsMapMutex.Lock`, `detailedTestSuiteExecutionsMapMutex.Unlock`

### ClearFlagRefreshOngoingOfDetailedTestSuiteExecution (method on `TestSuiteExecutionsModelStruct`)
- Signature: `func (TestSuiteExecutionsModelStruct) ClearFlagRefreshOngoingOfDetailedTestSuiteExecution(detailedTestSuiteExecutionMapKey DetailedTestSuiteExecutionMapKeyType)`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: ClearFlagRefreshOngoingOfDetailedTestSuiteExecution Clear the flag there is an ongoing refresh of the DetailedTestSuiteExecution-data
- External calls: `detailedTestSuiteExecutionsMapMutex.Lock`, `detailedTestSuiteExecutionsMapMutex.Unlock`, `fmt.Println`

### DeleteFromDetailedTestSuiteExecutionsMap (method on `TestSuiteExecutionsModelStruct`)
- Signature: `func (TestSuiteExecutionsModelStruct) DeleteFromDetailedTestSuiteExecutionsMap(detailedTestSuiteExecutionMapKey DetailedTestSuiteExecutionMapKeyType)`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: DeleteFromDetailedTestSuiteExecutionsMap Delete from the DetailedTestSuiteExecutions-Map
- External calls: `detailedTestSuiteExecutionsMapMutex.Lock`, `detailedTestSuiteExecutionsMapMutex.Unlock`

### GetNumberOfDetailedTestSuiteExecutionsRetrievedFromDatabase (method on `TestSuiteExecutionsModelStruct`)
- Signature: `func (TestSuiteExecutionsModelStruct) GetNumberOfDetailedTestSuiteExecutionsRetrievedFromDatabase() numberOfDetailedTestSuiteExecutionsRetrievedFromDatabase int`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: GetNumberOfDetailedTestSuiteExecutionsRetrievedFromDatabase Read all from the DetailedTestSuiteExecutions-Map
- External calls: `detailedTestSuiteExecutionsMapMutex.RLock`, `detailedTestSuiteExecutionsMapMutex.RUnlock`

### ReadFromDetailedTestSuiteExecutionsMap (method on `TestSuiteExecutionsModelStruct`)
- Signature: `func (TestSuiteExecutionsModelStruct) ReadFromDetailedTestSuiteExecutionsMap(detailedTestSuiteExecutionMapKey DetailedTestSuiteExecutionMapKeyType) (detailedTestSuiteExecution *fenixExecutionServerGuiGrpcApi.TestSuiteExecutionResponseMessage, existInMap bool)`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: ReadFromDetailedTestSuiteExecutionsMap Read from the DetailedTestSuiteExecutions-Map
- External calls: `detailedTestSuiteExecutionsMapMutex.Lock`, `detailedTestSuiteExecutionsMapMutex.Unlock`, `fmt.Println`

### SetFlagRefreshOngoingOfDetailedTestSuiteExecution (method on `TestSuiteExecutionsModelStruct`)
- Signature: `func (TestSuiteExecutionsModelStruct) SetFlagRefreshOngoingOfDetailedTestSuiteExecution(detailedTestSuiteExecutionMapKey DetailedTestSuiteExecutionMapKeyType)`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: SetFlagRefreshOngoingOfDetailedTestSuiteExecution Set the flag there is an ongoing refresh of the DetailedTestSuiteExecution-data
- External calls: `detailedTestSuiteExecutionsMapMutex.Lock`, `detailedTestSuiteExecutionsMapMutex.Unlock`, `fmt.Println`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
