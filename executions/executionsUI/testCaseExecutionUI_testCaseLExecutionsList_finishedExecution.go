package executionsUI

import (
	"FenixTesterGui/executions/executionsModel"
	"FenixTesterGui/headertable"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"strconv"
)

// CreateTableForTestCaseExecutionsWithFinishedExecution
// Create bindings to the data used by the table and then create the UI-table itself
func CreateTableForTestCaseExecutionsWithFinishedExecution() *fyne.Container {
	var tableForTestCaseExecutionsWithFinishedExecutionBindings []binding.DataMap

	// Create a binding for each TestExecutionWithFinishedExecutionRow data
	for _, tempTestCaseExecutionsFinishedExecutionDataAdaptedForUiTableReference := range executionsModel.TestCaseExecutionsFinishedExecutionMapAdaptedForUiTable {
		tableForTestCaseExecutionsWithFinishedExecutionBindings = append(
			tableForTestCaseExecutionsWithFinishedExecutionBindings,
			binding.BindStruct(tempTestCaseExecutionsFinishedExecutionDataAdaptedForUiTableReference))
	}

	executionsModel.TestCaseExecutionsFinishedExecutionTableOptions.Bindings = tableForTestCaseExecutionsWithFinishedExecutionBindings

	ht := headertable.NewSortingHeaderTable(&executionsModel.TestCaseExecutionsFinishedExecutionTableOptions)
	ExecutionsUIObject.FinishedExecutionTable = ht

	mySortTable := container.NewMax(ht)

	return mySortTable

}

