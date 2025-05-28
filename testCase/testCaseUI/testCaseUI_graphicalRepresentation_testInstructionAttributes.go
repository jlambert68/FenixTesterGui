package testCaseUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	testInstruction_SendTemplateToThisDomain_version_1_0 "github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTemplateToThisDomain/version_1_0"
	testInstruction_SendTestDataToThisDomain_version_1_0 "github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTestDataToThisDomain/version_1_0"
	"image/color"
	"log"
	"regexp"
	"sort"
)

// Generate the TestCaseAttributes Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateTestCaseAttributesAreaForTestCase(
	testCaseUuid string,
	testInstructionElementMatureUuid string) (
	testCaseAttributesArea fyne.CanvasObject,
	testInstructionAttributesAccordion *widget.Accordion,
	err error) {

	var immatureTestInstructionUuid string

	// Extract the current TestCase UI model
	testCase_ModelPtr, existsInMap := testCasesUiCanvasObject.TestCasesModelReference.TestCasesMap[testCaseUuid]
	if existsInMap == false {
		errorId := "07f8c5db-5a2a-4f1a-87ca-0c2e11f747a2"
		err := errors.New(fmt.Sprintf("testcase-model with TestCaseUuid '%s' is missing map for TestCasesMap [ErrorID: %s]", testCaseUuid, errorId))

		//TODO Send ERRORS over error-channel
		fmt.Println(err)

		return nil, nil, err

	}

	//testCase_ModelPtr.AttributesList = tempAttributesList
	//*attributesList = testCase_ModelPtr.AttributesList

	// If previous call to this method resulted in attributes then check if any of them were changed
	if testCase_ModelPtr.AttributesList != nil &&
		len(*testCase_ModelPtr.AttributesList) > 0 {
		for _, attribute := range *testCase_ModelPtr.AttributesList {
			if attribute.AttributeIsChanged == true {
				err = testCasesUiCanvasObject.TestCasesModelReference.SaveChangedTestCaseAttributeInTestCase(testCaseUuid)
				break
			}
		}
	}
	// Generate Data to be used for Attributes
	var attributesList testCaseModel.AttributeStructSliceReferenceType
	attributesList, err = testCasesUiCanvasObject.generateAttributeStringListData(testCaseUuid, testInstructionElementMatureUuid)
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
		//var firstTime bool = true

		// Get Immature TestInstruction Uuid
		var matureTestInstruction testCaseModel.MatureTestInstructionStruct
		matureTestInstruction, existsInMap = testCase_ModelPtr.MatureTestInstructionMap[testInstructionElementMatureUuid]
		if existsInMap == false {
			errorId := "93475afd-c095-4fd6-a277-36cf7d4e703a"
			err = errors.New(fmt.Sprintf("testcase-model with TestCaseUuid '%s' is missing "+
				"Mature TestInstruction '%s' [ErrorID: %s]",
				testCaseUuid,
				testInstructionElementMatureUuid,
				errorId))

			//TODO Send ERRORS over error-channel
			fmt.Println(err)

			return nil, nil, err
		}
		immatureTestInstructionUuid = matureTestInstruction.BasicTestInstructionInformation_NonEditableInformation.
			TestInstructionOriginalUuid

		// Loop attributes and create label en entry field for each attribute
		for attributeItemCounter, attributeItem := range attributesList {

			// First attribute-data or a new AttributeType is presented
			if attributeItemCounter == 0 || attributeItem.AttributeTypeName != previousAttributeTypeName {

				// Add  previous FormContainer and aa separator
				if attributeItemCounter != 0 || attributeItem.AttributeTypeName != previousAttributeTypeName {
					attributesContainer.Add(attributesFormContainer)
					attributesContainer.Add(widget.NewSeparator())

				}

				// Generate a new FormContainer
				attributesFormContainer = container.New(layout.NewFormLayout())

				//attributesContainer.Add(widget.NewLabel(attributeItem.AttributeTypeName))
				attributesFormContainer.Add(widget.NewLabel(attributeItem.AttributeTypeName + "::"))
				attributesFormContainer.Add(layout.NewSpacer()) //(layout.NewSpacer())

				// Generate a new FormContainer
				//attributesFormContainer = container.New(layout.NewFormLayout())
			}

			// Generate and add an 'attribute row' to be used in attributes
			testCasesUiCanvasObject.generateAttributeRow(
				testCaseUuid,
				attributeItem,
				&attributesList,
				attributesFormContainer,
				testCase_ModelPtr,
				testInstructionElementMatureUuid,
				immatureTestInstructionUuid)

			//attributesContainer.Add(attributesFormContainer)

			previousAttributeTypeName = attributeItem.AttributeTypeName

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

		attributesContainer = container.NewVBox(testCaseUIModel.currentTestCaseGraphicalStructure.
			currentTestCaseTestInstructionAttributesAccordionObject, widget.NewLabel(""), widget.NewSeparator())

	}

	// Create a container for the Accordion
	accordionContainer := container.NewVBox(testInstructionAttributesAccordion)

	// Wrap the container in a scrollContainer
	accordionScrollContainer := container.NewScroll(accordionContainer)

	return accordionScrollContainer, testInstructionAttributesAccordion, err
}

