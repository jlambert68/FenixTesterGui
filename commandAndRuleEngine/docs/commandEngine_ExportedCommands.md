# commandEngine_ExportedCommands.go

## File Overview
- Path: `commandAndRuleEngine/commandEngine_ExportedCommands.go`
- Package: `commandAndRuleEngine`
- Functions/Methods: `5`
- Imports: `2`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `DeleteElementFromTestCaseModel`
- `InitiateCommandChannelReader`
- `NewTestCaseModel`
- `SwapElementsInTestCaseModel`
- `VerifyIfElementCanBeSwapped`

## Imports
- `FenixTesterGui/testCase/testCaseModel`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### DeleteElementFromTestCaseModel (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) DeleteElementFromTestCaseModel(testCaseUuid string, elementId string) err error`
- Exported: `true`
- Control-flow features: `returns error`
- Doc: DeleteElementFromTestCaseModel Used, mostly from GUI, for Deleting an element from a TestCaseModel that is used within a TestCase
- Selector calls: `commandAndRuleEngine.executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel`

### InitiateCommandChannelReader (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) InitiateCommandChannelReader()`
- Exported: `true`
- Control-flow features: `go`
- Doc: InitiateCommandChannelReader Initiate the channel reader which is used for sending commands to CommandEngine
- Selector calls: `commandAndRuleEngine.startCommandChannelReader`

### NewTestCaseModel (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) NewTestCaseModel() (testCaseUuid string, err error)`
- Exported: `true`
- Control-flow features: `returns error`
- Doc: NewTestCaseModel Used, mostly from GUI, to for creating a new TestCase-Model to be used within a new TestCase
- Selector calls: `commandAndRuleEngine.executeCommandOnTestCaseModel_NewTestCaseModel`

### SwapElementsInTestCaseModel (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) SwapElementsInTestCaseModel(testcaseUuid string, elementUuidTobeSwappedOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) err error`
- Exported: `true`
- Control-flow features: `returns error`
- Doc: SwapElementsInTestCaseModel Used, mostly from GUI, for Swapping out an element, and in an element structure, from a TestCaseModel that is used within a TestCase
- Selector calls: `commandAndRuleEngine.executeCommandOnTestCaseModel_SwapOutElemenAndInNewElementInTestCaseModel`

### VerifyIfElementCanBeSwapped (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) VerifyIfElementCanBeSwapped(testCaseUuid string, elementUuidToBeSwappedOut string, elementTypeToBeSwappedIn fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum) (canBeSwappedIn bool, err error)`
- Exported: `true`
- Control-flow features: `returns error`
- Doc: VerifyIfElementCanBeSwapped Verify if an element can be swapped or not, regarding swap rules
- Selector calls: `commandAndRuleEngine.verifyIfElementCanBeSwapped`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
