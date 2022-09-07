package testCaseUI

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type attributeStruct struct {
	attributeName  string
	attributeValue string
}

var attributesList []attributeStruct

// Generate the TestCaseAttributes Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateTestCaseAttributesAreaForTestCase(testCaseUuid string, testInstructionElementOriginalUuid string) (testCaseAttributesArea fyne.CanvasObject, testInstructionAttributesAccordion *widget.Accordion, err error) {

	testCasesUiCanvasObject.generateAttributeStringListData(testInstructionElementOriginalUuid)

	/*
		// Get current TestCase-UI-model
		_, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]

		if existsInMap == true {
			errorId := "c4110d4f-3dca-48bd-a8e4-57cb040fe079"
			err = errors.New(fmt.Sprintf("testcase-UI-model with sourceUuid '%s' allready exist in 'TestCasesUiModelMap' [ErrorID: %s]", testCaseUuid, errorId))

			fmt.Println(err.Error()) //TODO send to error channel

			return nil, err
		}

	*/

	// Extract the current TestCase model
	testCaseModel, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]

	// Create Form-layout container to be used for attributes
	attributesContainer := container.New(layout.NewFormLayout())

	// Only add attributes if there are any, otherwise write simple text
	if len(attributesList) > 0 {
		// Loop attributes and create label en entry field for each attribut
		for _, attributeItem := range attributesList {
			attributesContainer.Add(widget.NewLabel(attributeItem.attributeName))

			newAttributeEntry := widget.NewEntry()
			newAttributeEntry.SetText(attributeItem.attributeValue)
			attributesContainer.Add(newAttributeEntry)
		}
	} else {
		// No attributes so return simple label
		newLabel := widget.NewLabel("No Attributes for TestInstruction")

		// Create the AccordionItem
		testInstructionAttributesAccordionItem := widget.NewAccordionItem("TestInstruction Attributes", newLabel)

		if existsInMap == true {
			// Accordion exist so just replace the AccordionItem
			testCaseModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject.RemoveIndex(0)
			testCaseModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject.Append(testInstructionAttributesAccordionItem)

			// Open the Accordion
			testCaseModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject.Open(0)

			return attributesContainer, testCaseModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject, err

		} else {
			// Accordion doesn't exit so create it and add the AccordionItem
			testInstructionAttributesAccordion = widget.NewAccordion(testInstructionAttributesAccordionItem)

			attributesContainer = container.NewVBox(testInstructionAttributesAccordion)

			// Open the Accordion
			testInstructionAttributesAccordion.Open(0)

		}

	}

	// Save Accordion to be able to update with new attributes for other TestInstruction

	if existsInMap == true {

		// Create the AccordionItem
		testInstructionAttributesAccordionItem := widget.NewAccordionItem("TestInstruction Attributes", attributesContainer)

		// Check if Accordion already  exist
		if testCaseModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject == nil {
			// Accordion doesn't exit so create it and add the AccordionItem
			testInstructionAttributesAccordion = widget.NewAccordion(testInstructionAttributesAccordionItem)

			// save the Accordion to UI-model
			testCaseModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject = testInstructionAttributesAccordion

			// Open the Accordion
			testCaseModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject.Open(0)

		} else {
			// Accordion exist so,  add new AccordionItem to be the Accordion

			// Accordion already exists so remove the old AccordionItem and add the new one
			testCaseModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject.RemoveIndex(0)
			testCaseModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject.Append(testInstructionAttributesAccordionItem)

		}

		// Open the Accordion and refresh
		testCaseModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject.Close(0)
		testCaseModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject.Open(0)
		testCaseModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject.Refresh()

		attributesContainer = container.NewVBox(testCaseModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject)

	}

	return attributesContainer, testInstructionAttributesAccordion, err
}

// Generate structure for 'binding.StringList' regarding Attribute values
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateAttributeStringListData(testInstructionElementOriginalUuid string) {

	// Clear variable
	attributesList = []attributeStruct{}

	// Used when creating th UI for first time
	if testInstructionElementOriginalUuid == "" {
		return
	}

	immatureTestInstructionAttributesMap, existInMap := testCasesUiCanvasObject.TestCasesModelReference.ImmatureTestInstructionAttributesMap[testInstructionElementOriginalUuid]
	//TODO change to read from TestCase instead

	if existInMap == false {
		errorId := "406439e8-1802-4a5f-b9ef-9024adf75ead"
		err := errors.New(fmt.Sprintf("testinstruction with uuid '%s' is missing in 'AvailableImmatureTestInstructionsMap' [ErrorID: %s]", testInstructionElementOriginalUuid, errorId))

		fmt.Println(err.Error()) //TODO send to error channel

		return

	}

	// Loop over attributes and append to slice of attributes with 'Name' and 'value'
	for _, testInstructionAttribute := range immatureTestInstructionAttributesMap {
		attributesList = append(attributesList, attributeStruct{
			attributeName:  testInstructionAttribute.TestInstructionAttributeName,
			attributeValue: testInstructionAttribute.TestInstructionAttributeName,
		})
	}

}
