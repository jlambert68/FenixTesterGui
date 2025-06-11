package testSuitesTabsUI

import (
	"FenixTesterGui/testSuites/testSuiteUI"
	"fyne.io/fyne/v2/container"
)

// TestSuiteTabs
// The Tab-object holding all TestSuiteTabs
var TestSuiteTabs *container.DocTabs

var TestSuiteUiMapPtr *map[*container.TabItem]*testSuiteUI.TestSuiteUiStruct
