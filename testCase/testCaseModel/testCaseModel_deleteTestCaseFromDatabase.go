package testCaseModel

import (
	"errors"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
)

// DeleteTestCaseAtThisDate - Mark the TestCase as deletedn by this date, in the Database
func (testCaseModel *TestCasesModelsStruct) DeleteTestCaseAtThisDate(
	testCaseUuid string) (err error) {

	// Save changed Attributes, if there are any, into the TestCase-model. Needs to call because last attributes change is not saved into model
	err = testCaseModel.SaveChangedTestCaseAttributeInTestCase(testCaseUuid)
	if err != nil {
		return err
	}

	// Get current TestCase
	currentTestCase, existsInMap := testCaseModel.TestCasesMapPtr[testCaseUuid]
	if existsInMap == false {

		errorId := "4c075798-ec6c-4486-8053-997ef0d0d8eb"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCasesMapPtr [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return err
	}

	// Populate data to be sent to TestCaseBuilderServer
	var gRPCDeleteTestCaseAtThisDateRequest *fenixGuiTestCaseBuilderServerGrpcApi.DeleteTestCaseAtThisDateRequest
	gRPCDeleteTestCaseAtThisDateRequest = &fenixGuiTestCaseBuilderServerGrpcApi.DeleteTestCaseAtThisDateRequest{
		UserIdentification: nil,
		DeleteThisTestCaseAtThisDate: &fenixGuiTestCaseBuilderServerGrpcApi.DeleteTestCaseAtThisDateRequestMessage{
			DomainUuid:      currentTestCase.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.GetDomainUuid(),
			DomainName:      currentTestCase.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.GetDomainName(),
			TestCaseUuid:    currentTestCase.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.GetTestCaseUuid(),
			TestCaseVersion: currentTestCase.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.GetTestCaseVersion(),
			DeletedDate:     currentTestCase.LocalTestCaseMessage.DeleteTimeStamp,
		},
	}

	// Send using gRPC
	returnMessage := testCaseModel.GrpcOutReference.SendDeleteTestCaseAtThisDate(gRPCDeleteTestCaseAtThisDateRequest)

	if returnMessage == nil || returnMessage.AckNack == false {

		errorId := "7f96e2f3-7470-4e08-9680-f66c61a11a5a"
		err = errors.New(fmt.Sprintf(returnMessage.Comments+"[ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return err
	}

	return err
}
