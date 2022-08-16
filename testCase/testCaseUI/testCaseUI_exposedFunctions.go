package testCaseUI

import (
	"errors"
	"fmt"
)

// GenerateShortUuidFromFullUuid
// Generate a short version of the UUID to be used in GUI
func (testCasesUiCanvasObject *TestCasesUiModelStruct) UpdateGraphicalRepresentationForTestCase(testCaseUuid string) (err error) {

	// Get current TestCase-UI-model
	currentTestCaseUIModel, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]

	if existsInMap == true {
		errorId := "ec24f71b-d828-4cd0-a319-9f118ebbdccf"
		err = errors.New(fmt.Sprintf("testcase-UI-model with uuid '%s' couldn't be foundin 'TestCasesUiModelMap' [ErrorID: %s]", testCaseUuid, errorId))

		return err
	}

	// Open and Update Accordion object and Tree-model
	currentTestCaseUIModel.currentTestCaseGraphicalStructure.currentTestCaseGraphicalAccordionObject.OpenAll()
	currentTestCaseUIModel.currentTestCaseGraphicalStructure.currentTestCaseGraphicalAccordionObject.Refresh()
	currentTestCaseUIModel.currentTestCaseGraphicalStructure.currentTestCaseGraphicalTreeComponent.Refresh()

	return err
}
