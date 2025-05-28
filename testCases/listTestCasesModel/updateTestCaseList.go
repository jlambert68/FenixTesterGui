package listTestCasesModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testCase/testCaseModel"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"time"
)

// LoadTestCaseThatCanBeEditedByUser
// Load list with TestCasesMapPtr that the user can edit
func LoadTestCaseThatCanBeEditedByUser(
	testCaseModeReference *testCaseModel.TestCasesModelsStruct,
	testCaseUpdatedMinTimeStamp time.Time,
	testCaseExecutionUpdatedMinTimeStamp time.Time) {

	var listTestCasesThatCanBeEditedResponseMessage *fenixGuiTestCaseBuilderServerGrpcApi.ListTestCasesThatCanBeEditedResponseMessage
	listTestCasesThatCanBeEditedResponseMessage = testCaseModeReference.GrpcOutReference.
		ListTestCasesThatCanBeEditedResponseMessage(testCaseUpdatedMinTimeStamp, testCaseExecutionUpdatedMinTimeStamp)

	if listTestCasesThatCanBeEditedResponseMessage.GetAckNackResponse().AckNack == false {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "ebfe8adb-c224-4071-869b-f79cefde0dd3",
			"error": listTestCasesThatCanBeEditedResponseMessage.GetAckNackResponse().Comments,
		}).Warning("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'loadTestCaseThatCanBeEditedByUser'")

		return
	}

	// Store the slice with TestCasesMapPtr that a user can edit as a Map
	storeTestCaseThatCanBeEditedByUser(
		listTestCasesThatCanBeEditedResponseMessage.GetTestCasesThatCanBeEditedByUser(),
		testCaseModeReference)

	// Store the slice with TestCasesMapPtr
	//testCaseModeReference.TestCasesThatCanBeEditedByUserSlice = listTestCasesThatCanBeEditedResponseMessage.GetTestCasesThatCanBeEditedByUser()
	/*
		testCaseModeReference.TestCasesThatCanBeEditedByUserSlice = nil
		for _, tempTestCasesThatCanBeEditedByUser := range testCaseModeReference.TestCasesThatCanBeEditedByUserMap {
			testCaseModeReference.TestCasesThatCanBeEditedByUserSlice = append(
				testCaseModeReference.TestCasesThatCanBeEditedByUserSlice, tempTestCasesThatCanBeEditedByUser)
		}
	*/

}

// Store TestCasesMapPtr That Can Be Edited By User
func storeTestCaseThatCanBeEditedByUser(
	testCasesThatCanBeEditedByUserAsSlice []*fenixGuiTestCaseBuilderServerGrpcApi.TestCaseThatCanBeEditedByUserMessage,
	testCaseModeReference *testCaseModel.TestCasesModelsStruct) {

	// Store the TestCaseThatCanBeEditedByUser-list in the TestCaseModel
	testCaseModeReference.TestCasesThatCanBeEditedByUserMap = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
		TestCaseThatCanBeEditedByUserMessage)

	// Store the Available TemplateRepositoryApiUrls as a map structure in TestCase-struct
	for _, testCaseThatCanBeEditedByUser := range testCasesThatCanBeEditedByUserAsSlice {

		testCaseModeReference.TestCasesThatCanBeEditedByUserMap[testCaseThatCanBeEditedByUser.GetTestCaseUuid()] =
			testCaseThatCanBeEditedByUser

	}

}
