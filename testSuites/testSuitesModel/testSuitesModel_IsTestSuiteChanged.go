package testSuitesModel

// IsTestSuiteChanged
// Checks if the TestSuite content has been changed from last saved occasion
func (testSuiteModel *TestSuiteModelStruct) IsTestSuiteChanged() (isTestSuiteChanged bool) {

	// Check 'TestSuiteDeletionDate'
	if testSuiteModel.testSuiteDeletionDate != testSuiteModel.TestSuiteUIModelBinding.TestSuiteDeletionDate {
		isTestSuiteChanged = true
	}

	// Check 'TestSuiteName'
	if testSuiteModel.testSuiteName != testSuiteModel.TestSuiteUIModelBinding.TestSuiteName {
		isTestSuiteChanged = true
	}

	// Check 'TestSuiteDescription'
	if testSuiteModel.testSuiteDescription != testSuiteModel.TestSuiteUIModelBinding.TestSuiteDescription {
		isTestSuiteChanged = true
	}

	// Check 'TestSuiteOwnerDomainUuid'
	if testSuiteModel.testSuiteOwnerDomainUuid != testSuiteModel.TestSuiteUIModelBinding.TestSuiteOwnerDomainUuid {
		isTestSuiteChanged = true
	}

	// Check 'TestSuiteIsNew'
	if testSuiteModel.testSuiteIsNew != testSuiteModel.TestSuiteUIModelBinding.TestSuiteIsNew {
		isTestSuiteChanged = true
	}

	// Check 'TestSuiteExecutionEnvironment'
	if testSuiteModel.testSuiteExecutionEnvironment != testSuiteModel.TestSuiteUIModelBinding.
		TestSuiteExecutionEnvironment {
		isTestSuiteChanged = true
	}

	return isTestSuiteChanged
}
