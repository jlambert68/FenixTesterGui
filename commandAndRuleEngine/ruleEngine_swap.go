package commandAndRuleEngine

import (
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Verify if anor element can be swapped or not, regarding swap rules
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) verifyIfElementCanBeSwapped(testCaseUuid string, elementUuidToBeSwappedOut string, elementTypeToBeSwappedIn fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum) (canBeSwapped bool, matchedSimpledRule string, matchedComplexRule string, err error) {

	// First verify towards simple rules
	canBeSwapped, matchedSimpledRule, err = commandAndRuleEngine.verifyIfComponentCanBeSwappedSimpleRules(testCaseUuid, elementUuidToBeSwappedOut)

	// Only check complex rules if simple rules was OK for swapping
	if !(canBeSwapped == true &&
		err == nil) {
		return canBeSwapped, matchedSimpledRule, "", err
	}

	// Verify towards complex rules
	matchedComplexRule, err = commandAndRuleEngine.verifyIfComponentCanBeSwappedWithComplexRules(testCaseUuid, elementUuidToBeSwappedOut, elementTypeToBeSwappedIn)

	return canBeSwapped, matchedSimpledRule, matchedComplexRule, err
}

// Swap an element, but first ensure that rules for swapping are used
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeSwapElement(testCaseUuid string, elementToSwapOutUuid string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (err error) {

	// Get ElementType to be swapped in
	topElementToBeSwappedIn, existInMap := immatureElementToSwapIn.ImmatureElementMap[immatureElementToSwapIn.FirstElementUuid]
	if existInMap == false {

		errorId := "0ba64da2-6dcd-4f4c-b742-0838eede4f49"
		err = errors.New(fmt.Sprintf("element referenced by first element ('%s')  doesn't exist in element-map for ImmatureElement. TestCase '%s' [ErrorID: %s]", immatureElementToSwapIn.FirstElementUuid, testCaseUuid, errorId))

		return err
	}

	elementTypeToBeSwappedIn := topElementToBeSwappedIn.TestCaseModelElementType

	// Verify that element is allowed, and can be swapped
	canBeSwapped, matchedSimpleRule, matchedComplexRule, err := commandAndRuleEngine.verifyIfElementCanBeSwapped(testCaseUuid, elementToSwapOutUuid, elementTypeToBeSwappedIn)

	// If there was an error from swap verification then exit
	if err != nil {
		return err
	}

	// If the component couldn't be swapped then exit with error message
	if canBeSwapped == false {
		err = errors.New("element couldn't be swapped due to swap rule '" + matchedSimpleRule + "' or that complex rules aren't met")

		return err
	}

	// Execute deletion of element
	err = commandAndRuleEngine.executeSwapElementBasedOnRule(testCaseUuid, elementToSwapOutUuid, immatureElementToSwapIn, matchedComplexRule)

	return err
}

// Execute a swap on an element based on specific rule
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeSwapElementBasedOnRule(testCaseUuid string, elementToBeSwappedIOutUuid string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct, matchedComplexRule string) (err error) {

	var matureElementToSwapIn testCaseModel.MatureElementStruct

	switch matchedComplexRule {
	case TCRuleSwap101:
		matureElementToSwapIn, err = commandAndRuleEngine.executeTCRuleSwap101(testCaseUuid, elementToBeSwappedIOutUuid, immatureElementToSwapIn)

	case TCRuleSwap102:
		matureElementToSwapIn, err = commandAndRuleEngine.executeTCRuleSwap102(testCaseUuid, elementToBeSwappedIOutUuid, immatureElementToSwapIn)

	case TCRuleSwap103:
		matureElementToSwapIn, err = commandAndRuleEngine.executeTCRuleSwap103(testCaseUuid, elementToBeSwappedIOutUuid, immatureElementToSwapIn)

	case TCRuleSwap104:
		matureElementToSwapIn, err = commandAndRuleEngine.executeTCRuleSwap104(testCaseUuid, elementToBeSwappedIOutUuid, immatureElementToSwapIn)

	case TCRuleSwap105:
		commandAndRuleEngine.executeTCRuleSwap105(testCaseUuid, elementToBeSwappedIOutUuid, immatureElementToSwapIn)

	case TCRuleSwap106:
		matureElementToSwapIn, err = commandAndRuleEngine.executeTCRuleSwap106(testCaseUuid, elementToBeSwappedIOutUuid, immatureElementToSwapIn)

	case TCRuleSwap107:
		matureElementToSwapIn, err = commandAndRuleEngine.executeTCRuleSwap107(testCaseUuid, elementToBeSwappedIOutUuid, immatureElementToSwapIn)

	case TCRuleSwap108:
		matureElementToSwapIn, err = commandAndRuleEngine.executeTCRuleSwap108(testCaseUuid, elementToBeSwappedIOutUuid, immatureElementToSwapIn)

	default:
		commandAndRuleEngine.logger.WithFields(logrus.Fields{
			"id":                 "eba1a213-aa42-4021-aaea-4b3107d5874c",
			"matchedComplexRule": matchedComplexRule,
		}).Error(" Unknown 'matchedComplexRule' was used when trying to swap")

		err = errors.New("'" + matchedComplexRule + "' is an unknown complex Swap rule")

		return err
	}

	// Exit if there was an error
	if err != nil {
		return err
	}

	// Move TestInstruction data to TestCase
	err = commandAndRuleEngine.addTestInstructionDataToTestCaseModel(testCaseUuid, immatureElementToSwapIn, &matureElementToSwapIn)

	return err

}

// Add All TestInstruction-data for the new TestInstruction into the TestCase-model
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) addTestInstructionDataToTestCaseModel(testCaseUuid string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct, matureElementToSwapIn *testCaseModel.MatureElementStruct) (err error) {

	// Extract TestCase to work with
	currentTestCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	if existsInMap == false {

		errorId := "ea7e4f3f-f6c8-4391-a191-116f60c6b5f5"
		err = errors.New(fmt.Sprintf("testCase-model with UUID ('%s') doesn't exist in TestModel-map [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err.Error()) //TODO Send to Error-channel

		return err
	}

	// Verify that TestInstruction doesn't exit in TestInstructionMap
	_, existsInMap = currentTestCase.MatureTestInstructionMap[matureElementToSwapIn.FirstElementUuid]
	if existsInMap == true {

		errorId := "9f659bc5-7088-4bf7-900e-c9e12b4ce36d"
		err = errors.New(fmt.Sprintf("Mature TestInstruction with UUID '%s' already exist in MatureTestInstructionMap in TestCase: %s [ErrorID: %s]", matureElementToSwapIn.FirstElementUuid, testCaseUuid, errorId))

		fmt.Println(err.Error()) //TODO Send to Error-channel

		return err
	}

	// Generate searchable map of

	// Loop over all elements in 'matureElementToSwapIn' and only process TestInstructions , TI or TIx
	for _, matureElement := range matureElementToSwapIn.MatureElementMap {

		// If found a TI or TIx, then process that one
		if matureElement.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			matureElement.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE {

			// Create timestamp when TestInstruction was added
			createdTimeStamp := timestamppb.Now()

			// Create a new Mature TestInstruction to be added
			newMatureTestInstruction := testCaseModel.MatureTestInstructionStruct{
				MatureBasicTestInstructionInformation: &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_MatureBasicTestInstructionInformationMessage{
					TestCaseUuid:                             testCaseUuid,
					TestInstructionMatureUuid:                matureElementToSwapIn.FirstElementUuid,
					ParentTestInstructionContainerUuid:       "",
					ParentTestInstructionContainerMatureUuid: "",
					ChosenDropZoneUuid:                       immatureElementToSwapIn.ChosenDropZoneColor,
					ChosenDropZoneName:                       matureElementToSwapIn.ChosenDropZoneName,
					TestInstructionType:                      matureElement.TestCaseModelElementType,
				},
				CreatedAndUpdatedInformation: &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_CreatedAndUpdatedInformationMessage{
					AddedToTestCaseTimeStamp:       createdTimeStamp,
					AddedToTestCaseByUserId:        "",
					LastUpdatedInTestCaseTimeStamp: createdTimeStamp,
					LastUpdatedInTestCaseByUserId:  commandAndRuleEngine.Testcases.CurrentUser,
					DeletedFromTestCaseTimeStamp: &timestamppb.Timestamp{
						Seconds: 0,
						Nanos:   0,
					},
					DeletedFromTestCaseByUserId: "",
				},
				TestInstructionAttributesList: make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage),
			}

			// Add attributes-data to newly created TestInstruction
			newTestInstructionAttributes := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage{
				BaseAttributeInformation: nil,
				AttributeInformation:     nil,
			}

			// Get all attributes for the immature TestInstruction
			immatureTestInstructionAttributesMap, existsInMap := commandAndRuleEngine.Testcases.ImmatureTestInstructionAttributesMap[matureElement.OriginalElementUuid]
			if existsInMap == false {

				errorId := "8713aa0a-60f0-4892-aaac-2302320e3019"
				err = errors.New(fmt.Sprintf("can't find Immature TestInstruction with UUID '%s' already exist in ImmatureTestInstructionAttributesMap [ErrorID: %s]", matureElement.OriginalElementUuid, errorId))

				fmt.Println(err.Error()) //TODO Send to Error-channel

				return err
			}

			// Loop alla attributes for the ImmatureTestInstruction
			for attributeUuid, attribute := range immatureTestInstructionAttributesMap {

				// Extract the correct DropZone
				dropZoneData, existsInMap := commandAndRuleEngine.Testcases.ImmatureDropZonesDataMap[matureElementToSwapIn.ChosenDropZoneUuid]
				if existsInMap == false {

					errorId := "7a255e6b-c6ac-40f3-9d09-5ef44a6fd7db"
					err = errors.New(fmt.Sprintf("dropZone with UUID '%s' couldn't be found in ImmatureDropZonesDataMap [ErrorID: %s]", matureElementToSwapIn.ChosenDropZoneName, errorId))

					fmt.Println(err.Error()) //TODO Send to Error-channel

					return err
				}

				// Extract attribute from DropZone-data if it exists
				attributeDataFromDropZone, existsInMap := dropZoneData.DropZonePreSetTestInstructionAttributesMap[attributeUuid]
				if existsInMap == true {
					// Attribute exist in DropZone data, so use that data as specified
					switch attributeDataFromDropZone.TestInstructionAttributeType {
					
					// Use the value from the DropZone when adding the attribute to the Model
					case fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionAttributeTypeEnum(fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionInformationMessage_AvailableDropZoneMessage_DropZonePreSetTestInstructionAttributeMessage_USE_DROPZONE_VALUE_FOR_ATTRIBUTE):
						
					// Don't add the attribute to the Model
					case fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionAttributeTypeEnum(fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionInformationMessage_AvailableDropZoneMessage_DropZonePreSetTestInstructionAttributeMessage_REMOVE_ATTRIBUTE_FROM_TESTINSTRUCTION):

					}

				} else {
					// Attribute doesn't exist in DropZone so just att the Attribute to the Model
					newTestInstructionBaseAttributeInformation := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_BaseAttributeInformationMessage{
						TestInstructionAttributeUuid:                  immatureTestInstructionAttributesMap[attributeUuid].TestInstructionAttributeUuid,
						TestInstructionAttributeName:                  immatureTestInstructionAttributesMap[attributeUuid].TestInstructionAttributeName,
						TestInstructionAttributeTypeUuid:              immatureTestInstructionAttributesMap[attributeUuid].TestInstructionAttributeTypeUuid,
						TestInstructionAttributeTypeName:              immatureTestInstructionAttributesMap[attributeUuid].TestInstructionAttributeTypeName,
						TestInstructionAttributeDescription:           immatureTestInstructionAttributesMap[attributeUuid].TestInstructionAttributeDescription,
						TestInstructionAttributeMouseOver:             immatureTestInstructionAttributesMap[attributeUuid].TestInstructionAttributeMouseOver,
						TestInstructionAttributeVisible:               immatureTestInstructionAttributesMap[attributeUuid].TestInstructionAttributeVisible,
						TestInstructionAttributeEnable:                immatureTestInstructionAttributesMap[attributeUuid].TestInstructionAttributeEnable,
						TestInstructionAttributeMandatory:             immatureTestInstructionAttributesMap[attributeUuid].TestInstructionAttributeMandatory,
						TestInstructionAttributeVisibleInTestCaseArea: immatureTestInstructionAttributesMap[attributeUuid].TestInstructionAttributeVisibleInTestCaseArea,
						TestInstructionAttributeIsDeprecated:          immatureTestInstructionAttributesMap[attributeUuid].TestInstructionAttributeIsDeprecated,
						TestInstructionAttributeType:                  fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionAttributeTypeEnum(immatureTestInstructionAttributesMap[attributeUuid].e,
					}

					// TODO handle other types then normal TEXTBOX
					newTestInstructionAttributeInformation := fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage{
						InputTextBoxProperty: &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeInputTextBoxProperty{
							TestInstructionAttributeInputTextBoUuid:  immatureTestInstructionAttributesMap[attributeUuid].TestInstructionAttributeUuid,
							TestInstructionAttributeInputTextBoxName: immatureTestInstructionAttributesMap[attributeUuid].TestInstructionAttributeName,
							TextBoxEditable:                          immatureTestInstructionAttributesMap[attributeUuid].TestInstructionAttributeEnable,
							TextBoxInputMask:                         "",
							TextBoxAttributeTypeUuid:                 immatureTestInstructionAttributesMap[attributeUuid].TestInstructionAttributeTypeUuid,
							TextBoxAttributeTypeName:                 immatureTestInstructionAttributesMap[attributeUuid].TestInstructionAttributeTypeName,
							TextBoxAttributeValue:                    immatureTestInstructionAttributesMap[attributeUuid].TestInstructionAttributeValueAsString,
						},
						InputComboBoxProperty: &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeInputComboBoxProperty{},
						InputFileSelectorProperty: &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeInputFileSelectorProperty{},
					}

				// Create the attribute object with all data
					newTestInstructionAttributes.BaseAttributeInformation = &newTestInstructionBaseAttributeInformation
					newTestInstructionAttributes.AttributeInformation = &newTestInstructionAttributeInformation

					// Save Attribute in TestInstruction
					newMatureTestInstruction.TestInstructionAttributesList[attributeUuid] = &newTestInstructionAttributes


				}
				// Save Mature TestInstruction in TestCase
				currentTestCase.MatureTestInstructionMap[matureElementToSwapIn.FirstElementUuid] = newMatureTestInstruction
			}
		}
	}

	// // Save TestCase back into model
	commandAndRuleEngine.Testcases.TestCases[testCaseUuid] = currentTestCase

	return err
}
