package testCaseUI

import (
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"sort"
)

// Generate the TestCaseAttributes Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateTestCaseAttributesAreaForTestCase(testCaseUuid string, testInstructionElementMatureUuid string) (testCaseAttributesArea fyne.CanvasObject, testInstructionAttributesAccordion *widget.Accordion, err error) {

	// Extract the current TestCase UI model
	testCase_Model, existsInMap := testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]
	if existsInMap == false {
		errorId := "07f8c5db-5a2a-4f1a-87ca-0c2e11f747a2"
		err := errors.New(fmt.Sprintf("testcase-model with TestCaseUuid '%s' is missing map for TestCases [ErrorID: %s]", testCaseUuid, errorId))

		//TODO Send ERRORS over error-channel
		fmt.Println(err)

		return nil, nil, err

	}

	//testCase_Model.AttributesList = tempAttributesList
	//*attributesList = testCase_Model.AttributesList

	// If previous call to this method resulted in attributes then check if any of them were changed
	if testCase_Model.AttributesList != nil &&
		len(*testCase_Model.AttributesList) > 0 {
		for _, attribute := range *testCase_Model.AttributesList {
			if attribute.AttributeIsChanged == true {
				err = testCasesUiCanvasObject.TestCasesModelReference.SaveChangedTestCaseAttributeInTestCase(testCaseUuid)
				break
			}
		}
	}
	// Generate Data to be used for  Attributes
	attributesList, err := testCasesUiCanvasObject.generateAttributeStringListData(testCaseUuid, testInstructionElementMatureUuid)
	if err != nil {
		return nil, nil, err
	}

	if err != nil {
		return nil, nil, err
	}

	// Extract the current TestCase UI model
	testCaseUIModel, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]

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
			if attributeItem.AttributeTypeName != previousAttributeTypeName {
				if firstTime == true {
					attributesContainer.Add(widget.NewLabel(attributeItem.AttributeTypeName))
					firstTime = false
				} else {
					attributesContainer.Add(widget.NewLabel(attributeItem.AttributeTypeName))
					attributesContainer.Add(attributesFormContainer)
					attributesFormContainer = container.New(layout.NewFormLayout())
				}
			}

			previousAttributeTypeName = attributeItem.AttributeTypeName

			// Add the label for the Entry-widget
			attributesFormContainer.Add(widget.NewLabel(attributeItem.AttributeName))

			// Add the Entry-widget
			newAttributeEntry := widget.NewEntry() // testCasesUiCanvasObject.NewAttributeEntry(attributeItem.attributeUuid)
			attributeItem.EntryRef = newAttributeEntry
			newAttributeEntry.SetText(attributeItem.AttributeValue)
			newAttributeEntry.OnChanged = func(newValue string) {
				// Find which attributes that we are dealing with
				var tempAttributeItem *testCaseModel.AttributeStruct

				for _, tempAttributeItem = range attributesList {
					if tempAttributeItem.EntryRef == newAttributeEntry {
						break
					}
				}
				if newValue != tempAttributeItem.AttributeValue {
					tempAttributeItem.AttributeIsChanged = true
					tempAttributeItem.AttributeChangedValue = newValue
				} else {
					tempAttributeItem.AttributeIsChanged = false
					tempAttributeItem.AttributeChangedValue = newValue
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
			testCaseUIModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject.RemoveIndex(0)
			testCaseUIModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject.Append(testInstructionAttributesAccordionItem)

			// Open the Accordion
			testCaseUIModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject.Open(0)

			return attributesContainer, testCaseUIModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject, err

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
		if testCaseUIModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject == nil {
			// Accordion doesn't exit so create it and add the AccordionItem
			testInstructionAttributesAccordion = widget.NewAccordion(testInstructionAttributesAccordionItem)

			// save the Accordion to UI-model
			testCaseUIModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject = testInstructionAttributesAccordion

			// Open the Accordion
			testCaseUIModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject.Open(0)

		} else {
			// Accordion exist so,  add new AccordionItem to be the Accordion

			// Accordion already exists so remove the old AccordionItem and add the new one
			testCaseUIModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject.RemoveIndex(0)
			testCaseUIModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject.Append(testInstructionAttributesAccordionItem)

		}

		// Open the Accordion and refresh
		testCaseUIModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject.Close(0)
		testCaseUIModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject.Open(0)
		testCaseUIModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject.Refresh()

		attributesContainer = container.NewVBox(testCaseUIModel.currentTestCaseGraphicalStructure.currentTestCaseTestInstructionAttributesAccordionObject)

	}

	return attributesContainer, testInstructionAttributesAccordion, err
}

// Generate structure for 'binding.StringList' regarding Attribute values
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateAttributeStringListData(testCaseUuid string, testInstructionElementMatureUuid string) (attributesListRef testCaseModel.AttributeStructSliceReference, err error) {

	// Extract TestCase-model
	currentTestCaseModel, existsInMap := testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid]

	if existsInMap == false {
		errorId := "50346c17-be7d-4929-b9f2-5367e464b0e7"
		err := errors.New(fmt.Sprintf("testcase-model with TestCaseUuid '%s' is missing map for TestCases [ErrorID: %s]", testCaseUuid, errorId))

		//TODO Send ERRORS over error-channel
		fmt.Println(err)

		return nil, err

	}

	// Used when creating th UI for first time

	if testInstructionElementMatureUuid == "" {

		// Clear variable
		attributesList := testCaseModel.AttributeStructSliceReference{}

		// Save AttributesList-reference back to TestCase
		currentTestCaseModel.AttributesList = &attributesList
		testCasesUiCanvasObject.TestCasesModelReference.TestCases[testCaseUuid] = currentTestCaseModel

		return attributesList, err
	}

	// Clear content of AttributesReference
	attributesList := testCaseModel.AttributeStructSliceReference{}

	// Extract the map for the TestInstructions Attributes
	matureTestInstruction, existInMap := currentTestCaseModel.MatureTestInstructionMap[testInstructionElementMatureUuid]

	//If the matureTestInstruction doesn't exist in map then it hasn't any attributes. So just return empty attributesLis
	if existInMap == false {

		*currentTestCaseModel.AttributesList = attributesList

		return attributesList, err

	}

	// Loop over attributes and append to slice of attributes with 'Name' and 'value'
	for _, testInstructionAttribute := range matureTestInstruction.TestInstructionAttributesList {
		attributesList = append(attributesList, &testCaseModel.AttributeStruct{
			AttributeUuid:                        testInstructionAttribute.BaseAttributeInformation.TestInstructionAttributeUuid,
			AttributeName:                        testInstructionAttribute.BaseAttributeInformation.TestInstructionAttributeName,
			AttributeValue:                       testInstructionAttribute.AttributeInformation.InputTextBoxProperty.TextBoxAttributeValue,
			AttributeTypeName:                    testInstructionAttribute.AttributeInformation.InputTextBoxProperty.TextBoxAttributeTypeName,
			TestInstructionElementMatureUuidUuid: testInstructionElementMatureUuid,
		})
	}

	// Sort Attributes in Name-order, within each Type
	sort.SliceStable(attributesList, func(i, j int) bool {
		if attributesList[i].AttributeTypeName != attributesList[j].AttributeTypeName {
			return attributesList[i].AttributeTypeName < attributesList[j].AttributeTypeName
		}

		return attributesList[i].AttributeName < attributesList[j].AttributeName
	})

	*currentTestCaseModel.AttributesList = attributesList

	//attributesListRef = &attributesList

	return attributesList, err
}
