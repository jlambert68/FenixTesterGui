package gui

import (
	"github.com/sirupsen/logrus"
)

// SetLogger
// Set to use the same logger reference as is used by central part of system
func (availableBuildingBlocksModel *availableBuildingBlocksModelStruct) SetLogger(logger *logrus.Logger) {

	//myUIServer = UIServerStruct{}
	availableBuildingBlocksModel.logger = logger

	return

}

// SetDialAddressString
// Set the Dial Address, which was received from environment variables
func (availableBuildingBlocksModel *availableBuildingBlocksModelStruct) SetDialAddressString(dialAddress string) {
	availableBuildingBlocksModel.fenixGuiBuilderServerAddressToDial = dialAddress

	return

}

// SetLogger
// Set to use the same logger reference as is used by central part of system
func (uiServer *GlobalUIServerStruct) SetLogger(logger *logrus.Logger) {

	//myUIServer = UIServerStruct{}
	uiServer.logger = logger

	return

}

// SetDialAddressString
// Set the Dial Address, which was received from environment variables
func (uiServer *GlobalUIServerStruct) SetDialAddressString(dialAddress string) {
	uiServer.fenixGuiBuilderServerAddressToDial = dialAddress

	return

}

// SetLogger
// Set to use the same logger reference as is used by central part of system
func (uiServer *UIServerStruct) SetLogger(logger *logrus.Logger) {

	//myUIServer = UIServerStruct{}
	uiServer.logger = logger

	return

}

// SetDialAddressString
// Set the Dial Address, which was received from environment variables
func (uiServer *UIServerStruct) SetDialAddressString(dialAddress string) {
	uiServer.fenixGuiBuilderServerAddressToDial = dialAddress

	return

}
