package executionsModelForExecutions

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/grpc_out_GuiExecutionServer"
	"errors"
	"fmt"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"strconv"
	"time"
)

// LoadAndCreateModelForTestCaseFinishedExecutions - Load TestCaseExecutions that have been finished Execution and transform them into model used
func (executionsModelObject *ExecutionsModelObjectStruct) LoadAndCreateModelForTestCaseWithFinishedExecutions(domainsToInclude []string) (err error) {

	// Prepare message to be sent to GuiExecutionServer to be able to get Ongoing TestCaseExecutions
	var listTestCasesWithFinishedExecutionsRequest *fenixExecutionServerGuiGrpcApi.ListTestCasesWithFinishedExecutionsRequest
	listTestCasesWithFinishedExecutionsRequest = &fenixExecutionServerGuiGrpcApi.ListTestCasesWithFinishedExecutionsRequest{
		UserAndApplicationRunTimeIdentification: &fenixExecutionServerGuiGrpcApi.UserAndApplicationRunTimeIdentificationMessage{
			ApplicationRunTimeUuid: sharedCode.ApplicationRunTimeUuid,
			UserIdOnComputer:       sharedCode.CurrentUserIdLogedInOnComputer,
			GCPAuthenticatedUser:   sharedCode.CurrentUserAuthenticatedTowardsGCP,
			ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
				grpc_out_GuiExecutionServer.GetHighestFenixGuiExecutionServerProtoFileVersion()),
		},
		DomainUuids:                    domainsToInclude,
		TestCaseExecutionFromTimeStamp: nil,
		TestCaseExecutionToTimeStamp:   nil,
	}

	// Load TestCases, from GuiExecutionServer, that exists on the TestCaseExecutionQueue
	var listTestCasesWithFinishedExecutionsResponse *fenixExecutionServerGuiGrpcApi.ListTestCasesWithFinishedExecutionsResponse
	listTestCasesWithFinishedExecutionsResponse = grpc_out_GuiExecutionServer.GrpcOutGuiExecutionServerObject.SendListTestCasesWithFinishedExecutions(listTestCasesWithFinishedExecutionsRequest)

	if listTestCasesWithFinishedExecutionsResponse.AckNackResponse.AckNack == false {
		ErrorID := "e43d3385-dc3e-426e-98fb-13a86d6375aa"
		err = errors.New(fmt.Sprintf("couldn't load TestCaseExecutions that have been finsihed its Execution, from GuiExecutionServer. Got message: '%s'. [ErrorID:'%s']", listTestCasesWithFinishedExecutionsResponse.AckNackResponse.Comments, ErrorID))

		fmt.Println(err) // TODO Send on Error-channel

		return err
	}

	// Save loaded data
	allTestCaseExecutionsFinishedExecution = allTestCaseExecutionsOngoingFinishedExecutionStruct{
		databaseReadTimeStamp:                   time.Now().UTC(),
		testCaseExecutionsBelongsToTheseDomains: domainsToInclude,
		testCaseExecutionsFinishedExecution:     listTestCasesWithFinishedExecutionsResponse.TestCaseWithFinishedExecution,
	}

	// Create Model from 'loaded' testCases under Execution

	// Initiate map for model
	AllTestCaseExecutionsFinishedExecutionModel =
		make(map[TestCaseExecutionMapKeyType]*fenixExecutionServerGuiGrpcApi.TestCaseWithFinishedExecutionMessage)

	// Initiate map-model for UI-table-data
	TestCaseExecutionsFinishedExecutionMapAdaptedForUiTable =
		make(map[TestCaseExecutionMapKeyType]*TestCaseExecutionsFinishedExecutionAdaptedForUiTableStruct)

	// Key to map: Should consist of 'TestCaseExecutionUuid' + 'TestCaseExecutionVersion'
	var testCaseExecutionMapKey TestCaseExecutionMapKeyType
	var existInMap bool

	// Loop all TestCasesExecutions that were received from GuiExecutionServer
	for _, tempTestCaseExecutionsFinishedExecution := range allTestCaseExecutionsFinishedExecution.testCaseExecutionsFinishedExecution {

		var testCaseExecutionsFinishedExecution *fenixExecutionServerGuiGrpcApi.TestCaseWithFinishedExecutionMessage
		testCaseExecutionsFinishedExecution = tempTestCaseExecutionsFinishedExecution

		// Create Key
		var testCaseExecutionVersionAsString string
		testCaseExecutionVersionAsString = strconv.Itoa(int(tempTestCaseExecutionsFinishedExecution.TestCaseExecutionBasicInformation.TestCaseExecutionVersion))

		testCaseExecutionMapKey = TestCaseExecutionMapKeyType(tempTestCaseExecutionsFinishedExecution.TestCaseExecutionBasicInformation.TestCaseExecutionUuid + testCaseExecutionVersionAsString)

		// Add TestCaseExecution to map
		AllTestCaseExecutionsFinishedExecutionModel[testCaseExecutionMapKey] = testCaseExecutionsFinishedExecution

		// Convert 'raw' TestCaseExecutionsFinishedExecutione-data into format to be used in UI
		var tempTestCaseExecutionFinishedExecutionAdaptedForUiTable TestCaseExecutionsFinishedExecutionAdaptedForUiTableStruct
		tempTestCaseExecutionFinishedExecutionAdaptedForUiTable = TestCaseExecutionsFinishedExecutionAdaptedForUiTableStruct{
			ShowDetailedTestCaseExecution: "false",

			// TestCaseExecutionBasicInformation
			DomainUuid:                          tempTestCaseExecutionsFinishedExecution.TestCaseExecutionBasicInformation.DomainUuid,
			DomainName:                          tempTestCaseExecutionsFinishedExecution.TestCaseExecutionBasicInformation.DomainName,
			TestSuiteUuid:                       tempTestCaseExecutionsFinishedExecution.TestCaseExecutionBasicInformation.TestSuiteUuid,
			TestSuiteName:                       tempTestCaseExecutionsFinishedExecution.TestCaseExecutionBasicInformation.TestSuiteName,
			TestSuiteVersion:                    strconv.Itoa(int(tempTestCaseExecutionsFinishedExecution.TestCaseExecutionBasicInformation.TestSuiteVersion)),
			TestSuiteExecutionUuid:              tempTestCaseExecutionsFinishedExecution.TestCaseExecutionBasicInformation.TestSuiteExecutionUuid,
			TestSuiteExecutionVersion:           strconv.Itoa(int(tempTestCaseExecutionsFinishedExecution.TestCaseExecutionBasicInformation.TestSuiteExecutionVersion)),
			TestCaseUuid:                        tempTestCaseExecutionsFinishedExecution.TestCaseExecutionBasicInformation.TestCaseUuid,
			TestCaseName:                        tempTestCaseExecutionsFinishedExecution.TestCaseExecutionBasicInformation.TestCaseName,
			TestCaseVersion:                     strconv.Itoa(int(tempTestCaseExecutionsFinishedExecution.TestCaseExecutionBasicInformation.TestCaseVersion)),
			TestCaseExecutionUuid:               tempTestCaseExecutionsFinishedExecution.TestCaseExecutionBasicInformation.TestCaseExecutionUuid,
			TestCaseExecutionVersion:            strconv.Itoa(int(tempTestCaseExecutionsFinishedExecution.TestCaseExecutionBasicInformation.TestCaseExecutionVersion)),
			PlacedOnTestExecutionQueueTimeStamp: tempTestCaseExecutionsFinishedExecution.TestCaseExecutionBasicInformation.PlacedOnTestExecutionQueueTimeStamp.AsTime().String(),
			ExecutionPriority:                   fenixExecutionServerGuiGrpcApi.ExecutionPriorityEnum_name[int32(tempTestCaseExecutionsFinishedExecution.TestCaseExecutionBasicInformation.ExecutionPriority)],

			// TestCaseExecutionDetails
			ExecutionStartTimeStamp:        tempTestCaseExecutionsFinishedExecution.TestCaseExecutionDetails.ExecutionStartTimeStamp.AsTime().String(),
			ExecutionStopTimeStamp:         tempTestCaseExecutionsFinishedExecution.TestCaseExecutionDetails.ExecutionStopTimeStamp.AsTime().String(),
			TestCaseExecutionStatus:        fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum_name[int32(tempTestCaseExecutionsFinishedExecution.TestCaseExecutionDetails.TestCaseExecutionStatus)],
			ExecutionHasFinished:           strconv.FormatBool(tempTestCaseExecutionsFinishedExecution.TestCaseExecutionDetails.ExecutionHasFinished),
			ExecutionStatusUpdateTimeStamp: tempTestCaseExecutionsFinishedExecution.TestCaseExecutionDetails.ExecutionStatusUpdateTimeStamp.AsTime().String(),
		}

		// Verify that key is not already used in map
		_, existInMap = TestCaseExecutionsFinishedExecutionMapAdaptedForUiTable[testCaseExecutionMapKey]
		if existInMap == true {

			errorId := "b32b17bd-f4cc-40fe-9181-3e5f0284bc88"
			err = errors.New(fmt.Sprintf("'testCaseExecutionMapKey', '%s' already exist in TestCaseExecutionsFinishedExecutionMapAdaptedForUiTable [ErrorID: %s]", testCaseExecutionMapKey, errorId))

			fmt.Println(err) // TODO Send on Error Channel

			return err
		}

		// Append to map for TestCaseExecutionsUnderExecution-data used by UI-table
		TestCaseExecutionsFinishedExecutionMapAdaptedForUiTable[testCaseExecutionMapKey] = &tempTestCaseExecutionFinishedExecutionAdaptedForUiTable

	}

	return err
}
