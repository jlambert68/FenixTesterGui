package testSuitesModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/soundEngine"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"github.com/jinzhu/copier"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
)

// SaveTestSuite
// Send TestSuite to TestCaseBuilderServer to saved in the Database
func (testSuiteModel *TestSuiteModelStruct) SaveTestSuite() (err error) {

	var isTestSuiteChanged bool
	var mandatoryFieldsHaveValues bool
	var mandatoryFieldsHaveValuesNotificationText string
	var valuesToBeHashed []string
	var messageHash string

	// Check if TestSuite is Changed
	isTestSuiteChanged = testSuiteModel.IsTestSuiteChanged()

	// Inform user that TestSuite is not changed
	if isTestSuiteChanged == false {

		// Trigger System Notification sound
		soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "Save TestSuite",
			Content: fmt.Sprintf("TestSuite is not changed"),
		})

		return nil
	}

	//testSuiteModel.NoneSavedTestSuiteUIModelBinding = createEmptyAndInitiatedTestSuiteModel(testSuiteModel.testCasesModel)

	// Copy fields from 'TestSuiteUIModelBinding' to  'NoneSavedTestSuiteUIModelBinding' using deep copy
	err = copier.CopyWithOption(&testSuiteModel.NoneSavedTestSuiteUIModelBinding, &testSuiteModel.TestSuiteUIModelBinding, copier.Option{DeepCopy: true})
	if err != nil {

		errorID := "85acf490-2599-4a8a-bf9e-1451eef60f78"

		errorMsg := fmt.Sprintf("error copying 'TestSuiteUIModelBinding' using 'copier'. error = '%s' [ErrorID: %s]",
			err.Error(),
			errorID)

		log.Fatalln(errorMsg)
	}

	// Check if all mandatory fields has values
	mandatoryFieldsHaveValues,
		mandatoryFieldsHaveValuesNotificationText = testSuiteModel.checkIfAllMandatoryFieldsHaveValues()

	// Inform user that mandatory field in TestSuite is missing
	if mandatoryFieldsHaveValues == false {

		// Trigger System Notification sound
		soundEngine.PlaySoundChannel <- soundEngine.UserNeedToRespondSound

		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title: "Mandatory Fields",
			Content: fmt.Sprintf("Mandatory field '%s' is empty",
				mandatoryFieldsHaveValuesNotificationText),
		})

		return nil
	}

	// Copy all UI-fields to model
	//testSuiteModel.copyUiFieldsToModel()

	// Keeps a list of what is to be saved. Used to ensure that older versions of the TestSuite can later be loaded when new functionality has been added to the client
	var supportedTestSuiteDataToBeStored testSuiteImplementedFunctionsToBeStoredStruct
	supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap = make(map[testSuiteImplementedFucntionsType]bool)

	// Generate 'TestSuiteBasicInformation' to be added to full gRPC-message
	var testSuiteBasicInformation *fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteBasicInformationMessage
	var testSuiteBasicInformationHash string
	testSuiteBasicInformation, testSuiteBasicInformationHash, err = testSuiteModel.
		generateTestSuiteBasicInformationMessageWhenSaving(&supportedTestSuiteDataToBeStored)
	valuesToBeHashed = append(valuesToBeHashed, testSuiteBasicInformationHash)

	// Generate 'DeleteDate' to be added to full gRPC-message
	var testSuiteDeleteDate string
	var testSuiteDeleteDateHash string
	testSuiteDeleteDate, testSuiteDeleteDateHash, err = testSuiteModel.generateTestSuiteDeleteDateMessageWhenSaving(
		&supportedTestSuiteDataToBeStored)
	valuesToBeHashed = append(valuesToBeHashed, testSuiteDeleteDateHash)

	// Convert 'TestSuiteType' into gRPC-message
	var testSuiteType *fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteTypeMessage
	var testSuiteTypeHash string
	testSuiteType, testSuiteTypeHash, err = testSuiteModel.
		generateTestSuiteTypeMessageWhenSaving(&supportedTestSuiteDataToBeStored)
	valuesToBeHashed = append(valuesToBeHashed, testSuiteTypeHash)

	// Convert 'supportedTestSuiteDataToBeStored'to be added to full gRPC-message
	var testSuiteImplementedFunctionsMap map[int32]bool
	var testSuiteImplementedFunctionsMapHash string
	testSuiteImplementedFunctionsMap = make(map[int32]bool)
	testSuiteImplementedFunctionsMap, testSuiteImplementedFunctionsMapHash, err = testSuiteModel.
		generateTestSuiteImplementedFunctionsMapWhenSaving(&supportedTestSuiteDataToBeStored)
	valuesToBeHashed = append(valuesToBeHashed, testSuiteImplementedFunctionsMapHash)

	// Create MessageHash
	messageHash = sharedCode.HashValues(valuesToBeHashed, false)

	// Generate full gRPC-message to be sent to TestCaseBuilder-Server
	var fullTestSuiteMessage fenixGuiTestCaseBuilderServerGrpcApi.FullTestSuiteMessage
	fullTestSuiteMessage = fenixGuiTestCaseBuilderServerGrpcApi.FullTestSuiteMessage{
		TestSuiteBasicInformation:        testSuiteBasicInformation,
		TestSuiteTestData:                nil,
		TestSuitePreview:                 nil,
		TestSuiteMetaData:                nil,
		TestCasesInTestSuite:             nil,
		DeletedDate:                      testSuiteDeleteDate,
		UpdatedByAndWhen:                 nil, // Used when loading TestSuite
		TestSuiteType:                    testSuiteType,
		TestSuiteImplementedFunctionsMap: testSuiteImplementedFunctionsMap,
		MessageHash:                      messageHash,
	}

	// Send using gRPC
	returnMessage := testSuiteModel.testCasesModel.GrpcOutReference.SendSaveFullTestSuite(&fullTestSuiteMessage)

	if returnMessage == nil || returnMessage.AckNack == false {

		// Trigger System Notification sound
		soundEngine.PlaySoundChannel <- soundEngine.InvalidNotificationSound

		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "Save TestSuite",
			Content: fmt.Sprintf("Got some error when trying to Save TestSuite"),
		})

		if returnMessage == nil {
			errorId := "6c8be10e-404e-49a3-b121-df05c952d5b8"
			err = errors.New(fmt.Sprintf("Got 'nil' back when trying to Save TestSuite [ErrorID: %s]", errorId))

		} else {
			errorId := "5f887e6f-db34-4ca0-8abf-4521657c2337"
			err = errors.New(fmt.Sprintf(returnMessage.Comments+"[ErrorID: %s]", errorId))

		}

		fmt.Println(err) // TODO Send on Error-channel

		return err
	}

	// Trigger System Notification sound
	soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   "Save TestSuite",
		Content: fmt.Sprintf("Success in saving TestSuite"),
	})

	// Update that the TestSuite is not New anymore and Update the TestSuiteVersion
	testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteIsNew = false
	testSuiteModel.TestSuiteUIModelBinding.TestSuiteIsNew = false

	// Copy data from 'NoneSavedTestSuiteUIModelBinding' to 'savedTestSuiteUIModelBinding' using deep copy
	err = copier.CopyWithOption(&testSuiteModel.savedTestSuiteUIModelBinding, &testSuiteModel.NoneSavedTestSuiteUIModelBinding, copier.Option{DeepCopy: true})
	if err != nil {

		errorID := "2304dabe-7e12-4cf4-a021-597793ace220"

		errorMsg := fmt.Sprintf("error copying 'NoneSavedTestSuiteUIModelBinding' using 'copier'. error = '%s' [ErrorID: %s]",
			err.Error(),
			errorID)

		log.Fatalln(errorMsg)
	}

	// Update The Hash for the saved TestSuite
	testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.testSuiteSavedMessageHash = messageHash

	// Set the new TestSuiteVersion
	testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.testSuiteVersion = testSuiteBasicInformation.TestSuiteVersion

	return err

}

