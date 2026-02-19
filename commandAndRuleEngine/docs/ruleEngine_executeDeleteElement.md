# ruleEngine_executeDeleteElement.go

## File Overview
- Path: `commandAndRuleEngine/ruleEngine_executeDeleteElement.go`
- Package: `commandAndRuleEngine`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `21`
- Imports: `4`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `FenixTesterGui/testCase/testCaseModel`
- `errors`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### executeTCRuleDeletion101 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleDeletion101(testCaseUuid string, uuidToDelete string) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: What to remove			Remove in structure				Result after deletion		Rule n= TIC(X)				B1-n-B1							B0							TCRuleDeletion101
- External calls: `commandAndRuleEngine.createNewBondB0Element`, `commandAndRuleEngine.recursiveDeleteOfChildElements`, `errors.New`

### executeTCRuleDeletion102 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleDeletion102(testCaseUuid string, uuidToDelete string) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: What to remove			Remove in structure				Result after deletion		Rule n=TI or TIC(X)			B11f-n-B11l						B10							TCRuleDeletion102
- External calls: `commandAndRuleEngine.createNewBondB10Element`, `commandAndRuleEngine.recursiveDeleteOfChildElements`, `errors.New`

### executeTCRuleDeletion103 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleDeletion103(testCaseUuid string, uuidToDelete string) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: What to remove			Remove in structure				Result after deletion		Rule n=TI or TIC(X)			B11fx-n-B11lx					B10*x*						TCRuleDeletion103
- External calls: `commandAndRuleEngine.createNewBondB10oxoElement`, `commandAndRuleEngine.recursiveDeleteOfChildElements`, `errors.New`

### executeTCRuleDeletion104 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleDeletion104(testCaseUuid string, uuidToDelete string) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: What to remove			Remove in structure				Result after deletion		Rule n=TI or TIC(X)			B11f-n-B11lx					B10x*						TCRuleDeletion104
- External calls: `commandAndRuleEngine.createNewBondB10xoElement`, `commandAndRuleEngine.recursiveDeleteOfChildElements`, `errors.New`

### executeTCRuleDeletion105 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleDeletion105(testCaseUuid string, uuidToDelete string) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: What to remove			Remove in structure				Result after deletion		Rule n=TI or TIC(X)			B11fx-n-B11l					B10*x						TCRuleDeletion105
- External calls: `commandAndRuleEngine.createNewBondB10oxElement`, `commandAndRuleEngine.recursiveDeleteOfChildElements`, `errors.New`

### executeTCRuleDeletion106 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleDeletion106(testCaseUuid string, uuidToDelete string) err error`
- Exported: `false`
- Control-flow features: `returns error`
- Doc: What to remove			Remove in structure				Result after deletion		Rule n=TI or TIC(X)			B11f-n-B12-X					B11f-X						TCRuleDeletion106
- External calls: `commandAndRuleEngine.executeTCRuleDeletion_106_107_108_109`

### executeTCRuleDeletion107 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleDeletion107(testCaseUuid string, uuidToDelete string) err error`
- Exported: `false`
- Control-flow features: `returns error`
- Doc: What to remove			Remove in structure				Result after deletion		Rule n=TI or TIC(X)			B11fx-n-B12x-X					B11fx-X						TCRuleDeletion107
- External calls: `commandAndRuleEngine.executeTCRuleDeletion_106_107_108_109`

### executeTCRuleDeletion108 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleDeletion108(testCaseUuid string, uuidToDelete string) err error`
- Exported: `false`
- Control-flow features: `returns error`
- Doc: What to remove			Remove in structure				Result after deletion		Rule n=TI or TIC(X)			B11f-n-B12x-X					B11fx-X						TCRuleDeletion108
- External calls: `commandAndRuleEngine.executeTCRuleDeletion_106_107_108_109`

### executeTCRuleDeletion109 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleDeletion109(testCaseUuid string, uuidToDelete string) err error`
- Exported: `false`
- Control-flow features: `returns error`
- Doc: What to remove			Remove in structure				Result after deletion		Rule n=TI or TIC(X)			B11fx-n-B12-X					B11fx-X						TCRuleDeletion109
- External calls: `commandAndRuleEngine.executeTCRuleDeletion_106_107_108_109`

### executeTCRuleDeletion110 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleDeletion110(testCaseUuid string, uuidToDelete string) err error`
- Exported: `false`
- Control-flow features: `returns error`
- Doc: What to remove			Remove in structure				Result after deletion		Rule n=TI or TIC(X)			X-B12-n-B11l					X-B11l						TCRuleDeletion110
- External calls: `commandAndRuleEngine.executeTCRuleDeletion_110_111_112_113`