func RemoveTestCaseExecutionFromFinishedTable(testCaseExecutionsFinishedDataRowAdaptedForUiTableReference *executionsModel.TestCaseExecutionsFinishedExecutionAdaptedForUiTableStruct) (err error) {

	// Key to map: Should consist of 'TestCaseExecutionUuid' + 'TestCaseExecutionVersion'
	var testCaseExecutionMapKey executionsModel.TestCaseExecutionMapKeyType

	testCaseExecutionMapKey = executionsModel.TestCaseExecutionMapKeyType(testCaseExecutionsFinishedDataRowAdaptedForUiTableReference.TestCaseExecutionUuid +
		testCaseExecutionsFinishedDataRowAdaptedForUiTableReference.TestCaseExecutionVersion)

	var tempTestCaseExecutionUuidDataItem binding.DataItem
	var tempTestCaseExecutionVersionFromDataItem binding.DataItem
	var tempTestCaseExecutionUuidDataItemValue string
	var tempTestCaseExecutionVersionFromDataItemValue string

	// Loop all binding data and find the one to be removed
	for binderSlicePosition, tempTestCaseExecutionFinishedDataRowBinding := range executionsModel.TestCaseExecutionsFinishedExecutionTableOptions.Bindings {
		fmt.Println(tempTestCaseExecutionFinishedDataRowBinding)

		dataMapBinding := executionsModel.TestCaseExecutionsFinishedExecutionTableOptions.Bindings[binderSlicePosition]

		// Extract first part if MapKey from 'Binded data'
		tempTestCaseExecutionUuidDataItem, err = dataMapBinding.GetItem("TestCaseExecutionUuid")
		if err != nil {
			// 'TestCaseExecutionUuid' doesn't exist within data
			errorId := "b4c68d95-da80-41f6-8758-c04feff21241"
			err = errors.New(fmt.Sprintf("'TestCaseExecutionUuid', doesn't exist among binded data in 'tempTestCaseExecutionFinishedDataRowBinding': '%s' [ErrorID: %s]", tempTestCaseExecutionFinishedDataRowBinding, errorId))

			fmt.Println(err) // TODO Send on Error Channel

			return err
		}
		tempTestCaseExecutionUuidDataItemValue, err = tempTestCaseExecutionUuidDataItem.(binding.String).Get()
		if err != nil {
			// Couldn't extract value for 'TestCaseExecutionUuid'
			errorId := "9ea02458-4d28-4b49-a97f-baae1c7e0eea"
			err = errors.New(fmt.Sprintf("couldn't get value from DataItem for 'TestCaseExecutionUuid', [ErrorID: %s]", errorId))

			fmt.Println(err) // TODO Send on Error Channel

			return err
		}

		tempTestCaseExecutionVersionFromDataItem, err = dataMapBinding.GetItem("TestCaseExecutionVersion")
		if err != nil {
			// 'TestCaseExecutionVersion' doesn't exist within data
			errorId := "ed95af2a-3557-4e92-928f-89003858dc80"
			err = errors.New(fmt.Sprintf("'TestCaseExecutionVersion', doesn't exist among binded data in 'tempTestCaseExecutionFinishedDataRowBinding': '%s' [ErrorID: %s]", tempTestCaseExecutionFinishedDataRowBinding, errorId))

			fmt.Println(err) // TODO Send on Error Channel

			return err
		}
		tempTestCaseExecutionVersionFromDataItemValue, err = tempTestCaseExecutionVersionFromDataItem.(binding.String).Get()
		if err != nil {
			// Couldn't extract value for 'TestCaseExecutionVersion'
			errorId := "fed1b2f2-2175-4cc3-b53b-47e7dc18a6eb"
			err = errors.New(fmt.Sprintf("couldn't get value from DataItem for 'TestCaseExecutionVersion', [ErrorID: %s]", errorId))

			fmt.Println(err) // TODO Send on Error Channel

			return err
		}

		// Check if this is the 'row' to delete
		if testCaseExecutionsFinishedDataRowAdaptedForUiTableReference.TestCaseExecutionUuid == tempTestCaseExecutionUuidDataItemValue &&
			testCaseExecutionsFinishedDataRowAdaptedForUiTableReference.TestCaseExecutionVersion == tempTestCaseExecutionVersionFromDataItemValue {

			// Remove the element at index 'binderSlicePosition' from slice.
			executionsModel.TestCaseExecutionsFinishedExecutionTableOptions.Bindings = remove(executionsModel.TestCaseExecutionsFinishedExecutionTableOptions.Bindings, binderSlicePosition)

			// Delete data from original data adapted for Table
			delete(executionsModel.TestCaseExecutionsFinishedExecutionMapAdaptedForUiTable, testCaseExecutionMapKey)

			ExecutionsUIObject.FinishedExecutionTable.Data.Refresh()

			break
		}

	}

	return err

}

