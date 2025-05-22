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
	var tempTestCase TestCaseModelStruct

	// Get current TestCase
	tempTestCase, existsInMap = testCaseModel.TestCases[testCaseUuid]
	if existsInMap == false {

		errorId := "d067efb8-8f4f-44d6-9e77-ad07c77b5c3c"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return err
	}

	var testCaseMetaDataForDomain *TestCaseMetaDataForDomainStruct
	var metaDataGroupsPtr *[]*MetaDataGroupStruct

	// Convert original structure based on json into TestCase-structure for TestCaseMetaData
	metaDataGroupsPtr = buildMetaDataGroups(testCaseMetaDataForDomain)

	// MetaDataObject for a newly selected Domain
	var testCaseMetaData TestCaseMetaDataStruct
	testCaseMetaData = TestCaseMetaDataStruct{
		CurrentSelectedDomainUuid:                             selectedDomainUuid,
		TestCaseMetaDataMessageJsonForTestCaseWhenLastSaved:   nil,
		TestCaseMetaDataMessageStructForTestCaseWhenLastSaved: nil,
		MetaDataGroupsSlicePtr:                                metaDataGroupsPtr,
	}

	// Store MetaData-object back into TestCase
	tempTestCase.TestCaseMetaDataPtr = &testCaseMetaData

	// Store TestCase back into TestCases-model
	testCaseModel.TestCases[testCaseUuid] = tempTestCase

	return err
}

func buildMetaDataGroups(testCaseMetaDataForDomain *TestCaseMetaDataForDomainStruct) *[]*MetaDataGroupStruct {

	// The MetaDataGroups for the TestCase
	var metaDataGroups []*MetaDataGroupStruct
	metaDataGroups = make([]*MetaDataGroupStruct, 0, len(testCaseMetaDataForDomain.MetaDataGroups))

	// Loop over Domains MetaDataGroups to convert into TestCase-version of MetaDataGroups
	for _, tempMetaDataGroup := range testCaseMetaDataForDomain.MetaDataGroups {

		// build each MetaDataInGroupStruct slice
		var metaDataInGroup []*MetaDataInGroupStruct
		metaDataInGroup = make([]*MetaDataInGroupStruct, 0, len(tempMetaDataGroup.MetaDataInGroup))

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

			metaDataInGroup = append(metaDataInGroup, &MetaDataInGroupStruct{
				MetaDataName:                            metaDataItem.MetaDataName,
				SelectType:                              selType,
				Mandatory:                               mandatory,
				AvailableMetaDataValues:                 append([]string(nil), metaDataItem.MetaDataValues...),
				SelectedMetaDataValueForSingleSelect:    "",
				SelectedMetaDataValuesForMultiSelect:    nil,
				SelectedMetaDataValuesForMultiSelectMap: nil,
			})
		}

		var metaDataGroup *MetaDataGroupStruct
		metaDataGroup = &MetaDataGroupStruct{
			MetaDataGroupName:  tempMetaDataGroup.MetaDataGroupName,
			MetaDataInGroupPtr: &metaDataInGroup,
		}
		metaDataGroups = append(metaDataGroups, metaDataGroup)
	}

	return &metaDataGroups
}

// ConvertTestCaseMetaData converts the JSON‐parsed TestCaseMetaDataForDomainStruct
// into the GUI‐friendly slice *[]*MetaDataGroupStruct.
func ConvertTestCaseMetaData(testCaseMetaDataForDomain *TestCaseMetaDataForDomainStruct) *[]*MetaDataGroupStruct {
	groups := make([]*MetaDataGroupStruct, 0, len(testCaseMetaDataForDomain.MetaDataGroups))

	for _, g := range testCaseMetaDataForDomain.MetaDataGroups {
		// build slice of MetaDataInGroupStruct pointers
		inGroup := make([]*MetaDataInGroupStruct, 0, len(g.MetaDataInGroup))
		for _, md := range g.MetaDataInGroup {
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
				MetaDataGroupName:                       g.MetaDataGroupName,
				MetaDataName:                            md.MetaDataName,
				SelectType:                              selType,
				Mandatory:                               mandatory,
				AvailableMetaDataValues:                 append([]string(nil), md.MetaDataValues...),
				SelectedMetaDataValueForSingleSelect:    "",
				SelectedMetaDataValuesForMultiSelect:    nil,
				SelectedMetaDataValuesForMultiSelectMap: nil,
			}

			inGroup = append(inGroup, item)
		}

		group := &MetaDataGroupStruct{
			MetaDataGroupName:  g.MetaDataGroupName,
			MetaDataInGroupPtr: &inGroup,
		}
		groups = append(groups, group)
	}

	return &groups
}
