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

// LoadAndCreateModelForTestCaseWithFinishedExecutions - Load TestCaseExecutions that have been finished their Executions and transform them into model used
func (executionsModelObject *ExecutionsModelObjectStruct) LoadAndCreateModelForTestCaseWithFinishedExecutions(domainsToInclude []string) (err error) {

	// Prepare message to be sent to GuiExecutionServer to be able to get Ongoing TestCaseExecutions
	var listTestCasesWithFinishedExecutionsRequest *fenixExecutionServerGuiGrpcApi.ListTestCasesWithFinishedExecutionsRequest
	listTestCasesWithFinishedExecutionsRequest = &fenixExecutionServerGuiGrpcApi.ListTestCasesWithFinishedExecutionsRequest{
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

	// Load TestCases, from GuiExecutionServer, that have been finished
	var listTestCasesWithFinishedExecutionsResponse *fenixExecutionServerGuiGrpcApi.ListTestCasesWithFinishedExecutionsResponse
	listTestCasesWithFinishedExecutionsResponse = grpc_out_GuiExecutionServer.GrpcOutGuiExecutionServerObject.SendListTestCasesWithFinishedExecutions(listTestCasesWithFinishedExecutionsRequest)

	if listTestCasesWithFinishedExecutionsResponse.AckNackResponse.AckNack == false {
		ErrorID := "6a45f94c-3ed1-4e08-84c8-6236c9e4a02f"
		err = errors.New(fmt.Sprintf("couldn't load TestCaseExecutions with finished Executions from GuiExecutionServer. Got message: '%s'. [ErrorID:'%s']", listTestCasesWithFinishedExecutionsResponse.AckNackResponse.Comments, ErrorID))

		fmt.Println(err) // TODO Send on Error-channel

		return err
	}

	// Save loaded data
	allTestCaseExecutionFinishedExecutions = allTestCaseExecutionsThatHaveBeenFinishedExecutedStruct{
		databaseReadTimeStamp:                    time.Now(),
		testCaseExecutionsBelongsToTheseDomains:  domainsToInclude,
		testCaseExecutionsWithFinishedExecutions: listTestCasesWithFinishedExecutionsResponse.TestCaseWithFinishedExecution,
	}

	// Create Model from 'loaded' testCases with finished Executions

	// Initiate map for model
	allTestCaseExecutionsThatHaveBeenFinishedExecutedModel = make(map[testCaseExecutionMapKeyType]*fenixExecutionServerGuiGrpcApi.TestCaseWithFinishedExecutionMessage)

	// Key to map: Should consist of 'TestCaseExecutionUuid' + 'TestCaseExecutionVersion'
	var testCaseExecutionMapKey testCaseExecutionMapKeyType

	// Loop all TestCasesExecutions that were received from GuiExecutionServer
	for _, tempTestCaseExecutionWithFinishedExecution := range allTestCaseExecutionFinishedExecutions.testCaseExecutionsWithFinishedExecutions {

		var testCaseWithFinishedExecutionMessage *fenixExecutionServerGuiGrpcApi.TestCaseWithFinishedExecutionMessage
		testCaseWithFinishedExecutionMessage = tempTestCaseExecutionWithFinishedExecution

		// Create Key
		var testCaseExecutionVersionAsString string
		testCaseExecutionVersionAsString = strconv.Itoa(int(tempTestCaseExecutionWithFinishedExecution.TestCaseExecutionBasicInformation.TestCaseExecutionVersion))

		testCaseExecutionMapKey = testCaseExecutionMapKeyType(tempTestCaseExecutionWithFinishedExecution.TestCaseExecutionBasicInformation.TestCaseExecutionUuid + testCaseExecutionVersionAsString)

		// Add TestCaseExecution to map
		allTestCaseExecutionsThatHaveBeenFinishedExecutedModel[testCaseExecutionMapKey] = testCaseWithFinishedExecutionMessage
	}

	return err
}
