package testSuitesModel

// Pointer to the object holding all TestSuites (visible in a Tab)
var TestSuitesModelPtr *TestSuitesModelStruct

// Holding all information about TestSuites
type TestSuitesModelStruct struct {
	TestSuitesMapPtr *map[string]*TestSuiteModelStruct // Map-key = 'TestSuiteUuid'
}

// TestSuiteModelStruct
// Holdning all information about a specific TestSuite
type TestSuiteModelStruct struct {
	testSuiteDeletionDate    string // Date for when the TestSuite is deleted. Can be date in the future
	testSuiteUuid            string // The TestSuites Uuid
	testSuiteName            string // The TestSuites Name
	testSuiteDescription     string // A description for the TestSuite
	testSuiteOwnerDomainUuid string // The Uuid for the Domain that owns the TestSuite
	testSuiteOwnerDomainName string // The Name of the Domain that owns the TestSuite
	testSuiteIsNew           bool   // Indicates that if this an new or existing TestSuite

	// TestSuite UI components bindings
	TestSuiteUIModelBinding TestSuiteUIModelBindingStruct
}

// Holding bindings to textboxes, dropDown, in UI
type TestSuiteUIModelBindingStruct struct {
	TestSuiteDeletionDate    string // Date that TestSuite will be deleted
	TestSuiteName            string
	TestSuiteDescription     string
	TestSuiteOwnerDomainUuid string
	TestSuiteIsNew           bool
}
