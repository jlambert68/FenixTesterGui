# testCaseModel_exposedFunctions.go

## File Overview
- Path: `testCase/testCaseModel/testCaseModel_exposedFunctions.go`
- Package: `testCaseModel`
- Functions/Methods: `11`
- Imports: `4`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `CreateTextualTestCase`
- `GenerateShortUuidFromFullUuid`
- `GetArrayOfTestCaseTreeNodeChildrenData`
- `GetTestCaseNameUuid`
- `GetTestCaseTreeNodeChildData`
- `GetTreeViewModelForTestCase`
- `GetUuidFromUiName`
- `ListAllAvailableBuildingBlocksInTestCase`
- `ListAvailableTestCases`
- `UpdateTreeViewModelForTestCase`
- `VerifyThatThereAreNoZombieElementsInTestCaseModel`

## Imports
- `errors`
- `fmt`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `strings`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### CreateTextualTestCase (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) CreateTextualTestCase(testCaseUuid string) (textualTestCaseSimple string, textualTestCaseComplex string, textualTestCaseExtended string, err error)`
- Exported: `true`
- Control-flow features: `if, for/range, switch, returns error`
- Doc: CreateTextualTestCase Create Textual TestCase Representations
- Internal calls: `int32`
- Selector calls: `errors.New`, `fmt.Sprintf`, `strings.Index`, `testCaseModel.GenerateShortUuidFromFullUuid`, `testCaseModel.UpdateTreeViewModelForTestCase`, `testCaseModel.recursiveTextualTestCaseModelExtractor`

### GenerateShortUuidFromFullUuid (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) GenerateShortUuidFromFullUuid(fullUuid string) shortUuid string`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: GenerateShortUuidFromFullUuid Generate a short version of the UUID to be used in GUI

### GetArrayOfTestCaseTreeNodeChildrenData (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) GetArrayOfTestCaseTreeNodeChildrenData(nodeUuid string, testCaseUuid string) childrenUuidSlice []string`
- Exported: `true`
- Control-flow features: `for/range`
- Doc: GetArrayOfTestCaseTreeNodeChildrenData Returns a slice of child-Uuid:s to a parent Uuid
- Selector calls: `testCaseModel.GetTreeViewModelForTestCase`

### GetTestCaseNameUuid (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) GetTestCaseNameUuid(testCaseUuid string) (testCaseName string, err error)`
- Exported: `true`
- Control-flow features: `if, returns error`
- Doc: GetTestCaseNameUuid Retrieve TestCaseName from TestCase based on UUID
- Selector calls: `errors.New`, `fmt.Sprintf`

### GetTestCaseTreeNodeChildData (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) GetTestCaseTreeNodeChildData(nodeUuid string, testCaseUuid string) treeNodeChildData TestCaseModelAdaptedForUiTreeDataStruct`
- Exported: `true`
- Control-flow features: `if`
- Doc: GetTestCaseTreeNodeChildData Returns a slice of child-Uuid:s to a parent Uuid
- Selector calls: `testCaseModel.GetTreeViewModelForTestCase`

### GetTreeViewModelForTestCase (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) GetTreeViewModelForTestCase(testCaseUuid string) (treeViewModel map[string][]TestCaseModelAdaptedForUiTreeDataStruct, err error)`
- Exported: `true`
- Control-flow features: `if, returns error`
- Doc: GetTreeViewModelForTestCase Updates, and returns, the model adapted for a Tree View representation of the TestCase
- Selector calls: `errors.New`, `fmt.Sprintf`

### GetUuidFromUiName (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) GetUuidFromUiName(testCaseUuid string, uiName string) (elementUuid string, err error)`
- Exported: `true`
- Control-flow features: `if, for/range, switch, returns error`
- Doc: GetUuidFromUiName Finds the UUID for from a UI-name like ' B0_BOND [3c8a3bc] [BOND] to live forever..'
- Selector calls: `errors.New`, `fmt.Sprintf`, `strings.Index`

### ListAllAvailableBuildingBlocksInTestCase (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) ListAllAvailableBuildingBlocksInTestCase(testCaseUuid string) (availableBuidlingBlocksInTestCaseList []string, err error)`
- Exported: `true`
- Control-flow features: `if, for/range, switch, returns error`
- Doc: ListAllAvailableBuildingBlocksInTestCase List ALL Building Blocks in TestCase
- Selector calls: `errors.New`, `fmt.Sprintf`, `testCaseModel.generateUINameForTestCaseElement`

### ListAvailableTestCases (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) ListAvailableTestCases() availableTestCasesAsList []string`
- Exported: `true`
- Control-flow features: `if, for/range`
- Doc: ListAvailableTestCases List all available TestCase in TestCasesModel

### UpdateTreeViewModelForTestCase (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) UpdateTreeViewModelForTestCase(testCaseUuid string) err error`
- Exported: `true`
- Control-flow features: `if, returns error`
- Doc: UpdateTreeViewModelForTestCase Updates, and returns, the model adapted for a Tree View representation of the TestCase
- Selector calls: `errors.New`, `fmt.Sprintf`, `testCaseModel.recursiveGraphicalTestCaseTreeModelExtractor`

### VerifyThatThereAreNoZombieElementsInTestCaseModel (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) VerifyThatThereAreNoZombieElementsInTestCaseModel(testCaseUuid string) err error`
- Exported: `true`
- Control-flow features: `if, for/range, returns error`
- Doc: VerifyThatThereAreNoZombieElementsInTestCaseModel Verify that all UUIDs are correct in TestCaseModel. Meaning that no empty uuid is allowed and they all are correct
- Selector calls: `errors.New`, `fmt.Sprintf`, `testCaseModel.recursiveZombieElementSearchInTestCaseModel`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
