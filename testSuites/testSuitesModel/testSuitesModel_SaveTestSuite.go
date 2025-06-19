package testSuitesModel

import (
	"FenixTesterGui/soundEngine"
	"fmt"
	"fyne.io/fyne/v2"
)

// SaveTestSuite
// Send TestSuite to TestCaseBuilderServer to saved in the Database
func (testSuiteModel *TestSuiteModelStruct) SaveTestSuite() (err error) {

	var isTestSuiteChanged bool
	var mandatoryFieldsHaveValues bool
	var mandatoryFieldsHaveValuesNotificationText string

	// Check if TestSuite is Changed
	isTestSuiteChanged = testSuiteModel.IsTestSuiteChanged()

	// Inform user that TestSuite is not changed
	if isTestSuiteChanged == false {

		// Trigger System Notification sound
		soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "Save TestSuite",
			Content: fmt.Sprintf("TestSuite is not changed"),
		})

		return nil
	}

	// Check if all mandatory fields has values
	// Check if TestSuite is Changed
	mandatoryFieldsHaveValues,
		mandatoryFieldsHaveValuesNotificationText = testSuiteModel.checkIfAllMandatoryFieldsHaveValues()

	// Inform user that mandatory field in TestSuite is missing
	if mandatoryFieldsHaveValues == false {

		// Trigger System Notification sound
		soundEngine.PlaySoundChannel <- soundEngine.UserNeedToRespondSound

		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title: "Mandatory Fields",
			Content: fmt.Sprintf("Mandatory field '%s' is empty",
				mandatoryFieldsHaveValuesNotificationText),
		})

		return nil
	}

	return err

}
