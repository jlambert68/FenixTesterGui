package executionsUI

import (
	"FenixTesterGui/executions/executionsModel"
	"FenixTesterGui/headertable"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"reflect"
)

func CreateTableForTestCaseExecutionsUnderExecution() *fyne.Container {
	var tableForTestCaseExecutionsUnderExecutionBindings []binding.DataMap

	// Create a binding for each TestExecutionUnderExecutionRow data
	/*
		for testDataRowCounter := 0; testDataRowCounter < len(executionsModel.TestCaseExecutionsUnderExecutionMapAdaptedForUiTable); testDataRowCounter++ {
			tableForTestCaseExecutionsUnderExecutionBindings = append(
				tableForTestCaseExecutionsUnderExecutionBindings,
				binding.BindStruct(&executionsModel.TestCaseExecutionsUnderExecutionMapAdaptedForUiTable[testDataRowCounter]))
		}
	*/

	for _, tempTestCaseExecutionsUnderExecutionDataAdaptedForUiTableReference := range executionsModel.TestCaseExecutionsUnderExecutionMapAdaptedForUiTable {
		tableForTestCaseExecutionsUnderExecutionBindings = append(
			tableForTestCaseExecutionsUnderExecutionBindings,
			binding.BindStruct(tempTestCaseExecutionsUnderExecutionDataAdaptedForUiTableReference))

	}

	executionsModel.TestCaseExecutionsUnderExecutionTableOptions.Bindings = tableForTestCaseExecutionsUnderExecutionBindings

	ht := headertable.NewSortingHeaderTable(&executionsModel.TestCaseExecutionsUnderExecutionTableOptions)
	ExecutionsUIObject.UnderExecutionTable = ht

	mySortTable := container.NewMax(ht)

	key1 := reflect.ValueOf(executionsModel.TestCaseExecutionsUnderExecutionMapAdaptedForUiTable).MapKeys()[1]
	fmt.Println(key1)
	value := executionsModel.TestCaseExecutionsUnderExecutionMapAdaptedForUiTable[executionsModel.TestCaseExecutionMapKeyType("d9c6fa2e-3d6a-477d-9727-a3083260777c1")]
	fmt.Println(value)
	//_ = RemoveTestCaseExecutionFromUnderExecutionTable(value)

	return mySortTable

}

func RemoveTestCaseExecutionFromUnderExecutionTable(testCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference *executionsModel.TestCaseExecutionsUnderExecutionAdaptedForUiTableStruct) (err error) {

	// Key to map: Should consist of 'TestCaseExecutionUuid' + 'TestCaseExecutionVersion'
	var testCaseExecutionMapKey executionsModel.TestCaseExecutionMapKeyType

	testCaseExecutionMapKey = executionsModel.TestCaseExecutionMapKeyType(testCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference.TestCaseExecutionUuid +
		testCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference.TestCaseExecutionVersion)

	var tempTestCaseExecutionUuidDataItem binding.DataItem
	var tempTestCaseExecutionVersionFromDataItem binding.DataItem
	var tempTestCaseExecutionUuidDataItemValue string
	var tempTestCaseExecutionVersionFromDataItemValue string

	// Loop all binding data and find the one to be removed
	for binderSlicePosition, tempTestCaseExecutionsUnderExecutionDataRowBinding := range executionsModel.TestCaseExecutionsUnderExecutionTableOptions.Bindings {
		fmt.Println(tempTestCaseExecutionsUnderExecutionDataRowBinding)

		dataMapBinding := executionsModel.TestCaseExecutionsUnderExecutionTableOptions.Bindings[binderSlicePosition]

		// Extract first part if MapKey from 'Binded data'
		tempTestCaseExecutionUuidDataItem, err = dataMapBinding.GetItem("TestCaseExecutionUuid")
		if err != nil {
			// 'TestCaseExecutionUuid' doesn't exist within data
			errorId := "329d10ca-b804-48f2-b91f-a3bf83f64386"
			err = errors.New(fmt.Sprintf("'TestCaseExecutionUuid', doesn't exist among binded data in 'tempTestCaseExecutionsUnderExecutionDataRowBinding': '%s' [ErrorID: %s]", tempTestCaseExecutionsUnderExecutionDataRowBinding, errorId))

			fmt.Println(err) // TODO Send on Error Channel

			return err
		}
		tempTestCaseExecutionUuidDataItemValue, err = tempTestCaseExecutionUuidDataItem.(binding.String).Get()
		if err != nil {
			// Couldn't extract value for 'TestCaseExecutionUuid'
			errorId := "dd8edb6c-f3be-4e69-9372-82e0251e7689"
			err = errors.New(fmt.Sprintf("couldn't get value from DataItem for 'TestCaseExecutionUuid', [ErrorID: %s]", errorId))

			fmt.Println(err) // TODO Send on Error Channel

			return err
		}

		tempTestCaseExecutionVersionFromDataItem, err = dataMapBinding.GetItem("TestCaseExecutionVersion")
		if err != nil {
			// 'TestCaseExecutionVersion' doesn't exist within data
			errorId := "04c9968a-c930-4f3e-8c6e-0d76515df1a5"
			err = errors.New(fmt.Sprintf("'TestCaseExecutionVersion', doesn't exist among binded data in 'tempTestCaseExecutionsUnderExecutionDataRowBinding': '%s' [ErrorID: %s]", tempTestCaseExecutionsUnderExecutionDataRowBinding, errorId))

			fmt.Println(err) // TODO Send on Error Channel

			return err
		}
		tempTestCaseExecutionVersionFromDataItemValue, err = tempTestCaseExecutionVersionFromDataItem.(binding.String).Get()
		if err != nil {
			// Couldn't extract value for 'TestCaseExecutionVersion'
			errorId := "aa2c7277-f0a6-4c9a-9517-1f3a41502a25"
			err = errors.New(fmt.Sprintf("couldn't get value from DataItem for 'TestCaseExecutionVersion', [ErrorID: %s]", errorId))

			fmt.Println(err) // TODO Send on Error Channel

			return err
		}

		// Check if this is the 'row' to delete
		if testCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference.TestCaseExecutionUuid == tempTestCaseExecutionUuidDataItemValue &&
			testCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference.TestCaseExecutionVersion == tempTestCaseExecutionVersionFromDataItemValue {

			// Remove the element at index 'binderSlicePosition' from slice.
			executionsModel.TestCaseExecutionsUnderExecutionTableOptions.Bindings = remove(executionsModel.TestCaseExecutionsUnderExecutionTableOptions.Bindings, binderSlicePosition)

			// Delete data from original data adapted for Table
			delete(executionsModel.TestCaseExecutionsUnderExecutionMapAdaptedForUiTable, testCaseExecutionMapKey)

			ExecutionsUIObject.UnderExecutionTable.Data.Refresh()

			break
		}

	}

	return err

}

// Remove item from the DataItem-slice and keep order
func remove(slice []binding.DataMap, s int) []binding.DataMap {
	return append(slice[:s], slice[s+1:]...)
}
