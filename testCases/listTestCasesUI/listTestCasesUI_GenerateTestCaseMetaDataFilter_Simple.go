package listTestCasesUI

import (
	"FenixTesterGui/testCase/testCaseModel"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/jlambert68/Fast_BitFilter_MetaData/boolbits/boolbits"
	"image/color"
	"log"
)

// generateSimpleTestCaseMetaDataFilterContainer
// Generates the GenerateTestCaseMetaDataFilterContainer containing the simple filter version
func generateSimpleTestCaseMetaDataFilterContainer(
	testCasesModel *testCaseModel.TestCasesModelsStruct) *fyne.Container {

	//var testCaseMetaDataFilterBottomContainer *fyne.Container

	// Initiate the 'simpleMetaDataFilterEntryMap'
	simpleMetaDataFilterEntryMap = make(map[string]*boolbits.Entry)

	// Generate the Top container, having the Domain Filter DropDown
	var testCaseMetaDataDomainFilterTopContainer *fyne.Container
	testCaseMetaDataDomainFilterTopContainer = generateSimpleTestCaseMetaDataDomainFilterTopContainer(testCasesModel)

	// Generate the Bottom container, having buttons for Filter TestCases-list- and clear MetaData-s
	var testCaseMetaDataDomainFilterBottomContainer *fyne.Container
	testCaseMetaDataDomainFilterBottomContainer = generateSimpleTestCaseMetaDataDomainFilterBottomContainer(testCasesModel)

	// Set selected Domain
	simpleTestCaseMetaDataSelectedDomainUuid = ""
	simpleTestCaseMetaDataSelectedDomainDisplayName = ""

	// Generate the main MetaData-filter area
	testCaseMainAreaForMetaDataFilterContainer = generateSimpleTestCaseMetaDataMainFilterContainer(
		simpleTestCaseMetaDataSelectedDomainUuid,
		simpleTestCaseMetaDataSelectedDomainDisplayName,
		testCasesModel)

	// Generate the full MetaDataFilter-container

	testCaseFullMetaDataFilterContainer = container.NewBorder(
		testCaseMetaDataDomainFilterTopContainer,
		testCaseMetaDataDomainFilterBottomContainer,
		nil,
		nil,
		testCaseMainAreaForMetaDataFilterContainer)

	return testCaseFullMetaDataFilterContainer

}

