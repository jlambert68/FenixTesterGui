# availableBuildingBlockTree_model.go

## File Overview
- Path: `gui/availableBuildingBlockTree_model.go`
- Package: `gui`
- Functions/Methods: `23`
- Imports: `7`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/testCase/testCaseModel`
- `errors`
- `fmt`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `sort`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### getAvailableBuildingBlocksModel (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) getAvailableBuildingBlocksModel() map[string][]string`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Gets the testCaseModel used to drive the Available Building Blocks-Tree

### makeTreeUIModel (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) makeTreeUIModel()`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Generate the testCaseModel used to drive the Available Building Blocks-Tree
- Selector calls: `availableBuildingBlocksModel.getPinnedBuildingBlocksTreeNamesFromModel`, `availableBuildingBlocksModel.getAvailableDomainTreeNamesFromModel`, `availableBuildingBlocksModel.getAvailableDomainsFromModel`, `sort.SliceStable`, `availableBuildingBlocksModel.generateUITreeNameForTestInstructionsHeader`, `availableBuildingBlocksModel.generateUITreeNameForTestInstructionContainersHeader`, `availableBuildingBlocksModel.getAvailableTestInstructionTypesFromModel`, `availableBuildingBlocksModel.getAvailableTestInstructionContainerTypesFromModel`

### loadAvailableBuildingBlocksFromServer (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) loadAvailableBuildingBlocksFromServer(testCaseModeReference *testCaseModel.TestCasesModelsStruct)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Load all Available Building Blocks from Gui-server
- Selector calls: `availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocks`, `errors.New`, `fmt.Sprintf`, `fmt.Println`, `fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionInformationMessage_AvailableDropZoneMessage_DropZonePreSetTestInstructionAttributeMessage_AttributeActionCommandEnum`, `availableBuildingBlocksModel.storeUsersAvailableExecutionDomains`, `testInstructionsAndTestContainersMessage.GetExecutionDomainsThatCanReceiveDirectTargetedTestInstructions`

### loadTemplateRepositoryApiUrls (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) loadTemplateRepositoryApiUrls(testCaseModeReference *testCaseModel.TestCasesModelsStruct)`
- Exported: `false`
- Control-flow features: `if`
- Doc: Load list with TemplateRepositoryApiUrls form GUI-server
- Selector calls: `allRepositoryApiUrlsResponseMessage.GetAckNackResponse`, `availableBuildingBlocksModel.storeTemplateRepositoryApiUrls`, `allRepositoryApiUrlsResponseMessage.GetRepositoryApiUrls`

### loadTestCaseAndTestSuiteMetaData (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) loadTestCaseAndTestSuiteMetaData(testCaseModeReference *testCaseModel.TestCasesModelsStruct)`
- Exported: `false`
- Control-flow features: `if`
- Doc: Load a list of TestCaseMetaData and TestSuiteMetaData per Domain form GUI-server
- Selector calls: `listTestCaseAndTestSuiteMetaDataResponseMessage.GetAckNackResponse`, `availableBuildingBlocksModel.storeTestCaseMetaDataPerDomain`, `listTestCaseAndTestSuiteMetaDataResponseMessage.GetTestCaseAndTestSuiteMetaDataForDomains`, `availableBuildingBlocksModel.storeTestSuiteMetaDataPerDomain`

### loadTestData (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) loadTestData(testCaseModeReference *testCaseModel.TestCasesModelsStruct)`
- Exported: `false`
- Control-flow features: `if`
- Doc: Load list with loadTestData form GUI-server
- Selector calls: `allTestDataForTestDataAreasResponseMessage.GetAckNackResponse`, `availableBuildingBlocksModel.storeTestData`, `allTestDataForTestDataAreasResponseMessage.GetTestDataFromSimpleTestDataAreaFiles`

### loadPinnedBuildingBlocksFromServer (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) loadPinnedBuildingBlocksFromServer()`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Load all Pinned Building Blocks from Gui-server
- Selector calls: `availableBuildingBlocksModel.loadModelWithPinnedBuildingBlocks`

### savePinnedBuildingBlocksFromServer (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) savePinnedBuildingBlocksFromServer() err error`
- Exported: `false`
- Control-flow features: `if, for/range, switch, returns error`
- Doc: Save all Pinned Building Blocks to Gui-server
- Selector calls: `fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum`, `errors.New`

### generateUITreeName (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) generateUITreeName(node AvailableBuildingBlocksForUITreeNodesStruct, domainName string) (treeName string, pinnedTreeName string)`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Generate UI Tree name for 'Domain', TestInstructionType, TestInstruction, TestInstructionContainerType and TestInstructionContainer for the Available Building Blocks UI-Tree

### generateUITreeNameForTestInstructionsHeader (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) generateUITreeNameForTestInstructionsHeader(domain AvailableBuildingBlocksForUITreeNodesStruct) treeName string`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Generate UI Tree name for 'TestInstructionsHeader' for Available Building Blocks

