package testCaseUI

import (
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Generate the MetaData Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) GenerateMetaDataAreaForTestCase(
	testCaseUuid string,
	domainUuidToGetMetaDataFor string) (
	testCaseMetaDataArea fyne.CanvasObject,
	accordion *widget.Accordion,
	err error) {

	var metaDataArea fyne.CanvasObject
	var metaDataAccordionItem *widget.AccordionItem

	var existsInMap bool

	//
	if len(domainUuidToGetMetaDataFor) > 0 {

		/*
			// Get current TestCase-UI-model
			_, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]

			if existsInMap == true {
				errorId := "bcb9d984-3106-42b6-9c23-288ec6d26224"
				err = errors.New(fmt.Sprintf("testcase-UI-model with sourceUuid '%s' allready exist in 'TestCasesUiModelMap' [ErrorID: %s]", testCaseUuid, errorId))

				return nil, err
			}


		*/
		// Get the MetaDataGroups depending on Domain
		var metaDataGroupsPtr *[]*testCaseModel.MetaDataGroupStruct
		var testCaseMetaDataForDomainsMap map[string]*testCaseModel.TestCaseMetaDataForDomainsForMapStruct
		var testCaseMetaDataForDomainPtr *testCaseModel.TestCaseMetaDataForDomainsForMapStruct
		var testCaseMetaDataForDomain testCaseModel.TestCaseMetaDataForDomainsForMapStruct
		testCaseMetaDataForDomainsMap = testCasesUiCanvasObject.TestCasesModelReference.TestCaseMetaDataForDomainsMap
		testCaseMetaDataForDomainPtr, existsInMap = testCaseMetaDataForDomainsMap[domainUuidToGetMetaDataFor]
		if existsInMap == false {

			errorId := "4e38ab00-2ffe-48ce-9c79-0f027227c4a4"
			err = errors.New(fmt.Sprintf("Domain with Uuid '%s' doesn't exist in'testCaseMetaDataForDomainsMap'. Should never happen [ErrorID: %s]",
				domainUuidToGetMetaDataFor, errorId))

			return nil, nil, err
		}

		testCaseMetaDataForDomain = *testCaseMetaDataForDomainPtr
		metaDataGroupsPtr = testCaseModel.ConvertTestCaseMetaData(testCaseMetaDataForDomain.TestCaseMetaDataForDomainPtr)

		// Get Object holding Selected data for TestCase
		var testCase testCaseModel.TestCaseModelStruct
		testCase, existsInMap = testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]
		if existsInMap == false {

			errorId := "7feb5fb3-0640-4866-8bd4-40c0011ceff1"
			err = errors.New(fmt.Sprintf("TestCase with Uuid '%s' doesn't exist in'TestCasesmap'. Should never happen [ErrorID: %s]",
				testCaseUuid, errorId))

			return nil, nil, err
		}

		// Get pointer to Structure holding selected values in TestCase
		var metaDataGroupInTestCasePtr *testCaseModel.TestCaseMetaDataStruct
		metaDataGroupInTestCasePtr = testCase.TestCaseMetaDataPtr
		if metaDataGroupInTestCasePtr == nil {
			metaDataGroupInTestCasePtr = &testCaseModel.TestCaseMetaDataStruct{
				CurrentSelectedDomainUuid:                             domainUuidToGetMetaDataFor,
				TestCaseMetaDataMessageJsonForTestCaseWhenLastSaved:   nil,
				TestCaseMetaDataMessageStructForTestCaseWhenLastSaved: nil,
				MetaDataGroupsSlicePtr:                                nil,
			}
		}

		// Generate TestCaseMeta-UI-object
		var metaDataGroupsAsCanvasObject fyne.CanvasObject
		metaDataGroupsAsCanvasObject = buildGUIFromMetaDataGroupsSlice(
			metaDataGroupsPtr,

			metaDataGroupInTestCasePtr)

		// Save back 'metaDataGroupInTestCasePtr' into the TestCase
		testCase.TestCaseMetaDataPtr = metaDataGroupInTestCasePtr

		// Save back the TestCase in TestCases-map
		testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid] = testCase

		myContainer := container.NewBorder(nil, nil, nil, nil, metaDataGroupsAsCanvasObject)
		fmt.Println("MinSize", myContainer.MinSize())
		fmt.Println("Size", myContainer.Size())

		if myContainer.MinSize().Width > 600 || myContainer.MinSize().Height > 600 {
			myContainerScroll := container.NewScroll(myContainer)
			myContainerScroll.SetMinSize(fyne.NewSize(600, 600))

			// Create an Accordion item for the MetaData
			metaDataAccordionItem = widget.NewAccordionItem("TestCase MetaData", myContainerScroll)

		} else {

			// Create an Accordion item for the MetaData
			metaDataAccordionItem = widget.NewAccordionItem("TestCase MetaData", myContainer)
		}

		metaDataAccordionItem.Detail.Hide()
	} else {

		myContainer := container.New(layout.NewGridLayout(1), widget.NewLabel("MetaData is available when 'Owner Domain' is selected"))

		// Create an Accordion item for the MetaData
		metaDataAccordionItem = widget.NewAccordionItem("TestCase MetaData", myContainer)
	}

	accordion = widget.NewAccordion(metaDataAccordionItem)

	// Create the VBox-container that will be returned
	metaDataArea = container.NewVBox(accordion, widget.NewLabel(""), widget.NewSeparator())

	return metaDataArea, accordion, err
}

