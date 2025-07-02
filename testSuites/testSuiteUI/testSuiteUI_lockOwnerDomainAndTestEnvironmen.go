package testSuiteUI

import (
	sharedCode "FenixTesterGui/common_code"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (testSuiteUiModel *TestSuiteUiStruct) generateLockOwnerDomainAndTestEnvironmentAreaContainer() (
	lockOwnerAndTestEnvironmentAreaContainer *fyne.Container,
	err error) {

	var lockButton *widget.Button
	var lockButtonFunction func()

	lockButtonFunction = func() {
		lockButton.Disable()
		lockButton.SetText("Locked")
		testSuiteUiModel.TestSuiteModelPtr.LockButtonHasBeenClicked()

		// Open up correct containers and
		testSuiteUiModel.lockUIUntilOwnerDomainAndTestEnvironmenIsSelected()

		// Lock down Selected OwnerDomain and TestEnvironment
		testSuiteUiModel.testCaseOwnerDomainCustomSelectComboBox.selectComboBox.Disable()
		if testSuiteUiModel.customTestEnvironmentSelectComboBox != nil {
			testSuiteUiModel.customTestEnvironmentSelectComboBox.selectComboBox.Disable()
		}

		// Close accordion for this part
		testSuiteUiModel.ownerDomainAndEnvironmentAccordion.CloseAll()
	}

	lockButton = widget.NewButton("Lock", func() {

		// Create popup asking user if TestSuite's Owner Domain and TestEnvironment should be locked
		question := widget.NewLabel("Should 'Owner Domain' and 'TestEnvironment' be locked?")
		dialog.ShowCustomConfirm(
			"Lock",
			"Lock",
			"Cancel",
			question,
			func(lock bool) {
				if lock {
					// Lock
					lockButtonFunction()

				} else {
					// user cancelled
				}
			},
			*sharedCode.FenixMasterWindowPtr,
		)
	})

	// Check if Lock should be applied
	if testSuiteUiModel.TestSuiteModelPtr.HasLockButtonBeenClickedAndBothOwnerDomainAndTestEnvironmentHaveValues() == true {
		lockButtonFunction()
	}

	lockOwnerAndTestEnvironmentAreaContainer = container.NewVBox(
		widget.NewLabel("Lock Owner and TestEnvironment"),
		container.NewHBox(lockButton, layout.NewSpacer()))

	return lockOwnerAndTestEnvironmentAreaContainer, nil
}
