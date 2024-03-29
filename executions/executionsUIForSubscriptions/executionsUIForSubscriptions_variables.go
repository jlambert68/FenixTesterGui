package executionsUIForSubscriptions

import (
	"FenixTesterGui/executions/executionsModelForSubscriptions"
	"FenixTesterGui/headertable"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ExecutionsUIModelStruct - Structure holding all object and references need to create UI for Executions
type ExecutionsUIModelStruct struct {
	ExecutionsToolUIBar *widget.Toolbar    // Toolbar used copy, cut, paste Building Blocks
	ExecutionsTabs      *container.AppTabs // The Tab-structure where each TestExecution has its own Tab
	//TestCasesUiModelMap     map[string]*testCaseGraphicalAreasStruct // Holds all UI sub-parts for a TestCase
	ExecutionsModelReference *executionsModelForSubscriptions.ExecutionsModelObjectStruct // A reference to the model for all TestExecutions
	OnQueueTable             *headertable.SortingHeaderTable
	UnderExecutionTable      *headertable.SortingHeaderTable
	FinishedExecutionTable   *headertable.SortingHeaderTable
}

// ExecutionsUIObject - The object that
var ExecutionsUIObject ExecutionsUIModelStruct
