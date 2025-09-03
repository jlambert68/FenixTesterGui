package testSuitesModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testCase/testCaseModel"
	uuidGenerator "github.com/google/uuid"
	"github.com/jlambert68/Fast_BitFilter_MetaData/boolbits/boolbits"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/jlambert68/FenixScriptEngine/testDataEngine"
	"strconv"
)

// GetTestSuiteUuid
// Gets the TestSuites Uuid
func (testSuiteModel *TestSuiteModelStruct) GetTestSuiteUuid() string {

	return testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.testSuiteUuid

}

// GetTestSuiteVersion
// Gets the TestSuites Version
func (testSuiteModel *TestSuiteModelStruct) GetTestSuiteVersion() string {

	return strconv.Itoa(int(testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.testSuiteVersion))

}

// GetCreatedByGcpLogin
// Gets the person that did log in towards GCP
func (testSuiteModel *TestSuiteModelStruct) GetCreatedByGcpLogin() string {

	return testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.createdByGcpLogin

}

// GetCreatedByComputerLogin
// Gets the person that is logged into the computer
func (testSuiteModel *TestSuiteModelStruct) GetCreatedByComputerLogin() string {

	return testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.createdByComputerLogin

}

// GetCreatedDate
// Gets the date when the TestSuite was first created
func (testSuiteModel *TestSuiteModelStruct) GetCreatedDate() string {

	return testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.createdDate

}

// GetLastChangedByGcpLogin
// Gets the person that did log in towards GCP when TestSuite was last changed and saved
func (testSuiteModel *TestSuiteModelStruct) GetLastChangedByGcpLogin() string {

	return testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.lastChangedByGcpLogin

}

// GetLastChangedByComputerLogin
// Gets the person that is logged into the computer when TestSuite was last changed and saved
func (testSuiteModel *TestSuiteModelStruct) GetLastChangedByComputerLogin() string {

	return testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.lastChangedByComputerLogin

}

// GetLastChangedDate
// Gets the date when the TestSuite was last changed and saved
func (testSuiteModel *TestSuiteModelStruct) GetLastChangedDate() string {

	return testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.lastChangedDate

}