// buildGUIFromSlice builds a fyne.CanvasObject from your slice pointer
func buildGUIFromMetaDataGroupsSlice(
	metaDataGroupsSourcePtr *[]*testCaseModel.MetaDataGroupStruct,
	metaDataGroupInTestCasePtr *testCaseModel.TestCaseMetaDataStruct) fyne.CanvasObject {

	var convertMetaDataToMapMap map[string]map[string]*NewMetaDataInGroupStruct
	convertMetaDataToMapMap = ConvertMetaDataToNewMap(metaDataGroupInTestCasePtr)

	// Create one “card” per MetaDataGroup
	var metaDataGroupCards []fyne.CanvasObject
	metaDataGroupCards = make([]fyne.CanvasObject, 0, len(*metaDataGroupsSourcePtr))

	var metaDataGroupFromTestCase map[string]*NewMetaDataInGroupStruct
	var newMetaDataItemInGroup *NewMetaDataInGroupStruct
	var metaDataGroupFromSourceExistInTestCaseMap bool
	var metaDataGroupItemFromSourceExistInTestCaseMap bool

	// Loop all MetaData-groups
	for _, metaDataGroupPtr := range *metaDataGroupsSourcePtr {

		// Get the MetaDataGroupName from the TestCase
		metaDataGroupFromTestCase, metaDataGroupFromSourceExistInTestCaseMap = convertMetaDataToMapMap[metaDataGroupPtr.MetaDataGroupName]

		// unpack the slice of *MetaDataInGroupStruct
		var metaDataItemsInGroupPtr *[]*testCaseModel.MetaDataInGroupStruct
		metaDataItemsInGroupPtr = metaDataGroupPtr.MetaDataInGroupPtr

		var metaDataItemsAsCanvasObject []fyne.CanvasObject
		metaDataItemsAsCanvasObject = make([]fyne.CanvasObject, 0, len(*metaDataGroupPtr.MetaDataInGroupPtr))

		// Loop all MetaDataItems in the MetaData-group
		for _, metaDataItemPtr := range *metaDataItemsInGroupPtr {

			if metaDataGroupFromSourceExistInTestCaseMap == true {
				newMetaDataItemInGroup, metaDataGroupItemFromSourceExistInTestCaseMap = metaDataGroupFromTestCase[metaDataItemPtr.MetaDataName]
			}

			var metaDataItem testCaseModel.MetaDataInGroupStruct
			metaDataItem = *metaDataItemPtr
			// append '*' to the label if it's mandatory
			label := metaDataItem.MetaDataName
			if metaDataItem.Mandatory {
				label += " *"
			}

			// Create correct widget depending on if the item is SingleSelect or MultiSelect
			switch metaDataItem.SelectType {
			case testCaseModel.MetaDataSelectType_SingleSelect:
				sel := widget.NewSelect(metaDataItem.AvailableMetaDataValues, func(val string) {

					fmt.Printf("Selected %q for %s\n", val, metaDataItem.MetaDataName)

					// store value in TestCase-version of the MetaData
					metaDataItem.SelectedMetaDataValueForSingleSelect = val

					// If the
					if metaDataGroupInTestCasePtr.MetaDataGroupsMapPtr == nil {
						// No 'MetaDataGroupsMap'

						var tempMetaDataInGroupMap map[string]*testCaseModel.MetaDataInGroupStruct
						tempMetaDataInGroupMap = make(map[string]*testCaseModel.MetaDataInGroupStruct)

						// Create MetaData for Group in TestCase
						var tempMetaDataInGroup testCaseModel.MetaDataInGroupStruct
						tempMetaDataInGroup = testCaseModel.MetaDataInGroupStruct{
							MetaDataGroupName:                       metaDataItem.MetaDataGroupName,
							MetaDataName:                            metaDataItem.MetaDataName,
							SelectType:                              metaDataItem.SelectType,
							Mandatory:                               metaDataItem.Mandatory,
							AvailableMetaDataValues:                 metaDataItem.AvailableMetaDataValues,
							SelectedMetaDataValueForSingleSelect:    val,
							SelectedMetaDataValuesForMultiSelect:    nil,
							SelectedMetaDataValuesForMultiSelectMap: nil,
						}

						// Add MetaData for Group to 'MetaDataGroupsMap' in TestCase
						tempMetaDataInGroupMap[metaDataItem.MetaDataGroupName] = &tempMetaDataInGroup

						// Create the 'MetaDataGroup'
						var testMetaDataGroup testCaseModel.MetaDataGroupStruct
						testMetaDataGroup = testCaseModel.MetaDataGroupStruct{
							MetaDataGroupName:     metaDataItem.MetaDataGroupName,
							MetaDataInGroupMapPtr: &tempMetaDataInGroupMap,
						}

						// Create the 'MetaDataGroupsMap' to be able to add the 'MetaDataGroup'
						var tempMetaDataGroupsMap map[string]*testCaseModel.MetaDataGroupStruct
						tempMetaDataGroupsMap = make(map[string]*testCaseModel.MetaDataGroupStruct)

						// Add the 'MetaDataGroup' to the 'MetaDataGroupsMap'
						tempMetaDataGroupsMap[metaDataItem.MetaDataGroupName] = &testMetaDataGroup

						metaDataGroupInTestCasePtr.MetaDataGroupsMapPtr = &tempMetaDataGroupsMap

					} else {
						// 'MetaDataGroupsMap' exists

					}

					fmt.Println(metaDataGroupInTestCasePtr)

				})
				// Extract Selected values from TestCase
				var selectedValue string
				if metaDataGroupFromSourceExistInTestCaseMap == true && metaDataGroupItemFromSourceExistInTestCaseMap == true {
					for _, availableValue := range metaDataItem.AvailableMetaDataValues {
						if value, ok := newMetaDataItemInGroup.SelectedMetaDataValuesForMultiSelectMap[availableValue]; ok {
							selectedValue = value
						}
					}
				}

				sel.PlaceHolder = "Choose..."
				// apply the existing selection if any
				if selectedValue != "" {
					sel.SetSelected(selectedValue)
				}

				// wrap in a 1-cell grid to force width
				w := calcSelectWidth(metaDataItem.AvailableMetaDataValues)
				wrapper := container.New(
					layout.NewGridWrapLayout(fyne.NewSize(w, sel.MinSize().Height)),
					sel,
				)

				metaDataItemsAsCanvasObject = append(metaDataItemsAsCanvasObject,
					container.NewVBox(
						widget.NewLabel(label),
						wrapper,
					),
				)

			case testCaseModel.MetaDataSelectType_MultiSelect:
				chk := widget.NewCheckGroup(metaDataItem.AvailableMetaDataValues, func(vals []string) {

					fmt.Printf("Multi-selected %v for %s\n", vals, metaDataItem.MetaDataName)

				})

				// Extract Selected values from TestCase
				var selectedValues []string
				if metaDataGroupFromSourceExistInTestCaseMap == true && metaDataGroupItemFromSourceExistInTestCaseMap == true {
					for _, availableValue := range metaDataItem.AvailableMetaDataValues {
						if _, ok := metaDataGroupFromTestCase[metaDataItem.MetaDataName].SelectedMetaDataValuesForMultiSelectMap[availableValue]; ok {
							selectedValues = append(selectedValues, availableValue)
						}
					}
				}

				// apply existing selections
				chk.Selected = append([]string(nil), selectedValues...)
				chk.Refresh()

				w := calcCheckGroupWidth(metaDataItem.AvailableMetaDataValues)
				wrapper := container.New(
					layout.NewGridWrapLayout(fyne.NewSize(w, chk.MinSize().Height)),
					chk,
				)

				metaDataItemsAsCanvasObject = append(metaDataItemsAsCanvasObject,
					container.NewVBox(
						widget.NewLabel(label),
						wrapper,
					),
				)

			default:
				// if you have NotSelected or other types, you can skip or handle here
				continue
			}
		}

		// each metadata
		content := container.New(
			layout.NewHBoxLayout(),
			metaDataItemsAsCanvasObject...,
		)

		var card *widget.Card
		card = widget.NewCard(metaDataGroupPtr.MetaDataGroupName, "", content)
		metaDataGroupCards = append(metaDataGroupCards, card)
	}

	// top‐level grid: each card cell is 220×180
	return container.New(
		layout.NewVBoxLayout(),
		metaDataGroupCards...,
	)
}

