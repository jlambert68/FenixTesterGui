package testSuitesModel

// ConvertTestSuiteMetaData converts the JSON‐parsed TestCaseMetaDataForDomainStruct
// into the GUI‐friendly slice *[]*MetaDataGroupStruct.
func ConvertTestSuiteMetaData(testSuiteMetaDataForDomain *TestSuiteMetaDataForDomainStruct) (
	*map[string]*MetaDataGroupStruct, []string) {
	groupsMap := make(map[string]*MetaDataGroupStruct)

	var tempMetaDataGroupsOrder []string

	for _, g := range testSuiteMetaDataForDomain.MetaDataGroups {

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
