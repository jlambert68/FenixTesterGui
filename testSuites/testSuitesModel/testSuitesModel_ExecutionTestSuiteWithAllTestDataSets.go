package testSuitesModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/grpc_out_GuiExecutionServer"
	"FenixTesterGui/soundEngine"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
)

// ExecuteOneTestSuiteWithAllItsTestDataSets
// Execute one TestSuit's all TestCases with all its TestDataSets
func (testSuiteModel *TestSuiteModelStruct) ExecuteOneTestSuiteWithAllItsTestDataSets() {

	// Create message to be sent to GuiExecutionServer
	var initiateSingleTestSuiteExecutionRequestMessage *fenixExecutionServerGuiGrpcApi.InitiateTestSuiteExecutionWithAllTestDataSetsRequestMessage
	initiateSingleTestSuiteExecutionRequestMessage = &fenixExecutionServerGuiGrpcApi.InitiateTestSuiteExecutionWithAllTestDataSetsRequestMessage{
		UserAndApplicationRunTimeIdentification: &fenixExecutionServerGuiGrpcApi.UserAndApplicationRunTimeIdentificationMessage{
			ApplicationRunTimeUuid: sharedCode.ApplicationRunTimeUuid,
			UserIdOnComputer:       sharedCode.CurrentUserIdLogedInOnComputer,
			GCPAuthenticatedUser:   sharedCode.CurrentUserAuthenticatedTowardsGCP,
			ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
				grpc_out_GuiExecutionServer.GetHighestFenixGuiExecutionServerProtoFileVersion()),
		},
		TestSuiteUuid:              testSuiteModel.GetTestSuiteUuid(),
		ExecutionStatusReportLevel: fenixExecutionServerGuiGrpcApi.ExecutionStatusReportLevelEnum_REPORT_ALL_STATUS_CHANGES_ON_EXECUTIONS, //fenixExecutionServerGuiGrpcApi.ExecutionStatusReportLevelEnum_REPORT_ALL_STATUS_CHANGES_ON_EXECUTIONS,
	}

	// Initiate TestCaseExecution
	var initiateSingleTestSuiteExecutionResponseMessage *fenixExecutionServerGuiGrpcApi.InitiateSingleTestSuiteExecutionResponseMessage
	initiateSingleTestSuiteExecutionResponseMessage = grpc_out_GuiExecutionServer.GrpcOutGuiExecutionServerObject.
		SendInitiateTestSuiteExecutionWithAllTestDataSets(initiateSingleTestSuiteExecutionRequestMessage)

	if initiateSingleTestSuiteExecutionResponseMessage.AckNackResponse.AckNack == false {

		errorId := "28d1c4d7-a8a7-42b5-843e-53fdd79b4d18"
		err := errors.New(fmt.Sprintf("couldn't execute TestSuite due to error: '%s', {error: %s} [ErrorID: %s]",
			initiateSingleTestSuiteExecutionResponseMessage.AckNackResponse.Comments, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return

	}

	// Notify User that the TestCase is execution
	// Notify the user

	// Trigger System Notification sound
	soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   "TestSuiteExecution",
		Content: "The TestSuite is sent for Execution of all TestDataSets. See execution Tab for status.",
	})

}