### executeTCRuleDeletion111 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleDeletion111(testCaseUuid string, uuidToDelete string) err error`
- Exported: `false`
- Control-flow features: `returns error`
- Doc: What to remove			Remove in structure				Result after deletion		Rule n=TI or TIC(X)			X-B12x-n-B11lx					X-B11lx						TCRuleDeletion111
- External calls: `commandAndRuleEngine.executeTCRuleDeletion_110_111_112_113`

### executeTCRuleDeletion112 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleDeletion112(testCaseUuid string, uuidToDelete string) err error`
- Exported: `false`
- Control-flow features: `returns error`
- Doc: What to remove			Remove in structure				Result after deletion		Rule n=TI or TIC(X)			X-B12-n-B11lx					X-B11lx						TCRuleDeletion112
- External calls: `commandAndRuleEngine.executeTCRuleDeletion_110_111_112_113`

### executeTCRuleDeletion113 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleDeletion113(testCaseUuid string, uuidToDelete string) err error`
- Exported: `false`
- Control-flow features: `returns error`
- Doc: What to remove			Remove in structure				Result after deletion		Rule n=TI or TIC(X)			X-B12x-n-B11l					X-B11lx						TCRuleDeletion113
- External calls: `commandAndRuleEngine.executeTCRuleDeletion_110_111_112_113`

### executeTCRuleDeletion114 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleDeletion114(testCaseUuid string, uuidToDelete string) err error`
- Exported: `false`
- Control-flow features: `returns error`
- Doc: What to remove			Remove in structure				Result after deletion		Rule n=TI or TIC(X)			X-B12-n-B12-X					X-B12-X						TCRuleDeletion114
- External calls: `commandAndRuleEngine.executeTCRuleDeletion_114_115_116_117`

### executeTCRuleDeletion115 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleDeletion115(testCaseUuid string, uuidToDelete string) err error`
- Exported: `false`
- Control-flow features: `returns error`
- Doc: What to remove			Remove in structure				Result after deletion		Rule n=TI or TIC(X)			X-B12x-n-B12x-X					X-B12x-X					TCRuleDeletion115
- External calls: `commandAndRuleEngine.executeTCRuleDeletion_114_115_116_117`

### executeTCRuleDeletion116 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleDeletion116(testCaseUuid string, uuidToDelete string) err error`
- Exported: `false`
- Control-flow features: `returns error`
- Doc: What to remove			Remove in structure				Result after deletion		Rule n=TI or TIC(X)			X-B12-n-B12x-X					X-B12x-X					TCRuleDeletion116
- External calls: `commandAndRuleEngine.executeTCRuleDeletion_114_115_116_117`

### executeTCRuleDeletion117 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleDeletion117(testCaseUuid string, uuidToDelete string) err error`
- Exported: `false`
- Control-flow features: `returns error`
- Doc: What to remove			Remove in structure				Result after deletion		Rule n=TI or TIC(X)			X-B12x-n-B12-X					X-B12x-X					TCRuleDeletion117
- External calls: `commandAndRuleEngine.executeTCRuleDeletion_114_115_116_117`

### executeTCRuleDeletion_106_107_108_109 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleDeletion_106_107_108_109(testCaseUuid string, uuidToDelete string) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Multi-deletion-rule function What to remove			Remove in structure				Result after deletion		Rule
- External calls: `commandAndRuleEngine.recursiveDeleteOfChildElements`, `errors.New`

### executeTCRuleDeletion_110_111_112_113 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleDeletion_110_111_112_113(testCaseUuid string, uuidToDelete string) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Multi-deletion-rule function What to remove			Remove in structure				Result after deletion		Rule
- External calls: `commandAndRuleEngine.recursiveDeleteOfChildElements`, `errors.New`

### executeTCRuleDeletion_114_115_116_117 (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) executeTCRuleDeletion_114_115_116_117(testCaseUuid string, uuidToDelete string) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Multi-deletion-rule function What to remove			Remove in structure				Result after deletion		Rule
- External calls: `commandAndRuleEngine.recursiveDeleteOfChildElements`, `errors.New`

### recursiveDeleteOfChildElements (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) recursiveDeleteOfChildElements(currentTestCase *testCaseModel.TestCaseModelStruct, elementsUuid string) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Remove all children, in TestCase-model, for specific Element
- External calls: `commandAndRuleEngine.recursiveDeleteOfChildElements`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
