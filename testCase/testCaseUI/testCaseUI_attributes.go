package testCaseUI

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Generate the TestCaseAttributes Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateTestCaseAttributesAreaForTestCase(testCaseUuid string) (testCaseAttributesArea fyne.CanvasObject, err error) {

	// Get current TestCase-UI-model
	_, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]

	if existsInMap == true {
		errorId := "c4110d4f-3dca-48bd-a8e4-57cb040fe079"
		err = errors.New(fmt.Sprintf("testcase-UI-model with uuid '%s' allready exist in 'TestCasesUiModelMap' [ErrorID: %s]", testCaseUuid, errorId))

		return nil, err
	}

	testCaseAttributesArea = widget.NewLabel("'testCaseAttributesArea'")

	return testCaseAttributesArea, err
}
