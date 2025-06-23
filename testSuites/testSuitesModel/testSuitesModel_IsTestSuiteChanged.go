package testSuitesModel

import (
	sharedCode "FenixTesterGui/common_code"
	"fmt"
	"log"
)

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

	// Check 'TestSuiteExecutionEnvironment'
	if testSuiteModel.testSuiteExecutionEnvironment != testSuiteModel.TestSuiteUIModelBinding.
		TestSuiteExecutionEnvironment {
		testSuiteIsChanged = true

		return testSuiteIsChanged
	}

	// First update 'TestSuiteMetaDataHash'
	var tempTestSuiteMetaDataHash string
	tempTestSuiteMetaDataHash = testSuiteModel.generateTesSuiteMetaDataHash()
	testSuiteModel.TestSuiteUIModelBinding.TestSuiteMetaDataHash = tempTestSuiteMetaDataHash

	// Second Check changes for 'TestSuiteMetaDataHash'
	if testSuiteModel.testSuiteMetaDataHash != testSuiteModel.TestSuiteUIModelBinding.
		TestSuiteMetaDataHash {
		testSuiteIsChanged = true

		return testSuiteIsChanged
	}

	// First update 'TestSuiteTesDataHash'
	var tempTestSuiteTestDataHash string
	tempTestSuiteTestDataHash = testSuiteModel.generateTesSuiteTestDataHash()
	testSuiteModel.TestSuiteUIModelBinding.TestSuiteTesDataHash = tempTestSuiteTestDataHash

	// Second Check changes for 'TestSuiteTesDataHash'
	if testSuiteModel.testSuiteTestDataHash != testSuiteModel.TestSuiteUIModelBinding.
		TestSuiteTesDataHash {
		testSuiteIsChanged = true

		return testSuiteIsChanged
	}

	return testSuiteIsChanged

}

func (testSuiteModel *TestSuiteModelStruct) generateTesSuiteMetaDataHash() (
	testSuiteMetaDataHash string) {

	var valuesToHash []string

	var mandatoryMetaDataFieldsMapKey string

	// Ensure that there are MetaData selected by the user, to be able to loop it
	if testSuiteModel.TestSuiteUIModelBinding.TestSuiteMetaDataPtr != nil {

		var tempTestSuiteMetaData TestSuiteMetaDataStruct
		tempTestSuiteMetaData = *testSuiteModel.TestSuiteUIModelBinding.TestSuiteMetaDataPtr

		if tempTestSuiteMetaData.MetaDataGroupsMapPtr != nil {

			// Loop Users selected MetaData-parameters
			for metaDataGroupName, metaDataGroupPtr := range *tempTestSuiteMetaData.MetaDataGroupsMapPtr {

				// Loop all MetaDataGroupItems for the MetaDataGroup
				for metaDataGroupItemName, tempMetaDataInGroupPtr := range *metaDataGroupPtr.MetaDataInGroupMapPtr {

					mandatoryMetaDataFieldsMapKey = ""

					switch tempMetaDataInGroupPtr.SelectType {
					case MetaDataSelectType_SingleSelect:
						if len(tempMetaDataInGroupPtr.SelectedMetaDataValueForSingleSelect) > 0 {
							mandatoryMetaDataFieldsMapKey = fmt.Sprintf("%s-%s-%s",
								metaDataGroupName,
								metaDataGroupItemName,
								tempMetaDataInGroupPtr.SelectedMetaDataValueForSingleSelect)

							// Add Selected value to slice of selected values
							valuesToHash = append(valuesToHash, mandatoryMetaDataFieldsMapKey)
						}

					case MetaDataSelectType_MultiSelect:
						if len(tempMetaDataInGroupPtr.SelectedMetaDataValuesForMultiSelect) > 0 {
							// Loop selected values
							for _, selectedValue := range tempMetaDataInGroupPtr.SelectedMetaDataValuesForMultiSelect {

								mandatoryMetaDataFieldsMapKey = fmt.Sprintf("%s-%s-%s",
									metaDataGroupName,
									metaDataGroupItemName,
									selectedValue,
								)

								// Add Selected value to slice of selected values
								valuesToHash = append(valuesToHash, mandatoryMetaDataFieldsMapKey)
							}
						}

					default:

						errorId := "cb77bda6-ac2d-4075-803c-d06b0696231d"
						log.Fatalln(fmt.Sprintf("MetaDataSelectType not implemented. [ErrorID: %s]", errorId))
					}

				}
			}
		}
	}

	// Hash slice with values
	testSuiteMetaDataHash = sharedCode.HashValues(valuesToHash, false)

	return testSuiteMetaDataHash
}

func (testSuiteModel *TestSuiteModelStruct) generateTesSuiteTestDataHash() (
	testSuiteMetaDataHash string) {

	var testDataValueToHash string
	var valuesToHash []string

	// Loop all TestDataGroups
	for _, testDataGroup := range testSuiteModel.TestSuiteUIModelBinding.TestDataPtr.ListTestDataGroups() {

		// For each TestDataGroup loop all its TestDataPoints
		for _, testDataPoint := range testSuiteModel.TestSuiteUIModelBinding.TestDataPtr.
			ListTestDataGroupPointsForAGroup(testDataGroup) {

			// For each TestDataPoint loop all its TestDataRows
			for _, testDataRow := range testSuiteModel.TestSuiteUIModelBinding.TestDataPtr.
				ListTestDataRowsForAGroupPoint(testDataGroup, testDataPoint) {

				// Create value to be hashed
				testDataValueToHash = fmt.Sprintf("%s-%s-%s",
					testDataGroup,
					testDataPoint,
					testDataRow)

				// Add Values to slice of values that will be hashed
				valuesToHash = append(valuesToHash, testDataValueToHash)

			}
		}
	}

	// Hash slice with values
	testSuiteMetaDataHash = sharedCode.HashValues(valuesToHash, false)

	return testSuiteMetaDataHash

}
