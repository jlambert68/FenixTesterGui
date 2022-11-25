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
	mySortTable := container.NewMax(ht)

	key1 := reflect.ValueOf(executionsModel.TestCaseExecutionsUnderExecutionMapAdaptedForUiTable).MapKeys()[1]
	fmt.Println(key1)
	value := executionsModel.TestCaseExecutionsUnderExecutionMapAdaptedForUiTable[executionsModel.TestCaseExecutionMapKeyType("d9c6fa2e-3d6a-477d-9727-a3083260777c1")]
	//executionsModel.TestCaseExecutionMapKeyType
	_ = RemoveBindingToTableDataForUnderExecutionTable(value)

	return mySortTable

}

// R
// afas
func RemoveBindingToTableDataForUnderExecutionTable(testCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference *executionsModel.TestCaseExecutionsUnderExecutionAdaptedForUiTableStruct) (err error) {

	// Key to map: Should consist of 'TestCaseExecutionUuid' + 'TestCaseExecutionVersion'
	//var testCaseExecutionMapKey executionsModel.TestCaseExecutionMapKeyType

	Testa att ta bort i orginaldatat och se om det påverkar det "bindade" datat
	var tempTestCaseExecutionUuidFromBindedData binding.DataItem
	var tempTestCaseExecutionVersionFromBindedData binding.DataItem

	// Loop all binding data and find the one to be removed
	for binderSlicePosition, tempTestCaseExecutionsUnderExecutionDataRowBinding := range executionsModel.TestCaseExecutionsUnderExecutionTableOptions.Bindings {

		// Extract first part if MapKey from 'Binded data'
		tempTestCaseExecutionUuidFromBindedData, err = tempTestCaseExecutionsUnderExecutionDataRowBinding.GetItem("TestCaseExecutionUuid")
		if err != nil {
			// 'TestCaseExecutionUuid' doesn't exist within data
			errorId := "329d10ca-b804-48f2-b91f-a3bf83f64386"
			err = errors.New(fmt.Sprintf("'TestCaseExecutionUuid', doesn't exist among binded data in 'tempTestCaseExecutionsUnderExecutionDataRowBinding': '%s' [ErrorID: %s]", tempTestCaseExecutionsUnderExecutionDataRowBinding, errorId))

			fmt.Println(err) // TODO Send on Error Channel

			return err
		}

		tempTestCaseExecutionVersionFromBindedData, err = tempTestCaseExecutionsUnderExecutionDataRowBinding.GetItem("TestCaseExecutionVersion")
		if err != nil {
			// 'TestCaseExecutionVersion' doesn't exist within data
			errorId := "04c9968a-c930-4f3e-8c6e-0d76515df1a5"
			err = errors.New(fmt.Sprintf("'TestCaseExecutionVersion', doesn't exist among binded data in 'tempTestCaseExecutionsUnderExecutionDataRowBinding': '%s' [ErrorID: %s]", tempTestCaseExecutionsUnderExecutionDataRowBinding, errorId))

			fmt.Println(err) // TODO Send on Error Channel

			return err
		}

		// If this is the 'binded' data that should be removed then remove it stopp looping
		fmt.Println(binderSlicePosition)
		fmt.Println(tempTestCaseExecutionUuidFromBindedData)
		fmt.Println(tempTestCaseExecutionVersionFromBindedData)

	}

	return err

}
