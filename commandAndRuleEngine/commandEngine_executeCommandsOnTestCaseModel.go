package commandAndRuleEngine

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"

	//"errors"
	uuidGenerator "github.com/google/uuid"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
)

// TestCaseCommandTypeEnum_NEW_TESTCASE
// Used for creating a new TestCase-Model to be used within a new TestCase
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_NewTestCaseModel() (testCaseUuid string, err error) {

	// Create new B0-Bind
	b0Bond := commandAndRuleEngine.createNewBondB0Element()

	// Initiate a TestCaseModel
	newTestCaseModel := testCaseModel.TestCaseModelStruct{
		LastLoadedTestCaseModelGRPCMessage: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage{},
		FirstElementUuid:                   "",
		TestCaseModelMap:                   nil,
	}

	// Initiate TestCaseModel-map
	newTestCaseModel.TestCaseModelMap = make(map[string]testCaseModel.MatureTestCaseModelElementStruct)

	// Add B0-bond to the TestCaseModel-map
	newTestCaseModel.TestCaseModelMap[b0Bond.MatureElementUuid] = testCaseModel.MatureTestCaseModelElementStruct{
		MatureTestCaseModelElementMessage:  b0Bond,
		MatureTestCaseModelElementMetaData: testCaseModel.MatureTestCaseModelElementMetaDataStruct{},
	}

	// Set the B0-bond as first element in TestCaseModel
	newTestCaseModel.FirstElementUuid = b0Bond.MatureElementUuid

	// Create a new Command
	newCommandEntry := fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage_TestCaseModelCommandMessage{
		TestCaseCommandType:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_NEW_TESTCASE,
		TestCaseCommandName:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_name[int32(fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_NEW_TESTCASE)],
		FirstParameter:           testCaseModel.NotApplicable,
		SecondParameter:          testCaseModel.NotApplicable,
		UserIdOnComputer:         sharedCode.CurrentUserIdLogedInOnComputer,
		GCPAuthenticatedUser:     sharedCode.CurrentUserAuthenticatedTowardsGCP,
		CommandExecutedTimeStamp: timestamppb.Now(),
	}

	// Add command to command stack
	newTestCaseModel.CommandStack = append(newTestCaseModel.CommandStack, newCommandEntry)

	// Generate new TestCase-UUID
	testCaseUuid = uuidGenerator.New().String()

	// If TestCasesMapPtr-map is not initialized then do that
	if commandAndRuleEngine.Testcases.TestCasesMapPtr == nil {

		var tempTestCasesMap map[string]*testCaseModel.TestCaseModelStruct
		tempTestCasesMap = make(map[string]*testCaseModel.TestCaseModelStruct)

		commandAndRuleEngine.Testcases.TestCasesMapPtr = &tempTestCasesMap
	}

	// TODO Add dropdown for user to chose among available Domains in available building blocks
	// Add BasicTestCaseInformation
	localTestCaseMessageStruct := testCaseModel.LocalTestCaseMessageStruct{
		BasicTestCaseInformationMessageNoneEditableInformation: fenixGuiTestCaseBuilderServerGrpcApi.BasicTestCaseInformationMessage_NonEditableBasicInformationMessage{
			TestCaseUuid:    testCaseUuid,
			DomainUuid:      "",
			DomainName:      "",
			TestCaseVersion: 1,
		},
		BasicTestCaseInformationMessageEditableInformation: fenixGuiTestCaseBuilderServerGrpcApi.BasicTestCaseInformationMessage_EditableBasicInformationMessage{
			TestCaseName:        "<New TestCase>",
			TestCaseDescription: "",
		},
		CreatedAndUpdatedInformation: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseBasicInformationMessage_CreatedAndUpdatedInformationMessage{
			AddedToTestCaseTimeStamp: &timestamppb.Timestamp{
				Seconds: 0,
				Nanos:   0,
			},
			AddedToTestCaseByUserId: "",
			LastUpdatedInTestCaseTimeStamp: &timestamppb.Timestamp{
				Seconds: 0,
				Nanos:   0,
			},
			LastUpdatedInTestCaseByUserId: "",
			DeletedFromTestCaseTimeStamp: &timestamppb.Timestamp{
				Seconds: 0,
				Nanos:   0,
			},
			DeletedFromTestCaseByUserId: "",
		},
	}
	newTestCaseModel.LocalTestCaseMessage = localTestCaseMessageStruct

	// Set that this is a new TestCase
	newTestCaseModel.ThisIsANewTestCase = true

	// Get the TestCasesModelMap
	var testCasesModelMap map[string]*testCaseModel.TestCaseModelStruct
	testCasesModelMap = *commandAndRuleEngine.Testcases.TestCasesMapPtr

	// Add the TestCaseModel into map of all TestCaseModels
	testCasesModelMap[testCaseUuid] = &newTestCaseModel

	// Add command Textual representations to Textual Representation Stack
	textualRepresentationSimple, textualRepresentationComplex, textualRepresentationExtended, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)

	// Initialize AttributesList
	var tempAttributeStructSliceReference []*testCaseModel.AttributeStruct
	tempAttributeStructSliceReference = make([]*testCaseModel.AttributeStruct, 0)
	var tempAttributesList *testCaseModel.AttributeStructSliceReferenceType
	tempAttributesList = (*testCaseModel.AttributeStructSliceReferenceType)(&tempAttributeStructSliceReference)
	newTestCaseModel.AttributesList = tempAttributesList

	if err == nil {
		newTestCaseModel.TextualTestCaseRepresentationSimpleStack = append(newTestCaseModel.TextualTestCaseRepresentationSimpleStack, textualRepresentationSimple)
		newTestCaseModel.TextualTestCaseRepresentationComplexStack = append(newTestCaseModel.TextualTestCaseRepresentationComplexStack, textualRepresentationComplex)
		newTestCaseModel.TextualTestCaseRepresentationExtendedStack = append(newTestCaseModel.TextualTestCaseRepresentationExtendedStack, textualRepresentationExtended)

		// Add the TestCaseModel back into map of all TestCaseModels
		testCasesModelMap[testCaseUuid] = &newTestCaseModel

	}

	return testCaseUuid, err

}