// calcSelectWidth returns the width needed to show the longest option
func calcSelectWidth(values []string) float32 {
	tmp := widget.NewSelect(values, nil)

	// Loop the values and check which has most characters
	var maxValue float32
	var indexForMaxValue int
	for valueIndex, value := range values {
		if maxValue < float32(len(value)) {
			maxValue = float32(len(value))
			indexForMaxValue = valueIndex
		}
	}

	tmp.PlaceHolder = values[indexForMaxValue] // ensure MinSize measures a non-empty string
	tmp.Refresh()                              // recalc cached size
	return tmp.MinSize().Width
}

// calcCheckGroupWidth returns the width needed to show the widest checkbox label
func calcCheckGroupWidth(values []string) float32 {
	tmp := widget.NewCheckGroup(values, nil)
	tmp.Refresh()
	return tmp.MinSize().Width
}

// NewMetaDataInGroupStruct
// Struct holding the available values, how they are selected and what was selected
type NewMetaDataInGroupStruct struct {
	MetaDataGroupName                       string                           // The name of the MetaData-Group
	MetaDataName                            string                           // The name of the MetaData-post
	SelectType                              testCaseModel.MetaDataSelectType // Is the MetaData-post single- or multi-select
	Mandatory                               bool                             // Is the MetaData-post mandatory or not
	SelectedMetaDataValueForSingleSelect    string                           // The value selected for single select
	SelectedMetaDataValuesForMultiSelectMap map[string]string                // The values selected for multi select
}

