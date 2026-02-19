# TestCaseExecutionsMapFunctions.go

## File Overview
- Path: `testCaseExecutions/testCaseExecutionsModel/TestCaseExecutionsMapFunctions.go`
- Package: `testCaseExecutionsModel`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `8`
- Imports: `4`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `AddToTestCaseExecutionsMap`
- `DeleteFromTestCaseExecutionsMap`
- `GetNumberOfTestCaseExecutionsRetrievedFromDatabase`
- `GetTestInstructionExecutionUuidFromTestInstructionUuid`
- `GetTestInstructionFromTestInstructionExecutionUuid`
- `InitiateTestCaseExecutionsMap`
- `ReadAllFromTestCaseExecutionsMap`
- `ReadFromTestCaseExecutionsMap`

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
- `testCaseExecutionsMapMutex`

## Functions and Methods
### AddToTestCaseExecutionsMap (method on `TestCaseExecutionsModelStruct`)
- Signature: `func (TestCaseExecutionsModelStruct) AddToTestCaseExecutionsMap(testCaseExecutionsMapKey TestCaseExecutionUuidType, testCaseExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage)`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: AddToTestCaseExecutionsMap Add to the TestCaseExecutions-Map
- External calls: `testCaseExecutionsMapMutex.Lock`, `testCaseExecutionsMapMutex.Unlock`

### DeleteFromTestCaseExecutionsMap (method on `TestCaseExecutionsModelStruct`)
- Signature: `func (TestCaseExecutionsModelStruct) DeleteFromTestCaseExecutionsMap(testCaseExecutionsMapKey TestCaseExecutionUuidType)`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: DeleteFromTestCaseExecutionsMap Delete from the TestCaseExecutions-Map
- External calls: `testCaseExecutionsMapMutex.Lock`, `testCaseExecutionsMapMutex.Unlock`

### GetNumberOfTestCaseExecutionsRetrievedFromDatabase (method on `TestCaseExecutionsModelStruct`)
- Signature: `func (TestCaseExecutionsModelStruct) GetNumberOfTestCaseExecutionsRetrievedFromDatabase() numberOfTestCaseExecutionsRetrievedFromDatabase int`
- Exported: `true`
- Control-flow features: `defer`
- Doc: GetNumberOfTestCaseExecutionsRetrievedFromDatabase Read all from the TestCaseExecutions-Map
- External calls: `testCaseExecutionsMapMutex.RLock`, `testCaseExecutionsMapMutex.RUnlock`, `testCaseExecutionsModel.ReadAllFromTestCaseExecutionsMap`

### GetTestInstructionExecutionUuidFromTestInstructionUuid (method on `TestCaseExecutionsModelStruct`)
- Signature: `func (TestCaseExecutionsModelStruct) GetTestInstructionExecutionUuidFromTestInstructionUuid(testCaseExecutionsMapKey TestCaseExecutionUuidType, testInstructionUuid RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType) (testInstructionExecutionUuid TestInstructionExecutionUuidType, existInMap bool)`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: GetTestInstructionExecutionUuidFromTestInstructionUuid Read from the TestCaseExecutions-Map and get the TestInstructionExecutionUuid + Version (mpKey) based on TestInstructionUuid
- Internal calls: `DetailedTestCaseExecutionMapKeyType`

### GetTestInstructionFromTestInstructionExecutionUuid (method on `TestCaseExecutionsModelStruct`)
- Signature: `func (TestCaseExecutionsModelStruct) GetTestInstructionFromTestInstructionExecutionUuid(testCaseExecutionsMapKey TestCaseExecutionUuidType, testInstructionExecutionUuid TestInstructionExecutionUuidType, lockMap bool) (testInstructionUuid RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType, testInstructionName string, existInMap bool)`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: GetTestInstructionFromTestInstructionExecutionUuid Read from the TestCaseExecutions-Map and get the TestInstruction Uuid and Name based on TestInstructionExecutionUuid + Version (mapKey)
- Internal calls: `DetailedTestCaseExecutionMapKeyType`
- External calls: `testCaseExecutionsMapMutex.RLock`, `testCaseExecutionsMapMutex.RUnlock`

### InitiateTestCaseExecutionsMap (method on `TestCaseExecutionsModelStruct`)
- Signature: `func (TestCaseExecutionsModelStruct) InitiateTestCaseExecutionsMap()`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: InitiateTestCaseExecutionsMap Add to the TestCaseExecutions-Map
- External calls: `testCaseExecutionsMapMutex.Lock`, `testCaseExecutionsMapMutex.Unlock`

### ReadAllFromTestCaseExecutionsMap (method on `TestCaseExecutionsModelStruct`)
- Signature: `func (TestCaseExecutionsModelStruct) ReadAllFromTestCaseExecutionsMap() testCaseExecutionsListMessage *[]*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage`
- Exported: `true`
- Control-flow features: `if, for/range, defer`
- Doc: ReadAllFromTestCaseExecutionsMap Read all from the TestCaseExecutions-Map
- External calls: `testCaseExecutionsMapMutex.RLock`, `testCaseExecutionsMapMutex.RUnlock`

### ReadFromTestCaseExecutionsMap (method on `TestCaseExecutionsModelStruct`)
- Signature: `func (TestCaseExecutionsModelStruct) ReadFromTestCaseExecutionsMap(testCaseExecutionsMapKey TestCaseExecutionUuidType) (testCaseExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage, existInMap bool)`
- Exported: `true`
- Control-flow features: `if, defer`
- Doc: ReadFromTestCaseExecutionsMap Read from the TestCaseExecutions-Map
- External calls: `testCaseExecutionsMapMutex.RLock`, `testCaseExecutionsMapMutex.RUnlock`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
