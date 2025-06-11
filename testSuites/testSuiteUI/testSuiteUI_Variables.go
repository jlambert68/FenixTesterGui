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
