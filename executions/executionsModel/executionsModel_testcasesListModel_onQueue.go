package executionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/grpc_out_GuiExecutionServer"
	"errors"
	"fmt"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"strconv"
	"time"
)

func (executionsModelObject *ExecutionsModelObjectStruct) LoadAndCreateModelForTestCasesOnExecutionQueue(domainsToInclude []string) (err error) {

	// Prepare message to be sent to GuiExecutionServer to be able to get TestCasesOnExecutionQueue
	var listTestCasesInExecutionQueueRequest *fenixExecutionServerGuiGrpcApi.ListTestCasesInExecutionQueueRequest
	listTestCasesInExecutionQueueRequest = &fenixExecutionServerGuiGrpcApi.ListTestCasesInExecutionQueueRequest{
		UserAndApplicationRunTimeIdentification: &fenixExecutionServerGuiGrpcApi.UserAndApplicationRunTimeIdentificationMessage{
			ApplicationRunTimeUuid: sharedCode.ApplicationRunTimeUuid,
			UserId:                 sharedCode.CurrentUserId,
			ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
				grpc_out_GuiExecutionServer.GetHighestFenixGuiExecutionServerProtoFileVersion()),
		},
		DomainUuids:                    domainsToInclude,
		TestCaseExecutionFromTimeStamp: nil,
		TestCaseExecutionToTimeStamp:   nil,
	}

	// Load TestCases, from GuiExecutionServer, that exists on the TestCaseExecutionQueue
	var listTestCasesInExecutionQueueResponse *fenixExecutionServerGuiGrpcApi.ListTestCasesInExecutionQueueResponse
	listTestCasesInExecutionQueueResponse = grpc_out_GuiExecutionServer.GrpcOutGuiExecutionServerObject.SendListTestCasesOnExecutionQueue(listTestCasesInExecutionQueueRequest)

	if listTestCasesInExecutionQueueResponse.AckNackResponse.AckNack == false {
		ErrorID := "22ea8c42-c185-456e-ad15-dc4ce0e1c05e"
		err = errors.New(fmt.Sprintf("couldn't load TestCaseExecutions on ExecutionQueue from GuiExecutionServer. Got message: '%s'. [ErrorID:'%s']", listTestCasesInExecutionQueueResponse.AckNackResponse.Comments, ErrorID))

		fmt.Println(err) // TODO Send on Error-channel

		return err
	}

	// Save loaded data
	allTestCaseExecutionsOnQueue = allTestCaseExecutionsOnQueueStruct{
		databaseReadTimeStamp:                   time.Now(),
		testCaseExecutionsBelongsToTheseDomains: domainsToInclude,
		testCaseExecutionsOnQueue:               listTestCasesInExecutionQueueResponse.TestCasesInExecutionQueue,
	}

	// Create Model from 'loaded' testCaseExecutions on Queue

	// Initiate map for model
	allTestCaseExecutionsOnQueueModel = make(map[testCaseExecutionMapKeyType]*fenixExecutionServerGuiGrpcApi.TestCaseExecutionBasicInformationMessage)

	// Key to map: Should consist of 'TestCaseExecutionUuid' + 'TestCaseExecutionVersion'
	var testCaseExecutionMapKey testCaseExecutionMapKeyType

	// Loop all TestCasesExecutions that were received from GuiExecutionServer
	for _, tempTestCaseExecutionsOnQueue := range allTestCaseExecutionsOnQueue.testCaseExecutionsOnQueue {

		var testCaseExecutionsOnQueue *fenixExecutionServerGuiGrpcApi.TestCaseExecutionBasicInformationMessage
		*testCaseExecutionsOnQueue = *tempTestCaseExecutionsOnQueue

		// Create Key
		var testCaseExecutionVersionAsString string
		testCaseExecutionVersionAsString = strconv.Itoa(int(tempTestCaseExecutionsOnQueue.TestCaseExecutionVersion))

		testCaseExecutionMapKey = testCaseExecutionMapKeyType(tempTestCaseExecutionsOnQueue.TestCaseExecutionUuid + testCaseExecutionVersionAsString)

		// Add TestCaseExecutions to map
		allTestCaseExecutionsOnQueueModel[testCaseExecutionMapKey] = testCaseExecutionsOnQueue
	}

	return err
}
