package testSuiteUI

import (
	"FenixTesterGui/testSuites/testSuitesModel"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
)

// Generate the MetaData Area for the TestCase
func (testSuiteUiModel TestSuiteUiStruct) GenerateMetaDataAreaForTestCase() (
	testSuiteMetaDataContainer *fyne.Container,
	err error) {

	// Get OwnerDomain
	var ownerDomainUuid string
	ownerDomainUuid = testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestSuiteOwnerDomainUuid

	var metaDataAccordionItem *widget.AccordionItem

	var existsInMap bool

	// Check if there is an OwnerDomain selected
	if len(testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestSuiteOwnerDomainUuid) > 0 {

		// Get the MetaDataGroups depending on Domain
		var metaDataGroupsMapPtr *map[string]*testSuitesModel.MetaDataGroupStruct
		//var metaDataGroupsMap map[string]*testSuitesModel.MetaDataGroupStruct
		var testSuiteMetaDataForDomainsMap map[string]*testSuitesModel.TestSuiteMetaDataForDomainsForMapStruct
		var testSuiteMetaDataForDomainPtr *testSuitesModel.TestSuiteMetaDataForDomainsForMapStruct
		var testSuiteMetaDataForDomain testSuitesModel.TestSuiteMetaDataForDomainsForMapStruct
		testSuiteMetaDataForDomainsMap = testSuitesModel.TestSuitesModelPtr.TestSuiteMetaDataForDomains.
			TestSuiteMetaDataForDomainsMap
		testSuiteMetaDataForDomainPtr, existsInMap = testSuiteMetaDataForDomainsMap[ownerDomainUuid]
		if existsInMap == false {

			testSuiteMetaDataContainer = container.NewVBox(widget.NewLabel("OwnerDomain doesn't have any TestSuite MetaData"))

			return testSuiteMetaDataContainer,

				err

		}

		testSuiteMetaDataForDomain = *testSuiteMetaDataForDomainPtr
		var tempMetaDataGroupsOrder []string
		metaDataGroupsMapPtr, tempMetaDataGroupsOrder = testSuitesModel.ConvertTestSuiteMetaData(testSuiteMetaDataForDomain.TestSuiteMetaDataForDomainPtr)

		// Get pointer to Structure holding selected values in TestSuite
		var metaDataGroupInTestSuitePtr *testSuitesModel.TestSuiteMetaDataStruct
		metaDataGroupInTestSuitePtr = testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestSuiteMetaDataPtr
		if metaDataGroupInTestSuitePtr == nil {
			metaDataGroupInTestSuitePtr = &testSuitesModel.TestSuiteMetaDataStruct{
				CurrentSelectedDomainUuid:                               ownerDomainUuid,
				TestSuiteMetaDataMessageJsonForTestSuiteWhenLastSaved:   nil,
				TestSuiteMetaDataMessageStructForTestSuiteWhenLastSaved: nil,
				MetaDataGroupsOrder:                                     tempMetaDataGroupsOrder,
				MetaDataGroupsMapPtr:                                    metaDataGroupsMapPtr,
			}
		}

		// Generate TestSuiteMeta-UI-object
		var metaDataGroupsAsCanvasObject fyne.CanvasObject
		metaDataGroupsAsCanvasObject = testSuiteUiModel.buildGUIFromMetaDataGroupsMap(
			tempMetaDataGroupsOrder,
			metaDataGroupsMapPtr,
			metaDataGroupInTestSuitePtr)

		myContainer := container.NewBorder(nil, nil, nil, nil, metaDataGroupsAsCanvasObject)

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

		// No OwnerDomain selected
		myContainer := container.New(layout.NewGridLayout(1), widget.NewLabel("MetaData is available when 'Owner Domain' is selected"))

		// Create an Accordion item for the MetaData
		metaDataAccordionItem = widget.NewAccordionItem("TestCase MetaData", myContainer)
	}

	var testSuiteMetaDataAccordion *widget.Accordion
	testSuiteMetaDataAccordion = widget.NewAccordion(metaDataAccordionItem)

	// Open all for the Accordion
	// Will only been done when a TestSuite hasn't been locked down due to user selected OwnerDomain and Environment
	if testSuiteUiModel.TestSuiteModelPtr.HasLockButtonBeenClickedAndBothOwnerDomainAndTestEnvironmentHaveValues() == false {
		testSuiteMetaDataAccordion.OpenAll()
	}

	// Create the VBox-container that will be returned
	testSuiteMetaDataContainer = container.NewVBox(testSuiteMetaDataAccordion, widget.NewLabel(""), widget.NewSeparator())

	return testSuiteMetaDataContainer, err
}

