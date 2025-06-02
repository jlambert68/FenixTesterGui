package listTestCasesUI

import (
	"FenixTesterGui/testCase/testCaseModel"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// GenerateTestCaseMetaDataFilterContainer
// Generates the GenerateTestCaseMetaDataFilterContainer containing a simple and an advanced filter version
func GenerateTestCaseMetaDataFilterContainer(
	testCasesModel *testCaseModel.TestCasesModelsStruct) *container.AppTabs {

	/*
		var testCaseMetaDataFilterTopContainer *fyne.Container
		var testCaseMetaDataFilterBottomContainer *fyne.Container

		var testCaseMainAreaForMetaDataFilterContainer *fyne.Container


	*/
	var metaDataFilterTabs *container.AppTabs
	var simpleMetaDataFilterTab *container.TabItem
	var advancedMetaDataFilterTab *container.TabItem

	// Generate the Tab for the Simple MetaDataFilter
	simpleMetaDataFilterTab = container.NewTabItem(
		"Simple MetaDataFilter",
		generateSimpleTestCaseMetaDataFilterContainer(testCasesModel))

	// Generate Tab for the Advanced MetaDataFilter
	advancedMetaDataFilterTab = container.NewTabItem(
		"Advanced MetaDataFilter",
		widget.NewLabel("Advanced MetaDataFilter"))

	// Generate the AppTabsContainer
	metaDataFilterTabs = container.NewAppTabs(simpleMetaDataFilterTab, advancedMetaDataFilterTab)

	return metaDataFilterTabs

}
