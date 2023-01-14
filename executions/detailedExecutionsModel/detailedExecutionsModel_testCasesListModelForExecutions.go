package detailedExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/executions/executionsModelForSubscriptions"
	"FenixTesterGui/grpc_out_GuiExecutionServer"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"strconv"
)

// Initiate the channels used when Adding or Removing items to/from OnQueue-table, UnderExecution-table or FinishedExecutions-table
func InitiateAndStartChannelsUsedByListModel() {

	// Initiate Channel used for Adding and Deleting Execution items in OnQueue-table
	//OnQueueTableAddRemoveChannel = make(chan OnQueueTableAddRemoveChannelStruct, MaximumNumberOfItemsForOnQueueTableAddRemoveChannel)

	// Initiate Channel used for Adding and Deleting Execution items in UnderExecution-table
	//UnderExecutionTableAddRemoveChannel = make(chan UnderExecutionTableAddRemoveChannelStruct, MaximumNumberOfItemsForUnderExecutionTableAddRemoveChannel)

	// Initiate Channel used for Adding and Deleting Execution items in FinishedExecutions-table
	//FinishedExecutionsTableAddRemoveChannel = make(chan FinishedExecutionsTableAddRemoveChannelStruct, MaximumNumberOfItemsForFinishedExecutionsTableAddRemoveChannel)

}

