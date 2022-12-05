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

func CreateTableForTestCaseExecutionsOnQueue() *fyne.Container {
	var tableForTestCaseExecutionsOnQueueBindings []binding.DataMap

	// Create a binding for each TestExecutionOnQueueRow data
	for _, tempTestCaseExecutionsOnQueueDataAdaptedForUiTableReference := range executionsModel.TestCaseExecutionsOnQueueMapAdaptedForUiTable {
		tableForTestCaseExecutionsOnQueueBindings = append(
			tableForTestCaseExecutionsOnQueueBindings,
			binding.BindStruct(tempTestCaseExecutionsOnQueueDataAdaptedForUiTableReference))
	}

	executionsModel.TestCaseExecutionsOnQueueTableOptions.Bindings = tableForTestCaseExecutionsOnQueueBindings

	ht := headertable.NewSortingHeaderTable(&executionsModel.TestCaseExecutionsOnQueueTableOptions)
	ExecutionsUIObject.OnQueueTable = ht

	mySortTable := container.NewMax(ht)

	/*
		first := container.NewHBox(widget.NewLabel("Första"), mySortTable, widget.NewLabel("Sista"))

		second := container.NewHBox(first)

		scrollableTable := container.NewScroll(second)

		scrollableTableContainer := container.NewMax(scrollableTable)
	*/
	//ht.Header.ScrollToTrailing()
	ht.Header.Refresh()
	//ht.Header.ScrollToLeading()

	return mySortTable

}

