package testSuitesModel

// IsTestSuiteChanged
// Checks if the TestSuite content has been changed from last saved occasion
func (testSuiteModel *TestSuiteModelStruct) IsTestSuiteChanged() (testSuiteIsChanged bool) {

	// Check 'TestSuiteDeletionDate'
	if testSuiteModel.testSuiteDeletionDate != testSuiteModel.TestSuiteUIModelBinding.TestSuiteDeletionDate {
		testSuiteIsChanged = true

		return testSuiteIsChanged
	}

	// Check 'TestSuiteName'
	if testSuiteModel.testSuiteName != testSuiteModel.TestSuiteUIModelBinding.TestSuiteName {
		testSuiteIsChanged = true

		return testSuiteIsChanged
	}

	// Check 'TestSuiteDescription'
	if testSuiteModel.testSuiteDescription != testSuiteModel.TestSuiteUIModelBinding.TestSuiteDescription {
		testSuiteIsChanged = true

		return testSuiteIsChanged
	}

	// Check 'TestSuiteOwnerDomainUuid'
	if testSuiteModel.testSuiteOwnerDomainUuid != testSuiteModel.TestSuiteUIModelBinding.TestSuiteOwnerDomainUuid {
		testSuiteIsChanged = true

		return testSuiteIsChanged
	}

	// Check 'TestSuiteIsNew'
	//if testSuiteModel.testSuiteIsNew != testSuiteModel.TestSuiteUIModelBinding.TestSuiteIsNew {
	//	testSuiteIsChanged = true

	//	return testSuiteIsChanged
	//}

	// Check 'TestSuiteExecutionEnvironment'
	if testSuiteModel.testSuiteExecutionEnvironment != testSuiteModel.TestSuiteUIModelBinding.
		TestSuiteExecutionEnvironment {
		testSuiteIsChanged = true

		return testSuiteIsChanged
	}

	return testSuiteIsChanged
}
