package testSuiteUI

import (
	"FenixTesterGui/testSuites/testSuitesCommandEngine"
	"fyne.io/fyne/v2"
)

// Lock or Unlock UI depending on if user has selected OwnerDomain and TestEnvironment
func (testSuiteUiModel *TestSuiteUiStruct) lockUIUntilOwnerDomainAndTestEnvironmenIsSelected() {

	var lockedHasBeenClickedAndOwnerDomainAndTestEnvironmentIsSelected bool
	var lockedButtonShoudlBeVisible bool

	// Check if  user has selected OwnerDomain and TestEnvironment
	lockedHasBeenClickedAndOwnerDomainAndTestEnvironmentIsSelected = testSuiteUiModel.TestSuiteModelPtr.
		HasLockButtonBeenClickedAndBothOwnerDomainAndTestEnvironmentHaveValues()

	if lockedHasBeenClickedAndOwnerDomainAndTestEnvironmentIsSelected == true {
		// Show Containers
		testSuiteUiModel.testSuiteMetaDataStackContainer.Show()
		testSuiteUiModel.testSuiteTestDataAreaContainer.Show()
		testSuiteUiModel.testCaseListAccordionItemContainer.Show()
	} else {
		// Hide Containers
		testSuiteUiModel.testSuiteMetaDataStackContainer.Hide()
		testSuiteUiModel.testSuiteTestDataAreaContainer.Hide()
		testSuiteUiModel.testCaseListAccordionItemContainer.Hide()
	}

	// Check if  lockButton should be visible
	lockedButtonShoudlBeVisible = testSuiteUiModel.TestSuiteModelPtr.
		DoBothOwnerDomainAndTestEnvironmentHaveValues()

	if lockedButtonShoudlBeVisible == true {
		// Show Containers
		testSuiteUiModel.lockOwnerAndTestEnvironmentAreaContainer.Show()
	} else {
		// Hide Containers
		testSuiteUiModel.lockOwnerAndTestEnvironmentAreaContainer.Hide()
	}

	// Refresh Tabs
	fyne.Do(func() {
		testSuitesCommandEngine.TestSuiteTabsRef.Refresh()
	})

}