func RemoveTestCaseExecutionFromOnQueueTable(testCaseExecutionsOnQueueDataRowAdaptedForUiTableReference *executionsModel.TestCaseExecutionsOnQueueAdaptedForUiTableStruct, onQueueTableChannelCommand executionsModel.OnQueueTableChannelCommandType) (err error) {

	// Key to map: Should consist of 'TestCaseExecutionUuid' + 'TestCaseExecutionVersion'
	var testCaseExecutionMapKey executionsModel.TestCaseExecutionMapKeyType

	testCaseExecutionMapKey = executionsModel.TestCaseExecutionMapKeyType(testCaseExecutionsOnQueueDataRowAdaptedForUiTableReference.TestCaseExecutionUuid +
		testCaseExecutionsOnQueueDataRowAdaptedForUiTableReference.TestCaseExecutionVersion)

	var tempTestCaseExecutionUuidDataItem binding.DataItem
	var tempTestCaseExecutionVersionFromDataItem binding.DataItem
	var tempTestCaseExecutionUuidDataItemValue string
	var tempTestCaseExecutionVersionFromDataItemValue string

	// Loop all binding data and find the one to be removed
	for binderSlicePosition, tempTestCaseExecutionOnQueueDataRowBinding := range executionsModel.TestCaseExecutionsOnQueueTableOptions.Bindings {
		fmt.Println(tempTestCaseExecutionOnQueueDataRowBinding)

		dataMapBinding := executionsModel.TestCaseExecutionsOnQueueTableOptions.Bindings[binderSlicePosition]

		// Extract first part if MapKey from 'Binded data'
		tempTestCaseExecutionUuidDataItem, err = dataMapBinding.GetItem("TestCaseExecutionUuid")
		if err != nil {
			// 'TestCaseExecutionUuid' doesn't exist within data
			errorId := "c4f14d9f-d479-493e-9b4a-f31a756a25f3"
			err = errors.New(fmt.Sprintf("'TestCaseExecutionUuid', doesn't exist among binded data in 'tempTestCaseExecutionOnQueueDataRowBinding': '%s' [ErrorID: %s]", tempTestCaseExecutionOnQueueDataRowBinding, errorId))

			fmt.Println(err) // TODO Send on Error Channel

			return err
		}
		tempTestCaseExecutionUuidDataItemValue, err = tempTestCaseExecutionUuidDataItem.(binding.String).Get()
		if err != nil {
			// Couldn't extract value for 'TestCaseExecutionUuid'
			errorId := "bb38c3d3-2091-48d2-838f-95fc045a4146"
			err = errors.New(fmt.Sprintf("couldn't get value from DataItem for 'TestCaseExecutionUuid', [ErrorID: %s]", errorId))

			fmt.Println(err) // TODO Send on Error Channel

			return err
		}

		tempTestCaseExecutionVersionFromDataItem, err = dataMapBinding.GetItem("TestCaseExecutionVersion")
		if err != nil {
			// 'TestCaseExecutionVersion' doesn't exist within data
			errorId := "c2ff0b36-4a17-423b-95c8-0ce8c0be8dcd"
			err = errors.New(fmt.Sprintf("'TestCaseExecutionVersion', doesn't exist among binded data in 'tempTestCaseExecutionOnQueueDataRowBinding': '%s' [ErrorID: %s]", tempTestCaseExecutionOnQueueDataRowBinding, errorId))

			fmt.Println(err) // TODO Send on Error Channel

			return err
		}
		tempTestCaseExecutionVersionFromDataItemValue, err = tempTestCaseExecutionVersionFromDataItem.(binding.String).Get()
		if err != nil {
			// Couldn't extract value for 'TestCaseExecutionVersion'
			errorId := "5e69ddfc-7447-428e-b3af-3e45a997a77d"
			err = errors.New(fmt.Sprintf("couldn't get value from DataItem for 'TestCaseExecutionVersion', [ErrorID: %s]", errorId))

			fmt.Println(err) // TODO Send on Error Channel

			return err
		}

		// Check if this is the 'row' to delete
		if testCaseExecutionsOnQueueDataRowAdaptedForUiTableReference.TestCaseExecutionUuid == tempTestCaseExecutionUuidDataItemValue &&
			testCaseExecutionsOnQueueDataRowAdaptedForUiTableReference.TestCaseExecutionVersion == tempTestCaseExecutionVersionFromDataItemValue {

			// Depending on channel command, act differently
			switch onQueueTableChannelCommand {

			case executionsModel.OnQueueTableAddRemoveChannelRemoveCommand_Flash:

				// Flash the row, to be deleted, in the table
				tableSizeHight, tableWidth := ExecutionsUIObject.OnQueueTable.Data.Length()

				if tableSizeHight > 0 {
					for columnCounter := 0; columnCounter < tableWidth; columnCounter++ {
						CellId := widget.TableCellID{
							Row: binderSlicePosition,
							Col: columnCounter,
						}
						var flashingTableCellsReference *headertable.FlashingTableCellStruct
						flashingTableCellsReference = ExecutionsUIObject.OnQueueTable.TableOpts.FlashingTableCellsReferenceMap[CellId]

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
					var onQueueTableAddRemoveChannelMessage executionsModel.OnQueueTableAddRemoveChannelStruct
					onQueueTableAddRemoveChannelMessage = executionsModel.OnQueueTableAddRemoveChannelStruct{
						ChannelCommand: executionsModel.OnQueueTableAddRemoveChannelRemoveCommand_Remove,
						RemoveCommandData: executionsModel.OnQueueRemoveCommandDataStruct{
							TestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference: testCaseExecutionsOnQueueDataRowAdaptedForUiTableReference},
					}

					// Put on channel
					executionsModel.OnQueueTableAddRemoveChannel <- onQueueTableAddRemoveChannelMessage

				}()

			case executionsModel.OnQueueTableAddRemoveChannelRemoveCommand_Remove:

				// Remove the element at index 'binderSlicePosition' from slice.
				executionsModel.TestCaseExecutionsOnQueueTableOptions.Bindings = remove(executionsModel.TestCaseExecutionsOnQueueTableOptions.Bindings, binderSlicePosition)

				// Delete data from original data adapted for Table
				delete(executionsModel.TestCaseExecutionsOnQueueMapAdaptedForUiTable, testCaseExecutionMapKey)

				// Resize the table based on its content
				ResizeTableColumns(ExecutionsUIObject.OnQueueTable)

				ExecutionsUIObject.OnQueueTable.Data.Refresh()

			default:
				// 'TestCaseExecutionVersion' doesn't exist within data
				errorId := "df7d5103-f9b8-4f19-82ee-281733c290f6"
				err = errors.New(fmt.Sprintf("unknown 'onQueueTableChannelCommand', '%s', in 'RemoveTestCaseExecutionFromOnQueueTable()': '%s' [ErrorID: %s]", onQueueTableChannelCommand, errorId))

				fmt.Println(err) // TODO Send on Error Channel

				return err

			}

			// End loop
			break
		}

	}

	return err

}

func AddTestCaseExecutionToOnQueueTable(testCaseExecutionBasicInformation *fenixExecutionServerGuiGrpcApi.TestCaseExecutionBasicInformationMessage) (err error) {

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

	// Append to map for TestCaseExecutionsOnQueue-data used by UI-table
	executionsModel.TestCaseExecutionsOnQueueMapAdaptedForUiTable[testCaseExecutionMapKey] = tempTestCaseExecutionsOnQueueAdaptedForUiTable

	// Add a new binding for TestExecutionOnQueueRow data in the first position of slice
	executionsModel.TestCaseExecutionsOnQueueTableOptions.Bindings = append(
		[]binding.DataMap{binding.BindStruct(tempTestCaseExecutionsOnQueueAdaptedForUiTable)},
		executionsModel.TestCaseExecutionsOnQueueTableOptions.Bindings...)

	// Resize the table based on its content
	ResizeTableColumns(ExecutionsUIObject.OnQueueTable)

	// Update TestCaseExecutionOnQueue-table
	ExecutionsUIObject.OnQueueTable.Data.Refresh()

	// Flash the newly added row in the table
	tableSizeHight, tableWidth := ExecutionsUIObject.OnQueueTable.Data.Length()

	if tableSizeHight > 0 {
		for columnCounter := 0; columnCounter < tableWidth; columnCounter++ {
			CellId := widget.TableCellID{
				Row: 0,
				Col: columnCounter,
			}
			var flashingTableCellsReference *headertable.FlashingTableCellStruct
			flashingTableCellsReference = ExecutionsUIObject.OnQueueTable.TableOpts.FlashingTableCellsReferenceMap[CellId]

			// Only call Flash-function when there is a reference, the reason for not having a reference is that Fynes table-engine only process visible table cells
			if flashingTableCellsReference != nil {
				headertable.FlashAddedRow(flashingTableCellsReference)
			}
		}
	}

	return err

}

// StartOnQueueTableAddRemoveChannelReader
// Start the channel reader and process messages from the channel
func StartOnQueueTableAddRemoveChannelReader() {

	var incomingOnQueueTableChannelCommand executionsModel.OnQueueTableAddRemoveChannelStruct
	var err error

	for {
		// Wait for incoming command over channel
		incomingOnQueueTableChannelCommand = <-executionsModel.OnQueueTableAddRemoveChannel

		switch incomingOnQueueTableChannelCommand.ChannelCommand {

		case executionsModel.OnQueueTableAddRemoveChannelAddCommand_AddAndFlash:
			_ = AddTestCaseExecutionToOnQueueTable(incomingOnQueueTableChannelCommand.AddCommandData.TestCaseExecutionBasicInformation)

		case executionsModel.OnQueueTableAddRemoveChannelRemoveCommand_Flash:
			_ = RemoveTestCaseExecutionFromOnQueueTable(incomingOnQueueTableChannelCommand.RemoveCommandData.TestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference, executionsModel.OnQueueTableAddRemoveChannelRemoveCommand_Flash)

		case executionsModel.OnQueueTableAddRemoveChannelRemoveCommand_Remove:
			_ = RemoveTestCaseExecutionFromOnQueueTable(incomingOnQueueTableChannelCommand.RemoveCommandData.TestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference, executionsModel.OnQueueTableAddRemoveChannelRemoveCommand_Remove)

		// No other command is supported
		default:

			errorId := "e6a0ffda-34cf-448e-bc96-6045d0825bc1"
			err = errors.New(fmt.Sprintf("unknown  'incomingOnQueueTableChannelCommand', '%s'. [ErrorID: %s]", incomingOnQueueTableChannelCommand, errorId))

			fmt.Println(err) //TODO Send on Error channel

		}
	}
}
