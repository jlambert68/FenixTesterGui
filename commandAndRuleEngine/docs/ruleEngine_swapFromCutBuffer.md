# ruleEngine_swapFromCutBuffer.go

## File Overview
- Path: `commandAndRuleEngine/ruleEngine_swapFromCutBuffer.go`
- Package: `commandAndRuleEngine`
- Functions/Methods: `2`
- Imports: `3`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `FenixTesterGui/testCase/testCaseModel`
- `errors`
- `fmt`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### verifyIfElementCanBeSwappedForCutBuffer (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) verifyIfElementCanBeSwappedForCutBuffer(testCaseUuid string, elementUuidToBeCutOut string) (canBeSwapped bool, matchedSimpledRule string, matchedComplexRule string, err error)`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Verify if anor element can be swapped or not, regarding swap rules
- Selector calls: `errors.New`, `fmt.Sprintf`, `commandAndRuleEngine.verifyIfElementCanBeSwapped`

### executeSwapElementFromCutBuffer (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeSwapElementFromCutBuffer(testCaseUuid string, elementToSwapOutUuid string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) err error`
- Exported: `false`
- Control-flow features: `returns error`
- Doc: Swap an element, but first ensure that rules for swapping are used
- Selector calls: `commandAndRuleEngine.executeSwapElement`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
