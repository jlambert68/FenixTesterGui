# commandEngine_loadAvailbleBondsIntoModel.go

## File Overview
- Path: `commandAndRuleEngine/commandEngine_loadAvailbleBondsIntoModel.go`
- Package: `commandAndRuleEngine`
- Functions/Methods: `2`
- Imports: `1`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `LoadAvailableBondsFromServer`

## Imports
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### LoadAvailableBondsFromServer (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) LoadAvailableBondsFromServer()`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: LoadAvailableBondsFromServer Load all Available Bonds from Gui-server
- Selector calls: `commandAndRuleEngine.loadModelWithAvailableBonds`

### loadModelWithAvailableBonds (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) loadModelWithAvailableBonds(availableImmatureBondsMessage *fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage)`
- Exported: `false`
- Control-flow features: `for/range`
- Doc: Load Model with Available Bonds

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