// Generates 'TestSuiteBasicInformation' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteBasicInformationMessageWhenSaving(
	supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct) (
	testSuiteBasicInformation *fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteBasicInformationMessage,
	testSuiteBasicInformationHash string,
	err error) {

	// This TestSuite has stored 'testSuiteBasicInformationIsSupported'
	supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap[testSuiteBasicInformationIsSupported] = true

	// Generate TestSuiteVersion
	var testSuiteVersion uint32
	if testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteIsNew == true {
		testSuiteVersion = 1
	} else {
		testSuiteVersion = testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.testSuiteVersion + 1
	}

	// Create 'testSuiteBasicInformation'
	testSuiteBasicInformation = &fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteBasicInformationMessage{
		DomainUuid:           testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteOwnerDomainUuid,
		DomainName:           testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteOwnerDomainName,
		TestSuiteUuid:        testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.testSuiteUuid,
		TestSuiteVersion:     testSuiteVersion,
		TestSuiteName:        testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteName,
		TestSuiteDescription: testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteDescription,
	}

	// Create the Hash of the Message
	var tempJson string
	tempJson = protojson.Format(testSuiteBasicInformation)
	testSuiteBasicInformationHash = sharedCode.HashSingleValue(tempJson)

	return testSuiteBasicInformation, testSuiteBasicInformationHash, err

}

// Generates 'UsersChosenTestDataForTestSuiteMessage' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteTestDataMessageWhenSaving(supportedTestSuiteDataToBeStored *[]testSuiteImplementedFucntionsType) (
	testSuiteTestData *fenixGuiTestCaseBuilderServerGrpcApi.UsersChosenTestDataForTestSuiteMessage,
	testSuiteTestDataHash string,
	err error) {

	// This TestSuite has stored 'testSuiteTestDataIsSupported'
	//*supportedTestSuiteDataToBeStored = append(*supportedTestSuiteDataToBeStored, testSuiteTestDataIsSupported)

	return testSuiteTestData, testSuiteTestDataHash, err

}

