package listTestCasesUI

import (
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"sync"
)

// The UI-table for the List with TestCase
var testCaseListTable *widget.Table

// The data source used to produce the UI-table for the List with TestCase
var testCaseListTableTable [][]string

// Keeps the number of TestCase that is shown in the list, after local filter is applied
var numberOfTestCasesAfterLocalFilters binding.String

// Keeps the number of TestCase that have been retrieved from the Database
var numberOfTestCasesInTheDatabaseSearch binding.String

var testCaseListTableHeader = []string{
	"DomainName", "TestCaseName", "TestCaseUuid", "TestCaseVersion", "Latest TestCaseExecution Status",
	"Latest TestCaseExecution TimeStamp", "Latest OK Finished TestCaseExecution TimeStamp", "DomainUuid"}

// Keeps track of the in which column the TestCaseUUID exist in the data source for the UI-table
const testCaseUuidColumnNumber uint8 = 2

// The row that the mouse is currently hovering above. Used for highlight that row in the UI-Table
var currentRowThatMouseIsHoveringAbove int16 = -1

// Use a mutex to synchronize access to 'currentRowThatMouseIsHoveringAbove'
var currentRowThatMouseIsHoveringAboveMutex sync.Mutex
