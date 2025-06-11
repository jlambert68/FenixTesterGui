package testSuitesTabsUI

import (
	"FenixTesterGui/testSuites/testSuiteUI"
	"fyne.io/fyne/v2/container"
)

// TestSuiteTabs
// The Tab-object holding all TestSuiteTabs
var TestSuiteTabs *container.DocTabs

// TestSuiteUiMapPtr
// Map with all 'open' Tab-items for TestSuites
var TestSuiteUiMapPtr *map[*container.TabItem]*testSuiteUI.TestSuiteUiStruct

// The Home-tab-item for the SuiteTabs
var testSuiteHomeTabItem *container.TabItem
