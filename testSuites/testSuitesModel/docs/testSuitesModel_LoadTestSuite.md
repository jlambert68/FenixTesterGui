# testSuitesModel_LoadTestSuite.go

## File Overview
- Path: `testSuites/testSuitesModel/testSuitesModel_LoadTestSuite.go`
- Package: `testSuitesModel`
- Functions/Methods: `10`
- Imports: `6`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `LoadFullTestSuiteFromDatabase`

## Imports
- `errors`
- `fmt`
- `github.com/jinzhu/copier`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `github.com/jlambert68/FenixScriptEngine/testDataEngine`
- `log`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### LoadFullTestSuiteFromDatabase (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) LoadFullTestSuiteFromDatabase(testSuiteUuid string) err error`
- Exported: `true`
- Control-flow features: `if, returns error`
- Doc: LoadFullTestSuiteFromDatabase - Load the TestSuite from the Database into model
- Selector calls: `errors.New`, `fmt.Sprintf`, `fmt.Println`, `detailedTestSuiteResponse.GetDetailedTestSuite`, `testSuiteModel.extractTestSuiteImplementedFunctionsMap`, `testSuiteModel.generateTestSuiteBasicInformationMessageWhenLoading`, `err.Error`, `testSuiteModel.generateTestSuiteTestDataMessageWhenLoading`

### extractTestSuiteImplementedFunctionsMap (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) extractTestSuiteImplementedFunctionsMap(testSuiteImplementedFunctionsGrpc map[int32]bool) (testSuiteImplementedFunctionsMap map[testSuiteImplementedFunctionsType]bool, err error)`
- Exported: `false`
- Control-flow features: `for/range, returns error`
- Doc: Extract 'TestSuiteImplementedFunctionsMap' from gRPC-message
- Internal calls: `testSuiteImplementedFunctionsType`

### generateTestSuiteBasicInformationMessageWhenLoading (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) generateTestSuiteBasicInformationMessageWhenLoading(supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct, testSuiteBasicInformation *fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteBasicInformationMessage, updatedByAndWhenMessage *fenixGuiTestCaseBuilderServerGrpcApi.UpdatedByAndWhenMessage, messageHash string) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Generates 'TestSuiteBasicInformation' from gRPC-message
- Selector calls: `testSuiteBasicInformation.GetTestSuiteName`, `testSuiteBasicInformation.GetTestSuiteDescription`, `testSuiteBasicInformation.GetDomainUuid`, `testSuiteBasicInformation.GetDomainName`, `testSuiteBasicInformation.GetTestSuiteExecutionEnvironment`, `testSuiteBasicInformation.GetTestSuiteUuid`, `testSuiteBasicInformation.GetTestSuiteVersion`, `updatedByAndWhenMessage.GetCreatedByGcpLogin`

### generateTestSuiteTestDataMessageWhenLoading (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) generateTestSuiteTestDataMessageWhenLoading(supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct, testSuiteTestData *fenixGuiTestCaseBuilderServerGrpcApi.UsersChosenTestDataForTestSuiteMessage) err error`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: Generates 'UsersChosenTestDataForTestSuiteMessage' from gRPC-message
- Selector calls: `testDataEngine.TestDataPointRowUuidType`, `testDataPointRowValueSummaryGrpc.GetTestDataPointRowUuid`, `testDataEngine.TestDataPointRowValuesSummaryType`, `testDataPointRowValueSummaryGrpc.GetTestDataPointRowValuesSummary`, `testDataEngine.TestDataDomainUuidType`, `testDataRowGrpc.GetTestDataDomainUuid`, `testDataEngine.TestDataDomainNameType`, `testDataRowGrpc.GetTestDataDomainName`

### generateTestSuitePreviewMessageWhenLoading (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) generateTestSuitePreviewMessageWhenLoading(supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct, testSuitePreview *fenixGuiTestCaseBuilderServerGrpcApi.TestSuitePreviewMessage) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Generates 'TestSuitePreview' from gRPC-message

### generateTestSuiteMetaDataMessageWhenLoading (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) generateTestSuiteMetaDataMessageWhenLoading(supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct, testSuiteMetaData *fenixGuiTestCaseBuilderServerGrpcApi.UserSpecifiedTestSuiteMetaDataMessage) err error`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: Generates 'TestSuiteMetaData' from gRPC-message
- Internal calls: `MetaDataSelectType`
- Selector calls: `testSuiteMetaData.GetMetaDataGroupsMap`, `tempMetaDataGroupFromGrpc.GetMetaDataInGroupMap`, `testSuiteMetaData.GetCurrentSelectedDomainUuid`, `testSuiteMetaData.GetCurrentSelectedDomainName`

### generateTestCasesInTestSuiteMessageWhenLoading (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) generateTestCasesInTestSuiteMessageWhenLoading(supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct, testCasesInTestSuiteFromGrpc *fenixGuiTestCaseBuilderServerGrpcApi.TestCasesInTestSuiteMessage) err error`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: Generates 'TestCasesInTestSuite' from gRPC-message
- Selector calls: `testCasesInTestSuiteFromGrpc.GetTestCasesInTestSuite`, `tempTestCasesInTestSuiteFromGrpc.GetTestCaseUuid`

### generateTestSuiteDeleteDateMessageWhenLoading (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) generateTestSuiteDeleteDateMessageWhenLoading(supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct, testSuiteDeleteDateFromGrpc string) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Generates 'TestSuiteDeleteData' from gRPC-message

### generateTestSuiteTypeMessageWhenLoading (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) generateTestSuiteTypeMessageWhenLoading(supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct, testSuiteTypeMessageFromGrpc *fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteTypeMessage) err error`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Generates 'TestCasesInTestSuite' from gRPC-message
- Internal calls: `TestSuiteTypeType`, `TestSuiteTypeNameType`
- Selector calls: `testSuiteTypeMessageFromGrpc.GetTestSuiteType`, `testSuiteTypeMessageFromGrpc.GetTestSuiteTypeName`

### generateTestSuiteImplementedFunctionsMapWhenLoading (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) generateTestSuiteImplementedFunctionsMapWhenLoading(supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct, testSuiteImplementedFunctionsMapFromGrpc map[int32]bool) err error`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: Generates 'TestSuiteImplementedFunctionsMap' from gRPC-message
- Internal calls: `testSuiteImplementedFunctionsType`
- Selector calls: `errors.New`, `fmt.Sprintf`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
