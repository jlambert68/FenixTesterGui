package testSuitesModel

import (
	sharedCode "FenixTesterGui/common_code"
	"encoding/json"
	"errors"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"google.golang.org/protobuf/encoding/protojson"
)

// LoadFullTestSuiteFromDatabase - Load the TestSuite from the Database into model
func (testSuiteModel *TestSuiteModelStruct) LoadFullTestSuiteFromDatabase(
	testSuiteUuid string) (err error) {

	// Send LoadTestSuite using gRPC
	var detailedTestSuiteResponse *fenixGuiTestCaseBuilderServerGrpcApi.GetDetailedTestSuiteResponse
	detailedTestSuiteResponse = testSuiteModel.testCasesModel.GrpcOutReference.LoadDetailedTestSuite(testSuiteUuid)

	// Exit if something was wrong
	if detailedTestSuiteResponse.AckNackResponse.AckNack == false {

		errorId := "84c34946-f840-47cb-8ce1-3dcb585e5941"
		err = errors.New(fmt.Sprintf(detailedTestSuiteResponse.AckNackResponse.Comments+"[ErrorID: %s]", testSuiteUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return err

	}

	// Check that correct TestSuite was loaded
	if detailedTestSuiteResponse.GetDetailedTestSuite().GetTestSuiteBasicInformation().GetTestSuiteUuid() != testSuiteUuid {

		errorId := "e57a57bd-9645-440b-9e12-0b19a6c4cf47"
		err = errors.New(fmt.Sprintf("Asked for TestSuite '%s'  but got '%s' [ErrorID: %s]",
			testSuiteUuid,
			detailedTestSuiteResponse.GetDetailedTestSuite().GetTestSuiteBasicInformation().GetTestSuiteUuid(),
			errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return err
	}

	// Extract a list of what was saved.
	//Used to ensure that older versions of the TestSuite can later be loaded when new functionality has been added to the client
	var supportedTestSuiteDataToBeStored testSuiteImplementedFunctionsToBeStoredStruct
	supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap = testSuiteModel.
		extractTestSuiteImplementedFunctionsMap(detailedTestSuiteResponse.GetDetailedTestSuite().
			GetTestSuiteImplementedFunctionsMap())

	// Object holding the data that can't be changed directly via the UI
	var tempTestSuiteModelDataThatCanNotBeChangedFromUI testSuiteModelDataThatCaNotBeChangedFromUIStruct
	tempTestSuiteModelDataThatCanNotBeChangedFromUI = testSuiteModelDataThatCaNotBeChangedFromUIStruct{
		testSuiteUuid: detailedTestSuiteResponse.GetDetailedTestSuite().GetTestSuiteBasicInformation().
			GetTestSuiteUuid(),
		testSuiteVersion: detailedTestSuiteResponse.GetDetailedTestSuite().GetTestSuiteBasicInformation().
			GetTestSuiteVersion(),
		createdByGcpLogin: detailedTestSuiteResponse.GetDetailedTestSuite().GetUpdatedByAndWhen().
			GetCreatedByGcpLogin(),
		createdByComputerLogin: detailedTestSuiteResponse.GetDetailedTestSuite().GetUpdatedByAndWhen().
			GetCreatedByComputerLogin(),
		createdDate: detailedTestSuiteResponse.GetDetailedTestSuite().GetUpdatedByAndWhen().
			GetCreatedDate(),
		lastChangedByGcpLogin: detailedTestSuiteResponse.GetDetailedTestSuite().GetUpdatedByAndWhen().
			GetGCPAuthenticatedUser(),
		lastChangedByComputerLogin: detailedTestSuiteResponse.GetDetailedTestSuite().GetUpdatedByAndWhen().
			GetUserIdOnComputer(),
		lastChangedDate: detailedTestSuiteResponse.GetDetailedTestSuite().GetUpdatedByAndWhen().
			GetUpdateTimeStamp(),
		testSuiteSavedMessageHash: "",
	}

	// Object holding info about if OwnerDomain and TestEnvironment has got any value, used for Locking parts of the GUI in the TestSuite
	var tempLockValuesForOwnerDomainAndTestEnvironment lockValuesForOwnerDomainAndTestEnvironmentStruct
	tempLockValuesForOwnerDomainAndTestEnvironment = lockValuesForOwnerDomainAndTestEnvironmentStruct{
		OwnerDomainHasValue:     true,
		TestEnvironmentHasValue: true,
		LockButtonHaBeenClicked: true,
	}

	// Structure that keeps data used after saving was successfully performed
	var tempSavedTestSuiteUIModelBinding TestSuiteUIModelBindingStruct
	tempSavedTestSuiteUIModelBinding = TestSuiteUIModelBindingStruct{
		TestSuiteDeletionDate: detailedTestSuiteResponse.GetDetailedTestSuite().
			GetDeletedDate(),
		TestSuiteName: detailedTestSuiteResponse.GetDetailedTestSuite().GetTestSuiteBasicInformation().
			GetTestSuiteName(),
		TestSuiteDescription: detailedTestSuiteResponse.GetDetailedTestSuite().GetTestSuiteBasicInformation().
			GetTestSuiteDescription(),
		TestSuiteOwnerDomainUuid: detailedTestSuiteResponse.GetDetailedTestSuite().GetTestSuiteBasicInformation().
			GetDomainUuid(),
		TestSuiteOwnerDomainName: detailedTestSuiteResponse.GetDetailedTestSuite().GetTestSuiteBasicInformation().
			GetDomainName(),
		TestSuiteExecutionEnvironment: "",
		TestSuiteIsNew:                false,
		TestSuiteTestDataHash: detailedTestSuiteResponse.GetDetailedTestSuite().GetTestSuiteTestData().
			GetHashOfThisMessageWithEmptyHashField(),
		TestDataPtr:           nil,
		TestSuiteMetaDataHash: "",
		TestSuiteMetaDataPtr:  nil,
		TestSuiteTypeHash:     "",
		TestSuiteType:         TestSuiteTypeStruct{},
	}

	// Create object that will hold complete TestSuite in memory
	testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI = tempTestSuiteModelDataThatCanNotBeChangedFromUI
	testSuiteModel.lockValuesForOwnerDomainAndTestEnvironment = tempLockValuesForOwnerDomainAndTestEnvironment
	testSuiteModel.savedTestSuiteUIModelBinding = tempSavedTestSuiteUIModelBinding

	return err

}

// Extract 'TestSuiteImplementedFunctionsMap' from gRPC-message
func (testSuiteModel *TestSuiteModelStruct) extractTestSuiteImplementedFunctionsMap(
	testSuiteImplementedFunctionsMapAsJsonString string) (
	testSuiteImplementedFunctionsMap map[testSuiteImplementedFucntionsType]bool,
	err error) {

	testSuiteImplementedFunctionsMap = make(map[testSuiteImplementedFucntionsType]bool)

	// UnMarshal the json
	err = json.Unmarshal([]byte(testSuiteImplementedFunctionsMapAsJsonString), testSuiteImplementedFunctionsMap)
	if err != nil {

		errorId := "8a3705bf-b074-4cd9-9466-986cff9d329a"
		err = errors.New(fmt.Sprintf("couldn't unmarshal 'testSuiteImplementedFunctionsMapAsJsonString'. Error = '%s'. [ErrorID: %s]",
			err.Error(),
			errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return nil, err
	}

	return testSuiteImplementedFunctionsMap, err
}

// Generates 'TestSuiteBasicInformation' from gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteBasicInformationMessageWhenLoading(
	supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct) (
	testSuiteBasicInformation *fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteBasicInformationMessage,
	testSuiteBasicInformationHash string,
	err error) {

	// Check if this TestSuite was stored with 'testSuiteBasicInformationIsSupported'
	if supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap[testSuiteBasicInformationIsSupported] == false {
		return nil, "", errors.New(fmt.Sprintf("TestSuiteBasicInformation is not supported"))
	}

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

// Generates 'UsersChosenTestDataForTestSuiteMessage' from gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteTestDataMessageWhenLoading(supportedTestSuiteDataToBeStored *[]testSuiteImplementedFucntionsType) (
	testSuiteTestData *fenixGuiTestCaseBuilderServerGrpcApi.UsersChosenTestDataForTestSuiteMessage,
	testSuiteTestDataHash string,
	err error) {

	// Check if this TestSuite was stored with 'testSuiteTestDataIsSupported'
	//*supportedTestSuiteDataToBeStored = append(*supportedTestSuiteDataToBeStored, testSuiteTestDataIsSupported)

	return testSuiteTestData, testSuiteTestDataHash, err

}

// Generates 'TestSuitePreview' from gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuitePreviewMessageWhenLoading(supportedTestSuiteDataToBeStored *[]testSuiteImplementedFucntionsType) (
	testSuitePreview *fenixGuiTestCaseBuilderServerGrpcApi.TestSuitePreviewMessage,
	testSuitePreviewHash string,
	err error) {

	// Check if this TestSuite was stored with 'testSuitePreviewIsSupported'
	//*supportedTestSuiteDataToBeStored = append(*supportedTestSuiteDataToBeStored, testSuitePreviewIsSupported)

	return testSuitePreview, testSuitePreviewHash, err

}

// Generates 'TestSuiteMetaData' from gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteMetaDataMessageWhenLoading(supportedTestSuiteDataToBeStored *[]testSuiteImplementedFucntionsType) (
	testSuiteMetaData *fenixGuiTestCaseBuilderServerGrpcApi.UserSpecifiedTestSuiteMetaDataMessage,
	testSuiteMetaDataHash string,
	err error) {

	// Check if this TestSuite was stored with 'testSuiteMetaDataIsSupported'
	//*supportedTestSuiteDataToBeStored = append(*supportedTestSuiteDataToBeStored, testSuiteMetaDataIsSupported)

	return testSuiteMetaData, testSuiteMetaDataHash, err

}

// Generates 'TestCasesInTestSuite' from gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestCasesInTestSuiteMessageWhenLoading(supportedTestSuiteDataToBeStored *[]testSuiteImplementedFucntionsType) (
	testCasesInTestSuite *fenixGuiTestCaseBuilderServerGrpcApi.TestCasesInTestSuiteMessage,
	testCasesInTestSuitenHash string,
	err error) {

	// Check if this TestSuite was stored with 'testCasesInTestSuiteIsSupported'
	//*supportedTestSuiteDataToBeStored = append(*supportedTestSuiteDataToBeStored, testCasesInTestSuiteIsSupported)

	return testCasesInTestSuite, testCasesInTestSuitenHash, err

}

// Generates 'TestSuiteDeleteData' from gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteDeleteDateMessageWhenLoading(
	supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct,
	testSuiteDeleteDateFromGrpc string) (
	err error) {

	// Check if this TestSuite was stored with 'deletedDateIsSupported'
	supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap[deletedDateIsSupported] = true

	// Create 'testSuiteDeleteDate'
	testSuiteDeleteDate = testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteDeletionDate

	// Create the Hash of the Message
	testSuiteDeleteDateHash = sharedCode.HashSingleValue(testSuiteDeleteDate)

	return testSuiteDeleteDate, testSuiteDeleteDateHash, err

}

// Generates 'TestCasesInTestSuite' from gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteTypeMessageWhenLoading(
	supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct,
	testSuiteTypeMessageFromGrpc *fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteTypeMessage) (
	err error) {

	var existInMap bool

	// Check if this TestSuite was stored with 'testSuiteTypeIsSupported'
	_, existInMap = supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap[testSuiteTypeIsSupported]

	if existInMap == false {
		return err
	}

	// Create 'testSuiteTypeMessage'
	var testSuiteTypeMessage TestSuiteTypeStruct
	testSuiteTypeMessage = TestSuiteTypeStruct{
		TestSuiteType:     TestSuiteTypeType(testSuiteTypeMessageFromGrpc.GetTestSuiteType()),
		TestSuiteTypeName: TestSuiteTypeNameType(testSuiteTypeMessageFromGrpc.GetTestSuiteTypeName()),
	}

	testSuiteModel.savedTestSuiteUIModelBinding.TestSuiteType = testSuiteTypeMessage

	return err

}

