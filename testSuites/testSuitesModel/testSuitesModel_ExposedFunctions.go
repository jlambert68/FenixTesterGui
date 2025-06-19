package testSuitesModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testCase/testCaseModel"
	uuidGenerator "github.com/google/uuid"
)

// GetTestSuiteUuid
// Gets the TestSuites Uuid
func (testSuiteModel *TestSuiteModelStruct) GetTestSuiteUuid() string {

	return testSuiteModel.testSuiteUuid

}

// GetCreatedByGcpLogin
// Gets the person that did log in towards GCP
func (testSuiteModel *TestSuiteModelStruct) GetCreatedByGcpLogin() string {

	return testSuiteModel.createdByGcpLogin

}

// GetCreatedByComputerLogin
// Gets the person that is logged into the computer
func (testSuiteModel *TestSuiteModelStruct) GetCreatedByComputerLogin() string {

	return testSuiteModel.createdByComputerLogin

}

// GetCreatedDate
// Gets the date when the TestSuite was first created
func (testSuiteModel *TestSuiteModelStruct) GetCreatedDate() string {

	return testSuiteModel.createdDate

}

// GetLastChangedByGcpLogin
// Gets the person that did log in towards GCP when TestSuite was last changed and saved
func (testSuiteModel *TestSuiteModelStruct) GetLastChangedByGcpLogin() string {

	return testSuiteModel.lastChangedByGcpLogin

}

// GetLastChangedByComputerLogin
// Gets the person that is logged into the computer when TestSuite was last changed and saved
func (testSuiteModel *TestSuiteModelStruct) GetLastChangedByComputerLogin() string {

	return testSuiteModel.lastChangedByComputerLogin

}

// GetLastChangedDate
// Gets the date when the TestSuite was last changed and saved
func (testSuiteModel *TestSuiteModelStruct) GetLastChangedDate() string {

	return testSuiteModel.lastChangedDate

}

// GenerateNewTestSuiteModelObject
// Generated s new TestSuiteModel-object
func GenerateNewTestSuiteModelObject(
	testCasesModel *testCaseModel.TestCasesModelsStruct) (
	newTestSuiteModel *TestSuiteModelStruct) {

	// Generate new TestSuite-UUID
	var testSuiteUuid string
	testSuiteUuid = uuidGenerator.New().String()

	// Generate the new TestSuiteModelStruct
	newTestSuiteModel = &TestSuiteModelStruct{
		testSuiteDeletionDate:         "",
		testSuiteUuid:                 testSuiteUuid,
		testSuiteName:                 "",
		testSuiteDescription:          "",
		testSuiteOwnerDomainUuid:      "",
		testSuiteOwnerDomainName:      "",
		createdByGcpLogin:             sharedCode.CurrentUserAuthenticatedTowardsGCP,
		createdByComputerLogin:        sharedCode.CurrentUserIdLogedInOnComputer,
		createdDate:                   "",
		lastChangedByGcpLogin:         "",
		lastChangedByComputerLogin:    "",
		lastChangedDate:               "",
		testSuiteExecutionEnvironment: "",
		testSuiteIsNew:                true,
		testSuiteMetaDataPtr: &TestSuiteMetaDataStruct{
			CurrentSelectedDomainUuid:                               "",
			TestSuiteMetaDataMessageJsonForTestSuiteWhenLastSaved:   nil,
			TestSuiteMetaDataMessageStructForTestSuiteWhenLastSaved: nil,
			MetaDataGroupsOrder:                                     nil,
			MetaDataGroupsMapPtr:                                    nil,
			SelectedTestSuiteMetaDataAsEntrySlice:                   nil,
		},
		lockValuesForOwnerDomainAndTestEnvironment: lockValuesForOwnerDomainAndTestEnvironmentStruct{
			OwnerDomainHasValue:     false,
			TestEnvironmentHasValue: false,
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
