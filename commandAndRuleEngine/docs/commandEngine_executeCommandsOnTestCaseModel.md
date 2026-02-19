# commandEngine_executeCommandsOnTestCaseModel.go

## File Overview
- Path: `commandAndRuleEngine/commandEngine_executeCommandsOnTestCaseModel.go`
- Package: `commandAndRuleEngine`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `7`
- Imports: `7`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/testCase/testCaseModel`
- `errors`
- `fmt`
- `github.com/google/uuid`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `google.golang.org/protobuf/types/known/timestamppb`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### executeCommandOnTestCaseModel_CopyElementInTestCaseModel (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_CopyElementInTestCaseModel(testCaseUuid string, elementIdToCopy string) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: TestCaseCommandTypeEnum_COPY_ELEMENT Used for copying an element  in a TestCaseModel that is used within a TestCase
- Internal calls: `int32`
- External calls: `commandAndRuleEngine.executeCopyElement`, `errors.New`, `fmt.Sprintf`, `timestamppb.Now`

### executeCommandOnTestCaseModel_CutElementInTestCaseModel (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_CutElementInTestCaseModel(testCaseUuid string, elementIdToCut string) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: TestCaseCommandTypeEnum_CUT_ELEMENT Used for cutting an element in a TestCaseModel that is used within a TestCase
- Internal calls: `int32`
- External calls: `commandAndRuleEngine.executeCutElement`, `errors.New`, `fmt.Sprintf`, `timestamppb.Now`

### executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel(testCaseUuid string, elementId string) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: TestCaseCommandTypeEnum_REMOVE_ELEMENT Used for Deleting an element from a TestCaseModel that is used within a TestCase
- Internal calls: `int32`
- External calls: `commandAndRuleEngine.executeDeleteElement`, `errors.New`, `fmt.Sprintf`, `timestamppb.Now`

### executeCommandOnTestCaseModel_NewTestCaseModel (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_NewTestCaseModel() (testCaseUuid string, err error)`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: TestCaseCommandTypeEnum_NEW_TESTCASE Used for creating a new TestCase-Model to be used within a new TestCase
- Internal calls: `int32`
- External calls: `commandAndRuleEngine.createNewBondB0Element`, `timestamppb.Now`, `uuidGenerator.New`

### executeCommandOnTestCaseModel_SwapInElementFromCopyBufferInTestCaseModel (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_SwapInElementFromCopyBufferInTestCaseModel(testCaseUuid string, elementIdToBeReplacedByCopyBuffer string) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_COPY_BUFFER_ELEMENT Used for Swapping in an element from Copy Buffer in a TestCaseModel that is used within a TestCase
- Internal calls: `int32`
- External calls: `commandAndRuleEngine.executeSwapElementForCopyBuffer`, `errors.New`, `fmt.Sprintf`, `timestamppb.Now`

### executeCommandOnTestCaseModel_SwapInElementFromCutBufferInTestCaseModel (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_SwapInElementFromCutBufferInTestCaseModel(testCaseUuid string, uuidToReplacedByCutBufferContent string) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_CUT_BUFFER_ELEMENT Used for Swapping in an element from Cut opy Buffer in a TestCaseModel that is used within a TestCase
- Internal calls: `int32`
- External calls: `commandAndRuleEngine.executeSwapElementFromCutBuffer`, `errors.New`, `fmt.Sprintf`, `timestamppb.Now`

### executeCommandOnTestCaseModel_SwapOutElemenAndInNewElementInTestCaseModel (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_SwapOutElemenAndInNewElementInTestCaseModel(testCaseUuid string, elementToSwapOutUuid string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_NEW_ELEMENT Used for Swapping out an element, and in an element structure, from a TestCaseModel that is used within a TestCase
- Internal calls: `int32`
- External calls: `commandAndRuleEngine.executeSwapElement`, `errors.New`, `fmt.Sprintf`, `timestamppb.Now`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
