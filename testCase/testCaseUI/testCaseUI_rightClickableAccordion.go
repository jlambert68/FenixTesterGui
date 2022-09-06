package testCaseUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"log"
)

type clickableAccordion struct {
	widget.Accordion
	isClickable            bool
	testCasesUiModelStruct *TestCasesUiModelStruct
	testCaseUuid           string
	testInstructionUuid    string
}

func (testCasesUiCanvasObject *TestCasesUiModelStruct) newClickableAccordion(accordionItem *widget.AccordionItem, isClickable bool, testCaseUuid, testInstructionUuid string) *clickableAccordion {
	accordion := &clickableAccordion{}
	accordion.ExtendBaseWidget(accordion)

	accordion.Append(accordionItem)

	accordion.isClickable = isClickable
	accordion.testCaseUuid = testCaseUuid
	accordion.testInstructionUuid = testInstructionUuid

	accordion.testCasesUiModelStruct = testCasesUiCanvasObject

	return accordion

}

func (b *clickableAccordion) TappedSecondary(_ *fyne.PointEvent) {

	if b.isClickable == false {
		return
	}
	log.Println("I have been Secondary tapped")

	b.testCasesUiModelStruct.generateTestCaseAttributesAreaForTestCase(b.testCaseUuid, b.testInstructionUuid)
	//generateTestCaseAttributesAreaForTestCase()

	/*
		// Generate the TestCaseAttributes Area for the TestCase
		testCaseAttributesArea, err := b.testCasesModelReference.generateTestCaseAttributesAreaForTestCase("") // "" used for first time creation

		if err != nil {
			return err
		}
		// Add newly created TestCaseAttributes Area to object for all graphical parts of one TestCase
		testCaseGraphicalAreas.TestCaseAttributesArea = testCaseAttributesArea
	*/

}

// **********************************************************************

type clickableAccordionItem struct {
	widget.AccordionItem
	isClickable bool
}

func (testCasesUiCanvasObject *TestCasesUiModelStruct) newClickableAccordionItem(title string, detail fyne.CanvasObject) *clickableAccordionItem {
	//accordionItem.ExtendBaseWidget(accordionItem)
	accordionItem := &clickableAccordionItem{
		AccordionItem: widget.AccordionItem{
			Title:  title,
			Detail: detail,
			Open:   true,
		},
		isClickable: true,
	}

	return accordionItem
}

func (b *clickableAccordionItem) TappedSecondary(_ *fyne.PointEvent) {

	if b.isClickable == false {
		return
	}
	log.Println("I'm an AccordionItem and I have been Secondary tapped")

}
