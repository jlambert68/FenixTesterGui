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

// LoadAndCreateModelForTestCasesOnExecutionQueue - Load TestCaseExecutions that waits on ExecutionQueue and transform them into model used
func (executionsModelObject *ExecutionsModelObjectStruct) LoadAndCreateModelForTestCasesOnExecutionQueue(domainsToInclude []string) (err error) {

	// Prepare message to be sent to GuiExecutionServer to be able to get TestCasesOnExecutionQueue
	var listTestCasesInExecutionQueueRequest *fenixExecutionServerGuiGrpcApi.ListTestCasesInExecutionQueueRequest
	listTestCasesInExecutionQueueRequest = &fenixExecutionServerGuiGrpcApi.ListTestCasesInExecutionQueueRequest{
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

	// Load TestCasesMapPtr, from GuiExecutionServer, that exists on the TestCaseExecutionQueue
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
		databaseReadTimeStamp:                   time.Now().UTC(),
		testCaseExecutionsBelongsToTheseDomains: domainsToInclude,
		testCaseExecutionsOnQueue:               listTestCasesInExecutionQueueResponse.TestCasesInExecutionQueue,
	}

	// Create Model from 'loaded' testCaseExecutions on Queue

	// Initiate map for model
	AllTestCaseExecutionsOnQueueModel =
		make(map[TestCaseExecutionMapKeyType]*fenixExecutionServerGuiGrpcApi.TestCaseExecutionBasicInformationMessage)

	// Initiate map-model for UI-table-data
	TestCaseExecutionsOnQueueMapAdaptedForUiTable =
		make(map[TestCaseExecutionMapKeyType]*TestCaseExecutionsOnQueueAdaptedForUiTableStruct)

	// Key to map: Should consist of 'TestCaseExecutionUuid' + 'TestCaseExecutionVersion'
	var testCaseExecutionMapKey TestCaseExecutionMapKeyType
	var existInMap bool

	// Loop all TestCasesExecutions that were received from GuiExecutionServer
	for _, tempTestCaseExecutionsOnQueue := range allTestCaseExecutionsOnQueue.testCaseExecutionsOnQueue {

		var testCaseExecutionsOnQueue *fenixExecutionServerGuiGrpcApi.TestCaseExecutionBasicInformationMessage
		testCaseExecutionsOnQueue = tempTestCaseExecutionsOnQueue

		// Create Key
		var testCaseExecutionVersionAsString string
		testCaseExecutionVersionAsString = strconv.Itoa(int(tempTestCaseExecutionsOnQueue.TestCaseExecutionVersion))

		testCaseExecutionMapKey = TestCaseExecutionMapKeyType(tempTestCaseExecutionsOnQueue.TestCaseExecutionUuid + testCaseExecutionVersionAsString)

		// Add TestCaseExecutions to map
		AllTestCaseExecutionsOnQueueModel[testCaseExecutionMapKey] = testCaseExecutionsOnQueue

		// Convert 'raw' TestCaseExecutionsOnQueue-data into format to be used in UI
		var tempTestCaseExecutionsOnQueueAdaptedForUiTable TestCaseExecutionsOnQueueAdaptedForUiTableStruct
		tempTestCaseExecutionsOnQueueAdaptedForUiTable = TestCaseExecutionsOnQueueAdaptedForUiTableStruct{
			ShowDetailedTestCaseExecution: "false",

			DomainUuid:                          tempTestCaseExecutionsOnQueue.DomainUuid,
			DomainName:                          tempTestCaseExecutionsOnQueue.DomainName,
			TestSuiteUuid:                       tempTestCaseExecutionsOnQueue.TestSuiteUuid,
			TestSuiteName:                       tempTestCaseExecutionsOnQueue.TestSuiteName,
			TestSuiteVersion:                    strconv.Itoa(int(tempTestCaseExecutionsOnQueue.TestSuiteVersion)),
			TestSuiteExecutionUuid:              tempTestCaseExecutionsOnQueue.TestSuiteExecutionUuid,
			TestSuiteExecutionVersion:           strconv.Itoa(int(tempTestCaseExecutionsOnQueue.TestSuiteExecutionVersion)),
			TestCaseUuid:                        tempTestCaseExecutionsOnQueue.TestCaseUuid,
			TestCaseName:                        tempTestCaseExecutionsOnQueue.TestCaseName,
			TestCaseVersion:                     strconv.Itoa(int(tempTestCaseExecutionsOnQueue.TestCaseVersion)),
			TestCaseExecutionUuid:               tempTestCaseExecutionsOnQueue.TestCaseExecutionUuid,
			TestCaseExecutionVersion:            strconv.Itoa(int(tempTestCaseExecutionsOnQueue.TestCaseExecutionVersion)),
			PlacedOnTestExecutionQueueTimeStamp: tempTestCaseExecutionsOnQueue.PlacedOnTestExecutionQueueTimeStamp.AsTime().String(),
			ExecutionPriority:                   fenixExecutionServerGuiGrpcApi.ExecutionPriorityEnum_name[int32(tempTestCaseExecutionsOnQueue.ExecutionPriority)],
		}

		// Verify that key is not already used in map
		_, existInMap = TestCaseExecutionsOnQueueMapAdaptedForUiTable[testCaseExecutionMapKey]
		if existInMap == true {

			errorId := "e5fba95a-2e69-4428-ac7d-a85c9a1819ac"
			err = errors.New(fmt.Sprintf("'testCaseExecutionMapKey', '%s' already exist in TestCaseExecutionsOnQueueMapAdaptedForUiTable [ErrorID: %s]", testCaseExecutionMapKey, errorId))

			fmt.Println(err) // TODO Send on Error Channel

			return err
		}

		// Append to map for TestCaseExecutionsUnderExecution-data used by UI-table
		TestCaseExecutionsOnQueueMapAdaptedForUiTable[testCaseExecutionMapKey] = &tempTestCaseExecutionsOnQueueAdaptedForUiTable

	}

	return err
}