// createEmptyAndInitiatedTestSuiteModel
// Generates a fully initiated TestSuiteModelStruct
func createEmptyAndInitiatedTestSuiteModel(
	testCasesModel *testCaseModel.TestCasesModelsStruct) (
	emptyAndInitiatedTestSuiteModel *TestSuiteModelStruct) {

	tempMetaDataGroupsMap1 := make(map[string]*MetaDataGroupStruct)
	tempMetaDataGroupsMap2 := make(map[string]*MetaDataGroupStruct)
	tempMetaDataGroupsMap3 := make(map[string]*MetaDataGroupStruct)
	tempTestCasesInTestSuiteMap1 := make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.TestCaseInTestSuiteMessage)
	tempTestCasesInTestSuiteMap2 := make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.TestCaseInTestSuiteMessage)
	tempTestCasesInTestSuiteMap3 := make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.TestCaseInTestSuiteMessage)

	// Generate the new and initiated TestSuiteModelStruct
	emptyAndInitiatedTestSuiteModel = &TestSuiteModelStruct{
		testSuiteModelDataThatCanNotBeChangedFromUI: testSuiteModelDataThatCaNotBeChangedFromUIStruct{
			testSuiteUuid:              "",
			createdByGcpLogin:          sharedCode.CurrentUserAuthenticatedTowardsGCP,
			createdByComputerLogin:     sharedCode.CurrentUserIdLogedInOnComputer,
			createdDate:                "",
			lastChangedByGcpLogin:      "",
			lastChangedByComputerLogin: "",
			lastChangedDate:            "",
		},
		lockValuesForOwnerDomainAndTestEnvironment: lockValuesForOwnerDomainAndTestEnvironmentStruct{
			OwnerDomainHasValue:     false,
			TestEnvironmentHasValue: false,
			LockButtonHaBeenClicked: false,
		},
		savedTestSuiteUIModelBinding: TestSuiteUIModelBindingStruct{
			TestSuiteDeletionDate:         "",
			TestSuiteName:                 "",
			TestSuiteDescription:          "",
			TestSuiteOwnerDomainUuid:      "",
			TestSuiteOwnerDomainName:      "",
			TestSuiteExecutionEnvironment: "",
			TestSuiteIsNew:                false,
			TestSuiteTestDataHash:         "",
			TestDataPtr:                   &testDataEngine.TestDataForGroupObjectStruct{},
			TestSuiteMetaDataHash:         "",
			TestSuiteMetaDataPtr: &TestSuiteMetaDataStruct{
				CurrentSelectedDomainUuid:                               "",
				TestSuiteMetaDataMessageJsonForTestSuiteWhenLastSaved:   &fenixGuiTestCaseBuilderServerGrpcApi.TestCaseAndTestSuiteMetaDataForOneDomainMessage{},
				TestSuiteMetaDataMessageStructForTestSuiteWhenLastSaved: &TestSuiteMetaDataForDomainStruct{},
				MetaDataGroupsOrder:                                     []string{},
				MetaDataGroupsMapPtr:                                    &tempMetaDataGroupsMap1,
				SelectedTestSuiteMetaDataAsEntrySlice:                   []*boolbits.Entry{},
			},
			TestCasesInTestSuiteHash: "",
			TestCasesInTestSuitePtr: &TestCasesInTestSuiteStruct{
				TestCasesInTestSuiteMapPtr: &tempTestCasesInTestSuiteMap1,
			},
			TestSuiteTypeHash: "",
			TestSuiteType: TestSuiteTypeStruct{
				TestSuiteType:     TestSuiteTypeIsStandard,
				TestSuiteTypeName: TestSuiteTypeNameMap[TestSuiteTypeIsStandard],
			},
		},
		NoneSavedTestSuiteUIModelBinding: TestSuiteUIModelBindingStruct{
			TestSuiteDeletionDate:         "",
			TestSuiteName:                 "",
			TestSuiteDescription:          "",
			TestSuiteOwnerDomainUuid:      "",
			TestSuiteOwnerDomainName:      "",
			TestSuiteExecutionEnvironment: "",
			TestSuiteIsNew:                false,
			TestSuiteTestDataHash:         "",
			TestDataPtr:                   &testDataEngine.TestDataForGroupObjectStruct{},
			TestSuiteMetaDataHash:         "",
			TestSuiteMetaDataPtr: &TestSuiteMetaDataStruct{
				CurrentSelectedDomainUuid:                               "",
				TestSuiteMetaDataMessageJsonForTestSuiteWhenLastSaved:   &fenixGuiTestCaseBuilderServerGrpcApi.TestCaseAndTestSuiteMetaDataForOneDomainMessage{},
				TestSuiteMetaDataMessageStructForTestSuiteWhenLastSaved: &TestSuiteMetaDataForDomainStruct{},
				MetaDataGroupsOrder:                                     []string{},
				MetaDataGroupsMapPtr:                                    &tempMetaDataGroupsMap2,
				SelectedTestSuiteMetaDataAsEntrySlice:                   []*boolbits.Entry{},
			},
			TestCasesInTestSuiteHash: "",
			TestCasesInTestSuitePtr: &TestCasesInTestSuiteStruct{
				TestCasesInTestSuiteMapPtr: &tempTestCasesInTestSuiteMap2,
			},
			TestSuiteTypeHash: "",
			TestSuiteType: TestSuiteTypeStruct{
				TestSuiteType:     TestSuiteTypeIsStandard,
				TestSuiteTypeName: TestSuiteTypeNameMap[TestSuiteTypeIsStandard],
			},
		},
		TestSuiteUIModelBinding: TestSuiteUIModelBindingStruct{
			TestSuiteDeletionDate:         "",
			TestSuiteName:                 "",
			TestSuiteDescription:          "",
			TestSuiteOwnerDomainUuid:      "",
			TestSuiteOwnerDomainName:      "",
			TestSuiteExecutionEnvironment: "",
			TestSuiteIsNew:                false,
			TestSuiteTestDataHash:         "",
			TestDataPtr:                   &testDataEngine.TestDataForGroupObjectStruct{},
			TestSuiteMetaDataHash:         "",
			TestSuiteMetaDataPtr: &TestSuiteMetaDataStruct{
				CurrentSelectedDomainUuid:                               "",
				TestSuiteMetaDataMessageJsonForTestSuiteWhenLastSaved:   &fenixGuiTestCaseBuilderServerGrpcApi.TestCaseAndTestSuiteMetaDataForOneDomainMessage{},
				TestSuiteMetaDataMessageStructForTestSuiteWhenLastSaved: &TestSuiteMetaDataForDomainStruct{},
				MetaDataGroupsOrder:                                     []string{},
				MetaDataGroupsMapPtr:                                    &tempMetaDataGroupsMap3,
				SelectedTestSuiteMetaDataAsEntrySlice:                   []*boolbits.Entry{},
			},
			TestCasesInTestSuiteHash: "",
			TestCasesInTestSuitePtr: &TestCasesInTestSuiteStruct{
				TestCasesInTestSuiteMapPtr: &tempTestCasesInTestSuiteMap3,
			},
			TestSuiteTypeHash: "",
			TestSuiteType: TestSuiteTypeStruct{
				TestSuiteType:     TestSuiteTypeIsStandard,
				TestSuiteTypeName: TestSuiteTypeNameMap[TestSuiteTypeIsStandard],
			},
		},
		testCasesModel: testCasesModel,
	}

	return emptyAndInitiatedTestSuiteModel
}

