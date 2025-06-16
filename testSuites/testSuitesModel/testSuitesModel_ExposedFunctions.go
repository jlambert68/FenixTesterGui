package testSuitesModel

import uuidGenerator "github.com/google/uuid"

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
func GenerateNewTestSuiteModelObject() (newTestSuiteModel *TestSuiteModelStruct) {

	// Generate new TestSuite-UUID
	var testSuiteUuid string
	testSuiteUuid = uuidGenerator.New().String()

	// Generate the new TestSuiteModelStruct
	newTestSuiteModel = &TestSuiteModelStruct{
		testSuiteDeletionDate:    "",
		testSuiteUuid:            testSuiteUuid,
		testSuiteName:            "",
		testSuiteDescription:     "",
		testSuiteOwnerDomainUuid: "",
		testSuiteOwnerDomainName: "",
		testSuiteIsNew:           true,

		TestSuiteUIModelBinding: TestSuiteUIModelBindingStruct{
			TestSuiteDeletionDate:    "",
			TestSuiteName:            "",
			TestSuiteDescription:     "",
			TestSuiteOwnerDomainUuid: "",
			TestSuiteIsNew:           true,
		},
	}

	return newTestSuiteModel

}
