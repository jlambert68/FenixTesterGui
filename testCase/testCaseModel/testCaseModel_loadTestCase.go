package testCaseModel

import (
	"errors"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
)

// LoadFullTestCaseFromDatabase - Load the TestCase from the Database into model
func (testCaseModel *TestCasesModelsStruct) LoadFullTestCaseFromDatabase(testCaseUuid string, currentActiveUser string) (err error) {

	// Send LoadTesCase using gRPC
	var detailedTestCaseResponse *fenixGuiTestCaseBuilderServerGrpcApi.GetDetailedTestCaseResponse
	detailedTestCaseResponse = testCaseModel.GrpcOutReference.LoadDetailedTestCase(currentActiveUser, testCaseUuid)

	// Exit if something was wrong
	if detailedTestCaseResponse.AckNackResponse.AckNack == false {

		errorId := "ba195459-8902-4727-ab81-ae48cd616eea"
		err = errors.New(fmt.Sprintf(detailedTestCaseResponse.AckNackResponse.Comments+"[ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return err

	}

	// Create object that will hold complete TestCase in memory
	var tempTestCaseModel TestCaseModelStruct
	tempTestCaseModel = TestCaseModelStruct{
		LastLoadedTestCaseModelGRPCMessage:         *detailedTestCaseResponse.DetailedTestCase.TestCaseBasicInformation.TestCaseModel, // fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage{},
		FirstElementUuid:                           detailedTestCaseResponse.DetailedTestCase.TestCaseBasicInformation.TestCaseModel.FirstMatureElementUuid,
		TestCaseModelMap:                           nil, // Created below
		TextualTestCaseRepresentationSimpleStack:   nil,
		TextualTestCaseRepresentationComplexStack:  nil,
		TextualTestCaseRepresentationExtendedStack: nil,
		CommandStack:                               nil,
		LastSavedCommandStack:                      lastSavedCommandStack{},
		CopyBuffer:                                 ImmatureElementStruct{},
		CutBuffer:                                  MatureElementStruct{},
		CutCommandInitiated:                        false,
		LocalTestCaseMessage: LocalTestCaseMessageStruct{
			BasicTestCaseInformationMessageNoneEditableInformation: *detailedTestCaseResponse.DetailedTestCase.
				TestCaseBasicInformation.BasicTestCaseInformation.GetNonEditableInformation(),
			BasicTestCaseInformationMessageEditableInformation: *detailedTestCaseResponse.DetailedTestCase.
				TestCaseBasicInformation.BasicTestCaseInformation.GetEditableInformation(),
			CreatedAndUpdatedInformation: *detailedTestCaseResponse.DetailedTestCase.
				TestCaseBasicInformation.GetCreatedAndUpdatedInformation(),
		},
		testCaseModelAdaptedForUiTree:     nil,
		CurrentSelectedTestCaseElement:    currentSelectedTestCaseElementStruct{},
		MatureTestInstructionMap:          nil, // Created below
		MatureTestInstructionContainerMap: nil, // Created below
		AttributesList:                    nil,
		ThisIsANewTestCase:                false,
		TestCaseHash:                      detailedTestCaseResponse.DetailedTestCase.MessageHash,
	}

	// Generate 'TestCaseModelMap'
	tempTestCaseModel.TestCaseModelMap = make(map[string]MatureTestCaseModelElementStruct)
	for _, tempMatureTestCaseModelElementMessage := range detailedTestCaseResponse.DetailedTestCase.
		TestCaseBasicInformation.TestCaseModel.TestCaseModelElements {

		var tempMatureTestCaseModelElement MatureTestCaseModelElementStruct
		tempMatureTestCaseModelElement = MatureTestCaseModelElementStruct{
			MatureTestCaseModelElementMessage: fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
				OriginalElementUuid:      tempMatureTestCaseModelElementMessage.OriginalElementUuid,
				OriginalElementName:      tempMatureTestCaseModelElementMessage.OriginalElementName,
				MatureElementUuid:        tempMatureTestCaseModelElementMessage.MatureElementUuid,
				PreviousElementUuid:      tempMatureTestCaseModelElementMessage.PreviousElementUuid,
				NextElementUuid:          tempMatureTestCaseModelElementMessage.NextElementUuid,
				FirstChildElementUuid:    tempMatureTestCaseModelElementMessage.FirstChildElementUuid,
				ParentElementUuid:        tempMatureTestCaseModelElementMessage.ParentElementUuid,
				TestCaseModelElementType: tempMatureTestCaseModelElementMessage.TestCaseModelElementType,
			},
			MatureTestCaseModelElementMetaData: MatureTestCaseModelElementMetaDataStruct{
				ChosenDropZoneUuid:        "",
				ChosenDropZoneColorString: "",
			},
		}
		tempTestCaseModel.TestCaseModelMap[tempMatureTestCaseModelElementMessage.MatureElementUuid] = tempMatureTestCaseModelElement

	}

	// Generate 'MatureTestInstructionMap'
	tempTestCaseModel.MatureTestInstructionMap = make(map[string]MatureTestInstructionStruct)
	for _, tempMatureTestInstructionMessage := range detailedTestCaseResponse.DetailedTestCase.MatureTestInstructions.
		MatureTestInstructions {

		// Add TestInstruction
		var tempMatureTestInstruction MatureTestInstructionStruct
		tempMatureTestInstruction = MatureTestInstructionStruct{
			BasicTestInstructionInformation_NonEditableInformation: tempMatureTestInstructionMessage.
				BasicTestInstructionInformation.NonEditableInformation,
			BasicTestInstructionInformation_EditableInformation: tempMatureTestInstructionMessage.
				BasicTestInstructionInformation.EditableInformation,
			BasicTestInstructionInformation_InvisibleBasicInformation: tempMatureTestInstructionMessage.
				BasicTestInstructionInformation.InvisibleBasicInformation,
			MatureBasicTestInstructionInformation: tempMatureTestInstructionMessage.
				MatureTestInstructionInformation.MatureBasicTestInstructionInformation,
			CreatedAndUpdatedInformation: tempMatureTestInstructionMessage.
				MatureTestInstructionInformation.CreatedAndUpdatedInformation,
			TestInstructionAttributesList: make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
				MatureTestInstructionInformationMessage_TestInstructionAttributeMessage),
		}

		// Add Attributes for TestInstruction
		for _, tempTestInstructionAttributeMessage := range tempMatureTestInstructionMessage.
			MatureTestInstructionInformation.TestInstructionAttributesList {
			var tempTestInstructionAttributes fenixGuiTestCaseBuilderServerGrpcApi.
				MatureTestInstructionInformationMessage_TestInstructionAttributeMessage
			tempTestInstructionAttributes = fenixGuiTestCaseBuilderServerGrpcApi.
				MatureTestInstructionInformationMessage_TestInstructionAttributeMessage{
				BaseAttributeInformation: tempTestInstructionAttributeMessage.BaseAttributeInformation,
				AttributeInformation:     tempTestInstructionAttributeMessage.AttributeInformation,
			}

			tempMatureTestInstruction.TestInstructionAttributesList[tempTestInstructionAttributeMessage.
				BaseAttributeInformation.TestInstructionAttributeUuid] = &tempTestInstructionAttributes
		}

		tempTestCaseModel.MatureTestInstructionMap[tempMatureTestInstructionMessage.MatureTestInstructionInformation.
			MatureBasicTestInstructionInformation.TestInstructionMatureUuid] = tempMatureTestInstruction

	}

	// Generate 'MatureTestInstructionContainerMap'
	tempTestCaseModel.MatureTestInstructionContainerMap = make(map[string]MatureTestInstructionContainerStruct)
	for _, tempMatureTestInstructionContainerMessage := range detailedTestCaseResponse.DetailedTestCase.
		MatureTestInstructionContainers.MatureTestInstructionContainers {

		// Add TestInstructionContainer
		var tempMatureTestInstructionContainer MatureTestInstructionContainerStruct
		tempMatureTestInstructionContainer = MatureTestInstructionContainerStruct{
			NonEditableInformation: tempMatureTestInstructionContainerMessage.
				BasicTestInstructionContainerInformation.NonEditableInformation,
			EditableInformation: tempMatureTestInstructionContainerMessage.
				BasicTestInstructionContainerInformation.EditableInformation,
			InvisibleBasicInformation: tempMatureTestInstructionContainerMessage.
				BasicTestInstructionContainerInformation.InvisibleBasicInformation,
			EditableTestInstructionContainerAttributes: tempMatureTestInstructionContainerMessage.
				BasicTestInstructionContainerInformation.EditableTestInstructionContainerAttributes,
			MatureTestInstructionContainerInformation: tempMatureTestInstructionContainerMessage.
				MatureTestInstructionContainerInformation.MatureTestInstructionContainerInformation,
			CreatedAndUpdatedInformation: tempMatureTestInstructionContainerMessage.
				MatureTestInstructionContainerInformation.CreatedAndUpdatedInformation,
		}

		tempTestCaseModel.MatureTestInstructionContainerMap[tempMatureTestInstructionContainerMessage.
			MatureTestInstructionContainerInformation.MatureTestInstructionContainerInformation.
			TestInstructionContainerMatureUuid] = tempMatureTestInstructionContainer
	}

	// Add TestCase to map with all TestCases
	if testCaseModel.TestCases == nil {
		testCaseModel.TestCases = make(map[string]TestCaseModelStruct)
	}
	testCaseModel.TestCases[tempTestCaseModel.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.TestCaseUuid] = tempTestCaseModel

	return err

}
