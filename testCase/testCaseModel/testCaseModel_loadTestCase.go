package testCaseModel

import (
	"errors"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
)

// LoadFullTestCaseFromDatabase - Load the TestCase from the Database into model
func (testCaseModel *TestCasesModelsStruct) LoadFullTestCaseFromDatabase(testCaseUuid string, currentActiveUser string) (err error) {

	// Send LoadTesCase using gRPC
	var detailedTestCaseResponse *fenixGuiTestCaseBuilderServerGrpcApi.GetDetailedTestCaseResponse
	detailedTestCaseResponse = testCaseModel.GrpcOutReference.LoadDetailedTestCase(currentActiveUser, testCaseUuid)

	// Exit if something was wrong
	if detailedTestCaseResponse.AckNackResponse.AckNack == false {

		errorId := "ba195459-8902-4727-ab81-ae48cd616eea"
		err = errors.New(fmt.Sprintf(detailedTestCaseResponse.AckNackResponse.Comments+"[ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return err

	}

	konvertera till testCaseModel, troligen via version av NewTestCase eller liknande

	return err

}
