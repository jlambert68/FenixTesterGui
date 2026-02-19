# testSuitesModel_SaveTestSuite.go

## File Overview
- Path: `testSuites/testSuitesModel/testSuitesModel_SaveTestSuite.go`
- Package: `testSuitesModel`
- Functions/Methods: `9`
- Imports: `11`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `SaveTestSuite`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/soundEngine`
- `errors`
- `fmt`
- `fyne.io/fyne/v2`
- `github.com/jinzhu/copier`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `github.com/jlambert68/FenixScriptEngine/testDataEngine`
- `google.golang.org/protobuf/encoding/protojson`
- `log`
- `strings`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### SaveTestSuite (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) SaveTestSuite() err error`
- Exported: `true`
- Control-flow features: `if, returns error`
- Doc: SaveTestSuite Send TestSuite to TestCaseBuilderServer to saved in the Database
- Selector calls: `testSuiteModel.IsTestSuiteChanged`, `fyne.CurrentApp`, `fmt.Sprintf`, `copier.CopyWithOption`, `err.Error`, `log.Fatalln`, `testSuiteModel.checkIfAllMandatoryFieldsHaveValues`, `testSuiteModel.generateTestSuiteBasicInformationMessageWhenSaving`

### generateTestSuiteBasicInformationMessageWhenSaving (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) generateTestSuiteBasicInformationMessageWhenSaving(supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct) (testSuiteBasicInformation *fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteBasicInformationMessage, testSuiteBasicInformationHash string, err error)`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Generates 'TestSuiteBasicInformation' to be added to full gRPC-message
- Selector calls: `protojson.Format`, `sharedCode.HashSingleValue`

### generateTestSuiteTestDataMessageWhenSaving (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) generateTestSuiteTestDataMessageWhenSaving(supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct) (testSuiteTestData *fenixGuiTestCaseBuilderServerGrpcApi.UsersChosenTestDataForTestSuiteMessage, testSuiteTestDataHash string, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: Generates 'UsersChosenTestDataForTestSuiteMessage' to be added to full gRPC-message
- Internal calls: `string`
- Selector calls: `protojson.Format`, `strings.ReplaceAll`, `sharedCode.HashSingleValue`

### generateTestSuitePreviewMessageWhenSaving (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) generateTestSuitePreviewMessageWhenSaving(supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct) (testSuitePreview *fenixGuiTestCaseBuilderServerGrpcApi.TestSuitePreviewMessage, testSuitePreviewHash string, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, switch, returns error`
- Doc: Generates 'TestSuitePreview' to be added to full gRPC-message
- Selector calls: `fmt.Sprintf`, `fenixGuiTestCaseBuilderServerGrpcApi.MetaDataSelectTypeEnum`, `log.Fatal`, `testSuiteModel.GetLastChangedByGcpLogin`, `testSuiteModel.GetLastChangedByComputerLogin`, `testSuiteModel.GetCreatedDate`

### generateTestSuiteMetaDataMessageWhenSaving (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) generateTestSuiteMetaDataMessageWhenSaving(supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct) (testSuiteMetaData *fenixGuiTestCaseBuilderServerGrpcApi.UserSpecifiedTestSuiteMetaDataMessage, testSuiteMetaDataHash string, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, switch, returns error`
- Doc: Generates 'TestSuiteMetaData' to be added to full gRPC-message
- Selector calls: `sharedCode.HashSingleValue`, `fenixGuiTestCaseBuilderServerGrpcApi.MetaDataSelectTypeEnum`, `fmt.Sprintf`, `log.Fatalln`, `sharedCode.HashValues`

### generateTestCasesInTestSuiteMessageWhenSaving (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) generateTestCasesInTestSuiteMessageWhenSaving(supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct) (testCasesInTestSuite *fenixGuiTestCaseBuilderServerGrpcApi.TestCasesInTestSuiteMessage, testCasesInTestSuiteHash string, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: Generates 'TestCasesInTestSuite' to be added to full gRPC-message
- Selector calls: `testCasesInTestSuiteMessage.GetDomainUuid`, `testCasesInTestSuiteMessage.GetDomainName`, `testCasesInTestSuiteMessage.GetTestCaseUuid`, `testCasesInTestSuiteMessage.GetTestCaseName`, `fmt.Sprintf`, `sharedCode.HashValues`, `sharedCode.HashSingleValue`

### generateTestSuiteDeleteDateMessageWhenSaving (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) generateTestSuiteDeleteDateMessageWhenSaving(supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct) (testSuiteDeleteDate string, testSuiteDeleteDateHash string, err error)`
- Exported: `false`
- Control-flow features: `returns error`
- Doc: Generates 'TestSuiteDeleteData' to be added to full gRPC-message
- Selector calls: `sharedCode.HashSingleValue`

### generateTestSuiteTypeMessageWhenSaving (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) generateTestSuiteTypeMessageWhenSaving(supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct) (testSuiteTypeMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteTypeMessage, testSuiteTypeHash string, err error)`
- Exported: `false`
- Control-flow features: `returns error`
- Doc: Generates 'TestCasesInTestSuite' to be added to full gRPC-message
- Internal calls: `string`
- Selector calls: `fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteTypeEnum`, `fmt.Sprintf`, `sharedCode.HashValues`

### generateTestSuiteImplementedFunctionsMapWhenSaving (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) generateTestSuiteImplementedFunctionsMapWhenSaving(supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct) (testSuiteImplementedFunctionsMap map[int32]bool, testSuiteImplementedFunctionsMapHash string, err error)`
- Exported: `false`
- Control-flow features: `for/range, returns error`
- Doc: Generates 'TestSuiteImplementedFunctionsMap' to be added to full gRPC-message
- Internal calls: `int32`
- Selector calls: `fmt.Sprintf`, `sharedCode.HashValues`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
