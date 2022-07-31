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
	LastLoadedTestCaseModelGRPCMessage        fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage
	FirstElementUuid                          string
	TestCaseModelMap                          map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage
	TextualTestCaseRepresentationSimpleStack  []string
	TextualTestCaseRepresentationComplexStack []string
	CommandStack                              []fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage_TestCaseModelCommandMessage
	LastSavedCommandStack                     lastSavedCommandStack
	CopyBuffer                                ImmatureElementStruct
	CutBuffer                                 MatureElementStruct
	CutCommandInitiated                       bool
}

type lastSavedCommandStack struct {
	savedTimeStamp time.Time
	commandStack   []fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage_TestCaseModelCommandMessage
}

type ImmatureElementStruct struct {
	FirstElementUuid   string
	ImmatureElementMap map[string]fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage
}

type MatureElementStruct struct {
	FirstElementUuid string
	MatureElementMap map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage
}

/*
type MatureElementStruct struct {
	FirstElementUuid string
	MatureElementMap map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage
}
*/

/*
type ImmatureElementStruct struct {
	FirstElementUuid   string
	ImmatureElementMap map[string]fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage
}

*/

const NotApplicable = "N/A"
