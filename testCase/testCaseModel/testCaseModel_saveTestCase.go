package testCaseModel

import (
	sharedCode "FenixTesterGui/common_code"
	"errors"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	testInstruction_SendTemplateToThisDomain_version_1_0 "github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTemplateToThisDomain/version_1_0"
	testInstruction_SendTestDataToThisDomain_version_1_0 "github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTestDataToThisDomain/version_1_0"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"strings"
)

// SaveFullTestCase - Save the TestCase to the Database
func (testCaseModel *TestCasesModelsStruct) SaveFullTestCase(testCaseUuid string, currentActiveUser string) (err error) {

	// Save changed Attributes, if there are any, into the TestCase-model. Needs to call because last attributes change is not saved into model
	err = testCaseModel.SaveChangedTestCaseAttributeInTestCase(testCaseUuid)
	if err != nil {
		return err
	}

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {

		errorId := "4c075798-ec6c-4486-8053-997ef0d0d8eb"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return err
	}

	//Loop all TestInstructions and Update

	// Create timestamp to be used
	timeStampForTestCaseUpdate := timestamppb.Now()

	// Convert map-messages into gRPC-version, mostly arrays
	var (
		gRPCMatureTestCaseModelElementMessage []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage
		gRPCMatureTestInstructions            []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionsMessage_MatureTestInstructionMessage
		gRPCMatureTestInstructionContainers   []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionContainersMessage_MatureTestInstructionContainerMessage
		gRPCTestCaseExtraInformation          *fenixGuiTestCaseBuilderServerGrpcApi.TestCaseExtraInformationMessage
		gRPCTestCaseTemplateFiles             *fenixGuiTestCaseBuilderServerGrpcApi.TestCaseTemplateFilesMessage
		gRPCTestCaseTestData                  *fenixGuiTestCaseBuilderServerGrpcApi.UsersChosenTestDataForTestCaseMessage
		finalHash                             string
	)
	gRPCMatureTestCaseModelElementMessage,
		gRPCMatureTestInstructions,
		gRPCMatureTestInstructionContainers,
		gRPCTestCaseExtraInformation,
		gRPCTestCaseTemplateFiles,
		gRPCTestCaseTestData,
		finalHash, err = testCaseModel.generateTestCaseForGrpcAndHash(testCaseUuid)
	if err != nil {
		return err
	}

	// Check if changes are done to TestCase, but is only done if the TestCase is not saved before
	if currentTestCase.ThisIsANewTestCase == true ||
		currentTestCase.TestCaseHash != finalHash ||
		currentTestCase.TestDataHash != gRPCTestCaseTestData.GetHashOfThisMessageWithEmptyHashField() {

		currentTestCase.TestCaseHash = finalHash
		currentTestCase.TestDataHash = gRPCTestCaseTestData.GetHashOfThisMessageWithEmptyHashField()

	} else {
		return nil

	}

	// Save full TestCase in DB
	//rpc SaveFullTestCase(FullTestCaseMessage) returns (AckNackResponse)
	gRPCFullTestCaseMessageToSend := fenixGuiTestCaseBuilderServerGrpcApi.FullTestCaseMessage{
		TestCaseBasicInformation: &fenixGuiTestCaseBuilderServerGrpcApi.TestCaseBasicInformationMessage{
			BasicTestCaseInformation: &fenixGuiTestCaseBuilderServerGrpcApi.BasicTestCaseInformationMessage{
				NonEditableInformation: &currentTestCase.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation,
				EditableInformation:    &currentTestCase.LocalTestCaseMessage.BasicTestCaseInformationMessageEditableInformation,
			},
			CreatedAndUpdatedInformation: &fenixGuiTestCaseBuilderServerGrpcApi.TestCaseBasicInformationMessage_CreatedAndUpdatedInformationMessage{
				AddedToTestCaseTimeStamp:       timeStampForTestCaseUpdate,
				AddedToTestCaseByUserId:        currentActiveUser,
				LastUpdatedInTestCaseTimeStamp: timeStampForTestCaseUpdate,
				LastUpdatedInTestCaseByUserId:  currentActiveUser,
				DeletedFromTestCaseTimeStamp: &timestamppb.Timestamp{
					Seconds: 0,
					Nanos:   0,
				},
				DeletedFromTestCaseByUserId: "",
			},
			TestCaseModel: &fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage{
				TestCaseModelAsString:  currentTestCase.TextualTestCaseRepresentationExtendedStack[0],
				FirstMatureElementUuid: currentTestCase.FirstElementUuid,
				TestCaseModelElements:  gRPCMatureTestCaseModelElementMessage,
				TestCaseModelCommands:  []*fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage_TestCaseModelCommandMessage{},
			},
			TestCaseMetaData: &fenixGuiTestCaseBuilderServerGrpcApi.TestCaseMetaDataMessage{},
			TestCaseFiles:    &fenixGuiTestCaseBuilderServerGrpcApi.TestCaseFilesMessage{},
			UserIdentification: &fenixGuiTestCaseBuilderServerGrpcApi.UserIdentificationMessage{
				UserIdOnComputer:     sharedCode.CurrentUserIdLogedInOnComputer,
				GCPAuthenticatedUser: sharedCode.CurrentUserAuthenticatedTowardsGCP,
				ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(
					testCaseModel.GrpcOutReference.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion()),
			},
		},
		MatureTestInstructions: &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionsMessage{
			MatureTestInstructions: gRPCMatureTestInstructions,
		},
		MatureTestInstructionContainers: &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionContainersMessage{
			MatureTestInstructionContainers: gRPCMatureTestInstructionContainers},
		MessageHash:              currentTestCase.TestCaseHash,
		TestCaseExtraInformation: gRPCTestCaseExtraInformation,
		TestCaseTemplateFiles:    gRPCTestCaseTemplateFiles,
		TestCaseTestData:         gRPCTestCaseTestData,
	}

	// Send using gRPC
	returnMessage := testCaseModel.GrpcOutReference.SendSaveFullTestCase(&gRPCFullTestCaseMessageToSend)

	if returnMessage == nil || returnMessage.AckNack == false {

		errorId := "cb68859b-5c99-48a5-8f8b-9af472a9a45a"
		err = errors.New(fmt.Sprintf(returnMessage.Comments+"[ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return err
	}

	// Update that the TestCase is not New anymore
	currentTestCase.ThisIsANewTestCase = false

	// Update The Hash for the TestCase
	currentTestCase.TestCaseHashWhenTestCaseWasSavedOrLoaded = gRPCFullTestCaseMessageToSend.MessageHash
	currentTestCase.TestDataHashWhenTestCaseWasSavedOrLoaded = gRPCFullTestCaseMessageToSend.GetTestCaseTestData().GetHashOfThisMessageWithEmptyHashField()

	// Save the TestCase back in Map
	testCaseModel.TestCases[testCaseUuid] = currentTestCase

	return err

}

// Convert the MatureTestCaseTestInstructions to its gRPC-version
func (testCaseModel *TestCasesModelsStruct) generateMatureTestInstructionsForGrpc(
	testCaseUuid string) (
	gRPCMatureTestInstructions []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionsMessage_MatureTestInstructionMessage,
	hashedSlice string,
	valuesToBeHashedSlice []string,
	err error) {

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {

		errorId := "48899cab-ce9d-48a2-947f-d7610a3bea81"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return nil, "", nil, err
	}

	// Loop map with all 'MatureTestInstructions' in the TestCase and create a slice
	for _, matureTestInstruction := range currentTestCase.MatureTestInstructionMap {

		var tempMatureTestInstruction MatureTestInstructionStruct
		tempMatureTestInstruction = matureTestInstruction

		// Generate Hashes for  'matureTestInstruction'
		tempJson := protojson.Format(tempMatureTestInstruction.BasicTestInstructionInformation_NonEditableInformation)
		valuesToBeHashedSlice = append(valuesToBeHashedSlice, tempJson)
		tempJson = protojson.Format(tempMatureTestInstruction.BasicTestInstructionInformation_EditableInformation)
		valuesToBeHashedSlice = append(valuesToBeHashedSlice, tempJson)
		tempJson = protojson.Format(tempMatureTestInstruction.BasicTestInstructionInformation_InvisibleBasicInformation)
		valuesToBeHashedSlice = append(valuesToBeHashedSlice, tempJson)
		tempJson = protojson.Format(tempMatureTestInstruction.MatureBasicTestInstructionInformation)
		valuesToBeHashedSlice = append(valuesToBeHashedSlice, tempJson)

		// Loop over all 'TestInstruction Attributes' in the TestInstruction and create slice
		var testInstructionAttributesList []*fenixGuiTestCaseBuilderServerGrpcApi.
			MatureTestInstructionInformationMessage_TestInstructionAttributeMessage
		for _, tempTestInstructionAttribute := range tempMatureTestInstruction.TestInstructionAttributesList {

			// Get all values from general map of Attributes on the TestCase

			testInstructionAttributesList = append(testInstructionAttributesList, tempTestInstructionAttribute)

			// Generate Hash for  'tempTestInstructionAttribute'
			tempJson := protojson.Format(tempTestInstructionAttribute)
			valuesToBeHashedSlice = append(valuesToBeHashedSlice, tempJson)

		}

		// Create one 'MatureTestInstructionMessage'
		MatureTestInstructionMessage := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionsMessage_MatureTestInstructionMessage{
			BasicTestInstructionInformation: &fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionInformationMessage{
				NonEditableInformation:    tempMatureTestInstruction.BasicTestInstructionInformation_NonEditableInformation,
				EditableInformation:       tempMatureTestInstruction.BasicTestInstructionInformation_EditableInformation,
				InvisibleBasicInformation: tempMatureTestInstruction.BasicTestInstructionInformation_InvisibleBasicInformation,
			},
			MatureTestInstructionInformation: &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage{
				MatureBasicTestInstructionInformation: tempMatureTestInstruction.MatureBasicTestInstructionInformation,
				CreatedAndUpdatedInformation:          tempMatureTestInstruction.CreatedAndUpdatedInformation,
				TestInstructionAttributesList:         testInstructionAttributesList,
			},
		}

		// Add 'MatureTestInstructionMessage' to slice
		gRPCMatureTestInstructions = append(gRPCMatureTestInstructions, &MatureTestInstructionMessage)
	}

	// Remove spaces before hashing, due to some bug that generates "double space" sometimes when running in non-debug-mode
	for index, textToReplaceIn := range valuesToBeHashedSlice {
		valuesToBeHashedSlice[index] = strings.ReplaceAll(textToReplaceIn, " ", "")
	}

	// Generate Hash of all sub-message-hashes
	hashedSlice = sharedCode.HashValues(valuesToBeHashedSlice, false)

	return gRPCMatureTestInstructions, hashedSlice, valuesToBeHashedSlice, err

}

// Convert the MatureTestCaseTestInstructionContainers to its gRPC-version
func (testCaseModel *TestCasesModelsStruct) generateMatureTestInstructionContainersForGrpc(
	testCaseUuid string) (
	gRPCMatureTestInstructionContainers []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionContainersMessage_MatureTestInstructionContainerMessage,
	hashedSlice string,
	valuesToBeHashedSlice []string,
	err error) {

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {

		errorId := "82040ba6-57c4-47e2-8fb5-c770db41d8f8"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return nil, "", nil, err
	}

	// Loop map with all 'MatureTestInstructionContainers' in the TestCase and create a slice
	for _, matureTestInstructionContainer := range currentTestCase.MatureTestInstructionContainerMap {
		var tempMatureTestInstructionContainer MatureTestInstructionContainerStruct
		tempMatureTestInstructionContainer = matureTestInstructionContainer

		// Generate Hashes for  'matureTestInstruction'
		tempJson := protojson.Format(tempMatureTestInstructionContainer.NonEditableInformation)
		valuesToBeHashedSlice = append(valuesToBeHashedSlice, tempJson)
		tempJson = protojson.Format(tempMatureTestInstructionContainer.EditableInformation)
		valuesToBeHashedSlice = append(valuesToBeHashedSlice, tempJson)
		tempJson = protojson.Format(tempMatureTestInstructionContainer.InvisibleBasicInformation)
		valuesToBeHashedSlice = append(valuesToBeHashedSlice, tempJson)
		tempJson = protojson.Format(tempMatureTestInstructionContainer.EditableTestInstructionContainerAttributes)
		valuesToBeHashedSlice = append(valuesToBeHashedSlice, tempJson)
		tempJson = protojson.Format(tempMatureTestInstructionContainer.MatureTestInstructionContainerInformation)
		valuesToBeHashedSlice = append(valuesToBeHashedSlice, tempJson)
		//tempJson = protojson.Format(tempMatureTestInstructionContainer.CreatedAndUpdatedInformation)
		//valuesToBeHashedSlice = append(valuesToBeHashedSlice, tempJson)

		// Create one 'MatureTestInstructionContainerMessage'
		MatureTestInstructionContainerMessage := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionContainersMessage_MatureTestInstructionContainerMessage{
			BasicTestInstructionContainerInformation: &fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionContainerInformationMessage{
				NonEditableInformation:                     tempMatureTestInstructionContainer.NonEditableInformation,
				EditableInformation:                        tempMatureTestInstructionContainer.EditableInformation,
				InvisibleBasicInformation:                  tempMatureTestInstructionContainer.InvisibleBasicInformation,
				EditableTestInstructionContainerAttributes: tempMatureTestInstructionContainer.EditableTestInstructionContainerAttributes,
			},
			MatureTestInstructionContainerInformation: &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionContainerInformationMessage{
				MatureTestInstructionContainerInformation: tempMatureTestInstructionContainer.MatureTestInstructionContainerInformation,
				CreatedAndUpdatedInformation:              tempMatureTestInstructionContainer.CreatedAndUpdatedInformation,
			},
		}

		//TODO change the row below to have the orignal date, but it need to be like this otherwise there is an error when sending over gRPC
		//MatureTestInstructionContainerMessage.BasicTestInstructionContainerInformation.NonEditableInformation.UpdatedTimeStamp = &timestamppb.Timestamp{
		//	Seconds: 0,
		//	Nanos:   0,
		//}

		// Add 'MatureTestInstructionContainerMessage' to slice
		gRPCMatureTestInstructionContainers = append(gRPCMatureTestInstructionContainers, &MatureTestInstructionContainerMessage)
	}

	// Remove spaces before hashing, due to some bug that generates "double space" sometimes when running in non-debug-mode
	for index, textToReplaceIn := range valuesToBeHashedSlice {
		valuesToBeHashedSlice[index] = strings.ReplaceAll(textToReplaceIn, " ", "")
	}

	// Generate Hash of all sub-message-hashes
	hashedSlice = sharedCode.HashValues(valuesToBeHashedSlice, false)

	return gRPCMatureTestInstructionContainers, hashedSlice, valuesToBeHashedSlice, err

}

// Convert the MatureTestCaseModelElementMessage-map into its gRPC-version
func (testCaseModel *TestCasesModelsStruct) generateTestCaseModelElementsForGrpc(
	testCaseUuid string) (
	gRPCMatureTestCaseModelElements []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage,
	hashedSlice string,
	valuesToBeHashedSlice []string,
	err error) {

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {

		errorId := "48899cab-ce9d-48a2-947f-d7610a3bea81"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return nil, "", nil, err
	}

	// Loop map with all 'MatureTestCaseModelElementMessage' in the TestCase and create a slice
	for _, matureTestCaseModelElement := range currentTestCase.TestCaseModelMap {
		var tempMatureTestCaseModelElement MatureTestCaseModelElementStruct
		tempMatureTestCaseModelElement = matureTestCaseModelElement
		gRPCMatureTestCaseModelElements = append(gRPCMatureTestCaseModelElements, &tempMatureTestCaseModelElement.MatureTestCaseModelElementMessage)

		// Generate Hash for  'matureTestCaseModelElement'
		tempJson := protojson.Format(&matureTestCaseModelElement.MatureTestCaseModelElementMessage)
		valuesToBeHashedSlice = append(valuesToBeHashedSlice, tempJson)

	}

	// Remove spaces before hashing, due to some bug that generates "double space" sometimes when running in non-debug-mode
	for index, textToReplaceIn := range valuesToBeHashedSlice {
		valuesToBeHashedSlice[index] = strings.ReplaceAll(textToReplaceIn, " ", "")
	}

	// Generate Hash of all sub-message-hashes
	hashedSlice = sharedCode.HashValues(valuesToBeHashedSlice, false)

	return gRPCMatureTestCaseModelElements, hashedSlice, valuesToBeHashedSlice, err
}

// Convert the TestCaseExtraInformationMessage into its gRPC-version
// Containing: 1) Textual Representation of TestCase
func (testCaseModel *TestCasesModelsStruct) generateTestCaseExtraInformationForGrpc(
	testCaseUuid string) (
	gRPCTestCaseExtraInformation *fenixGuiTestCaseBuilderServerGrpcApi.TestCaseExtraInformationMessage,
	hashedSlice string,
	valuesToBeHashedSlice []string,
	err error) {

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {

		errorId := "e6fbdfdc-e0dc-4dd8-8ab1-b6be82b9e9fe"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return nil, "", nil, err
	}

	// Secure that the number of Textual Models are the same
	var (
		numberSimpleModels   int
		numberComplexModels  int
		numberExtendedModels int
	)
	numberSimpleModels = len(currentTestCase.TextualTestCaseRepresentationSimpleStack)
	numberComplexModels = len(currentTestCase.TextualTestCaseRepresentationComplexStack)
	numberExtendedModels = len(currentTestCase.TextualTestCaseRepresentationExtendedStack)

	if numberSimpleModels != numberComplexModels && numberComplexModels != numberExtendedModels {

		errorId := "eb939008-eb50-40c6-91c5-64e4d1f597a1"
		err = errors.New(fmt.Sprintf("for testcase '%s', the number of Simple, Complex and Extended  models doesn't match. "+
			"'numberSimpleModels': '%d', 'numberComplexModels': '%d', 'numberExtendedModels': '%d', [ErrorID: %s]",
			testCaseUuid, numberSimpleModels, numberComplexModels, numberExtendedModels, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return nil, "", nil, err
	}

	var tempTestCaseTextualRepresentationHistory fenixGuiTestCaseBuilderServerGrpcApi.TestCaseExtraInformationMessage_TestCaseTextualRepresentationHistoryMessage
	tempTestCaseTextualRepresentationHistory = fenixGuiTestCaseBuilderServerGrpcApi.TestCaseExtraInformationMessage_TestCaseTextualRepresentationHistoryMessage{
		TextualTestCaseRepresentationSimpleHistory:        nil,
		TextualTestCaseRepresentationComplexHistory:       nil,
		TextualTestCaseRepresentationExtendedStackHistory: nil,
	}

	// Loop map with all 'tempTestCaseTextualRepresentationHistory' in the TestCase and add to gPRC-version
	for modelCounter := 0; modelCounter < numberSimpleModels; modelCounter++ {

		// Simple
		tempTestCaseTextualRepresentationHistory.TextualTestCaseRepresentationSimpleHistory =
			append(tempTestCaseTextualRepresentationHistory.TextualTestCaseRepresentationSimpleHistory,
				currentTestCase.TextualTestCaseRepresentationSimpleStack[modelCounter])

		// Complex
		tempTestCaseTextualRepresentationHistory.TextualTestCaseRepresentationComplexHistory =
			append(tempTestCaseTextualRepresentationHistory.TextualTestCaseRepresentationComplexHistory,
				currentTestCase.TextualTestCaseRepresentationComplexStack[modelCounter])

		// Extended
		tempTestCaseTextualRepresentationHistory.TextualTestCaseRepresentationExtendedStackHistory =
			append(tempTestCaseTextualRepresentationHistory.TextualTestCaseRepresentationExtendedStackHistory,
				currentTestCase.TextualTestCaseRepresentationExtendedStack[modelCounter])

	}

	// Generate Hash for  'tempTestCaseTextualRepresentationHistory'
	tempJson := protojson.Format(&tempTestCaseTextualRepresentationHistory)
	valuesToBeHashedSlice = append(valuesToBeHashedSlice, tempJson)

	// Remove spaces before hashing, due to some bug that generates "double space" sometimes when running in non-debug-mode
	for index, textToReplaceIn := range valuesToBeHashedSlice {
		valuesToBeHashedSlice[index] = strings.ReplaceAll(textToReplaceIn, " ", "")
	}

	// Generate Hash of all sub-message-hashes
	hashedSlice = sharedCode.HashValues(valuesToBeHashedSlice, false)

	// Create return message 'gRPCTestCaseExtraInformation'
	gRPCTestCaseExtraInformation = &fenixGuiTestCaseBuilderServerGrpcApi.TestCaseExtraInformationMessage{
		TestCaseTextualRepresentationHistory: &tempTestCaseTextualRepresentationHistory,
	}

	return gRPCTestCaseExtraInformation, hashedSlice, valuesToBeHashedSlice, err

}

// Convert the TestCaseTemplateFiles into its gRPC-version
func (testCaseModel *TestCasesModelsStruct) generateTestCaseTemplateFilesForGrpc(
	testCaseUuid string) (
	gRPCTestCaseTemplateFiles *fenixGuiTestCaseBuilderServerGrpcApi.TestCaseTemplateFilesMessage,
	hashedSlice string,
	err error) {

	var valuesToBeHashedSlice []string
	gRPCTestCaseTemplateFiles = &fenixGuiTestCaseBuilderServerGrpcApi.TestCaseTemplateFilesMessage{
		TestCaseTemplateFile: nil,
		HashForAllFiles:      "",
	}

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {

		errorId := "e6ceecbe-00e2-42af-9782-eb83af2d03c2"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map witsh all TestCases [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return nil, "", err
	}

	// Loop map with all 'TestCaseTemplateFiles' in the TestCase and add to gPRC-version
	for _, importedTemplateFileFromGitHub := range currentTestCase.ImportedTemplateFilesFromGitHub {

		// Create the gRPC-version of the 'ImportedTemplateFileFromGitHub'
		var tempTestCaseTemplateFileMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestCaseTemplateFileMessage
		tempTestCaseTemplateFileMessage = &fenixGuiTestCaseBuilderServerGrpcApi.TestCaseTemplateFileMessage{
			Name:                importedTemplateFileFromGitHub.Name,
			URL:                 importedTemplateFileFromGitHub.URL,
			DownloadURL:         importedTemplateFileFromGitHub.DownloadURL,
			SHA:                 importedTemplateFileFromGitHub.SHA,
			Size:                int64(importedTemplateFileFromGitHub.Size),
			FileContentAsString: importedTemplateFileFromGitHub.FileContentAsString,
			FileHash:            importedTemplateFileFromGitHub.FileHash,
		}

		// Generate Hash for  'importedTemplateFileFromGitHub'
		tempJson := protojson.Format(tempTestCaseTemplateFileMessage)
		valuesToBeHashedSlice = append(valuesToBeHashedSlice, tempJson)

		// Add to Slice of all gRPC-versions of all 'ImportedTemplateFileFromGitHub'
		gRPCTestCaseTemplateFiles.TestCaseTemplateFile = append(gRPCTestCaseTemplateFiles.TestCaseTemplateFile, tempTestCaseTemplateFileMessage)

	}

	// Remove spaces before hashing, due to some bug that generates "double space" sometimes when running in non-debug-mode
	for index, textToReplaceIn := range valuesToBeHashedSlice {
		valuesToBeHashedSlice[index] = strings.ReplaceAll(textToReplaceIn, " ", "")
	}

	// Generate Hash of all sub-message-hashes
	hashedSlice = sharedCode.HashValues(valuesToBeHashedSlice, true)

	// Add hash to gRPC-versions of 'TestCaseTemplateFiles'
	gRPCTestCaseTemplateFiles.HashForAllFiles = hashedSlice

	return gRPCTestCaseTemplateFiles, hashedSlice, err

}

// Convert the TestCaseTemplateFiles into its gRPC-version
func (testCaseModel *TestCasesModelsStruct) generateTestCaseTestDataForGrpc(
	testCaseUuid string) (
	gRPCUsersChosenTestDataForTestCase *fenixGuiTestCaseBuilderServerGrpcApi.UsersChosenTestDataForTestCaseMessage,
	err error) {

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {

		errorId := "f6354f22-5bca-41f8-8b92-b168b4381708"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map witsh all TestCases [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return nil, err
	}

	// The gRPC-version of 'testDataPointNameMap'
	var chosenTestDataPointsPerGroupMapGrpc map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
		TestDataPointNameMapMessage
	chosenTestDataPointsPerGroupMapGrpc = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
		TestDataPointNameMapMessage)

	if currentTestCase.TestData != nil {

		// Loop TestDataGroups
		for testDataGroupName, testDataPointNameMap := range currentTestCase.TestData.ChosenTestDataPointsPerGroupMap {

			var testDataPointNameMapMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestDataPointNameMapMessage
			testDataPointNameMapMessage = &fenixGuiTestCaseBuilderServerGrpcApi.TestDataPointNameMapMessage{}

			// The gRPC-version of 'testDataPointNameMap'
			var chosenTestDataRowsPerTestDataPointMapGprc map[string]*fenixGuiTestCaseBuilderServerGrpcApi.TestDataRowsMessage
			chosenTestDataRowsPerTestDataPointMapGprc = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.TestDataRowsMessage)

			// Extract TestDataPoints for the TestDataPointNameMap
			for testDataPointName, testDataPointsSlice := range *testDataPointNameMap {

				var testDataRowsGrpc *fenixGuiTestCaseBuilderServerGrpcApi.TestDataRowsMessage
				testDataRowsGrpc = &fenixGuiTestCaseBuilderServerGrpcApi.TestDataRowsMessage{}

				// The gRPC-version of 'testDataPointsSlice'
				var testDataRowMessageSliceGrc []*fenixGuiTestCaseBuilderServerGrpcApi.TestDataRowMessage

				// Loop the TestDataPoints
				for _, testDataPoint := range *testDataPointsSlice {

					// gRPC-version of 'SelectedTestDataPointUuidMap'
					var testDataPointRowValueSummaryMapGrpc map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
						TestDataPointRowValueSummaryMapMessage
					testDataPointRowValueSummaryMapGrpc = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
						TestDataPointRowValueSummaryMapMessage)

					// Loop the 'selected' TestDataRows in the TestDataPoint
					for _, testDataPointRowData := range testDataPoint.SelectedTestDataPointUuidMap {

						// Create gRPC-version of 'testDataPointRowData'
						var testDataPointRowDataGrpc *fenixGuiTestCaseBuilderServerGrpcApi.
							TestDataPointRowValueSummaryMapMessage
						testDataPointRowDataGrpc = &fenixGuiTestCaseBuilderServerGrpcApi.
							TestDataPointRowValueSummaryMapMessage{
							TestDataPointRowUuid:          string(testDataPointRowData.TestDataPointRowUuid),
							TestDataPointRowValuesSummary: string(testDataPointRowData.TestDataPointRowValuesSummary),
						}

						// Add 'testDataPointRowDataGrpc' to gRPC-version of 'SelectedTestDataPointUuidMap'
						testDataPointRowValueSummaryMapGrpc[string(testDataPointRowData.TestDataPointRowUuid)] =
							testDataPointRowDataGrpc
					}

					// Create the gRPC-version of 'testDataPoint'
					var testDataRowMessageGrpc *fenixGuiTestCaseBuilderServerGrpcApi.TestDataRowMessage
					testDataRowMessageGrpc = &fenixGuiTestCaseBuilderServerGrpcApi.TestDataRowMessage{
						TestDataDomainUuid:              string(testDataPoint.TestDataDomainUuid),
						TestDataDomainName:              string(testDataPoint.TestDataDomainName),
						TestDataAreaUuid:                string(testDataPoint.TestDataAreaUuid),
						TestDataAreaName:                string(testDataPoint.TestDataAreaName),
						TestDataPointName:               string(testDataPoint.TestDataPointName),
						TestDataPointRowValueSummaryMap: testDataPointRowValueSummaryMapGrpc,
					}

					// Append to slice of TestDataPoint-message
					testDataRowMessageSliceGrc = append(testDataRowMessageSliceGrc, testDataRowMessageGrpc)
				}

				// Add 'testDataRowMessageSliceGrc' into "single message"
				testDataRowsGrpc.TestDataRows = testDataRowMessageSliceGrc

				// Store "single-message" with TestDataRows in map 'chosenTestDataRowsPerTestDataPointMapGprc'
				chosenTestDataRowsPerTestDataPointMapGprc[string(testDataPointName)] = testDataRowsGrpc

			}

			// And 'chosenTestDataRowsPerTestDataPointMapGprc' into "single message"
			testDataPointNameMapMessage.ChosenTestDataRowsPerTestDataPointMap = chosenTestDataRowsPerTestDataPointMapGprc

			// Store "singe-message with Group-data in map 'chosenTestDataPointsPerGroupMapGrpc'
			chosenTestDataPointsPerGroupMapGrpc[string(testDataGroupName)] = testDataPointNameMapMessage

		}
	}

	gRPCUsersChosenTestDataForTestCase = &fenixGuiTestCaseBuilderServerGrpcApi.UsersChosenTestDataForTestCaseMessage{
		ChosenTestDataPointsPerGroupMap:     chosenTestDataPointsPerGroupMapGrpc,
		HashOfThisMessageWithEmptyHashField: ""}

	// Generate Hash of gRPC-message and add it to the message
	tempJson := protojson.Format(gRPCUsersChosenTestDataForTestCase)

	// Remove spaces before hashing, due to some bug that generates "double space" sometimes when running in non-debug-mode
	tempJson = strings.ReplaceAll(tempJson, " ", "")

	hashedJson := sharedCode.HashSingleValue(tempJson)
	gRPCUsersChosenTestDataForTestCase.HashOfThisMessageWithEmptyHashField = hashedJson

	return gRPCUsersChosenTestDataForTestCase, err

}

// Pack different parts of the TestCase into gRPC-version into one message together with Hash of TestCase
func (testCaseModel *TestCasesModelsStruct) generateTestCaseForGrpcAndHash(testCaseUuid string) (
	gRPCMatureTestCaseModelElementMessage []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage,
	gRPCMatureTestInstructions []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionsMessage_MatureTestInstructionMessage,
	gRPCMatureTestInstructionContainers []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionContainersMessage_MatureTestInstructionContainerMessage,
	gRPCTestCaseExtraInformation *fenixGuiTestCaseBuilderServerGrpcApi.TestCaseExtraInformationMessage,
	gRPCTestCaseTemplateFiles *fenixGuiTestCaseBuilderServerGrpcApi.TestCaseTemplateFilesMessage,
	gRPCTestCaseTestData *fenixGuiTestCaseBuilderServerGrpcApi.UsersChosenTestDataForTestCaseMessage,
	finalHash string,
	err error) {

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {

		errorId := "4c075798-ec6c-4486-8053-997ef0d0d8eb"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return nil, nil, nil, nil, nil, nil, "", err
	}

	// Initiate 'subHashPartsSlice', used for holding all Hashes and content, to be logged. Used for debugging when
	// there is a mismatch in Saved Hash and Loaded Hash for TestCase
	type subHashPartsMapValueType struct {
		nameOfContent        string
		hash                 string
		contentAsStringSlice []string
	}
	var subHashPartsSlice []subHashPartsMapValueType

	// Convert map-messages into gRPC-version, mostly arrays
	// TestCase-model
	var hashedMatureTestCaseModelElements string
	var valuesToBeHashedSlice []string
	gRPCMatureTestCaseModelElementMessage, hashedMatureTestCaseModelElements, valuesToBeHashedSlice, err =
		testCaseModel.generateTestCaseModelElementsForGrpc(testCaseUuid)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, "", err
	}

	// Add hash and values to slice
	var tempGrpcMatureTestCaseModelElementMessageHashData subHashPartsMapValueType
	tempGrpcMatureTestCaseModelElementMessageHashData = subHashPartsMapValueType{
		nameOfContent:        "gRPCMatureTestCaseModelElementMessage",
		hash:                 hashedMatureTestCaseModelElements,
		contentAsStringSlice: valuesToBeHashedSlice,
	}
	subHashPartsSlice = append(subHashPartsSlice, tempGrpcMatureTestCaseModelElementMessageHashData)

	// TestInstructions
	var hashedgRPCMatureTestInstructions string
	gRPCMatureTestInstructions, hashedgRPCMatureTestInstructions, valuesToBeHashedSlice, err =
		testCaseModel.generateMatureTestInstructionsForGrpc(testCaseUuid)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, "", err
	}
	// Add hash and values to slice
	var tempGrpcMatureTestInstructions subHashPartsMapValueType
	tempGrpcMatureTestInstructions = subHashPartsMapValueType{
		nameOfContent:        "gRPCMatureTestInstructions",
		hash:                 hashedgRPCMatureTestInstructions,
		contentAsStringSlice: valuesToBeHashedSlice,
	}
	subHashPartsSlice = append(subHashPartsSlice, tempGrpcMatureTestInstructions)

	// TestInstructionContainers
	var hashedgRPCMatureTestInstructionContainers string
	gRPCMatureTestInstructionContainers, hashedgRPCMatureTestInstructionContainers, valuesToBeHashedSlice, err =
		testCaseModel.generateMatureTestInstructionContainersForGrpc(testCaseUuid)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, "", err
	}
	// Add hash and values to slice
	var tempGrpcMatureTestInstructionContainers subHashPartsMapValueType
	tempGrpcMatureTestInstructionContainers = subHashPartsMapValueType{
		nameOfContent:        "gRPCMatureTestInstructionContainers",
		hash:                 hashedgRPCMatureTestInstructionContainers,
		contentAsStringSlice: valuesToBeHashedSlice,
	}
	subHashPartsSlice = append(subHashPartsSlice, tempGrpcMatureTestInstructionContainers)

	// TestCaseExtraInformation
	var hashedgRPCTestCaseExtraInformation string
	gRPCTestCaseExtraInformation, hashedgRPCTestCaseExtraInformation, valuesToBeHashedSlice, err =
		testCaseModel.generateTestCaseExtraInformationForGrpc(testCaseUuid)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, "", err
	}
	// Add hash and values to slice
	var tempGrpcTestCaseExtraInformation subHashPartsMapValueType
	tempGrpcTestCaseExtraInformation = subHashPartsMapValueType{
		nameOfContent:        "gRPCTestCaseExtraInformation",
		hash:                 hashedgRPCTestCaseExtraInformation,
		contentAsStringSlice: valuesToBeHashedSlice,
	}
	subHashPartsSlice = append(subHashPartsSlice, tempGrpcTestCaseExtraInformation)

	// TestCaseTemplateFiles
	var hashedgRPCTestCaseTemplateFiles string
	gRPCTestCaseTemplateFiles, hashedgRPCTestCaseTemplateFiles, err =
		testCaseModel.generateTestCaseTemplateFilesForGrpc(testCaseUuid)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, "", err
	}
	// Add hash and values to slice
	var tempGrpcTestCaseTemplateFiles subHashPartsMapValueType
	tempGrpcTestCaseTemplateFiles = subHashPartsMapValueType{
		nameOfContent:        "gRPCTestCaseTemplateFiles",
		hash:                 hashedgRPCTestCaseTemplateFiles,
		contentAsStringSlice: []string{},
	}
	subHashPartsSlice = append(subHashPartsSlice, tempGrpcTestCaseTemplateFiles)

	// TestCaseTestData
	gRPCTestCaseTestData, err =
		testCaseModel.generateTestCaseTestDataForGrpc(testCaseUuid)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, "", err
	}

	var valuesToReHash []string

	// NonEditableInformation, start by clearing Version because it s not the same after save to Database
	currentTestCase.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.TestCaseVersion = 0
	tempNonEditableInformation := fmt.Sprint(&currentTestCase.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation)
	hashNonEditableInformation := sharedCode.HashSingleValue(tempNonEditableInformation)
	valuesToReHash = append(valuesToReHash, hashNonEditableInformation)
	// Add hash and values to slice
	var tempNonEditableInformationValue subHashPartsMapValueType
	tempNonEditableInformationValue = subHashPartsMapValueType{
		nameOfContent:        "tempNonEditableInformation",
		hash:                 hashNonEditableInformation,
		contentAsStringSlice: []string{tempNonEditableInformation},
	}
	subHashPartsSlice = append(subHashPartsSlice, tempNonEditableInformationValue)

	// EditableInformation
	tempEditableInformation := fmt.Sprint(&currentTestCase.LocalTestCaseMessage.BasicTestCaseInformationMessageEditableInformation)
	hashEditableInformation := sharedCode.HashSingleValue(tempEditableInformation)
	valuesToReHash = append(valuesToReHash, hashEditableInformation)
	// Add hash and values to slice
	var tempEditableInformationValue subHashPartsMapValueType
	tempEditableInformationValue = subHashPartsMapValueType{
		nameOfContent:        "tempEditableInformation",
		hash:                 hashEditableInformation,
		contentAsStringSlice: []string{tempEditableInformation},
	}
	subHashPartsSlice = append(subHashPartsSlice, tempEditableInformationValue)

	// FirstMatureElementUuid
	tempFirstMatureElementUuid := fmt.Sprint(currentTestCase.FirstElementUuid)
	hashFirstMatureElementUuid := sharedCode.HashSingleValue(tempFirstMatureElementUuid)
	valuesToReHash = append(valuesToReHash, hashFirstMatureElementUuid)
	// Add hash and values to slice
	var tempFirstMatureElementUuidValue subHashPartsMapValueType
	tempFirstMatureElementUuidValue = subHashPartsMapValueType{
		nameOfContent:        "tempFirstMatureElementUuid",
		hash:                 hashFirstMatureElementUuid,
		contentAsStringSlice: []string{tempFirstMatureElementUuid},
	}
	subHashPartsSlice = append(subHashPartsSlice, tempFirstMatureElementUuidValue)

	valuesToReHash = append(valuesToReHash, hashedMatureTestCaseModelElements)

	valuesToReHash = append(valuesToReHash, hashedgRPCMatureTestInstructions)

	valuesToReHash = append(valuesToReHash, hashedgRPCMatureTestInstructionContainers)

	valuesToReHash = append(valuesToReHash, hashedgRPCTestCaseExtraInformation)

	valuesToReHash = append(valuesToReHash, hashedgRPCTestCaseTemplateFiles)

	finalHash = sharedCode.HashValues(valuesToReHash, false)
	// Add hash and values to slice
	var tempFinalHash subHashPartsMapValueType
	tempFinalHash = subHashPartsMapValueType{
		nameOfContent:        "finalHash",
		hash:                 finalHash,
		contentAsStringSlice: []string{finalHash},
	}
	subHashPartsSlice = append(subHashPartsSlice, tempFinalHash)

	// Log subHashPartsSlice to be used for bug findings when TestCase model is out of sync
	sharedCode.Logger.WithFields(logrus.Fields{
		"ID":                "d0034233-202b-469c-93e3-48903991fe23",
		"testCaseUuid":      testCaseUuid,
		"subHashPartsSlice": subHashPartsSlice,
	}).Debug("Used for bug findings when TestCase model gets of sync when saving TestCase and the Loading the same TestCase")

	return gRPCMatureTestCaseModelElementMessage,
		gRPCMatureTestInstructions,
		gRPCMatureTestInstructionContainers,
		gRPCTestCaseExtraInformation,
		gRPCTestCaseTemplateFiles,
		gRPCTestCaseTestData,
		finalHash,
		err

}

