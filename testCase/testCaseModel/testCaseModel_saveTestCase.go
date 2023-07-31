package testCaseModel

import (
	sharedCode "FenixTesterGui/common_code"
	"errors"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		finalHash                             string
	)
	gRPCMatureTestCaseModelElementMessage,
		gRPCMatureTestInstructions,
		gRPCMatureTestInstructionContainers,
		gRPCTestCaseExtraInformation,
		finalHash, err = testCaseModel.generateTestCaseForGrpcAndHash(testCaseUuid)
	if err != nil {
		return err
	}

	// Check if changes are done to TestCase, but is only done if the TestCase is not saved before
	if currentTestCase.ThisIsANewTestCase == true ||
		currentTestCase.TestCaseHash != finalHash {

		currentTestCase.TestCaseHash = finalHash

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
				UserId: currentActiveUser,
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

	// Save the TestCase back in Map
	testCaseModel.TestCases[testCaseUuid] = currentTestCase

	return err

}

// Convert the MatureTestCaseTestInstructions to its gRPC-version
func (testCaseModel *TestCasesModelsStruct) generateMatureTestInstructionsForGrpc(testCaseUuid string) (gRPCMatureTestInstructions []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionsMessage_MatureTestInstructionMessage, hashedSlice string, err error) {

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {

		errorId := "48899cab-ce9d-48a2-947f-d7610a3bea81"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return nil, "", err
	}

	var hashSlice []string

	// Loop map with all 'MatureTestInstructions' in the TestCase and create a slice
	for _, matureTestInstruction := range currentTestCase.MatureTestInstructionMap {

		var tempMatureTestInstruction MatureTestInstructionStruct
		tempMatureTestInstruction = matureTestInstruction

		// Generate Hashes for  'matureTestInstruction'
		tempJson := protojson.Format(tempMatureTestInstruction.BasicTestInstructionInformation_NonEditableInformation)
		hashSlice = append(hashSlice, tempJson)
		tempJson = protojson.Format(tempMatureTestInstruction.BasicTestInstructionInformation_EditableInformation)
		hashSlice = append(hashSlice, tempJson)
		tempJson = protojson.Format(tempMatureTestInstruction.BasicTestInstructionInformation_InvisibleBasicInformation)
		hashSlice = append(hashSlice, tempJson)
		tempJson = protojson.Format(tempMatureTestInstruction.MatureBasicTestInstructionInformation)
		hashSlice = append(hashSlice, tempJson)

		// Loop over all 'TestInstruction Attributes' in the TestInstruction and create slice
		var testInstructionAttributesList []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage
		for _, testInstructionAttribute := range tempMatureTestInstruction.TestInstructionAttributesList {

			testInstructionAttributesList = append(testInstructionAttributesList, testInstructionAttribute)

			// Generate Hash for  'testInstructionAttribute'
			tempJson := protojson.Format(testInstructionAttribute)
			hashSlice = append(hashSlice, tempJson)

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

	// Generate Hash of all sub-message-hashes
	hashedSlice = sharedCode.HashValues(hashSlice, false)

	return gRPCMatureTestInstructions, hashedSlice, err

}

// Convert the MatureTestCaseTestInstructionContainers to its gRPC-version
func (testCaseModel *TestCasesModelsStruct) generateMatureTestInstructionContainersForGrpc(testCaseUuid string) (gRPCMatureTestInstructionContainers []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionContainersMessage_MatureTestInstructionContainerMessage, hashedSlice string, err error) {

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {

		errorId := "82040ba6-57c4-47e2-8fb5-c770db41d8f8"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return nil, "", err
	}

	var hashSlice []string

	// Loop map with all 'MatureTestInstructionContainers' in the TestCase and create a slice
	for _, matureTestInstructionContainer := range currentTestCase.MatureTestInstructionContainerMap {
		var tempMatureTestInstructionContainer MatureTestInstructionContainerStruct
		tempMatureTestInstructionContainer = matureTestInstructionContainer

		// Generate Hashes for  'matureTestInstruction'
		tempJson := protojson.Format(tempMatureTestInstructionContainer.NonEditableInformation)
		hashSlice = append(hashSlice, tempJson)
		tempJson = protojson.Format(tempMatureTestInstructionContainer.EditableInformation)
		hashSlice = append(hashSlice, tempJson)
		tempJson = protojson.Format(tempMatureTestInstructionContainer.InvisibleBasicInformation)
		hashSlice = append(hashSlice, tempJson)
		tempJson = protojson.Format(tempMatureTestInstructionContainer.EditableTestInstructionContainerAttributes)
		hashSlice = append(hashSlice, tempJson)
		tempJson = protojson.Format(tempMatureTestInstructionContainer.MatureTestInstructionContainerInformation)
		hashSlice = append(hashSlice, tempJson)
		//tempJson = protojson.Format(tempMatureTestInstructionContainer.CreatedAndUpdatedInformation)
		//hashSlice = append(hashSlice, tempJson)

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

	// Generate Hash of all sub-message-hashes
	hashedSlice = sharedCode.HashValues(hashSlice, true)

	return gRPCMatureTestInstructionContainers, hashedSlice, err

}

// Convert the MatureTestCaseModelElementMessage-map into its gRPC-version
func (testCaseModel *TestCasesModelsStruct) generateTestCaseModelElementsForGrpc(testCaseUuid string) (gRPCMatureTestCaseModelElements []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage, hashedSlice string, err error) {

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {

		errorId := "48899cab-ce9d-48a2-947f-d7610a3bea81"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return nil, "", err
	}

	var hashSlice []string

	// Loop map with all 'MatureTestCaseModelElementMessage' in the TestCase and create a slice
	for _, matureTestCaseModelElement := range currentTestCase.TestCaseModelMap {
		var tempMatureTestCaseModelElement MatureTestCaseModelElementStruct
		tempMatureTestCaseModelElement = matureTestCaseModelElement
		gRPCMatureTestCaseModelElements = append(gRPCMatureTestCaseModelElements, &tempMatureTestCaseModelElement.MatureTestCaseModelElementMessage)

		// Generate Hash for  'matureTestCaseModelElement'
		tempJson := protojson.Format(&matureTestCaseModelElement.MatureTestCaseModelElementMessage)
		hashSlice = append(hashSlice, tempJson)

	}

	// Generate Hash of all sub-message-hashes
	hashedSlice = sharedCode.HashValues(hashSlice, false)

	return gRPCMatureTestCaseModelElements, hashedSlice, err
}

// Convert the TestCaseExtraInformationMessage into its gRPC-version
// Containing: 1) Textual Representation of TestCase
func (testCaseModel *TestCasesModelsStruct) generateTestCaseExtraInformationForGrpc(testCaseUuid string) (gRPCTestCaseExtraInformation *fenixGuiTestCaseBuilderServerGrpcApi.TestCaseExtraInformationMessage, hashedSlice string, err error) {

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {

		errorId := "e6fbdfdc-e0dc-4dd8-8ab1-b6be82b9e9fe"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return nil, "", err
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

		return nil, "", err
	}

	var hashSlice []string

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
	hashSlice = append(hashSlice, tempJson)

	// Generate Hash of all sub-message-hashes
	hashedSlice = sharedCode.HashValues(hashSlice, false)

	// Create return message 'gRPCTestCaseExtraInformation'
	gRPCTestCaseExtraInformation = &fenixGuiTestCaseBuilderServerGrpcApi.TestCaseExtraInformationMessage{
		TestCaseTextualRepresentationHistory: &tempTestCaseTextualRepresentationHistory,
	}

	return gRPCTestCaseExtraInformation, hashedSlice, err

}

// Pack different parts of the TestCase into gRPC-version into one message together with Hash of TestCase
func (testCaseModel *TestCasesModelsStruct) generateTestCaseForGrpcAndHash(testCaseUuid string) (
	gRPCMatureTestCaseModelElementMessage []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage,
	gRPCMatureTestInstructions []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionsMessage_MatureTestInstructionMessage,
	gRPCMatureTestInstructionContainers []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionContainersMessage_MatureTestInstructionContainerMessage,
	gRPCTestCaseExtraInformation *fenixGuiTestCaseBuilderServerGrpcApi.TestCaseExtraInformationMessage,
	finalHash string,
	err error) {

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {

		errorId := "4c075798-ec6c-4486-8053-997ef0d0d8eb"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return nil, nil, nil, nil, "", err
	}

	// Convert map-messages into gRPC-version, mostly arrays
	// TestCase-model
	var hashedMatureTestCaseModelElements string
	gRPCMatureTestCaseModelElementMessage, hashedMatureTestCaseModelElements, err = testCaseModel.generateTestCaseModelElementsForGrpc(testCaseUuid)
	if err != nil {
		return nil, nil, nil, nil, "", err
	}

	// TestInstructions
	var hashedgRPCMatureTestInstructions string
	gRPCMatureTestInstructions, hashedgRPCMatureTestInstructions, err = testCaseModel.generateMatureTestInstructionsForGrpc(testCaseUuid)
	if err != nil {
		return nil, nil, nil, nil, "", err
	}

	// TestInstructionContainers
	var hashedgRPCMatureTestInstructionContainers string
	gRPCMatureTestInstructionContainers, hashedgRPCMatureTestInstructionContainers, err = testCaseModel.generateMatureTestInstructionContainersForGrpc(testCaseUuid)
	if err != nil {
		return nil, nil, nil, nil, "", err
	}

	// TestCaseExtraInformation
	var hashedgRPCTestCaseExtraInformation string
	gRPCTestCaseExtraInformation, hashedgRPCTestCaseExtraInformation, err = testCaseModel.generateTestCaseExtraInformationForGrpc(testCaseUuid)
	if err != nil {
		return nil, nil, nil, nil, "", err
	}

	var valuesToReHash []string

	tempNonEditableInformation := fmt.Sprint(&currentTestCase.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation)
	hashNonEditableInformation := sharedCode.HashSingleValue(tempNonEditableInformation)
	valuesToReHash = append(valuesToReHash, hashNonEditableInformation)

	tempEditableInformation := fmt.Sprint(&currentTestCase.LocalTestCaseMessage.BasicTestCaseInformationMessageEditableInformation)
	hashEditableInformation := sharedCode.HashSingleValue(tempEditableInformation)
	valuesToReHash = append(valuesToReHash, hashEditableInformation)

	//tempTestCaseModelAsString := fmt.Sprint(currentTestCase.TextualTestCaseRepresentationExtendedStack)
	//hashTestCaseModelAsString := sharedCode.HashSingleValue(tempTestCaseModelAsString)
	//valuesToReHash = append(valuesToReHash, hashTestCaseModelAsString)

	tempFirstMatureElementUuid := fmt.Sprint(currentTestCase.FirstElementUuid)
	hashFirstMatureElementUuid := sharedCode.HashSingleValue(tempFirstMatureElementUuid)
	valuesToReHash = append(valuesToReHash, hashFirstMatureElementUuid)

	valuesToReHash = append(valuesToReHash, hashedMatureTestCaseModelElements)

	valuesToReHash = append(valuesToReHash, hashedgRPCMatureTestInstructions)

	valuesToReHash = append(valuesToReHash, hashedgRPCMatureTestInstructionContainers)

	valuesToReHash = append(valuesToReHash, hashedgRPCTestCaseExtraInformation)

	finalHash = sharedCode.HashValues(valuesToReHash, false)

	return gRPCMatureTestCaseModelElementMessage,
		gRPCMatureTestInstructions,
		gRPCMatureTestInstructionContainers,
		gRPCTestCaseExtraInformation,
		finalHash,
		err

}

// SaveChangedTestCaseAttributeInTestCase - Save changed Attributes into the TestCase-model under correct TestInstruction
func (testCaseModel *TestCasesModelsStruct) SaveChangedTestCaseAttributeInTestCase(testCaseUuid string) (err error) {

	// Extract current TestCase
	testCase, existInMap := testCaseModel.TestCases[testCaseUuid]
	if existInMap == false {

		errorId := "40fc730f-87d4-4c44-96ff-ab1003e40751"
		err := errors.New(fmt.Sprintf("testCase %s is missing in TestCases-map [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) //TODO Send error over error-channel
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
				tempTestInstructionAttribute.AttributeInformation.InputTextBoxProperty.TextBoxAttributeValue = attribute.AttributeChangedValue

			}
		}
	}

	return err

}
