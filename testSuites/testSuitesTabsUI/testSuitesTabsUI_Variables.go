package testSuitesTabsUI

import (
	"FenixTesterGui/testSuites/testSuiteUI"
	"fyne.io/fyne/v2/container"
)

// TestSuiteUiMapPtr
// Map with all 'open' Tab-items for TestSuites
var TestSuiteUiMapPtr *map[*container.TabItem]*testSuiteUI.TestSuiteUiStruct

// The Home-tab-item for the SuiteTabs
var testSuiteHomeTabItem *container.TabItem
