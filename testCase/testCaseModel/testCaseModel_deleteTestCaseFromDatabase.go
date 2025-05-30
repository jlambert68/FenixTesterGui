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

	var existsInMap bool

	// Get TestCasesMap
	var testCasesMap map[string]*TestCaseModelStruct
	testCasesMap = *testCaseModel.TestCasesMapPtr

	// Get current TestCase
	var currentTestCasePtr *TestCaseModelStruct
	currentTestCasePtr, existsInMap = testCasesMap[testCaseUuid]

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
			DomainUuid:      currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.GetDomainUuid(),
			DomainName:      currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.GetDomainName(),
			TestCaseUuid:    currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.GetTestCaseUuid(),
			TestCaseVersion: currentTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageNoneEditableInformation.GetTestCaseVersion(),
			DeletedDate:     currentTestCasePtr.LocalTestCaseMessage.DeleteTimeStamp,
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
