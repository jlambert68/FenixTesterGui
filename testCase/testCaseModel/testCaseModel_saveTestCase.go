package testCaseModel

import (
	"errors"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
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
	gRPCMatureTestCaseModelElementMessage, err := testCaseModel.generateTestCaseModelElementsForGrpc(testCaseUuid)
	if err != nil {
		return err
	}
	gRPCMatureTestInstructions, err := testCaseModel.generateatureTestInstructionsForGrpc(testCaseUuid)
	if err != nil {
		return err
	}

	// Save full TestCase in DB
	//rpc SaveFullTestCase(FullTestCaseMessage) returns (AckNackResponse)
	gRPCFullTestCaseMessageToSend := fenixGuiTestCaseBuilderServerGrpcApi.FullTestCaseMessage{
		TestCaseBasicInformation: &fenixGuiTestCaseBuilderServerGrpcApi.TestCaseBasicInformationMessage{
			BasicTestCaseInformation: &fenixGuiTestCaseBuilderServerGrpcApi.BasicTestCaseInformationMessage{
				NonEditableInformation:  &currentTestCase.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation,
				EditableInformation: &currentTestCase.LocalTestCaseMessage.BasicTestCaseInformationMessageEditableInformation,
			},
			CreatedAndUpdatedInformation: &fenixGuiTestCaseBuilderServerGrpcApi.TestCaseBasicInformationMessage_CreatedAndUpdatedInformationMessage{
				AddedToTestCaseTimeStamp: timeStampForTestCaseUpdate,
				AddedToTestCaseByUserId: currentActiveUser,
				LastUpdatedInTestCaseTimeStamp: timeStampForTestCaseUpdate,
				LastUpdatedInTestCaseByUserId: currentActiveUser,
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
			TestCaseFiles: &fenixGuiTestCaseBuilderServerGrpcApi.TestCaseFilesMessage{},
			UserIdentification: &fenixGuiTestCaseBuilderServerGrpcApi.UserIdentificationMessage{
				UserId:                       currentActiveUser,
				ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(
					testCaseModel.GrpcOutReference.GetHighestFenixGuiServerProtoFileVersion()),
			},
		},
		MatureTestInstructions: &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionsMessage{
			MatureTestInstructions: gRPCMatureTestInstructions,
		},
		MatureTestInstructionContainers: []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionContainerMessage{},
	}


	// Send using gRPC
	returnMessage := testCaseModel.GrpcOutReference.SendSaveFullTestCase(&gRPCFullTestCaseMessageToSend)

	if returnMessage.AckNack == false {

		errorId := "cb68859b-5c99-48a5-8f8b-9af472a9a45a"
		err = errors.New(fmt.Sprintf(returnMessage.Comments + "[ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return err
	}

	return err

}



// Convert the MatureTestCaseModelElementMessage-map into its gRPC-version
func (testCaseModel *TestCasesModelsStruct) generateatureTestInstructionsForGrpc(testCaseUuid string) (gRPCMatureTestInstructions []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionsMessage_MatureTestInstructionMessage, err error) {

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {

		errorId := "48899cab-ce9d-48a2-947f-d7610a3bea81"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return nil, err
	}

	// Loop map with all 'MatureTestCaseModelElementMessage' in the TestCase

	// Loop map with all 'MatureTestCaseModelElements' in the TestCase and create a slice
	for _, matureTestInstruction := range currentTestCase.MatureTestInstructionMap {

		MatureTestInstructionMessage := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionsMessage_MatureTestInstructionMessage{
			BasicTestInstructionInformation:  &fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionInformationMessage{
				NonEditableInformation:    &fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionInformationMessage_NonEditableBasicInformationMessage{
					DomainUuid:                  currentTestCase.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.DomainUuid,
					DomainName:                  currentTestCase.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.DomainName,
					TestInstructionOrignalUuid:   ,
					TestInstructionOriginalName: "",
					TestInstructionTypeUuid: string(matureTestInstruction.MatureBasicTestInstructionInformation.TestInstructionType),
					TestInstructionTypeName:     "",
					Deprecated:                  false,
					MajorVersionNumber:          0,
					MinorVersionNumber:          0,
					UpdatedTimeStamp:            nil,
					TestInstructionColor:        matureTestInstruction.,
					TCRuleDeletion:              "",
					TCRuleSwap:                  "",
				},
				EditableInformation:       &fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionInformationMessage_EditableBasicInformationMessage{
					TestInstructionDescription:   "",
					TestInstructionMouseOverText: "",
				},
				InvisibleBasicInformation: &fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionInformationMessage_InvisibleBasicInformationMessage{
					Enabled: true},
			},
			MatureTestInstructionInformation: nil,
		}
		gRPCMatureTestInstructions = append(gRPCMatureTestInstructions, &MatureTestInstructionMessage)
	}

	return gRPCMatureTestInstructions, err


}
// Convert the MatureTestCaseModelElementMessage-map into its gRPC-version
func (testCaseModel *TestCasesModelsStruct) generateTestCaseModelElementsForGrpc(testCaseUuid string) (gRPCMatureTestCaseModelElements []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage, err error) {

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {

		errorId := "48899cab-ce9d-48a2-947f-d7610a3bea81"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return nil, err
	}

	// Loop map with all 'MatureTestCaseModelElementMessage' in the TestCase and create a slice
	for _, matureTestCaseModelElement := range currentTestCase.TestCaseModelMap {
		gRPCMatureTestCaseModelElements = append(gRPCMatureTestCaseModelElements, &matureTestCaseModelElement.MatureTestCaseModelElementMessage)
	}

	return gRPCMatureTestCaseModelElements, err
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
