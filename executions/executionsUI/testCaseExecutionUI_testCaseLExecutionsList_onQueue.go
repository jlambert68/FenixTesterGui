package executionsUI

import (
	"FenixTesterGui/executions/executionsModel"
	"FenixTesterGui/headertable"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var data = [][]string{[]string{"top left", "top right"},
	[]string{"bottom left", "bottom right"}}

func createTable() (mytable *widget.Table) {

	mytable = widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})

	return mytable
}

var myTableOpts = headertable.TableOpts{
	RefWidth: "reference width",
	ColAttrs: []headertable.ColAttr{
		{
			Name:         "Name",
			Header:       "Name",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},
		{
			Name:         "Weight",
			Header:       "Weight",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},
		{
			Name:         "Type",
			Header:       "Type",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 80,
		},
		{
			Name:         "Color",
			Header:       "Color",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},
	},
}

type Animal struct {
	Name, Type, Color, Weight string
}

var animals = []Animal{
	{Name: "Frisky", Type: "cat", Color: "gray", Weight: "10"},
	{Name: "Ella", Type: "dog", Color: "brown", Weight: "50"},
	{Name: "Mickey", Type: "mouse", Color: "black", Weight: "1"},
	{Name: "Frisky", Type: "cat", Color: "gray", Weight: "10"},
	{Name: "Ella", Type: "dog", Color: "brown", Weight: "50"},
	{Name: "Mickey", Type: "mouse", Color: "black", Weight: "1"},
	{Name: "Frisky", Type: "cat", Color: "gray", Weight: "10"},
	{Name: "Ella", Type: "dog", Color: "brown", Weight: "50"},
	{Name: "Mickey", Type: "mouse", Color: "black", Weight: "1"},
	{Name: "Frisky", Type: "cat", Color: "gray", Weight: "10"},
	{Name: "Ella", Type: "dog", Color: "brown", Weight: "50"},
	{Name: "Mickey", Type: "mouse", Color: "black", Weight: "1"},
	{Name: "Frisky", Type: "cat", Color: "gray", Weight: "10"},
	{Name: "Ella", Type: "dog", Color: "brown", Weight: "50"},
	{Name: "Mickey", Type: "mouse", Color: "black", Weight: "1"},
}

func MySortTable() *fyne.Container {
	var AnimalBindings []binding.DataMap

	// Create a binding for each animal data
	for i := 0; i < len(animals); i++ {
		AnimalBindings = append(AnimalBindings, binding.BindStruct(&animals[i]))
	}
	myTableOpts.Bindings = AnimalBindings

	ht := headertable.NewSortingHeaderTable(&myTableOpts)
	mySortTable := container.NewMax(ht)

	return mySortTable

}

func CreateTableObject() (testCaseTextualModelAreaAccordion *fyne.Container) {

	mytable := createTable()

	newTableSize := fyne.NewSize(mytable.MinSize().Width, mytable.MinSize().Height*3)

	mytable.Resize(newTableSize)

	myTableContainer := container.New(layout.NewVBoxLayout(), mytable, widget.NewLabel("Test"))
	myTableContainer.Resize(newTableSize)
	fmt.Println(mytable.MinSize())

	// Create a Canvas Accordion type for grouping the Textual Representations
	testCaseTextualModelAreaAccordionItem := widget.NewAccordionItem("TestCaseExecutions on ExecutionQueue", myTableContainer)

	myAccordian := widget.NewAccordion(testCaseTextualModelAreaAccordionItem)

	myAccordian.Open(0)
	myAccordian.Refresh()

	//testCaseTextualModelAreaAccordion.Resize(newTableSize)

	//testCaseTextualModelAreaAccordion.Refresh()

	myContainer := container.New(layout.NewMaxLayout(), myAccordian)

	return myContainer
}

func CreateTableForTestCaseExecutionsOnQueue() *fyne.Container {
	var tableForTestCaseExecutionsOnQueueBindings []binding.DataMap

	// Create a binding for each TestExecutionOnQueueRow data
	for _, tempTestCaseExecutionsUnderExecutionDataAdaptedForUiTableReference := range executionsModel.TestCaseExecutionsOnQueueMapAdaptedForUiTable {
		tableForTestCaseExecutionsOnQueueBindings = append(
			tableForTestCaseExecutionsOnQueueBindings,
			binding.BindStruct(tempTestCaseExecutionsUnderExecutionDataAdaptedForUiTableReference))
	}

	executionsModel.TestCaseExecutionsOnQueueTableOptions.Bindings = tableForTestCaseExecutionsOnQueueBindings

	ht := headertable.NewSortingHeaderTable(&executionsModel.TestCaseExecutionsOnQueueTableOptions)
	mySortTable := container.NewMax(ht)

	/*
		first := container.NewHBox(widget.NewLabel("Första"), mySortTable, widget.NewLabel("Sista"))

		second := container.NewHBox(first)

		scrollableTable := container.NewScroll(second)

		scrollableTableContainer := container.NewMax(scrollableTable)
	*/
	// ht.Header.ScrollToLeading()

	return mySortTable

}
