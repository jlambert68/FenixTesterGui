package listTestSuiteExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/grpc_out_GuiExecutionServer"
	"FenixTesterGui/testSuiteExecutions/testSuiteExecutionsModel"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"time"
)

// LoadTestSuiteExecutionsThatCanBeViewedByUser
// Load list with TestSuiteExecutions that the user can view
func LoadTestSuiteExecutionsThatCanBeViewedByUser(
	latestUniqueTestSuiteExecutionDatabaseRowId int32,
	onlyRetrieveLimitedSizedBatch bool,
	batchSize int32,
	retrieveAllExecutionsForSpecificTestSuiteUuid bool,
	specificTestSuiteUuid string,
	testSuiteExecutionFromTimeStamp time.Time,
	testSuiteExecutionToTimeStamp time.Time,
	loadAllDataFromDatabase bool,
	updateGuiChannel *chan bool) {

	// Secure that the user has picked a TestSuiteExecution in the list before loading all executions for that TestSuite
	if retrieveAllExecutionsForSpecificTestSuiteUuid == true && len(specificTestSuiteUuid) == 0 {

		sharedCode.Logger.WithFields(logrus.Fields{
			"ID": "ea499116-ffdf-46d2-ae40-c7074bdac506",
			"retrieveAllExecutionsForSpecificTestSuiteUuid": retrieveAllExecutionsForSpecificTestSuiteUuid,
			"specificTestSuiteUuid":                         specificTestSuiteUuid,
		}).Fatal("'specificTestSuiteUuid' must have a value")

		return
	}

	var listTestSuiteExecutionsResponse *fenixExecutionServerGuiGrpcApi.ListTestSuiteExecutionsResponse

	if loadAllDataFromDatabase == false {
		// Only load one batch of data from the database
		listTestSuiteExecutionsResponse = grpc_out_GuiExecutionServer.GrpcOutGuiExecutionServerObject.
			SendListTestSuiteExecutionsThatCanBeViewed(
				latestUniqueTestSuiteExecutionDatabaseRowId,
				onlyRetrieveLimitedSizedBatch,
				batchSize,
				retrieveAllExecutionsForSpecificTestSuiteUuid,
				specificTestSuiteUuid,
				testSuiteExecutionFromTimeStamp,
				testSuiteExecutionToTimeStamp)

		if listTestSuiteExecutionsResponse.GetAckNackResponse().AckNack == false {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":    "0525519d-7321-4c60-b042-96c7a60ad1b7",
				"error": listTestSuiteExecutionsResponse.GetAckNackResponse().Comments,
			}).Warning("Problem to do gRPC-call to FenixGuiExecutionServer in 'LoadTestSuiteExecutionsThatCanBeViewedByUser'")

			return
		}

		// Store the TestSuiteExecutions in the Map
		if retrieveAllExecutionsForSpecificTestSuiteUuid == true {
			storeAllTestSuiteExecutionsForOneTestSuiteThatCanBeViewedByUser(
				listTestSuiteExecutionsResponse.GetTestSuiteExecutionsList(),
				&testSuiteExecutionsModel.TestSuiteExecutionsModel,
				listTestSuiteExecutionsResponse.
					GetLatestUniqueTestSuiteExecutionDatabaseRowId(),
				listTestSuiteExecutionsResponse.GetMoreRowsExists())
		} else {
			storeOneTestSuiteExecutionPerTestSuiteThatCanBeViewedByUser(
				listTestSuiteExecutionsResponse.GetTestSuiteExecutionsList(),
				&testSuiteExecutionsModel.TestSuiteExecutionsModel,
				listTestSuiteExecutionsResponse.
					GetLatestUniqueTestSuiteExecutionDatabaseRowId(),
				listTestSuiteExecutionsResponse.GetMoreRowsExists())
		}

	} else {
		// Load all data

		// Load first batch
		listTestSuiteExecutionsResponse = grpc_out_GuiExecutionServer.GrpcOutGuiExecutionServerObject.
			SendListTestSuiteExecutionsThatCanBeViewed(
				latestUniqueTestSuiteExecutionDatabaseRowId,
				onlyRetrieveLimitedSizedBatch,
				batchSize,
				retrieveAllExecutionsForSpecificTestSuiteUuid,
				specificTestSuiteUuid,
				testSuiteExecutionFromTimeStamp,
				testSuiteExecutionToTimeStamp)

		if listTestSuiteExecutionsResponse.GetAckNackResponse().AckNack == false {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":    "5b0f44d0-b172-48a1-a64d-5e38f5bfd8e7",
				"error": listTestSuiteExecutionsResponse.GetAckNackResponse().Comments,
			}).Warning("Problem to do gRPC-call to FenixGuiExecutionServer in 'LoadTestSuiteExecutionsThatCanBeViewedByUser'")

			return
		}

		// Store the TestSuiteExecutions in the Map
		if retrieveAllExecutionsForSpecificTestSuiteUuid == true {
			storeAllTestSuiteExecutionsForOneTestSuiteThatCanBeViewedByUser(
				listTestSuiteExecutionsResponse.GetTestSuiteExecutionsList(),
				&testSuiteExecutionsModel.TestSuiteExecutionsModel,
				listTestSuiteExecutionsResponse.
					GetLatestUniqueTestSuiteExecutionDatabaseRowId(),
				listTestSuiteExecutionsResponse.GetMoreRowsExists())
		} else {
			storeOneTestSuiteExecutionPerTestSuiteThatCanBeViewedByUser(
				listTestSuiteExecutionsResponse.GetTestSuiteExecutionsList(),
				&testSuiteExecutionsModel.TestSuiteExecutionsModel,
				listTestSuiteExecutionsResponse.
					GetLatestUniqueTestSuiteExecutionDatabaseRowId(),
				listTestSuiteExecutionsResponse.GetMoreRowsExists())
		}

		// Load rest of the data as go-routine
		go func(updateGuiChannel *chan bool) {
			listTestSuiteExecutionsResponse = grpc_out_GuiExecutionServer.GrpcOutGuiExecutionServerObject.
				SendListTestSuiteExecutionsThatCanBeViewed(
					testSuiteExecutionsModel.TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
						LatestUniqueTestSuiteExecutionDatabaseRowId,
					false,
					0,
					retrieveAllExecutionsForSpecificTestSuiteUuid,
					specificTestSuiteUuid,
					testSuiteExecutionFromTimeStamp,
					testSuiteExecutionToTimeStamp)

			if listTestSuiteExecutionsResponse.GetAckNackResponse().AckNack == false {
				sharedCode.Logger.WithFields(logrus.Fields{
					"ID":    "37485a00-077d-4a87-9c21-035e4dc45675",
					"error": listTestSuiteExecutionsResponse.GetAckNackResponse().Comments,
				}).Warning("Problem to do gRPC-call to FenixGuiExecutionServer in 'LoadTestSuiteExecutionsThatCanBeViewedByUser'")

				return
			}

			// Store the TestSuiteExecutions in the Map
			if retrieveAllExecutionsForSpecificTestSuiteUuid == true {
				storeAllTestSuiteExecutionsForOneTestSuiteThatCanBeViewedByUser(
					listTestSuiteExecutionsResponse.GetTestSuiteExecutionsList(),
					&testSuiteExecutionsModel.TestSuiteExecutionsModel,
					listTestSuiteExecutionsResponse.
						GetLatestUniqueTestSuiteExecutionDatabaseRowId(),
					listTestSuiteExecutionsResponse.GetMoreRowsExists())
			} else {
				storeOneTestSuiteExecutionPerTestSuiteThatCanBeViewedByUser(
					listTestSuiteExecutionsResponse.GetTestSuiteExecutionsList(),
					&testSuiteExecutionsModel.TestSuiteExecutionsModel,
					listTestSuiteExecutionsResponse.
						GetLatestUniqueTestSuiteExecutionDatabaseRowId(),
					listTestSuiteExecutionsResponse.GetMoreRowsExists())
			}

			*updateGuiChannel <- true

		}(updateGuiChannel)

	}

}

