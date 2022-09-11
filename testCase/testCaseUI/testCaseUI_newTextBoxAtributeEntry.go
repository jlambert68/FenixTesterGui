package testCaseUI

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type AttributeEntry struct {
	widget.Entry
	attributeUuid string
}

func (testCasesUiCanvasObject *TestCasesUiModelStruct) NewAttributeEntry(attributeUuid string) *AttributeEntry {
	myAttributeEntry := &AttributeEntry{}
	myAttributeEntry.ExtendBaseWidget(widget.NewEntry())

	myAttributeEntry.attributeUuid = attributeUuid

	return myAttributeEntry
}

// FocusGained - Single Click on colorRectangle
func (c *AttributeEntry) FocusGained(x fyne.Focusable) {

	fmt.Println("FocusGained")

	x.FocusGained()
}

// FocusLost - Single Click on colorRectangle
func (c *AttributeEntry) FocusLost(x fyne.Focusable) {

	fmt.Println("Lost Focus")

	x.FocusLost()
}

/*
// TypedRune - Single Click on colorRectangle
func (c *AttributeEntry) TypedRune(x fyne.Focusable) {

	fmt.Println("TypedRune")

	x.TypedRune(x.)
}

// TypedKey - Single Click on colorRectangle
func (c *AttributeEntry) TypedKey(x fyne.Focusable) {

	fmt.Println("Lost Focus")

	x.TypedKey(x)
}


// TypedRune is a hook called by the input handling logic on text input events if this object is focused.
TypedRune(rune)
// TypedKey is a hook called by the input handling logic on key events if this object is focused.
TypedKey(*KeyEvent)
*/
