package restAPI

import "github.com/sirupsen/logrus"

// SetLogger
// Set to use the same logger reference as is used by central part of system
func (restAPI *RestApiStruct) SetLogger(logger *logrus.Logger) {
	restAPI.logger = logger

	return

}

// SetDialAddressString
// Set the Dial Address, which was received from environment variables
func (restAPI *RestApiStruct) SetDialAddressString(dialAddress string) {
	restAPI.fenixGuiBuilderServerAddressToDial = dialAddress

	return

}