// TestCaseCommandTypeEnum_REMOVE_ELEMENT
// Used for Deleting an element from a TestCaseModel that is used within a TestCase
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_DeleteElementFromTestCaseModel(testCaseUuid string, elementId string) (err error) {

	var existsInMap bool

	// Try to Delete element
	err = commandAndRuleEngine.executeDeleteElement(testCaseUuid, elementId)

	// Exit if the element couldn't be deleted
	if err != nil {
		return err
	}

	// Create a new Command
	newCommandEntry := fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage_TestCaseModelCommandMessage{
		TestCaseCommandType:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_REMOVE_ELEMENT,
		TestCaseCommandName:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_name[int32(fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_REMOVE_ELEMENT)],
		FirstParameter:           elementId,
		SecondParameter:          testCaseModel.NotApplicable,
		UserIdOnComputer:         sharedCode.CurrentUserIdLogedInOnComputer,
		GCPAuthenticatedUser:     sharedCode.CurrentUserAuthenticatedTowardsGCP,
		CommandExecutedTimeStamp: timestamppb.Now(),
	}

	// Extract the TestCaseModel
	var currentTestCaseModelPtr *testCaseModel.TestCaseModelStruct
	var testCasesMap map[string]*testCaseModel.TestCaseModelStruct
	testCasesMap = *commandAndRuleEngine.Testcases.TestCasesMapPtr
	currentTestCaseModelPtr, existsInMap = testCasesMap[testCaseUuid]

	//currentTestCaseModel, existsInMap := commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid]
	if existsInMap == false {
		errorId := "9e42e135-e5c3-479c-8a09-0e33213a68d1"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCasesMapPtr [ErrorID: %s]", testCaseUuid, errorId))
	}

	// Clear 'clicked element'
	currentTestCaseModelPtr.CurrentSelectedTestCaseElement = testCaseModel.CurrentSelectedTestCaseElementStruct{
		CurrentSelectedTestCaseElementUuid: "",
		CurrentSelectedTestCaseElementName: "",
	}

	// Add command to command stack
	currentTestCaseModelPtr.CommandStack = append(currentTestCaseModelPtr.CommandStack, newCommandEntry)

	// Add the TestCaseModel back into map of all TestCaseModels
	//commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid] = currentTestCaseModel

	// Add command Textual representations to Textual Representation Stack
	textualRepresentationSimple, textualRepresentationComplex, textualRepresentationExtended, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)

	if err == nil {
		currentTestCaseModelPtr.TextualTestCaseRepresentationSimpleStack = append(currentTestCaseModelPtr.TextualTestCaseRepresentationSimpleStack, textualRepresentationSimple)
		currentTestCaseModelPtr.TextualTestCaseRepresentationComplexStack = append(currentTestCaseModelPtr.TextualTestCaseRepresentationComplexStack, textualRepresentationComplex)
		currentTestCaseModelPtr.TextualTestCaseRepresentationExtendedStack = append(currentTestCaseModelPtr.TextualTestCaseRepresentationExtendedStack, textualRepresentationExtended)

		// Add the TestCaseModel back into map of all TestCaseModels
		//commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid] = currentTestCaseModelPtr
	}

	return err
}

// TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_NEW_ELEMENT
// Used for Swapping out an element, and in an element structure, from a TestCaseModel that is used within a TestCase
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_SwapOutElemenAndInNewElementInTestCaseModel(testCaseUuid string, elementToSwapOutUuid string, immatureElementToSwapIn *testCaseModel.ImmatureElementStruct) (err error) {

	var existsInMap bool

	// Try to Swap out element
	err = commandAndRuleEngine.executeSwapElement(testCaseUuid, elementToSwapOutUuid, immatureElementToSwapIn)

	// Exit if the element couldn't be deleted
	if err != nil {
		return err
	}

	// Create a new Command
	newCommandEntry := fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage_TestCaseModelCommandMessage{
		TestCaseCommandType:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_NEW_ELEMENT,
		TestCaseCommandName:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_name[int32(fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_NEW_ELEMENT)],
		FirstParameter:           elementToSwapOutUuid,
		SecondParameter:          immatureElementToSwapIn.FirstElementUuid,
		UserIdOnComputer:         sharedCode.CurrentUserIdLogedInOnComputer,
		GCPAuthenticatedUser:     sharedCode.CurrentUserAuthenticatedTowardsGCP,
		CommandExecutedTimeStamp: timestamppb.Now(),
	}

	// Extract the TestCasesModel
	var testCasesMap map[string]*testCaseModel.TestCaseModelStruct
	testCasesMap = *commandAndRuleEngine.Testcases.TestCasesMapPtr

	var currentTestCaseModelPtr *testCaseModel.TestCaseModelStruct
	currentTestCaseModelPtr, existsInMap = testCasesMap[testCaseUuid]
	if existsInMap == false {
		errorId := "73cf671c-79e7-4a5e-8f42-d39cd86d94c9"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCasesMapPtr [ErrorID: %s]", testCaseUuid, errorId))
	}

	// Add command to command stack
	currentTestCaseModelPtr.CommandStack = append(currentTestCaseModelPtr.CommandStack, newCommandEntry)

	// Add the TestCaseModel back into map of all TestCaseModels
	// commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid] = currentTestCaseModelPtr

	// Add command Textual representations to Textual Representation Stack
	textualRepresentationSimple, textualRepresentationComplex, textualRepresentationExtended, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)

	// If no errors then add Simple and Complex Textual Representation to their stacks
	if err == nil {
		currentTestCaseModelPtr.TextualTestCaseRepresentationSimpleStack = append(currentTestCaseModelPtr.TextualTestCaseRepresentationSimpleStack, textualRepresentationSimple)
		currentTestCaseModelPtr.TextualTestCaseRepresentationComplexStack = append(currentTestCaseModelPtr.TextualTestCaseRepresentationComplexStack, textualRepresentationComplex)
		currentTestCaseModelPtr.TextualTestCaseRepresentationExtendedStack = append(currentTestCaseModelPtr.TextualTestCaseRepresentationExtendedStack, textualRepresentationExtended)

		// Add the TestCaseModel back into map of all TestCaseModels
		//commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid] = currentTestCaseModelPtr
	}

	return err
}

