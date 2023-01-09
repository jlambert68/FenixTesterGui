package executionsUIForSubscriptions

import (
	"FenixTesterGui/executions/executionsModelForSubscriptions"
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
	for _, tempTestCaseExecutionsOnQueueDataAdaptedForUiTableReference := range executionsModelForSubscriptions.TestCaseExecutionsOnQueueMapAdaptedForUiTable {
		tableForTestCaseExecutionsOnQueueBindings = append(
			tableForTestCaseExecutionsOnQueueBindings,
			binding.BindStruct(tempTestCaseExecutionsOnQueueDataAdaptedForUiTableReference))
	}

	executionsModelForSubscriptions.TestCaseExecutionsOnQueueTableOptions.Bindings = tableForTestCaseExecutionsOnQueueBindings

	ht := headertable.NewSortingHeaderTable(&executionsModelForSubscriptions.TestCaseExecutionsOnQueueTableOptions)
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

func RemoveTestCaseExecutionFromOnQueueTable(testCaseExecutionsOnQueueDataRowAdaptedForUiTableReference *executionsModelForSubscriptions.TestCaseExecutionsOnQueueAdaptedForUiTableStruct, onQueueTableChannelCommand executionsModelForSubscriptions.OnQueueTableChannelCommandType) (err error) {

	// Key to map: Should consist of 'TestCaseExecutionUuid' + 'TestCaseExecutionVersion'
	var testCaseExecutionMapKey executionsModelForSubscriptions.TestCaseExecutionMapKeyType

	testCaseExecutionMapKey = executionsModelForSubscriptions.TestCaseExecutionMapKeyType(testCaseExecutionsOnQueueDataRowAdaptedForUiTableReference.TestCaseExecutionUuid +
		testCaseExecutionsOnQueueDataRowAdaptedForUiTableReference.TestCaseExecutionVersion)

	var tempTestCaseExecutionUuidDataItem binding.DataItem
	var tempTestCaseExecutionVersionFromDataItem binding.DataItem
	var tempTestCaseExecutionUuidDataItemValue string
	var tempTestCaseExecutionVersionFromDataItemValue string

	// Loop all binding data and find the one to be removed
	for binderSlicePosition, tempTestCaseExecutionOnQueueDataRowBinding := range executionsModelForSubscriptions.TestCaseExecutionsOnQueueTableOptions.Bindings {
		fmt.Println(tempTestCaseExecutionOnQueueDataRowBinding)

		dataMapBinding := executionsModelForSubscriptions.TestCaseExecutionsOnQueueTableOptions.Bindings[binderSlicePosition]

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

			case executionsModelForSubscriptions.OnQueueTableAddRemoveChannelRemoveCommand_Flash:

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
					var onQueueTableAddRemoveChannelMessage executionsModelForSubscriptions.OnQueueTableAddRemoveChannelStruct
					onQueueTableAddRemoveChannelMessage = executionsModelForSubscriptions.OnQueueTableAddRemoveChannelStruct{
						ChannelCommand: executionsModelForSubscriptions.OnQueueTableAddRemoveChannelRemoveCommand_Remove,
						RemoveCommandData: executionsModelForSubscriptions.OnQueueRemoveCommandDataStruct{
							TestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference: testCaseExecutionsOnQueueDataRowAdaptedForUiTableReference},
					}

					// Put on channel
					executionsModelForSubscriptions.OnQueueTableAddRemoveChannel <- onQueueTableAddRemoveChannelMessage

				}()

			case executionsModelForSubscriptions.OnQueueTableAddRemoveChannelRemoveCommand_Remove:

				// Remove the element at index 'binderSlicePosition' from slice.
				executionsModelForSubscriptions.TestCaseExecutionsOnQueueTableOptions.Bindings = remove(executionsModelForSubscriptions.TestCaseExecutionsOnQueueTableOptions.Bindings, binderSlicePosition)

				// Delete data from original data adapted for Table
				delete(executionsModelForSubscriptions.TestCaseExecutionsOnQueueMapAdaptedForUiTable, testCaseExecutionMapKey)

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
	var tempTestCaseExecutionsOnQueueAdaptedForUiTable *executionsModelForSubscriptions.TestCaseExecutionsOnQueueAdaptedForUiTableStruct
	tempTestCaseExecutionsOnQueueAdaptedForUiTable = &executionsModelForSubscriptions.TestCaseExecutionsOnQueueAdaptedForUiTableStruct{
		ShowDetailedTestCaseExecution:       "true",
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

	// Verify that key is not already used in map.
	// Key to map: Should consist of 'TestCaseExecutionUuid' + 'TestCaseExecutionVersion'
	var testCaseExecutionMapKey executionsModelForSubscriptions.TestCaseExecutionMapKeyType

	testCaseExecutionMapKey = executionsModelForSubscriptions.TestCaseExecutionMapKeyType(tempTestCaseExecutionsOnQueueAdaptedForUiTable.TestCaseExecutionUuid +
		tempTestCaseExecutionsOnQueueAdaptedForUiTable.TestCaseExecutionVersion)

	// Define which TestCaseExecutions-table to look in and if the specific TestCaseExecution should exist in the tables
	var subscriptionsForTestCaseExecutionDetailsMap executionsModelForSubscriptions.SubscriptionsForTestCaseExecutionMapDetailsType
	subscriptionsForTestCaseExecutionDetailsMap =
		map[executionsModelForSubscriptions.SubscriptionTableType]executionsModelForSubscriptions.SubscriptionsForTestCaseExecutionMapDetailsStruct{
			executionsModelForSubscriptions.SubscriptionTableForTestCaseExecutionOnQueueTable:            executionsModelForSubscriptions.SubscriptionsForTestCaseExecutionMapDetailsStruct{ShouldExistInTable: false},
			executionsModelForSubscriptions.SubscriptionTableForTestCaseExecutionUnderExecutionTable:     executionsModelForSubscriptions.SubscriptionsForTestCaseExecutionMapDetailsStruct{ShouldExistInTable: false},
			executionsModelForSubscriptions.SubscriptionTableForTestCaseExecutionFinishedExecutionsTable: executionsModelForSubscriptions.SubscriptionsForTestCaseExecutionMapDetailsStruct{ShouldExistInTable: false},
		}

	// Create the map with key = 'TestCaseExecutionMapKey' to be sent to function
	var subscriptionsForTestCaseExecutionMap executionsModelForSubscriptions.SubscriptionsForTestCaseExecutionMapOverallType
	subscriptionsForTestCaseExecutionMap =
		map[executionsModelForSubscriptions.TestCaseExecutionMapKeyType]executionsModelForSubscriptions.SubscriptionsForTestCaseExecutionMapDetailsType{
			testCaseExecutionMapKey: subscriptionsForTestCaseExecutionDetailsMap,
		}

	// Verify if 'testCaseExecutionMapKey' is in use in any of OnQueue, UnderExecution or FinishedExecutions, depending on values sent in to functions
	err = verifyThatTestCaseExecutionIsNotInUse(subscriptionsForTestCaseExecutionMap)
	if err != nil {
		// Rule was not met due based on input parameters
		return nil
	}

	// Append to map for TestCaseExecutionsOnQueue-data used by UI-table
	executionsModelForSubscriptions.TestCaseExecutionsOnQueueMapAdaptedForUiTable[testCaseExecutionMapKey] = tempTestCaseExecutionsOnQueueAdaptedForUiTable

	// Add a new binding for TestExecutionOnQueueRow data in the first position of slice
	executionsModelForSubscriptions.TestCaseExecutionsOnQueueTableOptions.Bindings = append(
		[]binding.DataMap{binding.BindStruct(tempTestCaseExecutionsOnQueueAdaptedForUiTable)},
		executionsModelForSubscriptions.TestCaseExecutionsOnQueueTableOptions.Bindings...)

	// Resize the table based on its content
	ResizeTableColumns(ExecutionsUIObject.OnQueueTable)

	// Update TestCaseExecutionOnQueue-table
	ExecutionsUIObject.OnQueueTable.Data.Refresh()

	// Flash the newly added row in the table
	tableSizeHeight, tableWidth := ExecutionsUIObject.OnQueueTable.Data.Length()

	if tableSizeHeight > 0 {
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

	var incomingOnQueueTableChannelCommand executionsModelForSubscriptions.OnQueueTableAddRemoveChannelStruct
	var err error

	for {
		// Wait for incoming command over channel
		incomingOnQueueTableChannelCommand = <-executionsModelForSubscriptions.OnQueueTableAddRemoveChannel

		switch incomingOnQueueTableChannelCommand.ChannelCommand {

		case executionsModelForSubscriptions.OnQueueTableAddRemoveChannelAddCommand_AddAndFlash:
			_ = AddTestCaseExecutionToOnQueueTable(incomingOnQueueTableChannelCommand.AddCommandData.TestCaseExecutionBasicInformation)

		case executionsModelForSubscriptions.OnQueueTableAddRemoveChannelRemoveCommand_Flash:
			_ = RemoveTestCaseExecutionFromOnQueueTable(incomingOnQueueTableChannelCommand.RemoveCommandData.TestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference, executionsModelForSubscriptions.OnQueueTableAddRemoveChannelRemoveCommand_Flash)

		case executionsModelForSubscriptions.OnQueueTableAddRemoveChannelRemoveCommand_Remove:
			_ = RemoveTestCaseExecutionFromOnQueueTable(incomingOnQueueTableChannelCommand.RemoveCommandData.TestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference, executionsModelForSubscriptions.OnQueueTableAddRemoveChannelRemoveCommand_Remove)

		// No other command is supported
		default:

			errorId := "e6a0ffda-34cf-448e-bc96-6045d0825bc1"
			err = errors.New(fmt.Sprintf("unknown  'incomingOnQueueTableChannelCommand', '%s'. [ErrorID: %s]", incomingOnQueueTableChannelCommand, errorId))

			fmt.Println(err) //TODO Send on Error channel

		}
	}
}
