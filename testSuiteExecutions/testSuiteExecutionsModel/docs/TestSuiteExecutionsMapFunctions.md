# TestSuiteExecutionsMapFunctions.go

## File Overview
- Path: `testSuiteExecutions/testSuiteExecutionsModel/TestSuiteExecutionsMapFunctions.go`
- Package: `testSuiteExecutionsModel`
- Functions/Methods: `8`
- Imports: `4`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `AddToTestSuiteExecutionsMap`
- `DeleteFromTestSuiteExecutionsMap`
- `GetNumberOfTestSuiteExecutionsRetrievedFromDatabase`
- `InitiateTestSuiteExecutionsMap`
- `ReadAllFromTestSuiteExecutionsMap`
- `ReadAllFromTestSuiteExecutionsMapForTableList`
- `ReadFromTestSuiteExecutionsMap`
- `ReadFromTestSuiteExecutionsMapForTableList`

## Imports
- `FenixTesterGui/common_code`
- `github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `sync`

## Declared Types
- `readListOrDetailedMessagesTypeType`

## Declared Constants
- `ReadListMessages`
- `readDetailedMessages`

## Declared Variables
- `testSuiteExecutionsMapMutex`

## Functions and Methods
### InitiateTestSuiteExecutionsMap (method on `TestSuiteExecutionsModelStruct`)
- Signature: `func (TestSuiteExecutionsModelStruct) InitiateTestSuiteExecutionsMap()`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: InitiateTestSuiteExecutionsMap Add to the TestSuiteExecutions-Map
- Selector calls: `testSuiteExecutionsMapMutex.Lock`, `testSuiteExecutionsMapMutex.Unlock`

### ReadFromTestSuiteExecutionsMapForTableList (method on `TestSuiteExecutionsModelStruct`)
- Signature: `func (TestSuiteExecutionsModelStruct) ReadFromTestSuiteExecutionsMapForTableList(testSuiteExecutionsMapKey TestSuiteExecutionUuidType) (testSuiteExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage, existInMap bool)`
- Exported: `true`
- Control-flow features: `defer`
- Doc: ReadFromTestSuiteExecutionsMapForTableList Read from the TestSuiteExecutions-Map for the TableList for TestSuiteExecutions
- Selector calls: `testSuiteExecutionsModel.ReadFromTestSuiteExecutionsMap`

### ReadFromTestSuiteExecutionsMap (method on `TestSuiteExecutionsModelStruct`)
- Signature: `func (TestSuiteExecutionsModelStruct) ReadFromTestSuiteExecutionsMap(readListOrDetailedMessagesType readListOrDetailedMessagesTypeType, testSuiteExecutionsMapKey TestSuiteExecutionUuidType) (testSuiteExecutionsListMessages *fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage, testSuiteExecutionResponseMessage *fenixExecutionServerGuiGrpcApi.TestSuiteExecutionResponseMessage, existInMap bool)`
- Exported: `true`
- Control-flow features: `if, switch, defer`
- Doc: ReadFromTestSuiteExecutionsMap Read from the TestSuiteExecutions-Map
- Internal calls: `DetailedTestSuiteExecutionMapKeyType`
- Selector calls: `testSuiteExecutionsMapMutex.RLock`, `testSuiteExecutionsMapMutex.RUnlock`

### ReadAllFromTestSuiteExecutionsMapForTableList (method on `TestSuiteExecutionsModelStruct`)
- Signature: `func (TestSuiteExecutionsModelStruct) ReadAllFromTestSuiteExecutionsMapForTableList() testSuiteExecutionsListMessages *[]*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage`
- Exported: `true`
- Control-flow features: `defer`
- Doc: ReadAllFromTestSuiteExecutionsMapForTableList Read all from the TestSuiteExecutions-Map for the TableList for TestSuiteExecutions
- Selector calls: `testSuiteExecutionsModel.ReadAllFromTestSuiteExecutionsMap`

### ReadAllFromTestSuiteExecutionsMap (method on `TestSuiteExecutionsModelStruct`)
- Signature: `func (TestSuiteExecutionsModelStruct) ReadAllFromTestSuiteExecutionsMap(readListOrDetailedMessagesType readListOrDetailedMessagesTypeType) (testSuiteExecutionsListMessages *[]*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage, testSuiteExecutionResponseMessages *[]*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionResponseMessage)`
- Exported: `true`
- Control-flow features: `if, for/range, switch, defer`
- Doc: ReadAllFromTestSuiteExecutionsMap Read all from the TestSuiteExecutions-Map
- Selector calls: `testSuiteExecutionsMapMutex.RLock`, `testSuiteExecutionsMapMutex.RUnlock`

### GetNumberOfTestSuiteExecutionsRetrievedFromDatabase (method on `TestSuiteExecutionsModelStruct`)
- Signature: `func (TestSuiteExecutionsModelStruct) GetNumberOfTestSuiteExecutionsRetrievedFromDatabase() numberOfTestSuiteExecutionsRetrievedFromDatabase int`
- Exported: `true`
- Control-flow features: `defer`
- Doc: GetNumberOfTestSuiteExecutionsRetrievedFromDatabase Read all from the TestSuiteExecutions-Map
- Selector calls: `testSuiteExecutionsMapMutex.RLock`, `testSuiteExecutionsMapMutex.RUnlock`, `testSuiteExecutionsModel.ReadAllFromTestSuiteExecutionsMapForTableList`

### AddToTestSuiteExecutionsMap (method on `TestSuiteExecutionsModelStruct`)
- Signature: `func (TestSuiteExecutionsModelStruct) AddToTestSuiteExecutionsMap(testSuiteExecutionsMapKey TestSuiteExecutionUuidType, testSuiteExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage)`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: AddToTestSuiteExecutionsMap Add to the TestSuiteExecutions-Map
- Selector calls: `testSuiteExecutionsMapMutex.Lock`, `testSuiteExecutionsMapMutex.Unlock`

### DeleteFromTestSuiteExecutionsMap (method on `TestSuiteExecutionsModelStruct`)
- Signature: `func (TestSuiteExecutionsModelStruct) DeleteFromTestSuiteExecutionsMap(testSuiteExecutionsMapKey TestSuiteExecutionUuidType)`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: DeleteFromTestSuiteExecutionsMap Delete from the TestSuiteExecutions-Map
- Selector calls: `testSuiteExecutionsMapMutex.Lock`, `testSuiteExecutionsMapMutex.Unlock`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
