# availableBuildingBlockTree_load_gRPC_results_into_model.go

## File Overview
- Path: `gui/availableBuildingBlockTree_load_gRPC_results_into_model.go`
- Package: `gui`
- Functions/Methods: `14`
- Imports: `11`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/testCase/testCaseModel`
- `FenixTesterGui/testSuites/testSuitesModel`
- `encoding/json`
- `fmt`
- `github.com/jlambert68/Fast_BitFilter_MetaData/boolbits/bitmapper`
- `github.com/jlambert68/Fast_BitFilter_MetaData/boolbits/boolbits`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `github.com/jlambert68/FenixScriptEngine/testDataEngine`
- `github.com/sirupsen/logrus`
- `log`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### loadModelWithAvailableBuildingBlocks (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) loadModelWithAvailableBuildingBlocks(testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage, testCaseModeReference *testCaseModel.TestCasesModelsStruct)`
- Exported: `false`
- Control-flow features: `if`
- Doc: Load Available Building Blocks, TestInstructions and TestInstructionContainers, from GUI-server into testCaseModel
- Selector calls: `availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocksRegardingTestInstructions`, `availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocksRegardingTestInstructionContainers`, `availableBuildingBlocksModel.storeFullGrpcStructureForAvailableBuildingBlocks`, `availableBuildingBlocksModel.storeDomainsThatCanOwnTestCases`

### loadModelWithPinnedBuildingBlocks (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) loadModelWithPinnedBuildingBlocks(pinnedTestInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailablePinnedTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage)`
- Exported: `false`
- Control-flow features: `if`
- Doc: Load Pinned Building Blocks, TestInstructions and TestInstructionContainers, from GUI-server into testCaseModel
- Selector calls: `availableBuildingBlocksModel.loadModelWithPinnedBuildingBlocksRegardingTestInstructions`, `availableBuildingBlocksModel.loadModelWithPinnedBuildingBlocksRegardingTestInstructionContainers`

### loadModelWithAvailableBuildingBlocksRegardingTestInstructions (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) loadModelWithAvailableBuildingBlocksRegardingTestInstructions(testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Load all available TestInstructions Building Blocks
- Selector calls: `availableBuildingBlocksModel.generateUITreeName`

### loadModelWithAvailableBuildingBlocksRegardingTestInstructionContainers (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) loadModelWithAvailableBuildingBlocksRegardingTestInstructionContainers(testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Load all available TestInstructionContainers Building Blocks
- Selector calls: `availableBuildingBlocksModel.generateUITreeName`, `fmt.Println`

### loadModelWithPinnedBuildingBlocksRegardingTestInstructions (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) loadModelWithPinnedBuildingBlocksRegardingTestInstructions(pinnedTestInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailablePinnedTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Load all Pinned TestInstructions Building Blocks into testCaseModel

### loadModelWithPinnedBuildingBlocksRegardingTestInstructionContainers (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) loadModelWithPinnedBuildingBlocksRegardingTestInstructionContainers(pinnedTestInstructionsAndTestInstructionsContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailablePinnedTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Load all Pinned TestInstructions Building Blocks into testCaseModel

### storeFullGrpcStructureForAvailableBuildingBlocks (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) storeFullGrpcStructureForAvailableBuildingBlocks(testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.
	AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Store the full available Building Blocks Structure into the Available Building Blocks Model

### storeDomainsThatCanOwnTestCases (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) storeDomainsThatCanOwnTestCases(testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.
	AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage, testCaseModeReference *testCaseModel.TestCasesModelsStruct)`
- Exported: `false`
- Control-flow features: `for/range`
- Doc: Store list with Domains that can own a TestCase
- Selector calls: `testCaseModeReference.GenerateShortUuidFromFullUuid`, `tempDomainThatCanOwnTheTestCase.GetDomainUuid`, `tempDomainThatCanOwnTheTestCase.GetDomainName`

### storeTestCaseMetaDataPerDomain (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) storeTestCaseMetaDataPerDomain(testCaseMetaDataForDomainsToBeStored []*fenixGuiTestCaseBuilderServerGrpcApi.TestCaseAndTestSuiteMetaDataForOneDomainMessage, testCaseModeReference *testCaseModel.TestCasesModelsStruct)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Store list with TestCaseMEtaData per Domain
- Selector calls: `json.Unmarshal`, `testCaseMetaDataForDomain.GetDomainUuid`, `bitmapper.GenerateBitMaps`, `log.Fatalf`

### storeTestSuiteMetaDataPerDomain (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) storeTestSuiteMetaDataPerDomain(testSuiteMetaDataForDomainsToBeStored []*fenixGuiTestCaseBuilderServerGrpcApi.TestCaseAndTestSuiteMetaDataForOneDomainMessage)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Store list with TestSuiteMEtaData per Domain
- Selector calls: `json.Unmarshal`, `testSuiteMetaDataForDomain.GetDomainUuid`, `bitmapper.GenerateBitMaps`, `log.Fatalf`

### storeTemplateRepositoryApiUrls (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) storeTemplateRepositoryApiUrls(templateRepositoryApiUrlsToBeStored []*fenixGuiTestCaseBuilderServerGrpcApi.RepositoryApiUrlResponseMessage, testCaseModeReference *testCaseModel.TestCasesModelsStruct)`
- Exported: `false`
- Control-flow features: `for/range`
- Doc: Store list with TemplateRepositoryApiUrls
- Selector calls: `templateRepositoryApiUrlToBeStored.GetRepositoryApiUrlName`

### storeTestData (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) storeTestData(testDataFromSimpleTestDataAreaFiles []*fenixGuiTestCaseBuilderServerGrpcApi.TestDataFromOneSimpleTestDataAreaFileMessage, testCaseModeReference *testCaseModel.TestCasesModelsStruct)`
- Exported: `false`
- Control-flow features: `for/range`
- Doc: Store TestData that user can use within TestCasesMapPtr
- Selector calls: `rawHeader.GetShouldHeaderActAsFilter`, `rawHeader.GetHeaderName`, `rawHeader.GetHeaderUiName`, `simpleTestDataRow.GetTestDataValue`, `testDataFromOneSimpleTestDataAreaFile.GetTestDataDomainUuid`, `testDataFromOneSimpleTestDataAreaFile.GetTestDataDomainName`, `testDataFromOneSimpleTestDataAreaFile.GetTestDataDomainTemplateName`, `testDataFromOneSimpleTestDataAreaFile.GetTestDataAreaUuid`

### storeUsersAvailableExecutionDomains (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) storeUsersAvailableExecutionDomains(executionDomainsThatCanReceiveDirectTargetedTestInstructions []*fenixGuiTestCaseBuilderServerGrpcApi.
	ExecutionDomainsThatCanReceiveDirectTargetedTestInstructionsMessage, testCaseModeReference *testCaseModel.TestCasesModelsStruct)`
- Exported: `false`
- Control-flow features: `for/range`
- Doc: Store Users available ExecutionDomains to be used with Fenix-created TestInstructions that should be sent to other Domain then Fenix
- Selector calls: `executionDomainsThatCanReceiveDirectTargetedTestInstruction.GetNameUsedInGui`

### convertGrpcElementModelIntoTestCaseElementModel (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) convertGrpcElementModelIntoTestCaseElementModel(immatureGrpcElementModelMessage *fenixGuiTestCaseBuilderServerGrpcApi.ImmatureElementModelMessage) immatureTestCaseElementModel testCaseModel.ImmatureElementStruct`
- Exported: `false`
- Control-flow features: `for/range`
- Doc: Convert gRPC-message for TI or TIC into model used within the TestCase-model

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
