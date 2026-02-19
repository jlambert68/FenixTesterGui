# testCaseModel_saveTestCase.go

## File Overview
- Path: `testCase/testCaseModel/testCaseModel_saveTestCase.go`
- Package: `testCaseModel`
- Functions/Methods: `11`
- Imports: `11`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `SaveChangedTestCaseAttributeInTestCase`
- `SaveFullTestCase`

## Imports
- `FenixTesterGui/common_code`
- `errors`
- `fmt`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTemplateToThisDomain/version_1_0`
- `github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTestDataToThisDomain/version_1_0`
- `google.golang.org/protobuf/encoding/protojson`
- `google.golang.org/protobuf/types/known/timestamppb`
- `log`
- `os`
- `strings`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### SaveFullTestCase (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) SaveFullTestCase(testCaseUuid string, currentActiveUser string) err error`
- Exported: `true`
- Control-flow features: `if, returns error`
- Doc: SaveFullTestCase - Save the TestCase to the Database
- Selector calls: `testCaseModel.SaveChangedTestCaseAttributeInTestCase`, `errors.New`, `fmt.Sprintf`, `fmt.Println`, `timestamppb.Now`, `testCaseModel.generateTestCaseForGrpcAndHash`, `gRPCTestCaseTestData.GetHashOfThisMessageWithEmptyHashField`, `fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum`

### generateMatureTestInstructionsForGrpc (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) generateMatureTestInstructionsForGrpc(testCaseUuid string) (gRPCMatureTestInstructions []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionsMessage_MatureTestInstructionMessage, hashedSlice string, valuesToBeHashedSlice []string, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: Convert the MatureTestCaseTestInstructions to its gRPC-version
- Selector calls: `errors.New`, `fmt.Sprintf`, `fmt.Println`, `protojson.Format`, `strings.ReplaceAll`, `sharedCode.HashValues`

### generateMatureTestInstructionContainersForGrpc (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) generateMatureTestInstructionContainersForGrpc(testCaseUuid string) (gRPCMatureTestInstructionContainers []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionContainersMessage_MatureTestInstructionContainerMessage, hashedSlice string, valuesToBeHashedSlice []string, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: Convert the MatureTestCaseTestInstructionContainers to its gRPC-version
- Selector calls: `errors.New`, `fmt.Sprintf`, `fmt.Println`, `protojson.Format`, `strings.ReplaceAll`, `sharedCode.HashValues`

### generateTestCaseModelElementsForGrpc (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) generateTestCaseModelElementsForGrpc(testCaseUuid string) (gRPCMatureTestCaseModelElements []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage, hashedSlice string, valuesToBeHashedSlice []string, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: Convert the MatureTestCaseModelElementMessage-map into its gRPC-version
- Selector calls: `errors.New`, `fmt.Sprintf`, `fmt.Println`, `protojson.Format`, `strings.ReplaceAll`, `sharedCode.HashValues`

### generateTestCaseExtraInformationForGrpc (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) generateTestCaseExtraInformationForGrpc(testCaseUuid string) (gRPCTestCaseExtraInformation *fenixGuiTestCaseBuilderServerGrpcApi.TestCaseExtraInformationMessage, hashedSlice string, valuesToBeHashedSlice []string, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: Convert the TestCaseExtraInformationMessage into its gRPC-version Containing: 1) Textual Representation of TestCase
- Selector calls: `errors.New`, `fmt.Sprintf`, `fmt.Println`, `protojson.Format`, `strings.ReplaceAll`, `sharedCode.HashValues`

### generateTestCaseTemplateFilesForGrpc (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) generateTestCaseTemplateFilesForGrpc(testCaseUuid string) (gRPCTestCaseTemplateFiles *fenixGuiTestCaseBuilderServerGrpcApi.TestCaseTemplateFilesMessage, hashedSlice string, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: Convert the TestCaseTemplateFiles into its gRPC-version
- Internal calls: `int64`
- Selector calls: `errors.New`, `fmt.Sprintf`, `fmt.Println`, `protojson.Format`, `strings.ReplaceAll`, `sharedCode.HashValues`

