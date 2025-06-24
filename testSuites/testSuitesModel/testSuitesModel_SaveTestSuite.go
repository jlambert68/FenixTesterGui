package testSuitesModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/soundEngine"
	"encoding/json"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"google.golang.org/protobuf/encoding/protojson"
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

	// Keeps a list of what is to be saved. Used to ensure that older versions of the TestSuite can later be loaded
	var supportedTestSuiteDataToBeStored testSuiteDataToBeStoredStruct
	supportedTestSuiteDataToBeStored.supportedTestSuiteDataToBeStoredMap = make(map[supportedTestSuiteDataToBeStoredType]bool)

	// Generate 'TestSuiteBasicInformation' to be added to full gRPC-message
	var testSuiteBasicInformation *fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteBasicInformationMessage
	var testSuiteBasicInformationHash string
	testSuiteBasicInformation, testSuiteBasicInformationHash, err = testSuiteModel.
		generateTestSuiteBasicInformationMessage(&supportedTestSuiteDataToBeStored)
	valuesToBeHashed = append(valuesToBeHashed, testSuiteBasicInformationHash)

	// Generate 'DeleteDate' to be added to full gRPC-message
	var testSuiteDeleteDate string
	var testSuiteDeleteDateHash string
	testSuiteDeleteDate, testSuiteDeleteDateHash, err = testSuiteModel.generateTestSuiteDeleteDateMessage(
		&supportedTestSuiteDataToBeStored)
	valuesToBeHashed = append(valuesToBeHashed, testSuiteDeleteDateHash)

	// Convert 'supportedTestSuiteDataToBeStored' into json and the HAsh
	var supportedTestSuiteDataToBeStoredAsJsonByteArray []byte
	var supportedTestSuiteDataToBeStoredAsJsonString string
	var supportedTestSuiteDataToBeStoredAsHash string
	supportedTestSuiteDataToBeStoredAsJsonByteArray, err = json.Marshal(supportedTestSuiteDataToBeStored.supportedTestSuiteDataToBeStoredMap)
	if err != nil {

		errorId := "f6a1789c-d50a-4989-bebc-d10395377ec2"
		err = errors.New(fmt.Sprintf("couldn't marshal 'supportedTestSuiteDataToBeStored'. Error = '%s'. [ErrorID: %s]",
			err.Error(),
			errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return err
	}

	supportedTestSuiteDataToBeStoredAsJsonString = string(supportedTestSuiteDataToBeStoredAsJsonByteArray)
	supportedTestSuiteDataToBeStoredAsHash = sharedCode.HashSingleValue(supportedTestSuiteDataToBeStoredAsJsonString)

	// Add 'supportedTestSuiteDataToBeStoredAsHash' to hash-array
	valuesToBeHashed = append(valuesToBeHashed, supportedTestSuiteDataToBeStoredAsHash)

	// Create MessageHash
	messageHash = sharedCode.HashValues(valuesToBeHashed, false)

	// Generate full gRPC-message to be sent to TestCaseBuilder-Server
	var fullTestSuiteMessage fenixGuiTestCaseBuilderServerGrpcApi.FullTestSuiteMessage
	fullTestSuiteMessage = fenixGuiTestCaseBuilderServerGrpcApi.FullTestSuiteMessage{
		TestSuiteBasicInformation: testSuiteBasicInformation,
		TestSuiteTestData:         nil,
		TestSuitePreview:          nil,
		TestSuiteMetaData:         nil,
		TestCasesInTestSuite:      nil,
		DeletedDate:               testSuiteDeleteDate,
		SupportedTestSuiteDataToBeStoredAsJsonString: supportedTestSuiteDataToBeStoredAsJsonString,
		MessageHash: messageHash,
	}

	// Send using gRPC
	returnMessage := testSuiteModel.testCasesModel.GrpcOutReference.SendSaveFullTestSuite(&fullTestSuiteMessage)

	if returnMessage == nil || returnMessage.AckNack == false {

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

	// Update that the TestSuite is not New anymore and Update the TestSuiteVersion
	testSuiteModel.testSuiteIsNew = false
	testSuiteModel.testSuiteVersion = testSuiteBasicInformation.TestSuiteVersion

	// Update The Hash for the saved TestSuite
	testSuiteModel.testSuiteTestDataHash = messageHash

	return err

}

// Generates 'TestSuiteBasicInformation' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteBasicInformationMessage(
	supportedTestSuiteDataToBeStored *testSuiteDataToBeStoredStruct) (
	testSuiteBasicInformation *fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteBasicInformationMessage,
	testSuiteBasicInformationHash string,
	err error) {

	// This TestSuite has stored 'testSuiteBasicInformationIsSupported'
	supportedTestSuiteDataToBeStored.supportedTestSuiteDataToBeStoredMap[testSuiteBasicInformationIsSupported] = true

	// Generate TestSuiteVersion
	var testSuiteVersion uint32
	if testSuiteModel.testSuiteIsNew == true {
		testSuiteVersion = 1
	} else {
		testSuiteVersion = testSuiteModel.testSuiteVersion + 1
	}

	// Create 'testSuiteBasicInformation'
	testSuiteBasicInformation = &fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteBasicInformationMessage{
		DomainUuid:           testSuiteModel.testSuiteOwnerDomainUuid,
		DomainName:           testSuiteModel.testSuiteOwnerDomainName,
		TestSuiteUuid:        testSuiteModel.testSuiteUuid,
		TestSuiteVersion:     testSuiteVersion,
		TestSuiteName:        testSuiteModel.testSuiteName,
		TestSuiteDescription: testSuiteModel.testSuiteDescription,
	}

	// Create the Hash of the Message
	var tempJson string
	tempJson = protojson.Format(testSuiteBasicInformation)
	testSuiteBasicInformationHash = sharedCode.HashSingleValue(tempJson)

	return testSuiteBasicInformation, testSuiteBasicInformationHash, err

}

// Generates 'UsersChosenTestDataForTestSuiteMessage' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteTestDataMessage(supportedTestSuiteDataToBeStored *[]supportedTestSuiteDataToBeStoredType) (
	testSuiteTestData *fenixGuiTestCaseBuilderServerGrpcApi.UsersChosenTestDataForTestSuiteMessage,
	testSuiteTestDataHash string,
	err error) {

	// This TestSuite has stored 'testSuiteTestDataIsSupported'
	//*supportedTestSuiteDataToBeStored = append(*supportedTestSuiteDataToBeStored, testSuiteTestDataIsSupported)

	return testSuiteTestData, testSuiteTestDataHash, err

}

