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
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateMetaDataAreaForTestCase(
	testCaseUuid string,
	domainUuidToGetMetaDataFor string) (
	testCaseMetaDataArea fyne.CanvasObject, err error) {

	var metaDataAccordionItem *widget.AccordionItem
	var accordion *widget.Accordion
	var metaDataArea fyne.CanvasObject

	//
	if len(domainUuidToGetMetaDataFor) > 0 {

		// Get current TestCase-UI-model
		_, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]

		if existsInMap == true {
			errorId := "bcb9d984-3106-42b6-9c23-288ec6d26224"
			err = errors.New(fmt.Sprintf("testcase-UI-model with sourceUuid '%s' allready exist in 'TestCasesUiModelMap' [ErrorID: %s]", testCaseUuid, errorId))

			return nil, err
		}

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

			return nil, err
		}

		testCaseMetaDataForDomain = *testCaseMetaDataForDomainPtr
		metaDataGroupsPtr = testCaseModel.ConvertTestCaseMetaData(testCaseMetaDataForDomain.TestCaseMetaDataForDomainPtr)

		var metaDataGroupsAsCanvasObject fyne.CanvasObject
		metaDataGroupsAsCanvasObject = buildGUIFromMetaDataGroupsSlice(metaDataGroupsPtr)

		// Create an Accordion item for the MetaData
		metaDataAccordionItem = widget.NewAccordionItem("TestCase MetaData", metaDataGroupsAsCanvasObject)
	} else {
		// Create an Accordion item for the MetaData
		metaDataAccordionItem = widget.NewAccordionItem("TestCase MetaData", widget.NewLabel("MetaData is available when 'Owner Domain' is selected"))
	}

	accordion = widget.NewAccordion(metaDataAccordionItem)

	// Create the VBox-container that will be returned
	metaDataArea = container.NewVBox(accordion, widget.NewLabel(""), widget.NewSeparator())

	return metaDataArea, err
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
				metaDataItemsAsCanvasObject = append(metaDataItemsAsCanvasObject,
					container.NewVBox(
						widget.NewLabel(label),
						sel,
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

		// each metadata item cell is 200×60
		content := container.New(
			layout.NewGridWrapLayout(fyne.NewSize(200, 60)),
			metaDataItemsAsCanvasObject...,
		)

		var card *widget.Card
		card = widget.NewCard(metaDataGroupPtr.MetaDataGroupName, "", content)
		metaDataGroupCards = append(metaDataGroupCards, card)
	}

	// top‐level grid: each card cell is 220×180
	return container.New(
		layout.NewGridWrapLayout(fyne.NewSize(220, 180)),
		metaDataGroupCards...,
	)
}