### generateUITreeNameForTestInstructionContainersHeader (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) generateUITreeNameForTestInstructionContainersHeader(domain AvailableBuildingBlocksForUITreeNodesStruct) treeName string`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Generate UI Tree name for 'TestInstructionContainersHeader' for Available Building Blocks

### getAvailableDomainTreeNamesFromModel (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) getAvailableDomainTreeNamesFromModel() availableDomainTreeNamesList []string`
- Exported: `false`
- Control-flow features: `for/range`
- Doc: Extract all 'Domains', with Names suited for Tree-testCaseModel, for the testCaseModel tha underpins the UI Tree for Available Building Blocks
- Selector calls: `availableBuildingBlocksModel.getAvailableDomainsFromModel`

### getAvailableDomainsFromModel (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) getAvailableDomainsFromModel() availableDomains []AvailableBuildingBlocksForUITreeNodesStruct`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Extract all 'Domains', with Names suited for Tree-testCaseModel, for the testCaseModel tha underpins the UI Tree for Available Building Blocks

### getAvailableTestInstructionTypesFromModel (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) getAvailableTestInstructionTypesFromModel(domain AvailableBuildingBlocksForUITreeNodesStruct) availableTestInstructionTypes []AvailableBuildingBlocksForUITreeNodesStruct`
- Exported: `false`
- Control-flow features: `for/range`
- Doc: Extract all 'TestInstructionTypes', per Domain, with Names suited for Tree-testCaseModel

### getAvailableTestInstructionContainerTypesFromModel (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) getAvailableTestInstructionContainerTypesFromModel(domain AvailableBuildingBlocksForUITreeNodesStruct) availableTestInstructionContainerTypes []AvailableBuildingBlocksForUITreeNodesStruct`
- Exported: `false`
- Control-flow features: `for/range`
- Doc: Extract all 'TestInstructionContainerTypes', per Domain, with Names suited for Tree-testCaseModel

### getAvailableTestInstructionsFromModel (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) getAvailableTestInstructionsFromModel(domain AvailableBuildingBlocksForUITreeNodesStruct, testInstructionType AvailableBuildingBlocksForUITreeNodesStruct) availableTestInstructions []AvailableBuildingBlocksForUITreeNodesStruct`
- Exported: `false`
- Control-flow features: `for/range`
- Doc: Extract all 'TestInstructions', per TestInstructionType, with Names suited for Tree-testCaseModel

### getAvailableTestInstructionContainersFromModel (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) getAvailableTestInstructionContainersFromModel(domain AvailableBuildingBlocksForUITreeNodesStruct, testInstructionContainerType AvailableBuildingBlocksForUITreeNodesStruct) availableTestInstructionContainers []AvailableBuildingBlocksForUITreeNodesStruct`
- Exported: `false`
- Control-flow features: `for/range`
- Doc: Extract all 'TestInstructionContainers', per TestInstructionContainerType, with Names suited for Tree-testCaseModel

### getPinnedBuildingBlocksTreeNamesFromModel (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) getPinnedBuildingBlocksTreeNamesFromModel() pinnedBuildingBlocks []string`
- Exported: `false`
- Control-flow features: `for/range`
- Doc: Extract all 'Pinned TestInstructions' suited for Tree-testCaseModel

### verifyBeforePinTestInstructionOrTestInstructionContainer (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) verifyBeforePinTestInstructionOrTestInstructionContainer(nameInAvailableBuildingBlocksTree string, onlyForVerifying bool) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Verify that it is possible to Pin one Available Building Block (TestInstruction or TestInstructionContainer, if it isn't already pinned
- Selector calls: `errors.New`

### pinTestInstructionOrTestInstructionContainer (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) pinTestInstructionOrTestInstructionContainer(nameInAvailableBuildingBlocksTree string) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Pin one Available Building Block (TestInstruction or TestInstructionContainer, if it isn't already pinned
- Selector calls: `availableBuildingBlocksModel.verifyBeforePinTestInstructionOrTestInstructionContainer`

### verifyBeforeUnPinTestInstructionOrTestInstructionContainer (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) verifyBeforeUnPinTestInstructionOrTestInstructionContainer(pinnedNameInUITree string, onlyForVerifying bool) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Verify that it is possible to Unpin one pinned Available Building Block (TestInstruction or TestInstructionContainer
- Selector calls: `errors.New`

### unPinTestInstructionOrTestInstructionContainer (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) unPinTestInstructionOrTestInstructionContainer(pinnedNameInUITree string) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Unpin one pinned Available Building Block (TestInstruction or TestInstructionContainer
- Selector calls: `availableBuildingBlocksModel.verifyBeforeUnPinTestInstructionOrTestInstructionContainer`

### listAllAvailableBuidlingBlocks (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) listAllAvailableBuidlingBlocks() availableBuidlingBlocksList []string`
- Exported: `false`
- Control-flow features: `for/range, switch`
- Doc: List all Available Building Block (TestInstruction or TestInstructionContainer

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
