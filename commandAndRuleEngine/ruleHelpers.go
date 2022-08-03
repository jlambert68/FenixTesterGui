package commandAndRuleEngine

import (
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

// SetLogger
// Set to use the same logger reference as is used by central part of system
func (commandAndRuleEngineObject *CommandAndRuleEngineObjectStruct) SetLogger(logger *logrus.Logger) {

	//myUIServer = UIServerStruct{}
	commandAndRuleEngineObject.logger = logger

}

// SetTestCasesReference
// Set Available Bonds
func (commandAndRuleEngineObject *CommandAndRuleEngineObjectStruct) SetAvailableBondsMap(availableBondsMap map[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage) {

	commandAndRuleEngineObject.availableBondsMap = availableBondsMap

}