// buildGUIFromSlice builds a Container from your slice pointer
func (testSuiteUiModel TestSuiteUiStruct) buildGUIFromMetaDataGroupsMap(
	metaDataGroupsOrder []string,
	metaDataGroupsSourceMapPtr *map[string]*testSuitesModel.MetaDataGroupStruct,
	metaDataGroupInTestSuitePtr *testSuitesModel.TestSuiteMetaDataStruct) fyne.CanvasObject {

	// Get 'metaDataGroupInTestSuitePtr' from last saved version, which exists "untouched" when this is created
	metaDataGroupInTestSuitePtr = testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestSuiteMetaDataPtr

	// Get the 'metaDataGroupsSourceMap'
	var metaDataGroupsSourceMap map[string]*testSuitesModel.MetaDataGroupStruct
	metaDataGroupsSourceMap = *metaDataGroupsSourceMapPtr

	if len(*metaDataGroupsSourceMapPtr) != len(metaDataGroupsOrder) {
		log.Fatalln("ERROR: The number of MetaDataGroups in the 'metaDataGroupsSourceMap' doesn't match the number of MetaDataGroups in the 'metaDataGroupsOrder'")
	}

	var convertMetaDataToMapMap map[string]map[string]*NewMetaDataInGroupStruct
	convertMetaDataToMapMap = testSuiteUiModel.convertMetaDataToNewMap(metaDataGroupInTestSuitePtr)

	// Create one “card” per MetaDataGroup
	var metaDataGroupCards []fyne.CanvasObject
	metaDataGroupCards = make([]fyne.CanvasObject, 0, len(*metaDataGroupsSourceMapPtr))

	var metaDataGroupFromTestSuite map[string]*NewMetaDataInGroupStruct
	var newMetaDataItemInGroup *NewMetaDataInGroupStruct
	var metaDataGroupFromSourceExistInTestCaseMap bool
	var metaDataGroupItemFromSourceExistInTestCaseMap bool
	var existInMap bool

	// Loop all MetaData-groups
	for _, metaDataGroupName := range metaDataGroupsOrder {

		var metaDataGroupPtr *testSuitesModel.MetaDataGroupStruct
		metaDataGroupPtr = metaDataGroupsSourceMap[metaDataGroupName]

		// Get the MetaDataGroupName from the TestSuite
		metaDataGroupFromTestSuite, metaDataGroupFromSourceExistInTestCaseMap = convertMetaDataToMapMap[metaDataGroupPtr.MetaDataGroupName]

		// unpack the slice of *MetaDataInGroupStruct
		var metaDataItemsInGroupPtr *map[string]*testSuitesModel.MetaDataInGroupStruct
		metaDataItemsInGroupPtr = metaDataGroupPtr.MetaDataInGroupMapPtr

		// Get the metaDataItemsInGroupMap
		var metaDataItemsInGroupMap map[string]*testSuitesModel.MetaDataInGroupStruct
		metaDataItemsInGroupMap = *metaDataItemsInGroupPtr

		var metaDataItemsAsCanvasObject []fyne.CanvasObject
		metaDataItemsAsCanvasObject = make([]fyne.CanvasObject, 0, len(*metaDataGroupPtr.MetaDataInGroupMapPtr))

		// Loop all MetaDataItems in the MetaData-group
		for _, metaDataItemName := range metaDataGroupPtr.MetaDataInGroupOrder {

			//metaDataGroupPtr.MetaDataInGroupOrder är tom

			var metaDataItemPtr *testSuitesModel.MetaDataInGroupStruct
			metaDataItemPtr = metaDataItemsInGroupMap[metaDataItemName]

			if metaDataGroupFromSourceExistInTestCaseMap == true {
				newMetaDataItemInGroup, metaDataGroupItemFromSourceExistInTestCaseMap = metaDataGroupFromTestSuite[metaDataItemPtr.MetaDataName]
			}

			var metaDataItem testSuitesModel.MetaDataInGroupStruct
			metaDataItem = *metaDataItemPtr
			// append '*' to the label if it's mandatory
			label := metaDataItem.MetaDataName
			if metaDataItem.Mandatory {
				label += " *"
			}

			// Create correct widget depending on if the item is SingleSelect or MultiSelect
			switch metaDataItem.SelectType {

			case testSuitesModel.MetaDataSelectType_SingleSelect:

				var valueIsValidWarningBox *canvas.Rectangle

				// Generate Warnings-rectangle for valid value, or that value exist
				//var valueIsValidWarningBox *canvas.Rectangle
				var colorToUse color.NRGBA
				colorToUse = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
				valueIsValidWarningBox = canvas.NewRectangle(colorToUse)

				sel := widget.NewSelect(metaDataItem.AvailableMetaDataValues, func(val string) {

					// store value in TestCase-version of the MetaData
					metaDataItem.SelectedMetaDataValueForSingleSelect = val

					// If the 'MetaDataGroupsMap' exist
					if metaDataGroupInTestSuitePtr.MetaDataGroupsMapPtr == nil {
						// No 'MetaDataGroupsMap'

						var tempMetaDataInGroupMap map[string]*testSuitesModel.MetaDataInGroupStruct
						tempMetaDataInGroupMap = make(map[string]*testSuitesModel.MetaDataInGroupStruct)

						tempSelectedMetaDataValuesForMultiSelectMapPtr := make(map[string]string)

						// Create MetaData for Group in TestCase
						var tempMetaDataInGroup testSuitesModel.MetaDataInGroupStruct
						tempMetaDataInGroup = testSuitesModel.MetaDataInGroupStruct{
							MetaDataGroupName:                          metaDataItem.MetaDataGroupName,
							MetaDataName:                               metaDataItem.MetaDataName,
							SelectType:                                 metaDataItem.SelectType,
							Mandatory:                                  metaDataItem.Mandatory,
							AvailableMetaDataValues:                    metaDataItem.AvailableMetaDataValues,
							SelectedMetaDataValueForSingleSelect:       val,
							SelectedMetaDataValuesForMultiSelect:       nil,
							SelectedMetaDataValuesForMultiSelectMapPtr: &tempSelectedMetaDataValuesForMultiSelectMapPtr,
						}

						// Add MetaData for Group to 'MetaDataGroupsMap' in TestCase
						tempMetaDataInGroupMap[metaDataItem.MetaDataName] = &tempMetaDataInGroup

						// Create the 'MetaDataGroup'
						var testMetaDataGroup testSuitesModel.MetaDataGroupStruct
						testMetaDataGroup = testSuitesModel.MetaDataGroupStruct{
							MetaDataGroupName:     metaDataItem.MetaDataGroupName,
							MetaDataInGroupOrder:  metaDataGroupsOrder,
							MetaDataInGroupMapPtr: &tempMetaDataInGroupMap,
						}

						// Create the 'MetaDataGroupsMap' to be able to add the 'MetaDataGroup'
						var tempMetaDataGroupsMap map[string]*testSuitesModel.MetaDataGroupStruct
						tempMetaDataGroupsMap = make(map[string]*testSuitesModel.MetaDataGroupStruct)

						// Add the 'MetaDataGroup' to the 'MetaDataGroupsMap'
						tempMetaDataGroupsMap[metaDataItem.MetaDataGroupName] = &testMetaDataGroup

						metaDataGroupInTestSuitePtr.MetaDataGroupsMapPtr = &tempMetaDataGroupsMap

					} else {
						// 'MetaDataGroupsMap' exists, so get it
						var tempMetaDataGroupsMapPtr *map[string]*testSuitesModel.MetaDataGroupStruct
						var tempMetaDataGroupsMap map[string]*testSuitesModel.MetaDataGroupStruct
						tempMetaDataGroupsMapPtr = metaDataGroupInTestSuitePtr.MetaDataGroupsMapPtr
						tempMetaDataGroupsMap = *tempMetaDataGroupsMapPtr

						// Check for specific MetaDataGroup
						var tempMetaDataGroupPtr *testSuitesModel.MetaDataGroupStruct
						tempMetaDataGroupPtr, existInMap = tempMetaDataGroupsMap[metaDataItem.MetaDataGroupName]
						if existInMap == false {
							// Specific MetaDataGroupName doesn't exist

							var tempMetaDataInGroupMap map[string]*testSuitesModel.MetaDataInGroupStruct
							tempMetaDataInGroupMap = make(map[string]*testSuitesModel.MetaDataInGroupStruct)

							tempSelectedMetaDataValuesForMultiSelectMapPtr := make(map[string]string)

							// Create the specific MetaDataItem
							var tempMetaDataInGroup testSuitesModel.MetaDataInGroupStruct
							tempMetaDataInGroup = testSuitesModel.MetaDataInGroupStruct{
								MetaDataGroupName:                          metaDataItem.MetaDataGroupName,
								MetaDataName:                               metaDataItem.MetaDataName,
								SelectType:                                 metaDataItem.SelectType,
								Mandatory:                                  metaDataItem.Mandatory,
								AvailableMetaDataValues:                    metaDataItem.AvailableMetaDataValues,
								SelectedMetaDataValueForSingleSelect:       val,
								SelectedMetaDataValuesForMultiSelect:       nil,
								SelectedMetaDataValuesForMultiSelectMapPtr: &tempSelectedMetaDataValuesForMultiSelectMapPtr,
							}

							// Add MetaData for Group to 'MetaDataGroupsMap' in TestCase
							tempMetaDataInGroupMap[metaDataItem.MetaDataName] = &tempMetaDataInGroup

							// Create the 'MetaDataGroup'
							tempMetaDataGroupPtr = &testSuitesModel.MetaDataGroupStruct{
								MetaDataGroupName:     metaDataItem.MetaDataGroupName,
								MetaDataInGroupOrder:  metaDataGroupsOrder,
								MetaDataInGroupMapPtr: &tempMetaDataInGroupMap,
							}

							// Save MetaDataGroup with Item in 'MetaDataGroupsMap'
							tempMetaDataGroupsMap[metaDataItem.MetaDataGroupName] = tempMetaDataGroupPtr

						} else {

							// Specific MetaDataGroupName does exist
							var tempMetaDataGroup testSuitesModel.MetaDataGroupStruct
							tempMetaDataGroup = *tempMetaDataGroupPtr

							// Get map for MetaDataGroupItems
							var tempMetaDataInGroupMap map[string]*testSuitesModel.MetaDataInGroupStruct
							tempMetaDataInGroupMap = *tempMetaDataGroup.MetaDataInGroupMapPtr

							// Check if MetaDataItem exist
							var tempMetaDataInGroupPtr *testSuitesModel.MetaDataInGroupStruct
							tempMetaDataInGroupPtr, existInMap = tempMetaDataInGroupMap[metaDataItem.MetaDataName]

							if existInMap == false {

								tempSelectedMetaDataValuesForMultiSelectMapPtr := make(map[string]string)

								// Create the specific MetaDataItem, because it doesn't exist
								var tempMetaDataInGroup testSuitesModel.MetaDataInGroupStruct
								tempMetaDataInGroup = testSuitesModel.MetaDataInGroupStruct{
									MetaDataGroupName:                          metaDataItem.MetaDataGroupName,
									MetaDataName:                               metaDataItem.MetaDataName,
									SelectType:                                 metaDataItem.SelectType,
									Mandatory:                                  metaDataItem.Mandatory,
									AvailableMetaDataValues:                    metaDataItem.AvailableMetaDataValues,
									SelectedMetaDataValueForSingleSelect:       val,
									SelectedMetaDataValuesForMultiSelect:       nil,
									SelectedMetaDataValuesForMultiSelectMapPtr: &tempSelectedMetaDataValuesForMultiSelectMapPtr,
								}

								// Add MetaData for Group to 'MetaDataGroupsMap' in TestCase
								tempMetaDataInGroupMap[metaDataItem.MetaDataName] = &tempMetaDataInGroup

								// Create the 'MetaDataGroup'
								tempMetaDataGroupPtr = &testSuitesModel.MetaDataGroupStruct{
									MetaDataGroupName:     metaDataItem.MetaDataGroupName,
									MetaDataInGroupOrder:  metaDataGroupsOrder,
									MetaDataInGroupMapPtr: &tempMetaDataInGroupMap,
								}

								// Save MetaDataGroup with Item in 'MetaDataGroupsMap'
								tempMetaDataGroupsMap[metaDataItem.MetaDataGroupName] = tempMetaDataGroupPtr

							} else {
								// MetaDataItem does exist, so get it
								var tempMetaDataInGroup testSuitesModel.MetaDataInGroupStruct
								tempMetaDataInGroup = *tempMetaDataInGroupPtr

								// Set selected value for the TestDataItem
								tempMetaDataInGroup.SelectedMetaDataValueForSingleSelect = val

								// Add MetaData for Group to 'MetaDataGroupsMap' in TestCase
								tempMetaDataInGroupMap[metaDataItem.MetaDataName] = &tempMetaDataInGroup

							}
						}

					}

					// Save Changes to TestSuite regarding MetaData
					testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestSuiteMetaDataPtr = metaDataGroupInTestSuitePtr

					// Set Warning box that value is not selected
					if len(val) == 0 && metaDataItem.Mandatory == true {
						valueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
					} else {
						valueIsValidWarningBox.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
					}

				})

				// Extract Selected values from TestSuite
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

				// If MetaDataItemName == 'TestEnvironment' then overrule and set from separate TestEnvironment-section
				if metaDataItemName == "TestEnvironment" {
					sel.SetSelected(testSuiteUiModel.TestSuiteModelPtr.TestSuiteUIModelBinding.TestSuiteExecutionEnvironment)

					// Disable for change due to that TestEnvironment is set from separate TestEnvironment Section
					sel.Disable()

				}

				// Resize the DropDown

				// Create a custom SelectComboBox, with valueIsValidWarningBox
				var customMetadataSelectComboBox *customSelectComboBox
				customMetadataSelectComboBox = newCustomSelectComboBoxWidget(sel, valueIsValidWarningBox)

				// wrap in a 1-cell grid to force width
				w := calcSelectWidth(metaDataItem.AvailableMetaDataValues)
				wrapper := container.New(
					layout.NewGridWrapLayout(fyne.NewSize(w, sel.MinSize().Height)),
					customMetadataSelectComboBox,
				)

				// Set Warning box that value is not selected
				if len(customMetadataSelectComboBox.selectComboBox.Selected) == 0 && metaDataItem.Mandatory == true {
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

			case testSuitesModel.MetaDataSelectType_MultiSelect:

				var valueIsValidWarningBox *canvas.Rectangle

				// Generate Warnings-rectangle for valid value, or that value exist
				//var valueIsValidWarningBox *canvas.Rectangle
				var colorToUse color.NRGBA
				colorToUse = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
				valueIsValidWarningBox = canvas.NewRectangle(colorToUse)

				var chk *widget.CheckGroup
				chk = widget.NewCheckGroup(metaDataItem.AvailableMetaDataValues, func(vals []string) {

					// If the 'MetaDataGroupsMap' exist
					if metaDataGroupInTestSuitePtr.MetaDataGroupsMapPtr == nil {
						// No 'MetaDataGroupsMap'

						var tempMetaDataInGroupMap map[string]*testSuitesModel.MetaDataInGroupStruct
						tempMetaDataInGroupMap = make(map[string]*testSuitesModel.MetaDataInGroupStruct)

						// Create MetaData for Group in TestCase
						var tempMetaDataInGroup testSuitesModel.MetaDataInGroupStruct
						tempMetaDataInGroup = testSuitesModel.MetaDataInGroupStruct{
							MetaDataGroupName:                          metaDataItem.MetaDataGroupName,
							MetaDataName:                               metaDataItem.MetaDataName,
							SelectType:                                 metaDataItem.SelectType,
							Mandatory:                                  metaDataItem.Mandatory,
							AvailableMetaDataValues:                    metaDataItem.AvailableMetaDataValues,
							SelectedMetaDataValueForSingleSelect:       "",
							SelectedMetaDataValuesForMultiSelect:       vals,
							SelectedMetaDataValuesForMultiSelectMapPtr: nil,
						}

						// Loop the values and create the 'SelectedMetaDataValuesForMultiSelectMap'
						var tempSelectedMetaDataValuesForMultiSelectMap map[string]string
						tempSelectedMetaDataValuesForMultiSelectMap = make(map[string]string)
						for _, value := range vals {
							tempSelectedMetaDataValuesForMultiSelectMap[value] = value
						}
						// Add the map the 'MetaData'-object
						tempMetaDataInGroup.SelectedMetaDataValuesForMultiSelectMapPtr = &tempSelectedMetaDataValuesForMultiSelectMap

						// Add MetaData for Group to 'MetaDataGroupsMap' in TestCase
						tempMetaDataInGroupMap[metaDataItem.MetaDataName] = &tempMetaDataInGroup

						// Create the 'MetaDataGroup'
						var testMetaDataGroup testSuitesModel.MetaDataGroupStruct
						testMetaDataGroup = testSuitesModel.MetaDataGroupStruct{
							MetaDataGroupName:     metaDataItem.MetaDataGroupName,
							MetaDataInGroupMapPtr: &tempMetaDataInGroupMap,
						}

						// Create the 'MetaDataGroupsMap' to be able to add the 'MetaDataGroup'
						var tempMetaDataGroupsMap map[string]*testSuitesModel.MetaDataGroupStruct
						tempMetaDataGroupsMap = make(map[string]*testSuitesModel.MetaDataGroupStruct)

						// Add the 'MetaDataGroup' to the 'MetaDataGroupsMap'
						tempMetaDataGroupsMap[metaDataItem.MetaDataGroupName] = &testMetaDataGroup

						metaDataGroupInTestSuitePtr.MetaDataGroupsMapPtr = &tempMetaDataGroupsMap

					} else {
						// 'MetaDataGroupsMap' exists, so get it
						var tempMetaDataGroupsMapPtr *map[string]*testSuitesModel.MetaDataGroupStruct
						var tempMetaDataGroupsMap map[string]*testSuitesModel.MetaDataGroupStruct
						tempMetaDataGroupsMapPtr = metaDataGroupInTestSuitePtr.MetaDataGroupsMapPtr
						tempMetaDataGroupsMap = *tempMetaDataGroupsMapPtr

						// Check for specific MetaDataGroup
						var tempMetaDataGroupPtr *testSuitesModel.MetaDataGroupStruct
						tempMetaDataGroupPtr, existInMap = tempMetaDataGroupsMap[metaDataItem.MetaDataGroupName]
						if existInMap == false {
							// Specific MetaDataGroupName doesn't exist

							var tempMetaDataInGroupMap map[string]*testSuitesModel.MetaDataInGroupStruct
							tempMetaDataInGroupMap = make(map[string]*testSuitesModel.MetaDataInGroupStruct)

							// Create the specific MetaDataItem
							var tempMetaDataInGroup testSuitesModel.MetaDataInGroupStruct
							tempMetaDataInGroup = testSuitesModel.MetaDataInGroupStruct{
								MetaDataGroupName:                          metaDataItem.MetaDataGroupName,
								MetaDataName:                               metaDataItem.MetaDataName,
								SelectType:                                 metaDataItem.SelectType,
								Mandatory:                                  metaDataItem.Mandatory,
								AvailableMetaDataValues:                    metaDataItem.AvailableMetaDataValues,
								SelectedMetaDataValueForSingleSelect:       "",
								SelectedMetaDataValuesForMultiSelect:       vals,
								SelectedMetaDataValuesForMultiSelectMapPtr: nil,
							}

							// Loop the values and create the 'SelectedMetaDataValuesForMultiSelectMap'
							var tempSelectedMetaDataValuesForMultiSelectMap map[string]string
							tempSelectedMetaDataValuesForMultiSelectMap = make(map[string]string)
							for _, value := range vals {
								tempSelectedMetaDataValuesForMultiSelectMap[value] = value
							}
							// Add the map the 'MetaData'-object
							tempMetaDataInGroup.SelectedMetaDataValuesForMultiSelectMapPtr = &tempSelectedMetaDataValuesForMultiSelectMap

							// Add MetaData for Group to 'MetaDataGroupsMap' in TestCase
							tempMetaDataInGroupMap[metaDataItem.MetaDataName] = &tempMetaDataInGroup

							// Create the 'MetaDataGroup'
							tempMetaDataGroupPtr = &testSuitesModel.MetaDataGroupStruct{
								MetaDataGroupName:     metaDataItem.MetaDataGroupName,
								MetaDataInGroupMapPtr: &tempMetaDataInGroupMap,
							}

							// Save MetaDataGroup with Item in 'MetaDataGroupsMap'
							tempMetaDataGroupsMap[metaDataItem.MetaDataGroupName] = tempMetaDataGroupPtr

							metaDataGroupInTestSuitePtr.MetaDataGroupsMapPtr = &tempMetaDataGroupsMap

						} else {

							// Specific MetaDataGroupName does exist
							var tempMetaDataGroup testSuitesModel.MetaDataGroupStruct
							tempMetaDataGroup = *tempMetaDataGroupPtr

							// Get map for MetaDataGroupItems
							var tempMetaDataInGroupMap map[string]*testSuitesModel.MetaDataInGroupStruct
							tempMetaDataInGroupMap = *tempMetaDataGroup.MetaDataInGroupMapPtr

							// Check if MetaDataItem exist
							var tempMetaDataInGroupPtr *testSuitesModel.MetaDataInGroupStruct
							tempMetaDataInGroupPtr, existInMap = tempMetaDataInGroupMap[metaDataItem.MetaDataName]

							if existInMap == false {

								// Create the specific MetaDataItem, because it doesn't exist
								var tempMetaDataInGroup testSuitesModel.MetaDataInGroupStruct
								tempMetaDataInGroup = testSuitesModel.MetaDataInGroupStruct{
									MetaDataGroupName:                          metaDataItem.MetaDataGroupName,
									MetaDataName:                               metaDataItem.MetaDataName,
									SelectType:                                 metaDataItem.SelectType,
									Mandatory:                                  metaDataItem.Mandatory,
									AvailableMetaDataValues:                    metaDataItem.AvailableMetaDataValues,
									SelectedMetaDataValueForSingleSelect:       "",
									SelectedMetaDataValuesForMultiSelect:       vals,
									SelectedMetaDataValuesForMultiSelectMapPtr: nil,
								}

								// Loop the values and create the 'SelectedMetaDataValuesForMultiSelectMap'
								var tempSelectedMetaDataValuesForMultiSelectMap map[string]string
								tempSelectedMetaDataValuesForMultiSelectMap = make(map[string]string)
								for _, value := range vals {
									tempSelectedMetaDataValuesForMultiSelectMap[value] = value
								}

								// Add the map the 'MetaData'-object
								tempMetaDataInGroup.SelectedMetaDataValuesForMultiSelectMapPtr = &tempSelectedMetaDataValuesForMultiSelectMap

								// Add MetaData for Group to 'MetaDataGroupsMap' in TestCase
								tempMetaDataInGroupMap[metaDataItem.MetaDataName] = &tempMetaDataInGroup

								// Create the 'MetaDataGroup'
								tempMetaDataGroupPtr = &testSuitesModel.MetaDataGroupStruct{
									MetaDataGroupName:     metaDataItem.MetaDataGroupName,
									MetaDataInGroupOrder:  metaDataGroupsOrder,
									MetaDataInGroupMapPtr: &tempMetaDataInGroupMap,
								}

								// Save MetaDataGroup with Item in 'MetaDataGroupsMap'
								tempMetaDataGroupsMap[metaDataItem.MetaDataGroupName] = tempMetaDataGroupPtr

								metaDataGroupInTestSuitePtr.MetaDataGroupsMapPtr = &tempMetaDataGroupsMap

							} else {
								// MetaDataItem does exist, so get it
								var tempMetaDataInGroup testSuitesModel.MetaDataInGroupStruct
								tempMetaDataInGroup = *tempMetaDataInGroupPtr

								// Set selected value for the TestDataItem
								tempMetaDataInGroup.SelectedMetaDataValuesForMultiSelect = vals

								// Loop the values and create the 'SelectedMetaDataValuesForMultiSelectMap'
								var tempSelectedMetaDataValuesForMultiSelectMap map[string]string
								tempSelectedMetaDataValuesForMultiSelectMap = make(map[string]string)
								for _, value := range vals {
									tempSelectedMetaDataValuesForMultiSelectMap[value] = value
								}

								// Add the map the 'MetaData'-object
								tempMetaDataInGroup.SelectedMetaDataValuesForMultiSelectMapPtr = &tempSelectedMetaDataValuesForMultiSelectMap

								// Add MetaData for Group to 'MetaDataGroupsMap' in TestCase
								tempMetaDataInGroupMap[metaDataItem.MetaDataName] = &tempMetaDataInGroup

								// Create the 'MetaDataGroup'
								tempMetaDataGroupPtr = &testSuitesModel.MetaDataGroupStruct{
									MetaDataInGroupOrder:  metaDataGroupsOrder,
									MetaDataGroupName:     metaDataItem.MetaDataGroupName,
									MetaDataInGroupMapPtr: &tempMetaDataInGroupMap,
								}

								// Save MetaDataGroup with Item in 'MetaDataGroupsMap'
								tempMetaDataGroupsMap[metaDataItem.MetaDataGroupName] = tempMetaDataGroupPtr

								metaDataGroupInTestSuitePtr.MetaDataGroupsMapPtr = &tempMetaDataGroupsMap

								// Add MetaData for Group to 'MetaDataGroupsMap' in TestCase
								tempMetaDataInGroupMap[metaDataItem.MetaDataName] = &tempMetaDataInGroup

							}
						}

					}

					// Set Warning box that value is not selected
					if len(vals) == 0 && metaDataItem.Mandatory == true {
						valueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
					} else {
						valueIsValidWarningBox.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
					}

				})

				// Extract Selected values from TestCase
				var selectedValues []string
				if metaDataGroupFromSourceExistInTestCaseMap == true && metaDataGroupItemFromSourceExistInTestCaseMap == true {
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
				var customCheckBoxGroup *customAttributeCheckBoxGroup
				customCheckBoxGroup = newCustomAttributeCheckBoxGroupWidget(chk, valueIsValidWarningBox)

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

// NewMetaDataInGroupStruct
// Struct holding the available values, how they are selected and what was selected
type NewMetaDataInGroupStruct struct {
	MetaDataGroupName                       string                             // The name of the MetaData-Group
	MetaDataName                            string                             // The name of the MetaData-post
	SelectType                              testSuitesModel.MetaDataSelectType // Is the MetaData-post single- or multi-select
	Mandatory                               bool                               // Is the MetaData-post mandatory or not
	SelectedMetaDataValueForSingleSelect    string                             // The value selected for single select
	SelectedMetaDataValuesForMultiSelectMap map[string]string                  // The values selected for multi select
}

// convertMetaDataToNewMap transforms the TestsUITEMetaDataStruct.MetaDataGroupsSlicePtr
// into a nested map[groupName][metaDataName] => *NewMetaDataInGroupStruct.
func (testSuiteUiModel TestSuiteUiStruct) convertMetaDataToNewMap(
	ts *testSuitesModel.TestSuiteMetaDataStruct) map[string]map[string]*NewMetaDataInGroupStruct {
	result := make(map[string]map[string]*NewMetaDataInGroupStruct)

	if ts == nil {
		return result
	}

	if ts.MetaDataGroupsMapPtr == nil {
		return result
	}

	for _, grp := range *ts.MetaDataGroupsMapPtr {
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
