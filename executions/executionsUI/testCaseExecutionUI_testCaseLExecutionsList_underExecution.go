package executionsUI

import (
	"FenixTesterGui/executions/executionsModel"
	"FenixTesterGui/headertable"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"strconv"
	"time"
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

	return mySortTable

}

// RemoveTestCaseExecutionFromUnderExecutionTable
// Remove from both table-slice and from Map that Table-slice got its data from
func RemoveTestCaseExecutionFromUnderExecutionTable(
	testCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference *executionsModel.TestCaseExecutionsUnderExecutionAdaptedForUiTableStruct,
	underExecutionTableChannelCommand executionsModel.UnderExecutionTableChannelCommandType) (err error) {

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

			// Depending on channel command, act differently
			switch underExecutionTableChannelCommand {

			case executionsModel.UnderExecutionTableAddRemoveChannelRemoveCommand_Flash:

				// Flash the row, to be deleted, in the table
				tableSizeHight, tableWidth := ExecutionsUIObject.UnderExecutionTable.Data.Length()

				if tableSizeHight > 0 {
					for columnCounter := 0; columnCounter < tableWidth; columnCounter++ {
						CellId := widget.TableCellID{
							Row: binderSlicePosition,
							Col: columnCounter,
						}
						var flashingTableCellsReference *headertable.FlashingTableCellStruct
						flashingTableCellsReference = ExecutionsUIObject.UnderExecutionTable.TableOpts.FlashingTableCellsReferenceMap[CellId]

						// Only call Flash-function when there is a reference, the reason for not having a reference is that Fynes table-engine only process visible table cells
						if flashingTableCellsReference != nil {
							headertable.FlashRowToBeRemoved(flashingTableCellsReference)
						}
					}
				}

				// Trigger Delete in parallell
				go func() {
					// Wait for color animation to be finished
					time.Sleep(time.Millisecond * 1000)

					// Create Remove-message to be put on channel
					var underExecutionTableAddRemoveChannelMessage executionsModel.UnderExecutionTableAddRemoveChannelStruct
					underExecutionTableAddRemoveChannelMessage = executionsModel.UnderExecutionTableAddRemoveChannelStruct{
						ChannelCommand: executionsModel.UnderExecutionTableAddRemoveChannelRemoveCommand_Remove,
						RemoveCommandData: executionsModel.UnderExecutionRemoveCommandDataStruct{
							TestCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference: testCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference,
						},
					}

					// Put on channel
					executionsModel.UnderExecutionTableAddRemoveChannel <- underExecutionTableAddRemoveChannelMessage

				}()

			case executionsModel.UnderExecutionTableAddRemoveChannelRemoveCommand_Remove:

				// Remove the element at index 'binderSlicePosition' from slice.
				executionsModel.TestCaseExecutionsUnderExecutionTableOptions.Bindings = remove(executionsModel.TestCaseExecutionsUnderExecutionTableOptions.Bindings, binderSlicePosition)

				// Delete data from original data adapted for Table
				delete(executionsModel.TestCaseExecutionsUnderExecutionMapAdaptedForUiTable, testCaseExecutionMapKey)

				// Resize the table based on its content
				ResizeTableColumns(ExecutionsUIObject.UnderExecutionTable)

				ExecutionsUIObject.UnderExecutionTable.Data.Refresh()

			default:
				// 'TestCaseExecutionVersion' doesn't exist within data
				errorId := "6a1c96ef-bc86-4363-967c-2e6788f9a874"
				err = errors.New(fmt.Sprintf("unknown 'underExecutionTableChannelCommand', '%s', in 'RemoveTestCaseExecutionFromUnderExecutionTable()': '%s' [ErrorID: %s]", underExecutionTableChannelCommand, errorId))

				fmt.Println(err) // TODO Send on Error Channel

				return err

			}

			// End loop
			break
		}

	}

	return err

}

