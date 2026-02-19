# ruleEngine_executeSwapElement.go

## File Overview
- Path: `commandAndRuleEngine/ruleEngine_executeSwapElement.go`
- Package: `commandAndRuleEngine`
- Functions/Methods: `10`
- Imports: `6`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `FenixTesterGui/testCase/testCaseModel`
- `errors`
- `fmt`
- `github.com/google/uuid`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### verifySwapRuleAndConvertIntoMatureComponentElementModel (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) verifySwapRuleAndConvertIntoMatureComponentElementModel(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct, ruleNameToVerify string) (matureElementToSwapIn testCaseModel.MatureElementStruct, err error)`
- Exported: `false`
- Control-flow features: `if, returns error`
- Selector calls: `errors.New`, `fmt.Sprintf`, `commandAndRuleEngine.verifyIfElementCanBeSwapped`, `commandAndRuleEngine.transformImmatureElementModelIntoMatureElementModel`

### executeTCRuleSwap101 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleSwap101(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (matureElementToSwapIn testCaseModel.MatureElementStruct, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: TCRuleSwap101 What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
- Selector calls: `commandAndRuleEngine.verifySwapRuleAndConvertIntoMatureComponentElementModel`, `errors.New`, `commandAndRuleEngine.createNewBondB1fElement`, `commandAndRuleEngine.createNewBondB1lElement`

### executeTCRuleSwap102 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleSwap102(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (matureElementToSwapIn testCaseModel.MatureElementStruct, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: TCRuleSwap102 What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
- Selector calls: `commandAndRuleEngine.verifySwapRuleAndConvertIntoMatureComponentElementModel`, `errors.New`, `commandAndRuleEngine.createNewBondB11fElement`, `commandAndRuleEngine.createNewBondB11lElement`

### executeTCRuleSwap103 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleSwap103(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (matureElementToSwapIn testCaseModel.MatureElementStruct, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: TCRuleSwap103 What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
- Selector calls: `commandAndRuleEngine.verifySwapRuleAndConvertIntoMatureComponentElementModel`, `errors.New`, `commandAndRuleEngine.createNewBondB12Element`

### executeTCRuleSwap104 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleSwap104(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (matureElementToSwapIn testCaseModel.MatureElementStruct, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: TCRuleSwap104 What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
- Selector calls: `commandAndRuleEngine.verifySwapRuleAndConvertIntoMatureComponentElementModel`, `errors.New`, `commandAndRuleEngine.createNewBondB12Element`

### executeTCRuleSwap105 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleSwap105(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (matureElementToSwapIn testCaseModel.MatureElementStruct, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: TCRuleSwap105 What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
- Selector calls: `commandAndRuleEngine.verifySwapRuleAndConvertIntoMatureComponentElementModel`, `errors.New`, `commandAndRuleEngine.createNewBondB12Element`

### executeTCRuleSwap106 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleSwap106(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (matureElementToSwapIn testCaseModel.MatureElementStruct, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: TCRuleSwap106 What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
- Selector calls: `commandAndRuleEngine.verifySwapRuleAndConvertIntoMatureComponentElementModel`, `errors.New`, `commandAndRuleEngine.createNewBondB11fxElement`, `commandAndRuleEngine.createNewBondB11lxElement`

### executeTCRuleSwap107 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleSwap107(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (matureElementToSwapIn testCaseModel.MatureElementStruct, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: TCRuleSwap107 What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
- Selector calls: `commandAndRuleEngine.verifySwapRuleAndConvertIntoMatureComponentElementModel`, `errors.New`, `commandAndRuleEngine.createNewBondB11fxElement`, `commandAndRuleEngine.createNewBondB11lElement`

### executeTCRuleSwap108 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleSwap108(testCaseUuid string, uuidToSwapOut string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (matureElementToSwapIn testCaseModel.MatureElementStruct, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: TCRuleSwap108 What to swap in 	What to swap out	with	In the following structure		Result after swapping	Rule
- Selector calls: `commandAndRuleEngine.verifySwapRuleAndConvertIntoMatureComponentElementModel`, `errors.New`, `commandAndRuleEngine.createNewBondB11fElement`, `commandAndRuleEngine.createNewBondB11lxElement`

### transformImmatureElementModelIntoMatureElementModel (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) transformImmatureElementModelIntoMatureElementModel(immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (matureElementModel testCaseModel.MatureElementStruct, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: Transforms a immature element model into a mature element model. This means that new UUIDs are created for each element in the component
- Selector calls: `uuid.New`, `errors.New`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
