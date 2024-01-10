package testCaseUI

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

// Generate the BaseInformation Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateBaseInformationAreaForTestCase(testCaseUuid string) (testCaseBaseInformationArea fyne.CanvasObject, err error) {

	// Get current TestCase-UI-model
	_, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]

	if existsInMap == true {
		errorId := "4b062436-590a-4f2a-9004-181f3f575a4b"
		err = errors.New(fmt.Sprintf("testcase-UI-model with sourceUuid '%s' allready exist in 'TestCasesUiModelMap' [ErrorID: %s]", testCaseUuid, errorId))

		return nil, err
	}

	var tempBaseInformationAreaContainer *fyne.Container
	tempBaseInformationAreaContainer = container.New(layout.NewVBoxLayout())

	// Generate TestCaseName-UI-object
	var tempTestCaseNameArea fyne.CanvasObject
	tempTestCaseNameArea, err = testCasesUiCanvasObject.generateTestCaseNameArea(testCaseUuid)
	if err != nil {
		return nil, err
	}

	// Add the 'TestCaseName-UI-object' to the 'BaseInformationArea'
	tempBaseInformationAreaContainer.Add(tempTestCaseNameArea)

	// Generate TestCaseDomainOwner-UI-object
	var tempTestCaseDomainOwnerArea fyne.CanvasObject
	tempTestCaseDomainOwnerArea, err = testCasesUiCanvasObject.generateOwnerDomainForTestCaseArea(testCaseUuid)
	if err != nil {
		return nil, err
	}

	// Add the 'TestCaseDomainOwner-UI-object' to the 'BaseInformationArea'
	tempBaseInformationAreaContainer.Add(tempTestCaseDomainOwnerArea)

	// Generate TestCaseDescription-UI-object
	var tempTestCaseDescriptionArea fyne.CanvasObject
	tempTestCaseDescriptionArea, err = testCasesUiCanvasObject.generateTestCaseDescriptionArea(testCaseUuid)
	if err != nil {
		return tempBaseInformationAreaContainer, err
	}

	// Add the 'TestCaseDescription-UI-object' to the 'BaseInformationArea'
	tempBaseInformationAreaContainer.Add(tempTestCaseDescriptionArea)

	return tempBaseInformationAreaContainer, err
}