// TestCaseCommandTypeEnum_COPY_ELEMENT
// Used for copying an element  in a TestCaseModel that is used within a TestCase
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_CopyElementInTestCaseModel(testCaseUuid string, elementIdToCopy string) (err error) {

	var existsInMap bool

	// Try to Copy element to Copy Buffer
	err = commandAndRuleEngine.executeCopyElement(testCaseUuid, elementIdToCopy)

	// Exit if the element couldn't be deleted
	if err != nil {
		return err
	}

	// Create a new Command
	newCommandEntry := fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage_TestCaseModelCommandMessage{
		TestCaseCommandType:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_COPY_ELEMENT,
		TestCaseCommandName:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_name[int32(fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_COPY_ELEMENT)],
		FirstParameter:           elementIdToCopy,
		SecondParameter:          testCaseModel.NotApplicable,
		UserIdOnComputer:         sharedCode.CurrentUserIdLogedInOnComputer,
		GCPAuthenticatedUser:     sharedCode.CurrentUserAuthenticatedTowardsGCP,
		CommandExecutedTimeStamp: timestamppb.Now(),
	}

	// Get the TestCasesModel
	var testCasesMap map[string]*testCaseModel.TestCaseModelStruct
	testCasesMap = *commandAndRuleEngine.Testcases.TestCasesMapPtr

	// Extract the TestCaseModel
	currentTestCaseModelPtr, existsInMap := testCasesMap[testCaseUuid]
	if existsInMap == false {
		errorId := "2d6af5bd-5a1b-4cc0-b3e7-da21b5928c4f"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCasesMapPtr [ErrorID: %s]", testCaseUuid, errorId))

		return err
	}

	// Add command to command stack
	currentTestCaseModelPtr.CommandStack = append(currentTestCaseModelPtr.CommandStack, newCommandEntry)

	// If no errors then add the TestCaseModel back into map of all TestCaseModels
	//if err == nil {
	//	commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid] = currentTestCaseModelPtr
	//}

	return err
}

// TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_COPY_BUFFER_ELEMENT
// Used for Swapping in an element from Copy Buffer in a TestCaseModel that is used within a TestCase
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_SwapInElementFromCopyBufferInTestCaseModel(testCaseUuid string, elementIdToBeReplacedByCopyBuffer string) (err error) {

	// Try to Swap Element From Copy Buffer
	err = commandAndRuleEngine.executeSwapElementForCopyBuffer(testCaseUuid, elementIdToBeReplacedByCopyBuffer)

	// Exit if the element couldn't be deleted
	if err != nil {
		return err
	}

	// Create a new Command
	newCommandEntry := fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage_TestCaseModelCommandMessage{
		TestCaseCommandType:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_COPY_BUFFER_ELEMENT,
		TestCaseCommandName:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_name[int32(fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_COPY_BUFFER_ELEMENT)],
		FirstParameter:           elementIdToBeReplacedByCopyBuffer,
		SecondParameter:          testCaseModel.NotApplicable,
		UserIdOnComputer:         sharedCode.CurrentUserIdLogedInOnComputer,
		GCPAuthenticatedUser:     sharedCode.CurrentUserAuthenticatedTowardsGCP,
		CommandExecutedTimeStamp: timestamppb.Now(),
	}

	// Get the TestCasesModel
	var testCasesMap map[string]*testCaseModel.TestCaseModelStruct
	testCasesMap = *commandAndRuleEngine.Testcases.TestCasesMapPtr

	// Extract the TestCaseModel
	currentTestCaseModelPtr, existsInMap := testCasesMap[testCaseUuid]
	if existsInMap == false {
		errorId := "10ef5496-d92e-4e35-af41-e16c51c7df71"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCasesMapPtr [ErrorID: %s]", testCaseUuid, errorId))

		return err
	}

	// Add command to command stack
	currentTestCaseModelPtr.CommandStack = append(currentTestCaseModelPtr.CommandStack, newCommandEntry)

	// Indicate that Cut command has been initiated
	currentTestCaseModelPtr.CutCommandInitiated = true

	// If no errors then add the TestCaseModel back into map of all TestCaseModels
	//if err == nil {
	//	commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid] = currentTestCaseModelPtr
	//}

	return err
}

