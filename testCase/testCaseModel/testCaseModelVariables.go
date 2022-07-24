package testCaseModel

import (
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"time"
)

type TestCaseModelsStruct struct {
	TestCases map[string]TestCaseModelStruct
}

type TestCaseModelStruct struct {
	LastLoadedTestCaseModelGRPCMessage   fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage
	FirstElementUuid                     string
	TestCaseModelMap                     map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage
	TextualTestCaseRepresentationSimple  []string
	TextualTestCaseRepresentationComplex []string
	commandStack                         []commandStackStruct
	LastSavedCommandStack                lastSavedCommandStack
	copyBuffer                           fenixGuiTestCaseBuilderServerGrpcApi.ImmatureElementModelMessage
	cutBuffer                            matureElementStruct
}

type lastSavedCommandStack struct {
	savedTimeStamp time.Time
	commandStack   []commandStackStruct
}

type commandStackStruct struct {
	command           fenixGuiTestCaseBuilderServerGrpcApi.TestCaseCommandTypeEnum
	commandName       string
	commandParameter1 string
	commandParameter2 string
	updatedDateTime   time.Time
}

type matureElementStruct struct {
	firstElementUuid string
	matureElementMap map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage
}
