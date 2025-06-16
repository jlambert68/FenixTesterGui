package testSuiteUI

import (
	"fyne.io/fyne/v2/widget"
)

// The Select-items for Groups and TestDataPoints for a Group
var testDataPointGroupsSelect *widget.Select
var testDataPointGroupsSelectSelectedInMainTestSuiteArea string
var testDataPointsForAGroupSelect *widget.Select
var testDataPointForAGroupSelectSelectedInMainTestSuiteArea string
var testDataRowsForTestDataPointsSelect *widget.Select
var testDataRowForTestDataPointsSelectSelectedInMainTestSuiteArea string
var testDataAsRichText *widget.RichText

var showTestDataPointRows bool