// ConvertMetaDataToNewMap transforms the TestCaseMetaDataStruct.MetaDataGroupsSlicePtr
// into a nested map[groupName][metaDataName] => *NewMetaDataInGroupStruct.
func ConvertMetaDataToNewMap(tc *testCaseModel.TestCaseMetaDataStruct) map[string]map[string]*NewMetaDataInGroupStruct {
	result := make(map[string]map[string]*NewMetaDataInGroupStruct)

	if tc == nil {
		return result
	}

	if tc.MetaDataGroupsSlicePtr == nil {
		return result
	}

	for _, grp := range *tc.MetaDataGroupsSlicePtr {
		if grp == nil {
			continue
		}
		inner := make(map[string]*NewMetaDataInGroupStruct)
		if grp.MetaDataInGroupPtr != nil {
			for _, item := range *grp.MetaDataInGroupPtr {
				if item == nil {
					continue
				}
				// build the multi-select map
				selMap := make(map[string]string, len(item.SelectedMetaDataValuesForMultiSelect))
				for _, v := range item.SelectedMetaDataValuesForMultiSelect {
					selMap[v] = v
				}

				newItem := &NewMetaDataInGroupStruct{
					MetaDataGroupName:                       grp.MetaDataGroupName,
					MetaDataName:                            item.MetaDataName,
					SelectType:                              item.SelectType,
					Mandatory:                               item.Mandatory,
					SelectedMetaDataValueForSingleSelect:    item.SelectedMetaDataValueForSingleSelect,
					SelectedMetaDataValuesForMultiSelectMap: selMap,
				}
				inner[item.MetaDataName] = newItem
			}
		}
		result[grp.MetaDataGroupName] = inner
	}
	return result
}
