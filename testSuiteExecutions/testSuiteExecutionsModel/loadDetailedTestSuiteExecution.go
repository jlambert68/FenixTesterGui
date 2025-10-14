package testSuiteExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/grpc_out_GuiExecutionServer"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"strconv"
)

// LoadDetailedTestSuiteExecutionFromDatabase
// Load all Detailed TestSuiteExecution-data for specific execution
func LoadDetailedTestSuiteExecutionFromDatabase(testSuiteExecutionUuid string, testSuiteExecutionVersion uint32) {

	// Create DetailedTestSuiteExecutionMapKey
	var detailedTestSuiteExecutionMapKey DetailedTestSuiteExecutionMapKeyType
	detailedTestSuiteExecutionMapKey = DetailedTestSuiteExecutionMapKeyType(testSuiteExecutionUuid +
		strconv.Itoa(int(testSuiteExecutionVersion)))

	// Set the flag there is an ongoing refresh of the DetailedTestSuiteExecution-data
	TestSuiteExecutionsModel.SetFlagRefreshOngoingOfDetailedTestSuiteExecution(detailedTestSuiteExecutionMapKey)

	// Create request message for 'SendGetSingleTestSuiteExecution'
	var getSingleTestSuiteExecutionRequest *fenixExecutionServerGuiGrpcApi.GetSingleTestSuiteExecutionRequest
	getSingleTestSuiteExecutionRequest = &fenixExecutionServerGuiGrpcApi.GetSingleTestSuiteExecutionRequest{
		UserAndApplicationRunTimeIdentification: &fenixExecutionServerGuiGrpcApi.UserAndApplicationRunTimeIdentificationMessage{
			ApplicationRunTimeUuid: sharedCode.ApplicationRunTimeUuid,
			UserIdOnComputer:       sharedCode.CurrentUserIdLogedInOnComputer,
			GCPAuthenticatedUser:   sharedCode.CurrentUserAuthenticatedTowardsGCP,
			ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
				grpc_out_GuiExecutionServer.GetHighestFenixGuiExecutionServerProtoFileVersion()),
		},
		TestSuiteExecutionKey: &fenixExecutionServerGuiGrpcApi.TestSuiteExecutionKeyMessage{
			TestSuiteExecutionUuid:    testSuiteExecutionUuid,
			TestSuiteExecutionVersion: testSuiteExecutionVersion,
		},
	}
	// Call GuiExecution-server to load Detailed TestSuiteExecution
	var getSingleTestSuiteExecutionResponse *fenixExecutionServerGuiGrpcApi.GetSingleTestSuiteExecutionResponse
	getSingleTestSuiteExecutionResponse = grpc_out_GuiExecutionServer.GrpcOutGuiExecutionServerObject.
		SendGetSingleTestSuiteExecution(getSingleTestSuiteExecutionRequest)

	// Clear the flag there is an ongoing refresh of the DetailedTestSuiteExecution-data
	defer TestSuiteExecutionsModel.ClearFlagRefreshOngoingOfDetailedTestSuiteExecution(detailedTestSuiteExecutionMapKey)

	// Store the Detailed TestSuiteExecution
	if getSingleTestSuiteExecutionResponse != nil &&
		getSingleTestSuiteExecutionResponse.GetAckNackResponse().GetAckNack() == true {

		TestSuiteExecutionsModel.AddToDetailedTestSuiteExecutionsMap(
			detailedTestSuiteExecutionMapKey,
			getSingleTestSuiteExecutionResponse.GetTestSuiteExecutionResponse())

	} else {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "9f6c67a6-3cd3-4970-9914-f8665f637a60",
		}).Warning("Nothing to store in 'DetailedTestSuiteExecutionsMap', should not happen")

		// Clear the flag there is an ongoing refresh of the DetailedTestSuiteExecution-data
		defer TestSuiteExecutionsModel.ClearFlagRefreshOngoingOfDetailedTestSuiteExecution(detailedTestSuiteExecutionMapKey)

		return

	}

	// Clear the flag there is an ongoing refresh of the DetailedTestSuiteExecution-data
	TestSuiteExecutionsModel.ClearFlagRefreshOngoingOfDetailedTestSuiteExecution(detailedTestSuiteExecutionMapKey)

	// Extracts all LogPost-messages from a TestSuiteExecution and store them in a map per TInTICExecutionKey
	_ = TestSuiteExecutionsModel.ExtractAndStoreLogPostsAndValuesFromDetailedTestSuiteExecution(detailedTestSuiteExecutionMapKey)

	// Extract relation between TestInstructionUuid and TestSuiteExecutionUuid
	//_ = TestSuiteExecutionsModel.ExtractAndStoreRelationBetweenTestInstructionUuidAndTestSuiteExecutionUuid(detailedTestSuiteExecutionMapKey)

}