// Generates the top container having the Domain DropDown
func generateSimpleTestCaseMetaDataDomainFilterTopContainer(
	testCasesModel *testCaseModel.TestCasesModelsStruct) *fyne.Container {

	var valueIsValidWarningBox *canvas.Rectangle
	var newOwnerDomainSelect *widget.Select

	// Load Domains that can own the TestCase into options-array
	var options []string
	var uuidForDomainThatCanOwnTheTestCaseMap map[string]string
	uuidForDomainThatCanOwnTheTestCaseMap = make(map[string]string)
	for _, tempDomainsThatCanOwnTheTestCase := range testCasesModel.DomainsThatCanOwnTheTestCaseMap {
		options = append(options, tempDomainsThatCanOwnTheTestCase.DomainNameShownInGui)
		uuidForDomainThatCanOwnTheTestCaseMap[tempDomainsThatCanOwnTheTestCase.DomainNameShownInGui] =
			tempDomainsThatCanOwnTheTestCase.DomainUuid

	}

	// Create Form-layout container to be used for Name
	var testCaseOwnerDomainContainer *fyne.Container
	var testCaseOwnerDomainNameFormContainer *fyne.Container
	testCaseOwnerDomainContainer = container.New(layout.NewVBoxLayout())
	testCaseOwnerDomainNameFormContainer = container.New(layout.NewFormLayout())

	// Add Header to the Forms-container
	var headerLabel *widget.Label
	headerLabel = widget.NewLabel("Domain that 'Own' the TestCase")
	headerLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCaseOwnerDomainNameFormContainer.Add(headerLabel)

	// Generate Warnings-rectangle for valid value, or that value exist
	//var valueIsValidWarningBox *canvas.Rectangle
	var colorToUse color.NRGBA
	colorToUse = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	valueIsValidWarningBox = canvas.NewRectangle(colorToUse)

	// Add the DropDown box with all domains that can own the TestCase
	newOwnerDomainSelect = widget.NewSelect(options,
		func(value string) {
			// This function is called when an option is selected.

			// Set Warning box that value is not selected
			if len(value) == 0 {
				valueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
			} else {
				valueIsValidWarningBox.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
			}

			// Clear 'simpleMetaDataFilterEntryMap' holding all MetaData-Entry
			simpleMetaDataFilterEntryMap = make(map[string]*boolbits.Entry)

			// Set selected Domain
			simpleTestCaseMetaDataSelectedDomainUuid = uuidForDomainThatCanOwnTheTestCaseMap[value]
			simpleTestCaseMetaDataSelectedDomainDisplayName = value

			var newSimpleTestCaseMetaDataMainFilterContainer *fyne.Container
			newSimpleTestCaseMetaDataMainFilterContainer = generateSimpleTestCaseMetaDataMainFilterContainer(
				simpleTestCaseMetaDataSelectedDomainUuid,
				simpleTestCaseMetaDataSelectedDomainDisplayName,
				testCasesModel)

			// Remove old Main MetaData area when user change Domain
			testCaseFullMetaDataFilterContainer.Remove(testCaseMainAreaForMetaDataFilterContainer)

			// Add the new Main MetaData area when user change Domain
			testCaseFullMetaDataFilterContainer.Add(newSimpleTestCaseMetaDataMainFilterContainer)

			// Store 'newSimpleTestCaseMetaDataMainFilterContainer' in 'testCaseMainAreaForMetaDataFilterContainer'-variable
			testCaseMainAreaForMetaDataFilterContainer = newSimpleTestCaseMetaDataMainFilterContainer

			newSimpleTestCaseMetaDataMainFilterContainer.Refresh()

		})

	// Create a custom SelectComboBox, with valueIsValidWarningBox
	var customSelectComboBox *customMandatorySelectComboBox
	customSelectComboBox = newCustomMandatorySelectComboBoxWidget(newOwnerDomainSelect, valueIsValidWarningBox)

	// Add the Entry-widget to the Forms-container
	testCaseOwnerDomainNameFormContainer.Add(customSelectComboBox)

	// Create the VBox-container that will be returned
	testCaseOwnerDomainContainer = container.NewVBox(testCaseOwnerDomainNameFormContainer)

	return testCaseOwnerDomainContainer
}

// Generates the bottom container having the Filter TestCases-list- and clear MetaData.selection
func generateSimpleTestCaseMetaDataDomainFilterBottomContainer(
	testCasesModel *testCaseModel.TestCasesModelsStruct) (simpleTestCaseMetaDataDomainFilterBottomContainer *fyne.Container) {

	var filterTestCasesListButton *widget.Button
	var clearMetaDataSelectionButton *widget.Button

	// Create the Filter TestCases-list button
	filterTestCasesListButton = widget.NewButton("Filter TestCases-list", func() {

		fmt.Println("Filter TestCases-list")

	})

	// Create the Clear MetaData-selections button
	clearMetaDataSelectionButton = widget.NewButton("Clear Selected MetaData values", func() {

		var newSimpleTestCaseMetaDataMainFilterContainer *fyne.Container
		newSimpleTestCaseMetaDataMainFilterContainer = generateSimpleTestCaseMetaDataMainFilterContainer(
			simpleTestCaseMetaDataSelectedDomainUuid,
			simpleTestCaseMetaDataSelectedDomainDisplayName,
			testCasesModel)

		// Remove old Main MetaData area when user change Domain
		testCaseFullMetaDataFilterContainer.Remove(testCaseMainAreaForMetaDataFilterContainer)

		// Add the new Main MetaData area when user change Domain
		testCaseFullMetaDataFilterContainer.Add(newSimpleTestCaseMetaDataMainFilterContainer)

		// Store 'newSimpleTestCaseMetaDataMainFilterContainer' in 'testCaseMainAreaForMetaDataFilterContainer'-variable
		testCaseMainAreaForMetaDataFilterContainer = newSimpleTestCaseMetaDataMainFilterContainer

		newSimpleTestCaseMetaDataMainFilterContainer.Refresh()
	})

	// Container having the Buttons
	simpleTestCaseMetaDataDomainFilterBottomContainer = container.NewHBox(
		filterTestCasesListButton,
		clearMetaDataSelectionButton)

	return simpleTestCaseMetaDataDomainFilterBottomContainer

}

