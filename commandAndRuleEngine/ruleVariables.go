package commandAndRuleEngine

import (
	"FenixTesterGui/testCase/testCaseModel"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

type commandAndRuleEngineObjectStruct struct {
	logger            *logrus.Logger
	availableBondsMap map[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage
	testcaseModel     *testCaseModel.TestCaseModelStruct
}