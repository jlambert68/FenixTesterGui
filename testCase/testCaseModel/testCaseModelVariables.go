package testCaseModel

import (
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
)

type TestCaseModel struct {
	LastLoadedTestCaseModelGRPCMessage fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage
	testCaseModelMap                   map[string]fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementMessage
}
