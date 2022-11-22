package executionsUI_tabHandler

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// GenerateBaseCanvasObjectForTestCaseUI
// Create the Base-UI-canvas-object for the TestCases object. This base doesn't contain any specific TestCase-parts, and they will be added in other function
func (executionsUIObject *executionsUI.ExecutionsUIModelStruct) GenerateBaseCanvasObjectForTestCaseUI() (baseCanvasObjectForTestCaseUI fyne.CanvasObject) {

	// Create toolbar for TestCase area
	executionsUIObject.TestCaseToolUIBar = widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentRedoIcon(), func() {
			fmt.Println("Reload GUI TestCase from testCaseModel")
		}),

		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
			fmt.Println("Copy Node")
		}),

		widget.NewToolbarAction(theme.ContentCutIcon(), func() {
			fmt.Println("Cut Node")
		}),

		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {
			fmt.Println("Past Node")
		}),
	)

	// Create The Tab-object, where each TestCase will have its own Tab
	executionsUIObject.TestCasesTabs = container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")),
	)

	// Set the Tabs to be positioned in upper part of the object
	executionsUIObject.TestCasesTabs.SetTabLocation(container.TabLocationTop)

	// Create the complete TestCase UI area
	testCaseBorderedLayout := layout.NewBorderLayout(executionsUIObject.TestCaseToolUIBar, nil, nil, nil)
	baseCanvasObjectForTestCaseUI = container.New(testCaseBorderedLayout, executionsUIObject.TestCaseToolUIBar, executionsUIObject.TestCasesTabs)

	// Initiate map with TestCaseUI-models-Map
	executionsUIObject.TestCasesUiModelMap = make(map[string]*testCaseGraphicalAreasStruct)

	return baseCanvasObjectForTestCaseUI
}
