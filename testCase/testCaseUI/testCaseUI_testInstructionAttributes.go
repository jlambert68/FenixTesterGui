package testCaseUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type attributeStruct struct {
	attributeName  string
	attributeValue string
}

var attributesList []attributeStruct
var attributeListUI *widget.List
var bindedAttributeListData binding.StringList

// Generate the TestCaseAttributes Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateTestCaseAttributesAreaForTestCase(testCaseUuid string) (testCaseAttributesArea fyne.CanvasObject, err error) { //, testInstructionElementUuid string) (testCaseAttributesArea fyne.CanvasObject, err error) {
	/*
		att

		// Get current TestCase-UI-model
		_, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]

		if existsInMap == true {
			errorId := "c4110d4f-3dca-48bd-a8e4-57cb040fe079"
			err = errors.New(fmt.Sprintf("testcase-UI-model with sourceUuid '%s' allready exist in 'TestCasesUiModelMap' [ErrorID: %s]", testCaseUuid, errorId))

			return nil, err
		}

		// Extract immature element TODO change to read from TestCase instead
		testInstructionElement, existInMap :=  testCasesUiCanvasObject.TestCasesModelReference.AvailableImmatureTestInstructionsMap[testInstructionElementUuid]

		if existsInMap == false {
			errorId := "d2ab1a27-7398-49cc-9c52-f39d0eb0a9f2"
			err = errors.New(fmt.Sprintf("element with UUID '%s' doesn't exist in among available immature TestInstructions' [ErrorID: %s]", testCaseUuid, errorId))

			return nil, err
		}
		testCasesUiCanvasObject.TestCasesModelReference.


	*/

	//attributesContainer := container.New(layout.NewFormLayout(),
	testCaseAttributesArea = widget.NewLabel("'testCaseAttributesArea'")

	return testCaseAttributesArea, err
}