// SaveChangedTestCaseAttributeInTestCase - Save changed Attributes into the TestCase-model under correct TestInstruction
func (testCaseModel *TestCasesModelsStruct) SaveChangedTestCaseAttributeInTestCase(testCaseUuid string) (err error) {

	var atleastOneAttributeIsChanged bool

	// Extract current TestCase
	testCase, existInMap := testCaseModel.TestCases[testCaseUuid]
	if existInMap == false {

		errorId := "40fc730f-87d4-4c44-96ff-ab1003e40751"
		err := errors.New(fmt.Sprintf("testCase %s is missing in TestCases-map [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) //TODO Send error over error-channel
		return err
	}

	// If there are no attributes in  'AttributesList' for TestCase (No unsaved changes exists) then exist
	if testCase.AttributesList == nil || len(*testCase.AttributesList) == 0 {
		return err
	}

	// Extract testInstructionElementMatureUuidUuid
	attributesList := *testCase.AttributesList

	// If Nothing in attributes list then just exit
	if len(attributesList) == 0 {
		return err
	}

	testInstructionElementMatureUuidUuid := attributesList[0].TestInstructionElementMatureUuidUuid

	// Check if any attribute is changed
	if len(attributesList) > 0 {
		for _, attribute := range attributesList {
			if attribute.AttributeIsChanged == true {
				// Attribute is changed so save it,

				atleastOneAttributeIsChanged = true

				// Extract TestInstruction
				tempMatureTestInstruction, existInMap := testCase.MatureTestInstructionMap[testInstructionElementMatureUuidUuid]
				if existInMap == false {
					errorId := "83b64181-3a02-4b98-8eba-d1fbad61dcd5"
					err := errors.New(fmt.Sprintf("mature testInstruction %s is missing in MatureTestInstructionMap [ErrorID: %s]", testInstructionElementMatureUuidUuid, errorId))

					fmt.Println(err) //TODO Send error over error-channel
					return err
				}

				// Extract  Attribute
				tempTestInstructionAttribute, existInMap := tempMatureTestInstruction.TestInstructionAttributesList[attribute.AttributeUuid]
				if existInMap == false {
					errorId := "77e03442-7ccc-46c7-891e-0c5e0dd5bd1c"
					err := errors.New(fmt.Sprintf("testInstruction attribute %s is missing in MatureTestInstructionMap [ErrorID: %s]", attribute.AttributeUuid, errorId))

					fmt.Println(err) //TODO Send error over error-channel
					return err
				}

				// Save changed value for Attribute
				switch attribute.AttributeType {

				case fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionAttributeTypeEnum_TEXTBOX:
					tempTestInstructionAttribute.AttributeInformation.InputTextBoxProperty.
						TextBoxAttributeValue = attribute.AttributeChangedValue

					// Save back Attribute into TestInstruction
					tempMatureTestInstruction.TestInstructionAttributesList[attribute.AttributeUuid] = tempTestInstructionAttribute

					// Save back TestInstruction into TestCase
					testCase.MatureTestInstructionMap[testInstructionElementMatureUuidUuid] = tempMatureTestInstruction

				case fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionAttributeTypeEnum_COMBOBOX:
					tempTestInstructionAttribute.AttributeInformation.InputComboBoxProperty.
						ComboBoxAttributeValue = attribute.AttributeChangedValue
					tempTestInstructionAttribute.AttributeInformation.InputComboBoxProperty.
						ComboBoxAttributeValueUuid = sharedCode.ZeroUuid

					// SPECIAL
					// When the attribute is the ComboBox that the user use to chose to which ExecutionDomain the TestData for the TestExecution
					// should be sent to. Then set DomainUuid and ExecutionDomainUuid for TestInstruction
					if attribute.AttributeUuid == string(testInstruction_SendTestDataToThisDomain_version_1_0.
						TestInstructionAttributeUUID_FenixSentToUsersDomain_SendTestDataToThisDomain_SendTestDataToThisExecutionDomain) ||
						attribute.AttributeUuid == string(testInstruction_SendTemplateToThisDomain_version_1_0.
							TestInstructionAttributeUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisExecutionDomainComboBox) {

						var executionDomainsThatCanReceiveDirectTargetedTestInstructionsMap map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
							ExecutionDomainsThatCanReceiveDirectTargetedTestInstructionsMessage
						executionDomainsThatCanReceiveDirectTargetedTestInstructionsMap = *sharedCode.ExecutionDomainsThatCanReceiveDirectTargetedTestInstructionsMapPtr

						domainUuid := executionDomainsThatCanReceiveDirectTargetedTestInstructionsMap[attribute.AttributeChangedValue].
							GetDomainUuid()
						domainName := executionDomainsThatCanReceiveDirectTargetedTestInstructionsMap[attribute.AttributeChangedValue].
							GetDomainName()
						executionDomainUuid := executionDomainsThatCanReceiveDirectTargetedTestInstructionsMap[attribute.AttributeChangedValue].
							GetExecutionDomainUuid()
						executionDomainName := executionDomainsThatCanReceiveDirectTargetedTestInstructionsMap[attribute.AttributeChangedValue].
							GetExecutionDomainName()

						// Update TestInstruction with DomainUuid and DomainName
						tempMatureTestInstruction.BasicTestInstructionInformation_NonEditableInformation.DomainUuid = domainUuid
						tempMatureTestInstruction.BasicTestInstructionInformation_NonEditableInformation.DomainName = domainName

						// Update TestInstruction with ExecutionDomainUuid and ExecutionDomainName
						tempMatureTestInstruction.BasicTestInstructionInformation_NonEditableInformation.ExecutionDomainUuid = executionDomainUuid
						tempMatureTestInstruction.BasicTestInstructionInformation_NonEditableInformation.ExecutionDomainName = executionDomainName

					}

					// Save back Attribute into TestInstruction
					tempMatureTestInstruction.TestInstructionAttributesList[attribute.AttributeUuid] = tempTestInstructionAttribute

					// Save back TestInstruction into TestCase
					testCase.MatureTestInstructionMap[testInstructionElementMatureUuidUuid] = tempMatureTestInstruction

				case fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionAttributeTypeEnum_RESPONSE_VARIABLE_COMBOBOX:
					tempTestInstructionAttribute.AttributeInformation.ResponseVariableComboBoxProperty.
						ComboBoxAttributeValueAsString = attribute.AttributeChangedValue

					// The Uuid of the chosen allowed response variable type
					tempTestInstructionAttribute.AttributeInformation.ResponseVariableComboBoxProperty.
						ChosenResponseVariableTypeUuid = attribute.AttributeResponseVariableComboBoxProperty.
						AttributeResponseVariableComboBoxProperty.GetChosenResponseVariableTypeUuid()

					// The Name of the chosen allowed response variable type
					tempTestInstructionAttribute.AttributeInformation.ResponseVariableComboBoxProperty.
						ChosenResponseVariableTypeName = attribute.AttributeResponseVariableComboBoxProperty.
						AttributeResponseVariableComboBoxProperty.GetChosenResponseVariableTypeName()

					tempTestInstructionAttribute.AttributeInformation.ResponseVariableComboBoxProperty.
						ComboBoxAttributeValueAsString = attribute.AttributeResponseVariableComboBoxProperty.
						AttributeResponseVariableComboBoxProperty.GetComboBoxAttributeValueAsString()

					// Save back Attribute into TestInstruction
					tempMatureTestInstruction.TestInstructionAttributesList[attribute.AttributeUuid] = tempTestInstructionAttribute

					// Save back TestInstruction into TestCase
					testCase.MatureTestInstructionMap[testInstructionElementMatureUuidUuid] = tempMatureTestInstruction

				default:

					errorId := "8d5c40ca-1a88-4eae-8926-898d03e6806b"
					err = errors.New(fmt.Sprintf("Unhandled AttributeType '%s'. [ErrorID: %s]",
						attribute.AttributeType.String(),
						errorId))

					// Hard Exit
					log.Fatalln(err) //TODO Send error over error-channel

				}

			}
		}
	}

	// Save back TestCase if any Attribute was changed
	if atleastOneAttributeIsChanged == true {
		testCaseModel.TestCases[testCaseUuid] = testCase
	}

	return err

}