### generateTestCaseTestDataForGrpc (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) generateTestCaseTestDataForGrpc(testCaseUuid string) (gRPCUsersChosenTestDataForTestCase *fenixGuiTestCaseBuilderServerGrpcApi.UsersChosenTestDataForTestCaseMessage, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: Convert the TestCaseTemplateFiles into its gRPC-version
- Internal calls: `string`
- Selector calls: `errors.New`, `fmt.Sprintf`, `fmt.Println`, `protojson.Format`, `strings.ReplaceAll`, `sharedCode.HashSingleValue`

### generateTestCasePreviewMessageForGrpc (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) generateTestCasePreviewMessageForGrpc(testCaseUuid string) (gRPCTestCasePreviewMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestCasePreviewMessage, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, switch, returns error`
- Doc: Convert the TestCasePreviewMessage into its gRPC-version
- Selector calls: `errors.New`, `fmt.Sprintf`, `fmt.Println`, `testCaseStructureObject.GetTestCaseStructureObjectType`, `testCaseStructureObject.GetTestInstructionUuid`, `fenixGuiTestCaseBuilderServerGrpcApi.MetaDataSelectTypeEnum`, `log.Fatal`, `protojson.Format`

### generateUserSpecifiedTestCaseMetaDataMessageForGrpc (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) generateUserSpecifiedTestCaseMetaDataMessageForGrpc(testCaseUuid string, shouldBeSaved bool) (gRPCUserSpecifiedTestCaseMetaDataMessage *fenixGuiTestCaseBuilderServerGrpcApi.UserSpecifiedTestCaseMetaDataMessage, hashedSlice string, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, switch, returns error`
- Doc: Convert the UserSpecifiedTestCaseMetaDataMessage into its gRPC-version
- Selector calls: `errors.New`, `fmt.Sprintf`, `fmt.Println`, `testCaseModel.verifyMandatoryFieldsForMetaData`, `gRPCUserSpecifiedTestCaseMetaDataMessage.GetCurrentSelectedDomainUuid`, `fenixGuiTestCaseBuilderServerGrpcApi.MetaDataSelectTypeEnum`, `log.Fatalln`, `sharedCode.HashValues`

### generateTestCaseForGrpcAndHash (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) generateTestCaseForGrpcAndHash(testCaseUuid string, shouldBeSaved bool) (gRPCMatureTestCaseModelElementMessage []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage, gRPCMatureTestInstructions []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionsMessage_MatureTestInstructionMessage, gRPCMatureTestInstructionContainers []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionContainersMessage_MatureTestInstructionContainerMessage, gRPCTestCaseExtraInformation *fenixGuiTestCaseBuilderServerGrpcApi.TestCaseExtraInformationMessage, gRPCTestCaseTemplateFiles *fenixGuiTestCaseBuilderServerGrpcApi.TestCaseTemplateFilesMessage, gRPCTestCaseTestData *fenixGuiTestCaseBuilderServerGrpcApi.UsersChosenTestDataForTestCaseMessage, gRPCTestCasePreviewMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestCasePreviewMessage, gRPCUserSpecifiedTestCaseMetaDataMessage *fenixGuiTestCaseBuilderServerGrpcApi.UserSpecifiedTestCaseMetaDataMessage, finalHash string, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, defer, returns error`
- Doc: Pack different parts of the TestCase into gRPC-version into one message together with Hash of TestCase
- Selector calls: `errors.New`, `fmt.Sprintf`, `fmt.Println`, `testCaseModel.generateTestCaseModelElementsForGrpc`, `testCaseModel.generateMatureTestInstructionsForGrpc`, `testCaseModel.generateMatureTestInstructionContainersForGrpc`, `testCaseModel.generateTestCaseExtraInformationForGrpc`, `testCaseModel.generateTestCaseTemplateFilesForGrpc`

### SaveChangedTestCaseAttributeInTestCase (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) SaveChangedTestCaseAttributeInTestCase(testCaseUuid string) err error`
- Exported: `true`
- Control-flow features: `if, for/range, switch, returns error`
- Doc: SaveChangedTestCaseAttributeInTestCase - Save changed Attributes into the TestCase-model under correct TestInstruction
- Internal calls: `string`
- Selector calls: `errors.New`, `fmt.Sprintf`, `fmt.Println`, `log.Fatalln`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
