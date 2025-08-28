package listTestSuitesUI

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

// generateSimpleTestSuiteMetaDataFilterContainer
// Generates the GenerateTestSuiteMetaDataFilterContainer containing the simple filter version
func (listTestSuiteUIObject *ListTestSuiteUIStruct) generateSimpleTestSuiteMetaDataFilterContainer(
	testCasesModel *testCaseModel.TestCasesModelsStruct) *fyne.Container {

	// Generate the filter function used when user clicks on filter-button or when auto-filter is turned on
	listTestSuiteUIObject.filterOnMetaDataFunction = func(resultEntry *boolbits.Entry, testCasesModel *testCaseModel.TestCasesModelsStruct) {

		// Load a filtered list into the TestSuites-list-table
		listTestSuiteUIObject.loadTestSuiteListTableTable(resultEntry)
		listTestSuiteUIObject.calculateAndSetCorrectColumnWidths()
		listTestSuiteUIObject.updateTestSuitesListTable(testCasesModel)
	}

	// Generate the function used for calculating the current MetaData-filter
	listTestSuiteUIObject.calculateMetaDataFilterFunction = func() {

		var err error

		// Generate Initial Entry for the boolean arithmetic
		var resultEntry *boolbits.Entry
		resultEntry, err = boolbits.NewAllZerosEntry(64)

		if err != nil {
			errorID := "3d244190-17b5-48cc-80c2-8b46acf31e76"
			errorMessage := fmt.Sprintf("could not create initial Entry for boolean arithmetic [ErrorID=%s, err='%s']",
				errorID,
				err.Error())

			log.Fatalln(errorMessage)
		}

		// Check if there are any MetaData set
		if listTestSuiteUIObject.simpleMetaDataFilterEntryMap == nil || len(listTestSuiteUIObject.simpleMetaDataFilterEntryMap) == 0 {
			// No MetaData-filter is set by the user

			// Create Filter-Entry based on Domain
			var booleanANDResultsEntry *boolbits.Entry
			booleanANDResultsEntry, err = boolbits.NewAllOnesEntry(64)

			if err != nil {
				errorID := "5a96164d-a3d0-4019-9cef-6e474f63be6f"
				errorMessage := fmt.Sprintf("could not create initial Entry for boolean arithmetic [ErrorID=%s, err='%s']",
					errorID,
					err.Error())

				log.Fatalln(errorMessage)
			}

			// Generate BitSet for Domain, Group, Item and Value
			var domainBitSet *boolbits.BitSet
			var metaDataGroupBitSet *boolbits.BitSet
			var metaDataItemBitSet *boolbits.BitSet
			var metaDataValueBitSet *boolbits.BitSet
			var domainBitSetExistInMap bool

			// Get BitSets
			domainBitSet, domainBitSetExistInMap = testCasesModel.TestCaseMetaDataForDomains.UniqueMetaDataBitSets.
				DomainsBitSetMap[listTestSuiteUIObject.newMandatoryOwnerDomainSelect.dataValueRepresentingVisualizedData]

			if domainBitSetExistInMap == false {

				domainBitSet, err = boolbits.NewBitSet(64)

				if err != nil {
					errorID := "7ad69c47-b508-48e0-84df-f01ad30c66fd"
					errorMessage := fmt.Sprintf("could not create NewBitSet[ErrorID=%s, err='%s']",
						errorID,
						err.Error())

					log.Fatalln(errorMessage)
				}
			}

			metaDataGroupBitSet, err = boolbits.NewBitSet(64)
			if err != nil {
				errorID := "359a59d1-bd56-48c7-917c-84a254235598"
				errorMessage := fmt.Sprintf("could not create NewBitSet[ErrorID=%s, err='%s']",
					errorID,
					err.Error())

				log.Fatalln(errorMessage)
			}

			metaDataItemBitSet, err = boolbits.NewBitSet(64)
			if err != nil {
				errorID := "c7174a04-0a4b-4f62-85f1-7b1b3b8fdd84"
				errorMessage := fmt.Sprintf("could not create NewBitSet[ErrorID=%s, err='%s']",
					errorID,
					err.Error())

				log.Fatalln(errorMessage)
			}

			metaDataValueBitSet, err = boolbits.NewBitSet(64)
			if err != nil {
				errorID := "540f9ab4-9dc0-4028-875e-a06fb155e8c1"
				errorMessage := fmt.Sprintf("could not create NewBitSet[ErrorID=%s, err='%s']",
					errorID,
					err.Error())

				log.Fatalln(errorMessage)
			}

			var metaDataOnlyDomainEntry *boolbits.Entry
			metaDataOnlyDomainEntry, err = boolbits.NewEntry(domainBitSet, metaDataGroupBitSet, metaDataItemBitSet, metaDataValueBitSet)
			if err != nil {
				errorID := "9015af8a-94e5-4341-824c-a636567bb13a"
				errorMessage := fmt.Sprintf("could not create Domain-Entry [ErrorID=%s, err='%s']",
					errorID,
					err.Error())

				log.Fatalln(errorMessage)
			}

			resultEntry, err = booleanANDResultsEntry.And(metaDataOnlyDomainEntry)

			if err != nil {
				errorID := "5706d1df-462f-47bf-8451-1fdfb0e1fbe3"
				errorMessage := fmt.Sprintf("could not do boolean arithmetic, AND [ErrorID=%s, err='%s']",
					errorID,
					err.Error())

				log.Println(errorMessage)
			}

		} else {

			// MetaData-filter is selected by the user

			// Loop all Simple MetaDataEntry and make boolean 'OR' between all of them
			for _, simpleMetaDataEntry := range listTestSuiteUIObject.simpleMetaDataFilterEntryMap {

				// If multiple values per MEtaDataItem exist then process them with boolean OR
				var booleanOrResultsEntry *boolbits.Entry
				for valueIndex, tempValueEntryListToBeProcessedWithBooleanOr := range simpleMetaDataEntry.valueEntryListToBeProcessedWithBooleanOrSlice {

					if valueIndex == 0 {
						booleanOrResultsEntry = tempValueEntryListToBeProcessedWithBooleanOr

					} else {
						booleanOrResultsEntry, err = booleanOrResultsEntry.Or(tempValueEntryListToBeProcessedWithBooleanOr)

						if err != nil {
							errorID := "8657cc14-cdef-4ea9-9b5b-9143de0f5df0"
							errorMessage := fmt.Sprintf("could not do boolean arithmetic, OR [ErrorID=%s, err='%s']",
								errorID,
								err.Error())

							log.Println(errorMessage)
						}

					}
				}

				resultEntry, err = resultEntry.Or(booleanOrResultsEntry)

				if err != nil {
					errorID := "cee28686-0995-4492-94c3-01b884731e93"
					errorMessage := fmt.Sprintf("could not do boolean arithmetic, OR [ErrorID=%s, err='%s']",
						errorID,
						err.Error())

					log.Fatalln(errorMessage)
				}

			}
		}

		// Load a filtered list into the TestCases-list-table
		listTestSuiteUIObject.filterOnMetaDataFunction(resultEntry, testCasesModel)
	}

	// Initiate the 'simpleMetaDataFilterEntryMap'
	listTestSuiteUIObject.simpleMetaDataFilterEntryMap = make(map[string]simpleMetaDataFilterEntryMapStruct)

	// Generate the Top container, having the Domain Filter DropDown
	var testSuiteMetaDataDomainFilterTopContainer *fyne.Container
	testSuiteMetaDataDomainFilterTopContainer = listTestSuiteUIObject.generateSimpleTestSuiteMetaDataDomainFilterTopContainer(testCasesModel)

	// Generate the Bottom container, having buttons for Filter TestCases-list- and clear MetaData-s
	var testSuiteMetaDataDomainFilterBottomContainer *fyne.Container
	testSuiteMetaDataDomainFilterBottomContainer = listTestSuiteUIObject.generateSimpleTestSuiteMetaDataDomainFilterBottomContainer(testCasesModel)

	// Set selected Domain
	listTestSuiteUIObject.simpleTestSuiteMetaDataSelectedDomainUuid = ""
	listTestSuiteUIObject.simpleTestSuiteMetaDataSelectedDomainDisplayName = ""

	// Generate the main MetaData-filter area
	listTestSuiteUIObject.testSuiteMainAreaForMetaDataFilterContainer = listTestSuiteUIObject.generateSimpleTestSuiteMetaDataMainFilterContainer(
		listTestSuiteUIObject.simpleTestSuiteMetaDataSelectedDomainUuid,
		listTestSuiteUIObject.simpleTestSuiteMetaDataSelectedDomainDisplayName,
		testCasesModel)

	// Generate the full MetaDataFilter-container

	listTestSuiteUIObject.testSuiteFullMetaDataFilterContainer = container.NewBorder(
		testSuiteMetaDataDomainFilterTopContainer,
		testSuiteMetaDataDomainFilterBottomContainer,
		nil,
		nil,
		listTestSuiteUIObject.testSuiteMainAreaForMetaDataFilterContainer)

	return listTestSuiteUIObject.testSuiteFullMetaDataFilterContainer

}

