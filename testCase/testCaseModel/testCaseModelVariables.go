package testCaseModel

import (
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"time"
)

type TestCaseModelsStruct struct {
	TestCases   map[string]TestCaseModelStruct
	CurrentUser string
}

type TestCaseModelStruct struct {
	LastLoadedTestCaseModelGRPCMessage   fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage
	FirstElementUuid                     string
	TestCaseModelMap                     map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage
	TextualTestCaseRepresentationSimple  []string
	TextualTestCaseRepresentationComplex []string
	CommandStack                         []fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage_TestCaseModelCommandMessage
	LastSavedCommandStack                lastSavedCommandStack
	copyBuffer                           fenixGuiTestCaseBuilderServerGrpcApi.ImmatureElementModelMessage
	cutBuffer                            matureElementStruct
}

type lastSavedCommandStack struct {
	savedTimeStamp time.Time
	commandStack   []fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage_TestCaseModelCommandMessage
}

type matureElementStruct struct {
	firstElementUuid string
	matureElementMap map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage
}

const NotApplicable = "N/A"