// TestCaseCommandTypeEnum_CUT_ELEMENT
// Used for cutting an element in a TestCaseModel that is used within a TestCase
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_CutElementInTestCaseModel(testCaseUuid string, elementIdToCut string) (err error) {

	var existsInMap bool

	// Try to Copy element to Cut Buffer
	err = commandAndRuleEngine.executeCutElement(testCaseUuid, elementIdToCut)

	// Exit if the element couldn't be deleted
	if err != nil {
		return err
	}

	// Create a new Command
	newCommandEntry := fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage_TestCaseModelCommandMessage{
		TestCaseCommandType:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_CUT_ELEMENT,
		TestCaseCommandName:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_name[int32(fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_CUT_ELEMENT)],
		FirstParameter:           elementIdToCut,
		SecondParameter:          testCaseModel.NotApplicable,
		UserIdOnComputer:         sharedCode.CurrentUserIdLogedInOnComputer,
		GCPAuthenticatedUser:     sharedCode.CurrentUserAuthenticatedTowardsGCP,
		CommandExecutedTimeStamp: timestamppb.Now(),
	}

	// Get the TestCasesModel
	var testCasesMap map[string]*testCaseModel.TestCaseModelStruct
	testCasesMap = *commandAndRuleEngine.Testcases.TestCasesMapPtr

	// Extract the TestCaseModel
	var currentTestCaseModelPtr *testCaseModel.TestCaseModelStruct
	currentTestCaseModelPtr, existsInMap = testCasesMap[testCaseUuid]
	if existsInMap == false {
		errorId := "dc1cd5d3-e809-4465-aeda-cdf6ec44070f"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCasesMapPtr [ErrorID: %s]", testCaseUuid, errorId))

		return err
	}

	// Add command to command stack
	currentTestCaseModelPtr.CommandStack = append(currentTestCaseModelPtr.CommandStack, newCommandEntry)

	// If no errors then add the TestCaseModel back into map of all TestCaseModels
	//if err == nil {
	//	commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid] = currentTestCaseModel
	//}

	return err
}

// TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_CUT_BUFFER_ELEMENT
// Used for Swapping in an element from Cut opy Buffer in a TestCaseModel that is used within a TestCase
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) executeCommandOnTestCaseModel_SwapInElementFromCutBufferInTestCaseModel(testCaseUuid string, uuidToReplacedByCutBufferContent string) (err error) {

	var existsInMap bool

	// Try to Swap Element From Cut Buffer
	err = commandAndRuleEngine.executeSwapElementFromCutBuffer(testCaseUuid, uuidToReplacedByCutBufferContent, nil)

	// Exit if the element couldn't be Swapped in from Cut
	if err != nil {
		return err
	}

	// Create a new Command
	newCommandEntry := fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage_TestCaseModelCommandMessage{
		TestCaseCommandType:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_CUT_BUFFER_ELEMENT,
		TestCaseCommandName:      fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_name[int32(fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum_SWAP_OUT_ELEMENT_FOR_CUT_BUFFER_ELEMENT)],
		FirstParameter:           uuidToReplacedByCutBufferContent,
		SecondParameter:          testCaseModel.NotApplicable,
		UserIdOnComputer:         sharedCode.CurrentUserIdLogedInOnComputer,
		GCPAuthenticatedUser:     sharedCode.CurrentUserAuthenticatedTowardsGCP,
		CommandExecutedTimeStamp: timestamppb.Now(),
	}

	// Get the TestCasesModel
	var testCasesMap map[string]*testCaseModel.TestCaseModelStruct
	testCasesMap = *commandAndRuleEngine.Testcases.TestCasesMapPtr

	// Extract the TestCaseModel
	var currentTestCaseModelPtr *testCaseModel.TestCaseModelStruct
	currentTestCaseModelPtr, existsInMap = testCasesMap[testCaseUuid]
	if existsInMap == false {
		errorId := "e1f7b09a-1867-4c0d-a02a-2b513788d711"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCasesMapPtr [ErrorID: %s]", testCaseUuid, errorId))

		return err
	}

	// Add command to command stack
	currentTestCaseModelPtr.CommandStack = append(currentTestCaseModelPtr.CommandStack, newCommandEntry)

	// Indicate that Cut command has been finished, or is not active anymore
	currentTestCaseModelPtr.CutCommandInitiated = false

	// Clear Cut Buffer
	newEmptyCutBufferContet := testCaseModel.MatureElementStruct{
		FirstElementUuid: "",
		MatureElementMap: nil,
	}
	currentTestCaseModelPtr.CutBuffer = newEmptyCutBufferContet

	// If no errors then add the TestCaseModel back into map of all TestCaseModels
	//if err == nil {
	//	commandAndRuleEngine.Testcases.TestCasesMapPtr[testCaseUuid] = currentTestCaseModel
	//}

	return err
}
