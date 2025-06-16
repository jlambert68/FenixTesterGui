package testSuiteUI

import (
	"FenixTesterGui/testSuites/testSuitesModel"
	"fyne.io/fyne/v2/container"
)

// TestSuiteUiStruct
// Struct that holds all variables for one TestSuite-UI-object
type TestSuiteUiStruct struct {
	TestSuiteTabItem  *container.TabItem
	TestSuiteModelPtr *testSuitesModel.TestSuiteModelStruct

	// UI components
}

// mouseHasLeftSideSplitContainer
// Keep s track when mouse leaves or enters left side of the Split-container
var mouseHasLeftSideSplitContainer bool

// The space between the different Information-boxes
const spaceBetweenInformationBoxes = "           "
