# ruleEngine_verifyDelete.go

## File Overview
- Path: `commandAndRuleEngine/ruleEngine_verifyDelete.go`
- Package: `commandAndRuleEngine`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `2`
- Imports: `5`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `FenixTesterGui/testCase/testCaseModel`
- `errors`
- `fmt`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### verifyIfComponentCanBeDeletedSimpleRules (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) verifyIfComponentCanBeDeletedSimpleRules(testCaseUuid string, elementUuid string) (canBeDeleted bool, matchedRule string, err error)`
- Exported: `false`
- Control-flow features: `if, switch, returns error`
- Doc: Verify the simple rules if a component can be deleted or not
- External calls: `componentType.String`, `errors.New`, `fmt.Sprintf`

### verifyIfComponentCanBeDeletedWithComplexRules (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) verifyIfComponentCanBeDeletedWithComplexRules(testCaseUuid string, uuidToDelete string) (matchedRule string, err error)`
- Exported: `false`
- Control-flow features: `if, switch, returns error`
- Doc: Verify the complex rules if a component can be deleted or not Rules how deletion of an element is done
- External calls: `errors.New`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
