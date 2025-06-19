package testSuitesModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/soundEngine"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
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
	testSuiteModel.copyUiFieldsToModel()

	// Generate 'TestSuiteBasicInformation' to be added to full gRPC-message
	var testSuiteBasicInformation *fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteBasicInformationMessage
	var testSuiteBasicInformationHash string
	testSuiteBasicInformation, testSuiteBasicInformationHash, err = testSuiteModel.generateTestSuiteBasicInformationMessage()
	valuesToBeHashed = append(valuesToBeHashed, testSuiteBasicInformationHash)

	// Generate 'UsersChosenTestDataForTestSuiteMessage' to be added to full gRPC-message
	var testSuiteTestData *fenixGuiTestCaseBuilderServerGrpcApi.UsersChosenTestDataForTestSuiteMessage
	var testSuiteTestDataHash string
	testSuiteTestData, testSuiteTestDataHash, err = testSuiteModel.generateTestSuiteTestDataMessage()
	valuesToBeHashed = append(valuesToBeHashed, testSuiteTestDataHash)

	// Generate 'TestSuitePreview' to be added to full gRPC-message
	var testSuitePreview *fenixGuiTestCaseBuilderServerGrpcApi.TestSuitePreviewMessage
	var testSuitePreviewHash string
	testSuitePreview, testSuitePreviewHash, err = testSuiteModel.generateTestSuitePreviewMessage()
	valuesToBeHashed = append(valuesToBeHashed, testSuitePreviewHash)

	// Generate 'TestSuiteMetaData' to be added to full gRPC-message
	var testSuiteMetaData *fenixGuiTestCaseBuilderServerGrpcApi.UserSpecifiedTestSuiteMetaDataMessage
	var testSuiteMetaDataHash string
	testSuiteMetaData, testSuiteMetaDataHash, err = testSuiteModel.generateTestSuiteMetaDataMessage()
	valuesToBeHashed = append(valuesToBeHashed, testSuiteMetaDataHash)

	// Generate 'TestSuiteMetaData' to be added to full gRPC-message
	var testCasesInTestSuite *fenixGuiTestCaseBuilderServerGrpcApi.TestCasesInTestSuiteMessage
	var testCasesInTestSuitenHash string
	testCasesInTestSuite, testCasesInTestSuitenHash, err = testSuiteModel.generateTestCasesInTestSuiteMessage()
	valuesToBeHashed = append(valuesToBeHashed, testCasesInTestSuitenHash)

	// Generate 'DeleteDate' to be added to full gRPC-message
	var testSuiteDeleteDate string
	var testSuiteDeleteDateHash string
	testSuiteDeleteDate, testSuiteDeleteDateHash, err = testSuiteModel.generateTestSuiteDeleteDateMessage()
	valuesToBeHashed = append(valuesToBeHashed, testSuiteDeleteDateHash)

	// Create MessageHash
	messageHash = sharedCode.HashValues(valuesToBeHashed, false)

	// Generate full gRPC-message to be sent to TestCaseBuilder-Server
	var fullTestSuiteMessage fenixGuiTestCaseBuilderServerGrpcApi.FullTestSuiteMessage
	fullTestSuiteMessage = fenixGuiTestCaseBuilderServerGrpcApi.FullTestSuiteMessage{
		TestSuiteBasicInformation: testSuiteBasicInformation,
		TestSuiteTestData:         testSuiteTestData,
		TestSuitePreview:          testSuitePreview,
		TestSuiteMetaData:         testSuiteMetaData,
		TestCasesInTestSuite:      testCasesInTestSuite,
		DeletedDate:               testSuiteDeleteDate,
		MessageHash:               messageHash,
	}

	// Send using gRPC
	returnMessage := testSuiteModel.testCasesModel.GrpcOutReference.SendSaveFullTestSuite(&fullTestSuiteMessage)

	if returnMessage == nil || returnMessage.AckNack == false {

		errorId := "e7c11c11-27dd-4889-a41b-c0b9bf7bcada"
		err = errors.New(fmt.Sprintf(returnMessage.Comments+"[ErrorID: %s]", errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return err
	}

	// Update that the TestSuite is not New anymore
	testSuiteModel.testSuiteIsNew = false

	// Update The Hash for the TestCase
	currentTestCasePtr.TestCaseHashWhenTestCaseWasSavedOrLoaded = gRPCFullTestCaseMessageToSend.MessageHash
	currentTestCasePtr.TestDataHashWhenTestCaseWasSavedOrLoaded = gRPCFullTestCaseMessageToSend.GetTestCaseTestData().GetHashOfThisMessageWithEmptyHashField()

	return err

}

// Generates 'TestSuiteBasicInformation' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteBasicInformationMessage() (
	testSuiteBasicInformation *fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteBasicInformationMessage,
	testSuiteBasicInformationHash string,
	err error) {

	return testSuiteBasicInformation, testSuiteBasicInformationHash, err

}

// Generates 'UsersChosenTestDataForTestSuiteMessage' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteTestDataMessage() (
	testSuiteTestData *fenixGuiTestCaseBuilderServerGrpcApi.UsersChosenTestDataForTestSuiteMessage,
	testSuiteTestDataHash string,
	err error) {

	return testSuiteTestData, testSuiteTestDataHash, err

}

// Generates 'TestSuitePreview' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuitePreviewMessage() (
	testSuitePreview *fenixGuiTestCaseBuilderServerGrpcApi.TestSuitePreviewMessage,
	testSuitePreviewHash string,
	err error) {

	return testSuitePreview, testSuitePreviewHash, err

}

// Generates 'TestSuiteMetaData' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteMetaDataMessage() (
	testSuiteMetaData *fenixGuiTestCaseBuilderServerGrpcApi.UserSpecifiedTestSuiteMetaDataMessage,
	testSuiteMetaDataHash string,
	err error) {

	return testSuiteMetaData, testSuiteMetaDataHash, err

}

// Generates 'TestCasesInTestSuite' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestCasesInTestSuiteMessage() (
	testCasesInTestSuite *fenixGuiTestCaseBuilderServerGrpcApi.TestCasesInTestSuiteMessage,
	testCasesInTestSuitenHash string,
	err error) {

	return testCasesInTestSuite, testCasesInTestSuitenHash, err

}

// Generates 'TestSuiteDeleteData' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteDeleteDateMessage() (
	testSuiteDeleteDate string,
	testSuiteDeleteDateHash string,
	err error) {

	return testSuiteDeleteDate, testSuiteDeleteDate, err

}
