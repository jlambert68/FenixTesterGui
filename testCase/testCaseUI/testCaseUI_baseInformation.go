package testCaseUI

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
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

	//testCaseBaseInformationArea = widget.NewLabel("'testCaseBaseInformationArea'")
	testCaseBaseInformationArea, err = testCasesUiCanvasObject.generateTestCaseNameArea(testCaseUuid)

	return testCaseBaseInformationArea, err
}
