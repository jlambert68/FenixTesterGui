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

type matureElementStruct struct {
	firstElementUuid string
	matureElementMap map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage
}

type immatureElementStruct struct {
	firstElementUuid   string
	immatureElementMap map[string]fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage
}

// The Complex Deletion Rules
const (
	TCRuleDeletion101 = "TCRuleDeletion101"
	TCRuleDeletion102 = "TCRuleDeletion102"
	TCRuleDeletion103 = "TCRuleDeletion103"
	TCRuleDeletion104 = "TCRuleDeletion104"
	TCRuleDeletion105 = "TCRuleDeletion105"
	TCRuleDeletion106 = "TCRuleDeletion106"
	TCRuleDeletion107 = "TCRuleDeletion107"
	TCRuleDeletion108 = "TCRuleDeletion108"
	TCRuleDeletion109 = "TCRuleDeletion109"
	TCRuleDeletion110 = "TCRuleDeletion110"
	TCRuleDeletion111 = "TCRuleDeletion111"
	TCRuleDeletion112 = "TCRuleDeletion112"
	TCRuleDeletion113 = "TCRuleDeletion113"
	TCRuleDeletion114 = "TCRuleDeletion114"
	TCRuleDeletion115 = "TCRuleDeletion115"
	TCRuleDeletion116 = "TCRuleDeletion116"
	TCRuleDeletion117 = "TCRuleDeletion117"
)

// The Complex Swap Rules
const (
	TCRuleSwap101 = "TCRuleSwap101"
	TCRuleSwap102 = "TCRuleSwap102"
	TCRuleSwap103 = "TCRuleSwap103"
	TCRuleSwap104 = "TCRuleSwap104"
	TCRuleSwap105 = "TCRuleSwap105"
	TCRuleSwap106 = "TCRuleSwap106"
	TCRuleSwap107 = "TCRuleSwap107"
	TCRuleSwap108 = "TCRuleSwap108"
	TCRuleSwap109 = "TCRuleSwap109"
	TCRuleSwap110 = "TCRuleSwap110"
	TCRuleSwap111 = "TCRuleSwap111"
	TCRuleSwap112 = "TCRuleSwap112"
	TCRuleSwap113 = "TCRuleSwap113"
	TCRuleSwap114 = "TCRuleSwap114"
	TCRuleSwap115 = "TCRuleSwap115"
	TCRuleSwap116 = "TCRuleSwap116"
	TCRuleSwap117 = "TCRuleSwap117"
)
