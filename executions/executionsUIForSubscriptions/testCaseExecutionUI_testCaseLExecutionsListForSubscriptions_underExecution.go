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

func CreateTableForTestCaseExecutionsUnderExecution() *fyne.Container {
	var tableForTestCaseExecutionsUnderExecutionBindings []binding.DataMap

	// Create a binding for each TestExecutionUnderExecutionRow data
	/*
		for testDataRowCounter := 0; testDataRowCounter < len(executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionMapAdaptedForUiTable); testDataRowCounter++ {
			tableForTestCaseExecutionsUnderExecutionBindings = append(
				tableForTestCaseExecutionsUnderExecutionBindings,
				binding.BindStruct(&executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionMapAdaptedForUiTable[testDataRowCounter]))
		}
	*/

	for _, tempTestCaseExecutionsUnderExecutionDataAdaptedForUiTableReference := range executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionMapAdaptedForUiTable {
		tableForTestCaseExecutionsUnderExecutionBindings = append(
			tableForTestCaseExecutionsUnderExecutionBindings,
			binding.BindStruct(tempTestCaseExecutionsUnderExecutionDataAdaptedForUiTableReference))

	}

	executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionTableOptions.Bindings = tableForTestCaseExecutionsUnderExecutionBindings

	ht := headertable.NewSortingHeaderTable(&executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionTableOptions)
	ExecutionsUIObject.UnderExecutionTable = ht

	mySortTable := container.NewMax(ht)

	return mySortTable

}

