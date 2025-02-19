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
	retrieveAllExecutionsForSpecificTestCaseUuid bool,
	specificTestCaseUuid string,
	testCaseExecutionFromTimeStamp time.Time,
	testCaseExecutionToTimeStamp time.Time,
	loadAllDataFromDatabase bool) {

	// Secure that the user has picked a TestCaseExecution in the list before loading all executions for that TestCase
	if retrieveAllExecutionsForSpecificTestCaseUuid == true && len(specificTestCaseUuid) == 0 {

		sharedCode.Logger.WithFields(logrus.Fields{
			"ID": "ab8b70f7-9a0d-4e37-b2c8-aca1199a289d",
			"retrieveAllExecutionsForSpecificTestCaseUuid": retrieveAllExecutionsForSpecificTestCaseUuid,
			"specificTestCaseUuid":                         specificTestCaseUuid,
		}).Fatal("'specificTestCaseUuid' must have a value")

		return
	}

	var listTestCaseExecutionsResponse *fenixExecutionServerGuiGrpcApi.ListTestCaseExecutionsResponse

	if loadAllDataFromDatabase == false {
		// Only load one batch of data from the database
		listTestCaseExecutionsResponse = grpc_out_GuiExecutionServer.GrpcOutGuiExecutionServerObject.
			SendListTestCaseExecutionsThatCanBeViewed(
				latestUniqueTestCaseExecutionDatabaseRowId,
				onlyRetrieveLimitedSizedBatch,
				batchSize,
				retrieveAllExecutionsForSpecificTestCaseUuid,
				specificTestCaseUuid,
				testCaseExecutionFromTimeStamp,
				testCaseExecutionToTimeStamp)

		if listTestCaseExecutionsResponse.GetAckNackResponse().AckNack == false {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":    "320c6409-a68b-4cf0-adc1-aa65d8c51343",
				"error": listTestCaseExecutionsResponse.GetAckNackResponse().Comments,
			}).Warning("Problem to do gRPC-call to FenixGuiExecutionServer in 'LoadTestCaseExecutionsThatCanBeViewedByUser'")

			return
		}

		// Store information about latest row retrieved and if there are more rows
		testCaseExecutionsModel.TestCaseExecutionsModel.
			LatestUniqueTestCaseExecutionDatabaseRowId = listTestCaseExecutionsResponse.
			GetLatestUniqueTestCaseExecutionDatabaseRowId()
		testCaseExecutionsModel.TestCaseExecutionsModel.MoreRowsExists = listTestCaseExecutionsResponse.GetMoreRowsExists()

		// Store the TestCaseExecutions in the Map
		if retrieveAllExecutionsForSpecificTestCaseUuid == true {
			storeAllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUser(
				listTestCaseExecutionsResponse.GetTestCaseExecutionsList(),
				&testCaseExecutionsModel.TestCaseExecutionsModel)
		} else {
			storeTestCaseExecutionsThatCanBeViewedByUser(
				listTestCaseExecutionsResponse.GetTestCaseExecutionsList(),
				&testCaseExecutionsModel.TestCaseExecutionsModel)
		}

	} else {
		// Load all data

		// Load first batch
		listTestCaseExecutionsResponse = grpc_out_GuiExecutionServer.GrpcOutGuiExecutionServerObject.
			SendListTestCaseExecutionsThatCanBeViewed(
				latestUniqueTestCaseExecutionDatabaseRowId,
				onlyRetrieveLimitedSizedBatch,
				batchSize,
				retrieveAllExecutionsForSpecificTestCaseUuid,
				specificTestCaseUuid,
				testCaseExecutionFromTimeStamp,
				testCaseExecutionToTimeStamp)

		if listTestCaseExecutionsResponse.GetAckNackResponse().AckNack == false {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":    "320c6409-a68b-4cf0-adc1-aa65d8c51343",
				"error": listTestCaseExecutionsResponse.GetAckNackResponse().Comments,
			}).Warning("Problem to do gRPC-call to FenixGuiExecutionServer in 'LoadTestCaseExecutionsThatCanBeViewedByUser'")

			return
		}

		// Store information about latest row retrieved and if there are more rows
		testCaseExecutionsModel.TestCaseExecutionsModel.
			LatestUniqueTestCaseExecutionDatabaseRowId = listTestCaseExecutionsResponse.
			GetLatestUniqueTestCaseExecutionDatabaseRowId()
		testCaseExecutionsModel.TestCaseExecutionsModel.MoreRowsExists = listTestCaseExecutionsResponse.GetMoreRowsExists()

		// Store the TestCaseExecutions in the Map
		if retrieveAllExecutionsForSpecificTestCaseUuid == true {
			storeAllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUser(
				listTestCaseExecutionsResponse.GetTestCaseExecutionsList(),
				&testCaseExecutionsModel.TestCaseExecutionsModel)
		} else {
			storeTestCaseExecutionsThatCanBeViewedByUser(
				listTestCaseExecutionsResponse.GetTestCaseExecutionsList(),
				&testCaseExecutionsModel.TestCaseExecutionsModel)
		}

		// Load rest of the data as go-routine
		go func() {
			listTestCaseExecutionsResponse = grpc_out_GuiExecutionServer.GrpcOutGuiExecutionServerObject.
				SendListTestCaseExecutionsThatCanBeViewed(
					testCaseExecutionsModel.TestCaseExecutionsModel.LatestUniqueTestCaseExecutionDatabaseRowId,
					false,
					0,
					retrieveAllExecutionsForSpecificTestCaseUuid,
					specificTestCaseUuid,
					testCaseExecutionFromTimeStamp,
					testCaseExecutionToTimeStamp)

			if listTestCaseExecutionsResponse.GetAckNackResponse().AckNack == false {
				sharedCode.Logger.WithFields(logrus.Fields{
					"ID":    "320c6409-a68b-4cf0-adc1-aa65d8c51343",
					"error": listTestCaseExecutionsResponse.GetAckNackResponse().Comments,
				}).Warning("Problem to do gRPC-call to FenixGuiExecutionServer in 'LoadTestCaseExecutionsThatCanBeViewedByUser'")

				return
			}

			// Store information about latest row retrieved and if there are more rows
			testCaseExecutionsModel.TestCaseExecutionsModel.
				LatestUniqueTestCaseExecutionDatabaseRowId = listTestCaseExecutionsResponse.
				GetLatestUniqueTestCaseExecutionDatabaseRowId()
			testCaseExecutionsModel.TestCaseExecutionsModel.MoreRowsExists = listTestCaseExecutionsResponse.GetMoreRowsExists()

			// Store the TestCaseExecutions in the Map
			if retrieveAllExecutionsForSpecificTestCaseUuid == true {
				storeAllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUser(
					listTestCaseExecutionsResponse.GetTestCaseExecutionsList(),
					&testCaseExecutionsModel.TestCaseExecutionsModel)
			} else {
				storeTestCaseExecutionsThatCanBeViewedByUser(
					listTestCaseExecutionsResponse.GetTestCaseExecutionsList(),
					&testCaseExecutionsModel.TestCaseExecutionsModel)
			}

		}()

	}

}