// Generate and add an 'attribute row' to be used in attributes
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateAttributeRow(
	currentTestCaseUuid string,
	attributeItem *testCaseModel.AttributeStruct,
	attributesList *testCaseModel.AttributeStructSliceReferenceType,
	attributesFormContainer *fyne.Container,
	currentTestCase *testCaseModel.TestCaseModelStruct,
	testInstructionElementMatureUuid string,
	immatureTestInstructionUuid string) {

	// Add the label for the Attribute
	attributesFormContainer.Add(widget.NewLabel(attributeItem.AttributeName))

	// Depending on attribute chose correct UI-component
	switch attributeItem.AttributeType {
	case fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionAttributeTypeEnum_TEXTBOX:

		// Add the Entry-widget (TextBox)
		newAttributeEntry := widget.NewEntry()
		attributeItem.EntryRef = newAttributeEntry

		if attributeItem.AttributeTextBoxProperty.TextBoxEditable == true {
			newAttributeEntry.Enable()
		} else {
			newAttributeEntry.Disable()
		}
		newAttributeEntry.SetText(attributeItem.AttributeValue)
		newAttributeEntry.OnChanged = func(newValue string) {
			// Find which attributes that we are dealing with
			var tempAttributeItem *testCaseModel.AttributeStruct

			for _, tempAttributeItem = range *attributesList {
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

			// Set Warning box that value is not the correct one
			if attributeItem.CompileRegEx.MatchString(newValue) == false {
				attributeItem.AttributeValueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
				attributeItem.AttributeValueIsValid = false
			} else {
				attributeItem.AttributeValueIsValidWarningBox.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
				attributeItem.AttributeValueIsValid = true
			}

			// Update the attributesFormContainer
			attributesFormContainer.Refresh()

		}

		// Create The Attribute HBox-container
		//var attributeHboxContainer *fyne.Container
		//attributeHboxContainer = container.New(layout.NewHBoxLayout())

		// Create the 'AttributeValueIsValidWarningBox'
		var colorToUse color.NRGBA

		// Set Warning box that value is not the correct one
		if attributeItem.CompileRegEx.MatchString(attributeItem.AttributeValue) == false {
			colorToUse = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
			attributeItem.AttributeValueIsValid = false
		} else {
			colorToUse = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
			attributeItem.AttributeValueIsValid = true
		}

		attributeItem.AttributeValueIsValidWarningBox = canvas.NewRectangle(colorToUse)

		// Create a custom TextBoxEntry to use
		var customTextBoxEntry *customAttributeEntryWidget
		customTextBoxEntry = newCustomAttributeEntryWidget(newAttributeEntry, attributeItem.AttributeValueIsValidWarningBox)

		// Add the attribute container to the 'attributesFormContainer'
		attributesFormContainer.Add(customTextBoxEntry)

	case fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionAttributeTypeEnum_COMBOBOX:
		// Add the Select-widget (ComboBox)

		var err error

		// Get the available options from centralized stored list
		var optionsList []string

		var executionDomainsThatCanReceiveDirectTargetedTestInstructions map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
			ExecutionDomainsThatCanReceiveDirectTargetedTestInstructionsMessage
		executionDomainsThatCanReceiveDirectTargetedTestInstructions = *sharedCode.ExecutionDomainsThatCanReceiveDirectTargetedTestInstructionsMapPtr
		for _, tempExecutionDomain := range executionDomainsThatCanReceiveDirectTargetedTestInstructions {
			optionsList = append(optionsList, tempExecutionDomain.GetNameUsedInGui())
		}

		// SPECIAL
		// When the attribute is the ComboBox with Template, then change in lite with Templates if Template is in use or not
		if attributeItem.AttributeUuid == string(testInstruction_SendTemplateToThisDomain_version_1_0.
			TestInstructionAttributeUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateComboBox) {
			// Loop all Template to build the name used in the ComboBox

			optionsList = nil
			var templateName string
			var foundOldSelectedName bool

			for _, templateGitHubFile := range currentTestCase.ImportedTemplateFilesFromGitHub {

				templateName =
					fmt.Sprintf("%s [%s]",
						templateGitHubFile.Name,
						templateGitHubFile.FileHash[0:8])

				// Only Add the Template to the list if
				// it is not in use OR it is the chose Template for this ComboBox
				if templateGitHubFile.FileIsUsedInTestCase == false || attributeItem.AttributeValue == templateName {
					optionsList = append(optionsList, templateName)
				}

				if templateName == attributeItem.AttributeValue {
					foundOldSelectedName = true
				}

			}

			// Check if Old Selected was remove, if so clear attributeItem.AttributeValue
			if foundOldSelectedName == false {
				attributeItem.AttributeValue = ""
			}

		} else {

			// If there are values for the Combobox then use that
			if len(attributeItem.AttributeComboBoxProperty.GetComboBoxAllowedValues()) > 0 {
				optionsList = attributeItem.AttributeComboBoxProperty.GetComboBoxAllowedValues()
			}
		}

		// Create the Select
		newAttributeSelect := widget.NewSelect(optionsList, nil)
		attributeItem.SelectRef = newAttributeSelect

		// Set previous selected value if exist
		newAttributeSelect.SetSelected(attributeItem.AttributeValue)
		newAttributeSelect.OnChanged = func(newValue string) {

			// Generate Data to be used for Attributes
			attributesList = currentTestCase.AttributesList

			// Find which attributes that we are dealing with
			var tempAttributeItem *testCaseModel.AttributeStruct
			var foundAttribute bool

			for _, tempAttributeItem = range *attributesList {
				if tempAttributeItem.SelectRef == newAttributeSelect {
					foundAttribute = true
					break
				}
			}
			if foundAttribute == false {
				errorId := "bac28755-d9a8-4565-83b0-b01368ecbe9d"
				err = errors.New(fmt.Sprintf("Couldn't find Attribute in attribute list [ErrorID: %s]",
					errorId))

				//TODO Send ERRORS over error-channel

				// Hard exit
				log.Fatalln(err)

			}

			if newValue != tempAttributeItem.AttributeValue {
				tempAttributeItem.AttributeIsChanged = true
				tempAttributeItem.AttributeChangedValue = newValue

				// SPECIAL
				// When the attribute is the ComboBox that the user use to chose to which ExecutionDomain the TestData for the TestExecution
				// should be sent to. Then set DomainUuid and ExecutionDomainUuid in attribute fields for that. Used by the ExecutionServer
				switch tempAttributeItem.AttributeUuid {
				case string(testInstruction_SendTestDataToThisDomain_version_1_0.
					TestInstructionAttributeUUID_FenixSentToUsersDomain_SendTestDataToThisDomain_SendTestDataToThisExecutionDomain):

					var executionDomainsThatCanReceiveDirectTargetedTestInstructionsMap map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
						ExecutionDomainsThatCanReceiveDirectTargetedTestInstructionsMessage
					executionDomainsThatCanReceiveDirectTargetedTestInstructionsMap = *sharedCode.ExecutionDomainsThatCanReceiveDirectTargetedTestInstructionsMapPtr

					// Loop Attributes and find TextBox for DomainUuid and TextBox for ExecutionDomainUuid and populate from 'map'
					for _, tempAttribute := range *attributesList {

						// Check if this attribute used for DomainUuid, then set the value for DomainUuid
						if tempAttribute.AttributeUuid == string(testInstruction_SendTestDataToThisDomain_version_1_0.
							TestInstructionAttributeUUID_SendTestDataToThisDomain_SendTestDataToThisDomainTextBox) {

							// Set the value in the Attribute itself
							tempAttribute.AttributeValue = executionDomainsThatCanReceiveDirectTargetedTestInstructionsMap[newValue].
								GetDomainUuid()

							// Set the value in the UI-component for the Attribute
							tempAttribute.EntryRef.SetText(tempAttribute.AttributeValue)

							// Set that attribute was changed
							tempAttribute.AttributeIsChanged = true

						}

						// Check if this attribute used for ExecutionDomainUuid, then set the value for ExecutionDomainUuid
						if tempAttribute.AttributeUuid == string(testInstruction_SendTestDataToThisDomain_version_1_0.
							TestInstructionAttributeUUID_SendTestDataToThisDomain_SendTestDataToThisExecutionDomainTextBox) {

							tempAttribute.AttributeValue = executionDomainsThatCanReceiveDirectTargetedTestInstructionsMap[newValue].
								GetExecutionDomainUuid()

							// Set the value in the UI-component for the Attribute
							tempAttribute.EntryRef.SetText(tempAttribute.AttributeValue)

							// Set that attribute was changed
							tempAttribute.AttributeIsChanged = true
						}

					}

					// save back updated AttributeList
					currentTestCase.AttributesList = attributesList

					// Save back Updated TestCase
					testCasesUiCanvasObject.TestCasesModelReference.TestCasesMap[currentTestCaseUuid] = currentTestCase

				// SPECIAL
				// When the attribute is the ComboBox that the user use to chose to which ExecutionDomain the TestData for the TestExecution
				// should be sent to. Then set DomainUuid and ExecutionDomainUuid in attribute fields for that. Used by the ExecutionServer
				case string(testInstruction_SendTemplateToThisDomain_version_1_0.
					TestInstructionAttributeUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisExecutionDomainComboBox):

					var executionDomainsThatCanReceiveDirectTargetedTestInstructionsMap map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
						ExecutionDomainsThatCanReceiveDirectTargetedTestInstructionsMessage
					executionDomainsThatCanReceiveDirectTargetedTestInstructionsMap = *sharedCode.ExecutionDomainsThatCanReceiveDirectTargetedTestInstructionsMapPtr

					// Loop Attributes and find TextBox for DomainUuid and TextBox for ExecutionDomainUuid and populate from 'map'
					for _, tempAttribute := range *attributesList {

						// Check if this attribute used for DomainUuid, then set the value for DomainUuid
						if tempAttribute.AttributeUuid == string(testInstruction_SendTemplateToThisDomain_version_1_0.
							TestInstructionAttributeUUID_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisDomainTextBox) {

							// Set the value in the Attribute itself
							tempAttribute.AttributeValue = executionDomainsThatCanReceiveDirectTargetedTestInstructionsMap[newValue].
								GetDomainUuid()

							// Set the value in the UI-component for the Attribute
							tempAttribute.EntryRef.SetText(tempAttribute.AttributeValue)

							// Set that attribute was changed
							tempAttribute.AttributeIsChanged = true

						}

						// Check if this attribute used for ExecutionDomainUuid, then set the value for ExecutionDomainUuid
						if tempAttribute.AttributeUuid == string(testInstruction_SendTemplateToThisDomain_version_1_0.
							TestInstructionAttributeUUID_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisExecutionDomainTextBox) {

							tempAttribute.AttributeValue = executionDomainsThatCanReceiveDirectTargetedTestInstructionsMap[newValue].
								GetExecutionDomainUuid()

							// Set the value in the UI-component for the Attribute
							tempAttribute.EntryRef.SetText(tempAttribute.AttributeValue)

							// Set that attribute was changed
							tempAttribute.AttributeIsChanged = true
						}

					}

					// save back updated AttributeList
					currentTestCase.AttributesList = attributesList

					// Save back Updated TestCase
					testCasesUiCanvasObject.TestCasesModelReference.TestCasesMap[currentTestCaseUuid] = currentTestCase

					// SPECIAL
				// When the attribute is the ComboBox with Template, then change in lite with Templates if Template is in use or not
				case string(testInstruction_SendTemplateToThisDomain_version_1_0.
					TestInstructionAttributeUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateComboBox):

					var anyChange bool

					var oldSelectedTemplateName string
					var templateName string
					oldSelectedTemplateName = tempAttributeItem.AttributeValue

					// Set if  Template is not used
					for templateFileIndex, templateGitHubFile := range currentTestCase.ImportedTemplateFilesFromGitHub {

						templateName =
							fmt.Sprintf("%s [%s]",
								templateGitHubFile.Name,
								templateGitHubFile.FileHash[0:8])

						// Change old template to be available again
						if templateName == oldSelectedTemplateName {
							currentTestCase.ImportedTemplateFilesFromGitHub[templateFileIndex].FileIsUsedInTestCase = false

							// Loop Attributes and find TextBox for TemplateAsString to clear
							for _, tempAttribute := range *attributesList {

								// Check if this attribute used for TemplateAsString, then clear the value for TemplateAsString
								if tempAttribute.AttributeUuid == string(testInstruction_SendTemplateToThisDomain_version_1_0.
									TestInstructionAttributeUUID_FenixOwnedSendTemplateToThisDomain_FenixOwnedTemplateAsString) {

									// Set the value in the Attribute itself
									tempAttribute.AttributeValue = ""

									// Set the value in the UI-component for the Attribute
									tempAttribute.EntryRef.SetText(tempAttribute.AttributeValue)

									// Set that attribute was changed
									tempAttribute.AttributeIsChanged = true

									// Indicate that a change was done
									anyChange = true

									break
								}
							}

							break
						}
					}

					// Set if  Template is used
					for templateFileIndex, templateGitHubFile := range currentTestCase.ImportedTemplateFilesFromGitHub {

						templateName =
							fmt.Sprintf("%s [%s]",
								templateGitHubFile.Name,
								templateGitHubFile.FileHash[0:8])

						// Change old template to be available again
						if templateName == newValue {
							currentTestCase.ImportedTemplateFilesFromGitHub[templateFileIndex].FileIsUsedInTestCase = true

							// Loop Attributes and find TextBox for TemplateAsString to clear
							for _, tempAttribute := range *attributesList {

								// Check if this attribute used for TemplateAsString, then clear the value for TemplateAsString
								if tempAttribute.AttributeUuid == string(testInstruction_SendTemplateToThisDomain_version_1_0.
									TestInstructionAttributeUUID_FenixOwnedSendTemplateToThisDomain_FenixOwnedTemplateAsString) {

									// Set the value in the Attribute itself
									tempAttribute.AttributeValue = templateGitHubFile.FileContentAsString

									// Set the value in the UI-component for the Attribute
									tempAttribute.EntryRef.SetText(tempAttribute.AttributeValue)

									// Set that attribute was changed
									tempAttribute.AttributeIsChanged = true

									// Indicate that a change was done
									anyChange = true

									break
								}
							}

							break
						}
					}

					// Do a save if any change was done
					if anyChange == true {

						// save back updated AttributeList
						currentTestCase.AttributesList = attributesList

						// Save back Updated TestCase
						testCasesUiCanvasObject.TestCasesModelReference.TestCasesMap[currentTestCaseUuid] = currentTestCase
					}

				default:

				}

			} else {
				tempAttributeItem.AttributeIsChanged = false
				tempAttributeItem.AttributeChangedValue = newValue
			}

			// Set Warning box that value is not selected
			if attributeItem.CompileRegEx.MatchString(newValue) == false {
				attributeItem.AttributeValueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
				attributeItem.AttributeValueIsValid = false
			} else {
				attributeItem.AttributeValueIsValidWarningBox.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
				attributeItem.AttributeValueIsValid = true
			}

			// Update the attributesFormContainer
			attributesFormContainer.Refresh()

		}

		// Create the 'AttributeValueIsValidWarningBox'
		var colorToUse color.NRGBA

		// Set Warning box that value is not the correct one
		if attributeItem.CompileRegEx.MatchString(attributeItem.AttributeValue) == false {
			colorToUse = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
			attributeItem.AttributeValueIsValid = false
		} else {
			colorToUse = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
			attributeItem.AttributeValueIsValid = true
		}

		attributeItem.AttributeValueIsValidWarningBox = canvas.NewRectangle(colorToUse)

		// Create a custom SelectComboBox to use
		// Create a custom TextBoxEntry to use
		var customSelectComboBox *customAttributeSelectComboBox
		customSelectComboBox = newCustomAttributeSelectComboBoxWidget(newAttributeSelect, attributeItem.AttributeValueIsValidWarningBox)

		// Add the attribute HBox-container to the 'attributesFormContainer'
		attributesFormContainer.Add(customSelectComboBox)

	case fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionAttributeTypeEnum_RESPONSE_VARIABLE_COMBOBOX:
		// Add the Select-widget (ComboBox)

		// Get available values
		var allowedResponseVariablesTypeUuid []string
		var matureTestInstructionsWithCorrectResponseVariablesType []*testCaseModel.MatureTestInstructionWithCorrectResponseVariablesTypeStruct
		var err error

		// Get Immature TestInstruction
		var immatureTestInstruction *fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionMessage
		var existInMap bool
		immatureTestInstruction, existInMap = testCasesUiCanvasObject.TestCasesModelReference.
			AvailableImmatureTestInstructionsMap[immatureTestInstructionUuid]

		if existInMap == false {
			if err != nil {
				errorId := "94e939c2-aba6-428e-84f0-b012aa19f9bf"
				err = errors.New(fmt.Sprintf("Couldn't find Immature testInstruction '%s'. "+
					"Error message: '%s' [ErrorID: %s]",
					immatureTestInstructionUuid,
					errorId))

				//TODO Send ERRORS over error-channel
				fmt.Println(err)

				return
			}
		}

		// Generate 'allowedResponseVariablesTypeUuid' TODO secure that no duplicates is found
		for _, responseVariable := range immatureTestInstruction.GetResponseVariablesMapStructure().ResponseVariablesMap {
			allowedResponseVariablesTypeUuid = append(allowedResponseVariablesTypeUuid, responseVariable.ResponseVariableTypeUuid)
		}

		// Call to traverse the elements model to find all TestInstructions that has a matching response variable as
		// 'this' TestInstructions attribute need
		err = testCasesUiCanvasObject.recursiveTraverseUpwardsTestInstructionContainerElementsForResponseVariablesThatMatch(
			currentTestCase,
			testInstructionElementMatureUuid,
			&allowedResponseVariablesTypeUuid,
			&matureTestInstructionsWithCorrectResponseVariablesType,
			true,
			"",
			"")

		if err != nil {
			errorId := "c538c237-78ee-4fe8-9ef5-03f8761e2f90"
			err = errors.New(fmt.Sprintf("There was some problem when travesering the TestCase-model. "+
				"Error message: '%s' [ErrorID: %s]",
				err.Error(),
				errorId))

			//TODO Send ERRORS over error-channel
			fmt.Println(err)

			return
		}

		// Generate ComboBox options names
		var matureTestInstructionComboBoxOptionsName string
		var optionsList []string
		var optionsListLength int
		optionsListLength = len(matureTestInstructionsWithCorrectResponseVariablesType)
		for optionsIndex := optionsListLength - 1; optionsIndex >= 0; optionsIndex-- {

			// Extract
			matureTestInstructionWithCorrectResponseVariablesType := matureTestInstructionsWithCorrectResponseVariablesType[optionsIndex]

			// Create names to be used as options in ComboBox
			matureTestInstructionComboBoxOptionsName = fmt.Sprintf("%s [%s]",
				matureTestInstructionWithCorrectResponseVariablesType.MatureTestInstructionNameWithCorrectResponseVariablesType,
				sharedCode.GenerateShortUuidFromFullUuid(matureTestInstructionWithCorrectResponseVariablesType.
					MatureTestInstructionUuidWithCorrectResponseVariablesType))

			// Add the options name back to the object
			matureTestInstructionsWithCorrectResponseVariablesType[optionsIndex].
				MatureTestInstructionComboBoxOptionsName = matureTestInstructionComboBoxOptionsName

			// Add the object back to the slice
			matureTestInstructionsWithCorrectResponseVariablesType[optionsIndex] = matureTestInstructionWithCorrectResponseVariablesType

			// Add to the options array to be able to feed the Combobox
			optionsList = append(optionsList, matureTestInstructionComboBoxOptionsName)

		}

		// Save the 'matureTestInstructionsWithCorrectResponseVariablesType' back to the Attribute in TestCase-model
		attributeItem.AttributeResponseVariableComboBoxProperty.
			MatureTestInstructionsWithCorrectResponseVariablesType = &matureTestInstructionsWithCorrectResponseVariablesType

		newAttributeSelect := widget.NewSelect(optionsList, nil)
		attributeItem.SelectRef = newAttributeSelect

		// Set previous selected value if exist
		newAttributeSelect.SetSelected(attributeItem.AttributeValue)
		newAttributeSelect.OnChanged = func(newValue string) {
			// Find which attributes that we are dealing with
			var tempAttributeItem *testCaseModel.AttributeStruct
			var foundAttribute bool

			for _, tempAttributeItem = range *attributesList {
				if tempAttributeItem.SelectRef == newAttributeSelect {
					foundAttribute = true
					break
				}
			}
			if foundAttribute == false {
				errorId := "bac28755-d9a8-4565-83b0-b01368ecbe9d"
				err = errors.New(fmt.Sprintf("Couldn't find Attribute in attribute list [ErrorID: %s]",
					errorId))

				//TODO Send ERRORS over error-channel

				// Hard exit
				log.Fatalln(err)

			}

			// Loop Available TestInstructions with correct ResponseVariable and find the chosen
			var foundChosenTestInstruction bool
			for _, tempChosenTestInstruction := range *tempAttributeItem.AttributeResponseVariableComboBoxProperty.
				MatureTestInstructionsWithCorrectResponseVariablesType {

				foundChosenTestInstruction = true

				// Stop when we find the chosen TestInstruction
				if tempChosenTestInstruction.MatureTestInstructionComboBoxOptionsName == newValue {
					// Set Chosen TestInstruction
					tempAttributeItem.AttributeResponseVariableComboBoxProperty.AttributeResponseVariableComboBoxProperty.
						ChosenResponseVariableTypeUuid = tempChosenTestInstruction.
						MatureTestInstructionUuidWithCorrectResponseVariablesType
					tempAttributeItem.AttributeResponseVariableComboBoxProperty.AttributeResponseVariableComboBoxProperty.
						ChosenResponseVariableTypeName = tempChosenTestInstruction.
						MatureTestInstructionNameWithCorrectResponseVariablesType
					tempAttributeItem.AttributeResponseVariableComboBoxProperty.AttributeResponseVariableComboBoxProperty.
						ComboBoxAttributeValueAsString = tempChosenTestInstruction.
						MatureTestInstructionComboBoxOptionsName

					break
				}
			}

			if foundChosenTestInstruction == false {
				errorId := "baa28755-d9a8-4565-83b0-b02368ecbe9d"
				err = errors.New(fmt.Sprintf("Couldn't find Attribute in attribute list [ErrorID: %s]",
					errorId))

				//TODO Send ERRORS over error-channel

				// Hard exit
				log.Fatalln(err)
			}

			if newValue != tempAttributeItem.AttributeValue {
				tempAttributeItem.AttributeIsChanged = true
				tempAttributeItem.AttributeChangedValue = newValue
			} else {
				tempAttributeItem.AttributeIsChanged = false
				tempAttributeItem.AttributeChangedValue = newValue
			}

			// Set Warning box that value is not selected
			if attributeItem.CompileRegEx.MatchString(newValue) == false {
				attributeItem.AttributeValueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
				attributeItem.AttributeValueIsValid = false
			} else {
				attributeItem.AttributeValueIsValidWarningBox.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
				attributeItem.AttributeValueIsValid = true
			}

			// Update the attributesFormContainer
			attributesFormContainer.Refresh()

		}

		// Create the 'AttributeValueIsValidWarningBox'
		var colorToUse color.NRGBA

		// Set Warning box that value is not the correct one
		if attributeItem.CompileRegEx.MatchString(attributeItem.AttributeValue) == false {
			colorToUse = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
			attributeItem.AttributeValueIsValid = false
		} else {
			colorToUse = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
			attributeItem.AttributeValueIsValid = true
		}

		attributeItem.AttributeValueIsValidWarningBox = canvas.NewRectangle(colorToUse)

		// Create a custom SelectComboBox to use
		// Create a custom TextBoxEntry to use
		var customSelectComboBox *customAttributeSelectComboBox
		customSelectComboBox = newCustomAttributeSelectComboBoxWidget(newAttributeSelect, attributeItem.AttributeValueIsValidWarningBox)

		// Add the attribute HBox-container to the 'attributesFormContainer'
		attributesFormContainer.Add(customSelectComboBox)

	default:
		errorId := "c526f868-6c5c-4c4f-b64b-2d2c77272319"
		err := errors.New(fmt.Sprintf("Unhandled Attribute-type, %s [ErrorID: %s]",
			fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum_name[int32(
				attributeItem.AttributeType)],
			errorId))

		// Exit
		log.Fatalln(err)

	}

}

// Generate structure for 'binding.StringList' regarding Attribute values
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateAttributeStringListData(
	testCaseUuid string,
	testInstructionElementMatureUuid string) (attributesListRef testCaseModel.AttributeStructSliceReferenceType, err error) {

	// Extract TestCase-model
	currentTestCaseModel, existsInMap := testCasesUiCanvasObject.TestCasesModelReference.TestCasesMap[testCaseUuid]

	if existsInMap == false {
		errorId := "50346c17-be7d-4929-b9f2-5367e464b0e7"
		err := errors.New(fmt.Sprintf("testcase-model with TestCaseUuid '%s' is missing map for TestCasesMap [ErrorID: %s]", testCaseUuid, errorId))

		//TODO Send ERRORS over error-channel
		fmt.Println(err)

		return nil, err

	}

	// Used when creating the UI for first time

	if testInstructionElementMatureUuid == "" {

		// Clear variable
		attributesList := testCaseModel.AttributeStructSliceReferenceType{}

		// Save AttributesList-reference back to TestCase
		currentTestCaseModel.AttributesList = &attributesList
		testCasesUiCanvasObject.TestCasesModelReference.TestCasesMap[testCaseUuid] = currentTestCaseModel

		return attributesList, err
	}

	// Clear content of AttributesReference
	attributesList := testCaseModel.AttributeStructSliceReferenceType{}

	// Extract the map for the TestInstructions Attributes
	matureTestInstruction, existInMap := currentTestCaseModel.MatureTestInstructionMap[testInstructionElementMatureUuid]

	//If the matureTestInstruction doesn't exist in map then it hasn't any attributes. So just return empty attributesLis
	if existInMap == false {

		*currentTestCaseModel.AttributesList = attributesList

		return attributesList, err

	}

	// Loop over attributes and append to slice of attributes with 'Name' and 'value'
	for _, testInstructionAttribute := range matureTestInstruction.TestInstructionAttributesList {

		// Compile the RegEx for Attribute
		var compileRegEx *regexp.Regexp

		// Depending on Attribute Type pick correct attribute data
		switch testInstructionAttribute.BaseAttributeInformation.TestInstructionAttributeType {

		case fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionAttributeTypeEnum_TEXTBOX:

			compileRegEx = regexp.MustCompile(testInstructionAttribute.AttributeInformation.InputTextBoxProperty.TextBoxInputMask)

			attributesList = append(attributesList, &testCaseModel.AttributeStruct{
				AttributeUuid:                             testInstructionAttribute.BaseAttributeInformation.TestInstructionAttributeUuid,
				AttributeName:                             testInstructionAttribute.BaseAttributeInformation.TestInstructionAttributeName,
				AttributeValue:                            testInstructionAttribute.AttributeInformation.InputTextBoxProperty.TextBoxAttributeValue,
				AttributeChangedValue:                     "",
				AttributeTypeName:                         testInstructionAttribute.AttributeInformation.InputTextBoxProperty.TextBoxAttributeTypeName,
				AttributeType:                             testInstructionAttribute.BaseAttributeInformation.TestInstructionAttributeType,
				AttributeTextBoxProperty:                  testInstructionAttribute.AttributeInformation.GetInputTextBoxProperty(),
				AttributeComboBoxProperty:                 nil,
				AttributeResponseVariableComboBoxProperty: nil,
				EntryRef:                                  nil,
				SelectRef:                                 nil,
				AttributeIsChanged:                        false,
				TestInstructionElementMatureUuidUuid:      testInstructionElementMatureUuid,
				AttributeValueIsValidRegExAsString:        testInstructionAttribute.AttributeInformation.InputTextBoxProperty.TextBoxInputMask,
				CompileRegEx:                              compileRegEx,
				AttributeValueIsValid:                     false,
				AttributeValueIsValidWarningBox:           nil,
			})

		case fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionAttributeTypeEnum_COMBOBOX:

			compileRegEx = regexp.MustCompile(testInstructionAttribute.AttributeInformation.InputComboBoxProperty.ComboBoxInputMask)

			attributesList = append(attributesList, &testCaseModel.AttributeStruct{
				AttributeUuid:                             testInstructionAttribute.BaseAttributeInformation.TestInstructionAttributeUuid,
				AttributeName:                             testInstructionAttribute.BaseAttributeInformation.TestInstructionAttributeName,
				AttributeValue:                            testInstructionAttribute.AttributeInformation.InputComboBoxProperty.ComboBoxAttributeValue,
				AttributeChangedValue:                     "",
				AttributeTypeName:                         testInstructionAttribute.AttributeInformation.InputComboBoxProperty.ComboBoxAttributeTypeName,
				AttributeType:                             testInstructionAttribute.BaseAttributeInformation.TestInstructionAttributeType,
				AttributeTextBoxProperty:                  nil,
				AttributeComboBoxProperty:                 testInstructionAttribute.AttributeInformation.GetInputComboBoxProperty(),
				AttributeResponseVariableComboBoxProperty: nil,
				EntryRef:                                  nil,
				SelectRef:                                 nil,
				AttributeIsChanged:                        false,
				TestInstructionElementMatureUuidUuid:      testInstructionElementMatureUuid,
				AttributeValueIsValidRegExAsString:        testInstructionAttribute.AttributeInformation.InputComboBoxProperty.ComboBoxInputMask,
				CompileRegEx:                              compileRegEx,
				AttributeValueIsValid:                     false,
				AttributeValueIsValidWarningBox:           nil,
			})

		case fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionAttributeTypeEnum_RESPONSE_VARIABLE_COMBOBOX:
			var tempAttributeResponseVariableComboBoxProperty *testCaseModel.AttributeResponseVariableComboBoxPropertyStruct
			tempAttributeResponseVariableComboBoxProperty = &testCaseModel.AttributeResponseVariableComboBoxPropertyStruct{
				AttributeResponseVariableComboBoxProperty:              testInstructionAttribute.AttributeInformation.GetResponseVariableComboBoxProperty(),
				MatureTestInstructionsWithCorrectResponseVariablesType: nil,
			}

			compileRegEx = regexp.MustCompile(testInstructionAttribute.AttributeInformation.ResponseVariableComboBoxProperty.ComboBoxResponseVariableInputMask)

			attributesList = append(attributesList, &testCaseModel.AttributeStruct{
				AttributeUuid:                             testInstructionAttribute.BaseAttributeInformation.TestInstructionAttributeUuid,
				AttributeName:                             testInstructionAttribute.BaseAttributeInformation.TestInstructionAttributeName,
				AttributeValue:                            testInstructionAttribute.AttributeInformation.ResponseVariableComboBoxProperty.ComboBoxAttributeValueAsString,
				AttributeChangedValue:                     "",
				AttributeTypeName:                         testInstructionAttribute.AttributeInformation.ResponseVariableComboBoxProperty.ComboBoxAttributeTypeName,
				AttributeType:                             testInstructionAttribute.BaseAttributeInformation.TestInstructionAttributeType,
				AttributeTextBoxProperty:                  nil,
				AttributeComboBoxProperty:                 nil,
				AttributeResponseVariableComboBoxProperty: tempAttributeResponseVariableComboBoxProperty,
				EntryRef:                                  nil,
				SelectRef:                                 nil,
				AttributeIsChanged:                        false,
				TestInstructionElementMatureUuidUuid:      testInstructionElementMatureUuid,
				AttributeValueIsValidRegExAsString:        testInstructionAttribute.AttributeInformation.ResponseVariableComboBoxProperty.ComboBoxResponseVariableInputMask,
				CompileRegEx:                              compileRegEx,
				AttributeValueIsValid:                     false,
				AttributeValueIsValidWarningBox:           nil,
			})

		default:
			errorId := "c526f868-6c5c-4c4f-b64b-2d2c77272319"
			err = errors.New(fmt.Sprintf("Unhandled Attribute-type, %s [ErrorID: %s]",
				fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum_name[int32(
					testInstructionAttribute.BaseAttributeInformation.TestInstructionAttributeType)],
				errorId))

			//TODO Send ERRORS over error-channel
			fmt.Println(err)

			return nil, err

		}

	}

	// Sort Attributes in Name-order, within each Type
	sort.SliceStable(attributesList, func(i, j int) bool {

		/*
				var attributesListLower string
			var attributesListHigher string

			attributesListLower = attributesList[i].AttributeTypeName + attributesList[i].AttributeName
			attributesListHigher = attributesList[j].AttributeTypeName + attributesList[j].AttributeName

			return attributesListLower < attributesListHigher
		*/

		if attributesList[i].AttributeTypeName != attributesList[j].AttributeTypeName {
			return attributesList[i].AttributeTypeName < attributesList[j].AttributeTypeName
		}

		return attributesList[i].AttributeName < attributesList[j].AttributeName
	})

	*currentTestCaseModel.AttributesList = attributesList

	//attributesListRef = &attributesList

	return attributesList, err
}

// Traverse the element model to left and/or upward to the top element. This to be able to find all matching response variables.
// When the traverse logic comes to a TestInstructionContainer it will start to traverse down in this path, going first down and then right.
// This is use in UI for user to chose from. Then this information is used in runtime to get the value that a certain
// TestInstruction returned when it was executed
func (testCasesUiCanvasObject *TestCasesUiModelStruct) recursiveTraverseUpwardsTestInstructionContainerElementsForResponseVariablesThatMatch(
	currentTestCase *testCaseModel.TestCaseModelStruct,
	elementUuidToCheck string,
	allowedResponseVariablesTypeUuidPtr *[]string,
	testInstructionWithCorrectResponseVariablesTypePtr *[]*testCaseModel.MatureTestInstructionWithCorrectResponseVariablesTypeStruct,
	thisIsTheStartElement bool,
	previousProcessedElementUuid string,
	previousProcessedElementsParentUuid string) (
	err error) {

	// Extract current element
	currentElement, existInMap := currentTestCase.TestCaseModelMap[elementUuidToCheck]

	// If the element doesn't exit then there is something really wrong
	if existInMap == false {
		errorId := "b53ca94b-deaa-4136-88ec-5a434335cac0"
		err = errors.New(fmt.Sprintf("Element, '%s', couldn't be found in map  [ErrorID: %s]",
			elementUuidToCheck,
			errorId))

		//TODO Send ERRORS over error-channel
		fmt.Println(err)

		return err
	}

	// Check if this is the top Element of the full element tree
	if currentElement.MatureTestCaseModelElementMessage.
		ParentElementUuid == currentElement.MatureTestCaseModelElementMessage.MatureElementUuid {

		// This is the top element so start rewinding recursive loop
		return err

	}

	// Extract the parent element
	var parentElement testCaseModel.MatureTestCaseModelElementStruct
	parentElement, existInMap = currentTestCase.TestCaseModelMap[currentElement.MatureTestCaseModelElementMessage.ParentElementUuid]

	// If the element doesn't exit then there is something really wrong
	if existInMap == false {
		errorId := "fdfdcfec-551b-446e-91c9-7ce54825081d"
		err = errors.New(fmt.Sprintf("Parent Element, '%s', couldn't be found in map  [ErrorID: %s]",
			elementUuidToCheck,
			errorId))

		//TODO Send ERRORS over error-channel
		fmt.Println(err)

		return err
	}

	// Get parent TestInstructionContainer
	var parentTestInstructionContainer testCaseModel.MatureTestInstructionContainerStruct
	parentTestInstructionContainer, existInMap = currentTestCase.
		MatureTestInstructionContainerMap[parentElement.MatureTestCaseModelElementMessage.MatureElementUuid]

	// If the element doesn't exit then there is something really wrong
	if existInMap == false {
		errorId := "8260c569-da2f-4de9-95a6-29702dab165d"
		err = errors.New(fmt.Sprintf("Parent TestInstructionContainer, '%s', couldn't be found in map  [ErrorID: %s]",
			parentElement.MatureTestCaseModelElementMessage.MatureElementUuid,
			errorId))

		//TODO Send ERRORS over error-channel
		fmt.Println(err)

		return err
	}

	// Get parent TestInstructionContainer ExecutionType
	var parentTestInstructionContainerExecutionType fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionContainerExecutionTypeEnum
	parentTestInstructionContainerExecutionType = parentTestInstructionContainer.EditableTestInstructionContainerAttributes.
		GetTestInstructionContainerExecutionType()

	// Depending on ElementType do different task
	var testCaseModelElementType fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum
	testCaseModelElementType = currentElement.MatureTestCaseModelElementMessage.GetTestCaseModelElementType()
	switch testCaseModelElementType {

	// When TestInstruction then check Response Variables
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE:

		// Extract ResponseVariables from TestInstruction if this is not the start element
		if thisIsTheStartElement == false {

			err = testCasesUiCanvasObject.extractResponseVariablesFromTestInstruction(
				allowedResponseVariablesTypeUuidPtr,
				&currentElement,
				testInstructionWithCorrectResponseVariablesTypePtr)

			// When Error then exit recursively
			if err != nil {

				return err

			}

			/* Moved to separate function 'extractResponseVariablesFromTestInstruction'
			var tempResponseVariables map[string]*fenixGuiTestCaseBuilderServerGrpcApi.ResponseVariableMessage
			var tempImmatureTestInstructionMessage *fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionMessage

			tempImmatureTestInstructionMessage, existInMap = testCasesUiCanvasObject.TestCasesModelReference.
				AvailableImmatureTestInstructionsMap[currentElement.MatureTestCaseModelElementMessage.GetOriginalElementUuid()]

			if existInMap == false {
				errorId := "41aa60e3-f87e-4088-8433-0afa31c41bcc"
				err = errors.New(fmt.Sprintf("ImmatureTestInstruction, '%s', couldn't be found in "+
					"'AvailableImmatureTestInstructionsMap' + [ErrorID: %s]",
					currentElement.MatureTestCaseModelElementMessage.GetOriginalElementUuid(),
					errorId))

				//TODO Send ERRORS over error-channel
				fmt.Println(err)

				return err
			}

			tempResponseVariables = tempImmatureTestInstructionMessage.GetResponseVariablesMapStructure().
				ResponseVariablesMap

			// Loop TestInstructions ResponseVariables
			for _, tempResponseVariable := range tempResponseVariables {

				// Loop ResponseVariableTypes To match
				for _, allowedResponseVariableTypeUuid := range allowedResponseVariablesTypeUuid {

					// When TestInstructions Response Variable Type is the same as the ResponseVariableType then add to Slice
					if allowedResponseVariableTypeUuid == tempResponseVariable.GetResponseVariableTypeUuid() {
						// Add TestInstructions Original UUID Name
						var tempMatureTestInstructionWithCorrectResponseVariablesType *testCaseModel.
							MatureTestInstructionWithCorrectResponseVariablesTypeStruct

						tempMatureTestInstructionWithCorrectResponseVariablesType = &testCaseModel.
							MatureTestInstructionWithCorrectResponseVariablesTypeStruct{
							MatureTestInstructionUuidWithCorrectResponseVariablesType: currentElement.
								MatureTestCaseModelElementMessage.GetMatureElementUuid(),
							MatureTestInstructionNameWithCorrectResponseVariablesType: currentElement.
								MatureTestCaseModelElementMessage.GetOriginalElementName(),
							MatureTestInstructionComboBoxOptionsName: "",
						}
						*testInstructionWithCorrectResponseVariablesTypePtr = append(*testInstructionWithCorrectResponseVariablesTypePtr,
							tempMatureTestInstructionWithCorrectResponseVariablesType)
					}
				}

			}

			*/
		}

	// When TestInstructionContainer then do nothing due to that we travers horizontal or upwards in structure
	// TestInstructionContainer have a path downwards in tree
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE:

		// When Current Element is the same as Previous Parent Element, then we are att the
		// TestInstructionContainer that were just processed, then we need to decided how we should walk
		if currentElement.MatureTestCaseModelElementMessage.MatureElementUuid == previousProcessedElementUuid {

			// Check if elements are process in serial or in parallel
			switch parentTestInstructionContainerExecutionType {

			// Serial processed TestInstructionContainer -> Traverse up in this TestInstructionContainer
			case fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionContainerExecutionTypeEnum_SERIAL_PROCESSED:

				err = testCasesUiCanvasObject.recursiveTraverseUpwardsTestInstructionContainerElementsForResponseVariablesThatMatch(
					currentTestCase,
					currentElement.MatureTestCaseModelElementMessage.PreviousElementUuid,
					allowedResponseVariablesTypeUuidPtr,
					testInstructionWithCorrectResponseVariablesTypePtr,
					false,
					currentElement.MatureTestCaseModelElementMessage.MatureElementUuid,
					currentElement.MatureTestCaseModelElementMessage.ParentElementUuid,
				)

				return err

			// Parallel processed TestInstructionContainer -> Traverse from first to last element within this TestInstructionContainer
			case fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionContainerExecutionTypeEnum_PARALLELLED_PROCESSED:

				err = testCasesUiCanvasObject.recursiveTraverseUpwardsTestInstructionContainerElementsForResponseVariablesThatMatch(
					currentTestCase,
					parentElement.MatureTestCaseModelElementMessage.FirstChildElementUuid,
					allowedResponseVariablesTypeUuidPtr,
					testInstructionWithCorrectResponseVariablesTypePtr,
					false,
					currentElement.MatureTestCaseModelElementMessage.MatureElementUuid,
					currentElement.MatureTestCaseModelElementMessage.ParentElementUuid,
				)

				return err

			default:
				errorId := "aba4370f-2002-4c6b-8fbd-313cd4487438"
				err = errors.New(fmt.Sprintf("Unhandled 'TestCaseModelElementTypeEnum', '%d' [ErrorID: %s]",
					currentElement.MatureTestCaseModelElementMessage.GetTestCaseModelElementType(),
					errorId))

				// Hard Exit
				log.Fatalln(err)

			}

		}

		if currentElement.MatureTestCaseModelElementMessage.MatureElementUuid != previousProcessedElementsParentUuid {
			// Travers down this TestInstructionContainer to First child element
			err = testCasesUiCanvasObject.recursiveTraverseDownwardsTestInstructionContainerElementsForResponseVariablesThatMatch(
				currentTestCase,
				currentElement.MatureTestCaseModelElementMessage.FirstChildElementUuid,
				allowedResponseVariablesTypeUuidPtr,
				testInstructionWithCorrectResponseVariablesTypePtr)

			// When Error then exit recursively
			if err != nil {
				return err
			}
		}

	// When Bond the Traverse right if possible
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10oxo_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10ox_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10xo_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE:

	default:
		errorId := "07d8e9b3-fdd1-48d3-a1fe-d178aa92e3aa"
		err = errors.New(fmt.Sprintf("Unhandled 'TestCaseModelElementTypeEnum', '%d' [ErrorID: %s]",
			currentElement.MatureTestCaseModelElementMessage.GetTestCaseModelElementType(),
			errorId))

		// Hard Exit
		log.Fatalln(err)

		return err

	}

	// Different processing depending on if parent TestInstructionContainer is a serial or parallel executed one
	switch parentTestInstructionContainerExecutionType {

	// Serial processed TestInstructionContainer -> Traverse up in TestInstructionContainer
	case fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionContainerExecutionTypeEnum_SERIAL_PROCESSED:

		// If this is the start element and not the first element, then use previous element as next
		if currentElement.MatureTestCaseModelElementMessage.
			MatureElementUuid != currentElement.MatureTestCaseModelElementMessage.
			PreviousElementUuid && thisIsTheStartElement == true {

			// Traverse up/back in the Tree
			err = testCasesUiCanvasObject.recursiveTraverseUpwardsTestInstructionContainerElementsForResponseVariablesThatMatch(
				currentTestCase,
				currentElement.MatureTestCaseModelElementMessage.PreviousElementUuid,
				allowedResponseVariablesTypeUuidPtr,
				testInstructionWithCorrectResponseVariablesTypePtr,
				false,
				currentElement.MatureTestCaseModelElementMessage.MatureElementUuid,
				currentElement.MatureTestCaseModelElementMessage.ParentElementUuid,
			)

			return err
		}

		// Check if this is the first element in the TestInstructionContainer
		if currentElement.MatureTestCaseModelElementMessage.
			MatureElementUuid == currentElement.MatureTestCaseModelElementMessage.
			PreviousElementUuid {

			// Traverse up on level in the Tree
			err = testCasesUiCanvasObject.recursiveTraverseUpwardsTestInstructionContainerElementsForResponseVariablesThatMatch(
				currentTestCase,
				parentElement.MatureTestCaseModelElementMessage.MatureElementUuid,
				allowedResponseVariablesTypeUuidPtr,
				testInstructionWithCorrectResponseVariablesTypePtr,
				false,
				currentElement.MatureTestCaseModelElementMessage.MatureElementUuid,
				currentElement.MatureTestCaseModelElementMessage.ParentElementUuid,
			)

			return err

		} else {

			// This is not the first element in the TestInstructionContainer so
			// traverse one element upwards/back within TestInstructionContainer
			err = testCasesUiCanvasObject.recursiveTraverseUpwardsTestInstructionContainerElementsForResponseVariablesThatMatch(
				currentTestCase,
				currentElement.MatureTestCaseModelElementMessage.PreviousElementUuid,
				allowedResponseVariablesTypeUuidPtr,
				testInstructionWithCorrectResponseVariablesTypePtr,
				false,
				currentElement.MatureTestCaseModelElementMessage.MatureElementUuid,
				currentElement.MatureTestCaseModelElementMessage.ParentElementUuid,
			)

			return err
		}

	// Parallel processed TestInstructionContainer -> Traverse from first to last element within TestInstructionContainer
	case fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionContainerExecutionTypeEnum_PARALLELLED_PROCESSED:

		// Check if this is the last element in the TestInstructionContainer or this is the start element
		if currentElement.MatureTestCaseModelElementMessage.
			MatureElementUuid == currentElement.MatureTestCaseModelElementMessage.
			NextElementUuid || thisIsTheStartElement == true {

			// Traverse up on level in the Tree
			err = testCasesUiCanvasObject.recursiveTraverseUpwardsTestInstructionContainerElementsForResponseVariablesThatMatch(
				currentTestCase,
				currentElement.MatureTestCaseModelElementMessage.ParentElementUuid,
				allowedResponseVariablesTypeUuidPtr,
				testInstructionWithCorrectResponseVariablesTypePtr,
				false,
				currentElement.MatureTestCaseModelElementMessage.MatureElementUuid,
				currentElement.MatureTestCaseModelElementMessage.ParentElementUuid,
			)

			return err

		} else {

			// This is not the last element in the TestInstructionContainer so
			// traverse one element forward/right within TestInstructionContainer
			err = testCasesUiCanvasObject.recursiveTraverseUpwardsTestInstructionContainerElementsForResponseVariablesThatMatch(
				currentTestCase,
				currentElement.MatureTestCaseModelElementMessage.NextElementUuid,
				allowedResponseVariablesTypeUuidPtr,
				testInstructionWithCorrectResponseVariablesTypePtr,
				false,
				currentElement.MatureTestCaseModelElementMessage.MatureElementUuid,
				currentElement.MatureTestCaseModelElementMessage.ParentElementUuid,
			)

			return err
		}

	default:
		errorId := "cd4e981c-0336-443d-a086-ff832485baf6"
		err = errors.New(fmt.Sprintf("Unhandled 'TestCaseModelElementTypeEnum', '%d' [ErrorID: %s]",
			currentElement.MatureTestCaseModelElementMessage.GetTestCaseModelElementType(),
			errorId))

		// Hard Exit
		log.Fatalln(err)

	}

	return err
}

// When the traverse logic comes to a TestInstructionContainer it will start to traverse down in this path, going first down and then right.
// This is use in UI for user to chose from. Then this information is used in runtime to get the value that a certain
func (testCasesUiCanvasObject *TestCasesUiModelStruct) recursiveTraverseDownwardsTestInstructionContainerElementsForResponseVariablesThatMatch(
	currentTestCase *testCaseModel.TestCaseModelStruct,
	elementUuidToCheck string,
	allowedResponseVariablesTypeUuidPtr *[]string,
	testInstructionWithCorrectResponseVariablesTypePtr *[]*testCaseModel.MatureTestInstructionWithCorrectResponseVariablesTypeStruct) (
	err error) {

	// Extract current element
	currentElement, existInMap := currentTestCase.TestCaseModelMap[elementUuidToCheck]

	// If the element doesn't exit then there is something really wrong
	if existInMap == false {
		errorId := "b53ca94b-deaa-4136-88ec-5a434335cac0"
		err = errors.New(fmt.Sprintf("Element, '%s', couldn't be found in map  [ErrorID: %s]",
			elementUuidToCheck,
			errorId))

		//TODO Send ERRORS over error-channel
		fmt.Println(err)

		return err
	}

	// Depending on ElementType do different task
	var testCaseModelElementType fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum
	testCaseModelElementType = currentElement.MatureTestCaseModelElementMessage.GetTestCaseModelElementType()
	switch testCaseModelElementType {

	// When TestInstruction then check Response Variables and Traverse to Next element
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE:

		err = testCasesUiCanvasObject.extractResponseVariablesFromTestInstruction(
			allowedResponseVariablesTypeUuidPtr,
			&currentElement,
			testInstructionWithCorrectResponseVariablesTypePtr)

		// When Error then exit recursively
		if err != nil {
			return err
		}

		// Is this the last element, then exit
		if currentElement.MatureTestCaseModelElementMessage.
			MatureElementUuid == currentElement.MatureTestCaseModelElementMessage.NextElementUuid {

			return err
		}

		// Recursively traverse to next element
		err = testCasesUiCanvasObject.recursiveTraverseDownwardsTestInstructionContainerElementsForResponseVariablesThatMatch(
			currentTestCase,
			currentElement.MatureTestCaseModelElementMessage.NextElementUuid,
			allowedResponseVariablesTypeUuidPtr,
			testInstructionWithCorrectResponseVariablesTypePtr)

		return err

	// When TestInstructionContainer then first traverse Down to first child element.
	//When coming back then Traverse to Next element
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE:

		// Is this the last element, then exit
		if currentElement.MatureTestCaseModelElementMessage.
			MatureElementUuid == currentElement.MatureTestCaseModelElementMessage.NextElementUuid {

			return err
		}

		// Recursively traverse to first child element
		err = testCasesUiCanvasObject.recursiveTraverseDownwardsTestInstructionContainerElementsForResponseVariablesThatMatch(
			currentTestCase,
			currentElement.MatureTestCaseModelElementMessage.FirstChildElementUuid,
			allowedResponseVariablesTypeUuidPtr,
			testInstructionWithCorrectResponseVariablesTypePtr)

		// When Error then exit recursively
		if err != nil {
			return err
		}

		// Is this the Last element, then exit
		if currentElement.MatureTestCaseModelElementMessage.
			MatureElementUuid == currentElement.MatureTestCaseModelElementMessage.NextElementUuid {

			return err
		}

		// Recursively traverse to next element
		err = testCasesUiCanvasObject.recursiveTraverseDownwardsTestInstructionContainerElementsForResponseVariablesThatMatch(
			currentTestCase,
			currentElement.MatureTestCaseModelElementMessage.NextElementUuid,
			allowedResponseVariablesTypeUuidPtr,
			testInstructionWithCorrectResponseVariablesTypePtr)

		return err

		// When Bond the Traverse right if possible
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10oxo_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10ox_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10xo_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE:

		// Is this the Last element, then exit
		if currentElement.MatureTestCaseModelElementMessage.
			MatureElementUuid == currentElement.MatureTestCaseModelElementMessage.NextElementUuid {

			return err
		}

		// Recursively traverse to Next element
		err = testCasesUiCanvasObject.recursiveTraverseDownwardsTestInstructionContainerElementsForResponseVariablesThatMatch(
			currentTestCase,
			currentElement.MatureTestCaseModelElementMessage.NextElementUuid,
			allowedResponseVariablesTypeUuidPtr,
			testInstructionWithCorrectResponseVariablesTypePtr)

		return err

	default:
		errorId := "b212ab28-4701-46f8-a864-7f3506c63f75"
		err = errors.New(fmt.Sprintf("Unhandled 'TestCaseModelElementTypeEnum', '%d' [ErrorID: %s]",
			currentElement.MatureTestCaseModelElementMessage.GetTestCaseModelElementType(),
			errorId))

		// Hard Exit
		log.Fatalln(err)

		return err

	}

	return err

}

// // Extract ResponseVariables from TestInstruction, used within the two traverse functions
func (testCasesUiCanvasObject *TestCasesUiModelStruct) extractResponseVariablesFromTestInstruction(
	allowedResponseVariablesTypeUuidPtr *[]string,
	currentElement *testCaseModel.MatureTestCaseModelElementStruct,
	testInstructionWithCorrectResponseVariablesTypePtr *[]*testCaseModel.MatureTestInstructionWithCorrectResponseVariablesTypeStruct) (
	err error) {

	var allowedResponseVariablesTypeUuid []string
	allowedResponseVariablesTypeUuid = *allowedResponseVariablesTypeUuidPtr

	var tempResponseVariables map[string]*fenixGuiTestCaseBuilderServerGrpcApi.ResponseVariableMessage
	var tempImmatureTestInstructionMessage *fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionMessage

	// Extract Immature TestInstruction
	var existInMap bool
	tempImmatureTestInstructionMessage, existInMap = testCasesUiCanvasObject.TestCasesModelReference.
		AvailableImmatureTestInstructionsMap[currentElement.MatureTestCaseModelElementMessage.GetOriginalElementUuid()]

	if existInMap == false {
		errorId := "3bfeb6c1-eba1-4a69-9374-3bce36f48746"
		err = errors.New(fmt.Sprintf("ImmatureTestInstruction, '%s', couldn't be found in "+
			"'AvailableImmatureTestInstructionsMap' + [ErrorID: %s]",
			currentElement.MatureTestCaseModelElementMessage.GetOriginalElementUuid(),
			errorId))

		//TODO Send ERRORS over error-channel
		fmt.Println(err)

		return err
	}

	tempResponseVariables = tempImmatureTestInstructionMessage.GetResponseVariablesMapStructure().
		ResponseVariablesMap

	// Loop TestInstructions ResponseVariables
	for _, tempResponseVariable := range tempResponseVariables {

		// Loop ResponseVariableTypes To match
		for _, allowedResponseVariableTypeUuid := range allowedResponseVariablesTypeUuid {

			// When TestInstructions Response Variable Type is the same as the ResponseVariableType then add to Slice
			if allowedResponseVariableTypeUuid == tempResponseVariable.GetResponseVariableTypeUuid() {
				// Add TestInstructions Original UUID Name
				var tempMatureTestInstructionWithCorrectResponseVariablesType *testCaseModel.
					MatureTestInstructionWithCorrectResponseVariablesTypeStruct

				tempMatureTestInstructionWithCorrectResponseVariablesType = &testCaseModel.
					MatureTestInstructionWithCorrectResponseVariablesTypeStruct{
					MatureTestInstructionUuidWithCorrectResponseVariablesType: currentElement.
						MatureTestCaseModelElementMessage.GetMatureElementUuid(),
					MatureTestInstructionNameWithCorrectResponseVariablesType: currentElement.
						MatureTestCaseModelElementMessage.GetOriginalElementName(),
					MatureTestInstructionComboBoxOptionsName: "",
				}
				*testInstructionWithCorrectResponseVariablesTypePtr = append(*testInstructionWithCorrectResponseVariablesTypePtr,
					tempMatureTestInstructionWithCorrectResponseVariablesType)
			}
		}

	}
	return err
}