// MoveTestCaseExecutionFromOnQueueToUnderExecution
// Move TestCaseExecution from OnQueue-table to UnderExecution-table
func MoveTestCaseExecutionFromOnQueueToUnderExecution(
	testCaseExecutionsOnQueueDataRowAdaptedForUiTableReference *executionsModel.TestCaseExecutionsOnQueueAdaptedForUiTableStruct,
	testCaseExecutionDetails *fenixExecutionServerGuiGrpcApi.TestCaseExecutionDetailsMessage) (err error) {

	var existInMap bool

	// Key to map: Should consist of 'TestCaseExecutionUuid' + 'TestCaseExecutionVersion'
	var testCaseExecutionMapKey executionsModel.TestCaseExecutionMapKeyType

	testCaseExecutionMapKey = executionsModel.TestCaseExecutionMapKeyType(testCaseExecutionsOnQueueDataRowAdaptedForUiTableReference.TestCaseExecutionUuid +
		testCaseExecutionsOnQueueDataRowAdaptedForUiTableReference.TestCaseExecutionVersion)

	// Extract OnQueueData to be moved to UnderExecution
	var tempTestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference *executionsModel.TestCaseExecutionsOnQueueAdaptedForUiTableStruct
	tempTestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference, existInMap = executionsModel.TestCaseExecutionsOnQueueMapAdaptedForUiTable[testCaseExecutionMapKey]

	if existInMap == false {

		errorId := "2942f959-6ceb-4f78-b06c-a8ff314b11f3"
		err = errors.New(fmt.Sprintf("'testCaseExecutionMapKey', '%s' doesn't exist in TestCaseExecutionsOnQueueMapAdaptedForUiTable [ErrorID: %s]", testCaseExecutionMapKey, errorId))

		fmt.Println(err) // TODO Send on Error Channel

		return err
	}

	//Create the new object to be added to UnderExecution-table
	var testCaseExecutionUnderExecutionAdaptedForUiTable *executionsModel.TestCaseExecutionsUnderExecutionAdaptedForUiTableStruct
	testCaseExecutionUnderExecutionAdaptedForUiTable = &executionsModel.TestCaseExecutionsUnderExecutionAdaptedForUiTableStruct{
		DomainUuid:                          tempTestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference.DomainUuid,
		DomainName:                          tempTestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference.DomainName,
		TestSuiteUuid:                       tempTestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference.TestSuiteUuid,
		TestSuiteName:                       tempTestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference.TestSuiteName,
		TestSuiteVersion:                    tempTestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference.TestSuiteVersion,
		TestSuiteExecutionUuid:              tempTestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference.TestSuiteExecutionUuid,
		TestSuiteExecutionVersion:           tempTestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference.TestSuiteExecutionVersion,
		TestCaseUuid:                        tempTestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference.TestCaseUuid,
		TestCaseName:                        tempTestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference.TestCaseName,
		TestCaseVersion:                     tempTestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference.TestCaseVersion,
		TestCaseExecutionUuid:               tempTestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference.TestCaseExecutionUuid,
		TestCaseExecutionVersion:            tempTestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference.TestCaseExecutionVersion,
		PlacedOnTestExecutionQueueTimeStamp: tempTestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference.PlacedOnTestExecutionQueueTimeStamp,
		ExecutionPriority:                   tempTestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference.ExecutionPriority,
		ExecutionStartTimeStamp:             testCaseExecutionDetails.ExecutionStartTimeStamp.AsTime().String(),
		ExecutionStopTimeStamp:              testCaseExecutionDetails.ExecutionStopTimeStamp.AsTime().String(),
		TestCaseExecutionStatus:             testCaseExecutionDetails.TestCaseExecutionStatus.String(),
		ExecutionHasFinished:                strconv.FormatBool(testCaseExecutionDetails.ExecutionHasFinished),
		ExecutionStatusUpdateTimeStamp:      testCaseExecutionDetails.ExecutionStatusUpdateTimeStamp.AsTime().String(),
	}

	// if 'testCaseExecutionMapKey' already exist in TestCaseExecutionsUnderExecutionMapAdaptedForUiTable''
	_, existInMap = executionsModel.TestCaseExecutionsUnderExecutionMapAdaptedForUiTable[testCaseExecutionMapKey]
	if existInMap == true {

		errorId := "3177ff88-46f5-4572-b722-ac541c761a64"
		err = errors.New(fmt.Sprintf("'testCaseExecutionMapKey', '%s' already exist in TestCaseExecutionsUnderExecutionMapAdaptedForUiTable [ErrorID: %s]", testCaseExecutionMapKey, errorId))

		fmt.Println(err) // TODO Send on Error Channel

		return err
	}

	// Append to map for TestCaseExecutionsUnderExecution-data used by UI-table
	executionsModel.TestCaseExecutionsUnderExecutionMapAdaptedForUiTable[testCaseExecutionMapKey] = testCaseExecutionUnderExecutionAdaptedForUiTable

	// Add a new binding for TestExecutionUnderExecutionRow data in the first position of slice
	executionsModel.TestCaseExecutionsUnderExecutionTableOptions.Bindings = append(
		[]binding.DataMap{binding.BindStruct(testCaseExecutionUnderExecutionAdaptedForUiTable)},
		executionsModel.TestCaseExecutionsUnderExecutionTableOptions.Bindings...)

	// Resize the table based on its content
	ResizeTableColumns(ExecutionsUIObject.UnderExecutionTable)

	// Update TestCaseExecutionUnderExecution-table
	ExecutionsUIObject.UnderExecutionTable.Data.Refresh()

	// Flash the newly added row in the table
	tableSizeHight, tableWidth := ExecutionsUIObject.UnderExecutionTable.Data.Length()

	if tableSizeHight > 0 {
		for columnCounter := 0; columnCounter < tableWidth; columnCounter++ {
			CellId := widget.TableCellID{
				Row: 0,
				Col: columnCounter,
			}
			var flashingTableCellsReference *headertable.FlashingTableCellStruct
			flashingTableCellsReference = ExecutionsUIObject.UnderExecutionTable.TableOpts.FlashingTableCellsReferenceMap[CellId]

			// Only call Flash-function when there is a reference, the reason for not having a reference is that Fynes table-engine only process visible table cells
			if flashingTableCellsReference != nil {
				headertable.FlashAddedRow(flashingTableCellsReference)
			}
		}
	}

	// Remove the old Execution from OnQueue
	// Create Remove-message to be put on channel
	var onQueueTableAddRemoveChannelMessage executionsModel.OnQueueTableAddRemoveChannelStruct
	onQueueTableAddRemoveChannelMessage = executionsModel.OnQueueTableAddRemoveChannelStruct{
		ChannelCommand: executionsModel.OnQueueTableAddRemoveChannelRemoveCommand_Flash,
		RemoveCommandData: executionsModel.OnQueueRemoveCommandDataStruct{
			TestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference: testCaseExecutionsOnQueueDataRowAdaptedForUiTableReference},
	}

	// Put on channel
	executionsModel.OnQueueTableAddRemoveChannel <- onQueueTableAddRemoveChannelMessage

	return err
}

