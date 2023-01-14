package executionsModelForExecutions

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/grpc_out_GuiExecutionServer"
	"errors"
	"fmt"
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

	}

	fmt.Println(getSingleTestCaseExecutionResponse)

	return nil
}
