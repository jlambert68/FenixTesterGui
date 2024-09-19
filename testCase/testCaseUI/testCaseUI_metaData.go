package testCaseUI

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Generate the MetaData Area for the TestCase
func (testCasesUiCanvasObject *TestCasesUiModelStruct) generateMetaDataAreaForTestCase(
	testCaseUuid string) (
	testCaseMetaDataArea fyne.CanvasObject, err error) {

	var metaDataAccordionItem *widget.AccordionItem
	var accordion *widget.Accordion
	var metaDataArea fyne.CanvasObject

	// Get current TestCase-UI-model
	_, existsInMap := testCasesUiCanvasObject.TestCasesUiModelMap[testCaseUuid]

	if existsInMap == true {
		errorId := "bcb9d984-3106-42b6-9c23-288ec6d26224"
		err = errors.New(fmt.Sprintf("testcase-UI-model with sourceUuid '%s' allready exist in 'TestCasesUiModelMap' [ErrorID: %s]", testCaseUuid, errorId))

		return nil, err
	}

	// Create an Accordion item for the MetaData
	metaDataAccordionItem = widget.NewAccordionItem("TestCase MetaData", widget.NewLabel("'testCaseMetaDataArea'"))

	accordion = widget.NewAccordion(metaDataAccordionItem)

	// Create the VBox-container that will be returned
	metaDataArea = container.NewVBox(accordion, widget.NewLabel(""), widget.NewSeparator())

	return metaDataArea, err
}