// Store TestSuiteExecutions That Can Be Viewed By User
func storeOneTestSuiteExecutionPerTestSuiteThatCanBeViewedByUser(
	testSuiteExecutionsList []*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage,
	testSuiteExecutionsModelRef *testSuiteExecutionsModel.TestSuiteExecutionsModelStruct,
	latestUniqueTestSuiteExecutionDatabaseRowId int32,
	moreRowsExists bool) {

	// Store information about latest row retrieved and if there are more rows
	testSuiteExecutionsModel.TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
		LatestUniqueTestSuiteExecutionDatabaseRowId = latestUniqueTestSuiteExecutionDatabaseRowId
	testSuiteExecutionsModel.TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
		MoreRowsExists = moreRowsExists

	// Store the TestSuiteExecutionsThatCanBeViewedByUser as a map structure in TestSuiteExecution-struct
	for _, testSuiteExecutions := range testSuiteExecutionsList {

		testSuiteExecutionsModelRef.AddToTestSuiteExecutionsMap(
			testSuiteExecutionsModel.TestSuiteExecutionUuidType(testSuiteExecutions.GetTestSuiteExecutionUuid()),
			testSuiteExecutions)
	}

}

// Store All TestSuiteExecutions for one TestSuite, That Can Be Viewed By User
func storeAllTestSuiteExecutionsForOneTestSuiteThatCanBeViewedByUser(
	testSuiteExecutionsList []*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage,
	testSuiteExecutionsModelRef *testSuiteExecutionsModel.TestSuiteExecutionsModelStruct,
	latestUniqueTestSuiteExecutionDatabaseRowId int32,
	moreRowsExists bool) {

	var testSuiteUuid string
	var testSuiteExecutionUuid string

	// Store the TestSuiteExecutionsThatCanBeViewedByUser as a map structure in TestSuiteExecution-struct
	for _, testSuiteExecution := range testSuiteExecutionsList {

		// Extract keys
		testSuiteUuid = testSuiteExecution.GetTestSuiteUuid()
		testSuiteExecutionUuid = testSuiteExecution.GetTestSuiteExecutionUuid()

		// Store the TestSuiteExecution
		testSuiteExecutionsModel.AddTestSuiteExecutionsForOneTestSuiteUuid(
			testSuiteExecutionsModelRef,
			testSuiteExecutionsModel.TestSuiteUuidType(testSuiteUuid),
			testSuiteExecutionsModel.TestSuiteExecutionUuidType(testSuiteExecutionUuid),
			testSuiteExecution,
			latestUniqueTestSuiteExecutionDatabaseRowId,
			moreRowsExists)

	}

}