func AddTestCaseExecutionUnderExecutionTable(testCaseExecutionBasicInformation *fenixExecutionServerGuiGrpcApi.TestCaseExecutionBasicInformationMessage) (err error) {

	if testCaseExecutionBasicInformation == nil {
		return err
	}
	// Convert 'raw' TestCaseExecutionsOnQueue-data into format to be used in UI
	var tempTestCaseExecutionsOnQueueAdaptedForUiTable *executionsModel.TestCaseExecutionsOnQueueAdaptedForUiTableStruct
	tempTestCaseExecutionsOnQueueAdaptedForUiTable = &executionsModel.TestCaseExecutionsOnQueueAdaptedForUiTableStruct{
		DomainUuid:                          testCaseExecutionBasicInformation.DomainUuid,
		DomainName:                          testCaseExecutionBasicInformation.DomainName,
		TestSuiteUuid:                       testCaseExecutionBasicInformation.TestSuiteUuid,
		TestSuiteName:                       testCaseExecutionBasicInformation.TestSuiteName,
		TestSuiteVersion:                    strconv.Itoa(int(testCaseExecutionBasicInformation.TestSuiteVersion)),
		TestSuiteExecutionUuid:              testCaseExecutionBasicInformation.TestSuiteExecutionUuid,
		TestSuiteExecutionVersion:           strconv.Itoa(int(testCaseExecutionBasicInformation.TestSuiteExecutionVersion)),
		TestCaseUuid:                        testCaseExecutionBasicInformation.TestCaseUuid,
		TestCaseName:                        testCaseExecutionBasicInformation.TestCaseName,
		TestCaseVersion:                     strconv.Itoa(int(testCaseExecutionBasicInformation.TestCaseVersion)),
		TestCaseExecutionUuid:               testCaseExecutionBasicInformation.TestCaseExecutionUuid,
		TestCaseExecutionVersion:            strconv.Itoa(int(testCaseExecutionBasicInformation.TestCaseExecutionVersion)),
		PlacedOnTestExecutionQueueTimeStamp: testCaseExecutionBasicInformation.PlacedOnTestExecutionQueueTimeStamp.AsTime().String(),
		ExecutionPriority:                   fenixExecutionServerGuiGrpcApi.ExecutionPriorityEnum_name[int32(testCaseExecutionBasicInformation.ExecutionPriority)],
	}

	// Verify that key is not already used in map
	// Key to map: Should consist of 'TestCaseExecutionUuid' + 'TestCaseExecutionVersion'
	var testCaseExecutionMapKey executionsModel.TestCaseExecutionMapKeyType

	testCaseExecutionMapKey = executionsModel.TestCaseExecutionMapKeyType(tempTestCaseExecutionsOnQueueAdaptedForUiTable.TestCaseExecutionUuid +
		tempTestCaseExecutionsOnQueueAdaptedForUiTable.TestCaseExecutionVersion)

	var existInMap bool
	_, existInMap = executionsModel.TestCaseExecutionsOnQueueMapAdaptedForUiTable[testCaseExecutionMapKey]
	if existInMap == true {

		errorId := "c51f60c4-2f27-495d-8e5e-0be0900dad03"
		err = errors.New(fmt.Sprintf("'testCaseExecutionMapKey', '%s' already exist in TestCaseExecutionsOnQueueMapAdaptedForUiTable [ErrorID: %s]", testCaseExecutionMapKey, errorId))

		fmt.Println(err) // TODO Send on Error Channel

		return err
	}

	// Append to map for TestCaseExecutionsUnderExecution-data used by UI-table
	executionsModel.TestCaseExecutionsOnQueueMapAdaptedForUiTable[testCaseExecutionMapKey] = tempTestCaseExecutionsOnQueueAdaptedForUiTable

	// Add a binding for TestExecutionOnQueueRow data
	executionsModel.TestCaseExecutionsOnQueueTableOptions.Bindings = append(
		executionsModel.TestCaseExecutionsOnQueueTableOptions.Bindings,
		binding.BindStruct(tempTestCaseExecutionsOnQueueAdaptedForUiTable))

	// Update TestCaseExecutionOnQueue-table
	ExecutionsUIObject.OnQueueTable.Data.Refresh()

	return err

}