// Generates the top container having the Domain DropDown
func (listTestSuiteUIObject *ListTestSuiteUIStruct) generateSimpleTestSuiteMetaDataDomainFilterTopContainer(
	testCasesModel *testCaseModel.TestCasesModelsStruct) *fyne.Container {

	var valueIsValidWarningBox *canvas.Rectangle
	var newOwnerDomainSelect *widget.Select

	// Load Domains that can own the TestSuite into options-array
	var options []string
	var uuidForDomainThatCanOwnTheTestSuiteMap map[string]string
	uuidForDomainThatCanOwnTheTestSuiteMap = make(map[string]string)
	for _, tempDomainsThatCanOwnTheTestSuite := range testCasesModel.DomainsThatCanOwnTheTestCaseOrTestSuiteMap {
		options = append(options, tempDomainsThatCanOwnTheTestSuite.DomainNameShownInGui)
		uuidForDomainThatCanOwnTheTestSuiteMap[tempDomainsThatCanOwnTheTestSuite.DomainNameShownInGui] =
			tempDomainsThatCanOwnTheTestSuite.DomainUuid

	}

	// Create Form-layout container to be used for Name
	var testSuiteOwnerDomainContainer *fyne.Container
	var testSuiteOwnerDomainNameFormContainer *fyne.Container
	testSuiteOwnerDomainContainer = container.New(layout.NewVBoxLayout())
	testSuiteOwnerDomainNameFormContainer = container.New(layout.NewFormLayout())

	// Add Header to the Forms-container
	var headerLabel *widget.Label
	headerLabel = widget.NewLabel("Domain that 'Own' the TestCase")
	headerLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuiteOwnerDomainNameFormContainer.Add(headerLabel)

	// Generate Warnings-rectangle for valid value, or that value exist
	//var valueIsValidWarningBox *canvas.Rectangle
	var colorToUse color.NRGBA
	colorToUse = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	valueIsValidWarningBox = canvas.NewRectangle(colorToUse)

	// Add the DropDown box with all domains that can own the TestCase
	newOwnerDomainSelect = widget.NewSelect(options,
		func(value string) {
			// This function is called when an option is selected.

			// Set selected Domain
			listTestSuiteUIObject.simpleTestSuiteMetaDataSelectedDomainUuid = uuidForDomainThatCanOwnTheTestSuiteMap[value]
			listTestSuiteUIObject.simpleTestSuiteMetaDataSelectedDomainDisplayName = value

			// Set Warning box that value is not selected
			if len(value) == 0 {
				valueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}

				// Store data value representing the visual value
				listTestSuiteUIObject.newMandatoryOwnerDomainSelect.dataValueRepresentingVisualizedData = ""

			} else {
				valueIsValidWarningBox.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
				// Store data value representing the visual value
				listTestSuiteUIObject.newMandatoryOwnerDomainSelect.dataValueRepresentingVisualizedData = listTestSuiteUIObject.simpleTestSuiteMetaDataSelectedDomainUuid

			}

			// Clear 'simpleMetaDataFilterEntryMap' holding all MetaData-Entry
			listTestSuiteUIObject.simpleMetaDataFilterEntryMap = make(map[string]simpleMetaDataFilterEntryMapStruct)

			var newSimpleTestSuiteMetaDataMainFilterContainer *fyne.Container
			newSimpleTestSuiteMetaDataMainFilterContainer = listTestSuiteUIObject.generateSimpleTestSuiteMetaDataMainFilterContainer(
				listTestSuiteUIObject.simpleTestSuiteMetaDataSelectedDomainUuid,
				listTestSuiteUIObject.simpleTestSuiteMetaDataSelectedDomainDisplayName,
				testCasesModel)

			// Remove old Main MetaData area when user change Domain
			listTestSuiteUIObject.testSuiteFullMetaDataFilterContainer.Remove(listTestSuiteUIObject.testSuiteMainAreaForMetaDataFilterContainer)

			// Add the new Main MetaData area when user change Domain
			listTestSuiteUIObject.testSuiteFullMetaDataFilterContainer.Add(newSimpleTestSuiteMetaDataMainFilterContainer)

			// Store 'newSimpleTestSuiteMetaDataMainFilterContainer' in 'testSuiteMainAreaForMetaDataFilterContainer'-variable
			listTestSuiteUIObject.testSuiteMainAreaForMetaDataFilterContainer = newSimpleTestSuiteMetaDataMainFilterContainer

			newSimpleTestSuiteMetaDataMainFilterContainer.Refresh()

			// Check if auto-filter is enabled. If so then calculate the new TestSuite-liset
			if listTestSuiteUIObject.useAutoFilter == true {
				listTestSuiteUIObject.calculateMetaDataFilterFunction()
			}

		})

	// Create a custom SelectComboBox, with valueIsValidWarningBox
	//var customSelectComboBox *customMandatorySelectComboBox
	listTestSuiteUIObject.newMandatoryOwnerDomainSelect = newCustomMandatorySelectComboBoxWidget(newOwnerDomainSelect, valueIsValidWarningBox)

	// Add the Entry-widget to the Forms-container
	testSuiteOwnerDomainNameFormContainer.Add(listTestSuiteUIObject.newMandatoryOwnerDomainSelect)

	// Create the VBox-container that will be returned
	testSuiteOwnerDomainContainer = container.NewVBox(testSuiteOwnerDomainNameFormContainer)

	return testSuiteOwnerDomainContainer
}

