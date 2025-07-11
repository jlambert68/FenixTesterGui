package testDataSelectorForTestSuite

import (
	"fyne.io/fyne/v2"
	"github.com/jlambert68/FenixScriptEngine/testDataEngine"
)

// Updates the list that show the TestDataPoints for a specific Group in main window
func updateTestDataPointsForAGroupList(
	testDataPointGroupName testDataEngine.TestDataPointGroupNameType,
	testDataForGroupObject *testDataEngine.TestDataForGroupObjectStruct) {

	// Clear the slice that holds all TestDataPoints
	testDataForGroupObject.TestDataPointsForAGroup = nil

	// Extract the map with the TestDataPoints
	var tempTestDataPointNameMap testDataEngine.TestDataPointNameMapType
	tempTestDataPointNameMap = *testDataForGroupObject.ChosenTestDataPointsPerGroupMap[testDataPointGroupName]

	// Refill the slice with all TestDataPoints
	for testDataValueName, _ := range tempTestDataPointNameMap {

		testDataForGroupObject.TestDataPointsForAGroup = append(testDataForGroupObject.TestDataPointsForAGroup, testDataValueName)

	}

	// Refresh the List in the UI
	fyne.Do(func() {
		testDataPointsForAGroupList.Refresh()
	})

}
