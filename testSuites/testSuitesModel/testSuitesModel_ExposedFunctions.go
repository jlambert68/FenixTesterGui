package testSuitesModel

import uuidGenerator "github.com/google/uuid"

// GetTestSuiteUuid
// Gets the TestSuites Uuid
func (testSuiteModel *TestSuiteModelStruct) GetTestSuiteUuid() (testSuiteUuid string) {

	return testSuiteModel.testSuiteUuid

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
