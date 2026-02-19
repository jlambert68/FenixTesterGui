# ruleEngine_executeCopyElement.go

## File Overview
- Path: `commandAndRuleEngine/ruleEngine_executeCopyElement.go`
- Package: `commandAndRuleEngine`
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
### executeCopyFullELementStructure (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeCopyFullELementStructure(testCaseUuid string, uuidToCopy string) err error`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Selector calls: `errors.New`, `fmt.Sprintf`, `commandAndRuleEngine.recursiveCopyingOfFullElementStructure`

### recursiveCopyingOfFullElementStructure (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) recursiveCopyingOfFullElementStructure(currentTestCase *testCaseModel.TestCaseModelStruct, elementsUuid string, copiedElementStructure *testCaseModel.ImmatureElementStruct) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Copy the full structure of all children, in TestCase-model, for specific Element
- Selector calls: `errors.New`, `fmt.Sprintf`, `commandAndRuleEngine.recursiveCopyingOfFullElementStructure`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
