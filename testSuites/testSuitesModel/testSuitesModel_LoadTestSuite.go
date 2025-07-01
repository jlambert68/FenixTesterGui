package testSuitesModel

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/jlambert68/FenixScriptEngine/testDataEngine"
	"log"
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

	// Extract a list of what was saved from gRPC-message.
	//Used to ensure that older versions of the TestSuite can later be loaded when new functionality has been added to the client
	var supportedTestSuiteDataToBeStored testSuiteImplementedFunctionsToBeStoredStruct
	supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap, err = testSuiteModel.
		extractTestSuiteImplementedFunctionsMap(detailedTestSuiteResponse.GetDetailedTestSuite().
			GetTestSuiteImplementedFunctionsMap())

	// Generates 'TestSuiteBasicInformation' from gRPC-message
	err = testSuiteModel.generateTestSuiteBasicInformationMessageWhenLoading(
		&supportedTestSuiteDataToBeStored,
		detailedTestSuiteResponse.GetDetailedTestSuite().GetTestSuiteBasicInformation(),
		detailedTestSuiteResponse.GetDetailedTestSuite().GetUpdatedByAndWhen(),
		detailedTestSuiteResponse.GetDetailedTestSuite().GetMessageHash())

	if err != nil {
		errorId := "06c6bf54-c165-4410-ac8e-b741c9fd5932"
		err = errors.New(fmt.Sprintf("Couldn't generate 'TestSuiteBasicInformation' for TestSuite = '%s'. Error = '%s' [ErrorID: %s]",
			testSuiteUuid,
			err.Error(),
			errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return err
	}

	// Generates 'UsersChosenTestDataForTestSuiteMessage' from gRPC-message
	err = testSuiteModel.generateTestSuiteTestDataMessageWhenLoading(
		&supportedTestSuiteDataToBeStored,
		detailedTestSuiteResponse.GetDetailedTestSuite().GetTestSuiteTestData())

	if err != nil {
		errorId := "09552494-9c22-4120-9476-bf4c89d8bd53"
		err = errors.New(fmt.Sprintf("Couldn't generate 'UsersChosenTestDataForTestSuiteMessage' for TestSuite = '%s'. Error = '%s' [ErrorID: %s]",
			testSuiteUuid,
			err.Error(),
			errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return err
	}

	// Generates 'TestSuitePreview' from gRPC-message
	err = testSuiteModel.generateTestSuitePreviewMessageWhenLoading(
		&supportedTestSuiteDataToBeStored,
		detailedTestSuiteResponse.GetDetailedTestSuite().GetTestSuitePreview())

	if err != nil {
		errorId := "a5f9b5e3-e45a-483d-a59b-2ac7c9450836"
		err = errors.New(fmt.Sprintf("Couldn't generate 'TestSuitePreview' for TestSuite = '%s'. Error = '%s' [ErrorID: %s]",
			testSuiteUuid,
			err.Error(),
			errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return err
	}

	// Generates 'TestSuiteMetaData' from gRPC-message
	err = testSuiteModel.generateTestSuiteMetaDataMessageWhenLoading(
		&supportedTestSuiteDataToBeStored,
		detailedTestSuiteResponse.GetDetailedTestSuite().GetTestSuiteMetaData())

	if err != nil {
		errorId := "a085470e-10db-4eff-ab24-bf58fafbbb1c"
		err = errors.New(fmt.Sprintf("Couldn't generate 'TestSuiteMetaData' for TestSuite = '%s'. Error = '%s' [ErrorID: %s]",
			testSuiteUuid,
			err.Error(),
			errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return err
	}

	// Generates 'TestCasesInTestSuite' from gRPC-message
	err = testSuiteModel.generateTestCasesInTestSuiteMessageWhenLoading(
		&supportedTestSuiteDataToBeStored,
		detailedTestSuiteResponse.GetDetailedTestSuite().GetTestCasesInTestSuite())

	if err != nil {
		errorId := "f123cec8-db36-46af-ab92-1825ecd57cbb"
		err = errors.New(fmt.Sprintf("Couldn't generate 'TestCasesInTestSuite' for TestSuite = '%s'. Error = '%s' [ErrorID: %s]",
			testSuiteUuid,
			err.Error(),
			errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return err
	}

	// Generates 'TestSuiteDeleteData' from gRPC-message
	err = testSuiteModel.generateTestSuiteDeleteDateMessageWhenLoading(
		&supportedTestSuiteDataToBeStored,
		detailedTestSuiteResponse.GetDetailedTestSuite().GetDeletedDate())

	if err != nil {
		errorId := "cfea21a1-149a-4993-8234-c6f8b3a5316a"
		err = errors.New(fmt.Sprintf("Couldn't generate 'TestSuiteDeleteData' for TestSuite = '%s'. Error = '%s' [ErrorID: %s]",
			testSuiteUuid,
			err.Error(),
			errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return err
	}

	// Due to that DeepCopier can't copy 'TestDataPtr' it needs to be handled separate(I don't know why)
	var tempTestDataPtr *testDataEngine.TestDataForGroupObjectStruct
	tempTestDataPtr = testSuiteModel.savedTestSuiteUIModelBinding.TestDataPtr

	testSuiteModel.savedTestSuiteUIModelBinding.TestDataPtr = nil

	// Copy data from 'savedTestSuiteUIModelBinding' to 'TestSuiteUIModelBinding' using deep copy
	err = copier.CopyWithOption(&testSuiteModel.TestSuiteUIModelBinding, &testSuiteModel.savedTestSuiteUIModelBinding, copier.Option{DeepCopy: true})
	if err != nil {

		errorID := "34791be0-1117-4a46-9d28-479977689cd6"

		errorMsg := fmt.Sprintf("error copying 'savedTestSuiteUIModelBinding' using 'copier'. error = '%s' [ErrorID: %s]",
			err.Error(),
			errorID)

		log.Fatalln(errorMsg)
	}

	// Set back 'TestDataPtr' and put 'pointer-copy' in 'savedTestSuiteUIModelBinding'
	testSuiteModel.savedTestSuiteUIModelBinding.TestDataPtr = tempTestDataPtr
	testSuiteModel.TestSuiteUIModelBinding.TestDataPtr = tempTestDataPtr

	// Set that Domain and TestEnvironment is locked
	testSuiteModel.lockValuesForOwnerDomainAndTestEnvironment.OwnerDomainHasValue = true
	testSuiteModel.lockValuesForOwnerDomainAndTestEnvironment.TestEnvironmentHasValue = true
	testSuiteModel.lockValuesForOwnerDomainAndTestEnvironment.LockButtonHaBeenClicked = true

	return err

}

// Extract 'TestSuiteImplementedFunctionsMap' from gRPC-message
func (testSuiteModel *TestSuiteModelStruct) extractTestSuiteImplementedFunctionsMap(
	testSuiteImplementedFunctionsGrpc map[int32]bool) (
	testSuiteImplementedFunctionsMap map[testSuiteImplementedFunctionsType]bool,
	err error) {

	testSuiteImplementedFunctionsMap = make(map[testSuiteImplementedFunctionsType]bool)

	// Loop gRPC-map and add to local map
	for testSuiteImplementedFunction, testSuiteImplementedFunctionValue := range testSuiteImplementedFunctionsGrpc {
		testSuiteImplementedFunctionsMap[testSuiteImplementedFunctionsType(testSuiteImplementedFunction)] = testSuiteImplementedFunctionValue
	}

	return testSuiteImplementedFunctionsMap, err
}

// Generates 'TestSuiteBasicInformation' from gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteBasicInformationMessageWhenLoading(
	supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct,
	testSuiteBasicInformation *fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteBasicInformationMessage,
	updatedByAndWhenMessage *fenixGuiTestCaseBuilderServerGrpcApi.UpdatedByAndWhenMessage,
	messageHash string) (
	err error) {

	var existInMap bool

	// Check if this TestSuite was stored with 'testSuiteBasicInformationIsSupported'
	_, existInMap = supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap[testSuiteBasicInformationIsSupported]

	if existInMap == false {
		// Do nothing

		return err
	}

	// Copy values
	testSuiteModel.savedTestSuiteUIModelBinding.TestSuiteName = testSuiteBasicInformation.GetTestSuiteName()
	testSuiteModel.savedTestSuiteUIModelBinding.TestSuiteDescription = testSuiteBasicInformation.GetTestSuiteDescription()
	testSuiteModel.savedTestSuiteUIModelBinding.TestSuiteOwnerDomainUuid = testSuiteBasicInformation.GetDomainUuid()
	testSuiteModel.savedTestSuiteUIModelBinding.TestSuiteOwnerDomainName = testSuiteBasicInformation.GetDomainName()
	testSuiteModel.savedTestSuiteUIModelBinding.TestSuiteExecutionEnvironment = testSuiteBasicInformation.GetTestSuiteExecutionEnvironment()
	testSuiteModel.savedTestSuiteUIModelBinding.TestSuiteIsNew = false

	testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.testSuiteUuid = testSuiteBasicInformation.GetTestSuiteUuid()
	testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.testSuiteVersion = testSuiteBasicInformation.GetTestSuiteVersion()
	testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.createdByGcpLogin = updatedByAndWhenMessage.GetCreatedByGcpLogin()
	testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.createdByComputerLogin = updatedByAndWhenMessage.GetCreatedByComputerLogin()
	testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.createdDate = updatedByAndWhenMessage.GetCreatedDate()
	testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.lastChangedByGcpLogin = updatedByAndWhenMessage.GetGCPAuthenticatedUser()
	testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.lastChangedByComputerLogin = updatedByAndWhenMessage.GetUserIdOnComputer()
	testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.lastChangedDate = updatedByAndWhenMessage.GetUpdateTimeStamp()
	testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.testSuiteSavedMessageHash = messageHash

	return err

}

// Generates 'UsersChosenTestDataForTestSuiteMessage' from gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteTestDataMessageWhenLoading(
	supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct,
	testSuiteTestData *fenixGuiTestCaseBuilderServerGrpcApi.UsersChosenTestDataForTestSuiteMessage) (
	err error) {

	var existInMap bool

	// Check if this TestSuite was stored with 'testSuiteTestDataIsSupported'
	_, existInMap = supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap[testSuiteTestDataIsSupported]

	if existInMap == false {
		// Do nothing

		return err
	}

	// Generate users TestData for TestSuite
	var chosenTestDataPointsPerGroupMap map[testDataEngine.TestDataPointGroupNameType]*testDataEngine.TestDataPointNameMapType
	chosenTestDataPointsPerGroupMap = make(map[testDataEngine.TestDataPointGroupNameType]*testDataEngine.TestDataPointNameMapType)

	var testDataPointGroups []testDataEngine.TestDataPointGroupNameType

	var testData *testDataEngine.TestDataForGroupObjectStruct
	testData = &testDataEngine.TestDataForGroupObjectStruct{
		TestDataPointGroups:             nil,
		TestDataPointsForAGroup:         nil,
		ChosenTestDataPointsPerGroupMap: chosenTestDataPointsPerGroupMap,
		ShouldUpdateMainWindow: testDataEngine.ResponseChannelStruct{
			ShouldBeUpdated:        false,
			TestDataPointGroupName: "",
		},
	}

	if testSuiteTestData != nil {
		// User has TestData stored for the TestSuite

		// Loop all Groups with TestDataPoints in gRPC-message
		for testDataGroupNameGrpc, testDataGroupGrpc := range testSuiteTestData.ChosenTestDataPointsPerGroupMap {

			var testDataPointNameMap map[testDataEngine.TestDataValueNameType]*[]*testDataEngine.DataPointTypeForGroupsStruct
			testDataPointNameMap = make(map[testDataEngine.TestDataValueNameType]*[]*testDataEngine.DataPointTypeForGroupsStruct)

			var testDataPointNameMapAsObject testDataEngine.TestDataPointNameMapType

			for testDataPointNameGrpc, testDataPointGrpc := range testDataGroupGrpc.ChosenTestDataRowsPerTestDataPointMap {

				var dataPointTypeForGroups []*testDataEngine.DataPointTypeForGroupsStruct

				for _, testDataRowGrpc := range testDataPointGrpc.TestDataRows {

					// Initiate the maps in the struct, only 'selectedTestDataPointUuidMap' is filled with values
					var searchResultDataPointUuidMap map[testDataEngine.TestDataPointRowUuidType]testDataEngine.TestDataPointRowUuidStruct
					searchResultDataPointUuidMap = make(map[testDataEngine.TestDataPointRowUuidType]testDataEngine.TestDataPointRowUuidStruct)
					var availableTestDataPointUuidMap map[testDataEngine.TestDataPointRowUuidType]testDataEngine.TestDataPointRowUuidStruct
					availableTestDataPointUuidMap = make(map[testDataEngine.TestDataPointRowUuidType]testDataEngine.TestDataPointRowUuidStruct)
					var selectedTestDataPointUuidMap map[testDataEngine.TestDataPointRowUuidType]testDataEngine.TestDataPointRowUuidStruct
					selectedTestDataPointUuidMap = make(map[testDataEngine.TestDataPointRowUuidType]testDataEngine.TestDataPointRowUuidStruct)

					// Create the RowSummary for each row for 'dataPointTypeForGroup'
					for _, testDataPointRowValueSummaryGrpc := range testDataRowGrpc.TestDataPointRowValueSummaryMap {

						var testDataPointRowValuesSummary testDataEngine.TestDataPointRowUuidStruct
						testDataPointRowValuesSummary = testDataEngine.TestDataPointRowUuidStruct{
							TestDataPointRowUuid:          testDataEngine.TestDataPointRowUuidType(testDataPointRowValueSummaryGrpc.GetTestDataPointRowUuid()),
							TestDataPointRowValuesSummary: testDataEngine.TestDataPointRowValuesSummaryType(testDataPointRowValueSummaryGrpc.GetTestDataPointRowValuesSummary()),
						}

						// Add the RowUUid and the values summary to the map
						selectedTestDataPointUuidMap[testDataEngine.TestDataPointRowUuidType(testDataPointRowValueSummaryGrpc.GetTestDataPointRowUuid())] = testDataPointRowValuesSummary

					}

					// Create the 'dataPointTypeForGroup'
					var dataPointTypeForGroup *testDataEngine.DataPointTypeForGroupsStruct
					dataPointTypeForGroup = &testDataEngine.DataPointTypeForGroupsStruct{
						TestDataDomainUuid:            testDataEngine.TestDataDomainUuidType(testDataRowGrpc.GetTestDataDomainUuid()),
						TestDataDomainName:            testDataEngine.TestDataDomainNameType(testDataRowGrpc.GetTestDataDomainName()),
						TestDataAreaUuid:              testDataEngine.TestDataAreaUuidType(testDataRowGrpc.GetTestDataAreaUuid()),
						TestDataAreaName:              testDataEngine.TestDataAreaNameType(testDataRowGrpc.GetTestDataAreaName()),
						TestDataPointName:             testDataEngine.TestDataValueNameType(testDataRowGrpc.GetTestDataPointName()),
						SearchResultDataPointUuidMap:  searchResultDataPointUuidMap,
						AvailableTestDataPointUuidMap: availableTestDataPointUuidMap,
						SelectedTestDataPointUuidMap:  selectedTestDataPointUuidMap,
					}

					// Add 'dataPointTypeForGroup' to slice of 'dataPointTypeForGroups'
					dataPointTypeForGroups = append(dataPointTypeForGroups, dataPointTypeForGroup)

				}

				// Add the slice of TestDataRows to the map for TestDataPoints
				testDataPointNameMap[testDataEngine.TestDataValueNameType(testDataPointNameGrpc)] = &dataPointTypeForGroups

			}

			// Add to slice of GroupNames
			testDataPointGroups = append(testDataPointGroups, testDataEngine.TestDataPointGroupNameType(testDataGroupNameGrpc))

			// Move Map into object
			testDataPointNameMapAsObject = testDataPointNameMap

			chosenTestDataPointsPerGroupMap[testDataEngine.TestDataPointGroupNameType(testDataGroupNameGrpc)] = &testDataPointNameMapAsObject

		}

		testData.ChosenTestDataPointsPerGroupMap = chosenTestDataPointsPerGroupMap

		testData.TestDataPointGroups = testDataPointGroups

		testSuiteModel.savedTestSuiteUIModelBinding.TestDataPtr = testData

	} else {
		// User has no TestData stored for the TestSuite
		testData.ChosenTestDataPointsPerGroupMap = chosenTestDataPointsPerGroupMap

		testSuiteModel.savedTestSuiteUIModelBinding.TestDataPtr = testData
	}

	// Update The Hash for the TestSuite's TestData
	testSuiteModel.savedTestSuiteUIModelBinding.TestSuiteTestDataHash = testSuiteTestData.GetHashOfThisMessageWithEmptyHashField()

	return err

}

// Generates 'TestSuitePreview' from gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuitePreviewMessageWhenLoading(
	supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct,
	testSuitePreview *fenixGuiTestCaseBuilderServerGrpcApi.TestSuitePreviewMessage) (
	err error) {

	var existInMap bool

	// Check if this TestSuite was stored with 'testSuitePreviewIsSupported'
	_, existInMap = supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap[testSuitePreviewIsSupported]

	if existInMap == false {
		// Do nothing

		return err
	}

	return err

}

// Generates 'TestSuiteMetaData' from gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteMetaDataMessageWhenLoading(
	supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct,
	testSuiteMetaData *fenixGuiTestCaseBuilderServerGrpcApi.UserSpecifiedTestSuiteMetaDataMessage) (
	err error) {

	var existInMap bool

	// Check if this TestSuite was stored with 'testSuiteMetaDataIsSupported'
	_, existInMap = supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap[testSuiteMetaDataIsSupported]

	if existInMap == false {
		// Do nothing

		return err
	}

	// Generate the TestSuiteMetaData for the TestSuite-model from the gRPC-data
	var tempMetaDataGroupsMap map[string]*MetaDataGroupStruct
	var tempMetaDataGroupsOrder []string
	tempMetaDataGroupsMap = make(map[string]*MetaDataGroupStruct)

	// Loop MetaDataGroups in gPRC-data
	for tempMetaDataGroupNameFromGrpc, tempMetaDataGroupFromGrpc := range testSuiteMetaData.GetMetaDataGroupsMap() {

		tempMetaDataGroupsOrder = append(tempMetaDataGroupsOrder, tempMetaDataGroupNameFromGrpc)

		var tempMetaDataInGroupOrder []string

		var tempMetaDataInGroupMap map[string]*MetaDataInGroupStruct
		tempMetaDataInGroupMap = make(map[string]*MetaDataInGroupStruct)

		// Loop MetaDataGroupItems in MetaDataGroup
		for tempMetaDataGroupItemNameFromGrpc, tempMetaDataGroupItemFromGrpc := range tempMetaDataGroupFromGrpc.GetMetaDataInGroupMap() {

			tempMetaDataInGroupOrder = append(tempMetaDataInGroupOrder, tempMetaDataGroupItemNameFromGrpc)

			var tempMetaDataInGroup MetaDataInGroupStruct
			tempMetaDataInGroup = MetaDataInGroupStruct{
				MetaDataGroupName:                          tempMetaDataGroupNameFromGrpc,
				MetaDataName:                               tempMetaDataGroupItemNameFromGrpc,
				SelectType:                                 MetaDataSelectType(tempMetaDataGroupItemFromGrpc.SelectType),
				Mandatory:                                  tempMetaDataGroupItemFromGrpc.IsMandatory,
				AvailableMetaDataValues:                    tempMetaDataGroupItemFromGrpc.AvailableMetaDataValues,
				SelectedMetaDataValueForSingleSelect:       tempMetaDataGroupItemFromGrpc.SelectedMetaDataValueForSingleSelect,
				SelectedMetaDataValuesForMultiSelect:       tempMetaDataGroupItemFromGrpc.SelectedMetaDataValuesForMultiSelect,
				SelectedMetaDataValuesForMultiSelectMapPtr: nil,
			}

			// Generate the map holding the Selected values for the multi select
			var tempSelectedMetaDataValuesMap map[string]string
			tempSelectedMetaDataValuesMap = make(map[string]string)

			for tempMetaDataItemKeyFromGrpc, tempMetaDataItemValueFromGrpc := range tempMetaDataGroupItemFromGrpc.
				SelectedMetaDataValuesForMultiSelectMap {
				tempSelectedMetaDataValuesMap[tempMetaDataItemKeyFromGrpc] = tempMetaDataItemValueFromGrpc
			}

			// Store selected values map in main message
			tempMetaDataInGroup.SelectedMetaDataValuesForMultiSelectMapPtr = &tempSelectedMetaDataValuesMap

			// Add MetaDataGroupItem to 'MetaDataInGroupMap'
			tempMetaDataInGroupMap[tempMetaDataGroupItemNameFromGrpc] = &tempMetaDataInGroup

		}

		// Create the var MetaDataGroup-object
		var tempMetaDataGroup MetaDataGroupStruct
		tempMetaDataGroup = MetaDataGroupStruct{
			MetaDataGroupName:     tempMetaDataGroupNameFromGrpc,
			MetaDataInGroupOrder:  tempMetaDataInGroupOrder,
			MetaDataInGroupMapPtr: &tempMetaDataInGroupMap,
		}

		// Add  MetaDataGroup-object to 'tempMetaDataGroupsMap'
		tempMetaDataGroupsMap[tempMetaDataGroupNameFromGrpc] = &tempMetaDataGroup

	}

	// Create the full 'TestSuiteMetaData-object'
	var tempTestSuiteMetaData TestSuiteMetaDataStruct
	tempTestSuiteMetaData = TestSuiteMetaDataStruct{
		CurrentSelectedDomainUuid: testSuiteMetaData.GetCurrentSelectedDomainUuid(),
		TestSuiteMetaDataMessageJsonForTestSuiteWhenLastSaved: &fenixGuiTestCaseBuilderServerGrpcApi.TestCaseAndTestSuiteMetaDataForOneDomainMessage{
			DomainUuid:              testSuiteMetaData.GetCurrentSelectedDomainUuid(),
			DomainName:              testSuiteMetaData.GetCurrentSelectedDomainName(),
			TestCaseMetaDataAsJson:  "",
			TestSuiteMetaDataAsJson: "",
		},
		TestSuiteMetaDataMessageStructForTestSuiteWhenLastSaved: &TestSuiteMetaDataForDomainStruct{
			MetaDataGroups: nil,
		},
		MetaDataGroupsOrder:                   tempMetaDataGroupsOrder,
		MetaDataGroupsMapPtr:                  &tempMetaDataGroupsMap,
		SelectedTestSuiteMetaDataAsEntrySlice: nil,
	}

	// Add converted TestSuiteMetaData to 'TestSuiteModel'
	testSuiteModel.savedTestSuiteUIModelBinding.TestSuiteMetaDataPtr = &tempTestSuiteMetaData

	return err

}

// Generates 'TestCasesInTestSuite' from gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestCasesInTestSuiteMessageWhenLoading(
	supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct,
	testCasesInTestSuiteFromGrpc *fenixGuiTestCaseBuilderServerGrpcApi.TestCasesInTestSuiteMessage) (
	err error) {

	var existInMap bool

	// Check if this TestSuite was stored with 'testCasesInTestSuiteIsSupported'
	_, existInMap = supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap[testCasesInTestSuiteIsSupported]

	if existInMap == false {
		// Do nothing

		return err
	}

	return err

}

// Generates 'TestSuiteDeleteData' from gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteDeleteDateMessageWhenLoading(
	supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct,
	testSuiteDeleteDateFromGrpc string) (
	err error) {

	var existInMap bool

	// Check if this TestSuite was stored with 'deletedDateIsSupported'
	_, existInMap = supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap[deletedDateIsSupported]

	if existInMap == false {
		// Do nothing

		return err
	}

	// Create 'testSuiteDeleteDate'
	testSuiteModel.savedTestSuiteUIModelBinding.TestSuiteDeletionDate = testSuiteDeleteDateFromGrpc

	return err

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
		// Do nothing

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
			testSuiteImplementedFunctionsMap[testSuiteImplementedFunctionsType(testSuiteImplementedFunction)] = testSuiteImplementedFunctionValue
	}

	return err

}