// Generates 'TestSuiteImplementedFunctionsMap' from gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteImplementedFunctionsMapWhenLoading(
	supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct,
	testSuiteImplementedFunctionsMapFromGrpc map[int32]bool) (
	err error) {

	var existInMap bool

	// Check that this TestSuite has stored 'testSuiteImplementedFunctionsMapIsSupported'
	_, existInMap = supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap[testSuiteImplementedFunctionsMapIsSupported]
	if existInMap == false {
		errorId := "33ce8cb3-b557-4589-8136-4539b1d2828f"

		err = errors.New(fmt.Sprintf("'testSuiteImplementedFunctionsMap' is missing value 'testSuiteImplementedFunctionsMapIsSupported' [ErrorID: %s]",
			errorId))

		return err

	}

	// Loop 'testSuiteImplementedFunctionsMapFromGrpc' and convert from gRPC-message
	for testSuiteImplementedFunction, testSuiteImplementedFunctionValue := range testSuiteImplementedFunctionsMapFromGrpc {

		supportedTestSuiteDataToBeStored.
			testSuiteImplementedFunctionsMap[testSuiteImplementedFucntionsType(testSuiteImplementedFunction)] = testSuiteImplementedFunctionValue
	}

	return testSuiteImplementedFunctionsMapFromGrpc, testSuiteImplementedFunctionsMapHash, err

}
