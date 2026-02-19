# ruleEngine_swapFromCopyBuffer.go

## File Overview
- Path: `commandAndRuleEngine/ruleEngine_swapFromCopyBuffer.go`
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
### verifyIfElementCanBeSwappedForCopyBuffer (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) verifyIfElementCanBeSwappedForCopyBuffer(testCaseUuid string, elementUuid string) (canBeSwapped bool, matchedSimpledRule string, matchedComplexRule string, err error)`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Verify if an element can be swapped for copy Buffer or not, regarding swap rules
- Selector calls: `errors.New`, `fmt.Sprintf`, `commandAndRuleEngine.verifyIfComponentCanBeSwappedSimpleRules`, `commandAndRuleEngine.verifyIfComponentCanBeSwappedWithComplexRules`

### executeSwapElementForCopyBuffer (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeSwapElementForCopyBuffer(testCaseUuid string, elementToSwapOutUuid string) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Swap an element for content in Copy Buffer, but first ensure that rules for swapping are used
- Selector calls: `errors.New`, `fmt.Sprintf`, `commandAndRuleEngine.verifyIfElementCanBeSwapped`, `commandAndRuleEngine.executeSwapElementBasedOnRule`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
