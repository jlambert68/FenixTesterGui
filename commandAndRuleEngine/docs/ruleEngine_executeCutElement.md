# ruleEngine_executeCutElement.go

## File Overview
- Path: `commandAndRuleEngine/ruleEngine_executeCutElement.go`
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
### executeCutFullELementStructure (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeCutFullELementStructure(testCaseUuid string, uuidToBeCutOut string) err error`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Selector calls: `commandAndRuleEngine.executeDeleteElement`, `commandAndRuleEngine.recursiveCuttingOfFullElementStructure`, `errors.New`, `fmt.Sprintf`

### recursiveCuttingOfFullElementStructure (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) recursiveCuttingOfFullElementStructure(currentTestCase *testCaseModel.TestCaseModelStruct, elementsUuid string, copiedElementStructure *testCaseModel.MatureElementStruct) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Copy the full structure of all children, in TestCase-model, for specific Element
- Selector calls: `commandAndRuleEngine.recursiveCuttingOfFullElementStructure`, `errors.New`, `fmt.Sprintf`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
