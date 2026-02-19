# ruleEngine_createNewElements.go

## File Overview
- Path: `commandAndRuleEngine/ruleEngine_createNewElements.go`
- Package: `commandAndRuleEngine`
- Functions/Methods: `13`
- Imports: `2`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `github.com/google/uuid`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### createNewBondB0Element (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) createNewBondB0Element() newBondB0Element fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Create a new B0-bond to be used in the TestCase-model
- Selector calls: `uuidGenerator.New`

### createNewBondB1fElement (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) createNewBondB1fElement(parentElementUuid string) newBondB1fElement fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage`
- Exported: `false`
- Control-flow features: `if`
- Doc: Create a new B1f-bond to be used in the TestCase-model
- Selector calls: `uuidGenerator.New`

### createNewBondB1lElement (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) createNewBondB1lElement(parentElementUuid string) newBondB1lElement fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage`
- Exported: `false`
- Control-flow features: `if`
- Doc: Create a new B1l-bond to be used in the TestCase-model
- Selector calls: `uuidGenerator.New`

### createNewBondB10Element (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) createNewBondB10Element(parentElementUuid string) newBondB10Element fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage`
- Exported: `false`
- Control-flow features: `if`
- Doc: Create a new B10-bond to be used in the TestCase-model
- Selector calls: `uuidGenerator.New`

### createNewBondB10oxoElement (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) createNewBondB10oxoElement(parentElementUuid string) newBondB10oxoElement fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage`
- Exported: `false`
- Control-flow features: `if`
- Doc: Create a new B10*x*-bond to be used in the TestCase-model
- Selector calls: `uuidGenerator.New`

### createNewBondB10xoElement (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) createNewBondB10xoElement(parentElementUuid string) newBondB10xoElement fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage`
- Exported: `false`
- Control-flow features: `if`
- Doc: Create a new B10x*-bond to be used in the TestCase-model
- Selector calls: `uuidGenerator.New`

### createNewBondB10oxElement (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) createNewBondB10oxElement(parentElementUuid string) newBondB10oxElement fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage`
- Exported: `false`
- Control-flow features: `if`
- Doc: Create a new B10*x-bond to be used in the TestCase-model
- Selector calls: `uuidGenerator.New`

### createNewBondB11fElement (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) createNewBondB11fElement(parentElementUuid string) newBondB11fElement fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage`
- Exported: `false`
- Control-flow features: `if`
- Doc: Create a new B11f-Bond to be used in the TestCase-model
- Selector calls: `uuidGenerator.New`

### createNewBondB11lElement (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) createNewBondB11lElement(parentElementUuid string) newBondB11lElement fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage`
- Exported: `false`
- Control-flow features: `if`
- Doc: Create a new B11l-Bond to be used in the TestCase-model
- Selector calls: `uuidGenerator.New`

### createNewBondB11fxElement (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) createNewBondB11fxElement(parentElementUuid string) newBondB11fxElement fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage`
- Exported: `false`
- Control-flow features: `if`
- Doc: Create a new B11fx-Bond to be used in the TestCase-model
- Selector calls: `uuidGenerator.New`

### createNewBondB11lxElement (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) createNewBondB11lxElement(parentElementUuid string) newBondB11lxElement fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage`
- Exported: `false`
- Control-flow features: `if`
- Doc: Create a new B11lx-Bond to be used in the TestCase-model
- Selector calls: `uuidGenerator.New`

### createNewBondB12Element (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) createNewBondB12Element(parentElementUuid string) newBondB12Element fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage`
- Exported: `false`
- Control-flow features: `if`
- Doc: Create a new B12-Bond to be used in the TestCase-model
- Selector calls: `uuidGenerator.New`

### createNewBondB12xElement (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) createNewBondB12xElement(parentElementUuid string) newBondB12xElement fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage`
- Exported: `false`
- Control-flow features: `if`
- Doc: Create a new B12x-Bond to be used in the TestCase-model
- Selector calls: `uuidGenerator.New`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