// StartUnderExecutionTableAddRemoveChannelReader
// Start the channel reader and process messages from the channel
func StartUnderExecutionTableAddRemoveChannelReader() {

	var incomingUnderExecutionTableChannelCommand executionsModel.UnderExecutionTableAddRemoveChannelStruct
	var err error

	for {
		// Wait for incoming command over channel
		incomingUnderExecutionTableChannelCommand = <-executionsModel.UnderExecutionTableAddRemoveChannel

		switch incomingUnderExecutionTableChannelCommand.ChannelCommand {

		case executionsModel.UnderExecutionTableAddRemoveChannelAddCommand_MoveFromOnQueueToUnderExecution:
			MoveTestCaseExecutionFromOnQueueToUnderExecution(
				incomingUnderExecutionTableChannelCommand.AddCommandData.TestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference,
				incomingUnderExecutionTableChannelCommand.AddCommandData.TestCaseExecutionDetails)

		case executionsModel.UnderExecutionTableAddRemoveChannelRemoveCommand_Flash:
			RemoveTestCaseExecutionFromUnderExecutionTable(
				incomingUnderExecutionTableChannelCommand.RemoveCommandData.TestCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference,
				executionsModel.UnderExecutionTableAddRemoveChannelRemoveCommand_Flash)

		case executionsModel.UnderExecutionTableAddRemoveChannelRemoveCommand_Remove:
			RemoveTestCaseExecutionFromUnderExecutionTable(
				incomingUnderExecutionTableChannelCommand.RemoveCommandData.TestCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference,
				executionsModel.UnderExecutionTableAddRemoveChannelRemoveCommand_Remove)

		// No other command is supported
		default:

			errorId := "e6a0ffda-34cf-448e-bc96-6045d0825bc1"
			err = errors.New(fmt.Sprintf("unknown  'incomingUnderExecutionTableChannelCommand', '%s'. [ErrorID: %s]", incomingUnderExecutionTableChannelCommand, errorId))

			fmt.Println(err) //TODO Send on Error channel

		}
	}
}
