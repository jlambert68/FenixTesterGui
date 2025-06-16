package testSuitesModel

import "github.com/jlambert68/FenixScriptEngine/testDataEngine"

// Pointer to the object holding all TestSuites (visible in a Tab)
var TestSuitesModelPtr *TestSuitesModelStruct

// Holding all information about TestSuites
type TestSuitesModelStruct struct {
	TestSuitesMapPtr *map[string]*TestSuiteModelStruct // Map-key = 'TestSuiteUuid'
}

// TestSuiteModelStruct
// Holdning all information about a specific TestSuite
type TestSuiteModelStruct struct {
	testSuiteDeletionDate         string // Date for when the TestSuite is deleted. Can be date in the future
	testSuiteUuid                 string // The TestSuites Uuid
	testSuiteName                 string // The TestSuites Name
	testSuiteDescription          string // A description for the TestSuite
	testSuiteOwnerDomainUuid      string // The Uuid for the Domain that owns the TestSuite
	testSuiteOwnerDomainName      string // The Name of the Domain that owns the TestSuite
	createdByGcpLogin             string // The person that did log in towards GCP
	createdByComputerLogin        string // The person that is logged into the computer
	createdDate                   string // The date when the TestSuite was first created
	lastChangedByGcpLogin         string // The person that did log in towards GCP when TestSuite was last changed and saved
	lastChangedByComputerLogin    string // The person that is logged into the computer when TestSuite was last changed and saved
	lastChangedDate               string // The date when the TestSuite was last changed and saved
	testSuiteExecutionEnvironment string // The execution environment where the TestSuite will be executed

	testSuiteIsNew bool // Indicates that if this a new or existing TestSuite

	// TestSuite UI components bindings
	TestSuiteUIModelBinding TestSuiteUIModelBindingStruct
}

// TestSuiteUIModelBindingStruct
// Holding bindings to textboxes, dropDown, in UI
type TestSuiteUIModelBindingStruct struct {
	TestSuiteDeletionDate         string // Date that TestSuite will be deleted
	TestSuiteName                 string // The Name of the TestSuite
	TestSuiteDescription          string // A description for the TestSuite
	TestSuiteOwnerDomainUuid      string // The Uuid for the Domain that owns the TestSuite
	TestSuiteOwnerDomainName      string // The Name for the Domain that owns the TestSuite
	TestSuiteExecutionEnvironment string // The execution environment where the TestSuite will be executed

	TestSuiteIsNew bool
	TestDataPtr    *testDataEngine.TestDataForGroupObjectStruct
}