// RemoveTestCaseExecutionFromUnderExecutionTable
// Remove from both table-slice and from Map that Table-slice got its data from
func RemoveTestCaseExecutionFromUnderExecutionTable(
	testCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference *executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionAdaptedForUiTableStruct,
	underExecutionTableChannelCommand executionsModelForSubscriptions.UnderExecutionTableChannelCommandType) (err error) {

	// Key to map: Should consist of 'TestCaseExecutionUuid' + 'TestCaseExecutionVersion'
	var testCaseExecutionMapKey executionsModelForSubscriptions.TestCaseExecutionMapKeyType

	testCaseExecutionMapKey = executionsModelForSubscriptions.TestCaseExecutionMapKeyType(testCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference.TestCaseExecutionUuid +
		testCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference.TestCaseExecutionVersion)

	var tempTestCaseExecutionUuidDataItem binding.DataItem
	var tempTestCaseExecutionVersionFromDataItem binding.DataItem
	var tempTestCaseExecutionUuidDataItemValue string
	var tempTestCaseExecutionVersionFromDataItemValue string

	// Loop all binding data and find the one to be removed
	for binderSlicePosition, tempTestCaseExecutionsUnderExecutionDataRowBinding := range executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionTableOptions.Bindings {
		fmt.Println(tempTestCaseExecutionsUnderExecutionDataRowBinding)

		dataMapBinding := executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionTableOptions.Bindings[binderSlicePosition]

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

			case executionsModelForSubscriptions.UnderExecutionTableAddRemoveChannelRemoveCommand_Flash:

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
					var underExecutionTableAddRemoveChannelMessage executionsModelForSubscriptions.UnderExecutionTableAddRemoveChannelStruct
					underExecutionTableAddRemoveChannelMessage = executionsModelForSubscriptions.UnderExecutionTableAddRemoveChannelStruct{
						ChannelCommand: executionsModelForSubscriptions.UnderExecutionTableAddRemoveChannelRemoveCommand_Remove,
						RemoveCommandData: executionsModelForSubscriptions.UnderExecutionRemoveCommandDataStruct{
							TestCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference: testCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference,
						},
					}

					// Put on channel
					executionsModelForSubscriptions.UnderExecutionTableAddRemoveChannel <- underExecutionTableAddRemoveChannelMessage

				}()

			case executionsModelForSubscriptions.UnderExecutionTableAddRemoveChannelRemoveCommand_Remove:

				// Remove the element at index 'binderSlicePosition' from slice.
				executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionTableOptions.Bindings = remove(executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionTableOptions.Bindings, binderSlicePosition)

				// Delete data from original data adapted for Table
				delete(executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionMapAdaptedForUiTable, testCaseExecutionMapKey)

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
	testCaseExecutionsOnQueueDataRowAdaptedForUiTableReference *executionsModelForSubscriptions.TestCaseExecutionsOnQueueAdaptedForUiTableStruct,
	testCaseExecutionDetails *fenixExecutionServerGuiGrpcApi.TestCaseExecutionDetailsMessage) (err error) {

	var existInMap bool

	// Key to map: Should consist of 'TestCaseExecutionUuid' + 'TestCaseExecutionVersion'
	var testCaseExecutionMapKey executionsModelForSubscriptions.TestCaseExecutionMapKeyType

	testCaseExecutionMapKey = executionsModelForSubscriptions.TestCaseExecutionMapKeyType(testCaseExecutionsOnQueueDataRowAdaptedForUiTableReference.TestCaseExecutionUuid +
		testCaseExecutionsOnQueueDataRowAdaptedForUiTableReference.TestCaseExecutionVersion)

	// Secure that there is one TestCaseExecution in OnQueue and nowhere else
	// Define which TestCaseExecutions-table to look in and if the specific TestCaseExecution should exist in the tables
	var subscriptionsForTestCaseExecutionDetailsMap executionsModelForSubscriptions.SubscriptionsForTestCaseExecutionMapDetailsType
	subscriptionsForTestCaseExecutionDetailsMap =
		map[executionsModelForSubscriptions.SubscriptionTableType]executionsModelForSubscriptions.SubscriptionsForTestCaseExecutionMapDetailsStruct{
			executionsModelForSubscriptions.SubscriptionTableForTestCaseExecutionOnQueueTable:            executionsModelForSubscriptions.SubscriptionsForTestCaseExecutionMapDetailsStruct{ShouldExistInTable: true},
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
		// Rule was not met due based on input parameters.
		// Probably due to several due to same TestCaseExecution-status has been sent multiple times from ExecutionEngine
		return nil
	}

	// Extract OnQueueData to be moved to UnderExecution
	var tempTestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference *executionsModelForSubscriptions.TestCaseExecutionsOnQueueAdaptedForUiTableStruct
	tempTestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference, existInMap = executionsModelForSubscriptions.TestCaseExecutionsOnQueueMapAdaptedForUiTable[testCaseExecutionMapKey]

	if existInMap == false {

		errorId := "2942f959-6ceb-4f78-b06c-a8ff314b11f3"
		err = errors.New(fmt.Sprintf("'testCaseExecutionMapKey', '%s' doesn't exist in TestCaseExecutionsOnQueueMapAdaptedForUiTable [ErrorID: %s]", testCaseExecutionMapKey, errorId))

		fmt.Println(err) // TODO Send on Error Channel

		return err
	}

	// Create the new object to be added to UnderExecution-table
	var testCaseExecutionUnderExecutionAdaptedForUiTable *executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionAdaptedForUiTableStruct
	testCaseExecutionUnderExecutionAdaptedForUiTable = &executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionAdaptedForUiTableStruct{
		ShowDetailedTestCaseExecution:       tempTestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference.ShowDetailedTestCaseExecution,
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
	_, existInMap = executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionMapAdaptedForUiTable[testCaseExecutionMapKey]
	if existInMap == true {
		// Might happen when ExecutionEngine sends the same TestCaseExecution-status several time

		errorId := "3177ff88-46f5-4572-b722-ac541c761a64"
		err = errors.New(fmt.Sprintf("'testCaseExecutionMapKey', '%s' already exist in TestCaseExecutionsUnderExecutionMapAdaptedForUiTable [ErrorID: %s]", testCaseExecutionMapKey, errorId))

		fmt.Println(err) // TODO Send on Error Channel

		return err
	}

	// Append to map for TestCaseExecutionsUnderExecution-data used by UI-table
	executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionMapAdaptedForUiTable[testCaseExecutionMapKey] = testCaseExecutionUnderExecutionAdaptedForUiTable

	// Add a new binding for TestExecutionUnderExecutionRow data in the first position of slice
	executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionTableOptions.Bindings = append(
		[]binding.DataMap{binding.BindStruct(testCaseExecutionUnderExecutionAdaptedForUiTable)},
		executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionTableOptions.Bindings...)

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
	var onQueueTableAddRemoveChannelMessage executionsModelForSubscriptions.OnQueueTableAddRemoveChannelStruct
	onQueueTableAddRemoveChannelMessage = executionsModelForSubscriptions.OnQueueTableAddRemoveChannelStruct{
		ChannelCommand: executionsModelForSubscriptions.OnQueueTableAddRemoveChannelRemoveCommand_Flash,
		RemoveCommandData: executionsModelForSubscriptions.OnQueueRemoveCommandDataStruct{
			TestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference: testCaseExecutionsOnQueueDataRowAdaptedForUiTableReference},
	}

	// Put on channel
	executionsModelForSubscriptions.OnQueueTableAddRemoveChannel <- onQueueTableAddRemoveChannelMessage

	return err
}

func AddTestCaseExecutionUnderExecutionTable(testCaseExecutionBasicInformation *fenixExecutionServerGuiGrpcApi.TestCaseExecutionBasicInformationMessage) (err error) {

	if testCaseExecutionBasicInformation == nil {
		return err
	}
	// Convert 'raw' TestCaseExecutionsOnQueue-data into format to be used in UI
	var tempTestCaseExecutionsOnQueueAdaptedForUiTable *executionsModelForSubscriptions.TestCaseExecutionsOnQueueAdaptedForUiTableStruct
	tempTestCaseExecutionsOnQueueAdaptedForUiTable = &executionsModelForSubscriptions.TestCaseExecutionsOnQueueAdaptedForUiTableStruct{
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
	var testCaseExecutionMapKey executionsModelForSubscriptions.TestCaseExecutionMapKeyType

	testCaseExecutionMapKey = executionsModelForSubscriptions.TestCaseExecutionMapKeyType(tempTestCaseExecutionsOnQueueAdaptedForUiTable.TestCaseExecutionUuid +
		tempTestCaseExecutionsOnQueueAdaptedForUiTable.TestCaseExecutionVersion)

	// Define which TestCaseExecutions-table to look in and if the specific TestCaseExecution should exist in the tables
	var subscriptionsForTestCaseExecutionDetailsMap executionsModelForSubscriptions.SubscriptionsForTestCaseExecutionMapDetailsType
	subscriptionsForTestCaseExecutionDetailsMap =
		map[executionsModelForSubscriptions.SubscriptionTableType]executionsModelForSubscriptions.SubscriptionsForTestCaseExecutionMapDetailsStruct{
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

	// Append to map for TestCaseExecutionsUnderExecution-data used by UI-table
	executionsModelForSubscriptions.TestCaseExecutionsOnQueueMapAdaptedForUiTable[testCaseExecutionMapKey] = tempTestCaseExecutionsOnQueueAdaptedForUiTable

	// Add a binding for TestExecutionOnQueueRow data
	executionsModelForSubscriptions.TestCaseExecutionsOnQueueTableOptions.Bindings = append(
		executionsModelForSubscriptions.TestCaseExecutionsOnQueueTableOptions.Bindings,
		binding.BindStruct(tempTestCaseExecutionsOnQueueAdaptedForUiTable))

	// Update TestCaseExecutionOnQueue-table
	ExecutionsUIObject.OnQueueTable.Data.Refresh()

	return err

}

// StartUnderExecutionTableAddRemoveChannelReader
// Start the channel reader and process messages from the channel
func StartUnderExecutionTableAddRemoveChannelReader() {

	var incomingUnderExecutionTableChannelCommand executionsModelForSubscriptions.UnderExecutionTableAddRemoveChannelStruct
	var err error

	for {
		// Wait for incoming command over channel
		incomingUnderExecutionTableChannelCommand = <-executionsModelForSubscriptions.UnderExecutionTableAddRemoveChannel

		switch incomingUnderExecutionTableChannelCommand.ChannelCommand {

		case executionsModelForSubscriptions.UnderExecutionTableAddRemoveChannelAddCommand_MoveFromOnQueueToUnderExecution:
			_ = MoveTestCaseExecutionFromOnQueueToUnderExecution(
				incomingUnderExecutionTableChannelCommand.AddCommandData.TestCaseExecutionsOnQueueDataRowAdaptedForUiTableReference,
				incomingUnderExecutionTableChannelCommand.AddCommandData.TestCaseExecutionDetails)

		case executionsModelForSubscriptions.UnderExecutionTableAddRemoveChannelRemoveCommand_Flash:
			_ = RemoveTestCaseExecutionFromUnderExecutionTable(
				incomingUnderExecutionTableChannelCommand.RemoveCommandData.TestCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference,
				executionsModelForSubscriptions.UnderExecutionTableAddRemoveChannelRemoveCommand_Flash)

		case executionsModelForSubscriptions.UnderExecutionTableAddRemoveChannelRemoveCommand_Remove:
			_ = RemoveTestCaseExecutionFromUnderExecutionTable(
				incomingUnderExecutionTableChannelCommand.RemoveCommandData.TestCaseExecutionsUnderExecutionDataRowAdaptedForUiTableReference,
				executionsModelForSubscriptions.UnderExecutionTableAddRemoveChannelRemoveCommand_Remove)

		// No other command is supported
		default:

			errorId := "e6a0ffda-34cf-448e-bc96-6045d0825bc1"
			err = errors.New(fmt.Sprintf("unknown  'incomingUnderExecutionTableChannelCommand', '%s'. [ErrorID: %s]", incomingUnderExecutionTableChannelCommand, errorId))

			fmt.Println(err) //TODO Send on Error channel

		}
	}
}