// GenerateNewTestSuiteModelObject
// Generated s new TestSuiteModel-object
func GenerateNewTestSuiteModelObject(
	existingTestSuiteUuid string,
	testCasesModel *testCaseModel.TestCasesModelsStruct) (
	newTestSuiteModel *TestSuiteModelStruct) {

	// Generate new TestSuite-UUID, if not exist
	var testSuiteUuid string
	if existingTestSuiteUuid != "" {
		// Use existing TestSuiteUuid
		testSuiteUuid = existingTestSuiteUuid
	} else {
		// Generate a new TestSuiteUuid
		testSuiteUuid = uuidGenerator.New().String()
	}

	// Generate the new TestSuiteModelStruct
	newTestSuiteModel = createEmptyAndInitiatedTestSuiteModel(testCasesModel)
	newTestSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.testSuiteUuid = testSuiteUuid
	newTestSuiteModel.savedTestSuiteUIModelBinding.TestSuiteIsNew = true
	newTestSuiteModel.TestSuiteUIModelBinding.TestSuiteIsNew = true
	/*
		newTestSuiteModel = &TestSuiteModelStruct{
			testSuiteModelDataThatCanNotBeChangedFromUI: testSuiteModelDataThatCaNotBeChangedFromUIStruct{
				testSuiteUuid:              testSuiteUuid,
				createdByGcpLogin:          sharedCode.CurrentUserAuthenticatedTowardsGCP,
				createdByComputerLogin:     sharedCode.CurrentUserIdLogedInOnComputer,
				createdDate:                "",
				lastChangedByGcpLogin:      "",
				lastChangedByComputerLogin: "",
				lastChangedDate:            "",
			},
			lockValuesForOwnerDomainAndTestEnvironment: lockValuesForOwnerDomainAndTestEnvironmentStruct{
				OwnerDomainHasValue:     false,
				TestEnvironmentHasValue: false,
				LockButtonHaBeenClicked: false,
			},
			savedTestSuiteUIModelBinding: TestSuiteUIModelBindingStruct{
				TestSuiteDeletionDate:         "",
				TestSuiteName:                 "",
				TestSuiteDescription:          "",
				TestSuiteOwnerDomainUuid:      "",
				TestSuiteOwnerDomainName:      "",
				TestSuiteExecutionEnvironment: "",
				TestSuiteIsNew:                true,
				TestSuiteTestDataHash:         "",
				TestDataPtr:                   nil,
				TestSuiteMetaDataHash:         "",
				TestSuiteMetaDataPtr: &TestSuiteMetaDataStruct{
					CurrentSelectedDomainUuid:                               "",
					TestSuiteMetaDataMessageJsonForTestSuiteWhenLastSaved:   nil,
					TestSuiteMetaDataMessageStructForTestSuiteWhenLastSaved: nil,
					MetaDataGroupsOrder:                                     nil,
					MetaDataGroupsMapPtr:                                    nil,
					SelectedTestSuiteMetaDataAsEntrySlice:                   nil,
				},
			},
			NoneSavedTestSuiteUIModelBinding: TestSuiteUIModelBindingStruct{
				TestSuiteDeletionDate:         "",
				TestSuiteName:                 "",
				TestSuiteDescription:          "",
				TestSuiteOwnerDomainUuid:      "",
				TestSuiteOwnerDomainName:      "",
				TestSuiteExecutionEnvironment: "",
				TestSuiteIsNew:                false,
				TestSuiteTestDataHash:         "",
				TestDataPtr:                   nil,
				TestSuiteMetaDataHash:         "",
				TestSuiteMetaDataPtr: &TestSuiteMetaDataStruct{
					CurrentSelectedDomainUuid:                               "",
					TestSuiteMetaDataMessageJsonForTestSuiteWhenLastSaved:   nil,
					TestSuiteMetaDataMessageStructForTestSuiteWhenLastSaved: nil,
					MetaDataGroupsOrder:                                     nil,
					MetaDataGroupsMapPtr:                                    nil,
					SelectedTestSuiteMetaDataAsEntrySlice:                   nil,
				},
			},
			TestSuiteUIModelBinding: TestSuiteUIModelBindingStruct{
				TestSuiteDeletionDate:         "",
				TestSuiteName:                 "",
				TestSuiteDescription:          "",
				TestSuiteOwnerDomainUuid:      "",
				TestSuiteOwnerDomainName:      "",
				TestSuiteExecutionEnvironment: "",
				TestSuiteIsNew:                true,
				TestDataPtr:                   nil,
				TestSuiteMetaDataPtr: &TestSuiteMetaDataStruct{
					CurrentSelectedDomainUuid:                               "",
					TestSuiteMetaDataMessageJsonForTestSuiteWhenLastSaved:   nil,
					TestSuiteMetaDataMessageStructForTestSuiteWhenLastSaved: nil,
					MetaDataGroupsOrder:                                     nil,
					MetaDataGroupsMapPtr:                                    nil,
					SelectedTestSuiteMetaDataAsEntrySlice:                   nil,
				},
			},
			testCasesModel: testCasesModel,
		}

	*/

	return newTestSuiteModel

}

