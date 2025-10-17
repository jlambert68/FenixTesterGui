package testCaseExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/grpc_out_GuiExecutionServer"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"strconv"
)

// LoadDetailedTestCaseExecutionFromDatabase
// Load all Detailed TestCaseExecution-data for specific execution
func LoadDetailedTestCaseExecutionFromDatabase(testCaseExecutionUuid string, testCaseExecutionVersion uint32) {

	// Create DetailedTestCaseExecutionMapKey
	var detailedTestCaseExecutionMapKey DetailedTestCaseExecutionMapKeyType
	detailedTestCaseExecutionMapKey = DetailedTestCaseExecutionMapKeyType(testCaseExecutionUuid +
		strconv.Itoa(int(testCaseExecutionVersion)))

	// Set the flag there is an ongoing refresh of the DetailedTestCaseExecution-data
	TestCaseExecutionsModel.SetFlagRefreshOngoingOfDetailedTestCaseExecution(detailedTestCaseExecutionMapKey)

	// Create request message for 'SendGetSingleTestCaseExecution'
	var getSingleTestCaseExecutionRequest *fenixExecutionServerGuiGrpcApi.GetSingleTestCaseExecutionRequest
	getSingleTestCaseExecutionRequest = &fenixExecutionServerGuiGrpcApi.GetSingleTestCaseExecutionRequest{
		UserAndApplicationRunTimeIdentification: &fenixExecutionServerGuiGrpcApi.UserAndApplicationRunTimeIdentificationMessage{
			ApplicationRunTimeUuid: sharedCode.ApplicationRunTimeUuid,
			UserIdOnComputer:       sharedCode.CurrentUserIdLogedInOnComputer,
			GCPAuthenticatedUser:   sharedCode.CurrentUserAuthenticatedTowardsGCP,
			ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
				grpc_out_GuiExecutionServer.GetHighestFenixGuiExecutionServerProtoFileVersion()),
		},
		TestCaseExecutionKey: &fenixExecutionServerGuiGrpcApi.TestCaseExecutionKeyMessage{
			TestCaseExecutionUuid:    testCaseExecutionUuid,
			TestCaseExecutionVersion: testCaseExecutionVersion,
		},
	}
	// Call GuiExecution-server to load Detailed TestCaseExecution
	var getSingleTestCaseExecutionResponse *fenixExecutionServerGuiGrpcApi.GetSingleTestCaseExecutionResponse
	getSingleTestCaseExecutionResponse = grpc_out_GuiExecutionServer.GrpcOutGuiExecutionServerObject.
		SendGetSingleTestCaseExecution(getSingleTestCaseExecutionRequest)

	// Clear the flag there is an ongoing refresh of the DetailedTestCaseExecution-data
	defer TestCaseExecutionsModel.ClearFlagRefreshOngoingOfDetailedTestCaseExecution(detailedTestCaseExecutionMapKey)

	// Store the Detailed TestCaseExecution
	if getSingleTestCaseExecutionResponse != nil &&
		getSingleTestCaseExecutionResponse.GetAckNackResponse().GetAckNack() == true {

		TestCaseExecutionsModel.AddToDetailedTestCaseExecutionsMap(
			detailedTestCaseExecutionMapKey,
			getSingleTestCaseExecutionResponse.GetTestCaseExecutionResponse())

	} else {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "2e6879c4-11e9-44e7-864f-8f6787ff8b90",
		}).Warning("Nothing to store in 'DetailedTestCaseExecutionsMap', should not happen")

		// Clear the flag there is an ongoing refresh of the DetailedTestCaseExecution-data
		//defer TestCaseExecutionsModel.ClearFlagRefreshOngoingOfDetailedTestCaseExecution(detailedTestCaseExecutionMapKey)

		return

	}

	// Clear the flag there is an ongoing refresh of the DetailedTestCaseExecution-data
	TestCaseExecutionsModel.ClearFlagRefreshOngoingOfDetailedTestCaseExecution(detailedTestCaseExecutionMapKey)

	// Extracts all LogPost-messages from a TestCaseExecution and store them in a map per TInTICExecutionKey
	_ = TestCaseExecutionsModel.ExtractAndStoreLogPostsAndValuesFromDetailedTestCaseExecution(detailedTestCaseExecutionMapKey)

	// Extract relation between TestInstructionUuid and TestCaseExecutionUuid
	//_ = TestCaseExecutionsModel.ExtractAndStoreRelationBetweenTestInstructionUuidAndTestCaseExecutionUuid(detailedTestCaseExecutionMapKey)

}
