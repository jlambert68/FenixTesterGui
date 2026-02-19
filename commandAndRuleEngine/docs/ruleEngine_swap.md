# ruleEngine_swap.go

## File Overview
- Path: `commandAndRuleEngine/ruleEngine_swap.go`
- Package: `commandAndRuleEngine`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `5`
- Imports: `10`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/testCase/testCaseModel`
- `errors`
- `fmt`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions`
- `github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTemplateToThisDomain`
- `github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTemplateToThisDomain/version_1_0`
- `github.com/sirupsen/logrus`
- `google.golang.org/protobuf/types/known/timestamppb`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### addTestInstructionContainerDataToTestCaseModel (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) addTestInstructionContainerDataToTestCaseModel(testCaseUuid string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct, matureElementToSwapIn *testCaseModel.MatureElementStruct) err error`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: Add All TestInstructionContainer-data for the new TestInstructionContainer into the TestCase-model
- External calls: `err.Error`, `errors.New`, `fmt.Println`, `fmt.Sprintf`, `timestamppb.Now`

### addTestInstructionDataToTestCaseModel (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) addTestInstructionDataToTestCaseModel(testCaseUuid string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct, matureElementToSwapIn *testCaseModel.MatureElementStruct) err error`
- Exported: `false`
- Control-flow features: `if, for/range, switch, returns error`
- Doc: Add All TestInstruction-data for the new TestInstruction into the TestCase-model
- Internal calls: `string`
- External calls: `err.Error`, `errors.New`, `fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionAttributeTypeEnum`, `fmt.Println`, `fmt.Sprintf`, `tempExecutionDomain.GetNameUsedInGui`, `timestamppb.Now`

### executeSwapElement (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeSwapElement(testCaseUuid string, elementToSwapOutUuid string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Swap an element, but first ensure that rules for swapping are used
- External calls: `commandAndRuleEngine.executeSwapElementBasedOnRule`, `commandAndRuleEngine.verifyIfElementCanBeSwapped`, `errors.New`, `fmt.Sprintf`

### executeSwapElementBasedOnRule (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeSwapElementBasedOnRule(testCaseUuid string, elementToBeSwappedIOutUuid string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct, matchedComplexRule string) err error`
- Exported: `false`
- Control-flow features: `if, switch, returns error`
- Doc: Execute a swap on an element based on specific rule
- External calls: `commandAndRuleEngine.addTestInstructionContainerDataToTestCaseModel`, `commandAndRuleEngine.addTestInstructionDataToTestCaseModel`, `commandAndRuleEngine.executeTCRuleSwap101`, `commandAndRuleEngine.executeTCRuleSwap102`, `commandAndRuleEngine.executeTCRuleSwap103`, `commandAndRuleEngine.executeTCRuleSwap104`, `commandAndRuleEngine.executeTCRuleSwap105`, `commandAndRuleEngine.executeTCRuleSwap106`

### verifyIfElementCanBeSwapped (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) verifyIfElementCanBeSwapped(testCaseUuid string, elementUuidToBeSwappedOut string, elementTypeToBeSwappedIn fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum) (canBeSwapped bool, matchedSimpledRule string, matchedComplexRule string, err error)`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Verify if anor element can be swapped or not, regarding swap rules
- External calls: `commandAndRuleEngine.verifyIfComponentCanBeSwappedSimpleRules`, `commandAndRuleEngine.verifyIfComponentCanBeSwappedWithComplexRules`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
