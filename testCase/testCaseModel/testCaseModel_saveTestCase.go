package testCaseModel

import (
	sharedCode "FenixTesterGui/common_code"
	"errors"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	testInstruction_SendTemplateToThisDomain_version_1_0 "github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTemplateToThisDomain/version_1_0"
	testInstruction_SendTestDataToThisDomain_version_1_0 "github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTestDataToThisDomain/version_1_0"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"os"
	"strings"
)

// SaveFullTestCase - Save the TestCase to the Database
func (testCaseModel *TestCasesModelsStruct) SaveFullTestCase(testCaseUuid string, currentActiveUser string) (err error) {

	var existsInMap bool

	// Save changed Attributes, if there are any, into the TestCase-model. Needs to call because last attributes change is not saved into model
	err = testCaseModel.SaveChangedTestCaseAttributeInTestCase(testCaseUuid)
	if err != nil {
		return err
	}

	// Get current TestCase
	var testCasesMap map[string]*TestCaseModelStruct
	testCasesMap = *testCaseModel.TestCasesMapPtr

	// Get current TestCase
	var currentTestCasePtr *TestCaseModelStruct
	currentTestCasePtr, existsInMap = testCasesMap[testCaseUuid]

	if existsInMap == false {

		errorId := "4c075798-ec6c-4486-8053-997ef0d0d8eb"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCasesMapPtr [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return err
	}

	//Copy MetaData from Gui, to TestCase-model, and validate that all mandatory MetaData-fields has values

	// Create timestamp to be used
	timeStampForTestCaseUpdate := timestamppb.Now()

	// Convert map-messages into gRPC-version, mostly arrays
	var (
		gRPCMatureTestCaseModelElementMessage    []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage
		gRPCMatureTestInstructions               []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionsMessage_MatureTestInstructionMessage
		gRPCMatureTestInstructionContainers      []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionContainersMessage_MatureTestInstructionContainerMessage
		gRPCTestCaseExtraInformation             *fenixGuiTestCaseBuilderServerGrpcApi.TestCaseExtraInformationMessage
		gRPCTestCaseTemplateFiles                *fenixGuiTestCaseBuilderServerGrpcApi.TestCaseTemplateFilesMessage
		gRPCTestCaseTestData                     *fenixGuiTestCaseBuilderServerGrpcApi.UsersChosenTestDataForTestCaseMessage
		gRPCTestCasePreviewMessage               *fenixGuiTestCaseBuilderServerGrpcApi.TestCasePreviewMessage
		gRPCUserSpecifiedTestCaseMetaDataMessage *fenixGuiTestCaseBuilderServerGrpcApi.UserSpecifiedTestCaseMetaDataMessage
		finalHash                                string
	)
	gRPCMatureTestCaseModelElementMessage,
		gRPCMatureTestInstructions,
		gRPCMatureTestInstructionContainers,
		gRPCTestCaseExtraInformation,
		gRPCTestCaseTemplateFiles,
		gRPCTestCaseTestData,
		gRPCTestCasePreviewMessage,
		gRPCUserSpecifiedTestCaseMetaDataMessage,
		finalHash, err = testCaseModel.generateTestCaseForGrpcAndHash(testCaseUuid, true)
	if err != nil {
		return err
	}

	// Check if changes are done to TestCase, but is only done if the TestCase is not saved before
	if currentTestCasePtr.ThisIsANewTestCase == true ||
		currentTestCasePtr.TestCaseHash != finalHash ||
		currentTestCasePtr.TestDataHash != gRPCTestCaseTestData.GetHashOfThisMessageWithEmptyHashField() {

		currentTestCasePtr.TestCaseHash = finalHash
		currentTestCasePtr.TestDataHash = gRPCTestCaseTestData.GetHashOfThisMessageWithEmptyHashField()

	} else {
		return nil

	}

	// Save full TestCase in DB
	//rpc SaveFullTestCase(FullTestCaseMessage) returns (AckNackResponse)
	gRPCFullTestCaseMessageToSend := fenixGuiTestCaseBuilderServerGrpcApi.FullTestCaseMessage{
		TestCaseBasicInformation: &fenixGuiTestCaseBuilderServerGrpcApi.TestCaseBasicInformationMessage{
			BasicTestCaseInformation: &fenixGuiTestCaseBuilderServerGrpcApi.BasicTestCaseInformationMessage{
				NonEditableInformation: &currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation,
				EditableInformation:    &currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageEditableInformation,
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
				TestCaseModelAsString:  currentTestCasePtr.TextualTestCaseRepresentationExtendedStack[0],
				FirstMatureElementUuid: currentTestCasePtr.FirstElementUuid,
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
		TestCaseExtraInformation: gRPCTestCaseExtraInformation,
		TestCaseTemplateFiles:    gRPCTestCaseTemplateFiles,
		TestCaseTestData:         gRPCTestCaseTestData,
		TestCasePreview:          gRPCTestCasePreviewMessage,
		TestCaseMetaData:         gRPCUserSpecifiedTestCaseMetaDataMessage,
		MessageHash:              currentTestCasePtr.TestCaseHash,
		//DeletedDate:              "",
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
	currentTestCasePtr.ThisIsANewTestCase = false

	// Update The Hash for the TestCase
	currentTestCasePtr.TestCaseHashWhenTestCaseWasSavedOrLoaded = gRPCFullTestCaseMessageToSend.MessageHash
	currentTestCasePtr.TestDataHashWhenTestCaseWasSavedOrLoaded = gRPCFullTestCaseMessageToSend.GetTestCaseTestData().GetHashOfThisMessageWithEmptyHashField()

	// Save the TestCase back in Map
	//testCaseModel.TestCasesMapPtr[testCaseUuid] = currentTestCasePtr

	return err

}

// Convert the MatureTestCaseTestInstructions to its gRPC-version
func (testCaseModel *TestCasesModelsStruct) generateMatureTestInstructionsForGrpc(
	testCaseUuid string) (
	gRPCMatureTestInstructions []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionsMessage_MatureTestInstructionMessage,
	hashedSlice string,
	valuesToBeHashedSlice []string,
	err error) {

	var existsInMap bool

	// Get TestCasesMap
	var testCasesMap map[string]*TestCaseModelStruct
	testCasesMap = *testCaseModel.TestCasesMapPtr

	// Get current TestCase
	var currentTestCasePtr *TestCaseModelStruct
	currentTestCasePtr, existsInMap = testCasesMap[testCaseUuid]

	if existsInMap == false {

		errorId := "48899cab-ce9d-48a2-947f-d7610a3bea81"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCasesMapPtr [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return nil, "", nil, err
	}

	// Loop map with all 'MatureTestInstructions' in the TestCase and create a slice
	for _, matureTestInstruction := range currentTestCasePtr.MatureTestInstructionMap {

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

	var existsInMap bool

	// Get current TestCase
	var testCasesMap map[string]*TestCaseModelStruct
	testCasesMap = *testCaseModel.TestCasesMapPtr

	// Get current TestCase
	var currentTestCasePtr *TestCaseModelStruct
	currentTestCasePtr, existsInMap = testCasesMap[testCaseUuid]

	if existsInMap == false {

		errorId := "82040ba6-57c4-47e2-8fb5-c770db41d8f8"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCasesMapPtr [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return nil, "", nil, err
	}

	// Loop map with all 'MatureTestInstructionContainers' in the TestCase and create a slice
	for _, matureTestInstructionContainer := range currentTestCasePtr.MatureTestInstructionContainerMap {
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

	var existsInMap bool

	// Get current TestCase
	var testCasesMap map[string]*TestCaseModelStruct
	testCasesMap = *testCaseModel.TestCasesMapPtr

	// Get current TestCase
	var currentTestCasePtr *TestCaseModelStruct
	currentTestCasePtr, existsInMap = testCasesMap[testCaseUuid]

	if existsInMap == false {

		errorId := "48899cab-ce9d-48a2-947f-d7610a3bea81"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCasesMapPtr [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return nil, "", nil, err
	}

	// Loop map with all 'MatureTestCaseModelElementMessage' in the TestCase and create a slice
	for _, matureTestCaseModelElement := range currentTestCasePtr.TestCaseModelMap {
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

	var existsInMap bool

	// Get current TestCase
	var testCasesMap map[string]*TestCaseModelStruct
	testCasesMap = *testCaseModel.TestCasesMapPtr

	// Get current TestCase
	var currentTestCasePtr *TestCaseModelStruct
	currentTestCasePtr, existsInMap = testCasesMap[testCaseUuid]

	if existsInMap == false {

		errorId := "e6fbdfdc-e0dc-4dd8-8ab1-b6be82b9e9fe"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCasesMapPtr [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return nil, "", nil, err
	}

	// Secure that the number of Textual Models are the same
	var (
		numberSimpleModels   int
		numberComplexModels  int
		numberExtendedModels int
	)
	numberSimpleModels = len(currentTestCasePtr.TextualTestCaseRepresentationSimpleStack)
	numberComplexModels = len(currentTestCasePtr.TextualTestCaseRepresentationComplexStack)
	numberExtendedModels = len(currentTestCasePtr.TextualTestCaseRepresentationExtendedStack)

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
				currentTestCasePtr.TextualTestCaseRepresentationSimpleStack[modelCounter])

		// Complex
		tempTestCaseTextualRepresentationHistory.TextualTestCaseRepresentationComplexHistory =
			append(tempTestCaseTextualRepresentationHistory.TextualTestCaseRepresentationComplexHistory,
				currentTestCasePtr.TextualTestCaseRepresentationComplexStack[modelCounter])

		// Extended
		tempTestCaseTextualRepresentationHistory.TextualTestCaseRepresentationExtendedStackHistory =
			append(tempTestCaseTextualRepresentationHistory.TextualTestCaseRepresentationExtendedStackHistory,
				currentTestCasePtr.TextualTestCaseRepresentationExtendedStack[modelCounter])

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

	var existsInMap bool

	// Get current TestCase
	var testCasesMap map[string]*TestCaseModelStruct
	testCasesMap = *testCaseModel.TestCasesMapPtr

	// Get current TestCase
	var currentTestCasePtr *TestCaseModelStruct
	currentTestCasePtr, existsInMap = testCasesMap[testCaseUuid]

	if existsInMap == false {

		errorId := "e6ceecbe-00e2-42af-9782-eb83af2d03c2"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map witsh all TestCasesMapPtr [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return nil, "", err
	}

	// Loop map with all 'TestCaseTemplateFiles' in the TestCase and add to gPRC-version
	for _, importedTemplateFileFromGitHub := range currentTestCasePtr.ImportedTemplateFilesFromGitHub {

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

	var existsInMap bool

	// Get current TestCase
	var testCasesMap map[string]*TestCaseModelStruct
	testCasesMap = *testCaseModel.TestCasesMapPtr

	// Get current TestCase
	var currentTestCasePtr *TestCaseModelStruct
	currentTestCasePtr, existsInMap = testCasesMap[testCaseUuid]

	if existsInMap == false {

		errorId := "f6354f22-5bca-41f8-8b92-b168b4381708"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map witsh all TestCasesMapPtr [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return nil, err
	}

	// The gRPC-version of 'testDataPointNameMap'
	var chosenTestDataPointsPerGroupMapGrpc map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
		TestDataPointNameMapMessage
	chosenTestDataPointsPerGroupMapGrpc = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
		TestDataPointNameMapMessage)

	if currentTestCasePtr.TestData != nil {

		// Loop TestDataGroups
		for testDataGroupName, testDataPointNameMap := range currentTestCasePtr.TestData.ChosenTestDataPointsPerGroupMap {

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

// Convert the TestCasePreviewMessage into its gRPC-version
func (testCaseModel *TestCasesModelsStruct) generateTestCasePreviewMessageForGrpc(
	testCaseUuid string) (
	gRPCTestCasePreviewMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestCasePreviewMessage,
	err error) {

	var existsInMap bool

	// Get TestCasesMap
	var testCasesMap map[string]*TestCaseModelStruct
	testCasesMap = *testCaseModel.TestCasesMapPtr

	// Get current TestCase
	var currentTestCasePtr *TestCaseModelStruct
	currentTestCasePtr, existsInMap = testCasesMap[testCaseUuid]

	if existsInMap == false {

		errorId := "9f3050d4-920e-4065-b7c9-f9d41bee7fb6"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map witsh all TestCasesMapPtr [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return nil, err
	}

	// Generate the Domain that owns the TestCase
	var domainNameThatOwnsTestCase string
	domainNameThatOwnsTestCase = fmt.Sprintf("%s [%s]",
		currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.DomainName,
		currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.DomainUuid[0:8])

	// Generate the Complex Textual Model of the TestCase
	var complexTextualModel string
	complexTextualModel = currentTestCasePtr.TextualTestCaseRepresentationComplexStack[len(
		currentTestCasePtr.TextualTestCaseRepresentationComplexStack)-1]

	// Loop all TestInstructions and create Attributes for Preview-model
	var existInMap bool
	var tempMatureTestInstruction MatureTestInstructionStruct
	for index, testCaseStructureObject := range currentTestCasePtr.TestCasePreviewObject.TestCaseStructureObjects {

		// Check if this is a TestInstruction, if so then add the Attributes to the TestInstruction in Preview-model
		if testCaseStructureObject.GetTestCaseStructureObjectType() == fenixGuiTestCaseBuilderServerGrpcApi.
			TestCasePreviewStructureMessage_TestInstruction {

			// Extract the Mature TestInstruction
			tempMatureTestInstruction, existInMap = currentTestCasePtr.MatureTestInstructionMap[testCaseStructureObject.GetTestInstructionUuid()]
			if existInMap == false {
				errorId := "9f35c094-f35e-47a6-b005-9a637a6eabe9"
				err = errors.New(fmt.Sprintf("TestInstruction '%s' is missing in map with all TestCasesMapPtr [ErrorID: %s]",
					testCaseStructureObject.GetTestInstructionUuid(), errorId))

				// TODO Send on Error-channel
				fmt.Println(err)

				return
			}

			var tempPreviewAttributes []*fenixGuiTestCaseBuilderServerGrpcApi.
				TestCasePreviewStructureMessage_TestCaseStructureObjectMessage_TestInstructionAttributeMessage
			// Loop Attributes in the Mature TestInstruction and add to Preview-model
			for _, tempAttribute := range tempMatureTestInstruction.TestInstructionAttributesList {

				var tempPreviewAttribute *fenixGuiTestCaseBuilderServerGrpcApi.
					TestCasePreviewStructureMessage_TestCaseStructureObjectMessage_TestInstructionAttributeMessage
				tempPreviewAttribute = &fenixGuiTestCaseBuilderServerGrpcApi.
					TestCasePreviewStructureMessage_TestCaseStructureObjectMessage_TestInstructionAttributeMessage{
					AttributeName:      tempAttribute.BaseAttributeInformation.GetTestInstructionAttributeName(),
					AttributeValue:     "",
					AttributeGroupName: tempAttribute.BaseAttributeInformation.GetTestInstructionAttributeTypeName(),
				}

				// Get the value based on type
				switch tempAttribute.BaseAttributeInformation.GetTestInstructionAttributeType() {

				case fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionAttributeTypeEnum_TEXTBOX:
					tempPreviewAttribute.AttributeValue = tempAttribute.AttributeInformation.InputTextBoxProperty.
						GetTextBoxAttributeValue()

				case fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionAttributeTypeEnum_COMBOBOX:
					tempPreviewAttribute.AttributeValue = tempAttribute.AttributeInformation.InputComboBoxProperty.
						GetComboBoxAttributeValue()

				case fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionAttributeTypeEnum_RESPONSE_VARIABLE_COMBOBOX:
					tempPreviewAttribute.AttributeValue = tempAttribute.AttributeInformation.ResponseVariableComboBoxProperty.
						GetComboBoxAttributeValueAsString()

				case fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionAttributeTypeEnum_TESTCASE_BUILDER_SERVER_INJECTED_COMBOBOX:
					tempPreviewAttribute.AttributeValue = "<Unknown value for Attribute for TESTCASE_BUILDER_SERVER_INJECTED_COMBOBOX >"

				default:
					tempPreviewAttribute.AttributeValue = "<Unknown value for Attribute>"

				}

				// Add Attribute to slice of attributes for TestInstruction
				tempPreviewAttributes = append(tempPreviewAttributes, tempPreviewAttribute)
			}

			// Add the attributes to the TestInstruction
			testCaseStructureObject.TestInstructionAttributes = tempPreviewAttributes
		}

		// Save back the testCaseStructureObject
		currentTestCasePtr.TestCasePreviewObject.TestCaseStructureObjects[index] = testCaseStructureObject
	}

	// Generate Selected MetaData to be used when filtering TestCases
	var tempSelectedMetaDataValuesMap map[string]*fenixGuiTestCaseBuilderServerGrpcApi.TestCasePreviewStructureMessage_SelectedMetaDataValueMessage
	tempSelectedMetaDataValuesMap = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.TestCasePreviewStructureMessage_SelectedMetaDataValueMessage)

	// Get the MetaData for the TestCase
	var tempTestCaseMetaData TestCaseMetaDataStruct
	tempTestCaseMetaData = *currentTestCasePtr.TestCaseMetaDataPtr

	// Get the MetaDataGroupsMap
	var tempMetaDataGroupsMap map[string]*MetaDataGroupStruct
	tempMetaDataGroupsMap = *tempTestCaseMetaData.MetaDataGroupsMapPtr

	var selectedMetaDataValuesMapKey string

	// Loop all MetaDataGroups
	for _, tempMetaDataGroupPtr := range tempMetaDataGroupsMap {

		// Get the MetaDataGroup
		var tempMetaDataGroup MetaDataGroupStruct
		tempMetaDataGroup = *tempMetaDataGroupPtr

		// Loop all MetaDataGroupItems
		for _, tempMetaDataGroupItemPtr := range *tempMetaDataGroup.MetaDataInGroupMapPtr {

			// Get the MetaDataGroupItem
			var tempMetaDataGroupItem MetaDataInGroupStruct
			tempMetaDataGroupItem = *tempMetaDataGroupItemPtr

			// Is the Item a 'single select' or a 'multi select'
			switch tempMetaDataGroupItem.SelectType {

			case MetaDataSelectType_SingleSelect:

				if len(tempMetaDataGroupItem.SelectedMetaDataValueForSingleSelect) > 0 {
					// Add selected value to the 'SelectedMetaDataValuesMap'

					// Create the map-key
					selectedMetaDataValuesMapKey = fmt.Sprintf("%s.%s.%s.%s",
						currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.DomainUuid,
						tempMetaDataGroupItem.MetaDataGroupName,
						tempMetaDataGroupItem.MetaDataName,
						tempMetaDataGroupItem.SelectedMetaDataValueForSingleSelect)

					// Create the value to be inserted into the map
					var tempSelectedMetaDataValueMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestCasePreviewStructureMessage_SelectedMetaDataValueMessage
					tempSelectedMetaDataValueMessage = &fenixGuiTestCaseBuilderServerGrpcApi.TestCasePreviewStructureMessage_SelectedMetaDataValueMessage{
						OwnerDomainUuid:   currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.DomainUuid,
						OwnerDomainName:   domainNameThatOwnsTestCase,
						MetaDataGroupName: tempMetaDataGroupItem.MetaDataGroupName,
						MetaDataName:      tempMetaDataGroupItem.MetaDataName,
						MetaDataNameValue: tempMetaDataGroupItem.SelectedMetaDataValueForSingleSelect,
						SelectType:        fenixGuiTestCaseBuilderServerGrpcApi.MetaDataSelectTypeEnum(tempMetaDataGroupItem.SelectType),
						IsMandatory:       tempMetaDataGroupItem.Mandatory,
					}

					// dd selected value to the 'SelectedMetaDataValuesMap'
					tempSelectedMetaDataValuesMap[selectedMetaDataValuesMapKey] = tempSelectedMetaDataValueMessage
				}

			case MetaDataSelectType_MultiSelect:

				if len(tempMetaDataGroupItem.SelectedMetaDataValuesForMultiSelect) > 0 {
					// Add selected value to the 'SelectedMetaDataValuesMap'

					// Loop SelectedMetaDataValuesForMultiSelect
					for _, tempSelectedMetaDataValueForMultiSelect := range tempMetaDataGroupItem.SelectedMetaDataValuesForMultiSelect {

						// Create the map-key
						selectedMetaDataValuesMapKey = fmt.Sprintf("%s.%s.%s.%s",
							currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.DomainUuid,
							tempMetaDataGroupItem.MetaDataGroupName,
							tempMetaDataGroupItem.MetaDataName,
							tempSelectedMetaDataValueForMultiSelect)

						// Create the value to be inserted into the map
						var tempSelectedMetaDataValueMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestCasePreviewStructureMessage_SelectedMetaDataValueMessage
						tempSelectedMetaDataValueMessage = &fenixGuiTestCaseBuilderServerGrpcApi.TestCasePreviewStructureMessage_SelectedMetaDataValueMessage{
							OwnerDomainUuid:   currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.DomainUuid,
							OwnerDomainName:   domainNameThatOwnsTestCase,
							MetaDataGroupName: tempMetaDataGroupItem.MetaDataGroupName,
							MetaDataName:      tempMetaDataGroupItem.MetaDataName,
							MetaDataNameValue: tempSelectedMetaDataValueForMultiSelect,
							SelectType:        fenixGuiTestCaseBuilderServerGrpcApi.MetaDataSelectTypeEnum(tempMetaDataGroupItem.SelectType),
							IsMandatory:       tempMetaDataGroupItem.Mandatory,
						}

						// Add selected value to the 'SelectedMetaDataValuesMap'
						tempSelectedMetaDataValuesMap[selectedMetaDataValuesMapKey] = tempSelectedMetaDataValueMessage
					}
				}

			default:

				errorId := "f2a6571a-f267-4d83-83d3-247eb50d6002"

				errorMessage := fmt.Sprintf("Unknown SelectType for MetaDataGroupItem '%d' [ErrorID: %s]",
					tempMetaDataGroupItem.SelectType, errorId)

				log.Fatal(errorMessage)

			}

		}

	}

	// Generate the full TestCasePreview-object
	var tempTestCasePreview *fenixGuiTestCaseBuilderServerGrpcApi.TestCasePreviewStructureMessage
	tempTestCasePreview = &fenixGuiTestCaseBuilderServerGrpcApi.TestCasePreviewStructureMessage{
		TestCaseName:                    currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageEditableInformation.GetTestCaseName(),
		DomainThatOwnTheTestCase:        domainNameThatOwnsTestCase,
		TestCaseDescription:             currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageEditableInformation.GetTestCaseDescription(),
		TestCaseStructureObjects:        currentTestCasePtr.TestCasePreviewObject.TestCaseStructureObjects,
		ComplexTextualDescription:       complexTextualModel,
		TestCaseUuid:                    currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.GetTestCaseUuid(),
		TestCaseVersion:                 "",
		LastSavedByUserOnComputer:       sharedCode.CurrentUserIdLogedInOnComputer,
		LastSavedByUserGCPAuthorization: sharedCode.CurrentUserAuthenticatedTowardsGCP,
		LastSavedTimeStamp:              "",
		SelectedMetaDataValuesMap:       tempSelectedMetaDataValuesMap,
	}

	gRPCTestCasePreviewMessage = &fenixGuiTestCaseBuilderServerGrpcApi.TestCasePreviewMessage{
		TestCasePreview:     tempTestCasePreview,
		TestCasePreviewHash: "",
	}

	// Generate Hash of gRPC-message and add it to the message
	tempJson := protojson.Format(gRPCTestCasePreviewMessage)

	// Remove spaces before hashing, due to some bug that generates "double space" sometimes when running in non-debug-mode
	tempJson = strings.ReplaceAll(tempJson, " ", "")

	hashedJson := sharedCode.HashSingleValue(tempJson)
	gRPCTestCasePreviewMessage.TestCasePreviewHash = hashedJson

	return gRPCTestCasePreviewMessage, err

}

// Convert the UserSpecifiedTestCaseMetaDataMessage into its gRPC-version
func (testCaseModel *TestCasesModelsStruct) generateUserSpecifiedTestCaseMetaDataMessageForGrpc(
	testCaseUuid string, shouldBeSaved bool) (
	gRPCUserSpecifiedTestCaseMetaDataMessage *fenixGuiTestCaseBuilderServerGrpcApi.UserSpecifiedTestCaseMetaDataMessage,
	hashedSlice string,
	err error) {

	// SLice holding the values that will become the MetaDataSlice
	var valuesToBeHashedSlice []string
	var valueToBeHashed string

	var existsInMap bool

	// Get TestCasesMap
	var testCasesMap map[string]*TestCaseModelStruct
	testCasesMap = *testCaseModel.TestCasesMapPtr

	// Get current TestCase
	var currentTestCasePtr *TestCaseModelStruct
	currentTestCasePtr, existsInMap = testCasesMap[testCaseUuid]

	if existsInMap == false {

		errorId := "5b6834dd-0763-4d12-9d4a-d97281a93a46"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map witsh all TestCasesMapPtr [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return nil, "", err
	}

	// Generate the Domain that owns the TestCase
	var domainNameThatOwnsTestCase string
	if len(currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.DomainUuid) > 0 {
		domainNameThatOwnsTestCase = fmt.Sprintf("%s [%s]",
			currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.DomainName,
			currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.DomainUuid[0:8])

		gRPCUserSpecifiedTestCaseMetaDataMessage = &fenixGuiTestCaseBuilderServerGrpcApi.UserSpecifiedTestCaseMetaDataMessage{
			CurrentSelectedDomainUuid: currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.DomainUuid,
			CurrentSelectedDomainName: domainNameThatOwnsTestCase,
			MetaDataGroupsMap:         nil,
		}
	}

	// Verify that all mandatory MetaData-fields have been set
	err = testCaseModel.verifyMandatoryFieldsForMetaData(
		gRPCUserSpecifiedTestCaseMetaDataMessage.GetCurrentSelectedDomainUuid(),
		currentTestCasePtr,
		shouldBeSaved)

	// All mandatory fields have not been set
	if err != nil {
		return nil, "", err
	}

	// Check if there are any TestCaseMetaData, if not then just return
	if currentTestCasePtr.TestCaseMetaDataPtr == nil ||
		currentTestCasePtr.TestCaseMetaDataPtr.MetaDataGroupsMapPtr == nil {
		return nil, "", err
	}

	var tempMetaDataGroupsMap map[string]*fenixGuiTestCaseBuilderServerGrpcApi.MetaDataGroupMessage
	tempMetaDataGroupsMap = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.MetaDataGroupMessage)

	// Get the TestCaseMetaData from TestCaseMetaDataPtr
	//var TestCaseMetaData

	// Loop MetaDataGroups in the TestCase and extract each MetaDataGroup
	for tempMetaGroupNameInTestCase, tempMetaDataGroupInTestCasePtr := range *currentTestCasePtr.TestCaseMetaDataPtr.MetaDataGroupsMapPtr {

		// Get the MetaDataGroupInTestCaseMap
		var tempMetaDataGroupInTestCaseMap map[string]*MetaDataInGroupStruct
		tempMetaDataGroupInTestCaseMap = *tempMetaDataGroupInTestCasePtr.MetaDataInGroupMapPtr

		// Create the MetaDataInGroupItem-map
		var metaDataInGroupMessage map[string]*fenixGuiTestCaseBuilderServerGrpcApi.MetaDataInGroupMessage
		metaDataInGroupMessage = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.MetaDataInGroupMessage)

		// Loop items in MetaDataGroup
		for tempMetaDataGroupItemNameInTestCase, tempMetaDataGroupItemInTestCase := range tempMetaDataGroupInTestCaseMap {

			// Create the MetaDataGroupItem for the gRPC-message
			var metaDataGroupItem fenixGuiTestCaseBuilderServerGrpcApi.MetaDataInGroupMessage
			metaDataGroupItem = fenixGuiTestCaseBuilderServerGrpcApi.MetaDataInGroupMessage{
				MetaDataGroupName:                       tempMetaGroupNameInTestCase,
				MetaDataName:                            tempMetaDataGroupItemNameInTestCase,
				SelectType:                              fenixGuiTestCaseBuilderServerGrpcApi.MetaDataSelectTypeEnum(tempMetaDataGroupItemInTestCase.SelectType),
				IsMandatory:                             tempMetaDataGroupItemInTestCase.Mandatory,
				AvailableMetaDataValues:                 tempMetaDataGroupItemInTestCase.AvailableMetaDataValues,
				SelectedMetaDataValueForSingleSelect:    tempMetaDataGroupItemInTestCase.SelectedMetaDataValueForSingleSelect,
				SelectedMetaDataValuesForMultiSelect:    tempMetaDataGroupItemInTestCase.SelectedMetaDataValuesForMultiSelect,
				SelectedMetaDataValuesForMultiSelectMap: *tempMetaDataGroupItemInTestCase.SelectedMetaDataValuesForMultiSelectMapPtr,
			}

			// Add the MetaDataGroupItem for the gRPC-map-message
			metaDataInGroupMessage[tempMetaDataGroupItemNameInTestCase] = &metaDataGroupItem

			// Generate shared values, for SingleSelect and MultiSelect, to be hashed
			valueToBeHashed = fmt.Sprintf("%s-%s-%t-%d",
				tempMetaGroupNameInTestCase,
				tempMetaDataGroupItemNameInTestCase,
				tempMetaDataGroupItemInTestCase.Mandatory,
				tempMetaDataGroupItemInTestCase.SelectType)

			valuesToBeHashedSlice = append(valuesToBeHashedSlice, valueToBeHashed)

			// Add Available values to be hashed
			for _, tempMetaDataValue := range tempMetaDataGroupItemInTestCase.AvailableMetaDataValues {
				valueToBeHashed = fmt.Sprintf("%s",
					tempMetaDataValue)

				valuesToBeHashedSlice = append(valuesToBeHashedSlice, valueToBeHashed)
			}

			// Generate values to be hashed, depending on Single or Multiple Select
			switch tempMetaDataGroupItemInTestCase.SelectType {
			case MetaDataSelectType_SingleSelect:

				valueToBeHashed = fmt.Sprintf("%s",
					tempMetaDataGroupItemInTestCase.SelectedMetaDataValueForSingleSelect)

				valuesToBeHashedSlice = append(valuesToBeHashedSlice, valueToBeHashed)

			case MetaDataSelectType_MultiSelect:

				for _, tempMetaDataValue := range tempMetaDataGroupItemInTestCase.SelectedMetaDataValuesForMultiSelect {
					valueToBeHashed = fmt.Sprintf("%s",
						tempMetaDataValue)

					valuesToBeHashedSlice = append(valuesToBeHashedSlice, valueToBeHashed)
				}

			default:
				errorId := ""

				log.Fatalln(fmt.Sprintf("Unhandled tempMetaDataGroupItemInTestCase.SelectType",
					tempMetaDataGroupItemInTestCase.SelectType,
					errorId))
			}

		}

		// Create the MetaDataGroupMessage
		var metaDataGroupMessage fenixGuiTestCaseBuilderServerGrpcApi.MetaDataGroupMessage
		metaDataGroupMessage = fenixGuiTestCaseBuilderServerGrpcApi.MetaDataGroupMessage{
			MetaDataGroupName:  tempMetaGroupNameInTestCase,
			MetaDataInGroupMap: metaDataInGroupMessage,
		}

		// Add the MetaDataGroup-message for the gRPC-map-message
		tempMetaDataGroupsMap[tempMetaGroupNameInTestCase] = &metaDataGroupMessage

	}

	// Add the MetaDataGroupsMap to the gRPC-object
	gRPCUserSpecifiedTestCaseMetaDataMessage.MetaDataGroupsMap = tempMetaDataGroupsMap

	// Generate Hash of all sub-message-hashes
	hashedSlice = sharedCode.HashValues(valuesToBeHashedSlice, false)

	return gRPCUserSpecifiedTestCaseMetaDataMessage, hashedSlice, err
}

// Pack different parts of the TestCase into gRPC-version into one message together with Hash of TestCase
func (testCaseModel *TestCasesModelsStruct) generateTestCaseForGrpcAndHash(testCaseUuid string, shouldBeSaved bool) (
	gRPCMatureTestCaseModelElementMessage []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage,
	gRPCMatureTestInstructions []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionsMessage_MatureTestInstructionMessage,
	gRPCMatureTestInstructionContainers []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionContainersMessage_MatureTestInstructionContainerMessage,
	gRPCTestCaseExtraInformation *fenixGuiTestCaseBuilderServerGrpcApi.TestCaseExtraInformationMessage,
	gRPCTestCaseTemplateFiles *fenixGuiTestCaseBuilderServerGrpcApi.TestCaseTemplateFilesMessage,
	gRPCTestCaseTestData *fenixGuiTestCaseBuilderServerGrpcApi.UsersChosenTestDataForTestCaseMessage,
	gRPCTestCasePreviewMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestCasePreviewMessage,
	gRPCUserSpecifiedTestCaseMetaDataMessage *fenixGuiTestCaseBuilderServerGrpcApi.UserSpecifiedTestCaseMetaDataMessage,
	finalHash string,
	err error) {

	var existsInMap bool

	// Get TestCasesMap
	var testCasesMap map[string]*TestCaseModelStruct
	testCasesMap = *testCaseModel.TestCasesMapPtr

	// Get current TestCase
	var currentTestCasePtr *TestCaseModelStruct
	currentTestCasePtr, existsInMap = testCasesMap[testCaseUuid]

	if existsInMap == false {

		errorId := "4c075798-ec6c-4486-8053-997ef0d0d8eb"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCasesMapPtr [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			"",
			err
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
		return nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			"",
			err
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
		return nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			"",
			err
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
		return nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			"",
			err
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
		return nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			"",
			err
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
		return nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			"",
			err
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
		return nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			"",
			err
	}

	// Only Generate Preview when an actual SaveTestCase is performed
	if shouldBeSaved == true {

		// gRPCTestCasePreviewMessage
		gRPCTestCasePreviewMessage, err =
			testCaseModel.generateTestCasePreviewMessageForGrpc(testCaseUuid)
		if err != nil {
			return nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				"",
				err
		}
	}

	// Generate UserSpecifiedTestCaseMetaData, gRPCUserSpecifiedTestCaseMetaDataMessage
	var hashedgRPCTestCaseMetaData string
	gRPCUserSpecifiedTestCaseMetaDataMessage, hashedgRPCTestCaseMetaData, err =
		testCaseModel.generateUserSpecifiedTestCaseMetaDataMessageForGrpc(testCaseUuid, shouldBeSaved)
	if err != nil {
		return nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			"",
			err
	}

	var valuesToReHash []string

	// NonEditableInformation, start by clearing Version because it s not the same after save to Database
	currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.TestCaseVersion = 0
	//tempNonEditableInformation := fmt.Sprint(&currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation)

	// Generate Hash for  'tempNonEditableInformation'
	valuesToBeHashedSlice = nil
	tempJson := protojson.Format(&currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation)
	valuesToBeHashedSlice = append(valuesToBeHashedSlice, tempJson)

	// Remove spaces before hashing, due to some bug that generates "double space" sometimes when running in non-debug-mode
	for index, textToReplaceIn := range valuesToBeHashedSlice {
		valuesToBeHashedSlice[index] = strings.ReplaceAll(textToReplaceIn, " ", "")
	}

	hashNonEditableInformation := sharedCode.HashSingleValue(valuesToBeHashedSlice[0])
	valuesToReHash = append(valuesToReHash, hashNonEditableInformation)

	// Add hash and values to slice
	var tempNonEditableInformationValue subHashPartsMapValueType
	tempNonEditableInformationValue = subHashPartsMapValueType{
		nameOfContent:        "tempNonEditableInformation",
		hash:                 hashNonEditableInformation,
		contentAsStringSlice: []string{tempJson},
	}
	subHashPartsSlice = append(subHashPartsSlice, tempNonEditableInformationValue)

	// EditableInformation
	//tempEditableInformation := fmt.Sprint(&currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageEditableInformation)
	// Generate Hash for  'tempNonEditableInformation'
	valuesToBeHashedSlice = nil
	tempJson = protojson.Format(&currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageEditableInformation)
	valuesToBeHashedSlice = append(valuesToBeHashedSlice, tempJson)

	// Remove spaces before hashing, due to some bug that generates "double space" sometimes when running in non-debug-mode
	for index, textToReplaceIn := range valuesToBeHashedSlice {
		valuesToBeHashedSlice[index] = strings.ReplaceAll(textToReplaceIn, " ", "")
	}

	hashEditableInformation := sharedCode.HashSingleValue(valuesToBeHashedSlice[0])
	valuesToReHash = append(valuesToReHash, hashEditableInformation)

	// Add hash and values to slice
	var tempEditableInformationValue subHashPartsMapValueType
	tempEditableInformationValue = subHashPartsMapValueType{
		nameOfContent:        "tempEditableInformation",
		hash:                 hashEditableInformation,
		contentAsStringSlice: []string{tempJson},
	}
	subHashPartsSlice = append(subHashPartsSlice, tempEditableInformationValue)

	// FirstMatureElementUuid
	tempFirstMatureElementUuid := fmt.Sprint(currentTestCasePtr.FirstElementUuid)
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

	valuesToReHash = append(valuesToReHash, hashedgRPCTestCaseMetaData)

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
	//sharedCode.Logger.WithFields(logrus.Fields{
	//	"ID":                "d0034233-202b-469c-93e3-48903991fe23",
	//	"testCaseUuid":      testCaseUuid,
	//	"subHashPartsSlice": subHashPartsSlice,
	//}).Debug("Used for bug findings when TestCase model gets of sync when saving TestCase and the Loading the same TestCase")

	/*
		for _, subHashPartSlice := range subHashPartsSlice {
			fmt.Println(fmt.Sprintf("TestCaseUuid = '%s', NameOfContent = '%s', Hash = '%s', ContentAsStringSlice = '%s'",
				testCaseUuid,
				subHashPartSlice.nameOfContent,
				subHashPartSlice.hash,
				subHashPartSlice.contentAsStringSlice))
		}
	*/

	// Open the file in append mode. If it doesn't exist, create it.
	file, err := os.OpenFile("SaveLoadTestCaseProblems.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close() // Ensure the file is closed when the function exits

	for _, subHashPartSlice := range subHashPartsSlice {
		// Format the string as before
		line := fmt.Sprintf("TestCaseUuid = '%s', NameOfContent = '%s', Hash = '%s', ContentAsStringSlice = '%s'\n",
			testCaseUuid,
			subHashPartSlice.nameOfContent,
			subHashPartSlice.hash,
			subHashPartSlice.contentAsStringSlice)

		// Write the string to the file
		if _, err := file.WriteString(line); err != nil {
			fmt.Printf("Error writing to file: %v\n", err)

		}
	}

	return gRPCMatureTestCaseModelElementMessage,
		gRPCMatureTestInstructions,
		gRPCMatureTestInstructionContainers,
		gRPCTestCaseExtraInformation,
		gRPCTestCaseTemplateFiles,
		gRPCTestCaseTestData,
		gRPCTestCasePreviewMessage,
		gRPCUserSpecifiedTestCaseMetaDataMessage,
		finalHash,
		err

}

// SaveChangedTestCaseAttributeInTestCase - Save changed Attributes into the TestCase-model under correct TestInstruction
func (testCaseModel *TestCasesModelsStruct) SaveChangedTestCaseAttributeInTestCase(testCaseUuid string) (err error) {

	//var atLeastOneAttributeIsChanged bool

	var existInMap bool

	// Get TestCasesMap
	var testCasesMap map[string]*TestCaseModelStruct
	testCasesMap = *testCaseModel.TestCasesMapPtr

	// Get current TestCase
	var currentTestCasePtr *TestCaseModelStruct
	currentTestCasePtr, existInMap = testCasesMap[testCaseUuid]

	if existInMap == false {

		errorId := "40fc730f-87d4-4c44-96ff-ab1003e40751"
		err := errors.New(fmt.Sprintf("currentTestCasePtr %s is missing in TestCasesMapPtr-map [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) //TODO Send error over error-channel
		return err
	}

	// If there are no attributes in  'AttributesList' for TestCase (No unsaved changes exists) then exist
	if currentTestCasePtr.AttributesList == nil || len(*currentTestCasePtr.AttributesList) == 0 {
		return err
	}

	// Extract testInstructionElementMatureUuidUuid
	attributesList := *currentTestCasePtr.AttributesList

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

				//atLeastOneAttributeIsChanged = true

				// Extract TestInstruction
				tempMatureTestInstruction, existInMap := currentTestCasePtr.MatureTestInstructionMap[testInstructionElementMatureUuidUuid]
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
					currentTestCasePtr.MatureTestInstructionMap[testInstructionElementMatureUuidUuid] = tempMatureTestInstruction

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
					currentTestCasePtr.MatureTestInstructionMap[testInstructionElementMatureUuidUuid] = tempMatureTestInstruction

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
					currentTestCasePtr.MatureTestInstructionMap[testInstructionElementMatureUuidUuid] = tempMatureTestInstruction

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
	//if atLeastOneAttributeIsChanged == true {
	//	testCaseModel.TestCasesMapPtr[testCaseUuid] = currentTestCasePtr
	//}

	return err

}