// Store TestCaseExecutions That Can Be Viewed By User
func storeTestCaseExecutionsThatCanBeViewedByUser(
	testCaseExecutionsList []*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage,
	testCaseExecutionsModelRef *testCaseExecutionsModel.TestCaseExecutionsModelStruct) {

	// Store the TestCaseExecutionsThatCanBeViewedByUser as a map structure in TestCaseExecution-struct
	for _, testCaseExecutions := range testCaseExecutionsList {

		testCaseExecutionsModelRef.AddToTestCaseExecutionsMap(
			testCaseExecutionsModel.TestCaseExecutionUuidType(testCaseExecutions.GetTestCaseExecutionUuid()),
			testCaseExecutions)
	}

}

// Store All TestCaseExecutions for one TestCase, That Can Be Viewed By User
func storeAllTestCaseExecutionsForOneTestCaseThatCanBeViewedByUser(
	testCaseExecutionsList []*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage,
	testCaseExecutionsModelRef *testCaseExecutionsModel.TestCaseExecutionsModelStruct) {

	var testCaseUuid string

	// Store the TestCaseExecutionsThatCanBeViewedByUser as a map structure in TestCaseExecution-struct
	for _, testCaseExecutions := range testCaseExecutionsList {

		// Extract the TestCase that all executions should be stored under
		testCaseUuid =

			testCaseExecutionsModelRef.AddToAllTestCaseExecutionsForOneTestCaseMap(
				testCaseExecutionsModel.TestCaseExecutionUuidType(testCaseExecutions.GetTestCaseExecutionUuid()),
				testCaseExecutions)
	}

}
