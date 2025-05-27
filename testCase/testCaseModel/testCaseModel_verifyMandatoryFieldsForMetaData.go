package testCaseModel

import (
	"FenixTesterGui/soundEngine"
	"fmt"
	"fyne.io/fyne/v2"
)

// Validates that all mandatory MetaData fields has values for specified DomainUuid
func (testCaseModel *TestCasesModelsStruct) verifyMandatoryFieldsForMetaData(
	domainUuid string,
	currentTestCasePtr *TestCaseModelStruct) (err error) {

	var mandatoryMetaDataFieldsMapKey string

	// Get Mandatory MetaData for DomainUuid
	var mandatoryMetaDataFieldsMap map[string]bool
	mandatoryMetaDataFieldsMap = make(map[string]bool)

	mandatoryMetaDataFieldsMap = testCaseModel.generateMandatoryMetaDataFieldsMap(domainUuid)

	// Ensure that there are MetaData selected by the user, to be able to loop it
	if currentTestCasePtr.TestCaseMetaDataPtr != nil &&
		*currentTestCasePtr.TestCaseMetaDataPtr.MetaDataGroupsMapPtr != nil {

		// Loop Users selected MetaData-parameters
		for metaDataGroupName, metaDataGroupPtr := range *currentTestCasePtr.TestCaseMetaDataPtr.MetaDataGroupsMapPtr {

			// Loop all MetaDataGroupItems for the MetaDataGroup
			for metaDataGroupItemName, _ := range *metaDataGroupPtr.MetaDataInGroupMapPtr {

				// Create map-key to be able to 'remove' from 'metaDataGroupItemName'.
				// Is done to be able to find if some MetaDataItems was missed by the user
				mandatoryMetaDataFieldsMapKey = fmt.Sprintf("%s-%s",
					metaDataGroupName,
					metaDataGroupItemName)

				// Try to Remove Key from 'metaDataGroupItemName'
				delete(mandatoryMetaDataFieldsMap, mandatoryMetaDataFieldsMapKey)

			}
		}
	}

	// Check if there are any MetaDataItems left in the mandatory map, 'mandatoryMetaDataFieldsMap'
	if len(mandatoryMetaDataFieldsMap) > 0 {

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
	metaDataForDomainPtr, existInMap = testCaseModel.TestCaseMetaDataForDomainsMap[ownerDomainUuid]

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
	return
}
