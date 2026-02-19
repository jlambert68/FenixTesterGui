# commandEngine_executeCommandsOnTestCaseModel_test.go

## File Overview
- Path: `commandAndRuleEngine/commandEngine_executeCommandsOnTestCaseModel_test.go`
- Package: `commandAndRuleEngine`
- Functions/Methods: `6`
- Imports: `8`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `TestCopyElementCommandOnTestCaseModel`
- `TestNewTestCaseModelCommand`
- `TestRemoveElementCommandOnTestCaseModel`
- `TestSwapElementCommandOnTestCaseModel`
- `TestSwapElementFromCopyBufferCommandOnTestCaseModel`
- `TestSwapElementFromCutBufferCommandOnTestCaseModel`

## Imports
- `FenixTesterGui/gui/UnitTestTestData`
- `FenixTesterGui/testCase/testCaseModel`
- `fmt`
- `github.com/google/uuid`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `github.com/stretchr/testify/assert`
- `strconv`
- `testing`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### TestCopyElementCommandOnTestCaseModel
- Signature: `func TestCopyElementCommandOnTestCaseModel(t *testing.T)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: Test to Copy an element from the TestCaseModel
- Internal calls: `int32`
- Selector calls: `UnitTestTestData.InitLoggerForTest`, `assert.Equal`, `commandAndRuleEngine.executeCommandOnTestCaseModel_CopyElementInTestCaseModel`, `fmt.Sprint`, `uuidGenerator.New`

### TestNewTestCaseModelCommand
- Signature: `func TestNewTestCaseModelCommand(t *testing.T)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: Test to create a New TestCaseModel
- Internal calls: `int`
- Selector calls: `UnitTestTestData.InitLoggerForTest`, `assert.Equal`, `commandAndRuleEngine.executeCommandOnTestCaseModel_NewTestCaseModel`, `fmt.Sprint`, `strconv.Itoa`

### TestRemoveElementCommandOnTestCaseModel
- Signature: `func TestRemoveElementCommandOnTestCaseModel(t *testing.T)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: Test to Delete an element from the TestCaseModel
- Internal calls: `int32`
- Selector calls: `UnitTestTestData.InitLoggerForTest`, `assert.Equal`, `commandAndRuleEngine.executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel`, `fmt.Sprint`, `uuidGenerator.New`

### TestSwapElementCommandOnTestCaseModel
- Signature: `func TestSwapElementCommandOnTestCaseModel(t *testing.T)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: Test to Swap out an element and in another element in the TestCaseModel
- Internal calls: `int32`
- Selector calls: `UnitTestTestData.InitLoggerForTest`, `assert.Equal`, `commandAndRuleEngine.executeCommandOnTestCaseModel_SwapOutElemenAndInNewElementInTestCaseModel`, `fmt.Sprint`, `uuidGenerator.New`

### TestSwapElementFromCopyBufferCommandOnTestCaseModel
- Signature: `func TestSwapElementFromCopyBufferCommandOnTestCaseModel(t *testing.T)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: Test to Swap in element from Copy Buffer on the TestCaseModel
- Internal calls: `int32`
- Selector calls: `UnitTestTestData.InitLoggerForTest`, `assert.Equal`, `commandAndRuleEngine.executeCommandOnTestCaseModel_CopyElementInTestCaseModel`, `commandAndRuleEngine.executeCommandOnTestCaseModel_SwapInElementFromCopyBufferInTestCaseModel`, `fmt.Sprint`, `uuidGenerator.New`

### TestSwapElementFromCutBufferCommandOnTestCaseModel
- Signature: `func TestSwapElementFromCutBufferCommandOnTestCaseModel(t *testing.T)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: Test to Swap in element from Copy Buffer on the TestCaseModel
- Internal calls: `int32`
- Selector calls: `UnitTestTestData.InitLoggerForTest`, `assert.Equal`, `commandAndRuleEngine.executeCommandOnTestCaseModel_CutElementInTestCaseModel`, `commandAndRuleEngine.executeCommandOnTestCaseModel_SwapInElementFromCutBufferInTestCaseModel`, `fmt.Sprint`, `uuidGenerator.New`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