// OwnerDomainHasValue
// Store if OwnerDomain has any value selected by the user
func (testSuiteModel *TestSuiteModelStruct) OwnerDomainHasValue(hasValue bool) {
	testSuiteModel.lockValuesForOwnerDomainAndTestEnvironment.OwnerDomainHasValue = hasValue

	// Clear LockButton when false
	if hasValue == false {
		testSuiteModel.lockValuesForOwnerDomainAndTestEnvironment.LockButtonHaBeenClicked = false
	}

}

// TestEnvironmentHasValue
// Store if TestEnvironmentHasValue has any value selected by the user
func (testSuiteModel *TestSuiteModelStruct) TestEnvironmentHasValue(hasValue bool) {
	testSuiteModel.lockValuesForOwnerDomainAndTestEnvironment.TestEnvironmentHasValue = hasValue

	// Clear LockButton when false
	if hasValue == false {
		testSuiteModel.lockValuesForOwnerDomainAndTestEnvironment.LockButtonHaBeenClicked = false
	}
}

// LockButtonHasBeenClicked
// Store if LockButton has been clicked by the user
func (testSuiteModel *TestSuiteModelStruct) LockButtonHasBeenClicked() {
	testSuiteModel.lockValuesForOwnerDomainAndTestEnvironment.LockButtonHaBeenClicked = true
}

// DoBothOwnerDomainAndTestEnvironmentHaveValues
// Do both of OwnerDomain and TestEnvironmentHasValue have their values selected by the user
func (testSuiteModel *TestSuiteModelStruct) DoBothOwnerDomainAndTestEnvironmentHaveValues() (hasValue bool) {

	hasValue = testSuiteModel.lockValuesForOwnerDomainAndTestEnvironment.TestEnvironmentHasValue &&
		testSuiteModel.lockValuesForOwnerDomainAndTestEnvironment.OwnerDomainHasValue

	return hasValue
}

// HasLockButtonBeenClickedAndBothOwnerDomainAndTestEnvironmentHaveValues
// Has Locked been clicked and both of  OwnerDomain and TestEnvironmentHasValue have their values selected by the user
func (testSuiteModel *TestSuiteModelStruct) HasLockButtonBeenClickedAndBothOwnerDomainAndTestEnvironmentHaveValues() (hasValue bool) {

	hasValue = testSuiteModel.lockValuesForOwnerDomainAndTestEnvironment.TestEnvironmentHasValue &&
		testSuiteModel.lockValuesForOwnerDomainAndTestEnvironment.OwnerDomainHasValue &&
		testSuiteModel.lockValuesForOwnerDomainAndTestEnvironment.LockButtonHaBeenClicked

	return hasValue
}
