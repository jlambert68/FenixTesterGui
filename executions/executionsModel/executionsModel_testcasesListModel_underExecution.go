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

// LoadAndCreateModelForTestCaseUnderExecutions - Load TestCaseExecutions that are Under Execution and transform them into model used
func (executionsModelObject *ExecutionsModelObjectStruct) LoadAndCreateModelForTestCaseUnderExecutions(domainsToInclude []string) (err error) {

	// Prepare message to be sent to GuiExecutionServer to be able to get Ongoing TestCaseExecutions
	var listTestCasesUnderExecutionRequest *fenixExecutionServerGuiGrpcApi.ListTestCasesUnderExecutionRequest
	listTestCasesUnderExecutionRequest = &fenixExecutionServerGuiGrpcApi.ListTestCasesUnderExecutionRequest{
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
	var listTestCasesUnderExecutionResponse *fenixExecutionServerGuiGrpcApi.ListTestCasesUnderExecutionResponse
	listTestCasesUnderExecutionResponse = grpc_out_GuiExecutionServer.GrpcOutGuiExecutionServerObject.SendListTestCasesUnderExecution(listTestCasesUnderExecutionRequest)

	if listTestCasesUnderExecutionResponse.AckNackResponse.AckNack == false {
		ErrorID := "160b27ac-2687-4385-8dee-4fb129c10af6"
		err = errors.New(fmt.Sprintf("couldn't load TestCaseExecutions under Execution from GuiExecutionServer. Got message: '%s'. [ErrorID:'%s']", listTestCasesUnderExecutionResponse.AckNackResponse.Comments, ErrorID))

		fmt.Println(err) // TODO Send on Error-channel

		return err
	}

	// Save loaded data
	allTestCaseExecutionsUnderExecution = allTestCaseExecutionsOngoingUnderExecutionStruct{
		databaseReadTimeStamp:                   time.Now(),
		testCaseExecutionsBelongsToTheseDomains: domainsToInclude,
		testCaseExecutionsUnderExecution:        listTestCasesUnderExecutionResponse.TestCasesUnderExecution,
	}

	// Create Model from 'loaded' testCases under Execution

	// Initiate map for model
	allTestCaseExecutionsUnderExecutionModel = make(map[testCaseExecutionMapKeyType]*fenixExecutionServerGuiGrpcApi.TestCaseUnderExecutionMessage)

	// Key to map: Should consist of 'TestCaseExecutionUuid' + 'TestCaseExecutionVersion'
	var testCaseExecutionMapKey testCaseExecutionMapKeyType

	// Loop all TestCasesExecutions that were received from GuiExecutionServer
	for _, tempTestCaseExecutionsUnderExecution := range allTestCaseExecutionsUnderExecution.testCaseExecutionsUnderExecution {

		var testCaseExecutionsUnderExecution *fenixExecutionServerGuiGrpcApi.TestCaseUnderExecutionMessage
		testCaseExecutionsUnderExecution = tempTestCaseExecutionsUnderExecution

		// Create Key
		var testCaseExecutionVersionAsString string
		testCaseExecutionVersionAsString = strconv.Itoa(int(tempTestCaseExecutionsUnderExecution.TestCaseExecutionBasicInformation.TestCaseExecutionVersion))

		testCaseExecutionMapKey = testCaseExecutionMapKeyType(tempTestCaseExecutionsUnderExecution.TestCaseExecutionBasicInformation.TestCaseExecutionUuid + testCaseExecutionVersionAsString)

		// Add TestCaseExecution to map
		allTestCaseExecutionsUnderExecutionModel[testCaseExecutionMapKey] = testCaseExecutionsUnderExecution
	}

	return err
}
