package testDataSelectorForTestSuite

import (
	"FenixTesterGui/testDataSelector/newOrEditTestDataPointGroupUI"
	"FenixTesterGui/testSuites/testSuitesModel"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/jlambert68/FenixScriptEngine/testDataEngine"
)

func MainTestDataSelector(
	app fyne.App,
	parent fyne.Window,
	currentTestSuitePtr *testSuitesModel.TestSuiteModelStruct,
	testCaseUuid string,
	testDataSelectorsContainer *fyne.Container,
	testDataPointGroupsSelectInMainTestSuiteArea *widget.Select,
	testDataPointsForAGroupSelectInMainTestSuiteArea *widget.Select,
	testDataRowsForTestDataPointsSelectInMainTestSuiteArea *widget.Select) {

	parent.Hide()

	myWindow := app.NewWindow("TestData Management")
	myWindow.Resize(fyne.NewSize(600, 500))
	myWindow.CenterOnScreen()

	// When this window closed then show parent and send response to parent window
	myWindow.SetOnClosed(func() {
		parent.Show()
	})

	// Initiate TestData if it isn't already done
	if currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr == nil {
		currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr = &testDataEngine.TestDataForGroupObjectStruct{}
	}

	// Initiate 'chosenTestDataPointsPerGroupMap'
	if currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.ChosenTestDataPointsPerGroupMap == nil {
		currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.ChosenTestDataPointsPerGroupMap = make(map[testDataEngine.TestDataPointGroupNameType]*testDataEngine.TestDataPointNameMapType)
	}

	// Create List UI for 'testDataPointGroups'
	testDataPointGroupsList = widget.NewList(
		func() int { return len(currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.TestDataPointGroups) },
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			fyne.Do(func() {
				obj.(*widget.Label).SetText(string(currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.TestDataPointGroups[id]))
			})
		},
	)

	testDataPointGroupsList.OnSelected = func(id widget.ListItemID) {
		newOrEditTestDataPointGroupUI.SelectedIndexForGroups = id

		// Update List for  'testDataPointsForAGroup'
		updateTestDataPointsForAGroupList(currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.TestDataPointGroups[id], currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr)

		// Select correct Group in Select-dropdown
		newOrEditTestDataPointGroupUI.TestDataPointGroupsSelect.SetSelected(string(currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.TestDataPointGroups[id]))

		// UnSelect in DropDown- and List for TestDataPoints
		fyne.Do(func() {
			newOrEditTestDataPointGroupUI.TestDataPointsForAGroupSelect.ClearSelected()
		})
		testDataPointsForAGroupList.UnselectAll()
		newOrEditTestDataPointGroupUI.SelectedIndexForGroupTestDataPoints = -1

	}

	// Create function that converts a GroupSlice into a string slice
	testDataPointGroupsToStringSliceFunction := func() []string {
		var tempStringSlice []string

		for _, testDataPointGroup := range currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.TestDataPointGroups {
			tempStringSlice = append(tempStringSlice, string(testDataPointGroup))
		}

		return tempStringSlice
	}

	// Create function that converts a TestDataPointsSlice into a string slice
	testDataPointsToStringSliceFunction := func() []string {
		var tempStringSlice []string

		for _, testDataPointForAGroup := range currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.TestDataPointsForAGroup {
			tempStringSlice = append(tempStringSlice, string(testDataPointForAGroup))
		}

		return tempStringSlice
	}

	// Create the Group dropdown
	newOrEditTestDataPointGroupUI.TestDataPointGroupsSelect = widget.NewSelect(testDataPointGroupsToStringSliceFunction(), func(selected string) {

		// Find List-item to select
		for index, group := range currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.TestDataPointGroups {
			if string(group) == selected {
				newOrEditTestDataPointGroupUI.SelectedIndexForGroups = index

			}
		}

		// Select the correct TestDataPoint in the dropdown for TestDataPoints
		fyne.Do(func() {
			newOrEditTestDataPointGroupUI.TestDataPointsForAGroupSelect.SetOptions(testDataPointsToStringSliceFunction())
			newOrEditTestDataPointGroupUI.TestDataPointsForAGroupSelect.Refresh()
		})

		// Select the correct item in the Groups-List
		testDataPointGroupsList.Select(newOrEditTestDataPointGroupUI.SelectedIndexForGroups)
		fyne.Do(func() {
			testDataPointGroupsList.Refresh()
		})

		// UnSelect in DropDown- and List for TestDataPoints
		fyne.Do(func() {
			newOrEditTestDataPointGroupUI.TestDataPointsForAGroupSelect.ClearSelected()
		})
		testDataPointsForAGroupList.UnselectAll()
		newOrEditTestDataPointGroupUI.SelectedIndexForGroupTestDataPoints = -1

	})

	// Create the Groups TestDataPoints dropdown
	newOrEditTestDataPointGroupUI.TestDataPointsForAGroupSelect = widget.NewSelect(testDataPointsToStringSliceFunction(), func(selected string) {

		// Find List-item to select
		for index, group := range currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.TestDataPointsForAGroup {
			if string(group) == selected {
				newOrEditTestDataPointGroupUI.SelectedIndexForGroupTestDataPoints = index

			}
		}

		// Select the correct item in the TestDataPoints-List
		testDataPointsForAGroupList.Select(newOrEditTestDataPointGroupUI.SelectedIndexForGroupTestDataPoints)
		fyne.Do(func() {
			testDataPointsForAGroupList.Refresh()
		})
	})

	// Create List UI for 'testDataPointsForAGroup'
	testDataPointsForAGroupList = widget.NewList(
		func() int {
			return len(currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.TestDataPointsForAGroup)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(string(currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.TestDataPointsForAGroup[id]))
		},
	)

	testDataPointsForAGroupList.OnSelected = func(id widget.ListItemID) {
		newOrEditTestDataPointGroupUI.SelectedIndexForGroupTestDataPoints = id

		// Select correct Group in Select-dropdown
		newOrEditTestDataPointGroupUI.TestDataPointsForAGroupSelect.SetSelected(string(currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.TestDataPointsForAGroup[id]))

		// Select the correct TestDataPoint in the dropdown for TestDataPoints
		newOrEditTestDataPointGroupUI.TestDataPointsForAGroupSelect.SetOptions(testDataPointsToStringSliceFunction())
		newOrEditTestDataPointGroupUI.TestDataPointsForAGroupSelect.Refresh()

	}

	var testDataPointGroupsContainer *fyne.Container
	testDataPointGroupsContainer = container.NewBorder(newOrEditTestDataPointGroupUI.TestDataPointGroupsSelect,
		nil, nil, nil, testDataPointGroupsList)

	var testDataPointsForAGroupContainer *fyne.Container
	testDataPointsForAGroupContainer = container.NewBorder(newOrEditTestDataPointGroupUI.TestDataPointsForAGroupSelect,
		nil, nil, nil, testDataPointsForAGroupList)

	// Create Split Container used for 'testDataPointGroups' and 'testDataPointsForAGroup'
	var testDataGroupsAndPointsContainer *container.Split
	testDataGroupsAndPointsContainer = container.NewHSplit(testDataPointGroupsContainer, testDataPointsForAGroupContainer)

	var responseChannel chan testDataEngine.ResponseChannelStruct
	responseChannel = make(chan testDataEngine.ResponseChannelStruct)

	// The structure holding Group and TestDataPoints together
	//var newOrEditedChosenTestDataPointsPerGroupMap map[testDataEngine.TestDataPointGroupNameType][]TestDataPointRowUuidType
	//newOrEditedChosenTestDataPointsPerGroupMap = make(map[testDataEngine.TestDataPointGroupNameType][]TestDataPointRowUuidType)

	// Crete the 'New'-button for creating a new Group for TestDataPoints
	newButton := widget.NewButton("New", func() {
		myWindow.Hide()
		newOrEditTestDataPointGroupUI.ShowNewOrEditGroupWindow(
			app,
			myWindow,
			true,
			&responseChannel,
			"",
			&currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.ChosenTestDataPointsPerGroupMap,
			currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr)
	})

	// Crete the 'Edit'-button for editing an existing Group for TestDataPoints
	editButton := widget.NewButton("Edit", func() {
		if newOrEditTestDataPointGroupUI.SelectedIndexForGroups == -1 || len(currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.TestDataPointGroups) == 0 {
			dialog.ShowInformation("Error", "No selection made", myWindow)
			return
		}
		myWindow.Hide()

		newOrEditTestDataPointGroupUI.ShowNewOrEditGroupWindow(
			app,
			myWindow,
			false,
			&responseChannel,
			currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.TestDataPointGroups[newOrEditTestDataPointGroupUI.SelectedIndexForGroups],
			&currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.ChosenTestDataPointsPerGroupMap,
			currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr)
	})

	// Crete the 'Delete'-button for deleting an existing Group for TestDataPoints
	deleteButton := widget.NewButton("Delete", func() {
		if newOrEditTestDataPointGroupUI.SelectedIndexForGroups == -1 || len(currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.TestDataPointGroups) == 0 {
			dialog.ShowInformation("Error", "No selection made", myWindow)
			return
		}

		dialog.ShowConfirm("Confirm to Delete", fmt.Sprintf("Are you sure that you what to delete TestDataPointGroup '%s'?",
			currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.TestDataPointGroups[newOrEditTestDataPointGroupUI.SelectedIndexForGroups]), func(confirm bool) {
			if confirm {

				// Get the GroupName from the List to be deleted
				var groupNameToDelete testDataEngine.TestDataPointGroupNameType
				groupNameToDelete = currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.TestDataPointGroups[newOrEditTestDataPointGroupUI.SelectedIndexForGroups]

				// Delete the group
				delete(currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.ChosenTestDataPointsPerGroupMap, groupNameToDelete)

				// Rebuild the TestDataPointGroup-list
				currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.TestDataPointGroups = nil
				for testDataPointsGroupName, _ := range currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.ChosenTestDataPointsPerGroupMap {

					currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.TestDataPointGroups = append(currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.TestDataPointGroups, testDataPointsGroupName)
				}

				newOrEditTestDataPointGroupUI.SelectedIndexForGroups = -1

				testDataPointGroupsList.Refresh()
				testDataPointGroupsList.UnselectAll()

				// Clear the TestDataPointsList
				currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.TestDataPointsForAGroup = nil
				testDataPointsForAGroupList.Refresh()

				// UnSelect in DropDown- and List for TestDataPoints
				newOrEditTestDataPointGroupUI.TestDataPointsForAGroupSelect.ClearSelected()
				testDataPointsForAGroupList.UnselectAll()
				newOrEditTestDataPointGroupUI.SelectedIndexForGroupTestDataPoints = -1

				// When the delete Group is the one in the Group-Select, then Unselect it
				if string(groupNameToDelete) == newOrEditTestDataPointGroupUI.TestDataPointGroupsSelect.Selected {
					newOrEditTestDataPointGroupUI.TestDataPointGroupsSelect.ClearSelected()
					newOrEditTestDataPointGroupUI.TestDataPointGroupsSelect.SetOptions(currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.ListTestDataGroups())
					newOrEditTestDataPointGroupUI.TestDataPointGroupsSelect.Refresh()
				}

				// Update TestData on the TestCase
				//testCasesUiCanvasObject.TestCasesModelReference.TestCasesMapPtr[testCaseUuid] = currentTestSuitePtr

				// Update TestData-Selects in TestCase main area if the Deleted Group is selected in main TestCase area
				if string(groupNameToDelete) == testDataPointGroupsSelectInMainTestSuiteArea.Selected {

					// Reset Select options
					testDataPointGroupsSelectInMainTestSuiteArea.SetOptions(currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.ListTestDataGroups())
					testDataPointsForAGroupSelectInMainTestSuiteArea.SetOptions([]string{})
					testDataRowsForTestDataPointsSelectInMainTestSuiteArea.SetOptions([]string{})

					// Reset Selected
					testDataPointGroupsSelectInMainTestSuiteArea.ClearSelected()

					testDataPointGroupsSelectInMainTestSuiteArea.Refresh()
					testDataPointsForAGroupSelectInMainTestSuiteArea.Refresh()
					testDataRowsForTestDataPointsSelectInMainTestSuiteArea.Refresh()

					// If there are no TestDataGroups then hide the Selects in main TestCase window
					if len(currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.ListTestDataGroups()) == 0 {
						testDataSelectorsContainer.Hide()
					}
				} else {
					testDataPointGroupsSelectInMainTestSuiteArea.SetOptions(currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.ListTestDataGroups())
					testDataPointGroupsSelectInMainTestSuiteArea.Refresh()
				}

			}
		}, myWindow)
	})

	// Crete the 'Cllose'-button which closes the window
	closeButton := widget.NewButton("Close", func() {
		myWindow.Close()
	})

	// Create the container for handling TestDataGroups
	var buttonsContainer *fyne.Container
	buttonsContainer = container.NewHBox(newButton, editButton, deleteButton, closeButton)

	// Create the container that holds all UI components used for Groups and Points
	myContainer := container.NewBorder(nil, buttonsContainer, nil, nil, testDataGroupsAndPointsContainer)

	myWindow.SetContent(myContainer)

	// Function that updates new or changes lists in the UI
	go func() {

		var shouldListBeUpdated testDataEngine.ResponseChannelStruct
		var groupNameIndex int
		var groupNameIndexToSelect int

		for {

			shouldListBeUpdated = <-responseChannel

			// Update the List in main window if true as response
			if shouldListBeUpdated.ShouldBeUpdated == true {

				// Clear slice and variables used
				currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.TestDataPointGroups = nil
				groupNameIndex = 0
				groupNameIndexToSelect = 0

				for testDataPointsGroupName, _ := range currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.ChosenTestDataPointsPerGroupMap {

					currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.TestDataPointGroups = append(currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.TestDataPointGroups, testDataPointsGroupName)

					if testDataPointsGroupName == shouldListBeUpdated.TestDataPointGroupName {

						groupNameIndexToSelect = groupNameIndex

					}

					groupNameIndex = groupNameIndex + 1

				}

				fyne.Do(func() {
					testDataPointGroupsList.Refresh()
					testDataPointGroupsList.UnselectAll()
					testDataPointGroupsList.Select(groupNameIndexToSelect)
				})
				newOrEditTestDataPointGroupUI.SelectedIndexForGroups = groupNameIndexToSelect

				fyne.Do(func() {
					// Select the correct group in the dropdown for groups
					newOrEditTestDataPointGroupUI.TestDataPointGroupsSelect.SetOptions(testDataPointGroupsToStringSliceFunction())
					newOrEditTestDataPointGroupUI.TestDataPointGroupsSelect.SetSelected(string(shouldListBeUpdated.TestDataPointGroupName))
					newOrEditTestDataPointGroupUI.TestDataPointGroupsSelect.Refresh()

					// Select the correct TestDataPoint in the dropdown for TestDataPoints
					newOrEditTestDataPointGroupUI.TestDataPointsForAGroupSelect.SetOptions(testDataPointsToStringSliceFunction())
					newOrEditTestDataPointGroupUI.TestDataPointsForAGroupSelect.Refresh()
				})

				// Update TestData-Selects in TestCase main area
				testDataPointGroupsSelectInMainTestSuiteArea.SetOptions(testDataPointGroupsToStringSliceFunction())
				testDataPointsForAGroupSelectInMainTestSuiteArea.SetOptions(testDataPointsToStringSliceFunction())

				testDataPointGroupsSelectInMainTestSuiteArea.Refresh()
				testDataPointsForAGroupSelectInMainTestSuiteArea.Refresh()

				// Update TestData on the TestCase
				//testCasesUiCanvasObject.TestCasesModelReference.TestCasesMapPtr[testCaseUuid] = currentTestSuitePtr

				// If there are no TestDataGroups then hide the Selects in main TestCase window
				if len(currentTestSuitePtr.TestSuiteUIModelBinding.TestDataPtr.ListTestDataGroups()) == 0 {
					testDataSelectorsContainer.Hide()

				} else {
					testDataSelectorsContainer.Show()
				}
			}
		}
	}()

	myWindow.Show()

}
