package testCaseUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"log"
)

type clickableAccordion struct {
	widget.Accordion
	isClickable bool
}

func (testCasesUiCanvasObject *TestCasesUiModelStruct) newClickableAccordion(accordionItem *widget.AccordionItem, isClickable bool) *clickableAccordion {
	accordion := &clickableAccordion{}
	accordion.ExtendBaseWidget(accordion)

	accordion.Append(accordionItem)

	accordion.isClickable = isClickable

	return accordion
}

func (b *clickableAccordion) TappedSecondary(_ *fyne.PointEvent) {

	if b.isClickable == false {
		return
	}
	log.Println("I have been Secondary tapped")

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
