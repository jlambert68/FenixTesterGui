package testCaseUI

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Generate the BaseInformation Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateBaseInformationAreaForTestCase(
	testCaseUuid string) (
	testCaseBaseInformationArea fyne.CanvasObject,
	// Variables to be able to Set value in OwnerDomain-Dropdown
	tempCurrentOwnerDomainToBeChosenInDropDown string,
	newOwnerDomainSelect *widget.Select,
	valueIsValidWarningBox *canvas.Rectangle,
	// *****************
	err error) {

	// Get current TestCase-UI-model
	_, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]

	if existsInMap == true {
		errorId := "4b062436-590a-4f2a-9004-181f3f575a4b"
		err = errors.New(fmt.Sprintf("testcase-UI-model with sourceUuid '%s' allready exist in 'TestCasesUiModelMap' [ErrorID: %s]", testCaseUuid, errorId))

		return nil,
			"",
			nil,
			nil,
			err
	}

	var tempBaseInformationAreaContainer *fyne.Container
	tempBaseInformationAreaContainer = container.New(layout.NewVBoxLayout())

	// Generate "All Mandatory Fields" filled in as expected

	// Add the ' "AllMandatoryFields"-UI-object' to the 'BaseInformationArea'

	// Generate TestCaseDeleteDate-UI-object
	var tempTestCaseDeleteDateArea fyne.CanvasObject
	tempTestCaseDeleteDateArea, err = testCasesUiCanvasObject.generateTestCaseDeletionDateArea(testCaseUuid)
	if err != nil {
		return nil,
			"",
			nil,
			nil,
			err
	}

	// Add the 'TestCaseName-UI-object' to the 'BaseInformationArea'
	tempBaseInformationAreaContainer.Add(tempTestCaseDeleteDateArea)

	// Generate TestCaseName-UI-object
	var tempTestCaseNameArea fyne.CanvasObject
	tempTestCaseNameArea, err = testCasesUiCanvasObject.generateTestCaseNameArea(testCaseUuid)
	if err != nil {
		return nil,
			"",
			nil,
			nil,
			err
	}

	// Add the 'TestCaseName-UI-object' to the 'BaseInformationArea'
	tempBaseInformationAreaContainer.Add(tempTestCaseNameArea)

	// Generate TestCaseDomainOwner-UI-object
	var tempTestCaseDomainOwnerArea fyne.CanvasObject
	tempTestCaseDomainOwnerArea,
		tempCurrentOwnerDomainToBeChosenInDropDown,
		newOwnerDomainSelect,
		valueIsValidWarningBox,
		err = testCasesUiCanvasObject.generateOwnerDomainForTestCaseArea(testCaseUuid)
	if err != nil {
		return nil,
			"",
			nil,
			nil,
			err
	}

	// Add the 'TestCaseDomainOwner-UI-object' to the 'BaseInformationArea'
	tempBaseInformationAreaContainer.Add(tempTestCaseDomainOwnerArea)

	// Generate TestCaseDescription-UI-object
	var tempTestCaseDescriptionArea fyne.CanvasObject
	tempTestCaseDescriptionArea, err = testCasesUiCanvasObject.generateTestCaseDescriptionArea(testCaseUuid)
	if err != nil {
		return nil,
			"",
			nil,
			nil,
			err
	}

	// Add the 'TestCaseDescription-UI-object' to the 'BaseInformationArea'
	tempBaseInformationAreaContainer.Add(tempTestCaseDescriptionArea)

	// Generate TemplateList-UI-object
	var templateListArea fyne.CanvasObject
	templateListArea, err = testCasesUiCanvasObject.generateTemplateListForTestCaseArea(testCaseUuid)
	if err != nil {
		return nil,
			"",
			nil,
			nil,
			err
	}

	// Add the 'TemplateList-UI-object' to the 'BaseInformationArea'
	tempBaseInformationAreaContainer.Add(templateListArea)

	// Generate TestDataSelector-UI-object
	var testDataSelectorArea fyne.CanvasObject
	testDataSelectorArea, err = testCasesUiCanvasObject.generateSelectedTestDataForTestCaseArea(testCaseUuid)
	if err != nil {
		return nil,
			"",
			nil,
			nil,
			err
	}

	// Add the 'TestDataSelector-UI-object' to the 'BaseInformationArea'
	tempBaseInformationAreaContainer.Add(testDataSelectorArea)

	return tempBaseInformationAreaContainer,
		tempCurrentOwnerDomainToBeChosenInDropDown,
		newOwnerDomainSelect,
		valueIsValidWarningBox,
		err
}
