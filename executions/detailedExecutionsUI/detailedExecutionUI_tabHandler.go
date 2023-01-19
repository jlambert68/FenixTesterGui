package detailedTestCaseExecutionsUI

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// GenerateBaseUITabForExecutions
// Create the Base-UI-canvas-object for the Detailed TestCaseExecutions object.
// This base doesn't contain any specific TestCase-parts, and they will be added in other function
func (detailedTestCaseExecutionsUIObject *DetailedTestCaseExecutionsUIModelStruct) GenerateBaseUITabForDetailedTestCaseExecutions() (executionsCanvasObjectUI fyne.CanvasObject) {

	var testCaseExecutionsTabPage *fyne.Container
	// Create toolbar for Executions Area
	detailedTestCaseExecutionsUIObject.ExecutionsToolUIBar = widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentRedoIcon(), func() {
			fmt.Println("Reload Current Execution(s)")
			testCaseExecutionsTabPage = detailedTestCaseExecutionsUIObject.CreateDetailedTestCaseExecutionsTabPage()
			testCaseExecutionsTabPage.Refresh()
			detailedTestCaseExecutionsUIObject.TestCaseExecutionsTabs.Refresh()

		}),

		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
			fmt.Println("Show Executions in a read only undocked page")

		}),
	)

	// Generate TestCaseExecutions page
	//var testCaseExecutionsTabPage *fyne.Container
	testCaseExecutionsTabPage = detailedTestCaseExecutionsUIObject.CreateDetailedTestCaseExecutionsTabPage()

	// Create The Tab-object, where each TestCase will have its own Tab
	detailedTestCaseExecutionsUIObject.TestCaseExecutionsTabs = &container.DocTabs{
		BaseWidget:     widget.BaseWidget{},
		Items:          nil,
		CreateTab:      nil,
		CloseIntercept: nil,
		OnClosed:       nil,
		OnSelected:     nil,
		OnUnselected:   nil,
	}

	// Append TestCaseExecutions-Summary-page
	detailedTestCaseExecutionsUIObject.TestCaseExecutionsTabs.Append(&container.TabItem{
		Text:    "Detailed TestCase Executions Summary",
		Icon:    theme.HomeIcon(),
		Content: testCaseExecutionsTabPage,
	})

	detailedTestCaseExecutionsUIObject.TestCaseExecutionsTabs.OnSelected = func(tabItem *container.TabItem) {
		fmt.Println("OnSelected")
		fmt.Println(tabItem)
		testCaseExecutionsTabPage.Refresh()

	}

	detailedTestCaseExecutionsUIObject.TestCaseExecutionsTabs.CloseIntercept = func(tabItem *container.TabItem) {
		if tabItem.Text == "Detailed TestCase Executions Summary" {

			tabItem.Content = detailedTestCaseExecutionsUIObject.CreateDetailedTestCaseExecutionsTabPage()
			// tabItem.Content.Refresh()
		}

	}

	// Set the Tabs to be positioned in upper part of the object
	detailedTestCaseExecutionsUIObject.TestCaseExecutionsTabs.SetTabLocation(container.TabLocationTop)

	detailedTestCaseExecutionsUIObject.TestCaseExecutionsTabs.Refresh()

	// Create the complete Executions UI area
	var executionsBorderedLayout fyne.Layout
	executionsBorderedLayout = layout.NewBorderLayout(detailedTestCaseExecutionsUIObject.ExecutionsToolUIBar, nil, nil, nil)
	executionsCanvasObjectUI = container.New(executionsBorderedLayout, detailedTestCaseExecutionsUIObject.ExecutionsToolUIBar, detailedTestCaseExecutionsUIObject.TestCaseExecutionsTabs)

	return executionsCanvasObjectUI
}
