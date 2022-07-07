package testCaseModel

import (
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
)

type TestCaseModelStruct struct {
	LastLoadedTestCaseModelGRPCMessage fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage
	TestCaseModelMap                   map[string]fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementMessage
}
