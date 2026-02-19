# detailedTestCaseExecutionsMapFunctions.go

## File Overview
- Path: `testCaseExecutions/testCaseExecutionsModel/detailedTestCaseExecutionsMapFunctions.go`
- Package: `testCaseExecutionsModel`
- Functions/Methods: `6`
- Imports: `5`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `AddToDetailedTestCaseExecutionsMap`
- `ClearFlagRefreshOngoingOfDetailedTestCaseExecution`
- `DeleteFromDetailedTestCaseExecutionsMap`
- `GetNumberOfDetailedTestCaseExecutionsRetrievedFromDatabase`
- `ReadFromDetailedTestCaseExecutionsMap`
- `SetFlagRefreshOngoingOfDetailedTestCaseExecution`

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
- `detailedTestCaseExecutionsMapMutex`

## Functions and Methods
### ReadFromDetailedTestCaseExecutionsMap (method on `TestCaseExecutionsModelStruct`)
- Signature: `func (TestCaseExecutionsModelStruct) ReadFromDetailedTestCaseExecutionsMap(detailedTestCaseExecutionMapKey DetailedTestCaseExecutionMapKeyType) (detailedTestCaseExecution *fenixExecutionServerGuiGrpcApi.TestCaseExecutionResponseMessage, existInMap bool)`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: ReadFromDetailedTestCaseExecutionsMap Read from the DetailedTestCaseExecutions-Map
- Selector calls: `detailedTestCaseExecutionsMapMutex.Lock`, `detailedTestCaseExecutionsMapMutex.Unlock`, `fmt.Println`

### GetNumberOfDetailedTestCaseExecutionsRetrievedFromDatabase (method on `TestCaseExecutionsModelStruct`)
- Signature: `func (TestCaseExecutionsModelStruct) GetNumberOfDetailedTestCaseExecutionsRetrievedFromDatabase() numberOfDetailedTestCaseExecutionsRetrievedFromDatabase int`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: GetNumberOfDetailedTestCaseExecutionsRetrievedFromDatabase Read all from the DetailedTestCaseExecutions-Map
- Selector calls: `detailedTestCaseExecutionsMapMutex.RLock`, `detailedTestCaseExecutionsMapMutex.RUnlock`

### AddToDetailedTestCaseExecutionsMap (method on `TestCaseExecutionsModelStruct`)
- Signature: `func (TestCaseExecutionsModelStruct) AddToDetailedTestCaseExecutionsMap(detailedTestCaseExecutionMapKey DetailedTestCaseExecutionMapKeyType, detailedTestCaseExecution *fenixExecutionServerGuiGrpcApi.TestCaseExecutionResponseMessage)`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: AddToDetailedTestCaseExecutionsMap Add to the DetailedTestCaseExecutions-Map
- Selector calls: `detailedTestCaseExecutionsMapMutex.Lock`, `detailedTestCaseExecutionsMapMutex.Unlock`

### DeleteFromDetailedTestCaseExecutionsMap (method on `TestCaseExecutionsModelStruct`)
- Signature: `func (TestCaseExecutionsModelStruct) DeleteFromDetailedTestCaseExecutionsMap(detailedTestCaseExecutionMapKey DetailedTestCaseExecutionMapKeyType)`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: DeleteFromDetailedTestCaseExecutionsMap Delete from the DetailedTestCaseExecutions-Map
- Selector calls: `detailedTestCaseExecutionsMapMutex.Lock`, `detailedTestCaseExecutionsMapMutex.Unlock`

### SetFlagRefreshOngoingOfDetailedTestCaseExecution (method on `TestCaseExecutionsModelStruct`)
- Signature: `func (TestCaseExecutionsModelStruct) SetFlagRefreshOngoingOfDetailedTestCaseExecution(detailedTestCaseExecutionMapKey DetailedTestCaseExecutionMapKeyType)`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: SetFlagRefreshOngoingOfDetailedTestCaseExecution Set the flag there is an ongoing refresh of the DetailedTestCaseExecution-data
- Selector calls: `detailedTestCaseExecutionsMapMutex.Lock`, `detailedTestCaseExecutionsMapMutex.Unlock`, `fmt.Println`

### ClearFlagRefreshOngoingOfDetailedTestCaseExecution (method on `TestCaseExecutionsModelStruct`)
- Signature: `func (TestCaseExecutionsModelStruct) ClearFlagRefreshOngoingOfDetailedTestCaseExecution(detailedTestCaseExecutionMapKey DetailedTestCaseExecutionMapKeyType)`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: ClearFlagRefreshOngoingOfDetailedTestCaseExecution Clear the flag there is an ongoing refresh of the DetailedTestCaseExecution-data
- Selector calls: `detailedTestCaseExecutionsMapMutex.Lock`, `detailedTestCaseExecutionsMapMutex.Unlock`, `fmt.Println`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
