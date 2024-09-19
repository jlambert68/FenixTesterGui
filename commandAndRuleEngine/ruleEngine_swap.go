package commandAndRuleEngine

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	fenixTestInstructions "github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	testInstruction_SendTemplateToThisDomain "github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTemplateToThisDomain"
	testInstruction_SendTemplateToThisDomain_version_1_0 "github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTemplateToThisDomain/version_1_0"
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
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeSwapElement(
	testCaseUuid string,
	elementToSwapOutUuid string,
	immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (
	err error) {

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

	// Execute swap of element
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
		matureElementToSwapIn, err = commandAndRuleEngine.executeTCRuleSwap105(testCaseUuid, elementToBeSwappedIOutUuid, immatureElementToSwapIn)

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
	if err != nil {
		return err
	}

	// Move TestInstructionContainer data to TestCase
	err = commandAndRuleEngine.addTestInstructionContainerDataToTestCaseModel(testCaseUuid, immatureElementToSwapIn, &matureElementToSwapIn)
	if err != nil {
		return err
	}

	return err

}

// Add All TestInstruction-data for the new TestInstruction into the TestCase-model
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) addTestInstructionDataToTestCaseModel(
	testCaseUuid string,
	immatureElementToSwapIn *testCaseModel.ImmatureElementStruct,
	matureElementToSwapIn *testCaseModel.MatureElementStruct) (
	err error) {

	// Extract TestCase to work with
	currentTestCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	if existsInMap == false {

		errorId := "ea7e4f3f-f6c8-4391-a191-116f60c6b5f5"
		err = errors.New(fmt.Sprintf("testCase-model with UUID ('%s') doesn't exist in TestModel-map [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err.Error()) //TODO Send to Error-channel

		return err
	}

	// If 'currentTestCase.MatureTestInstructionMap' then initialize it
	if currentTestCase.MatureTestInstructionMap == nil {
		currentTestCase.MatureTestInstructionMap = make(map[string]testCaseModel.MatureTestInstructionStruct)
	}

	// Verify that TestInstruction doesn't exit in TestInstructionMap
	_, existsInMap = currentTestCase.MatureTestInstructionMap[matureElementToSwapIn.FirstElementUuid]
	if existsInMap == true {

		errorId := "9f659bc5-7088-4bf7-900e-c9e12b4ce36d"
		err = errors.New(fmt.Sprintf("Mature TestInstruction with UUID '%s' already exist in MatureTestInstructionMap in TestCase: %s [ErrorID: %s]", matureElementToSwapIn.FirstElementUuid, testCaseUuid, errorId))

		fmt.Println(err.Error()) //TODO Send to Error-channel

		return err
	}

	// Generate OptionsList for ExecutionsDomains
	var executionDomainsThatCanReceiveDirectTargetedTestInstructionsMap map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
		ExecutionDomainsThatCanReceiveDirectTargetedTestInstructionsMessage
	executionDomainsThatCanReceiveDirectTargetedTestInstructionsMap = *sharedCode.ExecutionDomainsThatCanReceiveDirectTargetedTestInstructionsMapPtr
	var optionsListForExecutionsDomains []string
	for _, tempExecutionDomain := range executionDomainsThatCanReceiveDirectTargetedTestInstructionsMap {
		optionsListForExecutionsDomains = append(optionsListForExecutionsDomains, tempExecutionDomain.GetNameUsedInGui())
	}

	// Used for extracting data to be used
	var tempParentTestInstructionContainerUuid string
	var tempParentTestInstructionContainerMatureUuid string

	// Generate searchable map of

	// Loop over all elements in 'matureElementToSwapIn' and only process TestInstructions , TI or TIx
	for _, matureElement := range matureElementToSwapIn.MatureElementMap {

		// If found a TI or TIx, then process that one
		if matureElement.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION ||
			matureElement.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE {

			// Create timestamp when TestInstruction was added
			createdTimeStamp := timestamppb.Now()

			// Extract parent Original TestInstructionContainer, if it exists
			if immatureElementToSwapIn.ImmatureElementMap[immatureElementToSwapIn.FirstElementUuid].OriginalElementUuid ==
				immatureElementToSwapIn.ImmatureElementMap[immatureElementToSwapIn.FirstElementUuid].ParentElementUuid {
				tempParentTestInstructionContainerUuid = immatureElementToSwapIn.ImmatureElementMap[immatureElementToSwapIn.FirstElementUuid].ParentElementUuid
			} else {
				tempParentTestInstructionContainerUuid = ""
			}

			// Extract parent Mature TestInstructionContainer
			if matureElement.OriginalElementUuid != matureElement.ParentElementUuid {
				tempParentTestInstructionContainerMatureUuid = matureElement.ParentElementUuid
			} else {
				tempParentTestInstructionContainerMatureUuid = ""
			}

			// Create a new Mature TestInstruction to be added
			var newMatureTestInstruction testCaseModel.MatureTestInstructionStruct

			newMatureTestInstruction = testCaseModel.MatureTestInstructionStruct{
				BasicTestInstructionInformation_NonEditableInformation: &fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionInformationMessage_NonEditableBasicInformationMessage{
					DomainUuid:                  commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionsMap[matureElement.OriginalElementUuid].BasicTestInstructionInformation.NonEditableInformation.DomainUuid,
					DomainName:                  commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionsMap[matureElement.OriginalElementUuid].BasicTestInstructionInformation.NonEditableInformation.DomainName,
					ExecutionDomainUuid:         commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionsMap[matureElement.OriginalElementUuid].BasicTestInstructionInformation.NonEditableInformation.ExecutionDomainUuid,
					ExecutionDomainName:         commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionsMap[matureElement.OriginalElementUuid].BasicTestInstructionInformation.NonEditableInformation.ExecutionDomainName,
					TestInstructionOriginalUuid: immatureElementToSwapIn.ImmatureElementMap[matureElement.OriginalElementUuid].OriginalElementUuid,
					TestInstructionOriginalName: immatureElementToSwapIn.ImmatureElementMap[matureElement.OriginalElementUuid].OriginalElementName,
					TestInstructionTypeUuid:     commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionsMap[matureElement.OriginalElementUuid].BasicTestInstructionInformation.NonEditableInformation.TestInstructionTypeUuid,
					TestInstructionTypeName:     commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionsMap[matureElement.OriginalElementUuid].BasicTestInstructionInformation.NonEditableInformation.TestInstructionTypeName,
					Deprecated:                  commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionsMap[matureElement.OriginalElementUuid].BasicTestInstructionInformation.NonEditableInformation.Deprecated,
					MajorVersionNumber:          commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionsMap[matureElement.OriginalElementUuid].BasicTestInstructionInformation.NonEditableInformation.MajorVersionNumber,
					MinorVersionNumber:          commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionsMap[matureElement.OriginalElementUuid].BasicTestInstructionInformation.NonEditableInformation.MinorVersionNumber,
					UpdatedTimeStamp:            commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionsMap[matureElement.OriginalElementUuid].BasicTestInstructionInformation.NonEditableInformation.UpdatedTimeStamp,
					TestInstructionColor:        commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionsMap[matureElement.OriginalElementUuid].BasicTestInstructionInformation.NonEditableInformation.TestInstructionColor,
					TCRuleDeletion:              commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionsMap[matureElement.OriginalElementUuid].BasicTestInstructionInformation.NonEditableInformation.TCRuleDeletion,
					TCRuleSwap:                  commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionsMap[matureElement.OriginalElementUuid].BasicTestInstructionInformation.NonEditableInformation.TCRuleSwap,
				},
				BasicTestInstructionInformation_EditableInformation: &fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionInformationMessage_EditableBasicInformationMessage{
					TestInstructionDescription:   commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionsMap[matureElement.OriginalElementUuid].BasicTestInstructionInformation.EditableInformation.TestInstructionDescription,
					TestInstructionMouseOverText: commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionsMap[matureElement.OriginalElementUuid].BasicTestInstructionInformation.EditableInformation.TestInstructionMouseOverText,
				},
				BasicTestInstructionInformation_InvisibleBasicInformation: &fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionInformationMessage_InvisibleBasicInformationMessage{
					Enabled: commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionsMap[matureElement.OriginalElementUuid].BasicTestInstructionInformation.InvisibleBasicInformation.Enabled,
				},
				MatureBasicTestInstructionInformation: &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_MatureBasicTestInstructionInformationMessage{
					TestCaseUuid:                             testCaseUuid,
					TestInstructionMatureUuid:                matureElement.MatureElementUuid,
					ParentTestInstructionContainerUuid:       tempParentTestInstructionContainerUuid,
					ParentTestInstructionContainerMatureUuid: tempParentTestInstructionContainerMatureUuid,
					ChosenDropZoneUuid:                       matureElementToSwapIn.ChosenDropZoneUuid,
					ChosenDropZoneName:                       matureElementToSwapIn.ChosenDropZoneName,
					ChosenDropZoneColor:                      matureElementToSwapIn.ChosenDropZoneColor,
					TestInstructionType:                      matureElement.TestCaseModelElementType,
				},
				CreatedAndUpdatedInformation: &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_CreatedAndUpdatedInformationMessage{
					AddedToTestCaseTimeStamp:       createdTimeStamp,
					AddedToTestCaseByUserId:        commandAndRuleEngine.Testcases.CurrentUser,
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

			// ************************************

			// Get all attributes for the immature TestInstruction
			immatureTestInstructionAttributesMap, _ := commandAndRuleEngine.Testcases.ImmatureTestInstructionAttributesMap[matureElement.OriginalElementUuid]

			/* Removed because there are TI without any DropZones
			if existsInMap == false {

			errorId := "8713aa0a-60f0-4892-aaac-2302320e3019"
			err = errors.New(fmt.Sprintf("can't find Immature TestInstruction with UUID '%s' in ImmatureTestInstructionAttributesMap [ErrorID: %s]", matureElement.OriginalElementUuid, errorId))

			fmt.Println(err.Error()) //TODO Send to Error-channel

			return err
			}
			*/

			// All create all Attributes for the TestInstruction, both from the original and from Fenix if needed
			var allImmatureTestInstructionAttributesMap map[string]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionAttributesMessage_TestInstructionAttributeMessage
			allImmatureTestInstructionAttributesMap = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionAttributesMessage_TestInstructionAttributeMessage)

			// Loop original attributes and add them
			for tempAttributeUuid, tempAttribute := range immatureTestInstructionAttributesMap {
				// Add Attribute to map
				allImmatureTestInstructionAttributesMap[tempAttributeUuid] = tempAttribute
			}

			// Check if this TestInstruction is a FenixTestInstruction-Add-On-type and should import the Attributes from Fenix
			if commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionsMap[matureElement.OriginalElementUuid].
				BasicTestInstructionInformation.NonEditableInformation.TestInstructionTypeUuid ==
				string(fenixTestInstructions.TestInstructionTypeUUID_FenixSentToUsersDomain_FenixSendTemplateAddOn) {

				// Get all attributes for the immature TestInstruction
				immatureFenixTestInstructionAttributesMap, _ := commandAndRuleEngine.Testcases.
					ImmatureTestInstructionAttributesMap[string(testInstruction_SendTemplateToThisDomain.
					TestInstructionUUID_FenixSentToUsersDomain_SendTemplateToThisDomain)]

				// Loop original Fenix-attributes and add them
				for tempAttributeUuid, tempAttribute := range immatureFenixTestInstructionAttributesMap {

					// Only add certain Attributes to map
					switch tempAttribute.TestInstructionAttributeUuid {

					// These Attributes should be EXCLUDED when adding the attributes
					case string(testInstruction_SendTemplateToThisDomain_version_1_0.
						TestInstructionAttributeUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisExecutionDomainComboBox),
						string(testInstruction_SendTemplateToThisDomain_version_1_0.
							TestInstructionAttributeUUID_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisDomainTextBox),
						string(testInstruction_SendTemplateToThisDomain_version_1_0.
							TestInstructionAttributeUUID_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisExecutionDomainTextBox):

					// Add Attribute to map of all attributes for TestInstruction
					default:
						allImmatureTestInstructionAttributesMap[tempAttributeUuid] = tempAttribute

					}

				}
			}

			// Loop alla attributes for the ImmatureTestInstruction
			for attributeUuid, attribute := range allImmatureTestInstructionAttributesMap {

				// Add attributes-data to newly created TestInstruction
				var newTestInstructionAttributes *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage
				newTestInstructionAttributes = &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage{
					BaseAttributeInformation: nil,
					AttributeInformation:     nil,
				}

				// Extract the correct DropZone among TestInstructions
				dropZoneData, dropZoneExistsInMap := commandAndRuleEngine.Testcases.
					ImmatureDropZonesDataMap[matureElementToSwapIn.ChosenDropZoneUuid]
				/* TODO This part should be remove
				if dropZoneExistsInMap == false {


					errorId := "7a255e6b-c6ac-40f3-9d09-5ef44a6fd7db"
					err = errors.New(fmt.Sprintf("dropZone with UUID '%s' couldn't be found in ImmatureDropZonesDataMap [ErrorID: %s]", matureElementToSwapIn.ChosenDropZoneName, errorId))

					fmt.Println(err.Error()) //TODO Send to Error-channel

					return err
				}

				*/

				// Extract attribute from DropZone-data Among TestInstructions if it exists, but only if there are any DropZones
				var attributeDataFromDropZone *fenixGuiTestCaseBuilderServerGrpcApi.
					ImmatureTestInstructionInformationMessage_AvailableDropZoneMessage_DropZonePreSetTestInstructionAttributeMessage
				var newTestInstructionBaseAttributeInformation *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_BaseAttributeInformationMessage
				if dropZoneExistsInMap == true {
					attributeDataFromDropZone, existsInMap = dropZoneData.DropZonePreSetTestInstructionAttributesMap[attributeUuid]
				}
				if existsInMap == true && dropZoneExistsInMap == true {
					// Attribute exist in DropZone data, so use that data as specified
					switch attributeDataFromDropZone.AttributeActionCommand {

					// Use the value from the DropZone when adding the attribute to the Model
					case fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionInformationMessage_AvailableDropZoneMessage_DropZonePreSetTestInstructionAttributeMessage_USE_DROPZONE_VALUE_FOR_ATTRIBUTE:
						newTestInstructionBaseAttributeInformation = &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_BaseAttributeInformationMessage{
							TestInstructionAttributeUuid:                  attribute.TestInstructionAttributeUuid,
							TestInstructionAttributeName:                  attribute.TestInstructionAttributeName,
							TestInstructionAttributeTypeUuid:              attribute.TestInstructionAttributeTypeUuid,
							TestInstructionAttributeTypeName:              attribute.TestInstructionAttributeTypeName,
							TestInstructionAttributeDescription:           attribute.TestInstructionAttributeDescription,
							TestInstructionAttributeMouseOver:             attribute.TestInstructionAttributeMouseOver,
							TestInstructionAttributeVisible:               attribute.TestInstructionAttributeVisible,
							TestInstructionAttributeEnable:                attribute.TestInstructionAttributeEnable,
							TestInstructionAttributeMandatory:             attribute.TestInstructionAttributeMandatory,
							TestInstructionAttributeVisibleInTestCaseArea: attribute.TestInstructionAttributeVisibleInTestCaseArea,
							TestInstructionAttributeIsDeprecated:          attribute.TestInstructionAttributeIsDeprecated,
							// TODO Change to dynamic type on next row
							TestInstructionAttributeType: fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionAttributeTypeEnum(fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionAttributeTypeEnum_TEXTBOX),
						}

						// TODO handle other types then normal TEXTBOX
						var newTestInstructionAttributeInformation *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage
						newTestInstructionAttributeInformation = &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage{
							InputTextBoxProperty: &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeInputTextBoxProperty{
								TestInstructionAttributeInputTextBoUuid:  attribute.TestInstructionAttributeUuid,
								TestInstructionAttributeInputTextBoxName: attribute.TestInstructionAttributeName,
								TextBoxEditable:                          attribute.TestInstructionAttributeEnable,
								TextBoxInputMask:                         attribute.TestInstructionAttributeInputMask,
								TextBoxAttributeTypeUuid:                 attribute.TestInstructionAttributeTypeUuid,
								TextBoxAttributeTypeName:                 attribute.TestInstructionAttributeTypeName,
								TextBoxAttributeValue:                    attributeDataFromDropZone.AttributeValueAsString,
							},
							InputComboBoxProperty:     &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeInputComboBoxProperty{},
							InputFileSelectorProperty: &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeInputFileSelectorProperty{},
						}

						// Create the attribute object with all data
						newTestInstructionAttributes.BaseAttributeInformation = newTestInstructionBaseAttributeInformation
						newTestInstructionAttributes.AttributeInformation = newTestInstructionAttributeInformation

						// Save Attribute in TestInstruction
						newMatureTestInstruction.TestInstructionAttributesList[attributeUuid] = newTestInstructionAttributes

					// Don't add the attribute to the Model
					case fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionInformationMessage_AvailableDropZoneMessage_DropZonePreSetTestInstructionAttributeMessage_REMOVE_ATTRIBUTE_FROM_TESTINSTRUCTION:
						// Do nothing

					// Shouldn't happen
					default:

						errorId := "4198f151-dc85-4e2d-9282-6a84974e1fb2"
						err = errors.New(fmt.Sprintf("unknown 'attributeDataFromDropZone.TestInstructionAttributeType' %s [ErrorID: %s]", attributeDataFromDropZone.TestInstructionAttributeType, errorId))

						fmt.Println(err.Error()) //TODO Send to Error-channel
						return err

					}

				} else {
					// Attribute doesn't exist in DropZone so just add the Attribute to the Model
					newTestInstructionBaseAttributeInformation = &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_BaseAttributeInformationMessage{
						TestInstructionAttributeUuid:                  attribute.TestInstructionAttributeUuid,
						TestInstructionAttributeName:                  attribute.TestInstructionAttributeName,
						TestInstructionAttributeTypeUuid:              attribute.TestInstructionAttributeTypeUuid,
						TestInstructionAttributeTypeName:              attribute.TestInstructionAttributeTypeName,
						TestInstructionAttributeDescription:           attribute.TestInstructionAttributeDescription,
						TestInstructionAttributeMouseOver:             attribute.TestInstructionAttributeMouseOver,
						TestInstructionAttributeVisible:               attribute.TestInstructionAttributeVisible,
						TestInstructionAttributeEnable:                attribute.TestInstructionAttributeEnable,
						TestInstructionAttributeMandatory:             attribute.TestInstructionAttributeMandatory,
						TestInstructionAttributeVisibleInTestCaseArea: attribute.TestInstructionAttributeVisibleInTestCaseArea,
						TestInstructionAttributeIsDeprecated:          attribute.TestInstructionAttributeIsDeprecated,
						//TestInstructionAttributeType: attribute.TestInstructionAttributeUIType,

						//TestInstructionAttributeType: attribute.TestInstructionAttributeUIType//fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionAttributeTypeEnum(fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionAttributeTypeEnum_TEXTBOX),
					}

					// Define the TestInstructionAttributeInformation
					var newTestInstructionAttributeInformation *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage
					newTestInstructionAttributeInformation = &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage{
						InputTextBoxProperty:             &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeInputTextBoxProperty{},
						InputComboBoxProperty:            &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeInputComboBoxProperty{},
						InputFileSelectorProperty:        &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeInputFileSelectorProperty{},
						ResponseVariableComboBoxProperty: &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeResponseVariableComboBoxProperty{},
					}

					// Set correct GUI-attribute type
					switch attribute.TestInstructionAttributeUIType {
					case "TEXTBOX":
						newTestInstructionBaseAttributeInformation.TestInstructionAttributeType = fenixGuiTestCaseBuilderServerGrpcApi.
							TestInstructionAttributeTypeEnum(fenixGuiTestCaseBuilderServerGrpcApi.
								TestInstructionAttributeTypeEnum_TEXTBOX)

						var inputTextBoxProperty *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeInputTextBoxProperty

						inputTextBoxProperty = &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeInputTextBoxProperty{
							TestInstructionAttributeInputTextBoUuid:  attribute.TestInstructionAttributeUuid,
							TestInstructionAttributeInputTextBoxName: attribute.TestInstructionAttributeName,
							TextBoxEditable:                          attribute.TestInstructionAttributeEnable,
							TextBoxInputMask:                         attribute.TestInstructionAttributeInputMask,
							TextBoxAttributeTypeUuid:                 attribute.TestInstructionAttributeTypeUuid,
							TextBoxAttributeTypeName:                 attribute.TestInstructionAttributeTypeName,
							TextBoxAttributeValue:                    attribute.TestInstructionAttributeValueAsString,
						}

						newTestInstructionAttributeInformation.InputTextBoxProperty = inputTextBoxProperty

						// If it shouldn't be any value in the Textbox then adjust that
						// TODO import TestDataProject to get correct UUID from there instead of using hardcoded value
						if newTestInstructionAttributeInformation.InputTextBoxProperty.TextBoxAttributeValue == "#NO_VALUE#" {
							newTestInstructionAttributeInformation.InputTextBoxProperty.TextBoxAttributeValue = ""
						}

					case "COMBOBOX":
						newTestInstructionBaseAttributeInformation.TestInstructionAttributeType = fenixGuiTestCaseBuilderServerGrpcApi.
							TestInstructionAttributeTypeEnum(fenixGuiTestCaseBuilderServerGrpcApi.
								TestInstructionAttributeTypeEnum_COMBOBOX)

						var comboBoxProperty *fenixGuiTestCaseBuilderServerGrpcApi.
							MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeInputComboBoxProperty

						comboBoxProperty = &fenixGuiTestCaseBuilderServerGrpcApi.
							MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeInputComboBoxProperty{
							TestInstructionAttributeComboBoxUuid: attribute.TestInstructionAttributeUuid,
							TestInstructionAttributeComboBoxName: attribute.TestInstructionAttributeName,
							ComboBoxEditable:                     attribute.TestInstructionAttributeEnable,
							ComboBoxInputMask:                    attribute.TestInstructionAttributeInputMask,
							ComboBoxAttributeTypeUuid:            attribute.TestInstructionAttributeTypeUuid,
							ComboBoxAttributeTypeName:            attribute.TestInstructionAttributeTypeName,
							ComboBoxAttributeValueUuid:           attribute.TestInstructionAttributeValueUuid,
							ComboBoxAttributeValue:               attribute.TestInstructionAttributeValueAsString,
							ComboBoxAllowedValues:                attribute.TestInstructionAttributeComboBoxPredefinedValues,
						}

						newTestInstructionAttributeInformation.InputComboBoxProperty = comboBoxProperty

						// If it shouldn't be any value in the Textbox then adjust that
						if newTestInstructionAttributeInformation.InputComboBoxProperty.ComboBoxAttributeValue == "#NO_VALUE#" {
							newTestInstructionAttributeInformation.InputComboBoxProperty.ComboBoxAttributeValue = ""
						}

					case "FILE_SELECTOR":
						newTestInstructionBaseAttributeInformation.TestInstructionAttributeType = fenixGuiTestCaseBuilderServerGrpcApi.
							TestInstructionAttributeTypeEnum(fenixGuiTestCaseBuilderServerGrpcApi.
								TestInstructionAttributeTypeEnum_FILE_SELECTOR)

					case "FUNCTION_SELECTOR":
						newTestInstructionBaseAttributeInformation.TestInstructionAttributeType = fenixGuiTestCaseBuilderServerGrpcApi.
							TestInstructionAttributeTypeEnum(fenixGuiTestCaseBuilderServerGrpcApi.
								TestInstructionAttributeTypeEnum_FUNCTION_SELECTOR)

					case "RESPONSE_VARIABLE_COMBOBOX":
						newTestInstructionBaseAttributeInformation.TestInstructionAttributeType = fenixGuiTestCaseBuilderServerGrpcApi.
							TestInstructionAttributeTypeEnum(fenixGuiTestCaseBuilderServerGrpcApi.
								TestInstructionAttributeTypeEnum_RESPONSE_VARIABLE_COMBOBOX)
						var responseVariableComboBoxProperty *fenixGuiTestCaseBuilderServerGrpcApi.
							MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeResponseVariableComboBoxProperty

						responseVariableComboBoxProperty = &fenixGuiTestCaseBuilderServerGrpcApi.
							MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeResponseVariableComboBoxProperty{
							TestInstructionAttributeResponseVariableComboBoxUuid: attribute.TestInstructionAttributeUuid,
							TestInstructionAttributeResponseVariableComboBoxName: attribute.TestInstructionAttributeName,
							ComboBoxAttributeTypeUuid:                            attribute.TestInstructionAttributeTypeUuid,
							ComboBoxAttributeTypeName:                            attribute.TestInstructionAttributeTypeName,
							AllowedResponseVariableType:                          []string{},
							ChosenResponseVariableTypeUuid:                       "",
							ChosenResponseVariableTypeName:                       "",
							ComboBoxAttributeValueAsString:                       "",
							ComboBoxResponseVariableInputMask:                    attribute.TestInstructionAttributeInputMask,
						}

						newTestInstructionAttributeInformation.ResponseVariableComboBoxProperty = responseVariableComboBoxProperty

						if newTestInstructionAttributeInformation.ResponseVariableComboBoxProperty.ComboBoxAttributeValueAsString == "#NO_VALUE#" {
							newTestInstructionAttributeInformation.ResponseVariableComboBoxProperty.ComboBoxAttributeValueAsString = ""
						}

					case "TESTCASE_BUILDER_SERVER_INJECTED_COMBOBOX":
						newTestInstructionBaseAttributeInformation.TestInstructionAttributeType = fenixGuiTestCaseBuilderServerGrpcApi.
							TestInstructionAttributeTypeEnum(fenixGuiTestCaseBuilderServerGrpcApi.
								TestInstructionAttributeTypeEnum_COMBOBOX)

						var responseVariableComboBoxProperty *fenixGuiTestCaseBuilderServerGrpcApi.
							MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeInputComboBoxProperty

						responseVariableComboBoxProperty = &fenixGuiTestCaseBuilderServerGrpcApi.
							MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeInputComboBoxProperty{
							TestInstructionAttributeComboBoxUuid: attribute.TestInstructionAttributeUuid,
							TestInstructionAttributeComboBoxName: attribute.TestInstructionAttributeName,
							ComboBoxEditable:                     false,
							ComboBoxInputMask:                    attribute.TestInstructionAttributeInputMask,
							ComboBoxAttributeTypeUuid:            attribute.TestInstructionAttributeTypeUuid,
							ComboBoxAttributeTypeName:            attribute.TestInstructionAttributeTypeName,
							ComboBoxAttributeValueUuid:           "",
							ComboBoxAttributeValue:               "",
							ComboBoxAllowedValues:                optionsListForExecutionsDomains,
						}

						newTestInstructionAttributeInformation.InputComboBoxProperty = responseVariableComboBoxProperty

						if newTestInstructionAttributeInformation.ResponseVariableComboBoxProperty.ComboBoxAttributeValueAsString == "#NO_VALUE#" {
							newTestInstructionAttributeInformation.ResponseVariableComboBoxProperty.ComboBoxAttributeValueAsString = ""
						}

					case "FENIX_OWNED_TESTER_GUI_INJECTED_COMBOBOX":
						newTestInstructionBaseAttributeInformation.TestInstructionAttributeType = fenixGuiTestCaseBuilderServerGrpcApi.
							TestInstructionAttributeTypeEnum(fenixGuiTestCaseBuilderServerGrpcApi.
								TestInstructionAttributeTypeEnum_COMBOBOX)

						var comboBoxAllowedValues []string

						// Check if this is a Templates Combobox served by Fenix shared TestInstructions
						if attribute.TestInstructionAttributeUuid == string(testInstruction_SendTemplateToThisDomain_version_1_0.
							TestInstructionAttributeUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateComboBox) {

							// Extract Templates
							for _, templateGitHubFile := range currentTestCase.ImportedTemplateFilesFromGitHub {

								// Only add File if it is not already in use
								if templateGitHubFile.FileIsUsedInTestCase == false {

									comboBoxAllowedValues = append(comboBoxAllowedValues,
										fmt.Sprintf("%s [%s]",
											templateGitHubFile.Name,
											templateGitHubFile.FileHash[0:8]))
								}
							}

						} else {
							comboBoxAllowedValues = []string{}
						}

						var responseVariableComboBoxProperty *fenixGuiTestCaseBuilderServerGrpcApi.
							MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeInputComboBoxProperty

						responseVariableComboBoxProperty = &fenixGuiTestCaseBuilderServerGrpcApi.
							MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeInputComboBoxProperty{
							TestInstructionAttributeComboBoxUuid: attribute.TestInstructionAttributeUuid,
							TestInstructionAttributeComboBoxName: attribute.TestInstructionAttributeName,
							ComboBoxEditable:                     false,
							ComboBoxInputMask:                    attribute.TestInstructionAttributeInputMask,
							ComboBoxAttributeTypeUuid:            attribute.TestInstructionAttributeTypeUuid,
							ComboBoxAttributeTypeName:            attribute.TestInstructionAttributeTypeName,
							ComboBoxAttributeValueUuid:           "",
							ComboBoxAttributeValue:               "",
							ComboBoxAllowedValues:                comboBoxAllowedValues,
						}

						newTestInstructionAttributeInformation.InputComboBoxProperty = responseVariableComboBoxProperty

						if newTestInstructionAttributeInformation.ResponseVariableComboBoxProperty.ComboBoxAttributeValueAsString == "#NO_VALUE#" {
							newTestInstructionAttributeInformation.ResponseVariableComboBoxProperty.ComboBoxAttributeValueAsString = ""
						}

					default:
						errorId := "117d964e-860f-4497-b367-02c041553615"
						err = errors.New(fmt.Sprintf("unknown 'attribute.TestInstructionAttributeUIType' %s [ErrorID: %s]", attribute.TestInstructionAttributeUIType, errorId))

						fmt.Println(err.Error()) //TODO Send to Error-channel
						return err

					}

					// Create the attribute object with all data
					newTestInstructionAttributes.BaseAttributeInformation = newTestInstructionBaseAttributeInformation
					newTestInstructionAttributes.AttributeInformation = newTestInstructionAttributeInformation

					// Save Attribute in TestInstruction
					newMatureTestInstruction.TestInstructionAttributesList[attributeUuid] = newTestInstructionAttributes

				}
				// Save Mature TestInstruction in TestCase
				currentTestCase.MatureTestInstructionMap[matureElement.MatureElementUuid] = newMatureTestInstruction
			}
		}
	}

	// // Save TestCase back into model
	commandAndRuleEngine.Testcases.TestCases[testCaseUuid] = currentTestCase

	return err
}

// Add All TestInstructionContainer-data for the new TestInstructionContainer into the TestCase-model
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) addTestInstructionContainerDataToTestCaseModel(testCaseUuid string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct, matureElementToSwapIn *testCaseModel.MatureElementStruct) (err error) {

	// Extract TestCase to work with
	currentTestCase, existsInMap := commandAndRuleEngine.Testcases.TestCases[testCaseUuid]
	if existsInMap == false {

		errorId := "bb0490d1-051a-468b-a8de-b0fd5299a45e"
		err = errors.New(fmt.Sprintf("testCase-model with UUID ('%s') doesn't exist in TestModel-map [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err.Error()) //TODO Send to Error-channel

		return err
	}

	// If 'currentTestCase.MatureTestInstructionContainerMap' then initialize it
	if currentTestCase.MatureTestInstructionContainerMap == nil {
		currentTestCase.MatureTestInstructionContainerMap = make(map[string]testCaseModel.MatureTestInstructionContainerStruct)
	}

	// Verify that TestInstructionContainer doesn't exit in TestInstructionContainerMap
	_, existsInMap = currentTestCase.MatureTestInstructionContainerMap[matureElementToSwapIn.FirstElementUuid]
	if existsInMap == true {

		errorId := "a5fc8c15-9788-45f6-b76c-08ac89a54f1d"
		err = errors.New(fmt.Sprintf("Mature TestInstructionContainer with UUID '%s' already exist in MatureTestInstructionContainerMap in TestCase: %s [ErrorID: %s]", matureElementToSwapIn.FirstElementUuid, testCaseUuid, errorId))

		fmt.Println(err.Error()) //TODO Send to Error-channel

		return err
	}

	// Loop over all elements in 'matureElementToSwapIn' and only process TestInstructionContainers , TIC or TICx
	for _, matureElement := range matureElementToSwapIn.MatureElementMap {

		// If found a TIC or TICx, then process that one
		if matureElement.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER ||
			matureElement.TestCaseModelElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE {

			// Create timestamp when TestInstructionContainer was added
			createdTimeStamp := timestamppb.Now()

			// Create a new Mature TestInstructionContainer to be added
			var newMatureTestInstructionContainer testCaseModel.MatureTestInstructionContainerStruct

			newMatureTestInstructionContainer = testCaseModel.MatureTestInstructionContainerStruct{
				NonEditableInformation: &fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionContainerInformationMessage_NonEditableBasicInformationMessage{
					DomainUuid:                       commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionContainersMap[matureElement.OriginalElementUuid].BasicTestInstructionContainerInformation.NonEditableInformation.DomainUuid,
					DomainName:                       commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionContainersMap[matureElement.OriginalElementUuid].BasicTestInstructionContainerInformation.NonEditableInformation.DomainName,
					TestInstructionContainerUuid:     commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionContainersMap[matureElement.OriginalElementUuid].BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerUuid,
					TestInstructionContainerName:     commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionContainersMap[matureElement.OriginalElementUuid].BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerName,
					TestInstructionContainerTypeUuid: commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionContainersMap[matureElement.OriginalElementUuid].BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerTypeUuid,
					TestInstructionContainerTypeName: commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionContainersMap[matureElement.OriginalElementUuid].BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerTypeName,
					Deprecated:                       commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionContainersMap[matureElement.OriginalElementUuid].BasicTestInstructionContainerInformation.NonEditableInformation.Deprecated,
					MajorVersionNumber:               commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionContainersMap[matureElement.OriginalElementUuid].BasicTestInstructionContainerInformation.NonEditableInformation.MajorVersionNumber,
					MinorVersionNumber:               commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionContainersMap[matureElement.OriginalElementUuid].BasicTestInstructionContainerInformation.NonEditableInformation.MinorVersionNumber,
					UpdatedTimeStamp:                 commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionContainersMap[matureElement.OriginalElementUuid].BasicTestInstructionContainerInformation.NonEditableInformation.UpdatedTimeStamp,
					TestInstructionContainerColor:    commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionContainersMap[matureElement.OriginalElementUuid].BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerColor,
					TCRuleDeletion:                   commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionContainersMap[matureElement.OriginalElementUuid].BasicTestInstructionContainerInformation.NonEditableInformation.TCRuleDeletion,
					TCRuleSwap:                       commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionContainersMap[matureElement.OriginalElementUuid].BasicTestInstructionContainerInformation.NonEditableInformation.TCRuleSwap,
				},
				EditableInformation: &fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionContainerInformationMessage_EditableBasicInformationMessage{
					TestInstructionContainerDescription:   commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionContainersMap[matureElement.OriginalElementUuid].BasicTestInstructionContainerInformation.EditableInformation.TestInstructionContainerDescription,
					TestInstructionContainerMouseOverText: commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionContainersMap[matureElement.OriginalElementUuid].BasicTestInstructionContainerInformation.EditableInformation.TestInstructionContainerMouseOverText,
				},
				InvisibleBasicInformation: &fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionContainerInformationMessage_InvisibleBasicInformationMessage{
					Enabled: commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionContainersMap[matureElement.OriginalElementUuid].BasicTestInstructionContainerInformation.InvisibleBasicInformation.Enabled,
				},
				EditableTestInstructionContainerAttributes: &fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionContainerInformationMessage_EditableTestInstructionContainerAttributesMessage{
					TestInstructionContainerExecutionType: commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionContainersMap[matureElement.OriginalElementUuid].BasicTestInstructionContainerInformation.EditableTestInstructionContainerAttributes.TestInstructionContainerExecutionType,
				},
				MatureTestInstructionContainerInformation: &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionContainerInformationMessage_MatureTestInstructionContainerInformationMessage{
					TestCaseUuid:                       testCaseUuid,
					TestInstructionContainerMatureUuid: matureElement.MatureElementUuid,
				},
				CreatedAndUpdatedInformation: &fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionContainerInformationMessage_CreatedAndUpdatedInformationMessage{
					AddedToTestCaseTimeStamp:       createdTimeStamp,
					AddedToTestCaseByUserId:        commandAndRuleEngine.Testcases.CurrentUser,
					LastUpdatedInTestCaseTimeStamp: createdTimeStamp,
					LastUpdatedInTestCaseByUserId:  "",
					DeletedFromTestCaseTimeStamp: &timestamppb.Timestamp{
						Seconds: 0,
						Nanos:   0,
					},
					DeletedFromTestCaseByUserId: "",
				},
			}

			// ************************************

			// Save Mature TestInstructionContainer in TestCase
			currentTestCase.MatureTestInstructionContainerMap[matureElement.MatureElementUuid] = newMatureTestInstructionContainer
		}
	}

	// // Save TestCase back into model
	commandAndRuleEngine.Testcases.TestCases[testCaseUuid] = currentTestCase

	return err
}
