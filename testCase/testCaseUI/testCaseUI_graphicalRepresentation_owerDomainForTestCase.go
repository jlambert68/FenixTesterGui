package testCaseUI

import (
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
)

// Generate the OwnerDomain Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateOwnerDomainForTestCaseArea(
	testCaseUuid string) (
	ownerDomainArea fyne.CanvasObject,
	tempCurrentOwnerDomainToBeChosenInDropDown string,
	newOwnerDomainSelect *widget.Select,
	valueIsValidWarningBox *canvas.Rectangle,
	err error) {

	var existsInMap bool

	// Get TestCasesMap
	var testCasesMap map[string]*testCaseModel.TestCaseModelStruct
	testCasesMap = *testCasesUiCanvasObject.TestCasesModelReference.TestCasesMapPtr

	// Get current TestCase
	var currentTestCasePtr *testCaseModel.TestCaseModelStruct
	currentTestCasePtr, existsInMap = testCasesMap[testCaseUuid]

	if existsInMap == false {
		errorId := "bb7fe228-2079-481f-89d3-8cf07a4da26a"
		err := errors.New(fmt.Sprintf("testcase-model with TestCaseUuid '%s' is missing map for TestCasesMapPtr [ErrorID: %s]", testCaseUuid, errorId))

		//TODO Send ERRORS over error-channel
		fmt.Println(err)

		return nil,
			"",
			nil,
			nil,
			err

	}

	// If TestCase already has a chosen OwnerDomain then set that value
	var tempCurrentOwnerDomain string

	var testCaseHasOwnerDomain bool

	if len(currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.DomainUuid) > 0 {
		testCaseHasOwnerDomain = true
		tempCurrentOwnerDomain = currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.DomainUuid
	}

	// Load Domains that can own the TestCase into options-array
	var options []string
	for _, tempDomainsThatCanOwnTheTestCase := range testCasesUiCanvasObject.TestCasesModelReference.DomainsThatCanOwnTheTestCaseMap {
		options = append(options, tempDomainsThatCanOwnTheTestCase.DomainNameShownInGui)

		// When TestCase has OwnerDomain find the one
		if testCaseHasOwnerDomain == true && tempDomainsThatCanOwnTheTestCase.DomainUuid == tempCurrentOwnerDomain {
			tempCurrentOwnerDomainToBeChosenInDropDown = tempDomainsThatCanOwnTheTestCase.DomainNameShownInGui
		}
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

			// Save TestCase back in Map
			// Get the latest version of TestCase
			tempTestCasePtr, _ := testCasesMap[testCaseUuid]

			// Store Domain in LocalTestCase in TestCase-model
			tempTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.DomainUuid =
				testCasesUiCanvasObject.TestCasesModelReference.DomainsThatCanOwnTheTestCaseMap[value].DomainUuid
			tempTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.DomainName =
				testCasesUiCanvasObject.TestCasesModelReference.DomainsThatCanOwnTheTestCaseMap[value].DomainName

			// Get the TestCaseUI-model
			var testCasesUiModelMap map[string]*testCaseGraphicalAreasStruct
			var testCaseGraphicalAreas *testCaseGraphicalAreasStruct
			testCasesUiModelMap = testCasesUiCanvasObject.TestCasesUiModelMap
			testCaseGraphicalAreas, _ = testCasesUiModelMap[testCaseUuid]

			//var testCaseMetaDataArea fyne.CanvasObject
			var metaDataAccordion *widget.Accordion
			_, metaDataAccordion, err = testCasesUiCanvasObject.GenerateMetaDataAreaForTestCase(
				tempTestCasePtr,
				testCaseUuid,
				testCasesUiCanvasObject.TestCasesModelReference.DomainsThatCanOwnTheTestCaseMap[value].DomainUuid)

			if err != nil {
				log.Println(err, metaDataAccordion)

				return
			}

			//var testCaseMetaDataAreaContainer fyne.Container
			//var ok binding.BoolList
			var TestCaseMetaDataAreaCanvasObject fyne.CanvasObject
			TestCaseMetaDataAreaCanvasObject = *testCaseGraphicalAreas.TestCaseMetaDataArea
			if testCaseMetaDataAreaContainer, ok := TestCaseMetaDataAreaCanvasObject.(*fyne.Container); ok {
				if ok != true {
					log.Fatalln("couldn't cast to fyne.Container")
				}

				metaDataAccordion.OpenAll()

				testCaseMetaDataAreaContainer.Objects[0] = metaDataAccordion

				testCaseMetaDataAreaContainer.Refresh()
			}

			/*
					metaDataAccordionItem2 := *metaDataAccordionItem
					if metaDataAccordionItemCanvas, ok :=  metaDataAccordionItem2.(*fyne.CanvasObject); ok {
						if ok != true {
							log.Fatalln("couldn't cast to fyne.Container")
						}
					testCaseMetaDataAreaContainer.Objects[0] = metaDataAccordionItem
				}


				testCaseMetaDataAreaContainer = *testCaseGraphicalAreas.TestCaseMetaDataArea
				testCaseMetaDataAreaContainer, err = (fyne.Container).(testCaseGraphicalAreas)
				//*testCaseGraphicalAreas.TestCaseMetaDataArea. = testCaseMetaDataArea
			*/
			// Save back TestCaseUI
			testCasesUiModelMap[testCaseUuid] = testCaseGraphicalAreas

			testCasesUiCanvasObject.TestCasesUiModelMap = testCasesUiModelMap

			// Clear the set MetaData in the TestCase
			tempTestCasePtr.TestCaseMetaDataPtr = nil

			//testCaseMetaDataArea.Refresh()

			// Store back TestCase-model in Map
			//testCasesUiCanvasObject.TestCasesModelReference.TestCasesMapPtr[testCaseUuid] = tempTestCasePtr

			// Set Warning box that value is not selected
			if len(value) == 0 {
				valueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
			} else {
				valueIsValidWarningBox.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
			}
		})

	/*
		// Set the Visible value for DropDown, if there is any
		if len(tempCurrentOwnerDomainToBeChosenInDropDown) > 0 {
			newOwnerDomainSelect.SetSelected(tempCurrentOwnerDomainToBeChosenInDropDown)
		}

		// Set correct warning box color
		if len(newOwnerDomainSelect.Selected) == 0 {
			valueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
		} else {
			valueIsValidWarningBox.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
		}
	*/

	// Create a custom SelectComboBox, with valueIsValidWarningBox
	var customSelectComboBox *customAttributeSelectComboBox
	customSelectComboBox = newCustomAttributeSelectComboBoxWidget(newOwnerDomainSelect, valueIsValidWarningBox)

	// Add the Entry-widget to the Forms-container
	testCaseOwnerDomainNameFormContainer.Add(customSelectComboBox)

	// Create the VBox-container that will be returned
	testCaseOwnerDomainContainer = container.NewVBox(testCaseOwnerDomainNameFormContainer)

	return testCaseOwnerDomainContainer,
		tempCurrentOwnerDomainToBeChosenInDropDown,
		newOwnerDomainSelect,
		valueIsValidWarningBox,
		err
}

// Sets the Selected value for the DropDown specifying the Owner-Domain of the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) setSelectedOwnerDomainForTestCaseArea(
	tempCurrentOwnerDomainToBeChosenInDropDown string,
	newOwnerDomainSelect *widget.Select,
	valueIsValidWarningBox *canvas.Rectangle) {

	// Set the Visible value for DropDown, if there is any
	if len(tempCurrentOwnerDomainToBeChosenInDropDown) > 0 {
		newOwnerDomainSelect.SetSelected(tempCurrentOwnerDomainToBeChosenInDropDown)
	}

	// Set correct warning box color
	if len(newOwnerDomainSelect.Selected) == 0 {
		valueIsValidWarningBox.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	} else {
		valueIsValidWarningBox.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
	}
}
