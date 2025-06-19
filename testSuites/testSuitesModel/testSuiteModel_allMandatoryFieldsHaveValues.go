package testSuitesModel

import (
	"fmt"
	"log"
)

// checkIfAllMandatoryFieldsHaveValues
// Checks if all mandatory fields in the TestSuite has gor any value
func (testSuiteModel *TestSuiteModelStruct) checkIfAllMandatoryFieldsHaveValues() (
	allMandatoryFieldsHaveValues bool,
	mandatoryFieldsHaveValuesNotificationText string) {

	// Check 'TestSuiteName'
	if len(testSuiteModel.TestSuiteUIModelBinding.TestSuiteName) == 0 {

		allMandatoryFieldsHaveValues = false
		mandatoryFieldsHaveValuesNotificationText = "TestSuite Name"

		return allMandatoryFieldsHaveValues, mandatoryFieldsHaveValuesNotificationText
	}

	// Check 'TestSuiteOwnerDomainUuid'
	if len(testSuiteModel.TestSuiteUIModelBinding.TestSuiteOwnerDomainUuid) == 0 {

		allMandatoryFieldsHaveValues = false
		mandatoryFieldsHaveValuesNotificationText = "Owner Domain "

		return allMandatoryFieldsHaveValues, mandatoryFieldsHaveValuesNotificationText
	}

	// Check 'TestSuiteExecutionEnvironment'
	if len(testSuiteModel.TestSuiteUIModelBinding.TestSuiteExecutionEnvironment) == 0 {

		allMandatoryFieldsHaveValues = false
		mandatoryFieldsHaveValuesNotificationText = "Execution Environment "

		return allMandatoryFieldsHaveValues, mandatoryFieldsHaveValuesNotificationText
	}

	// Check 'TestSuiteMetaData'
	var allMandatoryMetaDataFieldsHasValues bool
	allMandatoryMetaDataFieldsHasValues,
		mandatoryFieldsHaveValuesNotificationText = testSuiteModel.verifyMandatoryFieldsForMetaData()
	if allMandatoryMetaDataFieldsHasValues == false {

		allMandatoryFieldsHaveValues = false

		return allMandatoryFieldsHaveValues, mandatoryFieldsHaveValuesNotificationText

	}

	return true, ""
}

// Validates that all mandatory MetaData fields has values for specified DomainUuid
func (testSuiteModel *TestSuiteModelStruct) verifyMandatoryFieldsForMetaData() (
	allMandatoryFieldsHasValues bool,
	mandatoryMetaDataFieldsString string) {

	var mandatoryMetaDataFieldsMapKey string

	// Get Mandatory MetaData for DomainUuid
	var mandatoryMetaDataFieldsMap map[string]bool
	mandatoryMetaDataFieldsMap = make(map[string]bool)

	mandatoryMetaDataFieldsMap = testSuiteModel.generateMandatoryMetaDataFieldsMap()

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
							// Create map-key to be able to 'remove' from 'metaDataGroupItemName'.
							// Is done to be able to find if some MetaDataItems was missed by the user
							mandatoryMetaDataFieldsMapKey = fmt.Sprintf("%s-%s",
								metaDataGroupName,
								metaDataGroupItemName)
						}

					case MetaDataSelectType_MultiSelect:
						if len(tempMetaDataInGroupPtr.SelectedMetaDataValuesForMultiSelect) > 0 {
							// Create map-key to be able to 'remove' from 'metaDataGroupItemName'.
							// Is done to be able to find if some MetaDataItems was missed by the user
							mandatoryMetaDataFieldsMapKey = fmt.Sprintf("%s-%s",
								metaDataGroupName,
								metaDataGroupItemName)
						}

					default:

						errorId := "f42c7078-756f-43ee-87a0-351aec872bbe"
						log.Fatalln(fmt.Sprintf("MetaDataSelectType not implemented. [ErrorID: %s]", errorId))
					}

					// Try to Remove Key from 'metaDataGroupItemName'
					delete(mandatoryMetaDataFieldsMap, mandatoryMetaDataFieldsMapKey)
				}
			}
		}
	}

	// Check if there are any MetaDataItems left in the mandatory map, 'mandatoryMetaDataFieldsMap'
	if len(mandatoryMetaDataFieldsMap) > 0 {

		// Generate string with Mandatory fields
		for tempMandatoryMetaDataFieldsMapKey, _ := range mandatoryMetaDataFieldsMap {

			if len(mandatoryMetaDataFieldsString) == 0 {
				mandatoryMetaDataFieldsString = tempMandatoryMetaDataFieldsMapKey
			} else {
				mandatoryMetaDataFieldsString = mandatoryMetaDataFieldsString + ", " + tempMandatoryMetaDataFieldsMapKey
			}

		}

	} else {
		allMandatoryFieldsHasValues = true
	}

	return allMandatoryFieldsHasValues,
		mandatoryMetaDataFieldsString
}

// Generates a map containing all mandatory MetaDataFields for the selected DomainUuid
// ResponseMap map['GroupName-GroupItemName']bool
func (testSuiteModel *TestSuiteModelStruct) generateMandatoryMetaDataFieldsMap() (
	mandatoryMetaDataFieldsMap map[string]bool) {

	var existInMap bool
	var mandatoryMetaDataFieldsMapKey string

	// Initiate the response map
	mandatoryMetaDataFieldsMap = make(map[string]bool)

	// Extract Owner Domains MetaData
	var testSuiteMetaDataForDomainsForMapStructmetaDataForDomainPtr *TestSuiteMetaDataForDomainsForMapStruct
	testSuiteMetaDataForDomainsForMapStructmetaDataForDomainPtr, existInMap = TestSuitesModelPtr.TestSuiteMetaDataForDomains.TestSuiteMetaDataForDomainsMap[testSuiteModel.TestSuiteUIModelBinding.TestSuiteOwnerDomainUuid]

	if existInMap == false {
		// No MetaData fields, so no mandatory fields
		return mandatoryMetaDataFieldsMap
	}

	// Loop all MetaDa for the specified domain
	for _, metaDataGroup := range testSuiteMetaDataForDomainsForMapStructmetaDataForDomainPtr.TestSuiteMetaDataForDomainPtr.MetaDataGroups {

		// Loop MetaDataGroupItems for the MetaDataGroup
		for _, metaDataGroupItem := range metaDataGroup.MetaDataInGroup {

			// Check if MetaDataGroupItem is mandatory
			if metaDataGroupItem.Mandatory == "True" {

				// MetaDataGroupItem is mandatory so add to map with mandatory MetaDataGroupItems
				mandatoryMetaDataFieldsMapKey = fmt.Sprintf("%s-%s",
					metaDataGroup.MetaDataGroupName,
					metaDataGroupItem.MetaDataName)

				mandatoryMetaDataFieldsMap[mandatoryMetaDataFieldsMapKey] = true
			}
		}
	}
	return mandatoryMetaDataFieldsMap
}