func MoveTestCaseInstructionExecutionFromUnderExecutionToFinishedExecution(testCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference *executionsModel.TestCaseExecutionsUnderExecutionAdaptedForUiTableStruct, testCaseExecutionDetails *fenixExecutionServerGuiGrpcApi.TestCaseExecutionDetailsMessage) (err error) {

	var existInMap bool

	// Key to map: Should consist of 'TestCaseExecutionUuid' + 'TestCaseExecutionVersion'
	var testCaseExecutionMapKey executionsModel.TestCaseExecutionMapKeyType

	testCaseExecutionMapKey = executionsModel.TestCaseExecutionMapKeyType(testCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference.TestCaseExecutionUuid +
		testCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference.TestCaseExecutionVersion)

	// Extract UnderExecutionData to be moved to UnderExecution
	var tempTestCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference *executionsModel.TestCaseExecutionsUnderExecutionAdaptedForUiTableStruct
	tempTestCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference, existInMap = executionsModel.TestCaseExecutionsUnderExecutionMapAdaptedForUiTable[testCaseExecutionMapKey]

	if existInMap == false {

		errorId := "7433e805-5687-483c-9e5d-4dd5d5f5d0b7"
		err = errors.New(fmt.Sprintf("'testCaseExecutionMapKey', '%s' doesn't exist in TestCaseExecutionsOnQueueMapAdaptedForUiTable [ErrorID: %s]", testCaseExecutionMapKey, errorId))

		fmt.Println(err) // TODO Send on Error Channel

		return err
	}

	//Create the new object to be added to FinishedExecution-table
	var testCaseExecutionFinishedExecutionAdaptedForUiTable *executionsModel.TestCaseExecutionsFinishedExecutionAdaptedForUiTableStruct
	testCaseExecutionFinishedExecutionAdaptedForUiTable = &executionsModel.TestCaseExecutionsFinishedExecutionAdaptedForUiTableStruct{
		DomainUuid:                          tempTestCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference.DomainUuid,
		DomainName:                          tempTestCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference.DomainName,
		TestSuiteUuid:                       tempTestCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference.TestSuiteUuid,
		TestSuiteName:                       tempTestCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference.TestSuiteName,
		TestSuiteVersion:                    tempTestCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference.TestSuiteVersion,
		TestSuiteExecutionUuid:              tempTestCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference.TestSuiteExecutionUuid,
		TestSuiteExecutionVersion:           tempTestCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference.TestSuiteExecutionVersion,
		TestCaseUuid:                        tempTestCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference.TestCaseUuid,
		TestCaseName:                        tempTestCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference.TestCaseName,
		TestCaseVersion:                     tempTestCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference.TestCaseVersion,
		TestCaseExecutionUuid:               tempTestCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference.TestCaseExecutionUuid,
		TestCaseExecutionVersion:            tempTestCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference.TestCaseExecutionVersion,
		PlacedOnTestExecutionQueueTimeStamp: tempTestCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference.PlacedOnTestExecutionQueueTimeStamp,
		ExecutionPriority:                   tempTestCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference.ExecutionPriority,
		ExecutionStartTimeStamp:             testCaseExecutionDetails.ExecutionStartTimeStamp.AsTime().String(),
		ExecutionStopTimeStamp:              testCaseExecutionDetails.ExecutionStopTimeStamp.AsTime().String(),
		TestCaseExecutionStatus:             testCaseExecutionDetails.TestCaseExecutionStatus.String(),
		ExecutionHasFinished:                strconv.FormatBool(testCaseExecutionDetails.ExecutionHasFinished),
		ExecutionStatusUpdateTimeStamp:      testCaseExecutionDetails.ExecutionStatusUpdateTimeStamp.AsTime().String(),
	}

	// if 'testCaseExecutionMapKey' already exist in TestCaseExecutionsFinishedExecutionMapAdaptedForUiTable''
	_, existInMap = executionsModel.TestCaseExecutionsFinishedExecutionMapAdaptedForUiTable[testCaseExecutionMapKey]
	if existInMap == true {

		errorId := "2101afd8-4b1d-4f16-ae14-8458f42d7b81"
		err = errors.New(fmt.Sprintf("'testCaseExecutionMapKey', '%s' already exist in TestCaseExecutionsFinishedExecutionMapAdaptedForUiTable [ErrorID: %s]", testCaseExecutionMapKey, errorId))

		fmt.Println(err) // TODO Send on Error Channel

		return err
	}

	// Append to map for TestCaseExecutionsFinishedExecution-data used by UI-table
	executionsModel.TestCaseExecutionsFinishedExecutionMapAdaptedForUiTable[testCaseExecutionMapKey] = testCaseExecutionFinishedExecutionAdaptedForUiTable

	// Add a binding for TestExecutionFinishedExecutionRow data
	executionsModel.TestCaseExecutionsFinishedExecutionTableOptions.Bindings = append(
		executionsModel.TestCaseExecutionsFinishedExecutionTableOptions.Bindings,
		binding.BindStruct(testCaseExecutionFinishedExecutionAdaptedForUiTable))

	// Update TestCaseExecutionFinishedExecution-table
	ExecutionsUIObject.FinishedExecutionTable.Data.Refresh()

	// Remove the old Execution from UnderExecutions
	err = RemoveTestCaseExecutionFromUnderExecutionTable(testCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference)

	return err
}
