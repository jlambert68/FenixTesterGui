package executionsUIForSubscriptions

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// GenerateBaseUITabForExecutions
// Create the Base-UI-canvas-object for the TestCases object. This base doesn't contain any specific TestCase-parts, and they will be added in other function
func (executionsUIObject *ExecutionsUIModelStruct) GenerateBaseUITabForExecutions() (executionsCanvasObjectUI fyne.CanvasObject) {

	// Create toolbar for Executions Area
	executionsUIObject.ExecutionsToolUIBar = widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentRedoIcon(), func() {
			fmt.Println("Reload Current Execution(s)")

		}),

		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
			fmt.Println("Show Executions in a read only undocked page")

		}),
	)

	// Generate TestCaseExecutions page
	var testCaseExecutionsTabPage *fyne.Container
	testCaseExecutionsTabPage = executionsUIObject.CreateExecutionsListTabPageForSubsacriptions()

	/*
		// Create The Tab-object, where each TestCase will have its own Tab
		executionsUIObject.ExecutionsTabs = container.NewAppTabs(
			container.NewTabItemWithIcon("TestCase Executions", theme.HomeIcon(), testCaseExecutionsTabPage),
			container.NewTabItemWithIcon("TestSuite Executions", theme.HomeIcon(), widget.NewLabel("Home tab")),
		)
	*/

	// Create The Tab-object, where each TestCase will have its own Tab
	executionsUIObject.ExecutionsTabs = &container.AppTabs{
		Items:     nil,
		OnChanged: nil,
		OnSelected: func(tabItem *container.TabItem) {
			//tabItem.Content.Refresh()
			executionsUIObject.OnQueueTable.Header.ScrollToTrailing()
			tabItem.Content.Refresh()
		},
		OnUnselected: nil,
	}

	// Append TestCaseExecutions-List-page
	executionsUIObject.ExecutionsTabs.Append(&container.TabItem{
		Text:    "TestCase Executions",
		Icon:    theme.HomeIcon(),
		Content: testCaseExecutionsTabPage,
	})

	// Append TestSuiteExecutions-List-page
	executionsUIObject.ExecutionsTabs.Append(&container.TabItem{
		Text:    "TestSuite Executions",
		Icon:    theme.HomeIcon(),
		Content: widget.NewLabel("Tab for 'TestSuite Executions'"),
	})

	executionsUIObject.ExecutionsTabs.OnSelected = func(tabItem *container.TabItem) {
		fmt.Println(tabItem)
		//executionsUIObject.OnQueueTable.Header.ScrollToTrailing()
		executionsUIObject.OnQueueTable.Header.ScrollToLeading()
	}

	// Set the Tabs to be positioned in upper part of the object
	executionsUIObject.ExecutionsTabs.SetTabLocation(container.TabLocationTop)

	executionsUIObject.ExecutionsTabs.Refresh()

	// Create the complete Executions UI area
	var executionsBorderedLayout fyne.Layout
	executionsBorderedLayout = layout.NewBorderLayout(executionsUIObject.ExecutionsToolUIBar, nil, nil, nil)
	executionsCanvasObjectUI = container.New(executionsBorderedLayout, executionsUIObject.ExecutionsToolUIBar, executionsUIObject.ExecutionsTabs)

	// Initiate map for ExecutionsUI-models-Map
	//executionsUIObject.TestCasesUiModelMap = make(map[string]*testCaseGraphicalAreasStruct)

	return executionsCanvasObjectUI
}
