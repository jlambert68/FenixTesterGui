package testSuiteUI

import (
	sharedCode "FenixTesterGui/common_code"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (testSuiteUiModel *TestSuiteUiStruct) generateLockOwnerDomainAndTestEnvironmentAreaContainer() (
	lockOwnerAndTestEnvironmentAreaContainer *fyne.Container,
	err error) {

	var lockButton *widget.Button
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
					lockButton.Disable()
					lockButton.SetText("Locked")

				} else {
					// user cancelled
				}
			},
			*sharedCode.FenixMasterWindowPtr,
		)

	})

	lockOwnerAndTestEnvironmentAreaContainer = container.NewVBox(
		widget.NewLabel("Lock Owner and TestEnvironment"),
		lockButton)

	return lockOwnerAndTestEnvironmentAreaContainer, nil
}
