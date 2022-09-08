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

		}
	}

	return err
}
