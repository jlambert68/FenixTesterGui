package testCaseUI

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"sort"
)

type attributeStruct struct {
	attributeUuid                        string
	attributeName                        string
	attributeValue                       string
	attributeChangedValue                string
	attributeTypeName                    string
	entryRef                             *widget.Entry
	attributeIsChanged                   bool
	testInstructionElementMatureUuidUuid string
}

var attributesList []*attributeStruct

// Generate the TestCaseAttributes Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateTestCaseAttributesAreaForTestCase(testCaseUuid string, testInstructionElementMatureUuidUuid string) (testCaseAttributesArea fyne.CanvasObject, testInstructionAttributesAccordion *widget.Accordion, err error) {

	// If previous call to this method resulted in attributes then check if any of them were changed
	if len(attributesList) > 0 {
		for _, attribute := range attributesList {
			if attribute.attributeIsChanged == true {
				err = testCasesUiCanvasObject.saveChangedTestCaseAttributeInTestCase(testCaseUuid)
				break
			}
		}
	}

	if err != nil {
		return nil, nil, err
	}

	attributesList = testCasesUiCanvasObject.generateAttributeStringListData(testCaseUuid, testInstructionElementMatureUuidUuid)

	// Extract the current TestCase model
	testCaseModel, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]

	// Create Form-layout container to be used for attributes
	var attributesContainer *fyne.Container
	var attributesFormContainer *fyne.Container
	attributesContainer = container.New(layout.NewVBoxLayout())
	attributesFormContainer = container.New(layout.NewFormLayout())

	// Only add attributes if there are any, otherwise write simple text
	if len(attributesList) > 0 {
		var previousAttributeTypeName string
		var firstTime bool = true

		// Loop attributes and create label en entry field for each attribut
		for _, attributeItem := range attributesList {
			if attributeItem.attributeTypeName != previousAttributeTypeName {
				if firstTime == true {
					attributesContainer.Add(widget.NewLabel(attributeItem.attributeTypeName))
					firstTime = false
				} else {
					attributesContainer.Add(widget.NewLabel(attributeItem.attributeTypeName))
					attributesContainer.Add(attributesFormContainer)
					attributesFormContainer = container.New(layout.NewFormLayout())
				}
			}

			previousAttributeTypeName = attributeItem.attributeTypeName

			// Add the label for the Entry-widget
			attributesFormContainer.Add(widget.NewLabel(attributeItem.attributeName))

			// Add the Entry-widget
			newAttributeEntry := widget.NewEntry() // testCasesUiCanvasObject.NewAttributeEntry(attributeItem.attributeUuid)
			attributeItem.entryRef = newAttributeEntry
			newAttributeEntry.SetText(attributeItem.attributeValue)
			newAttributeEntry.OnChanged = func(newValue string) {
				// Find which attributes that we are dealing with
				var tempAttributeItem *attributeStruct
				for _, tempAttributeItem = range attributesList {
					if tempAttributeItem.entryRef == newAttributeEntry {
						break
					}
				}
				if newValue != tempAttributeItem.attributeValue {
					tempAttributeItem.attributeIsChanged = true
					tempAttributeItem.attributeChangedValue = newValue
				} else {
					tempAttributeItem.attributeIsChanged = false
					tempAttributeItem.attributeChangedValue = newValue
				}
			}
			attributesFormContainer.Add(newAttributeEntry)
		}

		// Handle last batch of batch Attributes
		attributesContainer.Add(attributesFormContainer)

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
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateAttributeStringListData(testCaseUuid string, testInstructionElementMatureUuid string) (attributesList []*attributeStruct) {

	// Clear variable
	attributesList = []*attributeStruct{}

	// Used when creating th UI for first time
	if testInstructionElementMatureUuid == "" {
		return
	}

	// Extract TestCase-model
	currentTestCaseModel, existsInMap := testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]

	if existsInMap == false {
		errorId := "50346c17-be7d-4929-b9f2-5367e464b0e7"
		err := errors.New(fmt.Sprintf("testcase-model with TestCaseUuid '%s' is missing map for TestCases [ErrorID: %s]", testCaseUuid, errorId))

		//TODO Send ERRORS over error-channel
		fmt.Println(err)

		return

	}

	// Extract the map for the TestInstructions Attributes
	matureTestInstruction, existInMap := currentTestCaseModel.MatureTestInstructionMap[testInstructionElementMatureUuid]

	if existInMap == false {
		errorId := "406439e8-1802-4a5f-b9ef-9024adf75ead"
		err := errors.New(fmt.Sprintf("testinstruction with uuid '%s' is missing in 'MatureTestInstructionMap' [ErrorID: %s]", testInstructionElementMatureUuid, errorId))

		//TODO Send ERRORS over error-channel
		fmt.Println(err.Error())

		return

	}

	// Loop over attributes and append to slice of attributes with 'Name' and 'value'
	for _, testInstructionAttribute := range matureTestInstruction.TestInstructionAttributesList {
		attributesList = append(attributesList, &attributeStruct{
			attributeUuid:                        testInstructionAttribute.BaseAttributeInformation.TestInstructionAttributeUuid,
			attributeName:                        testInstructionAttribute.BaseAttributeInformation.TestInstructionAttributeName,
			attributeValue:                       testInstructionAttribute.AttributeInformation.InputTextBoxProperty.TextBoxAttributeValue,
			attributeTypeName:                    testInstructionAttribute.AttributeInformation.InputTextBoxProperty.TextBoxAttributeTypeName,
			testInstructionElementMatureUuidUuid: testInstructionElementMatureUuid,
		})
	}

	// Sort Attributes in Name-order, within each Type
	sort.SliceStable(attributesList, func(i, j int) bool {
		if attributesList[i].attributeTypeName != attributesList[j].attributeTypeName {
			return attributesList[i].attributeTypeName < attributesList[j].attributeTypeName
		}

		return attributesList[i].attributeName < attributesList[j].attributeName
	})

	return attributesList
}

