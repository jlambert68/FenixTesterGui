package testSuitesCommandEngine

import (
	sharedCode "FenixTesterGui/common_code"
	"github.com/sirupsen/logrus"
)

// InitiateTestSuiteCommandChannelReader
// Initiate and start the Channel reader which is used for reading out commands for processing certain tasks regarding TestSuite
func InitiateTestSuiteCommandChannelReader() {

	go startTestSuiteCommandChannelReader()

}

// Channel reader which is used for reading out commands for processing certain tasks regarding TestSuite
func startTestSuiteCommandChannelReader() {

	var incomingChannelCommand CommandTestSuiteChannelStruct

	for {
		// Wait for incoming command over channel
		incomingChannelCommand = <-TestSuiteCommandChannel

		switch incomingChannelCommand.ChannelCommand {

		case TestSuiteChannelCommandRefreshTestSuiteTabsObject:
			testSuiteChannelCommandRefreshTestSuiteTabsObject(incomingChannelCommand)

		// No other command is supported
		default:
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":                                    "7683a442-a722-499b-ab86-7710c6d8cea3",
				"incomingChannelCommand.ChannelCommand": incomingChannelCommand.ChannelCommand,
			}).Fatal("Unsupported 'ChannelCommand'. This should not happen")

		}
	}

}

// Refresh Tabs-object for all TestSuites
func testSuiteChannelCommandRefreshTestSuiteTabsObject(incomingChannelCommand CommandTestSuiteChannelStruct) {

	// Refresh tabs
	TestSuiteTabsRef.Refresh()
}
