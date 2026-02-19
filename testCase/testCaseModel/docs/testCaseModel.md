# testCaseModel.go

## File Overview
- Path: `testCase/testCaseModel/testCaseModel.go`
- Package: `testCaseModel`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `6`
- Imports: `3`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `errors`
- `fmt`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### findElementInSliceAndRemove
- Signature: `func findElementInSliceAndRemove(sliceToWorkOn *[]string, uuid string) returnSlice *[]string`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Remove 'uuid' from slice

### generateUINameForTestCaseElement (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) generateUINameForTestCaseElement(element *MatureTestCaseModelElementStruct) elementUiName string`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Generate name to be used when presenting TestCase Element

### recursiveGraphicalTestCaseTreeModelExtractor (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) recursiveGraphicalTestCaseTreeModelExtractor(testCaseUuid string, currentElementsUuid string, treeViewNodeChildrenIn []TestCaseModelAdaptedForUiTreeDataStruct) (treeViewNodeChildrenOut []TestCaseModelAdaptedForUiTreeDataStruct, err error)`
- Exported: `false`
- Control-flow features: `if, switch, returns error`
- Doc: Generate the slice with the elements in the TestCase. Order is the same as in the Textual Representation of the TestCase
- External calls: `errors.New`, `fmt.Sprintf`, `testCaseModel.recursiveGraphicalTestCaseTreeModelExtractor`, `testCaseModel.reverseSliceOfNodeObjects`

### recursiveTextualTestCaseModelExtractor (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) recursiveTextualTestCaseModelExtractor(testCaseUuid string, elementsUuid string, testCaseModelElementsIn []fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage) (testCaseModelElementsIOut []fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage, err error)`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Generate the slice with the elements in the TestCase. Order is the same as in the Textual Representation of the TestCase
- External calls: `errors.New`, `testCaseModel.recursiveTextualTestCaseModelExtractor`

### recursiveZombieElementSearchInTestCaseModel (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) recursiveZombieElementSearchInTestCaseModel(testCaseUuid string, elementsUuid string, allUuidKeys []string) (processedAllUuidKeys []string, err error)`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Verify all children, in TestCaseElement-model and remove the found element from 'allUuidKeys'
- Internal calls: `findElementInSliceAndRemove`
- External calls: `errors.New`, `testCaseModel.recursiveZombieElementSearchInTestCaseModel`

### reverseSliceOfNodeObjects (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) reverseSliceOfNodeObjects(inSlice []TestCaseModelAdaptedForUiTreeDataStruct) outSlice []TestCaseModelAdaptedForUiTreeDataStruct`
- Exported: `false`
- Control-flow features: `for/range`
- Doc: Reverse a slice of strings

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
