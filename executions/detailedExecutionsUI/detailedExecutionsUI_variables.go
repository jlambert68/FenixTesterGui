package detailedTestCaseExecutionsUI

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// DetailedTestCaseExecutionsUIModelStruct - Structure holding all object and references need to create UI for TestCaseExecutions
type DetailedTestCaseExecutionsUIModelStruct struct {
	ExecutionsToolUIBar    *widget.Toolbar    // Toolbar used copy, cut, paste Building Blocks
	TestCaseExecutionsTabs *container.AppTabs // The Tab-structure where each TestCaseExecution has its own Tab
	//ExecutionsModelReference *executionsModelForSubscriptions.ExecutionsModelObjectStruct // A reference to the model for all TestExecutions

}

// ExecutionsUIObject - The object that
var DetailedTestCaseExecutionsUIObject DetailedTestCaseExecutionsUIModelStruct