func (testCasesUiCanvasObject *TestCasesUiModelStruct) saveChangedTestCaseAttributeInTestCase(testCaseUuid string) (err error) {

	// Extract current TestCase
	testCase, existInMap := testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]
	if existInMap == false {

		errorId := "40fc730f-87d4-4c44-96ff-ab1003e40751"
		err := errors.New(fmt.Sprintf("testCase %s is missing in TestCases-map [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) //TODO Send error over error-channel
		return err
	}

	// Extract testInstructionElementMatureUuidUuid
	testInstructionElementMatureUuidUuid := attributesList[0].testInstructionElementMatureUuidUuid

	// Check if any attribute is changed
	if len(attributesList) > 0 {
		for _, attribute := range attributesList {
			if attribute.attributeIsChanged == true {
				// Attribute is changed so save it,

				// Extract TestInstruction
				tempMatureTestInstruction, existInMap := testCase.MatureTestInstructionMap[testInstructionElementMatureUuidUuid]
				if existInMap == false {
					errorId := "83b64181-3a02-4b98-8eba-d1fbad61dcd5"
					err := errors.New(fmt.Sprintf("mature testInstruction %s is missing in MatureTestInstructionMap [ErrorID: %s]", testInstructionElementMatureUuidUuid, errorId))

					fmt.Println(err) //TODO Send error over error-channel
					return err
				}

				// Extract  Attribute
				tempTestInstructionAttribute, existInMap := tempMatureTestInstruction.TestInstructionAttributesList[attribute.attributeUuid]
				if existInMap == false {
					errorId := "77e03442-7ccc-46c7-891e-0c5e0dd5bd1c"
					err := errors.New(fmt.Sprintf("testInstruction attribute %s is missing in MatureTestInstructionMap [ErrorID: %s]", attribute.attributeUuid, errorId))

					fmt.Println(err) //TODO Send error over error-channel
					return err
				}

				// Save changed value for Attribute
				tempTestInstructionAttribute.AttributeInformation.InputTextBoxProperty.TextBoxAttributeValue = attribute.attributeChangedValue

			}
		}
	}

	return err

}
