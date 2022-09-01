package commandAndRuleEngine

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/grpc_out"
	"FenixTesterGui/testCase/testCaseModel"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

type CommandAndRuleEngineObjectStruct struct {
	logger            *logrus.Logger
	availableBondsMap map[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage
	Testcases         *testCaseModel.TestCasesModelsStruct
	// subSystemsCrossReferences *gui.SubSystemsCrossReferencesStruct
	GrpcOutReference        *grpc_out.GRPCOutStruct
	CommandChannelReference *sharedCode.ChannelType
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
)
