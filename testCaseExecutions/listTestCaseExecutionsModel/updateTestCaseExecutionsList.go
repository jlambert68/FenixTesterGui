package listTestCaseExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testCase/testCaseModel"
	"FenixTesterGui/testCaseExecutions/testCaseExecutionsModel"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"time"
)

// LoadTestCaseExecutionsThatCanBeViewedByUser
// Load list with TestCaseExecutions that the user can view
func LoadTestCaseExecutionsThatCanBeViewedByUser(
	testCaseModeReference *testCaseModel.TestCasesModelsStruct,
	testCaseUpdatedMinTimeStamp time.Time,
	testCaseExecutionUpdatedMinTimeStamp time.Time) {

	var listTestCasesThatCanBeEditedResponseMessage *fenixExecutionServerGuiGrpcApi.ListTestCaseExecutionsResponse
	listTestCasesThatCanBeEditedResponseMessage = testCaseModeReference.GrpcOutReference.
		ListTestCasesThatCanBeEditedResponseMessage(testCaseUpdatedMinTimeStamp, testCaseExecutionUpdatedMinTimeStamp)

	if listTestCasesThatCanBeEditedResponseMessage.GetAckNackResponse().AckNack == false {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "e703c704-8b96-4235-b403-a36e73e08a18",
			"error": listTestCasesThatCanBeEditedResponseMessage.GetAckNackResponse().Comments,
		}).Warning("Problem to do gRPC-call to FenixGuiExecutionServer in 'LoadTestCaseExecutionsThatCanBeViewedByUser'")

		return
	}

	// Store the slice with TestCases that a user can edit as a Map
	storeTestCaseExecutionsThatCanBeViewedByUser(
		listTestCasesThatCanBeEditedResponseMessage.GetTestCasesThatCanBeEditedByUser(),
		testCaseModeReference)

	// Store the slice with TestCases
	//testCaseModeReference.TestCasesThatCanBeEditedByUserSlice = listTestCasesThatCanBeEditedResponseMessage.GetTestCasesThatCanBeEditedByUser()
	testCaseModeReference.TestCasesThatCanBeEditedByUserSlice = nil
	for _, tempTestCasesThatCanBeEditedByUser := range testCaseModeReference.TestCasesThatCanBeEditedByUserMap {
		testCaseModeReference.TestCasesThatCanBeEditedByUserSlice = append(
			testCaseModeReference.TestCasesThatCanBeEditedByUserSlice, tempTestCasesThatCanBeEditedByUser)
	}

}

// Store TestCases That Can Be Edited By User
func storeTestCaseExecutionsThatCanBeViewedByUser(
	testCaseExecutionsList []*fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage,
	testCaseExecutionsModel *testCaseExecutionsModel.TestCaseExecutionsModelStruct) {

	// Store the TestCaseExecutionsThatCanBeViewedByUser-list in the TestCaseModel
	if testCaseExecutionsModel.TestCaseExecutionsThatCanBeViewedByUserMap == nil {
		testCaseExecutionsModel.TestCaseExecutionsThatCanBeViewedByUserMap = make(map[string]*fenixExecutionServerGuiGrpcApi.
			ListTestCaseExecutionsResponse)
	}

	// Store the TestCaseExecutionsThatCanBeViewedByUser as a map structure in TestCaseExecution-struct
	for _, testCaseExecutions := range testCaseExecutionsList {

		testCaseExecutionsModel.TestCaseExecutionsThatCanBeViewedByUserMap[testCaseExecutions.GetTestCaseUuid()] =
			testCaseExecutions

	}

}
