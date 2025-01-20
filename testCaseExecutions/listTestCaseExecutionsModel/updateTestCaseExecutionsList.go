package listTestCaseExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/grpc_out_GuiExecutionServer"
	"FenixTesterGui/testCaseExecutions/testCaseExecutionsModel"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"time"
)

// LoadTestCaseExecutionsThatCanBeViewedByUser
// Load list with TestCaseExecutions that the user can view
func LoadTestCaseExecutionsThatCanBeViewedByUser(
	latestUniqueTestCaseExecutionDatabaseRowId int32,
	onlyRetrieveLimitedSizedBatch bool,
	batchSize int32,
	testCaseExecutionFromTimeStamp time.Time,
	testCaseExecutionToTimeStamp time.Time,
	loadAllDataFromDatabase bool) {

	var listTestCaseExecutionsResponse *fenixExecutionServerGuiGrpcApi.ListTestCaseExecutionsResponse

	if loadAllDataFromDatabase == false {

		listTestCaseExecutionsResponse = grpc_out_GuiExecutionServer.GrpcOutGuiExecutionServerObject.
			SendListTestCaseExecutionsThatCanBeViewed(
				latestUniqueTestCaseExecutionDatabaseRowId,
				onlyRetrieveLimitedSizedBatch,
				batchSize,
				testCaseExecutionFromTimeStamp,
				testCaseExecutionToTimeStamp)

		if listTestCaseExecutionsResponse.GetAckNackResponse().AckNack == false {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":    "320c6409-a68b-4cf0-adc1-aa65d8c51343",
				"error": listTestCaseExecutionsResponse.GetAckNackResponse().Comments,
			}).Warning("Problem to do gRPC-call to FenixGuiExecutionServer in 'LoadTestCaseExecutionsThatCanBeViewedByUser'")

			return
		}

		// Store the slice with TestCases that a user can edit as a Map
		storeTestCaseExecutionsThatCanBeViewedByUser(
			listTestCaseExecutionsResponse.GetTestCaseExecutionsList(),
			&testCaseExecutionsModel.TestCaseExecutionsModel)

	} else {

		listTestCaseExecutionsResponse = grpc_out_GuiExecutionServer.GrpcOutGuiExecutionServerObject.
			SendListTestCaseExecutionsThatCanBeViewed(
				latestUniqueTestCaseExecutionDatabaseRowId,
				onlyRetrieveLimitedSizedBatch,
				batchSize,
				testCaseExecutionFromTimeStamp,
				testCaseExecutionToTimeStamp)

		if listTestCaseExecutionsResponse.GetAckNackResponse().AckNack == false {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":    "320c6409-a68b-4cf0-adc1-aa65d8c51343",
				"error": listTestCaseExecutionsResponse.GetAckNackResponse().Comments,
			}).Warning("Problem to do gRPC-call to FenixGuiExecutionServer in 'LoadTestCaseExecutionsThatCanBeViewedByUser'")

			return
		}

		// Store the slice with TestCases that a user can edit as a Map
		storeTestCaseExecutionsThatCanBeViewedByUser(
			listTestCaseExecutionsResponse.GetTestCaseExecutionsList(),
			&testCaseExecutionsModel.TestCaseExecutionsModel)

	}

	// Store the slice with TestCases
	/*
		//testCaseModeReference.TestCasesThatCanBeEditedByUserSlice = listTestCaseExecutionsResponse.GetTestCasesThatCanBeEditedByUser()
		testCaseModeReference.TestCasesThatCanBeEditedByUserSlice = nil
		for _, tempTestCasesThatCanBeEditedByUser := range testCaseModeReference.TestCasesThatCanBeEditedByUserMap {
			testCaseModeReference.TestCasesThatCanBeEditedByUserSlice = append(
				testCaseModeReference.TestCasesThatCanBeEditedByUserSlice, tempTestCasesThatCanBeEditedByUser)
		}


	*/

}

// Store TestCaseExecutions That Can Be Viewed By User
func storeTestCaseExecutionsThatCanBeViewedByUser(
	testCaseExecutionsList []*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage,
	testCaseExecutionsModel *testCaseExecutionsModel.TestCaseExecutionsModelStruct) {

	dd

	// Store the TestCaseExecutionsThatCanBeViewedByUser-list in the TestCaseModel
	if testCaseExecutionsModel.TestCaseExecutionsThatCanBeViewedByUserMap == nil {
		testCaseExecutionsModel.TestCaseExecutionsThatCanBeViewedByUserMap = make(map[string]*fenixExecutionServerGuiGrpcApi.
			TestCaseExecutionsListMessage)
	}

	// Store the TestCaseExecutionsThatCanBeViewedByUser as a map structure in TestCaseExecution-struct
	for _, testCaseExecutions := range testCaseExecutionsList {

		testCaseExecutionsModel.TestCaseExecutionsThatCanBeViewedByUserMap[testCaseExecutions.GetTestCaseExecutionUuid()] =
			testCaseExecutions

	}

}