// Generates 'TestSuitePreview' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuitePreviewMessage(supportedTestSuiteDataToBeStored *[]supportedTestSuiteDataToBeStoredType) (
	testSuitePreview *fenixGuiTestCaseBuilderServerGrpcApi.TestSuitePreviewMessage,
	testSuitePreviewHash string,
	err error) {

	// This TestSuite has stored 'testSuitePreviewIsSupported'
	//*supportedTestSuiteDataToBeStored = append(*supportedTestSuiteDataToBeStored, testSuitePreviewIsSupported)

	return testSuitePreview, testSuitePreviewHash, err

}

// Generates 'TestSuiteMetaData' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteMetaDataMessage(supportedTestSuiteDataToBeStored *[]supportedTestSuiteDataToBeStoredType) (
	testSuiteMetaData *fenixGuiTestCaseBuilderServerGrpcApi.UserSpecifiedTestSuiteMetaDataMessage,
	testSuiteMetaDataHash string,
	err error) {

	// This TestSuite has stored 'testSuiteMetaDataIsSupported'
	//*supportedTestSuiteDataToBeStored = append(*supportedTestSuiteDataToBeStored, testSuiteMetaDataIsSupported)

	return testSuiteMetaData, testSuiteMetaDataHash, err

}

// Generates 'TestCasesInTestSuite' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestCasesInTestSuiteMessage(supportedTestSuiteDataToBeStored *[]supportedTestSuiteDataToBeStoredType) (
	testCasesInTestSuite *fenixGuiTestCaseBuilderServerGrpcApi.TestCasesInTestSuiteMessage,
	testCasesInTestSuitenHash string,
	err error) {

	// This TestSuite has stored 'testCasesInTestSuiteIsSupported'
	//*supportedTestSuiteDataToBeStored = append(*supportedTestSuiteDataToBeStored, testCasesInTestSuiteIsSupported)

	return testCasesInTestSuite, testCasesInTestSuitenHash, err

}

// Generates 'TestSuiteDeleteData' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteDeleteDateMessage(
	supportedTestSuiteDataToBeStored *testSuiteDataToBeStoredStruct) (
	testSuiteDeleteDate string,
	testSuiteDeleteDateHash string,
	err error) {

	// This TestSuite has stored 'deletedDateIsSupported'
	supportedTestSuiteDataToBeStored.supportedTestSuiteDataToBeStoredMap[deletedDateIsSupported] = true

	// Create 'testSuiteDeleteDate'
	testSuiteDeleteDate = testSuiteModel.testSuiteDeletionDate

	// Create the Hash of the Message
	testSuiteDeleteDateHash = sharedCode.HashSingleValue(testSuiteDeleteDate)

	return testSuiteDeleteDate, testSuiteDeleteDateHash, err

}
