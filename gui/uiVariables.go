package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
)

type UIServerStruct struct {
	logger  *logrus.Logger
	fyneApp fyne.App
	tree    *widget.Label
	content *widget.Entry
}
