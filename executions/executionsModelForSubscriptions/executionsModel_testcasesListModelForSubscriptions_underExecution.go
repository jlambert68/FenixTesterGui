package executionsModelForSubscriptions

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
		databaseReadTimeStamp:                   time.Now().UTC(),
		testCaseExecutionsBelongsToTheseDomains: domainsToInclude,
		testCaseExecutionsUnderExecution:        listTestCasesUnderExecutionResponse.TestCasesUnderExecution,
	}

	// Create Model from 'loaded' testCases under Execution

	// Initiate map for model
	AllTestCaseExecutionsUnderExecutionModel =
		make(map[TestCaseExecutionMapKeyType]*fenixExecutionServerGuiGrpcApi.TestCaseUnderExecutionMessage)

	// Initiate map-model for UI-table-data
	TestCaseExecutionsUnderExecutionMapAdaptedForUiTable =
		make(map[TestCaseExecutionMapKeyType]*TestCaseExecutionsUnderExecutionAdaptedForUiTableStruct)

	// Key to map: Should consist of 'TestCaseExecutionUuid' + 'TestCaseExecutionVersion'
	var testCaseExecutionMapKey TestCaseExecutionMapKeyType
	var existInMap bool

	// Loop all TestCasesExecutions that were received from GuiExecutionServer
	for _, tempTestCaseExecutionsUnderExecution := range allTestCaseExecutionsUnderExecution.testCaseExecutionsUnderExecution {

		var testCaseExecutionsUnderExecution *fenixExecutionServerGuiGrpcApi.TestCaseUnderExecutionMessage
		testCaseExecutionsUnderExecution = tempTestCaseExecutionsUnderExecution

		// Create Key
		var testCaseExecutionVersionAsString string
		testCaseExecutionVersionAsString = strconv.Itoa(int(tempTestCaseExecutionsUnderExecution.TestCaseExecutionBasicInformation.TestCaseExecutionVersion))

		testCaseExecutionMapKey = TestCaseExecutionMapKeyType(tempTestCaseExecutionsUnderExecution.TestCaseExecutionBasicInformation.TestCaseExecutionUuid + testCaseExecutionVersionAsString)

		// Add TestCaseExecution to map
		AllTestCaseExecutionsUnderExecutionModel[testCaseExecutionMapKey] = testCaseExecutionsUnderExecution

		// Convert 'raw' TestCaseExecutionsUnderExecutione-data into format to be used in UI
		var tempTestCaseExecutionUnderExecutionAdaptedForUiTable TestCaseExecutionsUnderExecutionAdaptedForUiTableStruct
		tempTestCaseExecutionUnderExecutionAdaptedForUiTable = TestCaseExecutionsUnderExecutionAdaptedForUiTableStruct{
			// TestCaseExecutionBasicInformation
			DomainUuid:                          tempTestCaseExecutionsUnderExecution.TestCaseExecutionBasicInformation.DomainUuid,
			DomainName:                          tempTestCaseExecutionsUnderExecution.TestCaseExecutionBasicInformation.DomainName,
			TestSuiteUuid:                       tempTestCaseExecutionsUnderExecution.TestCaseExecutionBasicInformation.TestSuiteUuid,
			TestSuiteName:                       tempTestCaseExecutionsUnderExecution.TestCaseExecutionBasicInformation.TestSuiteName,
			TestSuiteVersion:                    strconv.Itoa(int(tempTestCaseExecutionsUnderExecution.TestCaseExecutionBasicInformation.TestSuiteVersion)),
			TestSuiteExecutionUuid:              tempTestCaseExecutionsUnderExecution.TestCaseExecutionBasicInformation.TestSuiteExecutionUuid,
			TestSuiteExecutionVersion:           strconv.Itoa(int(tempTestCaseExecutionsUnderExecution.TestCaseExecutionBasicInformation.TestSuiteExecutionVersion)),
			TestCaseUuid:                        tempTestCaseExecutionsUnderExecution.TestCaseExecutionBasicInformation.TestCaseUuid,
			TestCaseName:                        tempTestCaseExecutionsUnderExecution.TestCaseExecutionBasicInformation.TestCaseName,
			TestCaseVersion:                     strconv.Itoa(int(tempTestCaseExecutionsUnderExecution.TestCaseExecutionBasicInformation.TestCaseVersion)),
			TestCaseExecutionUuid:               tempTestCaseExecutionsUnderExecution.TestCaseExecutionBasicInformation.TestCaseExecutionUuid,
			TestCaseExecutionVersion:            strconv.Itoa(int(tempTestCaseExecutionsUnderExecution.TestCaseExecutionBasicInformation.TestCaseExecutionVersion)),
			PlacedOnTestExecutionQueueTimeStamp: tempTestCaseExecutionsUnderExecution.TestCaseExecutionBasicInformation.PlacedOnTestExecutionQueueTimeStamp.AsTime().String(),
			ExecutionPriority:                   fenixExecutionServerGuiGrpcApi.ExecutionPriorityEnum_name[int32(tempTestCaseExecutionsUnderExecution.TestCaseExecutionBasicInformation.ExecutionPriority)],

			// TestCaseExecutionDetails
			ExecutionStartTimeStamp:        tempTestCaseExecutionsUnderExecution.TestCaseExecutionDetails.ExecutionStartTimeStamp.AsTime().String(),
			ExecutionStopTimeStamp:         tempTestCaseExecutionsUnderExecution.TestCaseExecutionDetails.ExecutionStopTimeStamp.AsTime().String(),
			TestCaseExecutionStatus:        fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum_name[int32(tempTestCaseExecutionsUnderExecution.TestCaseExecutionDetails.TestCaseExecutionStatus)],
			ExecutionHasFinished:           strconv.FormatBool(tempTestCaseExecutionsUnderExecution.TestCaseExecutionDetails.ExecutionHasFinished),
			ExecutionStatusUpdateTimeStamp: tempTestCaseExecutionsUnderExecution.TestCaseExecutionDetails.ExecutionStatusUpdateTimeStamp.AsTime().String(),
		}

		// Verify that key is not already used in map
		_, existInMap = TestCaseExecutionsUnderExecutionMapAdaptedForUiTable[testCaseExecutionMapKey]
		if existInMap == true {

			errorId := "c7cbc0ed-97d7-48ba-bff4-2af7dd9d13b5"
			err = errors.New(fmt.Sprintf("'testCaseExecutionMapKey', '%s' already exist in TestCaseExecutionsUnderExecutionMapAdaptedForUiTable [ErrorID: %s]", testCaseExecutionMapKey, errorId))

			fmt.Println(err) // TODO Send on Error Channel

			return err
		}

		// Append to map for TestCaseExecutionsUnderExecution-data used by UI-table
		TestCaseExecutionsUnderExecutionMapAdaptedForUiTable[testCaseExecutionMapKey] = &tempTestCaseExecutionUnderExecutionAdaptedForUiTable

	}

	return err
}
