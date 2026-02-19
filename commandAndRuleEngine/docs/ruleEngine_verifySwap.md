# ruleEngine_verifySwap.go

## File Overview
- Path: `commandAndRuleEngine/ruleEngine_verifySwap.go`
- Package: `commandAndRuleEngine`
- Functions/Methods: `7`
- Imports: `6`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `FindElementInSliceAndRemove`

## Imports
- `FenixTesterGui/testCase/testCaseModel`
- `errors`
- `fmt`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `reflect`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### verifyIfComponentCanBeSwappedSimpleRules (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) verifyIfComponentCanBeSwappedSimpleRules(testCaseUuid string, elementUuid string) (canBeSwapped bool, matchedRule string, err error)`
- Exported: `false`
- Control-flow features: `if, switch, returns error`
- Doc: Verify the simple rules if a component can be Swapped or not
- Selector calls: `errors.New`, `componentType.String`

### verifyIfComponentCanBeSwappedWithComplexRules (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) verifyIfComponentCanBeSwappedWithComplexRules(testCaseUuid string, uuidToSwapOut string, elementTypeToBeSwappedIn fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum) (matchedRule string, err error)`
- Exported: `false`
- Control-flow features: `if, switch, returns error`
- Doc: Verify the Complex rules if a component can be Swapped or not
- Selector calls: `errors.New`

### verifyThatThereAreNoZombieElementsInComponent (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) verifyThatThereAreNoZombieElementsInComponent(immatureElement testCaseModel.ImmatureElementStruct) err error`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: Verify that all UUIDs are correct in component to be swapped in. Means that no empty uuid is allowed and they all are correct
- Selector calls: `commandAndRuleEngine.recursiveZombieElementSearchInComponentModel`, `errors.New`

### recursiveZombieElementSearchInComponentModel (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) recursiveZombieElementSearchInComponentModel(elementsUuid string, allUuidKeys []string, immatureElement *testCaseModel.ImmatureElementStruct) (processedAllUuidKeys []string, err error)`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Verify all children, in ImmatureEleemnt-model and remove the found element from 'allUuidKeys'
- Internal calls: `FindElementInSliceAndRemove`
- Selector calls: `errors.New`, `commandAndRuleEngine.recursiveZombieElementSearchInComponentModel`

### FindElementInSliceAndRemove
- Signature: `func FindElementInSliceAndRemove(sliceToWorkOn *[]string, uuid string) returnSlice *[]string`
- Exported: `true`
- Control-flow features: `if, for/range`
- Doc: Remove 'uuid' from slice

### verifyThatAllUuidsAreCorrectInComponent (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) verifyThatAllUuidsAreCorrectInComponent(immatureElement testCaseModel.ImmatureElementStruct) err error`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: Verify that all UUIDs are correct in component to be swapped in. Means that no empty uuid is allowed and they all are correct
- Selector calls: `reflect.ValueOf`, `e.NumField`, `e.Type`, `e.Field`, `varType.Kind`, `fmt.Printf`

### recursiveVerifyAllUuidOfChildElements (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) recursiveVerifyAllUuidOfChildElements(testCaseUuid string, elementsUuid string) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Verify all children, in new Element-model to be swapped in, that they contain correct UUIDs
- Selector calls: `errors.New`, `commandAndRuleEngine.recursiveDeleteOfChildElements`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
