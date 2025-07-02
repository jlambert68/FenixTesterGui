package testSuiteUI

import (
	"FenixTesterGui/testSuites/testSuitesModel"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// TestSuiteUiStruct
// Struct that holds all variables for one TestSuite-UI-object
type TestSuiteUiStruct struct {
	TestSuiteTabItem  *container.TabItem
	TestSuiteModelPtr *testSuitesModel.TestSuiteModelStruct

	// mouseHasLeftSideSplitContainer
	// Keep s track when mouse leaves or enters left side of the Split-container
	mouseHasLeftSideSplitContainer bool

	// UI components
	testSuiteInformationScroll                 *container.Scroll
	testSuiteInformationStackContainer         *fyne.Container
	testSuiteOwnerDomainContainer              *fyne.Container
	testCaseOwnerDomainCustomSelectComboBox    *customSelectComboBox
	testSuiteTestEnvironmentContainer          *fyne.Container
	customTestEnvironmentSelectComboBox        *customSelectComboBox
	testSuiteTestEnvironmentStackContainer     *fyne.Container
	lockOwnerAndTestEnvironmentAreaContainer   *fyne.Container
	testSuiteMetaDataContainer                 *fyne.Container
	testSuiteMetaDataStackContainer            *fyne.Container
	testSuiteTestDataAreaContainer             *fyne.Container
	testCaseListAccordionItemContainer         *fyne.Container
	preViewAndFilterTabsUsedForCreateTestSuite *container.AppTabs
	ownerDomainAndEnvironmentAccordion         *widget.Accordion
}

// The space between the different Information-boxes
const spaceBetweenInformationBoxes = "   "
