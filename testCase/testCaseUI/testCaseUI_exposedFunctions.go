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

	if existsInMap == false {
		errorId := "ec24f71b-d828-4cd0-a319-9f118ebbdccf"
		err = errors.New(fmt.Sprintf("testcase-UI-model with sourceUuid '%s' couldn't be foundin 'TestCasesUiModelMap' [ErrorID: %s]", testCaseUuid, errorId))

		return err
	}

	// Generate new version of graphical TestCase based on latest TestCase-Element Model
	testCaseGraphicalModelArea, _, accordion, err := testCasesUiCanvasObject.generateGraphicalRepresentationAreaForTestCase(testCaseUuid)

	currentTestCaseUIModel.TestCaseGraphicalModelArea = testCaseGraphicalModelArea

	// Open and Update Accordion object(Graphical) and Tree-model
	currentTestCaseUIModel.currentTestCaseGraphicalStructure.currentTestCaseGraphicalAccordionObject.RemoveIndex(0)
	currentTestCaseUIModel.currentTestCaseGraphicalStructure.currentTestCaseGraphicalAccordionObject.Append(accordion.Items[0])
	currentTestCaseUIModel.currentTestCaseGraphicalStructure.currentTestCaseGraphicalAccordionObject.Open(0)
	//currentTestCaseUIModel.currentTestCaseGraphicalStructure.currentTestCaseGraphicalAccordionObject.Refresh()

	//currentTestCaseUIModel.currentTestCaseGraphicalStructure.currentTestCaseGraphicalTreeComponent.OpenAllBranches()

	// Save TestCase UI-components-Map
	testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid] = currentTestCaseUIModel

	testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid].TestCaseGraphicalModelArea.Refresh()

	return err
}
