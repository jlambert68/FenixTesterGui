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

	// Create timestamp to be used
	timeStampForTestCaseUpdate := timestamppb.Now()

	// Convert map-messages into gRPC-version, mostly arrays
	// TestCase-model
	gRPCMatureTestCaseModelElementMessage, hashedMatureTestCaseModelElements, err := testCaseModel.generateTestCaseModelElementsForGrpc(testCaseUuid)
	if err != nil {
		return err
	}
	// TestInstructions
	gRPCMatureTestInstructions, hashedgRPCMatureTestInstructions, err := testCaseModel.generateMatureTestInstructionsForGrpc(testCaseUuid)
	if err != nil {
		return err
	}
	// TestInstructionContainers
	gRPCMatureTestInstructionContainers, hashedgRPCMatureTestInstructionContainers, err := testCaseModel.generateMatureTestInstructionContainersForGrpc(testCaseUuid)
	if err != nil {
		return err
	}

	var valuesToReHash []string

	tempNonEditableInformation := fmt.Sprint(&currentTestCase.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation)
	hashNonEditableInformation := sharedCode.HashSingleValue(tempNonEditableInformation)
	valuesToReHash = append(valuesToReHash, hashNonEditableInformation)

	tempEditableInformation := fmt.Sprint(&currentTestCase.LocalTestCaseMessage.BasicTestCaseInformationMessageEditableInformation)
	hashEditableInformation := sharedCode.HashSingleValue(tempEditableInformation)
	valuesToReHash = append(valuesToReHash, hashEditableInformation)

	tempTestCaseModelAsString := fmt.Sprint(currentTestCase.TextualTestCaseRepresentationExtendedStack)
	hashTestCaseModelAsString := sharedCode.HashSingleValue(tempTestCaseModelAsString)
	valuesToReHash = append(valuesToReHash, hashTestCaseModelAsString)

	tempFirstMatureElementUuid := fmt.Sprint(currentTestCase.FirstElementUuid)
	hashFirstMatureElementUuid := sharedCode.HashSingleValue(tempFirstMatureElementUuid)
	valuesToReHash = append(valuesToReHash, hashFirstMatureElementUuid)

	valuesToReHash = append(valuesToReHash, hashedMatureTestCaseModelElements)

	valuesToReHash = append(valuesToReHash, hashedgRPCMatureTestInstructions)

	valuesToReHash = append(valuesToReHash, hashedgRPCMatureTestInstructionContainers)

	finalHash := sharedCode.HashValues(valuesToReHash, false)

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
					testCaseModel.GrpcOutReference.GetHighestFenixGuiServerProtoFileVersion()),
			},
		},
		MatureTestInstructions: &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionsMessage{
			MatureTestInstructions: gRPCMatureTestInstructions,
		},
		MatureTestInstructionContainers: &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionContainersMessage{
			MatureTestInstructionContainers: gRPCMatureTestInstructionContainers},
		MessageHash: currentTestCase.TestCaseHash,
	}

	// Send using gRPC
	returnMessage := testCaseModel.GrpcOutReference.SendSaveFullTestCase(&gRPCFullTestCaseMessageToSend)

	if returnMessage.AckNack == false {

		errorId := "cb68859b-5c99-48a5-8f8b-9af472a9a45a"
		err = errors.New(fmt.Sprintf(returnMessage.Comments+"[ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return err
	}

	// Update that the TestCase is not New anymore
	currentTestCase.ThisIsANewTestCase = false
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

		// Generate Hashes for  'matureTestInstruction'
		tempJson := protojson.Format(matureTestInstruction.BasicTestInstructionInformation_NonEditableInformation)
		hashSlice = append(hashSlice, tempJson)
		tempJson = protojson.Format(matureTestInstruction.BasicTestInstructionInformation_EditableInformation)
		hashSlice = append(hashSlice, tempJson)
		tempJson = protojson.Format(matureTestInstruction.BasicTestInstructionInformation_InvisibleBasicInformation)
		hashSlice = append(hashSlice, tempJson)
		tempJson = protojson.Format(matureTestInstruction.MatureBasicTestInstructionInformation)
		hashSlice = append(hashSlice, tempJson)

		// Loop over all 'TestInstruction Attributes' in the TestInstruction and create slice
		var testInstructionAttributesList []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage
		for _, testInstructionAttribute := range matureTestInstruction.TestInstructionAttributesList {
			testInstructionAttributesList = append(testInstructionAttributesList, testInstructionAttribute)

			// Generate Hash for  'testInstructionAttribute'
			tempJson := protojson.Format(testInstructionAttribute)
			hashSlice = append(hashSlice, tempJson)

		}

		// Create one 'MatureTestInstructionMessage'
		MatureTestInstructionMessage := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionsMessage_MatureTestInstructionMessage{
			BasicTestInstructionInformation: &fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionInformationMessage{
				NonEditableInformation:    matureTestInstruction.BasicTestInstructionInformation_NonEditableInformation,
				EditableInformation:       matureTestInstruction.BasicTestInstructionInformation_EditableInformation,
				InvisibleBasicInformation: matureTestInstruction.BasicTestInstructionInformation_InvisibleBasicInformation,
			},
			MatureTestInstructionInformation: &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage{
				MatureBasicTestInstructionInformation: matureTestInstruction.MatureBasicTestInstructionInformation,
				CreatedAndUpdatedInformation:          matureTestInstruction.CreatedAndUpdatedInformation,
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

		// Generate Hashes for  'matureTestInstruction'
		tempJson := protojson.Format(matureTestInstructionContainer.NonEditableInformation)
		hashSlice = append(hashSlice, tempJson)
		tempJson = protojson.Format(matureTestInstructionContainer.EditableInformation)
		hashSlice = append(hashSlice, tempJson)
		tempJson = protojson.Format(matureTestInstructionContainer.InvisibleBasicInformation)
		hashSlice = append(hashSlice, tempJson)
		tempJson = protojson.Format(matureTestInstructionContainer.EditableTestInstructionContainerAttributes)
		hashSlice = append(hashSlice, tempJson)
		tempJson = protojson.Format(matureTestInstructionContainer.MatureTestInstructionContainerInformation)
		hashSlice = append(hashSlice, tempJson)
		tempJson = protojson.Format(matureTestInstructionContainer.CreatedAndUpdatedInformation)
		hashSlice = append(hashSlice, tempJson)

		// Create one 'MatureTestInstructionContainerMessage'
		MatureTestInstructionContainerMessage := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionContainersMessage_MatureTestInstructionContainerMessage{
			BasicTestInstructionContainerInformation: &fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionContainerInformationMessage{
				NonEditableInformation:                     matureTestInstructionContainer.NonEditableInformation,
				EditableInformation:                        matureTestInstructionContainer.EditableInformation,
				InvisibleBasicInformation:                  matureTestInstructionContainer.InvisibleBasicInformation,
				EditableTestInstructionContainerAttributes: matureTestInstructionContainer.EditableTestInstructionContainerAttributes,
			},
			MatureTestInstructionContainerInformation: &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionContainerInformationMessage{
				MatureTestInstructionContainerInformation: matureTestInstructionContainer.MatureTestInstructionContainerInformation,
				CreatedAndUpdatedInformation:              matureTestInstructionContainer.CreatedAndUpdatedInformation,
			},
		}

		// Add 'MatureTestInstructionContainerMessage' to slice
		gRPCMatureTestInstructionContainers = append(gRPCMatureTestInstructionContainers, &MatureTestInstructionContainerMessage)
	}

	// Generate Hash of all sub-message-hashes
	hashedSlice = sharedCode.HashValues(hashSlice, false)

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
		gRPCMatureTestCaseModelElements = append(gRPCMatureTestCaseModelElements, &matureTestCaseModelElement.MatureTestCaseModelElementMessage)

		// Generate Hash for  'matureTestCaseModelElement'
		tempJson := protojson.Format(&matureTestCaseModelElement.MatureTestCaseModelElementMessage)
		hashSlice = append(hashSlice, tempJson)

	}

	// Generate Hash of all sub-message-hashes
	hashedSlice = sharedCode.HashValues(hashSlice, false)

	return gRPCMatureTestCaseModelElements, hashedSlice, err
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
