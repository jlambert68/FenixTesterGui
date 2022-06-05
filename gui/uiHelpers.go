package gui

import "github.com/sirupsen/logrus"

// SetLogger
// Set to use the same logger reference as is used by central part of system
func (uiServer *UIServerStruct) SetLogger(logger *logrus.Logger) {
	uiServer.logger = logger

	return

}
