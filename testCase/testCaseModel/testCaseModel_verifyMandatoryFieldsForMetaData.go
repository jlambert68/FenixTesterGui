package testCaseModel

import (
	"FenixTesterGui/soundEngine"
	"fmt"
	"fyne.io/fyne/v2"
	"log"
)

// Validates that all mandatory MetaData fields has values for specified DomainUuid
func (testCaseModel *TestCasesModelsStruct) verifyMandatoryFieldsForMetaData(
	domainUuid string,
	currentTestCasePtr *TestCaseModelStruct,
	shouldBeSaved bool) (err error) {

	var mandatoryMetaDataFieldsMapKey string

	// Get Mandatory MetaData for DomainUuid
	var mandatoryMetaDataFieldsMap map[string]bool
	mandatoryMetaDataFieldsMap = make(map[string]bool)

	mandatoryMetaDataFieldsMap = testCaseModel.generateMandatoryMetaDataFieldsMap(domainUuid)

	// Ensure that there are MetaData selected by the user, to be able to loop it
	if currentTestCasePtr.TestCaseMetaDataPtr != nil {

		var tempTestCaseMetaData TestCaseMetaDataStruct
		tempTestCaseMetaData = *currentTestCasePtr.TestCaseMetaDataPtr

		if tempTestCaseMetaData.MetaDataGroupsMapPtr != nil {

			// Loop Users selected MetaData-parameters
			for metaDataGroupName, metaDataGroupPtr := range *currentTestCasePtr.TestCaseMetaDataPtr.MetaDataGroupsMapPtr {

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

						errorId := "57f87a09-3287-4127-a478-0675c7606386"
						log.Fatalln(fmt.Sprintf("MetaDataSelectType not implemented. [ErrorID: %s]", errorId))
					}

					// Try to Remove Key from 'metaDataGroupItemName'
					delete(mandatoryMetaDataFieldsMap, mandatoryMetaDataFieldsMapKey)
				}
			}
		}
	}

	// Check if there are any MetaDataItems left in the mandatory map, 'mandatoryMetaDataFieldsMap'
	if len(mandatoryMetaDataFieldsMap) > 0 && shouldBeSaved == true {

		err = fmt.Errorf("mandatory MetaDataFields are not set")

		// Generate string with Mandatory fields
		var mandatoryMetaDataFieldsString string
		for mandatoryMetaDataFieldsMapKey, _ := range mandatoryMetaDataFieldsMap {

			if len(mandatoryMetaDataFieldsString) == 0 {
				mandatoryMetaDataFieldsString = mandatoryMetaDataFieldsMapKey
			} else {
				mandatoryMetaDataFieldsString = mandatoryMetaDataFieldsString + ", " + mandatoryMetaDataFieldsMapKey
			}

		}

		// Notify the user

		// Trigger System Notification sound
		soundEngine.PlaySoundChannel <- soundEngine.UserNeedToRespondSound

		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title: "Save TestCase, failed",
			Content: fmt.Sprintf("All mandatory fields are not set: \n%s",
				mandatoryMetaDataFieldsString),
		})

	}

	return err
}

// Generates a map containing all mandatory MetaDataFields for the selected DomainUuid
// ResponseMap map['GroupName-GroupItemName']bool
func (testCaseModel *TestCasesModelsStruct) generateMandatoryMetaDataFieldsMap(
	ownerDomainUuid string) (
	mandatoryMetaDataFieldsMap map[string]bool) {

	var existInMap bool
	var mandatoryMetaDataFieldsMapKey string

	// Initiate the response map
	mandatoryMetaDataFieldsMap = make(map[string]bool)

	// Extract Owner Domains MetaData
	var metaDataForDomainPtr *TestCaseMetaDataForDomainsForMapStruct
	metaDataForDomainPtr, existInMap = testCaseModel.TestCaseMetaDataForDomains.TestCaseMetaDataForDomainsMap[ownerDomainUuid]

	if existInMap == false {
		// No MetaData fields, so no mandatory fields
		return mandatoryMetaDataFieldsMap
	}

	// Loop all MetaDa for the specified domain
	for _, metaDataGroup := range metaDataForDomainPtr.TestCaseMetaDataForDomainPtr.MetaDataGroups {

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
