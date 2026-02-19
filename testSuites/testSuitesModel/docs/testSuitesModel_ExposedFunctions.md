# testSuitesModel_ExposedFunctions.go

## File Overview
- Path: `testSuites/testSuitesModel/testSuitesModel_ExposedFunctions.go`
- Package: `testSuitesModel`
- Functions/Methods: `15`
- Imports: `7`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `DoBothOwnerDomainAndTestEnvironmentHaveValues`
- `GenerateNewTestSuiteModelObject`
- `GetCreatedByComputerLogin`
- `GetCreatedByGcpLogin`
- `GetCreatedDate`
- `GetLastChangedByComputerLogin`
- `GetLastChangedByGcpLogin`
- `GetLastChangedDate`
- `GetTestSuiteUuid`
- `GetTestSuiteVersion`
- `HasLockButtonBeenClickedAndBothOwnerDomainAndTestEnvironmentHaveValues`
- `LockButtonHasBeenClicked`
- `OwnerDomainHasValue`
- `TestEnvironmentHasValue`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/testCase/testCaseModel`
- `github.com/google/uuid`
- `github.com/jlambert68/Fast_BitFilter_MetaData/boolbits/boolbits`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `github.com/jlambert68/FenixScriptEngine/testDataEngine`
- `strconv`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### DoBothOwnerDomainAndTestEnvironmentHaveValues (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) DoBothOwnerDomainAndTestEnvironmentHaveValues() hasValue bool`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: DoBothOwnerDomainAndTestEnvironmentHaveValues Do both of OwnerDomain and TestEnvironmentHasValue have their values selected by the user

### GenerateNewTestSuiteModelObject
- Signature: `func GenerateNewTestSuiteModelObject(existingTestSuiteUuid string, testCasesModel *testCaseModel.TestCasesModelsStruct) newTestSuiteModel *TestSuiteModelStruct`
- Exported: `true`
- Control-flow features: `if`
- Doc: GenerateNewTestSuiteModelObject Generated s new TestSuiteModel-object
- Internal calls: `createEmptyAndInitiatedTestSuiteModel`
- Selector calls: `uuidGenerator.New`

### GetCreatedByComputerLogin (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) GetCreatedByComputerLogin() string`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: GetCreatedByComputerLogin Gets the person that is logged into the computer

### GetCreatedByGcpLogin (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) GetCreatedByGcpLogin() string`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: GetCreatedByGcpLogin Gets the person that did log in towards GCP

### GetCreatedDate (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) GetCreatedDate() string`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: GetCreatedDate Gets the date when the TestSuite was first created

### GetLastChangedByComputerLogin (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) GetLastChangedByComputerLogin() string`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: GetLastChangedByComputerLogin Gets the person that is logged into the computer when TestSuite was last changed and saved

### GetLastChangedByGcpLogin (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) GetLastChangedByGcpLogin() string`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: GetLastChangedByGcpLogin Gets the person that did log in towards GCP when TestSuite was last changed and saved

### GetLastChangedDate (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) GetLastChangedDate() string`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: GetLastChangedDate Gets the date when the TestSuite was last changed and saved

### GetTestSuiteUuid (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) GetTestSuiteUuid() string`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: GetTestSuiteUuid Gets the TestSuites Uuid

### GetTestSuiteVersion (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) GetTestSuiteVersion() string`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: GetTestSuiteVersion Gets the TestSuites Version
- Internal calls: `int`
- Selector calls: `strconv.Itoa`

### HasLockButtonBeenClickedAndBothOwnerDomainAndTestEnvironmentHaveValues (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) HasLockButtonBeenClickedAndBothOwnerDomainAndTestEnvironmentHaveValues() hasValue bool`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: HasLockButtonBeenClickedAndBothOwnerDomainAndTestEnvironmentHaveValues Has Locked been clicked and both of  OwnerDomain and TestEnvironmentHasValue have their values selected by the user

### LockButtonHasBeenClicked (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) LockButtonHasBeenClicked()`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: LockButtonHasBeenClicked Store if LockButton has been clicked by the user

### OwnerDomainHasValue (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) OwnerDomainHasValue(hasValue bool)`
- Exported: `true`
- Control-flow features: `if`
- Doc: OwnerDomainHasValue Store if OwnerDomain has any value selected by the user

### TestEnvironmentHasValue (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) TestEnvironmentHasValue(hasValue bool)`
- Exported: `true`
- Control-flow features: `if`
- Doc: TestEnvironmentHasValue Store if TestEnvironmentHasValue has any value selected by the user

### createEmptyAndInitiatedTestSuiteModel
- Signature: `func createEmptyAndInitiatedTestSuiteModel(testCasesModel *testCaseModel.TestCasesModelsStruct) emptyAndInitiatedTestSuiteModel *TestSuiteModelStruct`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: createEmptyAndInitiatedTestSuiteModel Generates a fully initiated TestSuiteModelStruct

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
