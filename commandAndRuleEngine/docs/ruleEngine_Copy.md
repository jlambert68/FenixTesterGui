# ruleEngine_Copy.go

## File Overview
- Path: `commandAndRuleEngine/ruleEngine_Copy.go`
- Package: `commandAndRuleEngine`
- Functions/Methods: `2`
- Imports: `2`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `errors`
- `fmt`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### verifyIfElementCanBeCopied (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) verifyIfElementCanBeCopied(testCaseUuid string, elementUuid string) (canBeCopied bool, matchedSimpldRule string, err error)`
- Exported: `false`
- Control-flow features: `returns error`
- Doc: Verify if an element can be copied or not, regarding copy rules
- Selector calls: `commandAndRuleEngine.verifyIfComponentCanBeCopiedSimpleRules`

### executeCopyElement (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeCopyElement(testCaseUuid string, elementUuid string) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Copy an element, but first ensure that rules for copying are used
- Selector calls: `commandAndRuleEngine.verifyIfElementCanBeCopied`, `errors.New`, `fmt.Sprintf`, `commandAndRuleEngine.executeCopyFullELementStructure`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
