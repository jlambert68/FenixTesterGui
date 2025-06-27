package testSuitesModel

import (
	"errors"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
)

// DeleteTestSuiteAtThisDate - Mark the TestSuite as deleted by this date, in the Database
func (testSuiteModel *TestSuiteModelStruct) DeleteTestSuiteAtThisDate(
	testSuiteUuid string) (err error) {

	// Populate data to be sent to TestCaseBuilderServer
	var gRPCDeleteTestCaseAtThisDateRequest *fenixGuiTestCaseBuilderServerGrpcApi.DeleteTestSuiteAtThisDateRequest
	gRPCDeleteTestCaseAtThisDateRequest = &fenixGuiTestCaseBuilderServerGrpcApi.DeleteTestSuiteAtThisDateRequest{
		UserIdentification: nil,
		DeleteThisTestSuiteAtThisDate: &fenixGuiTestCaseBuilderServerGrpcApi.DeleteTestSuiteAtThisDateRequestMessage{
			DomainUuid:       testSuiteModel.savedTestSuiteUIModelBinding.TestSuiteOwnerDomainUuid,
			DomainName:       testSuiteModel.savedTestSuiteUIModelBinding.TestSuiteOwnerDomainName,
			TestSuiteUuid:    testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.testSuiteUuid,
			TestSuiteVersion: testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.testSuiteVersion,
			DeletedDate:      testSuiteModel.TestSuiteUIModelBinding.TestSuiteDeletionDate,
		},
	}

	// Send using gRPC
	returnMessage := testSuiteModel.testCasesModel.GrpcOutReference.SendDeleteTestSuiteAtThisDate(
		gRPCDeleteTestCaseAtThisDateRequest)

	if returnMessage == nil || returnMessage.AckNack == false {

		errorId := "9e561420-be42-4deb-976b-3718d262a495"
		err = errors.New(fmt.Sprintf(returnMessage.Comments+"[ErrorID: %s]", testSuiteUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return err
	}

	return err
}