// Generates 'TestSuitePreview' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuitePreviewMessageWhenSaving(supportedTestSuiteDataToBeStored *[]testSuiteImplementedFucntionsType) (
	testSuitePreview *fenixGuiTestCaseBuilderServerGrpcApi.TestSuitePreviewMessage,
	testSuitePreviewHash string,
	err error) {

	// This TestSuite has stored 'testSuitePreviewIsSupported'
	//*supportedTestSuiteDataToBeStored = append(*supportedTestSuiteDataToBeStored, testSuitePreviewIsSupported)

	return testSuitePreview, testSuitePreviewHash, err

}

// Generates 'TestSuiteMetaData' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteMetaDataMessageWhenSaving(supportedTestSuiteDataToBeStored *[]testSuiteImplementedFucntionsType) (
	testSuiteMetaData *fenixGuiTestCaseBuilderServerGrpcApi.UserSpecifiedTestSuiteMetaDataMessage,
	testSuiteMetaDataHash string,
	err error) {

	// This TestSuite has stored 'testSuiteMetaDataIsSupported'
	//*supportedTestSuiteDataToBeStored = append(*supportedTestSuiteDataToBeStored, testSuiteMetaDataIsSupported)

	return testSuiteMetaData, testSuiteMetaDataHash, err

}

// Generates 'TestCasesInTestSuite' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestCasesInTestSuiteMessageWhenSaving(supportedTestSuiteDataToBeStored *[]testSuiteImplementedFucntionsType) (
	testCasesInTestSuite *fenixGuiTestCaseBuilderServerGrpcApi.TestCasesInTestSuiteMessage,
	testCasesInTestSuitenHash string,
	err error) {

	// This TestSuite has stored 'testCasesInTestSuiteIsSupported'
	//*supportedTestSuiteDataToBeStored = append(*supportedTestSuiteDataToBeStored, testCasesInTestSuiteIsSupported)

	return testCasesInTestSuite, testCasesInTestSuitenHash, err

}

// Generates 'TestSuiteDeleteData' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteDeleteDateMessageWhenSaving(
	supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct) (
	testSuiteDeleteDate string,
	testSuiteDeleteDateHash string,
	err error) {

	// This TestSuite has stored 'deletedDateIsSupported'
	supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap[deletedDateIsSupported] = true

	// Create 'testSuiteDeleteDate'
	testSuiteDeleteDate = testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteDeletionDate

	// Create the Hash of the Message
	testSuiteDeleteDateHash = sharedCode.HashSingleValue(testSuiteDeleteDate)

	return testSuiteDeleteDate, testSuiteDeleteDateHash, err

}

// Generates 'TestCasesInTestSuite' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteTypeMessageWhenSaving(
	supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct) (
	testSuiteTypeMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteTypeMessage,
	testSuiteTypeHash string,
	err error) {

	// This TestSuite has stored 'testSuiteTypeIsSupported'
	supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap[testSuiteTypeIsSupported] = true

	// Create 'testSuiteTypeMessage'
	testSuiteTypeMessage = &fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteTypeMessage{
		TestSuiteType: fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteTypeEnum(
			testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteType.TestSuiteType),
		TestSuiteTypeName: string(testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteType.TestSuiteTypeName),
	}

	// Create Hash for 'testSuiteTypeIsSupported'
	var valuesToBeHashed []string
	valuesToBeHashed = append(valuesToBeHashed, fmt.Sprintf("%d", testSuiteTypeMessage.TestSuiteType))
	valuesToBeHashed = append(valuesToBeHashed, fmt.Sprintf("%s", testSuiteTypeMessage.TestSuiteTypeName))

	// Create the Hash of the Message
	testSuiteTypeHash = sharedCode.HashValues(valuesToBeHashed, true)

	return testSuiteTypeMessage, testSuiteTypeHash, err

}

// Generates 'TestSuiteImplementedFunctionsMap' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteImplementedFunctionsMapWhenSaving(
	supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct) (
	testSuiteImplementedFunctionsMap map[int32]bool,
	testSuiteImplementedFunctionsMapHash string,
	err error) {

	// This TestSuite has stored 'testSuiteImplementedFunctionsMapIsSupported'
	supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap[testSuiteImplementedFunctionsMapIsSupported] = true

	// Create 'testSuiteImplementedFunctionsMap'
	testSuiteImplementedFunctionsMap = make(map[int32]bool)

	// Loop 'supportedTestSuiteDataToBeStored' and convert to gRPC-message
	var valuesToBeHashed []string
	for testSuiteImplementedFunction, testSuiteImplementedFunctionValue := range supportedTestSuiteDataToBeStored.
		testSuiteImplementedFunctionsMap {

		testSuiteImplementedFunctionsMap[int32(testSuiteImplementedFunction)] = testSuiteImplementedFunctionValue
		valuesToBeHashed = append(valuesToBeHashed, fmt.Sprintf("%d", testSuiteImplementedFunction))
	}

	// Create the Hash of the Message
	testSuiteImplementedFunctionsMapHash = sharedCode.HashValues(valuesToBeHashed, true)

	return testSuiteImplementedFunctionsMap, testSuiteImplementedFunctionsMapHash, err

}