// Generates the bottom container having the Filter TestSuites-list- and clear MetaData.selection
func (listTestSuiteUIObject *ListTestSuiteUIStruct) generateSimpleTestSuiteMetaDataDomainFilterBottomContainer(
	testCasesModel *testCaseModel.TestCasesModelsStruct) (simpleTestSuiteMetaDataDomainFilterBottomContainer *fyne.Container) {

	var filterTestSuitesListButton *widget.Button
	var clearMetaDataSelectionButton *widget.Button

	// Create the Filter TestSuites-list button
	filterTestSuitesListButton = widget.NewButton("Filter TestSuites-list", func() {

		// Trigger the function to calculate the MetaDataFilter
		listTestSuiteUIObject.calculateMetaDataFilterFunction()

	})

	// Create the Clear MetaData-selections button
	clearMetaDataSelectionButton = widget.NewButton("Clear Selected MetaData values", func() {

		var newSimpleTestSuiteMetaDataMainFilterContainer *fyne.Container
		newSimpleTestSuiteMetaDataMainFilterContainer = listTestSuiteUIObject.generateSimpleTestSuiteMetaDataMainFilterContainer(
			listTestSuiteUIObject.simpleTestSuiteMetaDataSelectedDomainUuid,
			listTestSuiteUIObject.simpleTestSuiteMetaDataSelectedDomainDisplayName,
			testCasesModel)

		// Remove old Main MetaData area when user change Domain
		listTestSuiteUIObject.testSuiteFullMetaDataFilterContainer.Remove(listTestSuiteUIObject.testSuiteMainAreaForMetaDataFilterContainer)

		// Add the new Main MetaData area when user change Domain
		listTestSuiteUIObject.testSuiteFullMetaDataFilterContainer.Add(newSimpleTestSuiteMetaDataMainFilterContainer)

		// Store 'newSimpleTestSuiteMetaDataMainFilterContainer' in 'testSuiteMainAreaForMetaDataFilterContainer'-variable
		listTestSuiteUIObject.testSuiteMainAreaForMetaDataFilterContainer = newSimpleTestSuiteMetaDataMainFilterContainer

		newSimpleTestSuiteMetaDataMainFilterContainer.Refresh()

		// clear the Domain-dropdown
		listTestSuiteUIObject.newMandatoryOwnerDomainSelect.selectComboBox.ClearSelected()
		listTestSuiteUIObject.newMandatoryOwnerDomainSelect.Refresh()

		// Clear selected MetaData-map
		listTestSuiteUIObject.simpleMetaDataFilterEntryMap = make(map[string]simpleMetaDataFilterEntryMapStruct)
	})

	// Create radio button for Auto-filter where each change in filter setting automatically filters the list
	var autoFilterRadioGroup *widget.RadioGroup
	autoFilterRadioGroup = widget.NewRadioGroup([]string{autoFilterRadioGroupOn, autoFilterRadioGroupOff}, func(selectedValue string) {

		if selectedValue == autoFilterRadioGroupOn {

			listTestSuiteUIObject.useAutoFilter = true
		} else {

			listTestSuiteUIObject.useAutoFilter = false
		}

	})
	autoFilterRadioGroup.SetSelected(autoFilterRadioGroupOn)

	// Container having the Buttons
	simpleTestSuiteMetaDataDomainFilterBottomContainer = container.NewHBox(
		filterTestSuitesListButton,
		clearMetaDataSelectionButton,
		autoFilterRadioGroup)

	return simpleTestSuiteMetaDataDomainFilterBottomContainer

}