// Generates the main container having all the MetaData-filters
func generateSimpleTestCaseMetaDataMainFilterContainer(
	domainUuidToGetMetaDataFor string,
	domainNameToGetMetaDataFor string,
	testCasesModel *testCaseModel.TestCasesModelsStruct) (metaDataFilterArea *fyne.Container) {
	var existInMap bool

	//
	if len(domainUuidToGetMetaDataFor) > 0 {

		// Check if there are any MetaData for the Domain
		var testCaseMetaDataForDomainsForMapPtr *testCaseModel.TestCaseMetaDataForDomainsForMapStruct
		testCaseMetaDataForDomainsForMapPtr, existInMap = testCasesModel.TestCaseMetaDataForDomains.TestCaseMetaDataForDomainsMap[domainUuidToGetMetaDataFor]
		if existInMap == false {
			metaDataFilterArea = container.NewVBox(widget.NewLabel(
				fmt.Sprintf("MetaData is not available for Domain having Uuid '%s'",
					domainNameToGetMetaDataFor)))

			return metaDataFilterArea
		}

		// Get the MetaDataGroups depending on Domain
		var metaDataGroupsPtr *map[string]*testCaseModel.MetaDataGroupStruct
		var tempMetaDataGroupsOrder []string
		metaDataGroupsPtr, tempMetaDataGroupsOrder = testCaseModel.ConvertTestCaseMetaData(testCaseMetaDataForDomainsForMapPtr.TestCaseMetaDataForDomainPtr)

		// Generate TestCaseMeta-UI-object
		var metaDataGroupsContainer *fyne.Container
		var metaDataGroupsScroll *container.Scroll
		metaDataGroupsContainer = buildGUIFromMetaDataGroupsMap(
			domainUuidToGetMetaDataFor,
			tempMetaDataGroupsOrder,
			metaDataGroupsPtr,
			testCasesModel)

		metaDataGroupsScroll = container.NewScroll(metaDataGroupsContainer)

		metaDataFilterArea = container.NewBorder(nil, nil, nil, nil, metaDataGroupsScroll)

	} else {

		metaDataFilterArea = container.New(layout.NewGridLayout(1), widget.NewLabel("MetaData is available when 'Owner Domain' is selected"))

	}

	return metaDataFilterArea
}

