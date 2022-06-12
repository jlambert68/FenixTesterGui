package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var commandStackList []string
var commandStackListUI *widget.List

func makeCommandStackUI() {

	commandStackList = []string{"MyValue 1", "MyValue 2", "MyValue 3"}

	bindedCommandListData := binding.NewStringList()
	bindedCommandListData.Set(commandStackList)

	commandStackListUI = widget.NewListWithData(bindedCommandListData,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	commandStackListUI.OnSelected = func(id widget.ListItemID) {
		commandStackListUI.Unselect(id)
		d, _ := bindedCommandListData.GetValue(id)
		w := myTestCase.fyneApp.NewWindow("Edit Data")

		itemName := widget.NewEntry()
		itemName.Text = d

		updateData := widget.NewButton("Update", func() {
			bindedCommandListData.SetValue(id, itemName.Text)
			w.Close()
		})

		cancel := widget.NewButton("Cancel", func() {
			w.Close()
		})

		deleteData := widget.NewButton("Delete", func() {
			var newData []string
			dt, _ := bindedCommandListData.Get()

			for index, item := range dt {
				if index != id {
					newData = append(newData, item)
				}
			}

			bindedCommandListData.Set(newData)

			w.Close()
		})

		w.SetContent(container.New(layout.NewVBoxLayout(), itemName, updateData, deleteData, cancel))
		w.Resize(fyne.NewSize(400, 200))
		w.CenterOnScreen()
		w.Show()

	}
	/*
		add := widget.NewButton("Add", func() {
			w := myTestCase.fyneApp.NewWindow("Add Data")

			itemName := widget.NewEntry()

			addData := widget.NewButton("Add", func() {
				bindedCommandListData.Append(itemName.Text)
				w.Close()
			})

			cancel := widget.NewButton("Cancel", func() {
				w.Close()
			})

			w.SetContent(container.New(layout.NewVBoxLayout(), itemName, addData, cancel))
			w.Resize(fyne.NewSize(400, 200))
			w.CenterOnScreen()
			w.Show()

		})


	*/

}
