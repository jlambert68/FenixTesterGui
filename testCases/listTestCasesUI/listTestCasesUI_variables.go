package listTestCasesUI

import (
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var testCaseListTable *widget.Table

var testCaseListTableTable [][]string

var numberOfTestCasesInTheSearch binding.String
