# ruleEngine_verifyCopy.go

## File Overview
- Path: `commandAndRuleEngine/ruleEngine_verifyCopy.go`
- Package: `commandAndRuleEngine`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `1`
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
### verifyIfComponentCanBeCopiedSimpleRules (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) verifyIfComponentCanBeCopiedSimpleRules(testCaseUuid string, elementUuid string) (canBeCopied bool, matchedRule string, err error)`
- Exported: `false`
- Control-flow features: `if, switch, returns error`
- Doc: Verify the simple rules if a component can be copied or not
- External calls: `componentType.String`, `errors.New`, `fmt.Sprintf`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
