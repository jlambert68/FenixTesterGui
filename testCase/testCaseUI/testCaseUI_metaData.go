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

		var metaDataGroupsAsCanvasObject fyne.CanvasObject
		metaDataGroupsAsCanvasObject = buildGUIFromMetaDataGroupsSlice(metaDataGroupsPtr)

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
func buildGUIFromMetaDataGroupsSlice(metaDataGroupsPtr *[]*testCaseModel.MetaDataGroupStruct) fyne.CanvasObject {

	// Create one “card” per MetaDataGroup
	var metaDataGroupCards []fyne.CanvasObject
	metaDataGroupCards = make([]fyne.CanvasObject, 0, len(*metaDataGroupsPtr))

	// Loop all MetaData-groups
	for _, metaDataGroupPtr := range *metaDataGroupsPtr {

		// unpack the slice of *MetaDataInGroupStruct
		var metaDataItemsInGroupPtr *[]*testCaseModel.MetaDataInGroupStruct
		metaDataItemsInGroupPtr = metaDataGroupPtr.MetaDataInGroupPtr

		var metaDataItemsAsCanvasObject []fyne.CanvasObject
		metaDataItemsAsCanvasObject = make([]fyne.CanvasObject, 0, len(*metaDataGroupPtr.MetaDataInGroupPtr))

		// Loop all MetaDataItems in the MetaData-group
		for _, metaDataItemPtr := range *metaDataItemsInGroupPtr {

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
				})
				sel.PlaceHolder = "Choose..."
				// apply the existing selection if any
				if metaDataItem.SelectedMetaDataValueForSingleSelect != "" {
					sel.SetSelected(metaDataItem.SelectedMetaDataValueForSingleSelect)
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
				// apply existing selections
				chk.Selected = append([]string(nil), metaDataItem.SelectedMetaDataValuesForMultiSelect...)
				chk.Refresh()

				metaDataItemsAsCanvasObject = append(metaDataItemsAsCanvasObject,
					container.NewVBox(
						widget.NewLabel(label),
						chk,
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
