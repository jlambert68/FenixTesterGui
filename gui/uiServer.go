package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
)

// Main UI server module
func (uiServerStruct *UIServerStruct) StartUIServer() {

	uiServerStruct.logger.WithFields(logrus.Fields{
		"id": "a4d2716f-ded1-4062-bffb-fd0c03d69ca3",
	}).Debug("Starting UI server")

	a := app.NewWithID("se.fenix.testcasebuilder")
	//a.Settings().SetTheme(&myTheme{})
	w := a.NewWindow("Fenix TestCase Builder")

	//list := &notelist{pref: a.Preferences()}
	//list.load()
	//notesUI := &ui{notes: list}

	w.SetContent(widget.NewLabel("Fenix TestCase Builder"))
	//notesUI.registerKeys(w)

	w.Resize(fyne.NewSize(400, 320))
	w.ShowAndRun()

}