// buildGUIFromSlice builds a fyne.CanvasObject from your slice pointer
func buildGUIFromMetaDataGroupsMap(
	domainUUid string,
	metaDataGroupsOrder []string,
	metaDataGroupsSourceMapPtr *map[string]*testCaseModel.MetaDataGroupStruct,
	testCasesModels *testCaseModel.TestCasesModelsStruct) *fyne.Container {

	var err error

	// Get the 'metaDataGroupsSourceMap'
	var metaDataGroupsSourceMap map[string]*testCaseModel.MetaDataGroupStruct
	metaDataGroupsSourceMap = *metaDataGroupsSourceMapPtr

	if len(*metaDataGroupsSourceMapPtr) != len(metaDataGroupsOrder) {
		log.Fatalln("ERROR: The number of MetaDataGroups in the 'metaDataGroupsSourceMap' doesn't match the number of MetaDataGroups in the 'metaDataGroupsOrder'")
	}

	//var convertMetaDataToMapMap map[string]map[string]*NewMetaDataInGroupStruct
	//convertMetaDataToMapMap = ConvertMetaDataToNewMap(metaDataGroupInTestCasePtr)

	// Create one “card” per MetaDataGroup
	var metaDataGroupCards []fyne.CanvasObject
	metaDataGroupCards = make([]fyne.CanvasObject, 0, len(*metaDataGroupsSourceMapPtr))

	var metaDataGroupFromTestCase map[string]*NewMetaDataInGroupStruct
	var newMetaDataItemInGroup *NewMetaDataInGroupStruct
	var metaDataGroupFromSourceExistInTestCaseMap bool
	var metaDataGroupItemFromSourceExistInTestCaseMap bool

	// Loop all MetaData-groups
	for _, metaDataGroupName := range metaDataGroupsOrder {

		var metaDataGroupPtr *testCaseModel.MetaDataGroupStruct
		metaDataGroupPtr = metaDataGroupsSourceMap[metaDataGroupName]

		// Get the MetaDataGroupName from the TestCase
		//metaDataGroupFromTestCase, metaDataGroupFromSourceExistInTestCaseMap = convertMetaDataToMapMap[metaDataGroupPtr.MetaDataGroupName]

		// unpack the slice of *MetaDataInGroupStruct
		var metaDataItemsInGroupPtr *map[string]*testCaseModel.MetaDataInGroupStruct
		metaDataItemsInGroupPtr = metaDataGroupPtr.MetaDataInGroupMapPtr

		// Get the metaDataItemsInGroupMap
		var metaDataItemsInGroupMap map[string]*testCaseModel.MetaDataInGroupStruct
		metaDataItemsInGroupMap = *metaDataItemsInGroupPtr

		var metaDataItemsAsCanvasObject []fyne.CanvasObject
		metaDataItemsAsCanvasObject = make([]fyne.CanvasObject, 0, len(*metaDataGroupPtr.MetaDataInGroupMapPtr))

		// Loop all MetaDataItems in the MetaData-group
		for _, metaDataItemName := range metaDataGroupPtr.MetaDataInGroupOrder {

			//metaDataGroupPtr.MetaDataInGroupOrder är tom

			var metaDataItemPtr *testCaseModel.MetaDataInGroupStruct
			metaDataItemPtr = metaDataItemsInGroupMap[metaDataItemName]

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

				var valueIsValidWarningBox *canvas.Rectangle

				// Generate Warnings-rectangle for valid value, or that value exist
				//var valueIsValidWarningBox *canvas.Rectangle
				var colorToUse color.NRGBA
				colorToUse = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
				valueIsValidWarningBox = canvas.NewRectangle(colorToUse)

				sel := widget.NewSelect(metaDataItem.AvailableMetaDataValues, func(val string) {

					// Set Warning box that value is not selected
					if len(val) == 0 && metaDataItem.Mandatory == true {
						valueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
					} else {
						valueIsValidWarningBox.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
					}

					// Remove all MetDataGroupItem with values from map with  'MetaData-Entries'
					var entryKey string
					for _, value := range metaDataItem.AvailableMetaDataValues {

						entryKey = fmt.Sprintf("%s.%s.%s.%s",
							domainUUid,
							metaDataItem.MetaDataGroupName,
							metaDataItem.MetaDataName,
							value)

						delete(simpleMetaDataFilterEntryMap, entryKey)
					}

					// If there is a value then add the MetaDataEntry for it
					if len(val) > 0 {
						// Add New pointer to selected value into 'simpleMetaDataFilterEntryMap'
						var metaDataEntry *boolbits.Entry
						var (
							domainBitSet *boolbits.BitSet
							groupBitSet  *boolbits.BitSet
							nameBitSet   *boolbits.BitSet
							valueBitSet  *boolbits.BitSet
						)

						// Get BitSet-parts for the Selected MetaData
						domainBitSet = testCasesModels.TestCaseMetaDataForDomains.UniqueMetaDataBitSets.DomainsBitSetMap[domainUUid]
						groupBitSet = testCasesModels.TestCaseMetaDataForDomains.UniqueMetaDataBitSets.MetaDataGroupsBitSetMap[metaDataItem.MetaDataGroupName]
						nameBitSet = testCasesModels.TestCaseMetaDataForDomains.UniqueMetaDataBitSets.MetaDataGroupItemsBitSetMap[metaDataItem.MetaDataName]
						valueBitSet = testCasesModels.TestCaseMetaDataForDomains.UniqueMetaDataBitSets.MetaDataGroupItemValuesBitSetMap[val]

						metaDataEntry, err = boolbits.NewEntry(domainBitSet, groupBitSet, nameBitSet, valueBitSet)

						if err != nil {
							errorID := "cf5a4b44-6559-4827-9ec8-377aefab0017"
							errorMessage := fmt.Sprintf("could not create MetaDataEntry for Domain '%s', MetaDataGroup '%s', MetaDataGroupItem '%s' and value '%s' [ErrorID=%s, err='%s']",
								domainUUid,
								metaDataItem.MetaDataGroupName,
								metaDataItem.MetaDataName,
								val,
								errorID,
								err.Error())

							log.Fatalln(errorMessage)
						}

						// Add MetaDataEntry to Entry-map
						entryKey = fmt.Sprintf("%s.%s.%s.%s",
							domainUUid,
							metaDataItem.MetaDataGroupName,
							metaDataItem.MetaDataName,
							val)

						simpleMetaDataFilterEntryMap[entryKey] = metaDataEntry

					}

				})
				// Extract Selected values from TestCase
				var selectedValue string
				if metaDataGroupFromSourceExistInTestCaseMap == true && metaDataGroupItemFromSourceExistInTestCaseMap == true {
					for _, availableValue := range metaDataItem.AvailableMetaDataValues {
						if newMetaDataItemInGroup.SelectedMetaDataValueForSingleSelect == availableValue {
							selectedValue = availableValue
							break
						}
					}
				}

				sel.PlaceHolder = "Choose..."
				// apply the existing selection if any
				if selectedValue != "" {
					sel.SetSelected(selectedValue)
				}

				// Resize the DropDown

				// Create a custom SelectComboBox, with valueIsValidWarningBox
				var customSelectComboBox *customMandatorySelectComboBox
				customSelectComboBox = newCustomMandatorySelectComboBoxWidget(sel, valueIsValidWarningBox)

				// wrap in a 1-cell grid to force width
				w := calcSelectWidth(metaDataItem.AvailableMetaDataValues)
				wrapper := container.New(
					layout.NewGridWrapLayout(fyne.NewSize(w, sel.MinSize().Height)),
					customSelectComboBox,
				)

				// Set Warning box that value is not selected
				if len(customSelectComboBox.selectComboBox.Selected) == 0 && metaDataItem.Mandatory == true {
					valueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
				} else {
					valueIsValidWarningBox.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
				}

				metaDataItemsAsCanvasObject = append(metaDataItemsAsCanvasObject,
					container.NewVBox(
						widget.NewLabel("   "+label),
						wrapper,
					),
				)

			case testCaseModel.MetaDataSelectType_MultiSelect:

				var valueIsValidWarningBox *canvas.Rectangle

				// Generate Warnings-rectangle for valid value, or that value exist
				//var valueIsValidWarningBox *canvas.Rectangle
				var colorToUse color.NRGBA
				colorToUse = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
				valueIsValidWarningBox = canvas.NewRectangle(colorToUse)

				var chk *widget.CheckGroup
				chk = widget.NewCheckGroup(metaDataItem.AvailableMetaDataValues, func(vals []string) {

					// Set Warning box that value is not selected
					if len(vals) == 0 && metaDataItem.Mandatory == true {
						valueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
					} else {
						valueIsValidWarningBox.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
					}

					// Remove all MetDataGroupItem with values from map with  'MetaData-Entries'
					var entryKey string
					for _, value := range metaDataItem.AvailableMetaDataValues {

						entryKey = fmt.Sprintf("%s.%s.%s.%s",
							domainUUid,
							metaDataItem.MetaDataGroupName,
							metaDataItem.MetaDataName,
							value)

						delete(simpleMetaDataFilterEntryMap, entryKey)
					}

					// If there is any value then add the MetaDataEntry for it all values
					if len(vals) > 0 {
						// Add New pointer to selected value into 'simpleMetaDataFilterEntryMap'
						var metaDataEntry *boolbits.Entry
						var (
							domainBitSet *boolbits.BitSet
							groupBitSet  *boolbits.BitSet
							nameBitSet   *boolbits.BitSet
							valueBitSet  *boolbits.BitSet
						)

						// Loop all selected values and create its MetaDataEntry
						for _, value := range vals {
							// Get BitSet-parts for the Selected MetaData
							domainBitSet = testCasesModels.TestCaseMetaDataForDomains.UniqueMetaDataBitSets.DomainsBitSetMap[domainUUid]
							groupBitSet = testCasesModels.TestCaseMetaDataForDomains.UniqueMetaDataBitSets.MetaDataGroupsBitSetMap[metaDataItem.MetaDataGroupName]
							nameBitSet = testCasesModels.TestCaseMetaDataForDomains.UniqueMetaDataBitSets.MetaDataGroupItemsBitSetMap[metaDataItem.MetaDataName]
							valueBitSet = testCasesModels.TestCaseMetaDataForDomains.UniqueMetaDataBitSets.MetaDataGroupItemValuesBitSetMap[value]

							metaDataEntry, err = boolbits.NewEntry(domainBitSet, groupBitSet, nameBitSet, valueBitSet)

							if err != nil {
								errorID := "d12a97d9-a3a0-4ae2-9385-62d81806afb8"
								errorMessage := fmt.Sprintf("could not create MetaDataEntry for Domain '%s', MetaDataGroup '%s', MetaDataGroupItem '%s' and value '%s' [ErrorID=%s, err='%s']",
									domainUUid,
									metaDataItem.MetaDataGroupName,
									metaDataItem.MetaDataName,
									value,
									errorID,
									err.Error())

								log.Fatalln(errorMessage)
							}

							// Add MetaDataEntry to Entry-map
							entryKey = fmt.Sprintf("%s.%s.%s.%s",
								domainUUid,
								metaDataItem.MetaDataGroupName,
								metaDataItem.MetaDataName,
								value)

							simpleMetaDataFilterEntryMap[entryKey] = metaDataEntry

						}

					}

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

				// Create a custom SelectComboBox, with valueIsValidWarningBox
				var customCheckBoxGroup *customMandatoryCheckBoxGroup
				customCheckBoxGroup = newCustomMandatoryCheckBoxGroupWidget(chk, valueIsValidWarningBox)

				w := calcCheckGroupWidth(metaDataItem.AvailableMetaDataValues)
				wrapper := container.New(
					layout.NewGridWrapLayout(fyne.NewSize(w, chk.MinSize().Height)),
					customCheckBoxGroup,
				)

				// Set Warning box that value is not selected
				if len(customCheckBoxGroup.checkBoxGroup.Selected) == 0 && metaDataItem.Mandatory == true {
					valueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
				} else {
					valueIsValidWarningBox.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
				}

				metaDataItemsAsCanvasObject = append(metaDataItemsAsCanvasObject,
					container.NewVBox(
						widget.NewLabel("   "+label),
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

	if tc.MetaDataGroupsMapPtr == nil {
		return result
	}

	for _, grp := range *tc.MetaDataGroupsMapPtr {
		if grp == nil {
			continue
		}
		inner := make(map[string]*NewMetaDataInGroupStruct)
		if grp.MetaDataInGroupMapPtr != nil {
			for _, item := range *grp.MetaDataInGroupMapPtr {
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
