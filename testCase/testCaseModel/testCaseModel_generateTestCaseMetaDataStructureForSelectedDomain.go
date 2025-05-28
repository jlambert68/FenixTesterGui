package testCaseModel

import (
	sharedCode "FenixTesterGui/common_code"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
)

// GenerateTestCaseMetaDataStructureForSelectedDomain - Verify if the Hash for the TestCase is the same as the one in the database
func (testCaseModel *TestCasesModelsStruct) GenerateTestCaseMetaDataStructureForSelectedDomain(
	testCaseUuid string,
	selectedDomainUuid string) (
	err error) {

	var existsInMap bool
	var tempTestCasePtr *TestCaseModelStruct

	// Get current TestCase
	tempTestCasePtr, existsInMap = testCaseModel.TestCasesMap[testCaseUuid]
	if existsInMap == false {

		errorId := "d067efb8-8f4f-44d6-9e77-ad07c77b5c3c"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCasesMap [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return err
	}

	var testCaseMetaDataForDomain *TestCaseMetaDataForDomainStruct
	var metaDataGroupsMapPtr *map[string]*MetaDataGroupStruct

	// Convert original structure based on json into TestCase-structure for TestCaseMetaData
	metaDataGroupsMapPtr = buildMetaDataGroups(testCaseMetaDataForDomain)

	// Get map keys
	var metaDataGroupsMapKeys []string
	for mapKey := range *metaDataGroupsMapPtr {
		metaDataGroupsMapKeys = append(metaDataGroupsMapKeys, mapKey)
	}

	// MetaDataObject for a newly selected Domain
	var testCaseMetaData TestCaseMetaDataStruct
	testCaseMetaData = TestCaseMetaDataStruct{
		CurrentSelectedDomainUuid:                             selectedDomainUuid,
		TestCaseMetaDataMessageJsonForTestCaseWhenLastSaved:   nil,
		TestCaseMetaDataMessageStructForTestCaseWhenLastSaved: nil,
		MetaDataGroupsOrder:                                   metaDataGroupsMapKeys,
		MetaDataGroupsMapPtr:                                  metaDataGroupsMapPtr,
	}

	// Store MetaData-object back into TestCase
	tempTestCasePtr.TestCaseMetaDataPtr = &testCaseMetaData

	// Store TestCase back into TestCasesMap-model
	testCaseModel.TestCasesMap[testCaseUuid] = tempTestCasePtr

	return err
}

func buildMetaDataGroups(testCaseMetaDataForDomain *TestCaseMetaDataForDomainStruct) *map[string]*MetaDataGroupStruct {

	// The MetaDataGroups for the TestCase
	var metaDataGroupsMap map[string]*MetaDataGroupStruct
	metaDataGroupsMap = make(map[string]*MetaDataGroupStruct)

	var metaDataGroupItemsMapKeys []string

	// Loop over Domains MetaDataGroups to convert into TestCase-version of MetaDataGroups
	for _, tempMetaDataGroup := range testCaseMetaDataForDomain.MetaDataGroups {

		// build each MetaDataInGroupStruct slice
		var metaDataInGroupMap map[string]*MetaDataInGroupStruct
		metaDataInGroupMap = make(map[string]*MetaDataInGroupStruct)

		// Loop over the Items in the MetaDataInGroup
		for _, metaDataItem := range tempMetaDataGroup.MetaDataInGroup {

			// parse select type
			var selType MetaDataSelectType
			switch metaDataItem.SelectType {

			case "SingleSelect":
				selType = MetaDataSelectType_SingleSelect

			case "MultiSelect":
				selType = MetaDataSelectType_MultiSelect

			default:
				// fallback or panic based on your needs

				sharedCode.Logger.WithFields(logrus.Fields{
					"id":                                "71f65ae6-4291-4274-b943-e5a40cf5792a",
					"metaDataItem":                      metaDataItem,
					"tempMetaDataGroup.MetaDataInGroup": tempMetaDataGroup.MetaDataInGroup,
					"tempMetaDataGroup":                 tempMetaDataGroup,
					"testCaseMetaDataForDomain":         testCaseMetaDataForDomain,
				}).Fatalln("unknown SelectType: " + metaDataItem.SelectType)
			}

			// parse mandatory
			var mandatory bool
			mandatory = false
			if metaDataItem.Mandatory == "True" {
				mandatory = true
			}

			metaDataInGroupMap[metaDataItem.MetaDataName] = &MetaDataInGroupStruct{
				MetaDataGroupName:                          tempMetaDataGroup.MetaDataGroupName,
				MetaDataName:                               metaDataItem.MetaDataName,
				SelectType:                                 selType,
				Mandatory:                                  mandatory,
				AvailableMetaDataValues:                    append([]string(nil), metaDataItem.MetaDataValues...),
				SelectedMetaDataValueForSingleSelect:       "",
				SelectedMetaDataValuesForMultiSelect:       nil,
				SelectedMetaDataValuesForMultiSelectMapPtr: nil,
			}

			// Get map keys
			metaDataGroupItemsMapKeys = append(metaDataGroupItemsMapKeys, metaDataItem.MetaDataName)
		}

		var metaDataGroup *MetaDataGroupStruct
		metaDataGroup = &MetaDataGroupStruct{
			MetaDataGroupName:     tempMetaDataGroup.MetaDataGroupName,
			MetaDataInGroupOrder:  metaDataGroupItemsMapKeys,
			MetaDataInGroupMapPtr: &metaDataInGroupMap,
		}

		metaDataGroupsMap[metaDataGroup.MetaDataGroupName] = metaDataGroup

	}

	return &metaDataGroupsMap
}

// ConvertTestCaseMetaData converts the JSON‐parsed TestCaseMetaDataForDomainStruct
// into the GUI‐friendly slice *[]*MetaDataGroupStruct.
func ConvertTestCaseMetaData(testCaseMetaDataForDomain *TestCaseMetaDataForDomainStruct) (
	*map[string]*MetaDataGroupStruct, []string) {
	groupsMap := make(map[string]*MetaDataGroupStruct)

	var tempMetaDataGroupsOrder []string

	for _, g := range testCaseMetaDataForDomain.MetaDataGroups {

		// Create the slice to put the orde of Items in
		var tempMetaDataItemsInGroupOrder []string

		tempMetaDataGroupsOrder = append(tempMetaDataGroupsOrder, g.MetaDataGroupName)

		// build slice of MetaDataInGroupStruct pointers
		inGroupMap := make(map[string]*MetaDataInGroupStruct)
		for _, md := range g.MetaDataInGroup {

			tempMetaDataItemsInGroupOrder = append(tempMetaDataItemsInGroupOrder, md.MetaDataName)

			// map the SelectType string to your enum
			var selType MetaDataSelectType
			switch md.SelectType {
			case "SingleSelect":
				selType = MetaDataSelectType_SingleSelect
			case "MultiSelect":
				selType = MetaDataSelectType_MultiSelect
			default:
				selType = MetaDataSelectType_NotSelected
			}

			// parse the Mandatory flag
			mandatory := false
			if md.Mandatory == "True" {
				mandatory = true
			}

			item := &MetaDataInGroupStruct{
				MetaDataGroupName:                          g.MetaDataGroupName,
				MetaDataName:                               md.MetaDataName,
				SelectType:                                 selType,
				Mandatory:                                  mandatory,
				AvailableMetaDataValues:                    append([]string(nil), md.MetaDataValues...),
				SelectedMetaDataValueForSingleSelect:       "",
				SelectedMetaDataValuesForMultiSelect:       nil,
				SelectedMetaDataValuesForMultiSelectMapPtr: nil,
			}

			inGroupMap[md.MetaDataName] = item

		}

		group := &MetaDataGroupStruct{
			MetaDataGroupName:     g.MetaDataGroupName,
			MetaDataInGroupOrder:  tempMetaDataItemsInGroupOrder,
			MetaDataInGroupMapPtr: &inGroupMap,
		}
		groupsMap[g.MetaDataGroupName] = group
	}

	return &groupsMap, tempMetaDataGroupsOrder
}
