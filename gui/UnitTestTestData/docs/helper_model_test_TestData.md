# helper_model_test_TestData.go

## File Overview
- Path: `gui/UnitTestTestData/helper_model_test_TestData.go`
- Package: `UnitTestTestData`
- Functions/Methods: `3`
- Imports: `5`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `InitLoggerForTest`
- `IsTestDataUsingCorrectTestDataProtoFileVersion`

## Imports
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `log`
- `os`
- `time`

## Declared Types
- None

## Declared Constants
- `loggingLevelForDebug`

## Declared Variables
- `PinnedTestInstructionsAndTestInstructionsContainersRespons_PBB001`
- `TestInstructionsAndTestInstructionsContainersRespons_ABB001`
- `TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_001`
- `TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_002`
- `TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_003`
- `TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_004`
- `TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_005`
- `TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_006`
- `TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_007`
- `TestInstructionsAndTestInstructionsContainersRespons_PBB001_ExpectedResultInModel_001`
- `TestInstructionsAndTestInstructionsContainersRespons_PBB001_ExpectedResultInModel_002`
- `TestInstructionsAndTestInstructionsContainersRespons_PBB001_ExpectedResultInModel_003`
- `TestInstructionsAndTestInstructionsContainersRespons_PBB002_ExpectedResultInModel_001`
- `TestInstructionsAndTestInstructionsContainersRespons_PBB002_ExpectedResultInModel_002`
- `TestInstructionsAndTestInstructionsContainersRespons_PBB002_ExpectedResultInModel_003`
- `TestInstructionsAndTestInstructionsContainersRespons_PBB002_ExpectedResultInModel_004`
- `TestInstructionsAndTestInstructionsContainersRespons_PBB003_ExpectedResultInModel_001`
- `TestInstructionsAndTestInstructionsContainersRespons_PBB003_ExpectedResultInModel_002`
- `TestInstructionsAndTestInstructionsContainersRespons_PBB003_ExpectedResultInModel_003`
- `TestInstructionsAndTestInstructionsContainersRespons_PBB003_ExpectedResultInModel_004`
- `highestFenixProtoFileVersion`

## Functions and Methods
### InitLoggerForTest
- Signature: `func InitLoggerForTest(filename string) myTestLogger *logrus.Logger`
- Exported: `true`
- Control-flow features: `if`
- Doc: Init the logger for UnitTests
- Selector calls: `logrus.StandardLogger`, `logrus.SetLevel`, `logrus.SetFormatter`, `os.OpenFile`, `log.Println`

### IsTestDataUsingCorrectTestDataProtoFileVersion
- Signature: `func IsTestDataUsingCorrectTestDataProtoFileVersion(usedProtoFileVersion fenixTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum) returnMessage *fenixTestCaseBuilderServerGrpcApi.AckNackResponse`
- Exported: `true`
- Control-flow features: `if`
- Doc: ******************************************************************************************************************* Check if testdata is using correct proto-file version
- Internal calls: `getHighestFenixTestDataProtoFileVersion`
- Selector calls: `fenixTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum`, `protoFileExpected.String`, `protoFileUsed.String`

### getHighestFenixTestDataProtoFileVersion
- Signature: `func getHighestFenixTestDataProtoFileVersion() int32`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Get the highest FenixProtoFileVersionEnumeration

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