// Generates the main container having all the MetaData-filters
func (listTestSuiteUIObject *ListTestSuiteUIStruct) generateSimpleTestSuiteMetaDataMainFilterContainer(
	domainUuidToGetMetaDataFor string,
	domainNameToGetMetaDataFor string,
	testCasesModel *testCaseModel.TestCasesModelsStruct) (metaDataFilterArea *fyne.Container) {
	var existInMap bool

	//
	if len(domainUuidToGetMetaDataFor) > 0 {

		// Check if there are any MetaData for the Domain
		var testSuiteMetaDataForDomainsForMapPtr *testCaseModel.TestCaseMetaDataForDomainsForMapStruct
		testSuiteMetaDataForDomainsForMapPtr, existInMap = testCasesModel.TestCaseMetaDataForDomains.TestCaseMetaDataForDomainsMap[domainUuidToGetMetaDataFor]
		if existInMap == false {
			metaDataFilterArea = container.NewVBox(widget.NewLabel(
				fmt.Sprintf("MetaData is not available for Domain having Uuid '%s'",
					domainNameToGetMetaDataFor)))

			return metaDataFilterArea
		}

		// Get the MetaDataGroups depending on Domain
		var metaDataGroupsPtr *map[string]*testCaseModel.MetaDataGroupStruct
		var tempMetaDataGroupsOrder []string
		metaDataGroupsPtr, tempMetaDataGroupsOrder = testCaseModel.ConvertTestCaseMetaData(testSuiteMetaDataForDomainsForMapPtr.TestCaseMetaDataForDomainPtr)

		// Generate TestCaseMeta-UI-object
		var metaDataGroupsContainer *fyne.Container
		var metaDataGroupsScroll *container.Scroll
		metaDataGroupsContainer = listTestSuiteUIObject.buildGUIFromMetaDataGroupsMap(
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
func (listTestSuiteUIObject *ListTestSuiteUIStruct) buildGUIFromMetaDataGroupsMap(
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
	//convertMetaDataToMapMap = convertMetaDataToNewMap(metaDataGroupInTestCasePtr)

	// Create one “card” per MetaDataGroup
	var metaDataGroupCards []fyne.CanvasObject
	metaDataGroupCards = make([]fyne.CanvasObject, 0, len(*metaDataGroupsSourceMapPtr))

	var metaDataGroupFromTestSuite map[string]*NewMetaDataInGroupStruct
	var newMetaDataItemInGroup *NewMetaDataInGroupStruct
	var metaDataGroupFromSourceExistInTestSuiteMap bool
	var metaDataGroupItemFromSourceExistInTestSuiteMap bool

	// Loop all MetaData-groups
	for _, metaDataGroupName := range metaDataGroupsOrder {

		var metaDataGroupPtr *testCaseModel.MetaDataGroupStruct
		metaDataGroupPtr = metaDataGroupsSourceMap[metaDataGroupName]

		// Get the MetaDataGroupName from the TestSuite
		//metaDataGroupFromTestSuite, metaDataGroupFromSourceExistInTestSuiteMap = convertMetaDataToMapMap[metaDataGroupPtr.MetaDataGroupName]

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

			if metaDataGroupFromSourceExistInTestSuiteMap == true {
				newMetaDataItemInGroup, metaDataGroupItemFromSourceExistInTestSuiteMap = metaDataGroupFromTestSuite[metaDataItemPtr.MetaDataName]
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

					entryKey = fmt.Sprintf("%s.%s.%s",
						domainUUid,
						metaDataItem.MetaDataGroupName,
						metaDataItem.MetaDataName)

					delete(listTestSuiteUIObject.simpleMetaDataFilterEntryMap, entryKey)

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

						var valueEntryListToBeProcessedWithBooleanOrSlice []*boolbits.Entry

						// Get BitSet-parts for the Selected MetaData
						domainBitSet = testCasesModels.TestCaseMetaDataForDomains.UniqueMetaDataBitSets.DomainsBitSetMap[domainUUid]
						groupBitSet = testCasesModels.TestCaseMetaDataForDomains.UniqueMetaDataBitSets.MetaDataGroupsBitSetMap[metaDataItem.MetaDataGroupName]
						nameBitSet = testCasesModels.TestCaseMetaDataForDomains.UniqueMetaDataBitSets.MetaDataGroupItemsBitSetMap[metaDataItem.MetaDataName]
						valueBitSet = testCasesModels.TestCaseMetaDataForDomains.UniqueMetaDataBitSets.MetaDataGroupItemValuesBitSetMap[val]

						metaDataEntry, err = boolbits.NewEntry(domainBitSet, groupBitSet, nameBitSet, valueBitSet)

						if err != nil {
							errorID := "29cfc2fb-08b4-4ddb-a330-73188b32d101"
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
						entryKey = fmt.Sprintf("%s.%s.%s",
							domainUUid,
							metaDataItem.MetaDataGroupName,
							metaDataItem.MetaDataName)

						valueEntryListToBeProcessedWithBooleanOrSlice = append(valueEntryListToBeProcessedWithBooleanOrSlice, metaDataEntry)

						// 'valueEntryListToBeProcessedWithBooleanOrSlice' to 'simpleMetaDataFilterEntryMap'
						listTestSuiteUIObject.simpleMetaDataFilterEntryMap[entryKey] = simpleMetaDataFilterEntryMapStruct{
							valueEntryListToBeProcessedWithBooleanOrSlice: valueEntryListToBeProcessedWithBooleanOrSlice}

					}

					// Check if auto-filter is enabled. If so then calculate the new TestSuite-liset
					if listTestSuiteUIObject.useAutoFilter == true {
						listTestSuiteUIObject.calculateMetaDataFilterFunction()
					}

				})
				// Extract Selected values from TestSuite
				var selectedValue string
				if metaDataGroupFromSourceExistInTestSuiteMap == true && metaDataGroupItemFromSourceExistInTestSuiteMap == true {
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
				w := listTestSuiteUIObject.calcSelectWidth(metaDataItem.AvailableMetaDataValues)
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
					entryKey = fmt.Sprintf("%s.%s.%s",
						domainUUid,
						metaDataItem.MetaDataGroupName,
						metaDataItem.MetaDataName)

					delete(listTestSuiteUIObject.simpleMetaDataFilterEntryMap, entryKey)

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

						// Add MetaDataEntry to Entry-map
						entryKey = fmt.Sprintf("%s.%s.%s",
							domainUUid,
							metaDataItem.MetaDataGroupName,
							metaDataItem.MetaDataName)

						// Loop all selected values and create its MetaDataEntry
						var valueEntryListToBeProcessedWithBooleanOrSlice []*boolbits.Entry
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

							valueEntryListToBeProcessedWithBooleanOrSlice = append(valueEntryListToBeProcessedWithBooleanOrSlice, metaDataEntry)

						}
						// 'valueEntryListToBeProcessedWithBooleanOrSlice' to 'simpleMetaDataFilterEntryMap'
						listTestSuiteUIObject.simpleMetaDataFilterEntryMap[entryKey] = simpleMetaDataFilterEntryMapStruct{
							valueEntryListToBeProcessedWithBooleanOrSlice: valueEntryListToBeProcessedWithBooleanOrSlice}

					}

					// Check if auto-filter is enabled. If so then calculate the new TestSuite-liset
					if listTestSuiteUIObject.useAutoFilter == true {
						listTestSuiteUIObject.calculateMetaDataFilterFunction()
					}

				})

				// Extract Selected values from TestSuite
				var selectedValues []string
				if metaDataGroupFromSourceExistInTestSuiteMap == true && metaDataGroupItemFromSourceExistInTestSuiteMap == true {
					for _, availableValue := range metaDataItem.AvailableMetaDataValues {
						if _, ok := metaDataGroupFromTestSuite[metaDataItem.MetaDataName].SelectedMetaDataValuesForMultiSelectMap[availableValue]; ok {
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

				w := listTestSuiteUIObject.calcCheckGroupWidth(metaDataItem.AvailableMetaDataValues)
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

	returnContainer := container.New(
		layout.NewVBoxLayout(),
		metaDataGroupCards...,
	)
	returnContainer.Add(layout.NewSpacer())

	return returnContainer
}

// calcSelectWidth returns the width needed to show the longest option
func (listTestSuiteUIObject *ListTestSuiteUIStruct) calcSelectWidth(values []string) float32 {
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
func (listTestSuiteUIObject *ListTestSuiteUIStruct) calcCheckGroupWidth(values []string) float32 {
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
func (listTestSuiteUIObject *ListTestSuiteUIStruct) ConvertMetaDataToNewMap(tc *testCaseModel.TestCaseMetaDataStruct) map[string]map[string]*NewMetaDataInGroupStruct {
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
