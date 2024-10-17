package templateViewer

import (
	"FenixTesterGui/importFilesFromGitHub"
	"FenixTesterGui/soundEngine"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/jlambert68/FenixScriptEngine/placeholderReplacementEngine"
	"github.com/jlambert68/FenixScriptEngine/testDataEngine"
	"strings"
)

func InitiateTemplateViewer(
	mainWindow fyne.Window,
	myApp fyne.App,
	importedFilesPtr *[]importFilesFromGitHub.GitHubFile,
	testDataForGroupObject *testDataEngine.TestDataForGroupObjectStruct,
	randomUuidForScriptEngine string,
	choseTemplateName string,
	testDataPointGroupName string,
	testDataPointName string,
	testDataRowName string) {

	// Disable the main window
	mainWindow.Hide()

	// The Select-items for Groups ans TestDataPoints for a Group
	var testDataPointGroupsSelect *widget.Select
	var testDataPointGroupsSelectSelected string
	var testDataPointsForAGroupSelect *widget.Select
	var testDataPointForAGroupSelectSelected string
	var testDataRowsForTestDataPointsSelect *widget.Select
	var testDataRowForTestDataPointsSelectSelected string

	// The slices for Groups, TestDataPoints for a Group and the specific TestDataRows for a TestDataPoint
	var testDataPointGroups []string
	var testDataPointsForAGroup []string
	var testDataRowsForATestDataPoint []string

	// Store reference to Fenix Main Window
	fenixMainWindow = mainWindow

	// Create the window for GitHub files
	templateViewerWindow = myApp.NewWindow("Imported Files Viewer")
	// Set initial size of the window
	templateViewerWindow.Resize(fyne.NewSize(800, 700))
	templateViewerWindow.CenterOnScreen()

	var leftContainer *fyne.Container
	var rightContainer *fyne.Container

	// Extract filenames for the fileSelectorDropdown
	var fileNames []string
	for _, file := range *importedFilesPtr {
		fileNames = append(fileNames, file.Name)
	}

	// Create UI component for 'fileSelectorDropdown'
	fileSelectorDropdown := widget.NewSelect(fileNames, nil)
	urlLabel := widget.NewLabel("")
	var richText *widget.RichText
	richText = &widget.RichText{
		BaseWidget: widget.BaseWidget{},
		Segments:   nil,
		Wrapping:   0,
		Scroll:     0,
		Truncation: 0,
	}
	var richTextWithValues *widget.RichText
	richTextWithValues = &widget.RichText{
		BaseWidget: widget.BaseWidget{},
		Segments:   nil,
		Wrapping:   0,
		Scroll:     0,
		Truncation: 0,
	}

	var testDataPointValues map[string]string // map[TestDataColumnDataNameType]TestDataValueType

	// Set the fileSelectorDropdown change handler
	var selectedFile string
	fileSelectorDropdown.OnChanged = func(selected string) {

		selectedFile = selected

		testDataPointValues, _, _, _, _, _ = testDataForGroupObject.
			GetTestDataPointValuesMapBasedOnGroupPointNameAndSummaryValue(
				testDataPointGroupsSelectSelected,
				testDataPointForAGroupSelectSelected,
				testDataRowForTestDataPointsSelectSelected)

		for _, file := range *importedFilesPtr {
			if file.Name == selected {
				urlLabel.SetText(file.URL)

				myContainerObjects := leftContainer.Objects
				for index, object := range myContainerObjects {
					if object == richText {
						//richText, _, _ = parseAndFormatText(file.FileContentAsString, &testDataPointValues, randomUuidForScriptEngine)
						richText, _, _ = placeholderReplacementEngine.ParseAndFormatPlaceholders(file.FileContentAsString, &testDataPointValues, randomUuidForScriptEngine)
						myContainerObjects[index] = richText
						leftContainer.Refresh()
					}
				}

				myContainerObjects = rightContainer.Objects
				for index, object := range myContainerObjects {
					if object == richTextWithValues {
						//_, richTextWithValues, _ = parseAndFormatText(file.FileContentAsString, &testDataPointValues, randomUuidForScriptEngine)
						_, richTextWithValues, _ = placeholderReplacementEngine.ParseAndFormatPlaceholders(file.FileContentAsString, &testDataPointValues, randomUuidForScriptEngine)
						myContainerObjects[index] = richTextWithValues
						rightContainer.Refresh()
					}
				}

				break
			}
		}
	}

	// Create function that converts a GroupSlice into a string slice
	getTestGroupsFromTestDataEngineFunction := func() []string {

		testDataPointGroups = testDataForGroupObject.ListTestDataGroups()

		return testDataPointGroups
	}

	// Create function that converts a TestDataPointsSlice into a string slice
	testDataPointsToStringSliceFunction := func(testDataGroup string) []string {

		if testDataGroup == "" {
			return []string{}
		}

		testDataPointsForAGroup = testDataForGroupObject.ListTestDataGroupPointsForAGroup(testDataGroup)

		return testDataPointsForAGroup
	}

	// Create function that converts a slice with the specific TestDataPoints into a string slice
	testDataRowSliceToStringSliceFunction := func(testDataGroup string, testDataGroupPoint string) []string {

		if testDataGroup == "" || testDataGroupPoint == "" {
			return []string{}
		}

		testDataRowsForATestDataPoint = testDataForGroupObject.ListTestDataRowsForAGroupPoint(testDataGroup, testDataGroupPoint)

		return testDataRowsForATestDataPoint
	}

	// Create the Group dropdown - <Name of the group>
	testDataPointGroupsSelect = widget.NewSelect(getTestGroupsFromTestDataEngineFunction(), func(selected string) {

		testDataPointGroupsSelectSelected = selected

		// Select the correct TestDataPoint in the dropdown for TestDataPoints
		testDataPointsForAGroupSelect.SetOptions(testDataPointsToStringSliceFunction(selected))
		testDataPointsForAGroupSelect.Refresh()

		// UnSelect in DropDown- and List for TestDataPoints
		testDataPointsForAGroupSelect.ClearSelected()

	})

	// Create the Groups TestDataPoints dropdown - <Sub Custody/Main TestData Area/SEK/AccTest/SE/CRDT/CH/Switzerland/BBH/EUR/EUR/SEK>
	testDataPointsForAGroupSelect = widget.NewSelect(
		testDataPointsToStringSliceFunction(testDataPointGroupsSelectSelected), func(selected string) {

			testDataPointForAGroupSelectSelected = selected

			// Select the correct TestDataPoint in the dropdown for TestDataPoints
			testDataRowsForTestDataPointsSelect.SetOptions(
				testDataRowSliceToStringSliceFunction(testDataPointGroupsSelect.Selected, selected))
			testDataRowsForTestDataPointsSelect.Refresh()

			// UnSelect in DropDown- and List for Specific TestDataPoints
			testDataRowsForTestDataPointsSelect.ClearSelected()

		})

	// Create the Groups Specific TestDataPoint dropdown - <All the specific values>
	testDataRowsForTestDataPointsSelect = widget.NewSelect(
		testDataRowSliceToStringSliceFunction(testDataPointGroupsSelectSelected, testDataPointForAGroupSelectSelected),
		func(selected string) {

			testDataRowForTestDataPointsSelectSelected = selected

			fileSelectorDropdown.SetSelected(selectedFile)

		})

	// Create UI component for 'TestDataGroupPointSelector'

	topContainer := container.NewVBox(
		fileSelectorDropdown,
		urlLabel,
		testDataPointGroupsSelect,
		testDataPointsForAGroupSelect,
		testDataRowsForTestDataPointsSelect)

	// Button for copy the "Template"
	var copyTemplateTextButton *widget.Button
	copyTemplateTextButton = widget.NewButton("Copy Template", func() {
		var templateText string
		templateText = getTextFromRichText(richText)

		clipboard := mainWindow.Clipboard()
		clipboard.SetContent(templateText)

		// Notify the user

		// Trigger System Notification sound
		soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "Clipboard",
			Content: "Template copied to clipboard!",
		})

	})

	// Button for copy the "Template with values"
	var copyTemplateWithValuesButton *widget.Button
	copyTemplateWithValuesButton = widget.NewButton("Copy Template", func() {
		var templateText string
		templateText = getTextFromRichText(richTextWithValues)

		clipboard := mainWindow.Clipboard()
		clipboard.SetContent(templateText)

		// Notify the user

		// Trigger System Notification sound
		soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "Clipboard",
			Content: "Template - with values- copied to clipboard!",
		})

	})

	// Placeholder for rightContainer - add your form view here
	rightContainer = container.NewBorder(copyTemplateWithValuesButton, nil, nil, nil, richTextWithValues)

	leftContainer = container.NewBorder(copyTemplateTextButton, nil, nil, nil, richText)

	// Set File if that was included in call
	if choseTemplateName != "" {
		fileSelectorDropdown.SetSelected(choseTemplateName)
	}

	// Set TestDataGroup, TestDataPoint and TestDataRow
	testDataPointGroupsSelect.SetSelected(testDataPointGroupName)
	testDataPointsForAGroupSelect.SetSelected(testDataPointName)
	testDataRowsForTestDataPointsSelect.SetSelected(testDataRowName)

	// Create split container
	split := container.NewHSplit(leftContainer, rightContainer)
	split.Offset = 0.5 // Adjust as needed

	fullContentContainer := container.NewBorder(topContainer, nil, nil, nil, split)

	templateViewerWindow.SetContent(fullContentContainer)

	// Set the callback function for window close event to show the Main window again
	templateViewerWindow.SetOnClosed(func() {
		fenixMainWindow.Show()
	})

	// Show the File Viewe Window
	templateViewerWindow.Show()
}

// Function to extract text from RichText
func getTextFromRichText(richText *widget.RichText) string {
	var sb strings.Builder

	// Iterate over each segment in the RichText
	for _, seg := range richText.Segments {
		switch s := seg.(type) {
		case *widget.TextSegment:
			sb.WriteString(s.Text)
		case *widget.HyperlinkSegment:
			sb.WriteString(s.Text)
			// Add cases for other segment types if needed
		}
	}

	return sb.String()
}