// RetrieveSingleTestCaseExecution
// Retrieves a TestCaseExecution and all of its data belonging to the execution
func RetrieveSingleTestCaseExecution(testCaseExecutionKey string) (err error) {

	// Exctract individual parts of the 'TestCaseExecutionKeyMessage'
	var testCaseExecutionUuid string
	var testCaseExecutionVersion int
	var testCaseExecutionVersionError error

	testCaseExecutionUuid = testCaseExecutionKey[:len(testCaseExecutionKey)-1]

	testCaseExecutionVersion, testCaseExecutionVersionError = strconv.Atoi(testCaseExecutionKey[len(testCaseExecutionKey)-1:])
	if testCaseExecutionVersionError != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"Id":                   "a7e1c59a-5e5f-47f2-ba7d-1a909eb90d68",
			"testCaseExecutionKey": testCaseExecutionKey,
			"testCaseExecutionKey[len(testCaseExecutionKey):]": testCaseExecutionKey[len(testCaseExecutionKey):],
		}).Error("Couldn't convert 'TestCaseExecutionVersion' from TestCaseExecutionKey into an integer")

		return testCaseExecutionVersionError
	}

	var testCaseExecutionKeyMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionKeyMessage
	testCaseExecutionKeyMessage = &fenixExecutionServerGuiGrpcApi.TestCaseExecutionKeyMessage{
		TestCaseExecutionUuid:    testCaseExecutionUuid,
		TestCaseExecutionVersion: uint32(testCaseExecutionVersion),
	}

	var getSingleTestCaseExecutionRequest *fenixExecutionServerGuiGrpcApi.GetSingleTestCaseExecutionRequest
	getSingleTestCaseExecutionRequest = &fenixExecutionServerGuiGrpcApi.GetSingleTestCaseExecutionRequest{
		UserAndApplicationRunTimeIdentification: &fenixExecutionServerGuiGrpcApi.UserAndApplicationRunTimeIdentificationMessage{
			ApplicationRunTimeUuid: sharedCode.ApplicationRunTimeUuid,
			UserId:                 sharedCode.CurrentUserId,
			ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
				grpc_out_GuiExecutionServer.GetHighestFenixGuiExecutionServerProtoFileVersion()),
		},
		TestCaseExecutionKey: testCaseExecutionKeyMessage,
	}

	// Do gRPC-call
	var getSingleTestCaseExecutionResponse *fenixExecutionServerGuiGrpcApi.GetSingleTestCaseExecutionResponse
	getSingleTestCaseExecutionResponse = grpc_out_GuiExecutionServer.GrpcOutGuiExecutionServerObject.
		SendGetSingleTestCaseExecution(getSingleTestCaseExecutionRequest)

	if getSingleTestCaseExecutionResponse.AckNackResponse.AckNack == false {
		return errors.New(getSingleTestCaseExecutionResponse.AckNackResponse.Comments)
	} else {
		// Add TestCaseExecution-details to repository
		var testCaseExecutionResponse *fenixExecutionServerGuiGrpcApi.TestCaseExecutionResponseMessage
		testCaseExecutionResponse = getSingleTestCaseExecutionResponse.TestCaseExecutionResponse

		// The definition used in SummaryTable to represent one TestCaseExecution and its current execution status
		type TestCaseExecutionsStatusForSummaryTableStruct struct {
			TestCaseUIName      string
			TestCaseStatusValue uint32
		}

		// The definition used in SummaryTable to represent one TestInstructionExecution and its current execution status
		type TestInstructionExecutionsStatusForSummaryTableStruct struct {
			TestInstructionExecutionUIName string
			TestInstructionStatusValue     uint32
		}

		// One TestCaseExecution and all of its data.
		type TestCaseExecutionsDetailsStruct struct {
			// The response message when a full TestCaseExecution is retrieved
			TestCaseExecutionDatabaseResponseMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionResponseMessage

			// The streamed status messages
			TestCaseExecutionsStatusUpdates        []*fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusMessage
			TestInstructionExecutionsStatusUpdates []*fenixExecutionServerGuiGrpcApi.TestInstructionExecutionStatusMessage

			// A map holding all TestInstructions with their execution status. Each slice is sorted by 'UniqueDatabaseRowCounter' ASC order
			// The slice data is used to show execution status and the last item in the slice is the one that has the current status
			// map[TestInstructionExecutionKey]*[]*fenixExecutionServerGuiGrpcApi.TestInstructionExecutionsInformationMessage
			// TestInstructionExecutionKey = TestInstructionExecutionUuid + TestInstructionExecutionVersion
			TestInstructionExecutionsStatusMap map[string]*[]*fenixExecutionServerGuiGrpcApi.TestInstructionExecutionsInformationMessage

			// Holding the information to be show in the SummaryTable for one TestCaseExecution
			TestCaseExecutionsStatusForSummaryTable TestCaseExecutionsStatusForSummaryTableStruct

			// The slice of all TestInstructionExecution, for one TestCaseExecution, and their current status. The order is the same as it is presented on screen
			TestInstructionExecutionsStatusForSummaryTable []TestInstructionExecutionsStatusForSummaryTableStruct
		}

		// TestCaseExecutionsDetailsMap
		// map[TestCaseExecutionMapKey]*TestCaseExecutionsDetailsStruct, TestCaseExecutionMapKey = TestCaseExecutionUuid + TestCaseExecutionVersionNumber
		var TestCaseExecutionsDetailsMap map[string]*TestCaseExecutionsDetailsStruct // m
		TestCaseExecutionsDetailsMap = make(map[string]*TestCaseExecutionsDetailsStruct)

		// Check if TestCaseExecution already exist
		var existInMap bool
		var testCaseExecutionsDetails *TestCaseExecutionsDetailsStruct
		var testInstructionExecutionsStatusMap map[string]*[]*fenixExecutionServerGuiGrpcApi.TestInstructionExecutionsInformationMessage
		testInstructionExecutionsStatusMap = make(map[string]*[]*fenixExecutionServerGuiGrpcApi.TestInstructionExecutionsInformationMessage)

		testCaseExecutionsDetails, existInMap = TestCaseExecutionsDetailsMap[testCaseExecutionKey]
		// If TestExecutionExecution doesn't exist in map then create a new instance
		if existInMap == false {
			testCaseExecutionsDetails = &TestCaseExecutionsDetailsStruct{
				TestCaseExecutionDatabaseResponseMessage:       testCaseExecutionResponse,
				TestCaseExecutionsStatusUpdates:                nil,
				TestInstructionExecutionsStatusUpdates:         nil,
				TestInstructionExecutionsStatusMap:             testInstructionExecutionsStatusMap,
				TestCaseExecutionsStatusForSummaryTable:        TestCaseExecutionsStatusForSummaryTableStruct{},
				TestInstructionExecutionsStatusForSummaryTable: nil,
			}

			// Add the TestCaseExecution to the Map
			TestCaseExecutionsDetailsMap[testCaseExecutionKey] = testCaseExecutionsDetails
		}

	}

	fmt.Println(getSingleTestCaseExecutionResponse)

	return nil
}

func CreateTableForDetailedTestCaseExecutionsList() *fyne.Container {
	var tableForTestCaseExecutionsOnQueueBindings []binding.DataMap

	// Create a binding for each TestExecutionOnQueueRow data
	for _, tempTestCaseExecutionsOnQueueDataAdaptedForUiTableReference := range executionsModelForSubscriptions.TestCaseExecutionsOnQueueMapAdaptedForUiTable {
		tableForTestCaseExecutionsOnQueueBindings = append(
			tableForTestCaseExecutionsOnQueueBindings,
			binding.BindStruct(tempTestCaseExecutionsOnQueueDataAdaptedForUiTableReference))
	}

	executionsModelForSubscriptions.TestCaseExecutionsOnQueueTableOptions.Bindings = tableForTestCaseExecutionsOnQueueBindings

	ht := NewTestCaseExecutionsSummaryTable(&executionsModelForSubscriptions.TestCaseExecutionsOnQueueTableOptions)
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
